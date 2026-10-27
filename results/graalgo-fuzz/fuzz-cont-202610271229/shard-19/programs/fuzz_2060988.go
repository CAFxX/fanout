package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 int8 }
type GenS2 struct { F0 uint16; F1 uint }
type GenS3 struct { F0 uint; F1 uint16; F2 int16 }
type GenS4 struct { F0 uint32; F1 bool }
type GenS5 struct { F0 int8; F1 string; F2 bool }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int8 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 uint16; F1 uint16 }
type GenS10 struct { F0 string; F1 int32; F2 uint64 }
type GenS11 struct { F0 int; F1 int32 }
func genFn0() uint16 {
    for fi6 := 0; fi6 < 16; fi6++ {
        var v7 float64 = ((1.5 - -2.25) + 1.5)
        v8 := []int{(fi6 & (fi6 * fi6)), 9223372036854775807, 127}
        _ = fi6
        _ = v7
        _ = v8
    }
    var v9 float32 = ((-2.25 / 1e10) / float32((1e10 + 1e10)))
    _ = v9
    return 611
}
func genFn1(p0 int64, p1 int64) int8 {
    p0 += (^(p1))
    for fi10 := 0; fi10 < 11; fi10++ {
        for fi11 := 0; fi11 < 19; fi11++ {
            var v12 uint64 = 886
            _ = fi11
            _ = v12
        }
        var v13 int8 = -90
        var v14 uint8 = 41
        _ = fi10
        _ = v13
        _ = v14
    }
    v15 := []int{735, 127, 9223372036854775807, -22}
    _ = p0
    _ = p1
    _ = v15
    return 3
}

func main() {
    var v1 int = 42
    v2 := []int{v1, v1, (v1 / ((v1) | 1)), (^(v1))}
    v3 := make(map[string]int)
    var v4 int16 = 1
    var v5 uint8 = 2
    v5 = (v5 - (v5 % (((v5 + 1)) | 1)))
    if (true && true) {
        switch "a" {
        case "foo bar":
            v5 -= (v5 + v5)
        case "\n\t":
            v5 = (v5 + v5)
        default:
            for fi16 := 0; fi16 < 11; fi16++ {
                if (false && (!(332 < 212))) {
                    for fi17 := 0; fi17 < 1; fi17++ {
                        fmt.Println(v4, fi17)
                        _ = len(v2)
                        _ = fi17
                    }
                    v5 += 218
                } else {
                    fmt.Println(fi16)
                    v1 = v1
                    var v18 int16 = v4
                    _ = v18
                }
                switch "\n\t" {
                case "hello":
                    _ = len(v3)
                }
                v2 = append(v2, -1)
                _ = fi16
            }
        }
        _ = len(v2)
    }
    v19 := v3[(("hello" + "a") + ("" + ""))]
    v20 := GenS1{F0: (^(v4)), F1: 1, F2: 100}
    if false {
        for fi21 := 0; fi21 < 3; fi21++ {
            var v22 uint64 = 1
            _ = fi21
            _ = v22
        }
        v23 := GenS2{F0: 837, F1: 0}
        _ = v23
    } else {
        v24 := GenS3{F0: 499, F1: genFn0(), F2: 32767}
        for fi25 := 0; fi25 < 12; fi25++ {
            var v26 int16 = int16(v1)
            _ = fi25
            _ = v26
        }
        for fi27 := 0; fi27 < 7; fi27++ {
            if true {
                if (-1 < int8(v1)) {
                    v4 -= v4
                    if false {
                        v1 += -9223372036854775808
                        if len(v2) > 0 {
                            v2[0] = 42
                        }
                        if (0.1 > 1e10) {
                            v28 := GenS3{F0: 1, F1: 1, F2: (v4 | int16(fi27))}
                            for fi29 := 0; fi29 < 12; fi29++ {
                                var v30 uint16 = 2
                                var v31 int16 = v4
                                _ = fi29
                                _ = v30
                                _ = v31
                            }
                            _ = v28
                        }
                    }
                } else {
                    var v32 string = "foo bar"
                    for fi33 := 0; fi33 < 13; fi33++ {
                        switch (-(fi33)) {
                        case 32768:
                            var v34 uint16 = 2
                            _ = v34
                        case 178:
                            switch "" {
                            case "hello":
                                v35 := GenS4{F0: 305, F1: ((255 == 0) || (-172 > -299))}
                                _ = v35
                            case "":
                                v5 = 127
                                v19 += fi33
                                var v36 bool = (1e10 == 0.0)
                                fmt.Println(v19, v1, fi27)
                                if (false || true) {
                                    if v36 {
                                        v37 := GenS5{F0: 127, F1: ("hello" + (v32 + ("" + ((((v32 + v32) + "a") + v32) + "a")))), F2: (!v36)}
                                        if v36 {
                                            var v38 int16 = (v4 + 0)
                                            if len(v2) > 0 {
                                                v2[0] = (v1 >> ((610) & 63))
                                            }
                                            _ = v38
                                        }
                                        var v39 float64 = (-2.25 - 3.14159)
                                        _ = v37
                                        _ = v39
                                    }
                                    v4 = v4
                                    var v40 uint16 = 42
                                    _ = v40
                                }
                                _ = v36
                            default:
                                v5 -= v5
                            }
                        }
                        _ = fi33
                    }
                    _ = v32
                }
                for fi41 := 0; fi41 < 20; fi41++ {
                    var v42 float64 = ((0.1 * float64(1.5)) - (0.1 / (0.0 + (float64((((1.5 - 1.5) + 0.0) - -2.25)) - float64(((0.0 / 0.1) * 0.0))))))
                    switch (v19 / ((v19) | 1)) {
                    case 801:
                        v1 -= (v1 % ((v19) | 1))
                        var v43 int16 = (v4 + (^(v4)))
                        v1 -= fi41
                        v5 += v5
                        v19 += (fi27 & (fi27 & int(v5)))
                        if false {
                            if false {
                                var v44 int64 = 42
                                _ = v44
                            }
                            switch (v1 ^ 355) {
                            case 9223372036854775807:
                                if len(v2) > 0 {
                                    v2[0] = -1
                                }
                            case 127:
                                var v45 int16 = 95
                                _ = v45
                            case 32768:
                                var v46 uint16 = 2
                                _ = v46
                            default:
                                v2 = append(v2, 32767)
                            }
                            var v47 float64 = (v42 - 0.1)
                            _ = v47
                        }
                        _ = v43
                    }
                    var v48 uint64 = 101
                    _ = fi41
                    _ = v42
                    _ = v48
                }
                fmt.Println(v1)
            }
            v1 = -9223372036854775808
            if (!((-112 > -9) && true)) {
                if true {
                    for fi49 := 0; fi49 < 17; fi49++ {
                        var v50 int16 = (v4 ^ (v4 ^ v4))
                        _ = fi49
                        _ = v50
                    }
                    v1 -= 229
                }
            } else {
                for fi51 := 0; fi51 < 20; fi51++ {
                    var v52 int8 = -123
                    _ = fi51
                    _ = v52
                }
                for fi53 := 0; fi53 < 9; fi53++ {
                    v19 = fi27
                    fmt.Println(fi53, v4)
                    fmt.Println(v4, fi53)
                    _ = fi53
                }
            }
            _ = fi27
        }
        _ = v24
    }
    v2 = append(v2, v19)
    fmt.Println(v4)
    if (true || ((127 > (v5 & 110)) || false)) {
        if ((true && (18446744073709551615 < 0)) && (42 <= 1)) {
            v54 := GenS4{F0: 2, F1: ((true && true) || false)}
            _ = len(v2)
            v55 := GenS1{F0: (v4 % (((v4 | (v4 % ((v4) | 1)))) | 1)), F1: 255, F2: genFn1(42, -1)}
            _ = v54
            _ = v55
        }
        switch (("" + "hello") + "\n\t") {
        case "\n\t":
            for fi56 := 0; fi56 < 3; fi56++ {
                var v57 int16 = (v4 >> ((4294967295) & 63))
                v4 += v57
                _ = fi56
                _ = v57
            }
        case "foo bar":
            v4 = v4
        case "a":
            fmt.Println(v19)
        default:
            v58 := GenS3{F0: 239, F1: 255, F2: (v4 + v4)}
            _ = v58
        }
        switch v1 {
        case -416:
            v5 = (v5 / ((2) | 1))
        }
    } else {
        for fi59 := 0; fi59 < 14; fi59++ {
            var v60 int64 = 1
            _ = fi59
            _ = v60
        }
    }
    switch (v19 % ((v19) | 1)) {
    case 127:
        if true {
            for fi61 := 0; fi61 < 10; fi61++ {
                v62 := GenS6{F0: -1}
                v5 -= v5
                _ = fi61
                _ = v62
            }
            if (!(int8(v1) != -1)) {
                if true {
                    v63 := v3[""]
                    fmt.Println(v19)
                    var v64 uint32 = 2147483647
                    _ = v63
                    _ = v64
                }
            }
            if len(v2) > 0 {
                v2[0] = -1
            }
        } else {
            switch "a" {
            case "foo bar":
                v1 -= v1
            case "":
                var v65 int16 = 1000
                _ = v65
            default:
                if (2 >= 2147483648) {
                    var v66 int = v19
                    _ = v66
                } else {
                    switch (v19 ^ 42) {
                    case 0:
                        v1 = (v19 % ((v19) | 1))
                    case 42:
                        for fi67 := 0; fi67 < 14; fi67++ {
                            for fi68 := 0; fi68 < 3; fi68++ {
                                switch "\n\t" {
                                case "hello":
                                    v19 += 42
                                }
                                switch "a" {
                                case "":
                                    if false {
                                        switch -2147483648 {
                                        case 2147483647:
                                            var v69 int32 = -531
                                            _ = v69
                                        case -1:
                                            v19 -= v19
                                        case 934:
                                            v1 = (fi67 + v19)
                                        default:
                                            if (!(false && (float32(float64((0.1 / 1.0))) >= float32(3.14159)))) {
                                                if true {
                                                    fmt.Println(fi67, v19)
                                                    v70 := GenS5{F0: -1, F1: "hello", F2: false}
                                                    _ = v70
                                                } else {
                                                    v4 = v4
                                                    var v71 uint64 = 54
                                                    _ = v71
                                                }
                                                if ("a" <= "") {
                                                    v72 := GenS7{F0: genFn1(293, 692)}
                                                    _ = v72
                                                } else {
                                                    var v73 string = "hello"
                                                    var v74 bool = false
                                                    _ = v73
                                                    _ = v74
                                                }
                                                switch fi68 {
                                                case 42:
                                                    v75 := GenS6{F0: 587}
                                                    _ = v75
                                                default:
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                }
                                            }
                                        }
                                        v76 := GenS5{F0: -126, F1: ("\n\t" + "foo bar"), F2: true}
                                        _ = v76
                                    }
                                case "foo bar":
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                }
                                _ = fi68
                            }
                            v4 += (v4 << ((281) & 63))
                            _ = fi67
                        }
                    case -2147483648:
                        for fi77 := 0; fi77 < 5; fi77++ {
                            v1 = v1
                            v78 := GenS1{F0: v4, F1: 705, F2: 83}
                            v5 = 0
                            _ = fi77
                            _ = v78
                        }
                    }
                    fmt.Println(v19, v4, v19)
                }
            }
            switch v1 {
            case 156:
                if (float32(3.14159) != ((float32(((1.5 / (0.1 + -2.25)) - float64(1e10))) / 1.0) * 0.0)) {
                    v5 -= (v5 / ((v5) | 1))
                }
            case 189:
                var v79 uint64 = 0
                _ = v79
            default:
                v19 = (v19 * v19)
            }
        }
    case -369:
        fmt.Println(v1)
    }
    switch "foo bar" {
    case "":
        var v80 int8 = 42
        _ = v80
    case "\n\t":
        if true {
            if false {
                fmt.Println(v1, v5)
                if false {
                    var v81 float32 = float32(float64(1e10))
                    _ = v81
                } else {
                    fmt.Println(v19, v5, v4)
                    var v82 uint16 = 42
                    _ = v82
                }
                for fi83 := 0; fi83 < 3; fi83++ {
                    var v84 uint = 4294967295
                    _ = fi83
                    _ = v84
                }
            }
        } else {
            var v85 int8 = 1
            switch -196 {
            case -599:
                v4 += 32767
            }
            fmt.Println(v19, v5, v1)
            _ = v85
        }
    default:
        for fi86 := 0; fi86 < 12; fi86++ {
            delete(v3, "hello")
            v87 := GenS8{F0: 832}
            _ = fi86
            _ = v87
        }
    }
    if (!true) {
        for fi88 := 0; fi88 < 20; fi88++ {
            if true {
                var v89 int16 = v4
                _ = v89
            } else {
                for fi90 := 0; fi90 < 5; fi90++ {
                    v91 := GenS9{F0: 0, F1: 65535}
                    var v92 float64 = (-2.25 * -2.25)
                    _ = fi90
                    _ = v91
                    _ = v92
                }
                for fi93 := 0; fi93 < 15; fi93++ {
                    var v94 int32 = 0
                    _ = fi93
                    _ = v94
                }
            }
            if true {
                v1 = (fi88 << ((18446744073709551615) & 63))
                for fi95 := 0; fi95 < 4; fi95++ {
                    v5 = v5
                    _ = fi95
                }
                var v96 int = fi88
                _ = v96
            }
            _ = fi88
        }
    } else {
        var v97 uint = 42
        var v98 int64 = 9223372036854775807
        var v99 int = v19
        _ = v97
        _ = v98
        _ = v99
    }
    if true {
        v100 := GenS10{F0: "hello", F1: 0, F2: 948}
        if true {
            for fi101 := 0; fi101 < 5; fi101++ {
                var v102 int16 = v4
                _ = fi101
                _ = v102
            }
            v19 = (v1 << ((0) & 63))
        }
        _ = v100
    } else {
        v103 := GenS4{F0: 1, F1: true}
        if false {
            v1 = v19
            var v104 int8 = 0
            _ = v104
        } else {
            v4 -= 550
            for fi105 := 0; fi105 < 2; fi105++ {
                var v106 int8 = 100
                v3["hello"] = 127
                _ = fi105
                _ = v106
            }
            fmt.Println(v5, v5)
        }
        switch (v19 | (v19 % ((v19) | 1))) {
        case 1:
            var v107 uint32 = 1
            _ = v107
        case 676:
            for fi108 := 0; fi108 < 14; fi108++ {
                v1 = 352
                switch 992 {
                case 21:
                    for fi109 := 0; fi109 < 5; fi109++ {
                        v5 = v5
                        var v110 bool = true
                        var v111 int16 = (^(v4))
                        _ = fi109
                        _ = v110
                        _ = v111
                    }
                case 128:
                    v5 = (v5 % (((v5 * (v5 - v5))) | 1))
                case -926:
                    var v112 int64 = 9223372036854775807
                    _ = v112
                default:
                    switch (("" + "\n\t") + ("" + "hello")) {
                    case "a":
                        v1 -= (fi108 * fi108)
                    case "":
                        v2 = append(v2, v19)
                    case "hello":
                        if (!true) {
                            for fi113 := 0; fi113 < 3; fi113++ {
                                fmt.Println(v5, fi113)
                                fmt.Println(fi113, v4, v5)
                                var v114 int16 = 148
                                _ = fi113
                                _ = v114
                            }
                            v115 := GenS11{F0: v1, F1: 295}
                            v3["hello"] = (fi108 | (v1 << ((74) & 63)))
                            _ = v115
                        } else {
                            switch "" {
                            case "":
                                v4 = 750
                            case "foo bar":
                                for fi116 := 0; fi116 < 13; fi116++ {
                                    v117 := v3["foo bar"]
                                    if false {
                                        var v118 string = ("hello" + ("" + ""))
                                        var v119 uint64 = 18446744073709551615
                                        _ = v118
                                        _ = v119
                                    }
                                    var v120 int64 = 1
                                    _ = fi116
                                    _ = v117
                                    _ = v120
                                }
                            case "a":
                                v4 += -32768
                            }
                            if true {
                                var v121 int64 = -1
                                _ = v121
                            }
                            var v122 uint = 255
                            _ = v122
                        }
                    }
                }
                _ = fi108
            }
        }
        _ = v103
    }
    fmt.Println(v5, v5)
    var v123 uint = 255
    v124 := v3["a"]
    fmt.Println("v1:", v1)
    fmt.Println("v123:", v123)
    fmt.Println("v124:", v124)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v123
    _ = v124
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
