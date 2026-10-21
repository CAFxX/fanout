package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int64; F2 uint8 }
type GenS2 struct { F0 uint32; F1 uint }
type GenS3 struct { F0 uint64; F1 int8; F2 int64 }
type GenS4 struct { F0 int; F1 int16 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 bool; F1 bool; F2 uint64 }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 int8; F1 int32 }
type GenS10 struct { F0 bool; F1 uint; F2 uint64 }
type GenS11 struct { F0 int64; F1 bool; F2 int64 }
type GenS12 struct { F0 uint; F1 uint32 }
type GenS13 struct { F0 int64; F1 string }
type GenS14 struct { F0 uint32; F1 uint16 }
type GenS15 struct { F0 string; F1 uint }
func genFn0(p0 bool, p1 bool) int8 {
    v6 := GenS1{F0: p0, F1: 9223372036854775807, F2: 255}
    p0 = true
    v7 := []int{979, 9223372036854775807}
    if len(v7) > 0 {
        v7[0] = 9223372036854775807
    }
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 127
}

func main() {
    v1 := make(map[string]string)
    var v2 int8 = 42
    v3 := []string{"foo bar"}
    v4 := make(map[uint32]uint32)
    var v5 int16 = -32768
    fmt.Println(v5, v2)
    v8 := GenS2{F0: 286, F1: 42}
    for fi9 := 0; fi9 < 1; fi9++ {
        v2 = v2
        if false {
            if ((!(18446744073709551615 < 42)) || ((-(v5)) != -415)) {
                switch fi9 {
                case 2147483647:
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                case -2147483648:
                    var v10 uint64 = 302
                    _ = v10
                default:
                    fmt.Println(fi9, v5, fi9)
                }
                for fi11 := 0; fi11 < 12; fi11++ {
                    switch (^(fi11)) {
                    case 1:
                        var v12 int64 = 0
                        _ = v12
                    case 127:
                        if ((v5 * (v5 & (v5 << ((1) & 63)))) <= v5) {
                            v5 += v5
                            if false {
                                if true {
                                    if (!false) {
                                        if len(v3) > 0 {
                                            v3[0] = ((("" + ((("foo bar" + "hello") + "foo bar") + "a")) + ("" + "")) + "")
                                        }
                                        v2 = (v2 ^ genFn0(((v2 >> ((0) & 63)) == 0), (0 > 4294967295)))
                                    } else {
                                        v5 = int16(v2)
                                        if (-9223372036854775808 < -782) {
                                            fmt.Println(v5)
                                        }
                                    }
                                } else {
                                    v13 := GenS3{F0: 2, F1: (v2 % ((v2) | 1)), F2: 9223372036854775807}
                                    if false {
                                        var v14 uint32 = 734
                                        if ((false || (!false)) && (2 == 65535)) {
                                            if (761 < 927) {
                                                if false {
                                                    var v15 float64 = (1.5 - 3.14159)
                                                    if false {
                                                        var v16 bool = (true || true)
                                                        v17 := GenS1{F0: (689 == 255), F1: 0, F2: 255}
                                                        var v18 uint64 = 433
                                                        _ = v16
                                                        _ = v17
                                                        _ = v18
                                                    }
                                                    _ = v15
                                                } else {
                                                    fmt.Println(fi11, v2)
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                }
                                                v5 = v5
                                                v19 := GenS3{F0: 889, F1: (v2 >> ((2) & 63)), F2: 566}
                                                if false {
                                                    fmt.Println(v2, v2, v5)
                                                }
                                                _ = v19
                                            } else {
                                                var v20 uint16 = 1000
                                                _ = v20
                                            }
                                            var v21 uint = 42
                                            _ = v21
                                        } else {
                                            v22 := GenS4{F0: fi11, F1: (v5 | int16(v2))}
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            var v23 uint16 = 693
                                            _ = v22
                                            _ = v23
                                        }
                                        fmt.Println(fi11, v2)
                                        _ = v14
                                    }
                                    _ = len(v1)
                                    _ = v13
                                }
                                var v24 uint32 = 2
                                v25 := GenS5{F0: (!true)}
                                v26 := GenS1{F0: false, F1: 9223372036854775807, F2: 255}
                                _ = v24
                                _ = v25
                                _ = v26
                            } else {
                                v27 := GenS4{F0: 128, F1: int16(v2)}
                                v2 -= v2
                                var v28 int8 = genFn0((!((!(!(!("a" < "\n\t")))) && (!(0.1 <= (1.5 - -2.25))))), false)
                                _ = v27
                                _ = v28
                            }
                        }
                    case 32767:
                        var v29 float64 = 0.0
                        _ = v29
                    }
                    if (1e10 < 0.1) {
                        if false {
                            fmt.Println(fi9, fi9)
                            var v30 uint32 = 2147483647
                            _ = v30
                        } else {
                            v2 = (v2 + -118)
                            v2 -= -37
                            fmt.Println(v2, fi11)
                            if true {
                                v5 = (v5 + 0)
                            }
                            v31 := GenS6{F0: 65535}
                            _ = v31
                        }
                        fmt.Println(fi9, fi11, v5)
                        if ((fi9 >> ((18446744073709551615) & 63)) < fi11) {
                            var v32 int64 = -124
                            _ = v32
                        } else {
                            if (127 < 17) {
                                var v33 float64 = 0.0
                                _ = v33
                            }
                            v2 -= (v2 + (v2 << ((18446744073709551615) & 63)))
                        }
                        v5 += (v5 * (v5 / (((v5 | (v5 << ((255) & 63)))) | 1)))
                        fmt.Println(v5, fi9)
                    }
                    _ = fi11
                }
            }
        }
        _ = fi9
    }
    var v34 uint = 4294967295
    switch 9223372036854775807 {
    case 4294967295:
        v5 = v5
    }
    if true {
        var v35 bool = true
        v2 = (v2 - (v2 * (^(v2))))
        _ = v35
    }
    v5 += v5
    if false {
        if true {
            v5 = v5
        } else {
            _ = len(v3)
        }
        switch -897 {
        case 2147483647:
            v34 -= 59
        case -358:
            if false {
                if false {
                    if (!false) {
                        var v36 uint8 = 2
                        switch (("\n\t" + ("a" + "hello")) + "a") {
                        case "\n\t":
                            fmt.Println(v2, v5)
                        }
                        _ = v36
                    } else {
                        if len(v3) > 0 {
                            v3[0] = "\n\t"
                        }
                        _ = len(v4)
                        switch 1 {
                        case 127:
                            v2 -= genFn0(false, true)
                        case -2147483648:
                            if true {
                                var v37 int64 = 440
                                _ = v37
                            } else {
                                v2 -= v2
                                if (4294967295 == 42) {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    var v38 uint16 = 42
                                    _ = v38
                                }
                            }
                        }
                    }
                    v39 := GenS7{F0: (true && true), F1: ((2 > 127) && false), F2: 18446744073709551615}
                    _ = v39
                }
            } else {
                for fi40 := 0; fi40 < 19; fi40++ {
                    if (true || (!true)) {
                        switch "foo bar" {
                        case "foo bar":
                            v34 = v34
                        }
                    } else {
                        fmt.Println(v34, v5, v5)
                        fmt.Println(v2, fi40, fi40)
                        fmt.Println(v2, v5)
                    }
                    switch -2147483648 {
                    case -2147483648:
                        v3 = append(v3, "hello")
                    case 380:
                        for fi41 := 0; fi41 < 1; fi41++ {
                            var v42 uint16 = 42
                            var v43 int64 = -1
                            if ("foo bar" > "foo bar") {
                                v44 := GenS7{F0: false, F1: ((^(v42)) > (v42 + 0)), F2: 337}
                                switch "hello" {
                                case "hello":
                                    switch int(v43) {
                                    case 42:
                                        var v45 uint64 = 0
                                        _ = v45
                                    case 0:
                                        var v46 uint16 = (v42 - 650)
                                        _ = v46
                                    case -1:
                                        v2 = (v2 - genFn0((!true), true))
                                    default:
                                        fmt.Println(fi41, fi41, v2)
                                    }
                                case "\n\t":
                                    if (!(1 > 4294967295)) {
                                        _ = len(v1)
                                        var v47 uint64 = 1
                                        v48 := GenS6{F0: (^(v42))}
                                        _ = v47
                                        _ = v48
                                    } else {
                                        if true {
                                            switch fi41 {
                                            case -2147483648:
                                                fmt.Println(v43)
                                            case 32768:
                                                _ = len(v3)
                                            case 0:
                                                v49 := v4[906]
                                                _ = v49
                                            }
                                            v50 := GenS8{F0: v5}
                                            var v51 int32 = 2147483647
                                            _ = v50
                                            _ = v51
                                        } else {
                                            switch "a" {
                                            case "\n\t":
                                                v2 -= v2
                                            case "":
                                                v52 := GenS9{F0: (v2 & v2), F1: int32(v42)}
                                                _ = v52
                                            case "a":
                                                v5 = (v5 / ((1000) | 1))
                                            default:
                                                v53 := GenS10{F0: (!false), F1: v34, F2: 581}
                                                _ = v53
                                            }
                                            v54 := v4[2147483648]
                                            v55 := GenS11{F0: (v43 / ((v43) | 1)), F1: (false && true), F2: v43}
                                            _ = v54
                                            _ = v55
                                        }
                                        if (!false) {
                                            var v56 int32 = -427
                                            v5 += v5
                                            switch fi41 {
                                            case -2147483648:
                                                var v57 int64 = 885
                                                _ = v57
                                            case 798:
                                                _ = len(v3)
                                            case 1:
                                                switch 127 {
                                                case 160:
                                                    v34 += (v34 << (((^(v34))) & 63))
                                                default:
                                                    switch ("a" + "a") {
                                                    case "\n\t":
                                                        switch fi41 {
                                                        case 9223372036854775807:
                                                            v34 -= 367
                                                            v43 = v43
                                                        case 42:
                                                            switch 32767 {
                                                            case -2147483648:
                                                                fmt.Println(v34)
                                                            case 9223372036854775807:
                                                                if true {
                                                                    switch "\n\t" {
                                                                    case "":
                                                                        _ = len(v1)
                                                                    case "foo bar":
                                                                        switch ("\n\t" + ((("foo bar" + "foo bar") + ("" + "foo bar")) + ("\n\t" + "\n\t"))) {
                                                                        case "a":
                                                                            var v58 float32 = (3.14159 / 3.14159)
                                                                            v3 = append(v3, "hello")
                                                                            if (!false) {
                                                                                switch "a" {
                                                                                case "foo bar":
                                                                                    if true {
                                                                                        var v59 float32 = (((3.14159 / 0.1) / v58) + v58)
                                                                                        v60 := GenS12{F0: v34, F1: 813}
                                                                                        _ = v59
                                                                                        _ = v60
                                                                                    } else {
                                                                                        switch "\n\t" {
                                                                                        case "\n\t":
                                                                                            v56 = (-(v56))
                                                                                        case "":
                                                                                            v61 := GenS13{F0: (v43 + (v43 | (v43 ^ (v43 >> (((^(v34))) & 63))))), F1: ("hello" + "hello")}
                                                                                            var v62 int8 = int8(v43)
                                                                                            v3 = append(v3, "")
                                                                                            _ = v61
                                                                                            _ = v62
                                                                                        default:
                                                                                            if false {
                                                                                                var v63 int32 = (v56 - 0)
                                                                                                v42 = 749
                                                                                                _ = v63
                                                                                            } else {
                                                                                                fmt.Println(v43, fi41)
                                                                                            }
                                                                                        }
                                                                                        if (563 <= (^(v34))) {
                                                                                            v2 = v2
                                                                                            if (!(404 != 114)) {
                                                                                                v43 = (v43 / ((v43) | 1))
                                                                                            }
                                                                                            v64 := GenS7{F0: false, F1: false, F2: 1}
                                                                                            var v65 float64 = 3.14159
                                                                                            switch "hello" {
                                                                                            case "foo bar":
                                                                                                switch int(v2) {
                                                                                                case 127:
                                                                                                    var v66 bool = true
                                                                                                    _ = v66
                                                                                                }
                                                                                            case "hello":
                                                                                                switch (("foo bar" + "\n\t") + "\n\t") {
                                                                                                case "foo bar":
                                                                                                    var v67 int32 = v56
                                                                                                    _ = v67
                                                                                                case "hello":
                                                                                                    v56 = (v56 / (((v56 * v56)) | 1))
                                                                                                default:
                                                                                                    _ = len(v1)
                                                                                                }
                                                                                            case "a":
                                                                                                v58 -= float32(v65)
                                                                                            default:
                                                                                                var v68 uint64 = 2
                                                                                                _ = v68
                                                                                            }
                                                                                            var v69 bool = (!(int(v5) != (fi40 >> (((v34 & 65535)) & 63))))
                                                                                            _ = v64
                                                                                            _ = v65
                                                                                            _ = v69
                                                                                        }
                                                                                        v58 = ((v58 * 3.14159) + (float32((1e10 / ((1e10 / 1.5) + float64(-2.25)))) * (v58 * v58)))
                                                                                    }
                                                                                    var v70 uint32 = 800
                                                                                    if true {
                                                                                        if false {
                                                                                            v2 -= v2
                                                                                        } else {
                                                                                            v71 := GenS14{F0: v70, F1: 527}
                                                                                            switch fi41 {
                                                                                            case -2147483648:
                                                                                                if (v2 < (v2 * (v2 >> ((v34) & 63)))) {
                                                                                                    v72 := GenS15{F0: "a", F1: (v34 >> ((v34) & 63))}
                                                                                                    v2 = (v2 / ((-128) | 1))
                                                                                                    var v73 int64 = int64(v70)
                                                                                                    _ = v72
                                                                                                    _ = v73
                                                                                                } else {
                                                                                                    switch (fi41 | fi40) {
                                                                                                    case 1:
                                                                                                        v3 = append(v3, "a")
                                                                                                    }
                                                                                                    v34 = (v34 + (v34 & (v34 | (v34 - (v34 >> ((255) & 63))))))
                                                                                                    var v74 int8 = genFn0(false, false)
                                                                                                    v70 += (v70 / ((2147483648) | 1))
                                                                                                    _ = v74
                                                                                                }
                                                                                            default:
                                                                                                v58 += (v58 * (v58 / float32(((0.1 + (1e10 + 1e10)) - 0.1))))
                                                                                            }
                                                                                            _ = v71
                                                                                        }
                                                                                        var v75 float64 = 0.1
                                                                                        _ = v75
                                                                                    }
                                                                                    if len(v3) >= 2 {
                                                                                        v3 = v3[1:len(v3)]
                                                                                    }
                                                                                    var v76 string = "foo bar"
                                                                                    _ = v70
                                                                                    _ = v76
                                                                                case "hello":
                                                                                    _ = len(v1)
                                                                                default:
                                                                                    var v77 int8 = genFn0(true, true)
                                                                                    _ = v77
                                                                                }
                                                                                v42 = v42
                                                                                var v78 float64 = 0.1
                                                                                _ = v78
                                                                            }
                                                                            _ = v58
                                                                        default:
                                                                            if ((0 != 128) && (true && false)) {
                                                                                var v79 int = fi41
                                                                                var v80 uint = (v34 | (v34 & (v34 * (^(v34)))))
                                                                                _ = v79
                                                                                _ = v80
                                                                            }
                                                                        }
                                                                    case "a":
                                                                        fmt.Println(v34, v43, v56)
                                                                    }
                                                                    fmt.Println(v42, v56, v43)
                                                                } else {
                                                                    switch (fi41 >> (((^(v34))) & 63)) {
                                                                    case 9223372036854775807:
                                                                        if false {
                                                                            v43 += -1
                                                                        }
                                                                        v42 += (v42 - v42)
                                                                    case 42:
                                                                        v81 := GenS2{F0: 1, F1: 573}
                                                                        _ = v81
                                                                    case 922:
                                                                        var v82 int8 = -128
                                                                        _ = v82
                                                                    default:
                                                                        fmt.Println(v34, v34, v2)
                                                                    }
                                                                    var v83 int32 = (v56 >> (((v34 * v34)) & 63))
                                                                    _ = v83
                                                                }
                                                            case 1:
                                                                var v84 int16 = 1000
                                                                _ = v84
                                                            }
                                                        }
                                                    default:
                                                        var v85 int16 = v5
                                                        _ = v85
                                                    }
                                                }
                                            }
                                            _ = v56
                                        }
                                    }
                                case "":
                                    var v86 int8 = v2
                                    _ = v86
                                default:
                                    v87 := GenS8{F0: v5}
                                    _ = v87
                                }
                                _ = v44
                            }
                            _ = fi41
                            _ = v42
                            _ = v43
                        }
                    case -621:
                        switch "\n\t" {
                        case "a":
                            for fi88 := 0; fi88 < 12; fi88++ {
                                delete(v1, "\n\t")
                                _ = fi88
                            }
                        }
                    }
                    _ = fi40
                }
                if ((("hello" + "hello") + "\n\t") != "a") {
                    v89 := GenS6{F0: 2}
                    var v90 float32 = 0.0
                    _ = v89
                    _ = v90
                } else {
                    if len(v3) > 0 {
                        v3[0] = "foo bar"
                    }
                    for fi91 := 0; fi91 < 4; fi91++ {
                        v3 = append(v3, ("a" + "foo bar"))
                        _ = fi91
                    }
                }
                var v92 string = "\n\t"
                _ = v92
            }
        default:
            if ("foo bar" > ("\n\t" + "\n\t")) {
                v34 = 369
                for fi93 := 0; fi93 < 19; fi93++ {
                    var v94 uint32 = 2
                    _ = fi93
                    _ = v94
                }
            }
        }
    }
    fmt.Println(v2)
    var v95 int = -917
    v96 := GenS4{F0: v95, F1: -778}
    v95 = v95
    var v97 uint8 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v95:", v95)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v2
    _ = v3
    _ = v34
    _ = v4
    _ = v5
    _ = v8
    _ = v95
    _ = v96
    _ = v97
}
