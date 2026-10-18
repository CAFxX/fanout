package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool; F2 int16 }
type GenS2 struct { F0 int64; F1 uint8; F2 uint }
type GenS3 struct { F0 int8; F1 int64; F2 uint64 }
type GenS4 struct { F0 uint; F1 uint }
type GenS5 struct { F0 int32 }
func genFn0(p0 string, p1 uint8) uint16 {
    var v4 int64 = 0
    v5 := []int{-567}
    _ = p0
    _ = p1
    _ = v4
    _ = v5
    return 42
}

func main() {
    var v1 int8 = -6
    v2 := []int{667, -1, 127}
    var v3 int64 = 0
    v6 := GenS1{F0: 18446744073709551615, F1: (444 != 253), F2: -485}
    if (true || true) {
        v7 := GenS2{F0: 529, F1: 0, F2: 0}
        _ = v7
    } else {
        if true {
            var v8 uint32 = 1
            _ = v8
        }
    }
    var v9 int16 = 1
    switch "\n\t" {
    case "\n\t":
        v9 = -34
    case "foo bar":
        for fi10 := 0; fi10 < 6; fi10++ {
            var v11 float64 = 1e10
            _ = fi10
            _ = v11
        }
    case "":
        v12 := GenS1{F0: 1, F1: true, F2: (v9 << ((42) & 63))}
        _ = v12
    }
    for fi13 := 0; fi13 < 7; fi13++ {
        v9 -= (v9 / ((v9) | 1))
        v14 := GenS1{F0: 2, F1: false, F2: v9}
        _ = fi13
        _ = v14
    }
    v9 = 0
    for fi15 := 0; fi15 < 2; fi15++ {
        if false {
            switch fi15 {
            case -9223372036854775808:
                v16 := GenS3{F0: 77, F1: v3, F2: 237}
                _ = v16
            default:
                fmt.Println(fi15, v3, v1)
            }
            v9 = (v9 - (v9 << ((4294967295) & 63)))
            fmt.Println(v1, v3, v3)
        } else {
            for fi17 := 0; fi17 < 6; fi17++ {
                _ = len(v2)
                v3 = (v3 + -1)
                var v18 int32 = int32(v3)
                _ = fi17
                _ = v18
            }
            if len(v2) > 0 {
                v2[0] = fi15
            }
            if (-9223372036854775808 >= (-(v3))) {
                var v19 uint64 = 713
                _ = v19
            } else {
                v20 := GenS3{F0: v1, F1: (v3 % (((^(v3))) | 1)), F2: 395}
                var v21 uint16 = 1
                var v22 uint = 65535
                _ = v20
                _ = v21
                _ = v22
            }
        }
        _ = fi15
    }
    v9 += v9
    if (!false) {
        var v23 bool = true
        _ = v23
    } else {
        v3 = (v3 / ((v3) | 1))
        var v24 int8 = v1
        _ = v24
    }
    v25 := GenS4{F0: 715, F1: 42}
    for fi26 := 0; fi26 < 13; fi26++ {
        fmt.Println(v1)
        v27 := GenS5{F0: int32(v3)}
        var v28 uint16 = 1
        _ = fi26
        _ = v27
        _ = v28
    }
    for fi29 := 0; fi29 < 15; fi29++ {
        var v30 int16 = 1
        fmt.Println(v1)
        if false {
            if len(v2) > 0 {
                v2[0] = fi29
            }
            if false {
                var v31 float32 = 0.0
                var v32 uint = 809
                v30 = 808
                _ = v31
                _ = v32
            }
            if (v3 >= (v3 - (^(v3)))) {
                v3 = v3
            }
        }
        _ = fi29
        _ = v30
    }
    v1 += v1
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v25
    _ = v3
    _ = v6
    _ = v9
}
