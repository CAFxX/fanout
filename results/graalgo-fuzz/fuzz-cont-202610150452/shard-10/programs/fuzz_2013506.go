package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 int8; F1 int; F2 int8 }
type GenS3 struct { F0 uint64; F1 uint16 }
type GenS4 struct { F0 bool }
func genFn0(p0 int, p1 int16) uint16 {
    switch "a" {
    case "hello":
        v7 := []int{42, (-(p0))}
        _ = v7
    default:
        v8 := GenS1{F0: 42}
        _ = v8
    }
    for fi9 := 0; fi9 < 5; fi9++ {
        for fi10 := 0; fi10 < 6; fi10++ {
            var v11 float32 = ((3.14159 / 3.14159) * ((1e10 - 1e10) + 0.0))
            _ = fi10
            _ = v11
        }
        var v12 int8 = 74
        var v13 int = (^(p0))
        _ = fi9
        _ = v12
        _ = v13
    }
    _ = p0
    _ = p1
    return 136
}
func genFn1(p0 int16, p1 int8, p2 int16) uint16 {
    for fi14 := 0; fi14 < 1; fi14++ {
        for fi15 := 0; fi15 < 7; fi15++ {
            v16 := GenS2{F0: 1, F1: 127, F2: (-(p1))}
            fmt.Println(fi15, p2)
            fmt.Println(fi15)
            _ = fi15
            _ = v16
        }
        _ = fi14
    }
    v17 := GenS1{F0: -9223372036854775808}
    p1 = p1
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    return 42
}

func main() {
    var v1 uint64 = 0
    v2 := []uint32{}
    v3 := make(map[uint32]int)
    var v4 int16 = 1
    v5 := make(map[int]uint32)
    v6 := make(map[uint32]uint32)
    if (-128 != -8) {
        v1 -= v1
        if (65535 == 714) {
            switch 32767 {
            case 32767:
                v18 := GenS1{F0: -1}
                _ = v18
            case 2147483647:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case 128:
                for fi19 := 0; fi19 < 5; fi19++ {
                    for fi20 := 0; fi20 < 12; fi20++ {
                        if (!(-2.25 <= float64(float32(0.1)))) {
                            v21 := GenS3{F0: 42, F1: genFn1(v4, 42, (-(v4)))}
                            _ = v21
                        } else {
                            var v22 float64 = 0.0
                            var v23 uint32 = 0
                            fmt.Println(v22, fi19)
                            v23 += (v23 & v23)
                            _ = v22
                            _ = v23
                        }
                        v24 := v3[114]
                        v1 = v1
                        _ = fi20
                        _ = v24
                    }
                    delete(v6, 293)
                    var v25 float64 = -2.25
                    _ = fi19
                    _ = v25
                }
            default:
                var v26 int16 = (-(v4))
                _ = v26
            }
            fmt.Println(v4, v4)
        }
        v27 := GenS3{F0: 460, F1: 65535}
        _ = v27
    } else {
        v4 += (^(v4))
    }
    v2 = append(v2, 257)
    v1 = (v1 ^ v1)
    switch "foo bar" {
    case "\n\t":
        fmt.Println(v1)
    case "a":
        var v28 int32 = -2147483648
        _ = v28
    }
    v4 = v4
    v29 := GenS2{F0: 57, F1: 465, F2: -104}
    delete(v6, 1)
    var v30 float32 = (1.5 + ((1e10 - 1e10) - 3.14159))
    v31 := GenS4{F0: (!((-2.25 < (1e10 + (3.14159 - 0.0))) || (!true)))}
    if (true || false) {
        fmt.Println(v4)
    }
    switch -9223372036854775808 {
    case 32768:
        v32 := GenS2{F0: -29, F1: -2147483648, F2: -128}
        _ = v32
    default:
        v30 -= -2.25
    }
    fmt.Println(v1, v4, v1)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
}
