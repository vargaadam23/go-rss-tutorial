-- +goose Up

CREATE TABLE posts (
    id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL, 
    title TEXT NOT NULL,
    description TEXT,
    published_at TIMESTAMP NOT NULL,
    url VARCHAR(100) UNIQUE NOT NULL,
    feed_id INT NOT NULL,
    CONSTRAINT FK_post_feed FOREIGN KEY (feed_id) REFERENCES feeds(id) ON DELETE CASCADE
);

-- +goose Down

DROP TABLE posts;