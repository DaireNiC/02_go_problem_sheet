# Go Problem Sheet 2

## About
### Daire Ní Chatháin: g00334757
A set of exercises demonstrating go web apps 

Provided by: https://github.com/data-representation/data-representation.github.io/edit/master/problems/go-web-applications.md

## How to Clone this repo

```bash
git clone https://github.com/DaireNiC/02_go_problem_sheet.git
```
## Navigate to the folder

```bash
cd 02_go_problem_sheet
```
## Build and the web app:

```go
go build ./complete_guessing_game.go
```

## Run the exe:

```bash
./complete_guessing_game.go
```

After running the exe you can also examine the response with curl. The -v flag is short hand for verbose mode. 
```curl
curl 127.0.0.1:8080 -v
curl 127.0.0.1:8080/guess -v
```
