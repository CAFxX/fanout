package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint64; F1 int16; F2 string }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint8; F1 int; F2 bool }
type GenS6 struct { F0 uint32; F1 int32; F2 int64 }
type GenS7 struct { F0 bool; F1 int64 }
type GenS8 struct { F0 uint8; F1 uint64 }
func genFn0(p0 uint16, p1 int16) string {
    p1 -= (p1 >> ((4294967295) & 63))
    var v7 uint8 = 255
    var v8 bool = ("" == "\n\t")
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return "a"
}
func genFn1(p0 int32, p1 int16) int64 {
    v9 := GenS1{F0: p1, F1: 2147483647, F2: 42}
    v10 := GenS1{F0: (p1 >> ((0) & 63)), F1: 0, F2: 0}
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return -9223372036854775808
}

func main() {
    var v1 int64 = 9223372036854775807
    v2 := []int{130, 127, 360, 2147483647}
    v3 := make(map[int]uint32)
    v4 := []string{"hello", ((("" + "hello") + "a") + "")}
    var v5 int8 = 1
    var v6 int8 = (v5 & 127)
    switch "\n\t" {
    case "":
        switch "\n\t" {
        case "a":
            v6 -= (v6 % ((v6) | 1))
        case "hello":
            for fi11 := 0; fi11 < 8; fi11++ {
                switch fi11 {
                case 32767:
                    v1 = (v1 ^ genFn1(962, 1))
                case -9223372036854775808:
                    if len(v4) > 0 {
                        v4[0] = ("\n\t" + "a")
                    }
                case 32768:
                    v5 = 100
                }
                switch (("" + "\n\t") + "a") {
                case "foo bar":
                    v12 := GenS1{F0: -1, F1: (^(fi11)), F2: 157}
                    _ = v12
                default:
                    switch "hello" {
                    case "a":
                        v13 := v3[(fi11 - fi11)]
                        _ = v13
                    case "":
                        v1 = int64(v5)
                    default:
                        var v14 float64 = (-2.25 + -2.25)
                        _ = v14
                    }
                }
                v15 := GenS2{F0: (fi11 & 1)}
                _ = fi11
                _ = v15
            }
        }
    case "\n\t":
        var v16 int = 0
        _ = v16
    case "foo bar":
        var v17 uint32 = 4294967295
        _ = v17
    }
    if (!(42 >= -9223372036854775808)) {
        var v18 int = int(v6)
        _ = v18
    } else {
        var v19 uint16 = 2
        for fi20 := 0; fi20 < 2; fi20++ {
            switch -1 {
            case -150:
                switch (-(fi20)) {
                case -225:
                    v21 := GenS1{F0: 43, F1: fi20, F2: 128}
                    _ = v21
                case 467:
                    var v22 float32 = 1e10
                    _ = v22
                case -241:
                    v19 = v19
                }
            case 9223372036854775807:
                fmt.Println(v6)
            }
            _ = fi20
        }
        v1 = v1
        _ = v19
    }
    if (1 < 128) {
        switch 127 {
        case 653:
            if false {
                var v23 int = 9223372036854775807
                switch 204 {
                case 896:
                    if false {
                        var v24 uint32 = 42
                        if (18446744073709551615 <= 0) {
                            switch "hello" {
                            case "foo bar":
                                var v25 bool = (float32((((0.1 / 1.0) - ((1e10 * -2.25) * (((1e10 + (3.14159 / 3.14159)) - 3.14159) * 1e10))) + (0.1 / 1.0))) >= 0.1)
                                _ = v25
                            }
                            if false {
                                v26 := GenS2{F0: 1}
                                fmt.Println(v23, v5)
                                _ = v26
                            }
                        }
                        var v27 float64 = 0.1
                        _ = v24
                        _ = v27
                    } else {
                        for fi28 := 0; fi28 < 20; fi28++ {
                            v29 := GenS1{F0: 42, F1: (v23 << ((0) & 63)), F2: 1}
                            _ = fi28
                            _ = v29
                        }
                        fmt.Println(v6, v23, v23)
                    }
                case 685:
                    v30 := v3[(-(v23))]
                    _ = v30
                }
                if (224 != 64) {
                    var v31 float32 = 3.14159
                    _ = v31
                } else {
                    if true {
                        v5 -= (v5 - -1)
                        switch "\n\t" {
                        case "hello":
                            switch "" {
                            case "\n\t":
                                if ((v6 ^ (v6 % ((v6) | 1))) >= 42) {
                                    if ((!(("a" + "a") < "foo bar")) || (true || true)) {
                                        for fi32 := 0; fi32 < 13; fi32++ {
                                            v33 := GenS2{F0: (fi32 - int(v1))}
                                            var v34 uint8 = 127
                                            _ = fi32
                                            _ = v33
                                            _ = v34
                                        }
                                        var v35 uint64 = 42
                                        _ = v35
                                    } else {
                                        var v36 uint16 = 2
                                        _ = v36
                                    }
                                    delete(v3, (-(v23)))
                                } else {
                                    switch "" {
                                    case "hello":
                                        var v37 uint16 = 2
                                        _ = v37
                                    case "a":
                                        v1 += 0
                                    case "\n\t":
                                        var v38 float64 = 0.0
                                        _ = v38
                                    default:
                                        var v39 uint = 2
                                        _ = v39
                                    }
                                }
                            case "hello":
                                var v40 int64 = v1
                                _ = v40
                            }
                        case "":
                            v23 += (-(v23))
                        }
                        switch "\n\t" {
                        case "foo bar":
                            delete(v3, (-(v23)))
                        case "a":
                            v41 := GenS3{F0: 18446744073709551615, F1: 32767, F2: "\n\t"}
                            _ = v41
                        default:
                            if ((0 <= 2) && ((-115 < (-(v6))) && (int16(v6) >= -298))) {
                                fmt.Println(v23)
                                v42 := GenS2{F0: (v23 ^ v23)}
                                if true {
                                    if ((v1 % ((v1) | 1)) == v1) {
                                        v6 += (v5 & v5)
                                    } else {
                                        v43 := GenS4{F0: 1}
                                        var v44 int16 = 1
                                        var v45 int16 = 1000
                                        _ = v43
                                        _ = v44
                                        _ = v45
                                    }
                                    v4 = append(v4, ("a" + "foo bar"))
                                    var v46 float32 = 0.1
                                    _ = v46
                                } else {
                                    if false {
                                        if ("\n\t" > ("hello" + "\n\t")) {
                                            var v47 int16 = -32768
                                            _ = v47
                                        }
                                        fmt.Println(v1, v6, v1)
                                    } else {
                                        v48 := GenS5{F0: 42, F1: (-(v23)), F2: true}
                                        _ = v48
                                    }
                                    v49 := GenS3{F0: 909, F1: 1000, F2: ("" + ("\n\t" + "\n\t"))}
                                    v6 = 127
                                    _ = v49
                                }
                                _ = v42
                            } else {
                                v50 := v3[(v23 << ((0) & 63))]
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                _ = v50
                            }
                        }
                    } else {
                        var v51 float64 = 0.0
                        switch "foo bar" {
                        case "a":
                            if false {
                                delete(v3, 9223372036854775807)
                            } else {
                                v52 := GenS4{F0: 2147483647}
                                _ = v52
                            }
                        case "hello":
                            fmt.Println(v51, v5, v5)
                        }
                        for fi53 := 0; fi53 < 1; fi53++ {
                            fmt.Println(fi53, fi53, v1)
                            var v54 uint32 = 645
                            v55 := GenS2{F0: fi53}
                            _ = fi53
                            _ = v54
                            _ = v55
                        }
                        _ = v51
                    }
                    v56 := GenS5{F0: 80, F1: 9223372036854775807, F2: true}
                    _ = v56
                }
                _ = v23
            }
        case 42:
            v57 := v3[32767]
            _ = v57
        }
        var v58 int8 = (v6 & v5)
        _ = v58
    } else {
        switch 42 {
        case 9223372036854775807:
            if (4294967295 <= 4294967295) {
                var v59 int = 637
                v60 := GenS6{F0: 889, F1: -354, F2: -9223372036854775808}
                _ = v59
                _ = v60
            } else {
                switch "hello" {
                case "":
                    v5 = v6
                case "a":
                    v1 -= 569
                }
                v1 = (v1 * v1)
            }
        }
        fmt.Println(v1, v6, v5)
        var v61 float64 = 1e10
        _ = v61
    }
    var v62 uint = 4294967295
    if true {
        var v63 uint = 42
        _ = v63
    }
    v3[2147483647] = 2147483647
    var v64 uint8 = 127
    var v65 uint = v62
    for fi66 := 0; fi66 < 9; fi66++ {
        var v67 int16 = 7
        v1 = 0
        switch "foo bar" {
        case "foo bar":
            fmt.Println(v67)
        default:
            if (!(!(true && false))) {
                for fi68 := 0; fi68 < 19; fi68++ {
                    v62 += v65
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi68
                }
            }
        }
        _ = fi66
        _ = v67
    }
    for fi69 := 0; fi69 < 20; fi69++ {
        switch ((("\n\t" + "\n\t") + ("foo bar" + (("\n\t" + "") + "hello"))) + "a") {
        case "\n\t":
            switch (("" + "\n\t") + ("a" + (("a" + ("" + "a")) + "hello"))) {
            case "\n\t":
                for fi70 := 0; fi70 < 5; fi70++ {
                    if true {
                        if (!true) {
                            fmt.Println(v1)
                            switch fi70 {
                            case 97:
                                switch (("hello" + ("" + "")) + ("foo bar" + (("a" + "") + ((("a" + (("hello" + ("a" + (("foo bar" + ("hello" + "\n\t")) + ("\n\t" + "a")))) + ("a" + "foo bar"))) + ("" + "\n\t")) + (("a" + ("\n\t" + "\n\t")) + ("" + "\n\t")))))) {
                                case "\n\t":
                                    var v71 bool = (0.0 > 1e10)
                                    _ = v71
                                case "":
                                    v1 = v1
                                case "hello":
                                    v72 := GenS1{F0: 32767, F1: 42, F2: v64}
                                    _ = v72
                                default:
                                    v73 := GenS3{F0: 46, F1: int16(fi70), F2: (("" + "a") + "foo bar")}
                                    _ = v73
                                }
                            case 182:
                                v74 := v3[(fi69 >> (((^(v62))) & 63))]
                                _ = v74
                            default:
                                v6 -= v5
                            }
                            var v75 int = 4294967295
                            v1 -= (-(v1))
                            _ = v75
                        }
                    } else {
                        switch "" {
                        case "foo bar":
                            v76 := v3[(fi69 / (((fi70 & fi69)) | 1))]
                            _ = v76
                        case "\n\t":
                            fmt.Println(v5, fi70)
                        default:
                            var v77 string = "foo bar"
                            _ = v77
                        }
                        v5 += 127
                    }
                    var v78 uint64 = 0
                    fmt.Println(v6, fi69)
                    _ = fi70
                    _ = v78
                }
            case "":
                v6 -= v5
            default:
                if true {
                    var v79 uint = v65
                    v6 = -1
                    _ = v79
                }
            }
        default:
            switch fi69 {
            case -1:
                for fi80 := 0; fi80 < 17; fi80++ {
                    switch "a" {
                    case "foo bar":
                        v5 = 100
                    case "hello":
                        switch (fi80 + fi80) {
                        case 4294967295:
                            switch ("foo bar" + "") {
                            case "hello":
                                switch (("\n\t" + "") + "foo bar") {
                                case "\n\t":
                                    if (true && (!true)) {
                                        v6 = (^(v6))
                                    }
                                case "a":
                                    var v81 string = "hello"
                                    _ = v81
                                default:
                                    var v82 int = fi80
                                    _ = v82
                                }
                            case "foo bar":
                                v83 := GenS2{F0: (^(fi80))}
                                _ = v83
                            default:
                                fmt.Println(fi69, v5, v65)
                            }
                        case 32768:
                            if (!false) {
                                fmt.Println(fi80)
                                var v84 bool = false
                                var v85 int = (-(fi80))
                                _ = v84
                                _ = v85
                            } else {
                                v1 -= (v1 & (-(v1)))
                                fmt.Println(v5)
                                var v86 uint8 = (v64 & 128)
                                v87 := GenS3{F0: 2, F1: -32768, F2: ""}
                                v1 += genFn1(1, 1000)
                                switch -528 {
                                case 32768:
                                    fmt.Println(v62, v86, v1)
                                case 541:
                                    if false {
                                        v65 += (v62 & v65)
                                        v3[(fi69 << (((v65 % ((255) | 1))) & 63))] = 2
                                        v6 = 100
                                    } else {
                                        var v88 float64 = ((0.0 * 1e10) - (1e10 + 1.5))
                                        _ = v88
                                    }
                                default:
                                    var v89 bool = true
                                    _ = v89
                                }
                                switch "\n\t" {
                                case "":
                                    var v90 float32 = (3.14159 + float32(0.1))
                                    _ = v90
                                case "\n\t":
                                    v91 := GenS2{F0: fi69}
                                    _ = v91
                                default:
                                    if false {
                                        if true {
                                            var v92 float64 = 3.14159
                                            var v93 int = 128
                                            _ = v92
                                            _ = v93
                                        } else {
                                            fmt.Println(fi69)
                                        }
                                    }
                                }
                                v6 = 10
                                v6 = (v6 * v5)
                                _ = v86
                                _ = v87
                            }
                        case 32767:
                            if (true && (v64 != 157)) {
                                v6 = v6
                                v62 -= 42
                            } else {
                                v65 = v65
                                var v94 bool = (fi69 >= (fi80 - fi69))
                                _ = v94
                            }
                        default:
                            switch "hello" {
                            case "\n\t":
                                if ((0.1 / 1.0) <= ((0.0 / float64((float32(3.14159) + (0.1 - 3.14159)))) * 1.5)) {
                                    v95 := GenS5{F0: (^(v64)), F1: (fi80 & -99), F2: true}
                                    var v96 bool = true
                                    _ = v95
                                    _ = v96
                                }
                            case "foo bar":
                                if false {
                                    var v97 uint32 = 278
                                    v5 = v5
                                    var v98 int32 = -1
                                    _ = v97
                                    _ = v98
                                }
                            default:
                                if true {
                                    fmt.Println(v62, v62)
                                    var v99 int64 = genFn1(42, 1)
                                    var v100 uint = v62
                                    _ = v100
                                    _ = v99
                                }
                            }
                        }
                    case "":
                        v101 := GenS7{F0: false, F1: genFn1(1, -32768)}
                        _ = v101
                    default:
                        v102 := GenS7{F0: (255 == 690), F1: -1}
                        _ = v102
                    }
                    var v103 float32 = 0.0
                    var v104 int64 = (v1 * (-(v1)))
                    v105 := GenS5{F0: (^(v64)), F1: -606, F2: true}
                    _ = fi80
                    _ = v103
                    _ = v104
                    _ = v105
                }
            default:
                v6 = v6
            }
        }
        v106 := GenS8{F0: v64, F1: 916}
        v107 := GenS8{F0: v64, F1: 1}
        _ = fi69
        _ = v106
        _ = v107
    }
    for fi108 := 0; fi108 < 19; fi108++ {
        fmt.Println(v6)
        v1 -= -1
        _ = len(v2)
        _ = fi108
    }
    v109 := GenS7{F0: (1 < 1), F1: genFn1(-2147483648, -1)}
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v62:", v62)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v109
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v62
    _ = v64
    _ = v65
}
