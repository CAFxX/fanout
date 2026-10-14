package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
func genFn0(p0 uint64, p1 uint, p2 uint64) int8 {
    switch 127 {
    case 42:
        switch ("foo bar" + "foo bar") {
        case "a":
            fmt.Println(p0, p0, p2)
        }
    case -2147483648:
        p1 = 65535
    case 128:
        p0 = 417
    }
    var v6 int8 = 83
    if (int64(p1) != 42) {
        p0 += 507
    }
    fmt.Println(p0, v6)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return (v6 * v6)
}
func genFn1(p0 uint64, p1 uint8) uint16 {
    var v7 uint16 = 492
    v8 := GenS1{F0: ((("a" + "hello") + ("a" + ("" + "foo bar"))) == "hello"), F1: 0}
    var v9 float64 = 0.1
    var v10 bool = (65535 == 2147483648)
    v11 := []int{679}
    _ = p0
    _ = p1
    _ = v10
    _ = v11
    _ = v7
    _ = v8
    _ = v9
    return v7
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []int{}
    var v3 uint16 = 0
    var v4 int = 1
    v5 := []int{v4, (v4 + v4), (v4 >> ((130) & 63))}
    delete(v1, 4294967295)
    _ = len(v1)
    for fi12 := 0; fi12 < 17; fi12++ {
        v5 = append(v5, 32767)
        _ = fi12
    }
    var v13 int64 = 9223372036854775807
    var v14 uint = 18446744073709551615
    v15 := GenS1{F0: (false || false), F1: 2147483647}
    v16 := GenS1{F0: (true && (0 > 1)), F1: 2}
    if (float64(float32(-2.25)) < (0.1 * 0.1)) {
        switch "hello" {
        case "":
            var v17 float64 = 3.14159
            _ = v17
        case "hello":
            var v18 int16 = 475
            _ = v18
        case "foo bar":
            var v19 int64 = 0
            _ = v19
        default:
            if ((637 <= 291) && (v3 == v3)) {
                var v20 int8 = 100
                for fi21 := 0; fi21 < 9; fi21++ {
                    var v22 int = 32768
                    if false {
                        v20 = v20
                    }
                    _ = fi21
                    _ = v22
                }
                v20 = (v20 & v20)
                _ = v20
            }
        }
        v23 := GenS1{F0: false, F1: 834}
        switch "hello" {
        case "foo bar":
            v13 = 1
        case "":
            if true {
                delete(v1, uint32(v14))
                fmt.Println(v13)
                var v24 int32 = 0
                _ = v24
            }
        case "hello":
            for fi25 := 0; fi25 < 20; fi25++ {
                for fi26 := 0; fi26 < 18; fi26++ {
                    _ = len(v5)
                    if false {
                        v14 = v14
                        switch 4294967295 {
                        case -1:
                            v4 = fi26
                        case 32767:
                            v14 -= v14
                        case -9223372036854775808:
                            v4 -= v4
                            var v27 int64 = (^(v13))
                            _ = v27
                        }
                        if len(v2) > 0 {
                            v2[0] = (v4 / ((v4) | 1))
                        }
                    }
                    var v28 uint32 = 2147483647
                    _ = fi26
                    _ = v28
                }
                _ = len(v1)
                _ = fi25
            }
        }
        _ = v23
    }
    var v29 int = (v4 - 1)
    v30 := GenS1{F0: ((2147483647 < 205) && ((v14 * v14) <= v14)), F1: 4294967295}
    for fi31 := 0; fi31 < 14; fi31++ {
        var v32 int32 = 1
        v13 -= (v13 ^ v13)
        _ = fi31
        _ = v32
    }
    var v33 string = "a"
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v33
    _ = v4
    _ = v5
}
