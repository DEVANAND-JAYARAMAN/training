package main

import "fmt"

// 1. Interface
type Device interface {
    TurnOn() string
}

// 2. TV implements Device
type TV struct{}
func (t TV) TurnOn() string {
    return "TV is turned on"
}

// 3. Speaker implements Device
type Speaker struct{}
func (s Speaker) TurnOn() string {
    return "Speaker is turned on"
}

// 4. Remote function
func Remote(d Device) {
    fmt.Println(d.TurnOn())
}

func main() {
    tv := TV{}
    speaker := Speaker{}

    Remote(tv)      // Output: TV is turned on
    Remote(speaker) // Output: Speaker is turned on
}
