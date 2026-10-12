package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint; F2 uint64 }
type GenS2 struct { F0 int32; F1 uint16; F2 string }
type GenS3 struct { F0 int64 }
func genFn0(p0 uint16, p1 uint64, p2 int) uint8 {
    p0 += (p0 + (p0 ^ (^(p0))))
    v6 := GenS1{F0: (p2 / ((p2) | 1)), F1: 0, F2: p1}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return 156
}

func main() {
    v1 := []int{}
    var v2 float64 = (0.1 / -2.25)
    var v3 float64 = float64(1e10)
    v4 := make(map[uint32]bool)
    var v5 int8 = 1
    var v7 bool = (-114 <= -129)
    switch "hello" {
    case "\n\t":
        fmt.Println(v7)
    case "":
        v8 := GenS1{F0: -9223372036854775808, F1: 2, F2: 598}
        _ = v8
    case "foo bar":
        fmt.Println(v3, v2, v5)
    }
    v5 += v5
    for fi9 := 0; fi9 < 13; fi9++ {
        var v10 float32 = (((((0.1 / (0.1 / float32(-2.25))) + 1e10) + 1.5) / 3.14159) - -2.25)
        switch "hello" {
        case "":
            for fi11 := 0; fi11 < 1; fi11++ {
                v7 = false
                v12 := GenS2{F0: 0, F1: 65535, F2: "a"}
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = fi11
                _ = v12
            }
        case "foo bar":
            var v13 bool = v7
            _ = v13
        case "a":
            fmt.Println(v7, v7)
        }
        v14 := GenS3{F0: 1}
        _ = fi9
        _ = v10
        _ = v14
    }
    for fi15 := 0; fi15 < 19; fi15++ {
        v5 -= v5
        if (!v7) {
            var v16 bool = false
            fmt.Println(fi15, v16, v2)
            var v17 uint = 18446744073709551615
            _ = v16
            _ = v17
        }
        _ = fi15
    }
    v5 -= (v5 / (((v5 % ((42) | 1))) | 1))
    var v18 bool = false
    var v19 int8 = -128
    v7 = (!(v7 || (0 != 1)))
    var v20 int = 73
    v18 = (true || v7)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v7
}
