package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }

func main() {
    v1 := make(map[uint32]int)
    var v2 int64 = -146
    v3 := make(map[uint32]uint32)
    v4 := []uint32{2147483648, 65535, 65535, 2147483648}
    var v5 bool = true
    v6 := make(map[int]uint32)
    v5 = v5
    v7 := GenS1{F0: 2}
    var v8 uint = 299
    fmt.Println(v2, v8, v5)
    v2 = v2
    v5 = (18446744073709551615 < 0)
    v5 = (!true)
    v9 := v1[42]
    var v10 uint = 18446744073709551615
    v4 = append(v4, 832)
    if false {
        if len(v4) > 0 {
            v4[0] = 521
        }
    }
    var v11 string = "a"
    v6[(v9 >> (((v10 % ((v8) | 1))) & 63))] = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
