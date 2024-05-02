FROM ubuntu
LABEL authors="aaditya"

RUN mkdir /app
WORKDIR /app
COPY auth-micro/build/run .

ENTRYPOINT ./run