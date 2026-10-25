package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 string }
type GenS2 struct { F0 uint; F1 uint8 }
type GenS3 struct { F0 string; F1 string; F2 string }
type GenS4 struct { F0 uint8; F1 string }
type GenS5 struct { F0 uint64; F1 int8 }
type GenS6 struct { F0 int8; F1 uint64 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 uint16; F1 uint64 }
type GenS9 struct { F0 uint32; F1 int8; F2 uint64 }
type GenS10 struct { F0 bool; F1 int16; F2 uint8 }
type GenS11 struct { F0 int8; F1 string }
type GenS12 struct { F0 int; F1 uint8 }
type GenS13 struct { F0 uint16; F1 bool }
type GenS14 struct { F0 int; F1 int }
type GenS15 struct { F0 uint8; F1 string }
type GenS16 struct { F0 int8; F1 uint8 }
type GenS17 struct { F0 uint8 }

func main() {
    v1 := make(map[string]uint32)
    var v2 int16 = 320
    var v3 bool = (true || false)
    var v4 uint16 = 65535
    var v5 uint32 = 0
    var v6 uint = 0
    switch ("foo bar" + ("\n\t" + "\n\t")) {
    case "foo bar":
        var v7 int = 42
        _ = v7
    case "hello":
        v8 := GenS1{F0: 32768, F1: "hello"}
        _ = v8
    default:
        var v9 int32 = 42
        _ = v9
    }
    switch 9223372036854775807 {
    case -1:
        if false {
            fmt.Println(v6, v6)
            fmt.Println(v2)
            v10 := v1["hello"]
            _ = v10
        } else {
            if (0 <= -1) {
                v11 := GenS2{F0: (v6 | (v6 / (((v6 + 4294967295)) | 1))), F1: 42}
                if v3 {
                    var v12 uint = (v6 - (^(v6)))
                    var v13 uint8 = 0
                    _ = v12
                    _ = v13
                } else {
                    switch 4294967295 {
                    case 9223372036854775807:
                        v5 = v5
                    default:
                        v3 = (true || false)
                    }
                }
                for fi14 := 0; fi14 < 11; fi14++ {
                    var v15 int64 = 1
                    if (!true) {
                        fmt.Println(v2, v3, v3)
                        v16 := []int{}
                        for fi17 := 0; fi17 < 4; fi17++ {
                            if ((42 > v15) || v3) {
                                if false {
                                    if true {
                                        v3 = (4294967295 > 423)
                                    } else {
                                        var v18 string = "hello"
                                        _ = v18
                                    }
                                    v3 = (uint64(v4) == 1)
                                } else {
                                    var v19 float32 = ((((1.5 * 3.14159) + float32((1e10 * 1e10))) * (1e10 / 1.0)) / 1e10)
                                    _ = v19
                                }
                                v2 = (v2 - 0)
                                var v20 int16 = v2
                                if (v3 || false) {
                                    switch (fi14 << ((v6) & 63)) {
                                    case -1:
                                        v3 = ((0.0 <= 0.1) && true)
                                    case -745:
                                        v20 += v20
                                    default:
                                        var v21 uint8 = 127
                                        _ = v21
                                    }
                                } else {
                                    if len(v16) >= 2 {
                                        v16 = v16[1:len(v16)]
                                    }
                                    var v22 uint32 = (v5 - 872)
                                    _ = v22
                                }
                                switch "foo bar" {
                                case "foo bar":
                                    v4 = v4
                                case "hello":
                                    v20 -= v20
                                case "a":
                                    var v23 int8 = 42
                                    v3 = true
                                    _ = v23
                                }
                                if len(v16) >= 2 {
                                    v16 = v16[1:len(v16)]
                                }
                                if v3 {
                                    v24 := GenS2{F0: v6, F1: 162}
                                    var v25 uint8 = 188
                                    v4 += (^(v4))
                                    _ = v24
                                    _ = v25
                                }
                                _ = v20
                            }
                            var v26 uint64 = 502
                            _ = fi17
                            _ = v26
                        }
                        _ = v16
                    } else {
                        var v27 int8 = 100
                        v4 -= v4
                        fmt.Println(v4)
                        _ = v27
                    }
                    for fi28 := 0; fi28 < 16; fi28++ {
                        if v3 {
                            switch fi14 {
                            case 9223372036854775807:
                                v15 += v15
                                if ((v3 && v3) || (((^(v15)) < (v15 << ((1) & 63))) && v3)) {
                                    v4 = (v4 - (v4 | (v4 << ((v6) & 63))))
                                    switch fi28 {
                                    case 32768:
                                        if false {
                                            switch ("foo bar" + (("foo bar" + "a") + (("a" + "") + ((("" + "a") + ("a" + "a")) + "\n\t")))) {
                                            case "foo bar":
                                                var v29 float32 = -2.25
                                                _ = v29
                                            case "\n\t":
                                                if v3 {
                                                    v5 -= v5
                                                    v4 += 42
                                                    v4 = (v4 << ((uint(v4)) & 63))
                                                    v2 += (v2 - (v2 - v2))
                                                }
                                            case "a":
                                                v3 = v3
                                            default:
                                                if (0.1 <= 0.0) {
                                                    switch "" {
                                                    case "foo bar":
                                                        v30 := []int{fi28}
                                                        _ = v30
                                                    case "hello":
                                                        if ((v15 | 247) <= v15) {
                                                            v31 := GenS2{F0: (v6 & uint(v4)), F1: uint8(v4)}
                                                            var v32 uint16 = v4
                                                            var v33 uint64 = 2
                                                            v34 := GenS3{F0: "a", F1: "\n\t", F2: ""}
                                                            var v35 string = "hello"
                                                            switch (fi28 | 128) {
                                                            case 2147483647:
                                                                if (v2 != 80) {
                                                                    fmt.Println(v3, v15)
                                                                    var v36 int16 = (v2 + 1)
                                                                    if v3 {
                                                                        if v3 {
                                                                            var v37 bool = (!false)
                                                                            _ = v37
                                                                        } else {
                                                                            var v38 uint64 = v33
                                                                            if (128 == 32768) {
                                                                                v2 = v2
                                                                                v39 := v1[(("\n\t" + "a") + ("foo bar" + "hello"))]
                                                                                if true {
                                                                                    fmt.Println(v15, v15)
                                                                                    v40 := []int{128, fi14}
                                                                                    switch 9223372036854775807 {
                                                                                    case 32767:
                                                                                        _ = len(v1)
                                                                                    case 128:
                                                                                        var v41 string = "foo bar"
                                                                                        _ = v41
                                                                                    case 42:
                                                                                        var v42 int64 = (v15 / (((^(v15))) | 1))
                                                                                        _ = v42
                                                                                    }
                                                                                    switch 4294967295 {
                                                                                    case 0:
                                                                                        if v3 {
                                                                                            var v43 bool = v3
                                                                                            _ = v43
                                                                                        } else {
                                                                                            v2 = (v36 / (((v2 ^ (-(v2)))) | 1))
                                                                                            v3 = ((!((!false) && ((v3 || true) && (!((v35 + v35) > (v35 + v35)))))) && ((fi28 - fi14) >= 4294967295))
                                                                                        }
                                                                                        if len(v40) >= 2 {
                                                                                            v40 = v40[1:len(v40)]
                                                                                        }
                                                                                        fmt.Println(fi14)
                                                                                        if (42 < 0) {
                                                                                            fmt.Println(v36)
                                                                                            fmt.Println(fi28, v2, v32)
                                                                                            var v44 uint = v6
                                                                                            fmt.Println(v35, v44)
                                                                                            switch fi28 {
                                                                                            case 128:
                                                                                                var v45 float64 = 1e10
                                                                                                _ = v45
                                                                                            case 32768:
                                                                                                _ = len(v1)
                                                                                            default:
                                                                                                if true {
                                                                                                    var v46 uint = (v44 | 248)
                                                                                                    _ = v46
                                                                                                } else {
                                                                                                    _ = len(v40)
                                                                                                    v4 = (v4 >> ((v44) & 63))
                                                                                                }
                                                                                            }
                                                                                            _ = v44
                                                                                        }
                                                                                    case -2147483648:
                                                                                        switch "" {
                                                                                        case "foo bar":
                                                                                            fmt.Println(v38)
                                                                                        case "hello":
                                                                                            switch (fi14 << ((v6) & 63)) {
                                                                                            case 846:
                                                                                                var v47 uint16 = v4
                                                                                                _ = v47
                                                                                            case 818:
                                                                                                if ((!(!((v32 % (((v32 * v32)) | 1)) >= v32))) && v3) {
                                                                                                    v39 += 576
                                                                                                    v40 = append(v40, (fi28 % ((fi28) | 1)))
                                                                                                    v35 = v35
                                                                                                }
                                                                                            case -1:
                                                                                                switch int(v15) {
                                                                                                case -394:
                                                                                                    fmt.Println(v5, v15, v33)
                                                                                                }
                                                                                            }
                                                                                        case "\n\t":
                                                                                            var v48 uint8 = 1
                                                                                            _ = v48
                                                                                        default:
                                                                                            var v49 float32 = 0.0
                                                                                            var v50 uint = uint(v38)
                                                                                            _ = v49
                                                                                            _ = v50
                                                                                        }
                                                                                    }
                                                                                    if len(v40) >= 2 {
                                                                                        v40 = v40[1:len(v40)]
                                                                                    }
                                                                                    _ = v40
                                                                                }
                                                                                _ = v39
                                                                            } else {
                                                                                var v51 float32 = 0.0
                                                                                _ = v51
                                                                            }
                                                                            var v52 string = (v35 + ("foo bar" + "hello"))
                                                                            _ = len(v1)
                                                                            _ = v38
                                                                            _ = v52
                                                                        }
                                                                        fmt.Println(v15, v36, v36)
                                                                    }
                                                                    _ = v36
                                                                }
                                                            case 300:
                                                                if v3 {
                                                                    var v53 string = "foo bar"
                                                                    var v54 uint8 = 22
                                                                    _ = v53
                                                                    _ = v54
                                                                } else {
                                                                    var v55 uint8 = 42
                                                                    v56 := GenS4{F0: v55, F1: "a"}
                                                                    v1[("a" + v35)] = uint32(v32)
                                                                    _ = v55
                                                                    _ = v56
                                                                }
                                                            default:
                                                                v57 := GenS2{F0: (^(v6)), F1: 128}
                                                                _ = v57
                                                            }
                                                            _ = v31
                                                            _ = v32
                                                            _ = v33
                                                            _ = v34
                                                            _ = v35
                                                        }
                                                    case "a":
                                                        var v58 int = fi14
                                                        _ = v58
                                                    default:
                                                        var v59 int64 = v15
                                                        _ = v59
                                                    }
                                                    if v3 {
                                                        switch "hello" {
                                                        case "":
                                                            v6 = 458
                                                        case "a":
                                                            var v60 int8 = 127
                                                            _ = v60
                                                        case "foo bar":
                                                            if (9223372036854775807 < (fi14 | int(v15))) {
                                                                var v61 int8 = 0
                                                                v4 += v4
                                                                _ = v61
                                                            } else {
                                                                v5 = v5
                                                                v3 = ((v3 && true) && (226 == uint16(v5)))
                                                            }
                                                        default:
                                                            switch (fi28 << ((v6) & 63)) {
                                                            case -2147483648:
                                                                v62 := []int{42, -1, 477, 9223372036854775807}
                                                                _ = v62
                                                            case 32768:
                                                                v4 += 2
                                                            case 299:
                                                                switch fi28 {
                                                                case -1:
                                                                    var v63 bool = false
                                                                    _ = v63
                                                                case -9223372036854775808:
                                                                    var v64 int8 = 1
                                                                    _ = v64
                                                                case 451:
                                                                    var v65 uint8 = 245
                                                                    _ = v65
                                                                default:
                                                                    fmt.Println(v15, v2)
                                                                }
                                                            default:
                                                                if ((-1 != -2147483648) || v3) {
                                                                    v66 := []int{(fi14 % (((fi14 & fi28)) | 1)), (fi28 - fi28)}
                                                                    if ((v4 % ((v4) | 1)) > v4) {
                                                                        v67 := GenS5{F0: 2, F1: -49}
                                                                        if true {
                                                                            v68 := GenS4{F0: 2, F1: "a"}
                                                                            _ = v68
                                                                        }
                                                                        v15 += (v15 + v15)
                                                                        switch "" {
                                                                        case "hello":
                                                                            switch "" {
                                                                            case "foo bar":
                                                                                var v69 string = ("a" + "foo bar")
                                                                                _ = v69
                                                                            case "hello":
                                                                                v6 = 4294967295
                                                                            }
                                                                        default:
                                                                            switch "" {
                                                                            case "\n\t":
                                                                                delete(v1, "\n\t")
                                                                            case "foo bar":
                                                                                fmt.Println(fi14)
                                                                            case "":
                                                                                if ((v3 || v3) || v3) {
                                                                                    if v3 {
                                                                                        var v70 int32 = 42
                                                                                        _ = v70
                                                                                    } else {
                                                                                        v15 += v15
                                                                                    }
                                                                                } else {
                                                                                    v4 -= (v4 | v4)
                                                                                }
                                                                            default:
                                                                                if true {
                                                                                    switch "hello" {
                                                                                    case "hello":
                                                                                        fmt.Println(fi28, v5, v6)
                                                                                    case "a":
                                                                                        v71 := v1["foo bar"]
                                                                                        _ = v71
                                                                                    case "\n\t":
                                                                                        switch ((((((("a" + "foo bar") + "") + ("\n\t" + "\n\t")) + "hello") + (((("hello" + "hello") + "\n\t") + ("hello" + "\n\t")) + "")) + "foo bar") + "a") {
                                                                                        case "hello":
                                                                                            if (((("hello" + "hello") + ("hello" + (("a" + "a") + "\n\t"))) + ((("\n\t" + ("a" + "\n\t")) + "hello") + "a")) != "foo bar") {
                                                                                                _ = len(v66)
                                                                                            } else {
                                                                                                var v72 int = fi28
                                                                                                _ = v72
                                                                                            }
                                                                                        case "a":
                                                                                            var v73 int16 = -32768
                                                                                            _ = v73
                                                                                        }
                                                                                    }
                                                                                }
                                                                            }
                                                                        }
                                                                        v2 -= v2
                                                                        var v74 uint = 18446744073709551615
                                                                        _ = v67
                                                                        _ = v74
                                                                    }
                                                                    var v75 uint64 = 12
                                                                    _ = v66
                                                                    _ = v75
                                                                } else {
                                                                    if v3 {
                                                                        switch "" {
                                                                        case "":
                                                                            v76 := []int{}
                                                                            _ = v76
                                                                        case "hello":
                                                                            switch fi14 {
                                                                            case 127:
                                                                                var v77 int = fi14
                                                                                _ = v77
                                                                            case -716:
                                                                                v6 = 703
                                                                            case 4294967295:
                                                                                fmt.Println(v3, v5, v15)
                                                                            }
                                                                            v3 = (uint64(v4) >= 116)
                                                                            var v78 float64 = 0.1
                                                                            _ = v78
                                                                        }
                                                                        v2 = (v2 / ((v2) | 1))
                                                                        switch "a" {
                                                                        case "foo bar":
                                                                            if (v5 == (v5 | (v5 * v5))) {
                                                                                switch "a" {
                                                                                case "hello":
                                                                                    v5 = 2147483648
                                                                                case "\n\t":
                                                                                    v1[("" + "\n\t")] = v5
                                                                                default:
                                                                                    var v79 uint16 = 584
                                                                                    _ = v79
                                                                                }
                                                                                if ((^(v4)) == v4) {
                                                                                    if true {
                                                                                        if false {
                                                                                            v4 -= (v4 % ((v4) | 1))
                                                                                            var v80 int32 = -720
                                                                                            _ = v80
                                                                                        } else {
                                                                                            var v81 uint32 = uint32(v6)
                                                                                            v4 -= (^(v4))
                                                                                            v6 = (v6 << ((v6) & 63))
                                                                                            _ = v81
                                                                                        }
                                                                                    }
                                                                                } else {
                                                                                    if (-166 != 1) {
                                                                                        var v82 float32 = 0.0
                                                                                        if true {
                                                                                            var v83 float32 = (v82 - float32(((((float64(3.14159) + (float64(v82) / ((3.14159 / 3.14159) / 1e10))) + float64(v82)) - 1.5) - (float64(-2.25) / (0.1 / (float64(0.1) - (1e10 * 0.0)))))))
                                                                                            var v84 uint16 = 2
                                                                                            v85 := GenS6{F0: 57, F1: 2}
                                                                                            _ = v83
                                                                                            _ = v84
                                                                                            _ = v85
                                                                                        }
                                                                                        _ = v82
                                                                                    } else {
                                                                                        var v86 uint8 = 128
                                                                                        var v87 string = "\n\t"
                                                                                        _ = v86
                                                                                        _ = v87
                                                                                    }
                                                                                    if v3 {
                                                                                        v88 := GenS1{F0: (fi14 & fi14), F1: "\n\t"}
                                                                                        _ = v88
                                                                                    }
                                                                                    var v89 float32 = 0.0
                                                                                    v3 = v3
                                                                                    var v90 uint32 = 806
                                                                                    v91 := GenS7{F0: (v2 & 973)}
                                                                                    _ = v89
                                                                                    _ = v90
                                                                                    _ = v91
                                                                                }
                                                                                v92 := GenS6{F0: -128, F1: 2}
                                                                                _ = v92
                                                                            } else {
                                                                                switch fi14 {
                                                                                case 4294967295:
                                                                                    var v93 uint8 = 33
                                                                                    _ = v93
                                                                                case 127:
                                                                                    v94 := GenS4{F0: 2, F1: ("a" + (("a" + ((("hello" + "foo bar") + ("foo bar" + "")) + ("hello" + ("hello" + "\n\t")))) + "hello"))}
                                                                                    _ = v94
                                                                                case 128:
                                                                                    switch (fi28 % ((-1) | 1)) {
                                                                                    case 845:
                                                                                        v4 = v4
                                                                                    case 127:
                                                                                        var v95 int32 = 2147483647
                                                                                        _ = v95
                                                                                    case 967:
                                                                                        fmt.Println(v3, fi14, v2)
                                                                                    default:
                                                                                        v5 = v5
                                                                                    }
                                                                                }
                                                                                v5 = v5
                                                                                var v96 bool = true
                                                                                v4 += v4
                                                                                var v97 float64 = 3.14159
                                                                                v5 = 2147483648
                                                                                _ = v96
                                                                                _ = v97
                                                                            }
                                                                        case "a":
                                                                            if false {
                                                                                if ((^(v6)) < v6) {
                                                                                    var v98 float32 = 0.0
                                                                                    v99 := GenS1{F0: (fi14 ^ fi28), F1: "a"}
                                                                                    _ = v98
                                                                                    _ = v99
                                                                                } else {
                                                                                    fmt.Println(v15, v5)
                                                                                    v100 := v1["foo bar"]
                                                                                    var v101 bool = v3
                                                                                    _ = v100
                                                                                    _ = v101
                                                                                }
                                                                            }
                                                                        }
                                                                    } else {
                                                                        var v102 int32 = 2147483647
                                                                        var v103 int = (-(fi28))
                                                                        if v3 {
                                                                            if (-1 == (v2 % ((1000) | 1))) {
                                                                                switch "" {
                                                                                case "":
                                                                                    fmt.Println(v5, fi14)
                                                                                case "a":
                                                                                    var v104 uint16 = 751
                                                                                    _ = v104
                                                                                }
                                                                                v105 := v1[""]
                                                                                _ = v105
                                                                            } else {
                                                                                var v106 int = -1
                                                                                _ = v106
                                                                            }
                                                                            v107 := v1["\n\t"]
                                                                            _ = v107
                                                                        } else {
                                                                            v2 -= v2
                                                                            switch fi14 {
                                                                            case 9223372036854775807:
                                                                                switch 128 {
                                                                                case 181:
                                                                                    var v108 int = (fi28 * fi28)
                                                                                    _ = v108
                                                                                case -1:
                                                                                    v109 := GenS2{F0: v6, F1: 127}
                                                                                    _ = v109
                                                                                }
                                                                            case 783:
                                                                                switch "foo bar" {
                                                                                case "\n\t":
                                                                                    switch -9223372036854775808 {
                                                                                    case 32767:
                                                                                        v102 = (v102 + (v102 << (((v6 + v6)) & 63)))
                                                                                    case -2147483648:
                                                                                        v6 = (v6 << ((v6) & 63))
                                                                                    default:
                                                                                        v110 := []int{(v103 + 4294967295), (-(fi14)), 2147483647, (-(fi28))}
                                                                                        _ = v110
                                                                                    }
                                                                                case "":
                                                                                    var v111 float64 = (float64(3.14159) - 1e10)
                                                                                    _ = v111
                                                                                default:
                                                                                    var v112 uint32 = (v5 - (^(v5)))
                                                                                    _ = v112
                                                                                }
                                                                            default:
                                                                                var v113 uint = (v6 - v6)
                                                                                _ = v113
                                                                            }
                                                                            _ = len(v1)
                                                                        }
                                                                        _ = v102
                                                                        _ = v103
                                                                    }
                                                                }
                                                            }
                                                        }
                                                    }
                                                    v5 = (v5 ^ 675)
                                                } else {
                                                    switch (("foo bar" + "\n\t") + "foo bar") {
                                                    case "foo bar":
                                                        v114 := GenS8{F0: v4, F1: uint64(v4)}
                                                        _ = v114
                                                    case "hello":
                                                        v4 = v4
                                                    case "a":
                                                        v6 = uint(v5)
                                                    default:
                                                        v115 := v1["\n\t"]
                                                        _ = v115
                                                    }
                                                    switch (fi28 ^ fi28) {
                                                    case -696:
                                                        var v116 uint16 = 0
                                                        _ = v116
                                                    }
                                                }
                                            }
                                            fmt.Println(fi28, v3)
                                            delete(v1, (("foo bar" + "hello") + ("" + ("a" + "hello"))))
                                        } else {
                                            var v117 uint8 = 42
                                            switch "" {
                                            case "hello":
                                                fmt.Println(fi28, v117)
                                            case "a":
                                                fmt.Println(v117, v3)
                                                v118 := []int{fi28, 1}
                                                _ = v118
                                            case "foo bar":
                                                v119 := GenS4{F0: uint8(v5), F1: "hello"}
                                                _ = v119
                                            }
                                            v120 := []int{}
                                            _ = v117
                                            _ = v120
                                        }
                                    case -403:
                                        v15 += 0
                                    }
                                    var v121 string = ""
                                    if (v2 <= -1) {
                                        v122 := GenS9{F0: v5, F1: -128, F2: 940}
                                        _ = v122
                                    } else {
                                        var v123 uint8 = 93
                                        var v124 int = 9223372036854775807
                                        if (v2 <= (v2 >> (((v6 | (v6 | (v6 % ((v6) | 1))))) & 63))) {
                                            v123 -= v123
                                            v121 = ""
                                        }
                                        v4 = (v4 + (v4 * v4))
                                        if false {
                                            v125 := GenS10{F0: v3, F1: -32768, F2: v123}
                                            switch (fi28 >> (((v6 & v6)) & 63)) {
                                            case 32768:
                                                fmt.Println(v5, v123)
                                            case 239:
                                                var v126 int16 = (v2 ^ v2)
                                                _ = v126
                                            default:
                                                switch "\n\t" {
                                                case "":
                                                    switch fi28 {
                                                    case -738:
                                                        fmt.Println(v3, v123, v2)
                                                    }
                                                }
                                                var v127 int8 = 100
                                                var v128 uint = 87
                                                _ = v127
                                                _ = v128
                                            }
                                            v129 := GenS11{F0: -104, F1: v121}
                                            v2 = (v2 << (((v6 * uint(v4))) & 63))
                                            v130 := []int{(fi14 >> (((v6 << ((v6) & 63))) & 63)), 925, fi14}
                                            var v131 bool = false
                                            _ = v125
                                            _ = v129
                                            _ = v130
                                            _ = v131
                                        } else {
                                            if v3 {
                                                switch (v121 + v121) {
                                                case "hello":
                                                    var v132 int8 = 0
                                                    v1[v121] = (^(v5))
                                                    _ = v132
                                                case "a":
                                                    if v3 {
                                                        var v133 uint16 = v4
                                                        v4 -= v133
                                                        var v134 float32 = (1.5 / (0.1 * (1.5 + 3.14159)))
                                                        _ = v133
                                                        _ = v134
                                                    }
                                                case "\n\t":
                                                    switch ((v121 + v121) + v121) {
                                                    case "hello":
                                                        switch (fi28 / (((v124 / (((fi28 - 233)) | 1))) | 1)) {
                                                        case -1:
                                                            if v3 {
                                                                var v135 int64 = -155
                                                                _ = v135
                                                            }
                                                        default:
                                                            switch v121 {
                                                            case "a":
                                                                switch ((("hello" + (("" + v121) + (v121 + ("" + "foo bar")))) + "foo bar") + ((v121 + (((("" + ("foo bar" + "hello")) + v121) + v121) + v121)) + v121)) {
                                                                case "a":
                                                                    v121 = ""
                                                                    v3 = v3
                                                                    v136 := GenS12{F0: (fi28 / (((fi28 & (fi14 - -709))) | 1)), F1: 102}
                                                                    _ = v136
                                                                case "\n\t":
                                                                    fmt.Println(fi14, v124, fi28)
                                                                default:
                                                                    var v137 float32 = 1e10
                                                                    _ = v137
                                                                }
                                                                v121 = v121
                                                            case "foo bar":
                                                                v138 := GenS13{F0: 1, F1: v3}
                                                                switch v121 {
                                                                case "a":
                                                                    var v139 bool = v3
                                                                    _ = v139
                                                                case "hello":
                                                                    var v140 uint8 = (v123 & v123)
                                                                    _ = v140
                                                                case "\n\t":
                                                                    var v141 uint64 = 4294967295
                                                                    _ = v141
                                                                default:
                                                                    v142 := []int{fi28, fi14, fi14, fi28}
                                                                    _ = v142
                                                                }
                                                                switch v121 {
                                                                case "":
                                                                    delete(v1, (("\n\t" + "hello") + ((((v121 + v121) + v121) + v121) + ("hello" + ((v121 + "hello") + (v121 + ("hello" + (((v121 + ("a" + v121)) + (v121 + v121)) + v121))))))))
                                                                    if (!((^(v6)) > 2)) {
                                                                        var v143 uint64 = 71
                                                                        if (v3 || (!(uint(v143) >= v6))) {
                                                                            var v144 int32 = -149
                                                                            v143 += v143
                                                                            _ = v144
                                                                        }
                                                                        _ = v143
                                                                    } else {
                                                                        var v145 float32 = ((float32(1e10) - (3.14159 * (((((0.1 - (1.5 * (0.0 - 1.5))) / float32((-2.25 * 0.1))) * float32(((float64((0.1 / 3.14159)) / 1e10) * ((-2.25 - (1.5 + (1e10 - 0.1))) + 0.0)))) + (1e10 * (1e10 * 1e10))) / 1.0))) * (0.0 * ((0.0 - float32(-2.25)) * 1e10)))
                                                                        switch (fi28 << ((uint(v123)) & 63)) {
                                                                        case 128:
                                                                            var v146 uint8 = (^(v123))
                                                                            _ = v146
                                                                        case -309:
                                                                            var v147 float64 = 3.14159
                                                                            _ = v147
                                                                        case 32768:
                                                                            v148 := GenS1{F0: fi14, F1: (("" + v121) + v121)}
                                                                            _ = v148
                                                                        default:
                                                                            fmt.Println(v2, v15)
                                                                        }
                                                                        _ = v145
                                                                    }
                                                                case "hello":
                                                                    v149 := []int{(fi28 ^ -9223372036854775808), v124, int(v5)}
                                                                    _ = v149
                                                                default:
                                                                    v3 = (!v3)
                                                                }
                                                                _ = v138
                                                            default:
                                                                var v150 string = v121
                                                                _ = v150
                                                            }
                                                        }
                                                    default:
                                                        v151 := GenS3{F0: (v121 + (((v121 + ((((v121 + v121) + "foo bar") + "a") + (v121 + (((v121 + v121) + "hello") + ((("foo bar" + v121) + v121) + ("" + v121)))))) + v121) + "\n\t")), F1: (("hello" + "hello") + ""), F2: v121}
                                                        _ = v151
                                                    }
                                                }
                                                var v152 int32 = 2147483647
                                                _ = v152
                                            } else {
                                                if false {
                                                    v153 := v1[((v121 + (v121 + (v121 + v121))) + v121)]
                                                    if (v15 == v15) {
                                                        v154 := []int{}
                                                        _ = v154
                                                    }
                                                    _ = v153
                                                }
                                                switch (fi14 % ((fi14) | 1)) {
                                                case 127:
                                                    v6 = (v6 / (((v6 * v6)) | 1))
                                                default:
                                                    switch v121 {
                                                    case "foo bar":
                                                        fmt.Println(v3, v5, v123)
                                                    case "\n\t":
                                                        if ((v121 + (v121 + "")) > v121) {
                                                            var v155 uint64 = 18446744073709551615
                                                            _ = v155
                                                        }
                                                    }
                                                }
                                            }
                                        }
                                        v156 := []int{fi14, fi28, (fi14 * -2147483648), (fi14 | (fi14 - 128))}
                                        _ = v123
                                        _ = v124
                                        _ = v156
                                    }
                                    var v157 bool = (7 == (v5 + 237))
                                    var v158 bool = v3
                                    v2 += 264
                                    _ = v121
                                    _ = v157
                                    _ = v158
                                }
                            }
                            v4 -= (v4 << ((v6) & 63))
                            v4 = (v4 % ((255) | 1))
                        } else {
                            switch int(v6) {
                            case 0:
                                v159 := GenS12{F0: fi28, F1: 0}
                                _ = v159
                            default:
                                var v160 uint16 = v4
                                _ = v160
                            }
                            if v3 {
                                v161 := v1["\n\t"]
                                fmt.Println(v15)
                                _ = v161
                            }
                            switch "hello" {
                            case "":
                                v162 := v1["hello"]
                                _ = v162
                            case "\n\t":
                                var v163 bool = v3
                                _ = v163
                            default:
                                v4 -= v4
                            }
                        }
                        _ = fi28
                    }
                    _ = fi14
                    _ = v15
                }
                _ = v11
            } else {
                fmt.Println(v5, v5, v5)
                if (!(v3 && true)) {
                    if (245 != 198) {
                        fmt.Println(v5)
                        v164 := GenS14{F0: 622, F1: 912}
                        _ = v164
                    }
                    var v165 int16 = -112
                    fmt.Println(v165)
                    _ = v165
                }
                _ = len(v1)
            }
        }
    case 4294967295:
        if v3 {
            switch "hello" {
            case "foo bar":
                v6 += v6
            }
            switch 22 {
            case -1:
                for fi166 := 0; fi166 < 6; fi166++ {
                    for fi167 := 0; fi167 < 19; fi167++ {
                        var v168 float64 = (0.1 - 1e10)
                        _ = fi167
                        _ = v168
                    }
                    _ = fi166
                }
            case 9223372036854775807:
                var v169 uint32 = v5
                _ = v169
            case -890:
                v3 = (true && v3)
            }
            v170 := GenS14{F0: int(v6), F1: -353}
            _ = v170
        }
    default:
        v171 := []int{}
        _ = v171
    }
    var v172 uint8 = 127
    var v173 int16 = 0
    _ = len(v1)
    if (v3 || (("" + "a") != "hello")) {
        _ = len(v1)
        fmt.Println(v173, v172)
    } else {
        if ((v3 && v3) && v3) {
            var v174 int = -1
            switch ("foo bar" + "") {
            case "":
                v175 := v1["hello"]
                _ = v175
            case "\n\t":
                v172 += (v172 >> (((v6 + uint(v172))) & 63))
            default:
                v176 := []int{9223372036854775807, v174, (v174 >> ((4294967295) & 63))}
                _ = v176
            }
            fmt.Println(v174, v5)
            _ = v174
        } else {
            v177 := []int{42, 32768}
            v6 = 1
            fmt.Println(v5, v3, v6)
            _ = v177
        }
        if (!(!(v3 || v3))) {
            var v178 string = "hello"
            _ = v178
        }
        var v179 int32 = -146
        _ = v179
    }
    v180 := GenS14{F0: 32768, F1: 42}
    var v181 uint32 = (v5 >> ((1) & 63))
    switch int(v172) {
    case 9223372036854775807:
        v172 -= (v172 & (v172 | 255))
    default:
        v182 := GenS11{F0: -128, F1: "a"}
        _ = v182
    }
    v183 := []int{401}
    var v184 bool = (v2 != v173)
    for fi185 := 0; fi185 < 1; fi185++ {
        var v186 bool = (171 < 9223372036854775807)
        if ("foo bar" <= "foo bar") {
            if true {
                var v187 uint32 = v181
                var v188 uint32 = v5
                var v189 int32 = 0
                _ = v187
                _ = v188
                _ = v189
            } else {
                if v186 {
                    var v190 bool = (-941 < 1)
                    for fi191 := 0; fi191 < 10; fi191++ {
                        if (!false) {
                            var v192 bool = (!((0.1 != 1.5) || (((-2.25 + 0.0) + (float64(((0.1 * -2.25) + -2.25)) - -2.25)) == -2.25)))
                            _ = v192
                        } else {
                            v193 := GenS15{F0: 2, F1: ""}
                            v194 := GenS16{F0: int8(fi191), F1: v172}
                            v183 = append(v183, -1)
                            _ = v193
                            _ = v194
                        }
                        _ = fi191
                    }
                    _ = v190
                }
                v5 = (v181 - 2147483648)
                var v195 int64 = 0
                _ = v195
            }
            _ = len(v183)
            v196 := GenS4{F0: (v172 & (v172 - 42)), F1: "hello"}
            _ = v196
        }
        v197 := GenS17{F0: v172}
        _ = fi185
        _ = v186
        _ = v197
    }
    var v198 uint8 = v172
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v172:", v172)
    fmt.Println("v173:", v173)
    fmt.Println("v181:", v181)
    fmt.Println("slicelen v183:", len(v183), v183)
    fmt.Println("v184:", v184)
    fmt.Println("v198:", v198)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v172
    _ = v173
    _ = v180
    _ = v181
    _ = v183
    _ = v184
    _ = v198
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
