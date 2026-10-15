package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 int16; F1 bool }
type GenS3 struct { F0 string; F1 int8; F2 uint16 }
type GenS4 struct { F0 uint64; F1 string; F2 bool }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int16; F1 uint64 }
type GenS7 struct { F0 uint16; F1 bool; F2 uint64 }
type GenS8 struct { F0 int8; F1 uint8; F2 uint64 }
type GenS9 struct { F0 int; F1 bool; F2 uint32 }
type GenS10 struct { F0 int16; F1 int; F2 string }
type GenS11 struct { F0 uint32; F1 uint64; F2 int16 }
type GenS12 struct { F0 int64; F1 int16 }
type GenS13 struct { F0 int8; F1 uint64 }
type GenS14 struct { F0 uint; F1 int64; F2 bool }
type GenS15 struct { F0 uint }
type GenS16 struct { F0 uint; F1 int32; F2 bool }
func genFn0(p0 string) int32 {
    v4 := []int{32768}
    v5 := GenS1{F0: 42, F1: 255}
    _ = p0
    _ = v4
    _ = v5
    return 0
}

func main() {
    v1 := []int{749, -637, -423}
    var v2 uint = 2
    v3 := make(map[int]bool)
    var v6 float64 = 0.1
    switch "a" {
    case "hello":
        for fi7 := 0; fi7 < 4; fi7++ {
            var v8 int16 = 40
            _ = fi7
            _ = v8
        }
    }
    v6 += (float64(-2.25) * (float64(-2.25) * float64(-2.25)))
    var v9 int32 = -516
    if (42 <= -128) {
        var v10 uint16 = 42
        if true {
            var v11 int8 = 0
            v11 -= v11
            v6 += float64(-2.25)
            _ = v11
        } else {
            v12 := GenS1{F0: int16(v9), F1: v2}
            _ = v12
        }
        var v13 string = ("foo bar" + ("" + "foo bar"))
        _ = v10
        _ = v13
    }
    if false {
        v3[716] = false
        v1 = append(v1, 4294967295)
        _ = len(v1)
    }
    if (!true) {
        switch 1 {
        case 42:
            v6 -= ((1.5 * v6) / v6)
        case -2147483648:
            if (true && (!false)) {
                for fi14 := 0; fi14 < 7; fi14++ {
                    var v15 uint8 = 127
                    _ = fi14
                    _ = v15
                }
                if true {
                    v2 += v2
                    var v16 int8 = 1
                    for fi17 := 0; fi17 < 1; fi17++ {
                        for fi18 := 0; fi18 < 14; fi18++ {
                            var v19 int16 = -360
                            _ = fi18
                            _ = v19
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        var v20 int32 = (v9 - (v9 + v9))
                        _ = fi17
                        _ = v20
                    }
                    _ = v16
                } else {
                    v6 = float64(3.14159)
                    delete(v3, 42)
                }
            } else {
                if false {
                    v21 := GenS1{F0: 1000, F1: (v2 - 2)}
                    var v22 int = 952
                    if ((((3.14159 - (v6 + ((v6 / 3.14159) * float64(1.5)))) >= 0.0) || true) || (581 != 206)) {
                        v23 := GenS1{F0: 1, F1: (v2 * (v2 >> ((v2) & 63)))}
                        for fi24 := 0; fi24 < 15; fi24++ {
                            switch ("hello" + "\n\t") {
                            case "\n\t":
                                v25 := GenS1{F0: 135, F1: v2}
                                _ = v25
                            case "hello":
                                var v26 int8 = -71
                                _ = v26
                            case "a":
                                var v27 int8 = 50
                                _ = v27
                            }
                            v28 := GenS2{F0: 1, F1: true}
                            _ = fi24
                            _ = v28
                        }
                        _ = v23
                    } else {
                        fmt.Println(v2, v2, v9)
                        fmt.Println(v6, v2, v2)
                        fmt.Println(v6, v6, v9)
                    }
                    _ = v21
                    _ = v22
                } else {
                    for fi29 := 0; fi29 < 11; fi29++ {
                        switch ("foo bar" + "a") {
                        case "hello":
                            v30 := GenS1{F0: 32767, F1: (^(v2))}
                            _ = v30
                        case "":
                            v31 := GenS3{F0: "\n\t", F1: 121, F2: 645}
                            _ = v31
                        default:
                            _ = len(v1)
                        }
                        _ = fi29
                    }
                }
                if (!false) {
                    if true {
                        v2 = v2
                        switch "" {
                        case "hello":
                            if (65535 >= 32767) {
                                var v32 uint64 = 0
                                v3[32767] = (v2 >= v2)
                                _ = v32
                            } else {
                                if true {
                                    v33 := GenS4{F0: 0, F1: "", F2: false}
                                    v9 -= (-(v9))
                                    v9 -= (v9 * (-(v9)))
                                    _ = v33
                                } else {
                                    var v34 string = "foo bar"
                                    _ = v34
                                }
                                fmt.Println(v2, v9)
                                if (!(!(!((v2 << (((^(v2))) & 63)) == (v2 ^ 42))))) {
                                    switch 32767 {
                                    case 0:
                                        fmt.Println(v9, v9)
                                    case 128:
                                        if ((1000 == 0) && false) {
                                            if true {
                                                if true {
                                                    var v35 bool = (4294967295 > 0)
                                                    v36 := GenS5{F0: v35}
                                                    _ = v35
                                                    _ = v36
                                                } else {
                                                    switch "\n\t" {
                                                    case "\n\t":
                                                        for fi37 := 0; fi37 < 17; fi37++ {
                                                            for fi38 := 0; fi38 < 1; fi38++ {
                                                                v2 -= (v2 + v2)
                                                                v9 -= (v9 ^ 1)
                                                                var v39 uint64 = 42
                                                                _ = fi38
                                                                _ = v39
                                                            }
                                                            v6 = v6
                                                            _ = len(v3)
                                                            _ = fi37
                                                        }
                                                    default:
                                                        v9 += (v9 + (v9 * v9))
                                                    }
                                                    v6 = v6
                                                }
                                                if len(v1) > 0 {
                                                    v1[0] = 810
                                                }
                                            } else {
                                                var v40 int16 = -798
                                                _ = v40
                                            }
                                        } else {
                                            fmt.Println(v9)
                                        }
                                    }
                                    switch ("foo bar" + "a") {
                                    case "":
                                        v2 += 4294967295
                                    case "hello":
                                        delete(v3, 139)
                                    default:
                                        v41 := GenS5{F0: true}
                                        _ = v41
                                    }
                                    v2 -= (v2 >> ((255) & 63))
                                }
                            }
                        case "foo bar":
                            v2 = v2
                        case "\n\t":
                            for fi42 := 0; fi42 < 17; fi42++ {
                                v2 = v2
                                var v43 uint16 = 255
                                _ = fi42
                                _ = v43
                            }
                        default:
                            for fi44 := 0; fi44 < 14; fi44++ {
                                switch ("a" + ((((("\n\t" + (("hello" + ("a" + "foo bar")) + "hello")) + "") + "foo bar") + ("a" + ("foo bar" + ""))) + "")) {
                                case "a":
                                    v45 := GenS6{F0: -61, F1: 42}
                                    _ = v45
                                case "foo bar":
                                    if (32767 != 1) {
                                        if (42 >= v2) {
                                            v46 := GenS3{F0: ("" + "hello"), F1: -1, F2: 255}
                                            v47 := GenS7{F0: 0, F1: (2147483647 == 504), F2: 18446744073709551615}
                                            var v48 bool = (-9223372036854775808 >= 1)
                                            _ = v46
                                            _ = v47
                                            _ = v48
                                        } else {
                                            v2 = v2
                                            v49 := GenS3{F0: "foo bar", F1: -128, F2: 94}
                                            v50 := GenS8{F0: 100, F1: 111, F2: 0}
                                            _ = v49
                                            _ = v50
                                        }
                                        var v51 int32 = (v9 & v9)
                                        var v52 int8 = int8(v51)
                                        _ = v51
                                        _ = v52
                                    }
                                }
                                v9 = v9
                                _ = len(v1)
                                _ = fi44
                            }
                        }
                        if (128 <= 0) {
                            if (!(65535 < 1)) {
                                if true {
                                    if (-998 == 1000) {
                                        v2 += 811
                                    } else {
                                        v9 = (v9 | genFn0("\n\t"))
                                        switch "a" {
                                        case "hello":
                                            for fi53 := 0; fi53 < 14; fi53++ {
                                                for fi54 := 0; fi54 < 9; fi54++ {
                                                    v9 -= genFn0("foo bar")
                                                    fmt.Println(v9)
                                                    _ = fi54
                                                }
                                                var v55 int32 = (-(v9))
                                                v56 := GenS9{F0: (-(fi53)), F1: ((v2 & v2) >= 339), F2: 2147483648}
                                                _ = fi53
                                                _ = v55
                                                _ = v56
                                            }
                                        case "foo bar":
                                            var v57 uint64 = 0
                                            _ = v57
                                        default:
                                            switch "\n\t" {
                                            case "":
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            case "a":
                                                v2 = 181
                                            case "\n\t":
                                                var v58 uint8 = 2
                                                _ = v58
                                            default:
                                                var v59 uint8 = 127
                                                _ = v59
                                            }
                                        }
                                        if false {
                                            fmt.Println(v6, v2)
                                            for fi60 := 0; fi60 < 13; fi60++ {
                                                if (!true) {
                                                    switch "" {
                                                    case "a":
                                                        v2 -= 65535
                                                    }
                                                    delete(v3, (fi60 & fi60))
                                                } else {
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                }
                                                switch "foo bar" {
                                                case "a":
                                                    if false {
                                                        var v61 int16 = 42
                                                        _ = v61
                                                    } else {
                                                        if ((fi60 < fi60) && (false || (!true))) {
                                                            fmt.Println(fi60, v2, fi60)
                                                            for fi62 := 0; fi62 < 10; fi62++ {
                                                                v63 := GenS8{F0: 0, F1: 255, F2: 873}
                                                                _ = fi62
                                                                _ = v63
                                                            }
                                                            if false {
                                                                var v64 int32 = (v9 * v9)
                                                                _ = v64
                                                            } else {
                                                                fmt.Println(v9, v2)
                                                            }
                                                        } else {
                                                            v65 := GenS10{F0: 1, F1: (fi60 + fi60), F2: ""}
                                                            if true {
                                                                switch "hello" {
                                                                case "":
                                                                    var v66 string = ("a" + "")
                                                                    _ = v66
                                                                }
                                                                var v67 float64 = v6
                                                                _ = v67
                                                            } else {
                                                                switch "" {
                                                                case "hello":
                                                                    v6 += 0.1
                                                                }
                                                                v68 := GenS11{F0: 0, F1: 18446744073709551615, F2: -32768}
                                                                v6 += v6
                                                                _ = v68
                                                            }
                                                            _ = v65
                                                        }
                                                        var v69 int = fi60
                                                        if true {
                                                            for fi70 := 0; fi70 < 11; fi70++ {
                                                                if (true && false) {
                                                                    v69 += -2147483648
                                                                    if false {
                                                                        v2 -= (v2 << ((1) & 63))
                                                                        v71 := v3[(fi70 ^ (^(fi70)))]
                                                                        if (!(("foo bar" + "\n\t") > ("" + ("hello" + ("hello" + ""))))) {
                                                                            switch fi70 {
                                                                            case 128:
                                                                                if v71 {
                                                                                    v71 = (127 == -128)
                                                                                    var v72 uint32 = 260
                                                                                    _ = v72
                                                                                } else {
                                                                                    v1 = append(v1, v69)
                                                                                    v69 += (fi70 + int(v9))
                                                                                }
                                                                            }
                                                                            v73 := GenS12{F0: -117, F1: 437}
                                                                            if false {
                                                                                var v74 uint = 991
                                                                                if false {
                                                                                    if true {
                                                                                        switch v69 {
                                                                                        case -591:
                                                                                            v3[fi60] = true
                                                                                        case -2147483648:
                                                                                            if true {
                                                                                                v74 -= 65535
                                                                                                v69 += (fi60 >> ((v74) & 63))
                                                                                            } else {
                                                                                                var v75 float64 = float64(0.0)
                                                                                                switch ("foo bar" + "\n\t") {
                                                                                                case "":
                                                                                                    var v76 uint64 = 4294967295
                                                                                                    _ = v76
                                                                                                case "foo bar":
                                                                                                    fmt.Println(v75, fi60, v75)
                                                                                                case "hello":
                                                                                                    v77 := GenS1{F0: 1, F1: 268}
                                                                                                    _ = v77
                                                                                                }
                                                                                                var v78 uint64 = 18446744073709551615
                                                                                                _ = v75
                                                                                                _ = v78
                                                                                            }
                                                                                        }
                                                                                        v6 = v6
                                                                                    }
                                                                                    if v71 {
                                                                                        _ = len(v1)
                                                                                        v79 := v3[(v69 & fi70)]
                                                                                        var v80 int32 = (v9 - 1)
                                                                                        _ = v79
                                                                                        _ = v80
                                                                                    } else {
                                                                                        if v71 {
                                                                                            fmt.Println(v6, v6)
                                                                                            var v81 int8 = 22
                                                                                            _ = v81
                                                                                        } else {
                                                                                            var v82 float64 = 1.5
                                                                                            switch fi60 {
                                                                                            case 2147483647:
                                                                                                v6 = float64(1e10)
                                                                                            case 127:
                                                                                                if ("" >= "\n\t") {
                                                                                                    var v83 float64 = 0.0
                                                                                                    _ = v83
                                                                                                }
                                                                                            case 1:
                                                                                                if (float32(v82) < (-2.25 - ((0.1 / 1e10) - (-2.25 + ((3.14159 + 1.5) - -2.25))))) {
                                                                                                    var v84 uint32 = 2147483648
                                                                                                    _ = v84
                                                                                                } else {
                                                                                                    v85 := GenS12{F0: 1, F1: 0}
                                                                                                    v86 := GenS12{F0: 1, F1: 1000}
                                                                                                    v9 += (v9 & (v9 + -196))
                                                                                                    _ = v85
                                                                                                    _ = v86
                                                                                                }
                                                                                            default:
                                                                                                switch (("hello" + ("foo bar" + "")) + (((("foo bar" + ("\n\t" + "a")) + ("" + "\n\t")) + ("a" + ("foo bar" + "hello"))) + "\n\t")) {
                                                                                                case "foo bar":
                                                                                                    fmt.Println(v6, v6)
                                                                                                }
                                                                                                v69 = (-(fi70))
                                                                                            }
                                                                                            if (((v71 && ((!v71) || ("foo bar" == (("\n\t" + ("hello" + "\n\t")) + "foo bar")))) || v71) && v71) {
                                                                                                var v87 int = v69
                                                                                                var v88 int16 = -1
                                                                                                v89 := GenS2{F0: (v88 & v88), F1: (v74 > v74)}
                                                                                                if false {
                                                                                                    if (v71 && false) {
                                                                                                        v90 := GenS9{F0: (fi60 + (fi60 << ((v2) & 63))), F1: (0 <= 83), F2: 42}
                                                                                                        v88 -= 409
                                                                                                        v71 = (float32(v82) == (0.0 / float32((0.1 * v6))))
                                                                                                        _ = v90
                                                                                                    } else {
                                                                                                        v88 += 1
                                                                                                        var v91 int32 = -964
                                                                                                        _ = v91
                                                                                                    }
                                                                                                    var v92 int8 = -128
                                                                                                    _ = v92
                                                                                                }
                                                                                                if ((v88 >> ((0) & 63)) <= 0) {
                                                                                                    var v93 bool = true
                                                                                                    _ = v93
                                                                                                } else {
                                                                                                    v71 = false
                                                                                                    v87 = fi60
                                                                                                    v94 := GenS10{F0: 42, F1: (fi70 * (v69 << ((42) & 63))), F2: "foo bar"}
                                                                                                    _ = v94
                                                                                                }
                                                                                                v88 = v88
                                                                                                var v95 int16 = v88
                                                                                                _ = v87
                                                                                                _ = v88
                                                                                                _ = v89
                                                                                                _ = v95
                                                                                            } else {
                                                                                                if len(v1) >= 2 {
                                                                                                    v1 = v1[1:len(v1)]
                                                                                                }
                                                                                                v74 = 42
                                                                                                v96 := v3[(fi60 & 121)]
                                                                                                _ = v96
                                                                                            }
                                                                                            _ = v82
                                                                                        }
                                                                                        fmt.Println(v71)
                                                                                        switch "hello" {
                                                                                        case "":
                                                                                            fmt.Println(v6, v2)
                                                                                        }
                                                                                    }
                                                                                }
                                                                                _ = v74
                                                                            } else {
                                                                                v97 := GenS13{F0: 42, F1: 42}
                                                                                var v98 int64 = -340
                                                                                _ = v97
                                                                                _ = v98
                                                                            }
                                                                            v99 := GenS13{F0: 0, F1: 4294967295}
                                                                            if (42 <= 1) {
                                                                                v9 = v9
                                                                            }
                                                                            _ = v73
                                                                            _ = v99
                                                                        }
                                                                        _ = v71
                                                                    } else {
                                                                        if (!true) {
                                                                            var v100 int = fi60
                                                                            _ = len(v3)
                                                                            _ = v100
                                                                        } else {
                                                                            if len(v1) > 0 {
                                                                                v1[0] = (fi60 | 888)
                                                                            }
                                                                            var v101 float32 = (-2.25 - (((0.0 - (((1e10 / 1.5) + (0.0 + float32((((v6 + (v6 / 1.0)) + v6) / v6)))) * 0.0)) / (3.14159 + (3.14159 * -2.25))) / float32(((v6 / v6) - 1.5))))
                                                                            var v102 uint64 = 0
                                                                            var v103 float64 = 0.0
                                                                            _ = v101
                                                                            _ = v102
                                                                            _ = v103
                                                                        }
                                                                        delete(v3, (fi60 << (((v2 & (v2 | v2))) & 63)))
                                                                    }
                                                                    if false {
                                                                        if len(v1) >= 2 {
                                                                            v1 = v1[1:len(v1)]
                                                                        }
                                                                        switch "hello" {
                                                                        case "hello":
                                                                            v6 += float64(3.14159)
                                                                        case "a":
                                                                            var v104 int32 = genFn0("foo bar")
                                                                            _ = v104
                                                                        case "\n\t":
                                                                            fmt.Println(fi60)
                                                                        default:
                                                                            if (368 >= int32(v2)) {
                                                                                v9 = (v9 | v9)
                                                                            }
                                                                        }
                                                                        if true {
                                                                            v6 -= v6
                                                                            var v105 uint8 = 0
                                                                            _ = v105
                                                                        }
                                                                    }
                                                                    fmt.Println(v69)
                                                                    if true {
                                                                        var v106 int = 4294967295
                                                                        fmt.Println(fi70, fi70, v6)
                                                                        v106 += (v106 | int(v2))
                                                                        _ = v106
                                                                    } else {
                                                                        if true {
                                                                            switch "" {
                                                                            case "":
                                                                                var v107 int64 = -9223372036854775808
                                                                                _ = v107
                                                                            case "foo bar":
                                                                                v6 = v6
                                                                            }
                                                                        }
                                                                        switch fi70 {
                                                                        case -2147483648:
                                                                            v108 := GenS5{F0: false}
                                                                            _ = v108
                                                                        case 42:
                                                                            v2 = v2
                                                                        }
                                                                        v1 = append(v1, 1)
                                                                        if len(v1) >= 2 {
                                                                            v1 = v1[1:len(v1)]
                                                                        }
                                                                        if ((-2.25 >= ((float32((((v6 + 3.14159) * float64(1.5)) * (((v6 + v6) / v6) + v6))) + 3.14159) + -2.25)) && (("hello" + "") == ("foo bar" + "\n\t"))) {
                                                                            v69 = (fi60 % (((v69 >> (((v2 & v2)) & 63))) | 1))
                                                                        }
                                                                    }
                                                                    fmt.Println(v69, v6, v6)
                                                                }
                                                                v6 = v6
                                                                _ = fi70
                                                            }
                                                            for fi109 := 0; fi109 < 14; fi109++ {
                                                                var v110 uint64 = 4294967295
                                                                if false {
                                                                    var v111 int64 = 9223372036854775807
                                                                    _ = v111
                                                                } else {
                                                                    if false {
                                                                        fmt.Println(v110)
                                                                        delete(v3, (fi60 + (v69 / ((v69) | 1))))
                                                                    } else {
                                                                        v112 := GenS4{F0: 589, F1: "\n\t", F2: (-1 != 42)}
                                                                        v113 := GenS1{F0: 0, F1: v2}
                                                                        _ = v112
                                                                        _ = v113
                                                                    }
                                                                    v69 = int(v110)
                                                                    v114 := GenS4{F0: v110, F1: "", F2: true}
                                                                    _ = v114
                                                                }
                                                                if len(v1) > 0 {
                                                                    v1[0] = (v69 >> ((383) & 63))
                                                                }
                                                                _ = fi109
                                                                _ = v110
                                                            }
                                                        } else {
                                                            v115 := GenS8{F0: 100, F1: 130, F2: 1}
                                                            var v116 float32 = ((float32((v6 / 1.0)) * float32(v6)) - ((((0.1 * 0.0) + 0.1) + ((3.14159 * 1e10) + -2.25)) / float32(v6)))
                                                            if true {
                                                                v9 += (v9 ^ (v9 >> (((v2 >> (((^(v2))) & 63))) & 63)))
                                                                var v117 uint64 = 1
                                                                _ = v117
                                                            }
                                                            _ = v115
                                                            _ = v116
                                                        }
                                                        _ = v69
                                                    }
                                                case "\n\t":
                                                    fmt.Println(v9)
                                                case "":
                                                    v118 := GenS14{F0: 517, F1: -425, F2: true}
                                                    _ = v118
                                                default:
                                                    v119 := GenS1{F0: 42, F1: v2}
                                                    _ = v119
                                                }
                                                var v120 float32 = 1.5
                                                _ = fi60
                                                _ = v120
                                            }
                                            v2 = v2
                                        }
                                    }
                                    for fi121 := 0; fi121 < 10; fi121++ {
                                        var v122 float64 = v6
                                        for fi123 := 0; fi123 < 12; fi123++ {
                                            if false {
                                                v124 := GenS15{F0: (^(v2))}
                                                if len(v1) > 0 {
                                                    v1[0] = 386
                                                }
                                                fmt.Println(v6, v6, v2)
                                                _ = v124
                                            }
                                            _ = fi123
                                        }
                                        var v125 int = (fi121 + 4294967295)
                                        _ = fi121
                                        _ = v122
                                        _ = v125
                                    }
                                    fmt.Println(v9, v9, v2)
                                } else {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    if false {
                                        var v126 bool = true
                                        v127 := GenS8{F0: 28, F1: 186, F2: 973}
                                        fmt.Println(v9, v9)
                                        _ = v126
                                        _ = v127
                                    }
                                    var v128 int64 = 1
                                    _ = v128
                                }
                                var v129 float64 = 0.1
                                _ = v129
                            }
                            v3[9223372036854775807] = false
                            if false {
                                var v130 uint = 1
                                v2 -= (^(v130))
                                _ = v130
                            }
                        } else {
                            fmt.Println(v6, v9, v2)
                            fmt.Println(v9, v2)
                        }
                    } else {
                        for fi131 := 0; fi131 < 2; fi131++ {
                            v2 += v2
                            delete(v3, int(v9))
                            _ = fi131
                        }
                        v9 = genFn0("\n\t")
                        if (!(!(173 > 65535))) {
                            for fi132 := 0; fi132 < 16; fi132++ {
                                fmt.Println(v6, fi132)
                                _ = fi132
                            }
                        } else {
                            v9 -= (v9 / (((v9 >> ((0) & 63))) | 1))
                        }
                    }
                    v9 += (v9 << (((v2 >> (((v2 | (v2 % (((v2 << (((^(v2))) & 63))) | 1)))) & 63))) & 63))
                }
                v2 = (v2 ^ v2)
            }
        default:
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
    }
    for fi133 := 0; fi133 < 18; fi133++ {
        v134 := GenS16{F0: v2, F1: (v9 ^ v9), F2: false}
        _ = fi133
        _ = v134
    }
    if len(v1) > 0 {
        v1[0] = 32768
    }
    for fi135 := 0; fi135 < 10; fi135++ {
        var v136 int8 = 42
        if (-32768 > -1) {
            var v137 float32 = 0.0
            if ((!false) || false) {
                var v138 float64 = v6
                v136 = (-(v136))
                _ = v138
            }
            _ = v137
        }
        if (834 == 643) {
            fmt.Println(v136)
        }
        _ = fi135
        _ = v136
    }
    v139 := v3[-225]
    switch (("hello" + "foo bar") + ("\n\t" + "a")) {
    case "":
        fmt.Println(v6, v2, v2)
    case "a":
        _ = len(v1)
    case "hello":
        v139 = v139
    default:
        var v140 int8 = 100
        _ = v140
    }
    v139 = v139
    v141 := GenS7{F0: 581, F1: v139, F2: 234}
    if (!true) {
        if (0 >= 4294967295) {
            if (686 <= 338) {
                fmt.Println(v6, v2, v9)
                var v142 float32 = 1.5
                _ = v142
            } else {
                var v143 int32 = v9
                _ = v143
            }
            v144 := v3[314]
            v144 = (42 >= 1)
            _ = v144
        } else {
            for fi145 := 0; fi145 < 3; fi145++ {
                fmt.Println(v6, fi145, v2)
                fmt.Println(fi145, v2)
                _ = fi145
            }
        }
        for fi146 := 0; fi146 < 15; fi146++ {
            var v147 int8 = 1
            _ = fi146
            _ = v147
        }
        v148 := GenS14{F0: v2, F1: 0, F2: (42 != 2)}
        _ = v148
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v139:", v139)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v139
    _ = v141
    _ = v2
    _ = v3
    _ = v6
    _ = v9
}
