package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	// connect to database
	connection, err := sql.Open("pgx", "host=localhost port=5432 dbname=test_connect user=postgres password=admin")
	if err != nil {
		log.Fatal(fmt.Sprintf("Unable to connect: %v\n", err))
	}
	defer connection.Close()

	log.Println("Connected to db")
	//test my connection
	err = connection.Ping()
	if err != nil {
		log.Fatal("Cannot pint db")
	}

	log.Println("Pinged db")
	//get rows from table
	err = getAllRows(connection)
	if err != nil {
		log.Fatal(err)
	}
	//insert a row
	query := `insert into users (first_name, last_name) values ($1, $2)`
	_, err = connection.Exec(query, "Jack", "Brown")

	if err != nil {
		log.Fatal(err)
	}
	//get rows from table again
	err = getAllRows(connection)
	if err != nil {
		log.Fatal(err)
	}
	//update a row
	stmt := `update users set first_name = $1 where first_name = $2`
	_, err = connection.Exec(stmt, "Jackie", "Jack")
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Updated to jackie")
	//get a row
	err = getAllRows(connection)
	if err != nil {
		log.Fatal(err)
	}
	// get one row by id
	var firstName, lastName string
	var id int
	fmt.Println("Query by id:")
	query = `select id, first_name, last_name from users where id = $1`
	row := connection.QueryRow(query, 1)
	err = row.Scan(&id, &firstName, &lastName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Record is", id, firstName, lastName)
	// delete a row

	query = `delete from users where id = $1`
	_, err = connection.Exec(query, 6)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Deleted a row")
	//get rows again

	err = getAllRows(connection)
	if err != nil {
		log.Fatal(err)
	}
}

func getAllRows(conn *sql.DB) error {
	rows, err := conn.Query("select id, first_name, last_name from users")
	if err != nil {
		log.Println(err)
		return err
	}
	defer rows.Close()

	var firstName, lastName string
	var id int

	for rows.Next() {
		err := rows.Scan(&id, &firstName, &lastName)
		if err != nil {
			log.Println(err)
			return err
		}
		fmt.Println("Record is", id, firstName, lastName)
	}

	if err = rows.Err(); err != nil {
		log.Fatal("Error scanning rows", err)
	}

	fmt.Println("----------------------------")

	return nil
}
