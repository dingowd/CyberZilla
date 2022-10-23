CREATE DATABASE IF NOT EXISTS users;

USE users;

CREATE TABLE IF NOT EXISTS usersgroups(
    group_id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(5) UNIQUE NOT NULL ,
    r_create BOOL NOT NULL,
    r_read BOOL NOT NULL,
    r_update BOOL NOT NULL,
    r_delete BOOL NOT NULL
    );
INSERT INTO usersgroups (name, r_create, r_read, r_update, r_delete)
VALUES
    ('admin', true, true, true, true),
    ('user', false, false, false, false);

CREATE TABLE IF NOT EXISTS users(
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(30) UNIQUE NOT NULL,
    group_id INT,
    pass VARCHAR(256) NOT NULL,
    email VARCHAR(128) NOT NULL,
    FOREIGN KEY (group_id) REFERENCES usersgroups (group_id)
    );
INSERT INTO users (name, group_id, pass, email)
VALUES ('admin', (select group_id from usersgroups where name = 'admin'), 'ghbdtn', 'example@example.com');