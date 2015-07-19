/*
Exercise 18
Analyze an MP3 File. Write a Go program that analyzes an MP3 file. Many MP3 files have a 128-byte
data structure at the end called an ID3 tag. These 128 bytes are literally packed with information
about the song: its name, the artist, which album it’s from, and so on. You can parse this data
structure by opening an MP3 file and doing a series of reads from a position near the end of the
file. According to the ID3 standard, if you start from the 128th-to-last byte of an MP3 file and read
three bytes, you should get the string TAG. If you don’t, there’s no ID3 tag for this MP3 file, and
nothing to do. If there is an ID3 tag present, then the 30 bytes after TAG contain the name of the
song, the 30 bytes after that contain the name of the artist, and so on. A sample song² is available
to test your program.
*/
package main

import (
	"fmt"
	"os"
	"log"	
	"strings"
)

func checkError(err error){
	if(err != nil){
		log.Fatal(err);
	}	
	
}

func main() {
	args := os.Args;
	if (len(args) !=2){
		log.Fatal("Pass MP3 Song Filename");
	}
	
	fn := args[1]
	file, err := os.Open(fn)
	checkError(err)	
	defer file.Close()
	
	stat, err := file.Stat()
	checkError(err)
	
	_, err = file.Seek(stat.Size()-128, 0)
	tag := make([]byte, 3)
	_, err = file.Read(tag)
	checkError(err)
	
	printDetails := func (s string, size int){
		details := make([]byte , size)
		_, err = file.Read(details)
		fmt.Println(s + ": " + string(details))
	}
	if(strings.Contains(string(tag), "TAG")){
		fmt.Println("Filename " + fn);				
		printDetails("Title", 30)
		printDetails("Artist", 30)
		printDetails("Album", 30)
		printDetails("Year", 4)		
	}else{
		fmt.Println("No information is available for " + fn)
	}
	
}
