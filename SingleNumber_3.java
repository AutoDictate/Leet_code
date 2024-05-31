import java.util.ArrayList;
import java.util.Arrays;
import java.util.LinkedHashMap;
import java.util.List;

public class SingleNumber_3 {
    public static void main(String[] args) {
        int[] nums = {1,2,1,3,2,5};
        System.out.println(Arrays.toString(singleNumber(nums)));
    }
    public static int[] singleNumber(int[] nums) {
        LinkedHashMap<Integer,Integer> h1 = new LinkedHashMap<>();
        for (int num : nums) {
            if (h1.containsKey(num)) {
                h1.put(num, h1.get(num) + 1);
            } else {
                h1.put(num, 1);
            }
        }
        List<Integer> l = new ArrayList<>();
        for(int i : h1.keySet()){
            if(h1.get(i)==1){
                l.add(i);
            }
        }
        int[] ans = new int[l.size()];
        for (int i =0;i< ans.length;i++) ans[i] = l.get(i);
        return ans;
    }
}
