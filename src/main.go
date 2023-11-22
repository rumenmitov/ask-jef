package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
    "log"
	"io"
	"net/http"
	"os"
	"strings"
)

// tracks if AI has finished with its response
var isFinished bool = false;

func main() {

    // get user input
    var input string = "";
    var model string = "luna-ai-llama2"

    if (len(os.Args) == 1) {
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
            input += " " + buf;
        }
    } else {
        for index, value := range os.Args {
            if value != "-m" {
                input += " " + value;
            } else {
                model = os.Args[index+1]
                index++;
            }
        }
    }

    // generate AI response
    fmt.Printf(Clear)
    fmt.Printf("\n%s", Reset)

	url := "http://localhost:8080/v1/chat/completions"

    payload := Payload {
        Model: model,
        Messages: []Message{
            {
                Role: "user",
                Content: input,
            },
        },
        Temperature: 0.9,
    };

    payload_str, err := json.Marshal(payload)
    if (err != nil) {
        fmt.Printf(Red)
        log.Printf("Error parsing user input to JSON: %s", err)
        fmt.Printf(Reset)
    }

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload_str))
	if err != nil {
        fmt.Printf(Red)
        log.Printf("Error fetching request: %s", err)
        fmt.Printf(Reset)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

    // loading animation in the background
    go animationLoad();

    resp, err := client.Do(req)
    if err != nil {
        isFinished = true
        fmt.Printf(Red)
        log.Printf("Error sending request: %s", err)
        fmt.Printf(Reset)
    }
    defer resp.Body.Close()
    isFinished = true

    fmt.Printf("\r                  ");

	body, err := io.ReadAll(resp.Body)
	if err != nil {
        fmt.Printf(Red)
        log.Printf("Error reading response body: %s", err)
        fmt.Printf(Reset)
		return
	}

	// print result
    var v Data;
    json.Unmarshal([]byte(body), &v)

    fmt.Printf("%s", B_Cyan);
    fmt.Printf("\r%s", v.Choices[0].Message.Content + Reset)
}
