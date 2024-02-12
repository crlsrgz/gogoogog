//--Summary:
//  Create a program to manage lending of library books.
//
//--Requirements:
//* The library must have books and members, and must include:
//  - Which books have been checked out
//  - What time the books were checked out
//  - What time the books were returned
//* Perform the following:
//  - Add at least 4 books and at least 3 members to the library
//  - Check out a book
//  - Check in a book
//  - Print out initial library information, and after each change
//* There must only ever be one copy of the library in memory at any time
//
//--Notes:
//* Use the `time` package from the standard library for check in/out times
//* Liberal use of type aliases, structs, and maps will help organize this project

package main

import (
	"fmt"
	"time"
)

type Title string
type Name string

type LendAudit struct {
	checkOut time.Time
	checkIn  time.Time
}

type Member struct {
	name  Name
	books map[Title]LendAudit
}
type BookEntry struct {
	total  int
	lended int
}
type Library struct {
	members map[Name]Member
	books   map[Title]BookEntry
}

func printMemberAudit(member *Member) {
	for title, audit := range member.books {
		var returnTime string
		if audit.checkIn.IsZero() {
			returnTime = "[not returned yet]"

		} else {
			returnTime = audit.checkOut.String()
		}
		fmt.Println(member.name, ":", title, ":", audit.checkOut.String(), "through", returnTime)

	}
}
func printMemberAudits(library *Library) {
	for _, member := range library.members {
		printMemberAudit(&member)
	}
}
func printLibraryBooks(library *Library) {
	fmt.Println()
	for title, book := range library.books {
		fmt.Println(title, "/total:", book.total, "/ lended:", book.lended)
	}

	fmt.Println()
}

func checkoutBook(library *Library, title Title, member *Member) bool {
	book, found := library.books[title]
	if !found {
		fmt.Println("Book is in the Librar catalog")
		return false

	}
	if book.lended == book.total {
		fmt.Println("no more books available")
		return false

	}
	book.lended += 1
	library.books[title] = book

	member.books[title] = LendAudit{checkOut: time.Now()}
	return true
}
func returnBook(library *Library, title Title, member *Member) bool {
	book, found := library.books[title]
	if !found {
		fmt.Println("Book not part of the library")
		return false
	}
	audit, found := member.books[title]
	if !found {
		fmt.Println("Member did not ckeck out this book")
		return false
	}
	book.lended -= 1
	library.books[title] = book

	audit.checkIn = time.Now()
	member.books[title] = audit
	return true
}

func main() {
	library := Library{
		books:   make(map[Title]BookEntry),
		members: make(map[Name]Member),
	}

	library.books["golang"] = BookEntry{
		total: 5,
	}
	library.books["java"] = BookEntry{
		total:  5,
		lended: 0,
	}
	library.books["cpp"] = BookEntry{
		total:  5,
		lended: 0,
	}
	library.books["c"] = BookEntry{
		total:  5,
		lended: 0,
	}
	library.members["John"] = Member{"John", make(map[Title]LendAudit)}
	library.members["Bill"] = Member{"John", make(map[Title]LendAudit)}
	library.members["Peter"] = Member{"John", make(map[Title]LendAudit)}

	fmt.Println("\nInitial:")
	printLibraryBooks(&library)
	printMemberAudits(&library)

	member := library.members["John"]
	checkedOut := checkoutBook(&library, "golang", &member)
	fmt.Println("Check out a book")
	if checkedOut {
		printLibraryBooks(&library)
		printMemberAudits(&library)
	}
	returned := returnBook(&library, "golang", &member)
	fmt.Println("\n Check in a book")
	if returned {
		printLibraryBooks(&library)
		printMemberAudits(&library)

	}

	member = library.members["John"]
	checkedOut = checkoutBook(&library, "cpp", &member)
	fmt.Println("Check out a book")
	if checkedOut {
		printLibraryBooks(&library)
		printMemberAudits(&library)
	}
}
