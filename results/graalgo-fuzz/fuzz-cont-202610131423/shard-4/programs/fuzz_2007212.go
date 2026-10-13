package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int8; F1 bool; F2 string }
type GenS5 struct { F0 int32 }
type GenS6 struct { F0 bool; F1 uint16 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 string }

func main() {
    v1 := make(map[uint32]bool)
    v2 := []string{"\n\t", "", ("hello" + "foo bar"), ""}
    var v3 uint16 = 1
    var v4 int = 42
    var v5 int32 = 42
    v6 := make(map[int]bool)
    fmt.Println(v4)
    var v7 uint = 1
    v7 += 18446744073709551615
    if (2147483647 != (v4 >> ((0) & 63))) {
        var v8 uint = 0
        v4 += (v4 % (((^(v4))) | 1))
        _ = v8
    } else {
        v4 = v4
        var v9 uint = uint(v3)
        _ = v9
    }
    switch "\n\t" {
    case "\n\t":
        v10 := GenS1{F0: ((("foo bar" + ("hello" + ("hello" + "\n\t"))) <= (("foo bar" + "\n\t") + "a")) && (!(((3.14159 + 0.1) != float32(3.14159)) && false))), F1: 2}
        _ = v10
    case "":
        var v11 uint = (v7 >> ((v7) & 63))
        _ = v11
    }
    v7 = v7
    for fi12 := 0; fi12 < 11; fi12++ {
        v13 := GenS1{F0: false, F1: v3}
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 12; fi14++ {
        var v15 int32 = 2147483647
        for fi16 := 0; fi16 < 10; fi16++ {
            if (!false) {
                var v17 uint16 = v3
                var v18 float64 = ((3.14159 + float64(3.14159)) * 1.5)
                fmt.Println(v3, v4)
                if (("hello" + ("" + "hello")) != "hello") {
                    var v19 uint64 = 370
                    fmt.Println(v3)
                    _ = v19
                }
                var v20 float64 = 0.1
                _ = v17
                _ = v18
                _ = v20
            }
            _ = fi16
        }
        for fi21 := 0; fi21 < 15; fi21++ {
            v15 = int32(v7)
            if true {
                var v22 string = "a"
                v23 := GenS1{F0: (300 < 496), F1: 255}
                v4 += (fi21 % (((fi21 >> ((v7) & 63))) | 1))
                _ = v22
                _ = v23
            } else {
                if (!true) {
                    _ = len(v2)
                }
                if ((42 <= 1) && true) {
                    v24 := GenS2{F0: 255}
                    if (false && (!false)) {
                        v4 = fi14
                    }
                    _ = v24
                }
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            }
            if (-1 > -1) {
                var v25 uint64 = 482
                var v26 uint32 = 2147483647
                if (!(v4 < int(v3))) {
                    fmt.Println(v25, fi21, v26)
                    fmt.Println(v7, v5, v26)
                } else {
                    fmt.Println(fi21, v15, v25)
                    v5 = (v5 | v5)
                    if true {
                        v26 -= v26
                        if false {
                            switch "\n\t" {
                            case "\n\t":
                                fmt.Println(v3, v26)
                            default:
                                var v27 string = ("\n\t" + ("hello" + "a"))
                                _ = v27
                            }
                            v6[9223372036854775807] = (false && true)
                            if len(v2) > 0 {
                                v2[0] = ("hello" + "a")
                            }
                        }
                    } else {
                        if ("\n\t" < "hello") {
                            if (!((false || true) || true)) {
                                var v28 uint = (v7 | v7)
                                var v29 uint64 = v25
                                v30 := GenS1{F0: false, F1: v3}
                                _ = v28
                                _ = v29
                                _ = v30
                            }
                            if len(v2) > 0 {
                                v2[0] = ("a" + ("\n\t" + "hello"))
                            }
                            if true {
                                switch "\n\t" {
                                case "a":
                                    var v31 int8 = 15
                                    fmt.Println(fi21, v5, v7)
                                    if (v4 >= 9223372036854775807) {
                                        switch fi21 {
                                        case 42:
                                            fmt.Println(v15, fi21)
                                        case -761:
                                            if (false || (117 < 128)) {
                                                var v32 int32 = v15
                                                var v33 int64 = -155
                                                _ = v32
                                                _ = v33
                                            }
                                        }
                                        v34 := GenS3{F0: 44}
                                        if true {
                                            delete(v1, v26)
                                            if false {
                                                switch fi21 {
                                                case 515:
                                                    var v35 string = ("hello" + (("hello" + "hello") + "a"))
                                                    _ = v35
                                                case -2147483648:
                                                    v36 := GenS4{F0: -5, F1: (true || false), F2: (("hello" + "foo bar") + (("a" + "") + "foo bar"))}
                                                    _ = v36
                                                default:
                                                    if (fi14 < fi14) {
                                                        v6[(v4 * -119)] = (((!false) && (452 == 0)) && false)
                                                        if (!(!(v26 == uint32(v25)))) {
                                                            if len(v2) > 0 {
                                                                v2[0] = "\n\t"
                                                            }
                                                        }
                                                        v15 -= v15
                                                    }
                                                }
                                                switch "a" {
                                                case "hello":
                                                    switch int(v25) {
                                                    case 127:
                                                        v37 := GenS4{F0: 55, F1: false, F2: (("\n\t" + "") + (("" + "hello") + ""))}
                                                        _ = v37
                                                    case 32767:
                                                        var v38 int64 = int64(v26)
                                                        _ = v38
                                                    case 0:
                                                        fmt.Println(v5)
                                                        v31 += v31
                                                    default:
                                                        if (false && (!true)) {
                                                            if false {
                                                                if len(v2) > 0 {
                                                                    v2[0] = "hello"
                                                                }
                                                                switch "" {
                                                                case "\n\t":
                                                                    var v39 int = 32768
                                                                    _ = v39
                                                                case "":
                                                                    if (!((!(false || (!(-1 != 32767)))) && ((592 >= (^(v26))) || true))) {
                                                                        var v40 float64 = 3.14159
                                                                        var v41 bool = true
                                                                        _ = v40
                                                                        _ = v41
                                                                    }
                                                                }
                                                                v2 = append(v2, ("" + "\n\t"))
                                                            }
                                                            var v42 uint64 = 42
                                                            var v43 float64 = 1.5
                                                            switch fi21 {
                                                            case -843:
                                                                var v44 int8 = int8(v25)
                                                                _ = v44
                                                            default:
                                                                switch ("a" + "hello") {
                                                                case "hello":
                                                                    fmt.Println(v25)
                                                                case "a":
                                                                    v45 := GenS5{F0: int32(v3)}
                                                                    _ = v45
                                                                default:
                                                                    var v46 float32 = 1.5
                                                                    _ = v46
                                                                }
                                                                switch ("" + "\n\t") {
                                                                case "foo bar":
                                                                    var v47 uint32 = (v26 % (((^(v26))) | 1))
                                                                    _ = v47
                                                                case "hello":
                                                                    if (!false) {
                                                                        var v48 uint8 = 2
                                                                        _ = v48
                                                                    } else {
                                                                        fmt.Println(v15)
                                                                    }
                                                                }
                                                                v49 := GenS5{F0: (-(v5))}
                                                                if true {
                                                                    fmt.Println(v15)
                                                                    var v50 int64 = -659
                                                                    _ = v50
                                                                } else {
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                    v51 := GenS4{F0: int8(v15), F1: true, F2: ""}
                                                                    _ = v51
                                                                }
                                                                _ = len(v6)
                                                                _ = v49
                                                            }
                                                            var v52 bool = ("\n\t" <= "")
                                                            _ = v42
                                                            _ = v43
                                                            _ = v52
                                                        } else {
                                                            v15 -= (v5 % (((v5 ^ 0)) | 1))
                                                            _ = len(v6)
                                                            fmt.Println(v15, fi14)
                                                        }
                                                    }
                                                }
                                            }
                                        } else {
                                            var v53 uint16 = (v3 - (v3 >> ((185) & 63)))
                                            if (!(false && ((v15 + (v5 | -2147483648)) <= v15))) {
                                                if false {
                                                    if (!false) {
                                                        v31 -= v31
                                                        if false {
                                                            if false {
                                                                var v54 int32 = v15
                                                                v26 = v26
                                                                if (0 < 42) {
                                                                    var v55 uint32 = v26
                                                                    if true {
                                                                        var v56 int = (fi21 | (v4 << ((65535) & 63)))
                                                                        if len(v2) > 0 {
                                                                            v2[0] = "\n\t"
                                                                        }
                                                                        _ = v56
                                                                    } else {
                                                                        v5 = 0
                                                                        switch fi14 {
                                                                        case -1:
                                                                            if (true || (("foo bar" + ("foo bar" + "foo bar")) >= "")) {
                                                                                v57 := GenS6{F0: (float64((1.5 + 0.1)) < 1.5), F1: (^(v53))}
                                                                                if false {
                                                                                    v6[v4] = false
                                                                                    if true {
                                                                                        if (v3 <= v53) {
                                                                                            var v58 float64 = (0.1 * (((0.1 / 1e10) / 1.0) * float64(((-2.25 * ((0.0 - 0.0) * 1.5)) - 0.1))))
                                                                                            if len(v2) >= 2 {
                                                                                                v2 = v2[1:len(v2)]
                                                                                            }
                                                                                            v59 := GenS5{F0: 1}
                                                                                            _ = v58
                                                                                            _ = v59
                                                                                        }
                                                                                        v60 := GenS7{F0: (("\n\t" + "") + "\n\t")}
                                                                                        _ = v60
                                                                                    }
                                                                                } else {
                                                                                    fmt.Println(fi14)
                                                                                    switch (fi21 + (fi21 | (fi21 * fi14))) {
                                                                                    case -2147483648:
                                                                                        v3 += (v53 & v3)
                                                                                        var v61 uint32 = (v26 - (v26 << (((v7 ^ (v7 & 0))) & 63)))
                                                                                        _ = v61
                                                                                    default:
                                                                                        var v62 float32 = -2.25
                                                                                        _ = v62
                                                                                    }
                                                                                }
                                                                                delete(v1, (v55 << ((v7) & 63)))
                                                                                _ = v57
                                                                            }
                                                                        case 32767:
                                                                            switch (v4 << ((18446744073709551615) & 63)) {
                                                                            case -9223372036854775808:
                                                                                v5 = 0
                                                                            case 1:
                                                                                v55 = v55
                                                                            case 32768:
                                                                                fmt.Println(v4)
                                                                            default:
                                                                                v26 += 226
                                                                            }
                                                                        }
                                                                    }
                                                                    _ = v55
                                                                } else {
                                                                    v26 = (v26 % ((v26) | 1))
                                                                    if (((v7 - (^(v7))) > (v7 - (v7 + v7))) && ((true && (128 == 1)) && (!(0 <= 32767)))) {
                                                                        v26 -= v26
                                                                        v63 := GenS3{F0: 255}
                                                                        _ = v63
                                                                    } else {
                                                                        v1[v26] = false
                                                                        switch "a" {
                                                                        case "":
                                                                            v3 -= v53
                                                                        default:
                                                                            var v64 float32 = 1e10
                                                                            _ = v64
                                                                        }
                                                                    }
                                                                    var v65 int64 = -524
                                                                    _ = v65
                                                                }
                                                                _ = v54
                                                            } else {
                                                                v25 += (v25 + v25)
                                                                switch "\n\t" {
                                                                case "a":
                                                                    var v66 uint16 = 65535
                                                                    _ = v66
                                                                case "foo bar":
                                                                    v67 := GenS4{F0: (-(v31)), F1: (!(1e10 <= 0.1)), F2: ("foo bar" + "foo bar")}
                                                                    _ = v67
                                                                default:
                                                                    v3 = (v53 ^ 1)
                                                                    v68 := GenS7{F0: ("foo bar" + (("a" + ("hello" + ("" + "hello"))) + "a"))}
                                                                    if len(v2) > 0 {
                                                                        v2[0] = "foo bar"
                                                                    }
                                                                    _ = v68
                                                                }
                                                                fmt.Println(v7)
                                                            }
                                                            v15 -= (-(v15))
                                                            var v69 int32 = v15
                                                            _ = v69
                                                        } else {
                                                            var v70 uint32 = v26
                                                            var v71 int16 = 978
                                                            switch -9223372036854775808 {
                                                            case -688:
                                                                var v72 bool = false
                                                                _ = v72
                                                            default:
                                                                var v73 float32 = 1.5
                                                                _ = v73
                                                            }
                                                            _ = v70
                                                            _ = v71
                                                        }
                                                        fmt.Println(v5)
                                                    } else {
                                                        v74 := GenS8{F0: "a"}
                                                        _ = v74
                                                    }
                                                    if false {
                                                        var v75 uint16 = (v3 % ((v3) | 1))
                                                        var v76 uint8 = 255
                                                        v75 -= (v75 * 732)
                                                        _ = v75
                                                        _ = v76
                                                    }
                                                    switch "" {
                                                    case "\n\t":
                                                        if true {
                                                            var v77 uint32 = uint32(v53)
                                                            _ = v77
                                                        }
                                                    case "a":
                                                        var v78 int32 = (-(v5))
                                                        _ = v78
                                                    default:
                                                        var v79 int = (^(v4))
                                                        _ = v79
                                                    }
                                                }
                                                var v80 uint8 = 0
                                                fmt.Println(v4, v7)
                                                _ = v80
                                            } else {
                                                v5 += 2147483647
                                            }
                                            v81 := GenS3{F0: 0}
                                            _ = v53
                                            _ = v81
                                        }
                                        _ = v34
                                    } else {
                                        switch fi21 {
                                        case 42:
                                            fmt.Println(v31)
                                        case 1:
                                            v15 -= (v5 + v15)
                                        }
                                        switch ("hello" + ("a" + (((("hello" + "a") + "a") + "a") + "hello"))) {
                                        case "hello":
                                            if (0 != v25) {
                                                var v82 int8 = -1
                                                _ = len(v6)
                                                _ = v82
                                            }
                                        default:
                                            fmt.Println(v7, v25)
                                        }
                                    }
                                    _ = v31
                                default:
                                    var v83 int16 = int16(v25)
                                    _ = v83
                                }
                                if len(v2) > 0 {
                                    v2[0] = ("\n\t" + ("\n\t" + "foo bar"))
                                }
                                var v84 int8 = int8(v25)
                                v84 = -54
                                v85 := GenS7{F0: (((("foo bar" + "\n\t") + "hello") + ("foo bar" + "hello")) + "foo bar")}
                                _ = v84
                                _ = v85
                            } else {
                                delete(v1, (v26 ^ (v26 << (((v7 << ((uint(v3)) & 63))) & 63))))
                            }
                        }
                        switch fi14 {
                        case -1:
                            fmt.Println(fi14)
                        default:
                            switch (-(fi21)) {
                            case -213:
                                v86 := GenS6{F0: (!(!((!(!true)) && (-9223372036854775808 > 0)))), F1: 0}
                                _ = v86
                            default:
                                v15 -= v15
                            }
                        }
                    }
                }
                _ = v25
                _ = v26
            }
            v87 := GenS2{F0: 2}
            switch -853 {
            case 42:
                delete(v1, 524)
            default:
                var v88 int64 = -924
                _ = v88
            }
            _ = fi21
            _ = v87
        }
        _ = fi14
        _ = v15
    }
    var v89 bool = (v3 < v3)
    var v90 int = v4
    v91 := GenS1{F0: (2147483647 < 0), F1: 670}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v89:", v89)
    fmt.Println("v90:", v90)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v89
    _ = v90
    _ = v91
}
