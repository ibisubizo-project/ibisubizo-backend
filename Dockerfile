FROM golang:1.10

# create a working directory
WORKDIR /go/src/github.com/ofonimefrancis/problemsApp/

# install packages

# add source code
ADD ./ /go/src/github.com/ofonimefrancis/problemsApp/


RUN echo $GOPATH

RUN go get -d -v

# build main.go
RUN go build main.go
# run the binary
EXPOSE 8000

CMD ["./main"]