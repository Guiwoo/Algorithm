#include <stdio.h>
#include <stdlib.h>

int compare(const void *a,const void *b) {
	if(*(int*)a>*(int*)b) return 1;
	else return -1;
}

int main(void){
	int arr[9];
	int total = 0;

	for (int i=0; i<9; i++) {
		scanf("%d",&arr[i]);
		total += arr[i];
	}

	qsort(arr,9,sizeof(arr[0]),compare);

	for (int i=0; i<9; i++) {
		for (int j=i+1; j<9; j++){
			if (total - arr[i] - arr[j] == 100) {
				for (int k=0; k<9; k++){
					if (k == i || k == j){
						continue;
					}
					printf("%d\n",arr[k]);
				}
				return 0;
			}
		}
	}

	return 0;	
}
