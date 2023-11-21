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
        for _, value := range os.Args {
            input += " " + value;
        }
    }

    // generate AI response
    fmt.Printf(Clear)
    fmt.Printf("\n%s", Reset)

	url := "http://localhost:8080/v1/chat/completions"
	payload := []byte(`{ "model": "luna-ai-llama2", "messages": [{"role": "user", "content": "` + input + `"}], "temperature": 0.9 }`)

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload))
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
