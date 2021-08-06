package utils

import (
	"database/sql"
	"log"
	"regoftw/conf"

	_ "github.com/mattn/go-sqlite3"
)

type Manager interface {
	CheckFunctionStatus(domain string, function string) int
	UpdateStatus(domain string, function string, status int) bool
	GenerateDataDomain(domain string) bool
	CanRunFunction(domain string, function string) bool
	ExistDomain(domain string) bool
	// TODO: Add other methods
}

type manager struct {
	db *sql.DB
}

var mgr Manager = nil

func GetDBManager() Manager {
	conf.GetCTX().GetWorkPlace()
	if mgr == nil {
		mgr = newManager(conf.GetCTX().GetWorkPlace())
	}
	return mgr
}

func newManager(path string) Manager {
	db, err := sql.Open("sqlite3", path+"/status.db")
	if err != nil {
		log.Fatal("Failed to init db:", err)
	}
	statement, _ := db.Prepare(`CREATE TABLE IF NOT EXISTS domainfunctionsstatus
							(id INTEGER PRIMARY KEY, domain TEXT, func TEXT, status NUMBER)`)
	statement.Exec()
	return &manager{db: db}
}

func (mgr *manager) CheckFunctionStatus(domain string, function string) int {
	rows, err := mgr.db.Query("SELECT status FROM domainfunctionsstatus WHERE domain=? and func=?", domain, function)
	if err != nil {
		return -10
	}
	var status int
	rows.Next()
	rows.Scan(&status)
	rows.Close()
	return status
}

// Exist dif mode to allow run to find new data
func (mgr *manager) CanRunFunction(domain string, function string) bool {
	canRun := false
	if mgr.CheckFunctionStatus(domain, function) <= 0 {
		canRun = true
	}
	return canRun
}

func (mgr *manager) ExistDomain(domain string) bool {
	exist := false
	var total int
	mgr.db.QueryRow("SELECT COUNT(*) FROM domainfunctionsstatus WHERE domain=?", domain).Scan(&total)

	if total > 0 {
		exist = true
	}
	return exist
}

func (mgr *manager) UpdateStatus(domain string, function string, status int) bool {
	statement, _ := mgr.db.Prepare("UPDATE domainfunctionsstatus set status=? where domain=? and func=?")
	res, _ := statement.Exec(status, domain, function)
	affect, _ := res.RowsAffected()
	if affect == 1 {
		return true
	}
	return false
}

func (mgr *manager) GenerateDataDomain(domain string) bool {
	success := true
	functions := []string{"metafinder", "emailfinder", "gotator", "roboxtractor", "analyticsrelationships"} // TODO: Add all values from enumerator
	for _, f := range functions {
		statement, err := mgr.db.Prepare("INSERT INTO domainfunctionsstatus (domain, func, status) VALUES (?, ?, ?)")
		if err != nil {
			success = false
		}
		// TODO: Status in a struct
		statement.Exec(domain, f, 0) // 0 = the function has not started, -1 error, 1 function running, 2 function finished
	}
	return success
}
