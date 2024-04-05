import java.io.*;
import java.util.*;

public class Main {
	public static void main(String[] args) throws IOException {
		BufferedReader br = new BufferedReader(new InputStreamReader(System.in));
		while (true) {
			String[] arr = br.readLine().split(" ");
			if (arr.length == 1) {
				break;
			}
			System.out.println(arr);
		}
	}
}

