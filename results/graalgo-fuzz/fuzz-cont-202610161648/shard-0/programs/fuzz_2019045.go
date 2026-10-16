package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint32; F1 uint64 }
type GenS5 struct { F0 int64; F1 uint; F2 int8 }
type GenS6 struct { F0 uint32; F1 int8 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int8; F1 uint8 }
type GenS9 struct { F0 int }
type GenS10 struct { F0 uint64; F1 int16; F2 string }
type GenS11 struct { F0 uint16; F1 int8; F2 int64 }
type GenS12 struct { F0 int64; F1 uint }
type GenS13 struct { F0 string; F1 bool; F2 uint64 }
func genFn0(p0 int8) int16 {
    var v5 int64 = 1
    for fi6 := 0; fi6 < 11; fi6++ {
        switch ("" + "foo bar") {
        case "foo bar":
            var v7 int16 = 42
            _ = v7
        }
        var v8 uint16 = 0
        _ = fi6
        _ = v8
    }
    _ = p0
    _ = v5
    return -1
}
func genFn1(p0 int) int {
    p0 = p0
    var v9 int = 982
    for fi10 := 0; fi10 < 3; fi10++ {
        v11 := []int{p0, (v9 << ((1) & 63)), (p0 >> ((0) & 63)), (p0 / ((71) | 1))}
        var v12 int = v9
        switch v12 {
        case 127:
            var v13 int8 = 127
            _ = v13
        default:
            var v14 float32 = ((((3.14159 - (float32(1.5) * (float32(float64((0.0 * 1e10))) - 3.14159))) * ((float32((0.0 - 0.0)) * (float32(1.5) * -2.25)) - 3.14159)) - 1e10) * (0.0 * (-2.25 - float32((0.1 - float64(0.1))))))
            _ = v14
        }
        _ = fi10
        _ = v11
        _ = v12
    }
    _ = p0
    _ = v9
    return (v9 - v9)
}

func main() {
    v1 := make(map[uint32]int)
    var v2 int = 1
    var v3 bool = ((!((-9223372036854775808 > 9223372036854775807) || true)) && (92 == 0))
    var v4 float64 = -2.25
    v15 := GenS1{F0: 2147483648}
    var v16 int32 = 0
    switch "hello" {
    case "\n\t":
        v17 := GenS1{F0: 2}
        _ = v17
    default:
        switch ("a" + "a") {
        case "hello":
            var v18 uint = 42
            _ = v18
        case "":
            for fi19 := 0; fi19 < 17; fi19++ {
                v16 -= 0
                v20 := GenS2{F0: v3}
                v21 := []int{fi19, fi19, (fi19 + -977)}
                _ = fi19
                _ = v20
                _ = v21
            }
        default:
            v22 := v1[331]
            _ = v22
        }
    }
    if (1 <= 255) {
        v23 := GenS3{F0: ("hello" + "foo bar")}
        for fi24 := 0; fi24 < 17; fi24++ {
            if v3 {
                switch genFn1((fi24 - (fi24 - (fi24 & (v2 >> ((172) & 63)))))) {
                case 9223372036854775807:
                    fmt.Println(v16)
                case 128:
                    for fi25 := 0; fi25 < 10; fi25++ {
                        if v3 {
                            switch (-(fi25)) {
                            case -9223372036854775808:
                                var v26 uint16 = 382
                                v2 -= fi24
                                v27 := GenS1{F0: 1}
                                _ = v26
                                _ = v27
                            }
                        } else {
                            var v28 uint64 = 696
                            v29 := GenS3{F0: "\n\t"}
                            var v30 int = (fi24 % ((fi25) | 1))
                            _ = v28
                            _ = v29
                            _ = v30
                        }
                        _ = fi25
                    }
                default:
                    var v31 int32 = (-(v16))
                    _ = v31
                }
            } else {
                v16 = v16
                fmt.Println(v3, v16)
                delete(v1, 0)
            }
            _ = fi24
        }
        for fi32 := 0; fi32 < 4; fi32++ {
            if true {
                v3 = v3
                _ = len(v1)
                for fi33 := 0; fi33 < 1; fi33++ {
                    v2 = fi32
                    var v34 int8 = -18
                    var v35 int16 = 0
                    _ = fi33
                    _ = v34
                    _ = v35
                }
            } else {
                var v36 uint = 738
                _ = v36
            }
            _ = fi32
        }
        _ = v23
    }
    if v3 {
        switch 127 {
        case 42:
            var v37 int32 = v16
            _ = v37
        case -9223372036854775808:
            var v38 int32 = v16
            _ = v38
        case -906:
            if (!(!(9223372036854775807 < -1))) {
                v39 := GenS3{F0: ((("hello" + "hello") + (("hello" + "") + ((("" + "foo bar") + ("\n\t" + (((("\n\t" + "foo bar") + "a") + ((("hello" + "") + "hello") + ("foo bar" + "foo bar"))) + "foo bar"))) + "\n\t"))) + "hello")}
                v3 = (!(634 < 678))
                for fi40 := 0; fi40 < 7; fi40++ {
                    var v41 uint8 = 127
                    var v42 bool = false
                    _ = fi40
                    _ = v41
                    _ = v42
                }
                _ = v39
            } else {
                fmt.Println(v3)
                var v43 int8 = -121
                switch "a" {
                case "":
                    switch (v2 / (((-(v2))) | 1)) {
                    case 128:
                        var v44 uint64 = 24
                        _ = v44
                    case 32767:
                        v45 := GenS4{F0: 2147483648, F1: 2}
                        _ = v45
                    case 32768:
                        var v46 int8 = (v43 | int8(v16))
                        _ = v46
                    }
                case "foo bar":
                    if v3 {
                        if (1000 != -133) {
                            var v47 uint8 = 255
                            v43 += (v43 % (((v43 % (((^(v43))) | 1))) | 1))
                            v4 -= (v4 + 0.1)
                            _ = v47
                        } else {
                            for fi48 := 0; fi48 < 19; fi48++ {
                                var v49 int = (v2 + (fi48 - (v2 - -9223372036854775808)))
                                v50 := []int{genFn1(fi48), (v2 + genFn1((v2 % ((v49) | 1)))), genFn1(v49)}
                                fmt.Println(v3)
                                _ = fi48
                                _ = v49
                                _ = v50
                            }
                            var v51 float64 = v4
                            var v52 int = v2
                            _ = v51
                            _ = v52
                        }
                        for fi53 := 0; fi53 < 19; fi53++ {
                            switch -1 {
                            case 1:
                                v43 += (v43 + v43)
                            }
                            _ = fi53
                        }
                        v54 := GenS1{F0: 0}
                        _ = v54
                    } else {
                        switch "a" {
                        case "a":
                            var v55 float32 = 1.5
                            _ = v55
                        }
                        if false {
                            delete(v1, 1)
                            var v56 bool = v3
                            _ = v56
                        } else {
                            v57 := []int{}
                            _ = v57
                        }
                    }
                default:
                    switch (("" + "foo bar") + ("hello" + ("foo bar" + (("\n\t" + ("\n\t" + "")) + (("" + "a") + ("a" + "a")))))) {
                    case "hello":
                        var v58 int8 = v43
                        _ = v58
                    default:
                        switch "a" {
                        case "":
                            _ = len(v1)
                        case "\n\t":
                            var v59 int32 = v16
                            _ = v59
                        case "foo bar":
                            v3 = false
                        }
                    }
                }
                _ = v43
            }
        default:
            var v60 uint32 = 2
            _ = v60
        }
        switch (-(v2)) {
        case 32767:
            for fi61 := 0; fi61 < 2; fi61++ {
                fmt.Println(fi61, v2, v3)
                _ = fi61
            }
        }
    } else {
        if true {
            v4 = 1.5
        } else {
            v62 := GenS5{F0: 1, F1: 42, F2: -128}
            _ = v62
        }
        var v63 int = genFn1((v2 + (v2 - (v2 + v2))))
        if v3 {
            var v64 uint8 = 128
            v65 := GenS1{F0: 65535}
            _ = v64
            _ = v65
        }
        _ = v63
    }
    fmt.Println(v2)
    if (0 <= 1) {
        for fi66 := 0; fi66 < 2; fi66++ {
            switch "\n\t" {
            case "hello":
                v3 = v3
            default:
                if false {
                    if (0.1 >= (((v4 * float64(0.0)) * v4) - ((v4 / v4) + v4))) {
                        if (v2 != (v2 * (v2 & v2))) {
                            v3 = (!true)
                        }
                    } else {
                        var v67 uint32 = 2147483647
                        _ = v67
                    }
                    for fi68 := 0; fi68 < 13; fi68++ {
                        switch ((("foo bar" + "hello") + "foo bar") + "hello") {
                        case "":
                            var v69 int64 = 0
                            _ = v69
                        case "a":
                            var v70 float64 = 3.14159
                            _ = v70
                        case "\n\t":
                            if (!false) {
                                v71 := GenS6{F0: 2, F1: -103}
                                switch "\n\t" {
                                case "\n\t":
                                    if v3 {
                                        fmt.Println(fi66)
                                        v2 += fi66
                                        v4 += 1e10
                                        v72 := []int{(v2 * (fi68 << ((1) & 63))), fi68, (v2 + (fi68 << ((255) & 63)))}
                                        _ = v72
                                    } else {
                                        v16 -= v16
                                    }
                                case "":
                                    var v73 int64 = -974
                                    _ = v73
                                default:
                                    var v74 int32 = v16
                                    _ = v74
                                }
                                _ = v71
                            }
                            fmt.Println(fi66, fi66)
                        default:
                            if (v3 && v3) {
                                v16 -= (v16 * (-(v16)))
                            }
                        }
                        var v75 int64 = -253
                        if false {
                            v76 := GenS7{F0: 0}
                            if false {
                                var v77 int32 = -765
                                if (v75 != v75) {
                                    v78 := GenS8{F0: 100, F1: 84}
                                    _ = v78
                                } else {
                                    v79 := GenS9{F0: (v2 + genFn1((v2 + v2)))}
                                    if v3 {
                                        v80 := GenS10{F0: 4294967295, F1: 1, F2: (("\n\t" + ("\n\t" + ((("a" + "foo bar") + ("\n\t" + "hello")) + ("foo bar" + "")))) + "hello")}
                                        _ = v80
                                    }
                                    _ = v79
                                }
                                _ = v77
                            }
                            _ = v76
                        }
                        _ = fi68
                        _ = v75
                    }
                }
            }
            for fi81 := 0; fi81 < 4; fi81++ {
                fmt.Println(v3, v4)
                _ = fi81
            }
            _ = fi66
        }
        switch ("foo bar" + ("foo bar" + "hello")) {
        case "foo bar":
            var v82 float64 = (v4 / 1.5)
            _ = v82
        case "hello":
            v83 := GenS7{F0: 1}
            _ = v83
        default:
            for fi84 := 0; fi84 < 16; fi84++ {
                var v85 uint = 737
                _ = fi84
                _ = v85
            }
        }
    }
    v86 := GenS11{F0: 936, F1: -1, F2: 0}
    var v87 int16 = 0
    v3 = v3
    if (1 > 169) {
        for fi88 := 0; fi88 < 8; fi88++ {
            v16 = (v16 ^ v16)
            if v3 {
                for fi89 := 0; fi89 < 10; fi89++ {
                    var v90 int32 = 2147483647
                    if (706 < 4294967295) {
                        if ((1e10 - 1.5) >= (0.0 * 1e10)) {
                            v91 := GenS12{F0: 552, F1: 42}
                            delete(v1, 1)
                            _ = v91
                        }
                        var v92 bool = (!true)
                        _ = v92
                    } else {
                        fmt.Println(v4, v90, fi88)
                        if (42 >= -111) {
                            _ = len(v1)
                        } else {
                            v90 = v16
                            v2 = fi89
                        }
                        v4 += v4
                    }
                    var v93 uint = 255
                    _ = fi89
                    _ = v90
                    _ = v93
                }
                for fi94 := 0; fi94 < 5; fi94++ {
                    v2 += (fi88 >> ((473) & 63))
                    v87 = v87
                    v95 := GenS13{F0: "foo bar", F1: v3, F2: 1}
                    _ = fi94
                    _ = v95
                }
            } else {
                v96 := []int{926, (v2 / ((genFn1((fi88 | -2147483648))) | 1)), fi88, (v2 << ((0) & 63))}
                _ = v96
            }
            _ = fi88
        }
        v4 = float64(0.1)
        fmt.Println(v2)
    }
    var v97 uint64 = 18446744073709551615
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v87:", v87)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v86
    _ = v87
    _ = v97
}
