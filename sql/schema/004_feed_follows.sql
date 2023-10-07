-- +goose Up

CREATE TABLE feed_follows (
    id INT PRIMARY KEY NOT NULL AUTO_INCREMENT,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL, 
    user_id INT NOT NULL,
    feed_id INT NOT NULL,
    CONSTRAINT PC_uid_fid UNIQUE(user_id, feed_id),
    CONSTRAINT FK_follow_feed FOREIGN KEY (feed_id) REFERENCES feeds(id) ON DELETE CASCADE,
    CONSTRAINT FK_follow_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down

DROP TABLE feeds_follows;