package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8 }
type GenS2 struct { F0 int16; F1 int32 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int16; F1 int32 }
type GenS7 struct { F0 uint32 }

func main() {
    v1 := make(map[string]string)
    var v2 int = 32767
    v3 := make(map[uint32]string)
    var v4 float64 = -2.25
    v5 := make(map[uint32]bool)
    var v6 int16 = 0
    for fi7 := 0; fi7 < 14; fi7++ {
        for fi8 := 0; fi8 < 8; fi8++ {
            var v9 uint = 255
            _ = fi8
            _ = v9
        }
        _ = fi7
    }
    var v10 uint8 = 198
    var v11 uint32 = 4294967295
    switch "a" {
    case "hello":
        switch "foo bar" {
        case "\n\t":
            v12 := v3[(v11 >> ((255) & 63))]
            _ = v12
        case "hello":
            if true {
                for fi13 := 0; fi13 < 7; fi13++ {
                    switch "foo bar" {
                    case "":
                        for fi14 := 0; fi14 < 16; fi14++ {
                            v11 = (^(v11))
                            var v15 int = (fi13 / ((fi14) | 1))
                            _ = len(v3)
                            _ = fi14
                            _ = v15
                        }
                    case "\n\t":
                        switch ("hello" + "") {
                        case "hello":
                            for fi16 := 0; fi16 < 5; fi16++ {
                                var v17 int64 = 0
                                var v18 string = ""
                                v5[2] = (false || true)
                                _ = fi16
                                _ = v17
                                _ = v18
                            }
                        case "":
                            if ((3.14159 < (-2.25 * 3.14159)) || (!(!false))) {
                                fmt.Println(v10)
                                for fi19 := 0; fi19 < 14; fi19++ {
                                    if false {
                                        v2 -= (v2 - -1)
                                    }
                                    v6 = (v6 & v6)
                                    _ = fi19
                                }
                                if ((fi13 / (((v2 * (fi13 >> ((65535) & 63)))) | 1)) > (fi13 / (((^(fi13))) | 1))) {
                                    var v20 uint64 = 0
                                    var v21 int64 = -9223372036854775808
                                    v20 -= (v20 | 852)
                                    _ = v20
                                    _ = v21
                                } else {
                                    var v22 int32 = -2147483648
                                    var v23 uint64 = 208
                                    _ = v22
                                    _ = v23
                                }
                            }
                        case "a":
                            var v24 bool = (32767 < 32767)
                            _ = v24
                        }
                    case "hello":
                        var v25 uint = 4294967295
                        _ = v25
                    default:
                        v6 = v6
                    }
                    _ = fi13
                }
                fmt.Println(v11, v11, v2)
            } else {
                v4 = (v4 / 1.0)
            }
        case "a":
            v10 += (v10 << ((468) & 63))
        }
    case "a":
        for fi26 := 0; fi26 < 5; fi26++ {
            var v27 uint32 = (v11 % ((uint32(v10)) | 1))
            switch "a" {
            case "foo bar":
                v28 := GenS1{F0: -9223372036854775808, F1: v10}
                _ = v28
            case "a":
                for fi29 := 0; fi29 < 9; fi29++ {
                    fmt.Println(v10)
                    _ = fi29
                }
            default:
                if (v10 <= (v10 * (v10 << ((0) & 63)))) {
                    v6 = 1000
                    for fi30 := 0; fi30 < 6; fi30++ {
                        v31 := v3[(^(v27))]
                        if (2147483648 <= uint32(v10)) {
                            var v32 uint = 0
                            if true {
                                if ((1 != int32(v11)) && false) {
                                    fmt.Println(v6, v10, v27)
                                    v33 := v3[191]
                                    _ = v33
                                } else {
                                    v11 -= v27
                                }
                                switch fi26 {
                                case 127:
                                    if (true && false) {
                                        var v34 int32 = -2147483648
                                        switch (v2 | (^(v2))) {
                                        case 2147483647:
                                            var v35 int32 = v34
                                            _ = v35
                                        case -1:
                                            if true {
                                                switch (-(fi30)) {
                                                case -195:
                                                    v36 := GenS1{F0: 0, F1: 2}
                                                    _ = v36
                                                }
                                                var v37 int32 = v34
                                                fmt.Println(v11, v27)
                                                _ = v37
                                            }
                                        case 127:
                                            v38 := []int{fi30}
                                            _ = v38
                                        default:
                                            if ((float64((-2.25 - 3.14159)) * (((v4 + v4) + (float64((float32(v4) / 1e10)) * ((((v4 - float64(1e10)) + v4) / ((v4 / (3.14159 / v4)) * v4)) * 1.5))) - (0.0 / v4))) >= ((-2.25 * float64(-2.25)) + v4)) {
                                                delete(v3, uint32(v10))
                                            } else {
                                                if true {
                                                    _ = len(v3)
                                                    v39 := GenS1{F0: 9223372036854775807, F1: v10}
                                                    _ = v39
                                                } else {
                                                    if (4294967295 > 287) {
                                                        switch fi30 {
                                                        case 207:
                                                            delete(v3, uint32(v32))
                                                            fmt.Println(fi30, v10)
                                                            fmt.Println(fi30, v11)
                                                        case 32767:
                                                            if false {
                                                                var v40 uint16 = 2
                                                                fmt.Println(v6)
                                                                if ((v6 % (((v6 + -32768)) | 1)) >= v6) {
                                                                    v4 += 0.1
                                                                } else {
                                                                    v11 += (v11 - 872)
                                                                    v32 -= v32
                                                                    if true {
                                                                        v6 += 1000
                                                                        v41 := []int{364, fi26, fi26}
                                                                        v11 = v27
                                                                        _ = v41
                                                                    }
                                                                }
                                                                fmt.Println(v31, fi26, v32)
                                                                _ = v40
                                                            }
                                                        case 1:
                                                            switch 127 {
                                                            case -755:
                                                                var v42 int16 = v6
                                                                _ = v42
                                                            default:
                                                                delete(v1, (v31 + v31))
                                                            }
                                                        }
                                                        var v43 int8 = 42
                                                        _ = v43
                                                    }
                                                    fmt.Println(v27)
                                                }
                                                v4 = 1e10
                                                delete(v3, 553)
                                            }
                                            v31 = v31
                                            v44 := GenS1{F0: 462, F1: (v10 % (((v10 * (v10 >> (((v32 >> ((v32) & 63))) & 63)))) | 1))}
                                            _ = v44
                                        }
                                        _ = v34
                                    }
                                default:
                                    v31 = "\n\t"
                                }
                            } else {
                                var v45 int32 = 1
                                fmt.Println(v2)
                                var v46 float64 = v4
                                _ = v45
                                _ = v46
                            }
                            _ = v32
                        }
                        _ = fi30
                        _ = v31
                    }
                }
            }
            switch (("" + ("foo bar" + (("hello" + "") + "a"))) + ("a" + "a")) {
            case "a":
                fmt.Println(fi26, v10, fi26)
            case "foo bar":
                delete(v5, (v11 >> ((56) & 63)))
            default:
                fmt.Println(v11, v4)
            }
            _ = fi26
            _ = v27
        }
    }
    if true {
        var v47 uint16 = 32767
        if true {
            switch -870 {
            case 0:
                switch ("" + "\n\t") {
                case "\n\t":
                    v48 := v3[(v11 >> ((2) & 63))]
                    _ = v48
                default:
                    fmt.Println(v47)
                }
            case 907:
                v6 = (v6 | v6)
            case 510:
                fmt.Println(v11, v47, v10)
            }
        }
        _ = v47
    } else {
        if (((v2 - -1) != (-(v2))) || (float32((float64(0.1) - 3.14159)) >= 0.0)) {
            v6 = v6
            var v49 float32 = 1e10
            for fi50 := 0; fi50 < 19; fi50++ {
                fmt.Println(v6, fi50)
                var v51 uint64 = 658
                _ = fi50
                _ = v51
            }
            _ = v49
        }
        v6 += (v6 / (((v6 & v6)) | 1))
        var v52 int32 = -2147483648
        _ = v52
    }
    if (("foo bar" + "foo bar") >= "") {
        var v53 uint = 18446744073709551615
        switch "\n\t" {
        case "a":
            var v54 bool = true
            _ = v54
        case "foo bar":
            switch (("a" + "\n\t") + ("foo bar" + "")) {
            case "":
                v55 := GenS1{F0: 9223372036854775807, F1: v10}
                _ = v55
            case "a":
                v6 += (v6 * v6)
            case "hello":
                v56 := v5[172]
                _ = v56
            }
        case "":
            for fi57 := 0; fi57 < 7; fi57++ {
                v11 = (^(v11))
                if true {
                    if true {
                        v58 := GenS1{F0: -689, F1: (v10 - (v10 | v10))}
                        switch (fi57 / ((fi57) | 1)) {
                        case 940:
                            v59 := GenS2{F0: (v6 / ((v6) | 1)), F1: 42}
                            _ = v59
                        case 2147483647:
                            var v60 bool = false
                            _ = v60
                        default:
                            v10 = (v10 | v10)
                        }
                        _ = v58
                    }
                    var v61 uint = v53
                    v2 = v2
                    _ = v61
                } else {
                    var v62 string = "a"
                    v4 = -2.25
                    for fi63 := 0; fi63 < 9; fi63++ {
                        v64 := v5[v11]
                        _ = fi63
                        _ = v64
                    }
                    _ = v62
                }
                v10 -= (v10 % ((uint8(v53)) | 1))
                _ = fi57
            }
        }
        _ = v53
    } else {
        if true {
            var v65 uint8 = v10
            v66 := GenS1{F0: -770, F1: 42}
            if false {
                for fi67 := 0; fi67 < 1; fi67++ {
                    v65 -= (v10 << ((792) & 63))
                    var v68 float64 = float64(3.14159)
                    _ = fi67
                    _ = v68
                }
                v1["foo bar"] = "\n\t"
            } else {
                v69 := GenS3{F0: v6}
                _ = v69
            }
            _ = v65
            _ = v66
        } else {
            v4 = (((v4 - ((v4 / v4) + (1e10 + -2.25))) - float64(1.5)) + (0.1 + 0.0))
            v11 += v11
        }
    }
    fmt.Println(v10)
    fmt.Println(v6, v11, v11)
    v70 := GenS3{F0: (v6 ^ v6)}
    for fi71 := 0; fi71 < 9; fi71++ {
        var v72 uint16 = 874
        v73 := v5[0]
        _ = fi71
        _ = v72
        _ = v73
    }
    if ((1 < 42) && (!true)) {
        if (-1 == 0) {
            v2 += v2
            var v74 float32 = float32((v4 + ((float64((1e10 / float32(float64(3.14159)))) - v4) + v4)))
            _ = v74
        } else {
            if true {
                fmt.Println(v2, v11, v11)
                v75 := GenS4{F0: v10}
                _ = v75
            } else {
                var v76 uint = 540
                switch ("" + ("hello" + "a")) {
                case "":
                    v76 += (v76 << (((v76 % ((255) | 1))) & 63))
                case "\n\t":
                    var v77 bool = true
                    _ = v77
                case "foo bar":
                    v2 = v2
                default:
                    switch (v2 / ((v2) | 1)) {
                    case 419:
                        switch (v2 | v2) {
                        case 0:
                            fmt.Println(v2)
                        case -2147483648:
                            var v78 int8 = -82
                            _ = v78
                        default:
                            if true {
                                if false {
                                    v3[(^(v11))] = "hello"
                                    switch v2 {
                                    case 826:
                                        v11 = (v11 ^ uint32(v76))
                                    case 860:
                                        v79 := GenS4{F0: 0}
                                        _ = v79
                                    }
                                } else {
                                    if true {
                                        if true {
                                            v80 := []int{(v2 | 396), (v2 * v2), (v2 / ((1) | 1)), v2}
                                            _ = v80
                                        }
                                        v81 := GenS5{F0: 89}
                                        if ((^(v2)) == (v2 / ((128) | 1))) {
                                            fmt.Println(v4, v6, v2)
                                            v4 = v4
                                            v82 := GenS6{F0: -1, F1: -521}
                                            _ = v82
                                        }
                                        _ = v81
                                    }
                                }
                                v83 := GenS4{F0: (v10 + 127)}
                                v5[(^(v11))] = (42 > (v10 % ((v10) | 1)))
                                _ = v83
                            } else {
                                fmt.Println(v11)
                                v76 = (v76 % ((v76) | 1))
                                v84 := GenS4{F0: (^(v10))}
                                _ = v84
                            }
                        }
                    case 1:
                        fmt.Println(v2, v76)
                    }
                }
                fmt.Println(v2)
                _ = v76
            }
        }
    } else {
        var v85 uint = 778
        var v86 float64 = (float64((3.14159 - (1.5 - ((((((float32(v4) + 0.0) / (1.5 / 0.1)) * 0.0) - float32(v4)) * 3.14159) / 1.0)))) - v4)
        v87 := GenS6{F0: 32767, F1: 0}
        _ = v85
        _ = v86
        _ = v87
    }
    if true {
        var v88 int = (v2 << ((18446744073709551615) & 63))
        var v89 int64 = 9223372036854775807
        if (!false) {
            v4 = (((1e10 / float64(float32(3.14159))) / 3.14159) - (-2.25 - ((v4 / v4) - (0.1 * v4))))
        } else {
            if true {
                switch (^(v88)) {
                case -387:
                    v10 += (^(v10))
                default:
                    var v90 uint16 = 42
                    _ = v90
                }
                delete(v3, v11)
                v91 := v1["a"]
                _ = v91
            } else {
                v92 := []int{-1}
                switch (v2 + (v88 >> ((65535) & 63))) {
                case 4294967295:
                    v2 = (v88 | v88)
                case 0:
                    for fi93 := 0; fi93 < 18; fi93++ {
                        for fi94 := 0; fi94 < 14; fi94++ {
                            fmt.Println(fi93)
                            v10 = v10
                            if (!false) {
                                v95 := GenS7{F0: 42}
                                var v96 uint16 = 65535
                                _ = v95
                                _ = v96
                            }
                            fmt.Println(v11, v11, v6)
                            if len(v92) > 0 {
                                v92[0] = int(v6)
                            }
                            v3[(v11 << ((42) & 63))] = "foo bar"
                            if (0 == int8(v10)) {
                                v4 -= v4
                                if true {
                                    var v97 uint8 = (v10 >> ((255) & 63))
                                    v98 := GenS7{F0: v11}
                                    _ = v97
                                    _ = v98
                                }
                            } else {
                                if len(v92) >= 2 {
                                    v92 = v92[1:len(v92)]
                                }
                                _ = len(v92)
                            }
                            _ = fi94
                        }
                        var v99 uint = 1
                        v100 := GenS5{F0: -93}
                        _ = fi93
                        _ = v100
                        _ = v99
                    }
                case 42:
                    var v101 bool = false
                    _ = v101
                default:
                    if false {
                        switch "hello" {
                        case "foo bar":
                            v4 = 0.1
                        default:
                            v102 := GenS1{F0: 822, F1: v10}
                            _ = v102
                        }
                        v89 = v89
                        for fi103 := 0; fi103 < 12; fi103++ {
                            fmt.Println(v88, v2, v10)
                            var v104 uint64 = 2
                            _ = fi103
                            _ = v104
                        }
                    }
                }
                _ = v92
            }
            _ = len(v3)
            switch ("\n\t" + ("\n\t" + ("\n\t" + (("" + "") + "hello")))) {
            case "hello":
                if (((1 != 4294967295) || false) && true) {
                    v105 := []int{(v2 & 4294967295), 127}
                    switch "hello" {
                    case "hello":
                        for fi106 := 0; fi106 < 16; fi106++ {
                            v3[(^(v11))] = "a"
                            if (745 <= 1) {
                                var v107 uint16 = 65535
                                v3[(v11 & uint32(v10))] = ""
                                _ = v107
                            } else {
                                for fi108 := 0; fi108 < 18; fi108++ {
                                    if (v89 == (v89 % ((-528) | 1))) {
                                        var v109 bool = (((true || (!false)) && true) || (("hello" + ("a" + "foo bar")) < "\n\t"))
                                        _ = v109
                                    }
                                    _ = len(v105)
                                    v110 := GenS6{F0: v6, F1: -157}
                                    var v111 float64 = (float64(3.14159) / 1.5)
                                    var v112 string = "foo bar"
                                    v88 -= (^(v2))
                                    _ = fi108
                                    _ = v110
                                    _ = v111
                                    _ = v112
                                }
                                fmt.Println(v10, v11)
                            }
                            _ = fi106
                        }
                    }
                    _ = v105
                }
            default:
                v11 -= (v11 / (((v11 - v11)) | 1))
            }
        }
        _ = v88
        _ = v89
    }
    var v113 uint64 = 4294967295
    var v114 uint16 = 42
    for fi115 := 0; fi115 < 15; fi115++ {
        fmt.Println(v6, v113, v4)
        for fi116 := 0; fi116 < 11; fi116++ {
            switch "" {
            case "":
                fmt.Println(v4, v2)
            }
            var v117 uint32 = (v11 ^ v11)
            _ = fi116
            _ = v117
        }
        fmt.Println(v4, v10)
        _ = fi115
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v113:", v113)
    fmt.Println("v114:", v114)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v113
    _ = v114
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v70
}
