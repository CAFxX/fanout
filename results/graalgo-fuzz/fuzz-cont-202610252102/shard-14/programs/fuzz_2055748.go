package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint }
type GenS2 struct { F0 int32; F1 int64; F2 bool }

func main() {
    var v1 float64 = -2.25
    var v2 int8 = 0
    var v3 int16 = 1
    v4 := []uint32{1, 65535}
    var v5 uint32 = 2147483647
    if true {
        v5 = 4294967295
        v2 -= (v2 ^ v2)
    }
    v2 += v2
    v2 = 86
    v6 := GenS1{F0: -79, F1: 18446744073709551615}
    for fi7 := 0; fi7 < 5; fi7++ {
        for fi8 := 0; fi8 < 4; fi8++ {
            v3 = 32767
            if false {
                v9 := GenS1{F0: (v2 >> ((18446744073709551615) & 63)), F1: 0}
                v5 = (v5 * v5)
                _ = v9
            }
            _ = fi8
        }
        for fi10 := 0; fi10 < 7; fi10++ {
            fmt.Println(v1, v3)
            var v11 uint64 = 4294967295
            var v12 int8 = (v2 << ((262) & 63))
            var v13 int8 = 127
            var v14 int = fi10
            var v15 uint = 1
            _ = fi10
            _ = v11
            _ = v12
            _ = v13
            _ = v14
            _ = v15
        }
        _ = fi7
    }
    if len(v4) > 0 {
        v4[0] = (v5 % ((65535) | 1))
    }
    v16 := GenS2{F0: 2147483647, F1: 0, F2: true}
    v3 = (v3 / (((v3 + 1000)) | 1))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
