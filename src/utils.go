package main

import (
    "fmt"
    "time"
    "os"
    "io"
    "log"
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

func updateModel(model string) {

    fp, err := os.Create(os.Getenv("HOME") + "/.config/ask-jef/ask.env")
    if err != nil {
        fmt.Printf(Red)
        log.Printf("Could not open .env file: %s", err)
        fmt.Printf(Reset)
    }

    defer fp.Close()

    io.WriteString(fp, "MODEL=" + model)
}
