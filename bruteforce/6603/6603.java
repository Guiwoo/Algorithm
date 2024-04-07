import java.io.*;
import java.util.*;
import java.lang.*;

public class Main {
    static StringBuilder sb = new StringBuilder();

	public static void main(String[] args) throws IOException {
		BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
		System.out.println("run");;;;

		while (true) {
			String[] arr = br.readLine().split(" ");
			if (arr.length == 1) {
				System.out.println("hit");
				break;
			}
			String[] output = new String[6];
			permutation(arr,output,1,0,0|(0<<1));
			sb.append("\n");
		}

		System.out.println(sb.toString());
	}

	public static void permutation(String[] arr,String[] output, int start,int depth,int flag){
	    if (depth >= 6) {
	         for (int i=0;i<output.length;i++){
	            sb.append(output[i]+" ");
	         }
	         sb.append("\n");
	        return;
	    }else{
			for (int i=start;i<arr.length;i++){ 
				if ((flag & 1<<i) != 0) {
					continue;
				}else{
					output[depth] = arr[i];
					permutation(arr,output,i+1,depth+1,flag | (1<<i));
				}
			}
		}
	    
	}
}

