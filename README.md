# Telegram shares bot
Telegram bot get shares signals from database

## Build Setup:
``` bash
# Clone repository:
git clone https://github.com/Paramosch/predictstock-backend-eng.git

# create database:
example shares_bot
```

## Create telegram bot
Use **BotFather** for create telegram bot and get token.

## Create configuration file .env
``` bash
# create .env file from example
cp .env.example .env 
```
configure database connection

configure **TELEGRAM_TOKEN** from token **BotFather**

## Run the Bot:
``` bash
# run build:
go build github.com/Paramosch/predictstock-backend-eng
```