FROM postgres:latest

ENV POSTGRES_DB=agg_db
ENV POSTGRES_USER=agg_user
ENV POSTGRES_PASSWORD=test

EXPOSE 5432

CMD ["postgres"]