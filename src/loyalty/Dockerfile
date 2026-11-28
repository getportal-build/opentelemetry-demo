FROM golang:1.24 AS build
WORKDIR /src
COPY . .
RUN go build -o /loyalty .
FROM gcr.io/distroless/static
COPY --from=build /loyalty /loyalty
ENTRYPOINT ["/loyalty"]
