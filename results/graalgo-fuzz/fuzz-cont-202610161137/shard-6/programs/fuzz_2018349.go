package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    v1 := make(map[int]string)
    var v2 bool = (1 > 1)
    v3 := make(map[int]int)
    v4 := []int{32768, 42}
    v5 := make(map[string]uint32)
    var v6 int32 = 42
    var v7 uint = 2
    v7 = v7
    delete(v1, int(v6))
    v8 := GenS1{F0: -32768}
    v2 = v2
    switch 0 {
    case 32768:
        for fi9 := 0; fi9 < 8; fi9++ {
            fmt.Println(v2, v6, fi9)
            for fi10 := 0; fi10 < 1; fi10++ {
                if (false || v2) {
                    var v11 float64 = 1.5
                    _ = v11
                }
                v7 -= (v7 / (((v7 >> ((222) & 63))) | 1))
                v12 := GenS1{F0: int16(fi10)}
                _ = fi10
                _ = v12
            }
            v13 := GenS1{F0: 0}
            _ = fi9
            _ = v13
        }
    case 1:
        fmt.Println(v2, v2, v6)
    }
    v2 = v2
    var v14 uint32 = 832
    for fi15 := 0; fi15 < 2; fi15++ {
        switch int(v7) {
        case 1:
            var v16 uint8 = 194
            _ = v16
        default:
            var v17 int = (fi15 % ((fi15) | 1))
            _ = v17
        }
        if v2 {
            v18 := v3[fi15]
            var v19 bool = v2
            fmt.Println(v18)
            _ = v18
            _ = v19
        }
        _ = fi15
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
