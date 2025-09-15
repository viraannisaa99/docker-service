use axum::{extract::{Path, Query}, response::Json, routing::get, Form, Router};
use serde::{Deserialize, Serialize};
use std::net::SocketAddr;
use tokio::net::TcpListener;

async fn hello(Path(name): Path<String>) -> String {
    format!("Hai from rust! Welcome {name}!")
}

#[tokio::main]
async fn main() {
    let app = Router::new()
        .route("/hello/{name}", get(hello))
        .route("/sum", get(sum));

    let addr: SocketAddr = "0.0.0.0:8080".parse().unwrap();
    println!("Server listening on http://{addr}");

    let listener = TcpListener::bind(addr).await.unwrap();
    axum::serve(listener, app).await.unwrap();
}
