CREATE TABLE posts (
    id INT UNIQUE NOT NULL ,
    user_id INT NOT NULL,
    title TEXT,
    body TEXT
);