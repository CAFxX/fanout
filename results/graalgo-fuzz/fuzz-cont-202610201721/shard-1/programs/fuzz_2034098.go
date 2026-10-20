package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 bool; F1 bool; F2 int }
type GenS3 struct { F0 bool; F1 int; F2 string }
type GenS4 struct { F0 int8; F1 uint8 }
type GenS5 struct { F0 int; F1 int32; F2 uint }
type GenS6 struct { F0 uint64 }
func genFn0(p0 string, p1 string, p2 uint) int32 {
    var v5 float32 = 1.5
    p2 = p2
    v5 = ((float32(0.0) / (v5 - ((((float32(1e10) - (((v5 / 0.1) * 0.1) + (0.0 - float32((0.1 * ((((1.5 - -2.25) + 0.0) / 0.1) / float64(float32(-2.25)))))))) - (v5 - 1e10)) * v5) + v5))) * 0.0)
    for fi6 := 0; fi6 < 15; fi6++ {
        v7 := []int{279, 9223372036854775807, (fi6 >> (((^(p2))) & 63))}
        _ = fi6
        _ = v7
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 42
}
func genFn1() int8 {
    var v8 int8 = -47
    var v9 float32 = float32(0.0)
    fmt.Println(v9, v8)
    _ = v8
    _ = v9
    return (v8 % ((-128) | 1))
}

func main() {
    v1 := make(map[string]int)
    var v2 uint16 = 65535
    v3 := []uint32{1}
    var v4 int8 = 127
    switch "foo bar" {
    case "hello":
        _ = len(v1)
    case "a":
        var v10 uint = 42
        _ = v10
    case "\n\t":
        v11 := GenS1{F0: v4}
        _ = v11
    }
    for fi12 := 0; fi12 < 19; fi12++ {
        v2 -= (v2 % (((v2 << ((4294967295) & 63))) | 1))
        _ = fi12
    }
    switch "foo bar" {
    case "\n\t":
        var v13 uint8 = 2
        _ = v13
    }
    v2 -= v2
    var v14 float64 = float64(0.1)
    var v15 int64 = 635
    var v16 bool = false
    for fi17 := 0; fi17 < 4; fi17++ {
        if ((v2 ^ (v2 - (v2 & (v2 | v2)))) > 650) {
            v18 := GenS2{F0: v16, F1: true, F2: fi17}
            _ = v18
        } else {
            if (v16 || ((1 >= -66) || (!v16))) {
                var v19 int = (fi17 | (fi17 & 128))
                for fi20 := 0; fi20 < 8; fi20++ {
                    var v21 string = "hello"
                    v16 = true
                    var v22 float32 = 1e10
                    _ = fi20
                    _ = v21
                    _ = v22
                }
                _ = v19
            } else {
                switch fi17 {
                case 545:
                    var v23 float32 = ((((1e10 / 3.14159) + float32((0.1 + ((float64(float32(float64(0.0))) * (v14 + (((((((float64(1.5) + ((float64(((1e10 + (float32(v14) - 0.0)) / 1.0)) / float64(0.1)) * 3.14159)) * v14) * v14) - v14) * v14) * float64(1.5)) - v14))) / (float64(0.0) * v14))))) / 1e10) / 3.14159)
                    _ = v23
                case -9223372036854775808:
                    v24 := GenS3{F0: true, F1: (fi17 >> ((828) & 63)), F2: (((("a" + "") + "") + "") + "\n\t")}
                    _ = v24
                case -463:
                    v25 := GenS2{F0: ((!(v15 != (-(v15)))) || (v2 != 32767)), F1: (uint8(v2) == 84), F2: fi17}
                    _ = v25
                default:
                    if v16 {
                        v16 = true
                        for fi26 := 0; fi26 < 3; fi26++ {
                            v27 := GenS4{F0: (v4 ^ (v4 * int8(fi26))), F1: 127}
                            var v28 float32 = (1.5 + (float32((((1.5 + (v14 - 1e10)) / 3.14159) - (((v14 * v14) - v14) / (0.1 - float64(3.14159))))) * (1e10 + float32(v14))))
                            v15 += v15
                            _ = fi26
                            _ = v27
                            _ = v28
                        }
                        v15 -= int64(v4)
                    } else {
                        v29 := GenS2{F0: (((-2.25 - 1e10) + (3.14159 / 0.1)) < 1.5), F1: (!(!false)), F2: -9223372036854775808}
                        for fi30 := 0; fi30 < 14; fi30++ {
                            var v31 bool = v16
                            _ = fi30
                            _ = v31
                        }
                        _ = v29
                    }
                }
                var v32 string = ""
                v33 := GenS5{F0: (fi17 | 247), F1: -1, F2: 757}
                _ = v32
                _ = v33
            }
            v34 := GenS2{F0: (false && v16), F1: (v16 || v16), F2: (fi17 - 1)}
            _ = v34
        }
        v15 = (^(v15))
        v15 = v15
        _ = fi17
    }
    if len(v3) > 0 {
        v3[0] = 1
    }
    _ = len(v1)
    if false {
        switch ("\n\t" + "a") {
        case "\n\t":
            if true {
                v16 = v16
                if (18446744073709551615 > 2) {
                    fmt.Println(v15)
                } else {
                    v4 = (-(v4))
                }
                for fi35 := 0; fi35 < 9; fi35++ {
                    var v36 int64 = -9223372036854775808
                    if v16 {
                        if ((-(v4)) != 100) {
                            var v37 float32 = (((0.1 / (1e10 + 3.14159)) * (1.5 - 1.5)) / (3.14159 + 1.5))
                            delete(v1, "foo bar")
                            _ = v37
                        }
                        v38 := GenS4{F0: (v4 & (^(v4))), F1: 127}
                        _ = v38
                    }
                    v39 := GenS1{F0: v4}
                    _ = fi35
                    _ = v36
                    _ = v39
                }
            }
        }
        v2 += (^(v2))
        switch ("foo bar" + "a") {
        case "foo bar":
            v40 := GenS6{F0: 4294967295}
            _ = v40
        case "":
            if v16 {
                for fi41 := 0; fi41 < 14; fi41++ {
                    delete(v1, "foo bar")
                    fmt.Println(v15)
                    _ = fi41
                }
                v1[(("\n\t" + ("hello" + ("\n\t" + "hello"))) + "")] = 0
                fmt.Println(v2, v2)
            }
        default:
            var v42 int8 = -1
            _ = v42
        }
    } else {
        _ = len(v3)
        var v43 uint16 = v2
        _ = v43
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
