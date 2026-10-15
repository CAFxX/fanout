package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int }
type GenS5 struct { F0 bool; F1 uint }
type GenS6 struct { F0 uint32; F1 int8; F2 uint32 }
func genFn0(p0 int32, p1 string, p2 int64) uint64 {
    p2 = (p2 >> ((255) & 63))
    if (!(42 > 645)) {
        for fi6 := 0; fi6 < 1; fi6++ {
            p2 -= (p2 * (p2 & -650))
            switch -1 {
            case 314:
                var v7 uint16 = 653
                _ = v7
            }
            v8 := GenS1{F0: p0, F1: 708}
            _ = fi6
            _ = v8
        }
    } else {
        var v9 uint = 596
        p2 -= p2
        _ = v9
    }
    if (181 <= 2) {
        v10 := GenS2{F0: 255}
        _ = v10
    }
    var v11 uint8 = 128
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 42
}
func genFn1(p0 int8, p1 int8, p2 int16) uint16 {
    v12 := GenS2{F0: 42}
    fmt.Println(p2, p1)
    fmt.Println(p0)
    switch -9223372036854775808 {
    case 128:
        fmt.Println(p0, p0, p2)
    case -9223372036854775808:
        var v13 string = ("a" + "")
        _ = v13
    case 604:
        var v14 string = "a"
        _ = v14
    default:
        v15 := []int{32768, -1, 128, 32768}
        _ = v15
    }
    switch "\n\t" {
    case "\n\t":
        v16 := []int{}
        _ = v16
    case "foo bar":
        v17 := []int{127}
        _ = v17
    case "hello":
        var v18 string = "\n\t"
        _ = v18
    default:
        switch "" {
        case "foo bar":
            if false {
                if true {
                    p2 += (p2 % ((1) | 1))
                    p0 = -128
                } else {
                    v19 := GenS2{F0: 0}
                    for fi20 := 0; fi20 < 13; fi20++ {
                        for fi21 := 0; fi21 < 20; fi21++ {
                            var v22 int = (fi21 % ((127) | 1))
                            _ = fi21
                            _ = v22
                        }
                        _ = fi20
                    }
                    _ = v19
                }
                for fi23 := 0; fi23 < 5; fi23++ {
                    v24 := []int{fi23, (fi23 ^ fi23), 447}
                    _ = fi23
                    _ = v24
                }
            }
        default:
            var v25 string = ("" + "foo bar")
            _ = v25
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 255
}

func main() {
    var v1 float32 = (0.1 * float32(0.1))
    var v2 int32 = 811
    v3 := make(map[int]uint32)
    var v4 string = "hello"
    var v5 int16 = 32767
    v26 := GenS3{F0: (18446744073709551615 > 737)}
    for fi27 := 0; fi27 < 4; fi27++ {
        v2 = v2
        v28 := GenS4{F0: fi27}
        v29 := GenS4{F0: (fi27 % (((fi27 >> ((42) & 63))) | 1))}
        _ = fi27
        _ = v28
        _ = v29
    }
    v30 := []int{9223372036854775807, 0, -2147483648, 0}
    if false {
        for fi31 := 0; fi31 < 13; fi31++ {
            for fi32 := 0; fi32 < 16; fi32++ {
                v33 := GenS1{F0: (v2 << ((950) & 63)), F1: 0}
                v34 := GenS5{F0: false, F1: 42}
                if (float32(0.1) <= v1) {
                    if len(v30) > 0 {
                        v30[0] = fi31
                    }
                } else {
                    if false {
                        var v35 string = "foo bar"
                        v2 += (-(v2))
                        _ = v35
                    } else {
                        v5 += (-(v5))
                        if len(v30) > 0 {
                            v30[0] = 424
                        }
                        v2 -= 1
                    }
                    v4 = (v4 + v4)
                    v1 = v1
                }
                v5 = (v5 | 1)
                _ = fi32
                _ = v33
                _ = v34
            }
            _ = fi31
        }
        v1 = (v1 * (3.14159 * 1e10))
        var v36 uint8 = 25
        _ = v36
    }
    for fi37 := 0; fi37 < 3; fi37++ {
        var v38 int8 = -109
        if len(v30) >= 2 {
            v30 = v30[1:len(v30)]
        }
        _ = fi37
        _ = v38
    }
    for fi39 := 0; fi39 < 15; fi39++ {
        fmt.Println(v4, v4, v2)
        for fi40 := 0; fi40 < 12; fi40++ {
            v5 = (-(v5))
            _ = fi40
        }
        v3[fi39] = 65535
        _ = fi39
    }
    v1 -= (((v1 * v1) - ((0.1 * v1) / ((v1 * v1) + (((((v1 * (float32(-2.25) * float32(0.1))) * 0.0) - v1) + ((-2.25 + v1) + (v1 * 0.1))) + float32(1.5))))) - ((0.1 + (3.14159 / (v1 - (((v1 - v1) - (((0.0 - v1) - 1.5) + v1)) - v1)))) / v1))
    if true {
        fmt.Println(v4, v2)
        var v41 float64 = ((((0.1 * (0.1 + 1e10)) * (0.0 + -2.25)) / 0.1) - float64(v1))
        _ = v41
    }
    var v42 float64 = (0.0 - 0.0)
    switch ((("\n\t" + "a") + v4) + v4) {
    case "\n\t":
        fmt.Println(v42, v4)
    case "":
        v43 := v3[363]
        _ = v43
    case "a":
        switch ("\n\t" + ((("hello" + "") + ("\n\t" + v4)) + "hello")) {
        case "\n\t":
            for fi44 := 0; fi44 < 11; fi44++ {
                if (false && true) {
                    var v45 uint64 = 1
                    for fi46 := 0; fi46 < 2; fi46++ {
                        fmt.Println(v42)
                        var v47 float32 = v1
                        _ = fi46
                        _ = v47
                    }
                    _ = v45
                }
                var v48 uint8 = 2
                if true {
                    v49 := GenS1{F0: v2, F1: 32767}
                    _ = v49
                }
                _ = fi44
                _ = v48
            }
        default:
            if true {
                fmt.Println(v42, v5, v1)
            } else {
                if true {
                    var v50 int = -2147483648
                    for fi51 := 0; fi51 < 4; fi51++ {
                        if (-2.25 <= (v42 / float64(float32(float64(1.5))))) {
                            v52 := GenS6{F0: 65535, F1: -24, F2: 1}
                            _ = v52
                        } else {
                            _ = len(v30)
                            var v53 uint8 = 127
                            _ = v53
                        }
                        _ = fi51
                    }
                    _ = v50
                }
            }
        }
    }
    var v54 int8 = -128
    switch ("" + v4) {
    case "foo bar":
        var v55 uint32 = 303
        _ = v55
    default:
        _ = len(v3)
    }
    var v56 uint = 141
    var v57 string = "a"
    var v58 bool = true
    var v59 uint32 = 2147483647
    switch v4 {
    case "foo bar":
        for fi60 := 0; fi60 < 16; fi60++ {
            var v61 uint64 = 2
            _ = fi60
            _ = v61
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v30:", len(v30), v30)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v30
    _ = v4
    _ = v42
    _ = v5
    _ = v54
    _ = v56
    _ = v57
    _ = v58
    _ = v59
}
