package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
    "errors"
	"github.com/google/uuid"
	"github.com/joho/godotenv"
)

// tracks if AI has finished with its response
var isFinished bool = false;

func main() {

    // get environment variables
    err := godotenv.Load(os.Getenv("HOME") + "/.config/ask-jef/ask.env")
    if err != nil {
        log.Println(err)
        return;
    }


    // get user input
    session := Session {
        AlreadyExists: false,
        Id: uuid.NewString(),
        Content: "",
    }
    var model string = "";
    var input string = "";
    var filesArr []UserFile;

    if (len(os.Args) == 1) {
        // case for no flags
        multiLineQuery(&input)
    } else {
        for i := 1; i < len(os.Args); i++ {
            value := os.Args[i]

            if value == "-m" {
                // flag for setting the model
                model = os.Args[i+1]
                i++;

            } else if value == "-f" {
                // flag for adding files
                f := UserFile { 
                    Name: os.Args[i+1], 
                    Contents: "",
                };

                contents, err := os.ReadFile(f.Name);
                if (err != nil) {
                    log.Println(err)
                }

                f.Contents = string(contents);
                filesArr = append(filesArr, f);
                i++;

            } else if value == "-s" {
                // flag for setting the session name
                session.Id = os.Args[i+1];

                session_file := os.Getenv("HOME") + "/.cache/ask-jef/" + session.Id;

                _, err := os.Stat(session_file);
                if err == nil {
                    session.AlreadyExists = true;

                    contents, err := os.ReadFile(session_file);
                    if err != nil {
                        log.Println(err)
                    }

                    session.Content = string(contents);
                }

                i++;

            } else if value == "-cat" {
                // flag for displaying the contents of a session
                session_file := os.Getenv("HOME") + "/.cache/ask-jef/" + os.Args[i+1];

                _, err := os.Stat(session_file);
                if err == nil {
                    contents, err := os.ReadFile(session_file);
                    if err != nil {
                        log.Println(err)
                    }

                    fmt.Printf("\n%s\n", B_Cyan + string(contents) + Reset)
                }

                return;

            } else if value == "-ls" {
                // flag for listing sessions
                session_dir := os.Getenv("HOME") + "/.cache/ask-jef/";

                dir, err := os.Open(session_dir)
                if err != nil {
                    log.Println(err)
                }

                files, err := dir.ReadDir(0)
                if err != nil {
                    log.Println(err)
                }

                fmt.Printf("\n%s", B_Purple + "Sessions:\n\n" + I_Cyan)

                for _, file := range files {
                    if file.IsDir() { continue; }

                    fmt.Printf("%s\n", file.Name())
                }

                fmt.Printf("%s", Reset)

                return;

            } else if value == "-rm" {
                // flag for deleting a session 
                session_file := os.Getenv("HOME") + "/.cache/ask-jef/" + os.Args[i+1];

                _, err := os.Stat(session_file);
                if err == nil {
                    err := os.Remove(session_file);
                    if err != nil {
                        log.Println(err)
                    }
                }

                return;

            } else if value == "-mv" {
                // flag for renaming a session 
                session_file := os.Getenv("HOME") + "/.cache/ask-jef/" + os.Args[i+1];
                new_name := os.Getenv("HOME") + "/.cache/ask-jef/" + os.Args[i+2]

                if new_name == "" {
                    log.Println(errors.New("Please provide a new name for the session!\n"))
                    return;
                }

                _, err := os.Stat(session_file);
                if err == nil {
                    err := os.Rename(session_file, new_name);
                    if err != nil {
                        log.Println(err)
                    }
                }

                return;

            } else {
                // get input from argument
                input += value;
            }
        }
        fmt.Printf("\n")
    }

    // parse user input

    if model == "" {
        model = os.Getenv("MODEL")
    }

    if model == "" {
        fmt.Printf("%s", B_Purple + "Enter model name (should be the same from LocalAI/models/ direcotry): " + Reset)
        fmt.Scanf("%s", &model)
    }

    updateModel(model);

    if (len(filesArr) == 0) && (input == "") {
        multiLineQuery(&input)
        if input == "" { return }
    }

    if session.Content != "" {
        input += "\nCosider the context:\n`" + session.Content + "`\nNow answer the following question:\n";
    }

    for _, file := range filesArr {
        input += "\nContents of file: `" + file.Name + "`\n";
        input += "`" + file.Contents + "`";
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
        log.Println(err)
    }

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(payload_str))
	if err != nil {
        log.Println(err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

    // loading animation in the background
    go animationLoad();

    resp, err := client.Do(req)
    if err != nil {
        isFinished = true
        log.Println(err)
    }
    defer resp.Body.Close()
    isFinished = true

    fmt.Printf("\r                  ");

	body, err := io.ReadAll(resp.Body)
	if err != nil {
        log.Println(err)
		return
	}

	// print result to stdout
    var v Data;
    json.Unmarshal([]byte(body), &v)

    fmt.Printf("%s", B_Cyan);
    fmt.Printf("\r%s", v.Choices[0].Message.Content + Reset)

    // save result to session file
    session_file := os.Getenv("HOME") + "/.cache/ask-jef/" + session.Id;

    f_out, err := 
        os.OpenFile(session_file, os.O_APPEND | os.O_CREATE | os.O_WRONLY, 0644);
    if err != nil {
        log.Println(err)
        return
    }

    defer f_out.Close()

    if session.AlreadyExists {
        _, err = f_out.WriteString("\n");
        if err != nil {
            log.Println(err)
        }
    }

    _, err = f_out.WriteString(v.Choices[0].Message.Content)
    if err != nil {
        log.Println(err)
    }
}
