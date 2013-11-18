package osRequests
/*
#cgo LDFLAGS: -lprocps
#include "readproc.h"

Array actualProcessList;
int actualProcessListSize, iterator;
void updateProcessList()
{
	actualProcessListSize = iterator = 0;
	actualProcessList = ProcReaderListProcess(&actualProcessListSize);
}

proc_t * nextProcessItem()
{
	if (iterator == actualProcessListSize) {
		return NULL;
	}
	iterator ++;
	return &actualProcessList.array[iterator - 1];
}
*/
import "C"
import (
	"fmt"
	"syscall"

)


func Openproc() {
	C.updateProcessList();
	var foo * C.proc_t;
	for foo = C.nextProcessItem() ; foo != nil; foo = C.nextProcessItem() {
		fmt.Printf("%s(%d): %d\n", (*foo).cmd,(*foo).tid, (*foo).resident);

/*
	printf("%20s:\t%5ld\t%5lld\t%5lld\n",
         proc_info.cmd, proc_info.resident,
         proc_info.utime, proc_info.stime);
         */

	}
}

func killProcess(process int) {
	syscall.Kill(process, syscall.SIGHUP);

}
