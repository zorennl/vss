#include "ncurses.h"

int main() {

  initscr();
  printw("Hello, Zoren!");
  refresh();
  getch();
  clear();
  printw("Fuck you");
  refresh();
  getch();
  endwin();
  
}
