CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    age VARCHAR(255) UNIQUE NOT NULL,
    birthday DATE NOT NULL,
    profile_picture_id VARCHAR(255) NOT NULL,
    title VARCHAR(255) NOT NULL,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);