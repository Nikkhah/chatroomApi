create table User
(
    id bigint NOT NULL AUTO_INCREMENT,
    username varchar(20) not null,
    password varchar(20) not null,
    PRIMARY KEY (id)
);

create table Chatroom
(
    id bigint NOT NULL AUTO_INCREMENT,
    name     varchar(20) not null,
    created_at   bigint      not null,
    updated_at   bigint      not null,
    user_id bigint not null,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES User(id)
);

create table User_Chatroom
(
    id bigint NOT NULL AUTO_INCREMENT,
    chatroom_id bigint not null,
    user_id bigint not null,
    PRIMARY KEY (id),
    FOREIGN KEY (user_id) REFERENCES User(id),
    FOREIGN KEY (chatroom_id) REFERENCES Chatroom(id)
);