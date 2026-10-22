package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 uint32 }
func genFn0(p0 uint32, p1 int32, p2 string) int16 {
    var v6 uint32 = (p0 - (p0 / ((p0) | 1)))
    v7 := []int{-1}
    var v8 float64 = (float64(1.5) + 0.0)
    var v9 int8 = 38
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return 32767
}

func main() {
    v1 := []uint32{4294967295, 65535, 65535, 42}
    v2 := make(map[string]uint32)
    var v3 float64 = 3.14159
    v4 := make(map[int]bool)
    v5 := make(map[uint32]bool)
    var v10 int32 = 2147483647
    if len(v1) > 0 {
        v1[0] = 1
    }
    for fi11 := 0; fi11 < 14; fi11++ {
        for fi12 := 0; fi12 < 16; fi12++ {
            v13 := GenS1{F0: 42}
            switch "" {
            case "":
                v10 += (v10 & -1)
            case "foo bar":
                v10 += (-(v10))
            case "\n\t":
                v4[(fi11 << ((1) & 63))] = true
                v14 := GenS2{F0: (fi12 + fi11), F1: 1}
                fmt.Println(fi11)
                _ = v14
            default:
                v15 := GenS1{F0: 135}
                _ = v15
            }
            delete(v2, "a")
            var v16 uint32 = 1
            if ((-32768 < 0) && (!(true || (1 < 2)))) {
                v16 -= v16
                var v17 int64 = -450
                v18 := GenS1{F0: 1}
                _ = v17
                _ = v18
            } else {
                v3 = (float64(float32(v3)) / (v3 - ((v3 / 1e10) * 0.1)))
            }
            v16 += v16
            _ = fi12
            _ = v13
            _ = v16
        }
        switch "\n\t" {
        case "a":
            _ = len(v1)
        }
        _ = fi11
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v1 = append(v1, 65535)
    var v19 string = "hello"
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    _ = len(v1)
    v10 += v10
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
