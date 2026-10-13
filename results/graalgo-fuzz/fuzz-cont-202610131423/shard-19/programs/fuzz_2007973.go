package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint64; F1 int16; F2 int32 }
type GenS3 struct { F0 uint64; F1 string }
type GenS4 struct { F0 int64; F1 int8 }
type GenS5 struct { F0 int8; F1 string }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int8; F1 int64; F2 int16 }
type GenS9 struct { F0 bool; F1 uint16; F2 int64 }
type GenS10 struct { F0 uint16 }
type GenS11 struct { F0 int64; F1 uint64 }
type GenS12 struct { F0 int; F1 string; F2 int32 }
type GenS13 struct { F0 bool; F1 uint8 }
type GenS14 struct { F0 int32; F1 bool }
type GenS15 struct { F0 string; F1 int; F2 int16 }
type GenS16 struct { F0 uint32; F1 uint32 }
func genFn0(p0 int8) int16 {
    if false {
        v5 := GenS1{F0: "hello"}
        var v6 uint64 = 4294967295
        _ = v5
        _ = v6
    }
    switch ("a" + "hello") {
    case "hello":
        p0 -= 0
    case "a":
        for fi7 := 0; fi7 < 18; fi7++ {
            var v8 uint8 = 127
            v9 := []int{244, (^(fi7))}
            _ = fi7
            _ = v8
            _ = v9
        }
    case "\n\t":
        var v10 int8 = (p0 | (p0 >> ((18446744073709551615) & 63)))
        _ = v10
    }
    if (false || ((true && (-1 <= -1)) || (!(4294967295 <= 675)))) {
        if true {
            var v11 bool = (true && (721 >= 0))
            switch "a" {
            case "hello":
                p0 = p0
            }
            _ = v11
        }
        if false {
            fmt.Println(p0)
            v12 := []int{42, 150, 169, 9223372036854775807}
            v12 = append(v12, -22)
            _ = v12
        }
    }
    for fi13 := 0; fi13 < 12; fi13++ {
        var v14 float64 = (float64(1e10) / 1.5)
        if true {
            var v15 int32 = -2147483648
            v16 := GenS2{F0: 1, F1: 1000, F2: v15}
            v14 += ((v14 + -2.25) / (3.14159 - 0.1))
            _ = v15
            _ = v16
        } else {
            v17 := GenS1{F0: ("" + "\n\t")}
            _ = v17
        }
        v18 := []int{128}
        _ = fi13
        _ = v14
        _ = v18
    }
    _ = p0
    return 1000
}
func genFn1(p0 uint64, p1 uint16) int16 {
    var v19 int8 = -43
    for fi20 := 0; fi20 < 5; fi20++ {
        v21 := GenS1{F0: ("foo bar" + ("hello" + "a"))}
        _ = fi20
        _ = v21
    }
    switch 32768 {
    case 2147483647:
        v22 := []int{42}
        _ = v22
    case 9223372036854775807:
        v23 := []int{776, 4294967295, 32767, 2147483647}
        _ = v23
    default:
        v19 -= v19
    }
    _ = p0
    _ = p1
    _ = v19
    return 316
}

func main() {
    var v1 int8 = 44
    var v2 bool = (18446744073709551615 == 530)
    v3 := make(map[int]int)
    var v4 int16 = 0
    var v24 bool = true
    for fi25 := 0; fi25 < 15; fi25++ {
        v26 := GenS3{F0: 754, F1: ("\n\t" + "a")}
        delete(v3, (fi25 + fi25))
        _ = fi25
        _ = v26
    }
    for fi27 := 0; fi27 < 2; fi27++ {
        fmt.Println(v1)
        _ = fi27
    }
    if false {
        if v2 {
            var v28 uint64 = 380
            _ = v28
        } else {
            v2 = v2
            fmt.Println(v4, v24)
        }
    } else {
        _ = len(v3)
        fmt.Println(v1, v4, v1)
    }
    v2 = true
    for fi29 := 0; fi29 < 7; fi29++ {
        switch (fi29 >> ((4294967295) & 63)) {
        case 32768:
            var v30 uint16 = 2
            _ = v30
        case 127:
            v31 := GenS4{F0: -1, F1: v1}
            _ = v31
        }
        var v32 int16 = 57
        if v2 {
            switch fi29 {
            case -17:
                if v2 {
                    for fi33 := 0; fi33 < 10; fi33++ {
                        v34 := GenS5{F0: (v1 - v1), F1: "\n\t"}
                        v35 := GenS5{F0: (v1 | 1), F1: (("foo bar" + "foo bar") + ("foo bar" + ""))}
                        switch 32767 {
                        case 128:
                            var v36 int64 = 42
                            _ = v36
                        case -138:
                            v2 = v24
                        case 1:
                            if (0 >= 427) {
                                var v37 uint64 = 192
                                var v38 uint16 = 844
                                _ = v37
                                _ = v38
                            } else {
                                v39 := GenS6{F0: 125}
                                v40 := GenS1{F0: (("a" + "hello") + (("\n\t" + (("" + ("foo bar" + "")) + "foo bar")) + ("foo bar" + "\n\t")))}
                                if (0 >= 42) {
                                    if v2 {
                                        v41 := v3[(fi29 / ((-379) | 1))]
                                        v2 = (v2 && true)
                                        v24 = v2
                                        _ = v41
                                    } else {
                                        v1 = v1
                                        fmt.Println(v2, v4, v32)
                                        if (true && (1 > 1)) {
                                            _ = len(v3)
                                        } else {
                                            if ((!v24) || true) {
                                                fmt.Println(v2)
                                                v42 := []int{(-(fi29)), (^(fi33)), (fi33 ^ -556)}
                                                var v43 int = (fi29 >> ((771) & 63))
                                                _ = v42
                                                _ = v43
                                            } else {
                                                v44 := []int{(-(fi29)), (-(fi33)), fi29}
                                                fmt.Println(v4, v4)
                                                _ = v44
                                            }
                                            v32 += v32
                                            v2 = false
                                        }
                                        if (!v2) {
                                            fmt.Println(v1, fi33, v1)
                                            _ = len(v3)
                                        } else {
                                            v4 -= v32
                                            if (((127 > 0) && v24) && v2) {
                                                switch fi33 {
                                                case 287:
                                                    v2 = v2
                                                default:
                                                    if v2 {
                                                        if (v2 || (0 != 1)) {
                                                            v45 := []int{}
                                                            v46 := GenS4{F0: 0, F1: v1}
                                                            _ = v45
                                                            _ = v46
                                                        }
                                                        fmt.Println(fi33, v2, fi33)
                                                        var v47 int64 = 0
                                                        _ = v47
                                                    } else {
                                                        switch fi29 {
                                                        case -2147483648:
                                                            v4 = (v32 & 65)
                                                        case 32768:
                                                            if (255 == 96) {
                                                                switch (fi33 >> ((255) & 63)) {
                                                                case -2147483648:
                                                                    fmt.Println(v4, fi33, v32)
                                                                case -9223372036854775808:
                                                                    v2 = ((2 >= 4294967295) || ((!(4294967295 <= 0)) && (-2.25 == (0.0 - 3.14159))))
                                                                case 768:
                                                                    if v2 {
                                                                        if v2 {
                                                                            v1 = v1
                                                                            switch -459 {
                                                                            case -2147483648:
                                                                                switch (fi33 & (fi29 << ((315) & 63))) {
                                                                                case 4294967295:
                                                                                    fmt.Println(fi29)
                                                                                case 127:
                                                                                    v48 := GenS2{F0: 4294967295, F1: v4, F2: 0}
                                                                                    _ = v48
                                                                                }
                                                                            case -69:
                                                                                var v49 float64 = 0.0
                                                                                _ = v49
                                                                            default:
                                                                                switch "\n\t" {
                                                                                case "":
                                                                                    var v50 int16 = (v32 / ((genFn1(93, 2)) | 1))
                                                                                    _ = v50
                                                                                case "a":
                                                                                    v2 = (2 < 687)
                                                                                }
                                                                                var v51 int8 = (^(v1))
                                                                                _ = v51
                                                                            }
                                                                            v32 -= genFn0(v1)
                                                                            if (!(0.0 != 3.14159)) {
                                                                                v52 := []int{fi29, (fi33 << ((0) & 63))}
                                                                                var v53 bool = v2
                                                                                v52 = append(v52, int(v1))
                                                                                _ = v52
                                                                                _ = v53
                                                                            }
                                                                            v1 -= (v1 - 5)
                                                                        } else {
                                                                            v1 -= (v1 >> ((1) & 63))
                                                                            v54 := []int{-606, (fi33 >> ((0) & 63))}
                                                                            if false {
                                                                                v54 = append(v54, 743)
                                                                                v54 = append(v54, fi29)
                                                                                var v55 uint8 = 82
                                                                                _ = v55
                                                                            } else {
                                                                                var v56 uint64 = 0
                                                                                var v57 uint16 = 32767
                                                                                v58 := GenS3{F0: uint64(v57), F1: ""}
                                                                                _ = v56
                                                                                _ = v57
                                                                                _ = v58
                                                                            }
                                                                            _ = v54
                                                                        }
                                                                    }
                                                                }
                                                                var v59 int64 = 42
                                                                v24 = (0 >= 0)
                                                                _ = v59
                                                            }
                                                        }
                                                        v24 = v2
                                                    }
                                                }
                                                if v24 {
                                                    v60 := GenS7{F0: "\n\t"}
                                                    if ("\n\t" <= ("a" + ("a" + ""))) {
                                                        var v61 string = (("foo bar" + "foo bar") + ("\n\t" + "hello"))
                                                        switch (^(fi29)) {
                                                        case 127:
                                                            fmt.Println(fi33)
                                                        case 945:
                                                            switch fi29 {
                                                            case 4294967295:
                                                                var v62 float64 = (-2.25 + 0.0)
                                                                _ = v62
                                                            }
                                                        default:
                                                            v63 := GenS7{F0: v61}
                                                            _ = v63
                                                        }
                                                        _ = v61
                                                    }
                                                    _ = v60
                                                } else {
                                                    v64 := GenS8{F0: (v1 % (((v1 | (v1 + (-(v1))))) | 1)), F1: 0, F2: genFn1(18446744073709551615, 2)}
                                                    var v65 uint8 = 42
                                                    _ = v64
                                                    _ = v65
                                                }
                                            } else {
                                                if true {
                                                    var v66 int64 = 42
                                                    v67 := GenS9{F0: v24, F1: 853, F2: v66}
                                                    v68 := GenS10{F0: 715}
                                                    _ = v66
                                                    _ = v67
                                                    _ = v68
                                                }
                                                v4 += (v4 / (((v32 + v4)) | 1))
                                                switch "" {
                                                case "hello":
                                                    v69 := []int{fi33}
                                                    _ = v69
                                                default:
                                                    v24 = false
                                                }
                                            }
                                        }
                                    }
                                    v70 := []int{fi29, (-(fi33)), (-(fi33)), fi29}
                                    var v71 int = fi33
                                    var v72 int8 = 100
                                    v72 -= (v72 & v1)
                                    _ = v70
                                    _ = v71
                                    _ = v72
                                }
                                _ = v39
                                _ = v40
                            }
                        default:
                            v24 = true
                        }
                        _ = fi33
                        _ = v34
                        _ = v35
                    }
                }
            default:
                if true {
                    v2 = v24
                    switch "\n\t" {
                    case "\n\t":
                        if false {
                            fmt.Println(v2)
                            fmt.Println(v1, v32)
                            v24 = v2
                        } else {
                            v24 = (32767 < (v4 ^ (v4 + v32)))
                            switch ("foo bar" + ("" + "a")) {
                            case "\n\t":
                                switch "hello" {
                                case "hello":
                                    v32 = v4
                                case "\n\t":
                                    fmt.Println(fi29, v1, v24)
                                }
                            case "hello":
                                switch (("hello" + "foo bar") + "foo bar") {
                                case "hello":
                                    v24 = (false && (956 >= 42))
                                case "a":
                                    v73 := GenS11{F0: -9223372036854775808, F1: 4294967295}
                                    _ = v73
                                default:
                                    v2 = true
                                }
                            case "":
                                fmt.Println(fi29, fi29)
                            }
                        }
                    case "foo bar":
                        v1 -= (^(v1))
                    }
                    var v74 int8 = -1
                    _ = v74
                } else {
                    v4 -= v4
                    v1 = (v1 ^ (v1 - (v1 + (v1 ^ (^(v1))))))
                    v4 = v32
                }
            }
            for fi75 := 0; fi75 < 10; fi75++ {
                var v76 uint8 = 76
                _ = fi75
                _ = v76
            }
            if (65535 <= 986) {
                fmt.Println(v24, v2, v32)
                for fi77 := 0; fi77 < 15; fi77++ {
                    if (0 > (-(v1))) {
                        v24 = (fi77 < fi77)
                        v78 := GenS12{F0: 4294967295, F1: "hello", F2: -1}
                        var v79 int32 = 0
                        _ = v78
                        _ = v79
                    } else {
                        v2 = (426 >= 2)
                        var v80 uint32 = 2
                        switch fi77 {
                        case -2147483648:
                            switch (fi77 * fi77) {
                            case 32767:
                                delete(v3, 1)
                            case 1:
                                v4 += (v32 / (((^(v4))) | 1))
                                var v81 float64 = 1.5
                                _ = v81
                            case -9223372036854775808:
                                v24 = true
                                var v82 float32 = (0.0 + 1e10)
                                _ = v82
                            default:
                                var v83 uint32 = 1
                                _ = v83
                            }
                        case 9223372036854775807:
                            if v24 {
                                switch ("hello" + "") {
                                case "\n\t":
                                    v24 = true
                                }
                            }
                        default:
                            var v84 int64 = 0
                            _ = v84
                        }
                        _ = v80
                    }
                    _ = fi77
                }
            } else {
                for fi85 := 0; fi85 < 8; fi85++ {
                    var v86 uint16 = 2
                    v1 += (^(v1))
                    _ = fi85
                    _ = v86
                }
                if v2 {
                    var v87 uint16 = 1
                    fmt.Println(v2)
                    if v24 {
                        var v88 uint = 0
                        _ = v88
                    } else {
                        v3[(fi29 << ((790) & 63))] = fi29
                        switch ("\n\t" + "foo bar") {
                        case "foo bar":
                            v89 := GenS13{F0: ((304 > 2) && (v24 || v2)), F1: 102}
                            _ = v89
                        case "":
                            var v90 int8 = v1
                            _ = v90
                        default:
                            v91 := GenS14{F0: 2147483647, F1: (9223372036854775807 == -720)}
                            _ = v91
                        }
                        if v24 {
                            v1 = (v1 % ((v1) | 1))
                            if false {
                                v32 = genFn1(504, (^(v87)))
                                v2 = (0 <= 127)
                            } else {
                                v92 := GenS3{F0: 4294967295, F1: ("" + ("foo bar" + (("\n\t" + "\n\t") + "foo bar")))}
                                v2 = v2
                                switch (("foo bar" + "a") + (("foo bar" + "foo bar") + ("foo bar" + ("a" + "hello")))) {
                                case "\n\t":
                                    if (!v2) {
                                        v93 := GenS8{F0: v1, F1: 9223372036854775807, F2: v32}
                                        if v2 {
                                            for fi94 := 0; fi94 < 9; fi94++ {
                                                v95 := GenS7{F0: (("foo bar" + "a") + "hello")}
                                                if v2 {
                                                    v2 = (!("a" > ""))
                                                }
                                                _ = fi94
                                                _ = v95
                                            }
                                            var v96 uint8 = 127
                                            _ = v96
                                        }
                                        _ = v93
                                    } else {
                                        switch fi29 {
                                        case -76:
                                            v1 -= 0
                                        default:
                                            fmt.Println(v87, v2, v24)
                                        }
                                        v97 := GenS15{F0: "\n\t", F1: fi29, F2: -32768}
                                        v1 = 0
                                        _ = v97
                                    }
                                case "":
                                    v24 = (v2 || (!v24))
                                case "foo bar":
                                    v32 = genFn1(42, (v87 ^ (v87 & v87)))
                                }
                                _ = v92
                            }
                        } else {
                            fmt.Println(v1, v1)
                        }
                    }
                    _ = v87
                }
            }
        } else {
            if (true && (!(18446744073709551615 >= 1))) {
                v98 := GenS16{F0: 4294967295, F1: 65535}
                _ = v98
            }
        }
        _ = fi29
        _ = v32
    }
    for fi99 := 0; fi99 < 8; fi99++ {
        v100 := []int{(fi99 >> ((18446744073709551615) & 63))}
        var v101 int64 = 0
        if v2 {
            for fi102 := 0; fi102 < 15; fi102++ {
                _ = len(v100)
                _ = fi102
            }
            v24 = ((3.14159 / ((1e10 / 1.5) * 0.1)) == 0.1)
        } else {
            var v103 uint16 = 255
            _ = v103
        }
        _ = fi99
        _ = v100
        _ = v101
    }
    switch "foo bar" {
    case "":
        _ = len(v3)
    case "foo bar":
        var v104 int64 = 22
        _ = v104
    default:
        v105 := []int{482, 9223372036854775807, 4294967295}
        _ = v105
    }
    var v106 float32 = 0.0
    switch (("a" + "foo bar") + (("" + ("" + ("hello" + ("hello" + (("" + "\n\t") + ("a" + "a")))))) + ("" + "foo bar"))) {
    case "a":
        if v24 {
            v107 := GenS10{F0: 65535}
            var v108 uint8 = 0
            fmt.Println(v4, v4, v1)
            _ = v107
            _ = v108
        }
    default:
        fmt.Println(v24, v4, v1)
    }
    var v109 int16 = v4
    fmt.Println(v106, v1, v109)
    v2 = v24
    delete(v3, 9223372036854775807)
    if (((v1 & (^(v1))) == v1) || false) {
        v2 = (2147483647 >= 111)
        var v110 uint8 = 202
        _ = v110
    }
    v111 := GenS6{F0: 689}
    fmt.Println("v1:", v1)
    fmt.Println("v106:", v106)
    fmt.Println("v109:", v109)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v106
    _ = v109
    _ = v111
    _ = v2
    _ = v24
    _ = v3
    _ = v4
}
