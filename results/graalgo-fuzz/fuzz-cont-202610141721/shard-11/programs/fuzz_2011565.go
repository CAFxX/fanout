package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint; F2 int }
type GenS2 struct { F0 int32; F1 int16 }
type GenS3 struct { F0 uint64; F1 int16 }
func genFn0(p0 uint32, p1 int8, p2 uint8) int32 {
    p2 = 136
    p1 = (p1 ^ p1)
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    v1 := []uint32{4294967295}
    v2 := []string{"", ("" + "a"), "hello", "\n\t"}
    v3 := make(map[int]int)
    var v4 int32 = 0
    for fi5 := 0; fi5 < 8; fi5++ {
        for fi6 := 0; fi6 < 5; fi6++ {
            fmt.Println(fi5, v4)
            fmt.Println(fi6)
            v7 := GenS1{F0: 564, F1: 155, F2: fi5}
            var v8 uint8 = 169
            fmt.Println(v8, fi5)
            _ = fi6
            _ = v7
            _ = v8
        }
        for fi9 := 0; fi9 < 2; fi9++ {
            var v10 float64 = ((0.1 - -2.25) - 0.1)
            _ = fi9
            _ = v10
        }
        fmt.Println(fi5)
        _ = fi5
    }
    for fi11 := 0; fi11 < 19; fi11++ {
        v12 := GenS2{F0: v4, F1: -1}
        _ = fi11
        _ = v12
    }
    v13 := GenS3{F0: 0, F1: 1}
    v4 -= 0
    var v14 float32 = ((1e10 / 1e10) - (0.0 / 1.0))
    v15 := GenS2{F0: 1, F1: 1000}
    _ = len(v2)
    v16 := GenS1{F0: 18446744073709551615, F1: 18446744073709551615, F2: 42}
    v4 -= genFn0(4294967295, 42, 42)
    var v17 float64 = 3.14159
    var v18 uint8 = 2
    for fi19 := 0; fi19 < 20; fi19++ {
        var v20 uint64 = 18446744073709551615
        var v21 int8 = -128
        _ = fi19
        _ = v20
        _ = v21
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
