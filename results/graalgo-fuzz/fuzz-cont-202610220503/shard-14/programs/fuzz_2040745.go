package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16; F2 string }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int; F1 uint64 }
type GenS4 struct { F0 uint; F1 int8 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint8; F1 uint16 }
type GenS8 struct { F0 int64; F1 bool; F2 string }
type GenS9 struct { F0 bool; F1 uint64; F2 string }
type GenS10 struct { F0 uint8; F1 int32 }
type GenS11 struct { F0 uint; F1 int16 }
type GenS12 struct { F0 int64; F1 int8 }
type GenS13 struct { F0 uint; F1 int32 }
type GenS14 struct { F0 int; F1 uint64 }
type GenS15 struct { F0 uint16; F1 int }
type GenS16 struct { F0 int; F1 uint }
type GenS17 struct { F0 int64; F1 int8 }
func genFn0(p0 uint64) uint8 {
    fmt.Println(p0)
    var v4 int = 881
    for fi5 := 0; fi5 < 15; fi5++ {
        var v6 float64 = 0.0
        v7 := GenS1{F0: 127, F1: 1, F2: "a"}
        _ = fi5
        _ = v6
        _ = v7
    }
    _ = p0
    _ = v4
    return 0
}
func genFn1() uint16 {
    var v8 string = "a"
    if true {
        var v9 int16 = 1
        v9 += (v9 << ((255) & 63))
        v10 := []int{-899, 19, -102}
        _ = v10
        _ = v9
    } else {
        if (651 >= 65535) {
            v11 := []int{-2147483648}
            fmt.Println(v8)
            for fi12 := 0; fi12 < 6; fi12++ {
                fmt.Println(fi12)
                if ((!false) && true) {
                    if true {
                        var v13 uint64 = 467
                        var v14 float32 = 0.0
                        v11 = append(v11, 161)
                        _ = v13
                        _ = v14
                    }
                } else {
                    v8 = v8
                    switch -9223372036854775808 {
                    case 2147483647:
                        var v15 bool = (!(269 == 42))
                        _ = v15
                    case 32767:
                        var v16 uint64 = 18446744073709551615
                        _ = v16
                    default:
                        var v17 bool = (!(false || true))
                        _ = v17
                    }
                    v8 = (("" + v8) + (v8 + v8))
                }
                _ = fi12
            }
            _ = v11
        }
        switch 9223372036854775807 {
        case 0:
            if true {
                v8 = (v8 + v8)
                v18 := GenS1{F0: 1, F1: 32767, F2: "foo bar"}
                for fi19 := 0; fi19 < 13; fi19++ {
                    switch v8 {
                    case "\n\t":
                        v8 = (("\n\t" + v8) + v8)
                    default:
                        v20 := GenS1{F0: 171, F1: 800, F2: "foo bar"}
                        _ = v20
                    }
                    v21 := GenS1{F0: 140, F1: 65535, F2: (("a" + v8) + v8)}
                    _ = fi19
                    _ = v21
                }
                _ = v18
            }
        case 32768:
            if ((!false) || (!false)) {
                v22 := []int{2147483647, 877}
                for fi23 := 0; fi23 < 1; fi23++ {
                    _ = len(v22)
                    for fi24 := 0; fi24 < 19; fi24++ {
                        v8 = ("a" + "")
                        v8 = ""
                        _ = fi24
                    }
                    _ = fi23
                }
                for fi25 := 0; fi25 < 3; fi25++ {
                    if len(v22) >= 2 {
                        v22 = v22[1:len(v22)]
                    }
                    fmt.Println(fi25)
                    var v26 int32 = -481
                    _ = fi25
                    _ = v26
                }
                _ = v22
            }
        case 2147483647:
            fmt.Println(v8)
        }
        var v27 int = -2147483648
        _ = v27
    }
    _ = v8
    return 1
}

func main() {
    v1 := []uint32{65535, 42, 2147483648}
    v2 := make(map[string]uint32)
    v3 := make(map[int]int)
    v28 := GenS2{F0: 0}
    var v29 int16 = 1
    v30 := v3[788]
    if (1 < 2147483647) {
        for fi31 := 0; fi31 < 15; fi31++ {
            if true {
                if (2 == 849) {
                    for fi32 := 0; fi32 < 13; fi32++ {
                        if false {
                            var v33 float64 = -2.25
                            v30 = 291
                            var v34 uint8 = 255
                            _ = v33
                            _ = v34
                        } else {
                            v35 := GenS2{F0: 255}
                            v36 := GenS1{F0: genFn0(425), F1: 65535, F2: "foo bar"}
                            _ = v35
                            _ = v36
                        }
                        v37 := GenS2{F0: 255}
                        _ = fi32
                        _ = v37
                    }
                } else {
                    if (-2147483648 >= 539) {
                        if (!(float32((1e10 * (1.5 + 0.0))) < 1e10)) {
                            v29 += v29
                            v2[("\n\t" + "\n\t")] = 4294967295
                            switch "foo bar" {
                            case "foo bar":
                                delete(v3, v30)
                            case "":
                                for fi38 := 0; fi38 < 3; fi38++ {
                                    var v39 uint = 4294967295
                                    v29 += (v29 - (^(v29)))
                                    _ = fi38
                                    _ = v39
                                }
                            default:
                                v40 := GenS2{F0: 42}
                                _ = v40
                            }
                        }
                        v30 = 128
                        v30 = v30
                    } else {
                        fmt.Println(v30, v30)
                        var v41 int16 = v29
                        var v42 int32 = -910
                        _ = v41
                        _ = v42
                    }
                    if (!true) {
                        v29 -= 1000
                        var v43 bool = (((false && false) || (float64((1.5 / 1e10)) <= -2.25)) && (18446744073709551615 != 301))
                        fmt.Println(v30)
                        _ = v43
                    } else {
                        var v44 float32 = 3.14159
                        v44 += 0.0
                        v45 := v2[((("\n\t" + "") + (("" + "foo bar") + ("\n\t" + ("" + "")))) + "hello")]
                        _ = v44
                        _ = v45
                    }
                    v29 += (v29 * v29)
                }
                switch ("hello" + "a") {
                case "foo bar":
                    v29 = 392
                default:
                    if false {
                        for fi46 := 0; fi46 < 15; fi46++ {
                            v47 := GenS3{F0: (v30 % ((fi46) | 1)), F1: 4294967295}
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                            switch "a" {
                            case "a":
                                if false {
                                    var v48 uint64 = 42
                                    var v49 int64 = 9223372036854775807
                                    v48 += (v48 + v48)
                                    _ = v48
                                    _ = v49
                                } else {
                                    v30 += fi31
                                }
                            default:
                                var v50 bool = false
                                if (42 == -911) {
                                    if (0.0 != 0.1) {
                                        v30 -= fi31
                                        if true {
                                            switch "" {
                                            case "hello":
                                                var v51 uint16 = 65535
                                                fmt.Println(fi31, v50)
                                                v51 -= (^(v51))
                                                v52 := GenS1{F0: 1, F1: v51, F2: "foo bar"}
                                                fmt.Println(v50, v51)
                                                _ = v51
                                                _ = v52
                                            case "a":
                                                var v53 uint8 = genFn0(4294967295)
                                                _ = v53
                                            case "\n\t":
                                                v50 = v50
                                            default:
                                                var v54 uint8 = genFn0(618)
                                                _ = v54
                                            }
                                            var v55 float32 = (float32(1.5) / 3.14159)
                                            _ = v55
                                        } else {
                                            v50 = v50
                                            if false {
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                v56 := GenS3{F0: v30, F1: 624}
                                                v1 = append(v1, 799)
                                                _ = v56
                                            } else {
                                                _ = len(v3)
                                                v57 := GenS4{F0: 65535, F1: 0}
                                                v50 = true
                                                _ = v57
                                            }
                                        }
                                    }
                                } else {
                                    if v50 {
                                        var v58 uint = 738
                                        _ = v58
                                    }
                                    v59 := GenS5{F0: 42}
                                    _ = v59
                                }
                                _ = v50
                            }
                            switch fi31 {
                            case 0:
                                if true {
                                    var v60 int = (v30 / ((285) | 1))
                                    if false {
                                        if len(v1) > 0 {
                                            v1[0] = 714
                                        }
                                    } else {
                                        fmt.Println(v30, v60, v30)
                                        _ = len(v1)
                                    }
                                    v30 -= fi46
                                    _ = v60
                                }
                            case 416:
                                switch v30 {
                                case -56:
                                    v61 := GenS6{F0: 65535}
                                    _ = v61
                                case -1:
                                    if true {
                                        v62 := GenS2{F0: 2}
                                        _ = v62
                                    } else {
                                        v30 -= (^(fi31))
                                        v29 -= v29
                                        var v63 uint8 = 42
                                        v63 = v63
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        _ = v63
                                    }
                                case 32767:
                                    if false {
                                        var v64 uint = 18446744073709551615
                                        v29 = (v29 >> (((v64 % ((v64) | 1))) & 63))
                                        v65 := v3[v30]
                                        _ = v64
                                        _ = v65
                                    } else {
                                        var v66 uint32 = 2147483647
                                        v66 = v66
                                        _ = v66
                                    }
                                }
                            }
                            if true {
                                fmt.Println(fi31, fi31, v30)
                            }
                            fmt.Println(v29, v29, v30)
                            if (42 != 65535) {
                                v30 = (^(fi46))
                                if (genFn1() != 0) {
                                    fmt.Println(fi31, fi46, fi46)
                                } else {
                                    if true {
                                        var v67 int8 = 42
                                        if true {
                                            v29 = v29
                                        }
                                        _ = v67
                                    } else {
                                        fmt.Println(v30, fi31)
                                        if false {
                                            v68 := v2[("" + "foo bar")]
                                            if true {
                                                fmt.Println(fi46)
                                                v69 := GenS7{F0: genFn0(18446744073709551615), F1: genFn1()}
                                                switch "hello" {
                                                case "foo bar":
                                                    switch ("\n\t" + "") {
                                                    case "":
                                                        v70 := GenS8{F0: 9223372036854775807, F1: (1 < -128), F2: "a"}
                                                        _ = v70
                                                    case "a":
                                                        v29 = 439
                                                    }
                                                default:
                                                    if (!(1 <= 1)) {
                                                        v71 := GenS9{F0: (-79 < -404), F1: 4294967295, F2: ""}
                                                        v72 := GenS5{F0: v68}
                                                        switch "hello" {
                                                        case "\n\t":
                                                            if false {
                                                                fmt.Println(v29, fi31, fi46)
                                                                v68 += v68
                                                                var v73 uint16 = 2
                                                                _ = v73
                                                            } else {
                                                                v74 := GenS10{F0: 127, F1: -1}
                                                                _ = v74
                                                            }
                                                        }
                                                        _ = v71
                                                        _ = v72
                                                    } else {
                                                        var v75 uint32 = v68
                                                        v29 += (v29 >> ((255) & 63))
                                                        fmt.Println(v68)
                                                        _ = v75
                                                    }
                                                }
                                                _ = v69
                                            } else {
                                                var v76 float64 = 3.14159
                                                _ = v76
                                            }
                                            v29 -= v29
                                            _ = v68
                                        }
                                    }
                                    switch fi46 {
                                    case 42:
                                        if false {
                                            v3[8] = fi31
                                        } else {
                                            var v77 int8 = 42
                                            delete(v3, (fi31 & fi31))
                                            _ = v77
                                        }
                                    case 32768:
                                        var v78 int32 = 2147483647
                                        v29 -= (-(v29))
                                        _ = v78
                                    case -1:
                                        switch (^(v30)) {
                                        case -1:
                                            v30 -= (^(fi46))
                                        default:
                                            _ = len(v1)
                                        }
                                    default:
                                        v29 += v29
                                    }
                                }
                                v29 = 42
                                fmt.Println(fi31, fi46)
                                var v79 uint32 = 723
                                v30 = 42
                                if true {
                                    var v80 int16 = v29
                                    var v81 int8 = 100
                                    if true {
                                        if ((((("foo bar" + "\n\t") + "") + "a") + "hello") < ("hello" + "a")) {
                                            var v82 int8 = (v81 / ((65) | 1))
                                            switch 128 {
                                            case 42:
                                                if (!true) {
                                                    v83 := GenS11{F0: 2, F1: v29}
                                                    _ = v83
                                                }
                                                _ = len(v2)
                                            }
                                            _ = v82
                                        } else {
                                            var v84 float32 = -2.25
                                            v85 := GenS2{F0: 255}
                                            if (false || true) {
                                                var v86 uint = 1
                                                v87 := GenS10{F0: 201, F1: -110}
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                _ = v86
                                                _ = v87
                                            }
                                            _ = v84
                                            _ = v85
                                        }
                                    }
                                    _ = v80
                                    _ = v81
                                } else {
                                    v1 = append(v1, v79)
                                    switch v30 {
                                    case -797:
                                        switch (fi31 >> ((479) & 63)) {
                                        case 490:
                                            v79 -= v79
                                        }
                                    case 50:
                                        var v88 uint32 = (^(v79))
                                        _ = v88
                                    }
                                    var v89 bool = true
                                    _ = v89
                                }
                                _ = v79
                            } else {
                                var v90 float32 = (1.5 / (0.1 - 3.14159))
                                _ = v90
                            }
                            var v91 int64 = 757
                            if ((v29 - v29) == v29) {
                                fmt.Println(v29, v30, v91)
                                fmt.Println(v91)
                                v29 = -32768
                                fmt.Println(v30)
                            }
                            _ = fi46
                            _ = v47
                            _ = v91
                        }
                    }
                }
            } else {
                var v92 uint8 = 2
                v93 := GenS9{F0: false, F1: 1, F2: "\n\t"}
                if false {
                    for fi94 := 0; fi94 < 4; fi94++ {
                        var v95 float64 = float64((0.1 * 0.0))
                        v29 = v29
                        _ = fi94
                        _ = v95
                    }
                }
                _ = v92
                _ = v93
            }
            v96 := GenS12{F0: 0, F1: -47}
            _ = fi31
            _ = v96
        }
    } else {
        switch (("\n\t" + ((("foo bar" + "hello") + "\n\t") + "foo bar")) + "foo bar") {
        case "\n\t":
            v97 := GenS1{F0: 2, F1: 1, F2: (("" + "a") + (("foo bar" + ((("hello" + "hello") + "a") + "")) + "a"))}
            _ = v97
        case "":
            if (((!(("" < "\n\t") || (1e10 == (((3.14159 - 1.5) * float64(1e10)) + (-2.25 / 0.1))))) || (1 <= 9223372036854775807)) && (165 == genFn0(0))) {
                v30 -= v30
                v98 := GenS2{F0: 0}
                var v99 uint32 = 1
                _ = v98
                _ = v99
            } else {
                var v100 int = (v30 - v30)
                fmt.Println(v29)
                _ = v100
            }
        case "hello":
            v29 = -32768
        }
    }
    for fi101 := 0; fi101 < 19; fi101++ {
        switch fi101 {
        case 127:
            v30 += v30
        default:
            for fi102 := 0; fi102 < 7; fi102++ {
                var v103 uint64 = 746
                v1 = append(v1, 304)
                if (1.5 != (3.14159 * (3.14159 - (((((1.5 / 1.0) - 0.0) / -2.25) + ((1e10 - (0.1 + 0.0)) / (0.0 + 3.14159))) * (0.0 - (-2.25 - 0.0)))))) {
                    var v104 uint16 = 11
                    var v105 int16 = 32767
                    _ = v104
                    _ = v105
                }
                _ = fi102
                _ = v103
            }
        }
        fmt.Println(v30)
        for fi106 := 0; fi106 < 20; fi106++ {
            v107 := GenS5{F0: 1}
            _ = fi106
            _ = v107
        }
        _ = fi101
    }
    if (float32(1e10) > 0.0) {
        v30 = int(v29)
        if false {
            v108 := v2["foo bar"]
            _ = v108
        }
        v1 = append(v1, 237)
    } else {
        var v109 int16 = (v29 << ((359) & 63))
        var v110 int = (v30 * v30)
        _ = v109
        _ = v110
    }
    v111 := GenS13{F0: 694, F1: -134}
    v29 += v29
    if (42 > 78) {
        switch ("a" + "hello") {
        case "hello":
            v2["a"] = 831
        case "a":
            v30 = v30
        }
        v2["\n\t"] = 65535
        v112 := GenS14{F0: (v30 >> ((255) & 63)), F1: 980}
        _ = v112
    } else {
        v29 = 32767
        for fi113 := 0; fi113 < 1; fi113++ {
            for fi114 := 0; fi114 < 5; fi114++ {
                if len(v1) > 0 {
                    v1[0] = 65535
                }
                _ = fi114
            }
            var v115 uint16 = 255
            switch (fi113 & (v30 & 719)) {
            case -1:
                for fi116 := 0; fi116 < 10; fi116++ {
                    if (!(83 < -128)) {
                        if false {
                            var v117 uint16 = v115
                            delete(v2, ("hello" + ""))
                            v118 := GenS4{F0: 42, F1: 0}
                            _ = v117
                            _ = v118
                        }
                    } else {
                        v30 -= fi116
                        v119 := GenS3{F0: fi116, F1: 18446744073709551615}
                        _ = v119
                    }
                    v120 := GenS15{F0: (v115 ^ (v115 >> ((42) & 63))), F1: fi116}
                    switch v30 {
                    case -2147483648:
                        switch fi116 {
                        case 564:
                            v121 := GenS13{F0: 18446744073709551615, F1: 671}
                            _ = v121
                        case 346:
                            v122 := GenS10{F0: 193, F1: 464}
                            _ = v122
                        default:
                            v115 -= 106
                        }
                    case 9223372036854775807:
                        var v123 float32 = float32((0.0 - 3.14159))
                        _ = v123
                    case 918:
                        v124 := GenS16{F0: fi113, F1: 65535}
                        _ = v124
                    }
                    _ = fi116
                    _ = v120
                }
            }
            _ = fi113
            _ = v115
        }
        v125 := GenS17{F0: 1, F1: 127}
        _ = v125
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v30:", v30)
    _ = v1
    _ = v111
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
}
