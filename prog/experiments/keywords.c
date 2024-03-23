#include "keywords.h"
#include <stdio.h>
#include <string.h>
#include <stdbool.h>

const char* keywords[] = {"if", "for", "let", "print", "puts"};
const int numKeywords = sizeof(keywords) / sizeof(char*);

bool isKeyword(char* word) {
    for (int i = 0; i < numKeywords; i++) {
        bool matchFound = strcmp(word, keywords[i]);
        if (matchFound) {
            return true;
        }
    }
    return false;
}

char* getKeyword(int index) {
    if (index >= 0 && index < numKeywords) {
        return keywords[index];
    }
    return "";
}