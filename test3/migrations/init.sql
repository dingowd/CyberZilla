CREATE DATABASE IF NOT EXISTS users;

CREATE TABLE IF NOT EXISTS u_groups(
    group_id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(5) UNIQUE NOT NULL ,
    r_create BOOL NOT NULL,
    r_view BOOL NOT NULL,
    r_change BOOL NOT NULL,
    r_delete BOOL NOT NULL
    );
INSERT INTO u_groups (name, r_create, r_view, r_change, r_delete)
VALUES
    ('admin', true, true, true, true),
    ('user', false, false, false, false);

CREATE TABLE IF NOT EXISTS users(
    id INT PRIMARY KEY AUTO_INCREMENT,
    group_id INT,
    name VARCHAR(30) UNIQUE NOT NULL,
    pass VARCHAR(256) NOT NULL,
    email VARCHAR(128) NOT NULL,
    FOREIGN KEY (group_id) REFERENCES u_groups (group_id)
    );
INSERT INTO users (name, group_id, pass, email)
VALUES ('admin', (select group_id from u_groups where name = 'admin'), 'ghbdtn', 'example@example.com');