package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 int8; F1 uint16; F2 bool }
type GenS3 struct { F0 int16; F1 int16; F2 int64 }
type GenS4 struct { F0 uint64; F1 uint64 }
type GenS5 struct { F0 int64; F1 string }
type GenS6 struct { F0 uint16; F1 int16 }
type GenS7 struct { F0 bool; F1 uint8 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 uint8; F1 int64 }
type GenS10 struct { F0 int; F1 int16; F2 string }

func main() {
    v1 := make(map[string]int)
    var v2 uint64 = 42
    var v3 uint32 = 579
    v4 := []int{-625, 4294967295, -9223372036854775808, -9223372036854775808}
    v5 := []string{}
    var v6 uint = 959
    v6 += v6
    for fi7 := 0; fi7 < 1; fi7++ {
        v6 = 42
        _ = fi7
    }
    var v8 uint64 = v2
    v8 += (v8 * v2)
    switch "\n\t" {
    case "\n\t":
        v9 := v1["hello"]
        _ = v9
    default:
        var v10 uint8 = 0
        _ = v10
    }
    v2 = (v2 | v2)
    if true {
        if (685 >= 127) {
            var v11 uint8 = 42
            for fi12 := 0; fi12 < 6; fi12++ {
                _ = len(v5)
                if len(v5) >= 2 {
                    v5 = v5[1:len(v5)]
                }
                _ = fi12
            }
            _ = v11
        }
    } else {
        v3 -= v3
    }
    var v13 int16 = 1000
    if false {
        var v14 uint16 = 829
        for fi15 := 0; fi15 < 14; fi15++ {
            _ = len(v5)
            _ = fi15
        }
        var v16 int64 = -1
        _ = v14
        _ = v16
    } else {
        if true {
            for fi17 := 0; fi17 < 7; fi17++ {
                v8 -= v8
                for fi18 := 0; fi18 < 5; fi18++ {
                    v19 := GenS1{F0: (v6 / ((95) | 1)), F1: fi18}
                    if true {
                        fmt.Println(fi18, v13)
                        v20 := GenS2{F0: 1, F1: 255, F2: ("" < "a")}
                        if true {
                            v5 = append(v5, "")
                        } else {
                            var v21 int16 = (v13 >> ((65535) & 63))
                            v22 := GenS2{F0: 0, F1: 255, F2: true}
                            v2 = 1
                            _ = v21
                            _ = v22
                        }
                        _ = len(v4)
                        _ = v20
                    } else {
                        var v23 int16 = (^(v13))
                        _ = v23
                    }
                    _ = fi18
                    _ = v19
                }
                _ = fi17
            }
            v2 -= (v2 + 2)
            if true {
                if (!(-1 > -1)) {
                    if true {
                        if ((!false) && (191 > 128)) {
                            v24 := GenS2{F0: -41, F1: 620, F2: false}
                            var v25 int8 = -1
                            _ = v24
                            _ = v25
                        } else {
                            fmt.Println(v2, v8)
                            _ = len(v1)
                            var v26 uint = v6
                            _ = v26
                        }
                        if (42 <= uint16(v6)) {
                            _ = len(v1)
                        } else {
                            switch "foo bar" {
                            case "foo bar":
                                for fi27 := 0; fi27 < 9; fi27++ {
                                    var v28 uint16 = 744
                                    delete(v1, "hello")
                                    _ = fi27
                                    _ = v28
                                }
                            }
                            switch "a" {
                            case "\n\t":
                                if false {
                                    var v29 int64 = -9223372036854775808
                                    if true {
                                        v2 += 2
                                        if true {
                                            if true {
                                                v2 = (v2 >> ((v6) & 63))
                                                v30 := GenS3{F0: v13, F1: 1000, F2: v29}
                                                v4 = append(v4, -572)
                                                _ = v30
                                            }
                                            var v31 uint32 = 65535
                                            v32 := GenS2{F0: -128, F1: 32767, F2: (!false)}
                                            _ = v31
                                            _ = v32
                                        }
                                    } else {
                                        for fi33 := 0; fi33 < 7; fi33++ {
                                            if false {
                                                v29 += (v29 ^ (v29 << ((42) & 63)))
                                                v34 := GenS4{F0: (v2 << ((v6) & 63)), F1: v8}
                                                for fi35 := 0; fi35 < 6; fi35++ {
                                                    v36 := GenS1{F0: v6, F1: fi35}
                                                    _ = fi35
                                                    _ = v36
                                                }
                                                _ = v34
                                            } else {
                                                if len(v5) > 0 {
                                                    v5[0] = ("hello" + "\n\t")
                                                }
                                                if false {
                                                    var v37 int32 = -1
                                                    v6 = (v6 | v6)
                                                    fmt.Println(fi33, v29, v29)
                                                    _ = v37
                                                }
                                                fmt.Println(v13)
                                            }
                                            v38 := v1[""]
                                            _ = fi33
                                            _ = v38
                                        }
                                        if true {
                                            switch 474 {
                                            case 4294967295:
                                                for fi39 := 0; fi39 < 2; fi39++ {
                                                    if true {
                                                        for fi40 := 0; fi40 < 13; fi40++ {
                                                            switch 1 {
                                                            case 63:
                                                                switch (fi40 | -9223372036854775808) {
                                                                case 9223372036854775807:
                                                                    if true {
                                                                        delete(v1, "\n\t")
                                                                        var v41 int64 = 42
                                                                        delete(v1, ((("" + "a") + ("a" + ("\n\t" + "foo bar"))) + "a"))
                                                                        v29 += int64(v6)
                                                                        v41 -= (v29 & -420)
                                                                        _ = v41
                                                                    }
                                                                default:
                                                                    fmt.Println(v29, fi40, v6)
                                                                }
                                                            case -9223372036854775808:
                                                                v6 += (^(v6))
                                                            }
                                                            v2 -= (v2 & 18446744073709551615)
                                                            v42 := GenS1{F0: (v6 ^ uint(v3)), F1: fi39}
                                                            _ = fi40
                                                            _ = v42
                                                        }
                                                    } else {
                                                        for fi43 := 0; fi43 < 11; fi43++ {
                                                            if true {
                                                                switch ("hello" + "\n\t") {
                                                                case "a":
                                                                    if true {
                                                                        v3 += (v3 << ((v6) & 63))
                                                                    }
                                                                case "hello":
                                                                    _ = len(v4)
                                                                default:
                                                                    v6 = v6
                                                                    _ = len(v1)
                                                                }
                                                                var v44 uint32 = v3
                                                                v13 -= (v13 + v13)
                                                                _ = v44
                                                            } else {
                                                                var v45 float32 = float32(((0.0 - -2.25) * 1e10))
                                                                if true {
                                                                    fmt.Println(v8)
                                                                    var v46 uint16 = 255
                                                                    _ = v46
                                                                } else {
                                                                    if (65535 == 223) {
                                                                        v2 = v2
                                                                        v13 -= -423
                                                                    }
                                                                    var v47 uint8 = 127
                                                                    _ = v47
                                                                }
                                                                _ = v45
                                                            }
                                                            v48 := GenS5{F0: v29, F1: "a"}
                                                            _ = fi43
                                                            _ = v48
                                                        }
                                                        var v49 int8 = 85
                                                        v6 += 4294967295
                                                        _ = v49
                                                    }
                                                    v50 := GenS6{F0: 0, F1: v13}
                                                    fmt.Println(v29, v3)
                                                    _ = fi39
                                                    _ = v50
                                                }
                                            case 127:
                                                for fi51 := 0; fi51 < 9; fi51++ {
                                                    v6 += (v6 % ((v6) | 1))
                                                    switch (fi51 | fi51) {
                                                    case -2147483648:
                                                        v8 += (v2 % (((v2 + 0)) | 1))
                                                    }
                                                    if (42 < 2147483647) {
                                                        if true {
                                                            var v52 float64 = (float64((0.0 / 1.5)) * -2.25)
                                                            if len(v5) >= 2 {
                                                                v5 = v5[1:len(v5)]
                                                            }
                                                            _ = v52
                                                        }
                                                        v6 = (v6 % ((uint(v2)) | 1))
                                                    } else {
                                                        for fi53 := 0; fi53 < 2; fi53++ {
                                                            if (216 > 1) {
                                                                if true {
                                                                    var v54 int = (fi51 + (-(fi53)))
                                                                    switch "\n\t" {
                                                                    case "":
                                                                        switch fi53 {
                                                                        case -9223372036854775808:
                                                                            var v55 uint64 = (^(v8))
                                                                            _ = v55
                                                                        case 0:
                                                                            fmt.Println(v54, v13)
                                                                        }
                                                                    case "\n\t":
                                                                        fmt.Println(v3, v13)
                                                                    case "hello":
                                                                        v29 -= 42
                                                                    }
                                                                    v56 := GenS2{F0: 1, F1: 419, F2: false}
                                                                    v5 = append(v5, "foo bar")
                                                                    _ = v54
                                                                    _ = v56
                                                                }
                                                                if false {
                                                                    if false {
                                                                        v8 = (^(v8))
                                                                    } else {
                                                                        var v57 int8 = 4
                                                                        switch ("a" + ("a" + "foo bar")) {
                                                                        case "foo bar":
                                                                            if true {
                                                                                var v58 uint = 4294967295
                                                                                if true {
                                                                                    v2 += (v8 * (^(v2)))
                                                                                    var v59 float64 = -2.25
                                                                                    _ = v59
                                                                                }
                                                                                v60 := GenS5{F0: v29, F1: (("\n\t" + (("a" + "foo bar") + (("a" + (("a" + "\n\t") + "\n\t")) + "foo bar"))) + "")}
                                                                                _ = v58
                                                                                _ = v60
                                                                            }
                                                                        }
                                                                        _ = v57
                                                                    }
                                                                    v61 := v1[("foo bar" + "foo bar")]
                                                                    _ = v61
                                                                } else {
                                                                    var v62 int8 = 42
                                                                    fmt.Println(fi51, v8)
                                                                    var v63 float32 = (1e10 + 3.14159)
                                                                    _ = v62
                                                                    _ = v63
                                                                }
                                                            }
                                                            if true {
                                                                v64 := GenS6{F0: uint16(v8), F1: v13}
                                                                fmt.Println(v8)
                                                                if true {
                                                                    var v65 uint = v6
                                                                    if true {
                                                                        if (((fi53 | (fi53 - (fi51 / ((int(v8)) | 1)))) <= fi53) && ((v6 << ((v65) & 63)) >= (v6 - (v6 * (v6 * v65))))) {
                                                                            var v66 uint64 = (v2 << ((25) & 63))
                                                                            _ = v66
                                                                        }
                                                                        fmt.Println(fi51)
                                                                        if false {
                                                                            if len(v5) > 0 {
                                                                                v5[0] = "hello"
                                                                            }
                                                                            if (false && ((v13 + (v13 ^ v13)) != (-(v13)))) {
                                                                                v6 -= (^(v65))
                                                                                var v67 float64 = (((-2.25 + 0.0) / 1e10) - 3.14159)
                                                                                if ((v2 ^ (v8 & v8)) < 42) {
                                                                                    if true {
                                                                                        v67 -= (v67 / v67)
                                                                                    }
                                                                                    v67 += ((v67 / v67) / v67)
                                                                                } else {
                                                                                    v68 := GenS7{F0: false, F1: 45}
                                                                                    _ = v68
                                                                                }
                                                                                _ = v67
                                                                            } else {
                                                                                _ = len(v5)
                                                                                v1["\n\t"] = (fi51 & (fi51 << ((v65) & 63)))
                                                                                fmt.Println(v2, v65)
                                                                            }
                                                                        }
                                                                    }
                                                                    delete(v1, "\n\t")
                                                                    _ = v65
                                                                } else {
                                                                    v13 -= v13
                                                                    v69 := GenS4{F0: (v2 / (((v2 | v8)) | 1)), F1: 455}
                                                                    if len(v5) > 0 {
                                                                        v5[0] = "hello"
                                                                    }
                                                                    _ = v69
                                                                }
                                                                _ = v64
                                                            }
                                                            _ = fi53
                                                        }
                                                    }
                                                    _ = fi51
                                                }
                                            }
                                            var v70 int32 = -2147483648
                                            _ = v70
                                        }
                                    }
                                    _ = v29
                                } else {
                                    for fi71 := 0; fi71 < 6; fi71++ {
                                        fmt.Println(fi71, v13)
                                        var v72 float32 = float32((3.14159 / ((0.0 * 0.0) - (-2.25 * (-2.25 - 3.14159)))))
                                        for fi73 := 0; fi73 < 8; fi73++ {
                                            switch 128 {
                                            case 0:
                                                var v74 int64 = -1
                                                _ = v74
                                            case -924:
                                                _ = len(v4)
                                            case 4294967295:
                                                v13 = 1
                                            }
                                            v75 := GenS8{F0: 127}
                                            _ = fi73
                                            _ = v75
                                        }
                                        _ = fi71
                                        _ = v72
                                    }
                                    fmt.Println(v2)
                                    v76 := v1["foo bar"]
                                    _ = v76
                                }
                            default:
                                for fi77 := 0; fi77 < 12; fi77++ {
                                    if false {
                                        for fi78 := 0; fi78 < 11; fi78++ {
                                            fmt.Println(v13, v6)
                                            if (1e10 > 0.1) {
                                                v6 += (v6 * 18446744073709551615)
                                                v1["foo bar"] = fi78
                                                v4 = append(v4, 42)
                                            }
                                            v6 = (^(v6))
                                            _ = fi78
                                        }
                                    } else {
                                        v1[(("" + (("" + (("\n\t" + "") + "")) + "a")) + ("foo bar" + "foo bar"))] = fi77
                                        switch "" {
                                        case "":
                                            v6 += v6
                                        case "\n\t":
                                            delete(v1, "foo bar")
                                        default:
                                            v3 = (v3 % ((v3) | 1))
                                        }
                                        v3 -= 698
                                    }
                                    v3 -= v3
                                    fmt.Println(v2)
                                    _ = fi77
                                }
                            }
                        }
                    }
                } else {
                    v1[(("" + "") + "\n\t")] = 128
                    if true {
                        switch "\n\t" {
                        case "hello":
                            for fi79 := 0; fi79 < 3; fi79++ {
                                v80 := GenS9{F0: 141, F1: 9223372036854775807}
                                var v81 int32 = -2147483648
                                switch "\n\t" {
                                case "\n\t":
                                    v13 -= int16(v2)
                                }
                                _ = fi79
                                _ = v80
                                _ = v81
                            }
                        case "\n\t":
                            for fi82 := 0; fi82 < 1; fi82++ {
                                _ = len(v1)
                                var v83 uint16 = 65535
                                var v84 float32 = 3.14159
                                _ = fi82
                                _ = v83
                                _ = v84
                            }
                        default:
                            var v85 uint16 = 65535
                            _ = v85
                        }
                    } else {
                        var v86 uint = v6
                        v87 := GenS9{F0: 59, F1: -1}
                        v13 = (v13 * 56)
                        _ = v86
                        _ = v87
                    }
                    v8 = 18446744073709551615
                }
            } else {
                switch 4294967295 {
                case 993:
                    v6 = v6
                case -1:
                    var v88 uint8 = 227
                    _ = v88
                case 1:
                    v13 += (v13 ^ v13)
                }
                v89 := GenS10{F0: 1, F1: (v13 + (v13 % ((v13) | 1))), F2: "foo bar"}
                var v90 string = ("foo bar" + ("hello" + ("hello" + (("foo bar" + "hello") + ("" + ("" + "\n\t"))))))
                _ = v89
                _ = v90
            }
        }
        delete(v1, "\n\t")
    }
    v13 -= (v13 + (v13 - (v13 & v13)))
    switch "" {
    case "a":
        v13 += v13
    }
    var v91 float32 = 1e10
    if false {
        switch 32767 {
        case -9223372036854775808:
            fmt.Println(v13, v2, v3)
        }
        var v92 uint8 = 0
        _ = v92
    }
    var v93 uint = v6
    var v94 int16 = (v13 ^ (v13 * v13))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v91:", v91)
    fmt.Println("v93:", v93)
    fmt.Println("v94:", v94)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v91
    _ = v93
    _ = v94
}
