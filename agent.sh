#!/bin/bash

cd agent && mkdir -p bin && go build -o bin/ ./... && ./bin/agent 