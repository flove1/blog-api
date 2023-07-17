create table if not exists users (
	id bigserial PRIMARY KEY,
	username varchar(50) UNIQUE NOT NULL CONSTRAINT username_not_empty CHECK(length(username)>5),
	first_name varchar(50) NOT NULL CONSTRAINT first_name_not_empty CHECK(length(first_name)>0),
	last_name varchar(50) NOT NULL CONSTRAINT last_name_not_empty CHECK(length(last_name)>0),
	hashed_password varchar(255) NOT NULL
);