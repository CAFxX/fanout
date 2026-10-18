package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint32; F2 uint8 }
type GenS2 struct { F0 string; F1 int32; F2 uint64 }
func genFn0(p0 int32, p1 string) uint64 {
    if true {
        v5 := GenS1{F0: 676, F1: 204, F2: 0}
        _ = v5
    } else {
        var v6 bool = false
        _ = v6
    }
    if false {
        p1 = (p1 + "\n\t")
        v7 := GenS2{F0: p1, F1: -901, F2: 2}
        _ = v7
    } else {
        var v8 int32 = (p0 + p0)
        if (1 < 32767) {
            v8 = 899
        }
        _ = v8
    }
    _ = p0
    _ = p1
    return 2
}

func main() {
    v1 := []int{958, 15}
    v2 := make(map[int]uint32)
    var v3 uint = 42
    v4 := []uint32{287, 116, 2147483647}
    fmt.Println(v3)
    switch "\n\t" {
    case "\n\t":
        fmt.Println(v3)
    case "":
        v3 = 18446744073709551615
    case "a":
        switch "" {
        case "a":
            fmt.Println(v3)
        default:
            for fi9 := 0; fi9 < 1; fi9++ {
                for fi10 := 0; fi10 < 9; fi10++ {
                    var v11 uint = v3
                    var v12 float64 = (1e10 / (1.5 / (0.1 * (0.1 - (3.14159 - 1e10)))))
                    _ = fi10
                    _ = v11
                    _ = v12
                }
                v13 := GenS1{F0: 1, F1: 969, F2: 0}
                _ = fi9
                _ = v13
            }
        }
    default:
        v3 -= v3
    }
    fmt.Println(v3)
    v2[4294967295] = 274
    var v14 uint = v3
    _ = len(v4)
    var v15 float32 = (1e10 + 0.0)
    _ = len(v2)
    v14 += (v3 & v14)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
}
