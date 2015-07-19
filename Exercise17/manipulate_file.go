/*
Exercise17
Manipulate a text file. A plain text file has the contents as shown below in “Original File”. Observe
that in this file, there exists a word ‘word’. Write a Go program that updates this file and the final
contents become as shown in “Modified Content” below.
*/
package main

import (
	"fmt"
	"os"
	"log"
	"io/ioutil"
	"strings"
)

func checkError(err error){
	if(err != nil){
		log.Fatal(err);
	}	
	
}
func main() {
	args := os.Args;
	if (len(args) !=4){
		log.Fatal("Pass Filename OldWord NewWord")
	}
	
	fn := args[1]
	w := args[2]
	nw := args[3]
	
	data, err := ioutil.ReadFile(fn)
	checkError(err)
	
	dataStr := string(data)
	dataStr = strings.Replace(dataStr, w, nw, -1)
	
	err = ioutil.WriteFile(fn, []byte(dataStr),  0666)
	checkError(err)
	
	fmt.Println("Successfully replaced \"" + w + "\" with \"" + nw + "\" in file \"" + fn + "\"")
	
}
