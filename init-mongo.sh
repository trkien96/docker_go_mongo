#!/usr/bin/env bash
echo 'Creating application user and db';
mongo golang_mongodb  --username root  --password root  --authenticationDatabase admin  --host localhost  --port 27017  --eval "db.createUser({user: 'root', pwd: 'root', roles:[{role:'dbOwner', db: 'golang_mongodb'}]});"
echo 'User: root create to database golang_mongodb';
mongo testDB  --username root  --password root  --authenticationDatabase admin  --host localhost  --port 27017  --eval "db.createUser({user: 'root', pwd: 'root', roles:[{role:'dbOwner', db: 'testDB'}]});"
echo 'User: root create to database testDB';
