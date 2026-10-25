package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int64 }
type GenS2 struct { F0 int64; F1 int8 }
type GenS3 struct { F0 int64; F1 uint32 }
type GenS4 struct { F0 int16; F1 uint64 }
type GenS5 struct { F0 uint8; F1 int64; F2 int64 }
type GenS6 struct { F0 uint64; F1 uint }
type GenS7 struct { F0 int32 }
type GenS8 struct { F0 int32 }
type GenS9 struct { F0 string; F1 bool }
type GenS10 struct { F0 uint8 }
type GenS11 struct { F0 uint32 }
func genFn0() int16 {
    var v4 int = -2147483648
    var v5 uint = 42
    _ = v4
    _ = v5
    return 0
}
func genFn1(p0 int32) int8 {
    v6 := GenS1{F0: 529, F1: 42}
    for fi7 := 0; fi7 < 6; fi7++ {
        var v8 int16 = 493
        for fi9 := 0; fi9 < 17; fi9++ {
            v10 := []int{(fi9 >> ((2) & 63)), (^(fi7))}
            var v11 uint16 = 255
            var v12 uint16 = (v11 / ((65535) | 1))
            _ = fi9
            _ = v10
            _ = v11
            _ = v12
        }
        _ = fi7
        _ = v8
    }
    p0 = -763
    var v13 uint = 1
    v13 = (v13 % (((v13 ^ v13)) | 1))
    _ = p0
    _ = v13
    _ = v6
    return 38
}

func main() {
    var v1 int = 2147483647
    v2 := make(map[int]uint32)
    var v3 uint16 = 0
    var v14 uint8 = 167
    var v15 uint = 643
    var v16 int32 = 2147483647
    if true {
        for fi17 := 0; fi17 < 5; fi17++ {
            fmt.Println(v1)
            if (-573 > 0) {
                fmt.Println(v3)
                v3 += 822
            }
            switch "hello" {
            case "a":
                if true {
                    v18 := GenS2{F0: int64(v1), F1: 42}
                    var v19 int8 = -48
                    _ = v18
                    _ = v19
                }
            case "\n\t":
                fmt.Println(v1, v16, fi17)
            }
            _ = fi17
        }
        v3 -= v3
    }
    if ((!(!((1 <= -1) && false))) || false) {
        v14 = (v14 + (v14 | uint8(v3)))
    }
    v20 := v2[(v1 / (((v1 * -1)) | 1))]
    if ((100 >= 127) || false) {
        switch "hello" {
        case "hello":
            for fi21 := 0; fi21 < 8; fi21++ {
                v22 := GenS3{F0: -1, F1: (^(v20))}
                var v23 int8 = 100
                if (true || false) {
                    var v24 uint64 = 0
                    v25 := GenS2{F0: 1, F1: -16}
                    _ = v24
                    _ = v25
                } else {
                    var v26 uint8 = 217
                    _ = v26
                }
                _ = fi21
                _ = v22
                _ = v23
            }
        case "a":
            fmt.Println(v3)
        default:
            if true {
                if ("" == "foo bar") {
                    v27 := GenS2{F0: -9223372036854775808, F1: 1}
                    for fi28 := 0; fi28 < 12; fi28++ {
                        if true {
                            if true {
                                fmt.Println(fi28)
                            }
                            if (!((!(false && true)) && ((v16 / (((v16 & (^(v16)))) | 1)) > 637))) {
                                for fi29 := 0; fi29 < 17; fi29++ {
                                    v3 = v3
                                    _ = fi29
                                }
                            }
                        } else {
                            if true {
                                v15 = 42
                                switch "" {
                                case "foo bar":
                                    var v30 bool = (18446744073709551615 >= (v15 ^ uint(v3)))
                                    _ = v30
                                case "a":
                                    v3 -= v3
                                default:
                                    v1 = fi28
                                }
                                v31 := []int{(v1 | fi28), v1}
                                _ = v31
                            } else {
                                for fi32 := 0; fi32 < 18; fi32++ {
                                    var v33 uint8 = (v14 + 42)
                                    v16 += (-(v16))
                                    if true {
                                        switch -127 {
                                        case 127:
                                            var v34 string = ""
                                            _ = v34
                                        default:
                                            switch (^(fi32)) {
                                            case 128:
                                                v1 += fi28
                                            case 42:
                                                if true {
                                                    v35 := GenS4{F0: -1, F1: 672}
                                                    var v36 float32 = (1e10 / (((0.0 + 0.1) + -2.25) / 3.14159))
                                                    _ = v35
                                                    _ = v36
                                                }
                                            }
                                        }
                                    } else {
                                        v37 := GenS2{F0: 1, F1: genFn1(1)}
                                        _ = v37
                                    }
                                    _ = fi32
                                    _ = v33
                                }
                                var v38 string = "hello"
                                delete(v2, 466)
                                _ = v38
                            }
                        }
                        for fi39 := 0; fi39 < 16; fi39++ {
                            v3 = 966
                            if true {
                                v40 := GenS5{F0: 255, F1: -144, F2: -606}
                                _ = v40
                            } else {
                                var v41 bool = true
                                switch (fi28 >> ((42) & 63)) {
                                case -34:
                                    if (v15 < 2) {
                                        var v42 uint16 = v3
                                        v3 += uint16(v20)
                                        v42 = v3
                                        v16 += int32(v20)
                                        v43 := GenS4{F0: 147, F1: 42}
                                        v44 := []int{(-(v1))}
                                        var v45 int32 = v16
                                        _ = v42
                                        _ = v43
                                        _ = v44
                                        _ = v45
                                    } else {
                                        if v41 {
                                            v46 := GenS1{F0: (v3 - 0), F1: -1}
                                            _ = v46
                                        }
                                    }
                                case -9223372036854775808:
                                    _ = len(v2)
                                    if ((v41 || v41) && v41) {
                                        fmt.Println(v20, v1, v14)
                                        if true {
                                            v15 -= v15
                                        } else {
                                            v47 := GenS6{F0: 0, F1: v15}
                                            fmt.Println(v1, fi28)
                                            _ = v47
                                        }
                                        var v48 int32 = v16
                                        _ = v48
                                    } else {
                                        v49 := GenS7{F0: (v16 >> ((v15) & 63))}
                                        fmt.Println(fi28, v20, v1)
                                        if v41 {
                                            var v50 string = "foo bar"
                                            v51 := GenS8{F0: (v16 & v16)}
                                            v52 := GenS6{F0: 42, F1: (v15 - (v15 * 42))}
                                            v20 += (v20 >> ((v15) & 63))
                                            _ = v50
                                            _ = v51
                                            _ = v52
                                        } else {
                                            if (42 >= 1000) {
                                                switch (("a" + "a") + ("" + "hello")) {
                                                case "hello":
                                                    var v53 int8 = 112
                                                    _ = v53
                                                }
                                            } else {
                                                v15 += uint(v20)
                                                switch "foo bar" {
                                                case "foo bar":
                                                    v54 := GenS8{F0: v16}
                                                    _ = v54
                                                case "a":
                                                    var v55 uint16 = 65535
                                                    _ = v55
                                                }
                                            }
                                            var v56 float32 = 3.14159
                                            _ = v56
                                        }
                                        _ = v49
                                    }
                                }
                                v16 = (v16 - (v16 | v16))
                                _ = v41
                            }
                            _ = fi39
                        }
                        for fi57 := 0; fi57 < 9; fi57++ {
                            fmt.Println(v20, fi28, v14)
                            v58 := []int{(fi57 | fi57), (v1 / ((fi28) | 1))}
                            v15 -= (v15 / ((v15) | 1))
                            _ = fi57
                            _ = v58
                        }
                        _ = fi28
                    }
                    _ = v27
                }
                for fi59 := 0; fi59 < 16; fi59++ {
                    v20 -= 4294967295
                    switch ("\n\t" + "hello") {
                    case "":
                        var v60 uint16 = 42
                        _ = v60
                    case "foo bar":
                        v61 := GenS9{F0: ("a" + ("a" + "hello")), F1: (float32(3.14159) == 0.1)}
                        _ = v61
                    case "a":
                        v15 -= 42
                    }
                    _ = fi59
                }
            }
        }
    }
    var v62 uint32 = (v20 / ((uint32(v14)) | 1))
    switch (^(v1)) {
    case 42:
        v63 := GenS6{F0: 264, F1: (v15 - 0)}
        _ = v63
    case 2147483647:
        switch ("foo bar" + ("a" + "hello")) {
        case "foo bar":
            v1 += 42
        default:
            var v64 uint32 = (v20 + (v62 / ((v62) | 1)))
            _ = v64
        }
    }
    fmt.Println(v14, v16)
    v65 := GenS4{F0: -460, F1: 42}
    if false {
        if false {
            v66 := []int{(v1 + v1)}
            _ = v66
        } else {
            v62 = (^(v20))
            v67 := GenS10{F0: (v14 & (^(v14)))}
            _ = v67
        }
        for fi68 := 0; fi68 < 13; fi68++ {
            if (false || ((1 <= 42) || (genFn0() > 1))) {
                for fi69 := 0; fi69 < 6; fi69++ {
                    v16 = 1
                    fmt.Println(v15, v62)
                    v14 = (v14 * v14)
                    _ = fi69
                }
                var v70 int16 = int16(v15)
                if (127 != 0) {
                    v71 := GenS5{F0: 127, F1: -1, F2: -418}
                    _ = v71
                }
                _ = v70
            }
            var v72 int64 = -743
            for fi73 := 0; fi73 < 16; fi73++ {
                v3 = (v3 | 65535)
                var v74 string = "foo bar"
                v75 := GenS11{F0: v62}
                _ = fi73
                _ = v74
                _ = v75
            }
            _ = fi68
            _ = v72
        }
    }
    for fi76 := 0; fi76 < 13; fi76++ {
        v77 := GenS9{F0: "a", F1: false}
        switch fi76 {
        case -9223372036854775808:
            fmt.Println(v15, fi76)
        case 1:
            switch "foo bar" {
            case "hello":
                switch (fi76 & v1) {
                case -497:
                    switch 4294967295 {
                    case 127:
                        v78 := []int{257, -155, -177}
                        _ = v78
                    }
                }
            case "a":
                if false {
                    if (0 >= genFn1(v16)) {
                        v79 := GenS3{F0: -9223372036854775808, F1: (^(v20))}
                        _ = v79
                    }
                    fmt.Println(fi76, v16, v15)
                    var v80 int = (fi76 % ((1) | 1))
                    _ = v80
                }
            default:
                fmt.Println(v1, fi76, v1)
            }
        default:
            for fi81 := 0; fi81 < 18; fi81++ {
                v16 = v16
                _ = fi81
            }
        }
        _ = fi76
        _ = v77
    }
    switch (("hello" + "") + "\n\t") {
    case "a":
        v3 = uint16(v14)
    case "hello":
        if ((((1e10 + (3.14159 / float32(1e10))) / 1e10) + (float32(float64((0.1 / 1.0))) + float32((-2.25 + -2.25)))) <= float32((0.0 + -2.25))) {
            var v82 int = (v1 ^ 128)
            if ((v62 ^ 4294967295) != (v62 & v20)) {
                for fi83 := 0; fi83 < 10; fi83++ {
                    var v84 uint = 2
                    var v85 int8 = -12
                    for fi86 := 0; fi86 < 6; fi86++ {
                        v16 -= v16
                        _ = fi86
                    }
                    _ = fi83
                    _ = v84
                    _ = v85
                }
            }
            fmt.Println(v20, v82)
            _ = v82
        } else {
            fmt.Println(v3, v16, v1)
            v87 := v2[(v1 >> ((v15) & 63))]
            var v88 string = "hello"
            _ = v87
            _ = v88
        }
    default:
        var v89 bool = true
        _ = v89
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v62:", v62)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v20
    _ = v3
    _ = v62
    _ = v65
}
