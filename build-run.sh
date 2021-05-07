#!/bin/bash
lsof -ti:3000 | xargs kill

go test --tags dev --cover ./... &&

cd ./frontend &&

chmod +x ./test-build.sh

./test-build.sh &&

cd ../ &&

docker-compose up --build