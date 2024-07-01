package main

import (
    "fmt"
    "log"
    "strings"
    "example.com/greetings"
)

func main() {

    // Request a greeting message.
    message, err := greetings.Hello("Riccardo")
    
    if (len(message) < 10)
        strings.Join(message, )

    // If an error was returned, print it to the console and
    // exit the program.
    if err != nil {
        log.Fatal(err)
    }

    // If no error was returned, print the returned message
    // to the console.
    fmt.Println(message)
}