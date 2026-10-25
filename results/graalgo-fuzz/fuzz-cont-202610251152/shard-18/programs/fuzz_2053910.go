package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int64; F1 uint64 }
func genFn0() int16 {
    v4 := GenS1{F0: (("foo bar" + "foo bar") + "hello")}
    var v5 int64 = 0
    _ = v4
    _ = v5
    return 563
}
func genFn1(p0 string, p1 uint16) int32 {
    if false {
        p1 = p1
    } else {
        p1 += p1
        v6 := GenS1{F0: "a"}
        _ = v6
    }
    var v7 bool = (127 == 1)
    p0 = (p0 + p0)
    _ = p0
    _ = p1
    _ = v7
    return 2147483647
}

func main() {
    var v1 uint8 = 22
    var v2 float64 = 1.5
    v3 := make(map[uint32]uint32)
    fmt.Println(v2, v2)
    var v8 int64 = -9223372036854775808
    if false {
        if ("foo bar" != "hello") {
            var v9 int = 425
            _ = v9
        }
        v1 += (v1 * v1)
    } else {
        var v10 uint64 = 0
        v1 += 255
        var v11 uint64 = (v10 | v10)
        _ = v10
        _ = v11
    }
    if false {
        v1 -= v1
    }
    switch 4294967295 {
    case -902:
        v2 -= -2.25
    case -12:
        var v12 float64 = float64((1.5 - ((((0.0 * 1.5) / 3.14159) + ((1.5 * float32(float64(1e10))) + float32(3.14159))) + -2.25)))
        _ = v12
    }
    v8 -= v8
    var v13 int32 = -331
    if (2 < 2147483647) {
        v14 := v3[0]
        if true {
            var v15 int16 = genFn0()
            fmt.Println(v2)
            v15 -= v15
            _ = v15
        }
        _ = v14
    }
    v8 = v8
    var v16 int8 = -55
    switch int(v13) {
    case -9223372036854775808:
        v1 = v1
    case -213:
        v17 := GenS2{F0: (^(v13))}
        _ = v17
    case 127:
        var v18 int64 = (v8 << ((4294967295) & 63))
        _ = v18
    default:
        var v19 uint8 = v1
        _ = v19
    }
    if (42 > v1) {
        if true {
            var v20 int = 1
            v21 := GenS2{F0: int32(v16)}
            _ = v20
            _ = v21
        } else {
            for fi22 := 0; fi22 < 6; fi22++ {
                var v23 int16 = 32767
                v16 = (v16 ^ v16)
                fmt.Println(v13, v1, v13)
                _ = fi22
                _ = v23
            }
            if true {
                var v24 int64 = 9223372036854775807
                _ = v24
            } else {
                switch "\n\t" {
                case "foo bar":
                    v25 := GenS3{F0: (v8 ^ (v8 << ((1) & 63))), F1: 856}
                    _ = v25
                default:
                    if (!false) {
                        v1 = v1
                        for fi26 := 0; fi26 < 6; fi26++ {
                            fmt.Println(v16, v1)
                            _ = fi26
                        }
                    } else {
                        for fi27 := 0; fi27 < 14; fi27++ {
                            v13 = 1
                            _ = fi27
                        }
                    }
                }
            }
        }
        v16 = v16
    } else {
        var v28 int32 = v13
        for fi29 := 0; fi29 < 13; fi29++ {
            v16 -= (v16 - v16)
            _ = fi29
        }
        _ = v28
    }
    if false {
        v30 := GenS2{F0: v13}
        _ = v30
    } else {
        var v31 uint64 = 18446744073709551615
        v32 := GenS1{F0: ("foo bar" + (("\n\t" + "hello") + ("a" + "a")))}
        _ = v31
        _ = v32
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v8
}
