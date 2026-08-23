create table movies (
	movie_id bigserial primary key,
	movie_name text,
	uploaded_by text,
	uploaded_on timestamp 
);

CREATE TABLE comments (
    comment_id BIGSERIAL PRIMARY KEY,
    movie_id BIGINT NOT NULL,
    comment TEXT NOT NULL,
    FOREIGN KEY (movie_id)
        REFERENCES movies(movie_id)
        ON DELETE CASCADE
);

--migration added sentiments 
alter table public."comments" 
add column sentiment varchar(20),
add message text ;


--added some other in movies schema  
select * from public.movies;
alter table public.movies
add column movie_link text,
add column image_url text, 
add column tarilor_link  text;

--rename column migration
select * from public.movies;

alter table public.movies 
rename column tarilor_link to new_column_name;

select * from public.movies;

--rename column name 
alter table public.movies 
rename column new_column_name to trailor_url;
