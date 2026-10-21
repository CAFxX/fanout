package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 uint8 }
type GenS2 struct { F0 int32; F1 string }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 uint32; F1 int64; F2 bool }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 bool; F1 uint16; F2 bool }
type GenS7 struct { F0 uint8; F1 int16; F2 int }
type GenS8 struct { F0 int64; F1 bool }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int64 = -1
    var v3 int32 = -1
    v3 = v3
    if true {
        if (true && (("" + "") >= ("" + "a"))) {
            v4 := []int{4294967295, -462, 2147483647}
            _ = v4
        } else {
            switch "a" {
            case "foo bar":
                if false {
                    var v5 int16 = 0
                    var v6 int = -278
                    v3 = (v3 ^ int32(v5))
                    _ = v5
                    _ = v6
                }
            }
        }
    }
    v2 = (^(v2))
    switch 0 {
    case 0:
        switch 1 {
        case 32768:
            v7 := []int{127, 32768, -741, 42}
            _ = v7
        case 128:
            v8 := GenS1{F0: 2, F1: -231, F2: 128}
            _ = v8
        case 210:
            if (false && (!(4294967295 != 4294967295))) {
                v9 := GenS2{F0: (v3 & (v3 << ((4294967295) & 63))), F1: "hello"}
                v10 := []int{}
                _ = v10
                _ = v9
            } else {
                v11 := GenS1{F0: 42, F1: 4294967295, F2: 1}
                if false {
                    v12 := []int{-2147483648, 17, 1, 2147483647}
                    for fi13 := 0; fi13 < 1; fi13++ {
                        for fi14 := 0; fi14 < 2; fi14++ {
                            if (242 >= 100) {
                                var v15 uint16 = 107
                                v3 = v3
                                var v16 string = ""
                                _ = v15
                                _ = v16
                            }
                            _ = fi14
                        }
                        for fi17 := 0; fi17 < 6; fi17++ {
                            v3 = (v3 << ((873) & 63))
                            _ = fi17
                        }
                        v12 = append(v12, fi13)
                        _ = fi13
                    }
                    for fi18 := 0; fi18 < 12; fi18++ {
                        v19 := GenS2{F0: 42, F1: "a"}
                        v20 := GenS1{F0: 1, F1: fi18, F2: 127}
                        _ = fi18
                        _ = v19
                        _ = v20
                    }
                    _ = v12
                }
                v21 := []int{}
                _ = v11
                _ = v21
            }
        default:
            for fi22 := 0; fi22 < 8; fi22++ {
                v2 += 1
                v23 := []int{(fi22 ^ (fi22 >> ((29) & 63)))}
                _ = fi22
                _ = v23
            }
        }
    default:
        _ = len(v1)
    }
    for fi24 := 0; fi24 < 13; fi24++ {
        switch -465 {
        case 188:
            if (1 > 1) {
                switch (fi24 * (fi24 << ((401) & 63))) {
                case 42:
                    var v25 float64 = (1e10 * float64(0.1))
                    _ = v25
                default:
                    if true {
                        _ = len(v1)
                        v2 -= -1
                        v26 := GenS2{F0: -1, F1: ""}
                        _ = v26
                    } else {
                        var v27 uint = 1
                        _ = v27
                    }
                }
            } else {
                v28 := GenS2{F0: 42, F1: "a"}
                var v29 float32 = 1.5
                switch int(v2) {
                case -821:
                    var v30 float32 = v29
                    _ = v30
                case 32767:
                    v3 -= -998
                case 42:
                    fmt.Println(v29, v2, v2)
                }
                _ = v28
                _ = v29
            }
        case 42:
            v2 = 1
        case 0:
            _ = len(v1)
        default:
            v31 := GenS2{F0: v3, F1: ("a" + "")}
            _ = v31
        }
        var v32 int = fi24
        _ = fi24
        _ = v32
    }
    v33 := GenS3{F0: 42}
    if false {
        v34 := GenS4{F0: 1, F1: -966, F2: (!true)}
        fmt.Println(v3)
        _ = v34
    } else {
        v2 -= v2
        v35 := GenS2{F0: v3, F1: (("hello" + "foo bar") + "hello")}
        _ = v35
    }
    fmt.Println(v3, v2)
    fmt.Println(v3, v2)
    fmt.Println(v3)
    v3 = v3
    for fi36 := 0; fi36 < 20; fi36++ {
        if (false && false) {
            v37 := GenS3{F0: 0}
            switch (fi36 | fi36) {
            case 4294967295:
                delete(v1, 981)
            case 128:
                if true {
                    switch 1 {
                    case 32768:
                        fmt.Println(fi36, v3)
                    case 9223372036854775807:
                        v2 = v2
                    case 706:
                        for fi38 := 0; fi38 < 4; fi38++ {
                            v2 += v2
                            v3 = (v3 >> ((18446744073709551615) & 63))
                            v39 := GenS3{F0: 852}
                            _ = fi38
                            _ = v39
                        }
                    default:
                        v40 := GenS4{F0: 65535, F1: (v2 % ((v2) | 1)), F2: (v2 > (v2 >> ((18446744073709551615) & 63)))}
                        _ = v40
                    }
                    v3 += (v3 % ((v3) | 1))
                }
            case -2147483648:
                v41 := []int{32767, 2147483647}
                _ = v41
            default:
                if false {
                    var v42 int64 = v2
                    _ = v42
                }
            }
            _ = v37
        }
        var v43 bool = false
        _ = fi36
        _ = v43
    }
    for fi44 := 0; fi44 < 2; fi44++ {
        var v45 uint32 = 2147483648
        for fi46 := 0; fi46 < 3; fi46++ {
            delete(v1, 42)
            if (!((false || (v2 == (v2 >> ((42) & 63)))) || true)) {
                v47 := GenS1{F0: 396, F1: (fi46 % ((fi44) | 1)), F2: 0}
                _ = v47
            } else {
                if true {
                    v48 := v1[(v45 + 134)]
                    _ = v48
                } else {
                    v49 := GenS3{F0: 2}
                    _ = v49
                }
            }
            v3 = (v3 - v3)
            var v50 bool = (100 <= 44)
            var v51 uint = 310
            _ = fi46
            _ = v50
            _ = v51
        }
        v3 -= v3
        _ = fi44
        _ = v45
    }
    var v52 string = "hello"
    var v53 uint32 = 2147483648
    if ((^(v3)) != v3) {
        fmt.Println(v53)
        switch 32768 {
        case 2147483647:
            v54 := GenS3{F0: 678}
            _ = v54
        case 0:
            var v55 int8 = 0
            _ = v55
        case 32768:
            for fi56 := 0; fi56 < 14; fi56++ {
                if (!(true && true)) {
                    for fi57 := 0; fi57 < 4; fi57++ {
                        v3 = v3
                        v3 -= (v3 / ((1) | 1))
                        _ = fi57
                    }
                    var v58 int16 = 1
                    _ = v58
                }
                if false {
                    v59 := GenS5{F0: -1}
                    v2 = -1
                    _ = v59
                } else {
                    delete(v1, v53)
                    switch (fi56 / ((9223372036854775807) | 1)) {
                    case -9223372036854775808:
                        if (false && false) {
                            fmt.Println(v53, v52, fi56)
                            v2 -= (v2 - -1)
                            v60 := []int{int(v3), (fi56 | 163), (fi56 | -2147483648)}
                            _ = v60
                        }
                    default:
                        fmt.Println(v3, v53, v52)
                    }
                    var v61 uint = 0
                    _ = v61
                }
                fmt.Println(v3)
                _ = fi56
            }
        default:
            v62 := []int{}
            _ = v62
        }
    } else {
        v52 = "foo bar"
        v3 -= (v3 >> ((190) & 63))
        var v63 uint32 = 2
        _ = v63
    }
    for fi64 := 0; fi64 < 16; fi64++ {
        if (!(2 != 2)) {
            if false {
                var v65 int = fi64
                v66 := GenS2{F0: 1, F1: v52}
                _ = v65
                _ = v66
            }
        } else {
            switch "hello" {
            case "":
                for fi67 := 0; fi67 < 15; fi67++ {
                    switch fi64 {
                    case -63:
                        switch (fi64 | fi67) {
                        case 42:
                            v68 := v1[v53]
                            _ = v68
                        case -2147483648:
                            switch v52 {
                            case "\n\t":
                                v69 := GenS4{F0: (v53 * 2147483648), F1: v2, F2: (42 <= -24)}
                                _ = v69
                            }
                        case 32768:
                            v70 := []int{fi67, fi67, (^(fi67)), -891}
                            _ = v70
                        default:
                            fmt.Println(v53)
                        }
                    }
                    var v71 bool = true
                    fmt.Println(v3)
                    _ = fi67
                    _ = v71
                }
            case "foo bar":
                v72 := GenS6{F0: (!true), F1: 1, F2: (!(false && (true && ((-2.25 - 1.5) <= 1.5))))}
                _ = v72
            default:
                if ((4294967295 > 4294967295) || (128 > 0)) {
                    for fi73 := 0; fi73 < 19; fi73++ {
                        v74 := []int{(fi64 << ((7) & 63)), fi64}
                        if true {
                            var v75 bool = (!false)
                            v52 = ("hello" + v52)
                            if (!(3.14159 <= (float32(1e10) / 1.5))) {
                                switch int(v2) {
                                case 42:
                                    var v76 int = (fi64 * fi73)
                                    _ = v76
                                case 907:
                                    var v77 int8 = -95
                                    _ = v77
                                default:
                                    v74 = append(v74, fi73)
                                }
                                v78 := GenS3{F0: 1}
                                _ = len(v74)
                                _ = v78
                            } else {
                                switch "foo bar" {
                                case "hello":
                                    var v79 int = -988
                                    _ = v79
                                case "a":
                                    if v75 {
                                        var v80 int16 = -1
                                        fmt.Println(v53)
                                        var v81 int = int(v53)
                                        v82 := GenS6{F0: (v75 && ((((127 > 127) || v75) || v75) || ((v2 << ((18446744073709551615) & 63)) > v2))), F1: 502, F2: v75}
                                        _ = v80
                                        _ = v81
                                        _ = v82
                                    } else {
                                        if len(v74) >= 2 {
                                            v74 = v74[1:len(v74)]
                                        }
                                    }
                                case "\n\t":
                                    v53 += 2
                                }
                                var v83 string = ("foo bar" + (("foo bar" + v52) + v52))
                                var v84 int32 = 960
                                _ = v83
                                _ = v84
                            }
                            switch "hello" {
                            case "a":
                                switch (-(fi73)) {
                                case -9223372036854775808:
                                    v3 -= (v3 % (((^(v3))) | 1))
                                case -783:
                                    v75 = v75
                                default:
                                    v3 += 539
                                }
                            case "\n\t":
                                if (!((0 < 128) || (160 < 255))) {
                                    if v75 {
                                        if true {
                                            var v85 float64 = (3.14159 - (1e10 * 1.5))
                                            if (42 == 127) {
                                                var v86 float32 = 1.5
                                                if v75 {
                                                    var v87 uint32 = 1
                                                    if len(v74) >= 2 {
                                                        v74 = v74[1:len(v74)]
                                                    }
                                                    var v88 float32 = v86
                                                    _ = v87
                                                    _ = v88
                                                } else {
                                                    fmt.Println(v75, v53)
                                                    fmt.Println(fi73)
                                                    if (0 > 0) {
                                                        var v89 int16 = 902
                                                        _ = v89
                                                    }
                                                }
                                                var v90 int = fi73
                                                _ = v86
                                                _ = v90
                                            } else {
                                                v53 = v53
                                                fmt.Println(v75, v52)
                                                switch (("\n\t" + v52) + v52) {
                                                case "\n\t":
                                                    var v91 uint64 = 42
                                                    _ = v91
                                                case "hello":
                                                    var v92 uint8 = 127
                                                    _ = v92
                                                default:
                                                    if v75 {
                                                        var v93 int64 = -9223372036854775808
                                                        v93 = (v93 % (((v93 >> ((396) & 63))) | 1))
                                                        _ = v93
                                                    } else {
                                                        v52 = "\n\t"
                                                        v2 = (v2 / (((v2 / (((v2 & (v2 << ((2) & 63)))) | 1))) | 1))
                                                        v94 := GenS5{F0: v2}
                                                        _ = v94
                                                    }
                                                }
                                            }
                                            v75 = false
                                            v85 = 3.14159
                                            _ = v85
                                        }
                                        var v95 float64 = (((1.5 * float64(((float32(0.0) * 0.0) + (1.5 + (float32(((0.0 * -2.25) * (((0.1 / 1.5) - (1.5 * 3.14159)) / float64(1e10)))) * 0.0))))) / 1.0) - (float64((float32((3.14159 * -2.25)) + 0.1)) / 1e10))
                                        if (fi73 > (^(fi73))) {
                                            v95 = (float64(-2.25) * 1.5)
                                            if ((v2 % (((v2 / ((v2) | 1))) | 1)) >= v2) {
                                                v75 = (!(!v75))
                                                if ((v3 % ((v3) | 1)) > (-(v3))) {
                                                    var v96 int64 = v2
                                                    switch (("a" + v52) + "a") {
                                                    case "foo bar":
                                                        switch (((("hello" + v52) + ("\n\t" + v52)) + v52) + ("a" + v52)) {
                                                        case "\n\t":
                                                            if (0 <= 227) {
                                                                v97 := GenS4{F0: 344, F1: (v2 / (((v96 - 0)) | 1)), F2: (-38 == 127)}
                                                                _ = v97
                                                            }
                                                        default:
                                                            var v98 string = v52
                                                            _ = v98
                                                        }
                                                    case "\n\t":
                                                        var v99 uint8 = 127
                                                        _ = v99
                                                    case "":
                                                        if v75 {
                                                            v100 := GenS1{F0: 0, F1: (fi64 - 2147483647), F2: 255}
                                                            v101 := GenS7{F0: 255, F1: 32767, F2: fi73}
                                                            _ = v100
                                                            _ = v101
                                                        }
                                                    }
                                                    var v102 int16 = 1000
                                                    _ = v102
                                                    _ = v96
                                                }
                                                var v103 uint64 = 0
                                                var v104 string = "a"
                                                _ = v103
                                                _ = v104
                                            }
                                            var v105 int32 = 0
                                            v3 += (v3 >> ((0) & 63))
                                            _ = v105
                                        } else {
                                            var v106 uint8 = 128
                                            v3 += -157
                                            _ = v106
                                        }
                                        _ = v95
                                    }
                                } else {
                                    switch fi73 {
                                    case 9223372036854775807:
                                        v52 = ("hello" + ("hello" + "foo bar"))
                                        fmt.Println(fi64, v3, v3)
                                        v75 = v75
                                    case -495:
                                        var v107 float64 = -2.25
                                        _ = v107
                                    case 0:
                                        v2 -= (v2 >> ((18446744073709551615) & 63))
                                    }
                                }
                                v75 = (113 < 1)
                            default:
                                var v108 int16 = 0
                                _ = v108
                            }
                            _ = v75
                        }
                        var v109 uint8 = 2
                        _ = fi73
                        _ = v109
                        _ = v74
                    }
                    for fi110 := 0; fi110 < 6; fi110++ {
                        v111 := GenS7{F0: 1, F1: 42, F2: fi64}
                        if (0 < 127) {
                            if true {
                                var v112 string = ("hello" + ("" + "hello"))
                                var v113 uint32 = (^(v53))
                                var v114 int16 = 42
                                _ = v112
                                _ = v113
                                _ = v114
                            } else {
                                v115 := []int{}
                                fmt.Println(v2, v52, v53)
                                _ = v115
                            }
                            if ((32767 == 771) || (((v52 + v52) + v52) < "")) {
                                fmt.Println(v53, fi64)
                                v116 := []int{fi64, -954, fi110}
                                _ = v116
                            }
                        }
                        _ = len(v1)
                        _ = fi110
                        _ = v111
                    }
                } else {
                    var v117 float64 = (((1e10 - float64(0.1)) / float64(1e10)) * 1.5)
                    _ = v117
                }
            }
            var v118 uint16 = 676
            for fi119 := 0; fi119 < 3; fi119++ {
                fmt.Println(v2, v118)
                _ = fi119
            }
            _ = v118
        }
        fmt.Println(v53, v53)
        _ = fi64
    }
    v120 := GenS8{F0: v2, F1: (42 == 1)}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v120
    _ = v2
    _ = v3
    _ = v33
    _ = v52
    _ = v53
}
