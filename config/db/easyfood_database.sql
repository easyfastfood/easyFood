/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;	
/*!40101 SET NAMES utf8 */;	
/*!50503 SET NAMES utf8mb4 */;	
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=1 */;	
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;

SET foreign_key_checks = 1;

-- Dumping database structure for easyfood
CREATE DATABASE IF NOT EXISTS `easyfood` /*!40100 DEFAULT CHARACTER SET utf8 */;
USE `easyfood`;

-- Dumping structure for table easyfood.categorias
CREATE TABLE IF NOT EXISTS `categorias` (
  `id` INT(11) unsigned NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- Dumping structure for table easyfood.restaurantes
CREATE TABLE IF NOT EXISTS `restaurantes` (
  `id` INT(11) unsigned NOT NULL AUTO_INCREMENT,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=UTF8;

INSERT INTO `restaurantes` (`id`) VALUES
	(1),
	(2);

-- Dumping data for table easyfood.categorias: ~0 rows (approximately)
/*!40000 ALTER TABLE `categorias` DISABLE KEYS */;
/*!40000 ALTER TABLE `categorias` ENABLE KEYS */;

-- Dumping structure for table easyfood.pratos
CREATE TABLE IF NOT EXISTS `pratos` (
  `id` INT(11) unsigned NOT NULL AUTO_INCREMENT,
  `id_restaurante` INT(11) unsigned NOT NULL DEFAULT 0,
  `id_categoria` INT(11) unsigned DEFAULT NULL,
  `nome` char(50) NOT NULL DEFAULT '0',
  `preco` decimal(10,2) unsigned NOT NULL DEFAULT 0.00,
  `tempo_de_preparo` tinyint(3) unsigned NOT NULL DEFAULT 0 COMMENT 'MINUTOS',
  PRIMARY KEY (`id`),
  KEY `id_restaurante` (`id_restaurante`),
  KEY `id_categoria` (`id_categoria`),
  CONSTRAINT `id_categoria` FOREIGN KEY (`id_categoria`) REFERENCES `categorias` (`id`) ON DELETE SET NULL ON UPDATE CASCADE,
  CONSTRAINT `id_restaurante` FOREIGN KEY (`id_restaurante`) REFERENCES `restaurantes` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- Dumping data for table easyfood.pratos: ~0 rows (approximately)
/*!40000 ALTER TABLE `pratos` DISABLE KEYS */;
INSERT INTO `pratos` (`id`, `id_restaurante`, `id_categoria`, `nome`, `preco`, `tempo_de_preparo`) VALUES
	(1, 1, NULL, 'tilápia', 45.00, 30),
	(2, 1, NULL, 'batata frita', 15.00, 15),
	(3, 1, NULL, 'batata frita com queijo', 18.00, 16),
	(4, 2, NULL, 'iscas de frango acebolada', 20.00, 20),
	(5, 2, NULL, 'filé parmegiana', 40.00, 40);
/*!40000 ALTER TABLE `pratos` ENABLE KEYS */;

-- Dumping data for table easyfood.restaurantes: ~0 rows (approximately)
/*!40000 ALTER TABLE `restaurantes` DISABLE KEYS */;
/*!40000 ALTER TABLE `restaurantes` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;

CREATE TABLE IF NOT EXISTS `usuarios` (
  `id` INT(11) unsigned NOT NULL AUTO_INCREMENT,
  `primeiro_nome` varchar(32) NOT NULL,
  `ultimo_nome` varchar(32) NOT NULL,
  `telefone` varchar(11) NOT NULL,
  `email` varchar(50) NOT NULL,
  `senha_hash` BINARY(64) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `usuarios` (`primeiro_nome`, `ultimo_nome`, `telefone`, `email`, `senha_hash`) VALUES
  ('Joao', 'Pedro', '31984464729', 'joaopedro@gmail.com', SHA1('joaopedro2010')),
  ('Maria', 'Lima', '35987432164', 'marialima@hotmail.com', SHA1('ml15122015')),
  ('Carlos', 'Antunes', '37984455792', 'carlos_antunes12@outlook.com', SHA1('carlitos1212'));
