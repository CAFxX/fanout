package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int16 }
func genFn0(p0 int32, p1 uint, p2 int16) uint {
    var v5 int = -413
    var v6 uint64 = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return p1
}

func main() {
    var v1 bool = (230 < 249)
    var v2 int16 = -32768
    v3 := make(map[int]uint32)
    v4 := []string{}
    v7 := v3[128]
    var v8 int8 = 42
    var v9 uint8 = 0
    v7 += (v7 | 4294967295)
    v8 -= 42
    v10 := GenS1{F0: 1}
    var v11 int32 = 948
    if (false || v1) {
        if v1 {
            v9 -= v9
            var v12 int = 42
            v11 += v11
            _ = v12
        } else {
            v13 := GenS1{F0: v11}
            var v14 uint8 = v9
            v7 = v7
            _ = v13
            _ = v14
        }
        var v15 int16 = (v2 | (v2 >> ((801) & 63)))
        _ = v15
    }
    fmt.Println(v2, v7, v11)
    v2 = v2
    var v16 int16 = (v2 << ((4294967295) & 63))
    if (1 >= 42) {
        v9 += (v9 >> ((genFn0(int32(v16), uint(v9), int16(v11))) & 63))
        v17 := GenS2{F0: -1}
        v16 -= v2
        _ = v17
    }
    switch "foo bar" {
    case "a":
        var v18 uint = genFn0(99, 2, (v16 << ((780) & 63)))
        _ = v18
    case "foo bar":
        fmt.Println(v9, v8)
    }
    fmt.Println(v8)
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
    _ = v9
}
