#!/bin/bash
curl -s http://127.0.0.1:10000/ok > /tmp/response
for i in {1..99}
do
  curl -s http://127.0.0.1:10000/ok >> /tmp/response
done
AZ1=$(cat /tmp/response | grep -Fo AZ1 | wc -l)
AZ2=$(cat /tmp/response | grep -Fo AZ2 | wc -l)
AZ3=$(cat /tmp/response | grep -Fo AZ3 | wc -l)

echo "Made 100 Http calls"
echo "AZ1 OK traffic: $AZ1"
echo "AZ2 OK traffic: $AZ2"
echo "AZ3 OK traffic: $AZ3"