package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 uint32 }
func genFn0(p0 int16) int64 {
    var v4 int = 1
    v4 = (v4 | v4)
    v5 := []int{v4, v4, v4}
    if true {
        var v6 uint8 = 2
        for fi7 := 0; fi7 < 15; fi7++ {
            for fi8 := 0; fi8 < 18; fi8++ {
                var v9 int = (fi8 / ((fi8) | 1))
                _ = fi8
                _ = v9
            }
            _ = fi7
        }
        v10 := GenS1{F0: -1, F1: 255, F2: 678}
        _ = v10
        _ = v6
    } else {
        v5 = append(v5, (v4 ^ 2147483647))
        for fi11 := 0; fi11 < 3; fi11++ {
            for fi12 := 0; fi12 < 1; fi12++ {
                var v13 float64 = (-2.25 + (float64(0.1) / 1e10))
                var v14 uint64 = 627
                _ = fi12
                _ = v13
                _ = v14
            }
            _ = fi11
        }
    }
    _ = p0
    _ = v4
    _ = v5
    return -788
}

func main() {
    var v1 uint16 = 2
    v2 := make(map[int]bool)
    v3 := make(map[string]int)
    v15 := []int{-9223372036854775808}
    fmt.Println(v1)
    switch 32768 {
    case 2147483647:
        var v16 int32 = -2147483648
        _ = v16
    case -1:
        var v17 bool = false
        _ = v17
    case 128:
        v18 := GenS1{F0: -1, F1: 235, F2: 2}
        _ = v18
    }
    delete(v2, 2147483647)
    var v19 int = -2147483648
    var v20 int64 = 42
    _ = len(v2)
    for fi21 := 0; fi21 < 3; fi21++ {
        var v22 int16 = -462
        _ = fi21
        _ = v22
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
