package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int; F2 uint64 }
type GenS2 struct { F0 uint64; F1 uint }
type GenS3 struct { F0 uint16 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 uint64 }
type GenS6 struct { F0 bool; F1 uint64 }
type GenS7 struct { F0 int64; F1 int }
type GenS8 struct { F0 int }
type GenS9 struct { F0 bool }
type GenS10 struct { F0 int8; F1 uint64; F2 uint8 }
type GenS11 struct { F0 int16; F1 string; F2 int }
type GenS12 struct { F0 uint16; F1 int16; F2 int64 }
type GenS13 struct { F0 uint64; F1 string }
type GenS14 struct { F0 int }
type GenS15 struct { F0 int64 }
type GenS16 struct { F0 uint8; F1 uint64; F2 string }

func main() {
    v1 := make(map[int]string)
    var v2 int = -874
    v3 := make(map[int]int)
    v4 := make(map[string]uint32)
    var v5 int = 127
    v6 := make(map[string]bool)
    var v7 uint16 = 65535
    v8 := []int{v2, -9223372036854775808, (v2 & (v5 / ((127) | 1)))}
    switch "a" {
    case "hello":
        fmt.Println(v7, v2, v7)
    default:
        v9 := GenS1{F0: 238, F1: v2, F2: 18446744073709551615}
        _ = v9
    }
    for fi10 := 0; fi10 < 11; fi10++ {
        fmt.Println(v5, v7)
        _ = fi10
    }
    switch "a" {
    case "\n\t":
        var v11 string = ""
        _ = v11
    case "":
        if true {
            fmt.Println(v2, v5, v5)
            v5 += v5
            v8 = append(v8, v5)
        }
    case "foo bar":
        if len(v8) >= 2 {
            v8 = v8[1:len(v8)]
        }
    }
    fmt.Println(v7, v5)
    v7 -= v7
    if false {
        var v12 int = (v2 >> ((339) & 63))
        v13 := GenS1{F0: 128, F1: v5, F2: 42}
        _ = v12
        _ = v13
    }
    var v14 int64 = 42
    if true {
        var v15 float32 = ((3.14159 / 1.0) / 1.5)
        v5 += (v2 * (v2 >> ((18446744073709551615) & 63)))
        v2 += v2
        _ = v15
    } else {
        v16 := GenS1{F0: 223, F1: (v5 & (v5 >> ((4294967295) & 63))), F2: 1}
        var v17 int32 = 42
        if (0.0 < 1.5) {
            v2 -= v5
            v17 -= -1
            var v18 uint64 = 760
            _ = v18
        }
        _ = v16
        _ = v17
    }
    for fi19 := 0; fi19 < 2; fi19++ {
        v7 = v7
        if (1.5 == 3.14159) {
            for fi20 := 0; fi20 < 5; fi20++ {
                v21 := GenS1{F0: 255, F1: (fi19 * fi19), F2: 1}
                if true {
                    v5 -= (v2 << ((4294967295) & 63))
                    v5 += -2147483648
                    v22 := v3[(v2 + v2)]
                    var v23 int = fi19
                    _ = v22
                    _ = v23
                }
                if ((false && false) || (!true)) {
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                } else {
                    v14 += (v14 / ((-683) | 1))
                    var v24 uint = 921
                    switch "foo bar" {
                    case "hello":
                        v25 := GenS2{F0: 23, F1: v24}
                        _ = v25
                    case "a":
                        v2 += (v5 & (v2 & 927))
                    case "\n\t":
                        if len(v8) > 0 {
                            v8[0] = v2
                        }
                    }
                    var v26 int64 = 1
                    fmt.Println(v24, v24, v5)
                    _ = v24
                    _ = v26
                }
                if (uint64(v7) <= 42) {
                    switch fi20 {
                    case 9223372036854775807:
                        v27 := GenS3{F0: 32767}
                        _ = v27
                    default:
                        if false {
                            if (0 > 0) {
                                if true {
                                    var v28 uint8 = 166
                                    v29 := GenS4{F0: 1}
                                    var v30 uint16 = (v7 + v7)
                                    switch ("a" + ("hello" + "hello")) {
                                    case "":
                                        v2 -= v2
                                    case "foo bar":
                                        fmt.Println(v7, v30)
                                    }
                                    _ = v28
                                    _ = v29
                                    _ = v30
                                } else {
                                    var v31 uint32 = 1
                                    v7 += v7
                                    switch ("hello" + ("foo bar" + ("\n\t" + "foo bar"))) {
                                    case "\n\t":
                                        if len(v8) > 0 {
                                            v8[0] = (fi20 - fi19)
                                        }
                                    case "":
                                        var v32 int64 = 0
                                        _ = v32
                                    case "a":
                                        switch (fi19 | (v5 - fi20)) {
                                        case 2147483647:
                                            var v33 bool = true
                                            switch (^(fi20)) {
                                            case -2147483648:
                                                v34 := GenS3{F0: (v7 / ((v7) | 1))}
                                                _ = v34
                                            }
                                            if false {
                                                switch 9223372036854775807 {
                                                case -2147483648:
                                                    if (127 == -122) {
                                                        v5 -= -605
                                                        switch -1 {
                                                        case 128:
                                                            fmt.Println(v31)
                                                        case 32768:
                                                            if (v33 && v33) {
                                                                if true {
                                                                    if (!false) {
                                                                        if v33 {
                                                                            fmt.Println(v33, v31)
                                                                            var v35 float64 = 0.0
                                                                            _ = v35
                                                                        } else {
                                                                            v2 += fi20
                                                                        }
                                                                        switch (v2 / ((fi20) | 1)) {
                                                                        case -9223372036854775808:
                                                                            if ((-36 >= -128) || (true && (0 != -532))) {
                                                                                if (1 < 127) {
                                                                                    var v36 int8 = 1
                                                                                    if (fi19 > v5) {
                                                                                        fmt.Println(v33, v36, fi19)
                                                                                    }
                                                                                    if (42 >= 1) {
                                                                                        v5 -= -2147483648
                                                                                    }
                                                                                    if len(v8) >= 2 {
                                                                                        v8 = v8[1:len(v8)]
                                                                                    }
                                                                                    _ = v36
                                                                                }
                                                                                v2 -= 1
                                                                            } else {
                                                                                var v37 uint = 255
                                                                                v7 = v7
                                                                                _ = v37
                                                                            }
                                                                        case -322:
                                                                            v7 = (v7 / (((^(v7))) | 1))
                                                                        case -2147483648:
                                                                            v33 = (v31 == v31)
                                                                            switch v2 {
                                                                            case 0:
                                                                                var v38 uint16 = v7
                                                                                _ = v38
                                                                            case -2147483648:
                                                                                if len(v8) > 0 {
                                                                                    v8[0] = 42
                                                                                }
                                                                            case 2147483647:
                                                                                var v39 int64 = v14
                                                                                _ = v39
                                                                            }
                                                                            if v33 {
                                                                                switch -9223372036854775808 {
                                                                                case 339:
                                                                                    v40 := GenS5{F0: 259}
                                                                                    _ = v40
                                                                                case -369:
                                                                                    var v41 uint32 = (^(v31))
                                                                                    _ = v41
                                                                                }
                                                                            }
                                                                        default:
                                                                            v5 -= 936
                                                                        }
                                                                    } else {
                                                                        if false {
                                                                            _ = len(v8)
                                                                            v5 += fi19
                                                                        } else {
                                                                            var v42 int64 = (v14 & -576)
                                                                            v43 := GenS6{F0: ((206 < 204) && (74 != -128)), F1: 1}
                                                                            v42 = v14
                                                                            _ = v42
                                                                            _ = v43
                                                                        }
                                                                        if ((-1 < -92) && ((1 <= 18446744073709551615) || false)) {
                                                                            var v44 bool = (((v31 + 67) <= (v31 / ((950) | 1))) && (32767 == 0))
                                                                            _ = v44
                                                                        } else {
                                                                            var v45 uint64 = uint64(v7)
                                                                            _ = v45
                                                                        }
                                                                        switch "a" {
                                                                        case "foo bar":
                                                                            if (!v33) {
                                                                                if ((v33 || (!false)) || ((v14 <= v14) && v33)) {
                                                                                    var v46 int64 = v14
                                                                                    _ = v46
                                                                                } else {
                                                                                    v47 := v1[(v2 % ((fi19) | 1))]
                                                                                    if true {
                                                                                        var v48 int64 = 9223372036854775807
                                                                                        if (v48 <= -9223372036854775808) {
                                                                                            if v33 {
                                                                                                v3[v2] = fi20
                                                                                                var v49 float64 = ((((1e10 + 0.0) - 3.14159) / 1e10) * 1.5)
                                                                                                _ = v49
                                                                                            } else {
                                                                                                var v50 uint64 = 2
                                                                                                var v51 string = v47
                                                                                                if true {
                                                                                                    switch (fi20 >> ((1) & 63)) {
                                                                                                    case 2147483647:
                                                                                                        v47 = ((v47 + v51) + "hello")
                                                                                                        var v52 bool = v33
                                                                                                        _ = v52
                                                                                                    case -221:
                                                                                                        fmt.Println(v48)
                                                                                                    }
                                                                                                    var v53 uint = 382
                                                                                                    v53 = v53
                                                                                                    _ = len(v8)
                                                                                                    fmt.Println(fi19, v14, fi20)
                                                                                                    _ = v53
                                                                                                } else {
                                                                                                    fmt.Println(v33, v7, fi19)
                                                                                                    fmt.Println(v48, v48)
                                                                                                }
                                                                                                _ = v50
                                                                                                _ = v51
                                                                                            }
                                                                                            var v54 int16 = -360
                                                                                            v55 := GenS7{F0: v14, F1: (v2 | (v2 >> ((42) & 63)))}
                                                                                            v5 += v2
                                                                                            _ = v54
                                                                                            _ = v55
                                                                                        }
                                                                                        _ = v48
                                                                                    } else {
                                                                                        switch v5 {
                                                                                        case -1:
                                                                                            v8 = append(v8, fi19)
                                                                                        case 0:
                                                                                            var v56 uint8 = 2
                                                                                            _ = v56
                                                                                        default:
                                                                                            fmt.Println(v47, v14, v5)
                                                                                        }
                                                                                    }
                                                                                    v57 := GenS8{F0: fi20}
                                                                                    if v33 {
                                                                                        v31 = uint32(v7)
                                                                                    }
                                                                                    _ = v47
                                                                                    _ = v57
                                                                                }
                                                                                if (!false) {
                                                                                    var v58 uint64 = 4294967295
                                                                                    _ = len(v8)
                                                                                    _ = v58
                                                                                } else {
                                                                                    switch v5 {
                                                                                    case 32767:
                                                                                        v59 := GenS9{F0: (false || (int32(v31) != 0))}
                                                                                        _ = v59
                                                                                    default:
                                                                                        switch fi19 {
                                                                                        case 32768:
                                                                                            v60 := GenS10{F0: 42, F1: 0, F2: uint8(v31)}
                                                                                            _ = v60
                                                                                        }
                                                                                    }
                                                                                    var v61 uint64 = 726
                                                                                    _ = v61
                                                                                }
                                                                                v14 -= (v14 % (((v14 << ((255) & 63))) | 1))
                                                                            } else {
                                                                                v62 := GenS2{F0: 718, F1: 255}
                                                                                var v63 int = (v5 ^ fi20)
                                                                                if (!(-2.25 > (1e10 / 1.5))) {
                                                                                    v64 := GenS11{F0: 1000, F1: ((("hello" + "a") + (("foo bar" + ("hello" + "a")) + "")) + "\n\t"), F2: -934}
                                                                                    v5 += fi19
                                                                                    var v65 uint = 2
                                                                                    _ = v64
                                                                                    _ = v65
                                                                                }
                                                                                _ = v62
                                                                                _ = v63
                                                                            }
                                                                            v66 := GenS12{F0: v7, F1: int16(v31), F2: 42}
                                                                            _ = v66
                                                                        case "hello":
                                                                            switch 32767 {
                                                                            case 32768:
                                                                                _ = len(v8)
                                                                            }
                                                                        case "":
                                                                            _ = len(v8)
                                                                        }
                                                                    }
                                                                    var v67 uint8 = 2
                                                                    var v68 uint = 2
                                                                    _ = v67
                                                                    _ = v68
                                                                } else {
                                                                    var v69 bool = v33
                                                                    _ = v69
                                                                }
                                                                v2 += 765
                                                                if (v14 <= (v14 - v14)) {
                                                                    var v70 uint64 = 83
                                                                    var v71 uint64 = 1
                                                                    _ = v70
                                                                    _ = v71
                                                                }
                                                            }
                                                        case -9223372036854775808:
                                                            switch "\n\t" {
                                                            case "foo bar":
                                                                var v72 float32 = -2.25
                                                                _ = v72
                                                            }
                                                        }
                                                    } else {
                                                        delete(v1, (fi20 / ((v5) | 1)))
                                                        v73 := GenS6{F0: true, F1: 1}
                                                        _ = v73
                                                    }
                                                    v7 = (v7 ^ v7)
                                                    v74 := GenS13{F0: 981, F1: "\n\t"}
                                                    _ = v74
                                                default:
                                                    v7 -= 42
                                                    switch "hello" {
                                                    case "hello":
                                                        switch (^(v5)) {
                                                        case -1:
                                                            var v75 uint32 = 286
                                                            _ = v75
                                                        default:
                                                            if (1.5 != -2.25) {
                                                                var v76 bool = ((!true) || v33)
                                                                switch (("a" + "hello") + "hello") {
                                                                case "foo bar":
                                                                    v7 = v7
                                                                case "hello":
                                                                    if v33 {
                                                                        var v77 float32 = 0.1
                                                                        var v78 uint64 = 4294967295
                                                                        if v76 {
                                                                            v78 -= (v78 ^ (v78 + v78))
                                                                        } else {
                                                                            _ = len(v8)
                                                                        }
                                                                        var v79 float64 = (-2.25 - float64(v77))
                                                                        _ = v77
                                                                        _ = v78
                                                                        _ = v79
                                                                    }
                                                                case "\n\t":
                                                                    v8 = append(v8, v5)
                                                                default:
                                                                    v31 -= 159
                                                                    var v80 string = ("foo bar" + "foo bar")
                                                                    if (874 < 0) {
                                                                        v31 -= (^(v31))
                                                                        if len(v8) >= 2 {
                                                                            v8 = v8[1:len(v8)]
                                                                        }
                                                                        v31 = v31
                                                                        var v81 bool = (52 <= v7)
                                                                        _ = v81
                                                                    }
                                                                    _ = v80
                                                                }
                                                                _ = v76
                                                            }
                                                        }
                                                    case "foo bar":
                                                        var v82 int8 = -105
                                                        _ = v82
                                                    default:
                                                        v8 = append(v8, (v2 / ((-226) | 1)))
                                                    }
                                                    v14 -= (v14 * (v14 / ((42) | 1)))
                                                    if v33 {
                                                        fmt.Println(fi19, v31, v5)
                                                    }
                                                }
                                                fmt.Println(v14, fi19)
                                                v83 := GenS7{F0: v14, F1: (-(fi20))}
                                                _ = v83
                                            } else {
                                                switch fi20 {
                                                case 566:
                                                    _ = len(v6)
                                                case 4294967295:
                                                    v5 += v2
                                                default:
                                                    var v84 bool = ((v7 / (((v7 ^ v7)) | 1)) == 944)
                                                    _ = v84
                                                }
                                                v7 = v7
                                            }
                                            _ = v33
                                        case -2147483648:
                                            if (-45 != 0) {
                                                var v85 int8 = int8(v7)
                                                var v86 int = fi19
                                                _ = v85
                                                _ = v86
                                            } else {
                                                var v87 string = "a"
                                                v7 += v7
                                                v88 := GenS4{F0: (v31 | v31)}
                                                _ = v87
                                                _ = v88
                                            }
                                        default:
                                            if true {
                                                _ = len(v8)
                                            } else {
                                                var v89 int16 = 81
                                                v90 := GenS2{F0: 18446744073709551615, F1: 18446744073709551615}
                                                switch int(v7) {
                                                case 1:
                                                    _ = len(v8)
                                                default:
                                                    v91 := GenS14{F0: int(v7)}
                                                    _ = v91
                                                }
                                                _ = v89
                                                _ = v90
                                            }
                                        }
                                    }
                                    _ = v31
                                }
                                v92 := GenS2{F0: 18446744073709551615, F1: 42}
                                _ = v92
                            }
                        }
                        var v93 string = ("hello" + "")
                        _ = v93
                    }
                }
                _ = fi20
                _ = v21
            }
            if true {
                v2 = (fi19 & (v2 | v5))
            } else {
                delete(v1, v5)
                fmt.Println(v2, fi19, v14)
                switch "a" {
                case "a":
                    v94 := GenS3{F0: 32767}
                    _ = v94
                case "hello":
                    var v95 uint16 = v7
                    _ = v95
                default:
                    v14 = v14
                }
            }
        } else {
            _ = len(v1)
        }
        _ = fi19
    }
    var v96 uint = 4294967295
    switch "" {
    case "foo bar":
        v7 += v7
    default:
        if false {
            var v97 uint16 = 0
            v98 := GenS12{F0: v7, F1: 1, F2: v14}
            _ = v97
            _ = v98
        } else {
            switch 172 {
            case -21:
                if true {
                    var v99 bool = (-1 == 1)
                    v99 = v99
                    _ = v99
                }
            }
            var v100 float32 = 0.0
            v101 := GenS15{F0: v14}
            _ = v100
            _ = v101
        }
    }
    v102 := GenS10{F0: -1, F1: 42, F2: 103}
    fmt.Println(v2, v5)
    v103 := GenS16{F0: 252, F1: 0, F2: ("hello" + "\n\t")}
    switch "a" {
    case "":
        v3[-346] = v5
    default:
        var v104 int8 = 1
        _ = v104
    }
    var v105 int64 = -530
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v105:", v105)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v96:", v96)
    _ = v1
    _ = v102
    _ = v103
    _ = v105
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v96
}
