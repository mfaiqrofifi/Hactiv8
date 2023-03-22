package domain

type BookDomain struct {
	BookId string `json:"book_id"`
	Title  string `json:"title"`
	Author string `json:"author"`
	Desc   string `json:"desc"`
}

var Books []BookDomain
