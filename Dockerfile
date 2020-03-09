#=============================================================
#--------------------- build stage ---------------------------
#=============================================================
FROM golang:1.13.5-stretch AS build_stage
ENV PACKAGE_PATH=github.com/eduardosz98
RUN mkdir -p /go/src/
WORKDIR /go/src/$PACKAGE_PATH/gorest-oracle
COPY . /go/src/$PACKAGE_PATH/gorest-oracle
RUN go get github.com/godror/godror
RUN go get github.com/joho/godotenv
RUN go build -o gorest-oracle
#=============================================================
#--------------------- final stage ---------------------------
#=============================================================
FROM oracle/instantclient:19 AS final_stage
ENV PACKAGE_PATH=github.com/eduardosz98
COPY --from=build_stage /go/src/$PACKAGE_PATH/gorest-oracle /go/src/$PACKAGE_PATH/gorest-oracle
WORKDIR /go/src/$PACKAGE_PATH/gorest-oracle
ENTRYPOINT ./gorest-oracle
EXPOSE 80