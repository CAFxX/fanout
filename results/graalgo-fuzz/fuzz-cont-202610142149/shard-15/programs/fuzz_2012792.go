package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 int16; F1 uint16; F2 uint16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint; F1 int; F2 uint }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int32; F1 uint64 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 uint64 }
type GenS10 struct { F0 bool }
type GenS11 struct { F0 uint8; F1 uint32; F2 int8 }
type GenS12 struct { F0 uint; F1 int8 }
type GenS13 struct { F0 int16; F1 uint64 }
type GenS14 struct { F0 int32; F1 int32; F2 string }
type GenS15 struct { F0 uint16; F1 uint64; F2 uint16 }
type GenS16 struct { F0 uint32; F1 int32 }
type GenS17 struct { F0 uint32; F1 uint }
type GenS18 struct { F0 int16 }
type GenS19 struct { F0 int8 }
type GenS20 struct { F0 uint; F1 int16; F2 int64 }
type GenS21 struct { F0 int }
type GenS22 struct { F0 uint64; F1 bool; F2 uint64 }
type GenS23 struct { F0 uint32 }
type GenS24 struct { F0 uint32; F1 uint32; F2 int32 }
type GenS25 struct { F0 bool; F1 uint8; F2 uint }
type GenS26 struct { F0 int8 }
type GenS27 struct { F0 uint8; F1 int32; F2 int64 }
type GenS28 struct { F0 bool; F1 uint32 }
type GenS29 struct { F0 int32; F1 uint }
type GenS30 struct { F0 int32; F1 uint; F2 int32 }
type GenS31 struct { F0 uint16; F1 uint16; F2 uint64 }
type GenS32 struct { F0 int8 }
type GenS33 struct { F0 int32 }
type GenS34 struct { F0 int16; F1 uint; F2 uint }
type GenS35 struct { F0 uint16; F1 uint8; F2 bool }
type GenS36 struct { F0 int64; F1 int32 }
type GenS37 struct { F0 int32; F1 int8 }
type GenS38 struct { F0 uint8; F1 int64; F2 uint8 }
type GenS39 struct { F0 int64; F1 uint; F2 int }
type GenS40 struct { F0 uint; F1 int16; F2 uint8 }
type GenS41 struct { F0 int8; F1 int64; F2 bool }
type GenS42 struct { F0 uint8; F1 int32; F2 uint64 }
type GenS43 struct { F0 uint16 }
type GenS44 struct { F0 int32 }
type GenS45 struct { F0 int16; F1 uint64 }
type GenS46 struct { F0 uint16; F1 bool }
func genFn0() int16 {
    if ("foo bar" == ("" + ((("foo bar" + ("a" + "\n\t")) + "foo bar") + (("foo bar" + (("\n\t" + "\n\t") + ("foo bar" + "\n\t"))) + ((("" + "") + "\n\t") + "foo bar"))))) {
        if (455 < 2) {
            switch ("\n\t" + "\n\t") {
            case "foo bar":
                var v4 int16 = 32767
                _ = v4
            default:
                v5 := []int{127}
                _ = v5
            }
        } else {
            var v6 int64 = 1
            v7 := []int{2147483647, 2147483647, 2147483647}
            v6 = v6
            _ = v6
            _ = v7
        }
        v8 := []int{-1000}
        if true {
            if len(v8) >= 2 {
                v8 = v8[1:len(v8)]
            }
            var v9 uint64 = 365
            _ = v9
        }
        _ = v8
    }
    switch 127 {
    case -9223372036854775808:
        v10 := []int{}
        _ = v10
    default:
        for fi11 := 0; fi11 < 6; fi11++ {
            var v12 int64 = 0
            _ = fi11
            _ = v12
        }
    }
    return -113
}

func main() {
    v1 := []string{"foo bar", "", "a"}
    v2 := make(map[uint32]bool)
    v3 := []string{("foo bar" + ("\n\t" + "hello")), (("\n\t" + (("\n\t" + ("foo bar" + "")) + ((("a" + "foo bar") + "hello") + ("hello" + ("a" + "foo bar"))))) + "foo bar")}
    for fi13 := 0; fi13 < 10; fi13++ {
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        var v14 int64 = -1
        switch fi13 {
        case 32768:
            switch fi13 {
            case 4294967295:
                v14 += -1
            case -656:
                v14 -= (v14 & (v14 - (v14 ^ v14)))
            case 42:
                v15 := GenS1{F0: int(v14), F1: 0}
                _ = v15
            default:
                if true {
                    v1 = append(v1, ("a" + (((("" + "hello") + "a") + "\n\t") + "hello")))
                    _ = len(v1)
                    for fi16 := 0; fi16 < 11; fi16++ {
                        var v17 uint8 = 127
                        var v18 int64 = (v14 + (v14 >> ((4294967295) & 63)))
                        var v19 uint16 = 1
                        var v20 int32 = -1
                        _ = fi16
                        _ = v17
                        _ = v18
                        _ = v19
                        _ = v20
                    }
                }
            }
        case -961:
            if false {
                for fi21 := 0; fi21 < 1; fi21++ {
                    switch fi13 {
                    case 1:
                        v22 := GenS1{F0: (fi21 & (fi21 | (fi13 / ((fi13) | 1)))), F1: 0}
                        _ = v22
                    }
                    _ = fi21
                }
                v14 -= v14
            }
        }
        _ = fi13
        _ = v14
    }
    var v23 bool = true
    for fi24 := 0; fi24 < 16; fi24++ {
        v25 := GenS1{F0: -811, F1: -2147483648}
        switch "" {
        case "a":
            v23 = (!true)
        case "":
            var v26 float64 = (1.5 + -2.25)
            _ = v26
        case "\n\t":
            fmt.Println(fi24)
        }
        _ = fi24
        _ = v25
    }
    switch 4294967295 {
    case 709:
        if v23 {
            if v23 {
                if false {
                    switch "a" {
                    case "":
                        var v27 int8 = 127
                        _ = v27
                    case "hello":
                        fmt.Println(v23)
                    case "a":
                        for fi28 := 0; fi28 < 20; fi28++ {
                            var v29 string = "foo bar"
                            if v23 {
                                _ = len(v2)
                                for fi30 := 0; fi30 < 19; fi30++ {
                                    v29 = v29
                                    var v31 uint = 255
                                    _ = fi30
                                    _ = v31
                                }
                                v29 = v29
                            }
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = fi28
                            _ = v29
                        }
                    default:
                        switch 32767 {
                        case 1:
                            if v23 {
                                var v32 bool = (v23 || v23)
                                v32 = (22 >= -15)
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                _ = v32
                            } else {
                                for fi33 := 0; fi33 < 10; fi33++ {
                                    var v34 uint16 = 32767
                                    _ = len(v2)
                                    if ((!(v23 || (!(42 == 1)))) && (548 == 273)) {
                                        switch "foo bar" {
                                        case "":
                                            var v35 int32 = 1
                                            _ = v35
                                        }
                                        if v23 {
                                            v36 := GenS2{F0: 695, F1: 42, F2: v34}
                                            fmt.Println(v23, fi33, v34)
                                            _ = v36
                                        }
                                        var v37 float64 = 3.14159
                                        _ = v37
                                    } else {
                                        var v38 uint64 = 1
                                        switch "hello" {
                                        case "foo bar":
                                            v39 := GenS3{F0: -32768}
                                            _ = v39
                                        case "":
                                            if (!v23) {
                                                for fi40 := 0; fi40 < 6; fi40++ {
                                                    switch 0 {
                                                    case 32768:
                                                        v41 := GenS4{F0: 86, F1: (fi33 | (fi33 | (fi33 & -2147483648))), F2: 84}
                                                        _ = v41
                                                    case 128:
                                                        if true {
                                                            v34 = (v34 + v34)
                                                            var v42 int32 = -759
                                                            _ = v42
                                                        }
                                                    default:
                                                        switch (^(fi40)) {
                                                        case -1:
                                                            var v43 uint = uint(v38)
                                                            _ = v43
                                                        case 127:
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                        default:
                                                            var v44 int64 = 9223372036854775807
                                                            _ = v44
                                                        }
                                                    }
                                                    v23 = (607 < 255)
                                                    v45 := GenS5{F0: (^(fi40))}
                                                    _ = fi40
                                                    _ = v45
                                                }
                                            }
                                        case "\n\t":
                                            if ((0 <= 1000) || (!(-102 <= -128))) {
                                                v38 = v38
                                            }
                                        }
                                        v1 = append(v1, "a")
                                        _ = v38
                                    }
                                    _ = fi33
                                    _ = v34
                                }
                                for fi46 := 0; fi46 < 11; fi46++ {
                                    switch 9223372036854775807 {
                                    case -1:
                                        var v47 int = fi46
                                        _ = v47
                                    default:
                                        v48 := GenS2{F0: -939, F1: 458, F2: 186}
                                        _ = v48
                                    }
                                    _ = fi46
                                }
                            }
                        case -9223372036854775808:
                            if (-9223372036854775808 >= -996) {
                                switch ("a" + "a") {
                                case "foo bar":
                                    if (!v23) {
                                        var v49 int16 = 1000
                                        switch -9223372036854775808 {
                                        case 2147483647:
                                            var v50 int16 = (v49 * (v49 << ((18446744073709551615) & 63)))
                                            _ = v50
                                        default:
                                            var v51 uint64 = 225
                                            _ = v51
                                        }
                                        if ((0 <= -369) || true) {
                                            switch "foo bar" {
                                            case "\n\t":
                                                switch -9223372036854775808 {
                                                case -921:
                                                    if ((!v23) || (94 < 127)) {
                                                        if v23 {
                                                            var v52 uint32 = 2147483648
                                                            _ = v52
                                                        } else {
                                                            var v53 int32 = 42
                                                            fmt.Println(v23, v23)
                                                            switch -2147483648 {
                                                            case 127:
                                                                if v23 {
                                                                    if ((!v23) && v23) {
                                                                        v49 = v49
                                                                    }
                                                                }
                                                            case 32767:
                                                                v54 := v2[2147483647]
                                                                _ = v54
                                                            default:
                                                                v49 -= -855
                                                            }
                                                            _ = v53
                                                        }
                                                        for fi55 := 0; fi55 < 10; fi55++ {
                                                            v49 = genFn0()
                                                            _ = fi55
                                                        }
                                                    }
                                                }
                                            }
                                            v23 = true
                                            v49 = 42
                                        } else {
                                            v23 = (!v23)
                                        }
                                        _ = v49
                                    } else {
                                        v3 = append(v3, ("a" + ""))
                                        v23 = true
                                        switch -541 {
                                        case 2147483647:
                                            var v56 int16 = 463
                                            _ = v56
                                        }
                                    }
                                case "\n\t":
                                    fmt.Println(v23)
                                default:
                                    if false {
                                        v57 := GenS1{F0: -327, F1: 42}
                                        var v58 int64 = 1
                                        _ = v57
                                        _ = v58
                                    }
                                }
                                v23 = (!false)
                                for fi59 := 0; fi59 < 15; fi59++ {
                                    var v60 string = "foo bar"
                                    if v23 {
                                        v60 = "a"
                                        if len(v3) > 0 {
                                            v3[0] = (v60 + "")
                                        }
                                        var v61 float64 = 1e10
                                        _ = v61
                                    } else {
                                        if ((3.14159 / 1.0) > 3.14159) {
                                            v23 = (true || v23)
                                            if len(v1) > 0 {
                                                v1[0] = v60
                                            }
                                        }
                                    }
                                    v60 = v60
                                    _ = fi59
                                    _ = v60
                                }
                            }
                        }
                    }
                    for fi62 := 0; fi62 < 19; fi62++ {
                        if (!true) {
                            v23 = v23
                            v63 := GenS6{F0: 42, F1: 0}
                            _ = v63
                        }
                        if v23 {
                            switch (-(fi62)) {
                            case -194:
                                v23 = v23
                            case 2147483647:
                                var v64 int32 = -2147483648
                                _ = v64
                            }
                            if len(v3) > 0 {
                                v3[0] = "hello"
                            }
                        } else {
                            fmt.Println(fi62, fi62)
                            if true {
                                var v65 uint64 = 16
                                v23 = (46 != 108)
                                switch fi62 {
                                case 522:
                                    for fi66 := 0; fi66 < 5; fi66++ {
                                        switch (("foo bar" + ("\n\t" + "hello")) + ("foo bar" + "a")) {
                                        case "":
                                            v67 := GenS7{F0: 32767}
                                            _ = v67
                                        default:
                                            if (uint16(v65) <= 111) {
                                                switch (fi66 & 2147483647) {
                                                case 0:
                                                    v65 = (v65 * v65)
                                                default:
                                                    switch "a" {
                                                    case "hello":
                                                        if v23 {
                                                            if true {
                                                                v65 = 283
                                                                v23 = v23
                                                            } else {
                                                                v68 := GenS4{F0: 2, F1: fi62, F2: 42}
                                                                v69 := GenS5{F0: fi66}
                                                                delete(v2, 621)
                                                                _ = v68
                                                                _ = v69
                                                            }
                                                        } else {
                                                            v70 := GenS8{F0: 359}
                                                            if v23 {
                                                                fmt.Println(fi66)
                                                                if (!true) {
                                                                    v65 -= (v65 - v65)
                                                                    switch fi62 {
                                                                    case 2147483647:
                                                                        v65 += v65
                                                                    default:
                                                                        delete(v2, 2147483648)
                                                                    }
                                                                    v71 := GenS2{F0: 42, F1: 32767, F2: 65535}
                                                                    _ = v71
                                                                }
                                                                var v72 float64 = (0.1 * 1e10)
                                                                v73 := v2[2147483648]
                                                                _ = v72
                                                                _ = v73
                                                            }
                                                            v3 = append(v3, ("foo bar" + "a"))
                                                            _ = v70
                                                        }
                                                    case "a":
                                                        if len(v1) > 0 {
                                                            v1[0] = "\n\t"
                                                        }
                                                    }
                                                    var v74 uint64 = v65
                                                    _ = v74
                                                }
                                                v75 := GenS1{F0: (fi62 % ((128) | 1)), F1: -658}
                                                if (!v23) {
                                                    v65 = v65
                                                    fmt.Println(v65, v65)
                                                    if ((fi62 - (fi62 - (fi66 - (fi66 & 331)))) == -9223372036854775808) {
                                                        if ((!v23) && v23) {
                                                            v76 := GenS9{F0: (^(v65))}
                                                            _ = v76
                                                        } else {
                                                            switch fi62 {
                                                            case 4294967295:
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                            case 128:
                                                                v65 = 4294967295
                                                            case -405:
                                                                var v77 int64 = 1
                                                                _ = v77
                                                            default:
                                                                v23 = v23
                                                            }
                                                            v78 := GenS6{F0: 42, F1: v65}
                                                            _ = v78
                                                        }
                                                        v23 = (!(1 < -1))
                                                        if (59 < 0) {
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                            if ((510 >= 128) || (!v23)) {
                                                                var v79 int = fi62
                                                                _ = v79
                                                            }
                                                            v65 += (v65 + (v65 << ((4294967295) & 63)))
                                                        } else {
                                                            switch fi62 {
                                                            case -2147483648:
                                                                var v80 uint32 = 90
                                                                _ = v80
                                                            case -654:
                                                                v65 = 119
                                                            default:
                                                                var v81 uint8 = 0
                                                                _ = v81
                                                            }
                                                            var v82 uint = 0
                                                            fmt.Println(v82)
                                                            _ = v82
                                                        }
                                                    }
                                                }
                                                _ = v75
                                            }
                                        }
                                        switch "hello" {
                                        case "\n\t":
                                            if true {
                                                if true {
                                                    v23 = v23
                                                    v83 := GenS3{F0: 0}
                                                    var v84 uint32 = 503
                                                    _ = v83
                                                    _ = v84
                                                } else {
                                                    var v85 float64 = 3.14159
                                                    _ = v85
                                                }
                                                if false {
                                                    v23 = false
                                                    var v86 string = ("a" + "hello")
                                                    if ((65535 > 179) && (((255 <= 1) || (fi62 >= 32768)) && (false || v23))) {
                                                        var v87 uint32 = 0
                                                        var v88 bool = (true && false)
                                                        _ = v87
                                                        _ = v88
                                                    }
                                                    _ = v86
                                                }
                                            } else {
                                                switch "\n\t" {
                                                case "":
                                                    v23 = (1 <= 875)
                                                    delete(v2, 2)
                                                    var v89 int8 = 127
                                                    _ = v89
                                                case "hello":
                                                    v23 = ((190 >= 42) && (!false))
                                                case "\n\t":
                                                    if (false && (((v23 || (0 >= 42)) && (577 >= 350)) && v23)) {
                                                        var v90 uint8 = 2
                                                        _ = v90
                                                    } else {
                                                        var v91 int64 = 1
                                                        v91 = 288
                                                        _ = v91
                                                    }
                                                default:
                                                    v92 := GenS5{F0: (-(fi66))}
                                                    _ = v92
                                                }
                                                var v93 float64 = (float64(3.14159) / -2.25)
                                                _ = v93
                                            }
                                        case "":
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                        default:
                                            var v94 uint64 = 18446744073709551615
                                            _ = v94
                                        }
                                        if (65535 < 65535) {
                                            v95 := GenS3{F0: 0}
                                            v1 = append(v1, (((("\n\t" + ((("" + ("" + ("\n\t" + "hello"))) + (("hello" + ("" + ("a" + "foo bar"))) + ("a" + "hello"))) + (("foo bar" + "") + "foo bar"))) + ("a" + "")) + "foo bar") + ("foo bar" + (("hello" + "hello") + (((("" + "a") + "\n\t") + ("\n\t" + ("\n\t" + ("a" + ("a" + "a"))))) + "hello")))))
                                            if v23 {
                                                var v96 int32 = 42
                                                v97 := GenS10{F0: (-2.25 >= (0.1 * (0.0 - 1.5)))}
                                                v23 = v23
                                                _ = v96
                                                _ = v97
                                            }
                                            switch 588 {
                                            case 443:
                                                var v98 uint8 = 2
                                                _ = v98
                                            case -1:
                                                v65 = 18446744073709551615
                                            default:
                                                if false {
                                                    v99 := GenS11{F0: 169, F1: 2147483647, F2: 42}
                                                    v65 = v65
                                                    var v100 float64 = 0.1
                                                    v101 := GenS12{F0: 2, F1: 42}
                                                    switch "\n\t" {
                                                    case "foo bar":
                                                        v23 = (v100 <= (v100 * (1e10 * 1.5)))
                                                    case "":
                                                        v100 = 0.0
                                                    case "\n\t":
                                                        if v23 {
                                                            switch (fi62 % (((fi62 | fi66)) | 1)) {
                                                            case 42:
                                                                var v102 int32 = 42
                                                                _ = v102
                                                            case 4294967295:
                                                                if false {
                                                                    v65 = v65
                                                                    v103 := GenS13{F0: genFn0(), F1: (^(v65))}
                                                                    _ = v103
                                                                } else {
                                                                    if v23 {
                                                                        switch (^(fi66)) {
                                                                        case -876:
                                                                            v23 = true
                                                                        }
                                                                        fmt.Println(v65, v100, fi66)
                                                                        var v104 string = ""
                                                                        v105 := GenS13{F0: 1, F1: 888}
                                                                        switch v104 {
                                                                        case "foo bar":
                                                                            fmt.Println(v104, v65, fi66)
                                                                        case "a":
                                                                            if (-9223372036854775808 == -31) {
                                                                                v23 = (((128 < 1) || (!v23)) || (!(2 < 1)))
                                                                            } else {
                                                                                var v106 float64 = v100
                                                                                v23 = (657 != 1)
                                                                                switch fi66 {
                                                                                case 33:
                                                                                    switch fi62 {
                                                                                    case 32767:
                                                                                        v106 = (v100 / 0.1)
                                                                                    case 32768:
                                                                                        if ((463 <= 2) && false) {
                                                                                            v107 := GenS14{F0: 42, F1: 2147483647, F2: (v104 + ((v104 + v104) + "\n\t"))}
                                                                                            if v23 {
                                                                                                var v108 float64 = (v106 * v100)
                                                                                                v109 := GenS12{F0: 801, F1: 1}
                                                                                                v104 = ("foo bar" + v104)
                                                                                                _ = v108
                                                                                                _ = v109
                                                                                            }
                                                                                            _ = v107
                                                                                        } else {
                                                                                            var v110 uint32 = 992
                                                                                            switch v104 {
                                                                                            case "\n\t":
                                                                                                v104 = v104
                                                                                            default:
                                                                                                var v111 uint32 = v110
                                                                                                _ = v111
                                                                                            }
                                                                                            var v112 int64 = 1
                                                                                            _ = v110
                                                                                            _ = v112
                                                                                        }
                                                                                    case 2147483647:
                                                                                        var v113 bool = v23
                                                                                        _ = v113
                                                                                    default:
                                                                                        var v114 float64 = (v100 / v106)
                                                                                        _ = v114
                                                                                    }
                                                                                }
                                                                                _ = v106
                                                                            }
                                                                        default:
                                                                            fmt.Println(fi66, fi62, v65)
                                                                        }
                                                                        v115 := GenS7{F0: -290}
                                                                        v104 = "hello"
                                                                        _ = v104
                                                                        _ = v105
                                                                        _ = v115
                                                                    } else {
                                                                        switch (("a" + "hello") + ("\n\t" + ("" + "foo bar"))) {
                                                                        case "a":
                                                                            var v116 float64 = (0.0 - v100)
                                                                            _ = v116
                                                                        case "\n\t":
                                                                            var v117 int = fi62
                                                                            v117 -= (v117 - (fi62 / ((-644) | 1)))
                                                                            v118 := GenS5{F0: -2147483648}
                                                                            if v23 {
                                                                                v23 = v23
                                                                                v119 := GenS9{F0: (v65 % (((v65 + (^(v65)))) | 1))}
                                                                                _ = v119
                                                                            }
                                                                            var v120 int = (v117 >> ((659) & 63))
                                                                            switch ("" + "a") {
                                                                            case "a":
                                                                                var v121 int8 = 1
                                                                                _ = v121
                                                                            default:
                                                                                if (!(65535 != 42)) {
                                                                                    if v23 {
                                                                                        v122 := GenS7{F0: 42}
                                                                                        _ = v122
                                                                                    } else {
                                                                                        fmt.Println(v120, v120)
                                                                                        switch 301 {
                                                                                        case 32768:
                                                                                            fmt.Println(fi66, v23)
                                                                                            if len(v1) >= 2 {
                                                                                                v1 = v1[1:len(v1)]
                                                                                            }
                                                                                            v123 := GenS6{F0: 909, F1: (v65 | v65)}
                                                                                            _ = v123
                                                                                        case 1:
                                                                                            v124 := GenS13{F0: 1, F1: 18446744073709551615}
                                                                                            _ = v124
                                                                                        case -1:
                                                                                            v125 := GenS5{F0: (v120 | (-(fi66)))}
                                                                                            _ = v125
                                                                                        default:
                                                                                            var v126 uint8 = 155
                                                                                            _ = v126
                                                                                        }
                                                                                    }
                                                                                    _ = len(v2)
                                                                                } else {
                                                                                    var v127 string = "a"
                                                                                    _ = v127
                                                                                }
                                                                            }
                                                                            _ = v117
                                                                            _ = v118
                                                                            _ = v120
                                                                        case "hello":
                                                                            v2[2147483648] = (0 <= -31)
                                                                        default:
                                                                            if v23 {
                                                                                switch (^(fi66)) {
                                                                                case -731:
                                                                                    v100 -= v100
                                                                                case 42:
                                                                                    fmt.Println(v100, v65, fi62)
                                                                                case 492:
                                                                                    if (!v23) {
                                                                                        if len(v3) >= 2 {
                                                                                            v3 = v3[1:len(v3)]
                                                                                        }
                                                                                    } else {
                                                                                        switch (fi66 / ((128) | 1)) {
                                                                                        case -824:
                                                                                            var v128 uint64 = v65
                                                                                            _ = v128
                                                                                        case 9223372036854775807:
                                                                                            var v129 float32 = 1.5
                                                                                            var v130 uint64 = v65
                                                                                            _ = v129
                                                                                            _ = v130
                                                                                        case 42:
                                                                                            var v131 int64 = 1
                                                                                            _ = v131
                                                                                        }
                                                                                    }
                                                                                default:
                                                                                    var v132 uint64 = 2
                                                                                    _ = v132
                                                                                }
                                                                                if v23 {
                                                                                    var v133 int16 = 1
                                                                                    var v134 uint8 = 255
                                                                                    fmt.Println(fi62, v65, v23)
                                                                                    _ = v133
                                                                                    _ = v134
                                                                                } else {
                                                                                    var v135 int64 = 0
                                                                                    var v136 bool = (v65 <= 4294967295)
                                                                                    v137 := GenS2{F0: 1000, F1: 791, F2: 0}
                                                                                    _ = v135
                                                                                    _ = v136
                                                                                    _ = v137
                                                                                }
                                                                                v138 := GenS15{F0: 164, F1: 1, F2: 2}
                                                                                v3 = append(v3, "\n\t")
                                                                                _ = v138
                                                                            }
                                                                        }
                                                                        v139 := GenS13{F0: -32768, F1: (v65 + (^(v65)))}
                                                                        _ = v139
                                                                    }
                                                                    v140 := GenS15{F0: 0, F1: v65, F2: 65535}
                                                                    if false {
                                                                        v65 += (v65 & 4294967295)
                                                                        var v141 int = -9223372036854775808
                                                                        v142 := GenS12{F0: 2, F1: 100}
                                                                        if (0.0 >= (0.0 + 1e10)) {
                                                                            v141 += (fi66 * fi62)
                                                                            var v143 int32 = -2147483648
                                                                            _ = v143
                                                                        } else {
                                                                            v65 = (v65 & v65)
                                                                            var v144 float64 = float64(1.5)
                                                                            _ = v144
                                                                        }
                                                                        if (((2 <= 2) && true) && (255 <= 128)) {
                                                                            var v145 float32 = (float32(0.0) + float32(((((((3.14159 / 3.14159) / 1.5) * (float64(3.14159) + v100)) - ((-2.25 / v100) - (v100 + v100))) / 1.0) / v100)))
                                                                            if len(v3) >= 2 {
                                                                                v3 = v3[1:len(v3)]
                                                                            }
                                                                            _ = v145
                                                                        }
                                                                        var v146 string = ""
                                                                        fmt.Println(v65, v146, v23)
                                                                        _ = v141
                                                                        _ = v142
                                                                        _ = v146
                                                                    } else {
                                                                        var v147 int8 = 127
                                                                        if len(v3) >= 2 {
                                                                            v3 = v3[1:len(v3)]
                                                                        }
                                                                        fmt.Println(v147, v147, fi66)
                                                                        _ = v147
                                                                    }
                                                                    v148 := v2[1]
                                                                    v149 := GenS16{F0: 0, F1: 42}
                                                                    _ = v140
                                                                    _ = v148
                                                                    _ = v149
                                                                }
                                                                var v150 int64 = -1
                                                                _ = v150
                                                            case 608:
                                                                if (!(v23 || v23)) {
                                                                    var v151 uint16 = 1
                                                                    if len(v3) > 0 {
                                                                        v3[0] = (("foo bar" + "\n\t") + "foo bar")
                                                                    }
                                                                    _ = v151
                                                                }
                                                            }
                                                        }
                                                    }
                                                    if (v23 || false) {
                                                        v152 := GenS1{F0: 32768, F1: -1}
                                                        var v153 int64 = -9223372036854775808
                                                        switch 953 {
                                                        case 32767:
                                                            if v23 {
                                                                var v154 int32 = 0
                                                                if v23 {
                                                                    switch int(v154) {
                                                                    case -637:
                                                                        if len(v3) > 0 {
                                                                            v3[0] = "\n\t"
                                                                        }
                                                                    default:
                                                                        v155 := GenS10{F0: false}
                                                                        _ = v155
                                                                    }
                                                                    v65 -= 18446744073709551615
                                                                    var v156 uint8 = 222
                                                                    _ = v156
                                                                }
                                                                fmt.Println(v100, fi66, v100)
                                                                _ = v154
                                                            } else {
                                                                v100 -= 1e10
                                                                switch "\n\t" {
                                                                case "hello":
                                                                    fmt.Println(v23)
                                                                default:
                                                                    switch fi66 {
                                                                    case 42:
                                                                        switch 32767 {
                                                                        case -97:
                                                                            v23 = (712 > -329)
                                                                        }
                                                                    default:
                                                                        v157 := GenS17{F0: 1, F1: 336}
                                                                        _ = v157
                                                                    }
                                                                }
                                                            }
                                                        case -768:
                                                            switch "foo bar" {
                                                            case "\n\t":
                                                                switch 357 {
                                                                case 42:
                                                                    var v158 int32 = -214
                                                                    _ = v158
                                                                case 2147483647:
                                                                    switch "a" {
                                                                    case "":
                                                                        v159 := GenS5{F0: (fi66 | (fi66 / ((fi62) | 1)))}
                                                                        _ = v159
                                                                    default:
                                                                        delete(v2, 0)
                                                                    }
                                                                default:
                                                                    v160 := GenS10{F0: v23}
                                                                    _ = v160
                                                                }
                                                            }
                                                        case 1:
                                                            switch fi66 {
                                                            case 944:
                                                                var v161 int16 = 32767
                                                                _ = v161
                                                            default:
                                                                fmt.Println(v100, v153, fi62)
                                                            }
                                                        default:
                                                            var v162 float32 = (0.0 - 1.5)
                                                            var v163 int = fi66
                                                            if v23 {
                                                                var v164 int32 = 42
                                                                if ((1000 != 1) || (127 >= 42)) {
                                                                    switch ("" + "\n\t") {
                                                                    case "a":
                                                                        switch (fi66 & fi66) {
                                                                        case 42:
                                                                            _ = len(v3)
                                                                        case 182:
                                                                            v162 = -2.25
                                                                        }
                                                                    case "\n\t":
                                                                        var v165 int8 = 1
                                                                        _ = v165
                                                                    }
                                                                }
                                                                _ = v164
                                                            }
                                                            _ = v162
                                                            _ = v163
                                                        }
                                                        _ = v152
                                                        _ = v153
                                                    } else {
                                                        v166 := GenS1{F0: int(v65), F1: int32(v65)}
                                                        _ = v166
                                                    }
                                                    _ = v100
                                                    _ = v101
                                                    _ = v99
                                                }
                                            }
                                            if false {
                                                fmt.Println(fi62, fi62)
                                                fmt.Println(fi62, v23)
                                                v23 = (!true)
                                            }
                                            switch fi66 {
                                            case 9223372036854775807:
                                                var v167 string = ("hello" + "\n\t")
                                                _ = v167
                                            case 4294967295:
                                                var v168 float32 = 0.1
                                                _ = v168
                                            default:
                                                v23 = v23
                                                var v169 int64 = 0
                                                _ = v169
                                            }
                                            _ = v95
                                        } else {
                                            if v23 {
                                                fmt.Println(fi66, fi66, fi66)
                                                var v170 int64 = int64(fi62)
                                                _ = v170
                                            }
                                            delete(v2, 1)
                                            v171 := GenS18{F0: -606}
                                            _ = v171
                                        }
                                        _ = fi66
                                    }
                                case 4294967295:
                                    for fi172 := 0; fi172 < 12; fi172++ {
                                        delete(v2, 2147483647)
                                        fmt.Println(fi172, v65, fi62)
                                        if (-695 <= -9223372036854775808) {
                                            switch (fi172 << ((2) & 63)) {
                                            case 212:
                                                if true {
                                                    fmt.Println(v65)
                                                    var v173 uint64 = (v65 + 516)
                                                    if v23 {
                                                        var v174 int64 = 0
                                                        v173 = v65
                                                        _ = v174
                                                    } else {
                                                        if (!(670 <= -574)) {
                                                            switch (fi62 - fi172) {
                                                            case -726:
                                                                v175 := GenS3{F0: genFn0()}
                                                                _ = v175
                                                            }
                                                        }
                                                        _ = len(v2)
                                                    }
                                                    _ = v173
                                                }
                                            case 93:
                                                if v23 {
                                                    var v176 uint = 898
                                                    v2[4294967295] = (!(586 < 665))
                                                    _ = v176
                                                }
                                            case 9223372036854775807:
                                                v1 = append(v1, ("" + "foo bar"))
                                            }
                                        }
                                        var v177 uint64 = v65
                                        _ = fi172
                                        _ = v177
                                    }
                                }
                                _ = v65
                            } else {
                                switch (fi62 % (((fi62 ^ 752)) | 1)) {
                                case 127:
                                    if (v23 || true) {
                                        v178 := GenS19{F0: -128}
                                        _ = v178
                                    } else {
                                        var v179 bool = false
                                        _ = v179
                                    }
                                case -126:
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                }
                                v180 := v2[448]
                                _ = v180
                            }
                        }
                        _ = len(v2)
                        _ = fi62
                    }
                }
            }
        }
    case 32767:
        if (((2 <= 65535) && v23) || v23) {
            v23 = v23
        } else {
            var v181 int8 = 127
            _ = v181
        }
    case -967:
        for fi182 := 0; fi182 < 15; fi182++ {
            if (!v23) {
                fmt.Println(fi182)
            } else {
                var v183 float64 = (1e10 * (1e10 + -2.25))
                for fi184 := 0; fi184 < 8; fi184++ {
                    v3 = append(v3, ("" + "\n\t"))
                    var v185 float32 = 1.5
                    _ = fi184
                    _ = v185
                }
                _ = v183
            }
            _ = fi182
        }
    default:
        fmt.Println(v23)
    }
    switch 42 {
    case 42:
        for fi186 := 0; fi186 < 7; fi186++ {
            var v187 uint16 = 42
            v187 += 320
            if v23 {
                _ = len(v3)
                var v188 uint32 = 2147483648
                if (128 > 127) {
                    switch "hello" {
                    case "hello":
                        var v189 uint = 1
                        _ = v189
                    case "\n\t":
                        v190 := GenS2{F0: -1, F1: (v187 * (^(v187))), F2: (v187 / ((v187) | 1))}
                        _ = v190
                    case "foo bar":
                        if v23 {
                            switch (fi186 / (((-(fi186))) | 1)) {
                            case -556:
                                var v191 uint8 = 255
                                _ = v191
                            }
                        } else {
                            if false {
                                var v192 int16 = 1000
                                v193 := GenS11{F0: 128, F1: (v188 ^ (^(v188))), F2: int8(v187)}
                                _ = v192
                                _ = v193
                            } else {
                                v188 -= (v188 & v188)
                            }
                            var v194 bool = (v23 && (!true))
                            _ = v194
                        }
                    default:
                        var v195 int64 = 113
                        _ = v195
                    }
                } else {
                    var v196 uint32 = (v188 * (^(v188)))
                    if (903 > 99) {
                        if v23 {
                            v188 = 4294967295
                        } else {
                            for fi197 := 0; fi197 < 19; fi197++ {
                                var v198 uint32 = 256
                                _ = fi197
                                _ = v198
                            }
                            for fi199 := 0; fi199 < 6; fi199++ {
                                v200 := GenS20{F0: 1, F1: 1000, F2: 2}
                                if (2 <= 65535) {
                                    v196 -= v196
                                    fmt.Println(fi199, fi186)
                                    v201 := GenS21{F0: (fi186 * fi199)}
                                    v202 := v2[uint32(v187)]
                                    v203 := GenS22{F0: 1, F1: false, F2: 18446744073709551615}
                                    _ = v201
                                    _ = v202
                                    _ = v203
                                } else {
                                    v2[714] = (100 == -3)
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                }
                                _ = fi199
                                _ = v200
                            }
                        }
                        for fi204 := 0; fi204 < 8; fi204++ {
                            fmt.Println(fi204)
                            _ = fi204
                        }
                    }
                    _ = v196
                }
                _ = v188
            }
            _ = fi186
            _ = v187
        }
    case 2147483647:
        v205 := v2[42]
        _ = v205
    }
    v23 = (!false)
    var v206 float64 = 1.5
    var v207 uint32 = 65535
    for fi208 := 0; fi208 < 18; fi208++ {
        var v209 uint16 = 42
        var v210 uint8 = 100
        _ = fi208
        _ = v209
        _ = v210
    }
    v211 := GenS6{F0: 2147483647, F1: 42}
    var v212 uint32 = (v207 % (((^(v207))) | 1))
    if v23 {
        if v23 {
            v207 += v207
            v206 -= v206
        } else {
            switch ("" + "") {
            case "hello":
                fmt.Println(v206, v212, v23)
            case "":
                if v23 {
                    _ = len(v1)
                } else {
                    v213 := GenS23{F0: (v207 * (^(v212)))}
                    v214 := v2[(v207 * v212)]
                    _ = v213
                    _ = v214
                }
            case "a":
                fmt.Println(v207, v206, v23)
            default:
                if len(v3) > 0 {
                    v3[0] = ""
                }
            }
        }
        for fi215 := 0; fi215 < 13; fi215++ {
            switch "" {
            case "\n\t":
                fmt.Println(v212)
            case "hello":
                if v23 {
                    if (v23 || ((!((v207 >> ((397) & 63)) >= (^(v212)))) || (v207 <= (v212 ^ v212)))) {
                        v23 = false
                        switch (^(fi215)) {
                        case 32768:
                            if (0 > 2) {
                                fmt.Println(v206, v206)
                                var v216 uint16 = 2
                                v2[v207] = v23
                                _ = v216
                            }
                        case -705:
                            v217 := GenS24{F0: v212, F1: (v207 % (((v212 * (^(v207)))) | 1)), F2: 2147483647}
                            _ = v217
                        }
                    } else {
                        if v23 {
                            v23 = (667 != 700)
                            fmt.Println(v207, v207)
                            var v218 string = "hello"
                            _ = v218
                        } else {
                            var v219 uint16 = 0
                            for fi220 := 0; fi220 < 20; fi220++ {
                                v221 := GenS17{F0: (v207 + v212), F1: 307}
                                _ = fi220
                                _ = v221
                            }
                            _ = v219
                        }
                        fmt.Println(fi215)
                        delete(v2, 42)
                    }
                    v207 = (^(v207))
                    if v23 {
                        var v222 string = ("a" + "a")
                        _ = v222
                    } else {
                        if v23 {
                            if v23 {
                                if false {
                                    v223 := GenS7{F0: 32767}
                                    var v224 string = "hello"
                                    var v225 uint16 = 890
                                    _ = v223
                                    _ = v224
                                    _ = v225
                                }
                            }
                            if false {
                                var v226 float32 = 1.5
                                switch fi215 {
                                case 9223372036854775807:
                                    fmt.Println(fi215)
                                case 32768:
                                    v226 = v226
                                case -972:
                                    v227 := GenS8{F0: 1}
                                    _ = v227
                                default:
                                    v206 = (v206 + v206)
                                }
                                _ = v226
                            } else {
                                v207 = (v212 | 2147483648)
                                var v228 uint32 = 4294967295
                                v229 := GenS3{F0: 1000}
                                _ = v228
                                _ = v229
                            }
                        } else {
                            var v230 int16 = 1
                            fmt.Println(v207, v23)
                            for fi231 := 0; fi231 < 14; fi231++ {
                                var v232 int = (fi231 + fi231)
                                v1 = append(v1, "hello")
                                _ = fi231
                                _ = v232
                            }
                            _ = v230
                        }
                        if v23 {
                            switch (fi215 * (fi215 - fi215)) {
                            case 127:
                                delete(v2, v207)
                            case -1:
                                delete(v2, v207)
                            }
                            for fi233 := 0; fi233 < 20; fi233++ {
                                var v234 int = (fi215 << ((354) & 63))
                                if (18446744073709551615 != 214) {
                                    v206 = (3.14159 * v206)
                                    var v235 int8 = 11
                                    _ = v235
                                }
                                _ = fi233
                                _ = v234
                            }
                        } else {
                            v236 := GenS2{F0: genFn0(), F1: 906, F2: 255}
                            var v237 int = fi215
                            v238 := GenS25{F0: ("hello" >= (("" + "hello") + (("a" + "") + "a"))), F1: 0, F2: 18446744073709551615}
                            _ = v236
                            _ = v237
                            _ = v238
                        }
                        v239 := GenS19{F0: 100}
                        _ = v239
                    }
                }
            case "":
                if v23 {
                    v212 -= (v207 >> ((18446744073709551615) & 63))
                    fmt.Println(v23, v206, fi215)
                    v206 = 0.1
                } else {
                    v240 := GenS26{F0: int8(v212)}
                    switch (fi215 & (fi215 >> ((124) & 63))) {
                    case 128:
                        if v23 {
                            if v23 {
                                var v241 uint32 = v212
                                var v242 int32 = -2147483648
                                _ = v241
                                _ = v242
                            } else {
                                if (!v23) {
                                    if true {
                                        if v23 {
                                            v212 -= 2147483648
                                            v207 = v207
                                        }
                                        v2[v207] = v23
                                    }
                                    for fi243 := 0; fi243 < 15; fi243++ {
                                        var v244 int16 = 836
                                        _ = fi243
                                        _ = v244
                                    }
                                } else {
                                    for fi245 := 0; fi245 < 6; fi245++ {
                                        _ = len(v2)
                                        _ = fi245
                                    }
                                    if v23 {
                                        fmt.Println(v207, v207)
                                        v246 := GenS18{F0: -239}
                                        v23 = v23
                                        _ = v246
                                    } else {
                                        switch 0 {
                                        case 4294967295:
                                            v247 := GenS22{F0: 478, F1: (-961 >= 1), F2: 33}
                                            _ = v247
                                        case -1:
                                            v248 := GenS23{F0: (^(v212))}
                                            _ = v248
                                        default:
                                            if len(v1) > 0 {
                                                v1[0] = "a"
                                            }
                                        }
                                        for fi249 := 0; fi249 < 12; fi249++ {
                                            v250 := GenS24{F0: v212, F1: 42, F2: 0}
                                            var v251 int32 = int32(fi249)
                                            _ = fi249
                                            _ = v250
                                            _ = v251
                                        }
                                    }
                                }
                                delete(v2, 1)
                                if v23 {
                                    var v252 string = "hello"
                                    _ = v252
                                }
                            }
                            for fi253 := 0; fi253 < 16; fi253++ {
                                v206 -= 1.5
                                v212 = (v207 & v207)
                                _ = fi253
                            }
                        } else {
                            v254 := GenS14{F0: 37, F1: 42, F2: "a"}
                            v255 := GenS23{F0: (v207 >> ((12) & 63))}
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = v254
                            _ = v255
                        }
                    case -2147483648:
                        v256 := GenS27{F0: 2, F1: 42, F2: -80}
                        _ = v256
                    case -236:
                        v257 := GenS26{F0: 100}
                        _ = v257
                    default:
                        if (!true) {
                            fmt.Println(v206, v207)
                        } else {
                            _ = len(v2)
                            for fi258 := 0; fi258 < 17; fi258++ {
                                if (10 < 4294967295) {
                                    v212 = 46
                                    var v259 int = (fi215 & (^(fi258)))
                                    if (((v259 & 127) == -1) || (42 != 42)) {
                                        switch (("foo bar" + (("" + ("hello" + "\n\t")) + ("\n\t" + "hello"))) + ("" + "foo bar")) {
                                        case "a":
                                            v260 := GenS8{F0: -993}
                                            _ = v260
                                        case "":
                                            if false {
                                                var v261 int8 = 42
                                                var v262 int16 = 32767
                                                _ = len(v2)
                                                fmt.Println(v212, v262)
                                                v1 = append(v1, "foo bar")
                                                _ = v261
                                                _ = v262
                                            } else {
                                                fmt.Println(v23, v23)
                                                v212 = 2147483647
                                            }
                                            if (9223372036854775807 <= -1) {
                                                fmt.Println(v206, v212, v206)
                                                if true {
                                                    switch "a" {
                                                    case "a":
                                                        if (v23 || ((v212 / ((0) | 1)) < v207)) {
                                                            switch "\n\t" {
                                                            case "hello":
                                                                var v263 bool = v23
                                                                _ = v263
                                                            }
                                                            var v264 bool = (float32(v206) != 0.1)
                                                            var v265 string = "a"
                                                            _ = v264
                                                            _ = v265
                                                        } else {
                                                            v266 := GenS28{F0: (v23 && v23), F1: (v207 & v212)}
                                                            v259 = (-(fi215))
                                                            if v23 {
                                                                if v23 {
                                                                    switch (fi215 % (((fi258 >> ((435) & 63))) | 1)) {
                                                                    case 9223372036854775807:
                                                                        fmt.Println(v259, v259)
                                                                    case 128:
                                                                        v267 := GenS29{F0: 924, F1: 19}
                                                                        _ = v267
                                                                    case -763:
                                                                        var v268 int64 = -9223372036854775808
                                                                        _ = v268
                                                                    default:
                                                                        v206 += v206
                                                                    }
                                                                } else {
                                                                    fmt.Println(v212, v23, fi258)
                                                                    if false {
                                                                        var v269 uint32 = v212
                                                                        v23 = true
                                                                        var v270 int16 = 32767
                                                                        v269 = (v207 << ((65535) & 63))
                                                                        _ = v269
                                                                        _ = v270
                                                                    } else {
                                                                        v259 = v259
                                                                        switch "" {
                                                                        case "":
                                                                            if v23 {
                                                                                var v271 uint64 = 130
                                                                                _ = v271
                                                                            }
                                                                        case "foo bar":
                                                                            v207 = v212
                                                                        default:
                                                                            if v23 {
                                                                                switch -2147483648 {
                                                                                case -9223372036854775808:
                                                                                    if (!v23) {
                                                                                        if true {
                                                                                            var v272 float64 = v206
                                                                                            if ((^(v207)) <= 0) {
                                                                                                fmt.Println(v259, v272, fi258)
                                                                                                fmt.Println(v207)
                                                                                                fmt.Println(v212, v206)
                                                                                            } else {
                                                                                                switch ("\n\t" + "foo bar") {
                                                                                                case "\n\t":
                                                                                                    v1 = append(v1, "\n\t")
                                                                                                case "foo bar":
                                                                                                    v212 = (v207 << ((4294967295) & 63))
                                                                                                case "":
                                                                                                    if (3.14159 == (float32(v206) - 3.14159)) {
                                                                                                        switch "\n\t" {
                                                                                                        case "":
                                                                                                            v23 = (-1 < 31)
                                                                                                        case "foo bar":
                                                                                                            _ = len(v2)
                                                                                                        case "a":
                                                                                                            var v273 float32 = 1e10
                                                                                                            _ = v273
                                                                                                        default:
                                                                                                            _ = len(v3)
                                                                                                        }
                                                                                                    }
                                                                                                default:
                                                                                                    switch fi215 {
                                                                                                    case 908:
                                                                                                        v207 += v212
                                                                                                    case 828:
                                                                                                        v212 = 2147483647
                                                                                                    case 127:
                                                                                                        var v274 float64 = v272
                                                                                                        _ = v274
                                                                                                    default:
                                                                                                        var v275 float32 = 1e10
                                                                                                        _ = v275
                                                                                                    }
                                                                                                }
                                                                                                var v276 int8 = 100
                                                                                                fmt.Println(v206, fi258)
                                                                                                _ = v276
                                                                                            }
                                                                                            var v277 uint16 = 65535
                                                                                            _ = v272
                                                                                            _ = v277
                                                                                        } else {
                                                                                            v206 -= -2.25
                                                                                            v278 := GenS30{F0: 1, F1: 18446744073709551615, F2: -363}
                                                                                            switch fi215 {
                                                                                            case -2147483648:
                                                                                                v207 = 600
                                                                                            case -9223372036854775808:
                                                                                                switch -1 {
                                                                                                case -2147483648:
                                                                                                    fmt.Println(v212, v207, v207)
                                                                                                case 42:
                                                                                                    v206 = float64((0.0 * 0.1))
                                                                                                    v279 := GenS31{F0: 2, F1: 255, F2: 18446744073709551615}
                                                                                                    v206 -= v206
                                                                                                    _ = v279
                                                                                                default:
                                                                                                    v1 = append(v1, "")
                                                                                                }
                                                                                            case 128:
                                                                                                var v280 float64 = v206
                                                                                                _ = v280
                                                                                            default:
                                                                                                v2[(v207 % ((v212) | 1))] = v23
                                                                                            }
                                                                                            _ = v278
                                                                                        }
                                                                                        var v281 int = (-(fi258))
                                                                                        var v282 int = 0
                                                                                        _ = v281
                                                                                        _ = v282
                                                                                    }
                                                                                    fmt.Println(v207, v212)
                                                                                }
                                                                            } else {
                                                                                fmt.Println(v259, v206, v206)
                                                                                var v283 uint64 = 18446744073709551615
                                                                                var v284 bool = v23
                                                                                v285 := GenS19{F0: 127}
                                                                                _ = v283
                                                                                _ = v284
                                                                                _ = v285
                                                                            }
                                                                        }
                                                                    }
                                                                    if false {
                                                                        var v286 int32 = 275
                                                                        if len(v3) >= 2 {
                                                                            v3 = v3[1:len(v3)]
                                                                        }
                                                                        _ = v286
                                                                    }
                                                                }
                                                            }
                                                            _ = v266
                                                        }
                                                    case "hello":
                                                        var v287 int8 = -1
                                                        _ = v287
                                                    case "\n\t":
                                                        if false {
                                                            v23 = v23
                                                        } else {
                                                            if len(v1) > 0 {
                                                                v1[0] = "foo bar"
                                                            }
                                                            var v288 uint8 = 127
                                                            _ = v288
                                                        }
                                                    }
                                                    _ = len(v3)
                                                    switch (-(fi215)) {
                                                    case 128:
                                                        switch "\n\t" {
                                                        case "":
                                                            var v289 float64 = 1.5
                                                            _ = v289
                                                        case "foo bar":
                                                            v290 := GenS32{F0: 127}
                                                            _ = v290
                                                        }
                                                    case 42:
                                                        v259 += fi258
                                                    case 882:
                                                        v291 := GenS33{F0: 796}
                                                        _ = v291
                                                    default:
                                                        if (v23 && v23) {
                                                            var v292 int8 = 0
                                                            var v293 bool = (42 > 1000)
                                                            var v294 uint64 = 42
                                                            switch -893 {
                                                            case 42:
                                                                switch "" {
                                                                case "\n\t":
                                                                    if ((v293 && (("foo bar" == "") && false)) && v23) {
                                                                        switch -950 {
                                                                        case -925:
                                                                            v295 := v2[v212]
                                                                            _ = v295
                                                                        case -720:
                                                                            var v296 uint64 = (v294 - v294)
                                                                            var v297 string = "a"
                                                                            var v298 int32 = -2147483648
                                                                            _ = v296
                                                                            _ = v297
                                                                            _ = v298
                                                                        case 891:
                                                                            if len(v3) > 0 {
                                                                                v3[0] = (("a" + "\n\t") + ("\n\t" + ""))
                                                                            }
                                                                        }
                                                                        v23 = v23
                                                                        if v293 {
                                                                            var v299 int64 = 0
                                                                            _ = v299
                                                                        } else {
                                                                            v300 := GenS19{F0: (v292 >> ((18446744073709551615) & 63))}
                                                                            v301 := GenS34{F0: -32768, F1: 2, F2: 0}
                                                                            if (!(!false)) {
                                                                                fmt.Println(v23)
                                                                                switch v259 {
                                                                                case 382:
                                                                                    v302 := GenS14{F0: 1, F1: 42, F2: "foo bar"}
                                                                                    v207 -= (v212 % (((v207 - (v212 >> ((4294967295) & 63)))) | 1))
                                                                                    v292 += v292
                                                                                    _ = v302
                                                                                case 4294967295:
                                                                                    if ((1e10 < -2.25) || (3.14159 > -2.25)) {
                                                                                        var v303 int32 = -50
                                                                                        _ = v303
                                                                                    } else {
                                                                                        v304 := GenS6{F0: 1, F1: v294}
                                                                                        _ = v304
                                                                                    }
                                                                                case -927:
                                                                                    var v305 uint64 = (v294 | v294)
                                                                                    _ = v305
                                                                                default:
                                                                                    v294 = (v294 / (((v294 / ((v294) | 1))) | 1))
                                                                                }
                                                                            } else {
                                                                                fmt.Println(v207, v292)
                                                                                var v306 uint8 = 2
                                                                                _ = len(v1)
                                                                                if (v23 && (v292 != (v292 ^ v292))) {
                                                                                    v307 := GenS24{F0: (v207 - v212), F1: v212, F2: 1}
                                                                                    var v308 int = fi258
                                                                                    _ = v307
                                                                                    _ = v308
                                                                                } else {
                                                                                    v292 = 127
                                                                                }
                                                                                v294 = (v294 + v294)
                                                                                _ = v306
                                                                            }
                                                                            _ = v300
                                                                            _ = v301
                                                                        }
                                                                        var v309 uint = 255
                                                                        fmt.Println(v259, v23)
                                                                        _ = v309
                                                                    }
                                                                case "foo bar":
                                                                    fmt.Println(v23, v212)
                                                                case "a":
                                                                    if v23 {
                                                                        if len(v3) > 0 {
                                                                            v3[0] = "a"
                                                                        }
                                                                        switch "a" {
                                                                        case "\n\t":
                                                                            var v310 uint32 = (v212 + 65535)
                                                                            _ = v310
                                                                        case "":
                                                                            if v293 {
                                                                                v311 := GenS35{F0: 648, F1: 0, F2: (0.1 >= (((v206 + v206) * (1.5 * (v206 - (v206 + (v206 - float64(float32((v206 - v206)))))))) * v206))}
                                                                                if (66 == 65535) {
                                                                                    if ((-(fi258)) > (fi258 + v259)) {
                                                                                        var v312 bool = ((v292 << ((766) & 63)) > 0)
                                                                                        if v23 {
                                                                                            v259 = fi215
                                                                                        } else {
                                                                                            v293 = v293
                                                                                        }
                                                                                        v207 = (v207 - (^(v207)))
                                                                                        _ = v312
                                                                                    }
                                                                                    switch "a" {
                                                                                    case "foo bar":
                                                                                        v207 += (v207 & (v207 >> ((50) & 63)))
                                                                                    }
                                                                                } else {
                                                                                    var v313 bool = ((-2.25 <= (float64(0.0) / 0.1)) || ("\n\t" > "a"))
                                                                                    _ = len(v1)
                                                                                    _ = v313
                                                                                }
                                                                                _ = v311
                                                                            } else {
                                                                                switch "a" {
                                                                                case "hello":
                                                                                    switch "foo bar" {
                                                                                    case "":
                                                                                        _ = len(v2)
                                                                                    default:
                                                                                        v314 := v2[(v212 % ((v207) | 1))]
                                                                                        _ = v314
                                                                                    }
                                                                                    if ((v292 / ((v292) | 1)) == v292) {
                                                                                        delete(v2, 65535)
                                                                                    }
                                                                                    if (v292 < v292) {
                                                                                        switch (fi258 >> ((65535) & 63)) {
                                                                                        case 336:
                                                                                            v315 := GenS36{F0: 0, F1: 1}
                                                                                            _ = v315
                                                                                        default:
                                                                                            v316 := GenS37{F0: 0, F1: (v292 | (v292 | v292))}
                                                                                            _ = v316
                                                                                        }
                                                                                    }
                                                                                case "a":
                                                                                    fmt.Println(fi215, v206, v294)
                                                                                case "\n\t":
                                                                                    v292 = -40
                                                                                default:
                                                                                    v23 = v23
                                                                                }
                                                                                v317 := GenS38{F0: 255, F1: -1, F2: 128}
                                                                                v207 -= (v212 - 4294967295)
                                                                                var v318 float32 = -2.25
                                                                                var v319 int32 = 220
                                                                                _ = v317
                                                                                _ = v318
                                                                                _ = v319
                                                                            }
                                                                        case "a":
                                                                            v320 := GenS1{F0: fi215, F1: -1}
                                                                            _ = v320
                                                                        default:
                                                                            switch fi258 {
                                                                            case -1:
                                                                                if true {
                                                                                    switch "hello" {
                                                                                    case "hello":
                                                                                        fmt.Println(v206)
                                                                                    case "":
                                                                                        v321 := GenS25{F0: (-9223372036854775808 < 1), F1: 28, F2: 592}
                                                                                        var v322 string = "a"
                                                                                        _ = v321
                                                                                        _ = v322
                                                                                    case "foo bar":
                                                                                        if ((0.1 * 1e10) != 0.1) {
                                                                                            v323 := GenS39{F0: 610, F1: 65535, F2: fi215}
                                                                                            switch (fi215 | fi258) {
                                                                                            case 0:
                                                                                                v294 = (v294 >> ((0) & 63))
                                                                                            case 32767:
                                                                                                var v324 float64 = 0.1
                                                                                                _ = v324
                                                                                            case 2147483647:
                                                                                                var v325 uint = 4294967295
                                                                                                _ = v325
                                                                                            }
                                                                                            _ = v323
                                                                                        }
                                                                                    }
                                                                                    if (596 < 1000) {
                                                                                        _ = len(v2)
                                                                                        _ = len(v2)
                                                                                    } else {
                                                                                        _ = len(v2)
                                                                                    }
                                                                                    var v326 int64 = -217
                                                                                    _ = v326
                                                                                } else {
                                                                                    var v327 uint = 65535
                                                                                    if (v293 || true) {
                                                                                        var v328 int = (^(fi258))
                                                                                        if (v293 || ("" >= "foo bar")) {
                                                                                            fmt.Println(v259)
                                                                                            var v329 uint8 = 29
                                                                                            _ = v329
                                                                                        } else {
                                                                                            v327 += (^(v327))
                                                                                            v294 -= v294
                                                                                            switch (v328 - 2147483647) {
                                                                                            case 2147483647:
                                                                                                v330 := GenS39{F0: 42, F1: v327, F2: (-(fi258))}
                                                                                                _ = v330
                                                                                            case 1:
                                                                                                if v293 {
                                                                                                    if true {
                                                                                                        var v331 int16 = 32767
                                                                                                        v212 = v207
                                                                                                        v332 := GenS12{F0: (v327 & (v327 << ((v327) & 63))), F1: -1}
                                                                                                        _ = v331
                                                                                                        _ = v332
                                                                                                    }
                                                                                                    v207 -= v207
                                                                                                    var v333 uint = v327
                                                                                                    _ = v333
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                        var v334 string = "a"
                                                                                        v334 = (v334 + "hello")
                                                                                        var v335 uint8 = uint8(v294)
                                                                                        v336 := GenS40{F0: (^(v327)), F1: 32767, F2: v335}
                                                                                        var v337 float64 = v206
                                                                                        var v338 bool = v293
                                                                                        _ = v328
                                                                                        _ = v334
                                                                                        _ = v335
                                                                                        _ = v336
                                                                                        _ = v337
                                                                                        _ = v338
                                                                                    }
                                                                                    _ = v327
                                                                                }
                                                                            default:
                                                                                v339 := GenS41{F0: (v292 << ((18446744073709551615) & 63)), F1: -1, F2: (true && (-692 < 1000))}
                                                                                _ = v339
                                                                            }
                                                                        }
                                                                        if (65535 > 4294967295) {
                                                                            v259 = (fi215 ^ (fi215 << ((1) & 63)))
                                                                        }
                                                                        v259 = (v259 & (v259 / ((v259) | 1)))
                                                                    }
                                                                default:
                                                                    v2[v212] = (v294 == 42)
                                                                }
                                                            default:
                                                                v206 = (v206 + ((((v206 / v206) / 1.5) * (-2.25 + (v206 - float64(-2.25)))) * 0.1))
                                                            }
                                                            if v293 {
                                                                var v340 int16 = 42
                                                                var v341 int = -783
                                                                _ = v340
                                                                _ = v341
                                                            }
                                                            _ = v292
                                                            _ = v293
                                                            _ = v294
                                                        }
                                                        v207 = v212
                                                        var v342 uint32 = v212
                                                        _ = v342
                                                    }
                                                    if len(v1) > 0 {
                                                        v1[0] = (("foo bar" + "") + ("\n\t" + ""))
                                                    }
                                                    var v343 float64 = ((((float64(0.1) / (((v206 - v206) / v206) * v206)) * 1.5) + 0.1) + 1e10)
                                                    _ = v343
                                                }
                                                var v344 float32 = 1.5
                                                var v345 int32 = 0
                                                v346 := GenS42{F0: 183, F1: (v345 * (v345 >> ((91) & 63))), F2: 761}
                                                _ = v344
                                                _ = v345
                                                _ = v346
                                            }
                                        }
                                        var v347 uint64 = 200
                                        _ = v347
                                    }
                                    _ = v259
                                } else {
                                    if (!(1 != 2)) {
                                        if ((32767 < 0) && ("\n\t" == ("a" + (("\n\t" + "\n\t") + (("hello" + ("a" + ("hello" + "\n\t"))) + "foo bar"))))) {
                                            fmt.Println(v23, fi215, v212)
                                            v207 += v207
                                        } else {
                                            v348 := v2[v212]
                                            _ = v348
                                        }
                                        var v349 bool = (!(0.1 >= 0.1))
                                        var v350 uint16 = 925
                                        if v23 {
                                            if false {
                                                v212 -= (v212 * v207)
                                                v349 = v23
                                                var v351 bool = v23
                                                _ = v351
                                            }
                                        } else {
                                            v350 = 0
                                            v350 = (v350 | 1)
                                            v2[1] = v349
                                            var v352 int16 = -32768
                                            if (127 < (fi258 - (fi215 % (((fi215 << ((395) & 63))) | 1)))) {
                                                var v353 bool = (0.0 < (3.14159 / 0.1))
                                                _ = v353
                                            }
                                            var v354 uint16 = (v350 - v350)
                                            _ = v352
                                            _ = v354
                                        }
                                        v212 += v207
                                        _ = v349
                                        _ = v350
                                    }
                                }
                                var v355 uint64 = 940
                                if (v23 || ((1 >= 107) || (4294967295 >= 658))) {
                                    fmt.Println(fi215, v23)
                                }
                                if ((v355 / (((^(v355))) | 1)) == 1) {
                                    v356 := GenS16{F0: v207, F1: 169}
                                    if (42 == 2147483647) {
                                        switch "foo bar" {
                                        case "hello":
                                            if false {
                                                v355 += 42
                                                switch "hello" {
                                                case "":
                                                    var v357 uint32 = v207
                                                    _ = v357
                                                case "\n\t":
                                                    v358 := GenS34{F0: -658, F1: 65535, F2: 1}
                                                    _ = v358
                                                case "foo bar":
                                                    v2[(v212 << ((995) & 63))] = true
                                                }
                                                if (!v23) {
                                                    fmt.Println(v212)
                                                    var v359 uint16 = 281
                                                    var v360 string = ("" + (("a" + "foo bar") + ""))
                                                    _ = v359
                                                    _ = v360
                                                } else {
                                                    fmt.Println(fi258, v212, v212)
                                                }
                                            }
                                        case "a":
                                            v212 = (v212 % ((417) | 1))
                                        case "":
                                            fmt.Println(v23, fi258)
                                            _ = len(v2)
                                            var v361 int64 = 9223372036854775807
                                            _ = v361
                                        default:
                                            var v362 int = 9223372036854775807
                                            _ = v362
                                        }
                                    }
                                    var v363 int64 = -9223372036854775808
                                    _ = v356
                                    _ = v363
                                } else {
                                    var v364 bool = (!((214 < 2) && v23))
                                    fmt.Println(v23, fi258)
                                    v365 := GenS30{F0: -2147483648, F1: 18446744073709551615, F2: 2147483647}
                                    _ = v364
                                    _ = v365
                                }
                                v23 = v23
                                _ = fi258
                                _ = v355
                            }
                            v366 := GenS43{F0: 718}
                            _ = v366
                        }
                    }
                    fmt.Println(v206, v206)
                    _ = v240
                }
            default:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            fmt.Println(v23, v206, v23)
            _ = fi215
        }
        v207 = v212
    }
    v367 := GenS44{F0: -2147483648}
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v23 = false
    if (float32(v206) >= (((0.1 + 1.5) - 0.0) - 0.0)) {
        var v368 int64 = 411
        v206 = 0.0
        if v23 {
            v369 := GenS42{F0: 128, F1: -869, F2: 185}
            if (v206 == float64(1.5)) {
                if (!v23) {
                    if ((v23 || (65535 >= 255)) || v23) {
                        fmt.Println(v207, v207)
                    } else {
                        v370 := GenS14{F0: 884, F1: 306, F2: "a"}
                        if ((v207 / ((65535) | 1)) == (v207 / ((v212) | 1))) {
                            v368 += -175
                            if ((!v23) || v23) {
                                var v371 uint32 = (v212 & (^(v207)))
                                if ((32767 > 947) || ((-2.25 != 1e10) && (2 <= 0))) {
                                    for fi372 := 0; fi372 < 15; fi372++ {
                                        v207 = v207
                                        _ = fi372
                                    }
                                    if v23 {
                                        v368 += -1
                                        v212 -= (v207 & 65535)
                                        fmt.Println(v207)
                                    } else {
                                        v373 := GenS45{F0: 942, F1: 222}
                                        v374 := GenS46{F0: 65535, F1: v23}
                                        var v375 uint8 = 42
                                        _ = v373
                                        _ = v374
                                        _ = v375
                                    }
                                    for fi376 := 0; fi376 < 10; fi376++ {
                                        var v377 float32 = (1.5 + float32(v206))
                                        v377 += v377
                                        _ = fi376
                                        _ = v377
                                    }
                                }
                                var v378 float32 = (((1e10 * float32((v206 / (v206 * -2.25)))) / (1e10 * (-2.25 - (float32(0.0) / -2.25)))) + (((1e10 - ((float32(v206) / 1e10) + ((1.5 + -2.25) + 3.14159))) - float32(v206)) / float32((v206 - 0.1))))
                                _ = v371
                                _ = v378
                            } else {
                                v379 := GenS15{F0: 789, F1: 2, F2: 1}
                                _ = v379
                            }
                        }
                        _ = v370
                    }
                    if (v23 || false) {
                        var v380 uint = 42
                        var v381 int16 = -497
                        _ = v380
                        _ = v381
                    } else {
                        fmt.Println(v23, v206)
                    }
                }
                var v382 float32 = (1e10 * -2.25)
                fmt.Println(v23, v368)
                _ = v382
            }
            _ = v369
        }
        _ = v368
    } else {
        if true {
            for fi383 := 0; fi383 < 10; fi383++ {
                v384 := GenS33{F0: 42}
                _ = fi383
                _ = v384
            }
        }
        var v385 bool = (-121 > 100)
        _ = v385
    }
    switch 32767 {
    case -591:
        var v386 int64 = 9223372036854775807
        _ = v386
    case 0:
        for fi387 := 0; fi387 < 8; fi387++ {
            var v388 uint = 65535
            var v389 float32 = 3.14159
            _ = fi387
            _ = v388
            _ = v389
        }
    default:
        switch -1 {
        case 128:
            v1 = append(v1, "a")
        case 4294967295:
            fmt.Println(v207, v212, v206)
        case -1:
            if (-128 == 127) {
                v390 := GenS42{F0: 1, F1: 903, F2: 2}
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                for fi391 := 0; fi391 < 10; fi391++ {
                    v207 = v207
                    _ = fi391
                }
                _ = v390
            }
        }
    }
    fmt.Println(v23)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v206:", v206)
    fmt.Println("v207:", v207)
    fmt.Println("v212:", v212)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v206
    _ = v207
    _ = v211
    _ = v212
    _ = v23
    _ = v3
    _ = v367
}
