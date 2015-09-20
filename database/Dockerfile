# Example Dockerfile for postgres database for "github.com/billychappell/blog" package for Golang.

FROM ubuntu

MAINTAINER chappell28@gmail.com

RUN apt-key adv --keyserver hkp://p80.pool.sks-keyservers.net:80 --recv-keys B97B0AFCAA1A47F044F244A07FCC7D46ACCC4CF8

RUN echo "deb http://apt.postgresql.org/pub/repos/apt/ precise-pgdg main" > /etc/apt/sources.list.d/pgdg.list

RUN apt-get update && apt-get install -y python-software-properties software-properties-common postgresql-9.4 postgresql-client-9.4 postgresql-contrib-9.4

USER postgres

RUN /etc/init.d/postgresql start &&\
	psql --command "CREATE USER docker WITH SUPERUSER PASSWORD 'docker';" &&\
	createdb -O docker docker

RUN echo "host all all 0.0.0.0/0 md5" >> /etc/postgresql/9.4/main/pg_hba.conf

RUN echo "listen_addresses='*'" >> /etc/postgresql/9.4/main/postgresql.conf

EXPOSE 5432

VOLUME ["/etc/postgresql", "/var/log/postgresql", "/var/lib/postgresql"]

RUN localedef -i en_US -c -f UTF-8 -A /usr/share/locale/locale.alias en_US.UTF-8

ENV LANG en_US.utf8

CMD ["/usr/lib/postgresql/9.4/bin/postgres", "-D", "/var/lib/postgresql/9.4/main" "-c", "config_file=/etc/postgresql/9.4/main/postgresql.conf"]
