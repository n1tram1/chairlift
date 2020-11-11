#ifndef RUNTIME_H
#define RUNTIME_H

#include <stdint.h>

#define __constructor __attribute__((constructor))
#define __destructor __attribute__((destructor))

#define WINDOW_WIDTH 640
#define WINDOW_HEIGHT 480

#define CHIP8_DISPLAY_WIDTH 64
#define CHIP8_DISPLAY_HEIGHT 32

void init_runtime(void);
void draw(uint8_t vx, uint8_t vy, uint16_t vI, uint8_t *ram, uint8_t n, uint8_t *collision);
void clear_display(void);
uint8_t random_uint8(void);

#endif /* ! RUNTIME_H */
