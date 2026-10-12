package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 int }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 uint8; F1 int }
type GenS4 struct { F0 uint64; F1 bool }
type GenS5 struct { F0 bool; F1 uint16 }
type GenS6 struct { F0 uint64; F1 uint8; F2 uint64 }
type GenS7 struct { F0 uint8; F1 int8; F2 uint32 }
type GenS8 struct { F0 int }
type GenS9 struct { F0 uint32; F1 uint8 }
type GenS10 struct { F0 string; F1 uint64; F2 uint }
type GenS11 struct { F0 int; F1 int; F2 string }
type GenS12 struct { F0 int16; F1 int32 }
type GenS13 struct { F0 uint }
func genFn0(p0 uint16, p1 bool) uint {
    v7 := GenS1{F0: 18446744073709551615, F1: 675, F2: 32767}
    for fi8 := 0; fi8 < 7; fi8++ {
        var v9 string = ("a" + "\n\t")
        v9 = ""
        _ = fi8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v7
    return 2
}
func genFn1() int16 {
    for fi10 := 0; fi10 < 1; fi10++ {
        switch "" {
        case "":
            var v11 int = (fi10 * fi10)
            _ = v11
        default:
            var v12 int32 = 1
            _ = v12
        }
        var v13 int = fi10
        v14 := []int{}
        _ = fi10
        _ = v13
        _ = v14
    }
    if false {
        var v15 int32 = 1
        if ((132 >= 1) || false) {
            fmt.Println(v15)
        } else {
            var v16 uint8 = 1
            for fi17 := 0; fi17 < 7; fi17++ {
                for fi18 := 0; fi18 < 2; fi18++ {
                    if (-772 == 9223372036854775807) {
                        v19 := []int{(fi17 << ((4294967295) & 63)), -2147483648}
                        var v20 uint8 = (v16 % ((v16) | 1))
                        v21 := GenS2{F0: 2, F1: v20}
                        _ = v19
                        _ = v20
                        _ = v21
                    } else {
                        var v22 int8 = 39
                        v23 := GenS3{F0: 128, F1: (^(fi17))}
                        if true {
                            var v24 float32 = (0.0 / -2.25)
                            v15 -= v15
                            v25 := GenS1{F0: 261, F1: 18446744073709551615, F2: fi17}
                            _ = v24
                            _ = v25
                        } else {
                            switch (fi18 + (^(fi18))) {
                            case 32768:
                                var v26 bool = (true || (-979 > 541))
                                _ = v26
                            case 4294967295:
                                v27 := []int{}
                                _ = v27
                            }
                            v16 = (v16 << ((42) & 63))
                            switch "hello" {
                            case "\n\t":
                                v16 -= (v16 << ((genFn0(2, false)) & 63))
                            case "foo bar":
                                if true {
                                    v28 := []int{(fi17 * 1)}
                                    if len(v28) > 0 {
                                        v28[0] = (fi17 ^ fi18)
                                    }
                                    _ = v28
                                } else {
                                    if (0.1 == -2.25) {
                                        v29 := GenS3{F0: v16, F1: 128}
                                        _ = v29
                                    } else {
                                        fmt.Println(fi17)
                                    }
                                    v22 = (v22 & 110)
                                    var v30 uint32 = 42
                                    _ = v30
                                }
                            }
                        }
                        _ = v22
                        _ = v23
                    }
                    v31 := GenS1{F0: 2, F1: 0, F2: 496}
                    _ = fi18
                    _ = v31
                }
                _ = fi17
            }
            _ = v16
        }
        _ = v15
    } else {
        var v32 string = "\n\t"
        v33 := GenS4{F0: 2, F1: true}
        v34 := GenS5{F0: true, F1: 42}
        _ = v32
        _ = v33
        _ = v34
    }
    var v35 uint8 = 0
    _ = v35
    return -32768
}

func main() {
    v1 := []uint32{703, 497, 0}
    var v2 uint8 = 1
    var v3 float32 = 0.1
    var v4 int64 = 42
    v5 := []string{}
    var v6 int8 = -110
    var v36 uint32 = 65535
    if false {
        switch -9223372036854775808 {
        case -1:
            for fi37 := 0; fi37 < 13; fi37++ {
                if true {
                    for fi38 := 0; fi38 < 6; fi38++ {
                        if false {
                            var v39 int = -2147483648
                            v6 += (v6 + (^(v6)))
                            if false {
                                v36 += v36
                                var v40 int64 = 9223372036854775807
                                _ = v40
                            } else {
                                switch (("hello" + (("\n\t" + "\n\t") + (("a" + "hello") + ("\n\t" + "hello")))) + "") {
                                case "a":
                                    _ = len(v1)
                                case "":
                                    v36 += uint32(v2)
                                    switch -200 {
                                    case 0:
                                        v41 := GenS1{F0: 65535, F1: 42, F2: v39}
                                        _ = v41
                                    case -2147483648:
                                        if (!true) {
                                            var v42 uint8 = 0
                                            _ = v42
                                        } else {
                                            var v43 uint8 = v2
                                            _ = len(v1)
                                            if len(v1) > 0 {
                                                v1[0] = v36
                                            }
                                            _ = v43
                                        }
                                    }
                                case "\n\t":
                                    switch ("a" + ("hello" + "\n\t")) {
                                    case "a":
                                        var v44 uint8 = v2
                                        _ = v44
                                    }
                                }
                            }
                            _ = v39
                        } else {
                            if true {
                                if ((v4 >= (^(v4))) && false) {
                                    if false {
                                        if (true && true) {
                                            var v45 uint16 = 255
                                            switch (fi38 & 0) {
                                            case -9223372036854775808:
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            case 426:
                                                v3 = 1.5
                                            case 0:
                                                var v46 int8 = 57
                                                _ = v46
                                            }
                                            if false {
                                                v6 = v6
                                                v36 = 1
                                            }
                                            v6 -= v6
                                            _ = v45
                                        }
                                        if len(v1) > 0 {
                                            v1[0] = v36
                                        }
                                    }
                                    switch "foo bar" {
                                    case "":
                                        var v47 bool = false
                                        _ = v47
                                    default:
                                        fmt.Println(v4)
                                    }
                                }
                                v48 := GenS6{F0: 18446744073709551615, F1: v2, F2: 165}
                                _ = v48
                            }
                            v6 += v6
                        }
                        switch (("foo bar" + "\n\t") + "hello") {
                        case "hello":
                            fmt.Println(v6, fi37, fi37)
                        case "foo bar":
                            if (0 < 20) {
                                v2 = v2
                                v49 := GenS7{F0: 255, F1: v6, F2: (v36 & uint32(v2))}
                                _ = v49
                            }
                        case "a":
                            v5 = append(v5, "\n\t")
                        default:
                            if true {
                                fmt.Println(v2, v4, fi38)
                                switch "foo bar" {
                                case "hello":
                                    var v50 float64 = (0.1 / (1e10 * 1e10))
                                    _ = v50
                                case "\n\t":
                                    if (!true) {
                                        fmt.Println(v2)
                                    }
                                case "foo bar":
                                    switch ("foo bar" + ("a" + "")) {
                                    case "a":
                                        v36 = (v36 & v36)
                                    }
                                }
                            }
                        }
                        if false {
                            if len(v1) > 0 {
                                v1[0] = 2147483648
                            }
                        } else {
                            var v51 float64 = 1.5
                            fmt.Println(fi37)
                            _ = v51
                        }
                        _ = fi38
                    }
                    for fi52 := 0; fi52 < 7; fi52++ {
                        fmt.Println(fi52, v36)
                        v53 := GenS8{F0: (fi37 & (fi37 | fi37))}
                        v3 += v3
                        _ = fi52
                        _ = v53
                    }
                    var v54 int = (fi37 / (((fi37 % ((fi37) | 1))) | 1))
                    _ = v54
                } else {
                    var v55 uint32 = (v36 ^ v36)
                    v56 := GenS1{F0: 18446744073709551615, F1: 42, F2: (fi37 ^ (fi37 % ((4294967295) | 1)))}
                    var v57 int8 = -128
                    _ = v55
                    _ = v56
                    _ = v57
                }
                for fi58 := 0; fi58 < 5; fi58++ {
                    var v59 int8 = (v6 & v6)
                    fmt.Println(v2, fi37, fi37)
                    _ = fi58
                    _ = v59
                }
                var v60 int = 1
                _ = fi37
                _ = v60
            }
        case -500:
            v4 = 42
        case -445:
            var v61 uint32 = v36
            _ = v61
        }
    }
    switch -620 {
    case 1:
        fmt.Println(v4, v6)
    case 9223372036854775807:
        if false {
            v2 = 2
        } else {
            for fi62 := 0; fi62 < 6; fi62++ {
                if true {
                    switch "hello" {
                    case "\n\t":
                        var v63 float32 = (v3 * v3)
                        _ = v63
                    case "":
                        if (-2.25 > (0.0 / 3.14159)) {
                            var v64 int16 = 32767
                            _ = v64
                        } else {
                            var v65 int8 = v6
                            v66 := GenS5{F0: true, F1: 446}
                            _ = v65
                            _ = v66
                        }
                    }
                    if false {
                        if false {
                            switch 1 {
                            case -9223372036854775808:
                                v36 = (v36 % ((v36) | 1))
                            case 42:
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                            case 0:
                                v36 += 1
                            default:
                                var v67 float64 = -2.25
                                _ = v67
                            }
                        }
                    } else {
                        v4 = 192
                        v68 := GenS9{F0: (^(v36)), F1: v2}
                        _ = v68
                    }
                    v69 := GenS1{F0: genFn0(65535, true), F1: 1, F2: (-(fi62))}
                    _ = v69
                }
                v70 := GenS10{F0: "", F1: 42, F2: 255}
                var v71 int32 = -504
                _ = fi62
                _ = v70
                _ = v71
            }
            v3 += ((float32(0.1) / 0.1) / (0.0 + (-2.25 + v3)))
            if true {
                v72 := GenS11{F0: -1, F1: 4294967295, F2: "a"}
                if false {
                    for fi73 := 0; fi73 < 8; fi73++ {
                        _ = len(v1)
                        var v74 string = "\n\t"
                        _ = fi73
                        _ = v74
                    }
                    var v75 int64 = v4
                    if (!false) {
                        if true {
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            v76 := GenS7{F0: v2, F1: (v6 - 100), F2: 42}
                            _ = v76
                        } else {
                            v36 -= v36
                            fmt.Println(v2, v3, v4)
                        }
                        var v77 uint16 = 255
                        _ = v77
                    }
                    _ = v75
                } else {
                    for fi78 := 0; fi78 < 14; fi78++ {
                        fmt.Println(v6)
                        if ((!true) && (false || true)) {
                            v6 = 85
                            var v79 bool = (855 <= 40)
                            _ = v79
                        }
                        v80 := GenS3{F0: (v2 & 240), F1: fi78}
                        _ = fi78
                        _ = v80
                    }
                    v4 -= int64(v2)
                }
                _ = v72
            }
        }
    case 42:
        var v81 uint16 = 0
        _ = v81
    }
    v4 += (v4 & v4)
    var v82 uint = 4294967295
    v1 = append(v1, (v36 ^ v36))
    v83 := GenS5{F0: true, F1: 0}
    for fi84 := 0; fi84 < 10; fi84++ {
        var v85 float32 = 0.1
        _ = fi84
        _ = v85
    }
    v86 := GenS12{F0: 1, F1: -2147483648}
    if false {
        switch "foo bar" {
        case "foo bar":
            if len(v1) > 0 {
                v1[0] = v36
            }
        case "a":
            if true {
                var v87 float32 = 0.1
                var v88 string = "a"
                fmt.Println(v88, v87, v3)
                _ = v87
                _ = v88
            } else {
                v82 += v82
            }
        default:
            for fi89 := 0; fi89 < 15; fi89++ {
                var v90 uint64 = 4294967295
                var v91 uint8 = (v2 >> (((v82 | v82)) & 63))
                if false {
                    for fi92 := 0; fi92 < 5; fi92++ {
                        v93 := GenS13{F0: uint(v90)}
                        _ = fi92
                        _ = v93
                    }
                }
                _ = fi89
                _ = v90
                _ = v91
            }
        }
        var v94 int16 = 780
        _ = v94
    } else {
        switch 0 {
        case -9223372036854775808:
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        case -1:
            var v95 float64 = (0.0 + 0.1)
            _ = v95
        case 511:
            for fi96 := 0; fi96 < 11; fi96++ {
                switch -9223372036854775808 {
                case -1:
                    _ = len(v1)
                case 76:
                    switch "" {
                    case "":
                        v97 := GenS5{F0: (v6 >= v6), F1: 0}
                        _ = v97
                    case "foo bar":
                        var v98 string = "a"
                        _ = v98
                    default:
                        if (1 < 0) {
                            var v99 int8 = int8(v4)
                            fmt.Println(v6, v6, v36)
                            v82 = v82
                            _ = v99
                        }
                    }
                default:
                    for fi100 := 0; fi100 < 4; fi100++ {
                        if true {
                            var v101 uint8 = (v2 / ((128) | 1))
                            fmt.Println(v36, v4)
                            var v102 int16 = -537
                            _ = v101
                            _ = v102
                        }
                        _ = fi100
                    }
                }
                if len(v5) > 0 {
                    v5[0] = "foo bar"
                }
                _ = fi96
            }
        default:
            fmt.Println(v36)
        }
        fmt.Println(v6, v6)
    }
    var v103 uint64 = 794
    v5 = append(v5, (((("a" + "hello") + ("foo bar" + (((("hello" + ("a" + "foo bar")) + "foo bar") + ("a" + (("hello" + "foo bar") + "\n\t"))) + "\n\t"))) + "") + ("" + ("\n\t" + ""))))
    if (18446744073709551615 >= 4294967295) {
        var v104 uint8 = v2
        for fi105 := 0; fi105 < 10; fi105++ {
            for fi106 := 0; fi106 < 1; fi106++ {
                v4 -= (v4 ^ int64(v104))
                var v107 bool = false
                if (v4 <= (v4 ^ (v4 | v4))) {
                    var v108 int8 = (v6 >> ((v82) & 63))
                    _ = v108
                }
                _ = fi106
                _ = v107
            }
            fmt.Println(v2, v4, v82)
            _ = fi105
        }
        _ = v104
    }
    v36 -= v36
    var v109 uint64 = 755
    v110 := GenS7{F0: (v2 + (v2 % (((v2 << ((v82) & 63))) | 1))), F1: (v6 / ((v6) | 1)), F2: (v36 >> ((uint(v109)) & 63))}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v103:", v103)
    fmt.Println("v109:", v109)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v82:", v82)
    _ = v1
    _ = v103
    _ = v109
    _ = v110
    _ = v2
    _ = v3
    _ = v36
    _ = v4
    _ = v5
    _ = v6
    _ = v82
    _ = v83
    _ = v86
}
