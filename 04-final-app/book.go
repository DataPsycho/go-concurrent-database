package main

import "fmt"

type Book struct {
	ID            int
	Title         string
	Author        string
	YearPublished int
}

func (b Book) String() string {
	return fmt.Sprintf(
		"Title\t\t%q\n"+
			"Author\t\t%q\n"+
			"Author\t\t%v\n", b.Title, b.Author, b.YearPublished,
	)
}

var books = []Book{
	{1, "A Simple Story", "Kellyer", 1990},
	{2, "A Double Plan", "Kumin", 1992},
	{3, "Rock Mozart", "Rob", 1989},
	{4, "Rolling Stone", "Dylan", 1992},
	{5, "Kora Jobabl", "Succker", 1880},
	{6, "Nomad", "Nunoc", 1892},
	{7, "Cumin Zinzer", "Koloy", 1983},
	{8, "Under Armour", "Roxy", 1905},
	{9, "Shock Attack", "Killer", 1900},
	{10, "Bug In your Hair", "Roland", 1900},
}
