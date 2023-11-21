package main

import (
    "fmt"
    "time"
)

func animationLoad() {

    for !isFinished {
        fmt.Printf("\r%s", "Jef is thinking   ")
        time.Sleep(1 * time.Second)
        fmt.Printf("\r%s", "Jef is thinking.  ")
        time.Sleep(1 * time.Second)
        fmt.Printf("\r%s", "Jef is thinking.. ")
        time.Sleep(1 * time.Second)
        fmt.Printf("\r%s", "Jef is thinking...")
        time.Sleep(1 * time.Second)
        fmt.Printf("\r%s", "Jef is thinking   ")
    }
}
