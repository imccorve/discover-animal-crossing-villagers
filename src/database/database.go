package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" // Loads driver anonymously
)

func addQuotes(str string) string {
	return "'" + str + "'"
}

func convertVillagersToDB(tableName string, db *sql.DB) {
	villagers, err := getVillagersFromNookipedia()
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, villager := range villagers {
		values := addQuotes(villager.ImageURL) + ", " + addQuotes(villager.Name) + ", " + addQuotes(villager.JapaneseName) +
			", " + addQuotes(villager.Species) + ", " + addQuotes(villager.Gender) + ", " + addQuotes(villager.Personality)
		insert(values, tableName, db)
	}
}

func insert(values string, tableName string, db *sql.DB) {
	_, err := db.Exec("INSERT INTO " + tableName + " VALUES(" + values + ");")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Adding new row")
	}
}

// func initDatabase(databaseName string, db *sql.DB) {
// 	_, err = db.Exec("CREATE DATABASE " + databaseName)
// 	if err != nil {
// 		fmt.Println(err.Error())
// 	} else {
// 		fmt.Println("Successfully created database..")
// 	}
// }

func main() {
	tableName := "VILLAGER"
	db, err := sql.Open("mysql", "root:___@tcp(127.0.0.1:3306)/")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Database created successfully")
	}

	// initDatabase(tableName, db)
	databaseName := "ANIMAL_CROSSING"
	_, err = db.Exec("CREATE DATABASE " + databaseName)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Successfully created database..")
	}

	_, err = db.Exec("USE " + databaseName)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Now using " + databaseName)
	}

	_, queryErr := db.Exec("CREATE TABLE IF NOT EXISTS VILLAGER(" +
		"imageURL varchar(90), " +
		"name varchar(50), " +
		"japanese_name varchar(50), " +
		"species varchar(20), " +
		"gender varchar(10), " +
		"personality varchar(10), " +
		"PRIMARY KEY (name))" +
		";")
	if queryErr != nil {
		fmt.Println(queryErr.Error())
	}

	convertVillagersToDB(tableName, db)
	defer db.Close()
}

// http://go-database-sql.org/accessing.html
