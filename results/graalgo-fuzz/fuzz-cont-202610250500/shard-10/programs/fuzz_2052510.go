package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32 }
type GenS2 struct { F0 bool; F1 int }
type GenS3 struct { F0 uint; F1 int32 }
type GenS4 struct { F0 int64; F1 int16; F2 string }
type GenS5 struct { F0 string; F1 uint16 }
type GenS6 struct { F0 int8; F1 int }
type GenS7 struct { F0 uint64; F1 uint; F2 int }
type GenS8 struct { F0 int16; F1 uint32; F2 uint16 }
func genFn0(p0 uint8, p1 int32, p2 uint32) string {
    v7 := []int{127, 32768, -9223372036854775808, 4294967295}
    p1 += (p1 * (p1 | p1))
    fmt.Println(p1, p2)
    var v8 float32 = -2.25
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return "foo bar"
}

func main() {
    v1 := []string{"a", "foo bar"}
    v2 := []int{0, 1}
    v3 := make(map[string]string)
    var v4 int8 = 127
    var v5 int8 = v4
    var v6 uint = 18446744073709551615
    v4 += 0
    if ((507 >= -1000) && false) {
        for fi9 := 0; fi9 < 16; fi9++ {
            for fi10 := 0; fi10 < 16; fi10++ {
                var v11 int = 128
                v12 := GenS1{F0: v11, F1: 4294967295}
                if true {
                    var v13 uint16 = 32767
                    v5 += (v4 / ((v4) | 1))
                    _ = v13
                }
                if (42 == 127) {
                    if true {
                        v4 = int8(fi9)
                    }
                    v5 += v5
                    if ((false || ((42 < 9223372036854775807) && (3.14159 <= 1.5))) && (231 != 2147483647)) {
                        var v14 int64 = 1
                        var v15 uint = (v6 >> ((v6) & 63))
                        if false {
                            fmt.Println(fi9)
                            fmt.Println(v15)
                        } else {
                            if len(v1) > 0 {
                                v1[0] = "\n\t"
                            }
                            v16 := GenS2{F0: true, F1: (fi9 << ((v6) & 63))}
                            var v17 bool = (true || true)
                            _ = v16
                            _ = v17
                        }
                        _ = v14
                        _ = v15
                    } else {
                        v18 := GenS2{F0: (true && true), F1: (v11 | (-(fi10)))}
                        switch (fi9 | fi10) {
                        case 127:
                            if len(v1) >= 2 {
                                v1 = v1[1:len(v1)]
                            }
                        default:
                            var v19 int = (v11 >> ((v6) & 63))
                            _ = v19
                        }
                        _ = v18
                    }
                }
                var v20 uint64 = 34
                v21 := GenS2{F0: false, F1: 32768}
                switch fi9 {
                case -9223372036854775808:
                    var v22 int32 = 0
                    _ = v22
                }
                _ = fi10
                _ = v11
                _ = v12
                _ = v20
                _ = v21
            }
            for fi23 := 0; fi23 < 19; fi23++ {
                fmt.Println(v6)
                fmt.Println(fi9, fi23)
                fmt.Println(v6, fi9)
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                if (42 == 2) {
                    fmt.Println(v5, fi23, fi23)
                    v5 = v5
                } else {
                    if (v6 < 252) {
                        v5 = (v5 * v5)
                        fmt.Println(fi9, v5)
                        var v24 bool = (((130 >= 128) || (!(1e10 == 1e10))) && false)
                        _ = v24
                    } else {
                        var v25 uint16 = 65535
                        var v26 string = "foo bar"
                        _ = v25
                        _ = v26
                    }
                    var v27 uint8 = 44
                    _ = v27
                }
                switch "foo bar" {
                case "\n\t":
                    fmt.Println(fi23, fi9)
                }
                if (-9223372036854775808 > 861) {
                    switch (fi9 | -2147483648) {
                    case 127:
                        v4 -= 100
                    case -2147483648:
                        v28 := v3["hello"]
                        _ = v28
                    case 921:
                        if ((!((!(true && (225 == 65535))) && (3.14159 != (1.5 * 1.5)))) && true) {
                            if len(v1) > 0 {
                                v1[0] = "\n\t"
                            }
                        }
                    default:
                        fmt.Println(v5, fi23, fi9)
                    }
                }
                _ = fi23
            }
            v29 := GenS2{F0: (!false), F1: fi9}
            _ = fi9
            _ = v29
        }
    }
    if ("a" > "") {
        if (!(true || true)) {
            fmt.Println(v5, v4)
            _ = len(v1)
            v4 = (v5 / ((1) | 1))
        }
        for fi30 := 0; fi30 < 13; fi30++ {
            switch "hello" {
            case "\n\t":
                fmt.Println(v5, v4, v6)
            case "foo bar":
                v31 := GenS2{F0: false, F1: int(v6)}
                _ = v31
            }
            fmt.Println(v6)
            var v32 uint16 = 642
            _ = fi30
            _ = v32
        }
    }
    _ = len(v2)
    if false {
        v33 := GenS1{F0: -958, F1: 42}
        _ = v33
    }
    v5 -= (v4 - 127)
    for fi34 := 0; fi34 < 4; fi34++ {
        if (!true) {
            for fi35 := 0; fi35 < 16; fi35++ {
                var v36 uint8 = 42
                _ = len(v3)
                _ = fi35
                _ = v36
            }
            var v37 int64 = -9223372036854775808
            switch "hello" {
            case "hello":
                for fi38 := 0; fi38 < 18; fi38++ {
                    fmt.Println(v4, v6, fi34)
                    var v39 bool = false
                    v40 := GenS2{F0: v39, F1: (-(fi34))}
                    v41 := GenS1{F0: (fi38 >> (((^(v6))) & 63)), F1: 42}
                    _ = fi38
                    _ = v39
                    _ = v40
                    _ = v41
                }
            case "a":
                fmt.Println(v4, v4, v4)
            default:
                for fi42 := 0; fi42 < 3; fi42++ {
                    fmt.Println(v37)
                    _ = fi42
                }
            }
            _ = v37
        } else {
            var v43 float64 = 3.14159
            _ = v43
        }
        for fi44 := 0; fi44 < 2; fi44++ {
            if (!(((false && true) && false) || true)) {
                switch "a" {
                case "\n\t":
                    switch "\n\t" {
                    case "":
                        var v45 string = "foo bar"
                        _ = v45
                    default:
                        v46 := GenS3{F0: (v6 >> (((v6 & 65535)) & 63)), F1: 0}
                        _ = v46
                    }
                case "":
                    v3["\n\t"] = "a"
                case "hello":
                    _ = len(v3)
                default:
                    switch "\n\t" {
                    case "hello":
                        if false {
                            v6 = (v6 << (((v6 ^ 0)) & 63))
                            if true {
                                v6 += v6
                            }
                        } else {
                            _ = len(v3)
                        }
                    case "":
                        fmt.Println(v5, fi34, v5)
                    case "\n\t":
                        v47 := GenS4{F0: 1, F1: -922, F2: ""}
                        _ = v47
                    default:
                        if true {
                            var v48 int32 = 42
                            if (!false) {
                                var v49 float32 = (1.5 - (0.0 - 1.5))
                                fmt.Println(v49)
                                var v50 string = "foo bar"
                                _ = len(v2)
                                _ = v49
                                _ = v50
                            }
                            if (3.14159 != (0.1 + 3.14159)) {
                                v51 := GenS4{F0: -9223372036854775808, F1: 825, F2: "\n\t"}
                                _ = v51
                            }
                            _ = v48
                        } else {
                            switch 492 {
                            case 935:
                                if true {
                                    var v52 bool = (!(!false))
                                    _ = v52
                                } else {
                                    v1 = append(v1, ((("a" + ("\n\t" + (("a" + ((("hello" + "a") + "a") + "")) + ""))) + "a") + ((("\n\t" + "foo bar") + "\n\t") + "")))
                                }
                            case 127:
                                var v53 uint8 = 1
                                _ = v53
                            default:
                                var v54 int = (-(fi34))
                                _ = v54
                            }
                            fmt.Println(v5, v5)
                            if len(v2) > 0 {
                                v2[0] = fi34
                            }
                        }
                    }
                }
                var v55 float64 = 0.0
                _ = v55
            }
            v5 -= v4
            var v56 uint8 = 1
            _ = fi44
            _ = v56
        }
        _ = fi34
    }
    for fi57 := 0; fi57 < 10; fi57++ {
        for fi58 := 0; fi58 < 13; fi58++ {
            if true {
                var v59 string = (("a" + "\n\t") + "a")
                v4 += (v4 + (-(v5)))
                v60 := GenS4{F0: -1, F1: -1, F2: ("\n\t" + v59)}
                _ = v59
                _ = v60
            } else {
                v61 := GenS5{F0: "\n\t", F1: 32767}
                v5 = v4
                v62 := GenS6{F0: v5, F1: 128}
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
                _ = v61
                _ = v62
            }
            _ = fi58
        }
        _ = fi57
    }
    v6 = (v6 & (v6 * (v6 / ((v6) | 1))))
    v4 = v4
    v63 := GenS7{F0: 2, F1: v6, F2: -980}
    switch ("foo bar" + (("a" + ("a" + ("hello" + ("" + "hello")))) + ("foo bar" + "\n\t"))) {
    case "a":
        if (true || (1 == 1)) {
            switch 339 {
            case 2147483647:
                fmt.Println(v5, v4, v4)
            case -333:
                if (false && (!(!((v5 > 28) && (((0.1 * 0.1) >= 0.1) && true))))) {
                    var v64 float32 = -2.25
                    v5 = v5
                    if true {
                        if true {
                            if (2147483648 > 2147483648) {
                                v6 = (v6 | 4294967295)
                                var v65 uint16 = 255
                                var v66 uint = (^(v6))
                                _ = v65
                                _ = v66
                            } else {
                                v64 += float32((0.0 / 3.14159))
                                fmt.Println(v4)
                            }
                        }
                        if true {
                            for fi67 := 0; fi67 < 11; fi67++ {
                                if true {
                                    for fi68 := 0; fi68 < 19; fi68++ {
                                        v69 := GenS6{F0: v4, F1: (-(fi67))}
                                        v70 := GenS8{F0: -737, F1: 385, F2: 1}
                                        var v71 int8 = (v5 & (v4 << (((^(v6))) & 63)))
                                        _ = fi68
                                        _ = v69
                                        _ = v70
                                        _ = v71
                                    }
                                    v64 = 0.0
                                } else {
                                    var v72 float64 = float64(0.1)
                                    fmt.Println(v72)
                                    _ = v72
                                }
                                _ = fi67
                            }
                            if true {
                                v64 = v64
                                var v73 uint32 = 2
                                v5 = 127
                                _ = v73
                            }
                        } else {
                            var v74 uint8 = 2
                            _ = v74
                        }
                        if true {
                            if (((3.14159 * 0.0) / 1e10) <= float64(float32((float64(v64) - 1.5)))) {
                                switch (("" + ("\n\t" + "\n\t")) + "hello") {
                                case "\n\t":
                                    v6 = (v6 - (v6 << (((v6 / ((v6) | 1))) & 63)))
                                case "":
                                    v75 := GenS2{F0: false, F1: 715}
                                    _ = v75
                                }
                                v3["hello"] = ("\n\t" + ("\n\t" + "a"))
                            }
                        } else {
                            v76 := GenS3{F0: (v6 + v6), F1: 1}
                            _ = v76
                        }
                    } else {
                        v77 := GenS3{F0: 2, F1: 287}
                        var v78 int = -2147483648
                        fmt.Println(v5, v4)
                        _ = v77
                        _ = v78
                    }
                    _ = v64
                }
            }
            var v79 uint64 = 2
            _ = v79
        }
    case "foo bar":
        if (v5 == v5) {
            for fi80 := 0; fi80 < 13; fi80++ {
                fmt.Println(v5, v6)
                if true {
                    v4 -= -1
                    if ((^(v6)) != 255) {
                        fmt.Println(v6, v4, v4)
                        v81 := GenS5{F0: "\n\t", F1: 32767}
                        v1 = append(v1, "")
                        _ = v81
                    } else {
                        v4 -= v4
                    }
                    for fi82 := 0; fi82 < 7; fi82++ {
                        var v83 uint8 = uint8(v6)
                        v5 -= v4
                        _ = fi82
                        _ = v83
                    }
                }
                _ = fi80
            }
        } else {
            fmt.Println(v5, v4)
        }
    }
    for fi84 := 0; fi84 < 13; fi84++ {
        var v85 int16 = 30
        _ = fi84
        _ = v85
    }
    var v86 int16 = 1
    v2 = append(v2, 4294967295)
    if true {
        v5 += v5
    } else {
        v4 = (v5 ^ v5)
    }
    v5 = v5
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v86:", v86)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v63
    _ = v86
}
