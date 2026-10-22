package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32; F1 int; F2 uint8 }
type GenS3 struct { F0 int64; F1 uint16 }
func genFn0(p0 uint32) int {
    p0 -= (p0 % ((2) | 1))
    fmt.Println(p0)
    var v7 int16 = -32768
    var v8 int64 = 470
    _ = p0
    _ = v7
    _ = v8
    return -47
}

func main() {
    var v1 bool = (!false)
    var v2 int16 = 32767
    var v3 bool = v1
    var v4 bool = true
    var v5 int16 = v2
    v6 := make(map[int]string)
    var v9 uint = 1
    v9 = v9
    var v10 int32 = -1
    var v11 uint32 = 1
    v12 := GenS1{F0: v9}
    delete(v6, 32768)
    var v13 int16 = -918
    if ((0.0 / 1e10) == 3.14159) {
        var v14 int8 = 100
        for fi15 := 0; fi15 < 6; fi15++ {
            v16 := GenS2{F0: (v11 >> (((v9 & (v9 + v9))) & 63)), F1: (fi15 / (((-(fi15))) | 1)), F2: 1}
            _ = fi15
            _ = v16
        }
        _ = v14
    } else {
        var v17 uint8 = 42
        _ = v17
    }
    var v18 int8 = 0
    var v19 int64 = -9223372036854775808
    switch "a" {
    case "\n\t":
        v20 := GenS3{F0: v19, F1: 62}
        _ = v20
    }
    v3 = v3
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
