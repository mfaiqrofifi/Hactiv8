package query

const (
	AddBook = `
	INSERT INTO books
	(
		title,
		author,
		description
	)
	VALUES($1,$2,$3)
	RETURNING *;
	`
)

const (
	GetBook = `
	SELECT * FROM books
	`
)

const (
	GetBookId = `
	SELECT * FROM books where id=$1
	`
)

const (
	UpdateBook = `
	UPDATE books
	SET title=$2,author=$3,description=$4
	WHERE id=$1
	`
)
const (
	DeleteBook = `
	DELETE FROM books
	WHERE id=$1
	`
)
