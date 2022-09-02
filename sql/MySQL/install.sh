#!/bin/bash

RUN="mysql -u root"
$RUN @create_db.sql
$RUN @create_users.sql

