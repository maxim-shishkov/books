package api

import (
	"context"
	"github.com/maxim-shishkov/books/pkg/database"
)

type GRPCServer struct{}

func (s *GRPCServer) GetBooks(ctx context.Context, req *BookRequest) (*BookResponse, error) {
	b, err := database.GetBooks(req.Author)

	books := make([]string, len(b))
	for i := 0; i < len(b); i++ {
		books[i] = b[i].Name
	}

	return &BookResponse{Books: books}, err
}

func (s *GRPCServer) GetAuthor(ctx context.Context, req *AuthorRequest) (*AuthorResponse, error) {
	author, err := database.GetAuthor(req.Book)
	return &AuthorResponse{Author: author}, err
}
