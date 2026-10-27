package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint16; F1 int8; F2 uint }
type GenS3 struct { F0 int32; F1 uint32 }
type GenS4 struct { F0 int16; F1 string; F2 int16 }
type GenS5 struct { F0 int; F1 uint16; F2 int32 }
func genFn0() uint8 {
    for fi7 := 0; fi7 < 10; fi7++ {
        var v8 string = "foo bar"
        v9 := GenS1{F0: -9223372036854775808}
        fmt.Println(v8)
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 uint64 = 42
    v11 := GenS2{F0: 65535, F1: -25, F2: 305}
    _ = v10
    _ = v11
    return 127
}
func genFn1(p0 string, p1 uint) uint16 {
    switch 4294967295 {
    case 4294967295:
        var v12 uint64 = 4294967295
        _ = v12
    case 243:
        var v13 uint32 = 1
        _ = v13
    }
    v14 := GenS1{F0: 1}
    v15 := GenS2{F0: 608, F1: 100, F2: 4294967295}
    _ = p0
    _ = p1
    _ = v14
    _ = v15
    return 643
}

func main() {
    v1 := make(map[string]uint32)
    var v2 float32 = (((((1e10 * (1e10 / -2.25)) + float32((0.0 - (float64((float32(1.5) + ((0.0 + 0.1) - (0.0 / (0.0 - float32(0.1)))))) / ((0.0 + -2.25) * (((1e10 + (-2.25 / 1e10)) / 1.5) / 1e10)))))) + 1.5) - 0.1) + 3.14159)
    var v3 float64 = float64(v2)
    v4 := []uint32{}
    var v5 uint32 = 43
    v6 := []uint32{v5, v5, 2147483647}
    switch ("hello" + ("" + "\n\t")) {
    case "\n\t":
        v5 -= (v5 & (v5 >> ((4294967295) & 63)))
    case "foo bar":
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    default:
        fmt.Println(v2, v3)
    }
    fmt.Println(v5, v2, v3)
    v5 = (v5 % ((v5) | 1))
    switch ("foo bar" + ((("hello" + "\n\t") + "\n\t") + "foo bar")) {
    case "\n\t":
        var v16 float64 = v3
        _ = v16
    }
    v3 = (v3 - 3.14159)
    fmt.Println(v5)
    v2 -= float32(1.5)
    fmt.Println(v5, v3)
    fmt.Println(v3, v2)
    v17 := GenS3{F0: -2147483648, F1: (v5 + (v5 >> ((1) & 63)))}
    v18 := GenS4{F0: 448, F1: (("a" + "foo bar") + ("a" + "a")), F2: 42}
    for fi19 := 0; fi19 < 5; fi19++ {
        v20 := GenS3{F0: 42, F1: (v5 >> ((762) & 63))}
        _ = fi19
        _ = v20
    }
    v21 := GenS5{F0: 1, F1: 32767, F2: 2147483647}
    v22 := GenS1{F0: -1}
    if false {
        for fi23 := 0; fi23 < 20; fi23++ {
            for fi24 := 0; fi24 < 5; fi24++ {
                fmt.Println(v2)
                fmt.Println(v2, v5)
                _ = fi24
            }
            _ = fi23
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v17
    _ = v18
    _ = v2
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
