package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 bool }
func genFn0(p0 int32, p1 int32) uint8 {
    fmt.Println(p1, p1)
    var v5 uint32 = 0
    var v6 uint16 = 621
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 255
}
func genFn1(p0 int, p1 int32, p2 int32) uint8 {
    for fi7 := 0; fi7 < 16; fi7++ {
        var v8 int32 = (p1 ^ (p1 >> ((1) & 63)))
        p0 = fi7
        _ = fi7
        _ = v8
    }
    var v9 int16 = -868
    v10 := GenS1{F0: 100, F1: int64(p0)}
    var v11 bool = false
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v9
    return 128
}

func main() {
    v1 := []string{""}
    var v2 uint16 = 2
    var v3 uint32 = 4294967295
    var v4 uint64 = 0
    v3 += v3
    fmt.Println(v3)
    fmt.Println(v4, v4)
    if true {
        var v12 float64 = -2.25
        for fi13 := 0; fi13 < 19; fi13++ {
            v12 += v12
            _ = fi13
        }
        var v14 int = 128
        _ = v12
        _ = v14
    } else {
        for fi15 := 0; fi15 < 11; fi15++ {
            v2 = 128
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            for fi16 := 0; fi16 < 13; fi16++ {
                v17 := GenS2{F0: ((^(v4)) < (v4 & 42))}
                v2 += (^(v2))
                var v18 int64 = -9223372036854775808
                _ = fi16
                _ = v17
                _ = v18
            }
            _ = fi15
        }
    }
    for fi19 := 0; fi19 < 12; fi19++ {
        v3 += (v3 >> ((18446744073709551615) & 63))
        v20 := GenS1{F0: 109, F1: int64(v3)}
        _ = fi19
        _ = v20
    }
    v1 = append(v1, "\n\t")
    var v21 int16 = 32767
    v22 := GenS2{F0: true}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
}
