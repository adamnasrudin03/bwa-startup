-- phpMyAdmin SQL Dump
-- version 5.0.4
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Mar 04, 2021 at 05:29 AM
-- Server version: 10.4.17-MariaDB
-- PHP Version: 7.4.13

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `startup_db`
--

-- --------------------------------------------------------

--
-- Table structure for table `campaigns`
--

CREATE TABLE `campaigns` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `short_description` varchar(255) NOT NULL,
  `description` text NOT NULL,
  `perks` text NOT NULL,
  `backer_count` int(11) NOT NULL,
  `goal_amount` int(11) NOT NULL,
  `current_amount` int(11) NOT NULL,
  `slug` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `campaigns`
--

INSERT INTO `campaigns` (`id`, `user_id`, `name`, `short_description`, `description`, `perks`, `backer_count`, `goal_amount`, `current_amount`, `slug`, `created_at`, `updated_at`) VALUES
(1, 2, 'Campaign punya Adam nasrudin', 'Cras faucibus magna vel blandit ultrices. Integer vitae aliquam nisl.', '            Cras faucibus magna vel blandit ultrices. Integer vitae aliquam nisl. Duis ullamcorper gravida tortor eget ornare. Aenean ante eros, dignissim nec dictum dapibus, auctor non elit. Fusce hendrerit est eu lectus gravida, eu posuere lectus lobortis. Quisque hendrerit odio et blandit consequat. Donec rutrum risus sit amet eros lacinia hendrerit. Curabitur congue dui sapien, et consectetur augue gravida at. Praesent erat tortor, venenatis at augue sit amet, ullamcorper mattis libero. Morbi feugiat pellentesque mauris et semper. Proin interdum tellus sit amet lorem tristique porttitor. Vivamus dapibus, eros vulputate euismod ultricies, tellus quam blandit enim, id porta augue augue et velit. Nam at efficitur est.\r\n\r\nDonec dapibus scelerisque accumsan. In hac habitasse platea dictumst. Proin et sapien sit amet ante vehicula vestibulum. Vivamus congue, tellus ut gravida consectetur, ligula felis tincidunt diam, sit amet placerat nulla nunc nec justo. Curabitur sit amet nibh at augue vestibulum accumsan vel eget massa. Cras eget sapien vel felis tempus eleifend non in eros. Curabitur id efficitur ex, eu volutpat sapien. Integer interdum sollicitudin velit, eget vehicula tortor fringilla sed. Nulla condimentum imperdiet massa ut interdum. Integer at ipsum sed odio lobortis dignissim. Morbi ac lectus suscipit, finibus ipsum sit amet, dapibus eros.\r\n        ', 'kecepatan yang ditawarkannya, tidak akan mengizinkan penggunanya untuk menghamburkan variabel atau data-data impor yang tidak digunakan, memiliki compiler yang dapat digunakan untuk mengkompilasi proyek dengan sangat cepat', 0, 200000, 0, 'campaign-punya-user-1', '2021-03-02 08:59:25', '2021-03-04 11:02:27'),
(2, 3, 'project name', 'project ini adalah', '                                    \r\n        PROJECT by fahriPROJECT by fahriPROJECT by fahriPROJECT by fahri PROJECT by fahri\r\n        \r\n        ', 'Keuntungan 1, Keuntungan 2, Keuntungan 3, Keuntungan 4', 1, 500000, 350000, 'project-by-fahri-16', '2021-03-02 09:44:21', '2021-03-04 11:27:24');

-- --------------------------------------------------------

--
-- Table structure for table `campaign_images`
--

CREATE TABLE `campaign_images` (
  `id` int(11) NOT NULL,
  `campaign_id` int(11) NOT NULL,
  `file_name` varchar(255) NOT NULL,
  `is_primary` tinyint(4) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `campaign_images`
--

INSERT INTO `campaign_images` (`id`, `campaign_id`, `file_name`, `is_primary`, `created_at`, `updated_at`) VALUES
(1, 1, 'images/campaign-1-undraw_profile_pic_ic5t.png', 1, '2021-03-02 11:41:42', '2021-03-02 11:41:42');

-- --------------------------------------------------------

--
-- Table structure for table `transactions`
--

CREATE TABLE `transactions` (
  `id` int(11) NOT NULL,
  `user_id` int(11) NOT NULL,
  `campaign_id` int(11) NOT NULL,
  `amount` int(11) NOT NULL,
  `status` varchar(255) NOT NULL,
  `code` varchar(255) NOT NULL,
  `payment_url` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `transactions`
--

INSERT INTO `transactions` (`id`, `user_id`, `campaign_id`, `amount`, `status`, `code`, `payment_url`, `created_at`, `updated_at`) VALUES
(15, 2, 1, 12345678, 'pending', '', 'https://app.sandbox.midtrans.com/snap/v2/vtweb/f1dd2953-c584-4b12-8e98-cf1c0b344329', '2021-03-02 13:56:19', '2021-03-02 13:56:21'),
(16, 2, 2, 200000, 'pending', '', 'https://app.sandbox.midtrans.com/snap/v2/vtweb/5bdafdb7-6721-4f14-a9b8-40e469b5b0cb', '2021-03-04 11:12:32', '2021-03-04 11:12:34'),
(18, 2, 2, 350000, 'paid', '', 'https://app.sandbox.midtrans.com/snap/v2/vtweb/4d433793-d2b1-4c6e-bba4-6c23b410e157', '2021-03-04 11:25:00', '2021-03-04 11:27:24');

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` int(11) NOT NULL,
  `name` varchar(255) NOT NULL,
  `occupation` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `password_hash` varchar(255) NOT NULL,
  `avatar_file_name` varchar(255) NOT NULL,
  `role` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `name`, `occupation`, `email`, `password_hash`, `avatar_file_name`, `role`, `created_at`, `updated_at`) VALUES
(1, 'Adam', 'Programmer', 'email@gmail.com', '$2a$04$QH.piCOpi0F6qBLixKk9weVlsqwXJXnzgF/31atA99TmM5j81iSHe', 'images/1-undraw_profile_pic_ic5t.png', 'admin', '2021-02-28 08:27:37', '2021-03-02 08:35:30'),
(2, 'Adam Nasrudin', 'Pengusaha', 'email2@gmail.com', '$2a$04$QH.piCOpi0F6qBLixKk9weVlsqwXJXnzgF/31atA99TmM5j81iSHe', '', 'user', '2021-02-28 08:27:37', '2021-03-02 08:35:10'),
(3, 'Fahri Maulana Azhari', 'Software Enginer', 'fahri@gmail.com', '$2a$04$Mq93Ix3D99YecS6IbOddReLhzbcxo/EZY92Rw/Xt1ul3jm8Ckhkm.', '', 'user', '2021-03-01 21:49:51', '2021-03-02 13:07:26');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `campaigns`
--
ALTER TABLE `campaigns`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `campaign_images`
--
ALTER TABLE `campaign_images`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `transactions`
--
ALTER TABLE `transactions`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `campaigns`
--
ALTER TABLE `campaigns`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=8;

--
-- AUTO_INCREMENT for table `campaign_images`
--
ALTER TABLE `campaign_images`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=10;

--
-- AUTO_INCREMENT for table `transactions`
--
ALTER TABLE `transactions`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=19;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=17;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
