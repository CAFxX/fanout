package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint16; F1 string }
type GenS4 struct { F0 int32; F1 int16 }
type GenS5 struct { F0 uint64; F1 uint32; F2 uint32 }
type GenS6 struct { F0 bool; F1 uint; F2 uint8 }
type GenS7 struct { F0 bool; F1 int8 }
type GenS8 struct { F0 string; F1 uint8 }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 int8; F1 uint16; F2 string }
type GenS11 struct { F0 string; F1 uint32; F2 int8 }
type GenS12 struct { F0 bool; F1 uint64 }
type GenS13 struct { F0 uint; F1 int32 }
type GenS14 struct { F0 uint16 }
type GenS15 struct { F0 uint8 }
type GenS16 struct { F0 string; F1 int8 }
type GenS17 struct { F0 bool; F1 int; F2 int32 }
func genFn0(p0 uint8, p1 int32, p2 uint64) uint32 {
    var v6 uint16 = 65535
    v7 := GenS1{F0: 2}
    fmt.Println(v6, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 65535
}
func genFn1() uint32 {
    switch 32768 {
    case -567:
        var v8 uint32 = 271
        _ = v8
    default:
        var v9 int8 = 42
        _ = v9
    }
    for fi10 := 0; fi10 < 7; fi10++ {
        var v11 int8 = 100
        v12 := GenS2{F0: 1000}
        _ = fi10
        _ = v11
        _ = v12
    }
    var v13 bool = (!true)
    _ = v13
    return 42
}

func main() {
    var v1 bool = true
    var v2 bool = (!(!true))
    v3 := make(map[uint32]int)
    var v4 int = 594
    v5 := []uint32{0, 2147483648}
    if (1e10 == (1e10 / 0.1)) {
        for fi14 := 0; fi14 < 12; fi14++ {
            v2 = (2147483648 != 1)
            _ = fi14
        }
    }
    switch "" {
    case "\n\t":
        _ = len(v5)
    case "foo bar":
        for fi15 := 0; fi15 < 19; fi15++ {
            v1 = false
            switch v4 {
            case -532:
                if ((("\n\t" + "a") < "a") || v1) {
                    var v16 uint8 = 69
                    v16 = (v16 + (^(v16)))
                    for fi17 := 0; fi17 < 11; fi17++ {
                        var v18 int = (^(fi17))
                        var v19 uint32 = 0
                        _ = fi17
                        _ = v18
                        _ = v19
                    }
                    _ = v16
                }
            case 32767:
                if v2 {
                    var v20 uint16 = 65535
                    var v21 float32 = (1.5 - 3.14159)
                    switch ("foo bar" + "hello") {
                    case "a":
                        v22 := GenS3{F0: v20, F1: (("" + "foo bar") + "\n\t")}
                        _ = v22
                    case "":
                        if (!false) {
                            v2 = true
                            var v23 uint32 = uint32(v20)
                            v24 := GenS3{F0: (v20 >> ((475) & 63)), F1: "a"}
                            _ = v23
                            _ = v24
                        } else {
                            fmt.Println(v20)
                        }
                    case "hello":
                        var v25 int = fi15
                        _ = v25
                    default:
                        v26 := GenS4{F0: 0, F1: 32767}
                        _ = v26
                    }
                    _ = v20
                    _ = v21
                } else {
                    v27 := GenS5{F0: 5, F1: 42, F2: 1}
                    _ = v27
                }
            default:
                v28 := GenS6{F0: v1, F1: 2, F2: 255}
                _ = v28
            }
            if (18446744073709551615 == 2) {
                v29 := v3[2147483647]
                v2 = (3.14159 <= (0.1 + ((-2.25 / -2.25) / 1e10)))
                _ = v29
            }
            _ = fi15
        }
    default:
        for fi30 := 0; fi30 < 2; fi30++ {
            _ = len(v3)
            for fi31 := 0; fi31 < 10; fi31++ {
                if v1 {
                    if false {
                        if (1 <= 9223372036854775807) {
                            if (v1 && v2) {
                                v32 := GenS6{F0: true, F1: 1, F2: 128}
                                if ((((142 == 2) || (false || true)) || v1) || (!v1)) {
                                    var v33 uint = 5
                                    v34 := v3[0]
                                    _ = v33
                                    _ = v34
                                }
                                v35 := GenS1{F0: 414}
                                _ = v32
                                _ = v35
                            }
                            if v2 {
                                var v36 uint64 = 77
                                if v1 {
                                    var v37 float32 = float32(-2.25)
                                    switch (fi30 + 0) {
                                    case 42:
                                        var v38 uint16 = 42
                                        _ = v38
                                    default:
                                        if v1 {
                                            if true {
                                                var v39 float32 = (0.1 / (((v37 / ((v37 / (v37 / ((1e10 * v37) + (float32(3.14159) * v37)))) - v37)) - ((3.14159 - (v37 / (1e10 + (0.0 - (float32(-2.25) * v37))))) / v37)) / 0.1))
                                                fmt.Println(v39, v4)
                                                v39 -= ((1e10 - 1.5) / v37)
                                                _ = v39
                                            } else {
                                                if ((-128 == -52) && (false && (42 <= -32768))) {
                                                    _ = len(v5)
                                                    fmt.Println(v4)
                                                    v5 = append(v5, 133)
                                                    if (genFn0(162, -2147483648, v36) < 65535) {
                                                        var v40 string = "foo bar"
                                                        var v41 int16 = 1000
                                                        var v42 uint16 = 182
                                                        _ = v40
                                                        _ = v41
                                                        _ = v42
                                                    }
                                                }
                                            }
                                            v2 = true
                                            var v43 float32 = v37
                                            _ = v43
                                        }
                                    }
                                    if v2 {
                                        switch (-(fi30)) {
                                        case -1:
                                            v4 = (^(fi30))
                                        case 142:
                                            switch "\n\t" {
                                            case "\n\t":
                                                delete(v3, genFn0(128, -2147483648, (v36 | v36)))
                                            case "a":
                                                v2 = v1
                                            case "hello":
                                                v44 := v3[65535]
                                                _ = v44
                                            }
                                        default:
                                            v45 := GenS7{F0: false, F1: 127}
                                            _ = v45
                                        }
                                        var v46 int8 = -1
                                        _ = v46
                                    } else {
                                        if v2 {
                                            var v47 float32 = v37
                                            if v1 {
                                                switch "hello" {
                                                case "\n\t":
                                                    v2 = (0.1 > 3.14159)
                                                }
                                                v48 := GenS8{F0: "\n\t", F1: 42}
                                                _ = v48
                                            }
                                            _ = v47
                                        }
                                        fmt.Println(v1, fi31)
                                    }
                                    _ = v37
                                } else {
                                    v49 := GenS4{F0: -1, F1: -1}
                                    _ = v49
                                }
                                _ = v36
                            } else {
                                var v50 float32 = -2.25
                                _ = v50
                            }
                            v51 := GenS8{F0: (("foo bar" + "") + ("" + (("" + "foo bar") + "\n\t"))), F1: 42}
                            _ = v51
                        } else {
                            v52 := GenS9{F0: (false || false)}
                            _ = v52
                        }
                        var v53 uint8 = 1
                        if (-2147483648 <= 42) {
                            if v1 {
                                if (!(true && (932 < 4294967295))) {
                                    v53 = 96
                                }
                                v5 = append(v5, 4294967295)
                            } else {
                                if v1 {
                                    v1 = v1
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                    if ((127 >= 255) || true) {
                                        v1 = false
                                        if (false && (!v2)) {
                                            delete(v3, genFn0(128, 334, 2))
                                        } else {
                                            var v54 int8 = -3
                                            _ = v54
                                        }
                                        var v55 int8 = int8(v4)
                                        _ = v55
                                    }
                                } else {
                                    if len(v5) > 0 {
                                        v5[0] = genFn0((v53 | (v53 - 255)), int32(v53), 1)
                                    }
                                    v56 := GenS10{F0: 127, F1: 2, F2: ""}
                                    var v57 int16 = 437
                                    _ = v56
                                    _ = v57
                                }
                                var v58 uint16 = 182
                                _ = v58
                            }
                            switch ("" + "hello") {
                            case "":
                                if (true || (1 <= 2)) {
                                    var v59 uint8 = (v53 / ((v53) | 1))
                                    fmt.Println(v4, v2)
                                    var v60 int = 554
                                    var v61 int16 = 0
                                    if v1 {
                                        var v62 uint64 = 2
                                        _ = v62
                                    }
                                    v5 = append(v5, 42)
                                    _ = v59
                                    _ = v60
                                    _ = v61
                                }
                            default:
                                switch fi30 {
                                case 32767:
                                    if (18446744073709551615 != 1) {
                                        switch fi30 {
                                        case 1:
                                            _ = len(v5)
                                        default:
                                            var v63 bool = true
                                            _ = v63
                                        }
                                    }
                                case -2147483648:
                                    var v64 uint = 873
                                    _ = v64
                                case -9223372036854775808:
                                    v1 = ((!(32 >= genFn0(v53, 2147483647, 191))) || v2)
                                }
                                v65 := GenS11{F0: "foo bar", F1: 785, F2: 108}
                                _ = len(v5)
                                _ = v65
                            }
                        }
                        _ = v53
                    } else {
                        if (2147483647 == 637) {
                            v2 = (0 > 660)
                            v1 = false
                            var v66 int32 = -2147483648
                            _ = v66
                        } else {
                            switch fi30 {
                            case 127:
                                var v67 int32 = -1
                                _ = v67
                            }
                        }
                        if (42 < 42) {
                            v68 := GenS6{F0: false, F1: 96, F2: 90}
                            if v1 {
                                v2 = v2
                                var v69 uint8 = 102
                                fmt.Println(v4, fi31)
                                _ = v69
                            } else {
                                v2 = v1
                            }
                            var v70 uint64 = 496
                            _ = v68
                            _ = v70
                        }
                    }
                    if v2 {
                        if len(v5) >= 2 {
                            v5 = v5[1:len(v5)]
                        }
                    } else {
                        v71 := GenS4{F0: 2147483647, F1: 42}
                        v72 := GenS12{F0: true, F1: 397}
                        var v73 int16 = 0
                        _ = v71
                        _ = v72
                        _ = v73
                    }
                } else {
                    fmt.Println(v2)
                    _ = len(v5)
                }
                _ = len(v5)
                if (2147483647 > (-(fi30))) {
                    switch "a" {
                    case "":
                        v4 -= 0
                    case "hello":
                        var v74 uint64 = 4294967295
                        _ = v74
                    case "foo bar":
                        v4 = fi30
                    }
                    var v75 int64 = 9223372036854775807
                    _ = v75
                } else {
                    v76 := v3[2]
                    fmt.Println(v1)
                    _ = v76
                }
                _ = fi31
            }
            _ = fi30
        }
    }
    if (100 == 127) {
        var v77 int32 = 2147483647
        var v78 int8 = -128
        for fi79 := 0; fi79 < 12; fi79++ {
            if (9223372036854775807 == 9223372036854775807) {
                v4 = v4
                if (42 > 2) {
                    if (1.5 >= ((0.0 - float64(((-2.25 / (0.0 - 3.14159)) * float32(1e10)))) + 1e10)) {
                        v5 = append(v5, 42)
                    } else {
                        fmt.Println(fi79, fi79, v4)
                        v80 := GenS13{F0: 42, F1: 541}
                        _ = v80
                    }
                }
                var v81 bool = (-95 < v78)
                _ = v81
            } else {
                v82 := GenS9{F0: v1}
                fmt.Println(v77, v77, fi79)
                _ = v82
            }
            var v83 bool = v1
            if len(v5) > 0 {
                v5[0] = 2
            }
            _ = fi79
            _ = v83
        }
        _ = v77
        _ = v78
    }
    for fi84 := 0; fi84 < 2; fi84++ {
        if false {
            if (true && v2) {
                fmt.Println(v4, v4)
                v3[65535] = 0
            } else {
                for fi85 := 0; fi85 < 15; fi85++ {
                    _ = len(v5)
                    _ = fi85
                }
            }
            var v86 int = -655
            _ = v86
        } else {
            fmt.Println(fi84)
            v2 = false
            v1 = ("hello" == ("hello" + (("hello" + "\n\t") + "foo bar")))
        }
        for fi87 := 0; fi87 < 2; fi87++ {
            var v88 uint16 = 65535
            _ = fi87
            _ = v88
        }
        fmt.Println(v4)
        _ = fi84
    }
    fmt.Println(v2)
    if (v4 >= v4) {
        var v89 uint16 = 2
        _ = v89
    }
    switch v4 {
    case 32767:
        var v90 bool = v2
        _ = v90
    case 830:
        v1 = v2
    }
    v2 = v2
    var v91 int64 = -398
    v92 := GenS1{F0: 2}
    if (18446744073709551615 <= 18446744073709551615) {
        v2 = v1
    }
    v1 = v2
    if ((v4 * 1) >= 128) {
        v2 = v2
        var v93 uint8 = 36
        var v94 uint = 65535
        _ = v93
        _ = v94
    } else {
        switch (("" + "foo bar") + "foo bar") {
        case "foo bar":
            fmt.Println(v2, v1, v2)
        case "\n\t":
            for fi95 := 0; fi95 < 19; fi95++ {
                var v96 uint32 = 2
                if false {
                    if v2 {
                        switch fi95 {
                        case -2147483648:
                            var v97 uint16 = 392
                            _ = v97
                        case 9223372036854775807:
                            if (!true) {
                                var v98 int8 = 32
                                _ = v98
                            } else {
                                v4 = (fi95 ^ fi95)
                                for fi99 := 0; fi99 < 11; fi99++ {
                                    v2 = (802 <= 0)
                                    v100 := GenS6{F0: (!(v96 > v96)), F1: 171, F2: 227}
                                    if (!false) {
                                        v2 = (uint(v96) >= 2)
                                        if (!(!true)) {
                                            v101 := GenS14{F0: 2}
                                            v102 := GenS13{F0: 42, F1: 1}
                                            _ = v101
                                            _ = v102
                                        } else {
                                            v2 = false
                                        }
                                    }
                                    _ = fi99
                                    _ = v100
                                }
                            }
                        case 988:
                            for fi103 := 0; fi103 < 4; fi103++ {
                                v91 = (v91 - 42)
                                if len(v5) > 0 {
                                    v5[0] = 857
                                }
                                v104 := GenS15{F0: 128}
                                delete(v3, (v96 % (((v96 + (v96 ^ v96))) | 1)))
                                _ = fi103
                                _ = v104
                            }
                        default:
                            if false {
                                delete(v3, v96)
                                v1 = ((-2.25 / (1e10 + float32(3.14159))) <= (-2.25 - 3.14159))
                            } else {
                                v1 = false
                                fmt.Println(fi95)
                                switch (fi95 / ((fi95) | 1)) {
                                case 32768:
                                    for fi105 := 0; fi105 < 14; fi105++ {
                                        _ = len(v3)
                                        v106 := GenS15{F0: 0}
                                        _ = fi105
                                        _ = v106
                                    }
                                case 1:
                                    var v107 string = "hello"
                                    _ = v107
                                default:
                                    var v108 int32 = 218
                                    _ = v108
                                }
                            }
                        }
                        v109 := GenS16{F0: "\n\t", F1: 42}
                        switch -1 {
                        case 2147483647:
                            v110 := v3[v96]
                            _ = v110
                        case 999:
                            v1 = false
                        default:
                            fmt.Println(v96)
                        }
                        _ = v109
                    }
                } else {
                    var v111 int8 = -128
                    if false {
                        if false {
                            var v112 string = ("foo bar" + "a")
                            switch -154 {
                            case 127:
                                if len(v5) > 0 {
                                    v5[0] = 0
                                }
                            case 4294967295:
                                var v113 string = v112
                                _ = v113
                            case 128:
                                var v114 int16 = -32768
                                _ = v114
                            }
                            _ = v112
                        } else {
                            var v115 int64 = 1
                            var v116 uint64 = 2
                            if v1 {
                                v1 = (((!false) && v1) && (-2.25 < -2.25))
                            } else {
                                switch (v4 ^ 9223372036854775807) {
                                case -9223372036854775808:
                                    switch fi95 {
                                    case 0:
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                    case -735:
                                        v91 -= (v91 - -114)
                                    case 1:
                                        v117 := GenS3{F0: 32767, F1: "foo bar"}
                                        _ = v117
                                    }
                                case 32768:
                                    var v118 int = v4
                                    _ = v118
                                case 2147483647:
                                    v1 = v1
                                }
                                _ = len(v5)
                            }
                            _ = v115
                            _ = v116
                        }
                        var v119 int32 = 42
                        var v120 uint16 = 1
                        _ = v119
                        _ = v120
                    }
                    v91 += (v91 & v91)
                    _ = v111
                }
                _ = fi95
                _ = v96
            }
        default:
            var v121 int = -9223372036854775808
            _ = v121
        }
    }
    var v122 int32 = 42
    var v123 uint = 2
    if v1 {
        v91 += -9223372036854775808
        for fi124 := 0; fi124 < 6; fi124++ {
            fmt.Println(v123, fi124)
            if v2 {
                if v1 {
                    for fi125 := 0; fi125 < 3; fi125++ {
                        v123 = v123
                        v2 = ((v122 & 1) > v122)
                        _ = fi125
                    }
                    if (!true) {
                        var v126 uint64 = 751
                        _ = v126
                    }
                } else {
                    for fi127 := 0; fi127 < 5; fi127++ {
                        v122 = (v122 << ((v123) & 63))
                        if v2 {
                            if false {
                                v91 += (v91 ^ (v91 | (-(v91))))
                                v4 -= 1
                                v4 -= fi127
                            }
                            v91 += 0
                            var v128 uint8 = 128
                            _ = v128
                        }
                        _ = fi127
                    }
                }
            } else {
                fmt.Println(v1, v1)
                v122 += v122
                if (((v122 * (-(v122))) != int32(v91)) || (!v2)) {
                    if v2 {
                        v2 = v1
                        v129 := GenS15{F0: 1}
                        v130 := GenS17{F0: true, F1: v4, F2: -2147483648}
                        _ = v129
                        _ = v130
                    } else {
                        if len(v5) > 0 {
                            v5[0] = uint32(v123)
                        }
                    }
                }
            }
            fmt.Println(v4, v2, v91)
            _ = fi124
        }
        for fi131 := 0; fi131 < 4; fi131++ {
            fmt.Println(v2)
            _ = fi131
        }
    }
    fmt.Println(v123, v2, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v122:", v122)
    fmt.Println("v123:", v123)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v122
    _ = v123
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v91
    _ = v92
}
