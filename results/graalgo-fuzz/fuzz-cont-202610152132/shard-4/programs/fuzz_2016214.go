package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 int64 }
type GenS2 struct { F0 uint; F1 int }
type GenS3 struct { F0 int32; F1 int64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int; F1 uint64; F2 bool }

func main() {
    v1 := []uint32{280, 171}
    var v2 int8 = 59
    var v3 string = ""
    var v4 float32 = 1.5
    var v5 uint16 = 65535
    var v6 uint = 1
    if (!true) {
        v6 = v6
    }
    v2 += (v2 % (((v2 ^ v2)) | 1))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v7 := GenS1{F0: 1, F1: v3, F2: -624}
    var v8 bool = (!false)
    fmt.Println(v6, v3)
    v9 := GenS2{F0: (^(v6)), F1: 2147483647}
    if len(v1) > 0 {
        v1[0] = 226
    }
    var v10 int = -488
    v2 -= (v2 << (((v6 + (^(v6)))) & 63))
    for fi11 := 0; fi11 < 4; fi11++ {
        for fi12 := 0; fi12 < 13; fi12++ {
            var v13 int = v10
            if v8 {
                v5 = (^(v5))
            }
            if len(v1) > 0 {
                v1[0] = 401
            }
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            if v8 {
                fmt.Println(v8, fi11, v3)
                v3 = v3
            }
            _ = fi12
            _ = v13
        }
        v14 := GenS3{F0: -2147483648, F1: -9223372036854775808}
        _ = fi11
        _ = v14
    }
    v15 := GenS3{F0: 2147483647, F1: 1}
    if (!v8) {
        if true {
            fmt.Println(v8, v10)
            _ = len(v1)
        } else {
            _ = len(v1)
        }
        v16 := GenS4{F0: 2}
        _ = v16
    } else {
        var v17 int32 = 174
        v18 := GenS4{F0: 105}
        var v19 bool = v8
        _ = v17
        _ = v18
        _ = v19
    }
    v1 = append(v1, 4294967295)
    v20 := GenS5{F0: (v10 & v10), F1: 1, F2: (uint32(v6) >= 4294967295)}
    v1 = append(v1, uint32(v5))
    fmt.Println(v2, v3)
    var v21 int = (-(v10))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v15
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
