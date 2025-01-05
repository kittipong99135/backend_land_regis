-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Jan 05, 2025 at 05:32 AM
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
-- Database: `agent`
--

-- --------------------------------------------------------

--
-- Table structure for table `tbl_accounts`
--

CREATE TABLE `tbl_accounts` (
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `account_id` varchar(10) NOT NULL,
  `username'type:varchar(50)` longtext DEFAULT NULL,
  `email` varchar(50) DEFAULT NULL,
  `phone_number` varchar(10) DEFAULT NULL,
  `otp_code` varchar(10) DEFAULT NULL,
  `otp_expiry` datetime DEFAULT NULL,
  `azure_ad_id` varchar(50) DEFAULT NULL,
  `auth_type` varchar(50) DEFAULT NULL,
  `firstname` varchar(50) DEFAULT NULL,
  `lastname` varchar(50) DEFAULT NULL,
  `status` varchar(20) DEFAULT NULL,
  `role_office_id` varchar(20) DEFAULT NULL,
  `role_website_id` varchar(20) DEFAULT NULL,
  `password` longtext DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

--
-- Dumping data for table `tbl_accounts`
--

INSERT INTO `tbl_accounts` (`created_at`, `updated_at`, `deleted_at`, `account_id`, `username'type:varchar(50)`, `email`, `phone_number`, `otp_code`, `otp_expiry`, `azure_ad_id`, `auth_type`, `firstname`, `lastname`, `status`, `role_office_id`, `role_website_id`, `password`) VALUES
('2025-01-04 13:11:00', '2025-01-05 10:04:53', NULL, '0000001', 'ktpkst', 'user1@email.com', '0666666666', '052525', NULL, '[test]', '[test]', 'john', 'doe', '', 'ROLE_TEST', 'ROLE_ADMIN', '$2a$10$6yWit3AUmdpUPZFZBmAzouCTioyC/0uGSXfWu1wuvyXw4LrFzP7t.'),
('2025-01-04 13:22:50', '2025-01-04 13:22:50', NULL, '0000002', 'prmkst', 'user2@email.com', '0888888888', '052525', NULL, '[test]', '[test]', 'tim', 'koock', '', 'ROLE_GUST', 'ROLE_GUST', '$2a$10$csq4kOAY6P8PYr4OaaffKebXKJ4ATAVPm64XWOKpXC31uZAzdo/KW'),
('2025-01-04 13:23:48', '2025-01-04 13:23:48', NULL, '0000003', 'wtnkst', 'user3@email.com', '0999999999', '052525', NULL, '[test]', '[test]', 'nine', 'tail', '', 'ROLE_GUST', 'ROLE_GUST', '$2a$10$28XQfjBjEMoUaj6jHDk0beGKR4vKMZ0JqC75d.8KB/4vu06kRNUWu');

-- --------------------------------------------------------

--
-- Table structure for table `tbl_permissions`
--

CREATE TABLE `tbl_permissions` (
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `permission_id` varchar(20) NOT NULL,
  `permission_name` varchar(20) DEFAULT NULL,
  `module` varchar(20) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

--
-- Dumping data for table `tbl_permissions`
--

INSERT INTO `tbl_permissions` (`created_at`, `updated_at`, `deleted_at`, `permission_id`, `permission_name`, `module`) VALUES
('2025-01-04 18:52:06', '2025-01-04 18:52:06', NULL, 'PERM_AGENT', 'Agent', 'ข้อมูลนายหน้า'),
('2025-01-04 18:53:03', '2025-01-04 18:53:03', NULL, 'PERM_LAND', 'Land', 'ข้อมูลที่ดิน'),
('2025-01-04 18:53:32', '2025-01-04 18:53:32', NULL, 'PERM_LAYER', 'Layer', 'จัดการ Layer'),
('2025-01-04 18:54:31', '2025-01-04 18:54:31', NULL, 'PERM_SETTINGPERMISSI', 'SettingPermission', 'ตั้งค่าการเข้าถึง'),
('2025-01-05 09:50:17', '2025-01-05 09:50:17', NULL, 'PERM_TEST', 'Test', 'ทดสอบ'),
('2025-01-04 18:53:52', '2025-01-04 18:53:52', NULL, 'PERM_USER', 'User', 'จัดการผู้ใช้งาน');

-- --------------------------------------------------------

--
-- Table structure for table `tbl_roles`
--

CREATE TABLE `tbl_roles` (
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `role_id` varchar(20) NOT NULL,
  `role_name` varchar(20) DEFAULT NULL,
  `role_ref` varchar(20) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

--
-- Dumping data for table `tbl_roles`
--

INSERT INTO `tbl_roles` (`created_at`, `updated_at`, `deleted_at`, `role_id`, `role_name`, `role_ref`) VALUES
('2025-01-04 18:31:16', '2025-01-04 18:31:16', NULL, 'ROLE_ADMIN', 'Admin', 'office'),
('2025-01-04 18:34:11', '2025-01-04 18:34:11', NULL, 'ROLE_ANALYST', 'Analyst', 'website'),
('2025-01-04 18:32:05', '2025-01-04 18:32:05', NULL, 'ROLE_CHECKER', 'Checker', 'website'),
('2025-01-04 13:06:20', '2025-01-04 13:06:20', NULL, 'ROLE_GUST', 'Gust', 'website'),
('2025-01-04 18:31:31', '2025-01-04 18:31:31', NULL, 'ROLE_SUPERADMIN', 'SuperAdmin', 'office'),
('2025-01-05 09:32:59', '2025-01-05 09:32:59', NULL, 'ROLE_TEST', 'test', 'website'),
('2025-01-05 10:06:38', '2025-01-05 10:06:38', NULL, 'ROLE_TEST2', 'TEST2', 'website');

-- --------------------------------------------------------

--
-- Table structure for table `tbl_role_permission`
--

CREATE TABLE `tbl_role_permission` (
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL,
  `role_permission_id` varchar(10) NOT NULL,
  `role_id` varchar(20) DEFAULT NULL,
  `permission_id` varchar(20) DEFAULT NULL,
  `permissions` longtext CHARACTER SET utf8mb4 COLLATE utf8mb4_bin DEFAULT NULL CHECK (json_valid(`permissions`))
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_general_ci;

--
-- Dumping data for table `tbl_role_permission`
--

INSERT INTO `tbl_role_permission` (`created_at`, `updated_at`, `deleted_at`, `role_permission_id`, `role_id`, `permission_id`, `permissions`) VALUES
('2025-01-04 20:44:08', '2025-01-05 10:09:07', NULL, 'RP00001', 'ROLE_ADMIN', 'PERM_AGENT', '{\"create\":false,\"view\":true,\"edit\":true,\"delete\":true}'),
('2025-01-04 20:44:08', '2025-01-05 10:09:07', NULL, 'RP00002', 'ROLE_ADMIN', 'PERM_LAND', '{\"create\":true,\"view\":true,\"edit\":true,\"delete\":true}'),
('2025-01-04 20:44:08', '2025-01-05 10:09:07', NULL, 'RP00003', 'ROLE_ADMIN', 'PERM_LAYER', '{\"create\":true,\"view\":true,\"edit\":true,\"delete\":true}'),
('2025-01-04 20:44:08', '2025-01-05 10:09:07', NULL, 'RP00004', 'ROLE_ADMIN', 'PERM_SETTINGPERMISSI', '{\"create\":true,\"view\":true,\"edit\":true,\"delete\":true}'),
('2025-01-04 20:44:08', '2025-01-05 10:09:07', NULL, 'RP00005', 'ROLE_ADMIN', 'PERM_USER', '{\"create\":true,\"view\":true,\"edit\":true,\"delete\":true}'),
('2025-01-05 09:45:48', '2025-01-05 11:02:39', NULL, 'RP00006', 'ROLE_TEST', 'PERM_AGENT', '{\"create\":false,\"view\":false,\"edit\":false,\"delete\":false}'),
('2025-01-05 09:45:48', '2025-01-05 11:02:39', NULL, 'RP00007', 'ROLE_TEST', 'PERM_LAND', '{\"create\":true,\"view\":true,\"edit\":true,\"delete\":true}'),
('2025-01-05 09:45:48', '2025-01-05 10:00:17', NULL, 'RP00008', 'ROLE_TEST', 'PERM_LAYER', '{\"create\":false,\"view\":false,\"edit\":false,\"delete\":true}'),
('2025-01-05 09:45:48', '2025-01-05 10:00:17', NULL, 'RP00009', 'ROLE_TEST', 'PERM_SETTINGPERMISSI', '{\"create\":true,\"view\":true,\"edit\":true,\"delete\":true}'),
('2025-01-05 09:45:48', '2025-01-05 10:00:17', NULL, 'RP00010', 'ROLE_TEST', 'PERM_USER', '{\"create\":false,\"view\":false,\"edit\":false,\"delete\":true}');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `tbl_accounts`
--
ALTER TABLE `tbl_accounts`
  ADD PRIMARY KEY (`account_id`),
  ADD UNIQUE KEY `idx_phone_number` (`phone_number`),
  ADD UNIQUE KEY `idx_email` (`email`),
  ADD UNIQUE KEY `idx_username` (`username'type:varchar(50)`) USING HASH,
  ADD KEY `idx_tbl_accounts_deleted_at` (`deleted_at`),
  ADD KEY `fk_tbl_accounts_role_office` (`role_office_id`),
  ADD KEY `fk_tbl_accounts_role_website` (`role_website_id`);

--
-- Indexes for table `tbl_permissions`
--
ALTER TABLE `tbl_permissions`
  ADD PRIMARY KEY (`permission_id`),
  ADD UNIQUE KEY `idx_permission_name` (`permission_name`),
  ADD UNIQUE KEY `uni_tbl_permissions_permission_name` (`permission_name`),
  ADD KEY `idx_tbl_permissions_deleted_at` (`deleted_at`);

--
-- Indexes for table `tbl_roles`
--
ALTER TABLE `tbl_roles`
  ADD PRIMARY KEY (`role_id`),
  ADD UNIQUE KEY `idx_role_name` (`role_name`),
  ADD UNIQUE KEY `uni_tbl_roles_role_name` (`role_name`),
  ADD KEY `idx_tbl_roles_deleted_at` (`deleted_at`);

--
-- Indexes for table `tbl_role_permission`
--
ALTER TABLE `tbl_role_permission`
  ADD PRIMARY KEY (`role_permission_id`),
  ADD UNIQUE KEY `idx_role_permission` (`role_id`,`permission_id`),
  ADD KEY `idx_tbl_role_permission_deleted_at` (`deleted_at`),
  ADD KEY `fk_tbl_role_permission_permission_data` (`permission_id`);

--
-- Constraints for dumped tables
--

--
-- Constraints for table `tbl_accounts`
--
ALTER TABLE `tbl_accounts`
  ADD CONSTRAINT `fk_tbl_accounts_role_office` FOREIGN KEY (`role_office_id`) REFERENCES `tbl_roles` (`role_id`),
  ADD CONSTRAINT `fk_tbl_accounts_role_website` FOREIGN KEY (`role_website_id`) REFERENCES `tbl_roles` (`role_id`);

--
-- Constraints for table `tbl_role_permission`
--
ALTER TABLE `tbl_role_permission`
  ADD CONSTRAINT `fk_tbl_role_permission_permission_data` FOREIGN KEY (`permission_id`) REFERENCES `tbl_permissions` (`permission_id`),
  ADD CONSTRAINT `fk_tbl_role_permission_role_data` FOREIGN KEY (`role_id`) REFERENCES `tbl_roles` (`role_id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
