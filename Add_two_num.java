import java.math.BigInteger;
import java.util.List;
import java.util.Scanner;
/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     int val;
 *     ListNode next;
 *     ListNode() {}
 *     ListNode(int val) { this.val = val; }
 *     ListNode(int val, ListNode next) { this.val = val; this.next = next; }
 * }
 */
class ListNode {
    int val;
    ListNode next;
    ListNode(){}
    ListNode(int val){this.val = val;}
    ListNode(int val, ListNode next){ this.val = val ; this.next = next;}
}
class Solution {
    public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
        StringBuilder s1 = new StringBuilder();
        StringBuilder s2 = new StringBuilder();

        while(l1 !=null){
            s1.append(l1.val);
            l1 = l1.next;
        }
        s1.reverse();
        while(l2 != null){
            s2.append(l2.val);
            l2 = l2.next;
        }
        s2.reverse();
        BigInteger sum;
        BigInteger b1 = new BigInteger(s1.toString());
        BigInteger b2 = new BigInteger(s2.toString());

        sum = b1.add(b2);
        String ans = sum.toString();

        ListNode head = null;
        ListNode tail = null;

        for (int i = ans.length() - 1; i >= 0; --i) {
            int temp = Integer.valueOf(""+ans.charAt(i));
            ListNode newNode = new ListNode(temp);

            if (head == null) {
                head = newNode;
                tail = newNode;
            } else {
                tail.next = newNode;
                tail = newNode;
            }
        }

        return head;
    }
}
