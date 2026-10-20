package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int32 }
type GenS2 struct { F0 int64; F1 uint8 }
type GenS3 struct { F0 bool; F1 int; F2 int32 }

func main() {
    var v1 int = 4294967295
    var v2 uint64 = 855
    v3 := []uint32{}
    var v4 int16 = 1000
    if false {
        var v5 uint8 = 1
        _ = v5
    } else {
        var v6 uint16 = 192
        v4 = v4
        v7 := GenS1{F0: 65535, F1: 65}
        _ = v6
        _ = v7
    }
    v4 += 602
    v1 = 84
    v8 := GenS2{F0: 671, F1: 42}
    var v9 uint64 = (v2 * 257)
    v4 = v4
    v4 = (v4 % ((-32768) | 1))
    var v10 bool = true
    var v11 uint64 = v2
    switch "" {
    case "foo bar":
        var v12 string = "foo bar"
        _ = v12
    }
    v11 = v9
    var v13 float64 = 1e10
    v2 += 42
    v14 := GenS3{F0: false, F1: 127, F2: -371}
    var v15 int8 = 127
    if (int16(v15) != 0) {
        v10 = (v10 && v10)
    } else {
        v3 = append(v3, 4294967295)
        v1 -= 32767
    }
    fmt.Println(v15)
    if false {
        v16 := GenS2{F0: 42, F1: 127}
        _ = v16
    } else {
        v17 := GenS1{F0: 2147483648, F1: int32(v4)}
        if ((v9 * v11) >= v11) {
            if len(v3) > 0 {
                v3[0] = 42
            }
            fmt.Println(v4, v1, v15)
        } else {
            var v18 bool = (!false)
            _ = v18
        }
        if len(v3) > 0 {
            v3[0] = 811
        }
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
