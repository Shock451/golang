/*

We can also delete items from a map using the built-in delete function:

delete(x, 1)

Let's look at an example program that uses a map:

*/

package main

import "fmt"

func main() {
  elements := make(map[string]string)
  elements["H"] = "Hydrogen"
  elements["He"] = "Helium"
  elements["Li"] = "Lithium"
  elements["Be"] = "Beryllium"
  elements["B"] = "Boron"
  elements["C"] = "Carbon"
  elements["N"] = "Nitrogen"
  elements["O"] = "Oxygen"
  elements["F"] = "Fluorine"
  elements["Ne"] = "Neon"

  fmt.Println(elements["Li"])


/*

elements is a map that represents the first 10 chemical elements indexed by their symbol. This is a very common way of using maps: as a lookup table or a dictionary. Suppose we tried to look up an element that doesn't exist:

fmt.Println(elements["Un"])

If you run this you should see nothing returned. Technically a map returns the zero value for the value type (which for strings is the empty string). Although we could check for the zero value in a condition (elements["Un"] == "") Go provides a better way:

name, ok := elements["Un"]
fmt.Println(name, ok)

Accessing an element of a map can return two values instead of just one. The first value is the result of the lookup, the second tells us whether or not the lookup was successful. In Go we often see code like this:
*/
        if name, ok := elements["Un"]; ok {
                fmt.Println(name, ok)
        }
}
/*
First we try to get the value from the map, then if it's successful we run the code inside of the block.

Like we saw with arrays there is also a shorter way to create maps:

elements := map[string]string{
  "H":  "Hydrogen",
  "He": "Helium",
  "Li": "Lithium",
  "Be": "Beryllium",
  "B":  "Boron",
  "C":  "Carbon",
  "N":  "Nitrogen",
  "O":  "Oxygen",
  "F":  "Fluorine",
  "Ne": "Neon",
}

*/
