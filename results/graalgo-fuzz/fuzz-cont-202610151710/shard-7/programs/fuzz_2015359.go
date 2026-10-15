package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16 }
func genFn0(p0 int64, p1 uint8) int64 {
    if (1 > 255) {
        var v5 uint64 = 18446744073709551615
        v6 := GenS1{F0: "a"}
        _ = v5
        _ = v6
    }
    fmt.Println(p1, p1)
    v7 := GenS1{F0: "\n\t"}
    p0 -= p0
    for fi8 := 0; fi8 < 10; fi8++ {
        v9 := []int{(fi8 | fi8), fi8}
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v7
    return (p0 & (-(p0)))
}
func genFn1(p0 uint, p1 int32) int8 {
    v10 := GenS2{F0: -1}
    for fi11 := 0; fi11 < 20; fi11++ {
        for fi12 := 0; fi12 < 2; fi12++ {
            v13 := GenS2{F0: -511}
            var v14 int16 = 42
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = fi11
    }
    p1 -= 0
    _ = p0
    _ = p1
    _ = v10
    return 1
}

func main() {
    var v1 int = 42
    var v2 float64 = 0.0
    v3 := []uint32{0}
    var v4 int64 = -1
    v2 += (-2.25 * ((v2 + float64((((1e10 + 1.5) / float32((v2 + 0.1))) - 1e10))) / v2))
    _ = len(v3)
    v1 -= v1
    v1 -= v1
    v2 -= v2
    if len(v3) > 0 {
        v3[0] = 4294967295
    }
    for fi15 := 0; fi15 < 3; fi15++ {
        var v16 uint64 = 0
        v16 -= 42
        _ = fi15
        _ = v16
    }
    v1 += (v1 | v1)
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
}
