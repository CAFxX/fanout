package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 int8 }
type GenS2 struct { F0 uint; F1 int32 }
type GenS3 struct { F0 int64; F1 int16 }
func genFn0(p0 uint32) int64 {
    var v6 int8 = 0
    fmt.Println(v6)
    v7 := []int{127}
    var v8 float64 = (0.1 * 0.1)
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 uint32 = 4294967295
    var v2 uint = 376
    var v3 int64 = 42
    v4 := make(map[string]int)
    v5 := make(map[uint32]bool)
    v9 := GenS1{F0: (v3 * (v3 >> (((v2 + v2)) & 63))), F1: v1, F2: 49}
    v1 = v1
    var v10 uint8 = 126
    for fi11 := 0; fi11 < 15; fi11++ {
        delete(v5, 2147483647)
        _ = fi11
    }
    v12 := []int{}
    v13 := GenS2{F0: 42, F1: -1}
    v2 += v2
    var v14 int = -1
    if (257 < 4294967295) {
        v3 = v3
        v2 += (v2 | 481)
        v15 := GenS2{F0: (v2 >> ((18446744073709551615) & 63)), F1: -2147483648}
        _ = v15
    } else {
        if (0 != 0) {
            switch "foo bar" {
            case "foo bar":
                v16 := GenS3{F0: (v3 + v3), F1: -32768}
                _ = v16
            case "a":
                v2 = (v2 + (^(v2)))
            case "\n\t":
                v1 = (v1 * v1)
            default:
                var v17 string = ("foo bar" + ("\n\t" + "\n\t"))
                _ = v17
            }
            fmt.Println(v14, v14, v1)
            var v18 uint8 = 12
            _ = v18
        } else {
            v19 := v5[v1]
            _ = v19
        }
        fmt.Println(v10, v1, v2)
    }
    for fi20 := 0; fi20 < 3; fi20++ {
        var v21 int = v14
        _ = fi20
        _ = v21
    }
    _ = len(v12)
    for fi22 := 0; fi22 < 15; fi22++ {
        var v23 int64 = 1
        v14 = fi22
        _ = fi22
        _ = v23
    }
    v10 = v10
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
