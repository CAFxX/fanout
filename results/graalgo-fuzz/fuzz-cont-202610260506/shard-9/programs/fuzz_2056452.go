package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 bool; F1 int16 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int32; F1 uint64; F2 string }
type GenS8 struct { F0 bool; F1 bool; F2 int }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 int; F1 uint8; F2 uint8 }
type GenS11 struct { F0 uint16; F1 bool; F2 uint64 }
type GenS12 struct { F0 uint; F1 int }
type GenS13 struct { F0 uint32; F1 string }
type GenS14 struct { F0 int8; F1 int16 }
type GenS15 struct { F0 int8 }
type GenS16 struct { F0 uint; F1 int64; F2 uint32 }
type GenS17 struct { F0 uint32 }
type GenS18 struct { F0 int32; F1 uint8 }
type GenS19 struct { F0 int; F1 uint8 }
type GenS20 struct { F0 bool; F1 int8; F2 string }
type GenS21 struct { F0 uint64; F1 string }
type GenS22 struct { F0 int16; F1 uint16 }

func main() {
    v1 := make(map[string]string)
    var v2 uint = 65535
    v3 := make(map[string]int)
    switch 409 {
    case 128:
        for fi4 := 0; fi4 < 17; fi4++ {
            for fi5 := 0; fi5 < 11; fi5++ {
                if false {
                    if (42 <= -9223372036854775808) {
                        v6 := GenS1{F0: 0}
                        var v7 string = "a"
                        _ = v6
                        _ = v7
                    } else {
                        switch "" {
                        case "hello":
                            switch "\n\t" {
                            case "hello":
                                switch "foo bar" {
                                case "a":
                                    var v8 int16 = -1
                                    _ = v8
                                case "foo bar":
                                    if false {
                                        switch "a" {
                                        case "a":
                                            v2 += v2
                                        case "":
                                            if true {
                                                var v9 uint16 = 2
                                                if (true && true) {
                                                    if false {
                                                        v9 -= (v9 & (v9 << ((255) & 63)))
                                                        var v10 uint16 = v9
                                                        if ((v10 % ((42) | 1)) != 23) {
                                                            v10 = v9
                                                        }
                                                        _ = v10
                                                    }
                                                    if false {
                                                        switch "hello" {
                                                        case "hello":
                                                            var v11 string = ("a" + ((("hello" + "foo bar") + "a") + "\n\t"))
                                                            _ = v11
                                                        default:
                                                            v9 += 0
                                                        }
                                                        var v12 int16 = 1
                                                        fmt.Println(fi5, fi4, v2)
                                                        _ = v12
                                                    } else {
                                                        switch fi5 {
                                                        case -596:
                                                            v2 = (^(v2))
                                                        default:
                                                            v13 := []int{(fi4 >> ((v2) & 63)), fi4}
                                                            _ = v13
                                                        }
                                                        var v14 int32 = 27
                                                        _ = v14
                                                    }
                                                    v15 := v3["a"]
                                                    _ = v15
                                                } else {
                                                    switch ("a" + ("" + ("a" + ("\n\t" + ("\n\t" + "\n\t"))))) {
                                                    case "\n\t":
                                                        switch "foo bar" {
                                                        case "foo bar":
                                                            if (("\n\t" == "hello") || (false || false)) {
                                                                v2 = 0
                                                                var v16 float32 = -2.25
                                                                _ = v16
                                                            } else {
                                                                switch (fi4 / ((189) | 1)) {
                                                                case 425:
                                                                    fmt.Println(fi4)
                                                                    var v17 uint16 = (v9 ^ 42)
                                                                    switch "" {
                                                                    case "hello":
                                                                        v9 -= 42
                                                                        v9 = v9
                                                                        switch (^(fi4)) {
                                                                        case 9223372036854775807:
                                                                            var v18 int8 = 127
                                                                            _ = v18
                                                                        case -2147483648:
                                                                            var v19 bool = false
                                                                            _ = v19
                                                                        default:
                                                                            if false {
                                                                                v20 := v3[""]
                                                                                v21 := GenS1{F0: int32(fi4)}
                                                                                switch "" {
                                                                                case "a":
                                                                                    if (-176 >= -1) {
                                                                                        var v22 uint8 = 48
                                                                                        _ = v22
                                                                                    } else {
                                                                                        v23 := GenS2{F0: 1}
                                                                                        _ = v23
                                                                                    }
                                                                                }
                                                                                _ = v20
                                                                                _ = v21
                                                                            }
                                                                        }
                                                                        var v24 bool = true
                                                                        if false {
                                                                            fmt.Println(v9, v2)
                                                                            v17 = uint16(v2)
                                                                        } else {
                                                                            var v25 uint = uint(v9)
                                                                            var v26 int8 = -51
                                                                            v26 += v26
                                                                            _ = v25
                                                                            _ = v26
                                                                        }
                                                                        _ = v24
                                                                    case "":
                                                                        switch "hello" {
                                                                        case "a":
                                                                            switch fi4 {
                                                                            case 215:
                                                                                if (18446744073709551615 <= 0) {
                                                                                    fmt.Println(v17)
                                                                                    v9 += v9
                                                                                    var v27 uint8 = 42
                                                                                    _ = v27
                                                                                } else {
                                                                                    v9 -= v17
                                                                                }
                                                                            case 127:
                                                                                v17 -= v17
                                                                            }
                                                                        default:
                                                                            v28 := []int{}
                                                                            _ = v28
                                                                        }
                                                                    case "a":
                                                                        var v29 float64 = 0.1
                                                                        _ = v29
                                                                    default:
                                                                        v17 += v17
                                                                    }
                                                                    if false {
                                                                        var v30 bool = true
                                                                        switch fi5 {
                                                                        case 4294967295:
                                                                            v2 = 1
                                                                        case -2147483648:
                                                                            if ((-2.25 - ((1.5 - (0.0 * (-2.25 / -2.25))) + (1.5 + -2.25))) > (1.5 + -2.25)) {
                                                                                v31 := []int{-9223372036854775808, (fi5 | (fi4 + fi5)), fi4, (fi5 % ((-998) | 1))}
                                                                                if v30 {
                                                                                    v32 := GenS2{F0: (v9 & (v17 << (((v2 >> ((v2) & 63))) & 63)))}
                                                                                    _ = len(v1)
                                                                                    fmt.Println(fi5)
                                                                                    _ = v32
                                                                                } else {
                                                                                    if v30 {
                                                                                        v17 += (v17 + 65535)
                                                                                    }
                                                                                    var v33 float64 = (0.0 + 0.1)
                                                                                    _ = v33
                                                                                }
                                                                                _ = v31
                                                                            }
                                                                            fmt.Println(v2)
                                                                            switch (fi5 - fi4) {
                                                                            case -724:
                                                                                fmt.Println(v9, v30)
                                                                            }
                                                                            v9 += 42
                                                                        case -9223372036854775808:
                                                                            var v34 string = "hello"
                                                                            _ = v34
                                                                        default:
                                                                            v9 += v9
                                                                        }
                                                                        _ = v30
                                                                    }
                                                                    switch fi5 {
                                                                    case -973:
                                                                        var v35 uint16 = (v9 | (v9 | 0))
                                                                        _ = v35
                                                                    case -1:
                                                                        v2 += 4294967295
                                                                    case 42:
                                                                        v36 := []int{-728, fi4}
                                                                        _ = v36
                                                                    default:
                                                                        switch (fi4 * fi4) {
                                                                        case 9223372036854775807:
                                                                            if ((931 == 0) || (0.0 <= (1e10 / 1e10))) {
                                                                                switch -1 {
                                                                                case -934:
                                                                                    v37 := GenS2{F0: 42}
                                                                                    _ = v37
                                                                                case 1:
                                                                                    v3["foo bar"] = (fi4 + fi4)
                                                                                    if false {
                                                                                        delete(v1, "")
                                                                                    } else {
                                                                                        switch "foo bar" {
                                                                                        case "foo bar":
                                                                                            if ((!((v9 % (((^(v9))) | 1)) != 948)) && true) {
                                                                                                if true {
                                                                                                    v2 -= 9
                                                                                                    var v38 int = fi4
                                                                                                    switch 1 {
                                                                                                    case -2147483648:
                                                                                                        v9 += v9
                                                                                                    case 1:
                                                                                                        switch "hello" {
                                                                                                        case "a":
                                                                                                            var v39 int32 = 0
                                                                                                            _ = v39
                                                                                                        case "foo bar":
                                                                                                            v40 := GenS3{F0: v38}
                                                                                                            _ = v40
                                                                                                        case "\n\t":
                                                                                                            fmt.Println(v17, fi4)
                                                                                                        }
                                                                                                    }
                                                                                                    _ = v38
                                                                                                }
                                                                                                var v41 int8 = 100
                                                                                                if true {
                                                                                                    v41 += 100
                                                                                                }
                                                                                                v9 -= (^(v17))
                                                                                                _ = v41
                                                                                            } else {
                                                                                                delete(v1, "foo bar")
                                                                                            }
                                                                                        case "hello":
                                                                                            fmt.Println(v2)
                                                                                        case "a":
                                                                                            switch fi4 {
                                                                                            case -9223372036854775808:
                                                                                                v42 := []int{}
                                                                                                _ = v42
                                                                                            case 42:
                                                                                                var v43 int32 = 0
                                                                                                _ = v43
                                                                                            default:
                                                                                                fmt.Println(fi5, v17, fi5)
                                                                                            }
                                                                                        }
                                                                                        var v44 float32 = 0.0
                                                                                        v44 += v44
                                                                                        _ = v44
                                                                                    }
                                                                                case -1:
                                                                                    var v45 uint8 = 42
                                                                                    _ = v45
                                                                                }
                                                                                var v46 float32 = (((float32((0.1 * (0.1 - (0.1 + float64((0.0 / 1.0)))))) * (-2.25 / (3.14159 - -2.25))) * 1e10) - -2.25)
                                                                                fmt.Println(fi4, v17, v17)
                                                                                _ = v46
                                                                            } else {
                                                                                v2 = 4294967295
                                                                                v47 := GenS2{F0: v17}
                                                                                _ = v47
                                                                            }
                                                                        default:
                                                                            if true {
                                                                                if false {
                                                                                    var v48 int32 = -2147483648
                                                                                    var v49 string = ((((("a" + "foo bar") + "foo bar") + ("a" + ("" + ("hello" + ("foo bar" + "hello"))))) + "a") + "foo bar")
                                                                                    v17 = 781
                                                                                    _ = v48
                                                                                    _ = v49
                                                                                } else {
                                                                                    var v50 uint32 = 1
                                                                                    delete(v3, "foo bar")
                                                                                    v51 := GenS3{F0: fi4}
                                                                                    _ = v50
                                                                                    _ = v51
                                                                                }
                                                                                switch "a" {
                                                                                case "a":
                                                                                    var v52 int = 127
                                                                                    v17 += 32767
                                                                                    _ = v52
                                                                                case "":
                                                                                    v53 := []int{4294967295, 181, fi5}
                                                                                    _ = v53
                                                                                case "\n\t":
                                                                                    var v54 uint8 = 55
                                                                                    _ = v54
                                                                                }
                                                                                v17 = v9
                                                                            } else {
                                                                                switch fi4 {
                                                                                case 32767:
                                                                                    var v55 uint32 = 277
                                                                                    _ = v55
                                                                                case 4294967295:
                                                                                    v56 := GenS4{F0: false}
                                                                                    _ = v56
                                                                                default:
                                                                                    switch "\n\t" {
                                                                                    case "":
                                                                                        v57 := []int{(fi5 & 320)}
                                                                                        _ = v57
                                                                                    case "a":
                                                                                        if ((false && (1 > (v2 + v2))) && true) {
                                                                                            switch "" {
                                                                                            case "":
                                                                                                var v58 uint8 = 128
                                                                                                v59 := []int{fi4, 2147483647}
                                                                                                _ = v58
                                                                                                _ = v59
                                                                                            case "foo bar":
                                                                                                _ = len(v1)
                                                                                            case "\n\t":
                                                                                                var v60 uint16 = (v9 << ((1) & 63))
                                                                                                _ = v60
                                                                                            }
                                                                                            v61 := GenS5{F0: false, F1: -136}
                                                                                            if (0.0 <= -2.25) {
                                                                                                switch fi4 {
                                                                                                case -2147483648:
                                                                                                    fmt.Println(fi4)
                                                                                                case -853:
                                                                                                    v62 := []int{4294967295}
                                                                                                    _ = v62
                                                                                                }
                                                                                                var v63 float64 = 1e10
                                                                                                _ = v63
                                                                                            }
                                                                                            fmt.Println(fi5, fi4, fi5)
                                                                                            if (-643 <= 68) {
                                                                                                if false {
                                                                                                    fmt.Println(fi5, v9, v17)
                                                                                                    v17 = v9
                                                                                                    v64 := GenS6{F0: 1}
                                                                                                    _ = v64
                                                                                                } else {
                                                                                                    var v65 uint8 = 193
                                                                                                    _ = v65
                                                                                                }
                                                                                            } else {
                                                                                                if (0 != 121) {
                                                                                                    if false {
                                                                                                        switch fi5 {
                                                                                                        case 45:
                                                                                                            v17 += (v9 >> (((^(v2))) & 63))
                                                                                                        }
                                                                                                        v2 = (^(v2))
                                                                                                        if true {
                                                                                                            v66 := []int{261}
                                                                                                            var v67 uint16 = (v9 - (v17 >> ((328) & 63)))
                                                                                                            switch (fi4 | (-(fi5))) {
                                                                                                            case -274:
                                                                                                                v2 -= v2
                                                                                                            case 9223372036854775807:
                                                                                                                switch "" {
                                                                                                                case "":
                                                                                                                    fmt.Println(v67, v17, v9)
                                                                                                                case "foo bar":
                                                                                                                    if true {
                                                                                                                        v68 := GenS4{F0: (!(127 < fi5))}
                                                                                                                        v67 += v67
                                                                                                                        _ = len(v3)
                                                                                                                        _ = v68
                                                                                                                    } else {
                                                                                                                        v69 := GenS2{F0: 32767}
                                                                                                                        _ = v69
                                                                                                                    }
                                                                                                                default:
                                                                                                                    fmt.Println(v9, v67, fi5)
                                                                                                                }
                                                                                                            }
                                                                                                            var v70 int32 = 759
                                                                                                            _ = v66
                                                                                                            _ = v67
                                                                                                            _ = v70
                                                                                                        }
                                                                                                    }
                                                                                                    if false {
                                                                                                        v71 := []int{fi4, fi5}
                                                                                                        v71 = append(v71, (fi4 ^ (fi4 ^ fi5)))
                                                                                                        v72 := GenS7{F0: -1, F1: 2, F2: "foo bar"}
                                                                                                        _ = v71
                                                                                                        _ = v72
                                                                                                    }
                                                                                                    var v73 float64 = 0.0
                                                                                                    _ = v73
                                                                                                }
                                                                                                if false {
                                                                                                    fmt.Println(v2, v9)
                                                                                                    v17 += v9
                                                                                                    if false {
                                                                                                        var v74 int64 = -9223372036854775808
                                                                                                        v75 := []int{2147483647, fi5}
                                                                                                        var v76 string = "a"
                                                                                                        var v77 uint8 = 128
                                                                                                        _ = v74
                                                                                                        _ = v75
                                                                                                        _ = v76
                                                                                                        _ = v77
                                                                                                    }
                                                                                                    if (1 == 4294967295) {
                                                                                                        v78 := GenS5{F0: true, F1: -302}
                                                                                                        v17 += 1
                                                                                                        _ = v78
                                                                                                    }
                                                                                                }
                                                                                                v2 += 872
                                                                                                v9 += 1
                                                                                            }
                                                                                            _ = v61
                                                                                        }
                                                                                    default:
                                                                                        switch "\n\t" {
                                                                                        case "\n\t":
                                                                                            v9 -= (v17 % ((v17) | 1))
                                                                                        case "a":
                                                                                            var v79 uint = v2
                                                                                            _ = v79
                                                                                        case "foo bar":
                                                                                            v2 += (v2 * (v2 << ((v2) & 63)))
                                                                                        default:
                                                                                            switch ((("a" + ("hello" + ("\n\t" + ""))) + "hello") + "\n\t") {
                                                                                            case "a":
                                                                                                v9 = uint16(v2)
                                                                                                v17 = (v9 & 565)
                                                                                                if (!((656 != 882) || (127 >= 100))) {
                                                                                                    v9 -= (v17 & v17)
                                                                                                    fmt.Println(fi5)
                                                                                                }
                                                                                            case "hello":
                                                                                                var v80 float32 = (3.14159 / (0.1 / 1.5))
                                                                                                _ = v80
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                }
                                                                                v81 := GenS8{F0: false, F1: (-126 > -112), F2: fi4}
                                                                                if false {
                                                                                    var v82 uint32 = 42
                                                                                    switch fi5 {
                                                                                    case 4294967295:
                                                                                        v82 -= v82
                                                                                    case 127:
                                                                                        if ((-453 < 80) || false) {
                                                                                            v9 = (v17 ^ v17)
                                                                                        } else {
                                                                                            switch "foo bar" {
                                                                                            case "\n\t":
                                                                                                switch fi5 {
                                                                                                case 4294967295:
                                                                                                    var v83 uint16 = 139
                                                                                                    _ = v83
                                                                                                case 32768:
                                                                                                    var v84 uint = 42
                                                                                                    _ = v84
                                                                                                default:
                                                                                                    v9 += v17
                                                                                                }
                                                                                            case "a":
                                                                                                v2 = (v2 ^ 404)
                                                                                            case "foo bar":
                                                                                                fmt.Println(fi4, v17, v82)
                                                                                            }
                                                                                            var v85 int8 = -48
                                                                                            var v86 int16 = -32768
                                                                                            _ = v85
                                                                                            _ = v86
                                                                                        }
                                                                                    case 255:
                                                                                        if false {
                                                                                            if (-4 >= -128) {
                                                                                                var v87 uint64 = 2
                                                                                                fmt.Println(v82)
                                                                                                switch "a" {
                                                                                                case "hello":
                                                                                                    var v88 int8 = 42
                                                                                                    _ = v88
                                                                                                case "\n\t":
                                                                                                    var v89 int64 = 9223372036854775807
                                                                                                    _ = v89
                                                                                                }
                                                                                                switch (-(fi5)) {
                                                                                                case 9223372036854775807:
                                                                                                    fmt.Println(fi4)
                                                                                                default:
                                                                                                    switch "foo bar" {
                                                                                                    case "foo bar":
                                                                                                        v90 := v1["\n\t"]
                                                                                                        var v91 int8 = 71
                                                                                                        fmt.Println(v17, fi5)
                                                                                                        _ = v90
                                                                                                        _ = v91
                                                                                                    case "\n\t":
                                                                                                        var v92 uint64 = (v87 + (v87 & v87))
                                                                                                        v93 := GenS9{F0: v2}
                                                                                                        var v94 int32 = 506
                                                                                                        _ = v92
                                                                                                        _ = v93
                                                                                                        _ = v94
                                                                                                    case "a":
                                                                                                        v95 := GenS10{F0: fi5, F1: 132, F2: 66}
                                                                                                        _ = v95
                                                                                                    }
                                                                                                }
                                                                                                v3[("a" + (("" + ((("" + ("\n\t" + "\n\t")) + "hello") + "foo bar")) + "foo bar"))] = (-(fi5))
                                                                                                _ = len(v3)
                                                                                                if true {
                                                                                                    fmt.Println(v87, fi4)
                                                                                                    fmt.Println(v17, v87, v17)
                                                                                                }
                                                                                                v82 -= 1
                                                                                                v96 := GenS8{F0: true, F1: true, F2: fi5}
                                                                                                _ = v87
                                                                                                _ = v96
                                                                                            }
                                                                                            v97 := GenS4{F0: true}
                                                                                            _ = v97
                                                                                        }
                                                                                    default:
                                                                                        fmt.Println(v9, v9, v17)
                                                                                    }
                                                                                    var v98 uint64 = 18446744073709551615
                                                                                    _ = v82
                                                                                    _ = v98
                                                                                } else {
                                                                                    switch fi5 {
                                                                                    case 389:
                                                                                        var v99 uint16 = 42
                                                                                        _ = v99
                                                                                    }
                                                                                }
                                                                                _ = v81
                                                                            }
                                                                        }
                                                                    }
                                                                    v2 -= v2
                                                                    _ = v17
                                                                }
                                                                v2 = (v2 - 1)
                                                            }
                                                        case "":
                                                            var v100 bool = true
                                                            _ = v100
                                                        }
                                                    default:
                                                        if (!(((1e10 / ((1.5 - 0.1) / 1.0)) * ((3.14159 + 1.5) * float32(0.1))) <= -2.25)) {
                                                            v2 = v2
                                                            v101 := GenS11{F0: (v9 - v9), F1: false, F2: 42}
                                                            _ = v101
                                                        }
                                                    }
                                                    v102 := v1["a"]
                                                    var v103 int = fi4
                                                    _ = v102
                                                    _ = v103
                                                }
                                                if false {
                                                    v2 = (v2 | v2)
                                                    v1["\n\t"] = "foo bar"
                                                    switch fi5 {
                                                    case -958:
                                                        v9 = 65535
                                                    default:
                                                        if false {
                                                            fmt.Println(fi5)
                                                            delete(v3, "a")
                                                            fmt.Println(fi5, fi5, fi4)
                                                            var v104 int = fi5
                                                            _ = v104
                                                        }
                                                    }
                                                }
                                                _ = v9
                                            } else {
                                                if true {
                                                    if (uint16(v2) == 22) {
                                                        v2 -= 191
                                                        v2 += (v2 * 0)
                                                    }
                                                    var v105 int8 = 42
                                                    v106 := GenS3{F0: fi5}
                                                    _ = v105
                                                    _ = v106
                                                } else {
                                                    var v107 uint32 = 2
                                                    v107 = v107
                                                    _ = v107
                                                }
                                            }
                                        case "hello":
                                            if false {
                                                v2 -= v2
                                                var v108 int = fi4
                                                _ = v108
                                            } else {
                                                v2 = v2
                                                fmt.Println(fi5, fi5, fi5)
                                                v2 -= (v2 << ((18446744073709551615) & 63))
                                            }
                                        default:
                                            v109 := v1[("a" + "\n\t")]
                                            _ = v109
                                        }
                                    } else {
                                        if false {
                                            switch (fi5 ^ fi5) {
                                            case -2147483648:
                                                var v110 float32 = ((3.14159 + 1.5) + 3.14159)
                                                _ = v110
                                            case -9223372036854775808:
                                                v111 := []int{(^(fi4)), (^(fi5)), (fi4 & 843)}
                                                _ = v111
                                            default:
                                                var v112 uint8 = 78
                                                _ = v112
                                            }
                                        } else {
                                            fmt.Println(fi4, v2)
                                            if ((((0.1 - (1e10 * 0.1)) - 3.14159) != (1.5 - ((1e10 + 0.1) - (3.14159 * 0.0)))) && (!(float64(float32(3.14159)) <= 0.0))) {
                                                var v113 int8 = 1
                                                v114 := []int{-9223372036854775808, fi5, fi4}
                                                var v115 string = (("foo bar" + ((((("a" + "a") + "hello") + "foo bar") + "a") + ("\n\t" + "\n\t"))) + "hello")
                                                _ = v113
                                                _ = v114
                                                _ = v115
                                            }
                                            v116 := []int{fi4, (fi4 + (-(fi4)))}
                                            _ = v116
                                        }
                                    }
                                case "":
                                    var v117 int16 = -1
                                    _ = v117
                                }
                                switch "" {
                                case "\n\t":
                                    v118 := []int{fi4}
                                    switch "hello" {
                                    case "a":
                                        if (65535 == 255) {
                                            v2 = (v2 % (((v2 | 18446744073709551615)) | 1))
                                            fmt.Println(fi4, fi5)
                                            v2 -= 2
                                        } else {
                                            v119 := GenS12{F0: 1, F1: -912}
                                            if true {
                                                var v120 int8 = -128
                                                _ = v120
                                            }
                                            fmt.Println(fi5, fi4, fi5)
                                            if (!(1 > 936)) {
                                                v121 := GenS5{F0: true, F1: -32768}
                                                _ = v121
                                            }
                                            _ = v119
                                        }
                                    case "\n\t":
                                        v118 = append(v118, (fi4 | fi5))
                                        switch "" {
                                        case "foo bar":
                                            switch fi4 {
                                            case 127:
                                                if false {
                                                    v2 += (v2 % ((18446744073709551615) | 1))
                                                    fmt.Println(fi4, fi5)
                                                    v2 = 42
                                                } else {
                                                    if len(v118) >= 2 {
                                                        v118 = v118[1:len(v118)]
                                                    }
                                                    v2 = (^(v2))
                                                    var v122 int = (-(fi5))
                                                    _ = v122
                                                }
                                            case 32768:
                                                v2 -= v2
                                            case 128:
                                                _ = len(v3)
                                            default:
                                                if (v2 >= v2) {
                                                    if (255 > 2) {
                                                        fmt.Println(fi4, fi5, fi5)
                                                    } else {
                                                        v2 = v2
                                                        v2 += (v2 / (((v2 * v2)) | 1))
                                                    }
                                                    if len(v118) >= 2 {
                                                        v118 = v118[1:len(v118)]
                                                    }
                                                }
                                            }
                                        case "\n\t":
                                            var v123 uint64 = 913
                                            _ = v123
                                        case "":
                                            fmt.Println(fi5)
                                        default:
                                            v124 := GenS13{F0: 361, F1: "hello"}
                                            _ = v124
                                        }
                                        var v125 bool = false
                                        _ = v125
                                    }
                                    _ = v118
                                case "hello":
                                    if false {
                                        v126 := []int{}
                                        if (true && ((!false) && true)) {
                                            v2 = (v2 & v2)
                                        }
                                        var v127 int = 32767
                                        _ = v126
                                        _ = v127
                                    }
                                }
                                if false {
                                    v128 := v1["a"]
                                    var v129 int8 = -26
                                    _ = v128
                                    _ = v129
                                } else {
                                    if (32767 <= 32767) {
                                        v130 := GenS14{F0: 100, F1: 42}
                                        v131 := GenS13{F0: 825, F1: "a"}
                                        v132 := GenS15{F0: 110}
                                        _ = v130
                                        _ = v131
                                        _ = v132
                                    }
                                    fmt.Println(fi5, fi4)
                                    var v133 int64 = 1
                                    _ = v133
                                }
                            case "\n\t":
                                v2 -= v2
                            default:
                                var v134 int = fi4
                                _ = v134
                            }
                        default:
                            fmt.Println(v2, fi4, v2)
                        }
                    }
                }
                v2 += (v2 / ((v2) | 1))
                _ = fi5
            }
            if (!(-486 > -690)) {
                v135 := []int{}
                for fi136 := 0; fi136 < 10; fi136++ {
                    fmt.Println(v2)
                    switch "foo bar" {
                    case "foo bar":
                        v135 = append(v135, (fi4 / ((fi4) | 1)))
                    case "a":
                        v2 = v2
                    case "":
                        var v137 uint8 = 1
                        _ = v137
                    }
                    var v138 string = "hello"
                    var v139 float64 = 0.1
                    _ = fi136
                    _ = v138
                    _ = v139
                }
                _ = v135
            }
            _ = fi4
        }
    default:
        var v140 float32 = ((float32(((0.0 / (3.14159 + 0.0)) - 1e10)) + 3.14159) + ((0.1 / (float32(0.1) - -2.25)) - ((float32(0.0) * ((1e10 + (0.1 - 1e10)) * (1e10 / (-2.25 * 1e10)))) / float32((-2.25 / (1e10 - -2.25))))))
        _ = v140
    }
    switch 42 {
    case 232:
        var v141 float32 = float32(1.5)
        _ = v141
    case 42:
        v142 := []int{1}
        _ = v142
    case 128:
        if (825 == 0) {
            v2 -= v2
            var v143 float64 = 1.5
            if false {
                v144 := []int{128, 4294967295, -2147483648}
                for fi145 := 0; fi145 < 19; fi145++ {
                    v143 = v143
                    if len(v144) > 0 {
                        v144[0] = (fi145 + int(v2))
                    }
                    _ = fi145
                }
                fmt.Println(v2, v2)
                _ = v144
            }
            _ = v143
        }
    default:
        v146 := GenS2{F0: 0}
        _ = v146
    }
    var v147 uint64 = 4294967295
    var v148 uint8 = 42
    var v149 uint32 = 2147483647
    fmt.Println(v147, v147, v2)
    v149 = v149
    v150 := GenS16{F0: (v2 ^ 729), F1: 42, F2: 1}
    if (!true) {
        if true {
            if (42 == 9223372036854775807) {
                v151 := v1[("" + "foo bar")]
                _ = v151
            }
            v147 = (v147 | (v147 * 18446744073709551615))
        }
        switch "a" {
        case "a":
            v152 := []int{4294967295, int(v149), -154}
            _ = v152
        case "hello":
            delete(v3, "foo bar")
        case "\n\t":
            if true {
                switch (("" + "foo bar") + "\n\t") {
                case "hello":
                    if false {
                        v2 = v2
                    }
                case "a":
                    var v153 int = -1
                    _ = v153
                case "\n\t":
                    v147 += 368
                default:
                    for fi154 := 0; fi154 < 13; fi154++ {
                        if (-2.25 > 1e10) {
                            var v155 float64 = -2.25
                            if false {
                                if (32767 != 0) {
                                    v156 := GenS11{F0: 255, F1: (!(true && true)), F2: 2}
                                    v157 := []int{fi154, fi154, -1, 32767}
                                    switch (^(fi154)) {
                                    case 128:
                                        if true {
                                            v158 := GenS17{F0: v149}
                                            _ = v158
                                        }
                                    case 376:
                                        if (1 >= v149) {
                                            switch fi154 {
                                            case 0:
                                                switch fi154 {
                                                case 128:
                                                    var v159 float32 = 0.1
                                                    _ = v159
                                                case 32768:
                                                    v148 += (v148 | 2)
                                                default:
                                                    for fi160 := 0; fi160 < 13; fi160++ {
                                                        v148 -= v148
                                                        fmt.Println(v148, v147)
                                                        v149 += (v149 % (((v149 * 65535)) | 1))
                                                        _ = fi160
                                                    }
                                                }
                                            case 128:
                                                v148 -= 2
                                            }
                                        }
                                    }
                                    _ = v156
                                    _ = v157
                                }
                            } else {
                                if true {
                                    var v161 float32 = -2.25
                                    if false {
                                        for fi162 := 0; fi162 < 20; fi162++ {
                                            v149 -= (^(v149))
                                            _ = fi162
                                        }
                                        fmt.Println(v147)
                                    } else {
                                        if false {
                                            var v163 uint16 = 133
                                            fmt.Println(v148, v148)
                                            fmt.Println(fi154, v147, v148)
                                            _ = v163
                                        }
                                        v147 = 419
                                    }
                                    _ = v161
                                }
                                fmt.Println(fi154, v149, v155)
                                var v164 bool = false
                                _ = v164
                            }
                            _ = v155
                        } else {
                            _ = len(v3)
                            v147 = v147
                        }
                        v165 := GenS18{F0: 1, F1: 2}
                        _ = fi154
                        _ = v165
                    }
                }
            } else {
                for fi166 := 0; fi166 < 4; fi166++ {
                    v167 := []int{fi166, 9223372036854775807}
                    _ = len(v167)
                    _ = fi166
                    _ = v167
                }
                v148 = 128
                v168 := []int{-9223372036854775808, 42}
                _ = v168
            }
        default:
            if false {
                if (!(false || ((0.0 / (float32((-2.25 * 0.0)) + ((-2.25 * (0.1 + ((1.5 - 3.14159) - ((0.0 + 1e10) / 0.1)))) + -2.25))) != 3.14159))) {
                    if (int32(v2) < -316) {
                        var v169 int64 = 42
                        _ = v169
                    } else {
                        switch "hello" {
                        case "":
                            v1["\n\t"] = ("foo bar" + "foo bar")
                        case "hello":
                            if (false && (v2 >= v2)) {
                                switch "a" {
                                case "hello":
                                    v170 := []int{959}
                                    _ = v170
                                default:
                                    v3["foo bar"] = 0
                                }
                            } else {
                                delete(v3, ("foo bar" + ("hello" + "\n\t")))
                                fmt.Println(v149)
                            }
                        case "foo bar":
                            for fi171 := 0; fi171 < 13; fi171++ {
                                var v172 int32 = 105
                                _ = fi171
                                _ = v172
                            }
                        }
                        switch "a" {
                        case "":
                            if false {
                                var v173 float32 = ((1.5 + (0.1 + float32(3.14159))) / 1.0)
                                v173 -= v173
                                for fi174 := 0; fi174 < 14; fi174++ {
                                    if (1 < -1) {
                                        v3["hello"] = (fi174 ^ -2147483648)
                                        v175 := GenS19{F0: fi174, F1: uint8(v149)}
                                        v176 := GenS13{F0: 2, F1: ""}
                                        _ = v175
                                        _ = v176
                                    } else {
                                        var v177 bool = true
                                        _ = v177
                                    }
                                    v178 := GenS5{F0: ((false || true) && (true || false)), F1: 42}
                                    v148 = 255
                                    _ = fi174
                                    _ = v178
                                }
                                _ = v173
                            }
                        }
                    }
                    for fi179 := 0; fi179 < 2; fi179++ {
                        if (865 < 32767) {
                            v180 := []int{fi179, 764}
                            var v181 int8 = 42
                            var v182 float64 = 0.1
                            _ = v180
                            _ = v181
                            _ = v182
                        } else {
                            v149 += (v149 | (v149 % (((v149 + v149)) | 1)))
                        }
                        for fi183 := 0; fi183 < 17; fi183++ {
                            v2 = (v2 + (v2 << ((v2) & 63)))
                            _ = fi183
                        }
                        v184 := []int{fi179, (fi179 ^ (-(fi179)))}
                        _ = fi179
                        _ = v184
                    }
                } else {
                    switch 58 {
                    case -1:
                        var v185 uint16 = uint16(v2)
                        _ = v185
                    case 9223372036854775807:
                        v2 -= 4294967295
                    default:
                        v186 := GenS4{F0: (128 < 128)}
                        _ = v186
                    }
                    v187 := GenS20{F0: (!(-521 < -281)), F1: int8(v2), F2: "foo bar"}
                    _ = v187
                }
                v188 := GenS14{F0: -128, F1: 42}
                _ = v188
            } else {
                v1[((("" + ("foo bar" + "hello")) + "\n\t") + "")] = ("" + "hello")
                var v189 int = 939
                switch v189 {
                case -1:
                    var v190 uint32 = 623
                    _ = v190
                case 125:
                    v147 = v147
                }
                _ = v189
            }
        }
        v149 = (v149 & v149)
    } else {
        delete(v3, ((("\n\t" + "hello") + ("\n\t" + "hello")) + "\n\t"))
        switch "foo bar" {
        case "hello":
            v191 := []int{-908, -51, 127}
            _ = v191
        }
    }
    fmt.Println(v149, v148)
    v147 = (v147 * 540)
    v192 := GenS21{F0: (v147 ^ v147), F1: "hello"}
    v2 = (^(v2))
    v149 = (v149 * 638)
    for fi193 := 0; fi193 < 6; fi193++ {
        var v194 float64 = 1.5
        var v195 uint8 = (v148 - (v148 + (v148 * (^(v148)))))
        if false {
            v147 += (v147 + (v147 / ((v147) | 1)))
            if false {
                v149 -= (v149 + uint32(v2))
            }
            if (0.0 != ((3.14159 * (0.0 / (-2.25 + (1e10 * float32(0.0))))) * float32((((v194 * v194) + v194) / 1e10)))) {
                fmt.Println(v195, v149, v195)
                var v196 float64 = ((v194 - float64(1e10)) + v194)
                v197 := GenS22{F0: 32767, F1: 74}
                _ = v196
                _ = v197
            } else {
                v198 := []int{fi193, 725, -545}
                for fi199 := 0; fi199 < 3; fi199++ {
                    var v200 float32 = 0.1
                    v149 = v149
                    _ = fi199
                    _ = v200
                }
                _ = v198
            }
        }
        _ = fi193
        _ = v194
        _ = v195
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v147:", v147)
    fmt.Println("v148:", v148)
    fmt.Println("v149:", v149)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v147
    _ = v148
    _ = v149
    _ = v150
    _ = v192
    _ = v2
    _ = v3
}
