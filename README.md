# Go Slack Pokebot

> Simple Slack bot to search for Pokemon info

## How to use

- You need to have **Go** configured in your machine, if you don't have take a look **[here](https://golang.org/doc/install)**

- If you don't have **Glide (Go Package Manager)**, get it **[here](https://github.com/Masterminds/glide)**

- Clone this **repo**

- Enter it and install the dependencies

```
glide install
```

- Create a **[new bot user integration](https://my.slack.com/services/new/bot)** on your Slack

- Create a `config/token.json` file

```
cp config/token_example.json config/token.json
```

- Enter the `config/token.json` file and put your **Bot API Key** there

- Run the **bot**

```
go run pokebot.go
```

## Commands

- `help:` Lists all other commands with description and examples
- `pokemon:` Gets info about the given Pokemon number or name **(in progress)**
