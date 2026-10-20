package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }

func main() {
    var v1 uint64 = 1
    v2 := make(map[string]bool)
    v3 := make(map[int]uint32)
    v1 -= v1
    var v4 float32 = 1e10
    if (-1 > 42) {
        var v5 int16 = 1
        fmt.Println(v1, v5, v4)
        _ = v5
    } else {
        var v6 string = ""
        for fi7 := 0; fi7 < 2; fi7++ {
            var v8 float64 = 3.14159
            var v9 bool = (!false)
            _ = fi7
            _ = v8
            _ = v9
        }
        var v10 int = -9223372036854775808
        _ = v10
        _ = v6
    }
    if ((v1 & 0) > (^(v1))) {
        var v11 int64 = 42
        _ = v11
    }
    v12 := []int{127, 32768, -9223372036854775808}
    v4 = 3.14159
    v13 := GenS1{F0: ((v1 | (v1 << ((255) & 63))) < v1), F1: 0}
    var v14 uint32 = 592
    for fi15 := 0; fi15 < 9; fi15++ {
        fmt.Println(v4)
        for fi16 := 0; fi16 < 3; fi16++ {
            var v17 int = (fi16 - 850)
            _ = fi16
            _ = v17
        }
        v18 := GenS1{F0: (false && (((0.1 - (1.5 / 1e10)) * -2.25) != ((1e10 + (0.1 + 3.14159)) / -2.25))), F1: 441}
        _ = fi15
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
