package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 string }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 uint64; F1 uint8; F2 bool }
type GenS5 struct { F0 int }
type GenS6 struct { F0 int32 }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 uint }
type GenS9 struct { F0 bool; F1 uint16; F2 int16 }
type GenS10 struct { F0 int32 }
type GenS11 struct { F0 uint; F1 uint8 }
type GenS12 struct { F0 bool }
func genFn0(p0 uint) string {
    p0 = p0
    v4 := GenS1{F0: 2, F1: "hello"}
    v5 := []int{}
    var v6 float64 = 1e10
    _ = p0
    _ = v4
    _ = v5
    _ = v6
    return ("" + "hello")
}

func main() {
    v1 := []int{32767, -2147483648}
    v2 := make(map[int]int)
    var v3 uint = 18446744073709551615
    v3 += (v3 * v3)
    if true {
        var v7 uint = 653
        v3 = v7
        var v8 string = "a"
        _ = v7
        _ = v8
    } else {
        var v9 float32 = 3.14159
        for fi10 := 0; fi10 < 11; fi10++ {
            for fi11 := 0; fi11 < 17; fi11++ {
                var v12 int32 = 1
                var v13 int32 = -2147483648
                _ = fi11
                _ = v12
                _ = v13
            }
            _ = fi10
        }
        v9 += v9
        _ = v9
    }
    v2[32767] = 127
    var v14 uint8 = 42
    if true {
        for fi15 := 0; fi15 < 3; fi15++ {
            var v16 int16 = 32767
            for fi17 := 0; fi17 < 18; fi17++ {
                v16 = (v16 | (v16 | v16))
                _ = fi17
            }
            _ = fi15
            _ = v16
        }
    }
    if (!(108 == 42)) {
        if false {
            var v18 int8 = 127
            if len(v1) > 0 {
                v1[0] = 32768
            }
            _ = v18
        } else {
            v14 = (^(v14))
            v19 := GenS1{F0: 792, F1: ("hello" + (("" + ("a" + "foo bar")) + ""))}
            _ = v19
        }
    }
    for fi20 := 0; fi20 < 3; fi20++ {
        var v21 int16 = 0
        _ = fi20
        _ = v21
    }
    _ = len(v1)
    if (int64(v3) != 42) {
        v22 := GenS1{F0: 32767, F1: "a"}
        if false {
            var v23 int8 = -1
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            _ = v23
        } else {
            for fi24 := 0; fi24 < 12; fi24++ {
                switch -2147483648 {
                case 2147483647:
                    var v25 uint = v3
                    _ = v25
                case -557:
                    fmt.Println(fi24, v14)
                case -9223372036854775808:
                    var v26 uint = uint(v14)
                    _ = v26
                default:
                    switch (fi24 / (((fi24 + (fi24 % (((fi24 << (((v3 % ((4294967295) | 1))) & 63))) | 1)))) | 1)) {
                    case 0:
                        for fi27 := 0; fi27 < 6; fi27++ {
                            var v28 uint = v3
                            switch "" {
                            case "a":
                                if (((float64(1e10) * (1.5 - 1.5)) / (3.14159 + 0.0)) <= 0.0) {
                                    if true {
                                        var v29 int16 = -32768
                                        _ = v29
                                    } else {
                                        fmt.Println(v28)
                                        fmt.Println(fi27, fi24, v3)
                                    }
                                    if ((true && (0 > 42)) && (!false)) {
                                        switch ("hello" + "a") {
                                        case "hello":
                                            fmt.Println(v3, fi24)
                                        default:
                                            if false {
                                                if (-128 < -128) {
                                                    if true {
                                                        if (-1 == -1) {
                                                            v30 := GenS1{F0: 32767, F1: "foo bar"}
                                                            v31 := GenS2{F0: (1 == 4294967295), F1: 842}
                                                            v32 := GenS2{F0: false, F1: -224}
                                                            _ = v30
                                                            _ = v31
                                                            _ = v32
                                                        } else {
                                                            var v33 uint32 = 42
                                                            _ = v33
                                                        }
                                                        var v34 uint = v3
                                                        _ = v34
                                                    }
                                                }
                                                var v35 int32 = -2147483648
                                                _ = v35
                                            } else {
                                                if (true || true) {
                                                    var v36 uint64 = 0
                                                    v37 := v2[9223372036854775807]
                                                    if false {
                                                        fmt.Println(v36, v14)
                                                    }
                                                    fmt.Println(v37, v37)
                                                    _ = v36
                                                    _ = v37
                                                } else {
                                                    v14 = v14
                                                }
                                            }
                                        }
                                    }
                                    v28 = (v28 % ((65535) | 1))
                                    v38 := GenS3{F0: (42 <= 2147483647)}
                                    _ = v38
                                }
                            case "\n\t":
                                v39 := GenS4{F0: 42, F1: v14, F2: false}
                                _ = v39
                            case "":
                                v14 += (^(v14))
                                var v40 int16 = 1
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                                _ = v40
                            }
                            _ = fi27
                            _ = v28
                        }
                    case 32767:
                        v3 = 781
                    case 4294967295:
                        for fi41 := 0; fi41 < 9; fi41++ {
                            _ = len(v1)
                            switch 42 {
                            case 9223372036854775807:
                                v42 := v2[(^(fi41))]
                                _ = v42
                            case 2147483647:
                                v43 := GenS5{F0: (-(fi24))}
                                _ = v43
                            }
                            v3 += (v3 + (v3 >> ((18446744073709551615) & 63)))
                            v3 -= 456
                            _ = fi41
                        }
                    default:
                        v44 := GenS6{F0: 1}
                        _ = v44
                    }
                }
                switch "" {
                case "a":
                    if true {
                        if true {
                            switch "" {
                            case "\n\t":
                                if true {
                                    if (85 < 1) {
                                        v3 -= (v3 & v3)
                                        v45 := GenS3{F0: (-407 > -223)}
                                        _ = v45
                                    } else {
                                        if len(v1) > 0 {
                                            v1[0] = -478
                                        }
                                        if ((fi24 + (-(fi24))) != 32768) {
                                            v46 := GenS7{F0: 1}
                                            var v47 uint32 = 657
                                            _ = v46
                                            _ = v47
                                        }
                                        var v48 uint64 = 0
                                        _ = v48
                                    }
                                    switch (("a" + "a") + "hello") {
                                    case "\n\t":
                                        if (((^(v14)) >= v14) && true) {
                                            for fi49 := 0; fi49 < 6; fi49++ {
                                                v50 := GenS8{F0: 2}
                                                _ = fi49
                                                _ = v50
                                            }
                                        }
                                    default:
                                        var v51 int32 = -694
                                        _ = v51
                                    }
                                    v52 := GenS3{F0: true}
                                    _ = v52
                                }
                            case "hello":
                                if true {
                                    var v53 uint16 = 274
                                    _ = v53
                                }
                            default:
                                v3 = (v3 % ((v3) | 1))
                            }
                            var v54 int64 = 1
                            _ = v54
                        }
                        switch "" {
                        case "\n\t":
                            v14 = 1
                        case "":
                            if true {
                                for fi55 := 0; fi55 < 20; fi55++ {
                                    var v56 uint8 = 0
                                    _ = fi55
                                    _ = v56
                                }
                            } else {
                                var v57 int16 = 42
                                if true {
                                    for fi58 := 0; fi58 < 15; fi58++ {
                                        var v59 int = 2147483647
                                        if (3.14159 == 3.14159) {
                                            var v60 uint32 = 4294967295
                                            v59 -= (fi58 * 32767)
                                            _ = v60
                                        } else {
                                            v59 += v59
                                            if len(v1) > 0 {
                                                v1[0] = v59
                                            }
                                        }
                                        _ = fi58
                                        _ = v59
                                    }
                                }
                                var v61 int64 = -9223372036854775808
                                _ = v57
                                _ = v61
                            }
                        case "foo bar":
                            v3 = 42
                        default:
                            var v62 int64 = 0
                            _ = v62
                        }
                    }
                default:
                    v14 = v14
                }
                _ = fi24
            }
            var v63 uint = (v3 & 1)
            _ = v63
        }
        _ = v22
    }
    switch 9223372036854775807 {
    case 32767:
        switch -2147483648 {
        case 128:
            for fi64 := 0; fi64 < 4; fi64++ {
                v14 = uint8(v3)
                _ = fi64
            }
        case -2147483648:
            switch 2147483647 {
            case 2147483647:
                if (0.1 > 1.5) {
                    var v65 string = (("a" + "a") + "hello")
                    switch (v65 + v65) {
                    case "foo bar":
                        v66 := GenS4{F0: 2, F1: (^(v14)), F2: ((3.14159 != 1.5) && true)}
                        _ = v66
                    case "\n\t":
                        switch 4294967295 {
                        case 127:
                            switch "" {
                            case "\n\t":
                                v14 += 255
                            case "hello":
                                var v67 int16 = -864
                                _ = v67
                            default:
                                var v68 uint = 45
                                _ = v68
                            }
                        }
                    case "a":
                        v2[0] = 127
                    }
                    v3 -= v3
                    _ = v65
                } else {
                    if (true && (!(!true))) {
                        v14 = v14
                        v3 = (^(v3))
                    } else {
                        fmt.Println(v14)
                        v3 = 4294967295
                        if ((-2.25 < 1e10) || (!true)) {
                            for fi69 := 0; fi69 < 6; fi69++ {
                                var v70 uint64 = 42
                                for fi71 := 0; fi71 < 7; fi71++ {
                                    var v72 float32 = (0.0 * (1e10 + (0.0 * 1e10)))
                                    v72 -= v72
                                    var v73 int8 = 1
                                    _ = fi71
                                    _ = v72
                                    _ = v73
                                }
                                _ = fi69
                                _ = v70
                            }
                            if false {
                                v14 = 42
                                var v74 uint64 = 356
                                _ = v74
                            } else {
                                for fi75 := 0; fi75 < 20; fi75++ {
                                    var v76 int32 = 1
                                    if true {
                                        if true {
                                            v76 = v76
                                            v77 := GenS9{F0: (false && ((65535 < 42) || ((1 >= 0) && (((!true) || true) && false)))), F1: 255, F2: 32767}
                                            v1 = append(v1, (^(fi75)))
                                            _ = v77
                                        } else {
                                            for fi78 := 0; fi78 < 10; fi78++ {
                                                v14 = (v14 + (^(v14)))
                                                _ = fi78
                                            }
                                        }
                                        v3 -= uint(v14)
                                        fmt.Println(v76, v76)
                                    }
                                    var v79 uint8 = v14
                                    _ = fi75
                                    _ = v76
                                    _ = v79
                                }
                                for fi80 := 0; fi80 < 15; fi80++ {
                                    switch 337 {
                                    case 128:
                                        var v81 int16 = 1
                                        _ = v81
                                    case 4294967295:
                                        v82 := GenS10{F0: -2147483648}
                                        _ = v82
                                    case 32768:
                                        _ = len(v2)
                                    default:
                                        v14 += v14
                                    }
                                    delete(v2, (-(fi80)))
                                    if true {
                                        if false {
                                            fmt.Println(v14, fi80)
                                            for fi83 := 0; fi83 < 7; fi83++ {
                                                v84 := GenS2{F0: (-13 > 65), F1: 0}
                                                v3 -= (^(v3))
                                                var v85 float64 = ((((((1e10 * 0.0) / 1.0) / 1e10) * (-2.25 - ((3.14159 + 1.5) - (1.5 / (0.1 + -2.25))))) * 1.5) / (1e10 / (3.14159 * (-2.25 - 0.0))))
                                                v14 -= v14
                                                _ = fi83
                                                _ = v84
                                                _ = v85
                                            }
                                        }
                                        fmt.Println(v3)
                                        if true {
                                            var v86 int = (-(fi80))
                                            for fi87 := 0; fi87 < 2; fi87++ {
                                                v14 = v14
                                                v86 = (v86 * v86)
                                                _ = fi87
                                            }
                                            v2[fi80] = (^(fi80))
                                            _ = v86
                                        }
                                    }
                                    _ = fi80
                                }
                            }
                        }
                    }
                    switch (("foo bar" + ("\n\t" + "foo bar")) + "\n\t") {
                    case "foo bar":
                        v3 += 889
                    case "a":
                        var v88 uint8 = (v14 >> ((v3) & 63))
                        _ = v88
                    case "hello":
                        if (true || (0 < 613)) {
                            switch "\n\t" {
                            case "\n\t":
                                var v89 int8 = 36
                                _ = v89
                            case "a":
                                for fi90 := 0; fi90 < 3; fi90++ {
                                    switch fi90 {
                                    case -942:
                                        if true {
                                            v14 += (v14 * v14)
                                            var v91 int = fi90
                                            switch fi90 {
                                            case 75:
                                                v1 = append(v1, (fi90 + (-(fi90))))
                                            case 4294967295:
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            }
                                            _ = v91
                                        } else {
                                            var v92 uint32 = uint32(v3)
                                            _ = v92
                                        }
                                    case 949:
                                        v93 := GenS1{F0: 2, F1: (("hello" + "\n\t") + "hello")}
                                        _ = v93
                                    default:
                                        v2[-118] = (^(fi90))
                                    }
                                    v94 := GenS1{F0: 784, F1: "hello"}
                                    if (true && false) {
                                        v1 = append(v1, (fi90 ^ (^(fi90))))
                                    } else {
                                        switch fi90 {
                                        case 9223372036854775807:
                                            if false {
                                                fmt.Println(fi90)
                                                var v95 uint32 = 2
                                                _ = v95
                                            } else {
                                                switch (fi90 | (-(fi90))) {
                                                case 2147483647:
                                                    v14 = 241
                                                case 42:
                                                    delete(v2, fi90)
                                                case 32768:
                                                    if (!(((0.1 * 3.14159) / float32(3.14159)) != 0.1)) {
                                                        var v96 uint = (v3 + v3)
                                                        v14 = v14
                                                        _ = v96
                                                    }
                                                }
                                                v3 = v3
                                            }
                                        case 0:
                                            for fi97 := 0; fi97 < 8; fi97++ {
                                                if false {
                                                    delete(v2, 823)
                                                } else {
                                                    v98 := v2[(^(fi90))]
                                                    _ = v98
                                                }
                                                v3 = v3
                                                var v99 uint16 = 1
                                                _ = fi97
                                                _ = v99
                                            }
                                        case 32768:
                                            delete(v2, (fi90 + (fi90 | fi90)))
                                        }
                                        fmt.Println(v3, v14)
                                        v14 = 42
                                    }
                                    _ = fi90
                                    _ = v94
                                }
                            case "hello":
                                for fi100 := 0; fi100 < 11; fi100++ {
                                    if true {
                                        v3 = 135
                                        if true {
                                            v101 := GenS9{F0: (42 != 596), F1: 244, F2: -552}
                                            _ = v101
                                        }
                                        var v102 int16 = 42
                                        _ = v102
                                    } else {
                                        var v103 int64 = 347
                                        for fi104 := 0; fi104 < 17; fi104++ {
                                            switch ("\n\t" + ("foo bar" + (("" + "") + (("a" + "hello") + "\n\t")))) {
                                            case "a":
                                                fmt.Println(fi100)
                                            case "foo bar":
                                                v103 += v103
                                            case "":
                                                delete(v2, -9223372036854775808)
                                            }
                                            v14 -= v14
                                            _ = fi104
                                        }
                                        _ = v103
                                    }
                                    if false {
                                        for fi105 := 0; fi105 < 9; fi105++ {
                                            var v106 uint16 = 255
                                            switch -325 {
                                            case -2147483648:
                                                fmt.Println(fi100, fi105)
                                            case 542:
                                                fmt.Println(fi100, v106, v106)
                                                v2[fi105] = fi100
                                                v3 = 255
                                            case -1:
                                                v107 := GenS5{F0: 263}
                                                _ = v107
                                            }
                                            _ = fi105
                                            _ = v106
                                        }
                                        for fi108 := 0; fi108 < 20; fi108++ {
                                            if false {
                                                delete(v2, (fi100 | 921))
                                                v109 := GenS11{F0: v3, F1: v14}
                                                _ = v109
                                            }
                                            _ = fi108
                                        }
                                        fmt.Println(v3, v14)
                                    } else {
                                        switch "a" {
                                        case "foo bar":
                                            var v110 int8 = 100
                                            _ = v110
                                        case "\n\t":
                                            v3 += (v3 << ((v3) & 63))
                                        }
                                    }
                                    _ = fi100
                                }
                            }
                            v2[453] = -9223372036854775808
                        }
                    default:
                        var v111 uint64 = 4294967295
                        _ = v111
                    }
                }
            case -820:
                v2[-2147483648] = -1
            }
        case 32768:
            var v112 bool = true
            _ = v112
        default:
            v2[127] = 9223372036854775807
        }
    case 4294967295:
        var v113 float64 = (((3.14159 * float64(0.1)) - 1.5) / 0.1)
        _ = v113
    case 32768:
        v14 = v14
    default:
        switch "\n\t" {
        case "foo bar":
            v114 := GenS12{F0: (503 > 776)}
            _ = v114
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
}
