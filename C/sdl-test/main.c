#include <SDL2/SDL_events.h>
#include <SDL2/SDL_stdinc.h>
#include <SDL2/SDL_video.h>

int main() {

  SDL_Window* win = SDL_CreateWindow("pong", SDL_WINDOWPOS_CENTERED, SDL_WINDOWPOS_CENTERED, 500, 500, 0);

  while (1) {
    SDL_Event event;
    while (SDL_PollEvent(&event)) {
      if (SDL_QUIT) {
        break;
      }
    }
  }

  return 0;

}
