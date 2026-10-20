package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint32 }
type GenS2 struct { F0 int16; F1 int8; F2 int32 }
func genFn0(p0 string, p1 string, p2 int8) int8 {
    fmt.Println(p0, p1)
    switch 4294967295 {
    case 1:
        p1 = p0
    case -9223372036854775808:
        v6 := []int{-593}
        _ = v6
    case 32767:
        fmt.Println(p2, p2, p2)
    }
    _ = p0
    _ = p1
    _ = p2
    return (p2 & p2)
}
func genFn1() bool {
    var v7 uint64 = 852
    v7 += 2
    _ = v7
    return true
}

func main() {
    v1 := []string{""}
    var v2 float32 = ((float32(float64(1e10)) * 0.1) * 0.0)
    v3 := make(map[uint32]int)
    var v4 uint8 = 255
    v5 := make(map[uint32]int)
    if (888 == 255) {
        fmt.Println(v4, v2)
        v8 := GenS1{F0: 0, F1: 0}
        _ = v8
    }
    var v9 int16 = 42
    var v10 uint8 = (v4 | v4)
    fmt.Println(v10, v4)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    if (false || true) {
        v4 += v4
        v10 = 1
    }
    v10 -= 32
    if (65535 < 42) {
        v10 = 102
        var v11 string = ("\n\t" + "")
        _ = v11
    }
    if (!false) {
        for fi12 := 0; fi12 < 16; fi12++ {
            fmt.Println(v2, fi12, fi12)
            if false {
                v13 := v5[2147483647]
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = v13
            } else {
                v5[4294967295] = (fi12 | fi12)
                v14 := GenS1{F0: 663, F1: 748}
                _ = v14
            }
            _ = fi12
        }
        var v15 uint8 = v4
        _ = v15
    }
    if true {
        if true {
            _ = len(v1)
            var v16 float64 = 1e10
            _ = v16
        }
        v17 := GenS2{F0: (v9 + (v9 * 42)), F1: 0, F2: 42}
        var v18 int = -701
        _ = v17
        _ = v18
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
