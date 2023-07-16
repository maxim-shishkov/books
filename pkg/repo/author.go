package repo

type Author struct {
	Id   uint64 `json:"id"  db:"id"`
	Name string `json:"name"  db:"name"`
}
