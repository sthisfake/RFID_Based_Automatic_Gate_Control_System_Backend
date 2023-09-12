#include <SPI.h>
#include <MFRC522.h>
#include <Servo.h>

#define SS_PIN 10
#define RST_PIN 9

MFRC522 mfrc522(SS_PIN, RST_PIN);
Servo servo;  // Create a servo object

void setup() {
  Serial.begin(9600);
  SPI.begin();
  mfrc522.PCD_Init(0);
  mfrc522.PCD_SetAntennaGain(mfrc522.RxGain_max);
  servo.attach(6);  // Attach the servo to pin 8
}

void loop() {
  if (mfrc522.PICC_IsNewCardPresent() && mfrc522.PICC_ReadCardSerial()) {

    String tagId = "";

    // Read tag ID
    for (byte i = 0; i < mfrc522.uid.size; i++) {
      tagId += (mfrc522.uid.uidByte[i] < 0x10 ? "0" : "") + String(mfrc522.uid.uidByte[i], HEX);
    }

    // Send data to computer
    Serial.print("Tag ID: ");
    Serial.println(tagId);

    // Rotate the servo 90 degrees for 2 seconds
    servo.write(0);  // Rotate servo to 0 degrees
    delay(5000);     // Wait for 2 seconds
    servo.write(90); // Rotate servo back to 90 degrees (center position)


    mfrc522.PICC_HaltA();
    mfrc522.PCD_StopCrypto1();
    delay(100);  // Delay to prevent reading the same card multiple times
  }
}
