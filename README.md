[![Contributors][contributors-shield]][contributors-url]
[![Issues][issues-shield]][issues-url]

<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/yuriykis/funny-filter">
  </a>

<h3 align="center">Funny-Filter (ff)</h3>

  <p align="center">
   funny-filter (ff) is a network filter CLI tool for Linux, designed to limit bandwidth and packets per second for specific IP addresses. It allows users to set and unset bandwidth and packet rate limits on network interfaces.
    <br />
    <a href="https://github.com/yuriykis/funny-filter/issues">Report Bug</a>
    ·
    <a href="https://github.com/yuriykis/funny-filter/issues">Request Feature</a>
  </p>
</div>

## Features
* Bandwidth Limiting: Set and unset bandwidth limits for specific IP addresses.
* Packets Limiting: Set and unset packet rate limits for specific IP addresses.

## Getting Started

### Prerequisites
* Linux 
* GO 1.21.4 or higher

### Installation
1. Clone the repo
   ```sh
   git clone https://github.com/yuriykis/funny-filter
    ```
2. Build the application binary
    ```sh
    make build
    ```
3. Optionally, you can adjust the build parameters in the Makefile:
    ```
    BINARY_NAME=ff
    GO=/path/to/go
    IP=80.249.99.148
    BW_LIMIT=100kbps
    P_LIMIT=10
    INTERFACE=enp0s5
    ```
4. Run the application with make
    ```sh
    make run
    ```
    ```sh
    make set-bw
    ```
    ```sh
    make unset-bw
    ```
    ```sh
    make set-p
    ```
    ```sh
    make unset-p
    ```
5. Install the application
    ```sh
    sudo make install
    ```

## Usage
### Commands
The tool supports various commands:

* bandwidth set: Set a bandwidth limit on a specific IP.
* bandwidth unset: Remove a bandwidth limit from a specific IP.
* packets set: Set a packet rate limit on a specific IP.
* packets unset: Remove a packet rate limit from a specific IP.

### Examples

#### Set a bandwidth limit on a specific IP
```sh
ff bandwidth set --dev enp0s5 --ip 80.249.99.148 --limit 100kbps
```
#### Remove a bandwidth limit from a specific IP
```sh
ff bandwidth unset --dev enp0s5 --ip 80.249.99.148 --limit 100kbps
```
#### Set a packet rate limit on a specific IP
```sh
ff packets set --ip 80.249.99.148 --limit 10
```
#### Remove a packet rate limit from a specific IP
```sh
ff packets unset --ip 80.249.99.148 --limit 10
```

## Testing
### To run the automated tests for the app, use:
```sh
make test
```

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->
[contributors-shield]: https://img.shields.io/github/contributors/yuriykis/bluetooth-keepalive.svg?style=for-the-badge
[contributors-url]: https://github.com/yuriykis/bluetooth-keepalive/graphs/contributors
[issues-shield]: https://img.shields.io/github/issues/yuriykis/bluetooth-keepalive.svg?style=for-the-badge
[issues-url]: https://github.com/yuriykis/bluetooth-keepalive/issues
[license-shield]: https://img.shields.io/github/license/yuriykis/bluetooth-keepalive.svg?style=for-the-badge
[license-url]: https://github.com/yuriykis/bluetooth-keepalive/blob/master/LICENSE.txt