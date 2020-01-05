-- +migrate Up
CREATE TABLE poep_main_db.poem_theme
(id SERIAL NOT NULL,
created_at TIMESTAMP NOT NULL default NOW(),
title VARCHAR(50) NOT NULL,
ntag INT NOT NULL,
detail TEXT NOT NULL,
npoem INT NOT NULL default 0,
answer_length_min INT NOT NULL,
answer_length_max INT NOT NULL,
theme_setter_name VARCHAR(20),
PRIMARY KEY(id));

CREATE TABLE poep_main_db.poem_tag
(id SERIAL NOT NULL,
tag VARCHAR(20) NOT NULL,
poem_theme_id BIGINT UNSIGNED NOT NULL,
PRIMARY KEY(id),
FOREIGN KEY(poem_theme_id) REFERENCES poem_theme(id));

CREATE TABLE poep_main_db.poem
(id SERIAL NOT NULL,
poem_theme_id BIGINT UNSIGNED NOT NULL,
nfav int NOT NULL default 0,
date_created TIMESTAMP NOT NULL default NOW(),
answerer_name VARCHAR(20),
answer_text TEXT NOT NULL,
PRIMARY KEY(id),
foreign key(poem_theme_id) references poem_theme(id));

-- +migrate Down
DROP TABLE poepdb.poem_theme;
DROP TABLE poepdb.poem_tag;
DROP TABLE poepdb.poem;