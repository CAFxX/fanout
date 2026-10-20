package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16; F1 int64; F2 int8 }

func main() {
    var v1 uint64 = 42
    v2 := make(map[uint32]string)
    var v3 bool = false
    var v4 uint32 = 4294967295
    v5 := GenS1{F0: 924}
    if (-45 == 63) {
        v4 = (v4 / (((v4 | v4)) | 1))
        var v6 int = 9223372036854775807
        switch -39 {
        case -2147483648:
            v6 = (v6 / (((v6 << ((0) & 63))) | 1))
        default:
            v1 = v1
        }
        _ = v6
    } else {
        var v7 float64 = float64(-2.25)
        for fi8 := 0; fi8 < 7; fi8++ {
            v3 = (!(2 < 70))
            var v9 uint = 4294967295
            _ = fi8
            _ = v9
        }
        _ = v7
    }
    var v10 uint = 4294967295
    if v3 {
        var v11 uint64 = (v1 + (v1 << ((528) & 63)))
        var v12 uint64 = (v11 | 0)
        _ = v11
        _ = v12
    }
    v10 += (v10 - 4294967295)
    var v13 bool = (2 == 42)
    var v14 int32 = -2147483648
    for fi15 := 0; fi15 < 3; fi15++ {
        v16 := GenS2{F0: 2, F1: 403, F2: int8(v1)}
        _ = fi15
        _ = v16
    }
    v13 = (!(!(2147483647 <= (v4 - (v4 - v4)))))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
