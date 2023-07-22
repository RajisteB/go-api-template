CREATE TABLE IF NOT EXISTS accounts(
    id VARCHAR UNIQUE NOT NULL PRIMARY KEY,
    first_name TEXT,
    last_name TEXT,
    email VARCHAR UNIQUE NOT NULL,
    password VARCHAR NOT NULL,
    country TEXT,
    avatar_url TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS merchants(
    id VARCHAR UNIQUE NOT NULL PRIMARY KEY,
    merchant_name TEXT,
    email VARCHAR UNIQUE NOT NULL,
    password VARCHAR NOT NULL,
    description TEXT,
    signed_merchant_terms BOOLEAN NOT NULL,
    country TEXT,
    rating FLOAT,
    avatar_url TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);