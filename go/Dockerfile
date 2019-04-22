FROM golang:latest 
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN go get cloud.google.com/go/profiler
RUN go get strconv
RUN go build -o main . 
CMD ["/app/main"]