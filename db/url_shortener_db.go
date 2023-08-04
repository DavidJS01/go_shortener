package url_db

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)

func connectToDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./db/url_shortener.db")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func GetUrlMapping(url string, getLongURL bool) (map[string]string, error) {
	db, err := connectToDB()
	var getURLQuery string
	var mappedURL string

	if err != nil {
		return nil, errors.New("Error while establishing a connection to the sqlite database")
	}
	// is the below toggle a code smell?
	if getLongURL {
		getURLQuery = fmt.Sprintf(`SELECT long_url FROM url_shortener WHERE shortened_url = '%s'`, url)
	} else {
		getURLQuery = fmt.Sprintf(`SELECT shortened_url FROM url_shortener WHERE long_url = '%s'`, url)
	}
	row := db.QueryRow(getURLQuery)
	err = row.Scan(&mappedURL)
	if err != nil {
		return nil, errors.New("Error while assigning url")
	}

	urlMapping := make(map[string]string)
	urlMapping[url] = mappedURL
	return urlMapping, nil
}
