import java.util.Arrays;

public class Zigzag_conversion {
    public static void main(String[] args) {
        convert("PAYPALISHIRING",3);
    }

    public static String convert(String s, int numRows) {
        int row = numRows, col = s.length()/2;
        String[][] ans = new String[row][col];
        int k =0, t = row/2;
        int r =0;
        for (int i = r; i < row; i++) {
            for (int j = 0; j < col; j++) {
                ans[i][j] = ""+s.charAt(k);
                ++k;
                System.out.print(ans[j][i]);
            }
            System.out.println();
            System.out.println(Arrays.deepToString(ans));
        }
        return "";
    }
}