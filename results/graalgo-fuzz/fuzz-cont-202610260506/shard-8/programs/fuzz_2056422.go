package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int8; F1 string }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 uint16; F1 string; F2 int32 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint32; F1 uint }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 bool; F1 uint8 }
type GenS10 struct { F0 uint16; F1 int64 }
type GenS11 struct { F0 string; F1 string }
type GenS12 struct { F0 string; F1 string }
type GenS13 struct { F0 uint }
type GenS14 struct { F0 uint; F1 int16; F2 int }
type GenS15 struct { F0 int32; F1 uint; F2 int32 }
type GenS16 struct { F0 uint; F1 int32 }
type GenS17 struct { F0 uint16; F1 uint }
type GenS18 struct { F0 uint32; F1 uint16; F2 uint }
type GenS19 struct { F0 uint16; F1 uint16; F2 int8 }
type GenS20 struct { F0 int16; F1 uint64; F2 uint16 }
type GenS21 struct { F0 bool }
type GenS22 struct { F0 int16 }
type GenS23 struct { F0 uint; F1 string; F2 uint32 }
type GenS24 struct { F0 int }
type GenS25 struct { F0 uint16 }
type GenS26 struct { F0 uint32; F1 uint8 }
type GenS27 struct { F0 bool; F1 int16 }
type GenS28 struct { F0 uint8; F1 uint64; F2 string }
type GenS29 struct { F0 int }
type GenS30 struct { F0 uint8 }

func main() {
    var v1 uint16 = 32767
    var v2 uint8 = 128
    var v3 uint32 = 23
    var v4 int8 = -128
    v5 := make(map[uint32]string)
    switch ("" + "a") {
    case "\n\t":
        if (!true) {
            if true {
                var v6 uint16 = uint16(v3)
                if false {
                    if ((1.5 / (((0.0 + (float64(1.5) * 0.0)) / 0.1) + (0.0 + float64((-2.25 - (1.5 * (3.14159 + 3.14159))))))) == 0.0) {
                        if true {
                            if (v4 != (v4 * v4)) {
                                v2 = v2
                            }
                            v7 := GenS1{F0: 0, F1: v2}
                            switch "foo bar" {
                            case "a":
                                if (!false) {
                                    v6 = 361
                                    v5[v3] = ("\n\t" + "a")
                                    if (42 == 859) {
                                        v1 -= v6
                                    } else {
                                        for fi8 := 0; fi8 < 7; fi8++ {
                                            if true {
                                                if ((v1 >> ((2) & 63)) > (^(v6))) {
                                                    var v9 string = ""
                                                    _ = v9
                                                }
                                                fmt.Println(fi8, v4, v2)
                                                switch ((("" + "\n\t") + "a") + "\n\t") {
                                                case "":
                                                    var v10 int = (fi8 - (fi8 % ((fi8) | 1)))
                                                    _ = v10
                                                case "\n\t":
                                                    if false {
                                                        fmt.Println(v4, v4)
                                                        switch "foo bar" {
                                                        case "":
                                                            for fi11 := 0; fi11 < 10; fi11++ {
                                                                switch "\n\t" {
                                                                case "hello":
                                                                    fmt.Println(v6)
                                                                case "a":
                                                                    if false {
                                                                        v6 += 700
                                                                        switch "\n\t" {
                                                                        case "\n\t":
                                                                            var v12 int8 = int8(v3)
                                                                            if true {
                                                                                switch 1 {
                                                                                case -912:
                                                                                    v13 := GenS2{F0: v4}
                                                                                    _ = v13
                                                                                case 32768:
                                                                                    var v14 int64 = 65
                                                                                    _ = v14
                                                                                case -2147483648:
                                                                                    var v15 uint32 = v3
                                                                                    _ = v15
                                                                                default:
                                                                                    var v16 uint16 = v1
                                                                                    v17 := GenS2{F0: (v4 | v4)}
                                                                                    switch "" {
                                                                                    case "\n\t":
                                                                                        v18 := GenS2{F0: -1}
                                                                                        _ = v18
                                                                                    case "":
                                                                                        v19 := []int{-596, (fi11 >> ((4294967295) & 63))}
                                                                                        _ = v19
                                                                                    }
                                                                                    var v20 uint8 = (v2 + v2)
                                                                                    switch fi11 {
                                                                                    case 42:
                                                                                        if (float64(1e10) >= float64((3.14159 / (1e10 * (((3.14159 * (0.0 + 1.5)) - 3.14159) - float32(-2.25)))))) {
                                                                                            if true {
                                                                                                var v21 uint32 = v3
                                                                                                _ = v21
                                                                                            }
                                                                                        } else {
                                                                                            fmt.Println(v1, v6, v12)
                                                                                        }
                                                                                    }
                                                                                    switch ("\n\t" + ("a" + "a")) {
                                                                                    case "":
                                                                                        switch (("\n\t" + "\n\t") + ("hello" + "\n\t")) {
                                                                                        case "hello":
                                                                                            var v22 uint16 = (v16 << ((uint(v1)) & 63))
                                                                                            _ = v22
                                                                                        case "":
                                                                                            v23 := GenS3{F0: v4, F1: ("a" + "a")}
                                                                                            _ = v23
                                                                                        }
                                                                                    case "hello":
                                                                                        v24 := []int{(fi8 - fi8), fi8, (-(fi8)), -1}
                                                                                        _ = v24
                                                                                    case "foo bar":
                                                                                        var v25 bool = (((!(!true)) || true) && (3.14159 <= 1e10))
                                                                                        _ = v25
                                                                                    }
                                                                                    if true {
                                                                                        var v26 uint32 = v3
                                                                                        _ = v26
                                                                                    } else {
                                                                                        if false {
                                                                                            if false {
                                                                                                var v27 string = ("foo bar" + ("a" + (((("foo bar" + ("foo bar" + "hello")) + "foo bar") + ("hello" + "a")) + (("" + "") + ("foo bar" + "hello")))))
                                                                                                var v28 int64 = -1
                                                                                                v2 = v20
                                                                                                if true {
                                                                                                    switch (fi8 | fi11) {
                                                                                                    case 0:
                                                                                                        v1 -= (v1 - 255)
                                                                                                    default:
                                                                                                        var v29 string = "a"
                                                                                                        _ = v29
                                                                                                    }
                                                                                                } else {
                                                                                                    v30 := []int{}
                                                                                                    var v31 int16 = -137
                                                                                                    fmt.Println(fi8, fi11, fi11)
                                                                                                    switch int(v6) {
                                                                                                    case 128:
                                                                                                        v32 := GenS2{F0: v12}
                                                                                                        _ = v32
                                                                                                    case 4294967295:
                                                                                                        switch "foo bar" {
                                                                                                        case "\n\t":
                                                                                                            switch v27 {
                                                                                                            case "hello":
                                                                                                                v33 := GenS4{F0: 64}
                                                                                                                _ = v33
                                                                                                            }
                                                                                                        case "hello":
                                                                                                            if len(v30) > 0 {
                                                                                                                v30[0] = (fi8 & 2147483647)
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                    v12 -= (-(v4))
                                                                                                    var v34 float32 = 0.0
                                                                                                    _ = v30
                                                                                                    _ = v31
                                                                                                    _ = v34
                                                                                                }
                                                                                                var v35 float64 = ((0.1 + (0.0 / (1.5 * -2.25))) + 0.1)
                                                                                                _ = v27
                                                                                                _ = v28
                                                                                                _ = v35
                                                                                            }
                                                                                            if (0 <= -42) {
                                                                                                _ = len(v5)
                                                                                                if (false && false) {
                                                                                                    var v36 uint32 = (v3 ^ v3)
                                                                                                    v6 -= 42
                                                                                                    if (!false) {
                                                                                                        var v37 float64 = (-2.25 * (1.5 - ((0.0 + (1e10 - float64(float32(((1e10 / 1.0) + 0.1))))) / 3.14159)))
                                                                                                        _ = v37
                                                                                                    } else {
                                                                                                        v12 = v4
                                                                                                        v38 := GenS5{F0: v1, F1: ("foo bar" + ""), F2: 311}
                                                                                                        var v39 string = "a"
                                                                                                        _ = v38
                                                                                                        _ = v39
                                                                                                    }
                                                                                                    _ = v36
                                                                                                } else {
                                                                                                    var v40 uint = 42
                                                                                                    _ = v40
                                                                                                }
                                                                                                v6 = 807
                                                                                            }
                                                                                            v20 = (v20 * 0)
                                                                                        }
                                                                                        if true {
                                                                                            var v41 int32 = -402
                                                                                            _ = v41
                                                                                        } else {
                                                                                            var v42 string = ("a" + ((("\n\t" + "") + ("foo bar" + (("\n\t" + ("\n\t" + (("foo bar" + "\n\t") + ("a" + (("" + "") + "hello"))))) + ("a" + (("\n\t" + ("a" + "foo bar")) + "foo bar"))))) + ((("\n\t" + "a") + "hello") + "")))
                                                                                            fmt.Println(fi11)
                                                                                            if true {
                                                                                                var v43 int8 = v4
                                                                                                _ = v43
                                                                                            } else {
                                                                                                v42 = v42
                                                                                            }
                                                                                            _ = v42
                                                                                        }
                                                                                        if false {
                                                                                            if false {
                                                                                                fmt.Println(v12, v1)
                                                                                                fmt.Println(v12)
                                                                                                if (v20 < 0) {
                                                                                                    v44 := GenS4{F0: (^(v4))}
                                                                                                    var v45 int16 = 0
                                                                                                    _ = v44
                                                                                                    _ = v45
                                                                                                } else {
                                                                                                    v46 := GenS6{F0: (^(v4))}
                                                                                                    v47 := v5[v3]
                                                                                                    var v48 uint8 = 2
                                                                                                    _ = v46
                                                                                                    _ = v47
                                                                                                    _ = v48
                                                                                                }
                                                                                            }
                                                                                        } else {
                                                                                            var v49 uint8 = (^(v20))
                                                                                            v3 = (v3 << ((0) & 63))
                                                                                            v2 -= v2
                                                                                            _ = v49
                                                                                        }
                                                                                    }
                                                                                    _ = v16
                                                                                    _ = v17
                                                                                    _ = v20
                                                                                }
                                                                                v50 := []int{445}
                                                                                v51 := v5[v3]
                                                                                _ = v50
                                                                                _ = v51
                                                                            }
                                                                            _ = v12
                                                                        }
                                                                    } else {
                                                                        v1 = v1
                                                                    }
                                                                default:
                                                                    v52 := []int{fi11, -1, (fi8 ^ fi8), (-(fi11))}
                                                                    _ = v52
                                                                }
                                                                _ = fi11
                                                            }
                                                        case "\n\t":
                                                            for fi53 := 0; fi53 < 5; fi53++ {
                                                                v2 = 42
                                                                switch ("\n\t" + "foo bar") {
                                                                case "\n\t":
                                                                    v1 -= (v1 | 42)
                                                                case "foo bar":
                                                                    fmt.Println(v2)
                                                                }
                                                                _ = fi53
                                                            }
                                                        case "a":
                                                            for fi54 := 0; fi54 < 20; fi54++ {
                                                                var v55 int64 = 1
                                                                if false {
                                                                    fmt.Println(v6, v1)
                                                                    var v56 uint64 = 18446744073709551615
                                                                    _ = v56
                                                                } else {
                                                                    switch ("foo bar" + "foo bar") {
                                                                    case "hello":
                                                                        var v57 float64 = float64(1.5)
                                                                        _ = v57
                                                                    case "a":
                                                                        fmt.Println(v2)
                                                                    }
                                                                }
                                                                _ = fi54
                                                                _ = v55
                                                            }
                                                        default:
                                                            for fi58 := 0; fi58 < 6; fi58++ {
                                                                v59 := []int{}
                                                                v2 -= v2
                                                                if (false && (true && true)) {
                                                                    if (-1 < 42) {
                                                                        v60 := GenS4{F0: (v4 / (((v4 / (((v4 ^ v4)) | 1))) | 1))}
                                                                        _ = v60
                                                                    } else {
                                                                        var v61 float32 = 0.1
                                                                        var v62 string = ("\n\t" + "hello")
                                                                        var v63 float64 = 1e10
                                                                        var v64 bool = (true || true)
                                                                        _ = v61
                                                                        _ = v62
                                                                        _ = v63
                                                                        _ = v64
                                                                    }
                                                                } else {
                                                                    v59 = append(v59, fi58)
                                                                    var v65 uint64 = 42
                                                                    _ = v65
                                                                }
                                                                _ = fi58
                                                                _ = v59
                                                            }
                                                        }
                                                    }
                                                case "hello":
                                                    v4 = (^(v4))
                                                }
                                            }
                                            var v66 int8 = 42
                                            v5[65535] = "hello"
                                            _ = fi8
                                            _ = v66
                                        }
                                        v67 := GenS5{F0: 1, F1: ("\n\t" + "a"), F2: 0}
                                        _ = v67
                                    }
                                }
                            case "foo bar":
                                for fi68 := 0; fi68 < 17; fi68++ {
                                    v69 := v5[(v3 >> ((65535) & 63))]
                                    if false {
                                        v69 = v69
                                        v70 := v5[2147483647]
                                        _ = v70
                                    } else {
                                        var v71 int32 = 91
                                        _ = v71
                                    }
                                    _ = fi68
                                    _ = v69
                                }
                            case "":
                                if (18446744073709551615 == 906) {
                                    switch 32768 {
                                    case -643:
                                        var v72 uint64 = 205
                                        _ = v72
                                    case 9223372036854775807:
                                        var v73 uint16 = (v1 * 2)
                                        _ = v73
                                    case -1:
                                        switch 1 {
                                        case 1:
                                            var v74 uint8 = (v2 & 1)
                                            _ = v74
                                        case -1:
                                            if true {
                                                if ((42 >= 255) && (1000 == 0)) {
                                                    var v75 uint = 462
                                                    if false {
                                                        switch 32767 {
                                                        case -531:
                                                            var v76 bool = false
                                                            _ = v76
                                                        default:
                                                            for fi77 := 0; fi77 < 17; fi77++ {
                                                                var v78 uint = v75
                                                                var v79 string = "\n\t"
                                                                var v80 uint16 = v1
                                                                _ = fi77
                                                                _ = v78
                                                                _ = v79
                                                                _ = v80
                                                            }
                                                        }
                                                        if false {
                                                            v75 = 347
                                                            if ((v2 >> ((2) & 63)) < (v2 + v2)) {
                                                                v75 -= (v75 % ((v75) | 1))
                                                                if (!(280 == 539)) {
                                                                    switch 116 {
                                                                    case -720:
                                                                        var v81 int32 = -564
                                                                        _ = v81
                                                                    case 2147483647:
                                                                        v1 -= uint16(v3)
                                                                    default:
                                                                        if (!true) {
                                                                            v82 := GenS1{F0: -841, F1: v2}
                                                                            v83 := GenS7{F0: 42, F1: (v75 % ((2) | 1))}
                                                                            fmt.Println(v4, v75, v1)
                                                                            _ = v82
                                                                            _ = v83
                                                                        } else {
                                                                            v84 := []int{127, -2147483648}
                                                                            _ = v84
                                                                        }
                                                                    }
                                                                    v85 := GenS6{F0: v4}
                                                                    _ = v85
                                                                }
                                                                v86 := []int{2147483647, 32767, 1}
                                                                _ = v86
                                                            } else {
                                                                if true {
                                                                    switch 817 {
                                                                    case 2147483647:
                                                                        var v87 int = 0
                                                                        _ = v87
                                                                    default:
                                                                        v5[(v3 ^ uint32(v2))] = ("foo bar" + "a")
                                                                    }
                                                                    v88 := GenS5{F0: v1, F1: "\n\t", F2: -2147483648}
                                                                    v89 := []int{42, 929, 128, 32768}
                                                                    _ = v88
                                                                    _ = v89
                                                                } else {
                                                                    if false {
                                                                        switch -651 {
                                                                        case 2147483647:
                                                                            delete(v5, (v3 % (((v3 / ((uint32(v1)) | 1))) | 1)))
                                                                        case 271:
                                                                            switch "\n\t" {
                                                                            case "":
                                                                                for fi90 := 0; fi90 < 5; fi90++ {
                                                                                    switch "a" {
                                                                                    case "a":
                                                                                        switch "a" {
                                                                                        case "hello":
                                                                                            switch "hello" {
                                                                                            case "hello":
                                                                                                v1 = uint16(v75)
                                                                                            default:
                                                                                                if false {
                                                                                                    v91 := GenS8{F0: (v3 - 44)}
                                                                                                    var v92 int32 = 42
                                                                                                    var v93 uint16 = 65535
                                                                                                    _ = v91
                                                                                                    _ = v92
                                                                                                    _ = v93
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    v75 += (v75 & v75)
                                                                                    if false {
                                                                                        v94 := GenS9{F0: false, F1: v2}
                                                                                        v95 := GenS4{F0: (v4 << ((1) & 63))}
                                                                                        fmt.Println(v1)
                                                                                        _ = v94
                                                                                        _ = v95
                                                                                    }
                                                                                    _ = fi90
                                                                                }
                                                                            }
                                                                        default:
                                                                            var v96 bool = false
                                                                            _ = v96
                                                                        }
                                                                        if true {
                                                                            v75 = (^(v75))
                                                                            for fi97 := 0; fi97 < 4; fi97++ {
                                                                                v6 -= 938
                                                                                if true {
                                                                                    var v98 int8 = v4
                                                                                    v99 := []int{(fi97 | 9223372036854775807), fi97, (fi97 % ((-635) | 1)), (^(fi97))}
                                                                                    v2 += v2
                                                                                    _ = v98
                                                                                    _ = v99
                                                                                }
                                                                                var v100 float32 = float32(float64(((1.5 - (((1e10 * ((float32(1.5) * 1e10) - (0.0 / 1.5))) - ((1.5 * -2.25) * (1e10 / float32(-2.25)))) / 1.5)) + ((0.0 * 0.0) - (float32((0.0 - 3.14159)) + 1e10)))))
                                                                                _ = fi97
                                                                                _ = v100
                                                                            }
                                                                        }
                                                                        v5[2147483647] = "foo bar"
                                                                    }
                                                                    switch -2147483648 {
                                                                    case -124:
                                                                        fmt.Println(v75)
                                                                    default:
                                                                        v3 += 464
                                                                    }
                                                                }
                                                                var v101 uint64 = 1
                                                                _ = v101
                                                            }
                                                            for fi102 := 0; fi102 < 2; fi102++ {
                                                                switch (fi102 >> ((879) & 63)) {
                                                                case 0:
                                                                    if (true || (("hello" != "") || (3.14159 == -2.25))) {
                                                                        v75 -= 2
                                                                    }
                                                                }
                                                                v3 = (v3 / (((^(v3))) | 1))
                                                                _ = fi102
                                                            }
                                                        }
                                                        var v103 uint64 = 18446744073709551615
                                                        _ = v103
                                                    } else {
                                                        v104 := GenS6{F0: -128}
                                                        _ = v104
                                                    }
                                                    _ = v75
                                                }
                                            } else {
                                                v5[(v3 >> ((1) & 63))] = "foo bar"
                                                fmt.Println(v4, v3)
                                            }
                                        case 42:
                                            v1 = v1
                                        default:
                                            var v105 uint64 = 0
                                            _ = v105
                                        }
                                    default:
                                        var v106 int = -460
                                        _ = v106
                                    }
                                    _ = len(v5)
                                    if true {
                                        fmt.Println(v2, v4, v2)
                                        for fi107 := 0; fi107 < 2; fi107++ {
                                            v2 = (^(v2))
                                            var v108 int32 = 1
                                            var v109 uint16 = 42
                                            _ = fi107
                                            _ = v108
                                            _ = v109
                                        }
                                        var v110 string = "hello"
                                        _ = v110
                                    }
                                }
                            }
                            _ = v7
                        } else {
                            v1 -= (^(v6))
                            v4 -= v4
                        }
                        v4 = (^(v4))
                        v2 += 42
                    } else {
                        for fi111 := 0; fi111 < 1; fi111++ {
                            v1 = 255
                            if (1 < 0) {
                                fmt.Println(fi111)
                            } else {
                                v3 += v3
                                var v112 uint32 = (^(v3))
                                for fi113 := 0; fi113 < 5; fi113++ {
                                    var v114 uint64 = 18446744073709551615
                                    _ = fi113
                                    _ = v114
                                }
                                _ = v112
                            }
                            _ = fi111
                        }
                    }
                    v115 := GenS10{F0: v1, F1: 9223372036854775807}
                    _ = v115
                } else {
                    if false {
                        if false {
                            var v116 bool = true
                            _ = v116
                        } else {
                            v117 := GenS4{F0: v4}
                            switch int(v3) {
                            case -509:
                                fmt.Println(v2, v4)
                            case -593:
                                if true {
                                    if (0 <= v3) {
                                        var v118 int = 701
                                        switch "a" {
                                        case "foo bar":
                                            if true {
                                                delete(v5, v3)
                                            } else {
                                                fmt.Println(v1, v6, v4)
                                                v119 := GenS4{F0: (v4 | v4)}
                                                _ = v119
                                            }
                                        case "":
                                            var v120 uint8 = v2
                                            _ = v120
                                        case "a":
                                            if false {
                                                v121 := []int{411, v118, v118}
                                                v122 := GenS11{F0: ("a" + (("\n\t" + ("foo bar" + ("a" + (("a" + "foo bar") + ("a" + ""))))) + "")), F1: "foo bar"}
                                                switch "foo bar" {
                                                case "foo bar":
                                                    var v123 float32 = 0.1
                                                    _ = v123
                                                case "\n\t":
                                                    fmt.Println(v1)
                                                case "a":
                                                    switch (v118 & 42) {
                                                    case 343:
                                                        if ((false || true) && ((v4 / ((v4) | 1)) != int8(v2))) {
                                                            v2 = v2
                                                            switch int(v2) {
                                                            case -987:
                                                                for fi124 := 0; fi124 < 5; fi124++ {
                                                                    for fi125 := 0; fi125 < 3; fi125++ {
                                                                        v126 := v5[(v3 << ((255) & 63))]
                                                                        _ = fi125
                                                                        _ = v126
                                                                    }
                                                                    switch ((((("hello" + "\n\t") + ("foo bar" + ("a" + ("hello" + "foo bar")))) + "hello") + "hello") + "") {
                                                                    case "":
                                                                        fmt.Println(v4, v2, v3)
                                                                    case "foo bar":
                                                                        if (v4 != (v4 - v4)) {
                                                                            v2 += 2
                                                                            v118 = (fi124 << ((42) & 63))
                                                                            var v127 uint32 = v3
                                                                            _ = v127
                                                                        }
                                                                    case "\n\t":
                                                                        switch "a" {
                                                                        case "a":
                                                                            v4 = (v4 << ((uint(v2)) & 63))
                                                                        case "hello":
                                                                            for fi128 := 0; fi128 < 19; fi128++ {
                                                                                switch 32768 {
                                                                                case -565:
                                                                                    fmt.Println(v1, fi128)
                                                                                case 2147483647:
                                                                                    var v129 bool = true
                                                                                    _ = v129
                                                                                case 42:
                                                                                    v3 = uint32(v1)
                                                                                    var v130 uint64 = 1
                                                                                    if (18446744073709551615 == (^(v130))) {
                                                                                        fmt.Println(fi128)
                                                                                        v131 := GenS9{F0: (42 == 1), F1: 255}
                                                                                        _ = v131
                                                                                    } else {
                                                                                        switch "foo bar" {
                                                                                        case "\n\t":
                                                                                            v118 -= 32767
                                                                                        case "foo bar":
                                                                                            v132 := v5[v3]
                                                                                            _ = v132
                                                                                        case "a":
                                                                                            var v133 int16 = -1
                                                                                            _ = v133
                                                                                        }
                                                                                        var v134 int32 = 22
                                                                                        v3 = v3
                                                                                        _ = v134
                                                                                    }
                                                                                    _ = len(v121)
                                                                                    _ = v130
                                                                                }
                                                                                v6 = 516
                                                                                v118 += fi128
                                                                                _ = fi128
                                                                            }
                                                                        case "":
                                                                            var v135 int64 = -9223372036854775808
                                                                            _ = v135
                                                                        default:
                                                                            v118 += v118
                                                                        }
                                                                    default:
                                                                        fmt.Println(v6)
                                                                    }
                                                                    _ = fi124
                                                                }
                                                            case 731:
                                                                v5[(v3 & v3)] = ("" + "hello")
                                                            case 127:
                                                                for fi136 := 0; fi136 < 14; fi136++ {
                                                                    v137 := GenS8{F0: v3}
                                                                    v138 := GenS1{F0: -467, F1: 255}
                                                                    for fi139 := 0; fi139 < 18; fi139++ {
                                                                        if false {
                                                                            v140 := GenS2{F0: int8(v2)}
                                                                            v141 := GenS12{F0: (("a" + ("foo bar" + ((("\n\t" + "hello") + ("" + "a")) + ("hello" + ("foo bar" + "\n\t"))))) + "hello"), F1: "a"}
                                                                            _ = v140
                                                                            _ = v141
                                                                        } else {
                                                                            var v142 string = "a"
                                                                            switch (v142 + v142) {
                                                                            case "foo bar":
                                                                                v121 = append(v121, fi136)
                                                                            }
                                                                            var v143 int8 = int8(v118)
                                                                            _ = v142
                                                                            _ = v143
                                                                        }
                                                                        v144 := v5[(^(v3))]
                                                                        var v145 float32 = 1e10
                                                                        _ = fi139
                                                                        _ = v144
                                                                        _ = v145
                                                                    }
                                                                    _ = fi136
                                                                    _ = v137
                                                                    _ = v138
                                                                }
                                                            default:
                                                                for fi146 := 0; fi146 < 18; fi146++ {
                                                                    v1 = v6
                                                                    var v147 float64 = 1e10
                                                                    for fi148 := 0; fi148 < 6; fi148++ {
                                                                        var v149 int8 = v4
                                                                        var v150 int16 = -996
                                                                        _ = len(v5)
                                                                        fmt.Println(v1)
                                                                        if false {
                                                                            v4 = (v4 * v4)
                                                                        } else {
                                                                            v151 := GenS2{F0: (v149 & v4)}
                                                                            v149 = v149
                                                                            v152 := GenS13{F0: 255}
                                                                            _ = v151
                                                                            _ = v152
                                                                        }
                                                                        if (v149 < (^(v149))) {
                                                                            var v153 float64 = (v147 / -2.25)
                                                                            v149 = v149
                                                                            _ = v153
                                                                        } else {
                                                                            fmt.Println(v150)
                                                                            if len(v121) > 0 {
                                                                                v121[0] = -684
                                                                            }
                                                                            v154 := GenS11{F0: ((((("foo bar" + ("foo bar" + "hello")) + "a") + (("hello" + (("" + "foo bar") + "")) + "a")) + "hello") + ""), F1: "foo bar"}
                                                                            _ = v154
                                                                        }
                                                                        v147 = float64(float32(((v147 - (((1.5 * ((0.0 - 0.1) / 0.1)) - v147) - v147)) + 3.14159)))
                                                                        v155 := GenS11{F0: "hello", F1: (("a" + "hello") + "foo bar")}
                                                                        v156 := GenS14{F0: 1, F1: -800, F2: fi148}
                                                                        if (0 <= -1) {
                                                                            var v157 int8 = v4
                                                                            if (true && ((!true) || true)) {
                                                                                v158 := GenS15{F0: -2147483648, F1: 218, F2: 447}
                                                                                switch fi146 {
                                                                                case 128:
                                                                                    var v159 float32 = (1e10 + 1.5)
                                                                                    _ = v159
                                                                                }
                                                                                var v160 float32 = (0.1 * ((-2.25 + ((float32(v147) / -2.25) + ((((0.1 / 1.0) - 1.5) + (-2.25 + (0.0 - float32(v147)))) / 1.0))) / 3.14159))
                                                                                _ = v158
                                                                                _ = v160
                                                                            } else {
                                                                                v1 = v1
                                                                                v150 -= (v150 | v150)
                                                                            }
                                                                            _ = v157
                                                                        } else {
                                                                            var v161 uint32 = 4294967295
                                                                            _ = v161
                                                                        }
                                                                        var v162 uint = 0
                                                                        _ = fi148
                                                                        _ = v149
                                                                        _ = v150
                                                                        _ = v155
                                                                        _ = v156
                                                                        _ = v162
                                                                    }
                                                                    _ = fi146
                                                                    _ = v147
                                                                }
                                                            }
                                                        }
                                                    }
                                                default:
                                                    if (42 != 136) {
                                                        _ = len(v5)
                                                        fmt.Println(v6, v6, v3)
                                                        if false {
                                                            v3 = v3
                                                        }
                                                    }
                                                }
                                                _ = v121
                                                _ = v122
                                            } else {
                                                if (!(float32(1e10) >= -2.25)) {
                                                    var v163 uint64 = 2
                                                    _ = v163
                                                } else {
                                                    switch (^(v118)) {
                                                    case 482:
                                                        var v164 int32 = -2147483648
                                                        _ = v164
                                                    case -9223372036854775808:
                                                        v165 := GenS16{F0: 674, F1: int32(v2)}
                                                        _ = v165
                                                    case 4294967295:
                                                        var v166 int = 128
                                                        _ = v166
                                                    }
                                                    v167 := GenS17{F0: 65535, F1: 65535}
                                                    v168 := v5[v3]
                                                    _ = v167
                                                    _ = v168
                                                }
                                                v2 = v2
                                            }
                                        }
                                        _ = v118
                                    }
                                    delete(v5, v3)
                                    switch 4294967295 {
                                    case 32767:
                                        v169 := []int{0, 193}
                                        _ = v169
                                    case -158:
                                        v2 += (v2 + 128)
                                    }
                                }
                            }
                            if ((!(!false)) || false) {
                                v170 := []int{}
                                v6 = (v1 >> ((42) & 63))
                                _ = v170
                            } else {
                                if true {
                                    v1 -= v6
                                    if true {
                                        var v171 int = -716
                                        _ = v171
                                    } else {
                                        var v172 uint64 = 18446744073709551615
                                        _ = v172
                                    }
                                    fmt.Println(v4, v4)
                                } else {
                                    v6 -= (v1 + v1)
                                }
                                for fi173 := 0; fi173 < 2; fi173++ {
                                    v174 := GenS4{F0: 113}
                                    _ = fi173
                                    _ = v174
                                }
                                v3 -= (v3 >> ((42) & 63))
                            }
                            _ = v117
                        }
                    } else {
                        if true {
                            v175 := GenS13{F0: 156}
                            _ = len(v5)
                            var v176 string = "\n\t"
                            _ = v175
                            _ = v176
                        }
                    }
                    var v177 uint = 42
                    _ = v177
                }
                _ = v6
            } else {
                if true {
                    v178 := []int{-85, 1, 4294967295, 1}
                    v179 := GenS18{F0: (v3 * v3), F1: 2, F2: 1}
                    _ = v178
                    _ = v179
                } else {
                    var v180 uint32 = v3
                    if true {
                        var v181 uint32 = v180
                        v2 += v2
                        v4 += v4
                        _ = v181
                    }
                    _ = v180
                }
                var v182 int8 = int8(v2)
                for fi183 := 0; fi183 < 2; fi183++ {
                    v2 += (v2 | 42)
                    if (!(1000 != int16(v1))) {
                        v184 := GenS19{F0: 42, F1: 533, F2: (v4 | (v4 / ((v182) | 1)))}
                        v185 := GenS20{F0: -32768, F1: 18446744073709551615, F2: 321}
                        _ = v184
                        _ = v185
                    }
                    switch fi183 {
                    case 128:
                        if (float64(((-2.25 / -2.25) / 1.5)) <= ((3.14159 * 0.0) + (1e10 - 3.14159))) {
                            fmt.Println(v3)
                            switch (fi183 | fi183) {
                            case -539:
                                v186 := GenS21{F0: false}
                                _ = v186
                            default:
                                switch int(v182) {
                                case 0:
                                    switch (fi183 >> ((19) & 63)) {
                                    case 1:
                                        if false {
                                            switch fi183 {
                                            case 32767:
                                                switch ("\n\t" + "a") {
                                                case "\n\t":
                                                    v187 := GenS19{F0: v1, F1: (^(v1)), F2: 127}
                                                    _ = v187
                                                }
                                            case 1:
                                                switch (^(fi183)) {
                                                case -321:
                                                    switch "" {
                                                    case "\n\t":
                                                        v188 := GenS10{F0: (^(v1)), F1: 257}
                                                        _ = v188
                                                    default:
                                                        var v189 int8 = (v4 >> ((985) & 63))
                                                        _ = v189
                                                    }
                                                }
                                            case -1:
                                                v190 := GenS22{F0: 1000}
                                                _ = v190
                                            }
                                            var v191 int32 = 42
                                            _ = v191
                                        } else {
                                            if true {
                                                v4 += v4
                                                var v192 uint16 = 1
                                                _ = v192
                                            }
                                            if true {
                                                if true {
                                                    v4 = v4
                                                    v3 = v3
                                                    v193 := GenS7{F0: 0, F1: 2}
                                                    _ = v193
                                                } else {
                                                    if ((32767 >= -414) || (!true)) {
                                                        switch fi183 {
                                                        case 128:
                                                            for fi194 := 0; fi194 < 4; fi194++ {
                                                                v195 := GenS5{F0: (v1 * 629), F1: ("foo bar" + ("a" + "hello")), F2: 2147483647}
                                                                v4 = 1
                                                                v2 = v2
                                                                _ = fi194
                                                                _ = v195
                                                            }
                                                        case 887:
                                                            switch (("" + ("hello" + "hello")) + "") {
                                                            case "":
                                                                for fi196 := 0; fi196 < 3; fi196++ {
                                                                    v197 := []int{fi196, (fi196 * (fi183 * -744)), (fi196 >> ((2) & 63)), fi196}
                                                                    v2 = (^(v2))
                                                                    _ = fi196
                                                                    _ = v197
                                                                }
                                                            }
                                                        }
                                                        v182 = (v4 % ((-128) | 1))
                                                    }
                                                    switch "hello" {
                                                    case "foo bar":
                                                        var v198 bool = true
                                                        _ = v198
                                                    case "a":
                                                        v3 = (v3 / (((v3 - v3)) | 1))
                                                    }
                                                    if true {
                                                        var v199 bool = (-2.25 >= 1.5)
                                                        v199 = (false && v199)
                                                        switch "hello" {
                                                        case "":
                                                            switch (fi183 - fi183) {
                                                            case 128:
                                                                var v200 int64 = -9223372036854775808
                                                                _ = v200
                                                            case 2147483647:
                                                                for fi201 := 0; fi201 < 7; fi201++ {
                                                                    var v202 int64 = 42
                                                                    v4 = (v182 + 1)
                                                                    _ = fi201
                                                                    _ = v202
                                                                }
                                                            case 4294967295:
                                                                var v203 int32 = 231
                                                                _ = v203
                                                            }
                                                        case "a":
                                                            if v199 {
                                                                v199 = v199
                                                            } else {
                                                                v204 := GenS6{F0: 114}
                                                                _ = v204
                                                            }
                                                        case "\n\t":
                                                            for fi205 := 0; fi205 < 16; fi205++ {
                                                                v206 := GenS23{F0: 661, F1: ("" + "a"), F2: 2147483648}
                                                                var v207 uint16 = (v1 << ((216) & 63))
                                                                switch ("\n\t" + ("hello" + (((("a" + "a") + ("a" + "foo bar")) + "") + "foo bar"))) {
                                                                case "foo bar":
                                                                    fmt.Println(fi205, v4)
                                                                case "\n\t":
                                                                    var v208 string = (("" + ("foo bar" + "")) + "foo bar")
                                                                    _ = v208
                                                                }
                                                                _ = fi205
                                                                _ = v206
                                                                _ = v207
                                                            }
                                                        }
                                                        _ = v199
                                                    }
                                                }
                                            } else {
                                                v209 := GenS21{F0: true}
                                                for fi210 := 0; fi210 < 13; fi210++ {
                                                    if false {
                                                        if (v2 == v2) {
                                                            v211 := GenS12{F0: ((("foo bar" + ("hello" + ("a" + ("foo bar" + "")))) + "a") + "hello"), F1: "a"}
                                                            if ((false && true) || (v4 >= -1)) {
                                                                switch (fi183 - (fi210 * 32767)) {
                                                                case 780:
                                                                    var v212 int8 = v182
                                                                    v213 := GenS6{F0: v4}
                                                                    fmt.Println(v4)
                                                                    _ = v212
                                                                    _ = v213
                                                                case -879:
                                                                    var v214 uint = 591
                                                                    _ = v214
                                                                default:
                                                                    if (!(v1 != 0)) {
                                                                        v182 = (v4 >> ((uint(v2)) & 63))
                                                                        v215 := GenS24{F0: fi183}
                                                                        _ = v215
                                                                    }
                                                                }
                                                            }
                                                            var v216 int8 = 100
                                                            _ = v211
                                                            _ = v216
                                                        }
                                                        v4 -= v182
                                                    } else {
                                                        var v217 uint64 = 0
                                                        v217 += v217
                                                        v4 += v4
                                                        v218 := GenS25{F0: (^(v1))}
                                                        _ = v217
                                                        _ = v218
                                                    }
                                                    v4 = v182
                                                    v1 += v1
                                                    v5[(^(v3))] = "hello"
                                                    var v219 int32 = 676
                                                    _ = fi210
                                                    _ = v219
                                                }
                                                fmt.Println(v182, fi183)
                                                _ = v209
                                            }
                                        }
                                    case -9223372036854775808:
                                        for fi220 := 0; fi220 < 13; fi220++ {
                                            fmt.Println(v4, v1, v4)
                                            fmt.Println(v4, fi220)
                                            v182 = -1
                                            _ = fi220
                                        }
                                    default:
                                        v221 := v5[(v3 ^ v3)]
                                        _ = v221
                                    }
                                case -1:
                                    fmt.Println(fi183)
                                default:
                                    for fi222 := 0; fi222 < 7; fi222++ {
                                        v182 += v182
                                        var v223 string = ("" + "")
                                        if (!(98 <= v2)) {
                                            v224 := GenS26{F0: (v3 >> ((18446744073709551615) & 63)), F1: 0}
                                            _ = v224
                                        }
                                        _ = fi222
                                        _ = v223
                                    }
                                }
                            }
                        } else {
                            v225 := GenS2{F0: v4}
                            if (!(0 != 437)) {
                                v3 = 0
                                if false {
                                    v1 -= (v1 * (v1 - v1))
                                }
                            }
                            if false {
                                switch int(v3) {
                                case 32767:
                                    var v226 uint16 = v1
                                    _ = v226
                                default:
                                    v2 -= (^(v2))
                                }
                                switch "hello" {
                                case "foo bar":
                                    var v227 int64 = 0
                                    _ = v227
                                case "hello":
                                    switch 127 {
                                    case -9223372036854775808:
                                        v2 += v2
                                    case -2147483648:
                                        var v228 int64 = 72
                                        _ = v228
                                    default:
                                        v182 += v182
                                    }
                                default:
                                    switch fi183 {
                                    case 127:
                                        for fi229 := 0; fi229 < 3; fi229++ {
                                            if (true && true) {
                                                v4 += (^(v182))
                                                v1 -= v1
                                            }
                                            _ = fi229
                                        }
                                    case -294:
                                        v230 := GenS10{F0: (^(v1)), F1: 1}
                                        _ = v230
                                    default:
                                        v3 = (v3 + (^(v3)))
                                    }
                                }
                                for fi231 := 0; fi231 < 15; fi231++ {
                                    v232 := []int{(fi183 + 0), (fi183 ^ fi231)}
                                    _ = fi231
                                    _ = v232
                                }
                            } else {
                                if true {
                                    v182 -= v182
                                } else {
                                    var v233 int32 = 1
                                    for fi234 := 0; fi234 < 12; fi234++ {
                                        if (!(int8(v1) >= (v182 & v182))) {
                                            v235 := []int{(fi183 << ((585) & 63)), fi234, 127}
                                            if true {
                                                var v236 int = fi183
                                                v237 := GenS27{F0: (0.1 >= 0.0), F1: 526}
                                                _ = v236
                                                _ = v237
                                            } else {
                                                var v238 int32 = int32(v2)
                                                v238 += v233
                                                var v239 uint16 = (v1 << ((4294967295) & 63))
                                                _ = v238
                                                _ = v239
                                            }
                                            _ = v235
                                        } else {
                                            v240 := GenS26{F0: (v3 % (((^(v3))) | 1)), F1: (^(v2))}
                                            var v241 uint8 = (^(v2))
                                            v242 := GenS18{F0: (v3 << ((4294967295) & 63)), F1: v1, F2: 4294967295}
                                            _ = v240
                                            _ = v241
                                            _ = v242
                                        }
                                        if (!true) {
                                            if (true && false) {
                                                var v243 uint8 = (v2 + v2)
                                                fmt.Println(fi234, v233, v4)
                                                v244 := GenS28{F0: (v243 << ((1) & 63)), F1: 2, F2: "a"}
                                                _ = v243
                                                _ = v244
                                            } else {
                                                switch (fi234 + fi183) {
                                                case 42:
                                                    if (-9223372036854775808 == 1) {
                                                        v233 = (v233 / ((v233) | 1))
                                                    }
                                                case 1:
                                                    v245 := GenS3{F0: (v182 >> ((65535) & 63)), F1: "a"}
                                                    _ = v245
                                                default:
                                                    v3 += 65535
                                                }
                                                if true {
                                                    var v246 string = ("foo bar" + (((("\n\t" + "") + ("hello" + "a")) + "hello") + ""))
                                                    var v247 int16 = 42
                                                    var v248 uint = 777
                                                    _ = v246
                                                    _ = v247
                                                    _ = v248
                                                } else {
                                                    switch (fi234 & -9223372036854775808) {
                                                    case 32768:
                                                        if (!((v3 >> ((18446744073709551615) & 63)) != v3)) {
                                                            switch ("hello" + "a") {
                                                            case "":
                                                                switch ("" + "a") {
                                                                case "\n\t":
                                                                    if false {
                                                                        delete(v5, v3)
                                                                    } else {
                                                                        v249 := GenS29{F0: (-(fi234))}
                                                                        v250 := GenS18{F0: (v3 << ((4294967295) & 63)), F1: (v1 - (v1 - 709)), F2: 4294967295}
                                                                        _ = v249
                                                                        _ = v250
                                                                    }
                                                                case "a":
                                                                    var v251 int32 = v233
                                                                    _ = v251
                                                                case "":
                                                                    if true {
                                                                        if (255 <= 42) {
                                                                            v252 := GenS13{F0: 2}
                                                                            fmt.Println(v182)
                                                                            var v253 bool = (!true)
                                                                            _ = v252
                                                                            _ = v253
                                                                        }
                                                                        fmt.Println(fi234, v233, v3)
                                                                    } else {
                                                                        v1 = (v1 - 300)
                                                                        fmt.Println(v4, fi183)
                                                                        if true {
                                                                            var v254 uint8 = v2
                                                                            v4 += (v4 + (v182 | (-(v182))))
                                                                            v233 = v233
                                                                            _ = v254
                                                                        }
                                                                    }
                                                                default:
                                                                    var v255 bool = true
                                                                    _ = v255
                                                                }
                                                            case "a":
                                                                var v256 int = (fi183 * 127)
                                                                _ = v256
                                                            }
                                                            v3 = (v3 & (v3 & 0))
                                                            v182 += 1
                                                            v257 := GenS9{F0: (true || (uint(v1) <= 544)), F1: 42}
                                                            _ = v257
                                                        } else {
                                                            if (v2 >= 13) {
                                                                v258 := GenS26{F0: (v3 >> ((42) & 63)), F1: v2}
                                                                _ = v258
                                                            } else {
                                                                var v259 uint8 = (v2 << ((27) & 63))
                                                                fmt.Println(v4, v4)
                                                                _ = v259
                                                            }
                                                            v2 += 37
                                                            if true {
                                                                var v260 uint = 4294967295
                                                                v261 := []int{fi183, fi234, (fi234 & (^(fi234))), fi183}
                                                                _ = v260
                                                                _ = v261
                                                            } else {
                                                                v1 = (v1 - (^(v1)))
                                                                switch -25 {
                                                                case 9223372036854775807:
                                                                    fmt.Println(v233, v1)
                                                                case 32767:
                                                                    var v262 int64 = -9223372036854775808
                                                                    _ = v262
                                                                case -944:
                                                                    var v263 int = (fi234 ^ 0)
                                                                    _ = v263
                                                                }
                                                                switch ("" + "") {
                                                                case "hello":
                                                                    v2 = v2
                                                                case "foo bar":
                                                                    v233 -= 42
                                                                default:
                                                                    var v264 uint32 = (^(v3))
                                                                    _ = v264
                                                                }
                                                            }
                                                            var v265 string = "a"
                                                            _ = v265
                                                        }
                                                    default:
                                                        v266 := GenS15{F0: v233, F1: 514, F2: 471}
                                                        _ = v266
                                                    }
                                                }
                                            }
                                        } else {
                                            var v267 bool = (true || true)
                                            v233 = v233
                                            fmt.Println(v267, v1)
                                            _ = v267
                                        }
                                        v268 := GenS5{F0: 317, F1: "hello", F2: v233}
                                        _ = fi234
                                        _ = v268
                                    }
                                    for fi269 := 0; fi269 < 14; fi269++ {
                                        if false {
                                            var v270 string = (((("a" + "") + "hello") + ("\n\t" + ("foo bar" + ""))) + ("a" + (((((((("" + "a") + "") + "hello") + ("hello" + "\n\t")) + "\n\t") + "foo bar") + "foo bar") + "foo bar")))
                                            _ = v270
                                        }
                                        _ = fi269
                                    }
                                    _ = v233
                                }
                                v3 += (v3 | 0)
                                switch -643 {
                                case -326:
                                    fmt.Println(v182)
                                case -691:
                                    if (!true) {
                                        v271 := []int{}
                                        _ = v271
                                    }
                                case 4294967295:
                                    fmt.Println(v182, v182)
                                default:
                                    var v272 int8 = 127
                                    _ = v272
                                }
                            }
                            _ = v225
                        }
                    case 42:
                        for fi273 := 0; fi273 < 10; fi273++ {
                            v274 := []int{(fi273 / ((128) | 1)), fi183, (fi273 - (fi183 - 921))}
                            _ = fi273
                            _ = v274
                        }
                    }
                    _ = fi183
                }
                _ = v182
            }
            v275 := GenS19{F0: v1, F1: v1, F2: v4}
            _ = v275
        }
    case "foo bar":
        v2 += (v2 << ((uint(v1)) & 63))
    default:
        var v276 int8 = v4
        _ = v276
    }
    var v277 int32 = 449
    if false {
        v277 = -2147483648
        fmt.Println(v2, v4, v2)
        var v278 int16 = 1000
        _ = v278
    }
    fmt.Println(v3, v2, v3)
    v1 = v1
    fmt.Println(v2)
    v279 := GenS30{F0: (v2 % ((v2) | 1))}
    var v280 uint = 4294967295
    v281 := []int{128}
    v3 -= v3
    var v282 float32 = 0.1
    var v283 int32 = (v277 | v277)
    var v284 int16 = 32767
    var v285 string = (("" + ("foo bar" + ("" + "foo bar"))) + ("foo bar" + "\n\t"))
    delete(v5, v3)
    var v286 uint32 = v3
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v277:", v277)
    fmt.Println("v280:", v280)
    fmt.Println("slicelen v281:", len(v281), v281)
    fmt.Println("v282:", v282)
    fmt.Println("v283:", v283)
    fmt.Println("v284:", v284)
    fmt.Println("v285:", v285)
    fmt.Println("v286:", v286)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v277
    _ = v279
    _ = v280
    _ = v281
    _ = v282
    _ = v283
    _ = v284
    _ = v285
    _ = v286
    _ = v3
    _ = v4
    _ = v5
}
