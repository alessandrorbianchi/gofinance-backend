create table if not exists users (
    id serial primary key not null,
    st_username varchar not null,
    st_password varchar not null,
    st_email varchar unique not null,
    dt_created_at timestamptz not null default now()
);

create table if not exists categories (
    id serial primary key not null,
    co_user_id int not null,
    st_title varchar not null,
    st_type varchar not null,
    st_description varchar not null,
    dt_created_at timestamptz not null default now()
);

alter table categories add foreign key (co_user_id) references users (id);

create table if not exists accounts (
    id serial primary key not null,
    co_user_id int not null,
    co_category_id int not null,
    st_title varchar not null,
    st_type varchar not null,
    st_description varchar not null,
    vl_value integer NOT NULL,
    dt_date date NOT NULL,
    dt_created_at timestamptz not null default now()
);

alter table accounts add foreign key (co_user_id) references users (id);
alter table accounts add foreign key (co_category_id) references categories (id);





