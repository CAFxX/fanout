package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 int16 }
type GenS2 struct { F0 uint64 }

func main() {
    v1 := make(map[int]uint32)
    var v2 uint = 4294967295
    v3 := make(map[string]uint32)
    v4 := make(map[uint32]string)
    var v5 uint16 = 32767
    if (uint32(v2) <= 910) {
        v5 -= (v5 << (((v2 & v2)) & 63))
        for fi6 := 0; fi6 < 1; fi6++ {
            var v7 string = ""
            _ = fi6
            _ = v7
        }
    } else {
        var v8 int64 = 9223372036854775807
        _ = v8
    }
    var v9 uint = v2
    var v10 uint64 = 4294967295
    fmt.Println(v10, v2, v10)
    if true {
        switch 4294967295 {
        case 32767:
            fmt.Println(v5, v5, v5)
        case 0:
            var v11 int8 = 127
            _ = v11
        case -414:
            v5 = v5
        }
        v12 := GenS1{F0: (128 < 127), F1: 65535, F2: 32767}
        _ = v12
    }
    for fi13 := 0; fi13 < 13; fi13++ {
        v14 := GenS2{F0: v10}
        _ = fi13
        _ = v14
    }
    var v15 int64 = -9223372036854775808
    v16 := GenS2{F0: 18446744073709551615}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
