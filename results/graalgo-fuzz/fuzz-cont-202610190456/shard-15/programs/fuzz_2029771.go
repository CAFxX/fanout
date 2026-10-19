package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32; F1 int64 }
type GenS3 struct { F0 uint64; F1 uint16 }
type GenS4 struct { F0 uint32; F1 int8 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int64; F1 uint32 }
type GenS7 struct { F0 uint64; F1 uint; F2 uint }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 uint; F1 uint64; F2 int64 }
func genFn0(p0 uint) uint16 {
    var v4 uint64 = 18446744073709551615
    v5 := GenS1{F0: 0}
    _ = p0
    _ = v4
    _ = v5
    return 255
}
func genFn1(p0 uint64, p1 uint64) int16 {
    v6 := GenS1{F0: 32767}
    p0 = (p1 | 2)
    _ = p0
    _ = p1
    _ = v6
    return -942
}

func main() {
    var v1 uint32 = 2
    var v2 uint8 = 1
    v3 := []string{("hello" + "a")}
    v7 := GenS1{F0: 42}
    fmt.Println(v1, v2)
    var v8 int64 = -683
    for fi9 := 0; fi9 < 8; fi9++ {
        v2 -= 127
        _ = fi9
    }
    v2 += 127
    v10 := GenS1{F0: 647}
    for fi11 := 0; fi11 < 2; fi11++ {
        var v12 float64 = (3.14159 - (-2.25 / 1.5))
        if true {
            v13 := GenS1{F0: -1}
            if true {
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                if true {
                    fmt.Println(v1)
                    for fi14 := 0; fi14 < 6; fi14++ {
                        _ = len(v3)
                        switch fi11 {
                        case -659:
                            var v15 uint = 706
                            _ = v15
                        case 4294967295:
                            if len(v3) > 0 {
                                v3[0] = "foo bar"
                            }
                        case 127:
                            _ = len(v3)
                        }
                        _ = fi14
                    }
                    if len(v3) > 0 {
                        v3[0] = (("a" + ("hello" + ("hello" + (((("\n\t" + "\n\t") + (("a" + "") + "a")) + "a") + ("\n\t" + (("" + "hello") + "hello")))))) + ((("hello" + ("foo bar" + ("a" + ""))) + "foo bar") + "a"))
                    }
                } else {
                    for fi16 := 0; fi16 < 9; fi16++ {
                        v12 = 0.0
                        var v17 uint16 = genFn0(0)
                        _ = len(v3)
                        _ = fi16
                        _ = v17
                    }
                }
            }
            _ = v13
        }
        _ = fi11
        _ = v12
    }
    v18 := GenS2{F0: (v1 - (v1 * v1)), F1: -1}
    for fi19 := 0; fi19 < 9; fi19++ {
        fmt.Println(v1)
        for fi20 := 0; fi20 < 15; fi20++ {
            v8 += v8
            _ = fi20
        }
        if (((!(255 > 99)) && true) || ("" >= "a")) {
            if true {
                if len(v3) > 0 {
                    v3[0] = ""
                }
                fmt.Println(v2, v8)
                for fi21 := 0; fi21 < 4; fi21++ {
                    v22 := GenS3{F0: 4294967295, F1: 176}
                    if (v8 >= 1) {
                        v8 = v8
                    }
                    v2 -= (v2 % ((1) | 1))
                    v1 += (v1 ^ 0)
                    var v23 float64 = -2.25
                    _ = fi21
                    _ = v22
                    _ = v23
                }
            } else {
                if (1 != 2) {
                    switch (fi19 ^ 578) {
                    case 879:
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    case 4294967295:
                        v2 = (v2 & v2)
                    case 42:
                        _ = len(v3)
                    }
                    v1 -= (v1 + (^(v1)))
                    var v24 float32 = 0.0
                    _ = v24
                } else {
                    switch 128 {
                    case -1:
                        for fi25 := 0; fi25 < 15; fi25++ {
                            v2 = (v2 / ((58) | 1))
                            v26 := GenS4{F0: (v1 + (v1 << ((1) & 63))), F1: -128}
                            _ = fi25
                            _ = v26
                        }
                    case -2147483648:
                        for fi27 := 0; fi27 < 11; fi27++ {
                            v28 := GenS3{F0: 0, F1: 180}
                            v29 := GenS3{F0: 1, F1: 1}
                            switch ("a" + ("" + "hello")) {
                            case "a":
                                v30 := GenS1{F0: genFn1(1, 4294967295)}
                                _ = v30
                            case "hello":
                                v31 := GenS5{F0: 1}
                                _ = v31
                            default:
                                if true {
                                    if false {
                                        v1 += v1
                                        var v32 uint16 = 601
                                        var v33 int8 = 1
                                        fmt.Println(v1, fi27, v33)
                                        if true {
                                            v2 = uint8(v1)
                                        }
                                        _ = v32
                                        _ = v33
                                    }
                                    v34 := GenS1{F0: genFn1(1, 4294967295)}
                                    switch "hello" {
                                    case "":
                                        fmt.Println(fi27)
                                    }
                                    _ = v34
                                } else {
                                    fmt.Println(fi19)
                                    if ((^(v8)) > v8) {
                                        switch "\n\t" {
                                        case "":
                                            var v35 int = fi27
                                            _ = v35
                                        case "hello":
                                            var v36 int64 = (v8 >> ((1) & 63))
                                            _ = v36
                                        default:
                                            var v37 uint32 = (v1 - 4294967295)
                                            _ = v37
                                        }
                                    } else {
                                        if (false && false) {
                                            _ = len(v3)
                                            switch (-(fi27)) {
                                            case 9223372036854775807:
                                                var v38 float32 = 1e10
                                                _ = v38
                                            }
                                            var v39 bool = true
                                            if v39 {
                                                switch fi19 {
                                                case -2147483648:
                                                    var v40 int32 = -2147483648
                                                    _ = v40
                                                default:
                                                    var v41 int = (fi19 ^ (fi27 / ((32768) | 1)))
                                                    if ((((("" + ((("a" + "a") + "hello") + "hello")) + "foo bar") + "foo bar") + "foo bar") < "hello") {
                                                        v1 -= v1
                                                        fmt.Println(v39)
                                                        switch (fi19 % ((fi19) | 1)) {
                                                        case 1:
                                                            v3 = append(v3, "a")
                                                        case 128:
                                                            if (!(!false)) {
                                                                if (v39 || v39) {
                                                                    v42 := GenS4{F0: v1, F1: 42}
                                                                    if (2147483647 > (v1 - (v1 * 99))) {
                                                                        v43 := GenS1{F0: -32768}
                                                                        _ = v43
                                                                    } else {
                                                                        var v44 int16 = 1000
                                                                        v44 -= (^(v44))
                                                                        if len(v3) > 0 {
                                                                            v3[0] = ""
                                                                        }
                                                                        _ = v44
                                                                    }
                                                                    v3 = append(v3, ("foo bar" + "foo bar"))
                                                                    _ = v42
                                                                } else {
                                                                    var v45 int64 = (-(v8))
                                                                    _ = v45
                                                                }
                                                                if (100 > int8(v41)) {
                                                                    if v39 {
                                                                        fmt.Println(v2)
                                                                    }
                                                                    if (!(uint64(v1) <= 0)) {
                                                                        v41 = (fi19 & (^(fi27)))
                                                                        var v46 int64 = 42
                                                                        fmt.Println(v1)
                                                                        _ = v46
                                                                    } else {
                                                                        if (46 <= 16) {
                                                                            var v47 string = ("\n\t" + ("hello" + "a"))
                                                                            var v48 float64 = 1.5
                                                                            _ = v47
                                                                            _ = v48
                                                                        } else {
                                                                            v41 -= fi19
                                                                            if v39 {
                                                                                v2 += 2
                                                                                if v39 {
                                                                                    fmt.Println(v8)
                                                                                    var v49 int = fi27
                                                                                    _ = v49
                                                                                }
                                                                                v50 := GenS6{F0: 0, F1: v1}
                                                                                _ = v50
                                                                            } else {
                                                                                v8 = v8
                                                                                v3 = append(v3, "hello")
                                                                                v1 -= (^(v1))
                                                                            }
                                                                        }
                                                                        v51 := GenS7{F0: 691, F1: 0, F2: 0}
                                                                        var v52 int = fi19
                                                                        v53 := GenS8{F0: 100}
                                                                        _ = v51
                                                                        _ = v52
                                                                        _ = v53
                                                                    }
                                                                    if true {
                                                                        switch (v41 % ((fi27) | 1)) {
                                                                        case -21:
                                                                            v8 = 1
                                                                        case -1:
                                                                            v54 := GenS3{F0: 34, F1: 2}
                                                                            _ = v54
                                                                        default:
                                                                            v41 += v41
                                                                        }
                                                                        switch fi19 {
                                                                        case -726:
                                                                            v8 -= (v8 + (^(v8)))
                                                                        case -2147483648:
                                                                            _ = len(v3)
                                                                        case 128:
                                                                            var v55 uint32 = 4294967295
                                                                            _ = v55
                                                                        }
                                                                        switch ("foo bar" + "hello") {
                                                                        case "\n\t":
                                                                            fmt.Println(fi19, v8)
                                                                        case "":
                                                                            if v39 {
                                                                                var v56 string = "\n\t"
                                                                                var v57 float32 = 0.1
                                                                                _ = v56
                                                                                _ = v57
                                                                            }
                                                                        }
                                                                    }
                                                                } else {
                                                                    if true {
                                                                        if len(v3) > 0 {
                                                                            v3[0] = ("a" + "a")
                                                                        }
                                                                        v3 = append(v3, ("a" + ("foo bar" + ((("\n\t" + "hello") + "foo bar") + "foo bar"))))
                                                                        _ = len(v3)
                                                                    }
                                                                }
                                                            }
                                                        }
                                                    } else {
                                                        var v58 int32 = 2147483647
                                                        var v59 uint = 255
                                                        _ = v58
                                                        _ = v59
                                                    }
                                                    _ = v41
                                                }
                                                v1 = v1
                                            } else {
                                                v2 = (v2 + 127)
                                                var v60 uint64 = 1
                                                _ = v60
                                            }
                                            _ = v39
                                        }
                                    }
                                }
                            }
                            _ = fi27
                            _ = v28
                            _ = v29
                        }
                    case 915:
                        for fi61 := 0; fi61 < 14; fi61++ {
                            var v62 uint16 = genFn0(709)
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            fmt.Println(fi61, fi61, v2)
                            _ = fi61
                            _ = v62
                        }
                    }
                }
                _ = len(v3)
                v1 = (v1 ^ v1)
            }
            switch (((("foo bar" + "foo bar") + "\n\t") + ("foo bar" + "hello")) + ("a" + "foo bar")) {
            case "foo bar":
                for fi63 := 0; fi63 < 2; fi63++ {
                    if (("foo bar" + ((("hello" + "") + "\n\t") + "foo bar")) >= "hello") {
                        var v64 uint16 = genFn0(581)
                        _ = v64
                    }
                    var v65 bool = (-686 == (fi19 << ((255) & 63)))
                    v65 = v65
                    v2 = (v2 - v2)
                    var v66 string = "foo bar"
                    switch 42 {
                    case 32768:
                        if (-589 > -2147483648) {
                            switch (v66 + ("a" + (v66 + v66))) {
                            case "":
                                var v67 uint = 0
                                _ = v67
                            default:
                                var v68 float64 = 3.14159
                                _ = v68
                            }
                            v65 = v65
                            var v69 int16 = -447
                            _ = v69
                        } else {
                            v70 := GenS1{F0: genFn1(1, 1)}
                            if (((^(v1)) >= uint32(v2)) && (!v65)) {
                                v1 -= (v1 & 0)
                                fmt.Println(fi63, v8, v66)
                                var v71 uint8 = 255
                                _ = v71
                            } else {
                                fmt.Println(fi19, v65, v1)
                                v65 = false
                            }
                            _ = v70
                        }
                    default:
                        v1 -= v1
                    }
                    _ = fi63
                    _ = v65
                    _ = v66
                }
            }
            for fi72 := 0; fi72 < 5; fi72++ {
                if false {
                    var v73 string = ("" + "")
                    _ = v73
                } else {
                    var v74 uint = 2
                    v8 += (v8 - (v8 + v8))
                    var v75 uint = (v74 ^ 772)
                    if (65535 >= 32767) {
                        var v76 uint = v75
                        _ = len(v3)
                        _ = v76
                    }
                    var v77 int16 = -441
                    var v78 uint = v74
                    if (v78 >= 662) {
                        v78 += 65535
                    } else {
                        var v79 uint32 = (v1 / ((v1) | 1))
                        var v80 int = (fi72 * fi72)
                        _ = v79
                        _ = v80
                    }
                    _ = v74
                    _ = v75
                    _ = v77
                    _ = v78
                }
                v81 := GenS9{F0: 1}
                _ = len(v3)
                _ = fi72
                _ = v81
            }
        }
        _ = fi19
    }
    var v82 int32 = 1
    switch ("\n\t" + "hello") {
    case "foo bar":
        v2 = (v2 ^ (v2 << ((969) & 63)))
    case "hello":
        switch "foo bar" {
        case "foo bar":
            v1 = 0
        default:
            if (!(!false)) {
                switch 4294967295 {
                case 2147483647:
                    var v83 float32 = (1e10 + 1e10)
                    _ = v83
                default:
                    _ = len(v3)
                }
                var v84 float32 = (-2.25 * (3.14159 + -2.25))
                _ = v84
            }
        }
    case "a":
        v85 := GenS10{F0: 65535, F1: 0, F2: v8}
        _ = v85
    }
    var v86 int = 127
    switch "a" {
    case "\n\t":
        var v87 int64 = v8
        _ = v87
    case "a":
        fmt.Println(v86, v8)
    }
    v8 += (v8 + 42)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v8:", v8)
    fmt.Println("v82:", v82)
    fmt.Println("v86:", v86)
    _ = v1
    _ = v10
    _ = v18
    _ = v2
    _ = v3
    _ = v7
    _ = v8
    _ = v82
    _ = v86
}
