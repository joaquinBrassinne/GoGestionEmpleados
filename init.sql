-- Initialize database schema
CREATE TABLE IF NOT EXISTS employees (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL
);

-- Insert some sample data
INSERT INTO employees (name, email) VALUES
('John Doe', 'john@example.com'),
('Jane Smith', 'jane@example.com');