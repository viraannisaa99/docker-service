FROM php:8.2-apache

ARG UID=1000
ARG GID=1000

ENV ACCEPT_EULA=Y \
    DEBIAN_FRONTEND=noninteractive \
    COMPOSER_HOME=/var/www/.composer \
    PATH="/opt/mssql-tools18/bin:${PATH}"

# --- deps utama + Microsoft repo (Debian 12/bookworm) ---
RUN apt-get update && apt-get install -y --no-install-recommends \
    curl ca-certificates gnupg2 \
    unixodbc unixodbc-dev \
    libgssapi-krb5-2 \
    libzip-dev zip unzip git \
    libpng-dev libjpeg62-turbo-dev libfreetype6-dev \
    $PHPIZE_DEPS \
 && mkdir -p /usr/share/keyrings \
 && curl -fsSL https://packages.microsoft.com/keys/microsoft.asc | gpg --dearmor -o /usr/share/keyrings/ms-prod.gpg \
 && echo "deb [arch=amd64 signed-by=/usr/share/keyrings/ms-prod.gpg] https://packages.microsoft.com/debian/12/prod bookworm main" > /etc/apt/sources.list.d/microsoft.list \
 && apt-get update && apt-get install -y --no-install-recommends \
    msodbcsql18 mssql-tools18 \
 && rm -rf /var/lib/apt/lists/*

# --- PHP ext inti ---
RUN docker-php-ext-configure gd --with-freetype --with-jpeg \
 && docker-php-ext-install -j"$(nproc)" gd zip pdo_mysql bcmath

 # Tambahan deps untuk ekstensi opsional
RUN apt-get update && apt-get install -y --no-install-recommends \
    libmemcached-dev zlib1g-dev libssl-dev libevent-dev \
    libmagickwand-dev \
 && rm -rf /var/lib/apt/lists/*

# --- SQL Server drivers via PECL ---
RUN pecl install sqlsrv pdo_sqlsrv \
 && docker-php-ext-enable sqlsrv pdo_sqlsrv

# --- Apache & user ---
RUN a2enmod rewrite \
 && groupmod -o -g "${GID}" www-data \
 && usermod  -o -u "${UID}" www-data

# --- Composer ---
COPY --from=composer:2 /usr/bin/composer /usr/bin/composer

# --- php.ini mini ---
RUN printf "memory_limit=512M\nupload_max_filesize=64M\npost_max_size=64M\ndate.timezone=Asia/Jakarta\n" > /usr/local/etc/php/conf.d/custom.ini

WORKDIR /var/www/html
RUN mkdir -p /var/www/.composer && chown -R www-data:www-data /var/www

EXPOSE 80
CMD ["apache2-foreground"]
