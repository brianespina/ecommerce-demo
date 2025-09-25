-- Enable pgcrypto for gen_random_uuid()
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- Create users table 
CREATE TABLE IF NOT EXISTS users (
    id SERIAL PRIMARY KEY,
    email VARCHAR(255) UNIQUE NOT NULL,
    password_hash TEXT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);

-- Create profiles table
CREATE TABLE IF NOT EXISTS profiles (
    id SERIAL PRIMARY KEY,
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    name VARCHAR(100),
    address TEXT,
    phone VARCHAR(20),
    avatar_url TEXT
);

-- Create sessions table
CREATE TABLE IF NOT EXISTS sessions (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    user_id INT REFERENCES users(id) ON DELETE CASCADE,
    session_token TEXT UNIQUE NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now(),
    expires_at TIMESTAMP WITH TIME ZONE NOT NULL
);

-- Index to quickly clean up expired sessions
CREATE INDEX IF NOT EXISTS idx_sessions_expires_at ON sessions(expires_at);

-- Create categories table
CREATE TABLE IF NOT EXISTS categories (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

-- Create products table
CREATE TABLE IF NOT EXISTS products (
    id SERIAL PRIMARY KEY,
    category_id INT REFERENCES categories(id) ON DELETE SET NULL,
    name TEXT NOT NULL,
    description TEXT,
    price NUMERIC(10,2) NOT NULL,
    stock INT NOT NULL DEFAULT 0,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT now()
);

-- Insert categories
INSERT INTO categories (name) VALUES
('Electronics'),
('Accessories'),
('Audio'),
('Monitors')
ON CONFLICT DO NOTHING;

-- Insert products with categories
INSERT INTO products (category_id, name, description, price, stock) VALUES
(1, 'Gaming Laptop', 'High-performance laptop with RTX GPU', 1299.99, 10),
(2, 'Mechanical Keyboard', 'RGB backlit mechanical keyboard', 89.99, 25),
(2, 'Wireless Mouse', 'Ergonomic wireless mouse', 39.99, 50),
(3, 'Noise Cancelling Headphones', 'Over-ear headphones with ANC', 199.99, 15),
(4, '4K Monitor', '27-inch 4K UHD IPS display', 349.99, 8);

