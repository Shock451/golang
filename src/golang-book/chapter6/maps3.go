/*
Maps are also often used to store general information. Let's modify our program so that instead of just storing the name of the element we store its standard state (state at room temperature) as well:

func main() {
  elements := map[string]map[string]string{
    "H": map[string]string{
      "name":"Hydrogen",
      "state":"gas",
    },
    "He": map[string]string{
      "name":"Helium",
      "state":"gas",
    },
    "Li": map[string]string{
      "name":"Lithium",
      "state":"solid",
    },
    "Be": map[string]string{
      "name":"Beryllium",
      "state":"solid",
    },
    "B":  map[string]string{
      "name":"Boron",
      "state":"solid",
    },
    "C":  map[string]string{
      "name":"Carbon",
      "state":"solid",
    },
    "N":  map[string]string{
      "name":"Nitrogen",
      "state":"gas",
    },
    "O":  map[string]string{
      "name":"Oxygen",
      "state":"gas",
    },
    "F":  map[string]string{
      "name":"Fluorine",
      "state":"gas",
    },
    "Ne":  map[string]string{
      "name":"Neon",
      "state":"gas",
    },
  }

  if el, ok := elements["Li"]; ok {
    fmt.Println(el["name"], el["state"])
  }
}

Notice that the type of our map has changed from map[string]string to map[string]map[string]string. We now have a map of strings to maps of strings to strings. The outer map is used as a lookup table based on the element's symbol, while the inner maps are used to store general information about the elements. Although maps are often used like this, in chapter 9 we will see a better way to store structured information.

*/
