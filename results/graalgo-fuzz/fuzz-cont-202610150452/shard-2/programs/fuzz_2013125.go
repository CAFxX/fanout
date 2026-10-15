package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 uint8 }
type GenS2 struct { F0 string; F1 uint64 }
func genFn0() int32 {
    var v6 uint64 = 0
    v6 = v6
    _ = v6
    return 561
}

func main() {
    v1 := []int{2147483647, 721, 32768, 32768}
    var v2 float64 = float64(((((1e10 - (0.1 - 0.0)) - 1.5) / (1e10 + 3.14159)) - 3.14159))
    v3 := make(map[int]int)
    v4 := []string{("a" + "\n\t")}
    v5 := []uint32{1, 359, 542, 817}
    v2 = v2
    _ = len(v3)
    var v7 bool = false
    var v8 uint64 = 1
    var v9 uint8 = 255
    if v7 {
        if ("" != "\n\t") {
            v10 := GenS1{F0: -381, F1: 447, F2: v9}
            _ = v10
        }
        var v11 bool = (2 >= 748)
        _ = v11
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        if v7 {
            v5 = append(v5, uint32(v9))
        }
        var v12 uint64 = v8
        _ = v12
    }
    if v7 {
        for fi13 := 0; fi13 < 6; fi13++ {
            var v14 uint8 = 0
            for fi15 := 0; fi15 < 3; fi15++ {
                v16 := GenS1{F0: -908, F1: 42, F2: v14}
                _ = fi15
                _ = v16
            }
            _ = fi13
            _ = v14
        }
        var v17 int8 = -73
        _ = v17
    }
    switch 9223372036854775807 {
    case -1:
        var v18 string = ""
        _ = v18
    case -214:
        v7 = (!v7)
    case -2147483648:
        v7 = (-742 <= 1)
    default:
        v19 := GenS2{F0: "", F1: uint64(v9)}
        _ = v19
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
