package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int; F1 int }
func genFn0(p0 uint, p1 int32) int {
    v4 := GenS1{F0: 1}
    for fi5 := 0; fi5 < 8; fi5++ {
        p0 = p0
        _ = fi5
    }
    switch 128 {
    case 9223372036854775807:
        var v6 int16 = -32768
        _ = v6
    case 42:
        var v7 string = "hello"
        _ = v7
    case -578:
        fmt.Println(p0, p0)
    default:
        switch -9223372036854775808 {
        case 2147483647:
            v8 := GenS1{F0: uint64(p0)}
            _ = v8
        case 624:
            for fi9 := 0; fi9 < 5; fi9++ {
                fmt.Println(fi9)
                _ = fi9
            }
        case 128:
            if (65535 < (p0 << ((p0) & 63))) {
                v10 := []int{int(p0)}
                _ = v10
            } else {
                fmt.Println(p0)
                if false {
                    p0 = (p0 >> ((p0) & 63))
                    if (42 < -76) {
                        v11 := []int{966, -2147483648}
                        _ = v11
                    }
                }
                p0 += p0
            }
        default:
            v12 := GenS1{F0: 461}
            _ = v12
        }
    }
    for fi13 := 0; fi13 < 9; fi13++ {
        p1 = (p1 >> ((p0) & 63))
        p1 += (p1 >> ((2) & 63))
        _ = fi13
    }
    if false {
        p0 += 2
    }
    _ = p0
    _ = p1
    _ = v4
    return -473
}

func main() {
    var v1 int = 127
    v2 := []int{v1, (-(v1)), v1}
    var v3 int8 = 12
    v3 = v3
    v3 = v3
    fmt.Println(v1, v1)
    if (255 != 95) {
        var v14 bool = (((float32(-2.25) / (1.5 * 1e10)) != (1e10 * (0.0 - 1.5))) || (18446744073709551615 == 65535))
        _ = v14
    }
    v15 := GenS1{F0: 165}
    for fi16 := 0; fi16 < 8; fi16++ {
        v3 = 21
        fmt.Println(v3)
        _ = fi16
    }
    switch (^(v1)) {
    case 128:
        for fi17 := 0; fi17 < 16; fi17++ {
            for fi18 := 0; fi18 < 16; fi18++ {
                var v19 float32 = 1e10
                _ = fi18
                _ = v19
            }
            _ = fi17
        }
    case 0:
        v20 := GenS1{F0: 2}
        _ = v20
    case 2147483647:
        v3 -= (v3 * v3)
    default:
        v21 := GenS1{F0: 18446744073709551615}
        _ = v21
    }
    v22 := GenS2{F0: 255}
    v3 += (v3 ^ v3)
    v23 := GenS3{F0: genFn0(0, 42), F1: (^(v1))}
    var v24 bool = (float32(-2.25) > 1.5)
    v25 := GenS1{F0: 42}
    v26 := GenS3{F0: genFn0(839, -1), F1: genFn0(255, 448)}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
}
