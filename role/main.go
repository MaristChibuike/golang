package main

import "fmt"

type content struct {
	title       string
	description string
	authors     []string
	year        int
}

// method to initialize content
func updateContents(c *content, title string, description string, authors []string, year int) {
	c.title = title
	c.description = description
	c.authors = authors
	c.year = year
}

func displayContents(c content) {
	fmt.Printf("Title: %s\nDescription: %s\nAuthors: %v\nYear: %d\n", c.title, c.description, c.authors, c.year)
}

func main() {
	// Initialize the content struct with information about a popular Igbo content
	var igboContent content = content{
		description: "Things Fall Apart is a novel written by Chinua Achebe. It is a classic in African literature.",
		authors:     []string{"Chinua Achebe"},
		title:       "Things Fall Apart",
		year:        1958,
	}

	// Display the details using the displayContents function
	displayContents(igboContent)
}
