FROM php:8.0-apache

ENV ACCEPT_EULA=Y \
    DEBIAN_FRONTEND=noninteractive \
    PATH="/opt/mssql-tools18/bin:${PATH}"

# Tool build PECL + ODBC
RUN apt-get update && apt-get install -y --no-install-recommends \
    curl ca-certificates gnupg2 apt-transport-https \
    unixodbc unixodbc-dev \
    libgssapi-krb5-2 \
    libzip-dev zip unzip git \
    libpng-dev libjpeg62-turbo-dev libfreetype6-dev \
    autoconf make gcc g++ pkg-config \
 && rm -rf /var/lib/apt/lists/*

# Add Microsoft repo (untuk Debian 11 / bullseye → php:8.0-apache base image)
RUN mkdir -p /usr/share/keyrings \
 && curl -fsSL https://packages.microsoft.com/keys/microsoft.asc | gpg --dearmor -o /usr/share/keyrings/ms-prod.gpg \
 && echo "deb [arch=amd64 signed-by=/usr/share/keyrings/ms-prod.gpg] https://packages.microsoft.com/debian/11/prod bullseye main" \
    > /etc/apt/sources.list.d/microsoft.list \
 && apt-get update && apt-get install -y --no-install-recommends \
    msodbcsql18 mssql-tools18 \
 && rm -rf /var/lib/apt/lists/*

# PHP ext umum (gd, zip)
RUN docker-php-ext-configure gd --with-freetype --with-jpeg \
 && docker-php-ext-install -j"$(nproc)" gd zip

# SQL Server drivers (specify versi biar aman di PHP 8.0)
RUN pecl install sqlsrv-5.10.1 pdo_sqlsrv-5.10.1 \
 && docker-php-ext-enable sqlsrv pdo_sqlsrv

# Apache rewrite buat CI/Laravel
RUN a2enmod rewrite

WORKDIR /var/www/html
RUN mkdir -p /var/www/.composer && chown -R www-data:www-data /var/www

EXPOSE 80
CMD ["apache2-foreground"]
