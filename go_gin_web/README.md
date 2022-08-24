gin-web脚手架







参考：https://blog.csdn.net/yhflyl/article/details/117635901

  public static ListNode reverseList(ListNode head) {
    ListNode pre = null;
    ListNode cur = head;
    ListNode tmp;
    while (cur != null) {
      tmp = cur.next;
      cur.next = pre;
      pre = cur;
      cur = tmp;
    }
    return pre;
  }
  
