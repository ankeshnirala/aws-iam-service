CREATE TABLE users (
    id INT(6) UNSIGNED AUTO_INCREMENT PRIMARY KEY,
    firstName VARCHAR(30) NOT NULL,
    lastName VARCHAR(30) NOT NULL,
    email VARCHAR(50),
	password VARCHAR(64) NOT NULL,
	role VARCHAR(50) NOT NULL,
    createdAt TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	CreatedBy VARCHAR(50)
);