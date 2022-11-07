create table car (
    id serial primary key,
    brand varchar(50),
    year int,
    color varchar(50),
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp
)

create table car_image (
    id serial primary key,
    image_url varchar,
    car_id int
)