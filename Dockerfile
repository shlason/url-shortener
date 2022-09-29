FROM shlason/ubuntu

RUN apt-get update -y

WORKDIR /home

# RUN mkdir go
# RUN mkdir go-app

# RUN apt-get install -y wget
# RUN apt-get install -y mysql-server

# RUN wget https://go.dev/dl/go1.18.1.linux-amd64.tar.gz
# RUN tar -C /usr/local -xzf go1.18.1.linux-amd64.tar.gz
# RUN rm go1.18.1.linux-amd64.tar.gz

ENV PATH=$PATH:/usr/local/go/bin
ENV GOPATH=$HOME/go

WORKDIR /go_app

COPY ./ ./

EXPOSE 443

RUN go mod tidy
RUN go build ./main.go
RUN service mysql start

CMD [ "./main" ]