#pragma once

#include "str.h"

string* readFile(char* path);
void lstrip(string* original);
int indexOf(string* str, char c);
int indexOfStr(string* str, char* s);
