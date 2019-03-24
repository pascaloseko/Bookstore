-- drop table event;
-- drop table location;

create table location(
  id serial primary key,
  name varchar(255),
  address varchar(255),
  country varchar(255),
  open_time int,
  close_time int
);

create table event(
  id serial primary key,
  name varchar(255),
  duration int,
  start_date timestamp not null,
  end_date timestamp not null,
  loc_id integer not null references location(id)
);