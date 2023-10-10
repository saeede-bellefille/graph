Challenge Solution
===

This project contains 5 modules

1. socket

    This module is a library for handling sockets used by other modules.
    It contains 2 structs, `Client` and `Server` which implement client side and server side sockets.

2. sender

    This module sends provided number of packets (each one between 50 bytes and 8k bytes) to `/send` via http POST method.
    It gets count from command line argument.

3. receiver

    This module receives data from `/send` http api and sends whole request body to broker via socket.
    This module uses [echo](https://echo.labstack.com/) for api.

4. broker

    This module receives data from socket, sends it to log and then sends it to destination module via socket

5. destination

    This module receives data from socket and prints number of received packets and total count of packets and size of each packet to output.
