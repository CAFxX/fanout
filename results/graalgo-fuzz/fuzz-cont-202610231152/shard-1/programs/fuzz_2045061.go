package main

import (
    "fmt"
)

type GenS1 struct { F0 int }

func main() {
    v1 := []int{2147483647, 4294967295, 127, 266}
    v2 := make(map[int]string)
    var v3 bool = ((13 > 0) && true)
    var v4 uint = 18446744073709551615
    var v5 bool = ("" < (("a" + "foo bar") + ""))
    if v3 {
        v3 = (42 == -128)
        v4 -= v4
    }
    var v6 int64 = -1
    var v7 uint8 = uint8(v4)
    v3 = (((v6 << (((v4 ^ 2)) & 63)) <= v6) || (-275 <= 2147483647))
    if false {
        v7 = 42
        v8 := GenS1{F0: 42}
        fmt.Println(v4, v3, v4)
        _ = v8
    } else {
        v2[-2147483648] = ("foo bar" + "hello")
    }
    fmt.Println(v3)
    _ = len(v2)
    v9 := GenS1{F0: -294}
    v5 = v3
    var v10 int8 = 0
    var v11 uint64 = 150
    fmt.Println(v10, v3, v10)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v9
}
