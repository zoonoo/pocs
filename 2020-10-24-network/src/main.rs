use std::net::{TcpListener, TcpStream};
use tokio;

fn main() {
    println!("Hello, world!");

    let server = TcpListener::bind(&local_addr);

    let server = server.incoming().for_each(|src| {
        let connection = TcpStream::connect(&remote_addr)
        .and_then(|move |dst| copy(src, dst));

        tokio::spawn(connection)
    });

    tokio::spawn(server);
}
