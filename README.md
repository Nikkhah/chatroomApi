chat application with golang and mysql
the structure of the code is based on hexagonal architecture and try to respect dependency inversion pattern. 

## to run the application:
1- install MySQL
2- set nvironment variables: MYSQL_USERNAME and MYSQL_PASSWORD
3- run script.sh to create database
4- run this commad: go build -o bin/chatroom
5- and run thw app: bin/chatroom

## curl command to create chatroom

curl --location 'http://localhost:2009/api/v1/create/chatroom' --header 'Content-Type: application/json' --data '{"name":"chat1", "user_id":"1"}'