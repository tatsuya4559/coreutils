#include <stdio.h>
#include <stdlib.h>
#include <utmp.h>
#include <fcntl.h>
#include <unistd.h>
#include <time.h>
#include "utmplib.h"

#define SHOWHOST

void show_info(struct utmp *);
void show_time(long);

int main(void) {
    struct utmp *utbufp;

    if (open_utmp(UTMP_FILE) == -1) {
        perror(UTMP_FILE);
        exit(1);
    }

    while ((utbufp = next_utmp()) != ((struct utmp*) NULL)) {
        show_info(utbufp);
    }
    close_utmp();

    return 0;
}

void show_info(struct utmp *utbufp) {
    if (utbufp->ut_type != USER_PROCESS) {
        return;
    }

    printf("%-8.8s", utbufp->ut_name);
    printf(" ");
    printf("%-8.8s", utbufp->ut_line);
    printf(" ");
    show_time(utbufp->ut_time);
    printf(" ");
#ifdef SHOWHOST
    printf("(%s)", utbufp->ut_host);
#endif
    printf("\n");
}

void show_time(long timeval) {
    char *cp = ctime(&timeval);
    printf("%12.12s", cp+4);
}
