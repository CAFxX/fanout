package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 uint8 }
type GenS2 struct { F0 bool; F1 uint; F2 uint }
type GenS3 struct { F0 uint; F1 int8 }
type GenS4 struct { F0 int64; F1 uint8; F2 int64 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint16; F1 string; F2 string }
type GenS9 struct { F0 bool }

func main() {
    var v1 int = 1
    v2 := make(map[string]uint32)
    v3 := make(map[uint32]uint32)
    v4 := []string{"a", "\n\t"}
    var v5 bool = false
    for fi6 := 0; fi6 < 19; fi6++ {
        v5 = true
        for fi7 := 0; fi7 < 16; fi7++ {
            fmt.Println(fi6, fi7)
            fmt.Println(fi6)
            switch "" {
            case "":
                var v8 uint16 = 255
                _ = v8
            case "\n\t":
                v5 = v5
            }
            v1 = (v1 / (((fi7 | fi7)) | 1))
            _ = fi7
        }
        var v9 int8 = 42
        _ = fi6
        _ = v9
    }
    var v10 uint8 = 128
    if (87 > 2147483647) {
        for fi11 := 0; fi11 < 2; fi11++ {
            var v12 float64 = -2.25
            v10 += v10
            for fi13 := 0; fi13 < 13; fi13++ {
                var v14 string = "foo bar"
                var v15 float64 = v12
                v5 = v5
                v16 := GenS1{F0: 4294967295, F1: 99, F2: v10}
                if v5 {
                    v10 -= v10
                    if ((v14 >= (v14 + "foo bar")) && ((v1 == (v1 ^ -13)) && v5)) {
                        fmt.Println(v5, v5, fi11)
                        if true {
                            if (v5 || true) {
                                var v17 int = (^(v1))
                                v18 := GenS1{F0: 556, F1: 1, F2: (v10 << ((70) & 63))}
                                _ = v17
                                _ = v18
                            } else {
                                if v5 {
                                    if false {
                                        switch v14 {
                                        case "":
                                            v12 = (3.14159 / (v12 + v12))
                                        case "foo bar":
                                            var v19 uint8 = v10
                                            _ = v19
                                        default:
                                            var v20 bool = (!v5)
                                            _ = v20
                                        }
                                        if len(v4) > 0 {
                                            v4[0] = v14
                                        }
                                    } else {
                                        if true {
                                            v12 = v15
                                            if (v5 && ((0.0 * 0.1) < 0.1)) {
                                                v2[(v14 + v14)] = 0
                                                v10 = (v10 | (v10 + v10))
                                                v21 := GenS2{F0: v5, F1: 4294967295, F2: 2}
                                                _ = len(v3)
                                                v22 := GenS3{F0: 743, F1: 37}
                                                v15 = v12
                                                switch v14 {
                                                case "":
                                                    v5 = (-467 < 42)
                                                case "\n\t":
                                                    v12 = 1e10
                                                default:
                                                    v12 = v12
                                                }
                                                _ = v21
                                                _ = v22
                                            }
                                        }
                                        var v23 bool = false
                                        _ = v23
                                    }
                                    switch v1 {
                                    case 625:
                                        if false {
                                            v24 := GenS2{F0: (2147483648 == 2), F1: 0, F2: 796}
                                            if (1 != 2147483648) {
                                                if len(v4) > 0 {
                                                    v4[0] = v14
                                                }
                                                if len(v4) >= 2 {
                                                    v4 = v4[1:len(v4)]
                                                }
                                                switch (v14 + v14) {
                                                case "hello":
                                                    v25 := GenS3{F0: 350, F1: -48}
                                                    var v26 uint64 = 0
                                                    _ = v25
                                                    _ = v26
                                                case "a":
                                                    var v27 int64 = 9223372036854775807
                                                    _ = v27
                                                }
                                                fmt.Println(v14)
                                            } else {
                                                switch fi11 {
                                                case 42:
                                                    if false {
                                                        fmt.Println(fi13, fi11)
                                                    } else {
                                                        var v28 uint64 = 42
                                                        _ = v28
                                                    }
                                                case 9223372036854775807:
                                                    var v29 int8 = 1
                                                    _ = v29
                                                case 620:
                                                    v15 += v12
                                                }
                                                fmt.Println(v5, v1)
                                                _ = len(v4)
                                                v4 = append(v4, "hello")
                                            }
                                            if (fi11 > fi13) {
                                                var v30 int = fi13
                                                fmt.Println(v10, fi11, fi11)
                                                _ = v30
                                            } else {
                                                v31 := GenS4{F0: 1, F1: (^(v10)), F2: 42}
                                                _ = v31
                                            }
                                            _ = v24
                                        }
                                    case 9223372036854775807:
                                        v32 := GenS4{F0: -1, F1: (v10 | v10), F2: 0}
                                        _ = v32
                                    }
                                }
                                var v33 float64 = v15
                                _ = v33
                            }
                            var v34 uint32 = 2147483648
                            switch fi13 {
                            case -852:
                                switch 32768 {
                                case 128:
                                    v35 := GenS5{F0: 42}
                                    _ = v35
                                }
                            case -2147483648:
                                v1 -= (fi13 | 32767)
                            case -9223372036854775808:
                                var v36 uint = 2
                                _ = v36
                            default:
                                v37 := v2[v14]
                                _ = v37
                            }
                            _ = v34
                        } else {
                            v38 := GenS6{F0: fi11}
                            v14 = v14
                            _ = v38
                        }
                    } else {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        fmt.Println(v5, fi11, v5)
                        if len(v4) > 0 {
                            v4[0] = v14
                        }
                        v10 = (v10 << ((711) & 63))
                        v1 -= (fi13 % (((v1 - fi11)) | 1))
                        v39 := GenS7{F0: 747}
                        var v40 int32 = 2147483647
                        _ = v39
                        _ = v40
                    }
                } else {
                    if v5 {
                        fmt.Println(fi11)
                    }
                    v14 = "a"
                    switch v14 {
                    case "hello":
                        v1 -= (fi13 << ((18446744073709551615) & 63))
                    case "foo bar":
                        v10 -= (v10 | v10)
                    case "\n\t":
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    }
                }
                var v41 uint = 42
                if v5 {
                    if (32767 == 42) {
                        v14 = v14
                    }
                    v5 = false
                } else {
                    v12 = (v12 / float64(0.1))
                }
                fmt.Println(fi13, v10, v14)
                _ = fi13
                _ = v14
                _ = v15
                _ = v16
                _ = v41
            }
            _ = fi11
            _ = v12
        }
    } else {
        v10 += (v10 / (((^(v10))) | 1))
        v10 = v10
        for fi42 := 0; fi42 < 20; fi42++ {
            v43 := GenS1{F0: 0, F1: 85, F2: (v10 + 127)}
            _ = fi42
            _ = v43
        }
    }
    if v5 {
        v5 = v5
        switch 127 {
        case -2147483648:
            for fi44 := 0; fi44 < 15; fi44++ {
                var v45 uint = 770
                fmt.Println(fi44)
                _ = fi44
                _ = v45
            }
        }
        v10 -= (v10 - (^(v10)))
    } else {
        var v46 uint = 18446744073709551615
        v46 -= v46
        _ = v46
    }
    delete(v3, 2147483648)
    v5 = true
    var v47 float64 = 0.0
    for fi48 := 0; fi48 < 5; fi48++ {
        v49 := v3[2147483647]
        _ = fi48
        _ = v49
    }
    var v50 uint16 = 973
    if v5 {
        v51 := GenS4{F0: 344, F1: (v10 | 128), F2: 9223372036854775807}
        var v52 uint16 = 9
        var v53 bool = (v1 == v1)
        _ = v51
        _ = v52
        _ = v53
    }
    if (false || false) {
        var v54 uint = 0
        var v55 float32 = 3.14159
        _ = v54
        _ = v55
    } else {
        for fi56 := 0; fi56 < 3; fi56++ {
            v5 = true
            _ = fi56
        }
        if ((^(v1)) < v1) {
            v57 := GenS8{F0: 0, F1: "a", F2: "a"}
            switch "hello" {
            case "hello":
                var v58 int32 = 246
                _ = v58
            case "foo bar":
                if (18446744073709551615 > uint(v50)) {
                    var v59 float64 = 0.1
                    if v5 {
                        var v60 int16 = -1
                        var v61 int32 = 586
                        var v62 int32 = v61
                        _ = v60
                        _ = v61
                        _ = v62
                    }
                    var v63 string = ("" + ("a" + "a"))
                    _ = v59
                    _ = v63
                }
            default:
                _ = len(v3)
            }
            _ = v57
        } else {
            var v64 uint64 = 18446744073709551615
            if len(v4) > 0 {
                v4[0] = "hello"
            }
            _ = v64
        }
        if ((18446744073709551615 < 1) || false) {
            v1 += 1
            switch ("" + "\n\t") {
            case "":
                if len(v4) >= 2 {
                    v4 = v4[1:len(v4)]
                }
            case "foo bar":
                var v65 uint16 = 1
                _ = v65
            case "a":
                v66 := GenS7{F0: 65535}
                _ = v66
            }
            if false {
                if false {
                    var v67 uint32 = 4294967295
                    _ = v67
                } else {
                    if (v5 && (!true)) {
                        v10 = 128
                    }
                    if v5 {
                        switch (v1 - -9223372036854775808) {
                        case 32768:
                            switch (v1 * v1) {
                            case -9223372036854775808:
                                v1 = v1
                            }
                        default:
                            switch ("foo bar" + ("a" + (("a" + "") + (("a" + ("foo bar" + "a")) + "\n\t")))) {
                            case "a":
                                var v68 uint = 65535
                                _ = v68
                            case "hello":
                                var v69 uint16 = v50
                                _ = v69
                            case "\n\t":
                                if v5 {
                                    if v5 {
                                        v70 := GenS4{F0: -878, F1: v10, F2: -401}
                                        for fi71 := 0; fi71 < 6; fi71++ {
                                            var v72 uint64 = uint64(v50)
                                            _ = fi71
                                            _ = v72
                                        }
                                        _ = v70
                                    }
                                } else {
                                    if ((!(42 < v50)) || v5) {
                                        switch v1 {
                                        case 32768:
                                            v5 = v5
                                        case -9223372036854775808:
                                            if len(v4) > 0 {
                                                v4[0] = (("" + ((("hello" + (("foo bar" + "a") + ("" + "\n\t"))) + "") + "a")) + "a")
                                            }
                                        case 42:
                                            if true {
                                                v5 = v5
                                                v4 = append(v4, "foo bar")
                                                if len(v4) > 0 {
                                                    v4[0] = (((("\n\t" + (("a" + "hello") + "hello")) + "hello") + "\n\t") + "foo bar")
                                                }
                                            }
                                        }
                                    }
                                    for fi73 := 0; fi73 < 8; fi73++ {
                                        switch "\n\t" {
                                        case "foo bar":
                                            if v5 {
                                                delete(v3, 2147483648)
                                            } else {
                                                var v74 uint = 609
                                                _ = v74
                                            }
                                        }
                                        _ = fi73
                                    }
                                }
                            default:
                                if v5 {
                                    var v75 uint16 = v50
                                    _ = v75
                                }
                            }
                        }
                    }
                }
                if (42 != (v10 + v10)) {
                    var v76 int16 = 510
                    switch 0 {
                    case 472:
                        v77 := GenS3{F0: 255, F1: 100}
                        _ = v77
                    case 127:
                        switch (v1 * -175) {
                        case 2147483647:
                            v76 = (^(v76))
                        case -1:
                            fmt.Println(v50)
                        case 1:
                            var v78 int64 = 42
                            _ = v78
                        }
                    }
                    delete(v3, 1)
                    _ = v76
                } else {
                    var v79 uint32 = 2
                    if (v5 || ((v5 && v5) && ((v47 * v47) != (v47 - (float64(3.14159) - v47))))) {
                        switch (v1 % (((v1 + v1)) | 1)) {
                        case -27:
                            var v80 string = "a"
                            _ = v80
                        case 461:
                            if (127 != 0) {
                                for fi81 := 0; fi81 < 5; fi81++ {
                                    switch "" {
                                    case "hello":
                                        var v82 bool = true
                                        _ = v82
                                    }
                                    var v83 int8 = 0
                                    v79 += (v79 - v79)
                                    _ = fi81
                                    _ = v83
                                }
                                if false {
                                    var v84 uint = 18446744073709551615
                                    if (-630 > 42) {
                                        if ((!v5) && v5) {
                                            v79 = v79
                                        }
                                        var v85 int32 = 0
                                        v50 = (v50 + v50)
                                        _ = v85
                                    }
                                    v79 += (v79 + v79)
                                    _ = v84
                                } else {
                                    for fi86 := 0; fi86 < 10; fi86++ {
                                        v50 -= v50
                                        _ = fi86
                                    }
                                    v10 = uint8(v79)
                                }
                            }
                        case 4294967295:
                            v50 = (v50 / ((1) | 1))
                        }
                    } else {
                        if ((float64(0.1) > -2.25) && (42 == 313)) {
                            v87 := GenS9{F0: true}
                            _ = v87
                        }
                        v4 = append(v4, "a")
                    }
                    _ = v79
                }
            } else {
                v1 = v1
                _ = len(v4)
                v47 -= (3.14159 + v47)
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v47:", v47)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v4
    _ = v47
    _ = v5
    _ = v50
}
