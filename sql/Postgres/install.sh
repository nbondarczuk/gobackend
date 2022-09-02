#!/bin/bash

createuser sysrun
createdb db1
psql -d db1 -U sysrun < create_user.sql
