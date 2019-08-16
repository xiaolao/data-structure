package linkedlist;

/**
 * 
 */
public class ListNode<E> {
    private E value;
    private ListNode<E> next;

    public ListNode(E value, ListNode<E> next) {
        this.next = next;
        this.value = value;
    }

    public E getValue() {
        return value;
    }

}

