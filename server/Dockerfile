FROM nginx:1.13.0
RUN mkdir /wwwroot
#COPY ./nginx.conf /etc/nginx/nginx.conf
#COPY ./nginx_dev.conf /etc/nginx/nginx_dev.conf
#COPY ./nginx_prod.conf /etc/nginx/nginx_prod.conf
COPY ./mime.types /etc/nginx/mime.types
#RUN chmod a+x /usr/local/bin/certbot-auto
#RUN certbot-auto --os-packages-only -n
#COPY ./entrypoint.sh /usr/local/bin/entrypoint.sh
#RUN chmod +x /usr/local/bin/entrypoint.sh
#ENTRYPOINT [ "/usr/local/bin/entrypoint.sh" ]
#RUN service nginx start