# DISYS_AuctionHouse
## How to get started
1. First you can build and run the docker compose file - this will automatically start the auction, replicas and frontends. Use the following command in your terminal:
```
docker-compose build && docker-compose up
```
2. Secondly you need to connect a client (thorugh which you bid and query the auction) to a front-end. There are two front-ends which have the ports - :1400 and :1401. You can choose either of them to type into [port] in the following command, which you type into a new terminal:
```
go run ./Client/client.go [Your name] [port]
```
Example:
```
go run ./Client/client.go Emil :1400
```
3. You are now ready to bid!

(note: each time a auction ends, you have to start over in order to start a new auction.)
