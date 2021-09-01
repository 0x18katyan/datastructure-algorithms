list1 = [0, 3, 4, 31]
list2 = [4, 6, 30]

def merge_sort(list1, list2):
    list3 = list1 + list2
    list3 = sorted(list3)
    return list3  

print(merge_sort(list1, list2))