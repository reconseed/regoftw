package utils

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Manager interface {
	CheckFunctionStatus(domain string, function string) int
	UpdateStatus(domain string, function string, status int) bool
	GenerateDataDomain(domain string) bool
	// TODO: Add other methods
}

type manager struct {
	db *sql.DB
}

var mgr Manager = nil

func GetManager(path string) Manager {
	if mgr == nil {
		mgr = newManager(path)
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
	functions := []string{"metafinder", "emailfinder", "gotator"} // TODO: Add all values from file
	for _, f := range functions {
		statement, err := mgr.db.Prepare("INSERT INTO domainfunctionsstatus (domain, func, status) VALUES (?, ?, ?)")
		if err != nil {
			success = false
		}
		statement.Exec(domain, f, 0) // 0 = the function has not started, -1 error, 1 function finished
	}
	return success
}
