import java.util.Collections;
import java.util.LinkedHashMap;
import java.util.Scanner;

public class Max_frequency {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        int n  = in.nextInt();
        int[] arr = new int[n];
        for (int i = 0; i < n; i++) arr[i] = in.nextInt();
        System.out.println("sss  " +maxFrequency(arr));
    }

    public static int maxFrequency(int[] arr){
        System.out.println("Function Calling");
        int max =1,count=0;
        LinkedHashMap<Integer,Integer> ar = new LinkedHashMap<>();
        for (int i = 0; i < arr.length; i++) {
            if(ar.containsKey(arr[i])){
                int t = ar.get(arr[i])+1;
                if(t>=max){
                    max = t;
                    System.out.println("m "+max);
                }
                ar.put(arr[i],t);
            }else{
                ar.put(arr[i],1);
            }
        }
        System.out.println("Max value : "+max);
        for (int n : arr){
            if(ar.get(n)>=max){
                count++;
            }
        }
        return count;
    }
}
