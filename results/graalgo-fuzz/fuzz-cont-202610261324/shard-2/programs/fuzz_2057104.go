package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32; F2 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int }
type GenS4 struct { F0 int8; F1 int32; F2 int32 }
type GenS5 struct { F0 bool; F1 string }
type GenS6 struct { F0 string; F1 int64 }
type GenS7 struct { F0 int16; F1 bool; F2 int8 }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint16; F1 uint16; F2 uint8 }
type GenS10 struct { F0 uint8; F1 string }
type GenS11 struct { F0 int32; F1 uint16; F2 uint64 }
type GenS12 struct { F0 uint64; F1 uint16 }
type GenS13 struct { F0 uint64; F1 int64; F2 string }

func main() {
    var v1 float64 = (-2.25 - 0.0)
    var v2 float64 = (v1 / v1)
    v3 := []string{((("" + ("hello" + "hello")) + ("a" + "")) + (("a" + (("hello" + "a") + ("foo bar" + ("a" + "foo bar")))) + ("foo bar" + ("hello" + "a"))))}
    var v4 string = "a"
    v5 := make(map[string]string)
    var v6 float64 = (v2 + float64((0.1 / (float32(v1) - 0.1))))
    for fi7 := 0; fi7 < 10; fi7++ {
        fmt.Println(fi7, v6, v4)
        fmt.Println(v1)
        v1 += 0.0
        _ = fi7
    }
    v5["foo bar"] = "hello"
    v1 -= (v6 / (0.0 * (v1 - float64(float32(-2.25)))))
    switch (("\n\t" + "") + (v4 + (v4 + v4))) {
    case "\n\t":
        v6 += (v6 / (v2 / 1e10))
    case "a":
        v5[v4] = v4
    case "hello":
        if (true && false) {
            var v8 int64 = -1
            if (false && (!true)) {
                v9 := GenS1{F0: 42, F1: 42, F2: 355}
                _ = v9
            } else {
                v4 = "foo bar"
                var v10 float32 = (0.1 - (-2.25 / (-2.25 - 0.1)))
                for fi11 := 0; fi11 < 6; fi11++ {
                    if (42 != 18446744073709551615) {
                        if (0 < 1) {
                            v12 := GenS2{F0: false}
                            for fi13 := 0; fi13 < 2; fi13++ {
                                switch (v4 + ((v4 + (v4 + "hello")) + (v4 + (v4 + (v4 + (v4 + v4)))))) {
                                case "foo bar":
                                    if (0 < 1) {
                                        v14 := GenS1{F0: 534, F1: -29, F2: -9223372036854775808}
                                        var v15 uint8 = 2
                                        var v16 int16 = 1
                                        _ = v14
                                        _ = v15
                                        _ = v16
                                    } else {
                                        fmt.Println(fi13, v4)
                                        v17 := GenS3{F0: (fi11 / (((fi13 << ((42) & 63))) | 1))}
                                        _ = v17
                                    }
                                case "":
                                    v18 := GenS3{F0: (fi13 << ((18446744073709551615) & 63))}
                                    _ = v18
                                case "a":
                                    v8 = 42
                                    v6 += v2
                                    v19 := GenS1{F0: fi11, F1: -1, F2: (fi13 >> ((500) & 63))}
                                    v20 := GenS4{F0: 127, F1: 1, F2: 42}
                                    v21 := v5[v4]
                                    _ = v19
                                    _ = v20
                                    _ = v21
                                default:
                                    var v22 uint16 = 255
                                    _ = v22
                                }
                                v23 := GenS5{F0: (false && ((0 == 1) && (2147483647 > 0))), F1: v4}
                                var v24 uint32 = 2
                                var v25 int8 = 61
                                var v26 string = v4
                                _ = fi13
                                _ = v23
                                _ = v24
                                _ = v25
                                _ = v26
                            }
                            _ = v12
                        }
                        for fi27 := 0; fi27 < 2; fi27++ {
                            _ = len(v5)
                            var v28 int16 = 42
                            v29 := GenS5{F0: (!(0 <= 2)), F1: v4}
                            _ = len(v3)
                            _ = fi27
                            _ = v28
                            _ = v29
                        }
                    }
                    _ = len(v5)
                    _ = fi11
                }
                _ = v10
            }
            v8 = 1
            _ = v8
        }
    }
    if true {
        fmt.Println(v6, v4)
    } else {
        var v30 uint32 = 2
        _ = v30
    }
    for fi31 := 0; fi31 < 5; fi31++ {
        var v32 string = (v4 + ((v4 + "foo bar") + "hello"))
        _ = fi31
        _ = v32
    }
    fmt.Println(v1, v4, v6)
    for fi33 := 0; fi33 < 11; fi33++ {
        for fi34 := 0; fi34 < 15; fi34++ {
            if (!(3.14159 == float32(v1))) {
                v6 = (v1 + v1)
                fmt.Println(fi34, v2)
                v35 := GenS6{F0: v4, F1: 0}
                _ = v35
            }
            v2 += (((v1 - float64(0.0)) - ((((1e10 / (-2.25 - v6)) / (0.0 + float64((3.14159 / 1.0)))) / v1) / v1)) * (v2 + ((v1 - (1e10 * 0.1)) / v2)))
            if (1 < 128) {
                if ((1 != 0) && false) {
                    v6 += 1.5
                    var v36 int8 = 42
                    var v37 uint16 = 176
                    _ = v36
                    _ = v37
                } else {
                    var v38 uint16 = 65535
                    v39 := GenS5{F0: false, F1: ("a" + (("a" + v4) + "hello"))}
                    if (!true) {
                        var v40 uint16 = (v38 / ((v38) | 1))
                        v1 = (v1 + 1.5)
                        switch v4 {
                        case "\n\t":
                            var v41 string = v4
                            _ = v41
                        case "foo bar":
                            var v42 bool = (42 <= 1)
                            _ = v42
                        default:
                            if ((1e10 / ((3.14159 * (-2.25 / 1.5)) - float32(1e10))) > -2.25) {
                                fmt.Println(fi34)
                                v43 := GenS5{F0: ((2 <= 610) && (1 < 4294967295)), F1: ((v4 + ("" + v4)) + (v4 + "a"))}
                                _ = v43
                            }
                        }
                        _ = v40
                    } else {
                        delete(v5, v4)
                        v1 -= 3.14159
                        v44 := GenS7{F0: 32767, F1: true, F2: 0}
                        _ = v44
                    }
                    _ = v38
                    _ = v39
                }
            } else {
                if false {
                    var v45 int16 = 0
                    v46 := v5[("hello" + "hello")]
                    if true {
                        v47 := GenS6{F0: (v4 + v46), F1: -66}
                        var v48 int64 = 674
                        if len(v3) > 0 {
                            v3[0] = ((v46 + "") + ("a" + ((v46 + "foo bar") + v4)))
                        }
                        _ = v47
                        _ = v48
                    } else {
                        var v49 int64 = -1
                        v50 := GenS7{F0: v45, F1: ((false && (1 >= 2147483647)) || (true || true)), F2: -128}
                        if len(v3) > 0 {
                            v3[0] = ("" + v4)
                        }
                        _ = v49
                        _ = v50
                    }
                    _ = v45
                    _ = v46
                }
                v3 = append(v3, ((v4 + v4) + (v4 + "\n\t")))
                v51 := GenS8{F0: 0}
                _ = v51
            }
            _ = fi34
        }
        var v52 int32 = -2147483648
        var v53 int16 = -32768
        _ = fi33
        _ = v52
        _ = v53
    }
    if true {
        var v54 float32 = 3.14159
        if false {
            v55 := GenS9{F0: 988, F1: 32767, F2: 42}
            v2 = (v2 * 0.0)
            _ = v55
        }
        _ = v54
    } else {
        v56 := GenS10{F0: 0, F1: (((v4 + ((v4 + ((v4 + v4) + v4)) + v4)) + "a") + "\n\t")}
        v2 += (0.1 - -2.25)
        _ = v56
    }
    var v57 string = "foo bar"
    v6 -= v2
    fmt.Println(v6, v2, v57)
    v57 = v4
    var v58 uint8 = 216
    switch v57 {
    case "a":
        delete(v5, (v4 + ((v4 + v57) + "\n\t")))
    case "\n\t":
        for fi59 := 0; fi59 < 13; fi59++ {
            v60 := GenS9{F0: 807, F1: 2, F2: v58}
            for fi61 := 0; fi61 < 16; fi61++ {
                v2 += v6
                switch -593 {
                case -54:
                    if (-9223372036854775808 >= 410) {
                        if ("" > v57) {
                            v5[("\n\t" + ("a" + "\n\t"))] = v4
                            v4 = (v4 + "hello")
                            v62 := GenS9{F0: 2, F1: 1, F2: v58}
                            _ = v62
                        }
                        if false {
                            var v63 bool = (true || (!true))
                            v4 = v4
                            fmt.Println(v1)
                            _ = v63
                        } else {
                            if (v58 != (v58 >> ((806) & 63))) {
                                var v64 int8 = -19
                                switch (fi61 / ((-9223372036854775808) | 1)) {
                                case -2147483648:
                                    var v65 uint8 = (v58 | v58)
                                    _ = v65
                                }
                                v66 := GenS6{F0: v57, F1: -221}
                                v67 := GenS10{F0: v58, F1: (("" + "\n\t") + (v4 + v4))}
                                v68 := GenS3{F0: 168}
                                var v69 uint = 407
                                v6 -= v2
                                v3 = append(v3, (v4 + v4))
                                _ = v64
                                _ = v66
                                _ = v67
                                _ = v68
                                _ = v69
                            } else {
                                v70 := GenS7{F0: 1000, F1: false, F2: 21}
                                _ = v70
                            }
                            fmt.Println(v4, v6)
                        }
                        var v71 uint16 = 255
                        _ = v71
                    } else {
                        if (42 != 32767) {
                            fmt.Println(v6)
                            fmt.Println(v57)
                        }
                        if false {
                            var v72 float64 = 1e10
                            v73 := GenS5{F0: false, F1: "foo bar"}
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = v72
                            _ = v73
                        } else {
                            if (!(!true)) {
                                if (0 != 1) {
                                    switch v4 {
                                    case "\n\t":
                                        var v74 int32 = 42
                                        _ = v74
                                    default:
                                        if true {
                                            var v75 uint64 = 1
                                            v76 := GenS11{F0: 42, F1: 1, F2: (^(v75))}
                                            v77 := GenS12{F0: uint64(v58), F1: 32767}
                                            if len(v3) > 0 {
                                                v3[0] = v57
                                            }
                                            var v78 uint32 = 0
                                            v1 = v2
                                            switch 4294967295 {
                                            case -496:
                                                if (65535 != 2) {
                                                    var v79 float32 = 1.5
                                                    _ = len(v5)
                                                    switch v4 {
                                                    case "hello":
                                                        if true {
                                                            if len(v3) >= 2 {
                                                                v3 = v3[1:len(v3)]
                                                            }
                                                            var v80 string = ""
                                                            v79 = v79
                                                            _ = v80
                                                        } else {
                                                            _ = len(v5)
                                                            var v81 int8 = 127
                                                            _ = v81
                                                        }
                                                    case "a":
                                                        switch (-(fi59)) {
                                                        case -1:
                                                            if (("a" <= "") || ((((false && false) && false) || true) && (!((32767 < -553) || ((v75 * (v75 % ((uint64(v58)) | 1))) <= (v75 << ((321) & 63))))))) {
                                                                if ((v57 >= (v4 + v4)) || false) {
                                                                    v75 += (^(v75))
                                                                    fmt.Println(v79, v6, v2)
                                                                    v82 := GenS3{F0: (fi61 + fi61)}
                                                                    _ = v82
                                                                }
                                                                var v83 uint = 4294967295
                                                                _ = v83
                                                            }
                                                        case 32767:
                                                            v57 = (v57 + v4)
                                                        }
                                                    case "":
                                                        if (v75 == v75) {
                                                            if false {
                                                                v3 = append(v3, (v4 + ((v4 + v57) + ((v57 + v57) + "\n\t"))))
                                                                fmt.Println(v4, v57, v57)
                                                                v4 = "foo bar"
                                                                var v84 float64 = (v6 + (((v6 - v1) - ((v2 * float64(float32(v2))) / ((v2 - 3.14159) / 0.1))) + 1.5))
                                                                _ = v84
                                                            } else {
                                                                switch v4 {
                                                                case "\n\t":
                                                                    if len(v3) >= 2 {
                                                                        v3 = v3[1:len(v3)]
                                                                    }
                                                                }
                                                            }
                                                            v85 := GenS11{F0: -2147483648, F1: uint16(v75), F2: (v75 % ((v75) | 1))}
                                                            if (!false) {
                                                                if true {
                                                                    var v86 uint16 = 1
                                                                    fmt.Println(v78, v86)
                                                                    _ = v86
                                                                } else {
                                                                    var v87 int = fi61
                                                                    switch 42 {
                                                                    case 9223372036854775807:
                                                                        v3 = append(v3, "")
                                                                    }
                                                                    v88 := GenS12{F0: (v75 * v75), F1: 2}
                                                                    _ = v87
                                                                    _ = v88
                                                                }
                                                            }
                                                            _ = v85
                                                        }
                                                    default:
                                                        if false {
                                                            _ = len(v5)
                                                        }
                                                    }
                                                    var v89 uint64 = (v75 / ((42) | 1))
                                                    v6 -= float64(1e10)
                                                    _ = v79
                                                    _ = v89
                                                } else {
                                                    v75 = v75
                                                }
                                            }
                                            _ = v75
                                            _ = v76
                                            _ = v77
                                            _ = v78
                                        } else {
                                            v58 += v58
                                            v6 += v2
                                            v90 := GenS4{F0: -1, F1: 0, F2: 562}
                                            _ = v90
                                        }
                                    }
                                }
                                var v91 uint64 = 0
                                _ = v91
                            }
                        }
                        v6 = v2
                    }
                }
                fmt.Println(v1, v2)
                _ = fi61
            }
            _ = fi59
            _ = v60
        }
    }
    if (true || true) {
        for fi92 := 0; fi92 < 17; fi92++ {
            var v93 bool = (fi92 == fi92)
            v94 := GenS13{F0: 42, F1: 568, F2: ("" + v57)}
            var v95 uint = 4294967295
            _ = fi92
            _ = v93
            _ = v94
            _ = v95
        }
        var v96 int64 = 42
        _ = len(v3)
        _ = v96
    } else {
        fmt.Println(v58)
        v4 = (v4 + v57)
        v1 -= (v1 / 1.0)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v57:", v57)
    fmt.Println("v58:", v58)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v57
    _ = v58
    _ = v6
}
