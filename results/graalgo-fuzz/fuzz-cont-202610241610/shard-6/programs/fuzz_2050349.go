package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8 }
func genFn0(p0 uint32, p1 uint8) uint32 {
    var v5 bool = true
    fmt.Println(p1)
    var v6 float32 = 3.14159
    fmt.Println(p0, v5)
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return (p0 >> ((792) & 63))
}

func main() {
    v1 := []uint32{4294967295}
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]bool)
    var v4 int64 = 708
    for fi7 := 0; fi7 < 9; fi7++ {
        switch fi7 {
        case 32767:
            v4 -= v4
        case 9223372036854775807:
            var v8 int16 = -32768
            _ = v8
        default:
            v4 -= -164
        }
        var v9 float64 = ((-2.25 + -2.25) * (3.14159 * 0.0))
        _ = fi7
        _ = v9
    }
    var v10 int16 = 1
    v10 -= v10
    if true {
        v11 := GenS1{F0: 0, F1: 2}
        v12 := GenS1{F0: 1, F1: 43}
        v4 -= (v4 >> ((18446744073709551615) & 63))
        _ = v11
        _ = v12
    }
    v1 = append(v1, 4294967295)
    if false {
        if false {
            var v13 uint64 = 405
            fmt.Println(v10)
            _ = v13
        } else {
            v4 = v4
        }
        var v14 uint32 = 809
        for fi15 := 0; fi15 < 19; fi15++ {
            switch "foo bar" {
            case "foo bar":
                var v16 int64 = (v4 + v4)
                _ = v16
            case "\n\t":
                v17 := GenS1{F0: 603, F1: 211}
                _ = v17
            case "":
                v18 := GenS1{F0: 255, F1: 42}
                _ = v18
            }
            _ = fi15
        }
        _ = v14
    } else {
        delete(v2, 2147483648)
        var v19 float32 = 3.14159
        _ = v19
    }
    for fi20 := 0; fi20 < 1; fi20++ {
        v4 = v4
        if ((4294967295 == 4294967295) && true) {
            v21 := v3["\n\t"]
            _ = v21
        }
        _ = fi20
    }
    var v22 int8 = -128
    var v23 float32 = 1.5
    var v24 int16 = v10
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v2
    _ = v22
    _ = v23
    _ = v24
    _ = v3
    _ = v4
}
