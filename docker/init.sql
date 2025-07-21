CREATE TABLE IF NOT EXISTS USERS (
    ID SERIAL PRIMARY KEY,
    EMAIL VARCHAR(64) NOT NULL UNIQUE,
    PASSWORD VARCHAR NOT NULL,
    ROLE VARCHAR(16) DEFAULT 'USER'
);

CREATE TABLE IF NOT EXISTS ADVERTISEMENT (
    ID SERIAL PRIMARY KEY,
    HEADER VARCHAR(64) NOT NULL,
    DESCRIPTION VARCHAR(512) NOT NULL,
    IMAGE_URL VARCHAR(256) NOT NULL,
    PRICE INT NOT NULL,
    CREATED_AT TIMESTAMP DEFAULT NOW(),
    OWNER_ID INT REFERENCES USERS(ID) NOT NULL
);

INSERT INTO USERS (EMAIL, PASSWORD, ROLE) VALUES
('test1@example.com', '$2a$14$QW8IuoILGZkU1WZcyu1IaOdPNZkgaZsYZoMgVCcKxRGm2zMoAZAVu', 'USER'), -- password: test123
('admin@example.com', '$2a$14$QW8IuoILGZkU1WZcyu1IaOdPNZkgaZsYZoMgVCcKxRGm2zMoAZAVu', 'ADMIN');

INSERT INTO ADVERTISEMENT (HEADER, DESCRIPTION, IMAGE_URL, PRICE, CREATED_AT, OWNER_ID) VALUES
('Lada 2112', 'В идеальном состоянии.', 'https://example.com/image1.jpg', 80000, NOW(), 1),
('Labubu', 'Почти новый, гарантия до конца года.', 'https://example.com/image2.jpg', 23000, NOW(), 1),
('M&M`s', 'Не вскрывался, новогодняя коллекция.', 'https://example.com/image3.jpg', 120, NOW(), 1),
('iPhone 16 Pro MAX', 'Смартфон Apple iPhone 16 Pro Max 256GB Natural Titanium (Nano+Nano)·', 'https://example.com/image4.jpg', 80000, NOW(), 1),
('Gopher', 'Талисман языка, философии дизайна Go', 'https://example.com/image7.jpg', 80000, NOW(), 1),
('Ноутбук ASUS', 'Для учебы и работы.', 'https://example.com/image5.jpg', 45000, NOW(), 2);
