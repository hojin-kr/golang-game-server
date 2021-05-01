FROM golang:1.14
COPY ./ /home/app
WORKDIR /home/app
EXPOSE 8888/tcp
RUN go get -v ./...
RUN go install -v ./...
CMD ["gamejam"]
