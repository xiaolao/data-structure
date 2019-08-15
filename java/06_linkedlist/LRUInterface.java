package linkedlist;

/**
 * interface of common linkedlist
 */
public interface LRUInterface<K, V> {

    // get the last use element in lru object.
    public V get(K key, V value);

    // push the element in lru.
    public void put(K key, V value);

}
