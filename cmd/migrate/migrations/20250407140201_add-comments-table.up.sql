CREATE TABLE IF NOT EXISTS comments (
    `id` INT UNSIGNED NOT NULL AUTO_INCREMENT,
    `post_id` INT UNSIGNED NOT NULL,
    `user_id` INT UNSIGNED NOT NULL,
    `text` VARCHAR(255) NOT NULL,

    PRIMARY KEY (id),
    FOREIGN KEY (post_id) REFERENCES posts(id),
    FOREIGN KEY (user_id) REFERENCES users(id)
);