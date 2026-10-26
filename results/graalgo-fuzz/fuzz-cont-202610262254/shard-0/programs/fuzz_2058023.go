package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 int8; F1 int; F2 uint32 }
type GenS4 struct { F0 bool; F1 int32 }
type GenS5 struct { F0 string; F1 int32 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint64; F1 int32 }
type GenS8 struct { F0 int16; F1 uint64; F2 int64 }
type GenS9 struct { F0 uint32; F1 uint16; F2 int }
type GenS10 struct { F0 int8; F1 uint8 }
type GenS11 struct { F0 string; F1 int64 }
type GenS12 struct { F0 uint; F1 uint8 }
type GenS13 struct { F0 bool; F1 int64; F2 int64 }
type GenS14 struct { F0 uint; F1 uint16; F2 bool }
func genFn0(p0 int, p1 int8) int8 {
    for fi6 := 0; fi6 < 18; fi6++ {
        v7 := GenS1{F0: "foo bar", F1: 471}
        p0 -= fi6
        if (true && false) {
            fmt.Println(p0, p0)
            fmt.Println(fi6)
            for fi8 := 0; fi8 < 2; fi8++ {
                v9 := GenS2{F0: 255, F1: p1}
                switch fi8 {
                case 0:
                    v10 := GenS2{F0: 4294967295, F1: (p1 % ((p1) | 1))}
                    _ = v10
                case 42:
                    fmt.Println(p1, p1)
                default:
                    var v11 uint16 = 32767
                    _ = v11
                }
                _ = fi8
                _ = v9
            }
        } else {
            for fi12 := 0; fi12 < 11; fi12++ {
                if (float32(0.0) >= 0.0) {
                    if true {
                        switch "\n\t" {
                        case "hello":
                            if (65535 == 42) {
                                p1 = 0
                                switch ("\n\t" + "\n\t") {
                                case "":
                                    fmt.Println(p0, p0, fi6)
                                case "a":
                                    p0 = p0
                                case "\n\t":
                                    if (!(!((false || ("\n\t" <= "a")) && false))) {
                                        if true {
                                            switch "hello" {
                                            case "foo bar":
                                                p1 -= (p1 ^ 42)
                                            case "a":
                                                var v13 uint64 = 42
                                                _ = v13
                                            case "hello":
                                                var v14 int32 = -465
                                                _ = v14
                                            }
                                        } else {
                                            var v15 int64 = -9223372036854775808
                                            v16 := GenS3{F0: 0, F1: fi12, F2: 0}
                                            v17 := GenS4{F0: (!(1000 == 1000)), F1: int32(p0)}
                                            var v18 uint16 = 746
                                            fmt.Println(p0, v18, v15)
                                            _ = v15
                                            _ = v16
                                            _ = v17
                                            _ = v18
                                        }
                                    } else {
                                        var v19 float32 = 3.14159
                                        switch ("a" + "foo bar") {
                                        case "a":
                                            p0 = (-(fi12))
                                        default:
                                            v20 := []int{-456}
                                            _ = v20
                                        }
                                        if (!(false && (-1 < -211))) {
                                            v21 := []int{}
                                            var v22 uint32 = 2
                                            if false {
                                                p0 = p0
                                                p1 = (p1 % (((^(p1))) | 1))
                                            }
                                            v22 += 65535
                                            _ = v21
                                            _ = v22
                                        } else {
                                            var v23 float64 = (((0.0 / float64((3.14159 + (v19 * (v19 - float32(float64((v19 - v19)))))))) / 1e10) + 1e10)
                                            p1 = p1
                                            _ = v23
                                        }
                                        v24 := []int{844, 2147483647, int(p1), fi12}
                                        _ = v19
                                        _ = v24
                                    }
                                }
                            }
                        }
                        if ("\n\t" != (((("" + "") + "hello") + ("a" + "\n\t")) + ((("foo bar" + "foo bar") + "\n\t") + ("hello" + "hello")))) {
                            fmt.Println(p0, p1)
                            if (true || (!true)) {
                                v25 := []int{fi12, fi12}
                                _ = v25
                            }
                            var v26 int = fi6
                            v27 := []int{}
                            var v28 string = ""
                            _ = v26
                            _ = v27
                            _ = v28
                        }
                        var v29 float32 = 3.14159
                        _ = v29
                    }
                }
                _ = fi12
            }
            v30 := []int{930, p0, (^(fi6)), p0}
            for fi31 := 0; fi31 < 2; fi31++ {
                switch ((("foo bar" + ("a" + "")) + "foo bar") + "hello") {
                case "":
                    var v32 int32 = 2147483647
                    _ = v32
                }
                _ = len(v30)
                v33 := GenS5{F0: (("" + "") + "hello"), F1: -1}
                _ = fi31
                _ = v33
            }
            _ = v30
        }
        _ = fi6
        _ = v7
    }
    if (!((86 <= 829) || true)) {
        v34 := GenS6{F0: 2}
        var v35 bool = (!false)
        _ = v34
        _ = v35
    }
    p0 = (p0 % (((p0 >> ((18446744073709551615) & 63))) | 1))
    if (1000 < 32767) {
        fmt.Println(p1, p0)
    }
    _ = p0
    _ = p1
    return (p1 << ((4294967295) & 63))
}
func genFn1() uint64 {
    var v36 uint8 = 128
    v36 = (v36 / (((^(v36))) | 1))
    _ = v36
    return 708
}

func main() {
    v1 := make(map[uint32]string)
    v2 := []uint32{1, 42, 65535, 252}
    var v3 bool = (true && (-1 != 100))
    var v4 float64 = 1.5
    v5 := make(map[uint32]uint32)
    switch "foo bar" {
    case "\n\t":
        var v37 int16 = -1
        _ = v37
    case "hello":
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    default:
        if (!(!(550 < 857))) {
            fmt.Println(v3, v3)
            if v3 {
                switch 746 {
                case 1:
                    switch ("\n\t" + "foo bar") {
                    case "":
                        v38 := GenS4{F0: (2 != 1), F1: 0}
                        _ = v38
                    default:
                        v3 = v3
                    }
                case 148:
                    for fi39 := 0; fi39 < 12; fi39++ {
                        for fi40 := 0; fi40 < 17; fi40++ {
                            v3 = (166 == 0)
                            v3 = false
                            _ = fi40
                        }
                        _ = fi39
                    }
                case 4294967295:
                    v1[2] = ("a" + "hello")
                }
                var v41 int32 = 2147483647
                _ = v41
            } else {
                var v42 float64 = float64((((3.14159 / float32(((float64(float32((v4 - -2.25))) / (1.5 * ((0.1 + (v4 / float64(1.5))) - v4))) + v4))) / 1.0) + ((1e10 + ((1.5 + 0.0) / 1.5)) + -2.25)))
                v3 = (-128 >= 127)
                _ = v42
            }
        }
    }
    v5[542] = 680
    var v43 int = 667
    if ((!true) || v3) {
        switch "\n\t" {
        case "":
            v44 := GenS3{F0: 1, F1: v43, F2: 4294967295}
            _ = v44
        default:
            v43 = (v43 ^ v43)
        }
        var v45 string = "hello"
        _ = v45
    }
    var v46 int = v43
    if v3 {
        if false {
            v3 = v3
        } else {
            v47 := GenS7{F0: 42, F1: 42}
            var v48 uint16 = 65535
            _ = v47
            _ = v48
        }
    } else {
        fmt.Println(v46, v4)
        if false {
            switch -2147483648 {
            case 32768:
                var v49 uint8 = 255
                _ = v49
            default:
                if len(v2) > 0 {
                    v2[0] = 65535
                }
            }
        } else {
            v43 -= (v46 << ((1) & 63))
            if v3 {
                for fi50 := 0; fi50 < 7; fi50++ {
                    var v51 uint8 = 210
                    _ = fi50
                    _ = v51
                }
                var v52 int = -19
                _ = v52
            } else {
                v53 := GenS6{F0: 1}
                for fi54 := 0; fi54 < 8; fi54++ {
                    v55 := GenS8{F0: 32767, F1: 965, F2: 0}
                    switch 42 {
                    case -2147483648:
                        v46 += 9223372036854775807
                    case 1:
                        fmt.Println(v4, v46, v3)
                    case -180:
                        v2 = append(v2, 59)
                    }
                    for fi56 := 0; fi56 < 20; fi56++ {
                        v5[412] = 0
                        _ = fi56
                    }
                    _ = fi54
                    _ = v55
                }
                _ = v53
            }
        }
    }
    _ = len(v5)
    v4 += -2.25
    for fi57 := 0; fi57 < 3; fi57++ {
        fmt.Println(v46)
        if (1 >= 2) {
            v58 := GenS9{F0: 599, F1: 32767, F2: 127}
            var v59 int64 = -341
            _ = v58
            _ = v59
        }
        _ = fi57
    }
    var v60 uint = 1
    v61 := GenS9{F0: 908, F1: 591, F2: v43}
    var v62 uint = v60
    var v63 int = 833
    var v64 int32 = 261
    v1[42] = "\n\t"
    v65 := GenS10{F0: 55, F1: 233}
    if (v4 >= (v4 * (float64(((float32((-2.25 * 3.14159)) - 0.1) - 1e10)) * v4))) {
        _ = len(v5)
    } else {
        switch -627 {
        case -9223372036854775808:
            switch "" {
            case "hello":
                var v66 uint8 = 113
                _ = v66
            }
        }
        for fi67 := 0; fi67 < 7; fi67++ {
            for fi68 := 0; fi68 < 6; fi68++ {
                var v69 int16 = -194
                var v70 int = -9223372036854775808
                _ = fi68
                _ = v69
                _ = v70
            }
            fmt.Println(v62, v60, v64)
            _ = fi67
        }
        switch "foo bar" {
        case "a":
            v71 := GenS11{F0: "\n\t", F1: -9223372036854775808}
            _ = v71
        case "":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        case "\n\t":
            for fi72 := 0; fi72 < 2; fi72++ {
                var v73 uint64 = 18446744073709551615
                for fi74 := 0; fi74 < 17; fi74++ {
                    var v75 float64 = ((((v4 + v4) * (v4 / (v4 / (v4 * v4)))) * v4) - v4)
                    v76 := GenS9{F0: 978, F1: 808, F2: (fi72 * fi74)}
                    if (false || (!(!((fi74 << ((18446744073709551615) & 63)) <= -504)))) {
                        switch (v63 | v43) {
                        case 127:
                            v63 -= fi72
                        case 128:
                            if true {
                                fmt.Println(v62)
                                v77 := GenS12{F0: 1, F1: 128}
                                _ = v77
                            } else {
                                switch "foo bar" {
                                case "hello":
                                    v78 := GenS13{F0: (100 > -26), F1: 42, F2: 42}
                                    _ = v78
                                case "foo bar":
                                    if ((!(!(-1 != 118))) && v3) {
                                        if true {
                                            var v79 int = fi74
                                            var v80 uint8 = 128
                                            _ = v79
                                            _ = v80
                                        } else {
                                            delete(v1, 91)
                                            fmt.Println(v62)
                                        }
                                        if (v3 && v3) {
                                            v81 := GenS2{F0: v60, F1: genFn0(v63, 0)}
                                            v73 = (v73 & v73)
                                            v75 += v75
                                            _ = v81
                                        }
                                        v82 := GenS7{F0: (v73 - 42), F1: (v64 * (v64 ^ (v64 >> (((^(v60))) & 63))))}
                                        _ = v82
                                    } else {
                                        fmt.Println(v64)
                                        v83 := GenS1{F0: "foo bar", F1: -262}
                                        _ = v83
                                    }
                                default:
                                    fmt.Println(v43, fi74, v60)
                                }
                                v84 := GenS9{F0: 0, F1: 255, F2: 39}
                                _ = v84
                            }
                        }
                        v85 := GenS14{F0: 300, F1: 255, F2: (!v3)}
                        switch int(v62) {
                        case 0:
                            if (0 <= 4294967295) {
                                var v86 int64 = int64(v64)
                                _ = v86
                            } else {
                                v63 += int(v64)
                                var v87 float32 = -2.25
                                v88 := GenS8{F0: 0, F1: (v73 + (v73 ^ genFn1())), F2: 9223372036854775807}
                                fmt.Println(v87, v63, v64)
                                _ = v87
                                _ = v88
                            }
                        case 128:
                            switch "hello" {
                            case "a":
                                var v89 int = (^(v63))
                                _ = v89
                            case "foo bar":
                                if v3 {
                                    var v90 uint64 = (v73 & (v73 ^ (v73 * 0)))
                                    var v91 int32 = v64
                                    _ = v90
                                    _ = v91
                                }
                            }
                        default:
                            v3 = false
                        }
                        _ = v85
                    }
                    _ = fi74
                    _ = v75
                    _ = v76
                }
                v64 -= (v64 ^ v64)
                _ = fi72
                _ = v73
            }
        default:
            for fi92 := 0; fi92 < 4; fi92++ {
                fmt.Println(v60)
                v4 = v4
                if v3 {
                    v64 -= v64
                    var v93 int16 = 0
                    _ = v93
                } else {
                    var v94 int16 = -759
                    for fi95 := 0; fi95 < 11; fi95++ {
                        fmt.Println(fi95, v63)
                        v94 = -42
                        _ = fi95
                    }
                    var v96 string = "hello"
                    _ = v94
                    _ = v96
                }
                _ = fi92
            }
        }
    }
    v64 = (v64 ^ (-(v64)))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v43:", v43)
    fmt.Println("v46:", v46)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v60:", v60)
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    fmt.Println("v64:", v64)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v46
    _ = v5
    _ = v60
    _ = v61
    _ = v62
    _ = v63
    _ = v64
    _ = v65
}
