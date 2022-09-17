#!/bin/bash

set -x

for i in {1..5}
do
   time go run client/client.go -n=4 -flags=c -pattern=alexander.com -prefix=vm
done

for i in {1..5}
do
   # time go run client/client.go -n=4 -flags=Ec -pattern="Windows N[A-Za-z][A-Za-z0-9]*" -prefix=vm
   time go run client/client.go -n=4 -flags=c -pattern="Windows NT" -prefix=vm
done

for i in {1..5}
do
   time go run client/client.go -n=4 -flags=c -pattern=http -prefix=vm
done

# for i in {1..5}
# do
#    time go run client/client.go -n=4 -flags=c -pattern=05/Feb/2024 -prefix=vm
# done