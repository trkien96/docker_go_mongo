#!/bin/bash
export $(grep -v '^#' .env | xargs)

echo "CREATING init-mongo.sh FILE..."
cat > ./init-mongo.sh <<EOF
#!/usr/bin/env bash
echo 'Creating application user and db';
mongo ${DB_DATABASE} \
 --username ${DB_USERNAME} \
 --password ${DB_PASSWORD} \
 --authenticationDatabase admin \
 --host ${DB_HOST} \
 --port ${DB_PORT} \
 --eval "db.createUser({user: '${DB_USERNAME}', pwd: '${DB_PASSWORD}', roles:[{role:'dbOwner', db: '${DB_DATABASE}'}]});"
echo 'User: ${DB_USERNAME} create to database ${DB_DATABASE}';
mongo testDB \
 --username ${DB_USERNAME} \
 --password ${DB_PASSWORD} \
 --authenticationDatabase admin \
 --host ${DB_HOST} \
 --port ${DB_PORT} \
 --eval "db.createUser({user: '${DB_USERNAME}', pwd: '${DB_PASSWORD}', roles:[{role:'dbOwner', db: 'testDB'}]});"
echo 'User: ${DB_USERNAME} create to database testDB';
EOF
echo "created..."