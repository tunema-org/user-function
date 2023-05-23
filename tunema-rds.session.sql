CREATE TABLE IF NOT EXISTS users(
    id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    username varchar(50),
    email varchar(250),
    password varchar(250),
    profile_img_url text
);

CREATE TABLE IF NOT EXISTS wishlists(
    user_id int,
    sample_id int,
    PRIMARY KEY (user_id, sample_id)
);

CREATE TABLE IF NOT EXISTS samples(
    id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id int,
    name varchar(250),
    bpm int,
    instrument varchar(250),
    key varchar(250),
    length varchar(250),
    sample_file_url text,
    cover_url text,
    price text,
    type text,
    CONSTRAINT fk_samples_user_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS orders(
    id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id int,
    created_at timestamp(0) without time zone NOT NULL,
    CONSTRAINT fk_orders_user_id FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS order_products(
    order_id int,
    sample_id int,
    PRIMARY KEY (order_id, sample_id),
    CONSTRAINT fk_order_products_order_id FOREIGN KEY (order_id) REFERENCES orders(id) ON DELETE CASCADE,
    CONSTRAINT fk_order_products_sample_id FOREIGN KEY (sample_id) REFERENCES samples(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS categories(
    id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name varchar(250)
);

CREATE TABLE IF NOT EXISTS tags(
    id int GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    category_id int,
    name varchar(250),
    CONSTRAINT fk_tags_category_id FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS sample_tags(
    sample_id int,
    tag_id int,
    PRIMARY KEY (sample_id, tag_id),
    CONSTRAINT fk_sample_tags_sample_id FOREIGN KEY (sample_id) REFERENCES samples(id) ON DELETE CASCADE,
    CONSTRAINT fk_sample_tags_tag_id FOREIGN KEY (tag_id) REFERENCES tags(id) ON DELETE CASCADE
);

