#!/bin/sh
set -e

echo "migrate_common"
./migrate_common

echo "users"
./users
