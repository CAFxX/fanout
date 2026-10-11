package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8; F2 uint }
type GenS2 struct { F0 int32; F1 int }
type GenS3 struct { F0 bool; F1 uint16 }
func genFn0(p0 uint16) uint64 {
    p0 = (p0 | 2)
    v6 := []int{-319, 32768}
    if len(v6) >= 2 {
        v6 = v6[1:len(v6)]
    }
    var v7 uint64 = 2
    _ = p0
    _ = v6
    _ = v7
    return 42
}

func main() {
    var v1 bool = (2 <= 83)
    var v2 int = 42
    v3 := make(map[string]string)
    var v4 int32 = 42
    var v5 float32 = (-2.25 / 1.0)
    v8 := GenS1{F0: "", F1: 0, F2: 0}
    v9 := []int{(^(v2)), (^(v2)), 320}
    var v10 string = "a"
    var v11 bool = v1
    v5 = 1e10
    for fi12 := 0; fi12 < 11; fi12++ {
        v13 := GenS2{F0: (v4 * (v4 - v4)), F1: 397}
        fmt.Println(v4)
        for fi14 := 0; fi14 < 4; fi14++ {
            switch 9223372036854775807 {
            case 4294967295:
                v11 = v11
            case 620:
                v15 := GenS3{F0: (((!((fi14 * fi12) > 42)) && (!false)) && v11), F1: 2}
                _ = v15
            }
            _ = fi14
        }
        _ = fi12
        _ = v13
    }
    var v16 int = v2
    v11 = v1
    v10 = (v10 + v10)
    v17 := v3[v10]
    for fi18 := 0; fi18 < 5; fi18++ {
        for fi19 := 0; fi19 < 18; fi19++ {
            v4 -= v4
            _ = fi19
        }
        var v20 bool = (v17 < v17)
        _ = fi18
        _ = v20
    }
    fmt.Println(v2, v17, v10)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
