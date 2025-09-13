-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
create extension if not exists "uuid-ossp"; --an extension responsible for generating uuid
CREATE TABLE IF NOT EXISTS categories (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    parent_id  uuid, -- for sub-directories
    name VARCHAR(255) UNIQUE NOT NULL
);

-- Parents
INSERT INTO categories (name, parent_id) VALUES ('Vehicles', NULL);
INSERT INTO categories (name, parent_id) VALUES ('Electronics', NULL);
INSERT INTO categories (name, parent_id) VALUES ('Real Estate', NULL);
INSERT INTO categories (name, parent_id) VALUES ('Jobs & Services', NULL);
INSERT INTO categories (name, parent_id) VALUES ('Home & Furniture', NULL);
INSERT INTO categories (name, parent_id) VALUES ('Fashion & Accessories', NULL);
INSERT INTO categories (name, parent_id) VALUES ('Sports & Hobbies', NULL);
INSERT INTO categories (name, parent_id) VALUES ('Animals & Pets', NULL);

-- Vehicles children
INSERT INTO categories (name, parent_id) SELECT 'Cars', id FROM categories WHERE name='Vehicles';
INSERT INTO categories (name, parent_id) SELECT 'Motorcycles', id FROM categories WHERE name='Vehicles';
INSERT INTO categories (name, parent_id) SELECT 'Trucks & Vans', id FROM categories WHERE name='Vehicles';
INSERT INTO categories (name, parent_id) SELECT 'Boats & Watercraft', id FROM categories WHERE name='Vehicles';

-- Electronics children
INSERT INTO categories (name, parent_id) SELECT 'Mobile Phones', id FROM categories WHERE name='Electronics';
INSERT INTO categories (name, parent_id) SELECT 'Computers & Laptops', id FROM categories WHERE name='Electronics';
INSERT INTO categories (name, parent_id) SELECT 'Cameras & Photography', id FROM categories WHERE name='Electronics';
INSERT INTO categories (name, parent_id) SELECT 'TV & Audio', id FROM categories WHERE name='Electronics';

-- Real Estate children
INSERT INTO categories (name, parent_id) SELECT 'Apartments', id FROM categories WHERE name='Real Estate';
INSERT INTO categories (name, parent_id) SELECT 'Villas & Houses', id FROM categories WHERE name='Real Estate';
INSERT INTO categories (name, parent_id) SELECT 'Land & Plots', id FROM categories WHERE name='Real Estate';
INSERT INTO categories (name, parent_id) SELECT 'Commercial Properties', id FROM categories WHERE name='Real Estate';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
drop table if exists cities;
-- +goose StatementEnd
