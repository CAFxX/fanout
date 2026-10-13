package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint8 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 uint16; F1 int }
type GenS4 struct { F0 int16; F1 uint64 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 uint64; F1 uint64; F2 int32 }
type GenS7 struct { F0 int16; F1 int16; F2 int16 }
type GenS8 struct { F0 uint; F1 string }
type GenS9 struct { F0 int64 }
type GenS10 struct { F0 uint; F1 uint; F2 bool }
type GenS11 struct { F0 int }
type GenS12 struct { F0 int32; F1 string; F2 uint64 }
type GenS13 struct { F0 int; F1 uint32 }
func genFn0() bool {
    var v6 int8 = -101
    if (1 != -1) {
        v7 := GenS1{F0: ("" + (("hello" + "a") + "foo bar")), F1: 2}
        switch "foo bar" {
        case "\n\t":
            switch "hello" {
            case "\n\t":
                var v8 int8 = -2
                _ = v8
            }
        case "hello":
            var v9 uint8 = 255
            _ = v9
        default:
            switch "\n\t" {
            case "foo bar":
                v10 := GenS2{F0: 0}
                _ = v10
            case "\n\t":
                if (2 <= 42) {
                    var v11 int32 = 42
                    var v12 int8 = v6
                    v13 := GenS2{F0: 41}
                    _ = v11
                    _ = v12
                    _ = v13
                }
            case "hello":
                v6 -= v6
            default:
                v6 = (v6 & (^(v6)))
            }
        }
        _ = v7
    } else {
        for fi14 := 0; fi14 < 19; fi14++ {
            var v15 uint32 = 65535
            _ = fi14
            _ = v15
        }
        v6 -= (v6 & (v6 << ((4294967295) & 63)))
        v16 := GenS3{F0: 1, F1: 9223372036854775807}
        _ = v16
    }
    if (42 != (v6 % ((v6) | 1))) {
        switch "a" {
        case "foo bar":
            v6 = v6
        case "hello":
            var v17 uint64 = 868
            _ = v17
        }
        v18 := []int{-2147483648, -51, 128}
        _ = v18
    } else {
        if true {
            var v19 float64 = (3.14159 - 1e10)
            v20 := []int{4294967295, 4294967295, -355}
            _ = v19
            _ = v20
        }
        switch "a" {
        case "\n\t":
            for fi21 := 0; fi21 < 8; fi21++ {
                var v22 uint64 = 2
                switch (fi21 + (^(fi21))) {
                case 32768:
                    for fi23 := 0; fi23 < 8; fi23++ {
                        fmt.Println(v22)
                        v24 := []int{0, fi21, 128, 32767}
                        if len(v24) >= 2 {
                            v24 = v24[1:len(v24)]
                        }
                        fmt.Println(v6, fi23)
                        switch "a" {
                        case "":
                            var v25 int = 128
                            _ = v25
                        case "a":
                            v22 = v22
                        case "foo bar":
                            var v26 uint8 = 255
                            _ = v26
                        default:
                            if (!true) {
                                if len(v24) > 0 {
                                    v24[0] = fi21
                                }
                                var v27 uint = 4294967295
                                _ = v27
                            } else {
                                if ((true || false) && true) {
                                    var v28 int8 = v6
                                    v29 := GenS2{F0: 2147483648}
                                    var v30 uint8 = 128
                                    _ = v28
                                    _ = v29
                                    _ = v30
                                }
                                var v31 int8 = 127
                                if (1 == 0) {
                                    v24 = append(v24, fi21)
                                } else {
                                    var v32 int8 = v31
                                    var v33 uint16 = 477
                                    switch -254 {
                                    case -827:
                                        if (238 >= 31) {
                                            _ = len(v24)
                                        } else {
                                            if len(v24) > 0 {
                                                v24[0] = fi23
                                            }
                                            v22 = (v22 / (((^(v22))) | 1))
                                            if len(v24) >= 2 {
                                                v24 = v24[1:len(v24)]
                                            }
                                            fmt.Println(v31)
                                            v34 := GenS1{F0: ((("foo bar" + "foo bar") + "a") + "a"), F1: 187}
                                            _ = v34
                                        }
                                    case 2147483647:
                                        v35 := GenS4{F0: 558, F1: 349}
                                        _ = v35
                                    default:
                                        var v36 uint64 = v22
                                        _ = v36
                                    }
                                    if len(v24) >= 2 {
                                        v24 = v24[1:len(v24)]
                                    }
                                    if true {
                                        v32 = v32
                                    } else {
                                        if ((fi21 > (fi23 / ((fi21) | 1))) || false) {
                                            v22 = 0
                                        } else {
                                            fmt.Println(v31)
                                        }
                                        v37 := GenS2{F0: 714}
                                        _ = v37
                                    }
                                    _ = v32
                                    _ = v33
                                }
                                var v38 uint16 = 0
                                _ = v31
                                _ = v38
                            }
                        }
                        var v39 string = "hello"
                        _ = fi23
                        _ = v24
                        _ = v39
                    }
                case 128:
                    for fi40 := 0; fi40 < 6; fi40++ {
                        switch (fi40 / ((-497) | 1)) {
                        case 722:
                            var v41 int16 = -32768
                            _ = v41
                        case 215:
                            if false {
                                if false {
                                    var v42 uint64 = v22
                                    v43 := GenS5{F0: false}
                                    v6 = (v6 >> ((722) & 63))
                                    _ = v42
                                    _ = v43
                                } else {
                                    var v44 uint16 = 32767
                                    _ = v44
                                }
                                v45 := GenS6{F0: 18446744073709551615, F1: 2, F2: 192}
                                _ = v45
                            } else {
                                v46 := GenS7{F0: -436, F1: 32767, F2: int16(fi21)}
                                v6 = (v6 | 42)
                                switch fi21 {
                                case -467:
                                    if true {
                                        var v47 string = ((("a" + ("\n\t" + (("hello" + ((("foo bar" + "hello") + "foo bar") + "\n\t")) + ("\n\t" + "\n\t")))) + "foo bar") + ("foo bar" + "hello"))
                                        _ = v47
                                    } else {
                                        v48 := []int{}
                                        _ = v48
                                    }
                                case 32767:
                                    if false {
                                        v49 := GenS8{F0: uint(v22), F1: "foo bar"}
                                        _ = v49
                                    } else {
                                        v6 = v6
                                        v22 = v22
                                        v50 := []int{}
                                        _ = v50
                                    }
                                }
                                if false {
                                    var v51 uint64 = v22
                                    v6 -= 42
                                    if (1 < 1) {
                                        if true {
                                            var v52 int = fi21
                                            _ = v52
                                        }
                                        if true {
                                            v53 := GenS1{F0: ("\n\t" + "a"), F1: 142}
                                            _ = v53
                                        }
                                    } else {
                                        var v54 int16 = -910
                                        var v55 int64 = -915
                                        var v56 uint64 = (^(v51))
                                        _ = v54
                                        _ = v55
                                        _ = v56
                                    }
                                    _ = v51
                                } else {
                                    v57 := GenS3{F0: 634, F1: (^(fi21))}
                                    _ = v57
                                }
                                v22 = v22
                                var v58 uint8 = 146
                                _ = v46
                                _ = v58
                            }
                        }
                        if ("\n\t" > "hello") {
                            v59 := []int{fi40, fi21}
                            if ((^(fi21)) == -1) {
                                if (("hello" + "\n\t") == "foo bar") {
                                    if true {
                                        fmt.Println(v6)
                                        var v60 string = ""
                                        v61 := GenS2{F0: 1}
                                        _ = v60
                                        _ = v61
                                    }
                                    var v62 uint8 = 131
                                    fmt.Println(v22, v22, fi21)
                                    var v63 uint32 = 2147483647
                                    var v64 int8 = v6
                                    _ = v62
                                    _ = v63
                                    _ = v64
                                }
                                _ = len(v59)
                            } else {
                                var v65 float32 = 3.14159
                                _ = v65
                            }
                            if len(v59) > 0 {
                                v59[0] = 4294967295
                            }
                            _ = v59
                        } else {
                            v66 := []int{(fi40 / (((fi40 & fi21)) | 1)), (fi21 | fi40), -1, (fi21 % ((fi40) | 1))}
                            _ = v66
                        }
                        v22 += (^(v22))
                        var v67 string = ""
                        _ = fi40
                        _ = v67
                    }
                }
                if (!(false && ((false || true) && (!(true || (!((v22 >> ((255) & 63)) < (v22 * v22)))))))) {
                    v6 = v6
                    switch 128 {
                    case 42:
                        var v68 float64 = 0.0
                        _ = v68
                    case 526:
                        for fi69 := 0; fi69 < 20; fi69++ {
                            v22 -= v22
                            if false {
                                fmt.Println(v6, v22, fi69)
                                v6 = (v6 % ((v6) | 1))
                            } else {
                                var v70 uint64 = 998
                                if true {
                                    v71 := []int{}
                                    if ((true && (42 < 1)) || (!true)) {
                                        _ = len(v71)
                                        var v72 float32 = (1e10 - (1.5 - (float32(-2.25) - (-2.25 + 3.14159))))
                                        _ = v72
                                    } else {
                                        var v73 uint = 95
                                        _ = v73
                                    }
                                    fmt.Println(v22)
                                    _ = v71
                                }
                                _ = v70
                            }
                            if (!true) {
                                fmt.Println(fi69)
                            }
                            _ = fi69
                        }
                    case -2147483648:
                        v74 := GenS9{F0: 561}
                        _ = v74
                    default:
                        v75 := GenS8{F0: 2, F1: "hello"}
                        _ = v75
                    }
                    v6 = 1
                }
                _ = fi21
                _ = v22
            }
        }
    }
    v6 = (v6 << ((622) & 63))
    var v76 float64 = ((((((3.14159 - 1.5) / -2.25) * 0.1) - float64((-2.25 + -2.25))) / (3.14159 / 1.0)) - 0.0)
    _ = v6
    _ = v76
    return ("a" > (("foo bar" + ("\n\t" + "")) + ("foo bar" + "")))
}
func genFn1() uint32 {
    var v77 int32 = -1
    v77 += 0
    if true {
        switch -1 {
        case 128:
            v78 := []int{}
            _ = v78
        case -680:
            v79 := GenS10{F0: 2, F1: 18446744073709551615, F2: (2147483647 <= 658)}
            _ = v79
        case -2147483648:
            v80 := []int{128, 2147483647, 9223372036854775807, 0}
            _ = v80
        }
    }
    for fi81 := 0; fi81 < 9; fi81++ {
        fmt.Println(fi81)
        _ = fi81
    }
    _ = v77
    return 4294967295
}

func main() {
    var v1 uint16 = 17
    v2 := make(map[string]uint32)
    var v3 uint = uint(v1)
    var v4 uint32 = 1
    v5 := []string{("hello" + "hello"), "", "hello", "hello"}
    if true {
        fmt.Println(v3, v1, v4)
    } else {
        v3 = (v3 / ((v3) | 1))
        _ = len(v5)
        if false {
            if false {
                for fi82 := 0; fi82 < 3; fi82++ {
                    v1 += (v1 | v1)
                    fmt.Println(v3)
                    _ = fi82
                }
                if (-9223372036854775808 > -2147483648) {
                    if ((0 > 1) && true) {
                        if (!(false && false)) {
                            var v83 int8 = 98
                            _ = v83
                        }
                    }
                    var v84 uint8 = 2
                    var v85 bool = false
                    _ = v84
                    _ = v85
                }
            }
        }
    }
    fmt.Println(v3)
    switch 128 {
    case 32767:
        switch -2147483648 {
        case 9223372036854775807:
            v3 += (v3 >> ((v3) & 63))
        default:
            v1 = v1
        }
    }
    if true {
        v3 = v3
        switch 2147483647 {
        case -2147483648:
            var v86 int = 32768
            _ = v86
        case 472:
            var v87 float32 = float32(float64(((0.0 + (-2.25 / 3.14159)) / (float32(-2.25) + 0.0))))
            _ = v87
        case 58:
            for fi88 := 0; fi88 < 1; fi88++ {
                if (v4 >= (v4 - v4)) {
                    switch fi88 {
                    case 127:
                        v4 = 84
                    case 4294967295:
                        var v89 uint8 = 53
                        _ = v89
                    default:
                        v4 -= 0
                    }
                }
                v3 = uint(v1)
                _ = fi88
            }
        default:
            v2[("" + "a")] = (v4 + 868)
        }
    }
    var v90 uint64 = 4294967295
    var v91 int16 = 32767
    v92 := GenS8{F0: uint(v1), F1: "hello"}
    v93 := GenS11{F0: 127}
    v90 -= (^(v90))
    for fi94 := 0; fi94 < 6; fi94++ {
        var v95 float64 = ((((-2.25 * float64(0.1)) + 3.14159) - 0.0) * float64(0.0))
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        if false {
            if false {
                fmt.Println(v90, v1, v3)
            } else {
                fmt.Println(v3, v95, v95)
                v96 := GenS12{F0: 847, F1: "foo bar", F2: 0}
                _ = v96
            }
            var v97 bool = false
            _ = v97
        }
        _ = fi94
        _ = v95
    }
    fmt.Println(v90)
    v90 += (v90 | 2)
    switch -337 {
    case 9223372036854775807:
        _ = len(v5)
    case 42:
        v98 := GenS10{F0: (v3 - v3), F1: (^(v3)), F2: (!("\n\t" != "foo bar"))}
        _ = v98
    case 0:
        switch "hello" {
        case "":
            v4 = 768
        case "\n\t":
            for fi99 := 0; fi99 < 8; fi99++ {
                var v100 int64 = 9223372036854775807
                switch -1 {
                case 0:
                    for fi101 := 0; fi101 < 14; fi101++ {
                        v102 := GenS6{F0: v90, F1: v90, F2: -1}
                        var v103 int64 = (v100 ^ (v100 ^ v100))
                        _ = len(v2)
                        var v104 int32 = 2147483647
                        _ = fi101
                        _ = v102
                        _ = v103
                        _ = v104
                    }
                default:
                    switch "a" {
                    case "":
                        for fi105 := 0; fi105 < 6; fi105++ {
                            var v106 int64 = -9223372036854775808
                            switch -1 {
                            case 1:
                                var v107 float64 = 0.0
                                _ = v107
                            case 42:
                                var v108 int16 = int16(v100)
                                _ = v108
                            case 493:
                                v109 := GenS13{F0: 1, F1: (v4 ^ v4)}
                                _ = v109
                            }
                            _ = fi105
                            _ = v106
                        }
                    default:
                        for fi110 := 0; fi110 < 15; fi110++ {
                            v1 = 2
                            switch fi99 {
                            case -844:
                                fmt.Println(fi110)
                            }
                            _ = fi110
                        }
                    }
                }
                _ = fi99
                _ = v100
            }
        case "a":
            switch "a" {
            case "foo bar":
                fmt.Println(v4, v3)
            case "a":
                v111 := v2[""]
                _ = v111
            }
        }
    }
    if false {
        for fi112 := 0; fi112 < 11; fi112++ {
            for fi113 := 0; fi113 < 3; fi113++ {
                var v114 int32 = -2147483648
                _ = fi113
                _ = v114
            }
            switch "a" {
            case "":
                v115 := GenS10{F0: v3, F1: (v3 + 1), F2: false}
                _ = v115
            default:
                for fi116 := 0; fi116 < 5; fi116++ {
                    v117 := GenS2{F0: 808}
                    var v118 float64 = 1.5
                    var v119 int16 = (v91 >> ((v3) & 63))
                    var v120 bool = (42 < 7)
                    _ = fi116
                    _ = v117
                    _ = v118
                    _ = v119
                    _ = v120
                }
            }
            _ = fi112
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v90:", v90)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v90
    _ = v91
    _ = v92
    _ = v93
}
