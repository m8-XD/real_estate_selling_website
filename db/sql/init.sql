-- sudo docker run --name postgres -e POSTGRES_PASSWORD=pass123 -d postgres

CREATE TABLE post (
    id SERIAL NOT NULL UNIQUE,
    header TEXT NOT NULL,
    description TEXT NOT NULL,
    price INT NOT NULL,
    name TEXT NOT NULL,
    phone TEXT NOT NULL,
    type INT NOT NULL,
    rooms TEXT NOT NULL,
    creator_id INT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP);

CREATE TABLE users (
    id SERIAL NOT NULL UNIQUE,
    name TEXT NOT NULL,
    mail TEXT NOT NULL UNIQUE,
    pass TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP);

CREATE TABLE types (
    id serial not null unique,
    name text not null unique,
    created_at TIMESTAMP,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP
);

INSERT INTO types (id, name) VALUES (0, 'Flat'), (1, 'Bungalow'), (2, 'Cottage'), (3, 'Other');
