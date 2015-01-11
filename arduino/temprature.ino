#define TEMP_OUT A0
void setup() {
  // put your setup code here, to run once:
  Serial.begin(9600);
  pinMode(TEMP_OUT,INPUT);
}

void loop() {
  // put your main code here, to run repeatedly:
  int i;
  float t;
  //char str[8];
  i = analogRead(TEMP_OUT);
  t = i*500/1023;
  //dtostrf(t,2,2,str);
  Serial.println(t);
  delay(1000);
}
