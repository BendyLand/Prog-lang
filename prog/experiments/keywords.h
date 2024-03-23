#include <stdbool.h>

extern const char* keywords[];
extern const int numKeywords;

extern bool isKeyword(char* word);
char* getKeyword(int index);