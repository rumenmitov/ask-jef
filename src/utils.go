package main

import (
    "fmt"
    "time"
    "os"
    "io"
    "log"
    "bufio"
    "strings"
)

func multiLineQuery(input *string, hasInput *bool) {
    fmt.Println(B_Purple + "What would you like to ask Jef?" + Reset);
    for {
        fmt.Printf("%s", Blue + "» " + I_Blue)
        var buf string;

        in := bufio.NewReader(os.Stdin)
        buf, err := in.ReadString('\n')
        if err != nil {
            fmt.Printf(Red)
            log.Printf("Could not get input: %s", err)
            fmt.Printf(Reset)
            return;
        }

        if buf == "\n" {
            break;
        }

        buf = strings.TrimSuffix(buf, "\n")

        if (buf != "") { 
            *hasInput = true;
            *input += " " + buf;
        }
    }
}

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
