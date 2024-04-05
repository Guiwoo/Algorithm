import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.Arrays;

public class Main {
    public static void main(String[] args) throws IOException {
        BufferedReader br = new BufferedReader(new InputStreamReader(System.in));

        int[] arr = new int[9];
        int total = 0;

        for (int i = 0; i < arr.length; i++) {
            int a = Integer.parseInt(br.readLine());
            arr[i] = a;
            total += a;
        }

        Arrays.sort(arr);

        for (int i = 0; i < (arr.length - 1); i++) {
            for (int j = i + 1; j < arr.length; j++) {
                if (total - arr[i] - arr[j] == 100) {
                    for (int k = 0; k < arr.length; k++) {
                        if (k == i || k == j) {
                            continue;
                        }
                        System.out.println(arr[k]);
                    }
                    return; // 결과 출력 후 프로그램 종료
                }
            }
        }

        br.close(); // BufferedReader 닫기
    }
}
