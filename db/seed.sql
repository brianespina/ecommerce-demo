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
    created_at TIMESTAMP DEFAULT NOW()
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

