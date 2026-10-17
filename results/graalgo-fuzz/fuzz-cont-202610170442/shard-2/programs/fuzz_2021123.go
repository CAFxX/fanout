package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int8; F1 uint }
func genFn0(p0 int16, p1 int64) uint8 {
    if (!(127 > 255)) {
        v5 := GenS1{F0: 42}
        v6 := GenS1{F0: 650}
        _ = v5
        _ = v6
    }
    var v7 float32 = 0.1
    switch -1 {
    case 42:
        for fi8 := 0; fi8 < 12; fi8++ {
            var v9 float64 = 1.5
            v10 := GenS1{F0: 931}
            _ = fi8
            _ = v10
            _ = v9
        }
    default:
        var v11 float64 = -2.25
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v7
    return 234
}
func genFn1() uint {
    var v12 uint32 = 1
    v13 := GenS2{F0: 100, F1: 255}
    fmt.Println(v12)
    _ = v12
    _ = v13
    return 65535
}

func main() {
    var v1 int64 = 1
    var v2 uint32 = 2147483647
    var v3 int16 = 32767
    var v4 float32 = 0.1
    fmt.Println(v3, v1)
    if (float32(1e10) != v4) {
        var v14 uint16 = 32767
        _ = v14
    }
    var v15 int64 = (v1 >> ((627) & 63))
    fmt.Println(v1, v15, v3)
    var v16 uint16 = 65535
    for fi17 := 0; fi17 < 4; fi17++ {
        if true {
            fmt.Println(v16)
        } else {
            var v18 uint32 = 2
            v2 = 795
            v15 -= int64(v2)
            _ = v18
        }
        var v19 int8 = int8(v2)
        _ = fi17
        _ = v19
    }
    v20 := GenS1{F0: 440}
    fmt.Println(v15, v15)
    for fi21 := 0; fi21 < 17; fi21++ {
        v22 := []int{(fi21 << ((65535) & 63)), fi21, 2147483647}
        _ = fi21
        _ = v22
    }
    v2 = (v2 << ((2) & 63))
    for fi23 := 0; fi23 < 15; fi23++ {
        if (false && false) {
            for fi24 := 0; fi24 < 18; fi24++ {
                fmt.Println(v16, v15, v4)
                _ = fi24
            }
        } else {
            switch ("a" + "foo bar") {
            case "":
                v25 := []int{(fi23 * (fi23 | 128)), (fi23 >> ((18446744073709551615) & 63))}
                _ = v25
            case "a":
                v3 -= v3
            default:
                v16 += v16
            }
        }
        for fi26 := 0; fi26 < 2; fi26++ {
            var v27 int32 = 2147483647
            v15 -= v1
            _ = fi26
            _ = v27
        }
        v4 += v4
        _ = fi23
    }
    v28 := GenS2{F0: 100, F1: 18446744073709551615}
    fmt.Println(v2, v16)
    fmt.Println(v15, v16, v4)
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v28
    _ = v3
    _ = v4
}
