class array(object):

    def __init__(self):

        self.arr = {} 
        self.size = 0

    def add(self, value):
        self.arr[self.size] = value
        self.size += 1

    def pop(self):
        del self.arr[self.size - 1]
        self.size -= 1

    def delete(self, index):
        for i in range(index, self.size - 1):
            self.arr[i] = self.arr[i + 1]
        del self.arr[self.size - 1]
        self.size -= 1

arr = array()

arr.add("a")
arr.add("b")
arr.add("c")
arr.add('d')
arr.add('e')
arr.add('f')

print(arr.arr)

arr.pop()
print(arr.arr)

arr.delete(1)
print(arr.arr)