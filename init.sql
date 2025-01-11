CREATE TABLE ru_en (
                       id SERIAL PRIMARY KEY,
                       title VARCHAR(50) UNIQUE,
                       translation VARCHAR(255)
);

INSERT INTO ru_en (title, translation) VALUES
                                           ('Привет', 'Hello'),
                                           ('Мир', 'World'),
                                           ('Книга', 'Book'),
                                           ('Стол', 'Table'),
                                           ('Яблоко', 'Apple'),
                                           ('Солнце', 'Sun'),
                                           ('Вода', 'Water'),
                                           ('Дом', 'House'),
                                           ('Кот', 'Cat'),
                                           ('Собака', 'Dog'),
                                           ('Человек', 'Human'),
                                           ('Школа', 'School'),
                                           ('Машина', 'Car'),
                                           ('Окно', 'Window'),
                                           ('Ручка', 'Pen');

CREATE TABLE reports (
                       id SERIAL PRIMARY KEY,
                       title VARCHAR(50),
                       description TEXT,
                       created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP

);

CREATE EXTENSION IF NOT EXISTS pg_trgm;
