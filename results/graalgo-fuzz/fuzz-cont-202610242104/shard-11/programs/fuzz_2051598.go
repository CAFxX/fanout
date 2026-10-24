package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string }
type GenS2 struct { F0 int; F1 int16 }
type GenS3 struct { F0 int64; F1 int32; F2 uint64 }
type GenS4 struct { F0 int32; F1 uint8 }
type GenS5 struct { F0 uint32; F1 string }
type GenS6 struct { F0 uint64 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint64 }
type GenS9 struct { F0 int16; F1 bool }
type GenS10 struct { F0 bool; F1 string }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 int16; F1 int }
type GenS13 struct { F0 string }
type GenS14 struct { F0 uint32; F1 uint64; F2 uint64 }
type GenS15 struct { F0 int64 }
type GenS16 struct { F0 uint16 }
type GenS17 struct { F0 uint32 }
type GenS18 struct { F0 int8 }
type GenS19 struct { F0 uint16 }
type GenS20 struct { F0 uint32; F1 string; F2 uint32 }
type GenS21 struct { F0 uint; F1 bool }
type GenS22 struct { F0 int16 }
type GenS23 struct { F0 int32; F1 bool }
type GenS24 struct { F0 string }

func main() {
    v1 := []string{((((("\n\t" + ((("a" + "\n\t") + "hello") + "a")) + "\n\t") + "hello") + "a") + ("\n\t" + "foo bar")), (("a" + ("a" + ((("\n\t" + ("" + ("foo bar" + "\n\t"))) + "hello") + (("hello" + "\n\t") + "hello")))) + "a")}
    v2 := []string{}
    var v3 float64 = 1.5
    v4 := GenS1{F0: false, F1: "a"}
    v3 -= 3.14159
    v5 := GenS2{F0: -2147483648, F1: 1000}
    fmt.Println(v3)
    switch "\n\t" {
    case "a":
        if true {
            for fi6 := 0; fi6 < 14; fi6++ {
                v2 = append(v2, "hello")
                var v7 uint8 = 19
                if len(v1) > 0 {
                    v1[0] = "hello"
                }
                _ = fi6
                _ = v7
            }
            for fi8 := 0; fi8 < 9; fi8++ {
                if false {
                    v9 := GenS3{F0: -519, F1: 42, F2: 18446744073709551615}
                    _ = v9
                } else {
                    for fi10 := 0; fi10 < 8; fi10++ {
                        v3 -= v3
                        switch fi10 {
                        case 127:
                            v11 := GenS3{F0: 0, F1: int32(fi8), F2: 42}
                            _ = v11
                        default:
                            var v12 bool = (1 != 61)
                            _ = v12
                        }
                        v3 -= 0.0
                        _ = fi10
                    }
                }
                for fi13 := 0; fi13 < 9; fi13++ {
                    if true {
                        var v14 int8 = -128
                        var v15 uint32 = 2147483647
                        if (!false) {
                            v15 = (^(v15))
                        } else {
                            v3 = v3
                            var v16 uint16 = 2
                            _ = v16
                        }
                        _ = v14
                        _ = v15
                    } else {
                        if true {
                            _ = len(v2)
                            v17 := GenS4{F0: -1, F1: 1}
                            switch (fi13 | fi13) {
                            case 2147483647:
                                var v18 uint16 = 910
                                _ = v18
                            case -15:
                                v3 = 1e10
                            default:
                                v3 += v3
                            }
                            if false {
                                if (!(1 == 2)) {
                                    v1 = append(v1, "foo bar")
                                }
                                if (true && (71 == 0)) {
                                    v19 := GenS5{F0: 976, F1: "foo bar"}
                                    _ = v19
                                } else {
                                    if true {
                                        if (true || (false || (false || false))) {
                                            if ((float64((0.0 / 3.14159)) - (v3 - ((float64(0.0) + 1e10) * (3.14159 - v3)))) != v3) {
                                                if (true || (0 == 179)) {
                                                    var v20 int = (fi8 % ((32767) | 1))
                                                    v3 -= float64(1.5)
                                                    _ = v20
                                                }
                                                if (true || (1.5 < -2.25)) {
                                                    if len(v1) > 0 {
                                                        v1[0] = "\n\t"
                                                    }
                                                    switch "foo bar" {
                                                    case "a":
                                                        if true {
                                                            _ = len(v2)
                                                            var v21 int16 = -684
                                                            _ = v21
                                                        } else {
                                                            v3 = 3.14159
                                                        }
                                                    case "foo bar":
                                                        v22 := GenS6{F0: 42}
                                                        _ = v22
                                                    }
                                                    v3 += v3
                                                } else {
                                                    v23 := GenS6{F0: 0}
                                                    fmt.Println(fi8)
                                                    _ = v23
                                                }
                                            } else {
                                                if (1 != -803) {
                                                    v3 -= v3
                                                }
                                                var v24 int = (fi8 % ((-977) | 1))
                                                v24 += -9223372036854775808
                                                _ = v24
                                            }
                                            _ = len(v1)
                                            v3 = v3
                                        }
                                        v3 = (v3 + -2.25)
                                    } else {
                                        v3 = (-2.25 * 3.14159)
                                    }
                                    if len(v2) > 0 {
                                        v2[0] = ("hello" + "\n\t")
                                    }
                                    if false {
                                        if true {
                                            _ = len(v1)
                                            switch "\n\t" {
                                            case "foo bar":
                                                if false {
                                                    if (1 != -1) {
                                                        v25 := GenS2{F0: 1, F1: -37}
                                                        var v26 int8 = 100
                                                        _ = v25
                                                        _ = v26
                                                    } else {
                                                        v27 := GenS7{F0: false}
                                                        v2 = append(v2, "\n\t")
                                                        if true {
                                                            fmt.Println(fi13, fi13, fi13)
                                                            switch "a" {
                                                            case "a":
                                                                v3 -= 0.0
                                                            case "hello":
                                                                switch 421 {
                                                                case -1:
                                                                    var v28 float32 = 0.1
                                                                    _ = v28
                                                                default:
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                }
                                                            case "":
                                                                if false {
                                                                    if (501 > 42) {
                                                                        if false {
                                                                            if (101 <= 1) {
                                                                                fmt.Println(fi13, fi8)
                                                                                var v29 uint32 = 2
                                                                                v3 = v3
                                                                                _ = v29
                                                                            } else {
                                                                                _ = len(v1)
                                                                            }
                                                                        }
                                                                        fmt.Println(fi13)
                                                                    }
                                                                    v30 := GenS1{F0: (-2.25 != 1e10), F1: "hello"}
                                                                    if (-489 < -1) {
                                                                        v31 := GenS8{F0: 18446744073709551615}
                                                                        _ = v31
                                                                    }
                                                                    _ = v30
                                                                }
                                                                if (v3 <= ((1.5 - (1.5 + 1e10)) + v3)) {
                                                                    var v32 int = -814
                                                                    _ = v32
                                                                } else {
                                                                    var v33 int32 = 42
                                                                    _ = len(v2)
                                                                    var v34 uint64 = 42
                                                                    _ = v33
                                                                    _ = v34
                                                                }
                                                            default:
                                                                if ((0 < 42) || (470 <= 73)) {
                                                                    var v35 int64 = 1
                                                                    _ = v35
                                                                } else {
                                                                    var v36 int8 = 42
                                                                    v36 = (v36 | 127)
                                                                    v3 += float64(3.14159)
                                                                    _ = v36
                                                                }
                                                            }
                                                        }
                                                        _ = v27
                                                    }
                                                } else {
                                                    v2 = append(v2, (("\n\t" + "") + ("hello" + ("" + ("hello" + "a")))))
                                                    v37 := GenS9{F0: -32768, F1: true}
                                                    _ = v37
                                                }
                                            }
                                            switch (fi8 - (fi8 + fi13)) {
                                            case -64:
                                                v2 = append(v2, ("hello" + "\n\t"))
                                            case -2147483648:
                                                v3 = ((v3 + (v3 / v3)) / float64(0.1))
                                            case 32767:
                                                v3 = 1e10
                                            default:
                                                if len(v1) > 0 {
                                                    v1[0] = "\n\t"
                                                }
                                            }
                                        }
                                    }
                                }
                                var v38 uint8 = 42
                                _ = v38
                            } else {
                                v3 += -2.25
                            }
                            var v39 int16 = 42
                            _ = v17
                            _ = v39
                        } else {
                            if (14 > 2) {
                                switch "foo bar" {
                                case "foo bar":
                                    v3 += ((3.14159 + float64(1e10)) / v3)
                                case "":
                                    if len(v1) > 0 {
                                        v1[0] = "hello"
                                    }
                                case "a":
                                    v40 := GenS10{F0: true, F1: "a"}
                                    _ = v40
                                }
                                switch fi13 {
                                case 874:
                                    if true {
                                        if true {
                                            switch "hello" {
                                            case "\n\t":
                                                v41 := GenS11{F0: -327}
                                                _ = v41
                                            }
                                        } else {
                                            v3 = v3
                                            _ = len(v2)
                                            if false {
                                                v3 -= v3
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            }
                                        }
                                        if len(v2) > 0 {
                                            v2[0] = "a"
                                        }
                                        v2 = append(v2, ("hello" + (("" + "\n\t") + "hello")))
                                        var v42 int8 = 18
                                        _ = v42
                                    }
                                }
                                v3 += 3.14159
                            } else {
                                if len(v1) >= 2 {
                                    v1 = v1[1:len(v1)]
                                }
                            }
                            switch 128 {
                            case 1:
                                var v43 int8 = -49
                                _ = v43
                            case 32767:
                                v3 -= (3.14159 / 1.5)
                            case 2147483647:
                                switch fi8 {
                                case -2147483648:
                                    v3 -= (3.14159 * 0.0)
                                case 42:
                                    var v44 int64 = 9223372036854775807
                                    _ = v44
                                default:
                                    if (1 == 10) {
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    }
                                }
                            default:
                                fmt.Println(fi13, fi13)
                            }
                        }
                        switch (fi8 % ((fi13) | 1)) {
                        case -1:
                            if true {
                                switch "" {
                                case "\n\t":
                                    v3 += (v3 * 3.14159)
                                    var v45 string = ""
                                    _ = v45
                                case "hello":
                                    var v46 int64 = 855
                                    _ = v46
                                }
                                switch ("" + "") {
                                case "\n\t":
                                    var v47 uint32 = 65535
                                    _ = v47
                                case "":
                                    v3 -= ((v3 * float64((1.5 + float32(0.1)))) + float64(float32(v3)))
                                }
                                switch fi13 {
                                case 0:
                                    v3 = 0.1
                                    v3 -= (0.1 * ((v3 / 1.0) / float64(1.5)))
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    v3 -= (((v3 - ((v3 / v3) - ((0.1 / v3) / v3))) / (v3 * v3)) / ((1e10 / v3) - -2.25))
                                default:
                                    if true {
                                        if false {
                                            var v48 uint64 = 2
                                            v48 -= 42
                                            fmt.Println(fi13, v48)
                                            _ = v48
                                        } else {
                                            var v49 float32 = 1.5
                                            if (true || false) {
                                                v50 := GenS12{F0: -1, F1: 1}
                                                _ = len(v2)
                                                v51 := GenS13{F0: "foo bar"}
                                                _ = v50
                                                _ = v51
                                            }
                                            var v52 float32 = v49
                                            _ = v49
                                            _ = v52
                                        }
                                    } else {
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        switch ("foo bar" + ("a" + "a")) {
                                        case "":
                                            var v53 bool = (1 == 1)
                                            _ = v53
                                        default:
                                            _ = len(v1)
                                        }
                                    }
                                }
                                switch "" {
                                case "":
                                    _ = len(v2)
                                case "foo bar":
                                    switch "foo bar" {
                                    case "hello":
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    case "foo bar":
                                        v3 -= ((0.1 - -2.25) / 1.0)
                                    case "\n\t":
                                        fmt.Println(v3, v3, v3)
                                    default:
                                        var v54 int8 = -128
                                        _ = v54
                                    }
                                case "hello":
                                    var v55 int32 = -1
                                    _ = v55
                                default:
                                    var v56 bool = (!false)
                                    _ = v56
                                }
                                if (4294967295 > 582) {
                                    if (!true) {
                                        v57 := GenS13{F0: "\n\t"}
                                        _ = v57
                                    }
                                    switch "foo bar" {
                                    case "foo bar":
                                        var v58 uint = 42
                                        if (!(((!true) && (9223372036854775807 == int64(fi8))) || ((true || true) || true))) {
                                            switch 694 {
                                            case 4294967295:
                                                v3 -= v3
                                            case 610:
                                                v3 -= v3
                                            case 128:
                                                var v59 bool = false
                                                _ = len(v1)
                                                _ = v59
                                            }
                                            if (fi8 < fi13) {
                                                var v60 float64 = (1.5 - 1.5)
                                                v60 += (v60 / float64(1e10))
                                                v3 = v3
                                                _ = len(v1)
                                                var v61 uint64 = 42
                                                if (-9223372036854775808 <= 0) {
                                                    var v62 string = ""
                                                    if (!(2147483647 > 2147483647)) {
                                                        var v63 string = (((v62 + v62) + "") + v62)
                                                        v62 = ((v62 + v63) + v62)
                                                        _ = v63
                                                    }
                                                    _ = v62
                                                }
                                                switch 32767 {
                                                case 9223372036854775807:
                                                    var v64 int64 = 42
                                                    _ = v64
                                                case 2147483647:
                                                    v65 := GenS6{F0: v61}
                                                    v66 := GenS6{F0: 885}
                                                    _ = v65
                                                    _ = v66
                                                default:
                                                    if (true || true) {
                                                        fmt.Println(fi13, fi8, v60)
                                                        if ((fi8 & fi13) != 127) {
                                                            v58 = v58
                                                            v58 = (^(v58))
                                                            var v67 uint64 = 0
                                                            _ = v67
                                                        }
                                                    } else {
                                                        v61 = (v61 + v61)
                                                    }
                                                    v60 = (v3 + (float64(((0.1 - (1.5 * 3.14159)) * ((float32((v3 * float64(-2.25))) * (1e10 / 0.1)) - (float32(v3) * 1.5)))) * (v3 - -2.25)))
                                                    v68 := GenS9{F0: 42, F1: false}
                                                    _ = v68
                                                }
                                                _ = v60
                                                _ = v61
                                            }
                                            var v69 float64 = v3
                                            _ = v69
                                        } else {
                                            var v70 uint16 = 255
                                            _ = v70
                                        }
                                        _ = v58
                                    case "":
                                        switch "a" {
                                        case "foo bar":
                                            if false {
                                                var v71 int16 = 0
                                                fmt.Println(fi8, v71)
                                                if ((^(fi8)) >= (fi8 * (fi13 % ((fi8) | 1)))) {
                                                    var v72 string = "foo bar"
                                                    if true {
                                                        v3 += ((((v3 * v3) * v3) + float64(((0.1 * (0.0 - (((((0.0 + 0.0) - (0.1 / 1e10)) + 0.1) * float32(((v3 / ((3.14159 * float64(-2.25)) * 0.1)) * v3))) + 1.5))) + (0.0 * (float32(((0.0 / ((v3 - 1e10) + (3.14159 * float64(0.0)))) - v3)) * -2.25))))) + v3)
                                                        switch 127 {
                                                        case -365:
                                                            v73 := GenS10{F0: (false && (!(v3 > v3))), F1: v72}
                                                            _ = v73
                                                        }
                                                        var v74 uint = 0
                                                        _ = v74
                                                    } else {
                                                        if len(v2) > 0 {
                                                            v2[0] = v72
                                                        }
                                                    }
                                                    fmt.Println(fi13, fi13)
                                                    v72 = v72
                                                    _ = v72
                                                } else {
                                                    var v75 uint64 = 0
                                                    _ = v75
                                                }
                                                _ = v71
                                            }
                                            if len(v1) > 0 {
                                                v1[0] = "foo bar"
                                            }
                                        case "":
                                            if (!(2 == 42)) {
                                                var v76 int64 = -821
                                                if true {
                                                    var v77 float64 = ((((v3 / (float64(float32(0.1)) + float64((float32(v3) / 1.5)))) / v3) + v3) + v3)
                                                    if true {
                                                        var v78 int32 = -1
                                                        v2 = append(v2, ("hello" + ("a" + "foo bar")))
                                                        _ = v78
                                                    }
                                                    v79 := GenS14{F0: 762, F1: 4294967295, F2: 694}
                                                    if (!false) {
                                                        var v80 int32 = -1
                                                        v81 := GenS15{F0: (v76 | (v76 ^ v76))}
                                                        _ = v80
                                                        _ = v81
                                                    } else {
                                                        v76 += 926
                                                    }
                                                    switch "hello" {
                                                    case "a":
                                                        v82 := GenS16{F0: 983}
                                                        _ = v82
                                                    case "foo bar":
                                                        if true {
                                                            var v83 int32 = 1
                                                            _ = v83
                                                        } else {
                                                            v3 = v3
                                                            switch fi8 {
                                                            case 127:
                                                                v3 = 0.0
                                                            case 941:
                                                                var v84 uint32 = 173
                                                                _ = v84
                                                            default:
                                                                switch "a" {
                                                                case "foo bar":
                                                                    switch (fi8 & fi13) {
                                                                    case 9223372036854775807:
                                                                        if (-57 < 59) {
                                                                            fmt.Println(v77, v76)
                                                                            v77 = (float64(0.0) - 0.0)
                                                                            fmt.Println(v77, v76, v76)
                                                                            v76 -= (v76 << ((2) & 63))
                                                                            v85 := GenS17{F0: 1}
                                                                            _ = len(v2)
                                                                            if len(v2) >= 2 {
                                                                                v2 = v2[1:len(v2)]
                                                                            }
                                                                            if false {
                                                                                switch ("" + "\n\t") {
                                                                                case "a":
                                                                                    if len(v2) > 0 {
                                                                                        v2[0] = "\n\t"
                                                                                    }
                                                                                    var v86 uint16 = 0
                                                                                    var v87 int16 = 42
                                                                                    v88 := GenS18{F0: 127}
                                                                                    var v89 float64 = v3
                                                                                    _ = v86
                                                                                    _ = v87
                                                                                    _ = v88
                                                                                    _ = v89
                                                                                case "\n\t":
                                                                                    if len(v2) > 0 {
                                                                                        v2[0] = ("foo bar" + ("hello" + "hello"))
                                                                                    }
                                                                                }
                                                                            } else {
                                                                                switch "\n\t" {
                                                                                case "\n\t":
                                                                                    switch 32768 {
                                                                                    case 450:
                                                                                        if len(v2) >= 2 {
                                                                                            v2 = v2[1:len(v2)]
                                                                                        }
                                                                                    case -852:
                                                                                        v90 := GenS18{F0: -128}
                                                                                        _ = v90
                                                                                    case 127:
                                                                                        var v91 float32 = 1e10
                                                                                        _ = v91
                                                                                    }
                                                                                case "foo bar":
                                                                                    fmt.Println(v3)
                                                                                }
                                                                                var v92 float32 = ((1e10 / 1.0) - (((1e10 * 1.5) + 1e10) / (1e10 / 0.1)))
                                                                                if (42 < 108) {
                                                                                    fmt.Println(fi13, v76, v3)
                                                                                } else {
                                                                                    v2 = append(v2, "a")
                                                                                    fmt.Println(v77, fi8, v76)
                                                                                }
                                                                                if false {
                                                                                    v93 := GenS18{F0: 1}
                                                                                    var v94 int16 = 32767
                                                                                    var v95 int = fi13
                                                                                    _ = v93
                                                                                    _ = v94
                                                                                    _ = v95
                                                                                } else {
                                                                                    v92 = float32((v3 + (float64(v92) - v3)))
                                                                                }
                                                                                _ = v92
                                                                            }
                                                                            var v96 uint32 = 65535
                                                                            _ = v85
                                                                            _ = v96
                                                                        }
                                                                    case -155:
                                                                        if false {
                                                                            fmt.Println(v76)
                                                                        } else {
                                                                            v3 -= 1.5
                                                                            v3 = float64((float32(((v3 / 0.1) - v3)) - 0.1))
                                                                            fmt.Println(v3)
                                                                            if false {
                                                                                var v97 int8 = 0
                                                                                v76 += v76
                                                                                _ = v97
                                                                            } else {
                                                                                fmt.Println(fi13)
                                                                                fmt.Println(v76, v76)
                                                                            }
                                                                            if ((!true) || true) {
                                                                                v98 := GenS15{F0: -1}
                                                                                _ = v98
                                                                            } else {
                                                                                v1 = append(v1, "foo bar")
                                                                                if len(v1) >= 2 {
                                                                                    v1 = v1[1:len(v1)]
                                                                                }
                                                                                switch 127 {
                                                                                case 4294967295:
                                                                                    fmt.Println(v3, fi13, fi13)
                                                                                case 2147483647:
                                                                                    if true {
                                                                                        if (!(("a" + "\n\t") >= ("foo bar" + "a"))) {
                                                                                            fmt.Println(fi8, v76, v77)
                                                                                        }
                                                                                        var v99 string = "hello"
                                                                                        switch 4294967295 {
                                                                                        case -1:
                                                                                            v77 += v77
                                                                                        case -517:
                                                                                            if ((1 != 0) || (0 != -2147483648)) {
                                                                                                v100 := GenS19{F0: 995}
                                                                                                v3 = v77
                                                                                                _ = v100
                                                                                            } else {
                                                                                                if len(v1) > 0 {
                                                                                                    v1[0] = ("" + v99)
                                                                                                }
                                                                                                if false {
                                                                                                    if (0 > 2) {
                                                                                                        if len(v1) >= 2 {
                                                                                                            v1 = v1[1:len(v1)]
                                                                                                        }
                                                                                                        fmt.Println(v76, fi13, v3)
                                                                                                    }
                                                                                                    v77 -= v3
                                                                                                }
                                                                                            }
                                                                                        default:
                                                                                            var v101 uint = 4294967295
                                                                                            _ = v101
                                                                                        }
                                                                                        _ = v99
                                                                                    } else {
                                                                                        switch fi13 {
                                                                                        case -9223372036854775808:
                                                                                            var v102 int16 = 947
                                                                                            _ = v102
                                                                                        case -136:
                                                                                            v77 = (v3 * float64(0.0))
                                                                                        default:
                                                                                            fmt.Println(v77)
                                                                                        }
                                                                                    }
                                                                                case -93:
                                                                                    v103 := GenS20{F0: 2, F1: "hello", F2: 42}
                                                                                    _ = v103
                                                                                default:
                                                                                    if ((127 < 127) || (255 != 1)) {
                                                                                        v1 = append(v1, "a")
                                                                                        var v104 uint32 = 531
                                                                                        if ((v76 << ((247) & 63)) != (v76 * (v76 >> ((534) & 63)))) {
                                                                                            switch (-(fi13)) {
                                                                                            case 32767:
                                                                                                v105 := GenS16{F0: 13}
                                                                                                _ = v105
                                                                                            default:
                                                                                                fmt.Println(v104, fi8, v76)
                                                                                            }
                                                                                        } else {
                                                                                            if ((-1 > (fi8 / ((fi8) | 1))) && false) {
                                                                                                fmt.Println(v77, fi8, v77)
                                                                                                fmt.Println(v104, v77, fi13)
                                                                                                var v106 float64 = v3
                                                                                                _ = v106
                                                                                            }
                                                                                            v107 := GenS12{F0: 217, F1: fi13}
                                                                                            _ = v107
                                                                                        }
                                                                                        var v108 int64 = 1
                                                                                        _ = v104
                                                                                        _ = v108
                                                                                    } else {
                                                                                        if (fi8 >= fi13) {
                                                                                            fmt.Println(fi13)
                                                                                        } else {
                                                                                            v3 += ((float64((0.1 + (-2.25 / -2.25))) * v3) * (v3 / 3.14159))
                                                                                        }
                                                                                        v109 := GenS21{F0: 1, F1: ((^(v76)) == v76)}
                                                                                        _ = v109
                                                                                    }
                                                                                }
                                                                            }
                                                                            v77 += -2.25
                                                                        }
                                                                    case 32767:
                                                                        var v110 float32 = ((0.1 + 0.1) + ((1e10 + -2.25) + 0.1))
                                                                        _ = v110
                                                                    }
                                                                case "":
                                                                    v111 := GenS2{F0: fi13, F1: 744}
                                                                    _ = v111
                                                                case "\n\t":
                                                                    var v112 int64 = (-(v76))
                                                                    _ = v112
                                                                default:
                                                                    v3 = 1e10
                                                                }
                                                            }
                                                        }
                                                    case "hello":
                                                        if len(v1) > 0 {
                                                            v1[0] = ""
                                                        }
                                                    default:
                                                        v113 := GenS15{F0: v76}
                                                        v3 -= float64((0.0 + (3.14159 / 1e10)))
                                                        _ = v113
                                                    }
                                                    _ = v77
                                                    _ = v79
                                                } else {
                                                    var v114 uint16 = 32767
                                                    v3 -= v3
                                                    _ = v114
                                                }
                                                switch (fi8 | (fi8 + fi13)) {
                                                case 2147483647:
                                                    if ((1 == -512) && (255 > 32767)) {
                                                        _ = len(v2)
                                                        fmt.Println(v76)
                                                        if (false && true) {
                                                            if false {
                                                                var v115 uint64 = 1
                                                                if (1e10 < 0.0) {
                                                                    if (4294967295 >= 232) {
                                                                        if false {
                                                                            v116 := GenS22{F0: 42}
                                                                            _ = v116
                                                                        } else {
                                                                            var v117 int8 = -128
                                                                            _ = v117
                                                                        }
                                                                        fmt.Println(v115)
                                                                        var v118 int64 = int64(fi13)
                                                                        _ = v118
                                                                    }
                                                                }
                                                                switch "\n\t" {
                                                                case "foo bar":
                                                                    fmt.Println(v115, fi8)
                                                                case "\n\t":
                                                                    switch "hello" {
                                                                    case "":
                                                                        v3 = 0.0
                                                                    case "\n\t":
                                                                        var v119 uint32 = 561
                                                                        _ = v119
                                                                    case "a":
                                                                        var v120 uint8 = 2
                                                                        _ = v120
                                                                    }
                                                                }
                                                                var v121 uint16 = 42
                                                                v115 = 704
                                                                _ = v115
                                                                _ = v121
                                                            }
                                                        } else {
                                                            v76 -= (v76 << ((18446744073709551615) & 63))
                                                        }
                                                    }
                                                case -233:
                                                    var v122 int64 = v76
                                                    _ = v122
                                                case 476:
                                                    v2 = append(v2, "")
                                                default:
                                                    var v123 bool = (2 != 82)
                                                    _ = v123
                                                }
                                                _ = v76
                                            } else {
                                                v2 = append(v2, "hello")
                                                v124 := GenS11{F0: 42}
                                                var v125 float32 = (float32(float64(0.0)) / 0.1)
                                                _ = v124
                                                _ = v125
                                            }
                                        }
                                    case "\n\t":
                                        if len(v2) > 0 {
                                            v2[0] = "\n\t"
                                        }
                                        var v126 float64 = -2.25
                                        fmt.Println(v3, v126, v126)
                                        _ = v126
                                    }
                                } else {
                                    var v127 uint16 = 777
                                    switch "foo bar" {
                                    case "hello":
                                        v1 = append(v1, (("\n\t" + "foo bar") + "a"))
                                    default:
                                        fmt.Println(fi13)
                                    }
                                    _ = v127
                                }
                            } else {
                                switch (fi8 + fi13) {
                                case 42:
                                    if len(v2) > 0 {
                                        v2[0] = ""
                                    }
                                default:
                                    var v128 bool = true
                                    _ = v128
                                }
                                v1 = append(v1, ("" + "foo bar"))
                                v3 -= (v3 + v3)
                            }
                        case 32768:
                            var v129 uint16 = 255
                            _ = v129
                        }
                    }
                    if (4294967295 < 2) {
                        v130 := GenS4{F0: -2147483648, F1: 255}
                        _ = v130
                    } else {
                        fmt.Println(fi13)
                        if (false || (2 != 725)) {
                            switch (fi8 % (((fi8 - (fi8 | (fi8 | fi8)))) | 1)) {
                            case 42:
                                switch "a" {
                                case "\n\t":
                                    var v131 bool = true
                                    _ = v131
                                }
                            case 998:
                                var v132 int16 = 543
                                _ = v132
                            }
                        } else {
                            var v133 uint8 = 2
                            v1 = append(v1, "foo bar")
                            fmt.Println(fi13, fi13, fi8)
                            v3 = 1.5
                            _ = v133
                        }
                    }
                    _ = fi13
                }
                _ = fi8
            }
        } else {
            for fi134 := 0; fi134 < 4; fi134++ {
                for fi135 := 0; fi135 < 4; fi135++ {
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                    _ = fi135
                }
                _ = fi134
            }
            v3 = 3.14159
        }
    case "hello":
        if (false || false) {
            v3 += (float64(1.5) / v3)
        } else {
            for fi136 := 0; fi136 < 11; fi136++ {
                fmt.Println(v3, v3)
                v3 = 0.0
                var v137 string = "foo bar"
                _ = fi136
                _ = v137
            }
        }
    case "":
        if len(v2) > 0 {
            v2[0] = "foo bar"
        }
    default:
        var v138 int8 = -1
        _ = v138
    }
    v3 -= float64((-2.25 / -2.25))
    v139 := GenS23{F0: -942, F1: (false || ((v3 / v3) != 0.0))}
    fmt.Println(v3)
    switch "foo bar" {
    case "a":
        v3 -= (v3 + 0.1)
    case "":
        if len(v1) > 0 {
            v1[0] = "\n\t"
        }
    case "hello":
        if true {
            v3 = ((1e10 - v3) * (float64((float32(v3) + -2.25)) / v3))
            v140 := GenS15{F0: 735}
            for fi141 := 0; fi141 < 6; fi141++ {
                for fi142 := 0; fi142 < 6; fi142++ {
                    fmt.Println(v3)
                    fmt.Println(v3, fi142, fi141)
                    fmt.Println(v3)
                    _ = fi142
                }
                v143 := GenS11{F0: -1}
                _ = fi141
                _ = v143
            }
            _ = v140
        }
    default:
        switch 0 {
        case 9223372036854775807:
            if true {
                var v144 int8 = -79
                v144 = (^(v144))
                _ = v144
            } else {
                switch ("foo bar" + "") {
                case "a":
                    if len(v2) >= 2 {
                        v2 = v2[1:len(v2)]
                    }
                case "hello":
                    for fi145 := 0; fi145 < 4; fi145++ {
                        v146 := GenS22{F0: -636}
                        _ = fi145
                        _ = v146
                    }
                default:
                    v3 = v3
                }
                var v147 uint32 = 2147483648
                v147 += 65535
                _ = v147
            }
        }
    }
    for fi148 := 0; fi148 < 8; fi148++ {
        if len(v2) > 0 {
            v2[0] = ("a" + ("\n\t" + "foo bar"))
        }
        v3 -= -2.25
        for fi149 := 0; fi149 < 11; fi149++ {
            v150 := GenS21{F0: 404, F1: (42 <= -42)}
            var v151 uint = 994
            _ = fi149
            _ = v150
            _ = v151
        }
        _ = fi148
    }
    _ = len(v1)
    _ = len(v1)
    v152 := GenS24{F0: "a"}
    var v153 int = 1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v153:", v153)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v139
    _ = v152
    _ = v153
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
