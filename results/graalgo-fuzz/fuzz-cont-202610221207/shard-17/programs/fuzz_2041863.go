package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint; F1 int16; F2 uint16 }
type GenS3 struct { F0 uint64; F1 uint }

func main() {
    var v1 float32 = (0.0 * 3.14159)
    var v2 float32 = v1
    v3 := []uint32{4294967295, 0, 831}
    v4 := make(map[string]int)
    v5 := []int{-1}
    for fi6 := 0; fi6 < 9; fi6++ {
        for fi7 := 0; fi7 < 15; fi7++ {
            if (!(!(127 >= 100))) {
                var v8 string = "\n\t"
                _ = v8
            }
            v1 = -2.25
            var v9 float64 = 1e10
            _ = fi7
            _ = v9
        }
        _ = fi6
    }
    var v10 int = -105
    if (-1 != -1) {
        v11 := GenS1{F0: 0}
        for fi12 := 0; fi12 < 6; fi12++ {
            v1 = 1e10
            _ = fi12
        }
        _ = v11
    } else {
        v1 = v2
        fmt.Println(v10, v1)
    }
    var v13 uint16 = 2
    var v14 int16 = -32768
    v15 := GenS2{F0: 42, F1: v14, F2: (v13 ^ v13)}
    v2 = (v2 - 3.14159)
    for fi16 := 0; fi16 < 5; fi16++ {
        v17 := GenS1{F0: (v14 * v14)}
        for fi18 := 0; fi18 < 11; fi18++ {
            var v19 uint8 = 127
            fmt.Println(v14)
            var v20 uint8 = 2
            _ = fi18
            _ = v19
            _ = v20
        }
        var v21 int16 = v14
        _ = fi16
        _ = v17
        _ = v21
    }
    v22 := GenS3{F0: 355, F1: 0}
    var v23 uint32 = 0
    v2 += v2
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
