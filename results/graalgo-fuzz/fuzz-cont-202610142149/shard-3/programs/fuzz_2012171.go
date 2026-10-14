package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 uint16; F1 uint64; F2 string }
type GenS3 struct { F0 int8; F1 int }
type GenS4 struct { F0 string }
func genFn0(p0 int8, p1 uint64, p2 string) int64 {
    var v4 uint8 = 42
    var v5 int8 = (p0 | p0)
    v4 -= (v4 << ((174) & 63))
    var v6 int16 = 1
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return 42
}

func main() {
    v1 := []int{42, 4294967295}
    var v2 uint32 = 2
    var v3 int32 = 1
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if false {
        _ = len(v1)
        v7 := GenS1{F0: 79}
        _ = v7
    } else {
        v2 -= (^(v2))
        fmt.Println(v3, v2)
    }
    fmt.Println(v2, v3)
    v3 = 0
    fmt.Println(v3, v3)
    var v8 float32 = 3.14159
    v1 = append(v1, 884)
    fmt.Println(v8)
    v9 := GenS2{F0: 1, F1: 42, F2: ""}
    for fi10 := 0; fi10 < 7; fi10++ {
        var v11 uint = 0
        v12 := GenS3{F0: 31, F1: -379}
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 string = "hello"
    v8 = 3.14159
    fmt.Println(v13, v8, v3)
    v14 := GenS4{F0: v13}
    if true {
        v2 = v2
        for fi15 := 0; fi15 < 9; fi15++ {
            if len(v1) > 0 {
                v1[0] = fi15
            }
            var v16 uint = 0
            _ = fi15
            _ = v16
        }
    } else {
        if (3.14159 > 1.5) {
            v2 = v2
        }
        var v17 int = -222
        var v18 uint16 = 310
        _ = v17
        _ = v18
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
