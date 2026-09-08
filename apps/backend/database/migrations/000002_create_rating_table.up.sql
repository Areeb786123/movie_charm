create table rating(
id bigserial primary key,
movie_id bigint not null,
rating int not null,
created_at timestamp default current_timestamp,
update_at timestamp default current_timestamp,
foreign key(movie_id) references movies(movie_id)
);