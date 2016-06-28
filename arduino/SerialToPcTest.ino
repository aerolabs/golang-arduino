/*
  Helice Project
  Serial Event example

 When new serial data arrives, this sketch adds it to a String.
 When a newline is received, the loop prints the string and
 clears it.

 This example code is in the public domain.

 http://www.arduino.cc/en/Tutorial/SerialEvent

 */

String inputString = "";         // a string to hold incoming data
boolean stringComplete = false;  // whether the string is complete
int count=34;

void setup() {
  // initialize serial:
  Serial.begin(9600);
  // reserve 200 bytes for the inputString:
  inputString.reserve(200);
}

void loop() {
  count++;
  if (count>1000000) count=0;
}

/*
  SerialEvent occurs whenever a new data comes in the
 hardware serial RX.  This routine is run between each
 time loop() runs, so using delay inside loop can delay
 response.  Multiple bytes of data may be available.
 */
 
void serialEvent() {
  while (Serial.available()) {
    // get the new byte:
    char inChar = (char)Serial.read();
    
    // add it to the inputString:
    // if the incoming character is a newline, set a flag
    // so the main loop can do something about it:
    if (inChar =='\n'){
      SendData(inputString);
      inputString ="";
    } else {
        inputString += inChar;
    }

  }
}

void SendData(String input){
      Serial.println("Starting the Arduino measurements for prop ID="+input);
      delay(1000);
      String numberString = String(count);
      Serial.println("Timer value: "+numberString );
      delay(1000);
      Serial.println("Some more data");
      delay(1000);
      Serial.println("Data transfer END");

}


