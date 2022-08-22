CREATE DATABASE IF NOT EXISTS `cards-api`;

USE `cards-api`;

CREATE TABLE IF NOT EXISTS `cards` (
  `id` int(4) NOT NULL AUTO_INCREMENT,
  `suit` varchar(255) DEFAULT NULL,
  `value` varchar(255) DEFAULT NULL,
  `code` varchar(8) DEFAULT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_cards_code` (`code`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS `decks` (
  `id` varchar(255) NOT NULL,
  `shuffled` tinyint(4) DEFAULT '0',
  `total` tinyint(4) DEFAULT '0',
  `remaining` tinyint(4) DEFAULT '0',
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

CREATE TABLE IF NOT EXISTS `deck_cards` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `deck_id` varchar(255) DEFAULT NULL,
  `card_id` int(4) DEFAULT NULL,
  `is_drawn` tinyint(4) DEFAULT '0',
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  PRIMARY KEY (`id`),
  KEY `fk_deck_cards_deck` (`deck_id`),
  KEY `fk_deck_cards_card` (`card_id`),
  CONSTRAINT `fk_deck_cards_card` FOREIGN KEY (`card_id`) REFERENCES `cards` (`id`),
  CONSTRAINT `fk_deck_cards_deck` FOREIGN KEY (`deck_id`) REFERENCES `decks` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

SET FOREIGN_KEY_CHECKS=0;
TRUNCATE TABLE `cards`;
SET FOREIGN_KEY_CHECKS=1;

INSERT INTO `cards` (`suit`, `value`, `code`) VALUES
    ('SPADES', 'ACE', 'AS'),
    ('SPADES', '2', '2S'),
    ('SPADES', '3', '3S'),
    ('SPADES', '4', '4S'),
    ('SPADES', '5', '5S'),
    ('SPADES', '6', '6S'),
    ('SPADES', '7', '7S'),
    ('SPADES', '8', '8S'),
    ('SPADES', '9', '9S'),
    ('SPADES', '10', '10S'),
    ('SPADES', 'JACK', 'JS'),
    ('SPADES', 'QUEEN', 'QS'),
    ('SPADES', 'KING', 'KS'),
    ('DIAMONDS', 'ACE', 'AD'),
    ('DIAMONDS', '2', '2D'),
    ('DIAMONDS', '3', '3D'),
    ('DIAMONDS', '4', '4D'),
    ('DIAMONDS', '5', '5D'),
    ('DIAMONDS', '6', '6D'),
    ('DIAMONDS', '7', '7D'),
    ('DIAMONDS', '8', '8D'),
    ('DIAMONDS', '9', '9D'),
    ('DIAMONDS', '10', '10D'),
    ('DIAMONDS', 'JACK', 'JD'),
    ('DIAMONDS', 'QUEEN', 'QD'),
    ('DIAMONDS', 'KING', 'KD'),
    ('CLUBS', 'ACE', 'AC'),
    ('CLUBS', '2', '2C'),
    ('CLUBS', '3', '3C'),
    ('CLUBS', '4', '4C'),
    ('CLUBS', '5', '5C'),
    ('CLUBS', '6', '6C'),
    ('CLUBS', '7', '7C'),
    ('CLUBS', '8', '8C'),
    ('CLUBS', '9', '9C'),
    ('CLUBS', '10', '10C'),
    ('CLUBS', 'JACK', 'JC'),
    ('CLUBS', 'QUEEN', 'QC'),
    ('CLUBS', 'KING', 'KC'),
    ('HEARTS', 'ACE', 'AH'),
    ('HEARTS', '2', '2H'),
    ('HEARTS', '3', '3H'),
    ('HEARTS', '4', '4H'),
    ('HEARTS', '5', '5H'),
    ('HEARTS', '6', '6H'),
    ('HEARTS', '7', '7H'),
    ('HEARTS', '8', '8H'),
    ('HEARTS', '9', '9H'),
    ('HEARTS', '10', '10H'),
    ('HEARTS', 'JACK', 'JH'),
    ('HEARTS', 'QUEEN', 'QH'),
    ('HEARTS', 'KING', 'KH');