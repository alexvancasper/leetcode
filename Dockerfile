FROM golang:1.20
RUN apt update && apt install -y graphviz &&  mkdir -p /gosrc
# COPY main.go go.mod /gosrc/
WORKDIR /gosrc
ENTRYPOINT ["go", "run", "main.go"]