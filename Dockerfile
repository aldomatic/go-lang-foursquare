FROM golang:1.6

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH
ENV GOBIN $GOPATH/bin

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
WORKDIR $GOPATH

EXPOSE 3000

RUN go get github.com/codegangsta/negroni
RUN go get github.com/gorilla/mux
RUN go get github.com/elbuo8/4square-venues

COPY app.go $GOPATH

RUN go install $GOPATH/app.go

CMD ["./bin/app"]