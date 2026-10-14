package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 uint; F1 int }
type GenS3 struct { F0 string; F1 string }
type GenS4 struct { F0 int16 }
type GenS5 struct { F0 int32; F1 uint64 }
type GenS6 struct { F0 string; F1 int32; F2 string }
type GenS7 struct { F0 int64 }
type GenS8 struct { F0 int16; F1 int }
type GenS9 struct { F0 uint32; F1 uint64 }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 int32 }
func genFn0(p0 int64, p1 uint32) bool {
    p1 -= p1
    v7 := []int{-1, -164, -2147483648, 912}
    for fi8 := 0; fi8 < 6; fi8++ {
        if (1000 >= 32767) {
            p0 = 165
            if false {
                var v9 int32 = 42
                switch "hello" {
                case "":
                    v9 = 0
                case "\n\t":
                    p0 = p0
                case "hello":
                    switch "a" {
                    case "foo bar":
                        p1 += p1
                    case "":
                        var v10 int = 2147483647
                        _ = v10
                    case "hello":
                        var v11 uint16 = 255
                        _ = v11
                    default:
                        if len(v7) > 0 {
                            v7[0] = (fi8 % (((fi8 & (^(fi8)))) | 1))
                        }
                    }
                }
                fmt.Println(p1, p0, v9)
                _ = v9
            } else {
                fmt.Println(p0, fi8)
            }
            if false {
                p1 += 0
                var v12 uint = 2
                _ = v12
            }
        }
        _ = fi8
    }
    fmt.Println(p1, p0)
    for fi13 := 0; fi13 < 20; fi13++ {
        p0 -= p0
        _ = fi13
    }
    _ = p0
    _ = p1
    _ = v7
    return true
}
func genFn1(p0 int, p1 uint8) uint32 {
    if (42 <= 2) {
        p1 = p1
        p1 = p1
    } else {
        for fi14 := 0; fi14 < 9; fi14++ {
            p1 += p1
            v15 := GenS1{F0: 18446744073709551615, F1: 938}
            _ = fi14
            _ = v15
        }
        var v16 uint8 = p1
        _ = v16
    }
    if (p0 == (^(p0))) {
        var v17 int64 = 670
        fmt.Println(p0)
        _ = v17
    } else {
        p0 = 128
        v18 := []int{p0, p0, p0}
        _ = v18
    }
    v19 := []int{p0, (p0 >> ((42) & 63)), (p0 * p0)}
    _ = p0
    _ = p1
    _ = v19
    return 0
}

func main() {
    v1 := make(map[string]uint32)
    var v2 uint16 = 32767
    var v3 uint16 = v2
    v4 := []string{("\n\t" + "hello"), "foo bar", ("a" + "\n\t"), ("hello" + "")}
    var v5 int = -2147483648
    v6 := make(map[string]int)
    for fi20 := 0; fi20 < 12; fi20++ {
        if (32767 != -695) {
            for fi21 := 0; fi21 < 11; fi21++ {
                v22 := GenS2{F0: 2, F1: 128}
                _ = fi21
                _ = v22
            }
            if (!false) {
                v23 := GenS3{F0: "hello", F1: "\n\t"}
                _ = v23
            } else {
                if false {
                    v4 = append(v4, ((("\n\t" + ("a" + "\n\t")) + "\n\t") + (("" + "a") + "hello")))
                    fmt.Println(fi20, v3, fi20)
                } else {
                    v24 := GenS2{F0: 18446744073709551615, F1: (fi20 << ((403) & 63))}
                    _ = v24
                }
                if len(v4) > 0 {
                    v4[0] = "\n\t"
                }
                switch fi20 {
                case 735:
                    var v25 int = 9223372036854775807
                    _ = v25
                case 127:
                    switch (-(v5)) {
                    case -2147483648:
                        v26 := GenS4{F0: 1}
                        _ = v26
                    case 1:
                        v4 = append(v4, "hello")
                    case -9223372036854775808:
                        var v27 int = v5
                        _ = v27
                    }
                }
            }
            for fi28 := 0; fi28 < 14; fi28++ {
                if false {
                    if (42 <= -2147483648) {
                        v5 = v5
                        _ = len(v6)
                        var v29 int16 = 1
                        _ = v29
                    } else {
                        v5 = (fi20 >> ((255) & 63))
                        var v30 uint16 = v3
                        v5 = (-(v5))
                        _ = v30
                    }
                    if false {
                        _ = len(v1)
                        switch v5 {
                        case 1:
                            fmt.Println(fi28, v2)
                        case 32768:
                            var v31 float64 = 0.1
                            _ = v31
                        }
                        v5 = fi20
                    } else {
                        v2 += v2
                    }
                } else {
                    v4 = append(v4, (("hello" + "") + (("\n\t" + "hello") + "foo bar")))
                    if (115 != 0) {
                        var v32 int8 = 67
                        _ = v32
                    }
                }
                _ = fi28
            }
        }
        _ = fi20
    }
    if (true || true) {
        var v33 int32 = 42
        v34 := GenS5{F0: v33, F1: 1}
        for fi35 := 0; fi35 < 17; fi35++ {
            v2 = 217
            _ = fi35
        }
        _ = v33
        _ = v34
    } else {
        var v36 uint = 4294967295
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        if true {
            v2 -= v2
            var v37 string = (("" + "foo bar") + ("hello" + ("foo bar" + ("\n\t" + "hello"))))
            _ = v37
        }
        _ = v36
    }
    v38 := GenS2{F0: 421, F1: -9223372036854775808}
    if (0 < 42) {
        if ((v2 & 2) >= 252) {
            switch (v5 ^ (v5 % ((v5) | 1))) {
            case 32768:
                switch (-(v5)) {
                case 29:
                    v3 -= (v3 % (((v3 | (v3 % ((1) | 1)))) | 1))
                }
            case 4294967295:
                if false {
                    var v39 uint16 = v3
                    v3 = (^(v39))
                    if true {
                        fmt.Println(v5, v5)
                    }
                    _ = v39
                } else {
                    fmt.Println(v3, v5, v5)
                }
            }
            var v40 int16 = 322
            v41 := GenS6{F0: "\n\t", F1: 248, F2: ("hello" + "hello")}
            _ = v40
            _ = v41
        } else {
            var v42 string = ""
            v5 = (v5 ^ v5)
            _ = v42
        }
        var v43 int32 = 42
        v2 -= 42
        _ = v43
    }
    for fi44 := 0; fi44 < 6; fi44++ {
        if true {
            var v45 int64 = -182
            v46 := GenS7{F0: (v45 | (v45 % ((v45) | 1)))}
            _ = v45
            _ = v46
        } else {
            var v47 int32 = -226
            if len(v4) > 0 {
                v4[0] = ("" + "hello")
            }
            if len(v4) > 0 {
                v4[0] = "\n\t"
            }
            _ = v47
        }
        _ = fi44
    }
    fmt.Println(v2, v5, v2)
    v3 = 1
    if (0 < 960) {
        for fi48 := 0; fi48 < 4; fi48++ {
            fmt.Println(v3)
            switch (fi48 / ((-291) | 1)) {
            case 4294967295:
                for fi49 := 0; fi49 < 8; fi49++ {
                    if true {
                        v50 := v6["a"]
                        switch ("a" + "") {
                        case "foo bar":
                            v3 = (v3 & v3)
                        case "a":
                            var v51 uint16 = (v3 % (((v2 / (((v2 - v2)) | 1))) | 1))
                            _ = v51
                        case "":
                            var v52 float32 = 1.5
                            _ = v52
                        }
                        var v53 uint64 = 1
                        if ((0.1 * 1e10) <= 0.0) {
                            v3 = 65535
                            v54 := v6[("\n\t" + ("" + "hello"))]
                            if ((true || false) || false) {
                                fmt.Println(v54)
                                v54 = v50
                                v4 = append(v4, "")
                            } else {
                                v50 += (^(fi48))
                                if false {
                                    v3 = v2
                                } else {
                                    if false {
                                        v55 := GenS8{F0: -955, F1: v50}
                                        switch (fi48 | (v50 - v50)) {
                                        case 164:
                                            fmt.Println(v3, v54, v50)
                                        default:
                                            fmt.Println(fi48, v5)
                                        }
                                        var v56 int64 = -1
                                        _ = v55
                                        _ = v56
                                    } else {
                                        v57 := GenS1{F0: 18446744073709551615, F1: 4294967295}
                                        _ = v57
                                    }
                                    v6[("foo bar" + "foo bar")] = -905
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                }
                                v58 := v6[("hello" + (("foo bar" + "hello") + (("a" + "\n\t") + ("hello" + "foo bar"))))]
                                _ = v58
                            }
                            _ = v54
                        } else {
                            if true {
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                fmt.Println(v3, v50)
                            } else {
                                v59 := GenS1{F0: 0, F1: 1}
                                _ = v59
                            }
                        }
                        v3 = (v3 << ((4294967295) & 63))
                        _ = v50
                        _ = v53
                    } else {
                        var v60 bool = false
                        _ = v60
                    }
                    fmt.Println(fi49, v5, v3)
                    v5 += fi49
                    _ = fi49
                }
            default:
                v3 = v3
            }
            v61 := GenS9{F0: 1, F1: 42}
            _ = fi48
            _ = v61
        }
        var v62 uint = 4294967295
        switch "a" {
        case "foo bar":
            var v63 uint32 = 721
            _ = v63
        case "a":
            v5 -= v5
        case "hello":
            var v64 uint8 = 128
            _ = v64
        default:
            var v65 string = ("" + "foo bar")
            _ = v65
        }
        _ = v62
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    switch "a" {
    case "foo bar":
        var v66 uint8 = 0
        _ = v66
    case "":
        fmt.Println(v5, v5)
    }
    var v67 uint = 18446744073709551615
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if false {
        var v68 float64 = 0.1
        var v69 bool = (65535 <= v67)
        _ = v68
        _ = v69
    } else {
        var v70 int8 = 100
        for fi71 := 0; fi71 < 7; fi71++ {
            if (false || (true && true)) {
                v6[("\n\t" + ("foo bar" + (((((("a" + "\n\t") + ((((("a" + "foo bar") + ((((("hello" + "\n\t") + "") + "a") + "") + ("a" + "hello"))) + "hello") + "\n\t") + "")) + "foo bar") + "hello") + "\n\t") + "foo bar")))] = v5
                var v72 uint64 = 7
                _ = v72
            } else {
                v4 = append(v4, "a")
            }
            var v73 uint = (v67 ^ (v67 >> (((v67 + v67)) & 63)))
            for fi74 := 0; fi74 < 2; fi74++ {
                v75 := GenS10{F0: 913}
                v67 += (v67 >> (((v67 >> ((v67) & 63))) & 63))
                _ = fi74
                _ = v75
            }
            _ = fi71
            _ = v73
        }
        var v76 int32 = 256
        _ = v70
        _ = v76
    }
    fmt.Println(v67)
    v3 += (^(v2))
    switch (v5 >> (((v67 / ((v67) | 1))) & 63)) {
    case 677:
        for fi77 := 0; fi77 < 8; fi77++ {
            var v78 uint = (v67 / (((v67 | (v67 | (^(v67))))) | 1))
            fmt.Println(v78, v78)
            switch 32767 {
            case 4294967295:
                v79 := GenS3{F0: "", F1: (("" + "foo bar") + "\n\t")}
                _ = v79
            default:
                for fi80 := 0; fi80 < 20; fi80++ {
                    if true {
                        var v81 int8 = -2
                        switch fi80 {
                        case 9223372036854775807:
                            var v82 float64 = 0.1
                            _ = v82
                        case -722:
                            v83 := GenS10{F0: 32767}
                            _ = v83
                        case 32767:
                            fmt.Println(v81, v5)
                        }
                        if true {
                            v84 := v1[("foo bar" + ("hello" + "hello"))]
                            _ = v84
                        }
                        _ = v81
                    }
                    v67 = v67
                    _ = fi80
                }
            }
            _ = fi77
            _ = v78
        }
    case -786:
        v85 := GenS6{F0: "a", F1: 1, F2: "a"}
        _ = v85
    }
    if false {
        v3 += v3
        v86 := GenS11{F0: 0}
        _ = v86
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v67:", v67)
    _ = v1
    _ = v2
    _ = v3
    _ = v38
    _ = v4
    _ = v5
    _ = v6
    _ = v67
}
