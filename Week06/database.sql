CREATE DATABASE dorm_parcel_db CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;

USE dorm_parcel_db;

CREATE TABLE parcels (
    id INT AUTO_INCREMENT PRIMARY KEY,
    tracking VARCHAR(50) NOT NULL,
    name VARCHAR(100) NOT NULL,
    status ENUM('รอรับ', 'รับแล้ว') DEFAULT 'รอรับ',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO parcels (tracking, name, status) VALUES 
('TH12345', 'ฟาติมะห์', 'รอรับ'),
('SP98765', 'พี่วัน', 'รับแล้ว');

