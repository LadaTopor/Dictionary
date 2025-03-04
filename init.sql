CREATE TABLE ru_en (
                       id SERIAL PRIMARY KEY,
                       title VARCHAR(50) UNIQUE,
                       translation VARCHAR(255)
);

CREATE TABLE reports (
                       id SERIAL PRIMARY KEY,
                       title VARCHAR(100) UNIQUE,
                       description VARCHAR(255),
                       created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                       updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE OR REPLACE FUNCTION update_timestamp() RETURNS TRIGGER AS
$$
BEGIN
    NEW.updated_at := CURRENT_TIMESTAMP;
    RETURN NEW;
END;
$$
LANGUAGE plpgsql;

CREATE TRIGGER update_reports_timestamp
    BEFORE UPDATE ON reports
    FOR EACH ROW
    EXECUTE FUNCTION update_timestamp();

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

INSERT INTO reports (title, description) VALUES
                                           ('Не по агъдау', 'Админ крутит пироги при разрерании'),
                                           ('test2', 'desctest2');
UPDATE reports
    SET title = 'зисимирули',
        description = 'they see me rollin'
WHERE id = 2;

CREATE EXTENSION IF NOT EXISTS pg_trgm;