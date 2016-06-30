# golang-arduino
## Connect a PC to Arduino through the serial port. Send data from one to the other.

This example has the arduino code that runs in the board, and the GO code that compiles to an executable (serial.exe) that runs on a Windows PC.

## steps:
* upload the file *arduino/SerialToPcTest.ino* to your arduino board
* compile the GO code in /serial (download this folder to your gopath/src/ folder and *go install serial*)
* or just download the /bin/serial.exe to your PC and run it
* open your browser and type *http://localhost:8080/send* to send data to arduino
* open another tab on your browser and type *http://localhost:8080/receive* to receive data from arduino
