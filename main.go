package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", "user=postgres password=postgres dbname=mvxt sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	s := &Store{DB: db}

	http.Handle("/movies", &AllMovies{
		Store: s,
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

type AllMovies struct {
	Store *Store
}

func (a *AllMovies) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	movies := a.Store.All()
	rw.Header().Add("content-type", "application/json; charset=utf-8")

	json.NewEncoder(rw).Encode(movies)
}

type Store struct {
	DB *sql.DB
}

func (s *Store) All() []Movie {
	rows, err := s.DB.Query("SELECT * FROM movies")
	if err != nil {
		log.Fatal(err)
	}

	var ms []Movie
	for rows.Next() {
		var (
			name     string
			director string
		)
		err := rows.Scan(&name, &director)
		if err != nil {
			log.Fatal(err)
		}
		ms = append(ms, Movie{Name: name, Director: director})
	}

	return ms
}

type Movie struct {
	Name     string `json:"name"`
	Director string `json:"director"`
}

func (m Movie) String() string {
	return fmt.Sprintf("{ Name = %s, Director = %s", m.Name, m.Director)
}
