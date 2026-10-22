package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint32; F1 int64 }
type GenS4 struct { F0 int64; F1 int16 }
type GenS5 struct { F0 uint16; F1 uint8 }
type GenS6 struct { F0 int64; F1 int32; F2 uint32 }

func main() {
    var v1 float64 = 3.14159
    v2 := make(map[string]bool)
    v3 := make(map[uint32]string)
    v4 := GenS1{F0: "foo bar"}
    switch "foo bar" {
    case "hello":
        if false {
            v5 := GenS1{F0: ""}
            fmt.Println(v1)
            _ = v5
        } else {
            v1 += v1
            var v6 uint32 = 2147483648
            _ = v6
        }
    default:
        for fi7 := 0; fi7 < 17; fi7++ {
            switch 32767 {
            case 127:
                if (!(!true)) {
                    v1 = 0.1
                    delete(v3, 4294967295)
                } else {
                    if (42 != 42) {
                        v8 := []int{}
                        var v9 float32 = 3.14159
                        _ = v8
                        _ = v9
                    } else {
                        for fi10 := 0; fi10 < 18; fi10++ {
                            v1 = v1
                            v1 -= v1
                            if false {
                                if true {
                                    fmt.Println(fi10, fi7)
                                    v11 := GenS1{F0: ((("foo bar" + "foo bar") + "hello") + "\n\t")}
                                    _ = v11
                                }
                                v12 := GenS2{F0: 1}
                                switch ("a" + "a") {
                                case "a":
                                    var v13 string = "hello"
                                    _ = v13
                                }
                                _ = v12
                            }
                            _ = fi10
                        }
                        v14 := GenS2{F0: 1}
                        _ = v14
                    }
                }
            }
            var v15 int16 = -507
            var v16 string = "\n\t"
            _ = fi7
            _ = v15
            _ = v16
        }
    }
    for fi17 := 0; fi17 < 10; fi17++ {
        switch "" {
        case "\n\t":
            var v18 uint8 = 1
            _ = v18
        case "hello":
            if ((45 < 543) && true) {
                v19 := GenS1{F0: ("foo bar" + (("foo bar" + ("\n\t" + "foo bar")) + "\n\t"))}
                fmt.Println(v1)
                if true {
                    if (!false) {
                        var v20 float64 = (v1 - 0.1)
                        _ = v20
                    }
                    v21 := []int{32768, fi17, (fi17 ^ (-(fi17)))}
                    var v22 uint16 = 1
                    _ = v21
                    _ = v22
                } else {
                    v1 += (((float64((-2.25 - (-2.25 / 1e10))) - 1e10) * v1) - 3.14159)
                    fmt.Println(v1, fi17)
                    var v23 int32 = -1
                    _ = v23
                }
                _ = v19
            }
        case "a":
            var v24 string = ((("hello" + "\n\t") + "\n\t") + "hello")
            _ = v24
        default:
            for fi25 := 0; fi25 < 18; fi25++ {
                v1 += v1
                v1 = float64(((0.0 - 0.1) - float32(v1)))
                _ = fi25
            }
        }
        switch fi17 {
        case 42:
            v26 := GenS2{F0: 26}
            _ = v26
        }
        _ = fi17
    }
    switch 4294967295 {
    case 42:
        var v27 uint = 65535
        _ = v27
    case 542:
        if true {
            var v28 float64 = v1
            var v29 int8 = 116
            for fi30 := 0; fi30 < 13; fi30++ {
                if (false || (42 == 0)) {
                    switch (((("foo bar" + ("" + "\n\t")) + "") + "hello") + "foo bar") {
                    case "a":
                        v1 -= v1
                    default:
                        v29 -= v29
                    }
                    var v31 string = ""
                    switch "hello" {
                    case "":
                        var v32 string = v31
                        _ = v32
                    case "hello":
                        v28 = 0.1
                    default:
                        v33 := GenS2{F0: -1}
                        _ = v33
                    }
                    _ = v31
                } else {
                    for fi34 := 0; fi34 < 3; fi34++ {
                        var v35 int32 = -895
                        if true {
                            var v36 uint64 = 18446744073709551615
                            var v37 uint32 = 0
                            _ = v36
                            _ = v37
                        }
                        _ = fi34
                        _ = v35
                    }
                    fmt.Println(v1)
                    v29 -= v29
                }
                var v38 uint16 = 0
                var v39 int16 = 32767
                _ = fi30
                _ = v38
                _ = v39
            }
            _ = v28
            _ = v29
        } else {
            for fi40 := 0; fi40 < 4; fi40++ {
                v41 := GenS3{F0: 2147483647, F1: 235}
                v1 -= 0.1
                v42 := GenS2{F0: 42}
                _ = fi40
                _ = v41
                _ = v42
            }
            if true {
                v43 := []int{-2147483648, 32768, 2147483647}
                var v44 uint16 = 2
                v43 = append(v43, 128)
                _ = v43
                _ = v44
            } else {
                var v45 bool = (-763 != 578)
                _ = v45
            }
            var v46 int16 = -681
            _ = v46
        }
    default:
        fmt.Println(v1)
    }
    v1 = ((v1 / (v1 / v1)) + v1)
    for fi47 := 0; fi47 < 19; fi47++ {
        if true {
            var v48 float64 = v1
            _ = v48
        } else {
            for fi49 := 0; fi49 < 10; fi49++ {
                var v50 int64 = -1
                v1 = (((v1 * -2.25) + (float64(1e10) / 1.5)) - (0.0 / (float64(float32((0.1 * 0.1))) - (v1 / ((float64(-2.25) * (1e10 + (3.14159 - 3.14159))) + v1)))))
                v51 := GenS2{F0: -128}
                _ = fi49
                _ = v50
                _ = v51
            }
            v52 := []int{}
            var v53 int64 = -9223372036854775808
            _ = v52
            _ = v53
        }
        switch (^(fi47)) {
        case -9223372036854775808:
            var v54 uint16 = 42
            _ = v54
        case 0:
            for fi55 := 0; fi55 < 6; fi55++ {
                v1 = (v1 + float64(3.14159))
                var v56 uint32 = 65535
                v1 = (v1 / 0.1)
                _ = fi55
                _ = v56
            }
        case 128:
            if (true || (0 > 296)) {
                _ = len(v2)
            } else {
                switch ("foo bar" + "hello") {
                case "a":
                    for fi57 := 0; fi57 < 3; fi57++ {
                        v58 := GenS1{F0: ""}
                        v1 += (0.1 * ((v1 / ((v1 - v1) + (1.5 - v1))) * ((float64(((3.14159 / 1.0) / 3.14159)) - (v1 * (float64(-2.25) / (((1e10 * v1) - (v1 / v1)) * v1)))) - 1.5)))
                        var v59 string = "hello"
                        _ = fi57
                        _ = v58
                        _ = v59
                    }
                }
            }
        default:
            if ((127 > 0) && ((!(176 <= 255)) && (!(255 < 26)))) {
                var v60 uint32 = 2147483648
                switch (("\n\t" + "\n\t") + "") {
                case "\n\t":
                    v61 := v3[(v60 ^ (v60 >> ((1) & 63)))]
                    _ = v61
                case "":
                    fmt.Println(v1, v1, v1)
                case "foo bar":
                    fmt.Println(v1, v60, v1)
                }
                _ = v60
            }
        }
        _ = fi47
    }
    var v62 uint16 = 1
    var v63 uint16 = (v62 * (v62 & v62))
    for fi64 := 0; fi64 < 7; fi64++ {
        fmt.Println(v1, fi64, fi64)
        _ = fi64
    }
    v62 += (v63 | v63)
    v62 = v63
    switch 0 {
    case 4294967295:
        switch 1 {
        case 42:
            if (!(0.1 < (1.5 * 0.1))) {
                for fi65 := 0; fi65 < 3; fi65++ {
                    v66 := []int{(fi65 | -453), fi65}
                    v62 = v63
                    var v67 float32 = (float32((v1 + v1)) / -2.25)
                    _ = fi65
                    _ = v66
                    _ = v67
                }
                v62 -= v63
                for fi68 := 0; fi68 < 17; fi68++ {
                    v69 := GenS3{F0: 2147483647, F1: 0}
                    v62 = v63
                    _ = fi68
                    _ = v69
                }
            }
        case -9223372036854775808:
            if ((0.1 * 0.1) >= (0.1 - 3.14159)) {
                v3[0] = "foo bar"
                var v70 uint64 = 2
                _ = v70
            }
        }
    }
    switch 32768 {
    case 42:
        if (!false) {
            for fi71 := 0; fi71 < 9; fi71++ {
                if true {
                    v72 := []int{}
                    v2["\n\t"] = (65535 <= (v62 * v62))
                    _ = v72
                }
                var v73 uint64 = 18446744073709551615
                _ = fi71
                _ = v73
            }
            if true {
                var v74 uint16 = (v63 % ((255) | 1))
                switch "" {
                case "hello":
                    for fi75 := 0; fi75 < 16; fi75++ {
                        if false {
                            if false {
                                if false {
                                    switch (fi75 | fi75) {
                                    case 32767:
                                        for fi76 := 0; fi76 < 8; fi76++ {
                                            v63 = 0
                                            v77 := GenS1{F0: "hello"}
                                            if (!(65535 != 730)) {
                                                v78 := []int{}
                                                switch ("a" + ("hello" + (("\n\t" + "\n\t") + "hello"))) {
                                                case "foo bar":
                                                    if false {
                                                        fmt.Println(fi76, fi76)
                                                        v62 = v74
                                                        fmt.Println(fi75)
                                                    } else {
                                                        fmt.Println(v62)
                                                    }
                                                case "\n\t":
                                                    v79 := GenS4{F0: 1, F1: 0}
                                                    _ = v79
                                                }
                                                var v80 string = ("a" + ("foo bar" + ("" + "\n\t")))
                                                _ = v78
                                                _ = v80
                                            } else {
                                                v63 += (v74 ^ v74)
                                                v81 := []int{fi75, (fi76 << ((4294967295) & 63)), 42}
                                                _ = v81
                                            }
                                            _ = fi76
                                            _ = v77
                                        }
                                    case 1:
                                        v1 += v1
                                    case -1:
                                        var v82 float64 = -2.25
                                        _ = v82
                                    }
                                    v62 += (^(v63))
                                    delete(v3, 747)
                                }
                                var v83 float64 = 0.0
                                _ = len(v3)
                                _ = v83
                            } else {
                                v84 := GenS4{F0: -1, F1: 1}
                                if (42 == 572) {
                                    var v85 bool = ((("foo bar" + ("foo bar" + "")) + "a") > "")
                                    v63 += (^(v74))
                                    _ = v85
                                } else {
                                    v63 -= (v63 << ((18446744073709551615) & 63))
                                }
                                _ = len(v2)
                                _ = v84
                            }
                            fmt.Println(v1, v1, v63)
                            if true {
                                v2[("" + (("foo bar" + "foo bar") + ("\n\t" + "")))] = (v62 < 2)
                                for fi86 := 0; fi86 < 3; fi86++ {
                                    switch -37 {
                                    case 42:
                                        v87 := GenS2{F0: 1}
                                        _ = v87
                                    default:
                                        if ((32767 > 32767) || (!(!false))) {
                                            if false {
                                                var v88 float32 = float32(float64(1e10))
                                                _ = v88
                                            } else {
                                                if ((1 >= 0) && (fi86 >= (fi75 - fi75))) {
                                                    v89 := GenS3{F0: 2, F1: 42}
                                                    _ = v89
                                                }
                                                _ = len(v3)
                                                v1 += ((float64(((float32((v1 / v1)) + 1e10) * 0.0)) * 0.0) / v1)
                                                v62 = v62
                                            }
                                        } else {
                                            switch "" {
                                            case "foo bar":
                                                v74 += v63
                                            case "":
                                                if (65535 <= 2) {
                                                    var v90 int = 1
                                                    var v91 int8 = -36
                                                    _ = v90
                                                    _ = v91
                                                } else {
                                                    v62 = v63
                                                    var v92 bool = false
                                                    v74 += (^(v63))
                                                    _ = v92
                                                }
                                            }
                                        }
                                    }
                                    var v93 uint16 = 283
                                    var v94 uint16 = (^(v62))
                                    v95 := GenS3{F0: 2147483647, F1: 1}
                                    v96 := GenS3{F0: 2147483648, F1: 9223372036854775807}
                                    _ = fi86
                                    _ = v93
                                    _ = v94
                                    _ = v95
                                    _ = v96
                                }
                            }
                        } else {
                            if true {
                                if (!true) {
                                    var v97 uint64 = 2
                                    var v98 uint16 = 603
                                    _ = v97
                                    _ = v98
                                }
                                var v99 int32 = -575
                                if (!(4294967295 > fi75)) {
                                    v62 -= v62
                                    if false {
                                        v1 += (v1 / -2.25)
                                    } else {
                                        v100 := GenS1{F0: ""}
                                        var v101 int8 = -59
                                        _ = v100
                                        _ = v101
                                    }
                                    fmt.Println(fi75, v74, v1)
                                } else {
                                    if (fi75 > (-(fi75))) {
                                        switch ("a" + "hello") {
                                        case "":
                                            v99 = -2147483648
                                        default:
                                            v74 = (v74 % ((v74) | 1))
                                        }
                                        delete(v3, 4294967295)
                                        v63 -= 42
                                    }
                                    for fi102 := 0; fi102 < 5; fi102++ {
                                        if false {
                                            var v103 float64 = -2.25
                                            _ = v103
                                        }
                                        var v104 int = 127
                                        _ = fi102
                                        _ = v104
                                    }
                                    var v105 uint16 = (v74 & (v62 / (((v62 + v74)) | 1)))
                                    _ = v105
                                }
                                _ = v99
                            }
                            var v106 int32 = -1
                            if true {
                                switch (fi75 * (fi75 % ((fi75) | 1))) {
                                case 1:
                                    v63 += (^(v62))
                                case 127:
                                    v107 := GenS4{F0: -1, F1: -1}
                                    _ = v107
                                case -104:
                                    var v108 int8 = 127
                                    _ = v108
                                }
                                v109 := v2[""]
                                _ = v109
                            } else {
                                v110 := GenS3{F0: 2147483647, F1: 0}
                                v111 := []int{}
                                if len(v111) >= 2 {
                                    v111 = v111[1:len(v111)]
                                }
                                _ = v110
                                _ = v111
                            }
                            _ = v106
                        }
                        v112 := GenS5{F0: 827, F1: 0}
                        _ = fi75
                        _ = v112
                    }
                default:
                    v113 := GenS2{F0: 0}
                    _ = v113
                }
                v114 := []int{42}
                _ = v114
                _ = v74
            } else {
                v63 += (v63 * 1)
                v115 := GenS4{F0: 9223372036854775807, F1: 1}
                _ = v115
            }
        } else {
            v62 += 65535
            var v116 uint16 = (^(v63))
            switch "\n\t" {
            case "":
                v117 := GenS5{F0: (v116 ^ (v63 * v116)), F1: 2}
                _ = v117
            }
            _ = v116
        }
    case 128:
        v118 := GenS5{F0: (v63 * 369), F1: 102}
        _ = v118
    case -545:
        var v119 int8 = -5
        _ = v119
    default:
        var v120 bool = true
        _ = v120
    }
    v121 := GenS6{F0: 42, F1: 42, F2: 366}
    var v122 uint16 = v63
    var v123 int32 = 1
    fmt.Println(v123)
    fmt.Println("v1:", v1)
    fmt.Println("v122:", v122)
    fmt.Println("v123:", v123)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v62:", v62)
    fmt.Println("v63:", v63)
    _ = v1
    _ = v121
    _ = v122
    _ = v123
    _ = v2
    _ = v3
    _ = v4
    _ = v62
    _ = v63
}
