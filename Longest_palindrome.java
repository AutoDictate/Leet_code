public class Longest_palindrome {
    public static void main(String[] args) {
        String s = longestPalindrome("cbbd");
        System.out.println(s);
    }

    public static String longestPalindrome(String s) {
        if(s.length()>1){
            String str="";
            String ans = "";
            int max = 0;
            for (int i = 0; i < s.length()-1; i++) {
                str +=s.charAt(i);
                for (int j = i+1; j < s.length(); j++) {
                    if(isPalindrome(str) && (str.length()>max)){
                        max = str.length();
                        ans = str;
                    }
                    str += ""+s.charAt(j);
                }
                str ="";
            }
            if(ans.isEmpty()){
                return ""+s.charAt(0);
            }
            return ans;
        }
        return s;
    }

    public static boolean isPalindrome(String s){
        return s.contentEquals(new StringBuilder(s).reverse());
    }
}
