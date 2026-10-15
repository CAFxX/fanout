package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint16; F1 int; F2 uint8 }

func main() {
    v1 := make(map[uint32]string)
    v2 := make(map[uint32]int)
    v3 := make(map[int]string)
    v4 := GenS1{F0: false, F1: 0}
    var v5 uint16 = 1
    var v6 bool = (0.1 == 0.1)
    var v7 float64 = -2.25
    for fi8 := 0; fi8 < 20; fi8++ {
        v7 = v7
        var v9 int64 = 1
        var v10 bool = true
        _ = fi8
        _ = v10
        _ = v9
    }
    for fi11 := 0; fi11 < 17; fi11++ {
        var v12 int8 = 87
        var v13 int = -9223372036854775808
        for fi14 := 0; fi14 < 19; fi14++ {
            var v15 uint = 426
            fmt.Println(v13, v7, v7)
            if (true && v6) {
                v15 = (v15 & (v15 | (v15 + (v15 & (v15 / (((v15 ^ (v15 | 558))) | 1))))))
                var v16 bool = v6
                _ = v16
            } else {
                v17 := GenS2{F0: 1}
                fmt.Println(v15, v12, v12)
                _ = v17
            }
            if (-356 < 9223372036854775807) {
                v18 := v3[fi11]
                _ = v18
            }
            v19 := GenS3{F0: 780}
            _ = fi14
            _ = v15
            _ = v19
        }
        _ = fi11
        _ = v12
        _ = v13
    }
    v6 = (!(0 == 0))
    if (-1 > 805) {
        fmt.Println(v6, v7)
    }
    v3[2147483647] = "\n\t"
    var v20 uint = 1
    v21 := GenS4{F0: v5, F1: 9223372036854775807, F2: 127}
    switch "hello" {
    case "\n\t":
        v22 := v2[0]
        _ = v22
    }
    _ = len(v2)
    fmt.Println(v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
