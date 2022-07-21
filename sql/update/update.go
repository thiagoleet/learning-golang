package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:@/cursogo")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Update
	stmtUpdate, _ := db.Prepare("update usuarios set nome = ? where id = ?")
	stmtUpdate.Exec("Uóxiton Stive", 1)
	stmtUpdate.Exec("Sheristone Ualeska", 2)

	// Delete
	stmtDelete, _ := db.Prepare("delete from usuarios where id = ?")
	stmtDelete.Exec(3)

}
