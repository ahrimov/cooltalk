DROP TABLE IF EXISTS  users;

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR (50) UNIQUE NOT NULL,
    password VARCHAR (50) NOT NULL,
    email VARCHAR (255) UNIQUE NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO users
    (username, password, email)
VALUES 
    ('test_user', '1', 'test@gmail.com'),
    ('antoxa', 't', 'antonxa@gmail.com');
