package main

import "fmt"

type Book struct {
	id     int64
	title  string
	author string
	year   uint16
	size   uint32
	rate   float32
}

func (book *Book) ID() int64 {
	return book.id
}

func (book *Book) SetID(id int64) {
	book.id = id
}

func (book *Book) Title() string {
	return book.title
}

func (book *Book) SetTitle(title string) {
	book.title = title
}

func (book *Book) Author() string {
	return book.author
}

func (book *Book) SetAuthor(author string) {
	book.author = author
}

func (book *Book) Year() uint16 {
	return book.year
}

func (book *Book) SetYear(year uint16) {
	book.year = year
}

func (book *Book) Size() uint32 {
	return book.size
}

func (book *Book) SetSize(size uint32) {
	book.size = size
}

func (book *Book) Rate() float32 {
	return book.rate
}

func (book *Book) SetRate(rate float32) {
	book.rate = rate
}

type comparingModeBooks int

const (
	CompareYear = iota
	CompareSize
	CompareRate
)

type ComparerBooks struct {
	mode comparingModeBooks
}

func NewComparerBooks(mode comparingModeBooks) *ComparerBooks {
	return &ComparerBooks{mode: mode}
}

func (comparerBooks *ComparerBooks) Compare(firstBook, secondBook *Book) bool {
	switch comparerBooks.mode {
	case CompareYear:
		return firstBook.year > secondBook.year
	case CompareSize:
		return firstBook.size > secondBook.size
	case CompareRate:
		return firstBook.rate > secondBook.rate
	}
	return false
}

func comparator(mode comparingModeBooks, firstBook, secondBook *Book) string {
	compareMode := NewComparerBooks(mode)
	result := compareMode.Compare(firstBook, secondBook)

	if result {
		return "Получена первая книга"
	}
	return "Получена вторая книга"
}

func main() {
	firstBook := Book{
		id:     1,
		title:  "First Book",
		author: "Блох Джошуа",
		year:   2021,
		size:   1500,
		rate:   0.95,
	}

	secondBook := Book{
		id:     2,
		title:  "Second Book",
		author: "Другой автор",
		year:   2022,
		size:   500,
		rate:   0.78,
	}

	secondResult := comparator(CompareYear, &firstBook, &secondBook)
	fmt.Println(secondResult)

	firstResult := comparator(CompareSize, &firstBook, &secondBook)
	fmt.Println(firstResult)

	threeResult := comparator(CompareRate, &firstBook, &secondBook)
	fmt.Println(threeResult)
}
