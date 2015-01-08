FROM debian:squeeze

RUN apt-get update
RUN apt-get -y upgrade

# Install and config apache and fastcgi
RUN DEBIAN_FRONTEND=noninteractive apt-get -y install apache2 curl mysql-client apache2-suexec libapache2-mod-fcgid
RUN a2enmod rewrite && \
	a2enmod suexec && \
	a2enmod include && \
	a2enmod fcgid && \
	a2dismod cgid

RUN rm /etc/apache2/envvars
ENV APACHE_RUN_USER www-data
ENV APACHE_RUN_GROUP www-data
ENV APACHE_LOG_DIR /var/log/apache2/
ENV APACHE_LOCK_DIR /var/lock/apache2
ENV APACHE_PID_FILE /var/run/apache2.pid

# Permissions
RUN chown www-data.www-data -R /var/lib/apache2/fcgid
RUN touch /var/run/docker-registry.pid && \
    chown www-data.www-data -R /var/run/docker-registry.pid

EXPOSE 80

COPY apache-config.conf /etc/apache2/sites-enabled/000-default
COPY public /var/www

CMD /usr/sbin/apache2ctl -D FOREGROUND
