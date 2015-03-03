// Author: Dinesh Appavoo (dineshappavoo) <dinesha.cit@gmail.com>

/*
 Algorithm REVERSE_STRING(S):
 1.Get the input string
 2.Find the length of the string
 3.Create a map of [string]boolean
 4.For i 1 to n iterate through the characters ch of the string
		if ch present in the map then return false
		else put ch to the map --> map[ch]=true
	5.return true
*/
#include<stdio.h>
void reverseString(char *);
int main() {
	char *str = "Hello Worlm";
	reverseString(str);
	printf("%s\n", str);
	return 0;
}

void reverseString(char *str) {
	char *end;
	end = str;
	if(str)
	{
		while(*end)
		{
			++end;
		}
		--end;//Last character will be null value
	}
	char tmp;
	printf("%c\n", *end);
	printf("%c\n", *str);
	while(str < end)
	{
	printf("%c\n", *end);
	printf("%c\n", *str);
		tmp = *str;
		*str = *end;
		*end = tmp;
		str++;
		end--;
	}
	printf("%c\n", *end);
	printf("%c\n", *str);
}
