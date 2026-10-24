package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint; F1 int8; F2 int32 }
type GenS3 struct { F0 bool; F1 uint32 }
type GenS4 struct { F0 int64 }
type GenS5 struct { F0 uint64 }
func genFn0(p0 int, p1 uint, p2 string) uint {
    p2 = p2
    switch (^(p0)) {
    case 127:
        fmt.Println(p0)
    default:
        var v4 uint8 = 1
        _ = v4
    }
    if (4294967295 <= 2) {
        p2 = ((((p2 + (p2 + "\n\t")) + (p2 + p2)) + p2) + "hello")
        v5 := []int{}
        p0 -= p0
        _ = v5
    }
    _ = p0
    _ = p1
    _ = p2
    return p1
}
func genFn1(p0 uint64, p1 bool) uint8 {
    p0 += p0
    if p1 {
        v6 := GenS1{F0: 2}
        fmt.Println(p1)
        _ = v6
    } else {
        var v7 int64 = 1
        var v8 uint64 = p0
        v9 := GenS2{F0: 65535, F1: 1, F2: -813}
        _ = v7
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    return 248
}

func main() {
    var v1 int16 = -976
    v2 := make(map[int]int)
    var v3 int8 = 0
    fmt.Println(v3)
    var v10 uint32 = 2
    v11 := GenS3{F0: true, F1: (^(v10))}
    var v12 int64 = int64(v10)
    if false {
        if false {
            v13 := GenS2{F0: 1, F1: -1, F2: -1}
            _ = v13
        }
    }
    var v14 uint = 297
    v14 -= genFn0(-397, v14, "foo bar")
    v15 := GenS4{F0: (v12 & int64(v10))}
    for fi16 := 0; fi16 < 13; fi16++ {
        if (false && (!false)) {
            v17 := GenS4{F0: v12}
            v18 := GenS5{F0: 2}
            _ = v17
            _ = v18
        }
        v19 := []int{(fi16 >> (((v14 + (^(v14)))) & 63)), (fi16 & (fi16 >> ((v14) & 63)))}
        _ = fi16
        _ = v19
    }
    var v20 string = "a"
    _ = len(v2)
    v21 := GenS4{F0: v12}
    v22 := v2[128]
    v23 := GenS2{F0: (v14 / ((923) | 1)), F1: v3, F2: 1}
    v24 := []int{(v22 ^ 32767), 1}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v24
    _ = v3
}
