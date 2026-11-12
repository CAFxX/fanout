package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint; F2 uint16 }
type GenS2 struct { F0 int; F1 uint16 }
type GenS3 struct { F0 uint64; F1 int }
type GenS4 struct { F0 int64; F1 uint16 }
func genFn0(p0 uint64, p1 int64) int16 {
    var v6 uint8 = 128
    switch "\n\t" {
    case "":
        v7 := []int{128}
        _ = v7
    case "a":
        v8 := GenS1{F0: false, F1: 18446744073709551615, F2: 1}
        _ = v8
    case "hello":
        switch "a" {
        case "a":
            v9 := []int{95, 760, -882, -371}
            _ = v9
        case "foo bar":
            var v10 uint64 = 0
            _ = v10
        default:
            if (!true) {
                if (!(0 != 2)) {
                    p0 -= p0
                    for fi11 := 0; fi11 < 5; fi11++ {
                        var v12 int16 = int16(fi11)
                        switch "foo bar" {
                        case "a":
                            p1 += int64(fi11)
                        case "\n\t":
                            v6 -= (^(v6))
                        }
                        var v13 float64 = (1e10 - -2.25)
                        _ = fi11
                        _ = v12
                        _ = v13
                    }
                }
            } else {
                var v14 uint32 = 65535
                var v15 int32 = 42
                _ = v14
                _ = v15
            }
        }
    }
    for fi16 := 0; fi16 < 11; fi16++ {
        var v17 float32 = float32(0.0)
        _ = fi16
        _ = v17
    }
    var v18 float64 = -2.25
    switch 0 {
    case 32768:
        for fi19 := 0; fi19 < 8; fi19++ {
            for fi20 := 0; fi20 < 3; fi20++ {
                var v21 int8 = -74
                v22 := []int{(fi20 << ((0) & 63)), fi19, fi20}
                var v23 string = ("a" + ((((("hello" + ("hello" + "a")) + ((("foo bar" + ("a" + "")) + "") + ("" + ("a" + "foo bar")))) + (("" + "foo bar") + "hello")) + "foo bar") + "foo bar"))
                _ = fi20
                _ = v21
                _ = v22
                _ = v23
            }
            _ = fi19
        }
    default:
        v18 = 1e10
    }
    _ = p0
    _ = p1
    _ = v18
    _ = v6
    return 443
}
func genFn1(p0 uint64, p1 int8, p2 uint64) int64 {
    v24 := GenS2{F0: -2147483648, F1: 62}
    var v25 uint64 = (p0 | p2)
    for fi26 := 0; fi26 < 4; fi26++ {
        v27 := []int{}
        _ = fi26
        _ = v27
    }
    v28 := GenS3{F0: v25, F1: 128}
    _ = p0
    _ = p1
    _ = p2
    _ = v24
    _ = v25
    _ = v28
    return -1
}

func main() {
    var v1 bool = true
    v2 := []string{}
    v3 := make(map[uint32]string)
    var v4 int16 = -32768
    var v5 int8 = 127
    _ = len(v3)
    var v29 float64 = 3.14159
    switch "a" {
    case "hello":
        delete(v3, 42)
    }
    var v30 uint8 = 255
    v31 := GenS2{F0: 873, F1: 255}
    if (v30 != v30) {
        v32 := GenS1{F0: false, F1: 563, F2: 65535}
        v5 = (v5 & (v5 + (v5 | v5)))
        _ = v32
    }
    fmt.Println(v5)
    if (!v1) {
        v5 -= (v5 * (v5 * -128))
        var v33 int64 = -9223372036854775808
        var v34 int16 = v4
        _ = v33
        _ = v34
    }
    for fi35 := 0; fi35 < 2; fi35++ {
        fmt.Println(v29, v29, v29)
        v5 = v5
        _ = fi35
    }
    for fi36 := 0; fi36 < 15; fi36++ {
        for fi37 := 0; fi37 < 2; fi37++ {
            v29 += (float64((1.5 + float32(v29))) * 0.1)
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = fi37
        }
        delete(v3, 65535)
        for fi38 := 0; fi38 < 7; fi38++ {
            v30 -= 128
            var v39 bool = false
            v40 := GenS4{F0: -1, F1: 364}
            _ = fi38
            _ = v39
            _ = v40
        }
        _ = fi36
    }
    _ = len(v3)
    v41 := GenS4{F0: 42, F1: 496}
    v42 := GenS3{F0: 387, F1: 127}
    v2 = append(v2, "\n\t")
    v5 = (v5 | (v5 + -63))
    var v43 uint64 = 42
    v44 := v3[36]
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v5
}
