package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int8 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int64; F1 uint64; F2 uint64 }
type GenS5 struct { F0 int; F1 string; F2 uint }
type GenS6 struct { F0 int16; F1 uint; F2 uint32 }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 string; F1 uint64 }
type GenS9 struct { F0 int64; F1 int64 }
func genFn0(p0 uint8) bool {
    var v7 int8 = 100
    if (0 < 590) {
        switch ("hello" + ("a" + "hello")) {
        case "":
            fmt.Println(v7)
        }
        fmt.Println(v7)
    }
    v8 := []int{-2147483648, 4294967295}
    var v9 float64 = ((0.0 - 1.5) - 1e10)
    _ = p0
    _ = v7
    _ = v8
    _ = v9
    return false
}
func genFn1() uint16 {
    v10 := []int{2147483647}
    for fi11 := 0; fi11 < 7; fi11++ {
        var v12 uint16 = 42
        _ = fi11
        _ = v12
    }
    _ = v10
    return 42
}

func main() {
    v1 := make(map[int]uint32)
    var v2 uint64 = 615
    v3 := make(map[int]uint32)
    var v4 uint8 = 128
    var v5 int64 = -9223372036854775808
    var v6 float32 = 0.1
    var v13 int16 = 1
    for fi14 := 0; fi14 < 17; fi14++ {
        for fi15 := 0; fi15 < 12; fi15++ {
            if ((^(v4)) >= (v4 << ((816) & 63))) {
                if (!(-66 == -56)) {
                    if (false && (127 < -1)) {
                        fmt.Println(v13, fi14)
                        var v16 uint16 = 55
                        _ = v16
                    }
                    fmt.Println(fi14)
                }
            } else {
                var v17 int8 = -51
                switch ("\n\t" + ("foo bar" + ("hello" + ((("" + "") + ("" + "a")) + "hello")))) {
                case "a":
                    v18 := v1[(fi15 * (fi15 & fi14))]
                    v6 -= float32(3.14159)
                    var v19 uint8 = 42
                    fmt.Println(v4)
                    if true {
                        var v20 int16 = (v13 | 32767)
                        _ = v20
                    } else {
                        v21 := GenS1{F0: (^(v5)), F1: 1}
                        _ = v21
                    }
                    _ = v18
                    _ = v19
                case "foo bar":
                    v4 -= 170
                default:
                    v4 += v4
                }
                _ = v17
            }
            fmt.Println(v5, v4, v13)
            v13 = -1
            if false {
                if (0 > 343) {
                    v3[(fi15 >> ((65535) & 63))] = 2147483647
                    v6 = v6
                } else {
                    var v22 int32 = 0
                    _ = v22
                }
            }
            _ = fi15
        }
        if true {
            v23 := GenS2{F0: 812}
            for fi24 := 0; fi24 < 15; fi24++ {
                if true {
                    v5 += v5
                } else {
                    var v25 uint32 = 692
                    v13 += (-(v13))
                    v26 := v1[fi24]
                    _ = v25
                    _ = v26
                }
                v27 := []int{fi24, 127, fi14, -1}
                if (!(v6 >= float32(0.1))) {
                    fmt.Println(v4)
                    if ((fi14 * 981) >= 128) {
                        var v28 int64 = (v5 + -520)
                        v4 += v4
                        v3[fi24] = 788
                        _ = v28
                    }
                    var v29 uint16 = 2
                    _ = v29
                } else {
                    var v30 uint16 = genFn1()
                    v3[fi24] = 0
                    _ = v30
                }
                _ = fi24
                _ = v27
            }
            _ = v23
        } else {
            switch (("\n\t" + (("a" + "hello") + (("foo bar" + "foo bar") + "a"))) + "a") {
            case "":
                var v31 int = (fi14 % (((fi14 - 32767)) | 1))
                _ = v31
            default:
                v2 += 42
            }
            if true {
                var v32 uint32 = 2
                var v33 float32 = v6
                _ = v32
                _ = v33
            }
        }
        _ = fi14
    }
    for fi34 := 0; fi34 < 5; fi34++ {
        var v35 int32 = 1
        for fi36 := 0; fi36 < 14; fi36++ {
            switch (fi36 / (((fi36 / (((fi36 & fi36)) | 1))) | 1)) {
            case -2147483648:
                fmt.Println(fi34, v4)
            }
            _ = fi36
        }
        _ = fi34
        _ = v35
    }
    v6 -= v6
    for fi37 := 0; fi37 < 20; fi37++ {
        v6 = v6
        switch (fi37 % ((int(v2)) | 1)) {
        case -133:
            for fi38 := 0; fi38 < 15; fi38++ {
                if (int8(v4) > 1) {
                    v39 := GenS3{F0: uint(v4)}
                    fmt.Println(v13)
                    switch 9223372036854775807 {
                    case 1:
                        v40 := []int{-869, fi38}
                        _ = v40
                    case -9223372036854775808:
                        var v41 float64 = 1.5
                        _ = v41
                    case -454:
                        if ((^(v2)) == 2) {
                            var v42 int64 = v5
                            v5 += (v5 - (v42 * v42))
                            switch (fi38 << ((818) & 63)) {
                            case 42:
                                v43 := GenS1{F0: 42, F1: 0}
                                _ = v43
                            }
                            switch "foo bar" {
                            case "foo bar":
                                fmt.Println(v6, v42)
                            case "":
                                if ((float32((0.0 * (0.0 / float64(v6)))) + v6) > (v6 + ((v6 + 0.0) * (v6 - (float32(((-2.25 - (0.0 - 1.5)) * (0.1 - (1.5 + -2.25)))) - 0.1))))) {
                                    switch int(v13) {
                                    case -291:
                                        var v44 string = "a"
                                        v45 := GenS1{F0: (v42 - (v42 | -955)), F1: -64}
                                        switch (^(fi38)) {
                                        case 1:
                                            var v46 uint8 = (v4 % ((v4) | 1))
                                            _ = v46
                                        case 932:
                                            v5 -= (v5 & int64(v2))
                                        case 2147483647:
                                            v4 -= v4
                                        }
                                        _ = v44
                                        _ = v45
                                    default:
                                        v2 -= (v2 % ((v2) | 1))
                                    }
                                    var v47 uint16 = 0
                                    v6 -= (v6 - v6)
                                    _ = v47
                                } else {
                                    v4 = 161
                                    if true {
                                        fmt.Println(fi37, v13, fi38)
                                        v13 -= (v13 % ((v13) | 1))
                                        v5 -= 837
                                    } else {
                                        v48 := []int{fi37, fi37, (fi38 ^ (^(fi38))), fi38}
                                        fmt.Println(v6, fi37, v2)
                                        _ = v48
                                    }
                                }
                            }
                            fmt.Println(v42)
                            _ = v42
                        } else {
                            if true {
                                v49 := GenS4{F0: (v5 % (((^(v5))) | 1)), F1: (^(v2)), F2: v2}
                                _ = v49
                            } else {
                                if false {
                                    if (0 <= (^(v2))) {
                                        v50 := []int{(fi37 >> ((4294967295) & 63)), (fi38 ^ -15)}
                                        if true {
                                            fmt.Println(v13, v4, fi38)
                                            v5 -= (v5 << ((65535) & 63))
                                            fmt.Println(v6)
                                            fmt.Println(v13)
                                        } else {
                                            switch "foo bar" {
                                            case "":
                                                v51 := GenS5{F0: fi37, F1: "\n\t", F2: 4294967295}
                                                _ = v51
                                            case "\n\t":
                                                v13 = (v13 + 42)
                                            default:
                                                v1[fi38] = 1
                                                if (fi38 <= (fi37 - (fi37 & fi38))) {
                                                    v13 = (v13 & v13)
                                                }
                                            }
                                            v4 -= (^(v4))
                                        }
                                        v5 -= int64(v4)
                                        v4 = v4
                                        fmt.Println(v2, v5, fi37)
                                        _ = v50
                                    } else {
                                        var v52 bool = true
                                        fmt.Println(v5, v2)
                                        v53 := GenS4{F0: (^(v5)), F1: v2, F2: v2}
                                        _ = v52
                                        _ = v53
                                    }
                                    var v54 bool = false
                                    var v55 uint8 = 127
                                    _ = v54
                                    _ = v55
                                }
                                v56 := GenS6{F0: (^(v13)), F1: 4294967295, F2: 2}
                                v57 := GenS7{F0: 916}
                                v58 := GenS6{F0: (v13 * -919), F1: 956, F2: 820}
                                _ = v56
                                _ = v57
                                _ = v58
                            }
                        }
                    }
                    var v59 bool = (!(!true))
                    if false {
                        var v60 float64 = 0.1
                        var v61 uint32 = 4294967295
                        _ = v60
                        _ = v61
                    } else {
                        switch fi37 {
                        case 470:
                            var v62 uint64 = v2
                            _ = v62
                        case 4294967295:
                            var v63 float32 = 0.0
                            if v59 {
                                v64 := GenS3{F0: 128}
                                _ = v64
                            } else {
                                var v65 int16 = 32767
                                var v66 uint = 965
                                var v67 uint8 = (v4 + v4)
                                var v68 uint = 1
                                _ = v65
                                _ = v66
                                _ = v67
                                _ = v68
                            }
                            _ = v63
                        case -2147483648:
                            fmt.Println(v59)
                        }
                        if (!(65535 >= 616)) {
                            v69 := GenS2{F0: -2147483648}
                            var v70 float64 = 1.5
                            _ = v69
                            _ = v70
                        }
                        var v71 int16 = v13
                        fmt.Println(v13, v4, v2)
                        _ = v71
                    }
                    if (0.0 <= 0.1) {
                        v72 := GenS2{F0: 42}
                        _ = v72
                    }
                    _ = v39
                    _ = v59
                } else {
                    var v73 int8 = 42
                    _ = v73
                }
                switch (("hello" + ("\n\t" + "a")) + "") {
                case "\n\t":
                    switch "" {
                    case "a":
                        var v74 int64 = v5
                        _ = v74
                    case "hello":
                        v2 -= (v2 * v2)
                    default:
                        v75 := GenS5{F0: 127, F1: ("foo bar" + "foo bar"), F2: 216}
                        v2 = (v2 | (^(v2)))
                        v76 := v1[fi38]
                        switch (("hello" + "\n\t") + "a") {
                        case "hello":
                            v1[fi38] = v76
                        case "a":
                            if false {
                                var v77 uint64 = (v2 ^ v2)
                                switch "a" {
                                case "a":
                                    var v78 uint64 = (v2 * v2)
                                    _ = v78
                                }
                                _ = v77
                            } else {
                                var v79 float32 = v6
                                v13 -= int16(v2)
                                fmt.Println(v13, v4, fi38)
                                _ = v79
                            }
                        default:
                            v80 := []int{32768, (fi38 & fi38)}
                            _ = v80
                        }
                        _ = v75
                        _ = v76
                    }
                default:
                    v81 := []int{694, (fi37 + (^(fi38)))}
                    v82 := GenS5{F0: 9223372036854775807, F1: "\n\t", F2: 0}
                    fmt.Println(v4, v4, v2)
                    v2 = 1
                    _ = v81
                    _ = v82
                }
                _ = fi38
            }
        default:
            _ = len(v1)
        }
        _ = fi37
    }
    switch 127 {
    case 8:
        v6 -= v6
    case -2147483648:
        var v83 int16 = 801
        _ = v83
    default:
        for fi84 := 0; fi84 < 15; fi84++ {
            fmt.Println(v13, v2, v5)
            if (false || true) {
                v13 += (v13 - 0)
            }
            v2 = (^(v2))
            _ = fi84
        }
    }
    for fi85 := 0; fi85 < 14; fi85++ {
        v5 -= 42
        v5 = v5
        _ = fi85
    }
    if (65535 >= 18446744073709551615) {
        for fi86 := 0; fi86 < 6; fi86++ {
            fmt.Println(v2, fi86, v2)
            _ = fi86
        }
        if (-2.25 > ((v6 - float32(float64(((v6 / -2.25) - ((0.0 * -2.25) / ((v6 / float32(((-2.25 * (((1e10 + 1e10) + (0.1 + (0.0 - (1.5 - -2.25)))) * 1.5)) / -2.25))) / v6)))))) + 0.1)) {
            var v87 uint32 = 0
            var v88 int16 = (v13 / (((v13 << ((0) & 63))) | 1))
            switch 59 {
            case 1:
                for fi89 := 0; fi89 < 3; fi89++ {
                    fmt.Println(v88)
                    _ = fi89
                }
            default:
                v87 = v87
            }
            _ = v87
            _ = v88
        } else {
            if true {
                v90 := v3[-223]
                v91 := []int{128, 9223372036854775807, 2147483647, -25}
                switch -120 {
                case -9223372036854775808:
                    for fi92 := 0; fi92 < 14; fi92++ {
                        for fi93 := 0; fi93 < 9; fi93++ {
                            switch 128 {
                            case 128:
                                var v94 float64 = -2.25
                                _ = v94
                            case 32767:
                                fmt.Println(v6, v13, v2)
                            case 42:
                                v13 -= 0
                            }
                            v13 -= (v13 / ((v13) | 1))
                            v90 -= (v90 - 373)
                            _ = fi93
                        }
                        for fi95 := 0; fi95 < 10; fi95++ {
                            var v96 int16 = int16(v2)
                            _ = fi95
                            _ = v96
                        }
                        _ = fi92
                    }
                case 32767:
                    var v97 uint8 = (^(v4))
                    _ = v97
                }
                _ = v90
                _ = v91
            }
            v98 := []int{}
            v99 := GenS3{F0: 367}
            _ = v98
            _ = v99
        }
        var v100 int16 = (-(v13))
        _ = v100
    }
    v101 := v1[-2147483648]
    var v102 uint8 = (v4 ^ (v4 % (((^(v4))) | 1)))
    v103 := GenS5{F0: 223, F1: ("\n\t" + ""), F2: 0}
    for fi104 := 0; fi104 < 6; fi104++ {
        v105 := []int{0, fi104}
        var v106 uint16 = 630
        _ = fi104
        _ = v105
        _ = v106
    }
    v102 = 171
    v102 -= v4
    if ((v4 / (((v102 + v4)) | 1)) >= v102) {
        v107 := []int{2147483647, -2147483648, -195, 283}
        _ = v107
    } else {
        if (!((0.0 == 0.1) || (!((v5 << ((972) & 63)) == (v5 - 9223372036854775807))))) {
            if true {
                switch -30 {
                case -616:
                    v108 := GenS8{F0: ("hello" + "hello"), F1: (^(v2))}
                    _ = v108
                case 42:
                    var v109 int8 = -124
                    _ = v109
                }
                var v110 float64 = 3.14159
                _ = v110
            }
        }
        fmt.Println(v102, v5)
        v4 = (^(v102))
    }
    v111 := GenS9{F0: (-(v5)), F1: 0}
    v101 = (v101 * v101)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v101:", v101)
    fmt.Println("v102:", v102)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v101
    _ = v102
    _ = v103
    _ = v111
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
