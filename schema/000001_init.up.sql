CREATE TABLE users
(
    id serial not null unique,
    name varchar(255) not null
);

CREATE TABLE msgs
(
    id serial not null unique,
    sender_id integer not null references users(id),
    receiver_id integer not null references users(id),
    message varchar(255) not null
);