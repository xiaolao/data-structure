from abc import ABCMeta, abstractmethod


class Array:

    def __init__(self, capacity: int):
        self.capacity = capacity
        # 模仿C中预先给数组分配固定大小内存
        self._data = [0] * capacity

    def __getitem__(self, index: int) -> object:
        return self._data[index]

    def __setitem__(self, index: int, value: object):
        self._data[index] = value

    def __len__(self) -> int:
        return len(self._data)

    def __iter__(self):
        for item in self._data:
            yield item

    def find(self, index: int) -> object:
        try:
            return self._data[index]
        except IndexError:
            raise Exception("index out of range")

    def insert(self, index: int, value: object) -> object:
        if index >= self.capacity or index < 0:
            raise Exception("index out of range")
        self._data.insert(index, value)
        return value

    def delete(self, index: int) -> object:
        if index >= self.capacity or index < 0:
            raise Exception("index out of range")
        self._data += [0]
        return self._data.pop(index)

    def print(self):
        for item in self:
            print(item)
