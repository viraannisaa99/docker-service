# Dockerfile (workspace-ready + user 'dev' + optional cargo-watch)
FROM rust:slim-bookworm

# --- tools/package dasar
RUN apt-get update && apt-get install -y --no-install-recommends \
    bash git pkg-config libssl-dev ca-certificates \
 && rm -rf /var/lib/apt/lists/*

# --- install cargo-watch (opsional buat hot reload pada dev)
RUN /usr/local/cargo/bin/cargo install cargo-watch

# --- bikin user non-root supaya prompt gak "I have no name!"
# kalau mau dinamis: build-arg UID/GID (default 1000:1000)
ARG UID=1000
ARG GID=1000
RUN groupadd -g ${GID} dev && useradd -u ${UID} -g ${GID} -ms /bin/bash dev

# --- env & PATH global (lokasi rustup/cargo)
ENV CARGO_HOME=/usr/local/cargo \
    RUSTUP_HOME=/usr/local/rustup \
    PATH=/usr/local/cargo/bin:/usr/local/rustup/bin:/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin

WORKDIR /workspace
# RUN mkdir -p /workspace && chown -R dev:dev /workspace /usr/local/cargo /usr/local/rustup

USER dev

CMD ["bash", "-lc", "sleep infinity"]
