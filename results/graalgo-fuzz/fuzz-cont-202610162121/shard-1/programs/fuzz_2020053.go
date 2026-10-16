package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int64; F2 int8 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32; F1 int16; F2 int16 }
type GenS4 struct { F0 bool }

func main() {
    v1 := []uint32{}
    v2 := make(map[uint32]uint32)
    v3 := []uint32{942, 388}
    delete(v2, 0)
    var v4 uint = 18446744073709551615
    fmt.Println(v4)
    switch "hello" {
    case "":
        fmt.Println(v4)
    case "hello":
        var v5 uint64 = 1
        _ = v5
    }
    v4 = (v4 / ((v4) | 1))
    v6 := GenS1{F0: 1, F1: 1, F2: int8(v4)}
    if true {
        if true {
            for fi7 := 0; fi7 < 1; fi7++ {
                v8 := GenS2{F0: 222}
                if (9223372036854775807 < fi7) {
                    fmt.Println(fi7, v4)
                }
                _ = fi7
                _ = v8
            }
            v4 = (v4 / (((v4 - 4294967295)) | 1))
        } else {
            for fi9 := 0; fi9 < 4; fi9++ {
                v2[4294967295] = 2147483647
                v4 += (v4 ^ 0)
                _ = fi9
            }
        }
        for fi10 := 0; fi10 < 2; fi10++ {
            var v11 uint32 = 0
            v12 := GenS2{F0: 18446744073709551615}
            _ = fi10
            _ = v11
            _ = v12
        }
        v4 = (v4 & (v4 ^ (^(v4))))
    } else {
        fmt.Println(v4)
        var v13 bool = ((1e10 <= (3.14159 / 1e10)) && (-75 > -128))
        _ = v13
    }
    for fi14 := 0; fi14 < 15; fi14++ {
        fmt.Println(fi14, fi14)
        v4 -= (v4 / ((4294967295) | 1))
        for fi15 := 0; fi15 < 7; fi15++ {
            v16 := GenS3{F0: 2147483647, F1: 1, F2: -32768}
            if len(v1) > 0 {
                v1[0] = 633
            }
            _ = fi15
            _ = v16
        }
        _ = fi14
    }
    v17 := GenS4{F0: (!((2 <= 42) && (38 > 1)))}
    var v18 int16 = 1
    var v19 uint = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v6
}
