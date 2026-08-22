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
