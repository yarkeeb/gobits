#include <stdio.h>
#include "byc.h"

int main(int argc, char** argv) {
    GoInt x = 12;
    GoInt y = 23;
    printf("About to call Go func!\n");
    PrintMessage();

    GoInt p = Multiply(x, y);
    printf("Product: %d\n", (int) p);
    printf("Works!\n");
    return 0;
}