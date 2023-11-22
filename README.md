# Ask Jef!
> Your terminal assistant 😉
***
## Table of Contents
1. [Installation](#installation)  
1.1 [Prerequisites](#prerequisites)  
1.2 [Building](#building)  
2. [Usage](#usage)  
***
## Installation
### Prerequisites
Before you start make sure you have set up the following:  
- [LocalAI](https://localai.io/)
- [Go](https://go.dev)
### Building
1. Clone this repository and go into it:
```sh
git clone https://github.com/rumenmitov/ask-jef.git
cd ask-jef
```
2. Run the installation script:
```sh
go build -o /usr/local/bin/ask src/*.go
```
## Usage
There two ways of running the program. Either:  
```sh
ask "your query here"
```
Or (if you want to use multiple lines for your query):  
```sh
ask
...
<enter>
<enter>
```
On the second _enter_ click your multi-line query will be sent to Jef, who will
return with a response shortly 😁.    

**NOTE:** Form multi-line queries quotation marks (_""_) are not needed!  
  
You may also specify a desired model by typing:
```sh
ask -m <model_name> "your query here"
```
  
**NOTE:** The default model is _luna-ai-llama2_.  
**NOTE:** The modeel name should correspond to the model in your _LocalAI/models/_ directory.  
