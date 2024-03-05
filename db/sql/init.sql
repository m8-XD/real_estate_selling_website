-- sudo docker run --name postgres -e POSTGRES_PASSWORD=pass123 -d postgres

CREATE TABLE post (
    id SERIAL NOT NULL UNIQUE,
    header TEXT NOT NULL,
    description TEXT NOT NULL,
    owner_name TEXT NOT NULL,
    owner_phone TEXT NOT NULL,
    creator_id UUID NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP);

CREATE TABLE users (
    id SERIAL NOT NULL UNIQUE,
    name TEXT NOT NULL,
    mail TEXT NOT NULL,
    pass TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP,
    deleted_at TIMESTAMP);
