package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 int16 }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 bool }
func genFn0(p0 uint32, p1 uint16, p2 int32) int64 {
    if (!true) {
        var v6 int16 = 0
        _ = v6
    }
    var v7 uint = 65535
    fmt.Println(p0, p1)
    p1 = (p1 - 573)
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 1
}

func main() {
    var v1 uint32 = 0
    var v2 int32 = 2147483647
    var v3 uint32 = (^(v1))
    var v4 int16 = 42
    v5 := make(map[string]bool)
    v8 := []int{int(v4), 98, 128}
    var v9 uint16 = 0
    v2 = (v2 << ((65535) & 63))
    if true {
        v10 := GenS1{F0: (v2 | v2), F1: 711, F2: 80}
        v9 = 1
        if false {
            v4 = (v4 >> ((939) & 63))
        }
        _ = v10
    } else {
        v1 += v3
    }
    if (524 < 81) {
        v11 := v5["foo bar"]
        v12 := GenS2{F0: uint8(v3), F1: -1}
        _ = v11
        _ = v12
    }
    var v13 int64 = 286
    v14 := GenS1{F0: v2, F1: v4, F2: 0}
    v15 := GenS3{F0: true}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
