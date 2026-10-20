package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint64 }
type GenS2 struct { F0 int64; F1 uint; F2 int16 }
type GenS3 struct { F0 int8; F1 uint32 }
type GenS4 struct { F0 int8; F1 int64; F2 int32 }

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]uint32)
    var v3 uint = 1
    var v4 bool = (!false)
    var v5 int16 = 32767
    if (!(!(true || ((v5 ^ (v5 - v5)) != (v5 * v5))))) {
        v5 += (v5 / ((v5) | 1))
        for fi6 := 0; fi6 < 1; fi6++ {
            v7 := GenS1{F0: 32768, F1: 95}
            _ = fi6
            _ = v7
        }
    }
    switch "hello" {
    case "":
        v3 = 255
    }
    for fi8 := 0; fi8 < 3; fi8++ {
        v4 = v4
        _ = fi8
    }
    var v9 uint16 = 32767
    if v4 {
        _ = len(v1)
        for fi10 := 0; fi10 < 16; fi10++ {
            v11 := GenS2{F0: 179, F1: 4294967295, F2: (v5 * (v5 & 674))}
            var v12 float32 = (-2.25 * ((3.14159 * ((((0.0 / -2.25) - -2.25) * -2.25) * (0.1 / -2.25))) * (0.0 * 1e10)))
            _ = fi10
            _ = v11
            _ = v12
        }
    } else {
        _ = len(v1)
    }
    v13 := GenS3{F0: -128, F1: 866}
    v14 := GenS2{F0: -117, F1: v3, F2: v5}
    var v15 int64 = -9223372036854775808
    for fi16 := 0; fi16 < 13; fi16++ {
        fmt.Println(v4)
        if v4 {
            fmt.Println(v9)
        }
        _ = fi16
    }
    v17 := GenS3{F0: 100, F1: 65535}
    if true {
        if v4 {
            var v18 int64 = (v15 / ((v15) | 1))
            _ = v18
        }
    }
    fmt.Println(v3, v9)
    var v19 int8 = -128
    if true {
        fmt.Println(v3, v3)
        v20 := GenS4{F0: (v19 % ((42) | 1)), F1: v15, F2: 2147483647}
        fmt.Println(v5, v9)
        _ = v20
    } else {
        v21 := GenS3{F0: v19, F1: 0}
        var v22 int16 = (^(v5))
        _ = v21
        _ = v22
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
