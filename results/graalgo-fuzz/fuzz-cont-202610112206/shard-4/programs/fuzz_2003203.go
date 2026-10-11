package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 uint32; F1 string; F2 uint8 }
type GenS3 struct { F0 uint64; F1 uint32; F2 string }
type GenS4 struct { F0 uint64; F1 int8; F2 string }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint64 }
type GenS8 struct { F0 uint16; F1 int; F2 int8 }
type GenS9 struct { F0 uint16; F1 uint64 }
type GenS10 struct { F0 uint8; F1 uint8 }

func main() {
    var v1 int = -1
    var v2 bool = false
    var v3 uint32 = 498
    v4 := make(map[string]bool)
    v5 := make(map[uint32]int)
    var v6 int32 = -297
    var v7 uint64 = 4294967295
    var v8 uint8 = 0
    var v9 int32 = int32(v8)
    var v10 string = "hello"
    switch "" {
    case "\n\t":
        v11 := GenS1{F0: v6, F1: 100}
        _ = v11
    case "":
        v12 := v5[0]
        _ = v12
    case "a":
        var v13 int16 = 42
        _ = v13
    }
    fmt.Println(v9)
    for fi14 := 0; fi14 < 11; fi14++ {
        switch (fi14 << ((2) & 63)) {
        case 122:
            switch fi14 {
            case 526:
                for fi15 := 0; fi15 < 8; fi15++ {
                    var v16 uint8 = (v8 ^ 128)
                    v17 := GenS2{F0: (v3 + (^(v3))), F1: v10, F2: (v16 | v8)}
                    var v18 string = v10
                    _ = fi15
                    _ = v16
                    _ = v17
                    _ = v18
                }
            }
        case 793:
            v19 := GenS2{F0: (^(v3)), F1: "foo bar", F2: v8}
            _ = v19
        default:
            if v2 {
                for fi20 := 0; fi20 < 11; fi20++ {
                    var v21 bool = v2
                    v22 := []int{}
                    if v21 {
                        var v23 float32 = 0.1
                        v4[v10] = v2
                        _ = v23
                    }
                    v24 := GenS3{F0: (v7 & 4294967295), F1: uint32(v7), F2: ""}
                    var v25 int32 = v6
                    var v26 int16 = -32768
                    v8 = v8
                    v27 := GenS4{F0: v7, F1: 42, F2: ((v10 + "\n\t") + (v10 + v10))}
                    _ = fi20
                    _ = v21
                    _ = v22
                    _ = v24
                    _ = v25
                    _ = v26
                    _ = v27
                }
                v10 = v10
            }
        }
        switch "" {
        case "foo bar":
            var v28 uint64 = v7
            _ = v28
        }
        _ = len(v4)
        _ = fi14
    }
    for fi29 := 0; fi29 < 17; fi29++ {
        v30 := []int{v1, v1, v1, -762}
        v31 := GenS2{F0: v3, F1: "foo bar", F2: (^(v8))}
        _ = fi29
        _ = v30
        _ = v31
    }
    for fi32 := 0; fi32 < 3; fi32++ {
        var v33 string = (v10 + "\n\t")
        switch "a" {
        case "a":
            if true {
                if true {
                    for fi34 := 0; fi34 < 9; fi34++ {
                        delete(v4, v33)
                        v35 := v4[""]
                        if ((uint(v3) < 42) && v2) {
                            var v36 float64 = 1.5
                            var v37 int32 = v9
                            _ = v36
                            _ = v37
                        }
                        _ = fi34
                        _ = v35
                    }
                    var v38 int32 = (^(v6))
                    _ = v38
                }
                fmt.Println(v8)
            } else {
                var v39 int16 = 42
                for fi40 := 0; fi40 < 15; fi40++ {
                    v7 = (^(v7))
                    var v41 uint8 = (^(v8))
                    if v2 {
                        switch (fi40 | v1) {
                        case 42:
                            v42 := GenS2{F0: uint32(v41), F1: "", F2: (v8 | (v8 & 152))}
                            _ = v42
                        case 144:
                            var v43 int32 = v9
                            _ = v43
                        case 4294967295:
                            if true {
                                if (!true) {
                                    v41 = 255
                                    v10 = ("a" + v33)
                                    fmt.Println(v10, v2)
                                } else {
                                    v1 -= -2147483648
                                }
                                v8 += v8
                            }
                        }
                        fmt.Println(v8)
                        v44 := GenS4{F0: v7, F1: -119, F2: "a"}
                        switch (v1 % (((fi40 % (((fi40 % ((fi40) | 1))) | 1))) | 1)) {
                        case 4294967295:
                            if false {
                                switch (fi32 + int(v41)) {
                                case 4294967295:
                                    fmt.Println(v41, v1)
                                case -1:
                                    v5[(v3 | v3)] = (fi40 | (fi40 - (fi32 >> ((2) & 63))))
                                default:
                                    v39 = 0
                                }
                                fmt.Println(fi32, fi40)
                            }
                        default:
                            v6 -= v9
                        }
                        switch v10 {
                        case "":
                            if ((48 != 0) || true) {
                                v7 = v7
                                var v45 float64 = -2.25
                                _ = v45
                            } else {
                                fmt.Println(v10, v7, v41)
                                v6 += (v9 % (((^(v6))) | 1))
                            }
                        case "foo bar":
                            v6 = (-(v6))
                            if true {
                                v46 := GenS5{F0: v6}
                                v3 = 0
                                v39 += v39
                                v10 = ""
                                var v47 uint32 = 591
                                _ = v46
                                _ = v47
                            }
                        default:
                            if (3.14159 > -2.25) {
                                switch fi40 {
                                case 4294967295:
                                    v41 -= uint8(v3)
                                case 32767:
                                    var v48 uint64 = v7
                                    _ = v48
                                case 9223372036854775807:
                                    if false {
                                        fmt.Println(v9, v3)
                                        var v49 uint8 = v41
                                        _ = v49
                                    } else {
                                        v7 = (v7 | (^(v7)))
                                    }
                                    v50 := []int{}
                                    var v51 float64 = (((0.0 / float64(-2.25)) - 1.5) * 3.14159)
                                    _ = v50
                                    _ = v51
                                }
                            } else {
                                delete(v4, v33)
                            }
                        }
                        if true {
                            var v52 string = (v33 + (v10 + ("\n\t" + ("" + (v10 + v33)))))
                            fmt.Println(v10, v6, v2)
                            fmt.Println(v8, v10, v1)
                            _ = v52
                        } else {
                            v3 = v3
                        }
                        if (42 >= 60) {
                            if true {
                                v53 := GenS6{F0: 127}
                                v39 += (^(v39))
                                switch 9223372036854775807 {
                                case -4:
                                    v54 := GenS5{F0: (v6 - v6)}
                                    _ = v54
                                case 519:
                                    fmt.Println(v3, v9)
                                }
                                _ = v53
                            }
                        }
                        v2 = true
                        _ = v44
                    }
                    v39 = -32768
                    switch (v1 << ((2) & 63)) {
                    case 9223372036854775807:
                        var v55 uint8 = (v41 >> ((2) & 63))
                        _ = v55
                    case -1:
                        v3 -= 0
                    }
                    switch fi40 {
                    case 1:
                        if v2 {
                            switch "hello" {
                            case "foo bar":
                                var v56 int8 = 101
                                _ = v56
                            case "hello":
                                if v2 {
                                    v10 = v33
                                    if (uint8(v7) >= (v41 + (v8 * v41))) {
                                        fmt.Println(fi40, fi40)
                                        v57 := []int{2147483647}
                                        switch v1 {
                                        case 32768:
                                            switch "a" {
                                            case "\n\t":
                                                if ((-125 != 74) || ((0 == 180) && (true && false))) {
                                                    var v58 int64 = 0
                                                    _ = v58
                                                } else {
                                                    if ((false && v2) && v2) {
                                                        var v59 bool = v2
                                                        v39 = (v39 * v39)
                                                        _ = v59
                                                    } else {
                                                        v60 := GenS1{F0: (v9 % ((v6) | 1)), F1: 42}
                                                        v41 = v41
                                                        _ = v60
                                                    }
                                                    var v61 bool = (!((!false) && true))
                                                    switch 0 {
                                                    case 42:
                                                        v62 := GenS5{F0: v6}
                                                        _ = v62
                                                    }
                                                    _ = v61
                                                }
                                            case "foo bar":
                                                fmt.Println(v2, v7)
                                            }
                                        case 42:
                                            if v2 {
                                                v63 := v5[4294967295]
                                                switch v10 {
                                                case "":
                                                    var v64 int32 = int32(v7)
                                                    _ = v64
                                                case "\n\t":
                                                    if (!(128 >= v63)) {
                                                        if (!(v2 && (v10 >= "foo bar"))) {
                                                            v65 := GenS6{F0: 100}
                                                            var v66 string = "\n\t"
                                                            if v2 {
                                                                if len(v57) > 0 {
                                                                    v57[0] = (fi32 % (((fi32 | (fi40 >> ((261) & 63)))) | 1))
                                                                }
                                                                var v67 uint64 = 1
                                                                var v68 uint8 = v41
                                                                _ = v67
                                                                _ = v68
                                                            }
                                                            v10 = "\n\t"
                                                            _ = v65
                                                            _ = v66
                                                        } else {
                                                            switch v63 {
                                                            case 4294967295:
                                                                v7 = (v7 << ((2) & 63))
                                                            }
                                                            switch v63 {
                                                            case -9223372036854775808:
                                                                v9 = v6
                                                            case 127:
                                                                v8 = (^(v8))
                                                            case -178:
                                                                v1 = int(v8)
                                                            }
                                                            if (!v2) {
                                                                fmt.Println(fi40, v8, v33)
                                                            }
                                                        }
                                                        var v69 uint16 = 65535
                                                        _ = v69
                                                    }
                                                default:
                                                    if (v7 == (v7 >> ((42) & 63))) {
                                                        fmt.Println(v33, v63, v3)
                                                    }
                                                }
                                                _ = v63
                                            }
                                        }
                                        _ = v57
                                    }
                                    switch 9223372036854775807 {
                                    case 336:
                                        var v70 int = v1
                                        if ((float64(0.1) + 0.0) != 3.14159) {
                                            fmt.Println(v10, v33, v2)
                                            v71 := GenS2{F0: 2147483647, F1: v33, F2: (^(v8))}
                                            switch "" {
                                            case "foo bar":
                                                v72 := GenS7{F0: 0}
                                                _ = v72
                                            case "hello":
                                                v6 += 42
                                            case "":
                                                if (42 != 53) {
                                                    fmt.Println(v70, v8)
                                                    v73 := GenS2{F0: 293, F1: (v10 + v33), F2: v8}
                                                    if (0.0 <= (0.1 - ((-2.25 - (float64(-2.25) - (1.5 - (1e10 * -2.25)))) / 0.1))) {
                                                        v39 = v39
                                                        var v74 float64 = 3.14159
                                                        if (v39 != v39) {
                                                            var v75 uint8 = 127
                                                            _ = v75
                                                        } else {
                                                            switch (fi40 & fi40) {
                                                            case -9223372036854775808:
                                                                if true {
                                                                    switch "a" {
                                                                    case "":
                                                                        if (0.1 < float32((0.0 - float64(3.14159)))) {
                                                                            v70 = int(v7)
                                                                            _ = len(v5)
                                                                        } else {
                                                                            v76 := GenS3{F0: (v7 >> ((0) & 63)), F1: (v3 - v3), F2: (v33 + v10)}
                                                                            fmt.Println(v7, v6)
                                                                            v10 = (((v10 + ((("hello" + v10) + (v10 + v33)) + v10)) + "foo bar") + v10)
                                                                            v77 := []int{2147483647, 2147483647, (^(v1)), -1}
                                                                            _ = v76
                                                                            _ = v77
                                                                        }
                                                                    case "foo bar":
                                                                        var v78 int = -287
                                                                        _ = v78
                                                                    }
                                                                } else {
                                                                    var v79 uint16 = 127
                                                                    _ = v79
                                                                }
                                                                fmt.Println(v3)
                                                            case 9223372036854775807:
                                                                switch v10 {
                                                                case "":
                                                                    var v80 uint64 = (v7 - (^(v7)))
                                                                    _ = v80
                                                                default:
                                                                    v81 := v5[v3]
                                                                    _ = v81
                                                                }
                                                            default:
                                                                v82 := GenS2{F0: (v3 * (v3 * 2)), F1: "\n\t", F2: v41}
                                                                _ = v82
                                                            }
                                                            v70 = (v1 | fi32)
                                                            var v83 string = ((((((v10 + v10) + (v10 + v10)) + "hello") + ("foo bar" + "hello")) + "foo bar") + "hello")
                                                            _ = v83
                                                        }
                                                        _ = v74
                                                    } else {
                                                        fmt.Println(v9)
                                                        v39 -= (v39 + (v39 >> ((18446744073709551615) & 63)))
                                                        _ = len(v5)
                                                        v84 := GenS8{F0: 0, F1: (v1 | v70), F2: -128}
                                                        _ = v84
                                                    }
                                                    _ = v73
                                                } else {
                                                    var v85 int = -632
                                                    v86 := GenS9{F0: 934, F1: v7}
                                                    var v87 bool = ((1 != 42) || false)
                                                    if ((-(v85)) > (-(v70))) {
                                                        var v88 uint64 = (v7 & (v7 / (((v7 + (^(v7)))) | 1)))
                                                        v89 := []int{}
                                                        _ = v88
                                                        _ = v89
                                                    } else {
                                                        v6 += (v6 - (v9 ^ (v9 % (((v6 & (v6 % ((2147483647) | 1)))) | 1))))
                                                    }
                                                    _ = v85
                                                    _ = v86
                                                    _ = v87
                                                }
                                                if (1e10 <= 3.14159) {
                                                    var v90 float32 = 1.5
                                                    var v91 uint16 = 1
                                                    _ = v90
                                                    _ = v91
                                                }
                                            default:
                                                delete(v4, "hello")
                                            }
                                            _ = v71
                                        } else {
                                            v10 = (((("" + v10) + v33) + "") + v33)
                                            if (v41 == v8) {
                                                v92 := []int{}
                                                _ = len(v4)
                                                _ = len(v4)
                                                fmt.Println(v39, v1, v39)
                                                fmt.Println(v2, v3, v3)
                                                v92 = append(v92, fi40)
                                                v2 = v2
                                                _ = v92
                                            } else {
                                                fmt.Println(v9, v7)
                                                var v93 uint32 = (v3 | 959)
                                                _ = v93
                                            }
                                        }
                                        _ = v70
                                    default:
                                        v3 = v3
                                    }
                                } else {
                                    var v94 string = "a"
                                    _ = len(v4)
                                    _ = v94
                                }
                            }
                            v95 := GenS5{F0: v6}
                            v6 -= -1
                            var v96 int64 = 0
                            var v97 uint8 = v41
                            _ = v95
                            _ = v96
                            _ = v97
                        }
                    default:
                        var v98 float32 = float32(((1e10 / float64(-2.25)) * float64((1.5 + -2.25))))
                        _ = v98
                    }
                    _ = fi40
                    _ = v41
                }
                switch v10 {
                case "":
                    v99 := GenS1{F0: v9, F1: 42}
                    _ = v99
                case "\n\t":
                    var v100 float64 = 1e10
                    _ = v100
                default:
                    var v101 int32 = v9
                    _ = v101
                }
                _ = v39
            }
        case "foo bar":
            v102 := GenS1{F0: -273, F1: 15}
            _ = v102
        }
        var v103 int16 = 32767
        _ = fi32
        _ = v103
        _ = v33
    }
    var v104 string = ((v10 + v10) + v10)
    v105 := GenS10{F0: v8, F1: (v8 % ((v8) | 1))}
    _ = len(v4)
    var v106 int32 = v9
    switch ("" + (v10 + "hello")) {
    case "hello":
        v9 = (v6 / (((v106 ^ -2147483648)) | 1))
    case "":
        var v107 float64 = 1.5
        _ = v107
    case "foo bar":
        for fi108 := 0; fi108 < 18; fi108++ {
            var v109 int16 = 42
            var v110 uint32 = v3
            v9 += (v6 ^ (v9 * (v6 % ((0) | 1))))
            _ = fi108
            _ = v109
            _ = v110
        }
    }
    v10 = ("foo bar" + v104)
    v106 = (v6 ^ int32(v1))
    if (!true) {
        if (v9 == 592) {
            fmt.Println(v6)
            if ((-(v1)) == v1) {
                v111 := GenS5{F0: v106}
                switch int(v7) {
                case -1:
                    v8 = v8
                case 32767:
                    v4["hello"] = (v2 && v2)
                }
                v1 = 42
                _ = v111
            }
        }
    } else {
        switch v1 {
        case 888:
            v104 = v104
        default:
            v112 := v4[v104]
            _ = v112
        }
        var v113 uint64 = (v7 % ((4294967295) | 1))
        _ = v113
    }
    _ = len(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v104:", v104)
    fmt.Println("v106:", v106)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v104
    _ = v105
    _ = v106
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
