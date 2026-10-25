package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint64; F1 uint }
func genFn0(p0 uint8, p1 uint32, p2 string) string {
    fmt.Println(p2, p2, p2)
    v6 := GenS1{F0: 1}
    p2 = p2
    v7 := []int{32768, 127}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return p2
}

func main() {
    var v1 uint32 = 16
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]string)
    v4 := make(map[string]int)
    var v5 uint8 = 127
    var v8 float64 = 0.0
    var v9 int16 = -532
    var v10 int32 = 2147483647
    for fi11 := 0; fi11 < 7; fi11++ {
        switch 32767 {
        case -9223372036854775808:
            var v12 bool = false
            _ = v12
        case 705:
            v13 := GenS1{F0: 1}
            _ = v13
        default:
            v14 := GenS2{F0: 4294967295}
            _ = v14
        }
        _ = fi11
    }
    v15 := GenS3{F0: 4294967295, F1: 255}
    fmt.Println(v9)
    v10 -= v10
    if false {
        if (-562 < 0) {
            var v16 int64 = -9223372036854775808
            var v17 float64 = 1.5
            fmt.Println(v9, v5)
            _ = v16
            _ = v17
        }
        fmt.Println(v8)
    }
    var v18 uint32 = (v1 / (((v1 / ((v1) | 1))) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v18:", v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
