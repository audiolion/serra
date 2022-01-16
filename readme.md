# serra

# Install

    go build .
    ./serra

# Todo

* termui overview
* prices since the beginning
* total card value column with history in mongodb

# What its not

* Gives a shit about conditions (NM, M, GD...)
* If the card is foil

# Cheatsheet

Find cards that increased prices

    db.cards.find({$expr: {$gt: [{$arrayElemAt: ["$serra_prices", -2]}, {$arrayElemAt: ["$serra_prices", -1]}]}}, {name:1})

Update card Price

		db.cards.update(
		{'_id':'8fa2ecf9-b53c-4f1d-9028-ca3820d043cb'},
		{$set:{'serra_updated':ISODate("2021-11-02T09:28:56.504Z")},
		$push: {"serra_prices": { date: ISODate("2021-11-02T09:28:56.504Z"), value: 0.1 }}});

Set value

    db.cards.aggregate([{ $group: { _id: { set: "$set" }, value: { $sum: { $multiply: ["$prices.eur", "$serra_count"] } }, count: { $sum: 1 } } }])

Color distribution

     db.cards.aggregate([{ $group: { _id: { color: "$colors" }, count: { $sum: 1 } } }])

Calculate value of all sets

    db.sets.aggregate({$match: {serra_prices: {$exists: true}}}, {$project: {name: 1, "totalValue": {$arrayElemAt: ["$serra_prices", -1]} }}, {$group: {_id: null, total: {$sum: "$totalValue.value" }}})

# MongoDB Operations

Do a database dump

    mongodump  -u root -p root --authenticationDatabase admin -d serra -o /backup/

Do a collection export to json

    mongoexport  -u root -p root --authenticationDatabase admin -d serra -c cards
