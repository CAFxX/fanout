package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 int16; F1 int }

func main() {
    var v1 uint64 = 18446744073709551615
    v2 := make(map[string]uint32)
    v3 := []uint32{2147483648, 2147483647}
    var v4 uint16 = 0
    v5 := v2["\n\t"]
    for fi6 := 0; fi6 < 11; fi6++ {
        v7 := GenS1{F0: int32(fi6), F1: v5}
        var v8 float32 = ((3.14159 / ((float32(1.5) - 1e10) - 1.5)) * float32(float64(-2.25)))
        _ = fi6
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: 2147483647, F1: v5}
    if (42 <= -62) {
        v10 := GenS1{F0: -1, F1: (v5 + (^(v5)))}
        _ = v10
    }
    for fi11 := 0; fi11 < 19; fi11++ {
        fmt.Println(fi11)
        var v12 int64 = 0
        _ = fi11
        _ = v12
    }
    for fi13 := 0; fi13 < 15; fi13++ {
        var v14 uint64 = v1
        if ((v5 & (v5 + (v5 | (v5 << ((65535) & 63))))) > (v5 % (((v5 | (v5 / ((uint32(v1)) | 1)))) | 1))) {
            if (106 == 0) {
                for fi15 := 0; fi15 < 1; fi15++ {
                    v4 = v4
                    _ = len(v2)
                    switch (fi15 | (fi15 / ((fi13) | 1))) {
                    case 0:
                        v5 = v5
                    }
                    _ = fi15
                }
            }
            v5 = v5
        }
        _ = fi13
        _ = v14
    }
    for fi16 := 0; fi16 < 5; fi16++ {
        v1 = (v1 & v1)
        fmt.Println(v4, fi16, fi16)
        v17 := GenS1{F0: 0, F1: v5}
        _ = fi16
        _ = v17
    }
    switch "foo bar" {
    case "foo bar":
        v18 := GenS2{F0: -32768, F1: 0}
        _ = v18
    case "":
        var v19 uint32 = (v5 % (((v5 * (^(v5)))) | 1))
        _ = v19
    case "\n\t":
        for fi20 := 0; fi20 < 19; fi20++ {
            v1 = uint64(v4)
            fmt.Println(fi20, v1)
            _ = fi20
        }
    default:
        v2["\n\t"] = v5
    }
    var v21 bool = (!false)
    v1 += (^(v1))
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    switch 417 {
    case 2147483647:
        v1 += v1
    case 206:
        var v22 uint16 = (v4 | (^(v4)))
        _ = v22
    }
    v23 := GenS2{F0: 42, F1: 32768}
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v21
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
