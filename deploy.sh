#!/bin/bash
fuser -k 443/tcp
fuser -k 8080/tcp
./url-shortener