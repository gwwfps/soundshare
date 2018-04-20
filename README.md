# soundshare

Problem: A Windows PC, a Mac, USB speakers, and no mixers

Solution: Extremely naive audio sharing

.NET program running on Windows sends raw captured WAV from NAudio loopback recording over UDP.

Python program on Mac plays everything received on the socket using PyAudio.
