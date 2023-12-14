# pi_radio

Radio with Raspberry Pi

## Requirenments

- Raspberry Pi (I use a Pi 4)
- USB Keypad (I use [this one](https://www.amazon.de/LogiLink-ID0184-Zusatztastatur-USB-Anschluss-LED-Aktivitätsanzeige/dp/B07KGLWY64/))
- speakers
- I2C screen (I use [this one](https://www.az-delivery.de/en/products/hd44780-2004-lcd-display-bundle-4x20-zeichen-mit-i2c-schnittstelle-gruen))

## Setup

1. wire it up

   **TODO**

2. [Install Raspberry pi OS lite (64 bit)](https://www.raspberrypi.com/documentation/computers/getting-started.html#raspberry-pi-imager)
3. [install Go](https://go.dev/doc/install)
4. download the source code

   ```bash
   git clone https://github.com/Letsric/pi_radio.git
   ```

5. install dependencies

   ```bash
   sudo apt update
   sudo apt install
   ```

   **TODO** Write down all dependencies

6. build the source code

   ```bash
   cd pi_radio
   go build
   ```

7. setup a Service
   
   edit the file `/etc/systemd/system/piradio.service` like this:
   
   ```
   [Unit]
   Description=Radio
   
   [Service]
   ExecStart=/home/eric/pi_radio/pi_radio
   Type=exec
   Restart=always
   
   [Install]
   WantedBy=multi-user.target
   ```
   
   then run this to make it auto-start:

   ```bash
   sudo systemctl enable --now piradio.service
   ```
