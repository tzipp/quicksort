list_one = [99, 2, 32, 4, 1, 5, 8, 31]
list_two = [99, 2, 32]
def quicksort(numbers): # list of numbers to sort
    if len(numbers) == 0 or len(numbers) == 1:
        return numbers
    elif len(numbers) == 2:
        if numbers[0] > numbers[1]:
            numbers[0], numbers[1] = numbers[1], numbers[0]
        return numbers 
    else:
        pivot = numbers[0]
        s_nums = []
        g_nums = []

        for num in numbers:
            if num < pivot:
                s_nums.append(num)
            elif num > pivot: 
                g_nums.append(num)        
        s_nums.append(pivot)
        return quicksort(s_nums) + quicksort(g_nums) 

print "Program Start" 
print quicksort(list_two)
print quicksort(list_one) 
        

