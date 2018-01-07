# Scratch

Scratch is a toy REST API that connects to a postgres DB.

## Setting up

First start postgres:

```
docker run -it --rm -p 5432:5432 postgres
```

To run psql, use:

```
 docker run -it --rm --link some-container-name:postgres postgres psql -h postgres -U postgres
```

Finally, run the following psql commands:

```
create database mvxt;
\c mvxt
create table movies (name varchar, director varchar);
insert into movies values ('seven samurai', 'kurozawa akira');
```
