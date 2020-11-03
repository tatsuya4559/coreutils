#ifndef _UTMPLIB_H
#define _UTMPLIB_H

int open_utmp(char*);
int reload_utmp(void);
void close_utmp(void);
struct utmp* next_utmp(void);

#endif /* utmplib.h */
