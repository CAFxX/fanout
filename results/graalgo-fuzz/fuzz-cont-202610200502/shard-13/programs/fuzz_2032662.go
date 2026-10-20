package main

import (
    "fmt"
)


func main() {
    v1 := []string{("a" + "\n\t"), ("" + "a"), ""}
    var v2 string = (("hello" + ("\n\t" + "")) + ("foo bar" + ("hello" + "hello")))
    var v3 int32 = 996
    v4 := make(map[int]string)
    delete(v4, 643)
    var v5 int32 = 1
    fmt.Println(v2, v2, v3)
    _ = len(v1)
    fmt.Println(v3, v5)
    fmt.Println(v5, v3, v3)
    var v6 int32 = v5
    fmt.Println(v6, v2)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
