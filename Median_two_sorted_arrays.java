import java.util.Arrays;
import java.util.Collections;
import java.util.LinkedList;

public class Median_two_sorted_arrays {
    public static void main(String[] args) {
        int[] a = {1,2};
        int[] b = {3,4};
        double ans = findMedianSortedArrays(a,b);
        System.out.println(ans);
    }
    public static double findMedianSortedArrays(int[] nums1, int[] nums2) {
        LinkedList<Integer> l1 = new LinkedList<>();
        for (int j : nums1) l1.add(j);
        for (int j : nums2) l1.add(j);

        Collections.sort(l1);
        System.out.println(Arrays.toString(l1.toArray()));
        if(l1.size()%2==0){
            return (double) (l1.get(l1.size() / 2) + l1.get((l1.size() / 2) - 1)) /2;
        }
        return (double)l1.get(l1.size()/2);
    }
}
