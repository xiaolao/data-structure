package linkedlist;

/**
 * common methods must contain in list.
 */
public interface ListInterface<ListNode> {

    public ListNode findByIndex(int index);

    public void insertToHead(V value);

    public void insertToTail(V value);

    public void insertBefore(ListNode node, V value);

    public void insertAfter(ListNode node, V value);
}
