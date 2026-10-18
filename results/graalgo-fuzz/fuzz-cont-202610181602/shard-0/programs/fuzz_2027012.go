package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 int16 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 bool; F1 int32; F2 uint16 }

func main() {
    var v1 uint8 = 128
    var v2 uint = 2
    v3 := []string{("\n\t" + ("" + "hello"))}
    v4 := []int{32767, 42, 128, 4294967295}
    var v5 float32 = 1e10
    v2 = 200
    var v6 int32 = 1
    if (100 < 21) {
        var v7 float64 = 0.1
        var v8 uint8 = 1
        _ = v7
        _ = v8
    } else {
        fmt.Println(v2, v1)
        var v9 uint = (v2 << ((65535) & 63))
        _ = v9
    }
    v10 := GenS1{F0: 127, F1: "", F2: -32768}
    v5 -= (v5 * ((3.14159 * float32(3.14159)) / ((float32((-2.25 - -2.25)) + v5) * v5)))
    v11 := GenS2{F0: v6}
    var v12 uint8 = (v1 % (((^(v1))) | 1))
    v13 := GenS3{F0: (1 < int64(v2)), F1: (v6 & -85), F2: 279}
    v12 += v12
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
