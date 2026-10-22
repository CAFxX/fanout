package main

import (
    "fmt"
)


func main() {
    var v1 int16 = 1
    v2 := make(map[int]int)
    v3 := make(map[uint32]string)
    v1 = v1
    if (2147483647 >= 333) {
        fmt.Println(v1)
        v1 = 32767
    }
    var v4 string = "a"
    v5 := []int{-43}
    v4 = (((("foo bar" + (v4 + "foo bar")) + "a") + v4) + v4)
    var v6 uint16 = 209
    v6 -= (v6 | (v6 >> ((4294967295) & 63)))
    var v7 uint16 = (^(v6))
    var v8 string = "a"
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
