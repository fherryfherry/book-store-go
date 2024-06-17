-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Waktu pembuatan: 17 Jun 2024 pada 15.54
-- Versi server: 8.3.0
-- Versi PHP: 8.3.4

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `book_store`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `books`
--

CREATE TABLE `books` (
  `id` int NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `code` varchar(100) NOT NULL,
  `title` varchar(255) NOT NULL,
  `description` varchar(255) DEFAULT NULL,
  `price` double DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data untuk tabel `books`
--

INSERT INTO `books` (`id`, `created_at`, `updated_at`, `deleted_at`, `code`, `title`, `description`, `price`) VALUES
(1, '2024-06-16 17:29:55', NULL, NULL, 'B001', 'The Lost Tides of Aetheria', 'In a world where the ocean holds secrets of ancient civilizations, a young cartographer discovers a map leading to a hidden realm beneath the waves. As he embarks on a perilous journey, he must navigate treacherous waters and confront mystical creatures t', 100),
(2, '2024-06-16 17:29:55', NULL, NULL, 'B002', 'Echoes of the Forgotten Realm', 'When an archaeologist unearths a mysterious artifact in the heart of an ancient city, she awakens dormant powers that can reshape reality. With dark forces seeking to control the artifact, she must team up with an unlikely ally to prevent the world from f', 200);

-- --------------------------------------------------------

--
-- Struktur dari tabel `customers`
--

CREATE TABLE `customers` (
  `id` int NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `first_name` varchar(55) DEFAULT NULL,
  `last_name` varchar(55) DEFAULT NULL,
  `email` varchar(150) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data untuk tabel `customers`
--

INSERT INTO `customers` (`id`, `created_at`, `updated_at`, `deleted_at`, `first_name`, `last_name`, `email`, `password`) VALUES
(1, '2024-06-16 16:24:02', NULL, NULL, 'Ferry', 'Ariawan', 'ferdevelop15@gmail.com', '$2a$14$64v8Mi3UbCezVOQZbjSAWulX4eBx7Y0Yko.7MtJD8OsjOhvELSrwC');

-- --------------------------------------------------------

--
-- Struktur dari tabel `orders`
--

CREATE TABLE `orders` (
  `id` int NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL,
  `order_no` varchar(255) NOT NULL,
  `customer_id` int NOT NULL,
  `customer_name` varchar(255) NOT NULL,
  `customer_email` varchar(255) NOT NULL,
  `shipment_address` varchar(255) DEFAULT NULL,
  `grand_total` double NOT NULL,
  `order_status` varchar(55) NOT NULL DEFAULT 'PENDING',
  `payment_status` varchar(55) NOT NULL DEFAULT 'UNPAID'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data untuk tabel `orders`
--

INSERT INTO `orders` (`id`, `created_at`, `updated_at`, `deleted_at`, `order_no`, `customer_id`, `customer_name`, `customer_email`, `shipment_address`, `grand_total`, `order_status`, `payment_status`) VALUES
(2, '2024-06-17 00:45:48', NULL, NULL, 'BS001', 1, 'Ferry', '', '', 0, 'COMPLETED', 'PAID'),
(3, '2024-06-17 00:45:56', NULL, NULL, 'BS002', 1, 'Ferry', '', '', 0, 'COMPLETED', 'PAID'),
(4, '2024-06-17 00:46:29', NULL, NULL, 'BS003', 1, 'Ferry', '', '', 0, 'COMPLETED', 'PAID'),
(7, '2024-06-17 00:48:43', NULL, NULL, 'BS004', 1, 'Ferry', '', '', 0, 'COMPLETED', 'PAID'),
(8, '2024-06-17 00:53:02', NULL, NULL, 'BS005', 1, 'Ferry', '', 'lorem ipsum dolor sit amet', 300, 'COMPLETED', 'PAID'),
(9, '2024-06-17 14:40:16', NULL, NULL, 'BS006', 1, 'Ferry', 'ferdevelop15@gmail.com', 'lorem ipsum dolor sit amet', 300, 'COMPLETED', 'PAID');

-- --------------------------------------------------------

--
-- Struktur dari tabel `order_details`
--

CREATE TABLE `order_details` (
  `id` int NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL,
  `order_id` int NOT NULL,
  `book_code` varchar(255) NOT NULL,
  `book_title` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL,
  `book_price` double NOT NULL,
  `qty` int DEFAULT NULL,
  `total_price` double NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

--
-- Dumping data untuk tabel `order_details`
--

INSERT INTO `order_details` (`id`, `created_at`, `updated_at`, `order_id`, `book_code`, `book_title`, `book_price`, `qty`, `total_price`) VALUES
(1, '2024-06-17 00:45:48', '0000-00-00 00:00:00', 2, 'B001', 'The Lost Tides of Aetheria', 100, 1, 100),
(2, '2024-06-17 00:45:48', '0000-00-00 00:00:00', 2, 'B002', 'The Lost Tides of Aetheria', 100, 1, 100),
(3, '2024-06-17 00:45:56', '0000-00-00 00:00:00', 3, 'B001', 'The Lost Tides of Aetheria', 100, 1, 100),
(4, '2024-06-17 00:45:56', '0000-00-00 00:00:00', 3, 'B002', 'The Lost Tides of Aetheria', 100, 1, 100),
(5, '2024-06-17 00:46:28', '0000-00-00 00:00:00', 4, '', 'The Lost Tides of Aetheria', 100, 1, 100),
(6, '2024-06-17 00:46:28', '0000-00-00 00:00:00', 4, 'B002', 'The Lost Tides of Aetheria', 100, 1, 100),
(7, '2024-06-17 00:48:43', '0000-00-00 00:00:00', 7, 'B001', 'The Lost Tides of Aetheria', 100, 1, 100),
(8, '2024-06-17 00:48:43', '0000-00-00 00:00:00', 7, 'B002', 'Echoes of the Forgotten Realm', 200, 1, 200),
(9, '2024-06-17 00:53:01', '0000-00-00 00:00:00', 8, 'B001', 'The Lost Tides of Aetheria', 100, 1, 100),
(10, '2024-06-17 00:53:01', '0000-00-00 00:00:00', 8, 'B002', 'Echoes of the Forgotten Realm', 200, 1, 200),
(11, '2024-06-17 14:40:16', '0000-00-00 00:00:00', 9, 'B001', 'The Lost Tides of Aetheria', 100, 1, 100),
(12, '2024-06-17 14:40:16', '0000-00-00 00:00:00', 9, 'B002', 'Echoes of the Forgotten Realm', 200, 1, 200);

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `books`
--
ALTER TABLE `books`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `code` (`code`);

--
-- Indeks untuk tabel `customers`
--
ALTER TABLE `customers`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `email` (`email`);

--
-- Indeks untuk tabel `orders`
--
ALTER TABLE `orders`
  ADD PRIMARY KEY (`id`);

--
-- Indeks untuk tabel `order_details`
--
ALTER TABLE `order_details`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `books`
--
ALTER TABLE `books`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT untuk tabel `customers`
--
ALTER TABLE `customers`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=7;

--
-- AUTO_INCREMENT untuk tabel `orders`
--
ALTER TABLE `orders`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=10;

--
-- AUTO_INCREMENT untuk tabel `order_details`
--
ALTER TABLE `order_details`
  MODIFY `id` int NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=13;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
