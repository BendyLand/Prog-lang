#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdbool.h>
#include "keywords.h"

#define MAX_LINES 100

/**  
 * States:
 * 0 = Parser
 * 1 = String
 * 2 = Comment
 */
typedef struct {
    char currentChar;
    char nextChar;
    int index;
    int state;
} Parser;

char* getFileContents();
void splitIntoLines(char**, char*);
void filterSignificantTokens(char** buffer);

int main(void) {
    char* data = getFileContents();
    char** buffer = (char**)malloc(sizeof(char*) * (MAX_LINES + 1));
    memset(buffer, 0, sizeof(char*) * (MAX_LINES + 1));

    splitIntoLines(buffer, data);
    filterSignificantTokens(buffer);
    int i = 0;
    while (buffer[i] != NULL) {
        printf("%s\n", buffer[i]);
        free(buffer[i]); 
        i++;
    }
    free(buffer);
    free(data);

    return 0;
}

void filterSignificantTokens(char** buffer) {
    char** result = (char**)malloc(sizeof(buffer) / sizeof(char*) + 1);
    if (result == NULL) {
        perror("Unable to allocate for result.");
        return;
    }
    int i = 0;
    while (buffer[i] != NULL) {
        char* line = (char*)malloc(strlen(buffer[i]) / sizeof(char) + 1);
        if (line == NULL) {
            perror("Unable to allocate memory for line.\n");
            return;
        }
        strcpy(line, buffer[i]);
        char* tokens  = strtok(line, " ");
        if (isKeyword(&tokens[0])) {
            strcpy(result[i], line); // problem
            printf("I'm here\n");
        }
    }
    memset(buffer, 0, 0);
    i = 0;
    while (result[i] != NULL) {
        int length = strlen(result[i]);
        buffer[i] = (char*)malloc(strlen(result[i]) / sizeof(char));
        if (buffer[i] == NULL) {
            perror("Memory allocation to buffer failed.\n");
            continue;
        }
        strcpy(buffer[i], result[i]);
        i++;
    }
}

void splitIntoLines(char** buffer, char* file) {
    int line = 0;
    int j = 0;
    char temp[100] = {0}; // lines must be under 100 characters long
    char c;

    int length = strlen(file);
    for (int i = 0; i < length; i++) {
        c = file[i];
        if (c == '\n' || i == length - 1) {
            temp[j] = '\0';
            buffer[line] = (char*)malloc(strlen(temp) + 1); // allocate for the line
            if (buffer[line] == NULL) {
                perror("Error creating line buffer.\n");
                continue;
            }
            strcpy(buffer[line], temp);
            memset(temp, 0, sizeof(temp));
            j = 0;
            line++;
        }
        else {
            temp[j] = c;
            j++;
        }
    }
    buffer[line] = NULL;
}

char* getFileContents() {
    FILE* fptr;
    long length;

    fptr = fopen("../../test.pr", "r");
    if (fptr == NULL) {
        perror("Problem opening file.\n");
        return NULL;
    }

    fseek(fptr, 0, SEEK_END);
    length = ftell(fptr);
    rewind(fptr);

    char* buffer = (char*)malloc(sizeof(char) * length + 1);
    if (buffer == NULL) {
        perror("Problem creating buffer");
        fclose(fptr);
        return NULL;
    }

    if (fread(buffer, sizeof(char), length, fptr) != length) {
        perror("Failed to read file.");
        fclose(fptr);
        free(buffer);
        return NULL;
    }

    buffer[length] = '\0';
    fclose(fptr);

    return buffer;
}
