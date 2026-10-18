package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int }
func genFn0(p0 int64) int8 {
    p0 -= (p0 << ((526) & 63))
    v5 := []int{32768, 128}
    v6 := GenS1{F0: (p0 - (p0 * (p0 << ((112) & 63))))}
    v7 := GenS1{F0: 42}
    var v8 int32 = 42
    _ = p0
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 100
}
func genFn1(p0 uint16, p1 bool) uint32 {
    if ((80 < 1) && p1) {
        p1 = (p1 || (!p1))
        for fi9 := 0; fi9 < 6; fi9++ {
            var v10 string = "\n\t"
            var v11 int = 9223372036854775807
            _ = fi9
            _ = v10
            _ = v11
        }
        var v12 float64 = 3.14159
        _ = v12
    } else {
        v13 := []int{}
        var v14 uint64 = 18446744073709551615
        v15 := GenS2{F0: -762}
        _ = v13
        _ = v14
        _ = v15
    }
    var v16 int32 = 2147483647
    _ = p0
    _ = p1
    _ = v16
    return 869
}

func main() {
    v1 := []int{-59, 127, 42}
    var v2 uint = 627
    v3 := make(map[uint32]int)
    var v4 int8 = -119
    var v17 uint = 65535
    v2 = (v2 << (((v17 * 255)) & 63))
    v18 := GenS1{F0: -23}
    var v19 int = 32767
    for fi20 := 0; fi20 < 13; fi20++ {
        var v21 int8 = v4
        fmt.Println(v17)
        _ = fi20
        _ = v21
    }
    var v22 int64 = 0
    _ = len(v3)
    for fi23 := 0; fi23 < 19; fi23++ {
        var v24 string = "a"
        _ = fi23
        _ = v24
    }
    var v25 uint = (^(v17))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v25:", v25)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v25
    _ = v3
    _ = v4
}
