import java.util.Scanner;

public class Two_sum {
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        int arr_size = in.nextInt();
        int[] nums = new int[arr_size];
        for (int i = 0; i < nums.length; i++) {
            nums[i] = in.nextInt();
        }
        int target = in.nextInt();
        int[] final_answer = twoSum(nums,target);
    }

    public static int[] twoSum(int[] nums, int target){
        return checkSum(0,0,nums,target);
    }

    public static int[] checkSum(int i,int j,int[] nums, int target){
        if(i!= nums.length){
            if(j!= nums.length && i!=j){
                if(nums[i]+nums[j] == target){
                    System.out.println(nums[i]+" and "+nums[j]);
                    return new int[]{i,j};
                }
                return checkSum(i,j+1,nums,target);
            }
            return checkSum(i+1,0,nums,target);
        }
        return new int[]{};
    }
}
