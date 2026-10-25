package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 uint8 }
type GenS2 struct { F0 int; F1 int; F2 int32 }
type GenS3 struct { F0 int; F1 string }
type GenS4 struct { F0 bool; F1 uint16; F2 uint32 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 string; F1 uint8; F2 string }
type GenS7 struct { F0 uint; F1 string; F2 uint }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 int32; F1 int }
type GenS10 struct { F0 int16; F1 int8 }
type GenS11 struct { F0 uint; F1 uint8; F2 uint }
type GenS12 struct { F0 uint16; F1 int64; F2 int16 }
type GenS13 struct { F0 bool; F1 uint64; F2 int8 }
type GenS14 struct { F0 int64; F1 int16; F2 uint64 }
func genFn0(p0 uint64, p1 uint, p2 int32) int32 {
    v6 := GenS1{F0: (p1 + p1), F1: p0, F2: 255}
    p2 += 2147483647
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    return p2
}
func genFn1(p0 int8) int8 {
    var v7 uint64 = 0
    switch 32767 {
    case 1:
        v8 := GenS2{F0: -98, F1: -1, F2: -1}
        _ = v8
    case 4294967295:
        for fi9 := 0; fi9 < 2; fi9++ {
            var v10 uint = 2
            v11 := GenS1{F0: 42, F1: 479, F2: 255}
            switch (fi9 % (((fi9 % ((fi9) | 1))) | 1)) {
            case 9223372036854775807:
                var v12 int = 0
                _ = v12
            case 649:
                for fi13 := 0; fi13 < 3; fi13++ {
                    var v14 uint8 = 128
                    _ = fi13
                    _ = v14
                }
            default:
                var v15 uint64 = v7
                _ = v15
            }
            _ = fi9
            _ = v10
            _ = v11
        }
    default:
        if false {
            v7 -= v7
            var v16 uint32 = 4294967295
            _ = v16
        }
    }
    v17 := []int{}
    fmt.Println(p0, v7)
    _ = p0
    _ = v17
    _ = v7
    return 42
}

func main() {
    var v1 uint = 358
    var v2 int32 = -602
    v3 := make(map[string]bool)
    var v4 uint32 = 0
    var v5 int = 1
    fmt.Println(v5)
    v5 += 0
    switch "a" {
    case "a":
        if false {
            v18 := GenS2{F0: 4294967295, F1: 2147483647, F2: 0}
            _ = v18
        } else {
            v3[("" + "foo bar")] = (false && true)
            v2 = genFn0(2, (^(v1)), v2)
            for fi19 := 0; fi19 < 5; fi19++ {
                var v20 uint16 = 42
                _ = fi19
                _ = v20
            }
        }
    case "":
        switch "foo bar" {
        case "a":
            delete(v3, "")
        default:
            fmt.Println(v5, v4, v2)
        }
    case "\n\t":
        if (v2 <= (v2 << (((^(v1))) & 63))) {
            switch "hello" {
            case "a":
                v5 = (v5 / ((v5) | 1))
            case "foo bar":
                fmt.Println(v4)
            case "\n\t":
                switch 0 {
                case 4294967295:
                    v21 := GenS3{F0: v5, F1: "\n\t"}
                    _ = v21
                case -2147483648:
                    var v22 uint64 = 151
                    _ = v22
                default:
                    switch (("hello" + "") + ("hello" + "foo bar")) {
                    case "hello":
                        fmt.Println(v5, v1, v2)
                    default:
                        v4 = (^(v4))
                    }
                }
            default:
                v23 := []int{128, v5, -785}
                _ = v23
            }
            v1 -= v1
        } else {
            var v24 int8 = 100
            var v25 uint32 = uint32(v1)
            for fi26 := 0; fi26 < 2; fi26++ {
                if true {
                    v25 = 2
                } else {
                    switch (("hello" + ("foo bar" + "foo bar")) + ("hello" + ("" + ("hello" + "\n\t")))) {
                    case "\n\t":
                        if (-182 == -1) {
                            var v27 int32 = (^(v2))
                            v28 := GenS2{F0: v5, F1: int(v27), F2: v27}
                            fmt.Println(v25, v1, fi26)
                            _ = v27
                            _ = v28
                        } else {
                            var v29 float64 = (3.14159 - 0.0)
                            if false {
                                v30 := []int{(v5 ^ -393), 42, v5}
                                if len(v30) > 0 {
                                    v30[0] = -9223372036854775808
                                }
                                var v31 uint64 = 0
                                _ = v30
                                _ = v31
                            }
                            _ = v29
                        }
                    case "a":
                        if true {
                            v2 = (v2 * (-(v2)))
                            v2 = genFn0(0, 42, v2)
                        }
                    }
                    for fi32 := 0; fi32 < 18; fi32++ {
                        if (0 == 733) {
                            v33 := GenS4{F0: (!false), F1: 65535, F2: (v25 / ((4294967295) | 1))}
                            _ = v33
                        } else {
                            v2 += 2147483647
                            v34 := GenS5{F0: 18446744073709551615}
                            _ = v34
                        }
                        v35 := GenS3{F0: v5, F1: ((("\n\t" + ("\n\t" + "a")) + ((("" + "foo bar") + ("foo bar" + "a")) + (("hello" + "foo bar") + ("foo bar" + ("a" + "a"))))) + "foo bar")}
                        var v36 int16 = 32767
                        _ = fi32
                        _ = v35
                        _ = v36
                    }
                    if ((!(1 != 847)) && false) {
                        for fi37 := 0; fi37 < 16; fi37++ {
                            switch ("foo bar" + "") {
                            case "":
                                if false {
                                    v4 = v4
                                }
                                fmt.Println(fi37, fi37, v25)
                                v38 := []int{}
                                if (false || false) {
                                    switch (v5 * (fi26 / ((fi37) | 1))) {
                                    case 32767:
                                        var v39 uint8 = 255
                                        _ = v39
                                    case -9223372036854775808:
                                        var v40 float32 = 0.1
                                        _ = v40
                                    default:
                                        if true {
                                            if (false || true) {
                                                switch (v5 * (v5 ^ fi26)) {
                                                case 908:
                                                    var v41 uint16 = 188
                                                    _ = v41
                                                }
                                                delete(v3, (("hello" + "\n\t") + "hello"))
                                                var v42 int16 = 1000
                                                _ = v42
                                            }
                                        } else {
                                            v43 := GenS6{F0: "foo bar", F1: 29, F2: ("" + ("" + "a"))}
                                            _ = v43
                                        }
                                    }
                                    v25 += 0
                                    v3[(("foo bar" + ("" + ("\n\t" + "a"))) + (("a" + ("foo bar" + ("foo bar" + "hello"))) + (("foo bar" + "foo bar") + "")))] = true
                                    v44 := GenS4{F0: false, F1: 32767, F2: v4}
                                    _ = v44
                                }
                                v5 += (v5 - fi26)
                                _ = v38
                            case "\n\t":
                                var v45 uint = 0
                                _ = len(v3)
                                if true {
                                    var v46 int32 = -2147483648
                                    v2 = (v2 - v46)
                                    v45 += 76
                                    var v47 int64 = -1
                                    switch (fi26 - 128) {
                                    case 9223372036854775807:
                                        if (32767 >= 622) {
                                            if true {
                                                v48 := []int{}
                                                _ = v48
                                            } else {
                                                var v49 uint16 = 178
                                                if (-32768 < int16(fi26)) {
                                                    var v50 uint32 = v4
                                                    v3["foo bar"] = true
                                                    fmt.Println(v46)
                                                    fmt.Println(v24, v2, v1)
                                                    _ = v50
                                                }
                                                v45 += (^(v1))
                                                _ = v49
                                            }
                                        } else {
                                            var v51 uint16 = 225
                                            switch "foo bar" {
                                            case "foo bar":
                                                v24 -= v24
                                            case "":
                                                if false {
                                                    var v52 string = ""
                                                    v25 -= v25
                                                    v47 -= (v47 / ((42) | 1))
                                                    _ = v52
                                                } else {
                                                    if (1000 != int16(v45)) {
                                                        fmt.Println(v1)
                                                        delete(v3, "\n\t")
                                                        if (v1 != v45) {
                                                            var v53 int16 = -300
                                                            var v54 int = (v5 >> ((255) & 63))
                                                            v55 := []int{(v54 ^ (v54 ^ (fi37 - 1))), v5}
                                                            _ = v53
                                                            _ = v54
                                                            _ = v55
                                                        } else {
                                                            var v56 float64 = (0.0 / float64(1.5))
                                                            v57 := GenS3{F0: v5, F1: "a"}
                                                            _ = v56
                                                            _ = v57
                                                        }
                                                    }
                                                }
                                            case "hello":
                                                v25 = v4
                                            default:
                                                switch ("a" + "hello") {
                                                case "foo bar":
                                                    fmt.Println(v47, fi37)
                                                }
                                            }
                                            switch 841 {
                                            case 598:
                                                v58 := []int{(^(fi37)), (^(v5))}
                                                _ = v58
                                            default:
                                                if (false || false) {
                                                    var v59 int8 = int8(v51)
                                                    if false {
                                                        if (uint8(v51) == 158) {
                                                            var v60 int = (fi26 + (v5 << (((^(v1))) & 63)))
                                                            v61 := []int{}
                                                            var v62 uint8 = 0
                                                            if (!false) {
                                                                v1 = (v1 | (v1 >> ((463) & 63)))
                                                                v2 += -921
                                                                v63 := GenS7{F0: v45, F1: ("" + ("\n\t" + "")), F2: (v1 - (v45 ^ v45))}
                                                                _ = v63
                                                            }
                                                            _ = v60
                                                            _ = v61
                                                            _ = v62
                                                        } else {
                                                            fmt.Println(v59, v59, v4)
                                                            switch fi26 {
                                                            case 32768:
                                                                v64 := GenS2{F0: fi37, F1: 128, F2: 42}
                                                                _ = v64
                                                            case -1:
                                                                fmt.Println(v5)
                                                            default:
                                                                if (!(false && (578 < v1))) {
                                                                    if false {
                                                                        fmt.Println(v4, v46, v47)
                                                                    }
                                                                    if true {
                                                                        if (v59 == (-(v59))) {
                                                                            var v65 uint = v45
                                                                            _ = v65
                                                                        }
                                                                    }
                                                                    v66 := []int{}
                                                                    _ = v66
                                                                }
                                                                switch 4294967295 {
                                                                case 42:
                                                                    fmt.Println(v2, v24)
                                                                case 2147483647:
                                                                    v45 -= (v45 & (v1 * 144))
                                                                case 0:
                                                                    v45 -= v1
                                                                default:
                                                                    if true {
                                                                        v67 := []int{32768, (v5 - 1)}
                                                                        v45 = v45
                                                                        _ = v67
                                                                    } else {
                                                                        fmt.Println(v45, v47, v45)
                                                                        var v68 float64 = (1e10 / -2.25)
                                                                        v1 = (^(v1))
                                                                        _ = v68
                                                                    }
                                                                }
                                                            }
                                                            if (!true) {
                                                                switch (-(fi26)) {
                                                                case 4294967295:
                                                                    v59 += (v59 + int8(v4))
                                                                case -1:
                                                                    if true {
                                                                        v69 := GenS8{F0: (v47 % ((42) | 1))}
                                                                        _ = v69
                                                                    }
                                                                case -915:
                                                                    fmt.Println(v45, v45, v4)
                                                                }
                                                            }
                                                        }
                                                        v1 -= 255
                                                        v47 += 1
                                                        v70 := GenS4{F0: (false || false), F1: (v51 << ((v1) & 63)), F2: 0}
                                                        _ = v70
                                                    }
                                                    _ = v59
                                                }
                                            }
                                            _ = v51
                                        }
                                    }
                                    var v71 float64 = (-2.25 - -2.25)
                                    if false {
                                        var v72 uint8 = 2
                                        switch "" {
                                        case "\n\t":
                                            v73 := GenS7{F0: (v1 | v1), F1: "\n\t", F2: 526}
                                            _ = v73
                                        default:
                                            v74 := GenS5{F0: 694}
                                            _ = v74
                                        }
                                        _ = v72
                                    } else {
                                        switch ("a" + "a") {
                                        case "hello":
                                            v75 := GenS9{F0: 2147483647, F1: 128}
                                            _ = v75
                                        case "":
                                            v5 += fi37
                                        default:
                                            switch (-(fi26)) {
                                            case 127:
                                                v76 := GenS10{F0: 789, F1: v24}
                                                _ = v76
                                            }
                                        }
                                        if ((float32(3.14159) - (float32((v71 - 3.14159)) / (3.14159 * -2.25))) >= (-2.25 - (1e10 - (-2.25 + 3.14159)))) {
                                            if true {
                                                v4 -= 184
                                            }
                                        } else {
                                            v5 = (v5 - -752)
                                            v77 := v3[("foo bar" + "a")]
                                            _ = v77
                                        }
                                    }
                                    switch v5 {
                                    case 4294967295:
                                        var v78 bool = ((39 <= uint8(v25)) || ((-(v47)) <= (v47 % (((v47 / (((v47 + (v47 / ((42) | 1)))) | 1))) | 1))))
                                        v47 = (v47 | (^(v47)))
                                        if ((2 >= 356) && v78) {
                                            switch "\n\t" {
                                            case "a":
                                                v79 := GenS11{F0: 18446744073709551615, F1: 42, F2: v1}
                                                _ = v79
                                            case "\n\t":
                                                if v78 {
                                                    var v80 uint = v45
                                                    v81 := GenS9{F0: (^(v46)), F1: (^(v5))}
                                                    v82 := GenS12{F0: 2, F1: 441, F2: 1}
                                                    _ = v80
                                                    _ = v81
                                                    _ = v82
                                                }
                                            default:
                                                var v83 int8 = genFn1(0)
                                                _ = v83
                                            }
                                            var v84 uint32 = (v25 << ((v1) & 63))
                                            var v85 uint8 = 1
                                            _ = v84
                                            _ = v85
                                        }
                                        _ = v78
                                    }
                                    _ = v46
                                    _ = v47
                                    _ = v71
                                }
                                _ = v45
                            default:
                                delete(v3, ("\n\t" + ""))
                            }
                            _ = fi37
                        }
                    }
                }
                if true {
                    if false {
                        var v86 uint = v1
                        _ = v86
                    } else {
                        v87 := []int{fi26, fi26}
                        if (1000 == 1) {
                            if len(v87) >= 2 {
                                v87 = v87[1:len(v87)]
                            }
                        } else {
                            if false {
                                delete(v3, "\n\t")
                            } else {
                                var v88 int32 = genFn0(2, 0, v2)
                                _ = len(v3)
                                fmt.Println(v25, v88, v2)
                                _ = v88
                            }
                        }
                        _ = v87
                    }
                }
                var v89 int16 = 716
                _ = fi26
                _ = v89
            }
            _ = v24
            _ = v25
        }
    }
    switch (^(v5)) {
    case 1:
        v4 = v4
    case -160:
        if true {
            _ = len(v3)
        }
    case 9223372036854775807:
        var v90 int16 = 467
        _ = v90
    default:
        var v91 float32 = -2.25
        _ = v91
    }
    v5 -= v5
    v92 := []int{(v5 & (v5 / (((v5 % (((v5 << ((v1) & 63))) | 1))) | 1))), 128, v5, -2147483648}
    v3["\n\t"] = ((v5 | 32767) == (v5 >> ((v1) & 63)))
    if len(v92) >= 2 {
        v92 = v92[1:len(v92)]
    }
    if len(v92) > 0 {
        v92[0] = (v5 >> (((v1 * 220)) & 63))
    }
    if (0 > -19) {
        if (!true) {
            var v93 uint16 = 255
            _ = len(v92)
            _ = v93
        } else {
            v92 = append(v92, (v5 ^ 4294967295))
            v1 = v1
        }
        var v94 uint16 = 42
        switch "\n\t" {
        case "\n\t":
            if (-345 > 114) {
                fmt.Println(v94)
            } else {
                v2 += 0
                v95 := GenS12{F0: (v94 << ((42) & 63)), F1: 9223372036854775807, F2: -842}
                for fi96 := 0; fi96 < 7; fi96++ {
                    if (!true) {
                        v94 -= (v94 + v94)
                    }
                    _ = fi96
                }
                _ = v95
            }
        case "a":
            if len(v92) > 0 {
                v92[0] = -403
            }
        default:
            v97 := GenS8{F0: -710}
            _ = v97
        }
        _ = v94
    }
    for fi98 := 0; fi98 < 17; fi98++ {
        if (!((true || true) || true)) {
            for fi99 := 0; fi99 < 3; fi99++ {
                delete(v3, "\n\t")
                v100 := GenS13{F0: ((v2 != v2) && true), F1: 70, F2: -1}
                if false {
                    v4 = (v4 << ((v1) & 63))
                    fmt.Println(fi98, fi99, fi98)
                }
                var v101 float32 = 0.0
                v92 = append(v92, (fi99 >> (((v1 >> (((v1 + v1)) & 63))) & 63)))
                _ = fi99
                _ = v100
                _ = v101
            }
            if (true || ((!(1 == 18446744073709551615)) && (-223 > -1))) {
                var v102 int32 = -1
                if false {
                    v103 := GenS14{F0: 9223372036854775807, F1: 634, F2: 2}
                    _ = v103
                } else {
                    var v104 float32 = 1.5
                    _ = v104
                }
                _ = v102
            } else {
                v105 := GenS2{F0: (fi98 + (^(fi98))), F1: v5, F2: v2}
                var v106 uint16 = 406
                _ = v105
                _ = v106
            }
        }
        fmt.Println(fi98, v2, v4)
        _ = fi98
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v92:", len(v92), v92)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v92
}
