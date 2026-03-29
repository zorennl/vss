#include "stdio.h"

int main() {
  int i;
  int x;
  int z;
  long double mydoubles[99999];

  for (i=0;i < 99999;i++) {
    mydoubles[i] = 9.999999999;
    printf("%Lf",mydoubles[i]);
    for (x=0;x < 99999;x++) {
      mydoubles[x] *= 244;
      printf("%Lf",mydoubles[x]);
      for (z=0;z < 99999;z++) {
        mydoubles[z] /= 33;
        printf("%Lf",mydoubles[z]);
      }
    }
  }
  printf("%zu",sizeof(mydoubles));
}
