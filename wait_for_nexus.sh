#!/bin/bash

HOSTNAME=$1

echo "HOSTNAME: $HOSTNAME"

for i in $(seq 1 300); do
    STATUS_CODE=$(curl -s -o /dev/null -w "%{http_code}" -L "http://${HOSTNAME}:8081/")

    # echo "${i}: ${STATUS_CODE}"

    if [ "${STATUS_CODE}" -eq "200" ]; then
        exit 0
    fi

    sleep 1s

done

exit 1