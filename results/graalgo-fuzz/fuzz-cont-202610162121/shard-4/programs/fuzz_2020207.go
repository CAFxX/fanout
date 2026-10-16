package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32; F2 int8 }
type GenS2 struct { F0 uint8; F1 uint32; F2 bool }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int32; F1 uint64; F2 bool }

func main() {
    var v1 int = 32767
    v2 := []uint32{702}
    var v3 int = v1
    v4 := GenS1{F0: -9223372036854775808, F1: -221, F2: 69}
    fmt.Println(v3)
    switch "\n\t" {
    case "foo bar":
        fmt.Println(v1, v1)
    case "":
        v5 := GenS1{F0: 0, F1: 889, F2: 87}
        _ = v5
    }
    v3 = 127
    switch "hello" {
    case "hello":
        fmt.Println(v3, v1)
    case "":
        if false {
            fmt.Println(v3)
            if ((float64(0.1) * -2.25) >= 3.14159) {
                if len(v2) > 0 {
                    v2[0] = 567
                }
            } else {
                v1 = 32768
            }
        } else {
            var v6 int32 = 252
            for fi7 := 0; fi7 < 7; fi7++ {
                if (true || (2 < 255)) {
                    var v8 float64 = ((0.0 + ((1.5 + 3.14159) * ((0.1 + 1e10) * (1e10 / ((3.14159 / 1e10) / 1.0))))) - 0.1)
                    _ = v8
                } else {
                    v6 -= (v6 >> ((0) & 63))
                    if false {
                        v3 = (v3 - fi7)
                    }
                }
                v9 := GenS1{F0: -1, F1: 42, F2: 100}
                _ = len(v2)
                _ = fi7
                _ = v9
            }
            _ = v6
        }
    default:
        v10 := GenS1{F0: 1, F1: 0, F2: -19}
        _ = v10
    }
    fmt.Println(v3, v1)
    var v11 uint = 0
    var v12 uint64 = 4294967295
    if false {
        v12 -= v12
        fmt.Println(v3)
        v13 := GenS2{F0: 42, F1: 65535, F2: (1 > uint32(v11))}
        _ = v13
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v14 uint = v11
        var v15 int = (v3 << ((v11) & 63))
        _ = v14
        _ = v15
    }
    if (-9223372036854775808 == 0) {
        var v16 bool = true
        _ = v16
    }
    if true {
        var v17 uint8 = 1
        v11 -= 18446744073709551615
        v18 := GenS1{F0: 513, F1: -1, F2: int8(v3)}
        _ = v17
        _ = v18
    } else {
        v3 = v1
        if (!false) {
            var v19 uint32 = 437
            v20 := GenS3{F0: -9223372036854775808}
            _ = v19
            _ = v20
        } else {
            v11 -= (v11 + 2)
        }
    }
    if len(v2) > 0 {
        v2[0] = 4294967295
    }
    for fi21 := 0; fi21 < 6; fi21++ {
        v1 += v3
        v22 := GenS4{F0: 1}
        for fi23 := 0; fi23 < 11; fi23++ {
            fmt.Println(v11, v1, v1)
            _ = fi23
        }
        _ = fi21
        _ = v22
    }
    switch v3 {
    case 32767:
        v24 := GenS5{F0: -1, F1: v12, F2: true}
        _ = v24
    default:
        var v25 int = 128
        _ = v25
    }
    if (!(2147483647 <= 0)) {
        switch "a" {
        case "":
            fmt.Println(v12, v12)
        default:
            v11 = (v11 | 667)
        }
    } else {
        var v26 int32 = -876
        _ = v26
    }
    switch 4294967295 {
    case 340:
        switch v1 {
        case 0:
            var v27 uint64 = 42
            _ = v27
        }
    case 2147483647:
        var v28 float32 = -2.25
        _ = v28
    default:
        v11 = v11
    }
    var v29 int64 = 746
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v29
    _ = v3
    _ = v4
}
