#!/bin/bash

# Set database credentials
DB_HOST="localhost"
DB_USER=${MYSQL_USERNAME}
DB_PASSWORD=${MYSQL_PASSWORD}
DB_NAME="chatroom_db"

# Create the database
mysql -h $DB_HOST -u $DB_USER -p$DB_PASSWORD -e "CREATE DATABASE IF NOT EXISTS $DB_NAME"

# Create the table
mysql -h $MYSQL_HOST -u $DB_USER -p$DB_PASSWORD $MYSQL_DB -e "
CREATE TABLE User (
    id bigint NOT NULL AUTO_INCREMENT,
    username varchar(20) not null,
    password varchar(20) not null,
    PRIMARY KEY (id)
);

CREATE TABLE Chatroom (
    id bigint NOT NULL AUTO_INCREMENT,
    name     varchar(20) not null,
    created_at   bigint      not null,
    updated_at   bigint      not null,
    user_id bigint not null,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES User(id)
);

CREATE TABLE User_Chatroom (
    id bigint NOT NULL AUTO_INCREMENT,
    chatroom_id bigint not null,
    user_id bigint not null,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES User(id),
    FOREIGN KEY (chatroom_id) REFERENCES Chatroom(id)
);
"