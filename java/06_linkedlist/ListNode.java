package linkedlist;

/**
 * node for any type.
 */
public class ListNode<V> {

    // the value of the node.
    private V value;

    // indicate the next node.
    private ListNode<V> next;

    public ListNode(V value, ListNode<V> next) {
        this.next = next;
        this.value = value;
    }

    public V getValue() {
        return value;
    }

    public ListNode<V> getNext() {
        return next;
    }
}
