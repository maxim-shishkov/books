package database

import (
	"errors"
	"fmt"
	"github.com/maxim-shishkov/books/pkg/repo"
)

func GetAuthor(book string) (name string, err error) {
	if book == "" {
		return "", errors.New("no book")
	}
	row := DB.QueryRow("SELECT a.name FROM author a JOIN book b ON b.author_id = a.id WHERE b.title =  $1", book)

	author := repo.Author{}
	err = row.Scan(&author.Name)
	if err != nil {
		return
	}
	name = author.Name
	return
}

func GetBooks(author string) (books []repo.Book, err error) {
	if author == "" {
		return nil, errors.New("no author")
	}
	rows, err := DB.Query("SELECT b.title FROM book b JOIN author a ON a.id = b.author_id WHERE a.name = $1", author)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var book repo.Book
		scanErr := rows.Scan(&book.Name)
		if scanErr != nil {
			fmt.Println(scanErr)
			continue
		}
		books = append(books, book)
	}
	return
}
