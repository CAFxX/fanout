package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int16; F2 uint16 }
type GenS2 struct { F0 int32; F1 int64; F2 int }
func genFn0(p0 int32) int64 {
    var v4 uint = 0
    p0 -= 792
    fmt.Println(p0, p0)
    _ = p0
    _ = v4
    return -1
}

func main() {
    var v1 float32 = -2.25
    var v2 int8 = -1
    var v3 int64 = -9223372036854775808
    for fi5 := 0; fi5 < 16; fi5++ {
        for fi6 := 0; fi6 < 17; fi6++ {
            v2 = v2
            _ = fi6
        }
        _ = fi5
    }
    var v7 int32 = -1
    var v8 int = 42
    switch int(v3) {
    case 32768:
        v9 := GenS1{F0: 478, F1: 797, F2: 925}
        _ = v9
    case -808:
        v1 = -2.25
    case 2147483647:
        var v10 uint16 = 65535
        _ = v10
    default:
        var v11 int8 = (-(v2))
        _ = v11
    }
    v3 = v3
    v12 := []int{}
    for fi13 := 0; fi13 < 12; fi13++ {
        for fi14 := 0; fi14 < 6; fi14++ {
            if true {
                v15 := GenS1{F0: 2, F1: 42, F2: 0}
                var v16 string = "foo bar"
                v2 = v2
                _ = v15
                _ = v16
            }
            v17 := GenS1{F0: 65535, F1: -32768, F2: 255}
            _ = fi14
            _ = v17
        }
        fmt.Println(v1)
        _ = fi13
    }
    switch v8 {
    case 9223372036854775807:
        var v18 float64 = 0.1
        _ = v18
    default:
        if false {
            v3 = v3
            for fi19 := 0; fi19 < 6; fi19++ {
                switch (^(v8)) {
                case -452:
                    v7 -= (v7 - 1)
                case 139:
                    v3 -= (v3 << ((1) & 63))
                }
                _ = fi19
            }
            for fi20 := 0; fi20 < 6; fi20++ {
                if len(v12) >= 2 {
                    v12 = v12[1:len(v12)]
                }
                _ = fi20
            }
        }
    }
    v21 := GenS1{F0: 1, F1: 42, F2: 0}
    if (((v7 & -1) == (v7 >> ((425) & 63))) || true) {
        v22 := GenS2{F0: (v7 % ((v7) | 1)), F1: (v3 - (v3 >> ((18446744073709551615) & 63))), F2: v8}
        v3 += genFn0(-2147483648)
        _ = v22
    } else {
        for fi23 := 0; fi23 < 10; fi23++ {
            _ = len(v12)
            fmt.Println(v7, fi23, v7)
            fmt.Println(fi23, v3, fi23)
            _ = fi23
        }
        v24 := GenS1{F0: 241, F1: 1, F2: 276}
        if len(v12) >= 2 {
            v12 = v12[1:len(v12)]
        }
        _ = v24
    }
    v3 -= (v3 & -374)
    var v25 int16 = 1
    for fi26 := 0; fi26 < 19; fi26++ {
        switch 42 {
        case -2147483648:
            v12 = append(v12, fi26)
        case 128:
            v7 -= (v7 | (v7 + v7))
        default:
            var v27 uint32 = 0
            _ = v27
        }
        _ = fi26
    }
    fmt.Println(v25, v25)
    var v28 int = 127
    _ = len(v12)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v12
    _ = v2
    _ = v21
    _ = v25
    _ = v28
    _ = v3
    _ = v7
    _ = v8
}
