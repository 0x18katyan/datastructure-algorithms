def reverse(string: str = "abcdefgh") -> str:

    if isinstance(string, str) == False:
        print("Input is not a string")
        return 
    elif len(string) < 2:
        print("Input is too short")
        return 
    rstring = ""

    for i in range(len(string) - 1, -1, -1):
        rstring += string[i]
    
    print("SomeString is {} and ReversedString is {}".format(somestring, rstring))
    
    return 

# somestring = "abcdefgh"
somestring = 1

reverse(somestring)

