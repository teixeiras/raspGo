package osRequests
/*
#include <stdio.h>
#include <utmp.h>
char getCurrentRunningLevel()
{
	struct utmp *ut;
	char character=0;

    setutent();
    while ((ut = getutent()) != NULL) {
            if (ut->ut_type == RUN_LVL) {
                    character = ut->ut_pid % 256;
                    endutent();
                    return character;
            }
    }

    endutent();
    return character;
}
*/
import "C"


