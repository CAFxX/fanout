package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 uint16; F1 bool; F2 uint }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int }
type GenS5 struct { F0 int64; F1 uint16 }
type GenS6 struct { F0 int8; F1 int32; F2 int64 }
type GenS7 struct { F0 int32; F1 int8 }
type GenS8 struct { F0 uint32; F1 uint32; F2 uint8 }
type GenS9 struct { F0 int32 }
func genFn0(p0 int, p1 uint16) uint64 {
    p1 = (^(p1))
    fmt.Println(p1)
    v7 := GenS1{F0: 2147483647, F1: 0}
    v8 := GenS1{F0: 4294967295, F1: 0}
    p1 = (p1 - (p1 - (p1 ^ (p1 >> ((65535) & 63)))))
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 60
}

func main() {
    var v1 uint16 = 1
    var v2 int = -1
    v3 := make(map[int]bool)
    var v4 int = v2
    v5 := []string{("\n\t" + "foo bar"), "hello", "\n\t"}
    var v6 uint = 288
    for fi9 := 0; fi9 < 15; fi9++ {
        if (32767 != 1) {
            if (!false) {
                if true {
                    var v10 uint = v6
                    if ((v6 | (v6 >> ((18446744073709551615) & 63))) >= v10) {
                        var v11 int64 = 937
                        switch (fi9 >> (((v10 / (((v6 & v6)) | 1))) & 63)) {
                        case 32767:
                            var v12 int64 = v11
                            _ = v12
                        default:
                            if true {
                                fmt.Println(v4)
                                switch ("foo bar" + ("" + ("a" + ("foo bar" + "foo bar")))) {
                                case "":
                                    _ = len(v5)
                                case "hello":
                                    var v13 int8 = int8(v11)
                                    _ = v13
                                case "foo bar":
                                    v5 = append(v5, ("" + ("hello" + ("a" + ("a" + "hello")))))
                                default:
                                    switch "hello" {
                                    case "hello":
                                        for fi14 := 0; fi14 < 13; fi14++ {
                                            if (255 > 128) {
                                                fmt.Println(fi9, v4)
                                                switch 4294967295 {
                                                case 127:
                                                    var v15 float64 = float64(1e10)
                                                    _ = v15
                                                }
                                                v16 := GenS1{F0: 42, F1: v11}
                                                _ = v16
                                            } else {
                                                var v17 uint32 = 611
                                                _ = v17
                                            }
                                            fmt.Println(v10, fi14, v11)
                                            _ = fi14
                                        }
                                    case "":
                                        v2 = fi9
                                    case "foo bar":
                                        var v18 uint64 = genFn0(v4, (v1 % ((255) | 1)))
                                        _ = v18
                                    default:
                                        var v19 uint32 = 65535
                                        _ = v19
                                    }
                                }
                                for fi20 := 0; fi20 < 1; fi20++ {
                                    var v21 uint16 = v1
                                    _ = fi20
                                    _ = v21
                                }
                            }
                        }
                        v2 = v4
                        _ = v11
                    } else {
                        fmt.Println(v6, fi9, v10)
                        var v22 int64 = 42
                        _ = v22
                    }
                    switch "a" {
                    case "foo bar":
                        v10 = 56
                    default:
                        fmt.Println(v2, v10, v10)
                    }
                    _ = v10
                } else {
                    for fi23 := 0; fi23 < 17; fi23++ {
                        v24 := GenS1{F0: 2, F1: 1}
                        _ = fi23
                        _ = v24
                    }
                    var v25 float32 = float32((float64((0.0 + 0.0)) / float64(3.14159)))
                    var v26 int8 = -1
                    _ = v25
                    _ = v26
                }
                if ((((v6 / ((v6) | 1)) <= (^(v6))) && true) || (false || ((1e10 * 1.5) < (3.14159 + 0.1)))) {
                    if true {
                        v27 := GenS2{F0: (^(v1)), F1: true, F2: 0}
                        if true {
                            fmt.Println(v6, v4, fi9)
                            if false {
                                if true {
                                    if ((!(false && false)) || false) {
                                        var v28 int16 = 1000
                                        for fi29 := 0; fi29 < 8; fi29++ {
                                            switch ("a" + "foo bar") {
                                            case "hello":
                                                _ = len(v5)
                                            case "":
                                                v1 = (v1 - 2)
                                            default:
                                                v2 += (v2 + (v2 % ((fi9) | 1)))
                                            }
                                            v1 += 926
                                            if false {
                                                if false {
                                                    var v30 uint64 = 791
                                                    switch 2147483647 {
                                                    case 144:
                                                        fmt.Println(v6, fi9)
                                                    default:
                                                        v31 := GenS3{F0: 1000}
                                                        _ = v31
                                                    }
                                                    _ = v30
                                                }
                                                if len(v5) > 0 {
                                                    v5[0] = "\n\t"
                                                }
                                                var v32 float32 = 1e10
                                                _ = v32
                                            } else {
                                                v33 := GenS2{F0: v1, F1: true, F2: v6}
                                                _ = v33
                                            }
                                            _ = fi29
                                        }
                                        switch fi9 {
                                        case 32768:
                                            var v34 uint = (v6 - v6)
                                            _ = v34
                                        }
                                        _ = v28
                                    }
                                }
                            }
                        } else {
                            var v35 int8 = 0
                            for fi36 := 0; fi36 < 18; fi36++ {
                                delete(v3, (fi36 - (fi36 / ((v4) | 1))))
                                _ = fi36
                            }
                            _ = v35
                        }
                        v6 -= v6
                        _ = v27
                    } else {
                        v37 := GenS1{F0: 42, F1: -9223372036854775808}
                        for fi38 := 0; fi38 < 8; fi38++ {
                            switch "a" {
                            case "foo bar":
                                v4 = -255
                            case "hello":
                                v4 = v2
                            case "\n\t":
                                var v39 float64 = (3.14159 + 1.5)
                                _ = v39
                            default:
                                v40 := GenS4{F0: v4}
                                _ = v40
                            }
                            v4 = fi38
                            _ = fi38
                        }
                        var v41 uint64 = 42
                        _ = v37
                        _ = v41
                    }
                }
            }
        }
        _ = fi9
    }
    v6 -= v6
    var v42 string = ("\n\t" + "hello")
    v4 = int(v1)
    for fi43 := 0; fi43 < 5; fi43++ {
        for fi44 := 0; fi44 < 4; fi44++ {
            if false {
                v4 = 823
                var v45 bool = true
                if (v45 || v45) {
                    var v46 bool = (!true)
                    if true {
                        var v47 float64 = 0.1
                        delete(v3, -9223372036854775808)
                        _ = v47
                    } else {
                        switch v42 {
                        case "a":
                            v48 := GenS1{F0: 1, F1: 364}
                            _ = v48
                        case "foo bar":
                            var v49 bool = v46
                            _ = v49
                        }
                    }
                    _ = v46
                } else {
                    switch 451 {
                    case -9223372036854775808:
                        var v50 uint8 = 0
                        _ = v50
                    }
                    v51 := GenS1{F0: 85, F1: -719}
                    fmt.Println(v4, v1)
                    _ = v51
                }
                v42 = (v42 + (("\n\t" + (("\n\t" + v42) + v42)) + "hello"))
                _ = v45
            } else {
                var v52 uint8 = 42
                switch "" {
                case "foo bar":
                    if (v42 != "a") {
                        v53 := v3[v2]
                        switch -1 {
                        case 223:
                            var v54 int8 = 1
                            _ = len(v5)
                            _ = v54
                        case -2147483648:
                            fmt.Println(fi43)
                        case 127:
                            if (v53 || (!v53)) {
                                v55 := GenS5{F0: 545, F1: v1}
                                switch (((("hello" + "foo bar") + v42) + (("\n\t" + "\n\t") + "foo bar")) + v42) {
                                case "hello":
                                    var v56 float32 = 0.0
                                    _ = v56
                                case "a":
                                    v52 = (v52 & v52)
                                case "":
                                    if v53 {
                                        var v57 int8 = 0
                                        if len(v5) > 0 {
                                            v5[0] = ((((v42 + (v42 + ("foo bar" + (v42 + (("" + v42) + ""))))) + (v42 + "")) + v42) + v42)
                                        }
                                        _ = v57
                                    } else {
                                        switch "hello" {
                                        case "":
                                            if v53 {
                                                if len(v5) >= 2 {
                                                    v5 = v5[1:len(v5)]
                                                }
                                                var v58 int64 = 686
                                                var v59 bool = (-1 < 0)
                                                _ = v58
                                                _ = v59
                                            } else {
                                                if ((^(v4)) != (^(fi44))) {
                                                    var v60 string = (v42 + "")
                                                    _ = v60
                                                } else {
                                                    v53 = v53
                                                    v53 = (v6 <= 4294967295)
                                                    if v53 {
                                                        if true {
                                                            var v61 int = v2
                                                            if false {
                                                                switch "\n\t" {
                                                                case "a":
                                                                    if false {
                                                                        if v53 {
                                                                            switch (("" + "a") + "a") {
                                                                            case "a":
                                                                                v1 = (v1 - v1)
                                                                            case "hello":
                                                                                if v53 {
                                                                                    fmt.Println(fi43, fi44)
                                                                                }
                                                                            case "foo bar":
                                                                                var v62 float32 = (3.14159 - ((0.1 - ((1.5 - (0.1 * -2.25)) - float32(float64(((-2.25 + 0.1) / -2.25))))) - (1e10 / 3.14159)))
                                                                                _ = v62
                                                                            default:
                                                                                var v63 int32 = 1
                                                                                _ = v63
                                                                            }
                                                                        }
                                                                    }
                                                                case "":
                                                                    var v64 uint = (v6 | 0)
                                                                    fmt.Println(v2)
                                                                    v53 = (2147483647 >= -2147483648)
                                                                    _ = v64
                                                                }
                                                                fmt.Println(v6, v2, v6)
                                                                var v65 uint16 = v1
                                                                _ = v65
                                                            } else {
                                                                if false {
                                                                    switch (v61 + v4) {
                                                                    case 42:
                                                                        switch -9223372036854775808 {
                                                                        case 0:
                                                                            var v66 uint8 = (v52 | 8)
                                                                            _ = v66
                                                                        }
                                                                    }
                                                                }
                                                                v67 := GenS3{F0: 1}
                                                                _ = v67
                                                            }
                                                            var v68 uint8 = v52
                                                            _ = v61
                                                            _ = v68
                                                        } else {
                                                            if true {
                                                                delete(v3, 127)
                                                                switch (v4 * (v4 << ((172) & 63))) {
                                                                case 127:
                                                                    if (true && (((0.0 + -2.25) <= (1.5 + 0.0)) && false)) {
                                                                        switch (fi44 | -117) {
                                                                        case -2147483648:
                                                                            if ((!v53) && (183 >= 9223372036854775807)) {
                                                                                fmt.Println(v1, v42, v2)
                                                                            }
                                                                        case -896:
                                                                            _ = len(v5)
                                                                        case 356:
                                                                            v52 = v52
                                                                        }
                                                                        if (9223372036854775807 >= -1) {
                                                                            v2 += 32768
                                                                            var v69 uint32 = 0
                                                                            _ = v69
                                                                        }
                                                                    }
                                                                case 128:
                                                                    v5 = append(v5, "a")
                                                                default:
                                                                    if (!v53) {
                                                                        var v70 int16 = 32767
                                                                        if ((v1 + (v1 + 42)) < 4) {
                                                                            v2 -= 0
                                                                            v53 = (false || true)
                                                                            v4 = fi43
                                                                        }
                                                                        switch ((v42 + "") + v42) {
                                                                        case "foo bar":
                                                                            v2 = (v4 & (v2 >> ((748) & 63)))
                                                                        case "":
                                                                            fmt.Println(v2, v42, v42)
                                                                        case "a":
                                                                            var v71 uint32 = 65535
                                                                            _ = v71
                                                                        default:
                                                                            var v72 uint64 = 1
                                                                            _ = v72
                                                                        }
                                                                        _ = v70
                                                                    }
                                                                }
                                                            } else {
                                                                fmt.Println(fi43, fi43)
                                                                v42 = (((v42 + (v42 + (v42 + (((v42 + "foo bar") + (v42 + (("\n\t" + (v42 + ((v42 + v42) + (v42 + (v42 + v42))))) + "hello"))) + v42)))) + v42) + (v42 + ""))
                                                            }
                                                            switch v2 {
                                                            case 6:
                                                                var v73 uint64 = 4294967295
                                                                _ = v73
                                                            }
                                                            var v74 int64 = -9223372036854775808
                                                            _ = v74
                                                        }
                                                        v42 = ("\n\t" + "\n\t")
                                                        fmt.Println(v1, fi44)
                                                    } else {
                                                        var v75 int = 4294967295
                                                        switch "" {
                                                        case "\n\t":
                                                            v76 := GenS2{F0: 1, F1: (true && true), F2: (^(v6))}
                                                            var v77 uint32 = 512
                                                            _ = v76
                                                            _ = v77
                                                        case "":
                                                            switch (((("" + ("\n\t" + "foo bar")) + v42) + (((v42 + ("" + (v42 + "a"))) + (((v42 + (v42 + ("hello" + (((v42 + v42) + v42) + "\n\t")))) + (("foo bar" + ("a" + "hello")) + ("foo bar" + "\n\t"))) + (((v42 + "foo bar") + ((("hello" + v42) + v42) + v42)) + "hello"))) + v42)) + (v42 + v42)) {
                                                            case "hello":
                                                                fmt.Println(fi44, fi43, v53)
                                                            case "":
                                                                v78 := v3[v4]
                                                                _ = v78
                                                            case "a":
                                                                if len(v5) > 0 {
                                                                    v5[0] = ""
                                                                }
                                                            }
                                                        default:
                                                            if v53 {
                                                                if v53 {
                                                                    var v79 uint8 = v52
                                                                    var v80 float64 = float64(((0.0 * 0.0) + (((-2.25 / 3.14159) * 0.0) * (0.1 * 1e10))))
                                                                    _ = v79
                                                                    _ = v80
                                                                } else {
                                                                    var v81 int32 = -140
                                                                    v75 = v2
                                                                    _ = v81
                                                                }
                                                            } else {
                                                                switch fi43 {
                                                                case 9223372036854775807:
                                                                    switch "a" {
                                                                    case "hello":
                                                                        v82 := GenS5{F0: 0, F1: 585}
                                                                        _ = v82
                                                                    case "":
                                                                        v83 := GenS5{F0: 42, F1: 152}
                                                                        _ = v83
                                                                    default:
                                                                        fmt.Println(v6, v53, v4)
                                                                    }
                                                                case 961:
                                                                    v2 += fi44
                                                                }
                                                                v84 := GenS5{F0: 9223372036854775807, F1: v1}
                                                                _ = v84
                                                            }
                                                        }
                                                        var v85 bool = (v53 || (!true))
                                                        v6 -= (v6 - 18446744073709551615)
                                                        var v86 int16 = -1
                                                        _ = v75
                                                        _ = v85
                                                        _ = v86
                                                    }
                                                }
                                            }
                                        default:
                                            if (!(!((2147483647 >= -2147483648) || (!v53)))) {
                                                v87 := GenS6{F0: 106, F1: -861, F2: -1}
                                                _ = v87
                                            }
                                        }
                                        v6 -= v6
                                    }
                                default:
                                    v4 = 263
                                }
                                _ = v55
                            } else {
                                var v88 int64 = int64(v2)
                                var v89 int32 = 229
                                var v90 uint32 = 4294967295
                                _ = v88
                                _ = v89
                                _ = v90
                            }
                            v53 = (v6 > (v6 * (v6 >> ((821) & 63))))
                        default:
                            v91 := GenS7{F0: 274, F1: 100}
                            _ = v91
                        }
                        _ = v53
                    } else {
                        v92 := GenS8{F0: 841, F1: 0, F2: (v52 % ((v52) | 1))}
                        fmt.Println(v52, v52, fi43)
                        _ = v92
                    }
                case "":
                    var v93 int = int(v1)
                    _ = v93
                case "a":
                    v4 = (v4 + int(v6))
                }
                v94 := GenS9{F0: 906}
                _ = v52
                _ = v94
            }
            v6 = v6
            var v95 uint = (v6 >> ((42) & 63))
            _ = fi44
            _ = v95
        }
        if (v1 > (v1 - 1)) {
            var v96 uint32 = 0
            _ = v96
        }
        var v97 string = v42
        _ = fi43
        _ = v97
    }
    var v98 uint32 = 2147483648
    var v99 bool = false
    _ = len(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v98:", v98)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v6
    _ = v98
    _ = v99
}
