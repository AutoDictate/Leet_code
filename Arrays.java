public class Arrays {
    public static void main(String[] args) {
        int[] arr = new int[3];
        String[] str = new String[4];
        Boolean isCorrect= ! Boolean.FALSE;

        arr[2] = 12;
        change(arr);
        System.out.println(arr[2]);
        System.out.println(str[3]);
        System.out.println(isCorrect);
    }

    public static void change(int[] arr){
        arr[2] = 2000;
    }
}
