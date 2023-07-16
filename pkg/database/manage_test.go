package database_test

import (
	"github.com/DATA-DOG/go-sqlmock"
	_ "github.com/lib/pq"
	"github.com/maxim-shishkov/books/pkg/database"
	"github.com/maxim-shishkov/books/pkg/repo"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_GetAuthor(t *testing.T) {
	fakeDB, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer fakeDB.Close()
	database.DB = fakeDB

	tests := []struct {
		request     string
		expected    string
		expectedErr bool
	}{
		{
			request:     "Капитанская дочка",
			expected:    "Пушкин",
			expectedErr: false,
		},
		{
			request:     "",
			expected:    "",
			expectedErr: true,
		},
	}

	for _, tt := range tests {
		rows := sqlmock.NewRows([]string{"name"}).AddRow(tt.expected)

		query := `SELECT a.name FROM author a JOIN book b ON b.author_id \= a.id WHERE b.title \= \$1`
		mock.ExpectQuery(query).WithArgs(tt.request).WillReturnRows(rows)

		author, err := database.GetAuthor(tt.request)

		assert.Equal(t, tt.expectedErr, err != nil)
		assert.Equal(t, tt.expected, author)
	}
}

func Test_GetBooks(t *testing.T) {
	fakeDB, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer fakeDB.Close()
	database.DB = fakeDB

	tests := []struct {
		request     string
		expected    []repo.Book
		expectedErr bool
	}{
		{
			request: "Пушкин",
			expected: []repo.Book{
				{
					Name: "Капитанская дочка",
				},
				{
					Name: "Пиковая дама",
				},
			},
			expectedErr: false,
		},
		{
			request:     "no author",
			expected:    nil,
			expectedErr: false,
		},
		{
			request:     "",
			expected:    nil,
			expectedErr: true,
		},
	}

	for _, tt := range tests {
		rows := sqlmock.NewRows([]string{"name"})

		for _, ex := range tt.expected {
			rows.AddRow(ex.Name)
		}

		mock.ExpectQuery(`ELECT b.title FROM book b JOIN author a ON a.id \= b.author_id WHERE a.name \= \$1`).
			WithArgs(tt.request).WillReturnRows(rows)

		books, err := database.GetBooks(tt.request)

		assert.Equal(t, tt.expectedErr, err != nil)
		if !tt.expectedErr {
			assert.Equal(t, tt.expected, books)
		}
	}
}
