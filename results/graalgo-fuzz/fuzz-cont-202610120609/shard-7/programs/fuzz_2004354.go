package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64 }
type GenS2 struct { F0 uint32 }

func main() {
    var v1 int = 4294967295
    var v2 int = 1
    var v3 int64 = -484
    var v4 int = (v1 | (^(v2)))
    v5 := []uint32{65535, 1, 2147483648}
    var v6 int32 = -2147483648
    switch (-(v1)) {
    case 1:
        var v7 int8 = 1
        _ = v7
    }
    v8 := GenS1{F0: (v3 ^ v3), F1: 2}
    _ = len(v5)
    v4 -= v2
    fmt.Println(v2, v3, v1)
    v9 := GenS1{F0: -9223372036854775808, F1: 18446744073709551615}
    for fi10 := 0; fi10 < 14; fi10++ {
        v1 = 127
        fmt.Println(v6)
        v11 := GenS1{F0: (v3 / (((v3 / (((v3 * v3)) | 1))) | 1)), F1: 4294967295}
        _ = fi10
        _ = v11
    }
    if true {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        var v12 bool = (589 > 42)
        _ = v12
    } else {
        v13 := GenS2{F0: 65535}
        _ = v13
    }
    v6 -= int32(v3)
    v14 := GenS1{F0: 360, F1: 18446744073709551615}
    for fi15 := 0; fi15 < 3; fi15++ {
        if true {
            if true {
                fmt.Println(v2, fi15, v4)
            } else {
                v16 := GenS1{F0: 42, F1: 21}
                var v17 uint64 = 0
                _ = v16
                _ = v17
            }
            v1 = 4294967295
        } else {
            var v18 uint32 = 65535
            v18 = (v18 | 340)
            fmt.Println(v6, v1, v4)
            _ = v18
        }
        switch 0 {
        case 32767:
            v6 += (v6 - -611)
        case 127:
            for fi19 := 0; fi19 < 6; fi19++ {
                var v20 bool = false
                v4 -= 128
                _ = fi19
                _ = v20
            }
        default:
            if false {
                var v21 uint32 = 351
                _ = v21
            } else {
                v22 := GenS1{F0: v3, F1: 18446744073709551615}
                v2 += fi15
                _ = v22
            }
        }
        _ = fi15
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
