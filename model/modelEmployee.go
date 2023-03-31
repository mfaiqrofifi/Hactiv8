package model

type BookDomain struct {
	ID     int    `json:"id" db:"id"`
	Title  string `json:"title" db:"title"`
	Author string `json:"author" db:"author"`
	Desc   string `json:"desc" db:"description"`
	// Division  string `json:"division" db:"division"`
}
