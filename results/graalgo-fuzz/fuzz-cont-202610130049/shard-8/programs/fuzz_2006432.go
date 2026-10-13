package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 string }
type GenS2 struct { F0 string; F1 uint16 }
type GenS3 struct { F0 int; F1 int32 }
type GenS4 struct { F0 int16; F1 uint8; F2 int32 }
type GenS5 struct { F0 int64; F1 int }
type GenS6 struct { F0 uint64; F1 int8 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 int64; F1 int32 }
type GenS10 struct { F0 int }
type GenS11 struct { F0 bool; F1 uint; F2 string }
type GenS12 struct { F0 uint }
type GenS13 struct { F0 uint }
type GenS14 struct { F0 int32; F1 string; F2 int }

func main() {
    var v1 int8 = 42
    var v2 uint = 2
    var v3 int64 = -9223372036854775808
    var v4 float32 = 3.14159
    v4 -= 1e10
    v5 := []int{32768, 139, 4294967295}
    switch 42 {
    case 128:
        for fi6 := 0; fi6 < 2; fi6++ {
            v2 -= 42
            _ = fi6
        }
    default:
        fmt.Println(v3, v4)
    }
    v2 -= 0
    if true {
        if false {
            if (true && false) {
                v7 := GenS1{F0: v1, F1: (v3 & v3), F2: "foo bar"}
                for fi8 := 0; fi8 < 10; fi8++ {
                    var v9 int8 = v1
                    v10 := GenS1{F0: v9, F1: v3, F2: "hello"}
                    _ = fi8
                    _ = v10
                    _ = v9
                }
                _ = v7
            }
            if true {
                if true {
                    if len(v5) > 0 {
                        v5[0] = 127
                    }
                    v3 = (v3 / (((v3 + v3)) | 1))
                } else {
                    var v11 uint32 = 260
                    var v12 float32 = v4
                    _ = v11
                    _ = v12
                }
                v4 += float32(3.14159)
                var v13 int = 127
                _ = v13
            }
            v14 := GenS2{F0: ("" + ("hello" + "hello")), F1: 2}
            _ = v14
        } else {
            var v15 uint64 = 922
            v16 := GenS3{F0: 127, F1: -256}
            _ = len(v5)
            _ = v15
            _ = v16
        }
        fmt.Println(v3, v4)
        if len(v5) > 0 {
            v5[0] = -86
        }
    } else {
        v17 := GenS4{F0: 1, F1: 127, F2: 1}
        _ = len(v5)
        v4 = float32((-2.25 * 1e10))
        _ = v17
    }
    switch "foo bar" {
    case "a":
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
    }
    fmt.Println(v3, v3)
    if (v3 != (v3 ^ -1)) {
        for fi18 := 0; fi18 < 12; fi18++ {
            switch ("" + ("a" + "")) {
            case "hello":
                var v19 uint32 = 0
                _ = v19
            case "a":
                v4 = v4
            case "foo bar":
                v4 = (float32((-2.25 + 0.0)) + v4)
            default:
                if len(v5) > 0 {
                    v5[0] = (fi18 | 127)
                }
            }
            v20 := GenS1{F0: 0, F1: (v3 & v3), F2: (((("foo bar" + "") + ("foo bar" + "\n\t")) + "a") + "foo bar")}
            _ = fi18
            _ = v20
        }
        var v21 uint8 = 166
        for fi22 := 0; fi22 < 10; fi22++ {
            v23 := GenS5{F0: v3, F1: fi22}
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            for fi24 := 0; fi24 < 1; fi24++ {
                fmt.Println(v21, v4, fi22)
                _ = fi24
            }
            _ = fi22
            _ = v23
        }
        _ = v21
    }
    for fi25 := 0; fi25 < 19; fi25++ {
        v2 += (v2 & v2)
        v26 := GenS3{F0: (-(fi25)), F1: 422}
        var v27 int64 = v3
        _ = fi25
        _ = v26
        _ = v27
    }
    var v28 bool = true
    for fi29 := 0; fi29 < 11; fi29++ {
        if v28 {
            var v30 uint16 = 2
            for fi31 := 0; fi31 < 14; fi31++ {
                v5 = append(v5, (fi31 ^ -1))
                if v28 {
                    var v32 float32 = v4
                    if v28 {
                        if (!true) {
                            fmt.Println(v4)
                            if (!(!(v28 || (!v28)))) {
                                var v33 uint64 = 1
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                switch "hello" {
                                case "\n\t":
                                    v5 = append(v5, -2147483648)
                                    v34 := GenS5{F0: (v3 % ((9223372036854775807) | 1)), F1: 0}
                                    _ = v34
                                case "hello":
                                    var v35 int32 = 967
                                    _ = v35
                                case "a":
                                    switch (("a" + "\n\t") + ("" + ("foo bar" + "\n\t"))) {
                                    case "\n\t":
                                        var v36 uint16 = (v30 * 2)
                                        _ = v36
                                    case "hello":
                                        if true {
                                            if (((1.5 < 0.1) || v28) && v28) {
                                                if v28 {
                                                    var v37 int32 = 42
                                                    v5 = append(v5, (fi31 % ((fi31) | 1)))
                                                    var v38 string = "hello"
                                                    if v28 {
                                                        fmt.Println(v28, v37)
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                    }
                                                    _ = v37
                                                    _ = v38
                                                } else {
                                                    fmt.Println(v4)
                                                    v33 = 0
                                                }
                                                v39 := GenS4{F0: 856, F1: 85, F2: 2147483647}
                                                v40 := GenS6{F0: v33, F1: int8(v30)}
                                                _ = v39
                                                _ = v40
                                            }
                                            if (32768 > int(v1)) {
                                                v3 += v3
                                                fmt.Println(v32, v4)
                                            }
                                        }
                                    case "a":
                                        if (-192 == -965) {
                                            if v28 {
                                                var v41 uint32 = uint32(v30)
                                                var v42 uint64 = (v33 >> ((uint(v33)) & 63))
                                                _ = v41
                                                _ = v42
                                            } else {
                                                v1 = 103
                                                v43 := GenS3{F0: fi31, F1: -233}
                                                if v28 {
                                                    if (false || v28) {
                                                        v4 += v32
                                                        fmt.Println(v30, v28, v3)
                                                        if v28 {
                                                            var v44 int32 = 0
                                                            _ = v44
                                                        } else {
                                                            fmt.Println(v30, fi31, fi31)
                                                        }
                                                    }
                                                } else {
                                                    v32 = 1.5
                                                }
                                                if true {
                                                    v45 := GenS7{F0: 42}
                                                    if len(v5) > 0 {
                                                        v5[0] = fi29
                                                    }
                                                    if v28 {
                                                        v5 = append(v5, fi29)
                                                        v30 = v30
                                                    }
                                                    _ = v45
                                                }
                                                fmt.Println(v4, fi29)
                                                _ = v43
                                            }
                                            fmt.Println(v32, v28, v28)
                                            v5 = append(v5, (fi31 % (((fi29 / (((fi31 >> ((v2) & 63))) | 1))) | 1)))
                                        } else {
                                            var v46 int = fi29
                                            if (v30 >= 0) {
                                                fmt.Println(v1)
                                                var v47 float32 = float32(-2.25)
                                                v48 := GenS2{F0: "foo bar", F1: v30}
                                                _ = v47
                                                _ = v48
                                            }
                                            _ = v46
                                        }
                                    default:
                                        var v49 uint = v2
                                        _ = v49
                                    }
                                }
                                var v50 string = (((("hello" + "foo bar") + ("foo bar" + "foo bar")) + "foo bar") + ("hello" + (("a" + ("hello" + "\n\t")) + "hello")))
                                switch fi29 {
                                case 32767:
                                    if v28 {
                                        v30 = v30
                                        switch (fi29 | (fi29 << ((v2) & 63))) {
                                        case -137:
                                            v1 -= (-(v1))
                                        case 32767:
                                            fmt.Println(v33, v33, v1)
                                        default:
                                            if (!v28) {
                                                v2 -= 4294967295
                                                var v51 float64 = 3.14159
                                                _ = v51
                                            }
                                        }
                                        v52 := GenS7{F0: 227}
                                        v53 := GenS4{F0: 32767, F1: 255, F2: -624}
                                        v3 += -9223372036854775808
                                        _ = v52
                                        _ = v53
                                    }
                                case 95:
                                    fmt.Println(v28, v3, v3)
                                    v33 = uint64(v2)
                                    v54 := GenS8{F0: 1}
                                    _ = v54
                                }
                                v2 = 65535
                                if v28 {
                                    v30 = (v30 / (((v30 * (v30 % ((65535) | 1)))) | 1))
                                    v5 = append(v5, fi29)
                                } else {
                                    if ((v50 + "a") >= v50) {
                                        var v55 uint64 = 379
                                        if ((^(v2)) < (v2 | 65535)) {
                                            v5 = append(v5, 1)
                                            v55 += v33
                                            switch (fi31 % ((fi31) | 1)) {
                                            case -2147483648:
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                            case -9223372036854775808:
                                                if v28 {
                                                    switch -828 {
                                                    case 42:
                                                        var v56 int64 = v3
                                                        _ = v56
                                                    case -68:
                                                        switch "a" {
                                                        case "hello":
                                                            v57 := GenS9{F0: (v3 * v3), F1: 42}
                                                            _ = v57
                                                        case "":
                                                            v4 = (v32 / ((3.14159 / float32(float64(-2.25))) - v4))
                                                        default:
                                                            fmt.Println(v33)
                                                        }
                                                    default:
                                                        var v58 float32 = v4
                                                        _ = v58
                                                    }
                                                }
                                            }
                                        } else {
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                        }
                                        _ = v55
                                    } else {
                                        v59 := GenS10{F0: (fi29 & fi29)}
                                        _ = v59
                                    }
                                    if v28 {
                                        v4 += (((v4 / 3.14159) + float32(-2.25)) / v4)
                                        if true {
                                            var v60 uint8 = 2
                                            if true {
                                                var v61 int32 = int32(v30)
                                                fmt.Println(v60, v1)
                                                _ = v61
                                            }
                                            _ = v60
                                        } else {
                                            var v62 int = 42
                                            v63 := GenS4{F0: -32768, F1: 255, F2: 42}
                                            switch v50 {
                                            case "hello":
                                                if (uint8(v33) < 18) {
                                                    v5 = append(v5, 0)
                                                }
                                            }
                                            v64 := GenS5{F0: -17, F1: (fi29 / ((2147483647) | 1))}
                                            _ = len(v5)
                                            fmt.Println(v62, v3)
                                            var v65 uint = (^(v2))
                                            _ = v62
                                            _ = v63
                                            _ = v64
                                            _ = v65
                                        }
                                        if (-2.25 >= float32(0.0)) {
                                            var v66 uint = v2
                                            if true {
                                                var v67 int = fi29
                                                v5 = append(v5, v67)
                                                if v28 {
                                                    if len(v5) >= 2 {
                                                        v5 = v5[1:len(v5)]
                                                    }
                                                    if true {
                                                        v68 := GenS11{F0: false, F1: 42, F2: (("a" + v50) + v50)}
                                                        v4 += 3.14159
                                                        if ((v3 + v3) != int64(v2)) {
                                                            v66 -= (^(v66))
                                                            var v69 uint = (v66 + 0)
                                                            _ = v69
                                                        }
                                                        _ = v68
                                                    } else {
                                                        v2 = v2
                                                    }
                                                    v50 = "\n\t"
                                                }
                                                _ = v67
                                            }
                                            switch ("hello" + v50) {
                                            case "foo bar":
                                                _ = len(v5)
                                            case "\n\t":
                                                v32 -= (v4 + v32)
                                            case "a":
                                                if len(v5) > 0 {
                                                    v5[0] = (fi29 | -1)
                                                }
                                            }
                                            var v70 uint = v2
                                            _ = v66
                                            _ = v70
                                        }
                                    } else {
                                        v71 := GenS12{F0: (v2 + v2)}
                                        var v72 string = "a"
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                        _ = v71
                                        _ = v72
                                    }
                                    if len(v5) > 0 {
                                        v5[0] = (fi29 & fi31)
                                    }
                                }
                                if len(v5) > 0 {
                                    v5[0] = fi29
                                }
                                _ = v33
                                _ = v50
                            }
                            switch "" {
                            case "a":
                                var v73 uint = 251
                                _ = v73
                            default:
                                v5 = append(v5, int(v1))
                            }
                        }
                        if v28 {
                            v32 = v4
                            v30 += (v30 | (v30 ^ v30))
                        } else {
                            if v28 {
                                var v74 int16 = 1000
                                _ = len(v5)
                                v5 = append(v5, (fi31 ^ (-(fi29))))
                                _ = v74
                            } else {
                                if len(v5) > 0 {
                                    v5[0] = 2147483647
                                }
                                var v75 int8 = v1
                                var v76 int64 = (v3 ^ v3)
                                _ = len(v5)
                                v75 = v75
                                _ = v75
                                _ = v76
                            }
                            if len(v5) > 0 {
                                v5[0] = 128
                            }
                            var v77 int64 = 9223372036854775807
                            _ = v77
                        }
                        if len(v5) > 0 {
                            v5[0] = fi29
                        }
                    } else {
                        var v78 uint64 = 42
                        v79 := GenS13{F0: 419}
                        var v80 float64 = 0.1
                        v5 = append(v5, (fi29 + fi29))
                        v81 := GenS1{F0: int8(v30), F1: v3, F2: ("\n\t" + "hello")}
                        v3 = (^(v3))
                        _ = len(v5)
                        switch ("a" + "a") {
                        case "\n\t":
                            v30 = v30
                        default:
                            var v82 int = (fi31 % (((fi29 >> ((v2) & 63))) | 1))
                            _ = v82
                        }
                        _ = v78
                        _ = v79
                        _ = v80
                        _ = v81
                    }
                    _ = v32
                }
                if false {
                    v4 = ((-2.25 - 0.1) + ((1.5 - v4) / 3.14159))
                    v2 -= v2
                    v83 := GenS10{F0: fi31}
                    if v28 {
                        var v84 uint = 0
                        _ = v84
                    } else {
                        _ = len(v5)
                    }
                    _ = v83
                }
                _ = fi31
            }
            _ = v30
        } else {
            v2 = v2
        }
        _ = fi29
    }
    var v85 uint64 = 42
    if (0.1 <= float64(1.5)) {
        var v86 int32 = -2147483648
        _ = v86
    }
    switch -9223372036854775808 {
    case 513:
        if len(v5) > 0 {
            v5[0] = 32767
        }
    case 886:
        if false {
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
            var v87 string = "\n\t"
            v5 = append(v5, 815)
            _ = v87
        } else {
            for fi88 := 0; fi88 < 5; fi88++ {
                _ = len(v5)
                if (!true) {
                    var v89 float64 = float64((1.5 - 0.0))
                    v85 -= 0
                    v2 -= 673
                    _ = v89
                } else {
                    v2 += v2
                    v5 = append(v5, fi88)
                }
                v28 = v28
                _ = fi88
            }
            var v90 uint = v2
            switch 9223372036854775807 {
            case 9223372036854775807:
                v91 := GenS2{F0: "\n\t", F1: 163}
                _ = v91
            case 2147483647:
                v92 := GenS1{F0: v1, F1: v3, F2: "foo bar"}
                _ = v92
            case 32768:
                for fi93 := 0; fi93 < 16; fi93++ {
                    var v94 uint = v2
                    for fi95 := 0; fi95 < 12; fi95++ {
                        v96 := GenS14{F0: 2147483647, F1: ("foo bar" + "hello"), F2: (^(fi95))}
                        _ = fi95
                        _ = v96
                    }
                    _ = fi93
                    _ = v94
                }
            }
            _ = v90
        }
    default:
        var v97 uint16 = 617
        _ = v97
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v85:", v85)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v85
}
