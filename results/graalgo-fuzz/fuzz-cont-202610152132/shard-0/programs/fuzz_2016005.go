package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 uint8 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint16; F1 int64 }
type GenS4 struct { F0 int64; F1 bool }
type GenS5 struct { F0 uint; F1 int32 }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int64; F1 uint }
type GenS9 struct { F0 uint; F1 bool; F2 bool }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 int8; F1 int }
type GenS12 struct { F0 uint64; F1 int8; F2 int64 }
type GenS13 struct { F0 int64; F1 string }
type GenS14 struct { F0 uint32; F1 uint8 }
func genFn0() uint64 {
    v4 := GenS1{F0: 0, F1: 275, F2: 128}
    var v5 int64 = -331
    v5 = (v5 ^ -1)
    fmt.Println(v5)
    _ = v4
    _ = v5
    return 643
}
func genFn1(p0 int8, p1 int) uint {
    fmt.Println(p1, p0)
    v6 := []int{(p1 << ((255) & 63)), (p1 ^ (p1 & (p1 * p1))), p1, (p1 / (((p1 - (p1 - (p1 ^ p1)))) | 1))}
    p0 = (p0 ^ (p0 % ((-1) | 1)))
    if (0.1 < (float64(1.5) / (1.5 * 1.5))) {
        v7 := GenS1{F0: -1, F1: 600, F2: 0}
        if true {
            for fi8 := 0; fi8 < 4; fi8++ {
                v9 := GenS1{F0: 527, F1: 1, F2: 59}
                v10 := GenS2{F0: true}
                _ = fi8
                _ = v10
                _ = v9
            }
        } else {
            p0 += 1
            var v11 int8 = p0
            _ = v11
        }
        p0 = (p0 >> ((2) & 63))
        _ = v7
    } else {
        p0 -= p0
        var v12 int8 = p0
        v12 -= v12
        _ = v12
    }
    if true {
        if false {
            v13 := GenS1{F0: -465, F1: 726, F2: 42}
            _ = v13
        } else {
            p0 = p0
        }
    } else {
        if len(v6) > 0 {
            v6[0] = (p1 * p1)
        }
        switch (p1 & (p1 >> ((415) & 63))) {
        case 32768:
            p0 = 127
        case 882:
            p0 += -17
        }
    }
    _ = p0
    _ = p1
    _ = v6
    return 306
}

func main() {
    var v1 uint8 = 127
    var v2 int = 1
    var v3 uint8 = (v1 & (v1 & v1))
    v14 := []int{127, (-(v2)), 245, (v2 ^ (v2 / ((32768) | 1)))}
    var v15 uint16 = 255
    if ((v15 * (v15 * v15)) >= v15) {
        _ = len(v14)
        v3 -= (v1 + (v1 * v1))
        switch "a" {
        case "a":
            _ = len(v14)
        case "hello":
            switch 0 {
            case 4294967295:
                if false {
                    _ = len(v14)
                } else {
                    v1 -= (v3 >> ((65535) & 63))
                    v3 -= v1
                }
            }
        }
    } else {
        var v16 int16 = 1
        if (int64(v15) < int64(v16)) {
            v17 := GenS3{F0: (v15 - v15), F1: -1}
            if len(v14) >= 2 {
                v14 = v14[1:len(v14)]
            }
            for fi18 := 0; fi18 < 12; fi18++ {
                if len(v14) >= 2 {
                    v14 = v14[1:len(v14)]
                }
                if len(v14) > 0 {
                    v14[0] = fi18
                }
                var v19 int8 = 42
                _ = fi18
                _ = v19
            }
            _ = v17
        }
        _ = v16
    }
    v2 = (v2 - v2)
    var v20 int64 = 42
    v20 = v20
    v3 += v3
    if true {
        if ((1.5 == (1.5 + (1.5 - float64(3.14159)))) && true) {
            if (1 <= 4294967295) {
                var v21 uint16 = 255
                _ = v21
            }
            v22 := GenS1{F0: 1000, F1: (v15 * (v15 * v15)), F2: 44}
            _ = v22
        } else {
            for fi23 := 0; fi23 < 20; fi23++ {
                fmt.Println(v20)
                if len(v14) >= 2 {
                    v14 = v14[1:len(v14)]
                }
                fmt.Println(v2, v20, v3)
                _ = fi23
            }
            fmt.Println(v15)
            if (true || (1 != v20)) {
                switch "a" {
                case "\n\t":
                    fmt.Println(v2, v15, v1)
                }
            } else {
                v14 = append(v14, v2)
                for fi24 := 0; fi24 < 14; fi24++ {
                    v25 := GenS3{F0: (v15 / ((v15) | 1)), F1: (v20 / ((9223372036854775807) | 1))}
                    _ = fi24
                    _ = v25
                }
            }
        }
        if false {
            if true {
                v26 := GenS4{F0: (-(v20)), F1: (!true)}
                v27 := GenS3{F0: v15, F1: v20}
                if (!true) {
                    fmt.Println(v2, v20, v1)
                    switch (v2 + -9223372036854775808) {
                    case 0:
                        for fi28 := 0; fi28 < 11; fi28++ {
                            for fi29 := 0; fi29 < 7; fi29++ {
                                v30 := GenS5{F0: 2, F1: -2147483648}
                                v31 := GenS6{F0: 412}
                                var v32 int64 = v20
                                _ = fi29
                                _ = v30
                                _ = v31
                                _ = v32
                            }
                            for fi33 := 0; fi33 < 20; fi33++ {
                                v15 = 32767
                                v1 = (v3 + v1)
                                _ = fi33
                            }
                            v34 := GenS2{F0: ((v20 << ((4294967295) & 63)) >= v20)}
                            _ = fi28
                            _ = v34
                        }
                    default:
                        for fi35 := 0; fi35 < 11; fi35++ {
                            var v36 float32 = (0.1 * -2.25)
                            if (v1 >= v1) {
                                v1 = v1
                            } else {
                                for fi37 := 0; fi37 < 13; fi37++ {
                                    if len(v14) >= 2 {
                                        v14 = v14[1:len(v14)]
                                    }
                                    _ = fi37
                                }
                                for fi38 := 0; fi38 < 2; fi38++ {
                                    v39 := GenS1{F0: 42, F1: v15, F2: v1}
                                    var v40 int64 = 42
                                    _ = fi38
                                    _ = v39
                                    _ = v40
                                }
                            }
                            for fi41 := 0; fi41 < 19; fi41++ {
                                v3 += v1
                                var v42 uint8 = uint8(v15)
                                v20 += (v20 - (^(v20)))
                                _ = fi41
                                _ = v42
                            }
                            _ = fi35
                            _ = v36
                        }
                    }
                } else {
                    for fi43 := 0; fi43 < 6; fi43++ {
                        v44 := GenS7{F0: 2147483647}
                        _ = fi43
                        _ = v44
                    }
                    if false {
                        v15 -= (v15 >> ((42) & 63))
                    }
                    if len(v14) > 0 {
                        v14[0] = (v2 & (-(v2)))
                    }
                }
                _ = v26
                _ = v27
            } else {
                _ = len(v14)
                v2 += (v2 & (-(v2)))
                if (false || true) {
                    for fi45 := 0; fi45 < 9; fi45++ {
                        v46 := GenS5{F0: genFn1(-128, 127), F1: -212}
                        var v47 uint8 = 0
                        _ = fi45
                        _ = v46
                        _ = v47
                    }
                } else {
                    for fi48 := 0; fi48 < 5; fi48++ {
                        if false {
                            v49 := GenS1{F0: -303, F1: v15, F2: (v1 / ((v3) | 1))}
                            if (2147483648 >= 42) {
                                var v50 int = v2
                                var v51 float64 = 1e10
                                _ = v50
                                _ = v51
                            } else {
                                for fi52 := 0; fi52 < 5; fi52++ {
                                    if true {
                                        var v53 int32 = 1
                                        _ = v53
                                    }
                                    switch "foo bar" {
                                    case "\n\t":
                                        if (!(!false)) {
                                            if true {
                                                if (((v15 + (^(v15))) < v15) && (0 > -1)) {
                                                    if ((1.5 * (1.5 - (((3.14159 + 0.1) - 1e10) + ((0.1 + (-2.25 + 1e10)) + float64(1.5))))) != 0.1) {
                                                        v54 := GenS8{F0: 0, F1: 18446744073709551615}
                                                        _ = v54
                                                    } else {
                                                        v55 := GenS5{F0: 0, F1: 0}
                                                        _ = v55
                                                    }
                                                } else {
                                                    var v56 bool = (!(((^(v20)) == (^(v20))) || false))
                                                    var v57 float32 = 1.5
                                                    v3 -= v1
                                                    _ = v56
                                                    _ = v57
                                                }
                                                var v58 int64 = -1
                                                if false {
                                                    var v59 uint8 = (^(v3))
                                                    _ = v59
                                                } else {
                                                    fmt.Println(v20, v15)
                                                }
                                                _ = v58
                                            }
                                            var v60 uint8 = v1
                                            if false {
                                                switch "hello" {
                                                case "a":
                                                    if ((float64(3.14159) < 0.1) && ((0.1 == (-2.25 + (0.0 + 3.14159))) || (-2.25 != ((1e10 * -2.25) - ((0.0 - (1e10 - 0.1)) - 0.0))))) {
                                                        v15 = 42
                                                        var v61 uint16 = v15
                                                        var v62 bool = true
                                                        _ = v61
                                                        _ = v62
                                                    } else {
                                                        switch "hello" {
                                                        case "foo bar":
                                                            switch (("" + ("\n\t" + "foo bar")) + ("a" + "foo bar")) {
                                                            case "foo bar":
                                                                if (!false) {
                                                                    var v63 bool = false
                                                                    v14 = append(v14, int(v15))
                                                                    switch (v2 << ((698) & 63)) {
                                                                    case 42:
                                                                        v64 := GenS9{F0: 18446744073709551615, F1: true, F2: v63}
                                                                        _ = v64
                                                                    case 9223372036854775807:
                                                                        var v65 int64 = v20
                                                                        _ = v65
                                                                    case -2147483648:
                                                                        switch "\n\t" {
                                                                        case "":
                                                                            switch fi52 {
                                                                            case 891:
                                                                                fmt.Println(v20)
                                                                                if len(v14) > 0 {
                                                                                    v14[0] = 9223372036854775807
                                                                                }
                                                                            case -2147483648:
                                                                                var v66 uint32 = 2147483648
                                                                                _ = v66
                                                                            }
                                                                        case "foo bar":
                                                                            var v67 uint16 = (v15 % ((v15) | 1))
                                                                            _ = v67
                                                                        }
                                                                    default:
                                                                        fmt.Println(v3, v2, v15)
                                                                    }
                                                                    v15 -= 635
                                                                    var v68 bool = v63
                                                                    _ = v63
                                                                    _ = v68
                                                                }
                                                            }
                                                        case "hello":
                                                            v2 -= fi52
                                                        default:
                                                            _ = len(v14)
                                                        }
                                                        v1 = (v3 * uint8(v15))
                                                        fmt.Println(v3, v1)
                                                    }
                                                case "":
                                                    if len(v14) >= 2 {
                                                        v14 = v14[1:len(v14)]
                                                    }
                                                case "hello":
                                                    if true {
                                                        switch "\n\t" {
                                                        case "\n\t":
                                                            v14 = append(v14, (fi52 ^ v2))
                                                        default:
                                                            switch fi48 {
                                                            case 1:
                                                                v69 := GenS10{F0: 42}
                                                                _ = v69
                                                            case 9223372036854775807:
                                                                v15 += v15
                                                            }
                                                        }
                                                        if true {
                                                            v3 = (v3 << ((genFn1(127, (v2 + -825))) & 63))
                                                            v70 := GenS11{F0: 127, F1: (fi48 << ((18446744073709551615) & 63))}
                                                            fmt.Println(v15)
                                                            var v71 float64 = 1.5
                                                            _ = v70
                                                            _ = v71
                                                        } else {
                                                            if (("" + "foo bar") > ("\n\t" + "hello")) {
                                                                var v72 uint16 = (^(v15))
                                                                v15 += (^(v72))
                                                                var v73 uint32 = 1
                                                                var v74 bool = true
                                                                switch (^(fi52)) {
                                                                case -9223372036854775808:
                                                                    v14 = append(v14, fi52)
                                                                case -897:
                                                                    if (v20 > v20) {
                                                                        if (v74 || true) {
                                                                            if (-2.25 != 0.0) {
                                                                                var v75 uint64 = uint64(v72)
                                                                                if (-883 <= -1) {
                                                                                    v14 = append(v14, v2)
                                                                                    switch ("hello" + "a") {
                                                                                    case "a":
                                                                                        if len(v14) > 0 {
                                                                                            v14[0] = v2
                                                                                        }
                                                                                    case "":
                                                                                        v73 += v73
                                                                                        var v76 uint16 = v15
                                                                                        if v74 {
                                                                                            var v77 uint = 0
                                                                                            _ = v77
                                                                                        }
                                                                                        _ = v76
                                                                                    case "\n\t":
                                                                                        if len(v14) > 0 {
                                                                                            v14[0] = (fi48 + (fi48 >> ((1) & 63)))
                                                                                        }
                                                                                    }
                                                                                    fmt.Println(v72)
                                                                                }
                                                                                _ = v75
                                                                            }
                                                                            if len(v14) > 0 {
                                                                                v14[0] = (-(fi52))
                                                                            }
                                                                            fmt.Println(v72, v74)
                                                                        }
                                                                    } else {
                                                                        if (fi52 > fi52) {
                                                                            if len(v14) > 0 {
                                                                                v14[0] = (^(v2))
                                                                            }
                                                                        }
                                                                    }
                                                                default:
                                                                    var v78 uint = uint(v73)
                                                                    _ = v78
                                                                }
                                                                switch (^(fi52)) {
                                                                case -9223372036854775808:
                                                                    v1 += 248
                                                                case 42:
                                                                    var v79 int16 = 1000
                                                                    _ = v79
                                                                }
                                                                _ = v72
                                                                _ = v73
                                                                _ = v74
                                                            }
                                                        }
                                                        fmt.Println(v20)
                                                    } else {
                                                        v1 = 0
                                                    }
                                                default:
                                                    var v80 float32 = 0.1
                                                    var v81 uint16 = v15
                                                    _ = v80
                                                    _ = v81
                                                }
                                                v82 := GenS12{F0: 919, F1: -2, F2: v20}
                                                _ = v82
                                            } else {
                                                var v83 int32 = -2147483648
                                                v2 = -9223372036854775808
                                                v83 += v83
                                                _ = v83
                                            }
                                            _ = v60
                                        }
                                    case "hello":
                                        v1 = uint8(v15)
                                    case "foo bar":
                                        v2 += 422
                                    }
                                    v14 = append(v14, v2)
                                    _ = fi52
                                }
                            }
                            if len(v14) > 0 {
                                v14[0] = -360
                            }
                            _ = v49
                        } else {
                            switch fi48 {
                            case 1:
                                var v84 int = (v2 ^ (fi48 << ((18446744073709551615) & 63)))
                                _ = v84
                            case 32768:
                                for fi85 := 0; fi85 < 2; fi85++ {
                                    v20 += (^(v20))
                                    switch fi48 {
                                    case 32767:
                                        var v86 float32 = 3.14159
                                        _ = v86
                                    case -970:
                                        if len(v14) > 0 {
                                            v14[0] = 127
                                        }
                                    case 4294967295:
                                        fmt.Println(v20, v15)
                                    }
                                    var v87 uint32 = 4294967295
                                    _ = fi85
                                    _ = v87
                                }
                            }
                            v14 = append(v14, 134)
                            var v88 string = ""
                            _ = v88
                        }
                        switch (("" + "a") + ("foo bar" + "hello")) {
                        case "\n\t":
                            var v89 uint32 = 500
                            _ = v89
                        case "foo bar":
                            var v90 float32 = 1e10
                            _ = v90
                        case "":
                            if len(v14) > 0 {
                                v14[0] = 899
                            }
                        default:
                            var v91 int64 = v20
                            _ = v91
                        }
                        v14 = append(v14, fi48)
                        _ = fi48
                    }
                    var v92 float64 = 1.5
                    var v93 uint8 = (v3 << ((255) & 63))
                    _ = v92
                    _ = v93
                }
            }
            var v94 int8 = 84
            _ = v94
        } else {
            if true {
                var v95 float64 = 1e10
                _ = len(v14)
                _ = v95
            } else {
                v96 := GenS13{F0: int64(v15), F1: ("" + "")}
                if true {
                    if len(v14) > 0 {
                        v14[0] = v2
                    }
                    var v97 uint = 13
                    if false {
                        v98 := GenS5{F0: (^(v97)), F1: 1}
                        _ = v98
                    } else {
                        if ((!true) || (0 == 1000)) {
                            if (genFn0() != genFn0()) {
                                v99 := GenS4{F0: 9223372036854775807, F1: (v97 >= (v97 ^ (v97 | (v97 + 65535))))}
                                _ = v99
                            } else {
                                var v100 int8 = -1
                                switch ("\n\t" + (("a" + "a") + "\n\t")) {
                                case "":
                                    if (259 > v15) {
                                        for fi101 := 0; fi101 < 13; fi101++ {
                                            v3 = v3
                                            var v102 float64 = (float64(-2.25) * -2.25)
                                            _ = fi101
                                            _ = v102
                                        }
                                        v14 = append(v14, (v2 ^ (v2 | -9223372036854775808)))
                                        for fi103 := 0; fi103 < 19; fi103++ {
                                            for fi104 := 0; fi104 < 1; fi104++ {
                                                v105 := GenS11{F0: v100, F1: fi104}
                                                _ = fi104
                                                _ = v105
                                            }
                                            _ = fi103
                                        }
                                    } else {
                                        if (127 < (v2 / (((v2 << (((v97 | 0)) & 63))) | 1))) {
                                            _ = len(v14)
                                        }
                                    }
                                case "a":
                                    v1 = (^(v1))
                                default:
                                    var v106 int16 = 32767
                                    _ = v106
                                }
                                switch v2 {
                                case 42:
                                    var v107 uint16 = (v15 + 95)
                                    _ = v107
                                case -760:
                                    var v108 string = "a"
                                    _ = v108
                                }
                                _ = v100
                            }
                        }
                        fmt.Println(v1, v2)
                    }
                    _ = v97
                } else {
                    v20 = -1
                    if ((^(v2)) != v2) {
                        _ = len(v14)
                        if true {
                            v109 := GenS14{F0: 65535, F1: (v1 ^ (v1 & v1))}
                            v3 = v1
                            _ = v109
                        }
                        v3 -= v3
                    } else {
                        v1 = (v3 | 1)
                    }
                    for fi110 := 0; fi110 < 6; fi110++ {
                        v111 := GenS4{F0: 620, F1: false}
                        _ = fi110
                        _ = v111
                    }
                }
                v1 = v3
                _ = v96
            }
            v14 = append(v14, (v2 >> ((42) & 63)))
            v112 := GenS10{F0: 1}
            _ = v112
        }
    }
    var v113 uint16 = v15
    fmt.Println(v2, v20)
    _ = len(v14)
    fmt.Println("v1:", v1)
    fmt.Println("v113:", v113)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v113
    _ = v14
    _ = v15
    _ = v2
    _ = v20
    _ = v3
}
