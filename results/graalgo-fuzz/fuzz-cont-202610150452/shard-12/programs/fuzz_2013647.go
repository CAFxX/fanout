package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8; F1 int16; F2 bool }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int64; F1 uint16; F2 uint8 }
type GenS6 struct { F0 uint64; F1 int }
type GenS7 struct { F0 uint; F1 bool }
type GenS8 struct { F0 uint16 }
type GenS9 struct { F0 uint32; F1 string }
type GenS10 struct { F0 int; F1 uint8; F2 int16 }
type GenS11 struct { F0 int; F1 uint }
type GenS12 struct { F0 int64 }
type GenS13 struct { F0 uint16; F1 int64; F2 string }
type GenS14 struct { F0 uint32 }
type GenS15 struct { F0 uint32; F1 string }
type GenS16 struct { F0 int32 }
type GenS17 struct { F0 int16 }
type GenS18 struct { F0 int64; F1 bool }
type GenS19 struct { F0 int16; F1 uint8 }
type GenS20 struct { F0 uint32; F1 uint64 }
type GenS21 struct { F0 int16; F1 int16 }
type GenS22 struct { F0 uint64; F1 int64; F2 bool }
type GenS23 struct { F0 int32; F1 bool }
type GenS24 struct { F0 int16 }
type GenS25 struct { F0 int64 }
type GenS26 struct { F0 int32 }
type GenS27 struct { F0 uint8; F1 string }
type GenS28 struct { F0 string; F1 uint8 }

func main() {
    var v1 int8 = 75
    var v2 int32 = -2147483648
    var v3 uint16 = 42
    v4 := []string{("a" + ("foo bar" + "")), "foo bar"}
    var v5 uint32 = 4294967295
    for fi6 := 0; fi6 < 20; fi6++ {
        switch (fi6 ^ fi6) {
        case -606:
            v7 := GenS1{F0: v2}
            _ = v7
        case -512:
            var v8 int16 = 1
            _ = v8
        case -444:
            v9 := GenS1{F0: (v2 ^ int32(fi6))}
            _ = v9
        }
        v10 := GenS2{F0: 127, F1: 32767, F2: true}
        _ = fi6
        _ = v10
    }
    v5 += 998
    v3 -= 65535
    if true {
        for fi11 := 0; fi11 < 18; fi11++ {
            v3 = 65535
            if false {
                for fi12 := 0; fi12 < 4; fi12++ {
                    var v13 uint64 = 4294967295
                    var v14 string = ""
                    _ = fi12
                    _ = v13
                    _ = v14
                }
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            }
            fmt.Println(v3, v1, v3)
            _ = fi11
        }
        var v15 uint8 = 2
        var v16 float32 = ((-2.25 - 0.0) + (0.0 + 3.14159))
        _ = v15
        _ = v16
    }
    if (v2 <= v2) {
        var v17 float32 = (1e10 / 1e10)
        _ = v17
    }
    fmt.Println(v1)
    var v18 int = 128
    if false {
        for fi19 := 0; fi19 < 7; fi19++ {
            fmt.Println(v18, v18)
            var v20 bool = (v5 >= v5)
            v1 += (v1 >> ((757) & 63))
            _ = fi19
            _ = v20
        }
        for fi21 := 0; fi21 < 6; fi21++ {
            if true {
                v4 = append(v4, "hello")
            }
            var v22 int64 = -1
            if true {
                var v23 uint8 = 1
                _ = v23
            } else {
                for fi24 := 0; fi24 < 10; fi24++ {
                    switch (v18 % (((fi21 - (v18 / (((fi24 / ((v18) | 1))) | 1)))) | 1)) {
                    case 32767:
                        if ((v5 / ((v5) | 1)) != 23) {
                            fmt.Println(fi24)
                            _ = len(v4)
                        } else {
                            switch "hello" {
                            case "a":
                                v25 := GenS3{F0: true}
                                _ = v25
                            case "":
                                fmt.Println(v22)
                            }
                            var v26 float64 = ((-2.25 + 1.5) + float64(0.1))
                            _ = v26
                        }
                        if (v1 <= (v1 / ((v1) | 1))) {
                            switch (fi21 % (((v18 ^ (fi24 - 0))) | 1)) {
                            case 9223372036854775807:
                                v1 -= 6
                            case 127:
                                if false {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    var v27 float64 = 1e10
                                    if len(v4) > 0 {
                                        v4[0] = "foo bar"
                                    }
                                    _ = v27
                                } else {
                                    switch "hello" {
                                    case "a":
                                        var v28 string = (("foo bar" + "hello") + ((("a" + "\n\t") + (("" + "") + ("\n\t" + ""))) + "foo bar"))
                                        _ = v28
                                    case "hello":
                                        if true {
                                            if (((3.14159 + float32(((0.1 * (-2.25 / 1.0)) + (-2.25 - 1.5)))) - 1e10) < (-2.25 / -2.25)) {
                                                var v29 int8 = (v1 - (-(v1)))
                                                v22 += (v22 + (v22 % (((-(v22))) | 1)))
                                                var v30 int32 = (v2 - 0)
                                                v29 = (v29 | (v1 & v1))
                                                if false {
                                                    fmt.Println(fi21)
                                                    if false {
                                                        fmt.Println(v5, v2)
                                                        switch (fi24 | 127) {
                                                        case 1:
                                                            v4 = append(v4, "hello")
                                                        default:
                                                            fmt.Println(v5, v29, fi24)
                                                        }
                                                    }
                                                    v31 := GenS4{F0: 0}
                                                    var v32 uint16 = (v3 % ((364) | 1))
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        fmt.Println(v1)
                                                    default:
                                                        if false {
                                                            var v33 uint32 = v5
                                                            _ = v33
                                                        } else {
                                                            switch (("foo bar" + "hello") + "a") {
                                                            case "hello":
                                                                if ((-(v1)) <= 100) {
                                                                    fmt.Println(v30, v2)
                                                                    v18 = fi24
                                                                }
                                                                v22 = v22
                                                                fmt.Println(v22)
                                                            case "\n\t":
                                                                var v34 float64 = (0.1 + float64(float32(1.5)))
                                                                _ = v34
                                                            default:
                                                                var v35 int32 = (v30 + v30)
                                                                _ = v35
                                                            }
                                                            var v36 float64 = 3.14159
                                                            switch (v18 & (fi21 / (((^(v18))) | 1))) {
                                                            case 342:
                                                                var v37 int16 = 32767
                                                                _ = v37
                                                            case -984:
                                                                var v38 uint8 = 127
                                                                _ = v38
                                                            case 2147483647:
                                                                fmt.Println(v30, fi24, fi21)
                                                            }
                                                            _ = v36
                                                        }
                                                    }
                                                    fmt.Println(v29, v2)
                                                    _ = v31
                                                    _ = v32
                                                }
                                                _ = v29
                                                _ = v30
                                            } else {
                                                var v39 float64 = (((-2.25 - 0.0) - (float64((0.0 - 1e10)) / 1e10)) / 1.5)
                                                var v40 uint32 = v5
                                                switch fi24 {
                                                case 128:
                                                    if (v2 < v2) {
                                                        v4 = append(v4, "")
                                                        if false {
                                                            var v41 uint16 = 2
                                                            _ = v41
                                                        } else {
                                                            var v42 uint16 = (v3 & uint16(v5))
                                                            v5 -= (^(v40))
                                                            _ = v42
                                                        }
                                                    }
                                                case -1:
                                                    v2 = (v2 >> ((2) & 63))
                                                default:
                                                    v5 = 42
                                                }
                                                switch "\n\t" {
                                                case "\n\t":
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                }
                                                var v43 float64 = -2.25
                                                var v44 uint16 = uint16(v40)
                                                v3 = 173
                                                _ = v39
                                                _ = v40
                                                _ = v43
                                                _ = v44
                                            }
                                        } else {
                                            _ = len(v4)
                                        }
                                    case "foo bar":
                                        switch (v18 * v18) {
                                        case 128:
                                            if (!false) {
                                                switch -9223372036854775808 {
                                                case -1:
                                                    var v45 bool = true
                                                    _ = v45
                                                case 9223372036854775807:
                                                    v4 = append(v4, "hello")
                                                case 32767:
                                                    v46 := GenS5{F0: v22, F1: (v3 % (((v3 / ((1) | 1))) | 1)), F2: 1}
                                                    _ = v46
                                                default:
                                                    var v47 float64 = ((0.1 - 0.1) + 0.1)
                                                    _ = v47
                                                }
                                            }
                                        case 32767:
                                            if true {
                                                var v48 float64 = (3.14159 / -2.25)
                                                switch fi24 {
                                                case -20:
                                                    v49 := GenS2{F0: 0, F1: int16(v5), F2: false}
                                                    _ = v49
                                                default:
                                                    var v50 int32 = 0
                                                    _ = v50
                                                }
                                                if len(v4) > 0 {
                                                    v4[0] = (("a" + "hello") + "")
                                                }
                                                _ = v48
                                            }
                                        case 9223372036854775807:
                                            v2 -= v2
                                            var v51 uint16 = (v3 * v3)
                                            _ = v51
                                        }
                                    }
                                    switch 530 {
                                    case -834:
                                        v52 := GenS4{F0: 32767}
                                        _ = v52
                                    default:
                                        switch (fi21 ^ v18) {
                                        case 127:
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        case 32767:
                                            if false {
                                                if true {
                                                    switch "foo bar" {
                                                    case "":
                                                        fmt.Println(fi24)
                                                    case "foo bar":
                                                        if (!(18446744073709551615 >= 42)) {
                                                            if len(v4) > 0 {
                                                                v4[0] = "foo bar"
                                                            }
                                                            var v53 uint32 = v5
                                                            var v54 uint16 = (v3 / (((^(v3))) | 1))
                                                            switch "\n\t" {
                                                            case "hello":
                                                                switch (("foo bar" + ("a" + ("a" + "a"))) + ("" + "\n\t")) {
                                                                case "a":
                                                                    switch (v18 % ((fi21) | 1)) {
                                                                    case 127:
                                                                        if false {
                                                                            fmt.Println(v2)
                                                                            v55 := GenS6{F0: 42, F1: int(v5)}
                                                                            v56 := GenS6{F0: 42, F1: (fi21 % ((fi24) | 1))}
                                                                            _ = v55
                                                                            _ = v56
                                                                        }
                                                                    }
                                                                case "foo bar":
                                                                    if (false && true) {
                                                                        if false {
                                                                            v57 := GenS7{F0: 18446744073709551615, F1: false}
                                                                            var v58 float32 = 3.14159
                                                                            if len(v4) > 0 {
                                                                                v4[0] = "\n\t"
                                                                            }
                                                                            _ = v57
                                                                            _ = v58
                                                                        }
                                                                        var v59 string = "\n\t"
                                                                        _ = v59
                                                                    } else {
                                                                        v60 := GenS8{F0: v54}
                                                                        fmt.Println(v5)
                                                                        _ = v60
                                                                    }
                                                                case "hello":
                                                                    v3 = v54
                                                                }
                                                            case "":
                                                                v1 = v1
                                                            }
                                                            _ = v53
                                                            _ = v54
                                                        }
                                                    case "hello":
                                                        v5 = v5
                                                    }
                                                    v5 = (v5 + 0)
                                                    v22 += v22
                                                }
                                                _ = len(v4)
                                                var v61 uint64 = 42
                                                _ = v61
                                            } else {
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                v62 := GenS9{F0: (^(v5)), F1: "\n\t"}
                                                _ = v62
                                            }
                                        case -412:
                                            var v63 int8 = 1
                                            _ = v63
                                        }
                                    }
                                    switch "" {
                                    case "hello":
                                        var v64 int16 = 1000
                                        _ = v64
                                    case "\n\t":
                                        v1 = -98
                                    case "foo bar":
                                        switch fi21 {
                                        case 127:
                                            v2 = 42
                                        case 809:
                                            var v65 uint32 = (v5 / ((2147483647) | 1))
                                            _ = v65
                                        default:
                                            if false {
                                                if ((v18 & v18) < (fi21 ^ (fi21 - 32768))) {
                                                    fmt.Println(v22, v22)
                                                    _ = len(v4)
                                                    if (false || false) {
                                                        v66 := GenS6{F0: 4294967295, F1: (v18 + (^(v18)))}
                                                        if (v22 != 10) {
                                                            var v67 int16 = 1000
                                                            if len(v4) > 0 {
                                                                v4[0] = "a"
                                                            }
                                                            _ = v67
                                                        } else {
                                                            var v68 uint = 490
                                                            fmt.Println(fi21, v2)
                                                            if (fi21 > -710) {
                                                                if len(v4) > 0 {
                                                                    v4[0] = "\n\t"
                                                                }
                                                            } else {
                                                                switch "hello" {
                                                                case "":
                                                                    if (!(!(-71 < v1))) {
                                                                        if len(v4) >= 2 {
                                                                            v4 = v4[1:len(v4)]
                                                                        }
                                                                    } else {
                                                                        if len(v4) > 0 {
                                                                            v4[0] = "\n\t"
                                                                        }
                                                                        var v69 uint64 = 18446744073709551615
                                                                        var v70 uint16 = (v3 & v3)
                                                                        var v71 int8 = 100
                                                                        fmt.Println(v18)
                                                                        if true {
                                                                            v5 = (v5 % (((^(v5))) | 1))
                                                                            if (!false) {
                                                                                v72 := GenS1{F0: v2}
                                                                                _ = v72
                                                                            }
                                                                            if false {
                                                                                switch ("foo bar" + "\n\t") {
                                                                                case "hello":
                                                                                    if len(v4) >= 2 {
                                                                                        v4 = v4[1:len(v4)]
                                                                                    }
                                                                                }
                                                                                _ = len(v4)
                                                                            }
                                                                            switch fi24 {
                                                                            case 128:
                                                                                if (2 == v5) {
                                                                                    if (!true) {
                                                                                        v2 -= v2
                                                                                        v73 := GenS10{F0: v18, F1: uint8(v68), F2: 32767}
                                                                                        _ = v73
                                                                                    } else {
                                                                                        var v74 uint16 = v70
                                                                                        switch "\n\t" {
                                                                                        case "":
                                                                                            var v75 uint = v68
                                                                                            _ = v75
                                                                                        }
                                                                                        _ = v74
                                                                                    }
                                                                                    v76 := GenS11{F0: (-(fi21)), F1: v68}
                                                                                    _ = v76
                                                                                } else {
                                                                                    v68 -= 42
                                                                                    var v77 uint = (v68 ^ (v68 & (v68 | (v68 >> (((v68 >> (((v68 * 0)) & 63))) & 63)))))
                                                                                    v77 = v68
                                                                                    _ = v77
                                                                                }
                                                                            default:
                                                                                var v78 uint16 = v3
                                                                                v2 -= v2
                                                                                v78 = v70
                                                                                _ = v78
                                                                            }
                                                                        } else {
                                                                            if true {
                                                                                var v79 uint32 = v5
                                                                                if true {
                                                                                    if (true && (!(!(71 == (^(v18)))))) {
                                                                                        v80 := GenS12{F0: (v22 / ((int64(v68)) | 1))}
                                                                                        var v81 int64 = 42
                                                                                        var v82 int = int(v2)
                                                                                        _ = v80
                                                                                        _ = v81
                                                                                        _ = v82
                                                                                    } else {
                                                                                        var v83 int32 = -2147483648
                                                                                        v84 := GenS12{F0: v22}
                                                                                        switch fi21 {
                                                                                        case 741:
                                                                                            _ = len(v4)
                                                                                        case 127:
                                                                                            if false {
                                                                                                if true {
                                                                                                    if (0 > 2) {
                                                                                                        v4 = append(v4, "a")
                                                                                                        v22 = (v22 & (v22 >> (((v68 / ((v68) | 1))) & 63)))
                                                                                                        if len(v4) >= 2 {
                                                                                                            v4 = v4[1:len(v4)]
                                                                                                        }
                                                                                                    } else {
                                                                                                        v79 -= 42
                                                                                                        if ((!(int(v1) <= v18)) || ((!true) || (v1 != -1))) {
                                                                                                            var v85 string = ((("a" + ((("a" + "\n\t") + ("foo bar" + (("\n\t" + "\n\t") + "a"))) + "hello")) + "") + "")
                                                                                                            var v86 int16 = 957
                                                                                                            _ = v85
                                                                                                            _ = v86
                                                                                                        }
                                                                                                        v87 := GenS1{F0: v83}
                                                                                                        switch (fi24 * fi21) {
                                                                                                        case 128:
                                                                                                            v4 = append(v4, ((("foo bar" + ("hello" + "\n\t")) + "a") + "hello"))
                                                                                                        case 739:
                                                                                                            fmt.Println(v1, fi21, v71)
                                                                                                        case 1:
                                                                                                            fmt.Println(v5, v68)
                                                                                                        default:
                                                                                                            v88 := GenS13{F0: v70, F1: v22, F2: ("" + "foo bar")}
                                                                                                            if len(v4) >= 2 {
                                                                                                                v4 = v4[1:len(v4)]
                                                                                                            }
                                                                                                            _ = v88
                                                                                                        }
                                                                                                        v89 := GenS1{F0: v83}
                                                                                                        _ = v87
                                                                                                        _ = v89
                                                                                                    }
                                                                                                    fmt.Println(v2)
                                                                                                } else {
                                                                                                    v5 += 2
                                                                                                    var v90 int8 = int8(v70)
                                                                                                    if (!(10 > 127)) {
                                                                                                        var v91 int32 = (v83 % ((509) | 1))
                                                                                                        v92 := GenS8{F0: 65535}
                                                                                                        fmt.Println(v3)
                                                                                                        v93 := GenS2{F0: 42, F1: 42, F2: true}
                                                                                                        v94 := GenS13{F0: 1, F1: (v22 - (v22 | v22)), F2: "foo bar"}
                                                                                                        var v95 uint = (v68 % ((875) | 1))
                                                                                                        _ = v91
                                                                                                        _ = v92
                                                                                                        _ = v93
                                                                                                        _ = v94
                                                                                                        _ = v95
                                                                                                    } else {
                                                                                                        switch "foo bar" {
                                                                                                        case "a":
                                                                                                            var v96 int8 = int8(v5)
                                                                                                            _ = v96
                                                                                                        default:
                                                                                                            v83 = 1
                                                                                                            fmt.Println(v1, v18)
                                                                                                            if (!false) {
                                                                                                                fmt.Println(v70, v3)
                                                                                                            } else {
                                                                                                                if (true && ("a" == "foo bar")) {
                                                                                                                    v97 := GenS7{F0: v68, F1: true}
                                                                                                                    _ = v97
                                                                                                                } else {
                                                                                                                    v22 = 42
                                                                                                                }
                                                                                                            }
                                                                                                            v22 -= (v22 / (((v22 & (v22 | (-(v22))))) | 1))
                                                                                                            v18 = fi24
                                                                                                            var v98 uint64 = (v69 + (v69 - 18446744073709551615))
                                                                                                            var v99 int = (v18 - v18)
                                                                                                            var v100 uint8 = 2
                                                                                                            _ = v100
                                                                                                            _ = v98
                                                                                                            _ = v99
                                                                                                        }
                                                                                                        v70 += v70
                                                                                                        var v101 uint64 = v69
                                                                                                        _ = v101
                                                                                                    }
                                                                                                    fmt.Println(v2, v18)
                                                                                                    v5 -= (v79 << ((v68) & 63))
                                                                                                    v3 = (v3 + v3)
                                                                                                    fmt.Println(v69, v5, v1)
                                                                                                    if len(v4) > 0 {
                                                                                                        v4[0] = "foo bar"
                                                                                                    }
                                                                                                    _ = v90
                                                                                                }
                                                                                                v69 -= 312
                                                                                            } else {
                                                                                                if false {
                                                                                                    v102 := GenS10{F0: (fi24 % ((int(v3)) | 1)), F1: 1, F2: 1}
                                                                                                    fmt.Println(v1, v83)
                                                                                                    var v103 string = "foo bar"
                                                                                                    _ = v102
                                                                                                    _ = v103
                                                                                                } else {
                                                                                                    switch "" {
                                                                                                    case "a":
                                                                                                        if (v5 <= (^(v5))) {
                                                                                                            var v104 uint16 = v3
                                                                                                            v4 = append(v4, "\n\t")
                                                                                                            var v105 uint = 255
                                                                                                            _ = v104
                                                                                                            _ = v105
                                                                                                        }
                                                                                                    case "foo bar":
                                                                                                        if len(v4) >= 2 {
                                                                                                            v4 = v4[1:len(v4)]
                                                                                                        }
                                                                                                    case "":
                                                                                                        switch 127 {
                                                                                                        case -267:
                                                                                                            if len(v4) >= 2 {
                                                                                                                v4 = v4[1:len(v4)]
                                                                                                            }
                                                                                                        case -2147483648:
                                                                                                            if len(v4) > 0 {
                                                                                                                v4[0] = ""
                                                                                                            }
                                                                                                        case -399:
                                                                                                            _ = len(v4)
                                                                                                        default:
                                                                                                            v106 := GenS14{F0: 1}
                                                                                                            _ = v106
                                                                                                        }
                                                                                                    default:
                                                                                                        v18 = (fi21 << (((v68 ^ 255)) & 63))
                                                                                                    }
                                                                                                    v2 -= (^(v83))
                                                                                                }
                                                                                            }
                                                                                        case 2147483647:
                                                                                            if len(v4) > 0 {
                                                                                                v4[0] = "\n\t"
                                                                                            }
                                                                                        default:
                                                                                            var v107 bool = ((v69 >> ((v68) & 63)) < uint64(v79))
                                                                                            v2 -= int32(v22)
                                                                                            _ = v107
                                                                                        }
                                                                                        _ = v83
                                                                                        _ = v84
                                                                                    }
                                                                                    v71 = v71
                                                                                    v4 = append(v4, "hello")
                                                                                } else {
                                                                                    var v108 uint32 = v5
                                                                                    _ = v108
                                                                                }
                                                                                v109 := GenS15{F0: 65535, F1: "\n\t"}
                                                                                if (false || true) {
                                                                                    fmt.Println(v70, v22, v22)
                                                                                    v5 -= (v5 - v5)
                                                                                }
                                                                                _ = v109
                                                                                _ = v79
                                                                            } else {
                                                                                v70 -= v70
                                                                                v110 := GenS3{F0: false}
                                                                                _ = v110
                                                                            }
                                                                            _ = len(v4)
                                                                        }
                                                                        _ = v69
                                                                        _ = v70
                                                                        _ = v71
                                                                    }
                                                                case "a":
                                                                    _ = len(v4)
                                                                case "foo bar":
                                                                    v68 -= 695
                                                                }
                                                                v22 = -9223372036854775808
                                                                v111 := GenS11{F0: 2147483647, F1: v68}
                                                                _ = v111
                                                            }
                                                            _ = v68
                                                        }
                                                        v2 -= -2147483648
                                                        _ = v66
                                                    }
                                                }
                                                v2 += -2147483648
                                                fmt.Println(v18)
                                            }
                                        }
                                    }
                                }
                            case 975:
                                v112 := GenS14{F0: v5}
                                _ = v112
                            }
                            if true {
                                _ = len(v4)
                            }
                            var v113 int64 = -195
                            v1 -= v1
                            if true {
                                if (0 > (^(v2))) {
                                    if (false && false) {
                                        switch ("a" + "\n\t") {
                                        case "a":
                                            var v114 float32 = 0.0
                                            _ = v114
                                        case "foo bar":
                                            if false {
                                                v22 += v22
                                                v2 -= (v2 * v2)
                                                if (false && ((3.14159 > (0.0 / 3.14159)) && false)) {
                                                    var v115 int8 = 0
                                                    fmt.Println(v2, fi24)
                                                    v22 += v22
                                                    _ = v115
                                                }
                                                _ = len(v4)
                                            }
                                        }
                                        var v116 string = "a"
                                        _ = v116
                                    } else {
                                        v5 -= v5
                                        v3 -= (^(v3))
                                        var v117 int64 = 0
                                        _ = v117
                                    }
                                    v118 := GenS5{F0: 0, F1: (v3 % ((v3) | 1)), F2: 42}
                                    _ = v118
                                } else {
                                    v22 += v113
                                    v119 := GenS16{F0: (v2 % ((1) | 1))}
                                    var v120 uint = 65535
                                    v113 += v22
                                    if true {
                                        if (float32(-2.25) == 1e10) {
                                            v121 := GenS1{F0: (v2 * int32(v5))}
                                            _ = v121
                                        }
                                        v113 = (v113 >> ((0) & 63))
                                        switch "\n\t" {
                                        case "foo bar":
                                            _ = len(v4)
                                        case "hello":
                                            var v122 uint8 = 128
                                            _ = v122
                                        }
                                    } else {
                                        switch "a" {
                                        case "a":
                                            if true {
                                                v1 -= v1
                                            } else {
                                                var v123 float64 = float64(((0.1 + (-2.25 - -2.25)) + 0.1))
                                                switch fi24 {
                                                case 32767:
                                                    var v124 string = ((("hello" + ((("\n\t" + ("hello" + "hello")) + (("" + "hello") + ("\n\t" + "hello"))) + ("foo bar" + "hello"))) + ("a" + "foo bar")) + "")
                                                    var v125 bool = ((v123 - v123) == (((v123 * v123) * v123) * v123))
                                                    if (-717 < -32768) {
                                                        v3 = (v3 | 429)
                                                        if (937 <= v18) {
                                                            fmt.Println(v120)
                                                            v113 = v113
                                                            var v126 uint8 = 255
                                                            _ = v126
                                                        }
                                                        _ = len(v4)
                                                    }
                                                    var v127 uint = 212
                                                    _ = v124
                                                    _ = v125
                                                    _ = v127
                                                case -217:
                                                    v128 := GenS9{F0: (v5 & v5), F1: ""}
                                                    _ = v128
                                                case -9223372036854775808:
                                                    switch fi24 {
                                                    case 9223372036854775807:
                                                        switch -2147483648 {
                                                        case 630:
                                                            if (int32(v5) > (v2 >> ((1) & 63))) {
                                                                if false {
                                                                    switch -9223372036854775808 {
                                                                    case 128:
                                                                        fmt.Println(v113)
                                                                    case 2147483647:
                                                                        fmt.Println(v22, fi24, v5)
                                                                    default:
                                                                        v113 -= v113
                                                                    }
                                                                    if (((v5 & v5) >= v5) || (!(true || (false && false)))) {
                                                                        v2 += v2
                                                                        fmt.Println(v22)
                                                                        if ((v3 % ((v3) | 1)) != (^(v3))) {
                                                                            _ = len(v4)
                                                                            v129 := GenS17{F0: -294}
                                                                            var v130 uint16 = 0
                                                                            _ = v129
                                                                            _ = v130
                                                                        }
                                                                    } else {
                                                                        var v131 float64 = float64(3.14159)
                                                                        fmt.Println(fi24, v131)
                                                                        _ = v131
                                                                    }
                                                                    v132 := GenS15{F0: 2147483648, F1: "foo bar"}
                                                                    _ = v132
                                                                }
                                                            }
                                                        case 4294967295:
                                                            var v133 uint64 = 4294967295
                                                            _ = v133
                                                        case 128:
                                                            v4 = append(v4, "a")
                                                            if len(v4) > 0 {
                                                                v4[0] = "a"
                                                            }
                                                        }
                                                    default:
                                                        var v134 uint32 = (v5 & 42)
                                                        _ = v134
                                                    }
                                                }
                                                v5 += v5
                                                _ = v123
                                            }
                                        case "hello":
                                            var v135 float32 = 3.14159
                                            _ = v135
                                        case "foo bar":
                                            var v136 string = "hello"
                                            _ = v136
                                        }
                                        var v137 int16 = -32768
                                        _ = v137
                                    }
                                    if false {
                                        var v138 int8 = 42
                                        if (!true) {
                                            var v139 bool = ("" < ("a" + ""))
                                            var v140 uint64 = 873
                                            if true {
                                                var v141 int = (-(v18))
                                                v142 := GenS10{F0: 0, F1: 128, F2: int16(v22)}
                                                fmt.Println(v113, v141)
                                                _ = v141
                                                _ = v142
                                            } else {
                                                v140 += uint64(v5)
                                                switch (((("a" + ("foo bar" + "hello")) + "") + "\n\t") + ("foo bar" + "foo bar")) {
                                                case "foo bar":
                                                    v143 := GenS18{F0: (v113 & v22), F1: v139}
                                                    _ = v143
                                                default:
                                                    if ((v139 || (!false)) || (!true)) {
                                                        var v144 int64 = (v113 & 0)
                                                        _ = v144
                                                    } else {
                                                        if len(v4) > 0 {
                                                            v4[0] = "foo bar"
                                                        }
                                                        switch ("\n\t" + "hello") {
                                                        case "\n\t":
                                                            var v145 uint8 = 127
                                                            _ = v145
                                                        default:
                                                            switch ("foo bar" + ((("hello" + "\n\t") + ("" + ("hello" + "a"))) + ("hello" + "\n\t"))) {
                                                            case "":
                                                                if ("foo bar" >= (("hello" + "hello") + "foo bar")) {
                                                                    var v146 uint = v120
                                                                    if len(v4) > 0 {
                                                                        v4[0] = "a"
                                                                    }
                                                                    v147 := GenS9{F0: v5, F1: "\n\t"}
                                                                    _ = v146
                                                                    _ = v147
                                                                }
                                                            }
                                                        }
                                                    }
                                                }
                                            }
                                            if (!false) {
                                                var v148 int8 = -128
                                                var v149 uint32 = v5
                                                _ = v148
                                                _ = v149
                                            } else {
                                                switch fi21 {
                                                case 1:
                                                    fmt.Println(v113, v140, v22)
                                                case 0:
                                                    var v150 uint16 = 353
                                                    _ = v150
                                                case 4294967295:
                                                    if (6 < 1000) {
                                                        if v139 {
                                                            fmt.Println(v22)
                                                        } else {
                                                            v151 := GenS11{F0: 9, F1: v120}
                                                            var v152 int16 = -318
                                                            _ = v151
                                                            _ = v152
                                                        }
                                                        var v153 float32 = (1.5 + (((float32(float64(3.14159)) / (0.1 / 3.14159)) / (float32(1e10) - 0.1)) - (3.14159 * (1e10 * 1e10))))
                                                        _ = v153
                                                    } else {
                                                        if (v139 || v139) {
                                                            var v154 uint64 = (v140 + 516)
                                                            _ = v154
                                                        } else {
                                                            if v139 {
                                                                v155 := GenS16{F0: (v2 * v2)}
                                                                _ = v155
                                                            } else {
                                                                v3 = 147
                                                            }
                                                            fmt.Println(fi24)
                                                            if len(v4) > 0 {
                                                                v4[0] = "foo bar"
                                                            }
                                                            switch int(v140) {
                                                            case 0:
                                                                v4 = append(v4, (("" + ("a" + "hello")) + ("a" + "hello")))
                                                            case 42:
                                                                fmt.Println(v138, fi21, v22)
                                                            case 266:
                                                                _ = len(v4)
                                                            default:
                                                                v5 = uint32(v120)
                                                            }
                                                        }
                                                    }
                                                }
                                                if len(v4) > 0 {
                                                    v4[0] = ("\n\t" + "")
                                                }
                                                v156 := GenS18{F0: v22, F1: v139}
                                                _ = v156
                                            }
                                            _ = v139
                                            _ = v140
                                        }
                                        _ = v138
                                    } else {
                                        v113 -= (^(v22))
                                        if ((v1 > 0) && false) {
                                            var v157 bool = true
                                            var v158 uint64 = 0
                                            _ = v157
                                            _ = v158
                                        }
                                    }
                                    v3 += (v3 * v3)
                                    _ = v119
                                    _ = v120
                                }
                            } else {
                                var v159 string = "a"
                                v113 = (v22 - (v22 % ((-1) | 1)))
                                _ = v159
                            }
                            var v160 int16 = 32767
                            v3 += v3
                            _ = v113
                            _ = v160
                        } else {
                            fmt.Println(v1)
                        }
                    case 2147483647:
                        fmt.Println(v2)
                    case -48:
                        v4 = append(v4, "")
                    default:
                        fmt.Println(fi24)
                    }
                    v5 = v5
                    _ = fi24
                }
                switch v18 {
                case 127:
                    v161 := GenS1{F0: (v2 << ((354) & 63))}
                    _ = v161
                case -2147483648:
                    var v162 int64 = v22
                    _ = v162
                case 32768:
                    v3 = 149
                }
            }
            _ = fi21
            _ = v22
        }
    } else {
        if true {
            var v163 uint16 = v3
            _ = v163
        }
        if len(v4) > 0 {
            v4[0] = "foo bar"
        }
    }
    var v164 uint64 = 18446744073709551615
    _ = len(v4)
    var v165 float64 = 0.1
    if (true && true) {
        if (v5 == v5) {
            var v166 uint16 = (v3 * 181)
            if true {
                var v167 int = v18
                if (v2 > v2) {
                    if (true || false) {
                        v168 := GenS19{F0: -32768, F1: 1}
                        v169 := GenS20{F0: (^(v5)), F1: (v164 ^ 2)}
                        for fi170 := 0; fi170 < 4; fi170++ {
                            var v171 uint32 = v5
                            _ = fi170
                            _ = v171
                        }
                        _ = v168
                        _ = v169
                    } else {
                        if false {
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        } else {
                            if true {
                                v172 := GenS6{F0: 4294967295, F1: v167}
                                v165 += (v165 / 1e10)
                                _ = v172
                            } else {
                                v173 := GenS7{F0: 0, F1: ((true || (42 != (v164 & 1))) && (float32(1.5) >= 3.14159))}
                                v174 := GenS13{F0: 480, F1: int64(v164), F2: "\n\t"}
                                switch ("\n\t" + "foo bar") {
                                case "":
                                    if true {
                                        if false {
                                            var v175 uint8 = 139
                                            var v176 uint32 = 622
                                            _ = v175
                                            _ = v176
                                        }
                                        if (42 < 0) {
                                            v3 = (v166 / ((v3) | 1))
                                            var v177 int = v18
                                            v5 -= 0
                                            _ = v177
                                        }
                                    }
                                }
                                _ = v173
                                _ = v174
                            }
                            if (!true) {
                                v165 = v165
                                if (((0 < (v2 / (((v2 - int32(v5))) | 1))) || false) && false) {
                                    var v178 uint = 2
                                    var v179 uint64 = 0
                                    var v180 int32 = (v2 + 819)
                                    _ = v178
                                    _ = v179
                                    _ = v180
                                } else {
                                    for fi181 := 0; fi181 < 11; fi181++ {
                                        if (v5 > (v5 - (v5 % ((v5) | 1)))) {
                                            fmt.Println(v167, v1, v1)
                                            var v182 bool = ((fi181 / (((v167 / ((fi181) | 1))) | 1)) != fi181)
                                            v5 = (v5 - 2)
                                            _ = v182
                                        }
                                        _ = fi181
                                    }
                                }
                                if ((v164 | (v164 - (v164 + (v164 * 0)))) == (v164 << ((0) & 63))) {
                                    var v183 int16 = 254
                                    _ = v183
                                } else {
                                    v3 = (v3 % (((v166 & 65535)) | 1))
                                    if len(v4) > 0 {
                                        v4[0] = "\n\t"
                                    }
                                }
                            } else {
                                var v184 uint8 = 238
                                if false {
                                    v185 := GenS8{F0: 196}
                                    v186 := GenS21{F0: 1, F1: 302}
                                    _ = v185
                                    _ = v186
                                }
                                v187 := GenS8{F0: (v166 ^ (v166 << ((355) & 63)))}
                                _ = v184
                                _ = v187
                            }
                        }
                    }
                } else {
                    v188 := GenS2{F0: 127, F1: 0, F2: (false && (!false))}
                    var v189 float32 = -2.25
                    var v190 int32 = 42
                    _ = v188
                    _ = v189
                    _ = v190
                }
                switch -333 {
                case 237:
                    switch v18 {
                    case 357:
                        v191 := GenS4{F0: 1}
                        _ = v191
                    default:
                        var v192 uint = 4294967295
                        _ = v192
                    }
                case 42:
                    var v193 float64 = v165
                    _ = v193
                }
                _ = v167
            }
            if ("foo bar" == "") {
                v1 = v1
                if ((v1 * -128) != int8(v5)) {
                    v18 = 655
                    var v194 int16 = -1
                    _ = v194
                }
                var v195 bool = false
                _ = v195
            } else {
                for fi196 := 0; fi196 < 3; fi196++ {
                    if (!(3.14159 != float32(v165))) {
                        v197 := GenS22{F0: 122, F1: -9223372036854775808, F2: (1 > 1)}
                        var v198 int8 = (v1 + v1)
                        _ = v197
                        _ = v198
                    } else {
                        for fi199 := 0; fi199 < 9; fi199++ {
                            switch "" {
                            case "a":
                                v5 += (v5 * v5)
                            }
                            v164 = (^(v164))
                            switch (fi196 ^ (fi196 >> ((42) & 63))) {
                            case 2147483647:
                                v200 := GenS23{F0: 0, F1: true}
                                _ = v200
                            case -49:
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            case 128:
                                _ = len(v4)
                            }
                            _ = fi199
                        }
                        fmt.Println(v1, v166)
                    }
                    _ = fi196
                }
            }
            _ = v166
        } else {
            for fi201 := 0; fi201 < 11; fi201++ {
                v202 := GenS13{F0: v3, F1: -9223372036854775808, F2: "a"}
                _ = fi201
                _ = v202
            }
        }
        v165 = 1e10
    } else {
        v4 = append(v4, ((("a" + "foo bar") + "") + ("\n\t" + ((("a" + ("foo bar" + "foo bar")) + "a") + ""))))
    }
    var v203 uint32 = (v5 / ((v5) | 1))
    if (((^(v2)) < v2) && (false && true)) {
        switch (v18 / ((-556) | 1)) {
        case 4294967295:
            var v204 uint = 1
            _ = v204
        case 42:
            switch 225 {
            case 32767:
                var v205 int64 = 448
                _ = v205
            case -9223372036854775808:
                for fi206 := 0; fi206 < 5; fi206++ {
                    v203 = v5
                    if len(v4) > 0 {
                        v4[0] = "hello"
                    }
                    if false {
                        var v207 uint16 = (v3 | (v3 - v3))
                        var v208 uint64 = (v164 << ((846) & 63))
                        v164 = v208
                        _ = v207
                        _ = v208
                    } else {
                        if true {
                            if true {
                                switch "a" {
                                case "\n\t":
                                    v209 := GenS24{F0: 1}
                                    _ = v209
                                case "foo bar":
                                    v165 += 3.14159
                                case "":
                                    v18 += fi206
                                default:
                                    var v210 int8 = (^(v1))
                                    _ = v210
                                }
                            } else {
                                switch "a" {
                                case "foo bar":
                                    switch ("a" + ("" + ("\n\t" + "\n\t"))) {
                                    case "foo bar":
                                        switch 0 {
                                        case 0:
                                            if ("foo bar" > "a") {
                                                v203 = v5
                                                if len(v4) > 0 {
                                                    v4[0] = "a"
                                                }
                                                v5 += (v203 - 2)
                                            } else {
                                                var v211 int16 = 0
                                                fmt.Println(v1)
                                                _ = len(v4)
                                                _ = v211
                                            }
                                        case -327:
                                            if (v2 == v2) {
                                                var v212 int16 = -32768
                                                _ = v212
                                            } else {
                                                v18 = 387
                                                if false {
                                                    for fi213 := 0; fi213 < 15; fi213++ {
                                                        switch "\n\t" {
                                                        case "":
                                                            var v214 float64 = (v165 * 0.1)
                                                            _ = v214
                                                        case "hello":
                                                            if true {
                                                                v215 := GenS25{F0: 42}
                                                                v18 = 0
                                                                _ = v215
                                                            } else {
                                                                _ = len(v4)
                                                                fmt.Println(v164, v1)
                                                            }
                                                        }
                                                        v216 := GenS21{F0: 0, F1: 1}
                                                        if true {
                                                            v1 = 127
                                                            var v217 bool = false
                                                            switch (fi206 / ((fi206) | 1)) {
                                                            case -409:
                                                                var v218 int = fi206
                                                                _ = v218
                                                            case 1:
                                                                v4 = append(v4, "hello")
                                                            }
                                                            _ = v217
                                                        }
                                                        _ = fi213
                                                        _ = v216
                                                    }
                                                    var v219 uint32 = 42
                                                    _ = v219
                                                } else {
                                                    for fi220 := 0; fi220 < 2; fi220++ {
                                                        fmt.Println(fi220)
                                                        _ = fi220
                                                    }
                                                    if true {
                                                        _ = len(v4)
                                                    }
                                                    if len(v4) > 0 {
                                                        v4[0] = ("hello" + "\n\t")
                                                    }
                                                }
                                            }
                                        case -90:
                                            fmt.Println(v2, v18, v18)
                                        }
                                    }
                                case "a":
                                    v5 = (v5 & v5)
                                case "":
                                    for fi221 := 0; fi221 < 9; fi221++ {
                                        switch ("a" + "foo bar") {
                                        case "foo bar":
                                            fmt.Println(v2, v165, v203)
                                        case "\n\t":
                                            fmt.Println(v3)
                                            if (1.5 >= float64((0.0 / ((1.5 + 1.5) / 3.14159)))) {
                                                v222 := GenS26{F0: v2}
                                                v223 := GenS12{F0: 42}
                                                v5 = v203
                                                _ = v222
                                                _ = v223
                                            } else {
                                                v224 := GenS8{F0: (v3 >> ((2) & 63))}
                                                switch fi206 {
                                                case -924:
                                                    if true {
                                                        v164 += 42
                                                        switch "\n\t" {
                                                        case "":
                                                            if len(v4) > 0 {
                                                                v4[0] = ""
                                                            }
                                                        }
                                                    }
                                                case 32767:
                                                    var v225 int = fi221
                                                    _ = v225
                                                case 808:
                                                    v5 = v203
                                                }
                                                _ = v224
                                            }
                                            v3 = (^(v3))
                                            var v226 float32 = (1.5 / (float32((v165 + 3.14159)) * 0.1))
                                            var v227 uint16 = (v3 % (((v3 & 0)) | 1))
                                            _ = v226
                                            _ = v227
                                        }
                                        _ = fi221
                                    }
                                default:
                                    for fi228 := 0; fi228 < 17; fi228++ {
                                        var v229 uint32 = (v203 << ((369) & 63))
                                        _ = fi228
                                        _ = v229
                                    }
                                }
                                var v230 float32 = 3.14159
                                var v231 string = ("a" + "")
                                _ = v230
                                _ = v231
                            }
                            switch ("hello" + (("" + "a") + "a")) {
                            case "foo bar":
                                var v232 float64 = v165
                                _ = v232
                            case "\n\t":
                                var v233 uint8 = 42
                                _ = v233
                            }
                            fmt.Println(v203, v5)
                        }
                    }
                    _ = fi206
                }
            case 1:
                for fi234 := 0; fi234 < 14; fi234++ {
                    for fi235 := 0; fi235 < 4; fi235++ {
                        var v236 int64 = -522
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        if len(v4) > 0 {
                            v4[0] = (("" + "\n\t") + "foo bar")
                        }
                        _ = fi235
                        _ = v236
                    }
                    _ = fi234
                }
            default:
                if ((42 == v2) || false) {
                    if true {
                        var v237 uint = 255
                        if false {
                            v238 := GenS27{F0: 255, F1: "a"}
                            fmt.Println(v18)
                            var v239 bool = (v237 != v237)
                            _ = v238
                            _ = v239
                        } else {
                            var v240 uint32 = uint32(v164)
                            var v241 uint8 = 229
                            _ = v240
                            _ = v241
                        }
                        v242 := GenS18{F0: 111, F1: true}
                        _ = v237
                        _ = v242
                    } else {
                        fmt.Println(v165, v2)
                        if true {
                            if false {
                                for fi243 := 0; fi243 < 18; fi243++ {
                                    fmt.Println(v203, v164, v3)
                                    _ = fi243
                                }
                                for fi244 := 0; fi244 < 7; fi244++ {
                                    v245 := GenS20{F0: (v203 << ((1) & 63)), F1: 2}
                                    for fi246 := 0; fi246 < 13; fi246++ {
                                        var v247 int16 = 42
                                        _ = fi246
                                        _ = v247
                                    }
                                    v164 += v164
                                    _ = fi244
                                    _ = v245
                                }
                                var v248 bool = (true || (!true))
                                _ = v248
                            } else {
                                v4 = append(v4, "foo bar")
                                fmt.Println(v203)
                                var v249 uint = 4294967295
                                _ = v249
                            }
                            v250 := GenS28{F0: ("\n\t" + ""), F1: 58}
                            v5 -= v203
                            _ = v250
                        }
                    }
                } else {
                    v165 = 0.1
                }
            }
        case 9223372036854775807:
            v251 := GenS7{F0: 437, F1: (-2.25 == 0.1)}
            _ = v251
        default:
            var v252 uint = 189
            _ = v252
        }
        v2 = v2
        switch ("" + "a") {
        case "\n\t":
            var v253 uint8 = 127
            _ = v253
        }
    }
    switch v18 {
    case -9223372036854775808:
        v254 := GenS21{F0: 0, F1: -1}
        _ = v254
    case 652:
        for fi255 := 0; fi255 < 3; fi255++ {
            for fi256 := 0; fi256 < 4; fi256++ {
                v257 := GenS7{F0: 0, F1: false}
                v258 := GenS21{F0: 781, F1: int16(v2)}
                var v259 float32 = 0.1
                _ = fi256
                _ = v257
                _ = v258
                _ = v259
            }
            if ((v5 << ((1) & 63)) != v203) {
                switch 817 {
                case 1:
                    if true {
                        var v260 int32 = 987
                        _ = v260
                    } else {
                        if (3.14159 > (0.0 / ((float32(v165) - (((1e10 + (3.14159 / (0.0 + (((-2.25 / 0.1) - float32((v165 * 1e10))) - -2.25)))) / 1.5) / 1e10)) * 0.0))) {
                            switch "\n\t" {
                            case "":
                                var v261 uint64 = (v164 / (((v164 * v164)) | 1))
                                _ = v261
                            }
                            var v262 int16 = -371
                            v262 = v262
                            _ = v262
                        } else {
                            v4 = append(v4, ("" + "foo bar"))
                            v2 -= v2
                            v203 = v203
                        }
                    }
                case 9223372036854775807:
                    fmt.Println(fi255)
                default:
                    fmt.Println(v2)
                }
                for fi263 := 0; fi263 < 13; fi263++ {
                    v4 = append(v4, "a")
                    _ = fi263
                }
            } else {
                fmt.Println(v5)
                v203 += v203
            }
            var v264 bool = false
            _ = fi255
            _ = v264
        }
    default:
        fmt.Println(v18)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v164:", v164)
    fmt.Println("v165:", v165)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v203:", v203)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v164
    _ = v165
    _ = v18
    _ = v2
    _ = v203
    _ = v3
    _ = v4
    _ = v5
}
