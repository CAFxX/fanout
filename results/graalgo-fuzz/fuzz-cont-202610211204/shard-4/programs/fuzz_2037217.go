package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16; F1 int8 }
type GenS4 struct { F0 uint; F1 bool; F2 int }
type GenS5 struct { F0 string }
type GenS6 struct { F0 int8; F1 string; F2 uint32 }
type GenS7 struct { F0 int64; F1 int8 }
type GenS8 struct { F0 int8; F1 uint8; F2 int8 }
type GenS9 struct { F0 int32 }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 int; F1 int32; F2 bool }
type GenS12 struct { F0 int16 }
type GenS13 struct { F0 int8; F1 int32 }
func genFn0(p0 int32, p1 int8, p2 string) bool {
    switch -984 {
    case -1:
        fmt.Println(p2, p0, p1)
    }
    v6 := GenS1{F0: 1000}
    p1 = -82
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return (489 == 4294967295)
}
func genFn1() uint8 {
    var v7 uint8 = 92
    for fi8 := 0; fi8 < 8; fi8++ {
        fmt.Println(v7, v7)
        if true {
            v9 := []int{(fi8 + (fi8 | fi8)), fi8, (-(fi8)), 898}
            fmt.Println(fi8, v7)
            v9 = append(v9, -375)
            _ = v9
        }
        _ = fi8
    }
    _ = v7
    return v7
}

func main() {
    var v1 int64 = -44
    var v2 int8 = -40
    var v3 uint = 620
    var v4 int8 = v2
    v5 := make(map[int]bool)
    v2 -= (-(v4))
    var v10 int8 = -128
    v10 = (v2 ^ (v4 | v2))
    switch 32768 {
    case 4294967295:
        var v11 bool = (!true)
        _ = v11
    case 128:
        switch 32768 {
        case 975:
            if true {
                fmt.Println(v3, v2, v10)
            }
        case 2147483647:
            v10 = v2
        case 0:
            if (0.1 <= float64(1.5)) {
                var v12 uint8 = 2
                var v13 bool = true
                _ = v12
                _ = v13
            } else {
                if (!true) {
                    var v14 float64 = 0.1
                    v1 += -723
                    var v15 uint = v3
                    _ = v14
                    _ = v15
                }
                for fi16 := 0; fi16 < 8; fi16++ {
                    if (true || ((((3.14159 - -2.25) / float64(1.5)) * (0.0 + ((0.1 + float64(1.5)) + 1e10))) >= (float64(float32((3.14159 - 3.14159))) + float64(float32(((0.1 / 3.14159) - 0.1)))))) {
                        v17 := GenS1{F0: 1000}
                        v10 = 0
                        _ = v17
                    } else {
                        v18 := GenS1{F0: 0}
                        _ = v18
                    }
                    v19 := GenS1{F0: -32768}
                    v3 -= (v3 * (v3 << ((v3) & 63)))
                    _ = fi16
                    _ = v19
                }
            }
        default:
            v20 := GenS2{F0: ("a" + "")}
            _ = v20
        }
    case 866:
        v3 -= v3
    }
    var v21 int = -2147483648
    v22 := []int{789, (v21 ^ (v21 ^ 127))}
    v2 += (-(v10))
    for fi23 := 0; fi23 < 16; fi23++ {
        if (!(262 < -2147483648)) {
            v2 += v10
            v22 = append(v22, (fi23 & 0))
        } else {
            v21 = int(v2)
        }
        _ = fi23
    }
    switch "a" {
    case "":
        var v24 int8 = v2
        _ = v24
    }
    var v25 uint8 = 225
    for fi26 := 0; fi26 < 15; fi26++ {
        v5[fi26] = false
        _ = fi26
    }
    var v27 uint8 = v25
    if (("a" >= "a") && true) {
        if (int32(v3) > -752) {
            for fi28 := 0; fi28 < 11; fi28++ {
                for fi29 := 0; fi29 < 13; fi29++ {
                    v10 += 1
                    var v30 bool = true
                    v5[fi28] = v30
                    var v31 uint16 = 0
                    v4 = v2
                    if v30 {
                        var v32 float64 = 0.0
                        v33 := GenS2{F0: (("foo bar" + (((((("" + "a") + "\n\t") + "hello") + ("hello" + ("foo bar" + ("" + (("\n\t" + "\n\t") + ""))))) + "\n\t") + "")) + "\n\t")}
                        _ = v32
                        _ = v33
                    } else {
                        v34 := GenS1{F0: int16(v1)}
                        if false {
                            switch "hello" {
                            case "\n\t":
                                fmt.Println(v31, v10, v2)
                            case "hello":
                                var v35 float64 = float64(1e10)
                                v4 = v10
                                v36 := GenS3{F0: (v31 - v31), F1: -1}
                                _ = v35
                                _ = v36
                            default:
                                switch "foo bar" {
                                case "\n\t":
                                    switch (fi28 % (((fi29 * int(v31))) | 1)) {
                                    case 117:
                                        if (42 != 42) {
                                            var v37 float32 = 1.5
                                            var v38 int16 = 42
                                            _ = v37
                                            _ = v38
                                        }
                                    case 42:
                                        var v39 uint32 = uint32(v27)
                                        _ = v39
                                    default:
                                        v27 = (v27 >> ((v3) & 63))
                                    }
                                    v40 := GenS1{F0: -1}
                                    v25 = (v25 * genFn1())
                                    _ = v40
                                case "":
                                    v3 += 1
                                case "hello":
                                    v41 := GenS2{F0: ""}
                                    _ = v41
                                default:
                                    v42 := GenS2{F0: "a"}
                                    _ = v42
                                }
                            }
                            delete(v5, (fi28 / ((fi29) | 1)))
                        }
                        v3 = v3
                        switch "hello" {
                        case "":
                            var v43 float32 = ((3.14159 * -2.25) * 3.14159)
                            _ = v43
                        case "\n\t":
                            var v44 uint16 = 65535
                            _ = v44
                        case "hello":
                            if true {
                                v4 = v2
                                v45 := GenS1{F0: int16(v1)}
                                v3 -= uint(v27)
                                var v46 uint8 = 231
                                var v47 int8 = (v4 * v10)
                                switch "" {
                                case "hello":
                                    fmt.Println(v21)
                                default:
                                    switch ("foo bar" + ("\n\t" + "a")) {
                                    case "":
                                        switch ("a" + "") {
                                        case "a":
                                            if ((0.1 + float32(0.1)) < (float32(-2.25) * -2.25)) {
                                                v2 -= (v2 - (v2 | (v4 & 1)))
                                                v2 -= (v10 << ((v3) & 63))
                                                if (v31 >= 451) {
                                                    v27 -= (v27 / (((v25 * (v46 ^ v46))) | 1))
                                                }
                                                v21 = int(v27)
                                            } else {
                                                v48 := GenS4{F0: v3, F1: v30, F2: v21}
                                                v49 := GenS5{F0: "a"}
                                                _ = v48
                                                _ = v49
                                            }
                                        default:
                                            v2 = (v4 % (((-(v2))) | 1))
                                        }
                                    }
                                }
                                _ = v45
                                _ = v46
                                _ = v47
                            }
                        default:
                            var v50 uint64 = 18446744073709551615
                            _ = v50
                        }
                        if (1e10 <= (float64((float32(0.1) / (-2.25 / 0.1))) / -2.25)) {
                            fmt.Println(v10, v25, v3)
                            if (int(v3) == fi28) {
                                v51 := v5[559]
                                _ = v51
                            } else {
                                v52 := GenS5{F0: "hello"}
                                if (v30 && (76 < (v1 * v1))) {
                                    var v53 uint16 = (v31 + 255)
                                    switch "foo bar" {
                                    case "hello":
                                        v53 -= (v53 + v31)
                                    case "foo bar":
                                        if false {
                                            _ = len(v22)
                                            if (42 < 42) {
                                                fmt.Println(v21, v21, fi29)
                                            } else {
                                                var v54 bool = false
                                                if false {
                                                    fmt.Println(v54, fi28, v21)
                                                    v53 -= (v31 % (((^(v31))) | 1))
                                                }
                                                if (v10 <= (v10 - (^(v2)))) {
                                                    var v55 int32 = 792
                                                    _ = v55
                                                } else {
                                                    fmt.Println(v54, v1)
                                                    _ = len(v22)
                                                }
                                                if true {
                                                    v56 := GenS5{F0: ("\n\t" + "hello")}
                                                    v57 := GenS6{F0: 42, F1: ((("" + "hello") + ("a" + "\n\t")) + ""), F2: 65535}
                                                    if ((^(v1)) >= 1) {
                                                        switch "hello" {
                                                        case "foo bar":
                                                            v54 = (!v54)
                                                        }
                                                    } else {
                                                        if v54 {
                                                            var v58 uint8 = 127
                                                            _ = v58
                                                        }
                                                    }
                                                    if (!(-775 >= int64(fi29))) {
                                                        switch v21 {
                                                        case 42:
                                                            if len(v22) > 0 {
                                                                v22[0] = fi29
                                                            }
                                                        default:
                                                            fmt.Println(v21, v30)
                                                        }
                                                        v59 := GenS1{F0: -1}
                                                        v27 = 243
                                                        _ = v59
                                                    }
                                                    v22 = append(v22, 160)
                                                    _ = v56
                                                    _ = v57
                                                }
                                                v1 = int64(v10)
                                                _ = v54
                                            }
                                        } else {
                                            v60 := GenS7{F0: v1, F1: 47}
                                            _ = v60
                                        }
                                    case "a":
                                        v61 := GenS8{F0: (v2 & int8(v3)), F1: 181, F2: (v2 >> (((^(v3))) & 63))}
                                        _ = v61
                                    }
                                    v21 -= 0
                                    _ = v53
                                } else {
                                    var v62 uint = (^(v3))
                                    var v63 bool = v30
                                    var v64 uint64 = 287
                                    _ = v62
                                    _ = v63
                                    _ = v64
                                }
                                fmt.Println(fi29, v27)
                                _ = len(v22)
                                var v65 uint64 = 982
                                _ = v52
                                _ = v65
                            }
                            v5[fi29] = true
                            var v66 float32 = -2.25
                            _ = v66
                        } else {
                            if len(v22) >= 2 {
                                v22 = v22[1:len(v22)]
                            }
                            var v67 uint64 = 4294967295
                            _ = v67
                        }
                        _ = v34
                    }
                    _ = fi29
                    _ = v30
                    _ = v31
                }
                _ = fi28
            }
            v68 := GenS1{F0: -1}
            if true {
                v10 += v10
                switch "\n\t" {
                case "":
                    if ((0.0 - ((3.14159 + 1.5) - 0.1)) >= -2.25) {
                        v69 := GenS2{F0: "hello"}
                        _ = v69
                    } else {
                        v22 = append(v22, v21)
                        v5[int(v1)] = (-32768 == 989)
                    }
                case "\n\t":
                    switch ("foo bar" + ("foo bar" + "hello")) {
                    case "\n\t":
                        if true {
                            v22 = append(v22, (v21 ^ v21))
                            if true {
                                switch ("a" + (("foo bar" + "a") + ("\n\t" + "a"))) {
                                case "\n\t":
                                    v70 := GenS9{F0: -465}
                                    _ = v70
                                default:
                                    var v71 float64 = 1.5
                                    _ = v71
                                }
                            } else {
                                for fi72 := 0; fi72 < 10; fi72++ {
                                    var v73 uint = v3
                                    var v74 int64 = (-(v1))
                                    for fi75 := 0; fi75 < 4; fi75++ {
                                        if true {
                                            switch "\n\t" {
                                            case "foo bar":
                                                var v76 uint16 = 58
                                                _ = v76
                                            case "hello":
                                                v2 = (-(v4))
                                                v77 := GenS10{F0: 255}
                                                if true {
                                                    v78 := GenS11{F0: (fi72 ^ 9223372036854775807), F1: 0, F2: true}
                                                    if len(v22) > 0 {
                                                        v22[0] = (fi75 * (^(v21)))
                                                    }
                                                    var v79 string = "hello"
                                                    _ = v78
                                                    _ = v79
                                                } else {
                                                    v5[(fi75 & v21)] = true
                                                    var v80 float32 = 3.14159
                                                    var v81 int = -2147483648
                                                    _ = v80
                                                    _ = v81
                                                }
                                                if (127 >= -128) {
                                                    v74 -= (v74 >> ((v3) & 63))
                                                }
                                                v3 -= (v73 >> ((969) & 63))
                                                _ = v77
                                            default:
                                                v1 = 176
                                            }
                                            var v82 bool = (-628 == int16(v74))
                                            _ = v82
                                        }
                                        v27 = uint8(v3)
                                        v83 := v5[(fi75 >> ((v3) & 63))]
                                        var v84 float32 = float32(((0.1 + 3.14159) - 0.0))
                                        _ = fi75
                                        _ = v83
                                        _ = v84
                                    }
                                    _ = fi72
                                    _ = v73
                                    _ = v74
                                }
                                for fi85 := 0; fi85 < 7; fi85++ {
                                    v86 := GenS12{F0: 42}
                                    for fi87 := 0; fi87 < 8; fi87++ {
                                        if true {
                                            var v88 uint8 = (v27 | v27)
                                            var v89 uint16 = 932
                                            v90 := GenS13{F0: 1, F1: -95}
                                            _ = v88
                                            _ = v89
                                            _ = v90
                                        } else {
                                            v27 = genFn1()
                                            var v91 float32 = 0.0
                                            if ((false && true) && false) {
                                                v91 = v91
                                                var v92 int8 = (v2 * v4)
                                                _ = v92
                                            }
                                            _ = v91
                                        }
                                        var v93 int32 = -583
                                        _ = fi87
                                        _ = v93
                                    }
                                    var v94 int8 = (v2 & v2)
                                    _ = fi85
                                    _ = v86
                                    _ = v94
                                }
                            }
                            for fi95 := 0; fi95 < 8; fi95++ {
                                v21 += v21
                                _ = fi95
                            }
                        }
                    case "":
                        var v96 int = v21
                        _ = v96
                    case "a":
                        _ = len(v22)
                    default:
                        var v97 int8 = v4
                        _ = v97
                    }
                case "a":
                    if (false && (true || false)) {
                        v98 := GenS13{F0: int8(v27), F1: 463}
                        fmt.Println(v1)
                        var v99 string = "a"
                        _ = v98
                        _ = v99
                    } else {
                        var v100 int64 = 42
                        _ = v100
                    }
                }
                if false {
                    v3 = (^(v3))
                    var v101 uint32 = 2
                    v102 := GenS11{F0: (v21 % (((v21 + (v21 & (v21 | 4294967295)))) | 1)), F1: -8, F2: (32767 == 155)}
                    _ = v101
                    _ = v102
                } else {
                    var v103 string = "\n\t"
                    v5[v21] = true
                    _ = v103
                }
            }
            _ = v68
        }
    } else {
        for fi104 := 0; fi104 < 18; fi104++ {
            for fi105 := 0; fi105 < 6; fi105++ {
                v1 = (v1 ^ v1)
                fmt.Println(v4, v21, v21)
                _ = fi105
            }
            var v106 string = ("" + ("foo bar" + ("foo bar" + "\n\t")))
            v1 = v1
            _ = fi104
            _ = v106
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v22:", len(v22), v22)
    fmt.Println("v25:", v25)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v10
    _ = v2
    _ = v21
    _ = v22
    _ = v25
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
