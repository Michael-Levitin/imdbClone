DROP TABLE IF EXISTS parts;
DROP TABLE IF EXISTS movies;
DROP TABLE IF EXISTS actors;

CREATE TABLE movies
(
    id          serial primary key,
    movie       varchar(150) not null UNIQUE,
    description varchar(1000),
    release     timestamp,
    rating      float,
    removed     bool,
    created_at  timestamp
);

comment on column movies.id is 'id записи';
comment on column movies.movie is 'название фильма';
comment on column movies.description is 'описание';
comment on column movies.release is 'дата выпуска';
comment on column movies.rating is 'рейтинг';
comment on column movies.removed is 'был удален';
comment on column movies.created_at is 'дата создания записи';


CREATE TABLE actors
(
    id         serial primary key,
    name       varchar(60) not null UNIQUE,
    dob        timestamp,
    removed    bool,
    created_at timestamp
);

comment on column actors.id is 'id записи';
comment on column actors.name is 'полное имя';
comment on column actors.dob is 'дата рождения';
comment on column actors.removed is 'был удален';
comment on column actors.created_at is 'дата создания записи';

CREATE TABLE parts
(
    movie_id INTEGER NOT NULL,
    actor_id INTEGER NOT NULL,
    CONSTRAINT unique_parts UNIQUE (movie_id, actor_id)
);

comment on column parts.movie_id is 'id фильма';
comment on column parts.actor_id is 'id актера';

ALTER TABLE parts
    ADD CONSTRAINT movie_id
        foreign key (movie_id) references movies (id);

ALTER TABLE parts
    ADD CONSTRAINT actor_id
        foreign key (actor_id) references actors (id);

