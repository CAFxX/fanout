package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    v1 := []int{4294967295, 0, -1, 9223372036854775807}
    v2 := []string{}
    v3 := []int{-605}
    var v4 uint64 = 0
    var v5 int64 = 9223372036854775807
    var v6 int = 519
    v4 = (v4 << ((759) & 63))
    var v7 uint = 2
    v8 := GenS1{F0: 1}
    v5 = v5
    var v9 float32 = -2.25
    v7 = (v7 ^ v7)
    switch (-(v6)) {
    case 1:
        if ((!false) || (!((8 == 2) || (644 > 986)))) {
            var v10 float32 = (((1e10 + float32(1.5)) - v9) * v9)
            _ = v10
        }
    case 9223372036854775807:
        var v11 string = "hello"
        _ = v11
    default:
        var v12 uint8 = 239
        _ = v12
    }
    var v13 uint = (v7 | (v7 | v7))
    v14 := GenS1{F0: 1}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
