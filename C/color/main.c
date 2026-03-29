#include "stdio.h"

int main(int argc, char **argv) {
  // FILE *config = fopen("/home/zoren/.config/kitty/theme.conf", "w");
  FILE *config = fopen("test.test", "w");
  fprintf(config,argv[1]);
}
