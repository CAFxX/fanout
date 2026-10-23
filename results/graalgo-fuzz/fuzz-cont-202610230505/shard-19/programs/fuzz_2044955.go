package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint64; F1 uint16 }
type GenS3 struct { F0 uint64 }
func genFn0(p0 int32, p1 int32) int16 {
    fmt.Println(p0, p1)
    v7 := []int{1, 1, 32768}
    for fi8 := 0; fi8 < 12; fi8++ {
        for fi9 := 0; fi9 < 4; fi9++ {
            v10 := GenS1{F0: 0}
            switch fi8 {
            case -293:
                v11 := GenS1{F0: 42}
                _ = v11
            case 32767:
                var v12 bool = (-45 == 9223372036854775807)
                _ = v12
            default:
                p1 += (p0 * p0)
            }
            _ = fi9
            _ = v10
        }
        var v13 float32 = (float32(0.0) / (0.1 * 1e10))
        var v14 bool = false
        _ = fi8
        _ = v13
        _ = v14
    }
    p0 -= p1
    p0 += p0
    _ = p0
    _ = p1
    _ = v7
    return -66
}

func main() {
    var v1 int16 = 42
    v2 := make(map[int]string)
    v3 := []uint32{}
    var v4 uint64 = 1
    var v5 int = 128
    v6 := []string{("foo bar" + "a"), "a", "hello"}
    v15 := GenS1{F0: 2147483648}
    v1 -= v1
    if (!true) {
        fmt.Println(v5, v4)
        v16 := GenS2{F0: (v4 / ((v4) | 1)), F1: 32767}
        v17 := GenS3{F0: 983}
        _ = v16
        _ = v17
    } else {
        var v18 int32 = 831
        _ = v18
    }
    var v19 uint16 = 65535
    for fi20 := 0; fi20 < 3; fi20++ {
        fmt.Println(v4, fi20)
        var v21 uint64 = v4
        _ = fi20
        _ = v21
    }
    fmt.Println(v4)
    var v22 float64 = float64((1e10 - 1.5))
    fmt.Println(v5, v1, v4)
    _ = len(v6)
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
