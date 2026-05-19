<header>
<p align="center">
  <a href="https://github.com/Nyveruus/doh-proxy">
    <img src=".assets/banner.jpg" width="200" alt="doh-proxy">
  </a>
</p>
<h1 align="center">DoH Proxy</h1>
<p align="center">
Keep DNS queries private and prevent SNI-based blocking & interception without per-host configuration
</p>
<p align="center">
  <img src="https://img.shields.io/github/contributors/nyveruus/doh-proxy?style=flat-square" alt="Contributors">
  <img src="https://img.shields.io/github/repo-size/nyveruus/doh-proxy?style=flat-square" alt="Repo Size">
  <img src="https://img.shields.io/github/issues/nyveruus/doh-proxy?style=flat-square" alt="Issues">
  <img src="https://img.shields.io/github/license/nyveruus/doh-proxy?style=flat-square" alt="License">
</p>
</header>

Minimalistic service for home networks written in Go that takes plain DNS queries in the local network and forwards them to an encrypted DoH endpoint over the internet (default Cloudflare). Preserves all headers including ECH for hosts that support it.

## Usage

```
$ sudo ./install.sh
$ systemctl enable --now doh-proxy
```
Set local IP as the DNS server in your router or DHCP settings.

## History

The idea developed when setting up a new home network with a Mikrotik router. After configuring its DNS resolver with DoH, it would strip ECH records causing SNI to be exposed in TLS handshakes. ECH needs to be preserved to prevent SNI-based blocking, so I wrote something light intended to run on a Raspberry Pi.
