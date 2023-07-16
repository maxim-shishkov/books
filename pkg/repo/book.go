package repo

type Book struct {
	Id       uint64 `json:"id" db:"id"`
	Name     string `json:"name" db:"name"`
	AuthorId uint64 `json:"author_id" db:"author_id"`
}
