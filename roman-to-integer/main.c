/*
 * https://leetcode.com/problems/roman-to-integer
*/

#include <stdio.h>
#include <string.h>

int romanToInt(char *s) {
    int result = 0;
    for (int i = 0; i < strlen(s); i++) {
        char c = s[i];
        char next = s[i + 1];

        switch (c) {
        case 'I':
            switch (next) {
            case 'V':
                result += 4;
                i++;
                break;

            case 'X':
                result += 9;
                i++;
                break;

            default:
                result++;
                break;
            }
            break;

        case 'V':
            result += 5;
            break;

        case 'X':
            switch (next) {
            case 'L':
                result += 40;
                i++;
                break;

            case 'C':
                result += 90;
                i++;
                break;

            default:
                result += 10;
                break;
            }
            break;

        case 'L':
            result += 50;
            break;

        case 'C':
            switch (next) {
            case 'D':
                result += 400;
                i++;
                break;
            
            case 'M':
                result += 900;
                i++;
                break;

            default:
                result += 100;
                break;
            }
            break;

        case 'D':
            result += 500;
            break;

        case 'M':
            result += 1000;
            break;

        default:
            break;
        }
    }

    return result;
}

int main() {
    int result = romanToInt("XIV");
    printf("%d\n", result);
}
