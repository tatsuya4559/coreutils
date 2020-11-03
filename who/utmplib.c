#include "utmplib.h"
#include <stdio.h>
#include <fcntl.h>
#include <unistd.h>
#include <utmp.h>

#define NBUFRECS 16
#define NULLUT ((struct utmp*) NULL)
#define UTSIZE (sizeof(struct utmp))

static char utmpbuf[NBUFRECS * UTSIZE];
static int num_recs;
static int cur_rec;
static int fd_utmp = -1;

int open_utmp(char *filename) {
    fd_utmp = open(filename, O_RDONLY);
    cur_rec = 0;
    num_recs = 0;
    return fd_utmp;
}

struct utmp* next_utmp(void) {
    struct utmp *recp;

    if (fd_utmp == -1) {
        return NULLUT;
    }
    if (cur_rec == num_recs && reload_utmp() == 0) {
        return NULLUT;
    }

    recp = (struct utmp*) &utmpbuf[cur_rec * UTSIZE];
    cur_rec++;
    return recp;
}

int reload_utmp(void) {
    int amt_read;
    amt_read = read(fd_utmp, utmpbuf, NBUFRECS * UTSIZE);
    num_recs = amt_read / UTSIZE;
    cur_rec = 0;
    return num_recs;
}

void close_utmp(void) {
    if (fd_utmp != -1) {
        close(fd_utmp);
    }
}
