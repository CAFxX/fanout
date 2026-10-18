package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint }

func main() {
    v1 := []int{32768}
    var v2 bool = true
    v3 := []string{"a", "hello", ""}
    var v4 uint8 = 184
    v5 := make(map[uint32]int)
    v4 += (v4 >> ((255) & 63))
    if (!v2) {
        var v6 uint8 = (v4 ^ 255)
        v7 := GenS1{F0: v6}
        var v8 uint64 = 333
        _ = v6
        _ = v7
        _ = v8
    }
    var v9 uint64 = 53
    fmt.Println(v4)
    v9 = v9
    var v10 uint32 = 1
    v2 = v2
    v11 := GenS1{F0: 128}
    v2 = true
    _ = len(v5)
    v10 = (v10 * (^(v10)))
    switch ("" + "hello") {
    case "":
        var v12 uint32 = v10
        _ = v12
    case "\n\t":
        var v13 uint64 = v9
        _ = v13
    }
    var v14 uint16 = 0
    if v2 {
        fmt.Println(v4, v14)
    } else {
        v15 := GenS2{F0: 281}
        _ = v15
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
