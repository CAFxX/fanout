package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 uint32 }
type GenS2 struct { F0 int32; F1 int64; F2 bool }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 string; F1 uint16 }
type GenS5 struct { F0 uint; F1 string; F2 int8 }
type GenS6 struct { F0 int32; F1 uint8 }
type GenS7 struct { F0 uint32; F1 int }
type GenS8 struct { F0 uint8; F1 int32; F2 int16 }
type GenS9 struct { F0 uint8 }
type GenS10 struct { F0 uint64; F1 uint8 }
type GenS11 struct { F0 string; F1 uint64; F2 bool }

func main() {
    var v1 int8 = 119
    var v2 bool = (!false)
    var v3 int32 = -1
    v4 := []int{-1, -1, -1}
    var v5 float32 = 1.5
    if (18446744073709551615 != 985) {
        switch 9223372036854775807 {
        case -9223372036854775808:
            v6 := GenS1{F0: "a", F1: 18446744073709551615, F2: 2}
            _ = v6
        case 32767:
            v5 = (((v5 - 0.1) / (0.0 / v5)) + (v5 + -2.25))
        case -2147483648:
            for fi7 := 0; fi7 < 9; fi7++ {
                for fi8 := 0; fi8 < 16; fi8++ {
                    if v2 {
                        v4 = append(v4, fi8)
                    } else {
                        var v9 uint64 = 0
                        var v10 int16 = 859
                        if v2 {
                            v11 := GenS1{F0: "\n\t", F1: 2, F2: 2}
                            v10 = (v10 & v10)
                            if ((fi7 / ((127) | 1)) == (^(fi7))) {
                                if (v1 != (v1 / ((42) | 1))) {
                                    v2 = v2
                                    if (255 <= 127) {
                                        v12 := GenS2{F0: 42, F1: 9223372036854775807, F2: true}
                                        var v13 uint16 = 255
                                        var v14 float32 = (float32(-2.25) + ((v5 - 3.14159) * v5))
                                        _ = v12
                                        _ = v13
                                        _ = v14
                                    }
                                    fmt.Println(v2, v3)
                                } else {
                                    v5 += 0.1
                                    v9 += (v9 - (^(v9)))
                                }
                                v5 -= (v5 - ((v5 * (1.5 - v5)) + (0.1 * 1e10)))
                            } else {
                                switch (fi7 + 2147483647) {
                                case -9223372036854775808:
                                    switch ("hello" + (("hello" + "") + "hello")) {
                                    case "":
                                        if len(v4) > 0 {
                                            v4[0] = (^(fi7))
                                        }
                                    case "hello":
                                        fmt.Println(v5, v2)
                                    default:
                                        v3 = v3
                                    }
                                case 4294967295:
                                    var v15 float32 = ((1e10 - float32(0.0)) - v5)
                                    _ = v15
                                case 128:
                                    var v16 bool = true
                                    _ = v16
                                default:
                                    if (90 != -9223372036854775808) {
                                        v17 := GenS2{F0: (v3 | (-(v3))), F1: -9223372036854775808, F2: v2}
                                        _ = v17
                                    } else {
                                        switch (fi7 & fi7) {
                                        case 9223372036854775807:
                                            fmt.Println(v2, v2)
                                        case 32768:
                                            var v18 uint32 = 21
                                            _ = v18
                                        default:
                                            v1 = (v1 << ((4294967295) & 63))
                                            v3 = (v3 - v3)
                                        }
                                        if false {
                                            if v2 {
                                                switch fi7 {
                                                case 42:
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                case 32767:
                                                    var v19 uint8 = 1
                                                    _ = v19
                                                case 128:
                                                    fmt.Println(v3, v2, v9)
                                                default:
                                                    if v2 {
                                                        v3 -= 588
                                                        _ = len(v4)
                                                        v3 -= (v3 << ((0) & 63))
                                                    }
                                                }
                                                if v2 {
                                                    var v20 string = ""
                                                    if v2 {
                                                        if v2 {
                                                            if len(v4) > 0 {
                                                                v4[0] = (fi7 % (((^(fi8))) | 1))
                                                            }
                                                            fmt.Println(fi7, v9, v3)
                                                            var v21 uint64 = (v9 % (((^(v9))) | 1))
                                                            _ = v21
                                                        } else {
                                                            v4 = append(v4, (fi8 ^ (fi8 ^ (-(fi7)))))
                                                        }
                                                        v10 = -32768
                                                        if len(v4) > 0 {
                                                            v4[0] = 0
                                                        }
                                                        if v2 {
                                                            var v22 bool = v2
                                                            _ = v22
                                                        } else {
                                                            var v23 float32 = (3.14159 * v5)
                                                            if v2 {
                                                                v3 = (v3 / (((v3 ^ v3)) | 1))
                                                            }
                                                            _ = v23
                                                        }
                                                    } else {
                                                        fmt.Println(v1, v2)
                                                        fmt.Println(v1)
                                                        v2 = v2
                                                    }
                                                    _ = v20
                                                }
                                                if len(v4) > 0 {
                                                    v4[0] = (fi8 / ((fi8) | 1))
                                                }
                                            } else {
                                                v24 := GenS3{F0: 42}
                                                _ = v24
                                            }
                                        }
                                    }
                                }
                                _ = len(v4)
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                v1 = v1
                            }
                            _ = v11
                        } else {
                            var v25 int = (fi8 - 32767)
                            switch (fi8 - (-(v25))) {
                            case 593:
                                switch "foo bar" {
                                case "hello":
                                    switch (v25 | (v25 | (fi7 | 32768))) {
                                    case 772:
                                        switch "\n\t" {
                                        case "":
                                            var v26 string = ("hello" + (("" + ("" + "hello")) + "a"))
                                            _ = v26
                                        case "\n\t":
                                            v4 = append(v4, (fi8 + (fi7 >> ((672) & 63))))
                                        default:
                                            if (v2 && (!v2)) {
                                                if v2 {
                                                    switch (fi7 / ((v25) | 1)) {
                                                    case 2147483647:
                                                        fmt.Println(v3)
                                                    case 32768:
                                                        if false {
                                                            if len(v4) > 0 {
                                                                v4[0] = fi7
                                                            }
                                                        } else {
                                                            var v27 float32 = v5
                                                            _ = len(v4)
                                                            _ = v27
                                                        }
                                                    case -237:
                                                        fmt.Println(fi8, v3, fi8)
                                                    }
                                                    _ = len(v4)
                                                    var v28 uint8 = 120
                                                    _ = v28
                                                }
                                            } else {
                                                var v29 float32 = float32((1.5 * (float64(3.14159) / 3.14159)))
                                                if v2 {
                                                    var v30 uint16 = 41
                                                    if false {
                                                        var v31 bool = (v9 == v9)
                                                        _ = v31
                                                    } else {
                                                        v1 = (v1 - v1)
                                                        if ((^(v30)) >= 0) {
                                                            v25 -= -1
                                                            switch (v25 % (((fi7 / (((v25 - fi7)) | 1))) | 1)) {
                                                            case 2147483647:
                                                                v9 = 903
                                                            case 782:
                                                                var v32 int8 = (v1 + 100)
                                                                _ = v32
                                                            case 4294967295:
                                                                fmt.Println(v29, v25)
                                                            default:
                                                                if true {
                                                                    v4 = append(v4, fi7)
                                                                    var v33 bool = false
                                                                    v9 += (v9 / ((v9) | 1))
                                                                    switch (fi7 - (v25 << ((671) & 63))) {
                                                                    case 252:
                                                                        v4 = append(v4, 837)
                                                                    case -920:
                                                                        fmt.Println(v25, v5)
                                                                    case -808:
                                                                        var v34 string = (("a" + ("hello" + ("hello" + ""))) + "")
                                                                        _ = v34
                                                                    default:
                                                                        if (!v2) {
                                                                            var v35 string = "foo bar"
                                                                            var v36 string = "hello"
                                                                            var v37 int64 = 9223372036854775807
                                                                            if (!v2) {
                                                                                v29 -= v5
                                                                            } else {
                                                                                if true {
                                                                                    if v33 {
                                                                                        fmt.Println(v37)
                                                                                    }
                                                                                    if (!((v33 && false) && (((v36 + (v36 + v36)) > "\n\t") && (v9 >= v9)))) {
                                                                                        if v33 {
                                                                                            switch (v36 + v36) {
                                                                                            case "a":
                                                                                                switch "\n\t" {
                                                                                                case "foo bar":
                                                                                                    var v38 float32 = (3.14159 + (v5 * float32(3.14159)))
                                                                                                    _ = v38
                                                                                                case "\n\t":
                                                                                                    v2 = ((v2 || true) && (1 < 27))
                                                                                                case "hello":
                                                                                                    v39 := GenS1{F0: ("hello" + "\n\t"), F1: 4294967295, F2: 516}
                                                                                                    _ = v39
                                                                                                }
                                                                                            case "hello":
                                                                                                var v40 int = (fi7 >> ((996) & 63))
                                                                                                _ = v40
                                                                                            }
                                                                                            var v41 uint32 = 2
                                                                                            _ = v41
                                                                                        }
                                                                                        var v42 uint16 = (v30 % ((v30) | 1))
                                                                                        var v43 int64 = v37
                                                                                        if v2 {
                                                                                            if v33 {
                                                                                                switch (fi7 + (fi8 | (-(v25)))) {
                                                                                                case -9223372036854775808:
                                                                                                    var v44 bool = (-2.25 > ((v29 - (v5 + v5)) + float32(3.14159)))
                                                                                                    _ = v44
                                                                                                case 566:
                                                                                                    if (18446744073709551615 < 2) {
                                                                                                        if false {
                                                                                                            v35 = ""
                                                                                                            var v45 int32 = v3
                                                                                                            v4 = append(v4, (fi7 % (((v25 % ((872) | 1))) | 1)))
                                                                                                            _ = v45
                                                                                                        }
                                                                                                    } else {
                                                                                                        v46 := GenS2{F0: (-(v3)), F1: 0, F2: (!false)}
                                                                                                        if len(v4) > 0 {
                                                                                                            v4[0] = (fi7 - fi7)
                                                                                                        }
                                                                                                        var v47 bool = v33
                                                                                                        _ = v46
                                                                                                        _ = v47
                                                                                                    }
                                                                                                default:
                                                                                                    fmt.Println(v10)
                                                                                                }
                                                                                                if (v33 || false) {
                                                                                                    if len(v4) >= 2 {
                                                                                                        v4 = v4[1:len(v4)]
                                                                                                    }
                                                                                                    fmt.Println(v5)
                                                                                                    if len(v4) > 0 {
                                                                                                        v4[0] = fi8
                                                                                                    }
                                                                                                } else {
                                                                                                    if ((-(v1)) > int8(v43)) {
                                                                                                        var v48 int = (fi7 | (^(v25)))
                                                                                                        if v2 {
                                                                                                            var v49 uint64 = (^(v9))
                                                                                                            fmt.Println(v3, v5)
                                                                                                            switch (fi7 % (((v25 + 9223372036854775807)) | 1)) {
                                                                                                            case 9223372036854775807:
                                                                                                                v4 = append(v4, (v25 | fi8))
                                                                                                            case 0:
                                                                                                                switch (fi7 + -272) {
                                                                                                                case 128:
                                                                                                                    fmt.Println(v37)
                                                                                                                case 220:
                                                                                                                    var v50 uint64 = 423
                                                                                                                    _ = v50
                                                                                                                case 6:
                                                                                                                    switch v35 {
                                                                                                                    case "foo bar":
                                                                                                                        v37 -= v43
                                                                                                                    case "":
                                                                                                                        var v51 int = (fi7 * (-(fi7)))
                                                                                                                        var v52 int = v25
                                                                                                                        fmt.Println(v10, v43)
                                                                                                                        v37 = v37
                                                                                                                        _ = v51
                                                                                                                        _ = v52
                                                                                                                    case "a":
                                                                                                                        if len(v4) >= 2 {
                                                                                                                            v4 = v4[1:len(v4)]
                                                                                                                        }
                                                                                                                    }
                                                                                                                }
                                                                                                            case -6:
                                                                                                                switch v36 {
                                                                                                                case "":
                                                                                                                    var v53 int32 = v3
                                                                                                                    _ = v53
                                                                                                                case "foo bar":
                                                                                                                    fmt.Println(v35, v33)
                                                                                                                case "\n\t":
                                                                                                                    var v54 float64 = (1e10 * (1e10 + 3.14159))
                                                                                                                    _ = v54
                                                                                                                default:
                                                                                                                    v10 += (v10 * 32767)
                                                                                                                    v37 -= -649
                                                                                                                    v49 -= v9
                                                                                                                    if len(v4) >= 2 {
                                                                                                                        v4 = v4[1:len(v4)]
                                                                                                                    }
                                                                                                                }
                                                                                                            default:
                                                                                                                _ = len(v4)
                                                                                                            }
                                                                                                            _ = v49
                                                                                                        } else {
                                                                                                            _ = len(v4)
                                                                                                            fmt.Println(fi8, v42, v5)
                                                                                                            v55 := GenS4{F0: "", F1: (v42 + (v30 << ((1) & 63)))}
                                                                                                            _ = v55
                                                                                                        }
                                                                                                        switch (fi7 & (v25 * fi7)) {
                                                                                                        case -2147483648:
                                                                                                            v56 := GenS5{F0: 0, F1: v35, F2: -93}
                                                                                                            _ = v56
                                                                                                        case 128:
                                                                                                            if true {
                                                                                                                if ((v10 / (((^(v10))) | 1)) > (v10 * 588)) {
                                                                                                                    v57 := GenS6{F0: -1, F1: 127}
                                                                                                                    _ = v57
                                                                                                                }
                                                                                                                if (false && (667 == 0)) {
                                                                                                                    var v58 string = "a"
                                                                                                                    _ = v58
                                                                                                                }
                                                                                                            }
                                                                                                        }
                                                                                                        _ = v48
                                                                                                    } else {
                                                                                                        v59 := GenS7{F0: 65535, F1: (fi7 | -625)}
                                                                                                        _ = v59
                                                                                                    }
                                                                                                    if (!((42 >= 255) || false)) {
                                                                                                        if false {
                                                                                                            switch (v35 + "") {
                                                                                                            case "a":
                                                                                                                var v60 uint8 = 127
                                                                                                                _ = v60
                                                                                                            }
                                                                                                        } else {
                                                                                                            v5 -= v5
                                                                                                        }
                                                                                                        v61 := GenS1{F0: ((v36 + "hello") + (("hello" + "hello") + v35)), F1: 42, F2: 2}
                                                                                                        if (true && true) {
                                                                                                            v36 = v35
                                                                                                            v62 := GenS3{F0: 2147483647}
                                                                                                            _ = len(v4)
                                                                                                            if len(v4) >= 2 {
                                                                                                                v4 = v4[1:len(v4)]
                                                                                                            }
                                                                                                            v42 = v30
                                                                                                            _ = v62
                                                                                                        } else {
                                                                                                            switch fi7 {
                                                                                                            case 32768:
                                                                                                                v9 -= v9
                                                                                                            default:
                                                                                                                v9 = (v9 + (v9 << ((18446744073709551615) & 63)))
                                                                                                            }
                                                                                                            fmt.Println(v33, v33, v1)
                                                                                                            v63 := GenS6{F0: (v3 | v3), F1: 255}
                                                                                                            v64 := GenS7{F0: 42, F1: fi7}
                                                                                                            _ = v63
                                                                                                            _ = v64
                                                                                                        }
                                                                                                        _ = v61
                                                                                                    } else {
                                                                                                        v29 = v29
                                                                                                    }
                                                                                                    v25 = fi7
                                                                                                }
                                                                                                v4 = append(v4, v25)
                                                                                            } else {
                                                                                                var v65 uint64 = v9
                                                                                                _ = v65
                                                                                            }
                                                                                        } else {
                                                                                            var v66 int8 = int8(v30)
                                                                                            var v67 bool = (v33 || v2)
                                                                                            _ = v66
                                                                                            _ = v67
                                                                                        }
                                                                                        v36 = v35
                                                                                        v68 := GenS6{F0: 0, F1: 0}
                                                                                        var v69 float64 = 0.1
                                                                                        _ = v42
                                                                                        _ = v43
                                                                                        _ = v68
                                                                                        _ = v69
                                                                                    } else {
                                                                                        if len(v4) >= 2 {
                                                                                            v4 = v4[1:len(v4)]
                                                                                        }
                                                                                        v37 = (v37 << ((95) & 63))
                                                                                        var v70 string = ("hello" + v35)
                                                                                        _ = v70
                                                                                    }
                                                                                    var v71 bool = ((0.0 / v29) != -2.25)
                                                                                    _ = v71
                                                                                } else {
                                                                                    var v72 uint8 = 76
                                                                                    var v73 uint8 = v72
                                                                                    if false {
                                                                                        var v74 bool = (!v2)
                                                                                        v2 = false
                                                                                        _ = v74
                                                                                    }
                                                                                    var v75 int32 = (v3 & (v3 % (((^(v3))) | 1)))
                                                                                    v4 = append(v4, (v25 / ((fi8) | 1)))
                                                                                    _ = v72
                                                                                    _ = v73
                                                                                    _ = v75
                                                                                }
                                                                                switch v35 {
                                                                                case "\n\t":
                                                                                    if len(v4) > 0 {
                                                                                        v4[0] = (fi7 & v25)
                                                                                    }
                                                                                }
                                                                            }
                                                                            fmt.Println(v36)
                                                                            _ = v35
                                                                            _ = v36
                                                                            _ = v37
                                                                        }
                                                                    }
                                                                    v76 := GenS8{F0: 42, F1: (^(v3)), F2: (-(v10))}
                                                                    _ = v33
                                                                    _ = v76
                                                                } else {
                                                                    if (1e10 != 0.0) {
                                                                        if v2 {
                                                                            var v77 string = ("" + "a")
                                                                            _ = v77
                                                                        }
                                                                        v9 -= (v9 >> ((65535) & 63))
                                                                    }
                                                                }
                                                            }
                                                        } else {
                                                            v25 = (fi8 - fi8)
                                                            v25 = (^(fi7))
                                                        }
                                                    }
                                                    v1 -= v1
                                                    _ = v30
                                                } else {
                                                    _ = len(v4)
                                                    var v78 int = -2147483648
                                                    v79 := GenS4{F0: "\n\t", F1: 727}
                                                    if v2 {
                                                        v80 := GenS6{F0: v3, F1: 74}
                                                        v78 -= v78
                                                        var v81 float64 = (1e10 - 1e10)
                                                        fmt.Println(fi8)
                                                        _ = v80
                                                        _ = v81
                                                    }
                                                    _ = v78
                                                    _ = v79
                                                }
                                                v10 -= v10
                                                fmt.Println(fi8, v25, fi7)
                                                fmt.Println(v3)
                                                _ = v29
                                            }
                                        }
                                    default:
                                        fmt.Println(fi8, fi7, v9)
                                    }
                                }
                                v25 = (v25 + v25)
                                v82 := GenS2{F0: (v3 << ((2) & 63)), F1: 0, F2: (-1 != 0)}
                                _ = v82
                            case 9223372036854775807:
                                switch v25 {
                                case 4294967295:
                                    v83 := GenS9{F0: 127}
                                    _ = v83
                                case -344:
                                    switch (v25 >> ((255) & 63)) {
                                    case -2147483648:
                                        switch "" {
                                        case "\n\t":
                                            v3 -= 649
                                        }
                                    case 4294967295:
                                        v84 := GenS2{F0: (v3 ^ 1), F1: 0, F2: v2}
                                        _ = v84
                                    case 32767:
                                        switch -1 {
                                        case -2147483648:
                                            var v85 int32 = int32(v25)
                                            _ = v85
                                        case -812:
                                            v86 := GenS3{F0: 1}
                                            _ = v86
                                        }
                                    }
                                default:
                                    if ("" != "hello") {
                                        if len(v4) > 0 {
                                            v4[0] = 0
                                        }
                                        var v87 int = v25
                                        _ = v87
                                    }
                                }
                            default:
                                v88 := GenS4{F0: "a", F1: 255}
                                v9 = (^(v9))
                                fmt.Println(fi8)
                                _ = v88
                            }
                            if v2 {
                                var v89 int8 = v1
                                if v2 {
                                    var v90 int8 = 42
                                    v3 += (v3 * v3)
                                    if (!(!((v90 - 100) == v1))) {
                                        v91 := GenS8{F0: 42, F1: -1, F2: v10}
                                        fmt.Println(v1, v5)
                                        _ = v91
                                    }
                                    _ = v90
                                }
                                var v92 uint64 = v9
                                _ = v89
                                _ = v92
                            }
                            _ = v25
                        }
                        _ = v10
                        _ = v9
                    }
                    _ = fi8
                }
                v93 := GenS1{F0: "foo bar", F1: 42, F2: 93}
                _ = fi7
                _ = v93
            }
        }
    } else {
        v2 = (false || (42 > -232))
        for fi94 := 0; fi94 < 20; fi94++ {
            v95 := GenS10{F0: 0, F1: 172}
            _ = fi94
            _ = v95
        }
        if v2 {
            if len(v4) > 0 {
                v4[0] = -1
            }
            var v96 int16 = 0
            _ = v96
        }
    }
    var v97 int32 = v3
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    for fi98 := 0; fi98 < 8; fi98++ {
        switch (fi98 ^ -2147483648) {
        case 128:
            v1 -= 0
        case 32768:
            v99 := GenS11{F0: ("" + "hello"), F1: 4294967295, F2: true}
            _ = v99
        case -1:
            v100 := GenS11{F0: "a", F1: 984, F2: false}
            _ = v100
        default:
            v3 -= (-(v3))
        }
        _ = fi98
    }
    if v2 {
        fmt.Println(v2)
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        for fi101 := 0; fi101 < 4; fi101++ {
            var v102 bool = v2
            _ = fi101
            _ = v102
        }
    } else {
        var v103 float64 = float64(v5)
        if len(v4) > 0 {
            v4[0] = -1
        }
        _ = v103
    }
    for fi104 := 0; fi104 < 18; fi104++ {
        fmt.Println(v1, fi104)
        _ = fi104
    }
    v5 -= v5
    if true {
        _ = len(v4)
        switch -9223372036854775808 {
        case -724:
            if (!v2) {
                var v105 uint8 = 128
                _ = v105
            } else {
                v5 -= 1e10
            }
        case -2147483648:
            v2 = true
        }
        v5 = (0.0 / v5)
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if len(v4) > 0 {
        v4[0] = 4294967295
    }
    v2 = (1 <= -1)
    if v2 {
        for fi106 := 0; fi106 < 14; fi106++ {
            if (!false) {
                for fi107 := 0; fi107 < 17; fi107++ {
                    var v108 int32 = v97
                    _ = fi107
                    _ = v108
                }
                for fi109 := 0; fi109 < 15; fi109++ {
                    v4 = append(v4, 2147483647)
                    v110 := GenS8{F0: 152, F1: (v97 | v97), F2: int16(fi109)}
                    switch "foo bar" {
                    case "foo bar":
                        fmt.Println(v2)
                    case "\n\t":
                        _ = len(v4)
                    }
                    _ = fi109
                    _ = v110
                }
            }
            v97 += (^(v3))
            var v111 bool = (255 != 0)
            _ = fi106
            _ = v111
        }
    }
    if true {
        _ = len(v4)
    } else {
        var v112 float32 = v5
        for fi113 := 0; fi113 < 11; fi113++ {
            for fi114 := 0; fi114 < 16; fi114++ {
                if v2 {
                    v5 += 3.14159
                    var v115 float32 = v112
                    v1 = -88
                    _ = v115
                }
                _ = fi114
            }
            _ = fi113
        }
        if v2 {
            v5 = ((((((((((((0.0 / (v112 / ((v112 + (v5 / v112)) / float32(float64(v5))))) * v5) * 1e10) + (v112 / (1e10 * 0.1))) + v5) / 3.14159) + (v112 - v112)) * v112) * ((float32((3.14159 + float64(((((3.14159 / v5) / v112) / ((v112 * -2.25) - v112)) * v112)))) - 1e10) + float32(0.0))) + (v5 / float32(1e10))) * v5) * -2.25)
            for fi116 := 0; fi116 < 8; fi116++ {
                v1 -= (v1 / ((v1) | 1))
                var v117 uint64 = 18446744073709551615
                for fi118 := 0; fi118 < 15; fi118++ {
                    switch ("foo bar" + "a") {
                    case "hello":
                        v112 -= -2.25
                    case "":
                        v119 := GenS5{F0: 42, F1: "\n\t", F2: (^(v1))}
                        _ = v119
                    }
                    if v2 {
                        var v120 float32 = -2.25
                        v3 -= -78
                        _ = v120
                    } else {
                        v4 = append(v4, (fi116 & (^(fi118))))
                        var v121 float64 = 1e10
                        _ = v121
                    }
                    v1 = (v1 ^ -87)
                    _ = fi118
                }
                _ = fi116
                _ = v117
            }
            for fi122 := 0; fi122 < 3; fi122++ {
                var v123 string = "\n\t"
                _ = len(v4)
                for fi124 := 0; fi124 < 10; fi124++ {
                    fmt.Println(v1, fi124)
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = fi124
                }
                _ = fi122
                _ = v123
            }
        }
        _ = v112
    }
    v5 = (1.5 + 3.14159)
    if (128 > 1) {
        v4 = append(v4, 704)
    } else {
        var v125 int8 = v1
        v126 := GenS5{F0: 18446744073709551615, F1: "", F2: 0}
        fmt.Println(v2, v1)
        _ = v125
        _ = v126
    }
    if true {
        var v127 int64 = -1
        _ = v127
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v97
}
