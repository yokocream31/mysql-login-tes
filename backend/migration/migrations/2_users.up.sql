create table if not exists users (
    id integer auto_increment primary key,
    username varchar(40),
    password varchar(40)
)