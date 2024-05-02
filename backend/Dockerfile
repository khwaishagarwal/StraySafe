FROM ubuntu
LABEL authors="aaditya"

RUN mkdir /app
WORKDIR /app
RUN mkdir img
COPY ../backend/build/build ./build
RUN chmod +x ./build

ENTRYPOINT ["./build"]