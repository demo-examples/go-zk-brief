#!/bin/sh

curl http://192.168.35.141:9090/servicelist -H "Accept:application/json" -d "key=\"1122-3344\"&destName=\"hello\"&zkidc=\"qa\""
#curl -d "key=1122-3434&destName=hello&zkidc=qa" http://127.0.0.1:9090/servicelist

echo ""

