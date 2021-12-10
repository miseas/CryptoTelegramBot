
<p align="center">
  <h1 align="center">BITCOIN-TELEGRAM-BOT</h1>
</p>


<!-- TABLE OF CONTENTS -->
<details open="open">
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#tada-inspiration">Inspiration</a>
    </li>
    <li>
      <a href="#star-getting-started">Getting Started</a>
      <ul>
        <li><a href="#what-you-will-need">What you will need</a></li>
        <li><a href="#computer-installation">Installation</a></li>
        <li><a href="#white_check_mark-add-btb-to-your-telegram-channel">Add BTB to your Telegram Channel</a></li>
      </ul>
    </li>
    <li><a href="#battery-usage">Usage</a></li>
  </ol>
</details>

## :star: Getting started

### What you will need:

- You are going to need a computer or server where to host the bot.
- Git
- Golang v1.13
- A device with Telegram

### Installation

Open a Terminal and copy these commands (Linux & Mac devices):

```bash
cd ~
git clone https://github.com/tomassirio/BitcoinTelegramBot.git
cd ./BitcoinTelegramBot
mv .env.example .env
go get github.com/tomassirio/bitcoinTelegram
go run main.go
```

##### Warning: 
This won't work unless you replace the **REPLACE_WITH_TOKEN** on the .env file with the Token granted by @BotFather

### Add BTB to your Telegram Channel

Open [@BotFather](https://telegram.me/botfather) on telegram and create a new bot with it's __/newbot__ command.

Assign it a name. This name won't be the one that is shown on each message, so you can name it whatever you want.

@BotFather will grant you a Token. This token is the one that will replace the **REPLACE_WITH_TOKEN** on the .env.example file on this repository. (Don't forget to rename that file to .env)

You can also play a little bit more with @BotFather. For example you can use the __/setcommands__ to define the uses your bot has on the '/' icon:

```
price - Gets symbol actual price. Default to BTCEUR
historic - Gets a percentage between Today's and Yesterday's price. Default to BTCEUR
summary - Gets both the price and historic values. Default to BTCEUR
```

## Usage

Once the bot is running and added to your Telegram Group, you can use any of the following commands:

```sh
    * /price :  Gets symbol last price. Default to BTCEUR
    * /historic : Gets a percentage between Today's and Yesterday's price
    * /summary : Gets both the price and historic values
```
