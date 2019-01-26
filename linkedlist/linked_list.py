class SingleNode:
    """单链表结构节点"""

    def __int__(self, data: int=0, next: SingleNode=None):
        self.__data = data
        self.__next = next

    @property
    def data(self) -> int:
        return self.__data

    @data.setter
    def data(self, data: int):
        self.__data = data

    @property
    def next(self) -> SingleNode:
        return self.__next

    @next.setter
    def next(self, next: SingleNode):
        self.__next = next


class BoubleNode:
    """双链表结构节点"""

    def __init__(self, data: int=0, prev: BoubleNode=None,
                 next: BoubleNode=None):
        self.__data = data
        self.__next = next
        self.__prev = prev

    @property
    def data(self) -> int:
        return self.__data

    @data.setter
    def data(self, data: int):
        self.__data = data

    @property
    def next(self) -> BoubleNode:
        return self.__next

    @next.setter
    def next(self, next: BoubleNode):
        self.__next = next

    @property
    def prev(self) -> BoubleNode:
        return self.__prev

    @prev.setter
    def prev(self, prev: BoubleNode):
        self.__prev = BoubleNode


class SingleLinkedList:
    """带头结点的单链表"""

    def __init__(self):
        self.length = 0
        self.__head = SingleNode()

    def insert_before(self, node: SingleNode, data: int) -> bool:
        """时间复杂度为o(n)"""
        if node is None or node == self.__head:
            return False
        prev = self.__head
        cur = prev.next
        while cur is not None:
            if cur == node:
                break
            prev = cur
            cur = cur.next
        if cur is None:
            return False
        new_node = SingleNode(data)
        new_node.next = cur
        prev.next = new_node
        self.length += 1
        return True

    def insert_after(self, node: SingleNode, data: int) -> bool:
        """时间复杂度为o(1)"""
        if node is None:
            return False
        new_node = SingleNode(data)
        new_node.next = node.next
        node.next = new_node
        self.length += 1
        return True

    def insert_to_head(self, data: int) -> bool:
        pass

    def insert_to_tail(self, data: int) -> bool:
        pass

    def find_by_index(self, index: int) -> SingleNode:
        pass

    def delete_node(self, node: SingleNode) -> bool:
        pass

    def reserve_self(self):
        pass

    def print(self):
        pass


class BoubleLinkedList:
    """带头结点的双链表"""

    def __int__(self):
        self.length = 0
        self.__head = BoubleNode()
