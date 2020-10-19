package main

import "time"

const (
	salt         = "writedrunkeditsober"
	secret       = "bGl2ZWZhc3RkaWV5b3VuZw=="
	tokenLife    = time.Hour
	tokenRefresh = tokenLife / 2
	dbname       = "database.db"
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
	comment TEXT NOT NULL,
	status INTEGER DEFAULT 1 );

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


`
