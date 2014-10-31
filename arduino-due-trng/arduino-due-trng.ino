// trng random number generator, adapted from
// http://www.kerrywong.com/2014/10/19/using-arduino-dues-true-random-number-generator/

void setup() {
  Serial.begin(230400);
  pmc_enable_periph_clk(ID_TRNG);
  trng_enable(TRNG);

  NVIC_DisableIRQ(TRNG_IRQn);
  NVIC_ClearPendingIRQ(TRNG_IRQn);
  NVIC_SetPriority(TRNG_IRQn, 0);
  NVIC_EnableIRQ(TRNG_IRQn);
  trng_enable_interrupt(TRNG);
}

void TRNG_Handler(void) {
  uint32_t stat = trng_get_interrupt_status(TRNG);

  if ((stat & TRNG_ISR_DATRDY) == TRNG_ISR_DATRDY) {
    uint32_t r = TRNG->TRNG_ODATA;
    Serial.write((uint8_t*)&r, sizeof(uint32_t));
  }
}

void loop() {
}
