
<p align="center">
  <h1 align="center">CRYPTO-TELEGRAM-BOT</h1>
</p>


<!-- TABLE OF CONTENTS -->
<details open="open">
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#minimum-requirements">Minimum requirements</a></li>
        <li><a href="#add-bot-to-your-telegram-channel">Add bot to your Telegram Channel</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    <li>
       <a href="#usage">Usage</a></li>
       <ul>
          <li><a href="#about-notifications">About notifications</a></li>
          <li><a href="#exchanges-supported">Exchanges supported</a></li>
      </ul>
    </li>
  </ol>
</details>

## Getting started

### Minimum requirements:

- Git
- Golang
- Telegram
- Computer/server to host the app

### Add bot to your Telegram Channel

Open [@BotFather](https://telegram.me/botfather) on telegram and create a new bot with it's __/newbot__ command.

Assign it a name. This name won't be the one that is shown on each message, so you can name it whatever you want.

@BotFather will grant you a Token. This token is the one that will replace the **REPLACE_WITH_TOKEN** on the .env.example file on this repository. (Don't forget to rename that file to .env)

You can use the __/setcommands__ to define the uses your bot has on the '/' icon:

```
price - Use price <symbol> . Gets symbol actual price. Default BTCEUR
historic - Use historic <symbol> . Gets a percentage between Today's and Yesterday's price. Default BTCEUR
summary - Use summary <symbol> . Gets both the price and historic values. Default BTCEUR
notification - Use add <symbol> <compare> <value>| remove <symbol> <compare> <value> | list . Get notified if the symbol es higher > or lower <
```

### Installation

Open a Terminal and copy these commands (Linux & Mac devices):

```bash
cd ~
git clone https://github.com/miseas/CryptoTelegramBot.git
cd ./CryptoTelegramBot
cp .env.example .env
go get cryptoTelegramBot
go run main.go
```

##### Warning: 
You must replace the **REPLACE_WITH_TOKEN** on the .env file with the Token granted by @BotFather


##  Usage

Once the bot is running, you can add it to a Telegram group and use any of the following commands:

```sh
    * /price: Use price <symbol> . Gets symbol actual price. Default BTCEUR
    * /historic: Use historic <symbol> . Gets a percentage between Today's and Yesterday's price. Default BTCEUR
    * /summary: Use summary <symbol> . Gets both the price and historic values. Default BTCEUR
    * /notification [add/remove/list] [a_symbol] [a_compare_symbol] [a_value]: . Get notified if the symbol es higher > or lower <
```

###  About Notifications

Notifications are still a WIP. You can add/remove or list your notifications. The notifications are checked every 1 min (or custom). After notified 3 times, it will be removed automatically from the list. 


### Exchanges supported 

(more will be added):
  - Binance
