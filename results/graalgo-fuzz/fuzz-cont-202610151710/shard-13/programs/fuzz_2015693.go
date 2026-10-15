package main

import (
    "fmt"
)


func main() {
    var v1 string = "hello"
    var v2 float32 = 3.14159
    var v3 int = 2147483647
    var v4 bool = (0 != -932)
    var v5 string = ("\n\t" + "a")
    var v6 uint16 = 42
    v7 := []int{1, v3}
    v1 = ("a" + (("hello" + v1) + (v5 + v1)))
    v5 = v1
    v2 += v2
    var v8 uint8 = 137
    v1 = "a"
    var v9 uint8 = (v8 % (((^(v8))) | 1))
    switch v5 {
    case "foo bar":
        v9 = v8
    }
    v4 = false
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
