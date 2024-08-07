package greetings

import (
    "errors"
    "fmt"
    "math/rand"
    "time"
)

func init() {
    rand.Seed(time.Now().UnixNano())
}


func Hello(name string) (string, error) {
    
    if name == "" {
        return name, errors.New("empty name")
    }
    
    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}

func randomFormat() string {
    
    formats := []string{
        "Hi, %v. Welcome!",
        "Great to see you, %v!",
        "Hail, %v! Well met!",
    }
    return formats[rand.Intn(len(formats))]
}