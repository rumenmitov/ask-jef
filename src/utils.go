package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

func getModels() ([]Model, error) {
    response, err := http.Get("http://localhost:8080/models");
    if err != nil {
        return nil, err;
    }

    defer response.Body.Close();

    response_body, err := io.ReadAll(response.Body);
    if err != nil {
        return nil, err;
    }

    var modelsJSON ModelsRequest;

    if err := json.Unmarshal(response_body, &modelsJSON); err != nil {
        return nil, err;
    }

    return modelsJSON.Data, nil;
}

func selectModel() (string, error) {

    available_models, err := getModels();
    if err != nil {
        return "", err;
    }

    if len(available_models) == 0 {
        log.Panic("Error! No models available! Please download them in your LocalAI/models directory!\n");
    }

    var model string;

    for {

        for _, availavailable_model := range available_models {
            if model == availavailable_model.Id {
                return model, nil;
            }
        }

        fmt.Printf("%s", B_Purple + "No valid model specified! Enter the name of the model you wish to use:\n" + Reset);

        for _, model := range available_models {
            fmt.Printf("%s", B_Purple + "- " + model.Id + "\n" + Reset);
        }

        fmt.Scanf("%s", &model);
    }

}

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
