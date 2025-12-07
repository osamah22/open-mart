-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
CREATE EXTENSION IF NOT EXISTS "uuid-ossp"; -- for UUIDs

CREATE TABLE IF NOT EXISTS categories (
    id uuid PRIMARY KEY DEFAULT uuid_generate_v4(),
    parent_id uuid REFERENCES categories(id),
    name VARCHAR(255) NOT NULL,
    slug VARCHAR(255) UNIQUE NOT NULL,
    logo_class VARCHAR(255) NOT NULL -- Font Awesome class
);

-- Parents
INSERT INTO categories (name, slug, parent_id, logo_class) 
VALUES 
    ('Vehicles', 'vehicles', NULL, 'fas fa-car'),
    ('Electronics', 'electronics', NULL, 'fas fa-tv'),
    ('Real Estate', 'real-estate', NULL, 'fas fa-home'),
    ('Jobs & Services', 'jobs-services', NULL, 'fas fa-briefcase'),
    ('Home & Furniture', 'home-furniture', NULL, 'fas fa-couch'),
    ('Fashion & Accessories', 'fashion-accessories', NULL, 'fas fa-tshirt'),
    ('Sports & Hobbies', 'sports-hobbies', NULL, 'fas fa-basketball-ball'),
    ('Animals & Pets', 'animals-pets', NULL, 'fas fa-dog');

-- Vehicles children
INSERT INTO categories (name, slug, parent_id, logo_class)
    SELECT 'Cars', 'cars', id, 'fas fa-car' FROM categories WHERE slug='vehicles';
INSERT INTO categories (name, slug, parent_id, logo_class)
    SELECT 'Motorcycles', 'motorcycles', id, 'fas fa-motorcycle' FROM categories WHERE slug='vehicles';
INSERT INTO categories (name, slug, parent_id, logo_class)
    SELECT 'Trucks & Vans', 'trucks-vans', id, 'fas fa-truck' FROM categories WHERE slug='vehicles';
INSERT INTO categories (name, slug, parent_id, logo_class)
    SELECT 'Boats & Watercraft', 'boats-watercraft', id, 'fas fa-ship' FROM categories WHERE slug='vehicles';

-- Electronics children
INSERT INTO categories (name, slug, parent_id, logo_class)
    SELECT 'Mobile Phones', 'mobile-phones', id, 'fas fa-mobile-alt' FROM categories WHERE slug='electronics';
INSERT INTO categories (name, slug, parent_id, logo_class)
    SELECT 'Computers & Laptops', 'computers-laptops', id, 'fas fa-laptop' FROM categories WHERE slug='electronics';
INSERT INTO categories (name, slug, parent_id, logo_class)
    SELECT 'Cameras & Photography', 'cameras-photography', id, 'fas fa-camera' FROM categories WHERE slug='electronics';
INSERT INTO categories (name, slug, parent_id, logo_class)
    SELECT 'TV & Audio', 'tv-audio', id, 'fas fa-tv' FROM categories WHERE slug='electronics';

-- Real Estate children
INSERT INTO categories (name, slug, parent_id, logo_class)
    SELECT 'Apartments', 'apartments', id, 'fas fa-building' FROM categories WHERE slug='real-estate';
INSERT INTO categories (name, slug, parent_id, logo_class)
    SELECT 'Villas & Houses', 'villas-houses', id, 'fas fa-home' FROM categories WHERE slug='real-estate';
INSERT INTO categories (name, slug, parent_id, logo_class)
    SELECT 'Land & Plots', 'land-plots', id, 'fas fa-tree' FROM categories WHERE slug='real-estate';
INSERT INTO categories (name, slug, parent_id, logo_class)
    SELECT 'Commercial Properties', 'commercial-properties', id, 'fas fa-store' FROM categories WHERE slug='real-estate';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
DROP TABLE IF EXISTS categories;
-- +goose StatementEnd
