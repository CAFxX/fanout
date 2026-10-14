package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint32; F2 int }
type GenS2 struct { F0 uint64; F1 string; F2 uint64 }
type GenS3 struct { F0 uint16 }

func main() {
    v1 := make(map[string]bool)
    var v2 int8 = 127
    var v3 uint8 = 2
    var v4 int16 = -1
    var v5 bool = false
    v4 -= v4
    v6 := GenS1{F0: (v2 % ((-12) | 1)), F1: 2147483648, F2: -451}
    v2 -= -128
    var v7 uint16 = uint16(v3)
    v2 += v2
    for fi8 := 0; fi8 < 5; fi8++ {
        var v9 uint = 496
        v1["hello"] = (!(fi8 == (fi8 << ((v9) & 63))))
        _ = fi8
        _ = v9
    }
    v2 = (^(v2))
    v2 = (v2 << ((748) & 63))
    var v10 int32 = 749
    v2 += -128
    var v11 uint8 = (v3 ^ v3)
    v12 := []int{12, -124, -633, 32768}
    v13 := GenS2{F0: uint64(v7), F1: ((("foo bar" + "\n\t") + "foo bar") + "\n\t"), F2: 42}
    v4 -= v4
    v14 := GenS3{F0: v7}
    if (!v5) {
        var v15 uint16 = (^(v7))
        _ = v15
    }
    if true {
        for fi16 := 0; fi16 < 4; fi16++ {
            var v17 int = (fi16 - (fi16 & (^(fi16))))
            _ = fi16
            _ = v17
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
