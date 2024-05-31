import java.util.LinkedHashMap;

class Longest_substr {
    public static void main(String[] args) {
        int n = lengthOfLongestSubstring("bbbbb");
        System.out.println(n);
    }
    public static int lengthOfLongestSubstring(String s) {

        int max =0;
        for(int i=0;i<s.length();i++){
            int t = substr(i,s);
            if(t>=max) max=t;
        }
        return max;
    }

    public static int substr(int j,String s){
        LinkedHashMap<Character,Integer> h1 = new LinkedHashMap<>();
        int count =0;
        for(int i = j;i<s.length();i++){
            if(! h1.containsKey(s.charAt(i))){
                h1.put(s.charAt(i),1);
                ++count;
            }else{
                return count;
            }
        }
        return count;
    }
}
