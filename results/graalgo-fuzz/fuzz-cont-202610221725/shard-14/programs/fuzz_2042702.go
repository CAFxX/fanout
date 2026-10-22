package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32 }
type GenS2 struct { F0 int32; F1 int16; F2 int64 }
type GenS3 struct { F0 uint64; F1 uint32; F2 uint16 }
type GenS4 struct { F0 int16; F1 int8; F2 uint16 }
type GenS5 struct { F0 uint16; F1 int }
type GenS6 struct { F0 int; F1 int8; F2 uint16 }
type GenS7 struct { F0 uint64; F1 int16; F2 uint64 }
type GenS8 struct { F0 uint64; F1 int64; F2 uint32 }
type GenS9 struct { F0 uint8; F1 int16; F2 uint32 }
type GenS10 struct { F0 uint64 }
func genFn0() uint64 {
    if true {
        var v6 int = 210
        if true {
            var v7 uint = 1
            v6 -= v6
            _ = v7
        }
        _ = v6
    }
    if true {
        switch "foo bar" {
        case "a":
            var v8 string = "a"
            _ = v8
        case "":
            for fi9 := 0; fi9 < 11; fi9++ {
                v10 := []int{636, (fi9 << ((4294967295) & 63)), -2147483648}
                v11 := GenS1{F0: 0, F1: 0}
                var v12 bool = false
                _ = fi9
                _ = v10
                _ = v11
                _ = v12
            }
        }
    } else {
        var v13 int8 = -1
        for fi14 := 0; fi14 < 9; fi14++ {
            v15 := GenS2{F0: 270, F1: -344, F2: 803}
            for fi16 := 0; fi16 < 17; fi16++ {
                v17 := []int{fi16, fi16, 32768, fi16}
                var v18 uint64 = 4294967295
                _ = fi16
                _ = v17
                _ = v18
            }
            var v19 string = ("foo bar" + ("a" + ""))
            _ = fi14
            _ = v15
            _ = v19
        }
        switch "\n\t" {
        case "\n\t":
            switch 950 {
            case 32767:
                var v20 int8 = -1
                _ = v20
            default:
                var v21 int32 = -698
                _ = v21
            }
        case "foo bar":
            v22 := GenS3{F0: 1, F1: 0, F2: 69}
            _ = v22
        default:
            if false {
                fmt.Println(v13)
                var v23 int16 = 876
                v24 := GenS3{F0: 970, F1: 42, F2: 942}
                _ = v23
                _ = v24
            } else {
                if true {
                    fmt.Println(v13)
                } else {
                    if true {
                        v13 += (-(v13))
                    } else {
                        v13 = v13
                        if (!false) {
                            var v25 int16 = -32768
                            for fi26 := 0; fi26 < 20; fi26++ {
                                v25 += v25
                                var v27 uint16 = 65535
                                _ = fi26
                                _ = v27
                            }
                            _ = v25
                        } else {
                            v28 := []int{9223372036854775807, 0, -701}
                            for fi29 := 0; fi29 < 12; fi29++ {
                                if len(v28) > 0 {
                                    v28[0] = (^(fi29))
                                }
                                _ = fi29
                            }
                            v30 := GenS4{F0: 1, F1: v13, F2: 22}
                            _ = v28
                            _ = v30
                        }
                    }
                    for fi31 := 0; fi31 < 2; fi31++ {
                        v32 := []int{32767, fi31}
                        if (3.14159 == ((1.5 * 1e10) + 1e10)) {
                            var v33 string = (("" + "hello") + "")
                            v33 = (("a" + v33) + "hello")
                            switch 9223372036854775807 {
                            case -9223372036854775808:
                                v32 = append(v32, (fi31 & fi31))
                            case -657:
                                v33 = (v33 + ("\n\t" + v33))
                            }
                            _ = v33
                        } else {
                            fmt.Println(fi31, fi31)
                        }
                        for fi34 := 0; fi34 < 4; fi34++ {
                            var v35 string = ("" + "foo bar")
                            fmt.Println(v35, v13)
                            v13 += v13
                            _ = fi34
                            _ = v35
                        }
                        _ = fi31
                        _ = v32
                    }
                }
                if false {
                    v13 = 42
                } else {
                    var v36 string = (("a" + "foo bar") + "foo bar")
                    _ = v36
                }
            }
        }
        _ = v13
    }
    var v37 uint16 = 32767
    _ = v37
    return 2
}
func genFn1(p0 uint16, p1 int16) bool {
    var v38 bool = true
    var v39 bool = ((p1 != p1) || ("a" >= ""))
    switch "a" {
    case "":
        v40 := GenS5{F0: p0, F1: 18}
        _ = v40
    case "foo bar":
        var v41 int8 = 88
        _ = v41
    case "a":
        for fi42 := 0; fi42 < 8; fi42++ {
            v43 := []int{(fi42 - 32768), -783, fi42}
            switch (("\n\t" + "") + "foo bar") {
            case "hello":
                _ = len(v43)
            case "foo bar":
                if len(v43) > 0 {
                    v43[0] = (fi42 / (((fi42 + fi42)) | 1))
                }
            default:
                for fi44 := 0; fi44 < 6; fi44++ {
                    switch "" {
                    case "hello":
                        fmt.Println(fi44, fi42)
                    case "a":
                        var v45 string = "a"
                        _ = v45
                    case "":
                        if len(v43) >= 2 {
                            v43 = v43[1:len(v43)]
                        }
                    default:
                        if (!v39) {
                            if len(v43) >= 2 {
                                v43 = v43[1:len(v43)]
                            }
                            switch (("a" + "a") + ("\n\t" + "foo bar")) {
                            case "hello":
                                var v46 bool = (int8(fi42) > 29)
                                _ = v46
                            case "foo bar":
                                var v47 uint = 65535
                                _ = v47
                            case "":
                                switch ("a" + ("a" + "a")) {
                                case "foo bar":
                                    if (!(fi44 >= fi44)) {
                                        var v48 string = "hello"
                                        _ = v48
                                    } else {
                                        p0 -= p0
                                        v38 = (false && (!((1e10 - ((1e10 - 1e10) / 1.5)) >= ((3.14159 * 0.0) * 3.14159))))
                                        var v49 uint = 0
                                        _ = v49
                                    }
                                default:
                                    var v50 int64 = 243
                                    fmt.Println(p1, fi44, v39)
                                    _ = len(v43)
                                    _ = v50
                                }
                            default:
                                switch ("hello" + "\n\t") {
                                case "hello":
                                    if (1 > 255) {
                                        p1 -= (p1 & (-(p1)))
                                        var v51 uint = 21
                                        _ = v51
                                    }
                                case "a":
                                    p1 += p1
                                }
                            }
                            var v52 uint8 = 128
                            if v39 {
                                v38 = v39
                                var v53 int16 = p1
                                var v54 bool = (((-2.25 + float32(((0.1 / 0.1) * 0.0))) != 0.1) && false)
                                _ = v53
                                _ = v54
                            }
                            _ = v52
                        }
                    }
                    var v55 string = "hello"
                    _ = fi44
                    _ = v55
                }
            }
            _ = fi42
            _ = v43
        }
    }
    _ = p0
    _ = p1
    _ = v38
    _ = v39
    return ((4294967295 >= 255) || true)
}

func main() {
    v1 := make(map[uint32]uint32)
    v2 := []int{32768, -9223372036854775808, 32767}
    var v3 int8 = 0
    v4 := []int{-2147483648, 42, -2147483648, 127}
    v5 := []string{"\n\t", "hello", "\n\t", "a"}
    var v56 bool = true
    var v57 int8 = -55
    delete(v1, 0)
    fmt.Println(v56, v57, v56)
    switch "" {
    case "hello":
        v3 -= 82
    default:
        for fi58 := 0; fi58 < 1; fi58++ {
            v3 = v3
            v1[0] = 914
            for fi59 := 0; fi59 < 11; fi59++ {
                v56 = v56
                fmt.Println(v56, fi58, v57)
                _ = fi59
            }
            _ = fi58
        }
    }
    var v60 uint32 = 985
    v61 := GenS1{F0: -1, F1: 876}
    for fi62 := 0; fi62 < 10; fi62++ {
        switch fi62 {
        case -9223372036854775808:
            fmt.Println(v57, fi62, fi62)
        }
        switch ("" + "a") {
        case "hello":
            switch "\n\t" {
            case "a":
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case "\n\t":
                v63 := v1[75]
                _ = v63
            case "foo bar":
                var v64 uint16 = 32767
                _ = v64
            default:
                var v65 uint16 = 65535
                _ = v65
            }
        case "\n\t":
            v1[(v60 * 0)] = v60
        case "":
            v66 := v1[(v60 - v60)]
            _ = v66
        }
        v57 -= int8(v60)
        _ = fi62
    }
    switch (("" + ("" + (("" + "") + "foo bar"))) + ((("a" + (("" + (("foo bar" + "a") + "foo bar")) + ("foo bar" + ("\n\t" + "a")))) + ("" + "")) + (("hello" + "hello") + ""))) {
    case "\n\t":
        _ = len(v1)
    case "hello":
        v67 := GenS6{F0: 32768, F1: v57, F2: 1}
        _ = v67
    default:
        var v68 int8 = v3
        _ = v68
    }
    fmt.Println(v3)
    for fi69 := 0; fi69 < 11; fi69++ {
        v3 -= (v3 - v57)
        if (true || true) {
            var v70 bool = v56
            if true {
                if (!false) {
                    fmt.Println(v57)
                    v57 = int8(v60)
                    for fi71 := 0; fi71 < 9; fi71++ {
                        v72 := GenS7{F0: 1, F1: 290, F2: 18446744073709551615}
                        var v73 string = ""
                        _ = fi71
                        _ = v72
                        _ = v73
                    }
                }
            }
            switch "hello" {
            case "\n\t":
                switch (fi69 & (fi69 << ((1) & 63))) {
                case -2147483648:
                    if true {
                        var v74 float64 = 0.0
                        if len(v4) > 0 {
                            v4[0] = (fi69 ^ 4294967295)
                        }
                        _ = v74
                    } else {
                        var v75 uint16 = uint16(v60)
                        if false {
                            v75 += (v75 ^ v75)
                            v76 := GenS4{F0: 42, F1: v57, F2: (v75 | v75)}
                            _ = v76
                        }
                        switch "" {
                        case "\n\t":
                            v56 = v70
                        case "a":
                            v77 := GenS8{F0: genFn0(), F1: -9223372036854775808, F2: (v60 + (^(v60)))}
                            _ = v77
                        }
                        _ = v75
                    }
                case -1:
                    switch (fi69 % ((42) | 1)) {
                    case -9223372036854775808:
                        v57 -= v57
                    case 668:
                        for fi78 := 0; fi78 < 8; fi78++ {
                            var v79 float64 = 1.5
                            v80 := GenS9{F0: 128, F1: 1, F2: v60}
                            _ = fi78
                            _ = v79
                            _ = v80
                        }
                    default:
                        v57 = (v57 >> ((254) & 63))
                    }
                default:
                    switch "a" {
                    case "a":
                        _ = len(v4)
                    default:
                        v81 := GenS1{F0: 2147483647, F1: -833}
                        _ = v81
                    }
                }
            }
            _ = v70
        }
        _ = fi69
    }
    v82 := GenS10{F0: 1}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v56:", v56)
    fmt.Println("v57:", v57)
    fmt.Println("v60:", v60)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v56
    _ = v57
    _ = v60
    _ = v61
    _ = v82
}
