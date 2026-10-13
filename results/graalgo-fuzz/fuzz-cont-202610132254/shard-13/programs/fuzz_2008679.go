package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32; F1 int; F2 string }
type GenS3 struct { F0 int; F1 int64; F2 uint }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 uint; F1 int; F2 uint }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 bool }
type GenS12 struct { F0 bool }
type GenS13 struct { F0 uint64 }
type GenS14 struct { F0 int8; F1 int16 }
type GenS15 struct { F0 uint16 }
type GenS16 struct { F0 uint32 }
type GenS17 struct { F0 int64; F1 int32 }
type GenS18 struct { F0 int32; F1 int16 }
type GenS19 struct { F0 uint; F1 int64 }
type GenS20 struct { F0 bool; F1 int64; F2 uint16 }
type GenS21 struct { F0 int8 }
type GenS22 struct { F0 uint }
type GenS23 struct { F0 int8; F1 int8; F2 int16 }
type GenS24 struct { F0 int; F1 int8 }
type GenS25 struct { F0 uint32 }
type GenS26 struct { F0 uint; F1 uint; F2 uint8 }
type GenS27 struct { F0 bool; F1 bool; F2 int8 }
type GenS28 struct { F0 int8; F1 int64; F2 int16 }
type GenS29 struct { F0 uint32; F1 uint16 }
type GenS30 struct { F0 uint8; F1 uint16; F2 uint }
type GenS31 struct { F0 int; F1 int32 }
type GenS32 struct { F0 uint64 }
type GenS33 struct { F0 uint16 }
type GenS34 struct { F0 bool; F1 int32; F2 uint }
type GenS35 struct { F0 int8; F1 uint32; F2 int8 }
type GenS36 struct { F0 string }
type GenS37 struct { F0 uint64 }
type GenS38 struct { F0 uint16; F1 uint; F2 uint32 }
type GenS39 struct { F0 uint64; F1 int32; F2 string }
type GenS40 struct { F0 string }
type GenS41 struct { F0 uint16; F1 uint8; F2 int32 }
type GenS42 struct { F0 int16; F1 bool }
type GenS43 struct { F0 int16; F1 uint8 }
type GenS44 struct { F0 bool; F1 int64 }
type GenS45 struct { F0 int64 }
type GenS46 struct { F0 int32 }
type GenS47 struct { F0 int32; F1 int; F2 string }
type GenS48 struct { F0 int8 }
type GenS49 struct { F0 uint32; F1 bool }
type GenS50 struct { F0 int64 }
func genFn0() uint64 {
    if false {
        var v5 int16 = 1
        _ = v5
    } else {
        var v6 uint16 = 0
        if ((!true) && ((65535 != v6) && false)) {
            v7 := GenS1{F0: ""}
            var v8 bool = true
            for fi9 := 0; fi9 < 11; fi9++ {
                var v10 uint = 880
                v11 := []int{}
                _ = fi9
                _ = v10
                _ = v11
            }
            _ = v7
            _ = v8
        }
        var v12 uint16 = (v6 * v6)
        _ = v12
        _ = v6
    }
    switch 32768 {
    case -588:
        var v13 uint16 = 159
        _ = v13
    case 966:
        v14 := []int{2147483647, 128}
        _ = v14
    }
    return 281
}
func genFn1(p0 int8, p1 bool, p2 uint16) int64 {
    v15 := []int{-1, 32768, 127, -435}
    v15 = append(v15, 42)
    _ = len(v15)
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return 42
}

func main() {
    v1 := []string{("foo bar" + ((("foo bar" + "") + ("hello" + "hello")) + "a")), ("foo bar" + "hello"), ("foo bar" + ("\n\t" + (("a" + "\n\t") + ("" + ""))))}
    v2 := []string{}
    v3 := make(map[string]string)
    v4 := make(map[int]uint32)
    var v16 bool = (true && false)
    for fi17 := 0; fi17 < 20; fi17++ {
        v18 := GenS1{F0: "\n\t"}
        v16 = (false || (!(306 < 0)))
        fmt.Println(fi17, v16)
        _ = fi17
        _ = v18
    }
    if (!v16) {
        var v19 string = "a"
        v19 = ("a" + (v19 + v19))
        if v16 {
            for fi20 := 0; fi20 < 8; fi20++ {
                v16 = true
                var v21 uint16 = 32767
                v21 = (v21 % ((65535) | 1))
                _ = fi20
                _ = v21
            }
            _ = len(v3)
            var v22 bool = v16
            _ = v22
        } else {
            v4[32767] = 2147483648
            v16 = (v16 && (v16 || (!false)))
            for fi23 := 0; fi23 < 7; fi23++ {
                v19 = ((v19 + ("\n\t" + v19)) + "a")
                fmt.Println(v16, v16)
                _ = fi23
            }
        }
        _ = v19
    } else {
        var v24 uint32 = 995
        for fi25 := 0; fi25 < 20; fi25++ {
            v26 := GenS1{F0: ("foo bar" + "a")}
            delete(v3, "")
            if len(v1) > 0 {
                v1[0] = ("" + ("hello" + "foo bar"))
            }
            _ = fi25
            _ = v26
        }
        _ = v24
    }
    if ("\n\t" >= "hello") {
        for fi27 := 0; fi27 < 3; fi27++ {
            v16 = (false || true)
            v16 = ((!(!(v16 && (42 <= 0)))) && v16)
            _ = fi27
        }
        v28 := GenS1{F0: (("hello" + (((((("a" + "\n\t") + ("a" + "hello")) + ("a" + "hello")) + "foo bar") + ("\n\t" + "")) + "")) + "hello")}
        _ = v28
    } else {
        switch -9223372036854775808 {
        case 2147483647:
            if v16 {
                if false {
                    v16 = true
                    if v16 {
                        var v29 float64 = (0.0 * 0.1)
                        switch "\n\t" {
                        case "\n\t":
                            v16 = v16
                        case "hello":
                            var v30 uint32 = 2147483648
                            _ = v30
                        case "foo bar":
                            if v16 {
                                if v16 {
                                    v4[-329] = 531
                                    delete(v3, "foo bar")
                                }
                                _ = len(v1)
                                switch -828 {
                                case 0:
                                    fmt.Println(v16)
                                case -36:
                                    var v31 uint8 = 255
                                    _ = v31
                                case -545:
                                    var v32 string = "foo bar"
                                    _ = v32
                                }
                            }
                        default:
                            v16 = ("a" < "a")
                        }
                        _ = v29
                    } else {
                        var v33 int32 = 0
                        if v16 {
                            for fi34 := 0; fi34 < 1; fi34++ {
                                if v16 {
                                    fmt.Println(fi34, fi34)
                                    _ = len(v4)
                                    if len(v2) > 0 {
                                        v2[0] = ""
                                    }
                                }
                                _ = fi34
                            }
                            _ = len(v3)
                        } else {
                            if v16 {
                                var v35 float32 = (1.5 + 0.1)
                                v16 = v16
                                v33 = (v33 - 2147483647)
                                _ = v35
                            } else {
                                for fi36 := 0; fi36 < 10; fi36++ {
                                    var v37 uint8 = 214
                                    if v16 {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        if v16 {
                                            var v38 bool = v16
                                            v33 = -883
                                            var v39 float64 = -2.25
                                            _ = v38
                                            _ = v39
                                        } else {
                                            v40 := GenS1{F0: ("foo bar" + "hello")}
                                            for fi41 := 0; fi41 < 15; fi41++ {
                                                var v42 uint64 = 42
                                                switch (("a" + ("foo bar" + "a")) + ("hello" + "a")) {
                                                case "\n\t":
                                                    var v43 int = (fi41 + (fi36 + (fi36 << ((4294967295) & 63))))
                                                    _ = v43
                                                case "hello":
                                                    if true {
                                                        if v16 {
                                                            v37 = v37
                                                            var v44 bool = true
                                                            var v45 uint64 = 4294967295
                                                            _ = v44
                                                            _ = v45
                                                        }
                                                    } else {
                                                        var v46 uint = 4294967295
                                                        var v47 uint8 = (^(v37))
                                                        _ = v46
                                                        _ = v47
                                                    }
                                                    switch fi41 {
                                                    case 42:
                                                        v37 -= v37
                                                    case -643:
                                                        var v48 uint32 = 957
                                                        _ = v48
                                                    case 127:
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                    default:
                                                        fmt.Println(v42)
                                                        v1 = append(v1, ("a" + ("foo bar" + "hello")))
                                                        var v49 uint32 = 2
                                                        if true {
                                                            if (v16 && ((440 != 32767) && (int32(fi41) < (v33 / ((v33) | 1))))) {
                                                                v37 += (v37 % (((v37 * (v37 ^ 0))) | 1))
                                                            }
                                                        } else {
                                                            v50 := GenS1{F0: ("\n\t" + (("foo bar" + "foo bar") + (("" + ("foo bar" + "")) + ((((("hello" + ("hello" + "foo bar")) + "a") + "") + "a") + (("hello" + "a") + "hello")))))}
                                                            v16 = (1000 == 341)
                                                            _ = v50
                                                        }
                                                        _ = v49
                                                    }
                                                    if (!(-1 < -997)) {
                                                        v42 = (^(v42))
                                                    } else {
                                                        v42 = v42
                                                    }
                                                }
                                                _ = fi41
                                                _ = v42
                                            }
                                            _ = v40
                                        }
                                    } else {
                                        var v51 uint32 = 0
                                        _ = v51
                                    }
                                    fmt.Println(fi36, v37, v33)
                                    _ = fi36
                                    _ = v37
                                }
                                if false {
                                    if (v33 > (v33 | v33)) {
                                        v52 := GenS1{F0: (("\n\t" + ("hello" + ("\n\t" + ("" + "\n\t")))) + (("foo bar" + ("a" + ((("" + "foo bar") + "foo bar") + "\n\t"))) + "\n\t"))}
                                        v53 := GenS1{F0: (("a" + ("foo bar" + "")) + "a")}
                                        v54 := GenS2{F0: 1, F1: -9223372036854775808, F2: ""}
                                        _ = v52
                                        _ = v53
                                        _ = v54
                                    }
                                } else {
                                    v55 := GenS3{F0: 0, F1: 42, F2: 42}
                                    _ = v55
                                }
                                if v16 {
                                    var v56 uint64 = genFn0()
                                    _ = v56
                                }
                            }
                            if v16 {
                                if true {
                                    v57 := GenS4{F0: 2}
                                    if (!(!false)) {
                                        if false {
                                            switch 32767 {
                                            case 42:
                                                v3[(("foo bar" + "a") + ("hello" + "\n\t"))] = "foo bar"
                                            case -9223372036854775808:
                                                v16 = (127 >= 42)
                                            default:
                                                for fi58 := 0; fi58 < 9; fi58++ {
                                                    var v59 uint16 = 42
                                                    for fi60 := 0; fi60 < 11; fi60++ {
                                                        if (!(int64(fi60) == int64(v59))) {
                                                            var v61 int32 = (^(v33))
                                                            if false {
                                                                fmt.Println(fi60)
                                                            }
                                                            var v62 bool = true
                                                            _ = v61
                                                            _ = v62
                                                        } else {
                                                            _ = len(v3)
                                                            var v63 uint32 = 65535
                                                            _ = v63
                                                        }
                                                        v64 := GenS1{F0: "\n\t"}
                                                        _ = fi60
                                                        _ = v64
                                                    }
                                                    _ = fi58
                                                    _ = v59
                                                }
                                            }
                                            var v65 uint64 = 0
                                            if ((32767 == 32767) && v16) {
                                                v33 += v33
                                                switch "a" {
                                                case "hello":
                                                    var v66 uint8 = 42
                                                    _ = v66
                                                default:
                                                    if len(v1) > 0 {
                                                        v1[0] = ""
                                                    }
                                                }
                                                switch 9223372036854775807 {
                                                case 1:
                                                    switch -9223372036854775808 {
                                                    case -9223372036854775808:
                                                        if v16 {
                                                            v33 = v33
                                                            switch "a" {
                                                            case "foo bar":
                                                                for fi67 := 0; fi67 < 18; fi67++ {
                                                                    v16 = ((-(fi67)) > -2147483648)
                                                                    _ = fi67
                                                                }
                                                            case "hello":
                                                                v68 := GenS2{F0: 2147483647, F1: 2147483647, F2: "\n\t"}
                                                                _ = v68
                                                            case "":
                                                                if v16 {
                                                                    v33 = -1
                                                                    var v69 int32 = 2147483647
                                                                    _ = v69
                                                                }
                                                            }
                                                            v70 := GenS5{F0: -262}
                                                            _ = v70
                                                        } else {
                                                            var v71 uint8 = 128
                                                            if v16 {
                                                                v72 := GenS2{F0: 0, F1: 32767, F2: "\n\t"}
                                                                _ = v72
                                                            } else {
                                                                if v16 {
                                                                    switch 421 {
                                                                    case 0:
                                                                        for fi73 := 0; fi73 < 8; fi73++ {
                                                                            v33 = v33
                                                                            _ = fi73
                                                                        }
                                                                    }
                                                                    var v74 int8 = -55
                                                                    var v75 uint16 = 477
                                                                    _ = v74
                                                                    _ = v75
                                                                } else {
                                                                    fmt.Println(v65, v33)
                                                                    var v76 uint = 18446744073709551615
                                                                    switch ("foo bar" + "a") {
                                                                    case "hello":
                                                                        v77 := GenS2{F0: 61, F1: 323, F2: "hello"}
                                                                        _ = v77
                                                                    case "a":
                                                                        var v78 uint8 = 255
                                                                        _ = v78
                                                                    default:
                                                                        v79 := GenS6{F0: 0}
                                                                        _ = v79
                                                                    }
                                                                    _ = v76
                                                                }
                                                                fmt.Println(v65, v71, v65)
                                                                fmt.Println(v71, v16)
                                                            }
                                                            _ = v71
                                                        }
                                                    }
                                                case -430:
                                                    switch "" {
                                                    case "hello":
                                                        fmt.Println(v65, v16, v33)
                                                    default:
                                                        for fi80 := 0; fi80 < 15; fi80++ {
                                                            v81 := GenS5{F0: (fi80 / (((-(fi80))) | 1))}
                                                            for fi82 := 0; fi82 < 17; fi82++ {
                                                                if false {
                                                                    var v83 int8 = -44
                                                                    _ = v83
                                                                } else {
                                                                    if (v33 < 0) {
                                                                        if (((v16 && (v16 || v16)) && (400 < 4294967295)) || (42 != 0)) {
                                                                            var v84 int64 = -114
                                                                            var v85 int32 = -2147483648
                                                                            v86 := GenS1{F0: (("foo bar" + ("" + "foo bar")) + "a")}
                                                                            _ = v84
                                                                            _ = v85
                                                                            _ = v86
                                                                        }
                                                                        if (((2 != 255) && (0 > 0)) && (v16 || v16)) {
                                                                            v87 := v4[fi82]
                                                                            _ = v87
                                                                        } else {
                                                                            var v88 uint16 = 32767
                                                                            var v89 uint32 = 725
                                                                            _ = v88
                                                                            _ = v89
                                                                        }
                                                                        switch "" {
                                                                        case "foo bar":
                                                                            v90 := GenS7{F0: v65}
                                                                            _ = v90
                                                                        case "a":
                                                                            v65 -= v65
                                                                        default:
                                                                            var v91 int8 = -128
                                                                            _ = v91
                                                                        }
                                                                        v65 += (v65 % ((v65) | 1))
                                                                    }
                                                                    var v92 int16 = 741
                                                                    _ = v92
                                                                }
                                                                v65 = v65
                                                                switch (fi82 ^ 32768) {
                                                                case -1:
                                                                    v93 := GenS8{F0: genFn0()}
                                                                    _ = v93
                                                                }
                                                                _ = fi82
                                                            }
                                                            _ = fi80
                                                            _ = v81
                                                        }
                                                    }
                                                case 4294967295:
                                                    var v94 int64 = 9223372036854775807
                                                    _ = v94
                                                default:
                                                    var v95 string = "a"
                                                    _ = v95
                                                }
                                            }
                                            _ = v65
                                        } else {
                                            for fi96 := 0; fi96 < 14; fi96++ {
                                                switch (fi96 * (^(fi96))) {
                                                case -9223372036854775808:
                                                    var v97 float64 = float64((3.14159 * (float32(-2.25) + ((float32(0.1) + -2.25) / (-2.25 + 3.14159)))))
                                                    _ = v97
                                                case 352:
                                                    v33 = v33
                                                case 1:
                                                    var v98 int64 = int64(fi96)
                                                    _ = v98
                                                }
                                                switch (fi96 >> ((357) & 63)) {
                                                case 1:
                                                    if (v16 && (0 != 18446744073709551615)) {
                                                        v33 = (v33 + (v33 / (((v33 - v33)) | 1)))
                                                        switch ("hello" + "hello") {
                                                        case "\n\t":
                                                            v99 := GenS8{F0: 42}
                                                            _ = v99
                                                        default:
                                                            v100 := GenS5{F0: fi96}
                                                            _ = v100
                                                        }
                                                        switch "hello" {
                                                        case "foo bar":
                                                            if (42 != 2) {
                                                                fmt.Println(v33, v16)
                                                                v33 = (v33 * v33)
                                                            } else {
                                                                var v101 int = 1
                                                                for fi102 := 0; fi102 < 7; fi102++ {
                                                                    if len(v2) > 0 {
                                                                        v2[0] = "foo bar"
                                                                    }
                                                                    var v103 int32 = v33
                                                                    var v104 uint8 = 9
                                                                    _ = fi102
                                                                    _ = v103
                                                                    _ = v104
                                                                }
                                                                delete(v4, (^(fi96)))
                                                                _ = v101
                                                            }
                                                        case "":
                                                            switch -2147483648 {
                                                            case 82:
                                                                if (!(0.1 < (1e10 + (float32(3.14159) + 1e10)))) {
                                                                    for fi105 := 0; fi105 < 1; fi105++ {
                                                                        v33 = -198
                                                                        v33 = 787
                                                                        v106 := GenS4{F0: 2}
                                                                        _ = fi105
                                                                        _ = v106
                                                                    }
                                                                }
                                                            case 42:
                                                                var v107 uint = 18446744073709551615
                                                                _ = v107
                                                            default:
                                                                fmt.Println(v16, v16, v33)
                                                            }
                                                        }
                                                    } else {
                                                        var v108 float64 = 3.14159
                                                        var v109 int64 = -591
                                                        _ = v108
                                                        _ = v109
                                                    }
                                                case -2147483648:
                                                    var v110 int32 = (v33 ^ (v33 & 1))
                                                    _ = v110
                                                default:
                                                    v33 -= 723
                                                }
                                                v16 = v16
                                                _ = fi96
                                            }
                                        }
                                        v16 = (!v16)
                                        var v111 int = 1
                                        _ = v111
                                    }
                                    v16 = true
                                    _ = v57
                                } else {
                                    v112 := v4[328]
                                    for fi113 := 0; fi113 < 19; fi113++ {
                                        fmt.Println(v112)
                                        _ = fi113
                                    }
                                    v114 := GenS4{F0: 0}
                                    _ = v112
                                    _ = v114
                                }
                            } else {
                                if (((1.5 / 1.5) - 0.1) != 3.14159) {
                                    v33 -= v33
                                    var v115 int16 = 298
                                    _ = v115
                                }
                            }
                        }
                        for fi116 := 0; fi116 < 8; fi116++ {
                            v33 = (v33 << ((255) & 63))
                            _ = fi116
                        }
                        _ = v33
                    }
                }
            }
        case -2147483648:
            for fi117 := 0; fi117 < 16; fi117++ {
                for fi118 := 0; fi118 < 16; fi118++ {
                    v119 := v3["hello"]
                    _ = fi118
                    _ = v119
                }
                for fi120 := 0; fi120 < 5; fi120++ {
                    v4[fi120] = 1
                    v121 := GenS5{F0: (fi117 | (^(fi117)))}
                    _ = fi120
                    _ = v121
                }
                _ = fi117
            }
        }
        for fi122 := 0; fi122 < 12; fi122++ {
            v123 := GenS8{F0: 753}
            for fi124 := 0; fi124 < 10; fi124++ {
                switch "" {
                case "foo bar":
                    if (18446744073709551615 >= 0) {
                        v16 = v16
                        if (v16 && false) {
                            v125 := GenS9{F0: 0, F1: fi124, F2: 304}
                            if false {
                                switch "foo bar" {
                                case "hello":
                                    v126 := GenS10{F0: 42}
                                    _ = v126
                                case "":
                                    var v127 uint32 = 0
                                    _ = v127
                                case "foo bar":
                                    v128 := GenS11{F0: v16}
                                    _ = v128
                                }
                                if (1.5 != 0.1) {
                                    if v16 {
                                        if ((1 != 42) && (((4294967295 != 891) || (1 >= 676)) && (3.14159 != 1.5))) {
                                            if (("a" != (("foo bar" + ("hello" + "foo bar")) + "")) && v16) {
                                                switch 42 {
                                                case -366:
                                                    if len(v2) > 0 {
                                                        v2[0] = ("a" + "\n\t")
                                                    }
                                                default:
                                                    v16 = v16
                                                }
                                            } else {
                                                var v129 float64 = 0.1
                                                if v16 {
                                                    var v130 int16 = 42
                                                    var v131 float32 = (-2.25 * 1e10)
                                                    if (47 < 255) {
                                                        var v132 uint16 = 829
                                                        _ = v132
                                                    } else {
                                                        var v133 int16 = v130
                                                        _ = len(v1)
                                                        fmt.Println(v16)
                                                        _ = v133
                                                    }
                                                    _ = v130
                                                    _ = v131
                                                }
                                                v134 := GenS7{F0: 42}
                                                _ = v129
                                                _ = v134
                                            }
                                            if false {
                                                v135 := GenS1{F0: (("a" + ("\n\t" + "\n\t")) + "\n\t")}
                                                if false {
                                                    switch (("\n\t" + "\n\t") + "\n\t") {
                                                    case "a":
                                                        v136 := GenS3{F0: (fi122 % ((fi122) | 1)), F1: 84, F2: 2}
                                                        _ = v136
                                                    case "":
                                                        v16 = ("\n\t" == ("a" + "hello"))
                                                    }
                                                    if v16 {
                                                        _ = len(v2)
                                                    }
                                                    if (!(!true)) {
                                                        if (3.14159 >= float32(3.14159)) {
                                                            if v16 {
                                                                v16 = true
                                                                v137 := GenS12{F0: (false || (1 != 0))}
                                                                switch ("hello" + ("hello" + "hello")) {
                                                                case "a":
                                                                    var v138 string = "a"
                                                                    _ = v138
                                                                case "foo bar":
                                                                    v139 := GenS2{F0: 1, F1: fi124, F2: "foo bar"}
                                                                    _ = v139
                                                                }
                                                                v16 = false
                                                                _ = v137
                                                            } else {
                                                                var v140 int8 = -16
                                                                v2 = append(v2, "a")
                                                                switch "hello" {
                                                                case "foo bar":
                                                                    v140 = 1
                                                                case "hello":
                                                                    var v141 uint = 255
                                                                    _ = v141
                                                                }
                                                                _ = v140
                                                            }
                                                        } else {
                                                            switch -1 {
                                                            case 0:
                                                                v142 := GenS7{F0: 4294967295}
                                                                _ = v142
                                                            }
                                                            v4[-733] = 65535
                                                        }
                                                    } else {
                                                        fmt.Println(fi122)
                                                        var v143 float32 = 0.1
                                                        fmt.Println(fi124)
                                                        v143 = 0.0
                                                        fmt.Println(fi122, fi124, v143)
                                                        _ = v143
                                                    }
                                                    if (!(1 >= 4294967295)) {
                                                        fmt.Println(fi124, fi124)
                                                        v144 := v3["\n\t"]
                                                        if (v16 || (226 > 42)) {
                                                            fmt.Println(v144, v16)
                                                        }
                                                        v144 = (v144 + v144)
                                                        _ = v144
                                                    } else {
                                                        var v145 int64 = -45
                                                        fmt.Println(fi124, v16)
                                                        v145 += v145
                                                        _ = v145
                                                    }
                                                    v16 = ((true && v16) || v16)
                                                } else {
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    _ = len(v2)
                                                    v16 = v16
                                                    fmt.Println(fi122)
                                                }
                                                fmt.Println(fi122, fi122)
                                                v146 := GenS10{F0: 4294967295}
                                                _ = v135
                                                _ = v146
                                            }
                                        }
                                        v16 = (!(210 > 129))
                                        var v147 uint64 = genFn0()
                                        _ = v147
                                    }
                                }
                            }
                            var v148 uint64 = 0
                            _ = v125
                            _ = v148
                        } else {
                            v149 := GenS1{F0: ("hello" + "a")}
                            v150 := GenS1{F0: (("a" + ("hello" + "foo bar")) + "\n\t")}
                            v1 = append(v1, "")
                            _ = v149
                            _ = v150
                        }
                        var v151 float64 = 1.5
                        _ = v151
                    }
                case "":
                    v152 := GenS2{F0: 2147483647, F1: fi122, F2: ("\n\t" + ("" + "\n\t"))}
                    _ = v152
                }
                v1 = append(v1, "hello")
                _ = fi124
            }
            if (!(495 == -1)) {
                for fi153 := 0; fi153 < 4; fi153++ {
                    v16 = false
                    if (-2.25 != 0.1) {
                        var v154 uint = 4294967295
                        v16 = v16
                        var v155 int64 = 222
                        _ = v154
                        _ = v155
                    }
                    if (!v16) {
                        var v156 uint8 = 103
                        var v157 uint64 = genFn0()
                        _ = v156
                        _ = v157
                    }
                    _ = fi153
                }
                var v158 string = ("hello" + (("" + "a") + ""))
                v158 = ((v158 + "a") + "a")
                _ = v158
            } else {
                switch (fi122 - (fi122 - -800)) {
                case -2147483648:
                    v159 := GenS5{F0: 1}
                    _ = v159
                case -118:
                    v160 := GenS13{F0: 491}
                    _ = v160
                case 0:
                    if (3.14159 == 0.1) {
                        switch ("\n\t" + "") {
                        case "a":
                            if true {
                                var v161 float32 = 1.5
                                _ = v161
                            } else {
                                for fi162 := 0; fi162 < 6; fi162++ {
                                    if v16 {
                                        v16 = (2 > 42)
                                    } else {
                                        if (918 != 4294967295) {
                                            switch 128 {
                                            case 42:
                                                v163 := v4[0]
                                                _ = v163
                                            default:
                                                var v164 uint64 = 990
                                                _ = v164
                                            }
                                            _ = len(v3)
                                            var v165 int32 = 522
                                            _ = v165
                                        }
                                        v16 = (4294967295 == 0)
                                        v16 = (87 >= 0)
                                    }
                                    if false {
                                        if (!(v16 && (0 == 127))) {
                                            switch (^(fi122)) {
                                            case -2147483648:
                                                v16 = (!false)
                                            case 10:
                                                if (!false) {
                                                    v1 = append(v1, "a")
                                                    if ((619 >= -1) && (3.14159 <= ((0.1 * float64((0.0 / ((1e10 + (3.14159 - 0.0)) - ((0.1 * 1e10) / 3.14159))))) - -2.25))) {
                                                        v16 = v16
                                                        switch "\n\t" {
                                                        case "\n\t":
                                                            v16 = (42 != 42)
                                                        case "foo bar":
                                                            v166 := GenS4{F0: 65535}
                                                            if (255 < 42) {
                                                                v3[("\n\t" + ((("" + "") + "\n\t") + ""))] = "\n\t"
                                                                var v167 int8 = 42
                                                                _ = v167
                                                            }
                                                            delete(v3, "\n\t")
                                                            _ = v166
                                                        }
                                                        var v168 int32 = 2147483647
                                                        _ = v168
                                                    } else {
                                                        if false {
                                                            var v169 int16 = 32767
                                                            _ = v169
                                                        }
                                                        if (false && (!(v16 || ((3.14159 - ((float32(3.14159) - 1e10) + 1e10)) < 1.5)))) {
                                                            if v16 {
                                                                var v170 uint16 = 21
                                                                var v171 int32 = 0
                                                                _ = v170
                                                                _ = v171
                                                            } else {
                                                                if true {
                                                                    if true {
                                                                        var v172 uint64 = 161
                                                                        _ = v172
                                                                    } else {
                                                                        fmt.Println(fi122)
                                                                        var v173 int = fi122
                                                                        switch ("" + "") {
                                                                        case "a":
                                                                            if true {
                                                                                var v174 float64 = float64(((1.5 * -2.25) * (1.5 - 1e10)))
                                                                                switch -2147483648 {
                                                                                case 32768:
                                                                                    v16 = v16
                                                                                case -2147483648:
                                                                                    v174 += (float64(-2.25) - 3.14159)
                                                                                case -9223372036854775808:
                                                                                    v16 = ((float32(1e10) / 0.1) > 0.0)
                                                                                }
                                                                                if (-123 > 1) {
                                                                                    var v175 int16 = -32768
                                                                                    _ = v175
                                                                                }
                                                                                _ = v174
                                                                            } else {
                                                                                var v176 int = fi162
                                                                                _ = v176
                                                                            }
                                                                            v173 += fi122
                                                                            var v177 int = (v173 << ((917) & 63))
                                                                            var v178 int32 = 2147483647
                                                                            v173 += (^(fi162))
                                                                            _ = v177
                                                                            _ = v178
                                                                        }
                                                                        _ = v173
                                                                    }
                                                                    fmt.Println(fi122)
                                                                } else {
                                                                    switch (fi162 - (fi122 << ((782) & 63))) {
                                                                    case -9223372036854775808:
                                                                        v16 = (0 < 0)
                                                                    case 32767:
                                                                        v179 := GenS14{F0: 127, F1: 32767}
                                                                        _ = v179
                                                                    default:
                                                                        v180 := GenS11{F0: v16}
                                                                        _ = v180
                                                                    }
                                                                    if len(v1) >= 2 {
                                                                        v1 = v1[1:len(v1)]
                                                                    }
                                                                    v16 = v16
                                                                }
                                                                v16 = false
                                                                v181 := GenS5{F0: 0}
                                                                switch fi122 {
                                                                case 920:
                                                                    fmt.Println(fi162, fi162)
                                                                case -348:
                                                                    var v182 uint16 = 980
                                                                    _ = v182
                                                                default:
                                                                    v16 = v16
                                                                }
                                                                v183 := GenS15{F0: 65535}
                                                                _ = v181
                                                                _ = v183
                                                            }
                                                            if v16 {
                                                                v184 := GenS3{F0: -149, F1: 9223372036854775807, F2: 37}
                                                                switch fi162 {
                                                                case -630:
                                                                    if len(v1) > 0 {
                                                                        v1[0] = ""
                                                                    }
                                                                default:
                                                                    switch "\n\t" {
                                                                    case "a":
                                                                        v185 := GenS9{F0: 162, F1: fi122, F2: 18446744073709551615}
                                                                        _ = v185
                                                                    case "":
                                                                        var v186 float64 = ((3.14159 / float64(float32(-2.25))) * (-2.25 + 0.1))
                                                                        _ = v186
                                                                    case "hello":
                                                                        _ = len(v4)
                                                                    }
                                                                }
                                                                _ = v184
                                                            } else {
                                                                if true {
                                                                    v187 := GenS16{F0: 494}
                                                                    if len(v2) > 0 {
                                                                        v2[0] = "a"
                                                                    }
                                                                    _ = v187
                                                                } else {
                                                                    switch (fi162 << ((0) & 63)) {
                                                                    case 42:
                                                                        if (true && false) {
                                                                            switch "" {
                                                                            case "\n\t":
                                                                                if (1 != 4294967295) {
                                                                                    if v16 {
                                                                                        if v16 {
                                                                                            v188 := v4[fi162]
                                                                                            v188 -= 310
                                                                                            v189 := GenS17{F0: 1, F1: 42}
                                                                                            _ = v188
                                                                                            _ = v189
                                                                                        } else {
                                                                                            if v16 {
                                                                                                if len(v2) >= 2 {
                                                                                                    v2 = v2[1:len(v2)]
                                                                                                }
                                                                                                var v190 uint8 = 128
                                                                                                _ = v190
                                                                                            }
                                                                                        }
                                                                                        fmt.Println(fi162)
                                                                                    } else {
                                                                                        if false {
                                                                                            var v191 int8 = 100
                                                                                            var v192 string = ""
                                                                                            _ = v191
                                                                                            _ = v192
                                                                                        }
                                                                                    }
                                                                                } else {
                                                                                    v193 := GenS18{F0: -2147483648, F1: 1}
                                                                                    var v194 int16 = -32768
                                                                                    fmt.Println(fi122, fi122, fi162)
                                                                                    var v195 uint8 = 62
                                                                                    _ = v193
                                                                                    _ = v194
                                                                                    _ = v195
                                                                                }
                                                                            case "a":
                                                                                if (232 == 436) {
                                                                                    if v16 {
                                                                                        fmt.Println(v16)
                                                                                    }
                                                                                    v16 = false
                                                                                }
                                                                            case "hello":
                                                                                if (1 <= 129) {
                                                                                    v16 = (4294967295 < 42)
                                                                                    v196 := GenS8{F0: 4294967295}
                                                                                    switch (-(fi122)) {
                                                                                    case -607:
                                                                                        _ = len(v4)
                                                                                    case 2147483647:
                                                                                        v16 = false
                                                                                    case 320:
                                                                                        var v197 int8 = 124
                                                                                        _ = v197
                                                                                    }
                                                                                    _ = v196
                                                                                } else {
                                                                                    switch fi122 {
                                                                                    case 985:
                                                                                        var v198 int8 = 42
                                                                                        _ = v198
                                                                                    case 127:
                                                                                        var v199 uint32 = 0
                                                                                        _ = v199
                                                                                    case -1:
                                                                                        switch (fi162 << ((4294967295) & 63)) {
                                                                                        case 42:
                                                                                            fmt.Println(fi122)
                                                                                        default:
                                                                                            v4[fi162] = 271
                                                                                        }
                                                                                    }
                                                                                    if len(v1) > 0 {
                                                                                        v1[0] = ("foo bar" + ("foo bar" + ""))
                                                                                    }
                                                                                    if ((("foo bar" <= ("hello" + "foo bar")) && (-128 <= -128)) && true) {
                                                                                        var v200 uint8 = 255
                                                                                        _ = v200
                                                                                    } else {
                                                                                        if (211 <= 882) {
                                                                                            v16 = true
                                                                                            v16 = false
                                                                                            var v201 uint16 = 708
                                                                                            var v202 int8 = 1
                                                                                            var v203 string = ("a" + ("" + "\n\t"))
                                                                                            if len(v2) >= 2 {
                                                                                                v2 = v2[1:len(v2)]
                                                                                            }
                                                                                            if (18446744073709551615 >= 18446744073709551615) {
                                                                                                if v16 {
                                                                                                    v16 = v16
                                                                                                    switch v203 {
                                                                                                    case "a":
                                                                                                        if (27 <= 127) {
                                                                                                            if len(v1) >= 2 {
                                                                                                                v1 = v1[1:len(v1)]
                                                                                                            }
                                                                                                        }
                                                                                                    case "foo bar":
                                                                                                        switch fi162 {
                                                                                                        case -335:
                                                                                                            switch v203 {
                                                                                                            case "foo bar":
                                                                                                                var v204 float32 = (((1.5 / (3.14159 + (1e10 - (1.5 - 1e10)))) - ((1.5 + -2.25) * 0.1)) / ((float32((((((1e10 - (3.14159 / float64((3.14159 / 1.0)))) * (((-2.25 / 0.1) - 0.1) * 1e10)) * float64(float32(0.0))) / float64(((1e10 * (0.1 * 1.5)) * (float32((1.5 + (0.0 - 3.14159))) - (0.1 * 3.14159))))) * (0.1 * (float64((1e10 - 1.5)) * float64(3.14159))))) * 1e10) + -2.25))
                                                                                                                _ = v204
                                                                                                            case "a":
                                                                                                                var v205 float64 = 0.1
                                                                                                                _ = v205
                                                                                                            case "\n\t":
                                                                                                                v206 := GenS3{F0: 42, F1: 1, F2: 4294967295}
                                                                                                                _ = v206
                                                                                                            default:
                                                                                                                var v207 uint = 2
                                                                                                                _ = v207
                                                                                                            }
                                                                                                        }
                                                                                                    case "":
                                                                                                        if ((^(v201)) <= v201) {
                                                                                                            v203 = ("foo bar" + v203)
                                                                                                        } else {
                                                                                                            var v208 float64 = 1.5
                                                                                                            v209 := GenS13{F0: 1}
                                                                                                            var v210 int16 = -779
                                                                                                            _ = v208
                                                                                                            _ = v209
                                                                                                            _ = v210
                                                                                                        }
                                                                                                    }
                                                                                                } else {
                                                                                                    v203 = ((v203 + v203) + v203)
                                                                                                    if len(v1) >= 2 {
                                                                                                        v1 = v1[1:len(v1)]
                                                                                                    }
                                                                                                    var v211 int = fi162
                                                                                                    _ = v211
                                                                                                }
                                                                                                var v212 uint = 0
                                                                                                var v213 uint32 = 42
                                                                                                _ = v212
                                                                                                _ = v213
                                                                                            } else {
                                                                                                if ("" < v203) {
                                                                                                    switch fi162 {
                                                                                                    case 2147483647:
                                                                                                        v202 -= 100
                                                                                                    case -1:
                                                                                                        v203 = v203
                                                                                                    case 42:
                                                                                                        v202 = v202
                                                                                                    }
                                                                                                } else {
                                                                                                    var v214 int16 = 251
                                                                                                    var v215 uint16 = 42
                                                                                                    _ = v214
                                                                                                    _ = v215
                                                                                                }
                                                                                                v203 = ((v203 + (((v203 + (v203 + v203)) + "") + "hello")) + (v203 + v203))
                                                                                                if true {
                                                                                                    var v216 uint8 = 1
                                                                                                    _ = v216
                                                                                                }
                                                                                            }
                                                                                            _ = v201
                                                                                            _ = v202
                                                                                            _ = v203
                                                                                        }
                                                                                        var v217 float32 = -2.25
                                                                                        if true {
                                                                                            v217 -= (3.14159 + (1.5 - (0.0 * v217)))
                                                                                        } else {
                                                                                            v217 -= v217
                                                                                            var v218 int = 2147483647
                                                                                            _ = v218
                                                                                        }
                                                                                        v16 = true
                                                                                        var v219 float64 = -2.25
                                                                                        _ = v217
                                                                                        _ = v219
                                                                                    }
                                                                                }
                                                                                switch "\n\t" {
                                                                                case "a":
                                                                                    switch "hello" {
                                                                                    case "":
                                                                                        v220 := GenS19{F0: 2, F1: genFn1(-54, ((false && (!(((((float32(1.5) - 1e10) - float32((3.14159 * (0.1 + 3.14159)))) - 0.1) == (1.5 / 1.5)) || (!(65535 == 18446744073709551615))))) && v16), 2)}
                                                                                        _ = v220
                                                                                    case "foo bar":
                                                                                        switch "hello" {
                                                                                        case "\n\t":
                                                                                            var v221 uint32 = 42
                                                                                            _ = v221
                                                                                        case "foo bar":
                                                                                            v16 = (0 < 2)
                                                                                        case "a":
                                                                                            delete(v3, "\n\t")
                                                                                            if v16 {
                                                                                                if v16 {
                                                                                                    switch "hello" {
                                                                                                    case "hello":
                                                                                                        var v222 bool = v16
                                                                                                        _ = v222
                                                                                                    case "":
                                                                                                        fmt.Println(fi162, fi162, fi122)
                                                                                                    case "\n\t":
                                                                                                        if (!(int16(fi162) >= 77)) {
                                                                                                            if v16 {
                                                                                                                var v223 int64 = genFn1(-35, false, 952)
                                                                                                                _ = v223
                                                                                                            }
                                                                                                            if (2 == 878) {
                                                                                                                var v224 int64 = 0
                                                                                                                if len(v2) >= 2 {
                                                                                                                    v2 = v2[1:len(v2)]
                                                                                                                }
                                                                                                                _ = v224
                                                                                                            } else {
                                                                                                                switch 128 {
                                                                                                                case -1:
                                                                                                                    if ((-2.25 / float32(1.5)) < (float32(1e10) / (1e10 + float32(0.0)))) {
                                                                                                                        var v225 uint16 = 1
                                                                                                                        v226 := GenS9{F0: 255, F1: (fi122 & 0), F2: 0}
                                                                                                                        if true {
                                                                                                                            var v227 int8 = 0
                                                                                                                            _ = v227
                                                                                                                        } else {
                                                                                                                            v228 := GenS20{F0: false, F1: genFn1(100, (v16 && v16), 32767), F2: (^(v225))}
                                                                                                                            _ = v228
                                                                                                                        }
                                                                                                                        _ = v225
                                                                                                                        _ = v226
                                                                                                                    }
                                                                                                                case 2147483647:
                                                                                                                    delete(v3, "")
                                                                                                                default:
                                                                                                                    var v229 int = fi162
                                                                                                                    _ = v229
                                                                                                                }
                                                                                                            }
                                                                                                        } else {
                                                                                                            fmt.Println(fi162, v16, fi122)
                                                                                                            if v16 {
                                                                                                                if (0 == 0) {
                                                                                                                    fmt.Println(fi122, fi122)
                                                                                                                    if (genFn0() != 964) {
                                                                                                                        v16 = v16
                                                                                                                        v230 := GenS7{F0: 151}
                                                                                                                        _ = v230
                                                                                                                    } else {
                                                                                                                        var v231 int64 = 9223372036854775807
                                                                                                                        _ = v231
                                                                                                                    }
                                                                                                                    if (genFn0() <= 42) {
                                                                                                                        switch fi122 {
                                                                                                                        case 42:
                                                                                                                            if (v16 || (42 != -1)) {
                                                                                                                                var v232 int8 = 77
                                                                                                                                if ((!v16) || v16) {
                                                                                                                                    v233 := v3["\n\t"]
                                                                                                                                    var v234 uint64 = 1
                                                                                                                                    v232 += v232
                                                                                                                                    _ = v233
                                                                                                                                    _ = v234
                                                                                                                                } else {
                                                                                                                                    _ = len(v3)
                                                                                                                                    var v235 uint = 255
                                                                                                                                    _ = v235
                                                                                                                                }
                                                                                                                                var v236 uint32 = 4294967295
                                                                                                                                var v237 int64 = 42
                                                                                                                                _ = v232
                                                                                                                                _ = v236
                                                                                                                                _ = v237
                                                                                                                            } else {
                                                                                                                                v16 = ("a" < "")
                                                                                                                            }
                                                                                                                        case 32767:
                                                                                                                            if v16 {
                                                                                                                                if (v16 && true) {
                                                                                                                                    switch (("hello" + "foo bar") + "foo bar") {
                                                                                                                                    case "a":
                                                                                                                                        var v238 bool = v16
                                                                                                                                        _ = v238
                                                                                                                                    }
                                                                                                                                    switch ("" + "") {
                                                                                                                                    case "\n\t":
                                                                                                                                        if v16 {
                                                                                                                                            if (-329 != -821) {
                                                                                                                                                delete(v3, "hello")
                                                                                                                                                switch fi162 {
                                                                                                                                                case 2147483647:
                                                                                                                                                    v16 = v16
                                                                                                                                                case 32767:
                                                                                                                                                    if len(v2) >= 2 {
                                                                                                                                                        v2 = v2[1:len(v2)]
                                                                                                                                                    }
                                                                                                                                                default:
                                                                                                                                                    fmt.Println(v16, v16, v16)
                                                                                                                                                }
                                                                                                                                            }
                                                                                                                                            fmt.Println(fi122, fi162)
                                                                                                                                        } else {
                                                                                                                                            v4[fi162] = 0
                                                                                                                                            v239 := GenS13{F0: 1}
                                                                                                                                            _ = v239
                                                                                                                                        }
                                                                                                                                    case "foo bar":
                                                                                                                                        var v240 int32 = -1
                                                                                                                                        _ = v240
                                                                                                                                    }
                                                                                                                                    v241 := GenS21{F0: 86}
                                                                                                                                    fmt.Println(fi162)
                                                                                                                                    var v242 uint16 = 544
                                                                                                                                    _ = v241
                                                                                                                                    _ = v242
                                                                                                                                }
                                                                                                                            }
                                                                                                                        }
                                                                                                                        var v243 int16 = 717
                                                                                                                        _ = v243
                                                                                                                    } else {
                                                                                                                        if v16 {
                                                                                                                            v244 := GenS10{F0: 1000}
                                                                                                                            var v245 int16 = -32768
                                                                                                                            _ = v244
                                                                                                                            _ = v245
                                                                                                                        } else {
                                                                                                                            var v246 int8 = -113
                                                                                                                            fmt.Println(v16, v246, fi162)
                                                                                                                            _ = v246
                                                                                                                        }
                                                                                                                        var v247 float64 = ((0.0 - 1.5) / 0.1)
                                                                                                                        var v248 float64 = v247
                                                                                                                        switch "\n\t" {
                                                                                                                        case "hello":
                                                                                                                            delete(v4, fi122)
                                                                                                                        }
                                                                                                                        _ = v247
                                                                                                                        _ = v248
                                                                                                                    }
                                                                                                                }
                                                                                                            } else {
                                                                                                                if (2147483647 != 2) {
                                                                                                                    v16 = v16
                                                                                                                } else {
                                                                                                                    if true {
                                                                                                                        switch (fi162 % (((fi122 + 59)) | 1)) {
                                                                                                                        case 127:
                                                                                                                            v249 := GenS22{F0: 4294967295}
                                                                                                                            _ = v249
                                                                                                                        default:
                                                                                                                            delete(v4, fi162)
                                                                                                                        }
                                                                                                                        v2 = append(v2, ("\n\t" + ((("" + "a") + "\n\t") + (("\n\t" + "hello") + ("a" + "\n\t")))))
                                                                                                                    } else {
                                                                                                                        v16 = v16
                                                                                                                        v16 = v16
                                                                                                                    }
                                                                                                                }
                                                                                                                v250 := GenS12{F0: true}
                                                                                                                _ = v250
                                                                                                            }
                                                                                                        }
                                                                                                    }
                                                                                                    fmt.Println(fi162, fi122)
                                                                                                    switch fi122 {
                                                                                                    case 969:
                                                                                                        var v251 uint16 = 2
                                                                                                        _ = v251
                                                                                                    default:
                                                                                                        fmt.Println(fi122)
                                                                                                    }
                                                                                                } else {
                                                                                                    switch (fi122 | 1) {
                                                                                                    case -513:
                                                                                                        if v16 {
                                                                                                            _ = len(v2)
                                                                                                            var v252 uint32 = 2147483648
                                                                                                            _ = v252
                                                                                                        }
                                                                                                    }
                                                                                                    v253 := GenS16{F0: 1}
                                                                                                    var v254 uint64 = 606
                                                                                                    v1 = append(v1, "\n\t")
                                                                                                    _ = v253
                                                                                                    _ = v254
                                                                                                }
                                                                                            }
                                                                                            switch fi122 {
                                                                                            case 127:
                                                                                                fmt.Println(fi122)
                                                                                                if ((643 == genFn1(33, (0 <= 42), 2)) || (42 != -1)) {
                                                                                                    if true {
                                                                                                        var v255 int16 = -32768
                                                                                                        v255 += v255
                                                                                                        _ = v255
                                                                                                    }
                                                                                                }
                                                                                            case 32767:
                                                                                                switch "a" {
                                                                                                case "a":
                                                                                                    switch ("hello" + (("" + "a") + "a")) {
                                                                                                    case "":
                                                                                                        v16 = false
                                                                                                        if (100 <= -128) {
                                                                                                            v256 := GenS23{F0: 1, F1: 0, F2: 1000}
                                                                                                            if (((v16 || v16) || false) || true) {
                                                                                                                v1 = append(v1, ("a" + ("" + "")))
                                                                                                                if (!(((("foo bar" + "a") + "") + "") == "\n\t")) {
                                                                                                                    v16 = ((-88 >= 1) || (!v16))
                                                                                                                    if (v16 || (v16 && (v16 || (464 != 2147483647)))) {
                                                                                                                        v257 := GenS24{F0: (fi122 | (fi122 % ((fi162) | 1))), F1: -24}
                                                                                                                        _ = v257
                                                                                                                    }
                                                                                                                    if len(v2) >= 2 {
                                                                                                                        v2 = v2[1:len(v2)]
                                                                                                                    }
                                                                                                                } else {
                                                                                                                    v16 = (0 != genFn0())
                                                                                                                    var v258 string = (("foo bar" + "a") + "a")
                                                                                                                    v258 = (((v258 + "hello") + "foo bar") + ((v258 + v258) + ("" + ((("hello" + v258) + (v258 + (v258 + v258))) + v258))))
                                                                                                                    _ = v258
                                                                                                                }
                                                                                                                v259 := GenS22{F0: 255}
                                                                                                                _ = v259
                                                                                                            }
                                                                                                            _ = v256
                                                                                                        }
                                                                                                        var v260 int = fi162
                                                                                                        if v16 {
                                                                                                            var v261 uint = 255
                                                                                                            switch 875 {
                                                                                                            case -2147483648:
                                                                                                                fmt.Println(fi122)
                                                                                                            case 32767:
                                                                                                                switch "" {
                                                                                                                case "\n\t":
                                                                                                                    var v262 bool = (true || (!(v16 && (0.0 != (1.5 + 0.1)))))
                                                                                                                    _ = v262
                                                                                                                }
                                                                                                            case 9223372036854775807:
                                                                                                                fmt.Println(v260, v260)
                                                                                                            default:
                                                                                                                if (124 != 553) {
                                                                                                                    switch fi122 {
                                                                                                                    case 42:
                                                                                                                        switch "\n\t" {
                                                                                                                        case "a":
                                                                                                                            v260 = 1
                                                                                                                        default:
                                                                                                                            v260 = 2147483647
                                                                                                                        }
                                                                                                                    case -2147483648:
                                                                                                                        v16 = true
                                                                                                                        fmt.Println(v261)
                                                                                                                        v263 := GenS25{F0: 42}
                                                                                                                        v261 = v261
                                                                                                                        _ = v263
                                                                                                                    }
                                                                                                                    v264 := v4[fi122]
                                                                                                                    var v265 int32 = 1
                                                                                                                    _ = v264
                                                                                                                    _ = v265
                                                                                                                } else {
                                                                                                                    delete(v4, v260)
                                                                                                                    if (2 <= 1) {
                                                                                                                        if (86 != 127) {
                                                                                                                            v16 = ((("" + (((("foo bar" + "hello") + "foo bar") + "hello") + "foo bar")) + (("" + "foo bar") + "hello")) >= "a")
                                                                                                                            v261 += 810
                                                                                                                            var v266 int32 = 257
                                                                                                                            v267 := GenS26{F0: 18446744073709551615, F1: v261, F2: 255}
                                                                                                                            switch (("foo bar" + "a") + "\n\t") {
                                                                                                                            case "hello":
                                                                                                                                var v268 int64 = 0
                                                                                                                                _ = v268
                                                                                                                            default:
                                                                                                                                v261 = (v261 | v261)
                                                                                                                                _ = len(v3)
                                                                                                                            }
                                                                                                                            _ = v266
                                                                                                                            _ = v267
                                                                                                                        } else {
                                                                                                                            v16 = (v261 > (v261 ^ v261))
                                                                                                                            v269 := GenS27{F0: (1.5 > ((1e10 / 1e10) * ((((float64((1e10 - 0.0)) + (-2.25 + ((3.14159 + -2.25) / (3.14159 / (3.14159 * -2.25))))) + 1e10) + 1e10) - 3.14159))), F1: false, F2: 1}
                                                                                                                            _ = v269
                                                                                                                        }
                                                                                                                        if (genFn1(54, true, 47) != -744) {
                                                                                                                            var v270 uint = (v261 / (((^(v261))) | 1))
                                                                                                                            _ = v270
                                                                                                                        }
                                                                                                                    }
                                                                                                                    v271 := GenS28{F0: 127, F1: genFn1(-128, false, 42), F2: 42}
                                                                                                                    var v272 uint32 = 4294967295
                                                                                                                    var v273 uint8 = 255
                                                                                                                    switch ("hello" + "\n\t") {
                                                                                                                    case "foo bar":
                                                                                                                        v4[fi162] = 13
                                                                                                                    case "hello":
                                                                                                                        v273 -= v273
                                                                                                                    default:
                                                                                                                        if (!true) {
                                                                                                                            var v274 string = ("" + ("a" + "foo bar"))
                                                                                                                            _ = v274
                                                                                                                        }
                                                                                                                    }
                                                                                                                    switch fi162 {
                                                                                                                    case -1:
                                                                                                                        v272 -= v272
                                                                                                                    case 784:
                                                                                                                        v275 := GenS29{F0: v272, F1: 42}
                                                                                                                        _ = v275
                                                                                                                    }
                                                                                                                    if ("\n\t" == ((("\n\t" + ("hello" + "")) + (("" + "a") + "hello")) + "foo bar")) {
                                                                                                                        var v276 int32 = -167
                                                                                                                        v276 = v276
                                                                                                                        v16 = ((42 < 2) && ((3.14159 / 3.14159) <= float32((-2.25 * -2.25))))
                                                                                                                        v260 = -505
                                                                                                                        _ = len(v3)
                                                                                                                        _ = v276
                                                                                                                    } else {
                                                                                                                        v277 := GenS8{F0: 523}
                                                                                                                        _ = v277
                                                                                                                    }
                                                                                                                    _ = v271
                                                                                                                    _ = v272
                                                                                                                    _ = v273
                                                                                                                }
                                                                                                            }
                                                                                                            _ = len(v1)
                                                                                                            _ = v261
                                                                                                        } else {
                                                                                                            if v16 {
                                                                                                                switch (^(v260)) {
                                                                                                                case 0:
                                                                                                                    switch ("hello" + "\n\t") {
                                                                                                                    case "hello":
                                                                                                                        var v278 uint32 = 792
                                                                                                                        _ = v278
                                                                                                                    default:
                                                                                                                        if v16 {
                                                                                                                            v279 := GenS12{F0: ((1 > 742) || (0 >= 2))}
                                                                                                                            v16 = true
                                                                                                                            v260 = (fi122 >> ((957) & 63))
                                                                                                                            v260 += (fi162 >> ((0) & 63))
                                                                                                                            switch "foo bar" {
                                                                                                                            case "a":
                                                                                                                                v280 := GenS1{F0: "hello"}
                                                                                                                                _ = v280
                                                                                                                            case "foo bar":
                                                                                                                                v4[fi122] = 65535
                                                                                                                            case "hello":
                                                                                                                                switch "" {
                                                                                                                                case "hello":
                                                                                                                                    v260 = v260
                                                                                                                                case "\n\t":
                                                                                                                                    _ = len(v4)
                                                                                                                                case "":
                                                                                                                                    v281 := GenS30{F0: 42, F1: 65535, F2: 18446744073709551615}
                                                                                                                                    _ = v281
                                                                                                                                default:
                                                                                                                                    var v282 int = (v260 >> ((555) & 63))
                                                                                                                                    _ = v282
                                                                                                                                }
                                                                                                                                var v283 int8 = 100
                                                                                                                                _ = v283
                                                                                                                            default:
                                                                                                                                v284 := GenS13{F0: 539}
                                                                                                                                _ = v284
                                                                                                                            }
                                                                                                                            _ = v279
                                                                                                                        }
                                                                                                                    }
                                                                                                                case 32767:
                                                                                                                    var v285 string = ("foo bar" + "")
                                                                                                                    _ = v285
                                                                                                                case 800:
                                                                                                                    v286 := GenS31{F0: (fi122 + 9223372036854775807), F1: -369}
                                                                                                                    _ = v286
                                                                                                                }
                                                                                                            } else {
                                                                                                                var v287 int = (v260 - (fi122 ^ (-(fi122))))
                                                                                                                v288 := GenS32{F0: 2}
                                                                                                                v287 += 4294967295
                                                                                                                _ = v287
                                                                                                                _ = v288
                                                                                                            }
                                                                                                            delete(v4, -2147483648)
                                                                                                            v260 = (-(v260))
                                                                                                        }
                                                                                                        v289 := GenS17{F0: -1, F1: 42}
                                                                                                        _ = v260
                                                                                                        _ = v289
                                                                                                    case "foo bar":
                                                                                                        if ((1e10 + -2.25) == 3.14159) {
                                                                                                            v290 := GenS13{F0: 604}
                                                                                                            _ = v290
                                                                                                        }
                                                                                                        v16 = (v16 || (!(-1 < 827)))
                                                                                                    case "\n\t":
                                                                                                        switch ("a" + "hello") {
                                                                                                        case "foo bar":
                                                                                                            var v291 bool = true
                                                                                                            _ = v291
                                                                                                        case "":
                                                                                                            var v292 uint64 = 2
                                                                                                            _ = v292
                                                                                                        case "a":
                                                                                                            v293 := GenS33{F0: 42}
                                                                                                            _ = v293
                                                                                                        }
                                                                                                    }
                                                                                                case "\n\t":
                                                                                                    if false {
                                                                                                        if false {
                                                                                                            if v16 {
                                                                                                                if len(v1) >= 2 {
                                                                                                                    v1 = v1[1:len(v1)]
                                                                                                                }
                                                                                                            } else {
                                                                                                                var v294 float64 = 1.5
                                                                                                                v294 = v294
                                                                                                                var v295 uint32 = 456
                                                                                                                _ = v294
                                                                                                                _ = v295
                                                                                                            }
                                                                                                        }
                                                                                                        switch ("a" + (("" + ("a" + ("hello" + "hello"))) + ("a" + "\n\t"))) {
                                                                                                        case "hello":
                                                                                                            var v296 float64 = float64(1.5)
                                                                                                            _ = v296
                                                                                                        case "":
                                                                                                            v16 = (0 <= 8)
                                                                                                        }
                                                                                                        var v297 string = ((("foo bar" + ("\n\t" + "\n\t")) + "hello") + "\n\t")
                                                                                                        _ = v297
                                                                                                    }
                                                                                                case "hello":
                                                                                                    switch (fi162 | fi122) {
                                                                                                    case -389:
                                                                                                        switch 32767 {
                                                                                                        case 2147483647:
                                                                                                            v16 = v16
                                                                                                        default:
                                                                                                            delete(v3, "")
                                                                                                        }
                                                                                                    case 797:
                                                                                                        v298 := GenS16{F0: 2147483647}
                                                                                                        _ = v298
                                                                                                    case 42:
                                                                                                        var v299 int64 = 0
                                                                                                        _ = v299
                                                                                                    default:
                                                                                                        if v16 {
                                                                                                            delete(v3, ("\n\t" + ""))
                                                                                                        } else {
                                                                                                            v16 = (v16 || v16)
                                                                                                            var v300 uint32 = 1
                                                                                                            _ = v300
                                                                                                        }
                                                                                                    }
                                                                                                    v16 = v16
                                                                                                    if v16 {
                                                                                                        v16 = true
                                                                                                        v16 = (1.5 <= (-2.25 - -2.25))
                                                                                                        v301 := GenS19{F0: 65535, F1: 922}
                                                                                                        switch "hello" {
                                                                                                        case "\n\t":
                                                                                                            v16 = v16
                                                                                                        default:
                                                                                                            v16 = ("hello" == "hello")
                                                                                                            v16 = v16
                                                                                                        }
                                                                                                        _ = v301
                                                                                                    }
                                                                                                }
                                                                                            case -937:
                                                                                                v302 := GenS34{F0: (255 <= 0), F1: -2147483648, F2: 2}
                                                                                                _ = v302
                                                                                            }
                                                                                        default:
                                                                                            v303 := GenS1{F0: ("" + "a")}
                                                                                            _ = v303
                                                                                        }
                                                                                    }
                                                                                default:
                                                                                    v16 = (fi162 >= (fi122 << ((255) & 63)))
                                                                                }
                                                                                var v304 uint32 = 2147483648
                                                                                _ = v304
                                                                            }
                                                                            v305 := v3["a"]
                                                                            v306 := GenS19{F0: 2, F1: 236}
                                                                            _ = len(v4)
                                                                            _ = v305
                                                                            _ = v306
                                                                        } else {
                                                                            v16 = (1 == -870)
                                                                            v16 = v16
                                                                            var v307 uint64 = genFn0()
                                                                            if false {
                                                                                v16 = v16
                                                                                if v16 {
                                                                                    var v308 int = 571
                                                                                    v309 := GenS33{F0: 32767}
                                                                                    if (0.0 >= (1.5 / 1.0)) {
                                                                                        switch "hello" {
                                                                                        case "foo bar":
                                                                                            var v310 int64 = 0
                                                                                            _ = v310
                                                                                        case "":
                                                                                            v16 = v16
                                                                                        default:
                                                                                            switch (("a" + "foo bar") + (("a" + "a") + ((("a" + ("" + "a")) + "hello") + "a"))) {
                                                                                            case "hello":
                                                                                                var v311 int16 = 1000
                                                                                                _ = v311
                                                                                            case "foo bar":
                                                                                                var v312 float64 = 0.1
                                                                                                _ = v312
                                                                                            case "\n\t":
                                                                                                var v313 int = fi122
                                                                                                _ = v313
                                                                                            }
                                                                                        }
                                                                                        switch "" {
                                                                                        case "\n\t":
                                                                                            var v314 bool = (!v16)
                                                                                            _ = v314
                                                                                        case "hello":
                                                                                            var v315 int16 = 615
                                                                                            _ = v315
                                                                                        default:
                                                                                            _ = len(v1)
                                                                                        }
                                                                                    } else {
                                                                                        fmt.Println(fi162, v16, v16)
                                                                                        if len(v2) > 0 {
                                                                                            v2[0] = "\n\t"
                                                                                        }
                                                                                        var v316 int16 = 344
                                                                                        v317 := GenS26{F0: 232, F1: 0, F2: 2}
                                                                                        fmt.Println(fi122, fi122)
                                                                                        v318 := v4[491]
                                                                                        _ = v316
                                                                                        _ = v317
                                                                                        _ = v318
                                                                                    }
                                                                                    v307 = genFn0()
                                                                                    var v319 uint16 = 42
                                                                                    v308 -= 42
                                                                                    v308 = fi122
                                                                                    _ = v308
                                                                                    _ = v309
                                                                                    _ = v319
                                                                                }
                                                                                v320 := GenS35{F0: 1, F1: 65535, F2: -7}
                                                                                _ = v320
                                                                            } else {
                                                                                if (v16 || (fi122 > (fi162 - (fi162 << ((65535) & 63))))) {
                                                                                    v307 = 1
                                                                                    v321 := GenS1{F0: ((("hello" + "a") + "foo bar") + "a")}
                                                                                    var v322 uint64 = v307
                                                                                    _ = v321
                                                                                    _ = v322
                                                                                } else {
                                                                                    switch fi162 {
                                                                                    case 32768:
                                                                                        if v16 {
                                                                                            if ((((!(958 < -1)) && true) || v16) || v16) {
                                                                                                if ((3.14159 >= -2.25) && ((fi162 + fi122) >= (-(fi162)))) {
                                                                                                    if v16 {
                                                                                                        v323 := GenS36{F0: "a"}
                                                                                                        switch (fi162 + (fi162 - (fi122 * (fi122 * fi162)))) {
                                                                                                        case -85:
                                                                                                            switch fi162 {
                                                                                                            case 0:
                                                                                                                var v324 uint32 = 42
                                                                                                                _ = v324
                                                                                                            default:
                                                                                                                if v16 {
                                                                                                                    fmt.Println(v16, fi162, v307)
                                                                                                                    if ((fi162 << ((4294967295) & 63)) <= (-(fi122))) {
                                                                                                                        var v325 bool = v16
                                                                                                                        v3["foo bar"] = "hello"
                                                                                                                        _ = v325
                                                                                                                    } else {
                                                                                                                        v16 = v16
                                                                                                                        var v326 int32 = -1
                                                                                                                        v327 := v3[(("" + "hello") + ("foo bar" + "hello"))]
                                                                                                                        v328 := GenS2{F0: 0, F1: (fi122 >> ((255) & 63)), F2: ((v327 + "\n\t") + (v327 + v327))}
                                                                                                                        if (((75 != 241) && v16) && (9223372036854775807 <= 59)) {
                                                                                                                            _ = len(v2)
                                                                                                                            if len(v1) > 0 {
                                                                                                                                v1[0] = v327
                                                                                                                            }
                                                                                                                            fmt.Println(v307)
                                                                                                                        } else {
                                                                                                                            var v329 uint64 = (v307 << ((2) & 63))
                                                                                                                            var v330 int32 = (v326 >> ((2) & 63))
                                                                                                                            _ = v329
                                                                                                                            _ = v330
                                                                                                                        }
                                                                                                                        v331 := GenS37{F0: (v307 & (v307 | (v307 * v307)))}
                                                                                                                        _ = v326
                                                                                                                        _ = v327
                                                                                                                        _ = v328
                                                                                                                        _ = v331
                                                                                                                    }
                                                                                                                    if v16 {
                                                                                                                        if true {
                                                                                                                            if (42 != 255) {
                                                                                                                                v16 = true
                                                                                                                                v307 -= (^(v307))
                                                                                                                                _ = len(v2)
                                                                                                                            }
                                                                                                                            v16 = v16
                                                                                                                            switch (fi122 * fi162) {
                                                                                                                            case 42:
                                                                                                                                v332 := GenS38{F0: 65535, F1: 1, F2: 2}
                                                                                                                                _ = v332
                                                                                                                            case -9223372036854775808:
                                                                                                                                fmt.Println(v16, v307)
                                                                                                                            case 128:
                                                                                                                                v333 := GenS12{F0: v16}
                                                                                                                                _ = v333
                                                                                                                            default:
                                                                                                                                v334 := v4[fi122]
                                                                                                                                _ = v334
                                                                                                                            }
                                                                                                                            v16 = (v16 && true)
                                                                                                                        }
                                                                                                                        v335 := GenS20{F0: false, F1: 1, F2: 42}
                                                                                                                        _ = v335
                                                                                                                    } else {
                                                                                                                        v16 = (42 >= v307)
                                                                                                                        v336 := GenS29{F0: 65535, F1: 32767}
                                                                                                                        if len(v2) >= 2 {
                                                                                                                            v2 = v2[1:len(v2)]
                                                                                                                        }
                                                                                                                        _ = v336
                                                                                                                    }
                                                                                                                } else {
                                                                                                                    v16 = (-32768 < 496)
                                                                                                                    switch fi162 {
                                                                                                                    case 4294967295:
                                                                                                                        v16 = (1 < 93)
                                                                                                                        if len(v2) > 0 {
                                                                                                                            v2[0] = ""
                                                                                                                        }
                                                                                                                    case 1:
                                                                                                                        v337 := GenS39{F0: 18446744073709551615, F1: 42, F2: ""}
                                                                                                                        _ = v337
                                                                                                                    case -2147483648:
                                                                                                                        v1 = append(v1, "foo bar")
                                                                                                                        switch "a" {
                                                                                                                        case "a":
                                                                                                                            if (("foo bar" != ("" + ("a" + "hello"))) && v16) {
                                                                                                                                v338 := GenS25{F0: 4294967295}
                                                                                                                                switch "hello" {
                                                                                                                                case "":
                                                                                                                                    var v339 bool = v16
                                                                                                                                    _ = v339
                                                                                                                                default:
                                                                                                                                    fmt.Println(v307, fi162)
                                                                                                                                }
                                                                                                                                var v340 uint32 = 42
                                                                                                                                _ = v338
                                                                                                                                _ = v340
                                                                                                                            } else {
                                                                                                                                if len(v1) > 0 {
                                                                                                                                    v1[0] = "\n\t"
                                                                                                                                }
                                                                                                                            }
                                                                                                                        case "foo bar":
                                                                                                                            v16 = (1 >= 251)
                                                                                                                            var v341 float32 = ((((0.0 * float32(0.1)) * -2.25) * (1e10 + float32(((float64((1e10 + 3.14159)) + 1e10) + 0.1)))) + 1e10)
                                                                                                                            v342 := GenS40{F0: "a"}
                                                                                                                            _ = v341
                                                                                                                            _ = v342
                                                                                                                        }
                                                                                                                    default:
                                                                                                                        v16 = v16
                                                                                                                    }
                                                                                                                    switch "foo bar" {
                                                                                                                    case "\n\t":
                                                                                                                        if len(v1) > 0 {
                                                                                                                            v1[0] = ""
                                                                                                                        }
                                                                                                                    case "foo bar":
                                                                                                                        var v343 uint64 = 83
                                                                                                                        _ = v343
                                                                                                                    case "":
                                                                                                                        if (!(!(-9223372036854775808 >= 402))) {
                                                                                                                            var v344 int8 = 1
                                                                                                                            v344 -= v344
                                                                                                                            _ = v344
                                                                                                                        }
                                                                                                                    default:
                                                                                                                        var v345 uint32 = 444
                                                                                                                        _ = v345
                                                                                                                    }
                                                                                                                }
                                                                                                            }
                                                                                                        case 0:
                                                                                                            var v346 uint8 = 2
                                                                                                            _ = v346
                                                                                                        case 1:
                                                                                                            if len(v1) > 0 {
                                                                                                                v1[0] = ""
                                                                                                            }
                                                                                                        default:
                                                                                                            v347 := GenS33{F0: 737}
                                                                                                            _ = v347
                                                                                                        }
                                                                                                        _ = v323
                                                                                                    }
                                                                                                    v307 = (v307 ^ v307)
                                                                                                    if len(v1) >= 2 {
                                                                                                        v1 = v1[1:len(v1)]
                                                                                                    }
                                                                                                }
                                                                                            } else {
                                                                                                var v348 bool = v16
                                                                                                if (fi122 < fi122) {
                                                                                                    var v349 uint32 = 1
                                                                                                    _ = v349
                                                                                                } else {
                                                                                                    if v16 {
                                                                                                        var v350 string = "foo bar"
                                                                                                        v307 += 47
                                                                                                        v351 := v3[v350]
                                                                                                        _ = v350
                                                                                                        _ = v351
                                                                                                    }
                                                                                                    v352 := GenS41{F0: 255, F1: 128, F2: -957}
                                                                                                    v16 = v16
                                                                                                    var v353 int16 = -416
                                                                                                    _ = v352
                                                                                                    _ = v353
                                                                                                }
                                                                                                v354 := GenS10{F0: 42}
                                                                                                v307 = genFn0()
                                                                                                _ = v348
                                                                                                _ = v354
                                                                                            }
                                                                                            v1 = append(v1, "\n\t")
                                                                                            var v355 uint16 = 1
                                                                                            var v356 int8 = -1
                                                                                            _ = v355
                                                                                            _ = v356
                                                                                        } else {
                                                                                            if len(v1) > 0 {
                                                                                                v1[0] = "a"
                                                                                            }
                                                                                        }
                                                                                    case 42:
                                                                                        switch fi122 {
                                                                                        case 9223372036854775807:
                                                                                            if ((-1 <= 0) && v16) {
                                                                                                var v357 int8 = 100
                                                                                                if (128 == 255) {
                                                                                                    fmt.Println(fi162, fi162)
                                                                                                    if false {
                                                                                                        v358 := GenS28{F0: 0, F1: 1, F2: -1}
                                                                                                        var v359 uint = 42
                                                                                                        _ = v358
                                                                                                        _ = v359
                                                                                                    } else {
                                                                                                        if (!true) {
                                                                                                            v360 := GenS36{F0: ("\n\t" + "hello")}
                                                                                                            v1 = append(v1, "\n\t")
                                                                                                            fmt.Println(v307)
                                                                                                            _ = v360
                                                                                                        }
                                                                                                    }
                                                                                                    if false {
                                                                                                        v16 = v16
                                                                                                    } else {
                                                                                                        if (349 <= 65535) {
                                                                                                            v361 := GenS40{F0: ("a" + "")}
                                                                                                            var v362 uint64 = 0
                                                                                                            _ = v361
                                                                                                            _ = v362
                                                                                                        }
                                                                                                        fmt.Println(v357, v307, fi162)
                                                                                                        v363 := GenS42{F0: 0, F1: ((42 != 4294967295) && v16)}
                                                                                                        _ = v363
                                                                                                    }
                                                                                                    if true {
                                                                                                        if len(v2) > 0 {
                                                                                                            v2[0] = "foo bar"
                                                                                                        }
                                                                                                        var v364 float32 = -2.25
                                                                                                        _ = len(v1)
                                                                                                        switch (("" + "a") + "a") {
                                                                                                        case "hello":
                                                                                                            if v16 {
                                                                                                                v307 = v307
                                                                                                            } else {
                                                                                                                var v365 uint32 = 2147483647
                                                                                                                var v366 int8 = (v357 % (((v357 - (v357 << ((4294967295) & 63)))) | 1))
                                                                                                                v367 := GenS3{F0: 32768, F1: -9223372036854775808, F2: 291}
                                                                                                                _ = v365
                                                                                                                _ = v366
                                                                                                                _ = v367
                                                                                                            }
                                                                                                        default:
                                                                                                            switch "foo bar" {
                                                                                                            case "":
                                                                                                                if v16 {
                                                                                                                    fmt.Println(v307, v16, fi162)
                                                                                                                    if (101 > v357) {
                                                                                                                        fmt.Println(fi122)
                                                                                                                    }
                                                                                                                    var v368 int16 = int16(v307)
                                                                                                                    _ = v368
                                                                                                                } else {
                                                                                                                    if v16 {
                                                                                                                        switch 42 {
                                                                                                                        case -496:
                                                                                                                            var v369 int64 = genFn1((v357 % ((v357) | 1)), (65535 >= 70), 0)
                                                                                                                            _ = v369
                                                                                                                        default:
                                                                                                                            var v370 uint8 = 0
                                                                                                                            _ = v370
                                                                                                                        }
                                                                                                                    }
                                                                                                                    v371 := GenS36{F0: ""}
                                                                                                                    var v372 int = (fi122 << ((213) & 63))
                                                                                                                    var v373 int16 = 32767
                                                                                                                    _ = v371
                                                                                                                    _ = v372
                                                                                                                    _ = v373
                                                                                                                }
                                                                                                            case "foo bar":
                                                                                                                var v374 uint64 = v307
                                                                                                                _ = v374
                                                                                                            default:
                                                                                                                v375 := GenS17{F0: 42, F1: -2147483648}
                                                                                                                _ = v375
                                                                                                            }
                                                                                                        }
                                                                                                        var v376 uint = 4294967295
                                                                                                        _ = v364
                                                                                                        _ = v376
                                                                                                    }
                                                                                                }
                                                                                                var v377 uint64 = (v307 - (v307 >> ((470) & 63)))
                                                                                                _ = v357
                                                                                                _ = v377
                                                                                            } else {
                                                                                                v16 = v16
                                                                                                v378 := GenS18{F0: -307, F1: 0}
                                                                                                fmt.Println(fi122, v16)
                                                                                                _ = v378
                                                                                            }
                                                                                        case 128:
                                                                                            v16 = false
                                                                                        case 322:
                                                                                            var v379 int16 = 0
                                                                                            _ = v379
                                                                                        default:
                                                                                            var v380 uint16 = 178
                                                                                            _ = v380
                                                                                        }
                                                                                        v307 += (v307 - v307)
                                                                                    case -2147483648:
                                                                                        switch (fi162 % ((fi162) | 1)) {
                                                                                        case 76:
                                                                                            var v381 int32 = 1
                                                                                            _ = v381
                                                                                        default:
                                                                                            v16 = false
                                                                                        }
                                                                                    }
                                                                                    fmt.Println(fi122, v16, v16)
                                                                                }
                                                                                if ((float32((1e10 - 0.0)) <= 1.5) && (true || false)) {
                                                                                    v382 := GenS43{F0: -511, F1: 42}
                                                                                    _ = v382
                                                                                } else {
                                                                                    v383 := GenS12{F0: true}
                                                                                    switch 267 {
                                                                                    case 32767:
                                                                                        v16 = (!(v16 && true))
                                                                                    }
                                                                                    _ = v383
                                                                                }
                                                                                var v384 uint16 = 65535
                                                                                _ = v384
                                                                            }
                                                                            v385 := GenS7{F0: (v307 ^ 1)}
                                                                            _ = v307
                                                                            _ = v385
                                                                        }
                                                                    case 4294967295:
                                                                        delete(v3, (("foo bar" + "\n\t") + "a"))
                                                                    case -154:
                                                                        v386 := GenS25{F0: 4294967295}
                                                                        _ = v386
                                                                    }
                                                                    if (42 <= 921) {
                                                                        v1 = append(v1, ((("" + "hello") + "foo bar") + (("\n\t" + ("a" + "a")) + "hello")))
                                                                        v16 = ((100 > 1) || v16)
                                                                    }
                                                                }
                                                            }
                                                            v387 := v3["a"]
                                                            _ = v387
                                                        } else {
                                                            if v16 {
                                                                switch "hello" {
                                                                case "hello":
                                                                    v388 := GenS44{F0: v16, F1: genFn1(1, v16, 737)}
                                                                    var v389 int = 245
                                                                    var v390 uint64 = 2
                                                                    v390 = genFn0()
                                                                    _ = v388
                                                                    _ = v389
                                                                    _ = v390
                                                                default:
                                                                    v391 := GenS38{F0: 881, F1: 4294967295, F2: 65535}
                                                                    _ = v391
                                                                }
                                                            }
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                            switch fi162 {
                                                            case -640:
                                                                v16 = v16
                                                            case 32767:
                                                                switch (^(fi122)) {
                                                                case 42:
                                                                    v392 := v3["\n\t"]
                                                                    _ = v392
                                                                case -858:
                                                                    var v393 int32 = 603
                                                                    _ = v393
                                                                }
                                                            case -2147483648:
                                                                fmt.Println(fi122, fi162)
                                                                if (1 != 2147483647) {
                                                                    v394 := GenS45{F0: genFn1(42, v16, 120)}
                                                                    var v395 uint16 = 151
                                                                    _ = v394
                                                                    _ = v395
                                                                }
                                                                var v396 uint32 = 42
                                                                _ = v396
                                                            }
                                                        }
                                                        var v397 uint8 = 1
                                                        _ = v397
                                                    }
                                                } else {
                                                    var v398 float64 = float64((1.5 / ((1.5 / 1.0) / 1.5)))
                                                    _ = v398
                                                }
                                            default:
                                                switch ("hello" + "hello") {
                                                case "hello":
                                                    var v399 int8 = -1
                                                    _ = v399
                                                }
                                            }
                                        } else {
                                            var v400 int = fi122
                                            var v401 int8 = 127
                                            v400 += fi162
                                            var v402 int16 = 1000
                                            _ = v400
                                            _ = v401
                                            _ = v402
                                        }
                                        var v403 float32 = 1.5
                                        if v16 {
                                            switch 128 {
                                            case -210:
                                                var v404 uint16 = 255
                                                _ = v404
                                            case 2147483647:
                                                if (394 < -1) {
                                                    fmt.Println(v403, fi122)
                                                    var v405 string = "foo bar"
                                                    switch (v405 + "\n\t") {
                                                    case "a":
                                                        v406 := v3[(((v405 + (v405 + "a")) + v405) + ((v405 + v405) + "\n\t"))]
                                                        if false {
                                                            if v16 {
                                                                var v407 uint16 = 71
                                                                var v408 float32 = v403
                                                                _ = v407
                                                                _ = v408
                                                            }
                                                            var v409 uint = 2
                                                            switch ((("a" + v405) + v406) + v405) {
                                                            case "a":
                                                                var v410 bool = (!(42 <= -78))
                                                                _ = v410
                                                            case "hello":
                                                                var v411 bool = v16
                                                                _ = v411
                                                            case "\n\t":
                                                                if true {
                                                                    var v412 float64 = -2.25
                                                                    v16 = (!(!(42 >= 68)))
                                                                    if v16 {
                                                                        v409 = v409
                                                                        v413 := GenS46{F0: 1}
                                                                        _ = v413
                                                                    }
                                                                    _ = v412
                                                                } else {
                                                                    v406 = ("hello" + "")
                                                                    var v414 int32 = -2147483648
                                                                    v415 := GenS9{F0: (v409 & (^(v409))), F1: fi162, F2: v409}
                                                                    _ = v414
                                                                    _ = v415
                                                                }
                                                            }
                                                            _ = v409
                                                        }
                                                        var v416 uint32 = 55
                                                        _ = v406
                                                        _ = v416
                                                    default:
                                                        var v417 float32 = v403
                                                        _ = v417
                                                    }
                                                    _ = v405
                                                } else {
                                                    var v418 uint16 = 255
                                                    var v419 int = (fi122 | 128)
                                                    v420 := GenS47{F0: 175, F1: (v419 / ((fi162) | 1)), F2: (("\n\t" + "hello") + ("hello" + (("foo bar" + ("\n\t" + "")) + ("hello" + "a"))))}
                                                    _ = v418
                                                    _ = v419
                                                    _ = v420
                                                }
                                            }
                                            v421 := GenS48{F0: -61}
                                            if false {
                                                v422 := GenS17{F0: genFn1(42, ((2 == 32767) && (42 == 1)), 593), F1: -1}
                                                v423 := GenS2{F0: 2, F1: (fi122 % ((fi162) | 1)), F2: (("a" + ("foo bar" + (("a" + ("\n\t" + "\n\t")) + "hello"))) + "")}
                                                if true {
                                                    var v424 float64 = (-2.25 * 3.14159)
                                                    if len(v1) > 0 {
                                                        v1[0] = ("foo bar" + (("foo bar" + "a") + "\n\t"))
                                                    }
                                                    _ = len(v3)
                                                    _ = v424
                                                }
                                                _ = v422
                                                _ = v423
                                            } else {
                                                v403 = ((v403 + 1e10) * v403)
                                                switch fi162 {
                                                case 128:
                                                    v425 := GenS28{F0: 1, F1: 9223372036854775807, F2: -1}
                                                    _ = v425
                                                case 2147483647:
                                                    _ = len(v3)
                                                case 42:
                                                    v426 := v3[""]
                                                    _ = v426
                                                }
                                                v403 = 3.14159
                                                if true {
                                                    if (0 <= 255) {
                                                        var v427 uint16 = 42
                                                        _ = v427
                                                    }
                                                    v403 = 3.14159
                                                    v403 = 0.1
                                                } else {
                                                    var v428 uint64 = 4294967295
                                                    _ = len(v2)
                                                    v428 += v428
                                                    _ = v428
                                                }
                                                v429 := v3[""]
                                                _ = v429
                                            }
                                            _ = v421
                                        }
                                        _ = v403
                                    }
                                    fmt.Println(v16, v16, fi122)
                                    _ = fi162
                                }
                                for fi430 := 0; fi430 < 3; fi430++ {
                                    if (572 > 4294967295) {
                                        var v431 uint64 = 2
                                        _ = v431
                                    }
                                    var v432 uint64 = 42
                                    _ = fi430
                                    _ = v432
                                }
                            }
                        default:
                            v16 = v16
                        }
                        v433 := GenS25{F0: 855}
                        v4[42] = 736
                        _ = v433
                    }
                }
            }
            _ = fi122
            _ = v123
        }
    }
    var v434 float32 = (3.14159 - 0.0)
    switch "\n\t" {
    case "foo bar":
        for fi435 := 0; fi435 < 12; fi435++ {
            v436 := GenS17{F0: 887, F1: -2147483648}
            _ = fi435
            _ = v436
        }
    case "\n\t":
        if (v16 || true) {
            var v437 int64 = 909
            if (true || (1 == -32768)) {
                for fi438 := 0; fi438 < 2; fi438++ {
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                    if (!false) {
                        for fi439 := 0; fi439 < 12; fi439++ {
                            fmt.Println(fi439)
                            var v440 int64 = (v437 << ((421) & 63))
                            v434 = (-2.25 - (((v434 * v434) - 0.1) + (v434 + 1.5)))
                            _ = len(v2)
                            _ = fi439
                            _ = v440
                        }
                        v16 = false
                    } else {
                        v441 := GenS19{F0: 744, F1: v437}
                        _ = v441
                    }
                    _ = fi438
                }
                var v442 string = ("" + ("foo bar" + ("a" + "hello")))
                v443 := GenS49{F0: 905, F1: true}
                _ = v442
                _ = v443
            }
            _ = v437
        } else {
            v444 := GenS23{F0: 100, F1: -123, F2: 459}
            fmt.Println(v434)
            fmt.Println(v16, v434)
            _ = v444
        }
    }
    fmt.Println(v434, v434)
    _ = len(v4)
    switch 2147483647 {
    case -2147483648:
        v434 = ((((v434 / float32(0.1)) - 0.1) + float32(1.5)) - ((3.14159 / 1.5) + (-2.25 / v434)))
    default:
        v16 = false
    }
    if false {
        if false {
            var v445 int = 0
            var v446 int = (v445 * (v445 >> ((255) & 63)))
            _ = v445
            _ = v446
        }
        var v447 uint8 = 255
        delete(v4, -511)
        _ = v447
    }
    v448 := GenS50{F0: 9223372036854775807}
    v16 = (!((271 < 65535) && v16))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v434:", v434)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v434
    _ = v448
}
