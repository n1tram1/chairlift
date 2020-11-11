#include "runtime.h"

#include <SDL2/SDL.h>
#include <err.h>
#include <stdio.h>
#include <stdlib.h>
#include <time.h>
#include <signal.h>
#include <string.h>

static SDL_Window *window;
static SDL_Renderer *renderer;

static uint8_t display_buffer[CHIP8_DISPLAY_WIDTH * CHIP8_DISPLAY_HEIGHT];

void sigint_handler(int signum)
{
	errx(1, "Received SIGINT");
}

void init_runtime()
{
	signal(SIGINT, sigint_handler);

	srandom(time(NULL));

	if (SDL_Init(SDL_INIT_VIDEO) != 0)
		errx(1, "SDL_Init: %s", SDL_GetError());


	window = SDL_CreateWindow(
			"chairlift",
			SDL_WINDOWPOS_UNDEFINED,
			SDL_WINDOWPOS_UNDEFINED,
			WINDOW_WIDTH,
			WINDOW_HEIGHT,
			SDL_WINDOW_SHOWN);

	renderer = SDL_CreateRenderer(window, -1, SDL_RENDERER_ACCELERATED);

	clear_display();
}

__destructor void fini()
{
	SDL_DestroyWindow(window);
	SDL_Quit();
}

void clear_display(void)
{
	SDL_SetRenderDrawColor(renderer, 0, 0, 0, 255);
	SDL_RenderClear(renderer);

	memset(display_buffer, 0, CHIP8_DISPLAY_WIDTH * CHIP8_DISPLAY_HEIGHT);
}

uint8_t get_pixel(uint8_t x, uint8_t y)
{
	return display_buffer[x + y * CHIP8_DISPLAY_WIDTH];
}

void set_pixel(uint8_t x, uint8_t y, uint8_t val)
{
	x %= CHIP8_DISPLAY_WIDTH;
	y %= CHIP8_DISPLAY_HEIGHT;

	display_buffer[x + y * CHIP8_DISPLAY_WIDTH] = val;
}

uint8_t draw_sprite(uint8_t vx, uint8_t vy, const uint8_t *sprite, uint8_t n)
{
	uint8_t collision = 0;
	uint8_t dest_x = vx;
	uint8_t dest_y = vy;

	for (uint8_t row = 0; row < n; row++) {
		uint8_t byte = sprite[row];

		for (uint8_t col = 0; col < 8; col++) {
			uint8_t bit = ((byte << col) & 0x80) != 0;
			uint8_t prev_pixel = get_pixel(vx + col, vy + row);
			
			uint8_t curr_pixel;

			curr_pixel = prev_pixel ^ bit;

			if (prev_pixel && !curr_pixel)
				collision = 1;

			set_pixel(vx + col, vy + row, curr_pixel);

			/* dest_y = ( + 1) % CHIP8_DISPLAY_WIDTH; */
		}

		/* dest_y = (dest_y + 1) % CHIP8_DISPLAY_HEIGHT; */
                /*  */
		/* dest_x = vx; */
	}

	return collision;
}

void apply_display_buffer(void)
{
	SDL_Rect r;
	r.w = WINDOW_WIDTH / CHIP8_DISPLAY_WIDTH;
	r.h = WINDOW_HEIGHT / CHIP8_DISPLAY_HEIGHT;

	for (int y = 0; y < CHIP8_DISPLAY_HEIGHT; y++) {
		for (int x = 0; x < CHIP8_DISPLAY_WIDTH; x++) {
			uint8_t bit = display_buffer[x + y * CHIP8_DISPLAY_WIDTH];

			r.x = x * r.w;
			r.y = y * r.h;

			if (bit) {
				SDL_SetRenderDrawColor(renderer, 255, 255, 255, 255);
			} else {
				SDL_SetRenderDrawColor(renderer, 0, 0, 0, 255);
			}

			SDL_RenderFillRect(renderer, &r);
		}
	}

	SDL_RenderPresent(renderer);
}

void print_display_buffer(void)
{
	putc(' ', stdout);
	for (int x = 0; x < CHIP8_DISPLAY_WIDTH; x++)
		putc('-', stdout);
	putc('\n', stdout);

	for (int y = 0; y < CHIP8_DISPLAY_HEIGHT; y++) {
		putc('|', stdout);
		for (int x = 0; x < CHIP8_DISPLAY_WIDTH; x++) {
			uint8_t bit = display_buffer[x + y * CHIP8_DISPLAY_WIDTH];

			if (bit)
				putc('+', stdout);
			else
				putc(' ', stdout);
		}
		putc('|', stdout);
		putc('\n', stdout);
	}

	putc(' ', stdout);
	for (int x = 0; x < CHIP8_DISPLAY_WIDTH; x++)
		putc('-', stdout);
	putc('\n', stdout);
}


void draw(uint8_t vx, uint8_t vy, uint16_t I, uint8_t *ram, uint8_t n, uint8_t *collision)
{
	const uint8_t *sprite = ram + I;

	*collision = draw_sprite(vx, vy , sprite, n);

	apply_display_buffer();

	printf("sprite @ I = 0x%x\n", I);
	print_display_buffer();
}

uint8_t random_uint8(void)
{
	return random();
}
