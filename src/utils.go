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

func multiLineQuery(input *string) {
    fmt.Println(B_Purple + "What would you like to ask Jef?" + Reset);
    for {
        fmt.Printf("%s", Blue + "» " + I_Blue)
        var buf string;

        in := bufio.NewReader(os.Stdin)
        buf, err := in.ReadString('\n')
        if err != nil {
            log.Println(err)
            return;
        }

        if buf == "\n" {
            break;
        }

        buf = strings.TrimSuffix(buf, "\n")

        if (buf != "") { 
            *input += " " + buf;
        }
    }
}

func updateModel(model string) {

    fp, err := os.Create(os.Getenv("HOME") + "/.config/ask-jef/ask.env")
    if err != nil {
        log.Println(err)
    }

    defer fp.Close()

    io.WriteString(fp, "MODEL=" + model)
}

func animationLoad() {

    for !isFinished {
        fmt.Printf("%s", B_Purple)
        fmt.Printf("\r%s", "Jef is thinking   ")
        time.Sleep(500 * time.Millisecond)
        fmt.Printf("\r%s", "Jef is thinking.  ")
        time.Sleep(500 * time.Millisecond)
        fmt.Printf("\r%s", "Jef is thinking.. ")
        time.Sleep(500 * time.Millisecond)
        fmt.Printf("\r%s", "Jef is thinking...")
        time.Sleep(500 * time.Millisecond)
        fmt.Printf("\r%s", "Jef is thinking   ")
        fmt.Printf("%s", Reset)
    }
}
