#include <stdio.h>
#include <sys/types.h>
#include <unistd.h>

int main()
{
    printf("hello from parent\n");
    daemon(0,0);

    // child
    pause();
    return 0;
}
