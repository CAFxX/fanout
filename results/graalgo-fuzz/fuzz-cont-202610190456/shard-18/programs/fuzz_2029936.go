package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 string; F1 uint32; F2 uint8 }
type GenS4 struct { F0 uint8; F1 uint32; F2 uint32 }
type GenS5 struct { F0 uint64; F1 int16; F2 int64 }
func genFn0(p0 bool, p1 int64) uint16 {
    if (true || p0) {
        var v5 float64 = ((0.1 + (float64(1.5) - 0.0)) + (-2.25 - 0.0))
        fmt.Println(p0, p1, p0)
        _ = v5
    } else {
        v6 := []int{-344, 128}
        fmt.Println(p1, p0)
        v7 := GenS1{F0: 32768}
        _ = v6
        _ = v7
    }
    switch (("hello" + "foo bar") + "") {
    case "foo bar":
        if (("a" + "hello") > "foo bar") {
            var v8 int = -2147483648
            for fi9 := 0; fi9 < 7; fi9++ {
                v10 := GenS2{F0: ((!(!(!false))) && p0)}
                _ = fi9
                _ = v10
            }
            var v11 float64 = float64((3.14159 * 1e10))
            _ = v11
            _ = v8
        }
    case "\n\t":
        fmt.Println(p0, p1)
    case "a":
        p0 = p0
    default:
        p1 -= p1
    }
    var v12 uint16 = 153
    for fi13 := 0; fi13 < 8; fi13++ {
        v14 := []int{fi13, 9223372036854775807, 32768, 0}
        if (239 != 65535) {
            for fi15 := 0; fi15 < 19; fi15++ {
                _ = len(v14)
                var v16 int8 = -1
                var v17 string = (("a" + (("hello" + ("foo bar" + "")) + "hello")) + "hello")
                _ = fi15
                _ = v16
                _ = v17
            }
            if len(v14) >= 2 {
                v14 = v14[1:len(v14)]
            }
        }
        _ = fi13
        _ = v14
    }
    _ = p0
    _ = p1
    _ = v12
    return (v12 * v12)
}
func genFn1(p0 int16) uint8 {
    var v18 uint16 = 0
    var v19 int64 = 42
    for fi20 := 0; fi20 < 16; fi20++ {
        if (0.0 < 3.14159) {
            var v21 uint16 = (^(v18))
            var v22 uint = 220
            _ = v21
            _ = v22
        }
        fmt.Println(v19)
        v19 = (-(v19))
        _ = fi20
    }
    _ = p0
    _ = v18
    _ = v19
    return uint8(v18)
}

func main() {
    var v1 uint = 37
    var v2 int8 = 1
    var v3 float64 = 0.1
    var v4 string = "a"
    var v23 uint32 = 2147483647
    v1 -= v1
    v24 := GenS3{F0: v4, F1: uint32(v1), F2: 0}
    switch -448 {
    case -2147483648:
        if true {
            var v25 uint32 = v23
            switch 32767 {
            case 128:
                if (false || false) {
                    for fi26 := 0; fi26 < 9; fi26++ {
                        var v27 int64 = 42
                        _ = fi26
                        _ = v27
                    }
                    v28 := GenS2{F0: false}
                    var v29 int32 = -1
                    _ = v28
                    _ = v29
                }
            case 32767:
                var v30 uint = (^(v1))
                _ = v30
            default:
                v25 = (v25 / (((^(v23))) | 1))
            }
            for fi31 := 0; fi31 < 6; fi31++ {
                var v32 int8 = (v2 ^ 0)
                if true {
                    v33 := GenS2{F0: false}
                    v34 := GenS3{F0: ("" + ("a" + v4)), F1: 674, F2: 127}
                    var v35 bool = ((0 >= 2) || true)
                    _ = v33
                    _ = v34
                    _ = v35
                } else {
                    if (true && false) {
                        v3 += (((v3 - v3) / ((-2.25 - float64(1e10)) - v3)) + 1e10)
                    }
                }
                _ = fi31
                _ = v32
            }
            _ = v25
        } else {
            v36 := GenS3{F0: "a", F1: v23, F2: 2}
            _ = v36
        }
    default:
        if true {
            v1 += v1
        } else {
            v37 := GenS3{F0: v4, F1: v23, F2: genFn1(32767)}
            for fi38 := 0; fi38 < 3; fi38++ {
                if (v4 >= "") {
                    v39 := GenS4{F0: uint8(v1), F1: 417, F2: (v23 + 526)}
                    v3 = 1e10
                    if false {
                        var v40 uint32 = (^(v23))
                        var v41 string = ((v4 + v4) + ("" + v4))
                        _ = v40
                        _ = v41
                    }
                    _ = v39
                }
                v42 := GenS2{F0: true}
                _ = fi38
                _ = v42
            }
            _ = v37
        }
    }
    var v43 int64 = 362
    v44 := GenS5{F0: 404, F1: 42, F2: (v43 + (v43 - (v43 << ((1) & 63))))}
    v1 = v1
    var v45 int8 = (v2 << (((^(v1))) & 63))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v45:", v45)
    _ = v1
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v43
    _ = v44
    _ = v45
}
