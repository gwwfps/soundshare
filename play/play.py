import socket
import pyaudio

p = pyaudio.PyAudio()
stream = p.open(format=p.get_format_from_width(4),
                channels=2,
                rate=48000,
                output=True)

sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
sock.bind(("192.168.1.207", 30242))
while True:
  data, addr = sock.recvfrom(32000)
  if len(data) > 0:
    stream.write(data)

stream.stop_stream()
sock.close()
stream.close()
p.terminate()
