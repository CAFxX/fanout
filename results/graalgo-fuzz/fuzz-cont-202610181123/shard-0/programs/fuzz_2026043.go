package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 string; F1 uint; F2 uint }
type GenS5 struct { F0 uint; F1 uint; F2 int }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint; F1 string; F2 int }
type GenS8 struct { F0 uint16; F1 string; F2 int }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 uint64; F1 uint16; F2 int8 }
type GenS11 struct { F0 int }
type GenS12 struct { F0 string }
type GenS13 struct { F0 uint; F1 int; F2 int }
type GenS14 struct { F0 uint32; F1 int32 }
type GenS15 struct { F0 int16; F1 uint8 }
type GenS16 struct { F0 uint64 }
type GenS17 struct { F0 int16; F1 int32 }
type GenS18 struct { F0 string; F1 uint16 }
type GenS19 struct { F0 bool }
type GenS20 struct { F0 uint32; F1 int; F2 uint16 }
type GenS21 struct { F0 int32 }
type GenS22 struct { F0 string; F1 string; F2 uint64 }

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[string]uint32)
    var v3 uint8 = 127
    v4 := make(map[uint32]string)
    v5 := make(map[string]string)
    v6 := []string{("\n\t" + "\n\t")}
    v3 = (v3 + 0)
    var v7 int = 4294967295
    fmt.Println(v7)
    v7 += v7
    switch v7 {
    case 41:
        if (42 <= -1) {
            for fi8 := 0; fi8 < 12; fi8++ {
                var v9 string = ("foo bar" + "hello")
                if true {
                    if true {
                        var v10 int8 = 42
                        _ = v10
                    } else {
                        if true {
                            switch (v7 & v7) {
                            case 2147483647:
                                switch fi8 {
                                case 600:
                                    delete(v5, v9)
                                default:
                                    fmt.Println(v9, v7, v9)
                                }
                            default:
                                if false {
                                    fmt.Println(fi8, v3)
                                    var v11 int = v7
                                    v3 = (^(v3))
                                    _ = v11
                                } else {
                                    v3 = 138
                                    var v12 int8 = 100
                                    v13 := GenS1{F0: (((v9 + v9) + v9) <= "hello"), F1: 255}
                                    _ = v12
                                    _ = v13
                                }
                            }
                            if (("hello" + "") != v9) {
                                for fi14 := 0; fi14 < 7; fi14++ {
                                    var v15 uint8 = (v3 * (v3 * 42))
                                    switch ("\n\t" + v9) {
                                    case "a":
                                        v16 := GenS1{F0: (!true), F1: 1}
                                        _ = v16
                                    case "foo bar":
                                        v15 -= 42
                                    }
                                    _ = fi14
                                    _ = v15
                                }
                            }
                        } else {
                            if (0 <= 125) {
                                delete(v5, (v9 + "foo bar"))
                                var v17 uint64 = 4294967295
                                if true {
                                    _ = len(v6)
                                    v18 := GenS2{F0: -718}
                                    var v19 int16 = 1
                                    _ = v18
                                    _ = v19
                                } else {
                                    switch v7 {
                                    case 42:
                                        for fi20 := 0; fi20 < 1; fi20++ {
                                            switch "" {
                                            case "foo bar":
                                                var v21 int32 = 1
                                                _ = v21
                                            case "":
                                                var v22 uint = 18446744073709551615
                                                if (false || (!(fi8 > fi20))) {
                                                    var v23 string = "foo bar"
                                                    switch v9 {
                                                    case "\n\t":
                                                        v24 := GenS2{F0: 42}
                                                        _ = v24
                                                    case "":
                                                        v25 := v1["hello"]
                                                        _ = v25
                                                    default:
                                                        v26 := GenS2{F0: 0}
                                                        _ = v26
                                                    }
                                                    _ = v23
                                                } else {
                                                    switch ("a" + "\n\t") {
                                                    case "":
                                                        v17 += (v17 + 483)
                                                    }
                                                    if len(v6) > 0 {
                                                        v6[0] = "a"
                                                    }
                                                    v27 := GenS3{F0: (9223372036854775807 > (fi8 | 141))}
                                                    _ = v27
                                                }
                                                switch (fi8 << ((18446744073709551615) & 63)) {
                                                case 128:
                                                    switch 1 {
                                                    case -211:
                                                        v28 := GenS2{F0: 1}
                                                        var v29 float64 = 0.1
                                                        var v30 int8 = 127
                                                        delete(v5, "foo bar")
                                                        var v31 int64 = -1
                                                        if (v9 < v9) {
                                                            var v32 uint64 = (v17 >> ((4294967295) & 63))
                                                            v29 = 1.5
                                                            v33 := GenS4{F0: v9, F1: (v22 | v22), F2: v22}
                                                            _ = v32
                                                            _ = v33
                                                        }
                                                        _ = v28
                                                        _ = v29
                                                        _ = v30
                                                        _ = v31
                                                    case 944:
                                                        if ((v17 >= v17) && (-128 == -35)) {
                                                            v3 = (v3 | 127)
                                                            if false {
                                                                v3 -= v3
                                                                if true {
                                                                    v34 := GenS4{F0: ("hello" + ((v9 + (v9 + (((v9 + "a") + "hello") + ("hello" + "")))) + (v9 + v9))), F1: v22, F2: v22}
                                                                    v35 := GenS5{F0: 179, F1: 987, F2: fi8}
                                                                    _ = v34
                                                                    _ = v35
                                                                } else {
                                                                    var v36 uint = v22
                                                                    _ = v36
                                                                }
                                                                if (!true) {
                                                                    v37 := v1[v9]
                                                                    fmt.Println(v7, v7)
                                                                    switch v9 {
                                                                    case "\n\t":
                                                                        v17 = (v17 ^ v17)
                                                                    case "a":
                                                                        switch int(v22) {
                                                                        case 1:
                                                                            v9 = ("hello" + (v9 + (v9 + ((("foo bar" + v9) + v9) + (v9 + v9)))))
                                                                        default:
                                                                            var v38 uint64 = v17
                                                                            _ = v38
                                                                        }
                                                                        if true {
                                                                            v39 := GenS6{F0: 9223372036854775807}
                                                                            _ = v39
                                                                        }
                                                                        v22 = uint(v17)
                                                                    case "":
                                                                        var v40 int64 = -9223372036854775808
                                                                        _ = v40
                                                                    }
                                                                    _ = v37
                                                                } else {
                                                                    if true {
                                                                        v3 = (^(v3))
                                                                        fmt.Println(fi20)
                                                                    } else {
                                                                        if (647 != 2) {
                                                                            v41 := GenS6{F0: v7}
                                                                            var v42 int8 = 123
                                                                            _ = v41
                                                                            _ = v42
                                                                        }
                                                                        if true {
                                                                            var v43 uint8 = (v3 / ((2) | 1))
                                                                            var v44 uint32 = 65535
                                                                            switch (fi8 | (fi8 >> ((v22) & 63))) {
                                                                            case -2147483648:
                                                                                if (!(((v9 + v9) + "hello") < ((v9 + v9) + (v9 + v9)))) {
                                                                                    if len(v6) > 0 {
                                                                                        v6[0] = v9
                                                                                    }
                                                                                    fmt.Println(fi8)
                                                                                }
                                                                            case 32768:
                                                                                switch "\n\t" {
                                                                                case "":
                                                                                    _ = len(v2)
                                                                                default:
                                                                                    if false {
                                                                                        var v45 int16 = 32767
                                                                                        if len(v6) >= 2 {
                                                                                            v6 = v6[1:len(v6)]
                                                                                        }
                                                                                        _ = v45
                                                                                    } else {
                                                                                        v46 := GenS1{F0: ((65535 >= 252) || (true || false)), F1: (^(v22))}
                                                                                        _ = v46
                                                                                    }
                                                                                }
                                                                            default:
                                                                                var v47 bool = (false && (false && true))
                                                                                _ = v47
                                                                            }
                                                                            _ = v43
                                                                            _ = v44
                                                                        } else {
                                                                            var v48 uint8 = v3
                                                                            v48 += (v48 + v3)
                                                                            var v49 int64 = 590
                                                                            _ = v48
                                                                            _ = v49
                                                                        }
                                                                        v7 += (fi20 + v7)
                                                                    }
                                                                    v7 = fi20
                                                                    var v50 uint8 = v3
                                                                    _ = v50
                                                                }
                                                            }
                                                            var v51 uint8 = (v3 % ((v3) | 1))
                                                            _ = v51
                                                        } else {
                                                            var v52 uint = 1
                                                            _ = v52
                                                        }
                                                    case 4294967295:
                                                        v53 := GenS7{F0: v22, F1: ("" + v9), F2: 0}
                                                        _ = v53
                                                    default:
                                                        fmt.Println(v9, v3)
                                                    }
                                                case 986:
                                                    var v54 int8 = -128
                                                    _ = v54
                                                }
                                                _ = v22
                                            case "hello":
                                                if (4294967295 == 779) {
                                                    if false {
                                                        var v55 uint16 = 97
                                                        if ((false && false) || true) {
                                                            switch (v9 + (v9 + "a")) {
                                                            case "":
                                                                if false {
                                                                    var v56 int32 = 42
                                                                    v57 := GenS6{F0: 2147483647}
                                                                    var v58 int64 = 0
                                                                    _ = v56
                                                                    _ = v57
                                                                    _ = v58
                                                                } else {
                                                                    switch "foo bar" {
                                                                    case "foo bar":
                                                                        if (255 != 708) {
                                                                            v59 := GenS3{F0: true}
                                                                            v60 := GenS8{F0: (v55 << ((2) & 63)), F1: v9, F2: v7}
                                                                            _ = v59
                                                                            _ = v60
                                                                        } else {
                                                                            v9 = ("\n\t" + v9)
                                                                            v9 = ((v9 + v9) + "\n\t")
                                                                            fmt.Println(v9, v3, fi8)
                                                                            fmt.Println(v7, v3, v9)
                                                                        }
                                                                    case "a":
                                                                        fmt.Println(v17, v9, v9)
                                                                    default:
                                                                        v9 = (v9 + ("foo bar" + ("foo bar" + v9)))
                                                                    }
                                                                    if len(v6) > 0 {
                                                                        v6[0] = "hello"
                                                                    }
                                                                }
                                                            case "foo bar":
                                                                var v61 uint64 = (v17 | v17)
                                                                _ = v61
                                                            case "\n\t":
                                                                if true {
                                                                    if len(v6) > 0 {
                                                                        v6[0] = (v9 + (v9 + v9))
                                                                    }
                                                                    if true {
                                                                        if ((!(false || true)) || (v3 < 161)) {
                                                                            v17 += v17
                                                                            v62 := GenS8{F0: 0, F1: v9, F2: fi8}
                                                                            v3 += v3
                                                                            _ = v62
                                                                        }
                                                                    }
                                                                    if (3.14159 != (1e10 + float32(-2.25))) {
                                                                        switch v7 {
                                                                        case 119:
                                                                            switch fi20 {
                                                                            case 127:
                                                                                var v63 int = (^(fi20))
                                                                                _ = v63
                                                                            case 1:
                                                                                v9 = ""
                                                                                v2[v9] = 582
                                                                            default:
                                                                                v17 = (^(v17))
                                                                                v55 -= 42
                                                                                switch (v7 >> ((2) & 63)) {
                                                                                case 1:
                                                                                    var v64 int64 = 0
                                                                                    _ = v64
                                                                                default:
                                                                                    if true {
                                                                                        fmt.Println(v9)
                                                                                        var v65 uint32 = 684
                                                                                        v66 := GenS9{F0: v55}
                                                                                        _ = v65
                                                                                        _ = v66
                                                                                    }
                                                                                }
                                                                                v9 = (v9 + v9)
                                                                            }
                                                                        default:
                                                                            fmt.Println(v9, fi20)
                                                                        }
                                                                        var v67 int8 = -128
                                                                        var v68 uint32 = 2
                                                                        _ = v67
                                                                        _ = v68
                                                                    }
                                                                } else {
                                                                    switch 4294967295 {
                                                                    case -2147483648:
                                                                        var v69 uint64 = v17
                                                                        _ = v69
                                                                    case -453:
                                                                        var v70 bool = true
                                                                        v9 = v9
                                                                        if len(v6) >= 2 {
                                                                            v6 = v6[1:len(v6)]
                                                                        }
                                                                        var v71 string = (v9 + "foo bar")
                                                                        if (-2.25 >= ((3.14159 - 0.1) * (3.14159 * 1e10))) {
                                                                            v70 = v70
                                                                            fmt.Println(v3, fi8, v7)
                                                                            v71 = (v9 + v71)
                                                                        } else {
                                                                            v9 = ""
                                                                            var v72 uint32 = 4294967295
                                                                            _ = v72
                                                                        }
                                                                        _ = v70
                                                                        _ = v71
                                                                    case 4294967295:
                                                                        if ((^(v3)) >= v3) {
                                                                            if false {
                                                                                if (-644 > 1) {
                                                                                    var v73 uint8 = v3
                                                                                    switch ("a" + "") {
                                                                                    case "hello":
                                                                                        if true {
                                                                                            if true {
                                                                                                var v74 int16 = 0
                                                                                                _ = v74
                                                                                            }
                                                                                            if (179 <= 2147483648) {
                                                                                                switch v9 {
                                                                                                case "":
                                                                                                    v75 := GenS10{F0: 2, F1: v55, F2: 1}
                                                                                                    _ = v75
                                                                                                case "hello":
                                                                                                    var v76 float32 = 1e10
                                                                                                    if true {
                                                                                                        if true {
                                                                                                            v77 := GenS1{F0: false, F1: 994}
                                                                                                            switch (fi20 % ((32768) | 1)) {
                                                                                                            case 32767:
                                                                                                                var v78 uint8 = (^(v73))
                                                                                                                _ = v78
                                                                                                            default:
                                                                                                                switch v9 {
                                                                                                                case "hello":
                                                                                                                    v17 = 42
                                                                                                                    fmt.Println(fi8, v3, v9)
                                                                                                                case "\n\t":
                                                                                                                    if (!false) {
                                                                                                                        v1[(v9 + "a")] = 457
                                                                                                                    } else {
                                                                                                                        _ = len(v2)
                                                                                                                        if false {
                                                                                                                            v79 := GenS11{F0: v7}
                                                                                                                            var v80 float64 = (float64(1e10) + 0.0)
                                                                                                                            v81 := v1["hello"]
                                                                                                                            v82 := GenS1{F0: false, F1: 1}
                                                                                                                            if (255 <= 586) {
                                                                                                                                switch v9 {
                                                                                                                                case "a":
                                                                                                                                    if (!true) {
                                                                                                                                        fmt.Println(v9)
                                                                                                                                        switch (v7 & v7) {
                                                                                                                                        case -228:
                                                                                                                                            v83 := GenS11{F0: fi8}
                                                                                                                                            _ = v83
                                                                                                                                        case 128:
                                                                                                                                            fmt.Println(v17, v7, v7)
                                                                                                                                        case -9223372036854775808:
                                                                                                                                            if ((v9 + "\n\t") < (v9 + ((("" + v9) + v9) + ""))) {
                                                                                                                                                var v84 int64 = 0
                                                                                                                                                if (!(v17 >= (^(v17)))) {
                                                                                                                                                    fmt.Println(fi8, v80)
                                                                                                                                                }
                                                                                                                                                switch fi8 {
                                                                                                                                                case -1:
                                                                                                                                                    v55 -= v55
                                                                                                                                                    var v85 int = (fi8 % (((v7 ^ -145)) | 1))
                                                                                                                                                    if false {
                                                                                                                                                        if true {
                                                                                                                                                            var v86 uint64 = (v17 / ((814) | 1))
                                                                                                                                                            v9 = ((("" + ("foo bar" + v9)) + "") + (v9 + "\n\t"))
                                                                                                                                                            v87 := GenS8{F0: v55, F1: (v9 + (v9 + v9)), F2: fi8}
                                                                                                                                                            _ = v86
                                                                                                                                                            _ = v87
                                                                                                                                                        }
                                                                                                                                                    }
                                                                                                                                                    _ = v85
                                                                                                                                                }
                                                                                                                                                v88 := GenS1{F0: true, F1: uint(v55)}
                                                                                                                                                var v89 bool = false
                                                                                                                                                if ((true && v89) && ((42 == v55) && (0 <= 1))) {
                                                                                                                                                    fmt.Println(v3, v80)
                                                                                                                                                } else {
                                                                                                                                                    v90 := GenS11{F0: fi8}
                                                                                                                                                    _ = v90
                                                                                                                                                }
                                                                                                                                                _ = v84
                                                                                                                                                _ = v88
                                                                                                                                                _ = v89
                                                                                                                                            }
                                                                                                                                            var v91 uint32 = 0
                                                                                                                                            var v92 int = -59
                                                                                                                                            _ = v91
                                                                                                                                            _ = v92
                                                                                                                                        }
                                                                                                                                        fmt.Println(v17, v9, v80)
                                                                                                                                    } else {
                                                                                                                                        fmt.Println(fi20, fi20)
                                                                                                                                        var v93 int32 = 578
                                                                                                                                        v94 := GenS5{F0: 1, F1: 498, F2: 1}
                                                                                                                                        _ = v93
                                                                                                                                        _ = v94
                                                                                                                                    }
                                                                                                                                default:
                                                                                                                                    v95 := v1[((v9 + v9) + ((v9 + (v9 + "foo bar")) + v9))]
                                                                                                                                    _ = v95
                                                                                                                                }
                                                                                                                            } else {
                                                                                                                                if false {
                                                                                                                                    if true {
                                                                                                                                        if len(v6) > 0 {
                                                                                                                                            v6[0] = (v9 + v9)
                                                                                                                                        }
                                                                                                                                        if true {
                                                                                                                                            var v96 uint16 = 892
                                                                                                                                            v97 := GenS12{F0: v9}
                                                                                                                                            if (!(42 != 127)) {
                                                                                                                                                v17 = (^(v17))
                                                                                                                                                switch (fi20 * (fi8 + -733)) {
                                                                                                                                                case 32768:
                                                                                                                                                    v3 += (v73 ^ (^(v3)))
                                                                                                                                                case 127:
                                                                                                                                                    var v98 uint64 = (v17 << ((42) & 63))
                                                                                                                                                    _ = v98
                                                                                                                                                case 4294967295:
                                                                                                                                                    var v99 uint64 = v17
                                                                                                                                                    _ = v99
                                                                                                                                                default:
                                                                                                                                                    _ = len(v2)
                                                                                                                                                }
                                                                                                                                                var v100 uint32 = (v81 ^ v81)
                                                                                                                                                var v101 bool = (!(v3 > (v73 | (v3 + v3))))
                                                                                                                                                fmt.Println(fi8, v3, v17)
                                                                                                                                                _ = v100
                                                                                                                                                _ = v101
                                                                                                                                            }
                                                                                                                                            _ = v96
                                                                                                                                            _ = v97
                                                                                                                                        }
                                                                                                                                    } else {
                                                                                                                                        v55 = 255
                                                                                                                                        _ = len(v4)
                                                                                                                                    }
                                                                                                                                    v102 := GenS13{F0: 255, F1: v7, F2: (fi20 ^ fi20)}
                                                                                                                                    _ = v102
                                                                                                                                } else {
                                                                                                                                    var v103 int8 = 106
                                                                                                                                    _ = v103
                                                                                                                                }
                                                                                                                                _ = len(v6)
                                                                                                                                var v104 float64 = v80
                                                                                                                                _ = v104
                                                                                                                            }
                                                                                                                            v80 -= 0.1
                                                                                                                            _ = v79
                                                                                                                            _ = v80
                                                                                                                            _ = v81
                                                                                                                            _ = v82
                                                                                                                        } else {
                                                                                                                            if len(v6) >= 2 {
                                                                                                                                v6 = v6[1:len(v6)]
                                                                                                                            }
                                                                                                                            v55 = uint16(v17)
                                                                                                                        }
                                                                                                                        var v105 uint8 = v3
                                                                                                                        v17 = (v17 % (((^(v17))) | 1))
                                                                                                                        fmt.Println(v73)
                                                                                                                        _ = v105
                                                                                                                    }
                                                                                                                case "a":
                                                                                                                    fmt.Println(fi20, v3)
                                                                                                                default:
                                                                                                                    v17 -= (^(v17))
                                                                                                                }
                                                                                                            }
                                                                                                            fmt.Println(v17)
                                                                                                            var v106 int64 = -101
                                                                                                            _ = v106
                                                                                                            _ = v77
                                                                                                        }
                                                                                                        var v107 uint16 = v55
                                                                                                        _ = v107
                                                                                                    } else {
                                                                                                        var v108 int = fi20
                                                                                                        fmt.Println(v7)
                                                                                                        _ = v108
                                                                                                    }
                                                                                                    _ = v76
                                                                                                default:
                                                                                                    var v109 uint16 = (v55 + v55)
                                                                                                    _ = v109
                                                                                                }
                                                                                            } else {
                                                                                                if ((((v9 + v9) + (v9 + ("" + (v9 + v9)))) + v9) >= v9) {
                                                                                                    fmt.Println(v9, fi20)
                                                                                                    fmt.Println(v7, v3)
                                                                                                    var v110 uint16 = (v55 >> ((1) & 63))
                                                                                                    _ = v110
                                                                                                }
                                                                                                if true {
                                                                                                    v7 = fi20
                                                                                                    switch ((v9 + v9) + "") {
                                                                                                    case "hello":
                                                                                                        v111 := GenS2{F0: int32(v3)}
                                                                                                        _ = v111
                                                                                                    case "a":
                                                                                                        v73 = 2
                                                                                                    case "foo bar":
                                                                                                        v6 = append(v6, "\n\t")
                                                                                                        switch v9 {
                                                                                                        case "\n\t":
                                                                                                            v3 += v73
                                                                                                        }
                                                                                                        v73 = uint8(v55)
                                                                                                    default:
                                                                                                        if len(v6) >= 2 {
                                                                                                            v6 = v6[1:len(v6)]
                                                                                                        }
                                                                                                    }
                                                                                                    v112 := GenS7{F0: 65535, F1: v9, F2: (v7 * v7)}
                                                                                                    _ = v112
                                                                                                } else {
                                                                                                    var v113 int8 = 1
                                                                                                    fmt.Println(v7, v3, v113)
                                                                                                    v9 = (("\n\t" + v9) + "foo bar")
                                                                                                    _ = v113
                                                                                                }
                                                                                                var v114 uint8 = (v3 >> ((4294967295) & 63))
                                                                                                _ = v114
                                                                                            }
                                                                                        }
                                                                                    case "\n\t":
                                                                                        v9 = v9
                                                                                    default:
                                                                                        switch v9 {
                                                                                        case "foo bar":
                                                                                            v6 = append(v6, v9)
                                                                                        default:
                                                                                            switch v9 {
                                                                                            case "foo bar":
                                                                                                switch "hello" {
                                                                                                case "\n\t":
                                                                                                    fmt.Println(v55, v9)
                                                                                                case "a":
                                                                                                    if (1.5 > (((0.1 + 1.5) * float64((0.1 * 0.1))) - 1.5)) {
                                                                                                        var v115 uint16 = (v55 % ((v55) | 1))
                                                                                                        v116 := v2[((v9 + (v9 + "a")) + (v9 + ("a" + v9)))]
                                                                                                        v117 := GenS14{F0: (v116 << ((0) & 63)), F1: -2147483648}
                                                                                                        _ = v115
                                                                                                        _ = v116
                                                                                                        _ = v117
                                                                                                    } else {
                                                                                                        var v118 uint = 4294967295
                                                                                                        _ = v118
                                                                                                    }
                                                                                                default:
                                                                                                    var v119 int32 = int32(fi20)
                                                                                                    _ = v119
                                                                                                }
                                                                                            case "\n\t":
                                                                                                v73 = (v3 ^ (v73 % ((v3) | 1)))
                                                                                            case "hello":
                                                                                                switch v9 {
                                                                                                case "":
                                                                                                    v120 := GenS14{F0: 2, F1: -1}
                                                                                                    _ = v120
                                                                                                }
                                                                                            default:
                                                                                                if false {
                                                                                                    v17 -= (v17 >> ((887) & 63))
                                                                                                    var v121 int8 = 84
                                                                                                    if (((v55 / ((v55) | 1)) < v55) || (2 >= 2)) {
                                                                                                        if len(v6) > 0 {
                                                                                                            v6[0] = v9
                                                                                                        }
                                                                                                    }
                                                                                                    var v122 int32 = 2147483647
                                                                                                    var v123 bool = true
                                                                                                    _ = v121
                                                                                                    _ = v122
                                                                                                    _ = v123
                                                                                                } else {
                                                                                                    v124 := GenS15{F0: 32767, F1: v3}
                                                                                                    var v125 int64 = 9223372036854775807
                                                                                                    if (0 == -128) {
                                                                                                        v126 := GenS3{F0: (false && false)}
                                                                                                        _ = v126
                                                                                                    } else {
                                                                                                        var v127 int32 = -2147483648
                                                                                                        v125 = v125
                                                                                                        fmt.Println(fi20, v127)
                                                                                                        switch (v9 + (v9 + v9)) {
                                                                                                        case "a":
                                                                                                            v55 -= 2
                                                                                                        case "":
                                                                                                            v3 = v3
                                                                                                        default:
                                                                                                            if (!(!(235 < 0))) {
                                                                                                                var v128 uint8 = uint8(v17)
                                                                                                                _ = len(v6)
                                                                                                                var v129 string = (((v9 + v9) + "a") + "a")
                                                                                                                _ = v128
                                                                                                                _ = v129
                                                                                                            }
                                                                                                        }
                                                                                                        var v130 int16 = -1
                                                                                                        fmt.Println(v125, v17, v127)
                                                                                                        _ = v127
                                                                                                        _ = v130
                                                                                                    }
                                                                                                    _ = v124
                                                                                                    _ = v125
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    if (fi8 <= -9223372036854775808) {
                                                                                        v7 -= 1
                                                                                        v3 = 255
                                                                                        if false {
                                                                                            fmt.Println(v55)
                                                                                            v131 := GenS2{F0: 603}
                                                                                            v132 := GenS13{F0: 173, F1: (fi20 >> ((42) & 63)), F2: (^(fi20))}
                                                                                            v133 := v2[v9]
                                                                                            v134 := GenS11{F0: fi8}
                                                                                            _ = v131
                                                                                            _ = v132
                                                                                            _ = v133
                                                                                            _ = v134
                                                                                        } else {
                                                                                            if (((255 == 201) || (true && true)) || (false || ((^(v55)) > v55))) {
                                                                                                v135 := GenS16{F0: v17}
                                                                                                var v136 uint32 = 1
                                                                                                _ = v135
                                                                                                _ = v136
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                    _ = v73
                                                                                }
                                                                                v137 := GenS17{F0: 32767, F1: 2147483647}
                                                                                switch int(v3) {
                                                                                case 0:
                                                                                    v7 = 516
                                                                                case -9223372036854775808:
                                                                                    v138 := v2["hello"]
                                                                                    _ = v138
                                                                                case 127:
                                                                                    v139 := GenS10{F0: 0, F1: 849, F2: 127}
                                                                                    if len(v6) > 0 {
                                                                                        v6[0] = v9
                                                                                    }
                                                                                    _ = v139
                                                                                }
                                                                                _ = v137
                                                                            } else {
                                                                                switch "foo bar" {
                                                                                case "a":
                                                                                    _ = len(v6)
                                                                                }
                                                                                switch fi8 {
                                                                                case 32768:
                                                                                    v140 := v4[2]
                                                                                    _ = v140
                                                                                default:
                                                                                    var v141 int8 = 38
                                                                                    _ = v141
                                                                                }
                                                                                var v142 float32 = (1.5 + 0.0)
                                                                                _ = v142
                                                                            }
                                                                            if true {
                                                                                switch v9 {
                                                                                case "a":
                                                                                    v55 -= (^(v55))
                                                                                case "foo bar":
                                                                                    switch fi8 {
                                                                                    case -249:
                                                                                        v3 = 16
                                                                                    case -9223372036854775808:
                                                                                        _ = len(v4)
                                                                                    case 42:
                                                                                        v7 += 432
                                                                                    }
                                                                                case "hello":
                                                                                    v9 = "foo bar"
                                                                                }
                                                                                if false {
                                                                                    v7 += fi20
                                                                                    var v143 float64 = (-2.25 / 1.0)
                                                                                    _ = v143
                                                                                }
                                                                            } else {
                                                                                v7 += (fi20 * v7)
                                                                                v9 = v9
                                                                                var v144 int32 = -1
                                                                                _ = v144
                                                                            }
                                                                            v145 := GenS16{F0: v17}
                                                                            _ = v145
                                                                        } else {
                                                                            if len(v6) > 0 {
                                                                                v6[0] = v9
                                                                            }
                                                                            v146 := GenS18{F0: v9, F1: uint16(v3)}
                                                                            _ = len(v6)
                                                                            _ = v146
                                                                        }
                                                                    default:
                                                                        var v147 bool = true
                                                                        _ = v147
                                                                    }
                                                                    v3 = (v3 * 128)
                                                                    v55 = 2
                                                                }
                                                            }
                                                            if ("" != v9) {
                                                                _ = len(v6)
                                                            }
                                                        } else {
                                                            v6 = append(v6, v9)
                                                        }
                                                        var v148 string = "foo bar"
                                                        _ = v148
                                                        _ = v55
                                                    } else {
                                                        var v149 uint8 = (v3 | 42)
                                                        v9 = "\n\t"
                                                        _ = v149
                                                    }
                                                    fmt.Println(v17, fi20, v17)
                                                }
                                            default:
                                                var v150 uint64 = v17
                                                v151 := v4[4294967295]
                                                _ = v150
                                                _ = v151
                                            }
                                            switch -14 {
                                            case 42:
                                                v17 = (v17 & 491)
                                            default:
                                                v3 = (^(v3))
                                            }
                                            fmt.Println(v17, v17)
                                            _ = fi20
                                        }
                                    case 127:
                                        _ = len(v4)
                                    case 32767:
                                        if false {
                                            v152 := GenS19{F0: true}
                                            var v153 float64 = 1e10
                                            v154 := v5[""]
                                            _ = v152
                                            _ = v153
                                            _ = v154
                                        }
                                    default:
                                        for fi155 := 0; fi155 < 9; fi155++ {
                                            _ = len(v5)
                                            _ = fi155
                                        }
                                    }
                                }
                                _ = v17
                            }
                            var v156 uint = 0
                            v156 = 1
                            _ = v156
                        }
                        v157 := GenS12{F0: "a"}
                        v9 = v9
                        _ = v157
                    }
                    if (v9 > v9) {
                        switch (v9 + "a") {
                        case "foo bar":
                            fmt.Println(v9, v3, fi8)
                        case "\n\t":
                            v9 = v9
                        default:
                            if (!true) {
                                v158 := GenS20{F0: 263, F1: v7, F2: 859}
                                switch (v7 * (v7 - (fi8 - (v7 * 128)))) {
                                case -2147483648:
                                    var v159 uint32 = 4294967295
                                    _ = v159
                                case 20:
                                    switch v9 {
                                    case "hello":
                                        var v160 int32 = -1
                                        _ = v160
                                    case "a":
                                        if len(v6) > 0 {
                                            v6[0] = ("hello" + ((v9 + v9) + "\n\t"))
                                        }
                                    case "foo bar":
                                        if true {
                                            v3 = v3
                                            if (((0.0 * (1.5 + float64((0.0 * 0.1)))) + (float64((1.5 * (((float32(3.14159) - 3.14159) * (1e10 - float32(0.0))) - (1.5 / -2.25)))) + -2.25)) == (1.5 - 1e10)) {
                                                if false {
                                                    v9 = v9
                                                }
                                                if len(v6) > 0 {
                                                    v6[0] = "hello"
                                                }
                                                v3 -= (v3 - (v3 | (v3 ^ v3)))
                                            }
                                        } else {
                                            v161 := GenS16{F0: 1}
                                            _ = v161
                                        }
                                    }
                                }
                                _ = v158
                            } else {
                                var v162 uint32 = 42
                                for fi163 := 0; fi163 < 8; fi163++ {
                                    var v164 bool = false
                                    var v165 bool = v164
                                    switch (fi163 - (fi8 * 128)) {
                                    case -277:
                                        v166 := GenS21{F0: 42}
                                        if (610 == -535) {
                                            fmt.Println(fi163, v3, v165)
                                            fmt.Println(fi163)
                                        }
                                        _ = v166
                                    case 42:
                                        if len(v6) > 0 {
                                            v6[0] = ((v9 + v9) + v9)
                                        }
                                    default:
                                        v167 := GenS18{F0: (v9 + (v9 + (v9 + ((v9 + "a") + v9)))), F1: 2}
                                        _ = v167
                                    }
                                    _ = fi163
                                    _ = v164
                                    _ = v165
                                }
                                _ = v162
                            }
                        }
                    }
                    if (((3.14159 > float64(0.0)) && (false && true)) && (!false)) {
                        v168 := GenS21{F0: 42}
                        for fi169 := 0; fi169 < 13; fi169++ {
                            var v170 int32 = -2147483648
                            switch v9 {
                            case "foo bar":
                                var v171 int16 = 1
                                _ = v171
                            case "\n\t":
                                fmt.Println(v170)
                            case "hello":
                                v170 += (v170 / (((v170 + -2147483648)) | 1))
                            }
                            _ = fi169
                            _ = v170
                        }
                        _ = v168
                    }
                }
                _ = fi8
                _ = v9
            }
        } else {
            v3 = v3
            v3 = v3
            v7 += (v7 % ((v7) | 1))
        }
    case -135:
        v172 := v2[("a" + "hello")]
        _ = v172
    default:
        v7 -= v7
    }
    var v173 float32 = 0.0
    var v174 int16 = 32767
    v174 = (v174 + -1)
    var v175 uint32 = 42
    v173 += (v173 * (((v173 * v173) - 0.1) / v173))
    v174 += v174
    switch v7 {
    case -822:
        v175 -= v175
    case 809:
        v176 := GenS22{F0: ("\n\t" + "hello"), F1: ("" + "foo bar"), F2: 1}
        _ = v176
    case -1:
        v6 = append(v6, "")
    default:
        var v177 int32 = -388
        _ = v177
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v173:", v173)
    fmt.Println("v174:", v174)
    fmt.Println("v175:", v175)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v173
    _ = v174
    _ = v175
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
