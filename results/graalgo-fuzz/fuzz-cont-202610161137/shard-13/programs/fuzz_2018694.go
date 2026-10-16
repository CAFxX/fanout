package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64 }
type GenS2 struct { F0 int8; F1 bool }
type GenS3 struct { F0 int; F1 uint64 }
func genFn0(p0 int64) int8 {
    fmt.Println(p0)
    fmt.Println(p0)
    var v5 uint16 = 32767
    var v6 uint16 = v5
    p0 -= p0
    _ = p0
    _ = v5
    _ = v6
    return -1
}
func genFn1(p0 string, p1 int64, p2 string) bool {
    p0 = p0
    var v7 uint32 = 4294967295
    v8 := GenS1{F0: p0, F1: 4294967295}
    fmt.Println(p1)
    v9 := []int{42}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return false
}

func main() {
    v1 := make(map[int]bool)
    var v2 string = ("a" + "\n\t")
    v3 := []uint32{2147483648}
    var v4 int16 = 32767
    v4 += (v4 << ((2) & 63))
    for fi10 := 0; fi10 < 1; fi10++ {
        v4 -= v4
        _ = fi10
    }
    switch v2 {
    case "\n\t":
        fmt.Println(v2, v4)
    default:
        if ((false && (true && (1 == 42))) && (-166 == 89)) {
            var v11 int16 = v4
            _ = v11
        }
    }
    if ((float32(((0.0 + 0.1) * (0.0 * 1.5))) > (float32((1.5 / 1e10)) + 1e10)) && (!true)) {
        v12 := GenS2{F0: -103, F1: false}
        v4 += v4
        _ = v12
    } else {
        v13 := GenS3{F0: 127, F1: 450}
        _ = v13
    }
    if false {
        v14 := GenS3{F0: -1, F1: 2}
        var v15 bool = (382 > 724)
        switch 0 {
        case 262:
            v15 = true
        default:
            fmt.Println(v15, v4, v4)
        }
        _ = v14
        _ = v15
    }
    for fi16 := 0; fi16 < 2; fi16++ {
        fmt.Println(v4)
        _ = fi16
    }
    fmt.Println(v2, v2)
    fmt.Println(v4, v4)
    var v17 float64 = (3.14159 * 1e10)
    fmt.Println(v4)
    v4 = (^(v4))
    if true {
        var v18 uint = 720
        _ = v18
    } else {
        v2 = v2
        if len(v3) > 0 {
            v3[0] = 2
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
