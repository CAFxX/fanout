package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 bool; F1 uint16 }
type GenS5 struct { F0 int; F1 uint16; F2 int8 }
type GenS6 struct { F0 uint32; F1 int8 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint32; F1 uint32 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 string }
type GenS11 struct { F0 int; F1 string; F2 uint8 }
type GenS12 struct { F0 int32; F1 int32; F2 uint32 }
type GenS13 struct { F0 uint }
type GenS14 struct { F0 int16 }
type GenS15 struct { F0 string }
type GenS16 struct { F0 string; F1 int16 }
type GenS17 struct { F0 uint8; F1 string }
type GenS18 struct { F0 uint; F1 int8; F2 int32 }
type GenS19 struct { F0 uint64; F1 string }
type GenS20 struct { F0 uint32 }
type GenS21 struct { F0 bool; F1 int; F2 int32 }
type GenS22 struct { F0 string; F1 uint8 }
type GenS23 struct { F0 string; F1 int8 }
type GenS24 struct { F0 uint16; F1 int8; F2 uint16 }
type GenS25 struct { F0 int8; F1 bool; F2 uint32 }

func main() {
    v1 := make(map[string]string)
    v2 := []int{127, -360}
    var v3 uint32 = 4294967295
    v4 := make(map[string]string)
    var v5 uint64 = 4294967295
    var v6 float64 = 0.1
    v4[(("hello" + "hello") + ("foo bar" + ("foo bar" + ("foo bar" + ""))))] = "a"
    _ = len(v2)
    fmt.Println(v3, v3, v6)
    v7 := GenS1{F0: false}
    if ((((float64(float32(3.14159)) * v6) + ((((v6 / (1e10 + 3.14159)) - v6) - v6) * -2.25)) * v6) >= (1.5 + (v6 + v6))) {
        for fi8 := 0; fi8 < 11; fi8++ {
            for fi9 := 0; fi9 < 9; fi9++ {
                if false {
                    if true {
                        if true {
                            v10 := v1["foo bar"]
                            var v11 uint = 1
                            delete(v4, (v10 + v10))
                            _ = v10
                            _ = v11
                        } else {
                            v3 = 42
                            if (2 >= 65535) {
                                switch fi9 {
                                case 128:
                                    if (((v6 - -2.25) < v6) && (1 <= 42)) {
                                        switch "hello" {
                                        case "\n\t":
                                            _ = len(v2)
                                        default:
                                            var v12 int8 = 42
                                            v13 := GenS2{F0: 1}
                                            switch "a" {
                                            case "hello":
                                                fmt.Println(v6, v12)
                                            case "\n\t":
                                                switch "hello" {
                                                case "hello":
                                                    var v14 int8 = (^(v12))
                                                    _ = v14
                                                case "\n\t":
                                                    v15 := GenS1{F0: true}
                                                    _ = v15
                                                }
                                            }
                                            _ = v12
                                            _ = v13
                                        }
                                        fmt.Println(v5, fi9)
                                        if (0 <= 1) {
                                            if true {
                                                v5 += v5
                                                v2 = append(v2, (^(fi8)))
                                                if true {
                                                    var v16 uint8 = 255
                                                    var v17 float32 = float32(float64((1.5 - (3.14159 + float32((float64(float32(float64(3.14159))) + v6))))))
                                                    _ = v16
                                                    _ = v17
                                                }
                                            } else {
                                                fmt.Println(fi8)
                                            }
                                            var v18 int = fi8
                                            _ = v18
                                        } else {
                                            switch ("hello" + "foo bar") {
                                            case "":
                                                delete(v4, (("foo bar" + ("foo bar" + ("\n\t" + "a"))) + ("" + "a")))
                                            case "foo bar":
                                                var v19 uint32 = v3
                                                _ = v19
                                            case "\n\t":
                                                v3 = (v3 | 918)
                                                v3 = 42
                                            }
                                            _ = len(v1)
                                        }
                                    }
                                case 0:
                                    v1["a"] = "\n\t"
                                    switch fi8 {
                                    case 4294967295:
                                        var v20 uint = 18446744073709551615
                                        _ = v20
                                    case 127:
                                        v5 -= (v5 * (v5 % (((^(v5))) | 1)))
                                    default:
                                        fmt.Println(fi8, fi9)
                                    }
                                    var v21 uint64 = v5
                                    fmt.Println(fi8, fi8)
                                    _ = v21
                                case -1:
                                    v22 := GenS3{F0: 0}
                                    _ = v22
                                default:
                                    switch -9223372036854775808 {
                                    case 195:
                                        v23 := v1["a"]
                                        _ = v23
                                    case 9223372036854775807:
                                        v24 := GenS3{F0: 1}
                                        _ = v24
                                    default:
                                        var v25 uint64 = v5
                                        _ = v25
                                    }
                                }
                                if false {
                                    v26 := GenS4{F0: true, F1: 255}
                                    v6 += (1e10 * -2.25)
                                    _ = v26
                                }
                            }
                            v27 := GenS5{F0: fi9, F1: 32767, F2: 1}
                            _ = v27
                        }
                        v28 := GenS3{F0: 27}
                        v6 -= float64(1e10)
                        _ = v28
                    } else {
                        fmt.Println(v3, v6, v3)
                        if (248 == 127) {
                            var v29 int = fi9
                            if (fi9 <= 42) {
                                v30 := GenS6{F0: 1, F1: -128}
                                var v31 uint = 1
                                v31 += v31
                                _ = v30
                                _ = v31
                            }
                            _ = v29
                        }
                    }
                    var v32 uint32 = 0
                    fmt.Println(v6, v3, fi8)
                    _ = v32
                }
                fmt.Println(fi9)
                _ = fi9
            }
            v6 = float64(float32(v6))
            fmt.Println(v5)
            _ = fi8
        }
        v33 := GenS7{F0: 32767}
        _ = v33
    } else {
        for fi34 := 0; fi34 < 19; fi34++ {
            if (true && false) {
                _ = len(v1)
                var v35 int = fi34
                _ = v35
            } else {
                var v36 uint32 = v3
                switch (^(fi34)) {
                case 42:
                    switch "" {
                    case "\n\t":
                        v36 = v36
                    default:
                        delete(v1, "foo bar")
                    }
                case -90:
                    v5 += (v5 * (^(v5)))
                case -1:
                    if (!("\n\t" > "foo bar")) {
                        v36 = (v36 << ((942) & 63))
                        if (255 > -76) {
                            var v37 uint32 = v3
                            var v38 int32 = 822
                            _ = v37
                            _ = v38
                        } else {
                            for fi39 := 0; fi39 < 6; fi39++ {
                                fmt.Println(v36, v36)
                                _ = fi39
                            }
                            var v40 uint64 = 2
                            var v41 int = fi34
                            _ = v40
                            _ = v41
                        }
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    } else {
                        switch "" {
                        case "hello":
                            if (true && false) {
                                v5 += v5
                                if true {
                                    var v42 string = "foo bar"
                                    switch v42 {
                                    case "a":
                                        if (!true) {
                                            var v43 int64 = -150
                                            _ = v43
                                        }
                                    case "":
                                        v6 -= (((v6 * v6) - v6) - float64(0.1))
                                    default:
                                        v42 = v42
                                    }
                                    _ = v42
                                }
                                switch fi34 {
                                case -1:
                                    var v44 uint8 = 42
                                    _ = v44
                                case -963:
                                    if (!false) {
                                        for fi45 := 0; fi45 < 13; fi45++ {
                                            v36 += 923
                                            _ = fi45
                                        }
                                        fmt.Println(v36, v36)
                                    }
                                }
                            }
                        case "a":
                            switch -9223372036854775808 {
                            case 549:
                                v46 := v4[((("a" + "foo bar") + "foo bar") + ("a" + (("a" + ("a" + (("foo bar" + ("\n\t" + ((("\n\t" + "hello") + "hello") + "hello"))) + ("foo bar" + "a")))) + ((("a" + ("\n\t" + (("a" + "\n\t") + "\n\t"))) + "") + (("a" + "\n\t") + (("" + "hello") + "\n\t"))))))]
                                _ = v46
                            case -474:
                                v36 = 577
                            }
                        case "":
                            for fi47 := 0; fi47 < 17; fi47++ {
                                var v48 int16 = 0
                                v49 := GenS8{F0: v3, F1: (v36 / ((2147483648) | 1))}
                                v50 := GenS9{F0: true}
                                _ = fi47
                                _ = v48
                                _ = v49
                                _ = v50
                            }
                        default:
                            v51 := GenS1{F0: (2 > 2)}
                            _ = v51
                        }
                        if ((^(v5)) < (v5 >> ((uint(v36)) & 63))) {
                            v52 := v1[("foo bar" + (("\n\t" + "") + ("" + ((("\n\t" + "\n\t") + "foo bar") + "foo bar"))))]
                            for fi53 := 0; fi53 < 3; fi53++ {
                                if (!(false || true)) {
                                    v1[v52] = (v52 + v52)
                                }
                                v3 = 404
                                _ = fi53
                            }
                            _ = v52
                        }
                        var v54 float64 = 0.0
                        _ = v54
                    }
                }
                v55 := GenS10{F0: ("a" + "foo bar")}
                _ = v36
                _ = v55
            }
            var v56 float32 = 1e10
            _ = fi34
            _ = v56
        }
        for fi57 := 0; fi57 < 11; fi57++ {
            v3 = v3
            _ = fi57
        }
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    v6 -= v6
    v3 -= (v3 - v3)
    v5 = (v5 % ((v5) | 1))
    if true {
        v5 = 4294967295
    }
    if false {
        var v58 uint32 = 2147483647
        v3 -= (v58 - 1)
        var v59 float32 = 1.5
        _ = v58
        _ = v59
    } else {
        if false {
            fmt.Println(v5, v5)
        } else {
            var v60 int64 = -1
            v6 -= v6
            _ = v60
        }
        if false {
            for fi61 := 0; fi61 < 13; fi61++ {
                fmt.Println(v5)
                if true {
                    v3 = (v3 >> ((1) & 63))
                    if false {
                        fmt.Println(fi61, v3, v5)
                    }
                    _ = len(v1)
                } else {
                    for fi62 := 0; fi62 < 20; fi62++ {
                        var v63 uint = 255
                        _ = len(v2)
                        _ = fi62
                        _ = v63
                    }
                }
                var v64 int16 = 32767
                _ = fi61
                _ = v64
            }
            var v65 uint = 679
            v5 = 42
            _ = v65
        } else {
            fmt.Println(v5)
            v6 += (-2.25 * v6)
            switch -9223372036854775808 {
            case 32767:
                v66 := v4[(("foo bar" + "hello") + ("hello" + ("" + "a")))]
                _ = v66
            case 128:
                switch "a" {
                case "":
                    switch 32767 {
                    case 1:
                        v67 := GenS11{F0: 4294967295, F1: (("a" + "a") + ""), F2: 1}
                        _ = v67
                    case 2147483647:
                        delete(v4, "hello")
                    case 32768:
                        v5 += (^(v5))
                    }
                case "\n\t":
                    if (!(!(true && false))) {
                        v3 = v3
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        v68 := GenS12{F0: -1, F1: 2147483647, F2: v3}
                        _ = v68
                    }
                default:
                    if (127 > 479) {
                        if true {
                            switch -9223372036854775808 {
                            case 292:
                                if false {
                                    if false {
                                        if (2 > v5) {
                                            v5 = 4294967295
                                        }
                                        var v69 int16 = 42
                                        _ = v69
                                    } else {
                                        fmt.Println(v6)
                                    }
                                    if true {
                                        switch 32768 {
                                        case 628:
                                            v3 += (^(v3))
                                        default:
                                            var v70 uint8 = 255
                                            _ = v70
                                        }
                                        var v71 int8 = -10
                                        _ = v71
                                    } else {
                                        var v72 int8 = 0
                                        var v73 int32 = -1
                                        var v74 uint16 = 65535
                                        _ = v72
                                        _ = v73
                                        _ = v74
                                    }
                                }
                            case -2147483648:
                                if false {
                                    v3 += (v3 % (((v3 * (v3 - 215))) | 1))
                                    v3 += (v3 / ((v3) | 1))
                                }
                            case 32767:
                                for fi75 := 0; fi75 < 10; fi75++ {
                                    v76 := GenS3{F0: 42}
                                    _ = fi75
                                    _ = v76
                                }
                            default:
                                v6 += float64(float32(v6))
                            }
                            for fi77 := 0; fi77 < 1; fi77++ {
                                v6 = 3.14159
                                var v78 int64 = -9223372036854775808
                                _ = fi77
                                _ = v78
                            }
                        } else {
                            v6 += (float64(0.1) / v6)
                            var v79 float64 = 0.0
                            _ = v79
                        }
                    } else {
                        v80 := GenS2{F0: 791}
                        fmt.Println(v3)
                        _ = v80
                    }
                }
            case -9223372036854775808:
                if true {
                    switch ("hello" + ((("foo bar" + "foo bar") + "hello") + "foo bar")) {
                    case "hello":
                        var v81 int8 = -49
                        _ = v81
                    case "foo bar":
                        v5 = 406
                    }
                    if (true || ((!false) && false)) {
                        v6 = float64(3.14159)
                        var v82 uint32 = 738
                        var v83 uint = 42
                        _ = v82
                        _ = v83
                    } else {
                        v5 = v5
                        v84 := GenS13{F0: 547}
                        if true {
                            v5 = (^(v5))
                            for fi85 := 0; fi85 < 8; fi85++ {
                                var v86 int16 = -32768
                                _ = fi85
                                _ = v86
                            }
                        } else {
                            if false {
                                v5 -= v5
                                if true {
                                    v6 = float64(-2.25)
                                    for fi87 := 0; fi87 < 4; fi87++ {
                                        v5 -= (v5 << ((132) & 63))
                                        switch 128 {
                                        case 42:
                                            if (42 < -834) {
                                                for fi88 := 0; fi88 < 16; fi88++ {
                                                    if ((^(v5)) <= 0) {
                                                        v89 := GenS14{F0: -826}
                                                        var v90 uint16 = 0
                                                        var v91 bool = (false || false)
                                                        _ = v89
                                                        _ = v90
                                                        _ = v91
                                                    } else {
                                                        if ((!false) || (42 >= 179)) {
                                                            var v92 int32 = -477
                                                            _ = v92
                                                        }
                                                        var v93 uint = 796
                                                        v94 := GenS4{F0: false, F1: uint16(v3)}
                                                        _ = v93
                                                        _ = v94
                                                    }
                                                    _ = fi88
                                                }
                                                var v95 bool = (945 < (^(v3)))
                                                _ = v95
                                            }
                                        case 32768:
                                            for fi96 := 0; fi96 < 8; fi96++ {
                                                if (v5 >= (v5 | 18446744073709551615)) {
                                                    var v97 uint = 65535
                                                    _ = v97
                                                } else {
                                                    var v98 int = (fi96 & fi87)
                                                    if (v5 > (v5 >> ((18446744073709551615) & 63))) {
                                                        if (1 > -116) {
                                                            if false {
                                                                switch "hello" {
                                                                case "foo bar":
                                                                    v98 -= (v98 >> ((1) & 63))
                                                                case "a":
                                                                    if ((((1.5 * 1e10) + -2.25) + 0.1) < float32((float64((1.5 * 1.5)) / -2.25))) {
                                                                        var v99 int64 = 9223372036854775807
                                                                        v100 := GenS11{F0: (v98 << ((18446744073709551615) & 63)), F1: ("foo bar" + (("hello" + (("\n\t" + "foo bar") + ("\n\t" + ""))) + ("" + "\n\t"))), F2: 0}
                                                                        _ = v100
                                                                        _ = v99
                                                                    }
                                                                case "hello":
                                                                    var v101 uint32 = (v3 + v3)
                                                                    fmt.Println(fi87)
                                                                    var v102 float64 = 0.0
                                                                    _ = v101
                                                                    _ = v102
                                                                }
                                                                v6 = v6
                                                                switch "" {
                                                                case "\n\t":
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                default:
                                                                    if false {
                                                                        delete(v1, (("foo bar" + "foo bar") + (("a" + "foo bar") + "hello")))
                                                                    }
                                                                }
                                                            } else {
                                                                var v103 uint64 = 333
                                                                if false {
                                                                    _ = len(v1)
                                                                    if true {
                                                                        _ = len(v2)
                                                                        if true {
                                                                            v3 += v3
                                                                        } else {
                                                                            switch 127 {
                                                                            case 32768:
                                                                                fmt.Println(v6)
                                                                            case 2147483647:
                                                                                v98 = fi96
                                                                            default:
                                                                                switch "foo bar" {
                                                                                case "foo bar":
                                                                                    v6 = float64(-2.25)
                                                                                case "":
                                                                                    var v104 bool = true
                                                                                    _ = v104
                                                                                }
                                                                            }
                                                                            v103 = (v103 << ((4294967295) & 63))
                                                                            var v105 int16 = -1
                                                                            _ = v105
                                                                        }
                                                                        if false {
                                                                            v103 += 746
                                                                            v106 := GenS15{F0: "hello"}
                                                                            v107 := GenS14{F0: -117}
                                                                            if true {
                                                                                v108 := GenS16{F0: "foo bar", F1: -241}
                                                                                _ = v108
                                                                            } else {
                                                                                v3 -= v3
                                                                            }
                                                                            _ = v106
                                                                            _ = v107
                                                                        } else {
                                                                            v5 = (v5 ^ (v103 - (^(v103))))
                                                                            v109 := GenS7{F0: 42}
                                                                            _ = v109
                                                                        }
                                                                        fmt.Println(v103, fi87, fi96)
                                                                        v110 := GenS17{F0: 127, F1: ("" + "")}
                                                                        _ = v110
                                                                    }
                                                                }
                                                                if false {
                                                                    var v111 uint64 = 4294967295
                                                                    if (!true) {
                                                                        var v112 int16 = 1
                                                                        if ((v98 | fi87) <= 32768) {
                                                                            var v113 uint64 = v103
                                                                            if len(v2) >= 2 {
                                                                                v2 = v2[1:len(v2)]
                                                                            }
                                                                            _ = v113
                                                                        } else {
                                                                            switch fi87 {
                                                                            case 0:
                                                                                switch v98 {
                                                                                case -798:
                                                                                    v112 = int16(v3)
                                                                                case 128:
                                                                                    if len(v2) >= 2 {
                                                                                        v2 = v2[1:len(v2)]
                                                                                    }
                                                                                    if (v5 == 42) {
                                                                                        var v114 float32 = -2.25
                                                                                        _ = v114
                                                                                    }
                                                                                    var v115 uint32 = 65535
                                                                                    _ = v115
                                                                                default:
                                                                                    var v116 int64 = -9223372036854775808
                                                                                    _ = v116
                                                                                }
                                                                            case 599:
                                                                                var v117 uint32 = 42
                                                                                _ = v117
                                                                            case 128:
                                                                                delete(v1, "\n\t")
                                                                            }
                                                                            v3 -= v3
                                                                            v3 = 4294967295
                                                                        }
                                                                        if true {
                                                                            var v118 int64 = -920
                                                                            v119 := GenS4{F0: (!false), F1: 1}
                                                                            switch (fi87 & 9223372036854775807) {
                                                                            case -9223372036854775808:
                                                                                var v120 int = (fi96 / ((4294967295) | 1))
                                                                                _ = v120
                                                                            case 0:
                                                                                v121 := GenS9{F0: true}
                                                                                _ = v121
                                                                            }
                                                                            _ = v118
                                                                            _ = v119
                                                                        } else {
                                                                            var v122 int8 = 0
                                                                            _ = v122
                                                                        }
                                                                        v111 += v103
                                                                        _ = v112
                                                                    } else {
                                                                        v98 -= fi96
                                                                        if ((!(128 != (v98 * 9223372036854775807))) || false) {
                                                                            v111 = v103
                                                                            v123 := GenS18{F0: 0, F1: 0, F2: 1}
                                                                            _ = v123
                                                                        } else {
                                                                            v124 := GenS7{F0: 0}
                                                                            _ = v124
                                                                        }
                                                                    }
                                                                    _ = v111
                                                                } else {
                                                                    v103 -= (^(v103))
                                                                    var v125 uint32 = (v3 & (v3 >> ((2) & 63)))
                                                                    if len(v2) > 0 {
                                                                        v2[0] = (fi96 ^ fi96)
                                                                    }
                                                                    v126 := GenS19{F0: v5, F1: "hello"}
                                                                    var v127 uint32 = v3
                                                                    _ = v125
                                                                    _ = v126
                                                                    _ = v127
                                                                }
                                                                _ = v103
                                                            }
                                                            var v128 bool = false
                                                            v129 := GenS1{F0: (v128 || (v128 || (((1.5 - 1.5) / ((v6 - v6) - v6)) > v6)))}
                                                            _ = v128
                                                            _ = v129
                                                        } else {
                                                            if false {
                                                                fmt.Println(v6, v3, v98)
                                                            } else {
                                                                fmt.Println(v6, v98, v5)
                                                            }
                                                            fmt.Println(v98)
                                                            v6 = (float64(float32(((((float64(0.1) * v6) - v6) / 0.1) * v6))) / v6)
                                                        }
                                                        var v130 string = "foo bar"
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                        _ = v130
                                                    } else {
                                                        v1["foo bar"] = ("foo bar" + "a")
                                                        v6 = v6
                                                        var v131 float32 = 0.1
                                                        _ = v131
                                                    }
                                                    v3 = v3
                                                    _ = v98
                                                }
                                                v3 = (v3 % ((v3) | 1))
                                                _ = fi96
                                            }
                                        case -1:
                                            var v132 int8 = -1
                                            _ = v132
                                        }
                                        _ = fi87
                                    }
                                    if (720 != 0) {
                                        var v133 int32 = -832
                                        _ = v133
                                    }
                                }
                                v134 := GenS20{F0: v3}
                                _ = v134
                            } else {
                                var v135 float32 = 3.14159
                                if (654 != 42) {
                                    if len(v2) > 0 {
                                        v2[0] = 0
                                    }
                                    if (-5 >= 345) {
                                        var v136 string = (("a" + ("foo bar" + "\n\t")) + "\n\t")
                                        _ = v136
                                    } else {
                                        for fi137 := 0; fi137 < 11; fi137++ {
                                            v135 = (v135 / 1.0)
                                            fmt.Println(v6)
                                            switch (-(fi137)) {
                                            case 42:
                                                v135 -= 3.14159
                                            case 9223372036854775807:
                                                for fi138 := 0; fi138 < 9; fi138++ {
                                                    var v139 uint = 710
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    var v140 uint = (v139 << (((v139 << (((v139 >> (((^(v139))) & 63))) & 63))) & 63))
                                                    _ = fi138
                                                    _ = v139
                                                    _ = v140
                                                }
                                            case -1:
                                                v5 = uint64(v3)
                                            }
                                            _ = fi137
                                        }
                                        v141 := GenS5{F0: 1, F1: 298, F2: -72}
                                        _ = v141
                                    }
                                } else {
                                    for fi142 := 0; fi142 < 19; fi142++ {
                                        if true {
                                            for fi143 := 0; fi143 < 14; fi143++ {
                                                v144 := GenS21{F0: true, F1: fi143, F2: 2147483647}
                                                fmt.Println(v3, v5, v5)
                                                v5 -= (v5 * (v5 & (v5 - 18446744073709551615)))
                                                _ = fi143
                                                _ = v144
                                            }
                                            v135 -= v135
                                        } else {
                                            var v145 string = ("\n\t" + "hello")
                                            v146 := GenS22{F0: v145, F1: 2}
                                            switch fi142 {
                                            case -2147483648:
                                                v5 = v5
                                            default:
                                                v6 += v6
                                            }
                                            _ = v145
                                            _ = v146
                                        }
                                        switch "hello" {
                                        case "foo bar":
                                            var v147 uint8 = 128
                                            _ = v147
                                        case "a":
                                            var v148 uint64 = v5
                                            _ = v148
                                        case "hello":
                                            v149 := GenS23{F0: "hello", F1: 127}
                                            _ = v149
                                        }
                                        _ = fi142
                                    }
                                    v6 += ((((1e10 - v6) / 1.0) * 0.0) / float64((1.5 * float32((0.1 / float64(0.1))))))
                                }
                                v6 = v6
                                _ = v135
                            }
                            if (!(128 <= 127)) {
                                for fi150 := 0; fi150 < 4; fi150++ {
                                    if false {
                                        switch "hello" {
                                        case "foo bar":
                                            fmt.Println(fi150, fi150)
                                        }
                                        v3 -= (^(v3))
                                        v5 = v5
                                    }
                                    var v151 uint64 = v5
                                    _ = fi150
                                    _ = v151
                                }
                                var v152 uint8 = 128
                                v6 += 1e10
                                _ = v152
                            } else {
                                v2 = append(v2, 833)
                            }
                            v153 := GenS11{F0: 9223372036854775807, F1: "foo bar", F2: 91}
                            _ = v153
                        }
                        _ = v84
                    }
                    var v154 float64 = (v6 + v6)
                    _ = v154
                }
            }
        }
    }
    if (-2.25 >= float32(v6)) {
        fmt.Println(v5, v5, v3)
    } else {
        if true {
            if ((false && false) && (859 < 255)) {
                fmt.Println(v6, v5, v6)
                if (112 == 128) {
                    fmt.Println(v6)
                    v3 = (v3 - (v3 | v3))
                }
            }
            if false {
                if (!false) {
                    v155 := GenS18{F0: 751, F1: -128, F2: 2147483647}
                    v156 := v1[("" + "hello")]
                    _ = v155
                    _ = v156
                }
                for fi157 := 0; fi157 < 10; fi157++ {
                    switch "\n\t" {
                    case "foo bar":
                        v158 := v1["a"]
                        _ = v158
                    case "\n\t":
                        v5 += (v5 & v5)
                    case "":
                        fmt.Println(v6, v5)
                    }
                    if (2 <= 53) {
                        v159 := v1["\n\t"]
                        delete(v4, v159)
                        _ = v159
                    } else {
                        var v160 uint64 = v5
                        _ = v160
                    }
                    _ = fi157
                }
                if true {
                    var v161 int8 = 0
                    _ = v161
                } else {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    var v162 uint16 = 0
                    var v163 int64 = 1
                    _ = v162
                    _ = v163
                }
            }
            if true {
                if (("a" + "\n\t") < "\n\t") {
                    for fi164 := 0; fi164 < 7; fi164++ {
                        v3 += v3
                        switch fi164 {
                        case -1:
                            switch (fi164 ^ (fi164 * fi164)) {
                            case -9223372036854775808:
                                fmt.Println(v5, fi164)
                            case 128:
                                v165 := v4["a"]
                                _ = v165
                            default:
                                for fi166 := 0; fi166 < 8; fi166++ {
                                    _ = len(v4)
                                    _ = fi166
                                }
                            }
                        case 227:
                            var v167 float32 = -2.25
                            _ = v167
                        default:
                            if false {
                                switch "" {
                                case "":
                                    if false {
                                        v3 += v3
                                        for fi168 := 0; fi168 < 2; fi168++ {
                                            v6 -= v6
                                            if len(v2) > 0 {
                                                v2[0] = 74
                                            }
                                            _ = fi168
                                        }
                                    }
                                case "a":
                                    fmt.Println(v5)
                                }
                            }
                        }
                        _ = fi164
                    }
                }
            }
        } else {
            var v169 int16 = 1000
            v3 -= v3
            _ = v169
        }
        if false {
            switch -2147483648 {
            case 0:
                var v170 uint16 = 20
                _ = v170
            default:
                v3 += v3
            }
            v5 = (v5 | 2)
            if (42 >= 0) {
                v171 := GenS22{F0: "\n\t", F1: 42}
                var v172 float32 = 0.1
                _ = v171
                _ = v172
            } else {
                switch ("foo bar" + "foo bar") {
                case "foo bar":
                    switch (("a" + (((("foo bar" + ("" + "")) + ((("a" + ("hello" + "foo bar")) + "a") + "")) + "") + ("a" + "hello"))) + "hello") {
                    case "hello":
                        switch 32767 {
                        case 9223372036854775807:
                            v5 = (v5 >> ((255) & 63))
                        case -681:
                            var v173 bool = false
                            _ = v173
                        case -2147483648:
                            var v174 uint64 = (v5 * (v5 * (v5 >> ((552) & 63))))
                            _ = v174
                        }
                    }
                case "":
                    v175 := GenS24{F0: 1, F1: 42, F2: 1}
                    _ = v175
                default:
                    for fi176 := 0; fi176 < 4; fi176++ {
                        var v177 int16 = -194
                        _ = fi176
                        _ = v177
                    }
                }
                if true {
                    if (!(!(1 <= 32767))) {
                        switch -764 {
                        case 32767:
                            fmt.Println(v6, v3)
                        case 128:
                            v6 -= (v6 * -2.25)
                        }
                    } else {
                        if (("hello" + "a") == ("" + "hello")) {
                            v5 = v5
                            v3 -= v3
                        }
                        if len(v2) > 0 {
                            v2[0] = -9223372036854775808
                        }
                        for fi178 := 0; fi178 < 20; fi178++ {
                            v179 := v4["foo bar"]
                            fmt.Println(v6, v5, v179)
                            _ = fi178
                            _ = v179
                        }
                    }
                    var v180 uint64 = uint64(v3)
                    v6 = ((v6 + (1e10 - v6)) + 1.5)
                    _ = v180
                } else {
                    fmt.Println(v3, v3, v5)
                    for fi181 := 0; fi181 < 17; fi181++ {
                        var v182 uint8 = 1
                        var v183 uint64 = 906
                        v3 -= v3
                        _ = fi181
                        _ = v182
                        _ = v183
                    }
                }
                _ = len(v2)
            }
        } else {
            var v184 int32 = 2147483647
            if false {
                v185 := GenS25{F0: -1, F1: true, F2: 42}
                _ = v185
            } else {
                var v186 uint16 = 0
                _ = v186
            }
            v187 := GenS23{F0: ("foo bar" + ("a" + "a")), F1: -1}
            _ = v184
            _ = v187
        }
        fmt.Println(v5, v5)
    }
    v5 = v5
    _ = len(v1)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
