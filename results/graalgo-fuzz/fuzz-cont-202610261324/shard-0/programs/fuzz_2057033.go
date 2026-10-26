package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int64; F2 uint16 }
type GenS2 struct { F0 uint; F1 int8; F2 uint8 }
type GenS3 struct { F0 uint; F1 uint64; F2 int32 }

func main() {
    var v1 uint32 = 1
    v2 := make(map[int]bool)
    var v3 uint8 = 165
    v4 := []string{}
    var v5 bool = (!true)
    var v6 int = 127
    v7 := GenS1{F0: "\n\t", F1: -854, F2: 800}
    for fi8 := 0; fi8 < 16; fi8++ {
        fmt.Println(v1)
        if ((^(v1)) < (v1 & (^(v1)))) {
            var v9 uint = 0
            var v10 int = v6
            v11 := GenS2{F0: (v9 & (v9 - v9)), F1: -128, F2: uint8(v1)}
            _ = v10
            _ = v11
            _ = v9
        }
        _ = fi8
    }
    var v12 int8 = -1
    v13 := GenS3{F0: 78, F1: 4294967295, F2: -2147483648}
    switch "" {
    case "\n\t":
        v5 = ((v1 % (((v1 | v1)) | 1)) < (v1 << ((108) & 63)))
    }
    for fi14 := 0; fi14 < 19; fi14++ {
        for fi15 := 0; fi15 < 11; fi15++ {
            var v16 bool = (9223372036854775807 == 9223372036854775807)
            v17 := v2[fi14]
            _ = fi15
            _ = v16
            _ = v17
        }
        v6 += (-(v6))
        _ = fi14
    }
    for fi18 := 0; fi18 < 9; fi18++ {
        v19 := GenS1{F0: (((("\n\t" + "a") + ("a" + "a")) + "hello") + "\n\t"), F1: 42, F2: 335}
        if v5 {
            v20 := GenS3{F0: 0, F1: 4294967295, F2: -1}
            var v21 int64 = 1
            _ = v20
            _ = v21
        }
        _ = fi18
        _ = v19
    }
    v22 := GenS1{F0: ("" + ""), F1: 9223372036854775807, F2: 255}
    var v23 uint8 = (v3 * (^(v3)))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v13
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
