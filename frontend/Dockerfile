FROM ubuntu
LABEL authors="aaditya"

RUN mkdir /app
WORKDIR /app
RUN mkdir img
COPY .main ./build
RUN chmod +x ./build

ENTRYPOINT ["./build"]