package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
func genFn0(p0 uint8, p1 int32, p2 uint16) uint16 {
    var v4 uint32 = 42
    v4 += v4
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    return p2
}
func genFn1() uint {
    v5 := []int{2147483647, -936, -2147483648}
    if true {
        var v6 int32 = 0
        _ = v6
    }
    _ = v5
    return 385
}

func main() {
    var v1 int = -889
    var v2 uint16 = 32767
    var v3 int32 = 681
    v7 := []int{v1}
    v8 := GenS1{F0: int8(v2)}
    _ = len(v7)
    var v9 uint32 = 65535
    if (v1 != v1) {
        if (!(127 >= 1)) {
            _ = len(v7)
        }
        v7 = append(v7, v1)
        v3 += v3
    }
    fmt.Println(v9)
    var v10 int64 = -1
    var v11 uint8 = 128
    switch v1 {
    case 32768:
        var v12 uint16 = v2
        _ = v12
    default:
        v1 += (v1 >> ((0) & 63))
    }
    if len(v7) > 0 {
        v7[0] = 2147483647
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
