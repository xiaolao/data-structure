interface ArrayInterface<T> {
    public int length();
    public T find(int index);
    public boolean insert(int index, T value);
    public T delete(int index);
}


public class Array<T> {
    private T[] data;
    private int size;

    /**
     * capacity: 数组的容量
     */
    public Array(int capacity) {
        data = (T[]) new Object[capacity];
        size = 0;
    }

    // 获取数据容量
    public int length() {
        return size;
    }

    // 通过下标查找元素
    public T find(int index) {
        checkIndex(index);
        return data[index];
    }

    public boolean insert(int index, T value) {
        checkIndex(index);
        if (size == data.length) {
            System.out.println("full array");
            return false;
        }
        for (int i = data.length; i > index; i--) {
            data[i] = data[i-1];
        }
        data[index] = value;
        size++;
        return true;
    }

    public T delete(int index) {
        checkIndex(index);
        T value = data[index];
        for (int i = index; i < size-1; i++) {
            data[i] = data[i+1];
        }
        size--;
        return value;
    }

    private void checkIndex(int index) {
        if (index < 0 || index >= data.length) {
            throw new IllegalArgumentException("index is out of range.");
        }
    }

}
