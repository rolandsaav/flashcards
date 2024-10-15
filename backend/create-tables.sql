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
    username VARCHAR(32) NOT NULL,
    hashed VARCHAR(128) NOT NULL,
    salt VARCHAR(32) NOT NULL,
    PRIMARY KEY (`id`)
);
DROP TABLE IF EXISTS sessions;
CREATE TABLE sessions (
    id INT AUTO_INCREMENT NOT NULL,
    user INT NOT NULL,
    expiration TIME NOT NULL,
    created TIME NOT NULL,
    PRIMARY KEY (`id`)
);
