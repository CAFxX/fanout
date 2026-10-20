package main

import (
    "fmt"
)


func main() {
    v1 := make(map[int]int)
    var v2 bool = false
    var v3 int = 9223372036854775807
    v4 := make(map[string]string)
    v5 := make(map[uint32]string)
    v2 = v2
    v3 += v3
    var v6 uint = 65535
    v3 -= (v3 ^ (v3 % (((-(v3))) | 1)))
    v5[588] = ("a" + ("" + (("a" + ((("foo bar" + "hello") + "a") + "a")) + "hello")))
    v7 := []int{v3, v3, v3, 716}
    v6 -= v6
    v6 += 323
    v2 = (-39 < 42)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
