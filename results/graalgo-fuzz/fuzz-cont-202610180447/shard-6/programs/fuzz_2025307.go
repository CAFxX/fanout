package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }

func main() {
    v1 := []uint32{65535, 245, 4294967295, 4294967295}
    v2 := []uint32{4294967295}
    var v3 float64 = float64(3.14159)
    var v4 int = 4294967295
    var v5 uint32 = 0
    var v6 bool = false
    if ((v4 ^ (v4 / (((v4 ^ v4)) | 1))) != (v4 & v4)) {
        switch v4 {
        case 128:
            var v7 uint64 = 544
            _ = v7
        case -9223372036854775808:
            var v8 int16 = -1
            _ = v8
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        if false {
            fmt.Println(v5, v4)
        } else {
            v6 = false
        }
    }
    v4 = v4
    if ((!true) && (v3 >= v3)) {
        var v9 int32 = 1
        fmt.Println(v4)
        v4 = -9223372036854775808
        _ = v9
    }
    v2 = append(v2, (v5 - v5))
    v5 = 2147483647
    var v10 int64 = -1
    if (v6 || v6) {
        v6 = ((v5 * v5) == (v5 + 722))
    }
    var v11 bool = false
    if (!v6) {
        var v12 int16 = -32768
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v12
    }
    _ = len(v2)
    v4 -= (v4 >> ((18446744073709551615) & 63))
    if false {
        v13 := GenS1{F0: 0, F1: v11}
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        _ = v13
    }
    var v14 uint8 = 255
    v2 = append(v2, v5)
    v6 = v11
    var v15 int32 = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
