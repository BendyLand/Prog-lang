#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#define MAX_LINES 100

char* getFileContents();
void splitIntoLines(char**, char*);
void tokenizeLine(char** dest, char* line);

int main(void) {
    char* data = getFileContents();
    char** buffer = (char**)malloc(sizeof(char*) * (MAX_LINES + 1));
    memset(buffer, 0, sizeof(char*) * (MAX_LINES + 1));

    splitIntoLines(buffer, data);

    int i = 0;
    while (buffer[i] != NULL) {
        printf("%s \n", buffer[i]);
        i++;
    }
    char** test = (char**)malloc(sizeof(char*) * sizeof(buffer) + 1);
    tokenizeLine(test, buffer[0]);
    i = 0;
    while (test[i] != NULL) {
        printf("%s\n", test[i]);
        free(test[i]);
        i++;
    }
    i = 0;
    while (buffer[i] != NULL) {
        free(buffer[i]); 
        i++;
    }

    free(data);

    return 0;
}

void tokenizeLine(char** dest, char* line) {
    int length = strlen(line);
    int j = 0;
    int n = 0;
    char buffer[100] = {0};
    char c;

    for (int i = 0; i < length; i++) {
        c = line[i];
        if (c == ' ') {
            dest[n] = (char*)malloc(sizeof(char) * 100 + 1);
            if (dest[n] == NULL) {
                perror("Unable to allocate memory for token");
                continue;
            }
            buffer[j] = '\0';
            strcpy(dest[n], buffer);
            memset(buffer, 0, 100);
            j = 0;
            n++;
        }
        else {
            buffer[j] = c;
            j++;
        }
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
