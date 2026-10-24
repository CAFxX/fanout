package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int64; F1 uint; F2 uint8 }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint32; F1 int16; F2 uint16 }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int8 }

func main() {
    v1 := []uint32{982}
    var v2 uint8 = 1
    var v3 int16 = 1
    var v4 uint32 = 642
    v5 := []string{}
    v6 := []int{}
    _ = len(v5)
    if (831 <= 2) {
        if false {
            fmt.Println(v2, v3, v3)
        }
    } else {
        for fi7 := 0; fi7 < 20; fi7++ {
            v1 = append(v1, (v4 + 2147483648))
            var v8 int32 = 0
            fmt.Println(fi7)
            _ = fi7
            _ = v8
        }
        var v9 float64 = 1.5
        fmt.Println(v9, v9, v4)
        _ = v9
    }
    if ((255 == 2) || (false && (false || true))) {
        if true {
            var v10 string = "hello"
            _ = v10
        } else {
            fmt.Println(v4, v3)
            var v11 uint = 255
            switch "a" {
            case "hello":
                fmt.Println(v4, v2)
            case "":
                _ = len(v5)
            default:
                v12 := GenS1{F0: 2}
                _ = v12
            }
            _ = v11
        }
        switch 0 {
        case 32767:
            var v13 uint = uint(v2)
            _ = v13
        default:
            if (false && (!(!(false || (false || true))))) {
                fmt.Println(v2, v3, v4)
            }
        }
    }
    switch 2147483647 {
    case 0:
        for fi14 := 0; fi14 < 6; fi14++ {
            v15 := GenS2{F0: 215, F1: 42, F2: v2}
            if false {
                var v16 float32 = 0.1
                _ = v16
            }
            _ = fi14
            _ = v15
        }
    case -945:
        var v17 int16 = (v3 & (-(v3)))
        _ = v17
    case -2147483648:
        v18 := GenS3{F0: 65535}
        _ = v18
    default:
        v19 := GenS4{F0: (v4 * 42), F1: (v3 ^ v3), F2: 121}
        _ = v19
    }
    var v20 int8 = 0
    switch "" {
    case "hello":
        var v21 int8 = (v20 + v20)
        _ = v21
    case "a":
        var v22 uint32 = v4
        _ = v22
    default:
        v23 := GenS5{F0: 42}
        _ = v23
    }
    v24 := GenS6{F0: v20}
    for fi25 := 0; fi25 < 19; fi25++ {
        v6 = append(v6, fi25)
        var v26 int = (fi25 * 9223372036854775807)
        _ = fi25
        _ = v26
    }
    v4 -= v4
    v3 = (v3 >> ((1) & 63))
    var v27 int16 = v3
    fmt.Println(v3)
    var v28 int64 = 42
    var v29 uint64 = 42
    if true {
        if len(v6) > 0 {
            v6[0] = -9223372036854775808
        }
    }
    var v30 uint16 = 2
    if (((0.0 + 0.1) - 0.0) >= 0.0) {
        var v31 int = -2147483648
        var v32 int8 = (v20 + 26)
        v33 := GenS4{F0: uint32(v29), F1: (v27 * (v3 - int16(v31))), F2: (v30 % ((1) | 1))}
        _ = v31
        _ = v32
        _ = v33
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v2
    _ = v20
    _ = v24
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
    _ = v6
}
