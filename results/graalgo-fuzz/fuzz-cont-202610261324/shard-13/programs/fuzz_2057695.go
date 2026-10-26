package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint32; F1 int32; F2 int8 }
type GenS4 struct { F0 int8; F1 uint64; F2 uint8 }
func genFn0(p0 int16, p1 uint32) uint32 {
    v7 := GenS1{F0: 80, F1: (p0 >> ((18446744073709551615) & 63))}
    var v8 uint8 = 2
    p0 = p0
    for fi9 := 0; fi9 < 10; fi9++ {
        p0 = p0
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 793
}

func main() {
    var v1 int8 = 1
    var v2 int32 = -2147483648
    v3 := make(map[string]int)
    v4 := []uint32{1, 65535, 4294967295}
    var v5 uint8 = 42
    var v6 uint = 592
    switch "a" {
    case "\n\t":
        v10 := GenS1{F0: (v5 ^ (v5 & v5)), F1: 0}
        _ = v10
    case "a":
        var v11 int32 = (v2 - v2)
        _ = v11
    default:
        if true {
            v2 -= 0
            v12 := GenS2{F0: 0}
            _ = v12
        }
    }
    for fi13 := 0; fi13 < 7; fi13++ {
        for fi14 := 0; fi14 < 9; fi14++ {
            v15 := GenS1{F0: uint8(v6), F1: -659}
            switch "\n\t" {
            case "foo bar":
                fmt.Println(v2)
            default:
                v16 := GenS2{F0: genFn0(-32768, 4294967295)}
                _ = v16
            }
            _ = fi14
            _ = v15
        }
        _ = fi13
    }
    for fi17 := 0; fi17 < 7; fi17++ {
        if true {
            switch -15 {
            case 2147483647:
                v18 := GenS1{F0: (v5 / ((v5) | 1)), F1: 0}
                _ = v18
            case 127:
                v2 = (v2 * v2)
            default:
                var v19 float64 = 1e10
                _ = v19
            }
            for fi20 := 0; fi20 < 17; fi20++ {
                v1 = (v1 % (((v1 ^ 0)) | 1))
                _ = fi20
            }
            switch fi17 {
            case 4294967295:
                var v21 uint64 = 398
                _ = v21
            case -178:
                var v22 int16 = 1
                _ = v22
            default:
                v1 += 127
            }
        }
        var v23 uint32 = 0
        _ = fi17
        _ = v23
    }
    var v24 uint64 = 740
    v25 := GenS3{F0: 65535, F1: (v2 & v2), F2: v1}
    var v26 uint = (v6 - (v6 >> ((4294967295) & 63)))
    for fi27 := 0; fi27 < 10; fi27++ {
        if (!true) {
            v6 -= v6
            var v28 float32 = 3.14159
            _ = v28
        }
        v29 := GenS4{F0: v1, F1: (v24 | (v24 + (v24 / ((v24) | 1)))), F2: 42}
        _ = fi27
        _ = v29
    }
    v5 += v5
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
