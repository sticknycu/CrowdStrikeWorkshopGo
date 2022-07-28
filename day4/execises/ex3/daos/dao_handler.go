package daos

import (
	"fmt"
	"http-rest/v/domain"
	"log"
)

var book_list = map[string]domain.Book{}

func AddBook(book domain.Book) error {
	val, ok := book_list[book.Title]
	if !ok {
		book_list[book.Title] = book
		log.Printf("Book %v is now in the database", book)
		return nil
	} else {
		errorMessage := "[ERROR] Book already exists in the database: " + val.Title
		log.Printf(errorMessage)
		return fmt.Errorf(errorMessage)
	}

}

func GetBook(title string) (domain.Book, error) {
	val, ok := book_list[title]
	if !ok {
		errorMessage := "[ERROR] Book does not exists in the database:" + title
		log.Println(errorMessage)
		return domain.Book{}, fmt.Errorf(errorMessage)
	} else {
		return val, nil
	}
}

func GetAll() map[string]domain.Book {
	return book_list
}

func EditTitleBook(title string, book domain.Book) (domain.Book, error) {
	val, ok := book_list[book.Title]
	if !ok {
		errorMessage := "[ERROR] Book does not exists in the database:" + title
		log.Println(errorMessage)
		return domain.Book{}, fmt.Errorf(errorMessage)
	} else {
		val.Title = title
		book_list[title] = val
		return val, nil
	}
}

func RemoveBook(title string) error {
	_, ok := book_list[title]
	if !ok {
		errorMessage := "[ERROR] Book does not exists in the database:" + title
		log.Println(errorMessage)
		return fmt.Errorf(errorMessage)
	} else {
		book_list[title] = domain.Book{}
		return nil
	}
}
