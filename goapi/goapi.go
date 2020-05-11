package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type User struct {
	ID        int    `json:"id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	CreatedAt string `json:"created_at"`
}

func main() {

	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/getusers", getusers).Methods("GET")

	log.Fatal(http.ListenAndServe(":3019", router))

}

func getusers(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Hi Aveek")

	sqlStatement := `select * from users;`
	var user User
	var response []User
	rows, err := queryDB(sqlStatement)
	fmt.Println(rows)
	for rows.Next() {
		err = rows.Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt)
		switch err {
		case sql.ErrNoRows:
			fmt.Println("No rows were returned!")
			return
		case nil:
			fmt.Println(user)
			response = append(response, user)
		default:
			panic(err)
			fmt.Println("Successfully Connected")
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func queryDB(sqlStatement string) (*sql.Rows, error) {
	var db *sql.DB
	db, err := sql.Open("postgres", "host=10.99.97.83 port=5432 user=postgresadmin password=admin123 dbname=postgresdb sslmode=disable")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	rows, err := db.Query(sqlStatement)
	return rows, err

}
