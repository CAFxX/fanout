package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint16; F1 int; F2 int16 }
func genFn0(p0 uint32) bool {
    v6 := []int{-13}
    p0 += (p0 % ((p0) | 1))
    _ = p0
    _ = v6
    return (2 < 254)
}

func main() {
    var v1 string = "hello"
    var v2 uint16 = 984
    v3 := []uint32{}
    v4 := make(map[int]int)
    var v5 string = ("foo bar" + "\n\t")
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v5, v2, v1)
    v1 = "hello"
    var v7 uint32 = 42
    var v8 int16 = -442
    v1 = "\n\t"
    if (0 == uint64(v7)) {
        if (710 == 255) {
            v9 := GenS1{F0: true}
            _ = v9
        }
    } else {
        v2 = (v2 + 139)
    }
    v1 = "hello"
    v10 := GenS1{F0: true}
    switch 128 {
    case -2147483648:
        v11 := v4[798]
        _ = v11
    default:
        v12 := GenS2{F0: (v2 >> ((42) & 63)), F1: -483, F2: v8}
        _ = v12
    }
    var v13 uint16 = v2
    v5 = v5
    fmt.Println(v1, v13, v5)
    var v14 uint16 = (v13 << ((65535) & 63))
    v14 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
