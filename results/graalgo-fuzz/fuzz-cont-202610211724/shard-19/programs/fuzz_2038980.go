package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8 }
type GenS2 struct { F0 int64; F1 uint; F2 uint64 }
type GenS3 struct { F0 int32; F1 uint64 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int; F1 int8; F2 uint }
type GenS7 struct { F0 uint16; F1 bool }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 int16 }
func genFn0() int {
    var v4 string = "a"
    for fi5 := 0; fi5 < 9; fi5++ {
        var v6 uint64 = 4294967295
        for fi7 := 0; fi7 < 11; fi7++ {
            var v8 int16 = 0
            var v9 string = v4
            switch (^(fi7)) {
            case 32768:
                fmt.Println(v8)
            case -642:
                v10 := []int{127, fi5, fi5, (fi5 ^ fi7)}
                _ = v10
            case -9223372036854775808:
                var v11 uint32 = 2147483647
                _ = v11
            default:
                v12 := GenS1{F0: v6, F1: 127}
                _ = v12
            }
            _ = fi7
            _ = v8
            _ = v9
        }
        _ = fi5
        _ = v6
    }
    var v13 int32 = 273
    if false {
        v13 = v13
    } else {
        v13 -= v13
    }
    _ = v13
    _ = v4
    return -287
}

func main() {
    v1 := []string{"\n\t", "foo bar", "", "a"}
    v2 := []int{177, 2147483647, 885, -633}
    var v3 int = -351
    fmt.Println(v3)
    v2 = append(v2, v3)
    if (false && (0 != 0)) {
        for fi14 := 0; fi14 < 17; fi14++ {
            switch "a" {
            case "\n\t":
                if (v3 != v3) {
                    var v15 uint = 1
                    v3 = (v3 >> (((^(v15))) & 63))
                    var v16 uint16 = 1
                    _ = v15
                    _ = v16
                }
            default:
                if true {
                    switch (("foo bar" + "") + ("hello" + "a")) {
                    case "hello":
                        v3 -= (fi14 << ((255) & 63))
                    case "foo bar":
                        switch 32767 {
                        case -2147483648:
                            if (32767 != 255) {
                                switch "a" {
                                case "a":
                                    fmt.Println(v3, v3)
                                case "":
                                    v17 := GenS1{F0: 2, F1: 127}
                                    _ = v17
                                case "foo bar":
                                    v18 := GenS2{F0: -1, F1: 42, F2: 42}
                                    _ = v18
                                }
                                v19 := GenS2{F0: -30, F1: 746, F2: 42}
                                if (456 < 287) {
                                    v3 += (v3 + (fi14 - (v3 >> ((920) & 63))))
                                }
                                _ = v19
                            } else {
                                if (-24 > 100) {
                                    if false {
                                        var v20 uint = 18446744073709551615
                                        v21 := GenS2{F0: 9223372036854775807, F1: 133, F2: 629}
                                        fmt.Println(fi14, v3)
                                        _ = v20
                                        _ = v21
                                    }
                                    var v22 uint32 = 65535
                                    _ = v22
                                }
                            }
                        default:
                            v3 -= genFn0()
                        }
                    case "a":
                        v1 = append(v1, "")
                    default:
                        v23 := GenS2{F0: -9223372036854775808, F1: 4294967295, F2: 4294967295}
                        _ = v23
                    }
                }
            }
            v3 = (fi14 + 207)
            _ = fi14
        }
    }
    if true {
        fmt.Println(v3)
    }
    v3 = (v3 >> ((434) & 63))
    v3 -= v3
    v24 := GenS1{F0: 2, F1: -128}
    switch genFn0() {
    case 837:
        if false {
            v3 = -2147483648
            if false {
                v25 := GenS1{F0: 4294967295, F1: 35}
                fmt.Println(v3)
                var v26 uint8 = 42
                _ = v25
                _ = v26
            }
        } else {
            v27 := GenS2{F0: 0, F1: 65535, F2: 822}
            var v28 uint = 493
            _ = v27
            _ = v28
        }
    case 515:
        if false {
            fmt.Println(v3)
        }
    case 128:
        v29 := GenS1{F0: 4294967295, F1: 1}
        _ = v29
    }
    v30 := GenS2{F0: 42, F1: 42, F2: 696}
    if (-1 > 42) {
        switch v3 {
        case -162:
            if (!((170 <= 255) && false)) {
                v31 := GenS2{F0: 42, F1: 255, F2: 1}
                _ = v31
            } else {
                var v32 int8 = -128
                for fi33 := 0; fi33 < 4; fi33++ {
                    v1 = append(v1, "")
                    _ = fi33
                }
                _ = v32
            }
        case -799:
            for fi34 := 0; fi34 < 18; fi34++ {
                switch ("" + (((("foo bar" + ((("hello" + "a") + "hello") + "")) + "foo bar") + (((("hello" + (("hello" + "hello") + ("\n\t" + ""))) + "") + "a") + "a")) + ("hello" + ("hello" + "foo bar")))) {
                case "hello":
                    fmt.Println(v3, fi34)
                case "foo bar":
                    if (float64(1e10) != -2.25) {
                        if len(v2) > 0 {
                            v2[0] = genFn0()
                        }
                        var v35 string = "\n\t"
                        _ = v35
                    } else {
                        var v36 int16 = 1
                        var v37 int8 = int8(fi34)
                        _ = v36
                        _ = v37
                    }
                case "":
                    if false {
                        if true {
                            v3 = (v3 & (fi34 * 9223372036854775807))
                            v3 = fi34
                            v3 = (^(v3))
                        } else {
                            for fi38 := 0; fi38 < 17; fi38++ {
                                fmt.Println(v3, fi34, fi38)
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                switch fi38 {
                                case 127:
                                    if true {
                                        v3 -= fi34
                                        v39 := GenS2{F0: 351, F1: 42, F2: 4294967295}
                                        _ = len(v2)
                                        _ = v39
                                    } else {
                                        fmt.Println(fi34, v3, fi34)
                                        v3 -= (fi34 / ((-9223372036854775808) | 1))
                                        var v40 uint8 = 230
                                        _ = v40
                                    }
                                }
                                if true {
                                    var v41 int8 = 1
                                    fmt.Println(v3, fi38)
                                    fmt.Println(fi34, fi38)
                                    _ = v41
                                } else {
                                    fmt.Println(v3, fi38, fi38)
                                    switch "\n\t" {
                                    case "\n\t":
                                        v3 = fi38
                                    case "a":
                                        if true {
                                            v3 = 392
                                            fmt.Println(v3)
                                            v42 := GenS1{F0: 1, F1: -1}
                                            v43 := GenS3{F0: -945, F1: 2}
                                            _ = v42
                                            _ = v43
                                        } else {
                                            if (true && false) {
                                                if true {
                                                    if (-1 != -9223372036854775808) {
                                                        v44 := GenS4{F0: 4294967295}
                                                        if (2 != 2) {
                                                            var v45 uint32 = 0
                                                            if len(v2) > 0 {
                                                                v2[0] = -963
                                                            }
                                                            var v46 int64 = 42
                                                            _ = v45
                                                            _ = v46
                                                        }
                                                        v47 := GenS5{F0: -184}
                                                        fmt.Println(v3)
                                                        _ = v44
                                                        _ = v47
                                                    } else {
                                                        v48 := GenS1{F0: 2, F1: -3}
                                                        _ = v48
                                                    }
                                                } else {
                                                    var v49 string = ("\n\t" + ("hello" + "foo bar"))
                                                    if ((!false) && true) {
                                                        if true {
                                                            v3 -= (v3 / ((-2147483648) | 1))
                                                            if ((v49 + v49) == v49) {
                                                                var v50 uint8 = 189
                                                                _ = v50
                                                            } else {
                                                                var v51 int64 = -133
                                                                v49 = v49
                                                                switch (fi38 * fi34) {
                                                                case -756:
                                                                    var v52 uint8 = 42
                                                                    var v53 float32 = (3.14159 + 1e10)
                                                                    fmt.Println(v49, v53)
                                                                    _ = v52
                                                                    _ = v53
                                                                case -2147483648:
                                                                    var v54 int = (v3 >> ((419) & 63))
                                                                    _ = v54
                                                                default:
                                                                    v55 := GenS4{F0: 0}
                                                                    _ = v55
                                                                }
                                                                var v56 int16 = 42
                                                                v51 = (v51 + v51)
                                                                _ = v51
                                                                _ = v56
                                                            }
                                                        }
                                                        v2 = append(v2, v3)
                                                        v49 = "foo bar"
                                                    }
                                                    _ = v49
                                                }
                                                fmt.Println(fi34, v3, fi34)
                                                var v57 int64 = -1
                                                _ = v57
                                            } else {
                                                v58 := GenS6{F0: 0, F1: 100, F2: 0}
                                                v59 := GenS7{F0: 42, F1: (2 >= 32767)}
                                                _ = v58
                                                _ = v59
                                            }
                                        }
                                        if false {
                                            if false {
                                                var v60 uint64 = 2
                                                var v61 uint64 = (v60 % ((v60) | 1))
                                                var v62 uint = 0
                                                _ = v60
                                                _ = v61
                                                _ = v62
                                            }
                                            v63 := GenS8{F0: -563}
                                            _ = v63
                                        } else {
                                            fmt.Println(fi38, fi38)
                                        }
                                        v3 -= v3
                                    default:
                                        if (true && (112 != 42)) {
                                            var v64 uint32 = 765
                                            if (201 < 4294967295) {
                                                var v65 float64 = 3.14159
                                                var v66 int = v3
                                                _ = v65
                                                _ = v66
                                            } else {
                                                if (617 <= 886) {
                                                    if (127 != 171) {
                                                        v64 = (^(v64))
                                                    }
                                                }
                                            }
                                            v3 -= (fi38 + 92)
                                            _ = v64
                                        } else {
                                            switch (((("foo bar" + ("hello" + ("foo bar" + "foo bar"))) + "\n\t") + "") + ("a" + "a")) {
                                            case "":
                                                v3 = v3
                                                v3 = fi34
                                            default:
                                                var v67 uint = 0
                                                _ = v67
                                            }
                                            var v68 int64 = -1
                                            switch "" {
                                            case "a":
                                                if (v3 == (fi38 % ((v3) | 1))) {
                                                    fmt.Println(v68, fi38)
                                                    v69 := GenS7{F0: 480, F1: (255 != 708)}
                                                    _ = v69
                                                } else {
                                                    var v70 bool = false
                                                    var v71 int32 = 103
                                                    v72 := GenS7{F0: 32767, F1: (v70 && false)}
                                                    if ((!false) && (v70 || (1e10 == float32((-2.25 + 0.1))))) {
                                                        var v73 uint = 224
                                                        _ = v73
                                                    } else {
                                                        v74 := GenS9{F0: -32768}
                                                        var v75 float64 = -2.25
                                                        _ = v74
                                                        _ = v75
                                                    }
                                                    _ = v70
                                                    _ = v71
                                                    _ = v72
                                                }
                                            }
                                            _ = v68
                                        }
                                    }
                                    v3 -= (v3 % ((genFn0()) | 1))
                                    v3 = (v3 + (fi34 & (fi34 & genFn0())))
                                    v3 = (v3 | 0)
                                }
                                switch "" {
                                case "\n\t":
                                    v3 += v3
                                case "a":
                                    _ = len(v2)
                                default:
                                    fmt.Println(v3, fi38, fi34)
                                }
                                _ = fi38
                            }
                            var v76 int64 = 42
                            var v77 uint8 = 255
                            _ = v76
                            _ = v77
                        }
                    }
                default:
                    if len(v1) > 0 {
                        v1[0] = "foo bar"
                    }
                }
                var v78 int8 = 0
                var v79 int32 = 42
                _ = fi34
                _ = v78
                _ = v79
            }
        case 127:
            if len(v1) > 0 {
                v1[0] = "\n\t"
            }
        }
        v3 -= (v3 - 916)
        v80 := GenS6{F0: 0, F1: -126, F2: 962}
        _ = v80
    }
    var v81 int32 = 1
    for fi82 := 0; fi82 < 20; fi82++ {
        fmt.Println(v3, v81, fi82)
        _ = fi82
    }
    switch "a" {
    case "hello":
        if len(v2) > 0 {
            v2[0] = v3
        }
    case "foo bar":
        var v83 uint16 = 65535
        _ = v83
    default:
        v3 = v3
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v81:", v81)
    _ = v1
    _ = v2
    _ = v24
    _ = v3
    _ = v30
    _ = v81
}
