package main

import "time"

const (
	salt         = "writedrunkeditsober"
	secret       = "bGl2ZWZhc3RkaWV5b3VuZw=="
	tokenLife    = time.Hour
	tokenRefresh = tokenLife / 2
)

const initialQuery = `

CREATE TABLE users (
	userId INTEGER PRIMARY KEY AUTOINCREMENT,
	role TEXT NOT NULL DEFAULT 'user',
	registered DATETIME DEFAULT CURRENT_TIMESTAMP,
	username TEXT NOT NULL UNIQUE,
	email TEXT NOT NULL UNIQUE,
	password TEXT NOT NULL,
	fullname TEXT NOT NULL,
	language TEXT NOT NULL DEFAULT 'en',
	status INTEGER NOT NULL DEFAULT 123,
	UNIQUE (username, email) );

CREATE TABLE categories (
	categoryId INTEGER PRIMARY KEY AUTOINCREMENT,
	created DATETIME DEFAULT CURRENT_TIMESTAMP,
	name TEXT NOT NULL,
	description TEXT NOT NULL );

CREATE TABLE posts (
	postId INTEGER PRIMARY KEY AUTOINCREMENT,
	posted DATETIME DEFAULT CURRENT_TIMESTAMP,
	userId INTEGER NOT NULL,
	title TEXT NOT NULL,
	text TEXT NOT NULL,
	categories TEXT NOT NULL,
	status INTEGER DEFAULT 1 );

CREATE TABLE postReactions (
	reactionId INTEGER PRIMARY KEY AUTOINCREMENT,
	reacted DATETIME DEFAULT CURRENT_TIMESTAMP,
	postId INTEGER NOT NULL,
	userId INTEGER NOT NULL,
	reaction TEXT DEFAULT 'idle',
	UNIQUE (userId, postId));

CREATE TABLE comments (
	commentId INTEGER PRIMARY KEY AUTOINCREMENT,
	commented DATETIME DEFAULT CURRENT_TIMESTAMP,
	postId INTEGER NOT NULL,
	userId INTEGER NOT NULL,
	comment TEXT NOT NULL );	

CREATE TABLE commentReactions (
	reactionId INTEGER PRIMARY KEY AUTOINCREMENT,
	reacted DATETIME DEFAULT CURRENT_TIMESTAMP,
	commentId INTEGER NOT NULL,
	userId INTEGER NOT NULL,
	reaction TEXT DEFAULT 'idle',
	UNIQUE (userId, commentId));

INSERT INTO users(email, username, password, fullname, role) 
	values(
		'azakost@gmail.com',
		'azakost',
		'$2a$04$kitMig4Sfj/Id0C85pysxu3MQbFMC0qXDn5j4RA8ZoI8P9GMcE8Vm',
		'Azamat Alimbayev',
		'admin'
	);
` +

	// Testing queries - should be deleted after
	`
INSERT INTO categories(name, description)
	values(
		'golang',
		'This category is for gophers!'
	);

INSERT INTO categories(name, description)
	values(
		'js',
		'JavaScript is a mother of all web devs!'
	);	

INSERT INTO posts(userId, title, text, categories) 
	values(
		'1',
		'The best title ever!',
		'This is a very long text written for testing purposes!',
		'"1","2"'
	);

INSERT INTO posts(userId, title, text, categories) 
	values(
		'2',
		'Hello!',
		'sddsd',
		'"2"'
	);

INSERT INTO comments(postId, userId, comment) 
	values('1',	'2', 'Sad! Not for udots!');

INSERT INTO comments(postId, userId, comment) 
	values('1',	'1', 'Best content!');

INSERT INTO comments(postId, userId, comment) 
	values('2',	'1', 'Bester content!');

INSERT INTO reactions(postId, userId, reaction) 
	values('1',	'1', 'like');

`
