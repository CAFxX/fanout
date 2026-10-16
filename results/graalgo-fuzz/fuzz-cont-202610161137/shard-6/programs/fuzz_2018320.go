package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint }
type GenS2 struct { F0 int16; F1 uint16; F2 int16 }
type GenS3 struct { F0 int32; F1 uint32; F2 bool }
type GenS4 struct { F0 int8; F1 uint32 }
type GenS5 struct { F0 int32; F1 bool }
type GenS6 struct { F0 string }
func genFn0(p0 uint) bool {
    p0 += 65535
    v7 := GenS1{F0: 2147483647, F1: p0}
    switch ("foo bar" + "foo bar") {
    case "hello":
        fmt.Println(p0)
    case "":
        fmt.Println(p0)
    case "foo bar":
        var v8 int64 = 0
        _ = v8
    default:
        switch "\n\t" {
        case "\n\t":
            for fi9 := 0; fi9 < 2; fi9++ {
                for fi10 := 0; fi10 < 6; fi10++ {
                    var v11 uint16 = 1
                    if ((v11 >> ((uint(v11)) & 63)) < v11) {
                        if (32767 >= -855) {
                            if ((-999 <= 257) && false) {
                                fmt.Println(fi9, v11, v11)
                                v12 := []int{(fi10 / (((fi9 % ((fi10) | 1))) | 1))}
                                _ = v12
                            } else {
                                var v13 bool = (!(0.0 >= 3.14159))
                                _ = v13
                            }
                        } else {
                            v11 = 504
                            if (!true) {
                                v14 := GenS1{F0: -974, F1: 161}
                                _ = v14
                            }
                        }
                    }
                    var v15 uint16 = 990
                    _ = fi10
                    _ = v11
                    _ = v15
                }
                var v16 int16 = -32768
                v16 = (v16 + v16)
                _ = fi9
                _ = v16
            }
        case "":
            for fi17 := 0; fi17 < 16; fi17++ {
                if false {
                    v18 := GenS2{F0: -1, F1: 255, F2: 0}
                    fmt.Println(fi17)
                    _ = v18
                }
                p0 -= (p0 / ((481) | 1))
                _ = fi17
            }
        case "foo bar":
            if true {
                v19 := GenS3{F0: 2147483647, F1: 596, F2: (true || true)}
                _ = v19
            } else {
                p0 -= p0
            }
        }
    }
    _ = p0
    _ = v7
    return false
}

func main() {
    v1 := make(map[string]uint32)
    var v2 float64 = (((float64(-2.25) / -2.25) - 1e10) + 0.1)
    v3 := []int{514, -615, 1, 2147483647}
    var v4 float64 = float64((((0.0 * ((0.1 - float32((0.1 + (3.14159 * (((v2 / 1.0) - v2) + v2))))) + 0.1)) * 0.1) + 0.0))
    var v5 float64 = (v2 - v4)
    var v6 uint32 = 638
    if true {
        var v20 uint32 = 175
        switch -9223372036854775808 {
        case -888:
            v20 = (v6 + v20)
        case 42:
            var v21 uint8 = 42
            _ = v21
        case 474:
            v6 += (^(v20))
        default:
            if (!true) {
                var v22 string = ""
                if (-1 >= 127) {
                    for fi23 := 0; fi23 < 20; fi23++ {
                        v2 = ((-2.25 / v2) / v4)
                        v24 := GenS2{F0: -586, F1: 2, F2: -937}
                        _ = fi23
                        _ = v24
                    }
                }
                fmt.Println(v2)
                _ = v22
            } else {
                v5 = (v4 + (v2 + (0.1 * v4)))
                var v25 uint64 = 961
                if (((183 == 127) || false) || true) {
                    v25 += (v25 | (v25 << ((17) & 63)))
                } else {
                    switch "hello" {
                    case "\n\t":
                        if len(v3) > 0 {
                            v3[0] = 730
                        }
                    default:
                        fmt.Println(v25)
                    }
                    var v26 int64 = 554
                    _ = v26
                }
                _ = v25
            }
        }
        _ = v20
    } else {
        var v27 uint32 = (v6 & (v6 | v6))
        fmt.Println(v4, v5, v4)
        _ = v27
    }
    v28 := v1["a"]
    var v29 float32 = 1e10
    fmt.Println(v4)
    delete(v1, "")
    v30 := GenS1{F0: 0, F1: 255}
    v31 := GenS1{F0: -2147483648, F1: 65535}
    if (!(127 >= 0)) {
        for fi32 := 0; fi32 < 15; fi32++ {
            v3 = append(v3, -991)
            var v33 uint8 = 13
            if true {
                if (!true) {
                    if false {
                        switch "\n\t" {
                        case "":
                            v6 -= 119
                        case "foo bar":
                            var v34 bool = (!(0.0 <= v29))
                            _ = v34
                        case "hello":
                            fmt.Println(v33)
                        }
                        for fi35 := 0; fi35 < 12; fi35++ {
                            var v36 uint32 = (v28 | uint32(v33))
                            _ = fi35
                            _ = v36
                        }
                    } else {
                        if false {
                            if len(v3) > 0 {
                                v3[0] = fi32
                            }
                            _ = len(v3)
                        } else {
                            var v37 int64 = -1
                            _ = v37
                        }
                    }
                } else {
                    var v38 int16 = 258
                    _ = v38
                }
                switch "\n\t" {
                case "":
                    var v39 uint8 = 0
                    _ = v39
                case "foo bar":
                    var v40 bool = false
                    _ = v40
                case "hello":
                    v41 := GenS1{F0: -1, F1: 255}
                    _ = v41
                default:
                    if (42 != 0) {
                        for fi42 := 0; fi42 < 9; fi42++ {
                            var v43 uint = 134
                            var v44 uint = v43
                            _ = fi42
                            _ = v43
                            _ = v44
                        }
                        if (4294967295 != 522) {
                            switch fi32 {
                            case 4294967295:
                                switch "foo bar" {
                                case "":
                                    switch "foo bar" {
                                    case "a":
                                        v28 = v6
                                    default:
                                        switch "" {
                                        case "foo bar":
                                            v28 = v6
                                        case "hello":
                                            v2 -= (v5 * v2)
                                        }
                                    }
                                case "hello":
                                    if ((!false) && true) {
                                        if true {
                                            var v45 float64 = 1e10
                                            var v46 int8 = 42
                                            _ = v45
                                            _ = v46
                                        }
                                        if (((0.0 / v2) - 3.14159) < v2) {
                                            v47 := GenS1{F0: -1, F1: 123}
                                            var v48 uint32 = (v6 | (v6 ^ v28))
                                            if false {
                                                var v49 uint8 = 198
                                                _ = v49
                                            }
                                            _ = v47
                                            _ = v48
                                        } else {
                                            for fi50 := 0; fi50 < 9; fi50++ {
                                                var v51 float32 = (float32(v5) + 0.1)
                                                var v52 uint8 = (^(v33))
                                                if true {
                                                    v53 := GenS4{F0: 127, F1: 4294967295}
                                                    var v54 int8 = 127
                                                    var v55 uint = 1
                                                    if (false || (true && false)) {
                                                        v56 := GenS1{F0: -471, F1: v55}
                                                        v55 += (^(v55))
                                                        if true {
                                                            v57 := v1[""]
                                                            fmt.Println(v5, v52, v33)
                                                            _ = v57
                                                        }
                                                        switch fi32 {
                                                        case -9223372036854775808:
                                                            switch ("hello" + "foo bar") {
                                                            case "\n\t":
                                                                v52 = (v52 * (v33 << ((v55) & 63)))
                                                            default:
                                                                switch "a" {
                                                                case "a":
                                                                    v54 -= v54
                                                                default:
                                                                    _ = len(v3)
                                                                }
                                                            }
                                                        }
                                                        var v58 string = ("" + "")
                                                        switch int(v28) {
                                                        case 2147483647:
                                                            v29 -= v29
                                                        case 140:
                                                            v3 = append(v3, fi32)
                                                        case 127:
                                                            v52 = v52
                                                        }
                                                        v28 = (v6 & v6)
                                                        var v59 int64 = 470
                                                        _ = v56
                                                        _ = v58
                                                        _ = v59
                                                    }
                                                    _ = v53
                                                    _ = v54
                                                    _ = v55
                                                }
                                                if (!false) {
                                                    v52 = v33
                                                }
                                                var v60 int8 = -128
                                                _ = fi50
                                                _ = v51
                                                _ = v52
                                                _ = v60
                                            }
                                        }
                                    }
                                }
                            case 525:
                                if true {
                                    switch -2147483648 {
                                    case 42:
                                        v61 := GenS4{F0: 89, F1: 607}
                                        _ = v61
                                    case -9223372036854775808:
                                        for fi62 := 0; fi62 < 3; fi62++ {
                                            v4 += 3.14159
                                            var v63 uint = 79
                                            v28 -= 438
                                            v4 += (v4 * v4)
                                            v64 := GenS4{F0: 75, F1: v28}
                                            _ = len(v1)
                                            delete(v1, ((("a" + "") + ("a" + "\n\t")) + ("" + "hello")))
                                            v65 := GenS5{F0: 150, F1: (2147483647 > 1)}
                                            var v66 bool = false
                                            _ = fi62
                                            _ = v63
                                            _ = v64
                                            _ = v65
                                            _ = v66
                                        }
                                    }
                                    var v67 int32 = -1
                                    _ = v67
                                } else {
                                    v3 = append(v3, 1)
                                }
                            }
                            for fi68 := 0; fi68 < 17; fi68++ {
                                v33 += (v33 - uint8(v6))
                                var v69 int64 = 9223372036854775807
                                _ = fi68
                                _ = v69
                            }
                        } else {
                            v2 += 0.0
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            for fi70 := 0; fi70 < 14; fi70++ {
                                var v71 bool = (693 >= 42)
                                var v72 int8 = 76
                                _ = fi70
                                _ = v71
                                _ = v72
                            }
                        }
                    }
                }
                _ = len(v3)
            } else {
                var v73 int32 = 71
                if true {
                    delete(v1, "\n\t")
                    switch ((("foo bar" + ("foo bar" + "")) + (("hello" + "foo bar") + "foo bar")) + "\n\t") {
                    case "\n\t":
                        v74 := GenS6{F0: ""}
                        _ = v74
                    case "":
                        fmt.Println(v33)
                    default:
                        if (((!(605 <= 65535)) && false) && false) {
                            var v75 bool = (false || ((v6 % (((^(v6))) | 1)) > v28))
                            switch (fi32 ^ 0) {
                            case 32768:
                                if false {
                                    v73 = v73
                                } else {
                                    for fi76 := 0; fi76 < 18; fi76++ {
                                        fmt.Println(v28)
                                        var v77 uint32 = 658
                                        _ = fi76
                                        _ = v77
                                    }
                                }
                            case 128:
                                var v78 int64 = -931
                                _ = v78
                            case 2147483647:
                                _ = len(v3)
                            }
                            _ = v75
                        }
                    }
                } else {
                    if true {
                        v1[("foo bar" + "\n\t")] = v6
                    } else {
                        _ = len(v3)
                        v79 := GenS3{F0: 42, F1: v28, F2: (32767 < 1000)}
                        _ = v79
                    }
                    for fi80 := 0; fi80 < 13; fi80++ {
                        var v81 uint = 18446744073709551615
                        v3 = append(v3, (fi32 % (((fi32 / ((fi32) | 1))) | 1)))
                        if ("" != "") {
                            if true {
                                v81 += 2
                                var v82 int32 = v73
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                _ = v82
                            }
                        }
                        v83 := GenS1{F0: (v73 - -2147483648), F1: v81}
                        v81 += v81
                        switch (fi32 & fi32) {
                        case 1:
                            var v84 float32 = v29
                            var v85 float64 = v5
                            _ = v84
                            _ = v85
                        default:
                            v86 := GenS1{F0: 872, F1: (v81 + 4294967295)}
                            _ = v86
                        }
                        _ = fi80
                        _ = v81
                        _ = v83
                    }
                }
                v33 = 199
                _ = v73
            }
            _ = fi32
            _ = v33
        }
    }
    v87 := v1["a"]
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v87:", v87)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v31
    _ = v4
    _ = v5
    _ = v6
    _ = v87
}
