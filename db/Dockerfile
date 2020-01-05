FROM mysql:8

RUN mkdir /usr/src/init
COPY ./init_mysql.sql /usr/src/init

#CMD mysql -u ${MYSQL_USER} -p ${MYSQL_ROOT_PASSWORD} -e 'source init_mysql.sql'