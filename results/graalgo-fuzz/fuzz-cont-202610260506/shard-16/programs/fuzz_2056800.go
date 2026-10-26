package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int32; F2 int64 }

func main() {
    var v1 int64 = 299
    var v2 int8 = -1
    var v3 uint16 = 32767
    v4 := []uint32{10}
    for fi5 := 0; fi5 < 3; fi5++ {
        v6 := GenS1{F0: 2, F1: 2147483647, F2: (v1 - v1)}
        v7 := GenS1{F0: (v3 & (v3 - v3)), F1: 42, F2: (v1 + v1)}
        if false {
            v1 = (^(v1))
        } else {
            for fi8 := 0; fi8 < 12; fi8++ {
                v2 += (v2 ^ 47)
                _ = fi8
            }
        }
        _ = fi5
        _ = v6
        _ = v7
    }
    if (1 == 171) {
        v3 = v3
    }
    var v9 uint64 = 4294967295
    v9 = (v9 | 617)
    switch ("foo bar" + "\n\t") {
    case "foo bar":
        var v10 float32 = 1.5
        _ = v10
    case "a":
        var v11 bool = false
        _ = v11
    default:
        v9 -= (v9 % ((uint64(v3)) | 1))
    }
    if len(v4) > 0 {
        v4[0] = 2147483647
    }
    v3 += (v3 % (((^(v3))) | 1))
    v1 -= 9223372036854775807
    var v12 int = 42
    var v13 uint8 = 1
    var v14 float32 = 1.5
    for fi15 := 0; fi15 < 18; fi15++ {
        var v16 float32 = v14
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi15
        _ = v16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
