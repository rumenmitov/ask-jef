package main

import (
	"bytes"
	"encoding/json"
	"fmt"
    "log"
	"io"
	"net/http"
	"os"
    "github.com/joho/godotenv"
)

// tracks if user has entered input
var hasInput bool = false;

// tracks if a file has been entered
var hasFile bool = false;

// tracks if AI has finished with its response
var isFinished bool = false;

func main() {

    // get environment variables
    err := godotenv.Load(os.Getenv("HOME") + "/.config/ask-jef/ask.env")
    if err != nil {
        fmt.Printf(Red)
        log.Printf("Could not load environment: %s", err)
        fmt.Printf(Reset)
        return;
    }


    // get user input
    var model string = "";
    var input string = "";

    if (len(os.Args) == 1) {
        multiLineQuery(&input, &hasInput)
    } else {
        for i := 1; i < len(os.Args); i++ {
            value := os.Args[i]

            if value == "-m" {
                model = os.Args[i+1]
                i+=1;
                continue;
            } else if value == "-f" {
                file_content, err := os.ReadFile(os.Args[i+1])
                if (err != nil) {
                    fmt.Printf(Red)
                    log.Printf("Error reading file: %s", err)
                    fmt.Printf(Reset)
                }
                input += "\n" + string(file_content) + "\n"
                i++;
                hasFile = true;
                continue
            } else {
                input += value + " ";
                hasInput = true;
            }
        }
        fmt.Printf("\n")
    }

    if model == "" {
        model = os.Getenv("MODEL")
    }

    if model == "" {
        fmt.Printf("%s", B_Purple + "Enter model name (should be the same from LocalAI/models/ direcotry): " + Reset)
        fmt.Scanf("%s", &model)
    }

    updateModel(model);

    if input == "" && !hasFile {
        return
    }

    if hasFile && !hasInput {
        multiLineQuery(&input, &hasInput)
        if (!hasInput) { return }
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
