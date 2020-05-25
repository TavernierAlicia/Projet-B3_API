-- phpMyAdmin SQL Dump
-- version 4.8.5
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1:3306
-- Generation Time: May 19, 2020 at 06:40 PM
-- Server version: 5.7.26
-- PHP Version: 7.2.18

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `orderndrink`
--

-- --------------------------------------------------------

--
-- Table structure for table `campaigns`
--

DROP TABLE IF EXISTS `campaigns`;
CREATE TABLE IF NOT EXISTS `campaigns` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `etab_id` int(11) NOT NULL,
  `pro_id` int(11) NOT NULL,
  `title` varchar(250) NOT NULL,
  `content` text NOT NULL,
  `date_begin` timestamp NOT NULL,
  `date_end` timestamp NOT NULL,
  `publish_date` timestamp NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `clients`
--

DROP TABLE IF EXISTS `clients`;
CREATE TABLE IF NOT EXISTS `clients` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `surname` varchar(100) NOT NULL,
  `profile_pic` varchar(250) DEFAULT NULL,
  `birth_date` date NOT NULL,
  `mail` varchar(250) NOT NULL,
  `phone_number` varchar(10) NOT NULL,
  `password` text NOT NULL,
  `token` varchar(250) NOT NULL,
  `inscription_date` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=10 DEFAULT CHARSET=latin1;

--
-- Dumping data for table `clients`
--

INSERT INTO `clients` (`id`, `name`, `surname`, `profile_pic`, `birth_date`, `mail`, `phone_number`, `password`, `token`, `inscription_date`) VALUES
(1, 'Jean', 'Dupont', '/path/to/pic', '1977-11-09', 'dupont-jean@gmail.com', '0607080900', '1234', 'hehe', '2020-04-28 17:10:43'),
(2, 'Jacques', 'Saint', '/path/to/pic', '2000-04-01', 'saintjacquesdu68@yahoo.com', '0607080900', '1234', 'haha', '2020-04-28 17:10:43'),
(3, 'Martin', 'Jean', '/path/to/pic', '1999-08-14', 'martin@live.fr', '0102030405', 'martin.tintin', 'hihi', '2020-04-29 01:37:38'),
(5, 'Test', 'Tests', '/pictures/rien.png', '2000-05-20', 'test@gmail.comm', '0102030405', '92f2fd99879b0c2466ab8648afb63c49032379c1', 'dcdb199e-2797-4041-8b26-08bc451dd47b', '2020-05-08 17:09:48'),
(6, 'Juliette', 'Jachere', '/path/to/pic', '1978-04-03', 'jj@caramail.fr', '0102030405', 'da39a3ee5e6b4b0d3255bfef95601890afd80709', '6d60e931-856c-4927-bca2-344be1cfe135', '2020-05-10 19:40:46'),
(7, 'Jean', 'Jachere', NULL, '1978-04-03', 'jj@caramail.fr', '0102030405', '7110eda4d09e062aa5e4a390b0a572ac0d2c0220', 'a5df21bd-438d-482d-a194-8a73e8a40ee1', '2020-05-15 17:46:00'),
(8, 'Jean', 'Jachere', NULL, '1978-04-03', 'jj@caramail.fr', '0102030405', '7110eda4d09e062aa5e4a390b0a572ac0d2c0220', '3354a837-b68d-4b4e-9e09-7b809b49bffe', '2020-05-15 17:46:24'),
(9, 'Testons', 'Tests', '/pictures/rien.png', '2000-05-20', 'test@gmail.comm', '0102030405', 'dd3bc42b2cbba792a371118cd1c87384c107bf6c', '6705c128-722e-430c-9cb4-d70f10f614fc', '2020-05-15 18:26:23');

-- --------------------------------------------------------

--
-- Table structure for table `commands`
--

DROP TABLE IF EXISTS `commands`;
CREATE TABLE IF NOT EXISTS `commands` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `client_id` int(11) NOT NULL,
  `etab_id` int(11) NOT NULL,
  `price` float DEFAULT NULL,
  `instructions` varchar(250) NOT NULL,
  `status` varchar(10) NOT NULL DEFAULT 'paid',
  `waiting_time` time DEFAULT NULL,
  `payment` varchar(100) NOT NULL,
  `tip` int(11) NOT NULL,
  `cmd_date` timestamp NOT NULL DEFAULT current_timestamp(),
  `reception_date` timestamp DEFAULT '0000-00-00 00:00:00',
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=17 DEFAULT CHARSET=latin1;

--
-- Dumping data for table `commands`
--

INSERT INTO `commands` (`id`, `client_id`, `etab_id`, `price`, `instructions`, `status`, `waiting_time`, `payment`, `tip`, `cmd_date`, `reception_date`) VALUES
(1, 1, 1, 5, '', 'pending', NULL, '', 0, '2020-05-12 22:13:48', NULL),
(2, 2, 1, 0, '', 'pending', NULL, '', 0, '2020-05-12 22:13:48', NULL),
(3, 5, 1, NULL, 'Pas de glacons', 'paid', '00:00:00', 'Gpay', 0, '2020-05-12 23:09:17', NULL),
(4, 5, 1, NULL, 'Pas de glacons', 'paid', '00:00:00', 'Gpay', 0, '2020-05-12 23:35:22', NULL),
(5, 5, 1, NULL, 'Pas de glacons', 'paid', '00:00:00', 'Gpay', 0, '2020-05-12 23:41:26', NULL),
(6, 5, 1, NULL, 'Pas de glacons', 'paid', '00:00:00', 'Gpay', 0, '2020-05-12 23:43:38', NULL),
(7, 5, 1, NULL, 'Pas de glacons', 'paid', '00:00:00', 'Gpay', 0, '2020-05-12 23:52:10', NULL),
(8, 5, 1, NULL, 'Pas de glacons', 'paid', '00:00:00', 'Gpay', 0, '2020-05-12 23:57:05', NULL),
(9, 5, 1, NULL, 'Pas de glacons', 'paid', '00:00:00', 'Gpay', 0, '2020-05-12 23:58:15', NULL),
(10, 5, 1, NULL, 'Pas de glacons', 'paid', '00:00:00', 'Gpay', 0, '2020-05-13 00:10:27', NULL),
(11, 5, 1, NULL, 'Pas de glacons', 'paid', '00:00:00', 'Gpay', 0, '2020-05-13 00:17:59', NULL),
(12, 5, 1, NULL, 'Pas de glacons', 'paid', '00:00:00', 'Gpay', 0, '2020-05-13 00:21:23', NULL),
(13, 5, 1, NULL, 'Pas de glacons', 'paid', '00:00:00', 'Gpay', 0, '2020-05-13 00:23:57', NULL),
(14, 5, 1, NULL, 'Pas de glacons', 'paid', '00:00:00', 'Gpay', 0, '2020-05-13 00:24:35', NULL),
(15, 5, 1, 25, 'Pas de glacons', 'paid', '00:00:00', 'Gpay', 0, '2020-05-13 00:31:52', NULL),
(16, 5, 1, 25, 'Pas de glacons', 'paid', '00:00:00', 'Gpay', 0, '2020-05-15 17:56:32', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `command_items`
--

DROP TABLE IF EXISTS `command_items`;
CREATE TABLE IF NOT EXISTS `command_items` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `command_id` int(11) NOT NULL,
  `item_id` int(11) NOT NULL,
  `price` float NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=16 DEFAULT CHARSET=latin1;

--
-- Dumping data for table `command_items`
--

INSERT INTO `command_items` (`id`, `command_id`, `item_id`, `price`) VALUES
(1, 1, 1, 5),
(2, 1, 1, 5),
(3, 1, 2, 10),
(4, 14, 1, 5),
(5, 14, 1, 5),
(6, 14, 2, 10),
(7, 14, 1, 5),
(8, 15, 1, 5),
(9, 15, 1, 5),
(10, 15, 2, 10),
(11, 15, 1, 5),
(12, 16, 1, 5),
(13, 16, 1, 5),
(14, 16, 2, 10),
(15, 16, 1, 5);

-- --------------------------------------------------------

--
-- Table structure for table `companies`
--

DROP TABLE IF EXISTS `companies`;
CREATE TABLE IF NOT EXISTS `companies` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(250) NOT NULL,
  `mail` varchar(250) NOT NULL,
  `phone_number` varchar(10) NOT NULL,
  `siret_siren` varchar(14) NOT NULL,
  `street_number` int(11) NOT NULL,
  `street_name` varchar(250) NOT NULL,
  `address_complement` varchar(250) DEFAULT NULL,
  `zip_code` varchar(5) NOT NULL,
  `city` varchar(250) NOT NULL,
  `country` varchar(100) NOT NULL,
  `iban` varchar(34) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;

--
-- Dumping data for table `companies`
--

INSERT INTO `companies` (`id`, `name`, `mail`, `phone_number`, `siret_siren`, `street_number`, `street_name`, `address_complement`, `zip_code`, `city`, `country`, `iban`) VALUES
(1, 'La maison du vin', 'lamaisonduvin_contact@gmail.com', '0102030405', 'fewt832hohdejj', 23, 'rue de la grande truanderie', '', '75011', 'Paris', 'France', 'hhhhjd7767889000');

-- --------------------------------------------------------

--
-- Table structure for table `etabs`
--

DROP TABLE IF EXISTS `etabs`;
CREATE TABLE IF NOT EXISTS `etabs` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `company_id` int(11) NOT NULL,
  `name` varchar(250) NOT NULL,
  `phone_number` varchar(10) NOT NULL,
  `mail` varchar(250) NOT NULL,
  `street_num` int(11) NOT NULL,
  `street_name` varchar(250) NOT NULL,
  `address_complement` varchar(250) NOT NULL,
  `city` varchar(250) NOT NULL,
  `zip` varchar(5) NOT NULL,
  `country` varchar(100) NOT NULL,
  `iban` varchar(34) DEFAULT NULL,
  `type` tinyint(1) NOT NULL,
  `subtype` varchar(250) NOT NULL,
  `description` text NOT NULL,
  `main_pic` varchar(250) NOT NULL,
  `longitude` float NOT NULL,
  `latitude` float NOT NULL,
  `date` timestamp NOT NULL DEFAULT current_timestamp(),
  `happy` time NOT NULL,
  `happy_end` time NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;

--
-- Dumping data for table `etabs`
--

INSERT INTO `etabs` (`id`, `company_id`, `name`, `phone_number`, `mail`, `street_num`, `street_name`, `address_complement`, `city`, `zip`, `country`, `iban`, `type`, `subtype`, `description`, `main_pic`, `longitude`, `latitude`, `date`, `happy`, `happy_end`) VALUES
(1, 1, 'La maison du vin - Chatelet', '0102030405', 'lamaisonduvinchatelet@gmail.com', 11, 'rue de la grande truanderie', '', 'Paris', '75006', 'France', 'hhhhjd7767889000', 1, 'Vin', 'Un bar de degustation de vin', '/home/path/to/pic/pic.jpeg', 2.35137, 48.829, '2020-05-09 02:46:19', '18:00:00', '00:00:00'),
(2, 1, 'Le choix du fermier', '0607080900', 'lechoixdufermier@yahoo.com', 23, 'rue de la pepiniere', '', 'Paris', '75017', 'France', 'hhhhjd7767889000', 1, 'Biere', 'un joli bar fermier avec du bon cidre', '/path/to/pic', 2.31245, 48.8188, '2020-05-09 02:48:47', '18:00:00', '22:00:00');

-- --------------------------------------------------------

--
-- Table structure for table `etab_pictures`
--

DROP TABLE IF EXISTS `etab_pictures`;
CREATE TABLE IF NOT EXISTS `etab_pictures` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `etab_id` int(11) NOT NULL,
  `path` varchar(250) NOT NULL,
  `description` varchar(250) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;

--
-- Dumping data for table `etab_pictures`
--

INSERT INTO `etab_pictures` (`id`, `etab_id`, `path`, `description`) VALUES
(1, 1, '/path/to/img/path/to/img', 'gneu'),
(2, 1, '/path/to/img/path/to/img', 'gneu');

-- --------------------------------------------------------

--
-- Table structure for table `favoris`
--

DROP TABLE IF EXISTS `favoris`;
CREATE TABLE IF NOT EXISTS `favoris` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_id` int(11) NOT NULL,
  `etab_id` int(11) NOT NULL,
  `date` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=14 DEFAULT CHARSET=latin1;

--
-- Dumping data for table `favoris`
--

INSERT INTO `favoris` (`id`, `user_id`, `etab_id`, `date`) VALUES
(1, 2, 1, '2020-05-08 21:09:40'),
(6, 5, 2, '2020-05-12 20:53:49');

-- --------------------------------------------------------

--
-- Table structure for table `items`
--

DROP TABLE IF EXISTS `items`;
CREATE TABLE IF NOT EXISTS `items` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(250) NOT NULL,
  `etab_id` int(11) NOT NULL,
  `type` varchar(250) NOT NULL,
  `description` varchar(250) NOT NULL,
  `price` float NOT NULL,
  `sale` float DEFAULT NULL,
  `picture_path` varchar(250) NOT NULL,
  `availlable` tinyint(1) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=3 DEFAULT CHARSET=latin1;

--
-- Dumping data for table `items`
--

INSERT INTO `items` (`id`, `name`, `etab_id`, `type`, `description`, `price`, `sale`, `picture_path`, `availlable`) VALUES
(1, 'Vin de table(10cl)', 1, 'Vin', 'Un vin de table comme on l\'attends', 5, NULL, '', 0),
(2, 'Chateau Montfort(10cl)', 1, 'Vin', 'Un vin subtil aux aromes delicats.', 10, NULL, '', 0);

-- --------------------------------------------------------

--
-- Table structure for table `preferences`
--

DROP TABLE IF EXISTS `preferences`;
CREATE TABLE IF NOT EXISTS `preferences` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `client_id` int(11) NOT NULL,
  `etab_id` int(11) NOT NULL,
  `date` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=latin1;

-- --------------------------------------------------------

--
-- Table structure for table `pros`
--

DROP TABLE IF EXISTS `pros`;
CREATE TABLE IF NOT EXISTS `pros` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(100) NOT NULL,
  `surname` varchar(100) NOT NULL,
  `mail` varchar(250) NOT NULL,
  `phone_number` varchar(10) NOT NULL,
  `password` text NOT NULL,
  `birth_date` date NOT NULL,
  `status` varchar(10) NOT NULL,
  `etab_id` int(11) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;

--
-- Dumping data for table `pros`
--

INSERT INTO `pros` (`id`, `name`, `surname`, `mail`, `phone_number`, `password`, `birth_date`, `status`, `etab_id`) VALUES
(1, 'Jean', 'Serlevin', 'jeanserlevin@live.com', '0607080900', '1234', '1977-11-20', 'admin', 1);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
