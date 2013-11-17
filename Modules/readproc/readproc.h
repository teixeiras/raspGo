#ifndef PROCREADER_DEFINE
#define PROCREADER_DEFINE
#include <proc/readproc.h>

typedef struct {
  proc_t *array;
  size_t used;
  size_t size;
} Array;

//extern "C" {
	Array ProcReaderListProcess(int * length);
//}
#endif