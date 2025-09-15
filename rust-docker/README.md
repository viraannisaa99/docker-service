# Struktur Folder

1. app/src.main.rs → entry point aplikasi, persis kayak main di java atau index.js di node
2. Cargo.toml → definisi project (name, version, dependencies, edition)
3. target/ → build-in folder seperti vendor di laravel atau node_modules di node

# DockerFile


# Jika ada error permission denied/ownership:

- chown -R dev:dev $CARGO_HOME $RUSTUP_HOME /workspace
- sudo chown -R $(id -u):$(id -g) ../../docker-apps/rust-workspace/target
- sudo chmod -R u+rwX ../../docker-apps/rust-workspace/target
- docker compose run --rm --user root rust-dev \
    bash -lc 'chown -R 1000:1000 /usr/local/cargo/registry /usr/local/cargo/git'