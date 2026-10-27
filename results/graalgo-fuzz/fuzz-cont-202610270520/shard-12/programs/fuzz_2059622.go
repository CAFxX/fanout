package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 int32; F1 uint64; F2 uint32 }
type GenS3 struct { F0 bool; F1 int16; F2 int16 }
type GenS4 struct { F0 uint32; F1 int8 }
type GenS5 struct { F0 int16 }
func genFn0(p0 bool, p1 uint16) int8 {
    fmt.Println(p1)
    v4 := GenS1{F0: 42, F1: 2147483648}
    _ = p0
    _ = p1
    _ = v4
    return 100
}
func genFn1(p0 bool, p1 string, p2 string) uint16 {
    p1 = ("hello" + "\n\t")
    switch p2 {
    case "hello":
        v5 := []int{0, 42, 702, 0}
        _ = v5
    case "foo bar":
        fmt.Println(p2)
    case "\n\t":
        switch 127 {
        case 0:
            v6 := GenS2{F0: -84, F1: 451, F2: 42}
            _ = v6
        case 38:
            p0 = ("" != "foo bar")
        case 127:
            fmt.Println(p2, p0, p0)
        }
    default:
        p0 = false
    }
    _ = p0
    _ = p1
    _ = p2
    return 2
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 uint64 = 1
    v3 := make(map[uint32]uint32)
    v7 := []int{1, -9223372036854775808, 0, -918}
    v8 := GenS3{F0: true, F1: 0, F2: -798}
    v9 := GenS4{F0: 2, F1: genFn0((!(510 <= 0)), 248)}
    var v10 uint64 = v2
    for fi11 := 0; fi11 < 4; fi11++ {
        var v12 int64 = 622
        v13 := GenS3{F0: false, F1: 1, F2: 32767}
        _ = fi11
        _ = v12
        _ = v13
    }
    var v14 float64 = 0.0
    var v15 int32 = 313
    v10 -= v10
    if false {
        v15 += -1
        v15 -= 0
        v16 := GenS5{F0: 0}
        _ = v16
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v7:", len(v7), v7)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v9
}
