package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint32; F1 string; F2 bool }
type GenS4 struct { F0 int8 }
type GenS5 struct { F0 int8; F1 int16 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int }
type GenS8 struct { F0 int64; F1 int; F2 bool }
type GenS9 struct { F0 uint32 }
type GenS10 struct { F0 uint32 }
type GenS11 struct { F0 int16; F1 uint8; F2 uint }

func main() {
    var v1 bool = (-9223372036854775808 >= 9223372036854775807)
    v2 := make(map[string]bool)
    var v3 float32 = float32(0.0)
    v4 := make(map[string]string)
    v4["hello"] = ("hello" + "")
    var v5 uint32 = 659
    for fi6 := 0; fi6 < 19; fi6++ {
        var v7 int8 = 1
        v8 := []int{fi6, fi6}
        _ = fi6
        _ = v7
        _ = v8
    }
    v9 := GenS1{F0: (!(!false)), F1: 2147483647}
    _ = len(v2)
    if (!(!(127 != 23))) {
        v3 -= ((v3 + 0.1) + 3.14159)
        var v10 float64 = (1.5 - (((0.1 + 1e10) * float64(v3)) / 1.5))
        switch "foo bar" {
        case "":
            var v11 int64 = -261
            _ = v11
        case "foo bar":
            v12 := GenS1{F0: v1, F1: 100}
            _ = v12
        }
        _ = v10
    } else {
        v13 := v2[""]
        if true {
            v3 = v3
            v3 -= (float32(1e10) / -2.25)
        }
        _ = v13
    }
    for fi14 := 0; fi14 < 20; fi14++ {
        v5 = (v5 ^ v5)
        var v15 int32 = 42
        _ = fi14
        _ = v15
    }
    if (-557 < -32768) {
        v3 -= float32(((1.5 / 0.1) - (0.1 * 1e10)))
        var v16 int = 9223372036854775807
        v17 := GenS2{F0: (v16 ^ (v16 ^ 933))}
        _ = v16
        _ = v17
    } else {
        if (("foo bar" + "foo bar") != "\n\t") {
            switch "hello" {
            case "hello":
                v5 -= 4294967295
            case "":
                if ("a" <= (("hello" + "hello") + ("" + "\n\t"))) {
                    v18 := GenS1{F0: v1, F1: v5}
                    switch -1 {
                    case 127:
                        if false {
                            v19 := GenS1{F0: (!v1), F1: 930}
                            v20 := GenS2{F0: 1}
                            _ = v19
                            _ = v20
                        }
                    case 1:
                        v21 := v2[("" + "foo bar")]
                        _ = v21
                    default:
                        v3 -= (float32((3.14159 / 1.5)) + (v3 * v3))
                    }
                    v1 = v1
                    _ = v18
                }
            default:
                var v22 float64 = 3.14159
                _ = v22
            }
        }
    }
    switch "foo bar" {
    case "a":
        v23 := GenS2{F0: 0}
        _ = v23
    default:
        fmt.Println(v5, v3)
    }
    switch "a" {
    case "\n\t":
        for fi24 := 0; fi24 < 3; fi24++ {
            if (42 < 2) {
                switch ("hello" + (("\n\t" + ("\n\t" + ("a" + "foo bar"))) + "hello")) {
                case "hello":
                    delete(v2, "hello")
                }
                v1 = ((fi24 & fi24) == fi24)
            } else {
                v25 := GenS1{F0: true, F1: (v5 % (((^(v5))) | 1))}
                fmt.Println(fi24, v1)
                v2[("\n\t" + "")] = ((fi24 & (fi24 - fi24)) < 1)
                _ = v25
            }
            if v1 {
                switch "foo bar" {
                case "a":
                    v5 -= (^(v5))
                case "\n\t":
                    _ = len(v4)
                case "hello":
                    v1 = v1
                default:
                    for fi26 := 0; fi26 < 10; fi26++ {
                        if v1 {
                            switch ((("a" + (("foo bar" + "a") + ("hello" + "foo bar"))) + ((("hello" + ((("" + "") + "foo bar") + "hello")) + (("\n\t" + (("" + "\n\t") + "foo bar")) + ("hello" + ("foo bar" + "a")))) + "\n\t")) + (("a" + "") + ((("\n\t" + ("\n\t" + "\n\t")) + ("" + "hello")) + "hello"))) {
                            case "":
                                var v27 int = fi26
                                _ = v27
                            default:
                                v28 := GenS3{F0: (v5 / ((v5) | 1)), F1: "a", F2: (!(-364 <= fi26))}
                                _ = v28
                            }
                        } else {
                            var v29 int64 = 1
                            _ = v29
                        }
                        if (true || v1) {
                            v3 -= float32(1.5)
                            if ((25 <= 5) || (127 >= 100)) {
                                v30 := GenS3{F0: v5, F1: "hello", F2: (!false)}
                                switch fi24 {
                                case 4294967295:
                                    if false {
                                        if true {
                                            v3 += 1.5
                                            var v31 bool = (42 >= 65535)
                                            if v1 {
                                                if ((2 > 1) || (0.0 <= (3.14159 * (1e10 + (0.1 - 3.14159))))) {
                                                    v31 = (((-(fi26)) > fi24) || v1)
                                                    var v32 float32 = 1.5
                                                    _ = v32
                                                } else {
                                                    var v33 uint16 = 255
                                                    v4["\n\t"] = ("a" + "\n\t")
                                                    _ = v33
                                                }
                                                v1 = (65535 == 1)
                                                var v34 int64 = 1
                                                _ = v34
                                            } else {
                                                var v35 uint64 = 2
                                                var v36 int16 = -867
                                                _ = v35
                                                _ = v36
                                            }
                                            _ = v31
                                        } else {
                                            if false {
                                                var v37 float32 = (v3 - ((float32((1e10 / 1.5)) * (v3 + ((v3 + (float32((0.0 - 3.14159)) + 1e10)) + ((v3 / v3) * v3)))) + ((0.1 - ((((v3 / (v3 - ((float32((1e10 * 1e10)) / (((0.0 * (3.14159 + v3)) - v3) * float32(1e10))) + 1.5))) * (v3 - v3)) - v3) - (3.14159 * float32(0.1)))) * v3)))
                                                v38 := []int{32768}
                                                v37 += v3
                                                _ = v37
                                                _ = v38
                                            } else {
                                                var v39 int64 = -381
                                                switch ("foo bar" + "foo bar") {
                                                case "a":
                                                    v5 -= (v5 * 65535)
                                                case "foo bar":
                                                    if v1 {
                                                        v2[("a" + "a")] = true
                                                        fmt.Println(fi24, fi26)
                                                        switch "a" {
                                                        case "foo bar":
                                                            switch "a" {
                                                            case "a":
                                                                var v40 string = "foo bar"
                                                                _ = v40
                                                            }
                                                        case "a":
                                                            if (39 > -127) {
                                                                switch ("foo bar" + ("foo bar" + "hello")) {
                                                                case "":
                                                                    v41 := GenS4{F0: -128}
                                                                    if (18446744073709551615 != 18446744073709551615) {
                                                                        switch "foo bar" {
                                                                        case "\n\t":
                                                                            v42 := GenS5{F0: 1, F1: -1}
                                                                            _ = v42
                                                                        }
                                                                    } else {
                                                                        var v43 int16 = -1
                                                                        _ = v43
                                                                    }
                                                                    _ = v41
                                                                }
                                                                var v44 int8 = 42
                                                                switch "" {
                                                                case "foo bar":
                                                                    v3 = (float32(-2.25) - v3)
                                                                case "":
                                                                    fmt.Println(fi26, v44)
                                                                default:
                                                                    _ = len(v2)
                                                                    v3 = v3
                                                                }
                                                                v45 := GenS6{F0: 18}
                                                                _ = v44
                                                                _ = v45
                                                            } else {
                                                                if (42 < 4294967295) {
                                                                    v46 := GenS7{F0: (fi24 / ((fi24) | 1))}
                                                                    var v47 int = (-(fi26))
                                                                    _ = v46
                                                                    _ = v47
                                                                }
                                                            }
                                                        }
                                                    } else {
                                                        _ = len(v4)
                                                    }
                                                    v48 := []int{(-(fi24)), fi26, fi24, fi26}
                                                    _ = v48
                                                case "":
                                                    v49 := GenS4{F0: 94}
                                                    _ = v49
                                                }
                                                var v50 uint = 730
                                                _ = v39
                                                _ = v50
                                            }
                                            v3 += float32(1.5)
                                        }
                                        v3 -= ((v3 + v3) / float32(float64(0.1)))
                                    }
                                case 366:
                                    var v51 float64 = 0.0
                                    if v1 {
                                        v52 := GenS1{F0: (!(909 >= 42)), F1: (v5 / ((v5) | 1))}
                                        if true {
                                            var v53 uint = 18446744073709551615
                                            v1 = v1
                                            if (!(!(255 <= 255))) {
                                                v54 := GenS5{F0: -1, F1: 32767}
                                                _ = v54
                                            } else {
                                                v55 := GenS8{F0: 0, F1: -2147483648, F2: v1}
                                                _ = v55
                                            }
                                            _ = v53
                                        }
                                        _ = v52
                                    } else {
                                        v56 := GenS5{F0: 0, F1: 1000}
                                        _ = v56
                                    }
                                    switch 127 {
                                    case -903:
                                        switch int(v5) {
                                        case 2147483647:
                                            var v57 uint8 = 2
                                            _ = v57
                                        case -240:
                                            var v58 uint64 = 0
                                            _ = v58
                                        case 42:
                                            var v59 uint = 4294967295
                                            _ = v59
                                        default:
                                            if (v51 == (v51 + ((0.0 - v51) * (float64(1e10) / v51)))) {
                                                if (18446744073709551615 <= 4294967295) {
                                                    var v60 bool = true
                                                    _ = v60
                                                } else {
                                                    v3 = v3
                                                }
                                                fmt.Println(v1)
                                            }
                                        }
                                    case 9223372036854775807:
                                        switch 0 {
                                        case 0:
                                            switch 127 {
                                            case 2147483647:
                                                var v61 int64 = 9223372036854775807
                                                _ = v61
                                            }
                                        default:
                                            var v62 int32 = 0
                                            _ = v62
                                        }
                                    default:
                                        fmt.Println(v5, v1, v51)
                                    }
                                    _ = v51
                                default:
                                    v5 += 303
                                }
                                v5 = (v5 | 504)
                                _ = v30
                            }
                            if true {
                                var v63 string = (("hello" + ("hello" + ("a" + ""))) + "a")
                                _ = v63
                            }
                        }
                        _ = fi26
                    }
                }
                v5 += 342
                v64 := GenS9{F0: 65535}
                _ = v64
            } else {
                v65 := v4["foo bar"]
                v66 := GenS3{F0: v5, F1: "a", F2: v1}
                fmt.Println(v5, v1)
                _ = v65
                _ = v66
            }
            _ = fi24
        }
    case "":
        if (v1 || v1) {
            v5 += v5
            var v67 bool = v1
            if ((((!v1) || (-77 == 857)) || false) && false) {
                for fi68 := 0; fi68 < 6; fi68++ {
                    v69 := []int{fi68}
                    _ = fi68
                    _ = v69
                }
            }
            _ = v67
        } else {
            v70 := GenS4{F0: 15}
            _ = v70
        }
    }
    if false {
        if (!(!(2 != 42))) {
            switch -2147483648 {
            case 32767:
                v1 = true
            case 42:
                v71 := GenS2{F0: 32768}
                _ = v71
            case 1:
                if true {
                    var v72 float32 = v3
                    _ = v72
                } else {
                    v5 -= v5
                }
            default:
                v1 = v1
            }
        }
    } else {
        var v73 uint16 = 32767
        if (255 < 128) {
            v3 = 1e10
            v74 := GenS10{F0: (v5 >> ((507) & 63))}
            _ = v74
        } else {
            fmt.Println(v73, v5, v1)
        }
        _ = v73
    }
    for fi75 := 0; fi75 < 6; fi75++ {
        v1 = true
        _ = fi75
    }
    if (uint16(v5) >= 32767) {
        if ((18446744073709551615 == 158) || true) {
            if true {
                v5 = v5
            }
        } else {
            for fi76 := 0; fi76 < 7; fi76++ {
                var v77 float32 = v3
                v5 += 2147483648
                v78 := GenS11{F0: -1, F1: 188, F2: 42}
                _ = fi76
                _ = v77
                _ = v78
            }
        }
        if (true && v1) {
            fmt.Println(v5)
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
