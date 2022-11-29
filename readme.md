# serra

Serra is my personal *Magic: The Gathering* collection tracker.

It began as a holiday project in winter 2021/2022 because I was frustrated of
Collection Tracker Websites that are:

* Pain to use
* Want ~$10 a month
* Don't have the features I want

So I started my own Collection Tracker using [Golang](https://golang.org),
[MongoDB](https://mongodb.com) and [Scryfall](https://scryfall.com) to have
an overview in what cards you own and what value they are.

## What Serra does

* Tracks prices
* Calculates statistics
* Query/filter all of your cards
* Shows what cards/sets do best in value development.

## What Serra does not

* Does not give a shit about conditions (NM, M, GD...)
* Does not track if card is foil or not (may come in the future)
* Is not configurable to have Dollar/US Prices

# Quickstart

on macOS you can use

    brew install noqqe/tap/serra

on Linux/BSD/Windows you can download binaries from

    https://github.com/noqqe/serra/releases

After that you need to spin up a MongoDB yourself or use the docker-compose
setup included in this Repo:

    docker-compose up -d
    export MONGODB_URI='mongodb://root:root@localhost:27017'

After that, you can add a card

    ./serra add usg/17

Start exploring :) (the more cards you add, the more fun it is)

# Usage

The overall usage is described in `--help` text. But below are some examples.
```
./serra
Usage:
  serra add <cardid>... [--count=<number>]
  serra remove <cardid>...
  serra cards [--rarity=<rarity>] [--set=<setcode>] [--sort=<sort>]
  serra card <cardid>...
  serra tops [--limit=<limit>]
  serra flops [--limit=<limit>]
  serra missing <setcode>
  serra set <setcode>
  serra sets
  serra update
  serra stats
```

## Add

To add a card to your collection.

![](https://github.com/noqqe/serra/blob/main/imgs/add.png)

## Cards

Query all of your cards with filters

![](https://github.com/noqqe/serra/blob/main/imgs/cards.png)

## Sets

List all your sets

![](https://github.com/noqqe/serra/blob/main/imgs/sets.png)

## Set

Show details of a single set

![](https://github.com/noqqe/serra/blob/main/imgs/set.png)

## Stats

Calculate some stats for all of your cards

![](https://github.com/noqqe/serra/blob/main/imgs/stats.png)

## Tops

Show what cards/set gained most value

![](https://github.com/noqqe/serra/blob/main/imgs/tops.png)

## Flops

Show what cards/set lost most value

![](https://github.com/noqqe/serra/blob/main/imgs/flops.png)

## Update

The update mechanism iterates over each card in your collection and fetches
its price. After all cards you own in a set are updated, the set value will
update. After all Sets are updated, the whole collection value is updated.

![](https://github.com/noqqe/serra/blob/main/imgs/update.png)

## Adding all those cards, manually?

Yes. While there are serveral OCR/Photo Scanners for mtg cards, I found they
are not accurate enough. They guess Editions wrong, they have problems with
blue/black cards and so on.

I add my cards using a tiny shell wrapper, since they are sorted by editions
anyways.

```
./add-card-wrapper.fish usg
read> 17
Updating Card "Herald of Serra" amount to 2

read> 18
...
```

Its basically typing 2-3 digit numbers and hitting enter. I was way faster
with this approach then Smartphone scanners.

# Development

## Install

    go build .
    ./serra

## MongoDB Operations

A few commands that do backups and exports of your data inside of the docker
container.

Do a database dump

    mongodump  -u root -p root --authenticationDatabase admin -d serra -o /backup/

Do a collection export to json

    mongoexport  -u root -p root --authenticationDatabase admin -d serra -c cards > /backup/cards.json
    mongoexport  -u root -p root --authenticationDatabase admin -d serra -c sets > /backup/sets.json
    mongoexport  -u root -p root --authenticationDatabase admin -d serra -c total > /backup/total.json
