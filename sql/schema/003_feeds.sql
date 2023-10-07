-- +goose Up

CREATE TABLE feeds (
    id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL, 
    url VARCHAR(100) UNIQUE NOT NULL,
    name TEXT NOT NULL,
    user_id INT NOT NULL,
    CONSTRAINT FK_feedUser FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down

DROP TABLE feeds;