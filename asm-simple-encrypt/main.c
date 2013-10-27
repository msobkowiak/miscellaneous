/*
 * Author: Monika Sobkowiak <monika.sobkowiak@gmail.com>
 * 
 * For documentation see encryp.S
 */

#include <stdio.h>
#include "encrypt.h"

char* str1 = NULL;
char* str2 = NULL;

int main()
{
    char phrase1[10] = "abc def0";	
    char phrase2[10] = "";
    str1 = phrase1;
    str2 = phrase2;
	
    int result;
	
    result = encrypt();
    printf(
        "An original phrase: %s\nChanged phrase: %s\nNumber of changed characters: %d\n",
        phrase1, phrase2, result
    );
	
    return 0;
}		
