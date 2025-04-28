-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Apr 28, 2025 at 05:53 PM
-- Server version: 10.4.32-MariaDB
-- PHP Version: 8.2.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `test_fullstack`
--

-- --------------------------------------------------------

--
-- Table structure for table `motorcycles`
--

CREATE TABLE `motorcycles` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `nama_motor` longtext DEFAULT NULL,
  `jenis_motor` longtext DEFAULT NULL,
  `nomor_plat_motor` longtext DEFAULT NULL,
  `qty_motor` int(10) UNSIGNED DEFAULT NULL,
  `harga_sewa_motor` int(10) UNSIGNED DEFAULT NULL,
  `image_motor` varchar(255) DEFAULT NULL,
  `tanggal_pinjam` date DEFAULT NULL,
  `tanggal_kembali` date DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `motorcycles`
--

INSERT INTO `motorcycles` (`id`, `created_at`, `updated_at`, `deleted_at`, `nama_motor`, `jenis_motor`, `nomor_plat_motor`, `qty_motor`, `harga_sewa_motor`, `image_motor`, `tanggal_pinjam`, `tanggal_kembali`) VALUES
(1, '2025-04-27 21:50:05.012', '2025-04-27 21:50:05.012', '2025-04-27 23:25:01.286', 'coba 4', 'coba 4', 'coba 4', 1000, 1000, 'coba 1', NULL, NULL),
(2, '2025-04-27 23:21:03.678', '2025-04-28 08:22:42.935', NULL, 'motor1', 'vario 160', 'kdnaslkdjslk', 10000, 10000, 'coba 1', '2025-04-27', '2025-04-29'),
(3, '2025-04-28 08:09:24.070', '2025-04-28 08:09:24.070', '2025-04-28 08:16:15.622', 'coba 2', 'coba 2', 'sadsadsad', 2000, 2000, 'coba 2', '0000-00-00', '0000-00-00'),
(4, '2025-04-28 09:43:21.155', '2025-04-28 09:43:21.155', '2025-04-28 09:52:28.497', 'COBA 4', 'COBA 4', 'COBA 4', 10, 1, 'uploads/ChatGPT Image Apr 5, 2025, 05_02_17 AM.png', '0000-00-00', '0000-00-00'),
(5, '2025-04-28 10:02:39.881', '2025-04-28 10:02:39.881', '2025-04-28 10:02:54.766', '', '', '', 0, 0, 'uploads/Chester 2.png', '0000-00-00', '0000-00-00');

-- --------------------------------------------------------

--
-- Table structure for table `permissions`
--

CREATE TABLE `permissions` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` longtext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `permissions`
--

INSERT INTO `permissions` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`) VALUES
(1, '2025-04-27 21:49:46.303', '2025-04-27 21:49:46.303', NULL, 'view_data'),
(2, '2025-04-27 21:49:46.305', '2025-04-27 21:49:46.305', NULL, 'edit_data');

-- --------------------------------------------------------

--
-- Table structure for table `pesan_motors`
--

CREATE TABLE `pesan_motors` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `user_id` bigint(20) UNSIGNED DEFAULT NULL,
  `motorcycle_id` bigint(20) UNSIGNED DEFAULT NULL,
  `tanggal_pinjam` datetime(3) DEFAULT NULL,
  `tanggal_kembali` datetime(3) DEFAULT NULL,
  `total_harga_sewa` int(10) UNSIGNED DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `pesan_motors`
--

INSERT INTO `pesan_motors` (`id`, `created_at`, `updated_at`, `deleted_at`, `user_id`, `motorcycle_id`, `tanggal_pinjam`, `tanggal_kembali`, `total_harga_sewa`) VALUES
(1, '2025-04-28 10:20:51.510', '2025-04-28 10:20:51.510', NULL, 2, 2, '2025-04-27 00:00:00.000', '2025-04-29 00:00:00.000', 20000),
(2, '2025-04-28 10:53:27.392', '2025-04-28 10:53:27.392', NULL, 2, 2, '2025-04-29 00:00:00.000', '2025-04-30 00:00:00.000', 10000),
(3, '2025-04-28 15:05:58.436', '2025-04-28 15:05:58.436', NULL, 2, 2, '2025-04-28 00:00:00.000', '2025-05-02 00:00:00.000', 40000),
(4, '2025-04-28 15:10:10.303', '2025-04-28 15:10:10.303', NULL, 2, 2, '2025-04-28 00:00:00.000', '2025-05-02 00:00:00.000', 40000),
(5, '2025-04-28 15:18:50.259', '2025-04-28 15:18:50.259', NULL, 2, 2, '2025-04-28 00:00:00.000', '2025-04-30 00:00:00.000', 20000),
(6, '2025-04-28 15:18:58.379', '2025-04-28 15:18:58.379', NULL, 2, 2, '2025-04-28 00:00:00.000', '2025-04-30 00:00:00.000', 20000),
(7, '2025-04-28 15:28:16.763', '2025-04-28 15:28:16.763', NULL, 2, 2, '2025-04-30 00:00:00.000', '2025-05-02 00:00:00.000', 20000),
(8, '2025-04-28 15:28:29.168', '2025-04-28 15:28:29.168', NULL, 2, 2, '2025-04-30 00:00:00.000', '2025-05-02 00:00:00.000', 20000),
(9, '2025-04-28 15:32:53.929', '2025-04-28 15:32:53.929', NULL, 2, 2, '2025-04-29 00:00:00.000', '2025-05-02 00:00:00.000', 30000),
(10, '2025-04-28 15:35:43.282', '2025-04-28 15:35:43.282', NULL, 2, 2, '2025-04-28 00:00:00.000', '2025-04-30 00:00:00.000', 20000),
(11, '2025-04-28 15:36:00.378', '2025-04-28 15:36:00.378', NULL, 2, 2, '2025-04-28 00:00:00.000', '2025-04-30 00:00:00.000', 20000),
(12, '2025-04-28 15:42:58.693', '2025-04-28 15:42:58.693', NULL, 2, 2, '2025-04-28 00:00:00.000', '2025-04-30 00:00:00.000', 20000),
(13, '2025-04-28 15:43:00.830', '2025-04-28 15:43:00.830', NULL, 2, 2, '2025-04-28 00:00:00.000', '2025-04-30 00:00:00.000', 20000),
(14, '2025-04-28 15:48:06.026', '2025-04-28 15:48:06.026', NULL, 2, 2, '2025-04-27 00:00:00.000', '2025-04-29 00:00:00.000', 20000),
(15, '2025-04-28 15:51:08.543', '2025-04-28 15:51:08.543', NULL, 2, 2, '2025-04-28 00:00:00.000', '2025-05-08 00:00:00.000', 100000);

-- --------------------------------------------------------

--
-- Table structure for table `roles`
--

CREATE TABLE `roles` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `name` longtext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `roles`
--

INSERT INTO `roles` (`id`, `created_at`, `updated_at`, `deleted_at`, `name`) VALUES
(1, '2025-04-27 21:49:46.306', '2025-04-27 21:49:46.306', NULL, 'admin'),
(2, '2025-04-27 21:49:46.310', '2025-04-27 21:49:46.310', NULL, 'user');

-- --------------------------------------------------------

--
-- Table structure for table `role_permissions`
--

CREATE TABLE `role_permissions` (
  `role_id` bigint(20) UNSIGNED NOT NULL,
  `permission_id` bigint(20) UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `role_permissions`
--

INSERT INTO `role_permissions` (`role_id`, `permission_id`) VALUES
(1, 1),
(1, 2),
(2, 1);

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `nama_lengkap` longtext DEFAULT NULL,
  `no_telp` longtext DEFAULT NULL,
  `nik` longtext DEFAULT NULL,
  `email` longtext DEFAULT NULL,
  `image` longtext DEFAULT NULL,
  `black_list` longtext DEFAULT NULL,
  `password` longtext DEFAULT NULL,
  `role_id` bigint(20) UNSIGNED DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `created_at`, `updated_at`, `deleted_at`, `nama_lengkap`, `no_telp`, `nik`, `email`, `image`, `black_list`, `password`, `role_id`) VALUES
(1, '2025-04-27 21:49:46.316', '2025-04-27 21:49:46.316', NULL, '', '', '', 'stefanus@example.com', '', '', '$2a$04$s3.VBqkgGEiRAoCZtlgopOQqqZHFIFjn4EjuSXxMEj47xQttSk2Zy', 1),
(2, '2025-04-27 21:49:46.318', '2025-04-27 21:49:46.318', NULL, '', '', '', 'budi@example.com', '', '', '$2a$04$r9WLXk1ZjCYk7QH3ZfPbJO3dsZtHtKcFZ3Jv5QjB9KKWaLLB13cCK', 2),
(3, '2025-04-27 22:22:32.735', '2025-04-27 22:22:32.735', NULL, 'coba 1', '', '', 'coba@gmail.com', '', '', '$2a$10$aSoJA//yQEaN320AhT6zTOVFY09COauypRZm58zMNiF5vSSQSn7Ce', 2);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `motorcycles`
--
ALTER TABLE `motorcycles`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_motorcycles_deleted_at` (`deleted_at`);

--
-- Indexes for table `permissions`
--
ALTER TABLE `permissions`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_permissions_deleted_at` (`deleted_at`);

--
-- Indexes for table `pesan_motors`
--
ALTER TABLE `pesan_motors`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_pesan_motors_deleted_at` (`deleted_at`),
  ADD KEY `fk_pesan_motors_user` (`user_id`),
  ADD KEY `fk_pesan_motors_motorcycle` (`motorcycle_id`);

--
-- Indexes for table `roles`
--
ALTER TABLE `roles`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_roles_deleted_at` (`deleted_at`);

--
-- Indexes for table `role_permissions`
--
ALTER TABLE `role_permissions`
  ADD PRIMARY KEY (`role_id`,`permission_id`),
  ADD KEY `fk_role_permissions_permission` (`permission_id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`),
  ADD KEY `idx_users_deleted_at` (`deleted_at`),
  ADD KEY `fk_users_role` (`role_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `motorcycles`
--
ALTER TABLE `motorcycles`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;

--
-- AUTO_INCREMENT for table `permissions`
--
ALTER TABLE `permissions`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `pesan_motors`
--
ALTER TABLE `pesan_motors`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=16;

--
-- AUTO_INCREMENT for table `roles`
--
ALTER TABLE `roles`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `pesan_motors`
--
ALTER TABLE `pesan_motors`
  ADD CONSTRAINT `fk_pesan_motors_motorcycle` FOREIGN KEY (`motorcycle_id`) REFERENCES `motorcycles` (`id`),
  ADD CONSTRAINT `fk_pesan_motors_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

--
-- Constraints for table `role_permissions`
--
ALTER TABLE `role_permissions`
  ADD CONSTRAINT `fk_role_permissions_permission` FOREIGN KEY (`permission_id`) REFERENCES `permissions` (`id`),
  ADD CONSTRAINT `fk_role_permissions_role` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`);

--
-- Constraints for table `users`
--
ALTER TABLE `users`
  ADD CONSTRAINT `fk_users_role` FOREIGN KEY (`role_id`) REFERENCES `roles` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
