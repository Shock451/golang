/*

Maps

A map is an unordered collection of key-value pairs. Also known as an associative array, a hash table or a dictionary, maps are used to look up a value by its associated key. Here's an example of a map in Go:

var x map[string]int

The map type is represented by the keyword map, followed by the key type in brackets and finally the value type. If you were to read this out loud you would say “x is a map of strings to ints.”

Like arrays and slices maps can be accessed using brackets. Try running the following program:

var x map[string]int
x["key"] = 10
fmt.Println(x)

You should see an error similar to this:

panic: runtime error: assignment to entry in nil map

goroutine 1 [running]:
main.main()
main.go:7 +0x4d

goroutine 2 [syscall]:
created by runtime.main
      C:/Users/ADMINI~1/AppData/Local/Temp/2/bindi
t269497170/src/pkg/runtime/proc.c:221
exit status 2

Up till now we have only seen compile-time errors. This is an example of a runtime error. As the name would imply, runtime errors happen when you run the program, while compile-time errors happen when you try to compile the program.

The problem with our program is that maps have to be initialized before they can be used. We should have written this:

x := make(map[string]int)
x["key"] = 10
fmt.Println(x["key"])

If you run this program you should see 10 displayed. The statement x["key"] = 10 is similar to what we saw with arrays but the key, instead of being an integer, is a string because the map's key type is string. We can also create maps with a key type of int:

x := make(map[int]int)
x[1] = 10
fmt.Println(x[1])

This looks very much like an array but there are a few differences. First the length of a map (found by doing len(x)) can change as we add new items to it. When first created it has a length of 0, after x[1] = 10 it has a length of 1. Second maps are not sequential. We have x[1], and with an array that would imply there must be an x[0], but maps don't have this requirement.

*/
