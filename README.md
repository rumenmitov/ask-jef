# Ask Jef!
> Your terminal assistant 😉  
  
![Example](https://github.com/rumenmitov/ask-jef/assets/108842741/606a45b6-557e-4259-8792-53e95cdecb74)
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
1. Clone this repository:
```sh
git clone https://github.com/rumenmitov/ask-jef.git
```
2. Run the installation script:
```sh
cd ask-jef
make install
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
```
On the _enter_ after an empty line, your multi-line query will be sent to Jef, who will
return with a response shortly 😁.    

**NOTE:** Form multi-line queries quotation marks (_""_) are not needed!  
  
You may also specify a desired model by typing:
```sh
ask -m <model_name> "your query here"
```
**NOTE:** The model name should correspond to the model in your _LocalAI/models/_ directory.  
**NOTE:** It is also possible to set the desired model in `$XDG_CONFIG_HOME/ask-jef/ask.env`.  
  
To enter the contents of a file you can type:  
```sh
ask -f <file_1> -f <file_2> "your query here"
```
  
Each response is saved in a unique file (referred to as a _session_) in the `$XDG_CACHE_HOME` directory.  
To save a response under a desired session name, please enter the command:
```sh
ask -s <session_name> "your query here"
```
**NOTE:** The above command could also be used to continue from where you left off in a previous session.  
  
To list the available sessions:  
```sh
ask -l
```
   
To print the results of a session:
```sh
ask -ss <session_name>
```
  
To delete a session:
```sh
ask -r <session_name>
```

***
![GitHub stars](https://img.shields.io/github/stars/rumenmitov/ask-jef?style=social)
  
    
To rename a session:
```sh
ask -mv <old_session_name> <new_session_name>
```
