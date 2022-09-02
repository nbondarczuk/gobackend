# gobackend

## Purpose

A really simple example of polimorfism used for backend DB connection check.
SQL and NoSQL databases are both supported as it overcomes the SQL interface
layer.

## Backends supported

The kinds of backend supported (so far):

- In memory cache db (always succeeds)
- MySQL 
- Postgres
- Oracle
- ETCD
- Mongo DB
- Dynamo DB

## Description of functionality

### version check

The tool allows to check version of each used backend DB. The backend
to be checked is to be used in command line parameter. It also pings
so int can be used to check the DB engine connectivity. The version
check is done by a very simple select operation, which in fact
depends on the partivular DB.

### ping timing

The next mode of the work is to do a ping to a particular backend 
giving out the timing.

## Configuration and usage

The credentials of the backend DB can be loaded from 
- config.yaml file
- env variables
- command line parameters
