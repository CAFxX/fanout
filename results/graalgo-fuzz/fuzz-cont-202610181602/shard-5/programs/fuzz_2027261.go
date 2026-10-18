package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32; F1 string; F2 bool }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 bool; F1 bool; F2 uint64 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 string; F1 int; F2 int32 }
type GenS9 struct { F0 bool; F1 int }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 uint16; F1 int8 }
type GenS12 struct { F0 string; F1 int64; F2 uint64 }
type GenS13 struct { F0 int32 }
type GenS14 struct { F0 int64 }
type GenS15 struct { F0 bool }

func main() {
    v1 := []uint32{}
    v2 := make(map[uint32]int)
    v3 := []string{"\n\t", "foo bar", ("hello" + "\n\t"), "\n\t"}
    var v4 uint16 = 32767
    var v5 uint = 18446744073709551615
    var v6 int64 = 0
    for fi7 := 0; fi7 < 6; fi7++ {
        fmt.Println(v4, v6)
        fmt.Println(v6)
        switch "foo bar" {
        case "":
            _ = len(v2)
        case "hello":
            if true {
                v5 = 4294967295
            } else {
                var v8 bool = (1e10 > (1.5 + -2.25))
                if v8 {
                    if (true && v8) {
                        switch (-(fi7)) {
                        case -249:
                            var v9 float32 = 3.14159
                            _ = v9
                        }
                        if true {
                            var v10 uint = (^(v5))
                            _ = v10
                        } else {
                            if len(v1) > 0 {
                                v1[0] = 42
                            }
                            fmt.Println(v6, v5)
                            v11 := GenS1{F0: fi7}
                            _ = v11
                        }
                    } else {
                        var v12 int = (^(fi7))
                        v3 = append(v3, "hello")
                        v2[233] = (fi7 / (((v12 | v12)) | 1))
                        _ = v12
                    }
                    if v8 {
                        v13 := v2[1]
                        _ = v13
                    } else {
                        v2[2147483648] = fi7
                        v14 := GenS1{F0: fi7}
                        v8 = (!((1e10 * 0.1) > (1e10 + 1e10)))
                        _ = v14
                    }
                    switch (fi7 & 0) {
                    case 4294967295:
                        var v15 int64 = (v6 | (-(v6)))
                        _ = v15
                    default:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                } else {
                    for fi16 := 0; fi16 < 5; fi16++ {
                        if v8 {
                            v17 := GenS2{F0: 65535, F1: "hello", F2: (100 == -27)}
                            v4 -= uint16(v5)
                            var v18 int32 = -2147483648
                            _ = v17
                            _ = v18
                        } else {
                            if true {
                                var v19 int = (-(fi16))
                                _ = v19
                            } else {
                                fmt.Println(v4, fi7)
                                switch "foo bar" {
                                case "hello":
                                    if false {
                                        v6 -= v6
                                        fmt.Println(v8, v4, fi7)
                                    } else {
                                        var v20 float32 = float32(1e10)
                                        if (131 > 181) {
                                            v20 = float32(0.1)
                                        } else {
                                            var v21 int64 = v6
                                            switch "foo bar" {
                                            case "foo bar":
                                                if v8 {
                                                    v6 -= v21
                                                } else {
                                                    if (v8 && false) {
                                                        fmt.Println(v5)
                                                        v22 := GenS3{F0: 1}
                                                        _ = v22
                                                    }
                                                }
                                            case "\n\t":
                                                v5 = (v5 << ((v5) & 63))
                                            case "":
                                                v3 = append(v3, "foo bar")
                                            }
                                            if v8 {
                                                switch 128 {
                                                case 127:
                                                    if v8 {
                                                        if (true && (!false)) {
                                                            switch (fi16 - fi7) {
                                                            case 42:
                                                                if false {
                                                                    var v23 int64 = v6
                                                                    v23 += 9223372036854775807
                                                                    v24 := GenS2{F0: 1, F1: "hello", F2: v8}
                                                                    _ = v23
                                                                    _ = v24
                                                                }
                                                            case -9223372036854775808:
                                                                v20 = (-2.25 - 0.1)
                                                                switch ("hello" + ("foo bar" + "hello")) {
                                                                case "\n\t":
                                                                    if v8 {
                                                                        if len(v1) > 0 {
                                                                            v1[0] = 0
                                                                        }
                                                                    }
                                                                case "hello":
                                                                    var v25 uint64 = 4294967295
                                                                    _ = v25
                                                                case "":
                                                                    var v26 string = ""
                                                                    switch v26 {
                                                                    case "foo bar":
                                                                        if (42 == -2147483648) {
                                                                            v27 := GenS2{F0: 2147483647, F1: v26, F2: v8}
                                                                            switch ((v26 + v26) + v26) {
                                                                            case "":
                                                                                v5 = (v5 * (v5 / ((v5) | 1)))
                                                                            case "hello":
                                                                                var v28 uint64 = 554
                                                                                _ = v28
                                                                            }
                                                                            _ = v27
                                                                        }
                                                                    default:
                                                                        delete(v2, 368)
                                                                    }
                                                                    _ = v26
                                                                default:
                                                                    var v29 int64 = (v6 << ((v5) & 63))
                                                                    _ = v29
                                                                }
                                                            }
                                                            var v30 int16 = -32768
                                                            if len(v1) > 0 {
                                                                v1[0] = 1
                                                            }
                                                            _ = v30
                                                        }
                                                        switch ("" + "a") {
                                                        case "a":
                                                            v20 = v20
                                                        case "":
                                                            if true {
                                                                _ = len(v2)
                                                                var v31 int64 = (v21 ^ v21)
                                                                v32 := GenS2{F0: 293, F1: "", F2: (-2147483648 >= -1)}
                                                                _ = v31
                                                                _ = v32
                                                            } else {
                                                                v20 = float32(1e10)
                                                                if false {
                                                                    var v33 uint64 = 2
                                                                    _ = v33
                                                                } else {
                                                                    fmt.Println(v20, fi16, v21)
                                                                    if v8 {
                                                                        var v34 uint16 = 1
                                                                        _ = v34
                                                                    } else {
                                                                        var v35 bool = v8
                                                                        var v36 int16 = -705
                                                                        _ = v35
                                                                        _ = v36
                                                                    }
                                                                }
                                                                v20 = ((3.14159 + ((float32(0.0) + (v20 * float32(3.14159))) - v20)) * 1.5)
                                                            }
                                                        case "\n\t":
                                                            v20 = float32(-2.25)
                                                        }
                                                        v6 -= v6
                                                    }
                                                case 42:
                                                    switch "hello" {
                                                    case "":
                                                        if v8 {
                                                            var v37 uint64 = 42
                                                            switch ("" + "") {
                                                            case "a":
                                                                v38 := GenS4{F0: (0.1 > 3.14159), F1: (v8 || (!v8)), F2: 930}
                                                                _ = v38
                                                            case "\n\t":
                                                                fmt.Println(v20, fi16)
                                                            default:
                                                                v39 := GenS5{F0: ("\n\t" + ("hello" + ""))}
                                                                _ = v39
                                                            }
                                                            _ = v37
                                                        }
                                                    }
                                                default:
                                                    if false {
                                                        if ((2 != 1) || false) {
                                                            switch "hello" {
                                                            case "a":
                                                                delete(v2, 2147483648)
                                                            default:
                                                                var v40 int64 = v6
                                                                _ = v40
                                                            }
                                                            v41 := GenS6{F0: 42}
                                                            _ = v41
                                                        } else {
                                                            v8 = (v20 > (v20 * float32(float64(v20))))
                                                            var v42 uint16 = (v4 - 1)
                                                            fmt.Println(v42)
                                                            _ = v42
                                                        }
                                                    } else {
                                                        var v43 uint32 = 2147483647
                                                        if len(v1) >= 2 {
                                                            v1 = v1[1:len(v1)]
                                                        }
                                                        if ("" <= (("" + "hello") + "foo bar")) {
                                                            if false {
                                                                if len(v3) >= 2 {
                                                                    v3 = v3[1:len(v3)]
                                                                }
                                                                v5 = v5
                                                            }
                                                            v21 += (v6 << ((v5) & 63))
                                                        }
                                                        switch "a" {
                                                        case "":
                                                            if v8 {
                                                                v44 := v2[(v43 * (v43 | 4294967295))]
                                                                if v8 {
                                                                    v8 = false
                                                                } else {
                                                                    var v45 bool = false
                                                                    if v8 {
                                                                        switch (("a" + "") + "") {
                                                                        case "hello":
                                                                            v4 -= v4
                                                                        case "":
                                                                            v46 := GenS7{F0: v4}
                                                                            _ = v46
                                                                        case "\n\t":
                                                                            fmt.Println(fi16, v21)
                                                                            v47 := GenS3{F0: -1}
                                                                            v2[v43] = fi7
                                                                            _ = v47
                                                                        }
                                                                        if v45 {
                                                                            var v48 float32 = (v20 - v20)
                                                                            if (251 > -9223372036854775808) {
                                                                                switch "foo bar" {
                                                                                case "":
                                                                                    switch "foo bar" {
                                                                                    case "hello":
                                                                                        if true {
                                                                                            if (-2.25 >= v48) {
                                                                                                if (v21 > -438) {
                                                                                                    if ((0.1 + -2.25) <= 1.5) {
                                                                                                        if v8 {
                                                                                                            var v49 uint = 693
                                                                                                            _ = v49
                                                                                                        }
                                                                                                        switch 0 {
                                                                                                        case -1:
                                                                                                            v43 -= (v43 | uint32(v5))
                                                                                                        case 42:
                                                                                                            var v50 float64 = ((0.0 + 1e10) / 1.5)
                                                                                                            _ = v50
                                                                                                        case 0:
                                                                                                            v51 := v2[(v43 & (v43 * (v43 / (((v43 & 505)) | 1))))]
                                                                                                            _ = v51
                                                                                                        }
                                                                                                    }
                                                                                                    v21 = (v21 / ((-1) | 1))
                                                                                                    var v52 int = (v44 ^ fi16)
                                                                                                    var v53 uint8 = 2
                                                                                                    _ = v52
                                                                                                    _ = v53
                                                                                                }
                                                                                            }
                                                                                        }
                                                                                    }
                                                                                case "hello":
                                                                                    switch v44 {
                                                                                    case 1:
                                                                                        switch "" {
                                                                                        case "foo bar":
                                                                                            var v54 string = "a"
                                                                                            _ = v54
                                                                                        case "hello":
                                                                                            var v55 float64 = ((3.14159 - (0.0 + (3.14159 / (-2.25 * (((1e10 - ((0.0 + ((0.1 + 1e10) * float64(-2.25))) + 1e10)) * ((float64(v20) - 1e10) - ((-2.25 + 1.5) - 1e10))) - (0.0 * (((0.1 + 0.0) - (0.0 * float64(v48))) / 3.14159))))))) * 1e10)
                                                                                            _ = v55
                                                                                        case "\n\t":
                                                                                            var v56 float32 = 0.1
                                                                                            _ = v56
                                                                                        }
                                                                                    default:
                                                                                        fmt.Println(v4)
                                                                                    }
                                                                                case "a":
                                                                                    v45 = true
                                                                                }
                                                                                fmt.Println(v4, v20, v21)
                                                                                v45 = ((0.1 < (3.14159 * 1.5)) || v45)
                                                                            }
                                                                            v20 += (v20 - v20)
                                                                            _ = v48
                                                                        }
                                                                        var v57 int = fi7
                                                                        switch (v57 & (^(fi7))) {
                                                                        case -1:
                                                                            v58 := GenS8{F0: ("a" + ("foo bar" + "a")), F1: (v57 + (v44 + -9223372036854775808)), F2: int32(v5)}
                                                                            _ = v58
                                                                        case 2147483647:
                                                                            v59 := GenS9{F0: true, F1: v44}
                                                                            _ = v59
                                                                        case 128:
                                                                            v60 := GenS3{F0: -670}
                                                                            _ = v60
                                                                        }
                                                                        _ = v57
                                                                    }
                                                                    if true {
                                                                        var v61 int16 = 42
                                                                        v43 = 1
                                                                        v62 := GenS9{F0: true, F1: 1}
                                                                        _ = v61
                                                                        _ = v62
                                                                    } else {
                                                                        if false {
                                                                            v43 = (v43 + 2)
                                                                            v63 := v2[v43]
                                                                            _ = v63
                                                                        }
                                                                        v45 = (!(1 == (v4 - (v4 * v4))))
                                                                        fmt.Println(v20, fi16)
                                                                        var v64 uint32 = (v43 / ((v43) | 1))
                                                                        var v65 uint64 = 305
                                                                        v66 := GenS1{F0: fi7}
                                                                        var v67 bool = v45
                                                                        v68 := GenS10{F0: 9223372036854775807}
                                                                        fmt.Println(v67, v43, v4)
                                                                        if (!v45) {
                                                                            switch "hello" {
                                                                            case "a":
                                                                                switch "a" {
                                                                                case "foo bar":
                                                                                    var v69 uint = 18446744073709551615
                                                                                    _ = v69
                                                                                case "\n\t":
                                                                                    v4 = (^(v4))
                                                                                }
                                                                                var v70 string = ""
                                                                                _ = v70
                                                                            }
                                                                            switch "" {
                                                                            case "":
                                                                                if ("\n\t" != "hello") {
                                                                                    var v71 uint64 = (v65 / (((v65 >> ((v5) & 63))) | 1))
                                                                                    switch "foo bar" {
                                                                                    case "a":
                                                                                        if false {
                                                                                            v72 := GenS5{F0: ("foo bar" + ("hello" + "\n\t"))}
                                                                                            var v73 float32 = float32(1e10)
                                                                                            _ = v72
                                                                                            _ = v73
                                                                                        } else {
                                                                                            v6 += -9223372036854775808
                                                                                        }
                                                                                    default:
                                                                                        v21 = (v21 + -628)
                                                                                    }
                                                                                    _ = len(v2)
                                                                                    _ = v71
                                                                                }
                                                                                fmt.Println(v4)
                                                                                var v74 int8 = 0
                                                                                _ = v74
                                                                            default:
                                                                                switch 9223372036854775807 {
                                                                                case -992:
                                                                                    var v75 uint = uint(v64)
                                                                                    _ = v75
                                                                                default:
                                                                                    if v67 {
                                                                                        if (true && true) {
                                                                                            switch "hello" {
                                                                                            case "\n\t":
                                                                                                if len(v1) > 0 {
                                                                                                    v1[0] = (v43 >> ((42) & 63))
                                                                                                }
                                                                                                var v76 string = "a"
                                                                                                _ = len(v2)
                                                                                                _ = v76
                                                                                            }
                                                                                            var v77 uint32 = (^(v43))
                                                                                            _ = v77
                                                                                        }
                                                                                        var v78 float32 = ((1e10 * -2.25) * v20)
                                                                                        _ = v78
                                                                                    }
                                                                                }
                                                                            }
                                                                            var v79 int64 = v6
                                                                            v4 += (^(v4))
                                                                            _ = len(v2)
                                                                            _ = v79
                                                                        }
                                                                        _ = v64
                                                                        _ = v65
                                                                        _ = v66
                                                                        _ = v67
                                                                        _ = v68
                                                                    }
                                                                    _ = v45
                                                                }
                                                                _ = v44
                                                            } else {
                                                                v80 := GenS11{F0: (^(v4)), F1: -128}
                                                                _ = v80
                                                            }
                                                        }
                                                        var v81 int = -9223372036854775808
                                                        if (!((!(fi7 > 129)) && (0 != 151))) {
                                                            if len(v1) >= 2 {
                                                                v1 = v1[1:len(v1)]
                                                            }
                                                        }
                                                        if false {
                                                            switch -9223372036854775808 {
                                                            case -619:
                                                                if v8 {
                                                                    v82 := GenS5{F0: ""}
                                                                    v83 := v2[0]
                                                                    if v8 {
                                                                        if len(v3) > 0 {
                                                                            v3[0] = ("\n\t" + (("\n\t" + ("a" + ("" + "foo bar"))) + ""))
                                                                        }
                                                                    } else {
                                                                        v84 := GenS11{F0: v4, F1: -63}
                                                                        v20 += ((float32((0.1 * 1e10)) * v20) - (-2.25 - 3.14159))
                                                                        v43 -= (v43 % (((v43 * 42)) | 1))
                                                                        _ = v84
                                                                    }
                                                                    _ = v82
                                                                    _ = v83
                                                                }
                                                            case 1:
                                                                if len(v1) > 0 {
                                                                    v1[0] = 65535
                                                                }
                                                            case -9223372036854775808:
                                                                var v85 uint = v5
                                                                _ = v85
                                                            default:
                                                                var v86 uint16 = (v4 % ((v4) | 1))
                                                                _ = v86
                                                            }
                                                            var v87 int8 = 96
                                                            switch (-(v81)) {
                                                            case -2147483648:
                                                                v6 = v6
                                                            default:
                                                                v88 := GenS10{F0: (v21 >> ((v5) & 63))}
                                                                _ = v88
                                                            }
                                                            _ = v87
                                                        }
                                                        v89 := GenS1{F0: fi16}
                                                        if false {
                                                            if v8 {
                                                                fmt.Println(v81, v43, fi7)
                                                                fmt.Println(fi7, v4, v6)
                                                                switch -840 {
                                                                case 342:
                                                                    if (false && v8) {
                                                                        var v90 string = "foo bar"
                                                                        _ = v90
                                                                    } else {
                                                                        v2[1] = 128
                                                                    }
                                                                default:
                                                                    switch "hello" {
                                                                    case "a":
                                                                        if v8 {
                                                                            if (!(42 >= -651)) {
                                                                                v21 = v21
                                                                                if len(v1) > 0 {
                                                                                    v1[0] = (v43 << (((v5 % ((900) | 1))) & 63))
                                                                                }
                                                                            } else {
                                                                                var v91 bool = v8
                                                                                var v92 uint = v5
                                                                                _ = v91
                                                                                _ = v92
                                                                            }
                                                                            if len(v1) > 0 {
                                                                                v1[0] = v43
                                                                            }
                                                                            var v93 uint64 = 1
                                                                            v6 = (v6 ^ (v21 ^ v6))
                                                                            switch fi7 {
                                                                            case 127:
                                                                                fmt.Println(v81, v81, v21)
                                                                            }
                                                                            v21 = 1
                                                                            _ = v93
                                                                        } else {
                                                                            var v94 uint64 = 1
                                                                            _ = len(v2)
                                                                            v95 := GenS5{F0: ""}
                                                                            _ = v94
                                                                            _ = v95
                                                                        }
                                                                    default:
                                                                        v96 := GenS12{F0: "", F1: 646, F2: 214}
                                                                        _ = v96
                                                                    }
                                                                }
                                                            }
                                                            v81 = v81
                                                        } else {
                                                            if false {
                                                                v4 = (v4 % ((v4) | 1))
                                                                var v97 uint8 = 2
                                                                var v98 uint64 = 0
                                                                v8 = (1 > int32(v98))
                                                                if (!true) {
                                                                    v81 = (^(fi16))
                                                                    var v99 uint32 = v43
                                                                    v100 := GenS7{F0: v4}
                                                                    v4 = v4
                                                                    _ = v100
                                                                    _ = v99
                                                                } else {
                                                                    delete(v2, v43)
                                                                    v101 := GenS2{F0: v43, F1: "\n\t", F2: (!false)}
                                                                    v8 = false
                                                                    _ = v101
                                                                }
                                                                _ = v97
                                                                _ = v98
                                                            }
                                                        }
                                                        var v102 int64 = (v21 / ((421) | 1))
                                                        var v103 uint32 = 2147483647
                                                        _ = v102
                                                        _ = v103
                                                        _ = v43
                                                        _ = v81
                                                        _ = v89
                                                    }
                                                }
                                                fmt.Println(fi7, v20, v20)
                                            }
                                            _ = v21
                                        }
                                        v4 = 32767
                                        if v8 {
                                            v104 := GenS3{F0: -270}
                                            _ = len(v2)
                                            switch (fi7 % (((fi16 + fi7)) | 1)) {
                                            case 2147483647:
                                                var v105 uint64 = 903
                                                _ = v105
                                            case 127:
                                                var v106 uint64 = 2
                                                _ = v106
                                            case 232:
                                                var v107 uint32 = 42
                                                _ = v107
                                            default:
                                                v108 := GenS13{F0: -834}
                                                _ = v108
                                            }
                                            _ = v104
                                        } else {
                                            var v109 int64 = (^(v6))
                                            _ = v109
                                        }
                                        _ = v20
                                    }
                                }
                                var v110 bool = (v8 && true)
                                _ = v110
                            }
                        }
                        v111 := GenS14{F0: (v6 + -1)}
                        _ = fi16
                        _ = v111
                    }
                    v6 = v6
                }
                var v112 string = "\n\t"
                _ = v112
                _ = v8
            }
        case "a":
            var v113 uint32 = 324
            _ = v113
        }
        _ = fi7
    }
    v6 -= (v6 - (-(v6)))
    var v114 int = 264
    switch "" {
    case "a":
        var v115 int8 = 127
        _ = v115
    case "":
        v116 := GenS15{F0: false}
        _ = v116
    default:
        v5 = (v5 | v5)
    }
    var v117 int64 = -9223372036854775808
    if len(v3) > 0 {
        v3[0] = "a"
    }
    switch "\n\t" {
    case "\n\t":
        v4 -= v4
    }
    for fi118 := 0; fi118 < 5; fi118++ {
        v5 -= uint(v4)
        switch (^(v114)) {
        case 4294967295:
            fmt.Println(v114, v117)
        default:
            if false {
                _ = len(v1)
                for fi119 := 0; fi119 < 9; fi119++ {
                    v5 -= (^(v5))
                    fmt.Println(fi118, v5, v6)
                    _ = fi119
                }
                fmt.Println(v5, v5, fi118)
            } else {
                v4 = (v4 ^ 771)
            }
        }
        v114 += (fi118 - (v114 << (((v5 >> (((^(v5))) & 63))) & 63)))
        _ = fi118
    }
    v6 = 0
    v4 = (v4 * 65535)
    v2[653] = 9223372036854775807
    var v120 int16 = -745
    for fi121 := 0; fi121 < 12; fi121++ {
        v6 -= (v117 + 1)
        _ = len(v1)
        var v122 int64 = (v6 << ((65535) & 63))
        _ = fi121
        _ = v122
    }
    v123 := GenS12{F0: "foo bar", F1: v6, F2: 1}
    var v124 int = int(v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v114:", v114)
    fmt.Println("v117:", v117)
    fmt.Println("v120:", v120)
    fmt.Println("v124:", v124)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v114
    _ = v117
    _ = v120
    _ = v123
    _ = v124
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
