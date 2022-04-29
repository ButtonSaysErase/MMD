package main

import (
	"flag"
	"fmt"
	"log"

	id3 "github.com/mikkyang/id3-go"
)

func main() {
	//var mp3 string
	mp3 := flag.String("mp3", "", "input mp3 file.")
	flag.Parse()

	fmt.Println("This is a test of the ID3 tag module.")
	file, err := id3.Open(*mp3)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fmt.Println("Artist: " + file.Artist())
	fmt.Println("Album: " + file.Album())
	fmt.Println("Year: " + file.Year())
	fmt.Println("Genre: " + file.Genre())
	fmt.Println("Title: " + file.Title())
}
