CREATE TABLE IF NOT EXISTS users(
    id SERIAL PRIMARY KEY,
    passport_serie INTEGER NOT NULL,
    passport_number INTEGER NOT NULL,
    surname VARCHAR(100) NOT NULL,
    name VARCHAR(100) NOT NULL,
    patronymic VARCHAR(100),
    address VARCHAR(255) NOT NULL,
    UNIQUE(passport_serie, passport_number)
);

CREATE TABLE IF NOT EXISTS tasks(
    id SERIAL PRIMARY KEY,
    user_id INTEGER NOT NULL,
    title VARCHAR(100) NOT NULL,
    description TEXT NOT NULL,
    start_at TIMESTAMP NOT NULL,
    end_at TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);