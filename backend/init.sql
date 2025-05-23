CREATE DATABASE IF NOT EXISTS SAAS;

USE SAAS;

CREATE TABLE IF NOT EXISTS users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    first_name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    has_upgraded BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE IF NOT EXISTS user_sessions (
    user_id INT NOT NULL,
    expires_at DATETIME NOT NULL,
    token VARCHAR(255) PRIMARY KEY NOT NULL UNIQUE,
    FOREIGN KEY (user_id) REFERENCES users(id)
);
