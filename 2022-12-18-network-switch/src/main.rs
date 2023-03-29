use pcap::{Capture, Device};

fn main() {
    let main_device = Device::lookup().unwrap().unwrap();
    println!("device name : {}", main_device.name);
    let mut cap = Capture::from_device(main_device)
        .unwrap()
        .promisc(true)
        .snaplen(5000)
        .open()
        .unwrap();

    while let Ok(packet) = cap.next_packet() {
        println!("received packet! {:?}", packet);
    }
}
