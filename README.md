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
ask <your_query_here>
```
Or:  
```sh
ask
<your_query_here>
<can_stretch_over>
<multiple_lines>
<enter>
<enter>
```
On the second _enter_ click your multi-line query will be sent to Jef, who will
return with a response shortly 😁.    
  
**NOTE:** Do not type in _""_ or _\\_ as support for them has not been added yet!
