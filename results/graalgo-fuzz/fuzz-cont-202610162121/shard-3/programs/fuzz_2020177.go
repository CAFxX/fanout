package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint64; F2 uint16 }

func main() {
    var v1 bool = false
    v2 := make(map[uint32]int)
    var v3 int8 = -89
    var v4 string = ("a" + "hello")
    v1 = false
    v2[2147483647] = 32768
    var v5 int8 = 1
    var v6 float32 = float32(3.14159)
    var v7 string = v4
    v8 := GenS1{F0: 1, F1: 0, F2: 670}
    v1 = (0 >= 1)
    if v1 {
        switch "foo bar" {
        case "\n\t":
            v9 := GenS1{F0: 127, F1: 0, F2: 0}
            _ = v9
        }
        v3 = v3
        var v10 uint16 = 65535
        _ = v10
    }
    v4 = v7
    fmt.Println(v1, v6, v1)
    var v11 uint = 18446744073709551615
    v12 := []int{32768, 42}
    var v13 float64 = 1.5
    var v14 uint32 = 2147483648
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
