package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16; F2 uint8 }
type GenS2 struct { F0 int; F1 int64 }
func genFn0(p0 int16, p1 int, p2 int64) uint64 {
    fmt.Println(p0, p0, p0)
    for fi6 := 0; fi6 < 4; fi6++ {
        fmt.Println(fi6, fi6)
        v7 := []int{}
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    return 376
}
func genFn1() int16 {
    var v8 int = -957
    for fi9 := 0; fi9 < 5; fi9++ {
        var v10 string = ""
        _ = fi9
        _ = v10
    }
    _ = v8
    return 32767
}

func main() {
    v1 := []uint32{2147483648, 2147483647, 1}
    v2 := make(map[int]int)
    var v3 int64 = -591
    var v4 bool = (817 != 427)
    var v5 bool = true
    v11 := GenS1{F0: -92, F1: 1, F2: 0}
    v3 -= -9223372036854775808
    fmt.Println(v4)
    v4 = false
    v12 := GenS2{F0: -325, F1: (v3 << ((2) & 63))}
    switch "\n\t" {
    case "\n\t":
        var v13 uint16 = 705
        _ = v13
    default:
        var v14 uint64 = 4294967295
        _ = v14
    }
    v2[32768] = 0
    if (!(2147483647 <= 350)) {
        var v15 uint64 = 221
        _ = v15
    }
    switch "hello" {
    case "":
        var v16 float32 = (1.5 - 1e10)
        _ = v16
    case "foo bar":
        for fi17 := 0; fi17 < 19; fi17++ {
            _ = len(v2)
            v4 = v5
            _ = fi17
        }
    }
    fmt.Println(v5, v5, v4)
    v3 = (v3 ^ v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
