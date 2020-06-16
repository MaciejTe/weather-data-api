FROM golang:1.14.4-alpine as build
COPY . /app/
WORKDIR /app
RUN CGO_ENABLED=0 go build -o /app/weather_app .

FROM scratch as run
COPY --from=build /app/weather_app /bin/weather_app
ENTRYPOINT ["/bin/weather_app"]
