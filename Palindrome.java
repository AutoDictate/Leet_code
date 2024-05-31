import java.util.Scanner;

public class Palindrome {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        int n = in.nextInt();

        System.out.println(isPalindrome(n));
    }

    public static boolean isPalindrome(int n){
        StringBuilder sb = new StringBuilder();
        sb.append(n);
        String s = sb.reverse().toString();
        String ss = ""+n;

        if(s.equals(ss)){
            return true;
        }
        return false;
    }
}
