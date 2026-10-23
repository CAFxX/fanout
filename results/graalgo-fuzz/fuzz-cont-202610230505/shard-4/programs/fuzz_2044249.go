package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 int16; F1 uint8; F2 int16 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 int; F1 uint32; F2 int16 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint32; F1 int32; F2 int8 }
type GenS7 struct { F0 uint8; F1 int32 }
type GenS8 struct { F0 uint; F1 int; F2 uint8 }
type GenS9 struct { F0 uint16; F1 bool }
type GenS10 struct { F0 int }
type GenS11 struct { F0 uint64 }
type GenS12 struct { F0 int16; F1 uint16; F2 int8 }
type GenS13 struct { F0 int16; F1 int64; F2 string }
type GenS14 struct { F0 int32; F1 int16; F2 int16 }
type GenS15 struct { F0 uint32; F1 uint }
func genFn0() int {
    if (127 > 42) {
        var v7 string = "hello"
        _ = v7
    }
    var v8 uint16 = 1
    v8 -= 2
    _ = v8
    return -924
}
func genFn1(p0 int16, p1 string) int16 {
    switch genFn0() {
    case 128:
        var v9 bool = true
        _ = v9
    default:
        v10 := []int{}
        _ = v10
    }
    var v11 int16 = 427
    switch (p1 + p1) {
    case "\n\t":
        if (0 == 80) {
            var v12 string = ("\n\t" + "")
            v13 := GenS1{F0: "hello"}
            _ = v12
            _ = v13
        }
    case "hello":
        switch p1 {
        case "foo bar":
            if ((((float64(1.5) * (((3.14159 - (-2.25 + (1e10 + float64(-2.25)))) * 3.14159) * 3.14159)) / 1.0) * -2.25) < 0.1) {
                var v14 int = 568
                var v15 int64 = 1
                _ = v14
                _ = v15
            } else {
                v16 := []int{int(v11), 9223372036854775807, 32768}
                for fi17 := 0; fi17 < 18; fi17++ {
                    for fi18 := 0; fi18 < 1; fi18++ {
                        fmt.Println(v11)
                        v11 -= p0
                        v19 := GenS1{F0: "a"}
                        _ = fi18
                        _ = v19
                    }
                    _ = fi17
                }
                p1 = ((p1 + ("" + "hello")) + (p1 + "foo bar"))
                _ = v16
            }
        case "":
            v20 := []int{32767, 4294967295, 38}
            _ = v20
        }
    case "a":
        for fi21 := 0; fi21 < 12; fi21++ {
            v11 = v11
            var v22 bool = ((((0.1 - (-2.25 * 1.5)) - -2.25) + 1.5) >= -2.25)
            _ = fi21
            _ = v22
        }
    }
    _ = p0
    _ = p1
    _ = v11
    return (^(p0))
}

func main() {
    var v1 uint8 = 1
    v2 := make(map[int]bool)
    v3 := make(map[string]uint32)
    var v4 uint64 = 0
    var v5 float64 = float64(0.1)
    var v6 int32 = -2147483648
    var v23 uint64 = v4
    v1 += (v1 % ((87) | 1))
    v5 = v5
    var v24 uint32 = 179
    v25 := v2[0]
    v26 := []int{-451, genFn0(), -9223372036854775808, 42}
    if len(v26) > 0 {
        v26[0] = -1
    }
    if ((!v25) || (false && (33 == 32767))) {
        var v27 int16 = 1000
        var v28 float32 = 0.0
        if false {
            v29 := GenS1{F0: "a"}
            _ = v29
        } else {
            var v30 uint8 = uint8(v24)
            for fi31 := 0; fi31 < 17; fi31++ {
                for fi32 := 0; fi32 < 5; fi32++ {
                    if true {
                        v28 += v28
                        var v33 uint = 42
                        v2[(fi31 / ((32768) | 1))] = v25
                        v30 += (v1 % ((v30) | 1))
                        _ = v33
                    } else {
                        if v25 {
                            if (2 != 381) {
                                _ = len(v3)
                                if true {
                                    v34 := GenS2{F0: (-(v27)), F1: v30, F2: genFn1((v27 >> ((4294967295) & 63)), "a")}
                                    v5 = v5
                                    _ = v34
                                }
                                v35 := GenS2{F0: (v27 & (v27 * (v27 & (v27 << ((494) & 63))))), F1: v1, F2: 32767}
                                _ = v35
                            } else {
                                switch ("hello" + (("a" + ("a" + "foo bar")) + "a")) {
                                case "\n\t":
                                    v6 = v6
                                case "":
                                    var v36 uint = 0
                                    _ = v36
                                case "foo bar":
                                    if false {
                                        switch 143 {
                                        case 4294967295:
                                            if (("a" + "hello") <= "\n\t") {
                                                if (!v25) {
                                                    var v37 int8 = int8(v27)
                                                    switch (fi31 | fi31) {
                                                    case -281:
                                                        fmt.Println(v4, v5)
                                                    case 32768:
                                                        var v38 int16 = genFn1(32767, "hello")
                                                        _ = v38
                                                    }
                                                    if (((^(v24)) <= (v24 & (v24 & (v24 / ((17) | 1))))) && v25) {
                                                        switch ("\n\t" + "") {
                                                        case "foo bar":
                                                            if true {
                                                                if false {
                                                                    switch "a" {
                                                                    case "":
                                                                        v37 += -1
                                                                    default:
                                                                        delete(v3, ("foo bar" + (("foo bar" + (("\n\t" + "") + ("hello" + "hello"))) + "a")))
                                                                    }
                                                                    if len(v26) >= 2 {
                                                                        v26 = v26[1:len(v26)]
                                                                    }
                                                                    v39 := GenS1{F0: ("hello" + ("" + "a"))}
                                                                    _ = v39
                                                                }
                                                                var v40 int64 = int64(v23)
                                                                if ("\n\t" >= ("a" + "a")) {
                                                                    if true {
                                                                        fmt.Println(v28, v37, v1)
                                                                        _ = len(v26)
                                                                        switch (fi32 - (fi31 | fi31)) {
                                                                        case 42:
                                                                            fmt.Println(fi31)
                                                                        case 32767:
                                                                            v23 = 18446744073709551615
                                                                        case 249:
                                                                            var v41 int8 = (v37 << ((18446744073709551615) & 63))
                                                                            _ = v41
                                                                        }
                                                                    } else {
                                                                        v30 += v1
                                                                        if v25 {
                                                                            if false {
                                                                                v42 := GenS3{F0: "foo bar"}
                                                                                _ = v42
                                                                            } else {
                                                                                v25 = (false && true)
                                                                                v43 := GenS4{F0: (-(fi32)), F1: v24, F2: (v27 / ((v27) | 1))}
                                                                                switch "a" {
                                                                                case "foo bar":
                                                                                    switch ("a" + ("" + "foo bar")) {
                                                                                    case "":
                                                                                        switch (fi32 % ((fi31) | 1)) {
                                                                                        case 9223372036854775807:
                                                                                            switch (("hello" + ("foo bar" + ("\n\t" + (("" + "foo bar") + ("" + ""))))) + "foo bar") {
                                                                                            case "hello":
                                                                                                switch 32767 {
                                                                                                case -648:
                                                                                                    v4 = v23
                                                                                                }
                                                                                            case "":
                                                                                                if false {
                                                                                                    if v25 {
                                                                                                        _ = len(v3)
                                                                                                        v44 := GenS3{F0: "hello"}
                                                                                                        if (2 <= 1) {
                                                                                                            v45 := GenS5{F0: v27}
                                                                                                            switch fi31 {
                                                                                                            case 4294967295:
                                                                                                                switch "foo bar" {
                                                                                                                case "":
                                                                                                                    switch "" {
                                                                                                                    case "":
                                                                                                                        _ = len(v26)
                                                                                                                        v46 := GenS6{F0: v24, F1: (^(v6)), F2: (v37 & v37)}
                                                                                                                        _ = v46
                                                                                                                    case "\n\t":
                                                                                                                        fmt.Println(v5)
                                                                                                                    case "hello":
                                                                                                                        var v47 uint32 = v24
                                                                                                                        _ = v47
                                                                                                                    default:
                                                                                                                        switch ("foo bar" + "\n\t") {
                                                                                                                        case "":
                                                                                                                            v23 = v23
                                                                                                                        }
                                                                                                                        if v25 {
                                                                                                                            var v48 int16 = (v27 % ((v27) | 1))
                                                                                                                            _ = v48
                                                                                                                        } else {
                                                                                                                            var v49 uint32 = v24
                                                                                                                            _ = v49
                                                                                                                        }
                                                                                                                    }
                                                                                                                case "foo bar":
                                                                                                                    v23 = v4
                                                                                                                case "a":
                                                                                                                    v5 = v5
                                                                                                                }
                                                                                                            default:
                                                                                                                if v25 {
                                                                                                                    var v50 uint = 1
                                                                                                                    if true {
                                                                                                                        v27 += (v27 & (v27 - genFn1(int16(v30), ("foo bar" + ((("hello" + ("" + ("foo bar" + ""))) + ("\n\t" + "hello")) + ("\n\t" + "\n\t"))))))
                                                                                                                        var v51 uint32 = 311
                                                                                                                        _ = v51
                                                                                                                    }
                                                                                                                    v28 += (v28 / float32((v5 * v5)))
                                                                                                                    _ = v50
                                                                                                                } else {
                                                                                                                    if false {
                                                                                                                        if v25 {
                                                                                                                            v30 = v1
                                                                                                                            v52 := GenS3{F0: (("foo bar" + ("" + ("\n\t" + "a"))) + ((("\n\t" + "\n\t") + ("" + "a")) + ("\n\t" + "foo bar")))}
                                                                                                                            _ = v52
                                                                                                                        } else {
                                                                                                                            fmt.Println(v5, v4, v4)
                                                                                                                            var v53 uint16 = 65535
                                                                                                                            switch fi32 {
                                                                                                                            case 42:
                                                                                                                                v5 += v5
                                                                                                                            }
                                                                                                                            switch ("\n\t" + ("a" + "\n\t")) {
                                                                                                                            case "hello":
                                                                                                                                v26 = append(v26, fi32)
                                                                                                                            }
                                                                                                                            _ = v53
                                                                                                                        }
                                                                                                                        if len(v26) > 0 {
                                                                                                                            v26[0] = (-(fi31))
                                                                                                                        }
                                                                                                                        v4 += (^(v23))
                                                                                                                        switch "hello" {
                                                                                                                        case "":
                                                                                                                            var v54 uint64 = 2
                                                                                                                            _ = v54
                                                                                                                        case "\n\t":
                                                                                                                            var v55 uint64 = v23
                                                                                                                            _ = v55
                                                                                                                        case "foo bar":
                                                                                                                            if false {
                                                                                                                                var v56 string = "a"
                                                                                                                                _ = v56
                                                                                                                            }
                                                                                                                        default:
                                                                                                                            var v57 int32 = (v6 % ((v6) | 1))
                                                                                                                            _ = v57
                                                                                                                        }
                                                                                                                    } else {
                                                                                                                        switch fi32 {
                                                                                                                        case 367:
                                                                                                                            fmt.Println(v4)
                                                                                                                        }
                                                                                                                        delete(v2, (fi32 % (((^(fi31))) | 1)))
                                                                                                                        var v58 uint64 = uint64(v24)
                                                                                                                        if v25 {
                                                                                                                            v1 = (v30 & (^(v30)))
                                                                                                                        }
                                                                                                                        v59 := GenS7{F0: v1, F1: (v6 / ((v6) | 1))}
                                                                                                                        _ = v58
                                                                                                                        _ = v59
                                                                                                                    }
                                                                                                                }
                                                                                                            }
                                                                                                            if (("" + "a") <= "") {
                                                                                                                v24 -= (v24 >> ((4294967295) & 63))
                                                                                                                var v60 uint = 2
                                                                                                                _ = v60
                                                                                                            } else {
                                                                                                                var v61 int8 = 60
                                                                                                                _ = v61
                                                                                                            }
                                                                                                            _ = v45
                                                                                                        }
                                                                                                        _ = v44
                                                                                                    }
                                                                                                    v62 := GenS6{F0: 556, F1: v6, F2: (v37 << ((uint(v1)) & 63))}
                                                                                                    _ = v62
                                                                                                } else {
                                                                                                    v6 = v6
                                                                                                    var v63 uint16 = 255
                                                                                                    var v64 uint64 = uint64(v24)
                                                                                                    fmt.Println(v24)
                                                                                                    switch "" {
                                                                                                    case "\n\t":
                                                                                                        v65 := GenS3{F0: ("a" + ("foo bar" + ""))}
                                                                                                        _ = v65
                                                                                                    case "hello":
                                                                                                        if true {
                                                                                                            fmt.Println(v28, v6)
                                                                                                            var v66 int64 = (^(v40))
                                                                                                            v24 -= (v24 >> ((4294967295) & 63))
                                                                                                            _ = v66
                                                                                                        } else {
                                                                                                            var v67 bool = v25
                                                                                                            if ((v40 * v40) < int64(v6)) {
                                                                                                                v68 := GenS6{F0: (v24 << ((uint(v64)) & 63)), F1: 1, F2: (v37 >> ((2) & 63))}
                                                                                                                var v69 int = int(v63)
                                                                                                                _ = v68
                                                                                                                _ = v69
                                                                                                            } else {
                                                                                                                v2[fi32] = ((!(!((v6 ^ int32(v30)) < v6))) || v25)
                                                                                                                fmt.Println(v4, fi31, v67)
                                                                                                                v70 := GenS7{F0: 111, F1: (v6 | -583)}
                                                                                                                fmt.Println(v64, v67, v4)
                                                                                                                _ = v70
                                                                                                            }
                                                                                                            fmt.Println(v25, v28, v4)
                                                                                                            if (false && (true && ((!false) || (v40 != v40)))) {
                                                                                                                v23 += (v4 / ((197) | 1))
                                                                                                                var v71 int16 = 42
                                                                                                                if ((v37 - (^(v37))) >= 42) {
                                                                                                                    var v72 float64 = (1e10 * (v5 / (0.0 - 3.14159)))
                                                                                                                    _ = v72
                                                                                                                }
                                                                                                                _ = v71
                                                                                                            }
                                                                                                            _ = v67
                                                                                                        }
                                                                                                    }
                                                                                                    _ = v63
                                                                                                    _ = v64
                                                                                                }
                                                                                                var v73 uint = 18446744073709551615
                                                                                                switch (fi31 * (^(fi31))) {
                                                                                                case 32768:
                                                                                                    var v74 uint = uint(v23)
                                                                                                    fmt.Println(v1, v24)
                                                                                                    _ = v74
                                                                                                case 2147483647:
                                                                                                    fmt.Println(v1, v23, v6)
                                                                                                case 4294967295:
                                                                                                    _ = len(v26)
                                                                                                default:
                                                                                                    switch (-(fi32)) {
                                                                                                    case 4294967295:
                                                                                                        v75 := v2[fi32]
                                                                                                        _ = v75
                                                                                                    default:
                                                                                                        var v76 uint8 = (v1 * v30)
                                                                                                        _ = v76
                                                                                                    }
                                                                                                }
                                                                                                _ = v73
                                                                                            }
                                                                                        case -749:
                                                                                            if ((-2147483648 == v6) && ("\n\t" < "a")) {
                                                                                                v77 := GenS1{F0: "\n\t"}
                                                                                                _ = v77
                                                                                            } else {
                                                                                                if len(v26) > 0 {
                                                                                                    v26[0] = (-(fi32))
                                                                                                }
                                                                                                fmt.Println(v25, v40, v4)
                                                                                                var v78 float64 = v5
                                                                                                var v79 uint8 = 110
                                                                                                _ = v78
                                                                                                _ = v79
                                                                                            }
                                                                                        default:
                                                                                            switch 32768 {
                                                                                            case 1:
                                                                                                if ((!v25) && true) {
                                                                                                    var v80 uint8 = v30
                                                                                                    v81 := GenS1{F0: "hello"}
                                                                                                    if len(v26) > 0 {
                                                                                                        v26[0] = genFn0()
                                                                                                    }
                                                                                                    var v82 float32 = -2.25
                                                                                                    _ = v80
                                                                                                    _ = v81
                                                                                                    _ = v82
                                                                                                }
                                                                                            case -1:
                                                                                                v4 -= v23
                                                                                            case -131:
                                                                                                if (502 <= 42) {
                                                                                                    var v83 int64 = int64(fi32)
                                                                                                    fmt.Println(v4, v28, v30)
                                                                                                    v84 := GenS2{F0: v27, F1: v30, F2: genFn1(v27, "a")}
                                                                                                    var v85 int64 = v83
                                                                                                    v4 = (v23 % ((503) | 1))
                                                                                                    var v86 uint32 = v24
                                                                                                    if len(v26) > 0 {
                                                                                                        v26[0] = fi31
                                                                                                    }
                                                                                                    var v87 uint16 = 233
                                                                                                    _ = v83
                                                                                                    _ = v84
                                                                                                    _ = v85
                                                                                                    _ = v86
                                                                                                    _ = v87
                                                                                                }
                                                                                            default:
                                                                                                fmt.Println(v6, v23)
                                                                                            }
                                                                                        }
                                                                                    default:
                                                                                        v5 = ((((v5 + v5) + ((-2.25 * v5) + v5)) / (float64((((((float32((float64(-2.25) / float64((-2.25 / 3.14159)))) * (((((v28 * v28) - v28) + ((1e10 / (v28 + v28)) - v28)) - v28) / v28)) + ((v28 / 0.1) * v28)) + v28) / 0.1) - (float32(v5) / (0.0 / (float32(v5) - 0.1))))) / ((1.5 - 1.5) / (v5 + 1e10)))) - v5)
                                                                                        v23 -= (v23 % (((^(v23))) | 1))
                                                                                    }
                                                                                }
                                                                                var v88 int32 = 2147483647
                                                                                v89 := GenS2{F0: (v27 + (v27 + 1)), F1: (^(v30)), F2: genFn1((v27 ^ v27), (("\n\t" + "\n\t") + ("foo bar" + "hello")))}
                                                                                _ = v43
                                                                                _ = v88
                                                                                _ = v89
                                                                            }
                                                                            v90 := GenS8{F0: 18446744073709551615, F1: fi32, F2: v30}
                                                                            var v91 uint8 = v1
                                                                            _ = v90
                                                                            _ = v91
                                                                        } else {
                                                                            v4 = 42
                                                                            switch "foo bar" {
                                                                            case "hello":
                                                                                v24 = v24
                                                                                var v92 bool = true
                                                                                _ = v92
                                                                            case "foo bar":
                                                                                fmt.Println(fi32, fi31)
                                                                            case "\n\t":
                                                                                v93 := v3[("\n\t" + "\n\t")]
                                                                                _ = v93
                                                                            default:
                                                                                if (false && v25) {
                                                                                    v94 := GenS4{F0: fi32, F1: (v24 ^ v24), F2: 0}
                                                                                    var v95 int8 = -23
                                                                                    _ = v94
                                                                                    _ = v95
                                                                                } else {
                                                                                    switch "hello" {
                                                                                    case "foo bar":
                                                                                        v96 := GenS9{F0: 592, F1: v25}
                                                                                        _ = v96
                                                                                    default:
                                                                                        var v97 uint64 = 4294967295
                                                                                        _ = v97
                                                                                    }
                                                                                    v30 = 2
                                                                                    v98 := GenS5{F0: (v27 & -313)}
                                                                                    _ = v98
                                                                                }
                                                                            }
                                                                            fmt.Println(v30, v40)
                                                                        }
                                                                    }
                                                                    if ((false && false) && true) {
                                                                        _ = len(v3)
                                                                        switch "" {
                                                                        case "hello":
                                                                            v99 := GenS10{F0: (fi31 ^ genFn0())}
                                                                            _ = v99
                                                                        }
                                                                        v100 := GenS11{F0: (v23 ^ (^(v23)))}
                                                                        _ = v100
                                                                    }
                                                                    if false {
                                                                        switch "foo bar" {
                                                                        case "hello":
                                                                            if false {
                                                                                var v101 int64 = (v40 - (v40 & v40))
                                                                                var v102 uint16 = 0
                                                                                var v103 uint = 255
                                                                                var v104 int64 = (v101 >> ((uint(v24)) & 63))
                                                                                _ = v101
                                                                                _ = v102
                                                                                _ = v103
                                                                                _ = v104
                                                                            } else {
                                                                                v26 = append(v26, (-(fi31)))
                                                                                var v105 string = ("a" + "a")
                                                                                fmt.Println(v6)
                                                                                _ = v105
                                                                            }
                                                                            switch int(v1) {
                                                                            case -1:
                                                                                v37 += (v37 / ((v37) | 1))
                                                                            case 117:
                                                                                switch fi32 {
                                                                                case 1:
                                                                                    fmt.Println(v1, v25)
                                                                                }
                                                                            case -535:
                                                                                fmt.Println(v30, fi31)
                                                                                if v25 {
                                                                                    v106 := v2[-1]
                                                                                    var v107 int8 = -128
                                                                                    _ = v106
                                                                                    _ = v107
                                                                                } else {
                                                                                    fmt.Println(v5, v28, v1)
                                                                                }
                                                                                switch "foo bar" {
                                                                                case "hello":
                                                                                    v108 := GenS12{F0: (v27 | v27), F1: 255, F2: (-(v37))}
                                                                                    _ = v108
                                                                                case "":
                                                                                    v109 := GenS12{F0: 0, F1: 0, F2: (^(v37))}
                                                                                    _ = v109
                                                                                default:
                                                                                    _ = len(v26)
                                                                                }
                                                                                var v110 uint64 = (^(v4))
                                                                                _ = v110
                                                                            default:
                                                                                v4 = (v23 | (v23 % ((18446744073709551615) | 1)))
                                                                            }
                                                                        case "a":
                                                                            fmt.Println(v40, fi31, v28)
                                                                        }
                                                                        v111 := GenS13{F0: (v27 | v27), F1: (-(v40)), F2: ("foo bar" + ("a" + "foo bar"))}
                                                                        _ = v111
                                                                    } else {
                                                                        var v112 float64 = (((float64(v28) - float64(v28)) * ((0.0 / (float64(v28) + 0.1)) + v5)) * float64(3.14159))
                                                                        _ = v112
                                                                    }
                                                                } else {
                                                                    switch "" {
                                                                    case "":
                                                                        fmt.Println(fi31, v1)
                                                                    default:
                                                                        if v25 {
                                                                            v5 = (v5 / (1e10 + float64(v28)))
                                                                            switch "hello" {
                                                                            case "a":
                                                                                v113 := GenS8{F0: 4294967295, F1: (-(fi31)), F2: v1}
                                                                                _ = v113
                                                                            case "hello":
                                                                                delete(v3, "a")
                                                                            }
                                                                            var v114 bool = (v25 && ((-(v37)) > (v37 % ((int8(v40)) | 1))))
                                                                            _ = v114
                                                                        }
                                                                    }
                                                                    v1 = 1
                                                                    var v115 uint64 = v23
                                                                    v116 := GenS14{F0: v6, F1: int16(v6), F2: (v27 >> ((255) & 63))}
                                                                    _ = v115
                                                                    _ = v116
                                                                }
                                                                _ = v40
                                                            }
                                                        case "\n\t":
                                                            fmt.Println(v27, v6, v23)
                                                        default:
                                                            _ = len(v2)
                                                        }
                                                        var v117 uint = 807
                                                        _ = v117
                                                    } else {
                                                        v4 -= v4
                                                    }
                                                    if ((((v6 << ((255) & 63)) < (v6 % ((v6) | 1))) && false) || false) {
                                                        _ = len(v26)
                                                    } else {
                                                        v118 := GenS4{F0: fi31, F1: (^(v24)), F2: v27}
                                                        var v119 int = 478
                                                        if v25 {
                                                            _ = len(v2)
                                                            var v120 float64 = v5
                                                            _ = v120
                                                        }
                                                        _ = v118
                                                        _ = v119
                                                    }
                                                    _ = v37
                                                }
                                                if v25 {
                                                    if false {
                                                        var v121 bool = v25
                                                        _ = v121
                                                    } else {
                                                        var v122 bool = (42 > 935)
                                                        v123 := GenS5{F0: (v27 * (v27 * v27))}
                                                        _ = v122
                                                        _ = v123
                                                    }
                                                }
                                                v124 := GenS12{F0: 42, F1: 834, F2: 100}
                                                _ = v124
                                            }
                                        case 127:
                                            fmt.Println(v28, v27)
                                        case 128:
                                            v5 += v5
                                        default:
                                            v1 += v1
                                        }
                                        v2[fi32] = (!v25)
                                        fmt.Println(v23)
                                        var v125 float32 = float32(0.0)
                                        var v126 uint16 = 991
                                        var v127 int32 = v6
                                        _ = v125
                                        _ = v126
                                        _ = v127
                                    } else {
                                        _ = len(v2)
                                        v26 = append(v26, genFn0())
                                        var v128 uint = 373
                                        _ = v128
                                    }
                                }
                            }
                        }
                    }
                    var v129 int64 = int64(v23)
                    var v130 int8 = 42
                    _ = fi32
                    _ = v129
                    _ = v130
                }
                v24 -= (v24 * v24)
                _ = fi31
            }
            v5 = v5
            _ = v30
        }
        _ = v27
        _ = v28
    }
    if (65535 > 2) {
        var v131 int16 = 0
        v6 = (v6 * (v6 % ((int32(v23)) | 1)))
        _ = v131
    } else {
        var v132 uint64 = 851
        _ = v132
    }
    v4 = v23
    v6 = v6
    var v133 float64 = float64(1.5)
    v134 := GenS15{F0: v24, F1: 255}
    var v135 uint32 = (^(v24))
    var v136 int16 = -1
    fmt.Println(v6, v1, v136)
    v25 = v25
    fmt.Println("v1:", v1)
    fmt.Println("v133:", v133)
    fmt.Println("v135:", v135)
    fmt.Println("v136:", v136)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v26:", len(v26), v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v133
    _ = v134
    _ = v135
    _ = v136
    _ = v2
    _ = v23
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
