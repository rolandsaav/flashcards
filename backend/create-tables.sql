DROP TABLE IF EXISTS flashcard;
CREATE TABLE flashcard (
  id         INT AUTO_INCREMENT NOT NULL,
  ownerId   INT NOT NULL,
  term      VARCHAR(128) NOT NULL,
  definition     VARCHAR(255) NOT NULL,
  PRIMARY KEY (`id`)
);

