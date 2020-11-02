#include <pthread.h>
#include <stdio.h>

void *hello_world(void *args) {
    int *n = args;
    printf("Hello from thread %d\n", *n);
    pthread_exit(NULL);
}

int main(int argc, char const *argv[]) {
    pthread_t thread[5];
    int nums[5];
    for(int n = 0; n < 5; n++){
        nums[n] = n + 1;
        if (pthread_create(&thread[n], NULL, hello_world, &nums[n])) {
            printf("Error creating thread\n");
        }
    }
    for(int n = 0; n < 5; n++){
        if (pthread_join(thread[n], NULL)) {
            printf("Error joining thread\n");
        }
    }
}
