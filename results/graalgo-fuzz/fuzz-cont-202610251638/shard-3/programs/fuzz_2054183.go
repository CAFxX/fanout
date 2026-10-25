package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 uint32; F1 int; F2 string }
type GenS3 struct { F0 uint16; F1 bool }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 int16; F1 int16; F2 int64 }
func genFn0(p0 int8) uint {
    p0 = p0
    switch "hello" {
    case "hello":
        for fi5 := 0; fi5 < 11; fi5++ {
            for fi6 := 0; fi6 < 6; fi6++ {
                var v7 int8 = (-(p0))
                var v8 bool = (1 > 18446744073709551615)
                fmt.Println(v7)
                switch 1 {
                case 2147483647:
                    var v9 uint16 = 42
                    _ = v9
                case -14:
                    v7 = -128
                    var v10 uint = 894
                    v11 := []int{618, (-(fi5)), fi5}
                    var v12 float64 = float64((float32(float64(float32((-2.25 * 1.5)))) / 1.5))
                    _ = v10
                    _ = v11
                    _ = v12
                case -2147483648:
                    var v13 uint16 = 32767
                    _ = v13
                }
                fmt.Println(v7, v8)
                var v14 uint = 42
                var v15 float64 = (0.0 / -2.25)
                if false {
                    v16 := []int{fi5, (fi5 * (fi6 * fi6)), 127}
                    v17 := GenS1{F0: v8, F1: 2}
                    _ = v16
                    _ = v17
                }
                _ = fi6
                _ = v14
                _ = v15
                _ = v7
                _ = v8
            }
            v18 := GenS1{F0: ((("foo bar" + "foo bar") + (("a" + "\n\t") + "\n\t")) > "hello"), F1: 2}
            var v19 int16 = -1
            _ = fi5
            _ = v18
            _ = v19
        }
    default:
        var v20 uint32 = 465
        _ = v20
    }
    for fi21 := 0; fi21 < 8; fi21++ {
        var v22 uint = 13
        switch fi21 {
        case -481:
            for fi23 := 0; fi23 < 16; fi23++ {
                if (859 >= 449) {
                    var v24 uint64 = 864
                    var v25 int8 = (p0 & (p0 ^ 42))
                    _ = v24
                    _ = v25
                } else {
                    var v26 uint16 = 65535
                    v27 := []int{128, -1, -9223372036854775808, (fi23 & (fi21 % (((-(fi23))) | 1)))}
                    if (false && ((p0 >> ((v22) & 63)) == 42)) {
                        var v28 uint32 = 1
                        v22 = (v22 * (v22 & 201))
                        var v29 int32 = -212
                        _ = v28
                        _ = v29
                    } else {
                        var v30 uint8 = 109
                        var v31 uint16 = v26
                        _ = v30
                        _ = v31
                    }
                    _ = v26
                    _ = v27
                }
                if (0 >= 42) {
                    var v32 float32 = 1.5
                    v22 = (v22 >> (((v22 * 65535)) & 63))
                    _ = v32
                } else {
                    v22 = (v22 ^ (v22 - (v22 | 18446744073709551615)))
                    v33 := GenS2{F0: 0, F1: (fi23 - fi23), F2: "\n\t"}
                    _ = v33
                }
                _ = fi23
            }
        case -674:
            v34 := GenS3{F0: 105, F1: false}
            _ = v34
        }
        _ = fi21
        _ = v22
    }
    v35 := GenS3{F0: 0, F1: false}
    v36 := []int{}
    _ = p0
    _ = v35
    _ = v36
    return 481
}
func genFn1(p0 string, p1 int32, p2 uint16) uint32 {
    v37 := GenS4{F0: p1}
    p0 = ("" + ("hello" + p0))
    var v38 uint16 = 949
    p1 += p1
    _ = p0
    _ = p1
    _ = p2
    _ = v37
    _ = v38
    return 322
}

func main() {
    var v1 string = "\n\t"
    var v2 int8 = 0
    v3 := make(map[string]int)
    var v4 bool = true
    var v39 uint16 = 2
    if true {
        v3[("a" + v1)] = 2147483647
    }
    var v40 int16 = -526
    for fi41 := 0; fi41 < 9; fi41++ {
        for fi42 := 0; fi42 < 16; fi42++ {
            if v4 {
                v43 := GenS5{F0: 42, F1: v40, F2: 589}
                _ = v43
            }
            v4 = (127 <= 255)
            _ = fi42
        }
        var v44 int16 = -75
        _ = fi41
        _ = v44
    }
    v40 = 1000
    v40 -= v40
    v40 = 32767
    var v45 bool = (!(!true))
    fmt.Println(v4, v45, v39)
    switch 0 {
    case 9223372036854775807:
        for fi46 := 0; fi46 < 18; fi46++ {
            if (!true) {
                var v47 float32 = (1.5 / ((float32(float64((float32(-2.25) - 0.1))) + -2.25) + 3.14159))
                _ = v47
            } else {
                v48 := []int{-9223372036854775808, -2147483648, (fi46 * 582)}
                var v49 uint32 = 42
                _ = v48
                _ = v49
            }
            var v50 int = (fi46 + (fi46 + fi46))
            fmt.Println(v45, v2)
            _ = fi46
            _ = v50
        }
    default:
        for fi51 := 0; fi51 < 18; fi51++ {
            var v52 uint32 = 325
            var v53 uint8 = 201
            _ = fi51
            _ = v52
            _ = v53
        }
    }
    v45 = v4
    fmt.Println(v2, v4)
    v54 := v3[(((("\n\t" + "hello") + v1) + "a") + "a")]
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("v40:", v40)
    fmt.Println("v45:", v45)
    fmt.Println("v54:", v54)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v45
    _ = v54
}
