package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int16 }
type GenS2 struct { F0 uint8; F1 uint8; F2 bool }
type GenS3 struct { F0 bool; F1 string }
type GenS4 struct { F0 string; F1 uint; F2 uint64 }
type GenS5 struct { F0 int16; F1 int; F2 int16 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint8 }
type GenS8 struct { F0 uint32; F1 uint32; F2 int32 }
type GenS9 struct { F0 int16; F1 uint; F2 uint64 }
type GenS10 struct { F0 string; F1 uint16 }
type GenS11 struct { F0 int32 }
type GenS12 struct { F0 string; F1 uint; F2 int }

func main() {
    var v1 float32 = 1e10
    v2 := []int{}
    var v3 float64 = (3.14159 + 1e10)
    var v4 uint8 = 42
    var v5 uint32 = 0
    var v6 int16 = -425
    v6 -= v6
    switch "foo bar" {
    case "\n\t":
        var v7 uint32 = v5
        _ = v7
    }
    v8 := GenS1{F0: (^(v6)), F1: 42}
    fmt.Println(v4, v3, v4)
    for fi9 := 0; fi9 < 14; fi9++ {
        if (1e10 != v1) {
            for fi10 := 0; fi10 < 8; fi10++ {
                var v11 string = ("a" + "")
                _ = fi10
                _ = v11
            }
        } else {
            v6 = -956
            var v12 int8 = 12
            v13 := GenS2{F0: (v4 | v4), F1: (^(v4)), F2: true}
            _ = v12
            _ = v13
        }
        _ = fi9
    }
    if false {
        for fi14 := 0; fi14 < 1; fi14++ {
            if false {
                var v15 bool = (((false || (2147483647 < (fi14 / (((fi14 ^ fi14)) | 1)))) && (97 == -628)) || true)
                v16 := GenS2{F0: 42, F1: v4, F2: false}
                _ = v15
                _ = v16
            } else {
                v17 := GenS1{F0: (v6 / ((v6) | 1)), F1: (v6 & 32767)}
                fmt.Println(v4)
                _ = v17
            }
            v5 += (v5 % ((v5) | 1))
            v6 = v6
            _ = fi14
        }
    } else {
        switch 32768 {
        case 32767:
            if (42 == 2147483647) {
                v2 = append(v2, 2147483647)
                _ = len(v2)
            }
        case 127:
            if false {
                for fi18 := 0; fi18 < 13; fi18++ {
                    for fi19 := 0; fi19 < 16; fi19++ {
                        if len(v2) > 0 {
                            v2[0] = fi19
                        }
                        if false {
                            v1 -= 0.0
                            fmt.Println(v5)
                            if (!((!true) || true)) {
                                if (!false) {
                                    switch -2147483648 {
                                    case 728:
                                        var v20 uint8 = 2
                                        _ = v20
                                    case 128:
                                        switch ("a" + ("a" + ("foo bar" + ("\n\t" + ("hello" + "\n\t"))))) {
                                        case "foo bar":
                                            v4 = (v4 >> ((446) & 63))
                                        case "hello":
                                            if (int8(v6) <= 100) {
                                                v5 = (v5 - (v5 + (^(v5))))
                                                v2 = append(v2, (fi18 / ((-1) | 1)))
                                                fmt.Println(v5, v1, fi18)
                                            } else {
                                                if true {
                                                    var v21 float64 = ((float64(v1) + 0.0) - 1.5)
                                                    if len(v2) > 0 {
                                                        v2[0] = fi18
                                                    }
                                                    switch "hello" {
                                                    case "":
                                                        if false {
                                                            if false {
                                                                switch (("" + "hello") + "a") {
                                                                case "foo bar":
                                                                    var v22 bool = false
                                                                    _ = v22
                                                                case "hello":
                                                                    switch "" {
                                                                    case "a":
                                                                        fmt.Println(v3, v1)
                                                                    case "foo bar":
                                                                        var v23 int16 = int16(v5)
                                                                        v23 -= v6
                                                                        _ = v23
                                                                    default:
                                                                        _ = len(v2)
                                                                    }
                                                                case "a":
                                                                    v24 := GenS2{F0: 1, F1: v4, F2: (!(!true))}
                                                                    _ = v24
                                                                default:
                                                                    v25 := GenS1{F0: (v6 & (v6 ^ (v6 << ((1) & 63)))), F1: 1}
                                                                    _ = v25
                                                                }
                                                                fmt.Println(v6)
                                                                switch (fi19 & fi18) {
                                                                case 0:
                                                                    if false {
                                                                        v26 := GenS3{F0: (127 != 42), F1: ""}
                                                                        _ = v26
                                                                    }
                                                                    if false {
                                                                        var v27 bool = false
                                                                        v3 = ((float64(v1) / (-2.25 / v3)) * (v21 / (v21 / 1e10)))
                                                                        v28 := GenS3{F0: (v5 > v5), F1: "foo bar"}
                                                                        _ = v27
                                                                        _ = v28
                                                                    } else {
                                                                        fmt.Println(fi19, fi18)
                                                                        v21 += v21
                                                                    }
                                                                    var v29 int64 = 855
                                                                    _ = v29
                                                                case 42:
                                                                    switch (("\n\t" + "hello") + "a") {
                                                                    case "hello":
                                                                        if len(v2) > 0 {
                                                                            v2[0] = fi19
                                                                        }
                                                                    case "\n\t":
                                                                        var v30 float64 = v3
                                                                        _ = v30
                                                                    default:
                                                                        fmt.Println(v4, v6, v3)
                                                                    }
                                                                case 127:
                                                                    var v31 uint64 = 1
                                                                    _ = v31
                                                                default:
                                                                    if true {
                                                                        _ = len(v2)
                                                                        v21 -= (v21 + (0.0 - 3.14159))
                                                                        if false {
                                                                            switch 1 {
                                                                            case -843:
                                                                                if false {
                                                                                    var v32 int64 = -1
                                                                                    fmt.Println(v32, v6)
                                                                                    _ = v32
                                                                                }
                                                                            default:
                                                                                if len(v2) >= 2 {
                                                                                    v2 = v2[1:len(v2)]
                                                                                }
                                                                            }
                                                                            var v33 int16 = (-(v6))
                                                                            if false {
                                                                                v34 := GenS4{F0: ("" + "\n\t"), F1: 694, F2: 2}
                                                                                if true {
                                                                                    if false {
                                                                                        if len(v2) >= 2 {
                                                                                            v2 = v2[1:len(v2)]
                                                                                        }
                                                                                        var v35 uint16 = 32767
                                                                                        _ = v35
                                                                                    } else {
                                                                                        var v36 uint32 = v5
                                                                                        var v37 uint16 = 32767
                                                                                        if false {
                                                                                            switch "" {
                                                                                            case "\n\t":
                                                                                                v21 = (float64(-2.25) - ((v3 / -2.25) * ((1e10 - (1e10 - (0.1 - 1e10))) + v3)))
                                                                                            }
                                                                                            v38 := GenS4{F0: ("hello" + "a"), F1: 452, F2: 442}
                                                                                            fmt.Println(v1, fi19, v1)
                                                                                            _ = v38
                                                                                        } else {
                                                                                            _ = len(v2)
                                                                                        }
                                                                                        _ = v36
                                                                                        _ = v37
                                                                                    }
                                                                                    fmt.Println(v33, v21, fi19)
                                                                                    if (true && (true && true)) {
                                                                                        v21 -= 1.5
                                                                                        if true {
                                                                                            switch -1 {
                                                                                            case 2147483647:
                                                                                                if ((fi19 <= fi18) || false) {
                                                                                                    if (!false) {
                                                                                                        fmt.Println(v33)
                                                                                                        fmt.Println(fi18, fi18, v1)
                                                                                                        if len(v2) >= 2 {
                                                                                                            v2 = v2[1:len(v2)]
                                                                                                        }
                                                                                                    } else {
                                                                                                        switch (("a" + "\n\t") + ("hello" + "")) {
                                                                                                        case "\n\t":
                                                                                                            fmt.Println(v21)
                                                                                                            var v39 uint8 = v4
                                                                                                            fmt.Println(fi18, v3)
                                                                                                            _ = v39
                                                                                                        case "a":
                                                                                                            fmt.Println(v5, v33, v1)
                                                                                                        }
                                                                                                        var v40 int32 = -2147483648
                                                                                                        var v41 int32 = int32(fi19)
                                                                                                        _ = v40
                                                                                                        _ = v41
                                                                                                    }
                                                                                                    var v42 int32 = 780
                                                                                                    _ = v42
                                                                                                }
                                                                                            case -1:
                                                                                                fmt.Println(fi19)
                                                                                            }
                                                                                            if true {
                                                                                                v3 = v21
                                                                                                _ = len(v2)
                                                                                                fmt.Println(fi19, v5, v1)
                                                                                            } else {
                                                                                                v5 -= v5
                                                                                                var v43 string = ""
                                                                                                if false {
                                                                                                    if (255 <= 18446744073709551615) {
                                                                                                        var v44 int32 = -1
                                                                                                        _ = v44
                                                                                                    } else {
                                                                                                        var v45 float64 = 3.14159
                                                                                                        if len(v2) > 0 {
                                                                                                            v2[0] = (fi19 + (-(fi18)))
                                                                                                        }
                                                                                                        var v46 int = (fi18 - (fi19 - 9223372036854775807))
                                                                                                        var v47 int = fi19
                                                                                                        fmt.Println(v21, v6, v3)
                                                                                                        _ = v45
                                                                                                        _ = v46
                                                                                                        _ = v47
                                                                                                    }
                                                                                                    switch (-(fi19)) {
                                                                                                    case -9223372036854775808:
                                                                                                        switch fi19 {
                                                                                                        case 4294967295:
                                                                                                            fmt.Println(v6, fi19, v6)
                                                                                                        case 127:
                                                                                                            v2 = append(v2, (fi18 / (((fi18 % ((fi18) | 1))) | 1)))
                                                                                                        }
                                                                                                    case 1:
                                                                                                        fmt.Println(v43)
                                                                                                    default:
                                                                                                        var v48 int16 = v33
                                                                                                        _ = v48
                                                                                                    }
                                                                                                    var v49 float64 = (v3 * v21)
                                                                                                    _ = v49
                                                                                                }
                                                                                                v50 := GenS3{F0: true, F1: v43}
                                                                                                v51 := GenS3{F0: false, F1: v43}
                                                                                                _ = v43
                                                                                                _ = v50
                                                                                                _ = v51
                                                                                            }
                                                                                        }
                                                                                        v52 := GenS4{F0: (("" + "\n\t") + "\n\t"), F1: 255, F2: 18446744073709551615}
                                                                                        var v53 uint32 = v5
                                                                                        _ = v52
                                                                                        _ = v53
                                                                                    }
                                                                                }
                                                                                switch fi18 {
                                                                                case 4294967295:
                                                                                    v6 = v33
                                                                                case 2147483647:
                                                                                    v2 = append(v2, fi18)
                                                                                default:
                                                                                    v54 := GenS1{F0: v33, F1: v33}
                                                                                    _ = v54
                                                                                }
                                                                                v55 := GenS5{F0: 601, F1: 1, F2: (v6 % (((v6 | v33)) | 1))}
                                                                                if len(v2) > 0 {
                                                                                    v2[0] = (-(fi19))
                                                                                }
                                                                                v1 -= 1.5
                                                                                if len(v2) >= 2 {
                                                                                    v2 = v2[1:len(v2)]
                                                                                }
                                                                                v56 := GenS2{F0: (v4 - v4), F1: 0, F2: true}
                                                                                _ = v34
                                                                                _ = v55
                                                                                _ = v56
                                                                            } else {
                                                                                v57 := GenS6{F0: (v4 % (((v4 & v4)) | 1))}
                                                                                v33 = v6
                                                                                v1 += ((v1 + (1.5 / v1)) / (0.1 * v1))
                                                                                _ = v57
                                                                            }
                                                                            _ = v33
                                                                        } else {
                                                                            v58 := GenS6{F0: (v4 % ((42) | 1))}
                                                                            _ = v58
                                                                        }
                                                                        fmt.Println(v21)
                                                                    } else {
                                                                        var v59 int = -488
                                                                        v60 := GenS4{F0: "a", F1: 10, F2: 4294967295}
                                                                        _ = v59
                                                                        _ = v60
                                                                    }
                                                                }
                                                                if len(v2) >= 2 {
                                                                    v2 = v2[1:len(v2)]
                                                                }
                                                                var v61 int = 206
                                                                _ = v61
                                                            } else {
                                                                v4 = v4
                                                            }
                                                            v5 = (v5 & (v5 * (v5 & 527)))
                                                            v62 := GenS7{F0: v4}
                                                            var v63 float64 = v21
                                                            if len(v2) > 0 {
                                                                v2[0] = (fi18 | (^(fi18)))
                                                            }
                                                            _ = v62
                                                            _ = v63
                                                        } else {
                                                            _ = len(v2)
                                                        }
                                                    case "foo bar":
                                                        v5 = 727
                                                    case "hello":
                                                        v21 = float64(v1)
                                                    }
                                                    _ = v21
                                                }
                                                v6 = v6
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                            }
                                        case "":
                                            v64 := GenS8{F0: (^(v5)), F1: (v5 ^ v5), F2: 2147483647}
                                            var v65 string = ("\n\t" + "\n\t")
                                            _ = v64
                                            _ = v65
                                        default:
                                            if true {
                                                switch "\n\t" {
                                                case "\n\t":
                                                    if false {
                                                        if false {
                                                            v66 := GenS9{F0: v6, F1: 0, F2: 18446744073709551615}
                                                            if (992 == 2) {
                                                                fmt.Println(v3)
                                                                var v67 string = ""
                                                                _ = v67
                                                            } else {
                                                                if (9223372036854775807 != 42) {
                                                                    _ = len(v2)
                                                                    var v68 int = fi18
                                                                    if false {
                                                                        fmt.Println(v5, v4)
                                                                    } else {
                                                                        if len(v2) > 0 {
                                                                            v2[0] = (v68 >> ((4294967295) & 63))
                                                                        }
                                                                        v69 := GenS10{F0: ("a" + "foo bar"), F1: 1}
                                                                        _ = v69
                                                                    }
                                                                    v1 = (1.5 * (v1 - 1e10))
                                                                    _ = v68
                                                                } else {
                                                                    v70 := GenS11{F0: 42}
                                                                    if true {
                                                                        if len(v2) > 0 {
                                                                            v2[0] = fi18
                                                                        }
                                                                    }
                                                                    switch (("" + ("a" + "hello")) + ("\n\t" + "hello")) {
                                                                    case "":
                                                                        if (1.5 > 0.1) {
                                                                            var v71 int32 = -756
                                                                            var v72 int64 = -1
                                                                            var v73 uint16 = 42
                                                                            _ = v71
                                                                            _ = v72
                                                                            _ = v73
                                                                        } else {
                                                                            fmt.Println(v5)
                                                                        }
                                                                    case "hello":
                                                                        var v74 uint = 255
                                                                        _ = v74
                                                                    }
                                                                    fmt.Println(v5, v4)
                                                                    var v75 uint64 = 413
                                                                    var v76 float32 = 0.1
                                                                    _ = v70
                                                                    _ = v75
                                                                    _ = v76
                                                                }
                                                            }
                                                            _ = v66
                                                        }
                                                    }
                                                case "":
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                }
                                                v77 := GenS12{F0: "\n\t", F1: 1, F2: 2147483647}
                                                _ = v77
                                            } else {
                                                v2 = append(v2, (fi18 << ((4294967295) & 63)))
                                                v78 := GenS2{F0: (v4 ^ 42), F1: (v4 >> ((399) & 63)), F2: (false || true)}
                                                _ = v78
                                            }
                                        }
                                    case -9223372036854775808:
                                        if len(v2) > 0 {
                                            v2[0] = fi19
                                        }
                                    default:
                                        v1 = (v1 + v1)
                                    }
                                    var v79 int16 = v6
                                    _ = v79
                                }
                            } else {
                                v3 = v3
                            }
                        } else {
                            v5 += (^(v5))
                            if true {
                                var v80 uint32 = 4294967295
                                _ = v80
                            } else {
                                v81 := GenS9{F0: (v6 / ((-1) | 1)), F1: 42, F2: 0}
                                var v82 string = "hello"
                                var v83 string = (v82 + "\n\t")
                                _ = v81
                                _ = v82
                                _ = v83
                            }
                            if ((!(!false)) && false) {
                                v1 -= v1
                                v4 = 0
                            }
                            fmt.Println(fi18, v1)
                            var v84 float64 = 0.1
                            _ = v84
                        }
                        _ = fi19
                    }
                    var v85 uint32 = (v5 % (((^(v5))) | 1))
                    _ = fi18
                    _ = v85
                }
                if len(v2) > 0 {
                    v2[0] = 0
                }
            }
        }
        v1 -= float32(0.1)
        var v86 uint16 = 42
        _ = v86
    }
    v87 := GenS8{F0: (v5 * (v5 & v5)), F1: 42, F2: 1}
    v5 = (v5 / ((2147483647) | 1))
    v4 += v4
    if (!false) {
        v3 = (v3 - 0.1)
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    for fi88 := 0; fi88 < 14; fi88++ {
        v6 += (-(v6))
        var v89 float64 = 3.14159
        var v90 float64 = 0.0
        _ = fi88
        _ = v89
        _ = v90
    }
    _ = len(v2)
    v6 = 42
    var v91 uint = 255
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v87
    _ = v91
}
