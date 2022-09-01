CREATE OR REPLACE USER 'sysrun'@localhost IDENTIFIED BY 'sysrun';

grant all privileges on DB1.* TO 'sysrun'@localhost identified by 'sysrun';

flush privileges;
