DROP TABLE IF EXISTS flashcard;
DROP TABLE IF EXISTS flashcards;
CREATE TABLE flashcards (
  id         INT AUTO_INCREMENT NOT NULL,
  ownerId   INT NOT NULL,
  term      VARCHAR(128) NOT NULL,
  definition     VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`)
);

DROP TABLE IF EXISTS users;
CREATE TABLE users (
    id INT AUTO_INCREMENT NOT NULL,
    username VARCHAR(128) NOT NULL,
    hashed VARCHAR(128) NOT NULL,
    salt VARCHAR(128) NOT NULL,
    PRIMARY KEY (`id`)
);
DROP TABLE IF EXISTS sessions;
CREATE TABLE sessions (
    id INT AUTO_INCREMENT NOT NULL,
    user_id INT NOT NULL,
    token VARCHAR(128) NOT NULL,
    expiration DATETIME NOT NULL,
    created_at DATETIME NOT NULL,
    expired BOOLEAN DEFAULT 0,
    PRIMARY KEY (`id`)
);
