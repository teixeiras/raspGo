#include <stdio.h>
#include <string.h>
#include <stdlib.h>

#include "readproc.h"

void initArray(Array *a, size_t initialSize) {
  a->array = (proc_t *)malloc(initialSize * sizeof(proc_t));
  a->used = 0;
  a->size = initialSize;
}

void insertArray(Array *a, proc_t element) {
  if (a->used == a->size) {
    a->size *= 2;
    a->array = (proc_t *)realloc(a->array, a->size * sizeof(proc_t));
  }
  a->array[a->used++] = element;
}

void freeArray(Array *a) {
  free(a->array);
  a->array = NULL;
  a->used = a->size = 0;
}

Array ProcReaderListProcess(int * length)
{

  PROCTAB* proc = openproc(PROC_FILLMEM | PROC_FILLSTAT | PROC_FILLSTATUS);
	proc_t proc_info;
	(*length) = 0;
	Array array;

	initArray(&array, 1);  // initially 5 elements	
	memset(&proc_info, 0, sizeof(proc_info));
	while (readproc(proc, &proc_info) != NULL) {
		insertArray(&array, proc_info); 
		memset(&proc_info, 0, sizeof(proc_info));
	  (*length) ++;
	}
	closeproc(proc);
  	return array;
}