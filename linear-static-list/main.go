package main

import (
	"fmt"
	"os"
	"os/exec"
)

type Book struct {
	Id          uint32
	Description string
	Author      string
	Year        uint8
}

type List struct {
	Books [10]Book
	Total int
}

func initialize(l *List) {
	l.Total = 0
}

func getTotalOfBooks(l *List) (t int) {
	t = l.Total
	return
}

func getBooks(l *List) {
	if l.Total == 0 {
		fmt.Println("It's empty...")
		return
	}

	fmt.Println("#############################")
	for i := 0; i < l.Total; i++ {
		fmt.Printf("Id: %d", l.Books[i].Id)
		fmt.Printf("Description: %s", l.Books[i].Description)
		fmt.Printf("Author: %s", l.Books[i].Author)
		fmt.Printf("Year: %d", l.Books[i].Year)
		fmt.Println("#############################")
	}
}

func addBookAtTheEnd(l *List, b Book) {
	if l.Total == len(l.Books) {
		fmt.Println("It's full...")
		return
	}
}

func main() {
	var op int
	for true {
		menu()

		fmt.Scanf("%d", &op)
		if op == 0 {
			return
		}

		var l List
		initialize(&l)

		switch op {
		case 1:
			cleanBash()
			t := getTotalOfBooks(&l)
			fmt.Println("Total: ", t)
			fmt.Print("Press any key to continue...")
			fmt.Scanf("%s")
			cleanBash()
			break
		case 2:
			cleanBash()
			getBooks(&l)
			fmt.Print("Press any key to continue...")
			fmt.Scanf("%s")
			cleanBash()
			break
		case 3:
			cleanBash()
			getBooks(&l)
			fmt.Print("Press any key to continue...")
			fmt.Scanf("%s")
			cleanBash()
			break
		default:
			cleanBash()
			fmt.Println("Invalid option...")
			fmt.Print("Press any key to continue...")
			fmt.Scanf("%s")
			cleanBash()
			break
		}
	}
}

func menu() {
	fmt.Println("########### Hello ###########")
	fmt.Println("# (1) Get total of book     #")
	fmt.Println("# (2) Get books information #")
	fmt.Println("# (3) Exit                  #")
	fmt.Println("#############################")
	fmt.Print("Take an option: ")
}

func cleanBash() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
