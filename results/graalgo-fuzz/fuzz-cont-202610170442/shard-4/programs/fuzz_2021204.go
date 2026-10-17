package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint64; F2 int8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int; F1 uint }
type GenS4 struct { F0 uint8; F1 bool }
type GenS5 struct { F0 bool; F1 int16 }
func genFn0() uint8 {
    v6 := []int{}
    if (false && (!(2147483647 == 2147483647))) {
        var v7 int16 = 1
        for fi8 := 0; fi8 < 8; fi8++ {
            for fi9 := 0; fi9 < 4; fi9++ {
                if (32767 < 57) {
                    var v10 float64 = 1.5
                    var v11 uint8 = 255
                    v11 = 212
                    _ = v10
                    _ = v11
                }
                switch "a" {
                case "foo bar":
                    var v12 int32 = -21
                    _ = v12
                default:
                    switch ("\n\t" + (((("" + "foo bar") + "") + "a") + ((("\n\t" + "") + ("a" + ("a" + "\n\t"))) + "foo bar"))) {
                    case "hello":
                        if (!(1 < -967)) {
                            var v13 uint64 = 1
                            var v14 int = fi8
                            _ = v13
                            _ = v14
                        }
                    case "":
                        v15 := GenS1{F0: 255, F1: 389, F2: 0}
                        _ = v15
                    case "a":
                        v7 += int16(fi8)
                        v16 := GenS2{F0: -106}
                        v7 -= (v7 >> ((315) & 63))
                        _ = v16
                    default:
                        if ((v7 & (^(v7))) > (^(v7))) {
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            fmt.Println(fi9, fi8, fi8)
                        }
                    }
                }
                v17 := GenS1{F0: 42, F1: 425, F2: -128}
                v7 += v7
                switch (fi8 | fi8) {
                case 127:
                    var v18 int = fi8
                    _ = v18
                case -9223372036854775808:
                    fmt.Println(fi8, v7, fi9)
                }
                _ = fi9
                _ = v17
            }
            v19 := GenS2{F0: 0}
            v20 := GenS3{F0: (fi8 - 42), F1: 255}
            _ = fi8
            _ = v19
            _ = v20
        }
        for fi21 := 0; fi21 < 16; fi21++ {
            var v22 uint32 = 1
            for fi23 := 0; fi23 < 6; fi23++ {
                var v24 int32 = 633
                _ = fi23
                _ = v24
            }
            if (2 == 48) {
                v22 -= 0
                v6 = append(v6, (fi21 << ((1) & 63)))
                for fi25 := 0; fi25 < 12; fi25++ {
                    if (false || true) {
                        var v26 uint64 = 42
                        _ = v26
                    }
                    if len(v6) >= 2 {
                        v6 = v6[1:len(v6)]
                    }
                    _ = fi25
                }
            } else {
                v22 = 2147483647
                v27 := GenS3{F0: fi21, F1: 18446744073709551615}
                _ = v27
            }
            _ = fi21
            _ = v22
        }
        _ = v7
    } else {
        for fi28 := 0; fi28 < 1; fi28++ {
            var v29 int8 = 1
            v29 -= (v29 ^ v29)
            v30 := GenS1{F0: 1, F1: 887, F2: (-(v29))}
            _ = fi28
            _ = v29
            _ = v30
        }
        v6 = append(v6, 584)
    }
    _ = v6
    return 244
}
func genFn1() int16 {
    var v31 int8 = -128
    for fi32 := 0; fi32 < 14; fi32++ {
        v33 := []int{fi32, fi32, (fi32 % ((fi32) | 1)), (^(fi32))}
        _ = fi32
        _ = v33
    }
    _ = v31
    return 32767
}

func main() {
    var v1 float32 = ((-2.25 - float32(1.5)) + 3.14159)
    var v2 int64 = 42
    var v3 float32 = float32((-2.25 - 3.14159))
    v4 := make(map[string]bool)
    var v5 float32 = v1
    if (0 != 834) {
        v4["hello"] = (false && false)
    }
    v34 := GenS4{F0: 255, F1: false}
    for fi35 := 0; fi35 < 20; fi35++ {
        if ((((1e10 + 0.0) / (v3 * float32(0.0))) - ((((1.5 - v5) + 1.5) * ((1.5 - (0.0 / 1.0)) * v3)) / (v5 * (0.1 * ((v1 / (v1 * ((v3 * float32(0.0)) - v1))) / v5))))) <= 0.0) {
            fmt.Println(fi35, fi35)
            if true {
                switch (^(fi35)) {
                case 32767:
                    switch "foo bar" {
                    case "":
                        v36 := GenS5{F0: (((-2.25 - v1) + v1) == v1), F1: 482}
                        _ = v36
                    case "hello":
                        var v37 float64 = 0.1
                        _ = v37
                    case "foo bar":
                        fmt.Println(v1, v1)
                    }
                }
            } else {
                v38 := v4["a"]
                var v39 int32 = 954
                v38 = (((float32((0.1 - -2.25)) + v1) * ((float32((-2.25 / -2.25)) + float32((0.1 / (1.5 - -2.25)))) * 1.5)) > 0.1)
                _ = v38
                _ = v39
            }
        }
        _ = fi35
    }
    var v40 int16 = -32768
    var v41 float64 = (((0.1 / float64(v3)) - ((0.0 + -2.25) * 1.5)) * float64((-2.25 + 1.5)))
    var v42 uint16 = 65535
    v41 = -2.25
    v4["hello"] = true
    var v43 int8 = 1
    fmt.Println(v5, v1)
    var v44 int8 = v43
    v42 = (v42 ^ v42)
    v5 += -2.25
    v5 = (1.5 * v5)
    var v45 int = 1
    v42 += 2
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v40:", v40)
    fmt.Println("v41:", v41)
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v45:", v45)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v40
    _ = v41
    _ = v42
    _ = v43
    _ = v44
    _ = v45
    _ = v5
}
