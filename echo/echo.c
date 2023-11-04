#include <stdio.h>
#include <string.h>
#include <stdbool.h>


static bool show_trailing_newline = true;
static bool unescape_backslash = false;

static void print_word(char *word) {
  for (int i=0; word[i]; i++) {
    char c = word[i];

    if (unescape_backslash && c == '\\') {
      // \でエスケープされている場合
      i++;
      c = word[i];
      switch (c) {
        case 'n':
          printf("\n");
          break;
        case 't':
          printf("\t");
          break;
        default:
          printf("\\%c", c);
          break;
      }
    } else {
      printf("%c", c);
    }
  }
}

int main(int argc, char **argv)
{
  if (argc == 1) {
    return 0;
  }


  for (int i=1; i<argc; i++) {
    char *arg = argv[i];

    if (strcmp(arg, "-n") == 0) {
      show_trailing_newline = false;
      continue;
    }

    if (strcmp(arg, "-e") == 0) {
      unescape_backslash = true;
      continue;
    }

    print_word(arg);

    if (i != argc-1) {
      printf(" ");
    }
  }

  if (show_trailing_newline) {
    printf("\n");
  }

  return 0;
}
