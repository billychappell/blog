drop table posts;
drop table comments;

create table posts (
	id			serial primary key,
	title 		text,
	createdat 	timestamp not null,
	author  	varchar(255),
	description varchar(255),
	content		text,
	imageurl varchar(255)
);

create table comments (
	id			serial primary key,
	uuiud		varchar(64) not null unique,
	post 		integer refereneces posts(id),
	content		text,
	author		varchar(255),
	createdat	timestamp not null
);