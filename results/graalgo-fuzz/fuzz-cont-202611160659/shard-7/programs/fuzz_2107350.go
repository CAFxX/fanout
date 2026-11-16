package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 int8; F1 int8; F2 uint }
func genFn0(p0 string, p1 uint64) uint32 {
    var v4 int64 = 42
    fmt.Println(p1, p0)
    _ = p0
    _ = p1
    _ = v4
    return 0
}

func main() {
    var v1 int8 = 127
    v2 := []uint32{0, 2147483647, 331, 814}
    var v3 int16 = -1
    v5 := GenS1{F0: 697, F1: 0}
    var v6 uint = 2
    var v7 int64 = 764
    for fi8 := 0; fi8 < 20; fi8++ {
        var v9 uint = (v6 + (v6 + v6))
        if len(v2) > 0 {
            v2[0] = 15
        }
        _ = fi8
        _ = v9
    }
    v10 := GenS2{F0: v7, F1: 42}
    var v11 int32 = -2147483648
    switch ("hello" + "") {
    case "\n\t":
        for fi12 := 0; fi12 < 4; fi12++ {
            v6 = (v6 / ((323) | 1))
            var v13 bool = true
            fmt.Println(v6)
            _ = fi12
            _ = v13
        }
    default:
        var v14 string = "hello"
        _ = v14
    }
    v15 := GenS3{F0: v1, F1: (v1 >> (((v6 % ((v6) | 1))) & 63)), F2: v6}
    fmt.Println(v3, v6, v11)
    v16 := GenS3{F0: (v1 | v1), F1: (v1 << (((v6 | 54)) & 63)), F2: v6}
    var v17 int32 = 2147483647
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v5
    _ = v6
    _ = v7
}
