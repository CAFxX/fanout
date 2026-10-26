package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }

func main() {
    v1 := []int{0, -2147483648, 60, 32767}
    var v2 int32 = 2147483647
    v3 := make(map[string]bool)
    v4 := []int{-317, -1}
    var v5 uint16 = 2
    for fi6 := 0; fi6 < 1; fi6++ {
        if len(v4) > 0 {
            v4[0] = (fi6 | 857)
        }
        _ = fi6
    }
    var v7 uint64 = 328
    v7 -= 42
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v7 -= (^(v7))
    var v8 float64 = 1e10
    for fi9 := 0; fi9 < 18; fi9++ {
        v10 := GenS1{F0: 2147483647}
        var v11 uint64 = v7
        if (!(100 != 127)) {
            var v12 int16 = -1
            var v13 float32 = float32(v8)
            _ = v12
            _ = v13
        }
        _ = fi9
        _ = v10
        _ = v11
    }
    fmt.Println(v7)
    var v14 int16 = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
}
