package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint16; F1 string }
type GenS4 struct { F0 int }
type GenS5 struct { F0 uint; F1 int64 }
type GenS6 struct { F0 uint; F1 uint16 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 int16; F1 int64; F2 int16 }
type GenS10 struct { F0 bool; F1 int; F2 int32 }
type GenS11 struct { F0 bool; F1 int16; F2 uint8 }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 string; F1 uint32 }
type GenS14 struct { F0 int8; F1 int32 }
type GenS15 struct { F0 int16; F1 bool; F2 int64 }
type GenS16 struct { F0 string; F1 string }
type GenS17 struct { F0 int; F1 int32 }
type GenS18 struct { F0 int32; F1 uint }
type GenS19 struct { F0 bool; F1 uint32; F2 uint }
type GenS20 struct { F0 uint8; F1 int64 }
type GenS21 struct { F0 uint64; F1 uint16 }
type GenS22 struct { F0 uint16; F1 uint16 }
type GenS23 struct { F0 int64; F1 uint8 }
type GenS24 struct { F0 uint32 }
type GenS25 struct { F0 uint16; F1 int8 }
type GenS26 struct { F0 uint8 }
type GenS27 struct { F0 int; F1 int16; F2 uint8 }
type GenS28 struct { F0 uint8; F1 uint16 }
type GenS29 struct { F0 int32 }
type GenS30 struct { F0 int; F1 int8; F2 bool }
type GenS31 struct { F0 int32; F1 uint8; F2 uint }
type GenS32 struct { F0 int8; F1 uint8 }
type GenS33 struct { F0 uint32 }
func genFn0(p0 bool, p1 string) uint8 {
    if (306 <= 2) {
        var v4 int32 = -97
        if p0 {
            var v5 int8 = -29
            _ = v5
        }
        _ = v4
    } else {
        p1 = "hello"
    }
    var v6 uint64 = 42
    var v7 uint = 1
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 255
}
func genFn1(p0 uint8, p1 uint8) int16 {
    switch 230 {
    case 863:
        if true {
            for fi8 := 0; fi8 < 20; fi8++ {
                v9 := GenS1{F0: 498}
                v10 := []int{fi8, (fi8 << ((420) & 63))}
                _ = fi8
                _ = v10
                _ = v9
            }
        } else {
            fmt.Println(p1, p0)
        }
    default:
        for fi11 := 0; fi11 < 9; fi11++ {
            for fi12 := 0; fi12 < 10; fi12++ {
                var v13 int32 = -1
                p1 = 1
                v13 += v13
                _ = fi12
                _ = v13
            }
            _ = fi11
        }
    }
    fmt.Println(p0, p1)
    fmt.Println(p1, p1)
    _ = p0
    _ = p1
    return 42
}

func main() {
    v1 := []string{"hello", "hello", "\n\t"}
    var v2 bool = (!(975 >= 18446744073709551615))
    var v3 int8 = 100
    v3 = (v3 >> ((65535) & 63))
    v2 = (v2 && (-32768 <= -675))
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    switch 32767 {
    case 32767:
        fmt.Println(v2)
    default:
        var v14 float64 = ((0.1 * (1e10 / (3.14159 - (-2.25 - 0.1)))) * 0.0)
        _ = v14
    }
    v15 := GenS1{F0: 4294967295}
    var v16 float64 = 1e10
    if (!(2 != 42)) {
        var v17 uint64 = 1
        _ = len(v1)
        _ = v17
    } else {
        var v18 int32 = 42
        for fi19 := 0; fi19 < 10; fi19++ {
            for fi20 := 0; fi20 < 16; fi20++ {
                fmt.Println(v16, v18, v16)
                if (!false) {
                    v3 -= v3
                    switch "foo bar" {
                    case "hello":
                        switch "a" {
                        case "\n\t":
                            v1 = append(v1, "")
                        case "hello":
                            v21 := GenS1{F0: 109}
                            fmt.Println(fi19)
                            v2 = (32767 < -709)
                            if (42 == 42) {
                                v2 = false
                                if true {
                                    fmt.Println(fi20, fi19)
                                }
                            }
                            _ = v21
                        default:
                            if v2 {
                                fmt.Println(fi19, v2)
                                if false {
                                    if ((v3 % (((-(v3))) | 1)) > (v3 & v3)) {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        var v22 uint8 = 197
                                        if ((v22 % ((v22) | 1)) < v22) {
                                            if (0 <= 42) {
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            }
                                            var v23 uint8 = v22
                                            var v24 int = (^(fi19))
                                            v1 = append(v1, "foo bar")
                                            _ = v23
                                            _ = v24
                                        }
                                        _ = v22
                                    }
                                    if v2 {
                                        if v2 {
                                            v2 = v2
                                        }
                                        if len(v1) > 0 {
                                            v1[0] = ("\n\t" + "foo bar")
                                        }
                                        switch ("\n\t" + "foo bar") {
                                        case "a":
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                        }
                                    } else {
                                        switch ("foo bar" + "hello") {
                                        case "hello":
                                            v16 += (v16 * v16)
                                        case "":
                                            fmt.Println(v16, fi20)
                                        case "foo bar":
                                            var v25 int = (fi19 & (^(fi20)))
                                            _ = v25
                                        default:
                                            var v26 float32 = (((float32(v16) + 0.0) - 1.5) * 1.5)
                                            _ = v26
                                        }
                                        v27 := GenS1{F0: 65535}
                                        v28 := GenS2{F0: 42}
                                        v18 = (v18 & v18)
                                        _ = v27
                                        _ = v28
                                    }
                                } else {
                                    var v29 bool = ((!(false || false)) && (255 != 0))
                                    var v30 float64 = v16
                                    _ = v29
                                    _ = v30
                                }
                                if v2 {
                                    var v31 uint16 = 32767
                                    v16 += v16
                                    v1 = append(v1, "\n\t")
                                    _ = v31
                                }
                            }
                        }
                    case "a":
                        if (0.1 <= 1.5) {
                            var v32 uint = 300
                            var v33 int16 = 1
                            if v2 {
                                v18 = v18
                                var v34 bool = true
                                fmt.Println(v34, fi19)
                                _ = v34
                            } else {
                                var v35 int32 = 18
                                var v36 int64 = int64(v3)
                                switch 0 {
                                case 2147483647:
                                    switch "\n\t" {
                                    case "":
                                        switch "hello" {
                                        case "hello":
                                            v37 := GenS1{F0: 2147483648}
                                            _ = v37
                                        case "":
                                            var v38 int16 = v33
                                            _ = v38
                                        }
                                    default:
                                        if true {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            if v2 {
                                                if true {
                                                    v36 = (v36 - -1)
                                                    v36 = v36
                                                } else {
                                                    _ = len(v1)
                                                    if len(v1) > 0 {
                                                        v1[0] = "a"
                                                    }
                                                }
                                                v39 := GenS3{F0: 671, F1: ""}
                                                fmt.Println(fi20, v16, v36)
                                                if false {
                                                    if (false || v2) {
                                                        var v40 bool = ("a" <= "\n\t")
                                                        _ = v40
                                                    } else {
                                                        v41 := GenS2{F0: 19}
                                                        fmt.Println(v18, v33)
                                                        v18 -= int32(v33)
                                                        _ = v41
                                                    }
                                                } else {
                                                    if (genFn0((true && (!(!(506 > 473)))), "\n\t") == 0) {
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        if v2 {
                                                            v42 := GenS4{F0: int(v3)}
                                                            _ = v42
                                                        } else {
                                                            switch (fi20 & fi19) {
                                                            case 2147483647:
                                                                if (!true) {
                                                                    switch 0 {
                                                                    case 127:
                                                                        var v43 int64 = (^(v36))
                                                                        _ = v43
                                                                    case 880:
                                                                        v33 += (v33 | (v33 / (((v33 % (((v33 - (v33 * v33))) | 1))) | 1)))
                                                                    default:
                                                                        switch "" {
                                                                        case "a":
                                                                            var v44 float32 = 0.0
                                                                            _ = v44
                                                                        case "":
                                                                            var v45 string = ((("\n\t" + ("a" + (("a" + ("foo bar" + "foo bar")) + ("hello" + "")))) + "hello") + "foo bar")
                                                                            _ = v45
                                                                        default:
                                                                            v46 := GenS3{F0: 65535, F1: ""}
                                                                            _ = v46
                                                                        }
                                                                    }
                                                                }
                                                            case 42:
                                                                var v47 int = fi20
                                                                _ = v47
                                                            case -2147483648:
                                                                v33 = v33
                                                            }
                                                        }
                                                        v16 -= v16
                                                    }
                                                }
                                                v35 = v18
                                                _ = v39
                                            } else {
                                                v48 := GenS1{F0: 1000}
                                                v3 = (v3 | (v3 >> (((v32 / ((v32) | 1))) & 63)))
                                                v35 = int32(v33)
                                                _ = v48
                                            }
                                            if false {
                                                v2 = (!false)
                                            }
                                            v3 += (v3 & v3)
                                            v49 := GenS1{F0: 246}
                                            _ = v49
                                        }
                                    }
                                }
                                if ((v32 | (v32 | (v32 * 370))) >= (v32 & v32)) {
                                    v3 += (v3 << ((v32) & 63))
                                    if v2 {
                                        fmt.Println(fi19, v32)
                                        v33 = (v33 >> (((v32 + v32)) & 63))
                                        v2 = v2
                                        switch (("\n\t" + "\n\t") + "foo bar") {
                                        case "foo bar":
                                            var v50 int8 = (v3 ^ v3)
                                            _ = v50
                                        case "a":
                                            v2 = v2
                                            var v51 float64 = 3.14159
                                            var v52 uint64 = 977
                                            _ = v51
                                            _ = v52
                                        case "":
                                            v16 -= 1e10
                                        }
                                    }
                                    if (v32 != v32) {
                                        v53 := GenS5{F0: 1, F1: (v36 + v36)}
                                        var v54 uint16 = 995
                                        var v55 string = ""
                                        _ = v53
                                        _ = v54
                                        _ = v55
                                    } else {
                                        var v56 float64 = v16
                                        if (!(!(fi20 <= fi20))) {
                                            if (((^(v35)) == (v18 | (v35 + (v18 << (((v32 << (((v32 * (v32 ^ (v32 / ((v32) | 1))))) & 63))) & 63))))) || (!v2)) {
                                                var v57 string = "\n\t"
                                                if (4294967295 == 917) {
                                                    var v58 string = v57
                                                    v59 := GenS6{F0: v32, F1: 32767}
                                                    fmt.Println(fi20, v36)
                                                    _ = v58
                                                    _ = v59
                                                } else {
                                                    if (v2 && v2) {
                                                        v16 += v16
                                                        if true {
                                                            if (188 <= 0) {
                                                                if (65535 >= 42) {
                                                                    v36 = (v36 >> ((v32) & 63))
                                                                    if ((3.14159 * float32(v16)) <= 0.1) {
                                                                        var v60 uint16 = 824
                                                                        v35 -= (v35 / ((v18) | 1))
                                                                        if v2 {
                                                                            v18 = (v35 % (((v35 ^ v18)) | 1))
                                                                            v56 += 3.14159
                                                                        }
                                                                        _ = len(v1)
                                                                        v61 := GenS5{F0: (v32 ^ v32), F1: v36}
                                                                        _ = v60
                                                                        _ = v61
                                                                    } else {
                                                                        switch fi19 {
                                                                        case 128:
                                                                            v36 = (v36 | (v36 ^ (v36 & -192)))
                                                                        }
                                                                        if (v2 || false) {
                                                                            var v62 uint32 = 2147483648
                                                                            _ = v62
                                                                        } else {
                                                                            var v63 uint8 = 127
                                                                            v64 := GenS4{F0: (^(fi19))}
                                                                            v33 -= 482
                                                                            switch (((("foo bar" + v57) + v57) + v57) + (v57 + v57)) {
                                                                            case "":
                                                                                if len(v1) >= 2 {
                                                                                    v1 = v1[1:len(v1)]
                                                                                }
                                                                            case "foo bar":
                                                                                v65 := GenS7{F0: 42}
                                                                                _ = v65
                                                                            case "a":
                                                                                var v66 int64 = (-(v36))
                                                                                _ = v66
                                                                            }
                                                                            v67 := GenS3{F0: 829, F1: (("hello" + (v57 + "\n\t")) + "hello")}
                                                                            _ = v63
                                                                            _ = v64
                                                                            _ = v67
                                                                        }
                                                                        if len(v1) >= 2 {
                                                                            v1 = v1[1:len(v1)]
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                        }
                                                        v18 += (v35 ^ v35)
                                                    }
                                                    var v68 int64 = -313
                                                    v69 := GenS7{F0: -1}
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    var v70 uint16 = 202
                                                    _ = v68
                                                    _ = v69
                                                    _ = v70
                                                }
                                                _ = v57
                                            } else {
                                                _ = len(v1)
                                                v32 = (v32 & 18446744073709551615)
                                            }
                                            if (3.14159 > v16) {
                                                v56 += (v16 - v16)
                                                v3 -= 44
                                            }
                                        }
                                        _ = v56
                                    }
                                    _ = len(v1)
                                    fmt.Println(v16)
                                    switch "foo bar" {
                                    case "\n\t":
                                        if v2 {
                                            v18 = 2147483647
                                            var v71 uint8 = 218
                                            if ((-(v35)) > -1) {
                                                var v72 int64 = (-(v36))
                                                switch fi19 {
                                                case 2147483647:
                                                    v18 = (v35 / (((v18 * (v18 & v35))) | 1))
                                                default:
                                                    _ = len(v1)
                                                }
                                                switch (fi19 * (fi19 | fi19)) {
                                                case -2147483648:
                                                    if true {
                                                        v71 = v71
                                                    }
                                                case -140:
                                                    if false {
                                                        v73 := GenS6{F0: 122, F1: 159}
                                                        v16 = v16
                                                        _ = v73
                                                    }
                                                case -294:
                                                    v33 = v33
                                                }
                                                var v74 bool = false
                                                _ = len(v1)
                                                _ = v72
                                                _ = v74
                                            } else {
                                                var v75 uint32 = 65535
                                                if (v2 || (!true)) {
                                                    switch "a" {
                                                    case "\n\t":
                                                        if len(v1) > 0 {
                                                            v1[0] = "\n\t"
                                                        }
                                                    default:
                                                        v2 = false
                                                    }
                                                    if len(v1) >= 2 {
                                                        v1 = v1[1:len(v1)]
                                                    }
                                                    switch fi20 {
                                                    case -1:
                                                        var v76 uint64 = 4294967295
                                                        _ = v76
                                                    case 4294967295:
                                                        v77 := GenS7{F0: int64(fi20)}
                                                        _ = v77
                                                    }
                                                    switch "foo bar" {
                                                    case "foo bar":
                                                        switch "" {
                                                        case "a":
                                                            fmt.Println(v33, v75, v32)
                                                        case "":
                                                            v78 := GenS2{F0: uint8(v32)}
                                                            _ = v78
                                                        }
                                                        var v79 uint64 = 388
                                                        var v80 float32 = 1.5
                                                        _ = v79
                                                        _ = v80
                                                    }
                                                } else {
                                                    _ = len(v1)
                                                    if false {
                                                        var v81 uint64 = 0
                                                        v35 += int32(v3)
                                                        v82 := GenS8{F0: (fi20 - -2147483648)}
                                                        _ = v81
                                                        _ = v82
                                                    }
                                                }
                                                if (v71 > (v71 * v71)) {
                                                    _ = len(v1)
                                                } else {
                                                    v36 -= v36
                                                }
                                                _ = v75
                                            }
                                            v83 := GenS9{F0: (v33 - (v33 * v33)), F1: 42, F2: v33}
                                            _ = v71
                                            _ = v83
                                        } else {
                                            if v2 {
                                                var v84 uint = v32
                                                if v2 {
                                                    v3 = int8(v18)
                                                    fmt.Println(v35, v84)
                                                }
                                                switch "hello" {
                                                case "hello":
                                                    v85 := GenS4{F0: fi19}
                                                    if ((!v2) && true) {
                                                        switch "foo bar" {
                                                        case "a":
                                                            if true {
                                                                v18 -= (v18 | v18)
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                            } else {
                                                                v86 := GenS9{F0: v33, F1: 9223372036854775807, F2: 32767}
                                                                if (0 != (v18 - v35)) {
                                                                    fmt.Println(v3, v32)
                                                                    switch 32767 {
                                                                    case -7:
                                                                        v87 := GenS7{F0: (-(v36))}
                                                                        _ = v87
                                                                    case 1:
                                                                        v84 = (v84 ^ v84)
                                                                        var v88 int32 = v35
                                                                        var v89 uint64 = 1
                                                                        v33 += -1
                                                                        _ = v88
                                                                        _ = v89
                                                                    case 32768:
                                                                        _ = len(v1)
                                                                    default:
                                                                        v90 := GenS10{F0: (v3 <= (^(v3))), F1: fi19, F2: v35}
                                                                        _ = v90
                                                                    }
                                                                    if v2 {
                                                                        var v91 int8 = v3
                                                                        _ = v91
                                                                    } else {
                                                                        fmt.Println(v16, v18, v84)
                                                                        if (v2 || v2) {
                                                                            var v92 uint64 = 4294967295
                                                                            if true {
                                                                                if (!true) {
                                                                                    fmt.Println(v18, v18)
                                                                                    fmt.Println(v16, v18)
                                                                                }
                                                                                var v93 int8 = v3
                                                                                var v94 uint32 = 42
                                                                                _ = v93
                                                                                _ = v94
                                                                            }
                                                                            if v2 {
                                                                                _ = len(v1)
                                                                            } else {
                                                                                fmt.Println(v36, v35)
                                                                                var v95 int64 = v36
                                                                                v2 = (0 >= 367)
                                                                                _ = v95
                                                                            }
                                                                            _ = v92
                                                                        }
                                                                        if v2 {
                                                                            v3 = v3
                                                                            var v96 int = (fi19 * (fi19 / (((fi20 - fi20)) | 1)))
                                                                            if v2 {
                                                                                var v97 uint32 = 2147483648
                                                                                _ = len(v1)
                                                                                switch ("hello" + "hello") {
                                                                                case "hello":
                                                                                    v98 := GenS11{F0: v2, F1: (v33 % ((42) | 1)), F2: 42}
                                                                                    _ = v98
                                                                                }
                                                                                _ = v97
                                                                            }
                                                                            if ((v3 + -82) >= int8(v33)) {
                                                                                if (v18 <= v18) {
                                                                                    if (!(-1 >= (^(fi20)))) {
                                                                                        fmt.Println(v84, v36)
                                                                                        fmt.Println(v3, v84)
                                                                                        v84 += v84
                                                                                    } else {
                                                                                        switch (("" + "hello") + "") {
                                                                                        case "hello":
                                                                                            var v99 uint8 = 176
                                                                                            var v100 int16 = genFn1(v99, 2)
                                                                                            _ = v100
                                                                                            _ = v99
                                                                                        }
                                                                                        var v101 string = ("\n\t" + "")
                                                                                        _ = v101
                                                                                    }
                                                                                    var v102 float64 = v16
                                                                                    var v103 uint16 = 596
                                                                                    fmt.Println(v3, v36)
                                                                                    _ = v102
                                                                                    _ = v103
                                                                                } else {
                                                                                    var v104 uint16 = 845
                                                                                    fmt.Println(v32, fi20, v35)
                                                                                    fmt.Println(v2, v84, v32)
                                                                                    _ = v104
                                                                                }
                                                                                _ = len(v1)
                                                                                switch "a" {
                                                                                case "hello":
                                                                                    v105 := GenS12{F0: 344}
                                                                                    _ = v105
                                                                                default:
                                                                                    fmt.Println(v33)
                                                                                }
                                                                            }
                                                                            var v106 int8 = 31
                                                                            fmt.Println(v16, v35)
                                                                            var v107 uint32 = 2
                                                                            _ = len(v1)
                                                                            _ = v106
                                                                            _ = v107
                                                                            _ = v96
                                                                        }
                                                                    }
                                                                } else {
                                                                    if (!(fi20 == (fi20 >> ((v32) & 63)))) {
                                                                        switch ("\n\t" + "\n\t") {
                                                                        case "foo bar":
                                                                            fmt.Println(v16)
                                                                            fmt.Println(fi20, v33, fi19)
                                                                        default:
                                                                            v108 := GenS10{F0: (!(v2 || true)), F1: (^(fi19)), F2: (v18 << (((^(v84))) & 63))}
                                                                            _ = v108
                                                                        }
                                                                        v36 = v36
                                                                        var v109 uint = 37
                                                                        _ = len(v1)
                                                                        v84 -= (^(v109))
                                                                        _ = v109
                                                                    }
                                                                    var v110 int64 = (^(v36))
                                                                    var v111 int64 = -9223372036854775808
                                                                    _ = v110
                                                                    _ = v111
                                                                }
                                                                _ = v86
                                                            }
                                                        case "\n\t":
                                                            switch fi20 {
                                                            case 9223372036854775807:
                                                                v112 := GenS7{F0: v36}
                                                                if true {
                                                                    if (((v33 > -756) && (fi20 >= (fi19 << ((v84) & 63)))) && (!v2)) {
                                                                        _ = len(v1)
                                                                        var v113 int8 = (v3 - 100)
                                                                        _ = v113
                                                                    }
                                                                    v33 = (v33 * genFn1(127, 140))
                                                                }
                                                                if (0.0 == (3.14159 - -2.25)) {
                                                                    fmt.Println(v2, v84, v18)
                                                                    var v114 uint16 = 42
                                                                    v35 += (v35 + (^(v35)))
                                                                    _ = v114
                                                                }
                                                                _ = v112
                                                            case -2147483648:
                                                                if (v2 || v2) {
                                                                    var v115 uint64 = 42
                                                                    v116 := GenS13{F0: "foo bar", F1: 196}
                                                                    if (1 <= uint32(v115)) {
                                                                        v35 += (v18 << ((v84) & 63))
                                                                    } else {
                                                                        var v117 int8 = v3
                                                                        v33 -= v33
                                                                        _ = v117
                                                                    }
                                                                    var v118 int32 = v18
                                                                    var v119 uint64 = 42
                                                                    _ = v115
                                                                    _ = v116
                                                                    _ = v118
                                                                    _ = v119
                                                                }
                                                            case 4294967295:
                                                                v33 = (v33 << (((v84 * 447)) & 63))
                                                            default:
                                                                v120 := GenS8{F0: fi19}
                                                                _ = v120
                                                            }
                                                        case "":
                                                            var v121 int32 = 1
                                                            _ = v121
                                                        }
                                                        var v122 int8 = int8(v33)
                                                        var v123 float64 = v16
                                                        _ = v122
                                                        _ = v123
                                                    } else {
                                                        fmt.Println(v18, fi20, fi19)
                                                        switch (fi19 ^ fi20) {
                                                        case 316:
                                                            fmt.Println(v35, v2, v36)
                                                        case 0:
                                                            if ((-(v3)) < v3) {
                                                                var v124 int32 = (v35 ^ v35)
                                                                v125 := GenS14{F0: v3, F1: -2147483648}
                                                                switch 9223372036854775807 {
                                                                case -998:
                                                                    v1 = append(v1, "a")
                                                                default:
                                                                    switch fi19 {
                                                                    case 32767:
                                                                        v124 += v124
                                                                    default:
                                                                        switch (fi19 | fi20) {
                                                                        case -416:
                                                                            switch 2147483647 {
                                                                            case -9223372036854775808:
                                                                                var v126 uint16 = 42
                                                                                _ = v126
                                                                            case 127:
                                                                                fmt.Println(v35, fi19)
                                                                            case 675:
                                                                                if (((3.14159 - (3.14159 * 3.14159)) > 3.14159) && ((((fi20 | (fi20 | fi20)) > fi20) || (4294967295 > 255)) || (v16 > v16))) {
                                                                                    switch int(v36) {
                                                                                    case -809:
                                                                                        fmt.Println(v16, v33)
                                                                                        v127 := GenS15{F0: (v33 ^ (^(v33))), F1: (v84 <= v32), F2: v36}
                                                                                        _ = v127
                                                                                    case -149:
                                                                                        v33 = (v33 - -252)
                                                                                    }
                                                                                    fmt.Println(v2)
                                                                                    v36 -= v36
                                                                                    if v2 {
                                                                                        v128 := GenS15{F0: (^(v33)), F1: v2, F2: v36}
                                                                                        var v129 uint32 = 681
                                                                                        v130 := GenS16{F0: "foo bar", F1: "hello"}
                                                                                        _ = v128
                                                                                        _ = v129
                                                                                        _ = v130
                                                                                    } else {
                                                                                        v131 := GenS17{F0: int(v124), F1: 42}
                                                                                        _ = v131
                                                                                    }
                                                                                    fmt.Println(v32, v2, v35)
                                                                                } else {
                                                                                    if len(v1) >= 2 {
                                                                                        v1 = v1[1:len(v1)]
                                                                                    }
                                                                                }
                                                                            default:
                                                                                if (1 == 65535) {
                                                                                    if v2 {
                                                                                        if len(v1) >= 2 {
                                                                                            v1 = v1[1:len(v1)]
                                                                                        }
                                                                                        if true {
                                                                                            fmt.Println(v35)
                                                                                        }
                                                                                        if v2 {
                                                                                            var v132 uint16 = 1
                                                                                            _ = v132
                                                                                        } else {
                                                                                            var v133 uint = 440
                                                                                            if ((3.14159 / (((3.14159 + 1e10) * (-2.25 * ((0.1 * 0.1) - 3.14159))) * float32(v16))) >= 0.0) {
                                                                                                if false {
                                                                                                    var v134 int64 = (v36 | 9223372036854775807)
                                                                                                    _ = v134
                                                                                                } else {
                                                                                                    var v135 float32 = 0.0
                                                                                                    _ = v135
                                                                                                }
                                                                                            }
                                                                                            v36 -= (v36 * (v36 << ((v84) & 63)))
                                                                                            fmt.Println(v36)
                                                                                            if true {
                                                                                                var v136 uint16 = 42
                                                                                                v137 := GenS18{F0: v18, F1: v32}
                                                                                                _ = v136
                                                                                                _ = v137
                                                                                            }
                                                                                            _ = v133
                                                                                        }
                                                                                        if (v2 || (32767 != 325)) {
                                                                                            fmt.Println(v32, v18)
                                                                                            var v138 string = ("" + "")
                                                                                            if (true && false) {
                                                                                                switch fi19 {
                                                                                                case 2147483647:
                                                                                                    var v139 bool = (0 <= 2)
                                                                                                    _ = v139
                                                                                                default:
                                                                                                    if (v2 && v2) {
                                                                                                        if (2 < 2) {
                                                                                                            v36 = v36
                                                                                                        }
                                                                                                    } else {
                                                                                                        v33 = (-(v33))
                                                                                                        fmt.Println(v35)
                                                                                                        var v140 uint8 = 154
                                                                                                        v36 = (v36 ^ v36)
                                                                                                        var v141 uint8 = v140
                                                                                                        if len(v1) > 0 {
                                                                                                            v1[0] = ("" + "")
                                                                                                        }
                                                                                                        _ = v140
                                                                                                        _ = v141
                                                                                                    }
                                                                                                }
                                                                                                if (332 > 2) {
                                                                                                    var v142 int64 = 886
                                                                                                    var v143 int = 9223372036854775807
                                                                                                    var v144 uint8 = genFn0((!true), v138)
                                                                                                    _ = v142
                                                                                                    _ = v143
                                                                                                    _ = v144
                                                                                                }
                                                                                                if ((-(v33)) != v33) {
                                                                                                    var v145 bool = (v2 && (42 >= 0))
                                                                                                    _ = v145
                                                                                                } else {
                                                                                                    if ((!(!false)) || v2) {
                                                                                                        if true {
                                                                                                            var v146 uint8 = 127
                                                                                                            fmt.Println(v18)
                                                                                                            fmt.Println(v33, fi19)
                                                                                                            _ = v146
                                                                                                        } else {
                                                                                                            v2 = (!false)
                                                                                                            if (231 != 65535) {
                                                                                                                v1 = append(v1, (("hello" + "") + "\n\t"))
                                                                                                                switch v138 {
                                                                                                                case "":
                                                                                                                    v147 := GenS19{F0: (v3 <= (v3 ^ v3)), F1: 42, F2: v32}
                                                                                                                    _ = v147
                                                                                                                case "\n\t":
                                                                                                                    v148 := GenS10{F0: v2, F1: fi20, F2: (v124 ^ v18)}
                                                                                                                    _ = v148
                                                                                                                default:
                                                                                                                    var v149 int = -2147483648
                                                                                                                    _ = v149
                                                                                                                }
                                                                                                                var v150 float64 = (float64(1.5) + v16)
                                                                                                                _ = v150
                                                                                                            } else {
                                                                                                                switch fi20 {
                                                                                                                case 1:
                                                                                                                    v1 = append(v1, ("foo bar" + (v138 + "foo bar")))
                                                                                                                }
                                                                                                                v35 += v35
                                                                                                                v151 := GenS16{F0: v138, F1: v138}
                                                                                                                _ = v151
                                                                                                            }
                                                                                                        }
                                                                                                    } else {
                                                                                                        var v152 bool = (false && (v2 || ((v36 % ((v36) | 1)) < (v36 + (v36 | (v36 ^ v36))))))
                                                                                                        var v153 uint32 = 614
                                                                                                        v154 := GenS2{F0: 54}
                                                                                                        _ = v152
                                                                                                        _ = v153
                                                                                                        _ = v154
                                                                                                    }
                                                                                                    var v155 bool = true
                                                                                                    _ = len(v1)
                                                                                                    _ = v155
                                                                                                }
                                                                                            }
                                                                                            _ = v138
                                                                                        } else {
                                                                                            if len(v1) >= 2 {
                                                                                                v1 = v1[1:len(v1)]
                                                                                            }
                                                                                            if v2 {
                                                                                                if v2 {
                                                                                                    v156 := GenS15{F0: v33, F1: (((fi20 | (fi19 ^ -9223372036854775808)) == fi19) && (v18 == -1)), F2: (v36 - v36)}
                                                                                                    v157 := GenS20{F0: 0, F1: 42}
                                                                                                    fmt.Println(v2, v18, v124)
                                                                                                    _ = v156
                                                                                                    _ = v157
                                                                                                } else {
                                                                                                    if (2 != 0) {
                                                                                                        v16 += v16
                                                                                                        var v158 uint = v32
                                                                                                        _ = v158
                                                                                                    } else {
                                                                                                        v2 = (783 <= 547)
                                                                                                    }
                                                                                                    v2 = (!((!true) && v2))
                                                                                                }
                                                                                                var v159 string = "foo bar"
                                                                                                switch fi19 {
                                                                                                case -349:
                                                                                                    v160 := GenS8{F0: (fi19 ^ fi19)}
                                                                                                    _ = v160
                                                                                                case -1:
                                                                                                    v161 := GenS21{F0: 250, F1: 42}
                                                                                                    _ = v161
                                                                                                default:
                                                                                                    v162 := GenS17{F0: (fi20 + 9223372036854775807), F1: (-(v124))}
                                                                                                    _ = v162
                                                                                                }
                                                                                                _ = v159
                                                                                            }
                                                                                            v36 += (v36 - v36)
                                                                                        }
                                                                                    }
                                                                                    var v163 int8 = v3
                                                                                    if (20 >= 1) {
                                                                                        v124 -= (^(v124))
                                                                                        v16 -= (v16 + v16)
                                                                                        if len(v1) > 0 {
                                                                                            v1[0] = ""
                                                                                        }
                                                                                    } else {
                                                                                        var v164 int32 = (v18 / (((v35 / ((-1) | 1))) | 1))
                                                                                        v165 := GenS22{F0: 2, F1: 42}
                                                                                        v166 := GenS4{F0: fi20}
                                                                                        _ = v164
                                                                                        _ = v165
                                                                                        _ = v166
                                                                                    }
                                                                                    _ = v163
                                                                                }
                                                                            }
                                                                        case -372:
                                                                            fmt.Println(v124, v18)
                                                                        case 2147483647:
                                                                            v18 = v18
                                                                        default:
                                                                            v167 := GenS23{F0: (v36 ^ v36), F1: 255}
                                                                            _ = v167
                                                                        }
                                                                    }
                                                                }
                                                                _ = v124
                                                                _ = v125
                                                            } else {
                                                                var v168 int = 32768
                                                                switch -141 {
                                                                case 32767:
                                                                    v84 -= (^(v32))
                                                                case 32768:
                                                                    v169 := GenS22{F0: 196, F1: 2}
                                                                    v18 += -975
                                                                    v35 = -1
                                                                    _ = v169
                                                                case 377:
                                                                    v170 := GenS2{F0: 8}
                                                                    _ = v170
                                                                }
                                                                if ((0 > genFn0(v2, "foo bar")) && ((v36 << ((v32) & 63)) < (v36 >> (((^(v84))) & 63)))) {
                                                                    if len(v1) > 0 {
                                                                        v1[0] = "\n\t"
                                                                    }
                                                                    if len(v1) >= 2 {
                                                                        v1 = v1[1:len(v1)]
                                                                    }
                                                                    if true {
                                                                        switch fi19 {
                                                                        case 4294967295:
                                                                            var v171 int32 = 0
                                                                            _ = v171
                                                                        case 0:
                                                                            var v172 float32 = 3.14159
                                                                            _ = v172
                                                                        case -263:
                                                                            var v173 int = (fi19 % (((fi20 >> (((v32 + (v32 >> (((v32 * v32)) & 63)))) & 63))) | 1))
                                                                            _ = v173
                                                                        default:
                                                                            v174 := GenS24{F0: 2}
                                                                            _ = v174
                                                                        }
                                                                        if v2 {
                                                                            v175 := GenS25{F0: 32767, F1: v3}
                                                                            _ = v175
                                                                        } else {
                                                                            switch fi19 {
                                                                            case -524:
                                                                                if true {
                                                                                    v32 = v84
                                                                                    if false {
                                                                                        var v176 uint32 = 0
                                                                                        v177 := GenS1{F0: (v176 / (((^(v176))) | 1))}
                                                                                        v178 := GenS26{F0: 42}
                                                                                        _ = v176
                                                                                        _ = v177
                                                                                        _ = v178
                                                                                    } else {
                                                                                        switch "\n\t" {
                                                                                        case "":
                                                                                            v18 += (v18 ^ v18)
                                                                                        case "hello":
                                                                                            v179 := GenS27{F0: (-(fi20)), F1: 42, F2: 2}
                                                                                            _ = v179
                                                                                        case "\n\t":
                                                                                            if ((v2 || ((-2.25 + 3.14159) < v16)) || v2) {
                                                                                                if len(v1) > 0 {
                                                                                                    v1[0] = ("hello" + "foo bar")
                                                                                                }
                                                                                                switch "" {
                                                                                                case "":
                                                                                                    var v180 float64 = (1e10 + -2.25)
                                                                                                    _ = v180
                                                                                                case "a":
                                                                                                    fmt.Println(v35, v18, fi19)
                                                                                                    v181 := GenS3{F0: 1, F1: "hello"}
                                                                                                    if len(v1) >= 2 {
                                                                                                        v1 = v1[1:len(v1)]
                                                                                                    }
                                                                                                    _ = v181
                                                                                                }
                                                                                                var v182 int8 = v3
                                                                                                _ = v182
                                                                                            } else {
                                                                                                v168 = v168
                                                                                                v183 := GenS3{F0: 2, F1: ("\n\t" + "a")}
                                                                                                var v184 int16 = (v33 ^ genFn1(2, 2))
                                                                                                _ = len(v1)
                                                                                                var v185 uint16 = 255
                                                                                                fmt.Println(fi20)
                                                                                                var v186 uint64 = 18446744073709551615
                                                                                                var v187 string = "\n\t"
                                                                                                v168 += (fi20 << ((v32) & 63))
                                                                                                fmt.Println(v186, v184)
                                                                                                _ = v183
                                                                                                _ = v184
                                                                                                _ = v185
                                                                                                _ = v186
                                                                                                _ = v187
                                                                                            }
                                                                                        default:
                                                                                            v188 := GenS9{F0: v33, F1: (v36 * v36), F2: v33}
                                                                                            _ = v188
                                                                                        }
                                                                                        if v2 {
                                                                                            if len(v1) >= 2 {
                                                                                                v1 = v1[1:len(v1)]
                                                                                            }
                                                                                            v36 = (v36 + 177)
                                                                                            var v189 float32 = float32(v16)
                                                                                            _ = v189
                                                                                        } else {
                                                                                            v33 -= genFn1(234, 1)
                                                                                        }
                                                                                        var v190 uint32 = 676
                                                                                        _ = v190
                                                                                    }
                                                                                }
                                                                            default:
                                                                                if len(v1) >= 2 {
                                                                                    v1 = v1[1:len(v1)]
                                                                                }
                                                                            }
                                                                            var v191 int8 = (v3 - (v3 - (v3 - (v3 & 127))))
                                                                            var v192 float64 = v16
                                                                            _ = v191
                                                                            _ = v192
                                                                        }
                                                                        v193 := GenS28{F0: 236, F1: 1}
                                                                        _ = v193
                                                                    } else {
                                                                        if (!((v32 >> ((v32) & 63)) != 4294967295)) {
                                                                            fmt.Println(fi19)
                                                                        } else {
                                                                            v33 = genFn1(0, genFn0((true || (v2 || false)), "hello"))
                                                                            fmt.Println(v36)
                                                                            var v194 float32 = 0.1
                                                                            v1 = append(v1, ("foo bar" + (("\n\t" + "a") + "a")))
                                                                            _ = v194
                                                                        }
                                                                        if len(v1) > 0 {
                                                                            v1[0] = ""
                                                                        }
                                                                    }
                                                                    switch "foo bar" {
                                                                    case "a":
                                                                        _ = len(v1)
                                                                    default:
                                                                        v32 -= v84
                                                                        if ((v16 * float64(-2.25)) >= ((0.0 + (((1.5 - 1e10) + 1e10) * 1.5)) - (1e10 - v16))) {
                                                                            fmt.Println(v32, v33)
                                                                            v195 := GenS29{F0: v18}
                                                                            _ = v195
                                                                        } else {
                                                                            switch "a" {
                                                                            case "hello":
                                                                                v196 := GenS29{F0: v35}
                                                                                _ = v196
                                                                            case "a":
                                                                                v16 += v16
                                                                            default:
                                                                                var v197 int = fi20
                                                                                _ = v197
                                                                            }
                                                                            fmt.Println(v33, v18)
                                                                            v16 += v16
                                                                        }
                                                                        v198 := GenS30{F0: fi19, F1: 42, F2: v2}
                                                                        _ = v198
                                                                    }
                                                                } else {
                                                                    v16 -= (v16 * (v16 + 1e10))
                                                                    var v199 uint32 = 2
                                                                    _ = v199
                                                                }
                                                                _ = v168
                                                            }
                                                        case 32767:
                                                            var v200 int8 = 82
                                                            _ = v200
                                                        default:
                                                            v18 = -1
                                                        }
                                                        var v201 int = (-(fi19))
                                                        _ = v201
                                                    }
                                                    switch "\n\t" {
                                                    case "":
                                                        var v202 float32 = (((1e10 - 1.5) + 0.0) - 1e10)
                                                        _ = v202
                                                    case "hello":
                                                        v203 := GenS9{F0: genFn1(0, 127), F1: (v36 << (((^(v84))) & 63)), F2: int16(v18)}
                                                        _ = v203
                                                    }
                                                    _ = v85
                                                }
                                                _ = v84
                                            } else {
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                v204 := GenS17{F0: fi19, F1: v18}
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                _ = v204
                                            }
                                            var v205 int64 = 0
                                            v35 -= (-(v18))
                                            _ = v205
                                        }
                                    }
                                } else {
                                    v16 -= v16
                                    switch ("\n\t" + "\n\t") {
                                    case "a":
                                        _ = len(v1)
                                    default:
                                        fmt.Println(v35)
                                    }
                                    switch ("\n\t" + "") {
                                    case "foo bar":
                                        v206 := GenS31{F0: int32(fi19), F1: 255, F2: v32}
                                        _ = v206
                                    }
                                }
                                _ = v35
                                _ = v36
                            }
                            if (849 == 2) {
                                if v2 {
                                    var v207 int8 = (v3 ^ 35)
                                    _ = v207
                                }
                            }
                            v2 = (901 != 2147483648)
                            v33 -= (v33 | v33)
                            if ((v2 && (v3 > (v3 * 100))) && false) {
                                var v208 string = ""
                                var v209 uint16 = 0
                                _ = v208
                                _ = v209
                            }
                            _ = v32
                            _ = v33
                        }
                    default:
                        var v210 string = "\n\t"
                        _ = v210
                    }
                    v3 = (v3 / (((v3 % (((v3 & 42)) | 1))) | 1))
                } else {
                    var v211 uint16 = 32767
                    fmt.Println(v3, fi20)
                    _ = v211
                }
                var v212 uint32 = 2147483647
                _ = fi20
                _ = v212
            }
            _ = fi19
        }
        _ = v18
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    fmt.Println(v2)
    var v213 uint16 = 656
    fmt.Println(v213, v2, v3)
    if (761 >= 2147483647) {
        v214 := GenS32{F0: v3, F1: 224}
        v1 = append(v1, (("\n\t" + "foo bar") + ("" + ("a" + "foo bar"))))
        for fi215 := 0; fi215 < 2; fi215++ {
            v216 := GenS27{F0: 42, F1: 0, F2: 127}
            for fi217 := 0; fi217 < 17; fi217++ {
                if v2 {
                    v218 := GenS19{F0: (!v2), F1: 2147483647, F2: 466}
                    v219 := GenS1{F0: 2147483647}
                    _ = v218
                    _ = v219
                }
                var v220 float32 = ((1.5 + 3.14159) - 3.14159)
                _ = fi217
                _ = v220
            }
            if true {
                var v221 uint = 917
                _ = v221
            }
            _ = fi215
            _ = v216
        }
        _ = v214
    } else {
        var v222 string = "a"
        v223 := GenS27{F0: 128, F1: -1, F2: 202}
        _ = v222
        _ = v223
    }
    v2 = v2
    if (v2 || false) {
        v16 = 3.14159
    } else {
        if v2 {
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
        }
        if ((!(42 <= -1)) || v2) {
            var v224 int16 = 1
            v225 := GenS33{F0: 2147483647}
            _ = v224
            _ = v225
        }
        v1 = append(v1, ("a" + "\n\t"))
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v213:", v213)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v15
    _ = v16
    _ = v2
    _ = v213
    _ = v3
}
