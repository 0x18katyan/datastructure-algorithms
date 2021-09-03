def merge_sort(arr):
    
    if len(arr) > 1:

        mid = len(arr) // 2
        L = arr[:mid]
        R = arr[mid:]


        merge_sort(L)
        merge_sort(R)        
    
        i = j = k = 0
        while i < len(L) and j < len(R):
            print("arr is", arr)
            print("i is {}, j is {} , k is {}".format(i, j , k))

            if L[i] < R[j]:
                arr[k] = L[i]
                i += 1
            else:
                arr[k] = R[j]
                j += 1                
            k += 1
        
        # print("arr is", arr)
        # print("i is {}, j is {} , k is {}".format(i, j , k))

        while i < len(L):
            arr[k] = L[i]
            k += 1
            i += 1
        
        while j < len(R):
            arr[k] = R[j]
            k += 1
            j += 1


arr = [123, 45, 1, 2,67, 8, 10, 1,456,234,456,9]

print("Length of original arr is", len(arr))

merge_sort(arr)

print("Length of sorted arr is", len(arr))
print("Sorted Array: ", arr)