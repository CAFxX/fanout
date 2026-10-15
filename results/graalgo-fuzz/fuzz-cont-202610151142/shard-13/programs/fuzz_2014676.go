package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint32; F1 int64; F2 string }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint; F1 int }
type GenS5 struct { F0 int8; F1 int8; F2 string }
func genFn0(p0 int64) uint32 {
    v6 := GenS1{F0: 9223372036854775807}
    v7 := GenS2{F0: 42, F1: (p0 | (p0 << ((42) & 63))), F2: ""}
    var v8 int64 = (p0 % (((p0 & (p0 / ((p0) | 1)))) | 1))
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    return 775
}

func main() {
    var v1 uint64 = 18446744073709551615
    var v2 int = 32767
    var v3 int32 = -911
    var v4 uint8 = 201
    var v5 float32 = (float32(0.1) / 1.0)
    if false {
        for fi9 := 0; fi9 < 6; fi9++ {
            v10 := GenS3{F0: -128}
            var v11 uint8 = (^(v4))
            var v12 float32 = (v5 + (0.1 / v5))
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
        }
        v13 := GenS2{F0: uint32(v4), F1: -450, F2: "a"}
        _ = v13
    } else {
        v2 = (^(v2))
        for fi14 := 0; fi14 < 12; fi14++ {
            if (fi14 < -1) {
                if (false && true) {
                    for fi15 := 0; fi15 < 2; fi15++ {
                        v3 = v3
                        _ = fi15
                    }
                    for fi16 := 0; fi16 < 10; fi16++ {
                        switch ("\n\t" + ("\n\t" + "hello")) {
                        case "a":
                            fmt.Println(v1, v2, v5)
                        case "\n\t":
                            v2 = fi16
                        case "hello":
                            v17 := GenS4{F0: 151, F1: 0}
                            _ = v17
                        default:
                            fmt.Println(v3, v2, fi14)
                        }
                        _ = fi16
                    }
                    v18 := GenS4{F0: 65535, F1: 32768}
                    _ = v18
                }
                v19 := []int{}
                v4 = 255
                _ = v19
            } else {
                v20 := []int{}
                var v21 uint32 = 4294967295
                _ = v20
                _ = v21
            }
            _ = fi14
        }
        v22 := []int{(-(v2)), 127, v2, v2}
        _ = v22
    }
    var v23 float32 = 0.0
    var v24 bool = (3.14159 <= (1.5 * 3.14159))
    if (410 <= 1000) {
        var v25 string = "\n\t"
        var v26 uint16 = 69
        v27 := []int{v2, v2, int(v3), v2}
        _ = v25
        _ = v26
        _ = v27
    }
    fmt.Println(v1)
    fmt.Println(v2, v5)
    v28 := []int{-1, 2147483647}
    var v29 int16 = 1
    fmt.Println(v1, v29)
    v4 += (v4 | v4)
    _ = len(v28)
    if true {
        v30 := GenS5{F0: -1, F1: 1, F2: ""}
        _ = v30
    }
    for fi31 := 0; fi31 < 17; fi31++ {
        for fi32 := 0; fi32 < 15; fi32++ {
            fmt.Println(fi31)
            _ = fi32
        }
        _ = fi31
    }
    for fi33 := 0; fi33 < 7; fi33++ {
        switch v2 {
        case 4294967295:
            if ((float32((1.5 - (-2.25 + 1e10))) >= float32((3.14159 / (1e10 / -2.25)))) && v24) {
                switch "hello" {
                case "a":
                    v4 += (^(v4))
                default:
                    var v34 string = ((("a" + ("foo bar" + ("foo bar" + ""))) + "foo bar") + "foo bar")
                    _ = v34
                }
                var v35 float32 = (0.1 - v23)
                for fi36 := 0; fi36 < 15; fi36++ {
                    v2 -= (fi36 - (fi36 & -688))
                    _ = fi36
                }
                _ = v35
            }
        case 32768:
            var v37 string = "\n\t"
            _ = v37
        }
        v23 += v5
        v1 = 18446744073709551615
        _ = fi33
    }
    var v38 float64 = 0.0
    if (2 != 110) {
        v39 := GenS2{F0: uint32(v1), F1: 42, F2: (("hello" + "a") + "hello")}
        _ = v39
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("slicelen v28:", len(v28), v28)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("v38:", v38)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v28
    _ = v29
    _ = v3
    _ = v38
    _ = v4
    _ = v5
}
