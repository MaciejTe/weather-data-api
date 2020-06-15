#!/bin/sh
while true; do
  go build -o weather_app
  $@ &
  PID=$!
  inotifywait -r -e modify .
  kill $PID
done
