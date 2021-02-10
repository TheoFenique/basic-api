CREATE TABLE `articles`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `username` varchar(255) NULL,
  `title` varchar(255) NULL,
  `content` text NULL,
  `creation_date` datetime(0) NULL,
  PRIMARY KEY (`id`)
);

CREATE TABLE `comments`  (
  `id` int NOT NULL,
  `username` varchar(255) NULL,
  `content` text NULL,
  `creation_date` datetime(0) NULL,
  `article_id` int NULL,
  PRIMARY KEY (`id`)
);

CREATE TABLE `users`  (
  `id` int NOT NULL AUTO_INCREMENT,
  `firstname` varchar(255) NOT NULL,
  `lastname` varchar(255) NOT NULL,
  `dob` datetime NOT NULL,
  `city` varchar(255) NOT NULL,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
);

ALTER TABLE `articles` ADD CONSTRAINT `fk_articles_comments_1` FOREIGN KEY (`id`) REFERENCES `comments` (`id`);
ALTER TABLE `comments` ADD CONSTRAINT `fk_comments_users_1` FOREIGN KEY (`username`) REFERENCES `users` (`username`);
ALTER TABLE `users` ADD CONSTRAINT `fk_users_articles_1` FOREIGN KEY (`id`) REFERENCES `articles` (`id`);

