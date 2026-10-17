package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 string }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int8 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 bool; F1 int64; F2 uint64 }
type GenS6 struct { F0 uint; F1 uint32; F2 uint64 }
type GenS7 struct { F0 uint8; F1 uint32 }
type GenS8 struct { F0 uint32; F1 int8 }
type GenS9 struct { F0 uint32; F1 int32 }
type GenS10 struct { F0 bool; F1 int16; F2 bool }
type GenS11 struct { F0 int32; F1 uint64 }
type GenS12 struct { F0 uint16; F1 int64; F2 bool }
type GenS13 struct { F0 int16; F1 int64; F2 int64 }

func main() {
    v1 := []int{42, 128, -477, 1}
    var v2 string = (("a" + "a") + ("hello" + ("a" + "\n\t")))
    var v3 uint8 = 255
    v4 := make(map[int]uint32)
    var v5 float64 = (float64((0.0 - 0.1)) / 1.0)
    var v6 int16 = 185
    if false {
        v3 -= 42
    } else {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v7 := GenS1{F0: 255, F1: (v6 & v6), F2: v2}
        _ = v7
    }
    v8 := v4[-2147483648]
    var v9 float32 = 3.14159
    if false {
        v10 := GenS1{F0: 65535, F1: v6, F2: v2}
        if false {
            if ("hello" > "foo bar") {
                var v11 uint8 = (v3 % (((^(v3))) | 1))
                fmt.Println(v6)
                for fi12 := 0; fi12 < 4; fi12++ {
                    if ((v8 + v8) == (v8 - v8)) {
                        var v13 bool = true
                        v14 := GenS1{F0: 906, F1: (-(v6)), F2: ""}
                        var v15 uint = 42
                        _ = v13
                        _ = v14
                        _ = v15
                    } else {
                        fmt.Println(v8, v9)
                    }
                    _ = fi12
                }
                _ = v11
            } else {
                switch -1 {
                case 42:
                    for fi16 := 0; fi16 < 9; fi16++ {
                        for fi17 := 0; fi17 < 9; fi17++ {
                            var v18 int16 = v6
                            if ((fi16 >> ((65535) & 63)) != (-(fi17))) {
                                v19 := GenS2{F0: 155}
                                switch "a" {
                                case "\n\t":
                                    fmt.Println(v6, v3)
                                case "":
                                    v20 := GenS1{F0: 65535, F1: (v6 & v6), F2: v2}
                                    _ = v20
                                }
                                v3 = v3
                                _ = v19
                            } else {
                                v21 := GenS3{F0: -1}
                                _ = v21
                            }
                            fmt.Println(v2, v5, v8)
                            _ = fi17
                            _ = v18
                        }
                        _ = fi16
                    }
                }
                fmt.Println(v2, v3)
                fmt.Println(v5, v6)
            }
            if false {
                fmt.Println(v8, v8, v2)
                fmt.Println(v6, v9)
            }
            v22 := v4[-2147483648]
            _ = v22
        }
        for fi23 := 0; fi23 < 4; fi23++ {
            switch "foo bar" {
            case "\n\t":
                v24 := GenS1{F0: 2, F1: (v6 - v6), F2: (((v2 + ((v2 + v2) + v2)) + v2) + "foo bar")}
                _ = v24
            case "hello":
                switch v2 {
                case "\n\t":
                    delete(v4, int(v6))
                case "":
                    var v25 uint = 4294967295
                    _ = v25
                default:
                    v26 := GenS4{F0: (v8 % (((v8 + (v8 - (^(v8))))) | 1))}
                    _ = v26
                }
            default:
                var v27 uint = 498
                _ = v27
            }
            for fi28 := 0; fi28 < 1; fi28++ {
                switch v2 {
                case "a":
                    var v29 int = 1
                    _ = v29
                case "foo bar":
                    if (0 == -555) {
                        var v30 uint = 145
                        if false {
                            var v31 int64 = 0
                            switch "hello" {
                            case "\n\t":
                                if ((-2147483648 > 42) || false) {
                                    if true {
                                        _ = len(v1)
                                    } else {
                                        var v32 uint16 = 255
                                        _ = v32
                                    }
                                    switch (^(fi28)) {
                                    case 128:
                                        v30 = v30
                                        v33 := v4[1]
                                        var v34 int16 = v6
                                        switch ("\n\t" + v2) {
                                        case "\n\t":
                                            var v35 uint = (v30 ^ v30)
                                            _ = v35
                                        case "foo bar":
                                            if (255 == 1) {
                                                v36 := GenS2{F0: -313}
                                                _ = v36
                                            }
                                        case "hello":
                                            switch "" {
                                            case "hello":
                                                var v37 uint64 = 4294967295
                                                _ = v37
                                            case "foo bar":
                                                delete(v4, fi28)
                                            }
                                        default:
                                            if (!true) {
                                                fmt.Println(v30)
                                            } else {
                                                _ = len(v1)
                                                v31 = v31
                                            }
                                        }
                                        _ = v33
                                        _ = v34
                                    case -2147483648:
                                        v38 := GenS5{F0: false, F1: (v31 / ((v31) | 1)), F2: 0}
                                        _ = v38
                                    case 4294967295:
                                        if true {
                                            v31 = v31
                                            v9 -= float32(v5)
                                        }
                                    default:
                                        var v39 int32 = 331
                                        _ = v39
                                    }
                                } else {
                                    v40 := GenS2{F0: -37}
                                    v41 := GenS6{F0: v30, F1: v8, F2: 4294967295}
                                    v3 += (v3 - v3)
                                    _ = v40
                                    _ = v41
                                }
                            }
                            v3 += (v3 + v3)
                            v42 := GenS5{F0: (uint16(v3) > 32767), F1: v31, F2: 443}
                            var v43 uint16 = 32767
                            if false {
                                v44 := GenS7{F0: v3, F1: (v8 + 2147483647)}
                                var v45 string = "a"
                                _ = v44
                                _ = v45
                            } else {
                                var v46 bool = (fi23 == (-(fi23)))
                                switch "" {
                                case "":
                                    if v46 {
                                        switch 4294967295 {
                                        case 127:
                                            v47 := GenS8{F0: uint32(v30), F1: 0}
                                            _ = v47
                                        case 9223372036854775807:
                                            switch ("\n\t" + "\n\t") {
                                            case "foo bar":
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                            case "":
                                                fmt.Println(v46)
                                            default:
                                                var v48 uint = uint(v3)
                                                _ = v48
                                            }
                                        case 469:
                                            switch 0 {
                                            case -985:
                                                v43 -= uint16(v3)
                                            }
                                        }
                                        var v49 string = "a"
                                        _ = v49
                                    }
                                }
                                var v50 int8 = -128
                                v2 = v2
                                _ = v46
                                _ = v50
                            }
                            if (!((v3 | v3) == (v3 & 42))) {
                                if false {
                                    if (true && (1e10 >= float64((v9 / ((0.1 / float32(((v5 / v5) + float64(-2.25)))) * ((1.5 + -2.25) + v9)))))) {
                                        if (true || (!(!(!((^(v30)) < (v30 % ((uint(v3)) | 1))))))) {
                                            var v51 bool = true
                                            switch (fi28 + fi23) {
                                            case 128:
                                                v51 = v51
                                                v5 -= v5
                                            default:
                                                v52 := GenS9{F0: v8, F1: 0}
                                                if v51 {
                                                    var v53 float64 = v5
                                                    _ = v53
                                                }
                                                _ = v52
                                            }
                                            v9 += 3.14159
                                            _ = v51
                                        }
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                        if len(v1) >= 2 {
                                            v1 = v1[1:len(v1)]
                                        }
                                    } else {
                                        v54 := GenS1{F0: (v30 & v30), F1: v6, F2: ("hello" + v2)}
                                        _ = v54
                                    }
                                    v43 += v43
                                    v43 += v43
                                } else {
                                    var v55 int64 = (v31 - (v31 ^ -1))
                                    _ = v55
                                }
                            } else {
                                var v56 uint32 = (v8 >> ((uint(v3)) & 63))
                                var v57 string = ""
                                var v58 uint64 = 2
                                _ = v56
                                _ = v57
                                _ = v58
                            }
                            _ = v31
                            _ = v42
                            _ = v43
                        } else {
                            switch "a" {
                            case "":
                                v9 -= (v9 + v9)
                            }
                            switch fi28 {
                            case -1:
                                v8 -= v8
                            case 190:
                                fmt.Println(v6, fi28)
                                v8 = v8
                                if (false && true) {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                    if (!true) {
                                        v6 += (^(v6))
                                    }
                                } else {
                                    v59 := GenS5{F0: (563 == 955), F1: -1, F2: 549}
                                    _ = v59
                                }
                            }
                            fmt.Println(v2, v5)
                            if (!(1 < 65)) {
                                v8 += 932
                                v60 := GenS10{F0: (127 > (v3 ^ 9)), F1: v6, F2: false}
                                _ = v60
                            } else {
                                if true {
                                    v3 = (^(v3))
                                    v2 = v2
                                }
                            }
                        }
                        _ = v30
                    }
                }
                v5 = v5
                var v61 int64 = 42
                _ = fi28
                _ = v61
            }
            _ = fi23
        }
        _ = v10
    }
    var v62 string = v2
    fmt.Println(v2, v6)
    v62 = ("a" + v62)
    if true {
        for fi63 := 0; fi63 < 11; fi63++ {
            var v64 string = ("foo bar" + ((v62 + "foo bar") + v62))
            v5 -= v5
            v6 = v6
            _ = fi63
            _ = v64
        }
    }
    for fi65 := 0; fi65 < 17; fi65++ {
        if ((v3 >> ((uint(v3)) & 63)) == (^(v3))) {
            v66 := GenS10{F0: ((fi65 % (((-(fi65))) | 1)) >= fi65), F1: (v6 + v6), F2: (490 < 4)}
            _ = v66
        }
        v3 = (v3 << ((255) & 63))
        for fi67 := 0; fi67 < 10; fi67++ {
            if ((v6 << ((2) & 63)) > v6) {
                v68 := v4[fi67]
                v2 = v2
                var v69 float32 = v9
                switch 2147483647 {
                case 9223372036854775807:
                    v2 = v2
                case 128:
                    v3 -= v3
                case -2147483648:
                    if true {
                        fmt.Println(v5)
                    } else {
                        v69 = 1.5
                        var v70 string = v62
                        _ = v70
                    }
                }
                _ = v68
                _ = v69
            }
            switch fi67 {
            case 4294967295:
                v71 := GenS11{F0: 1, F1: 42}
                _ = v71
            case 1:
                fmt.Println(fi65)
            }
            if (0 == -124) {
                v5 += v5
            } else {
                var v72 uint = 4294967295
                _ = v72
            }
            v73 := GenS8{F0: v8, F1: 127}
            switch (fi67 / ((-9223372036854775808) | 1)) {
            case 424:
                if (!((fi65 <= fi67) && (42 != 0))) {
                    v5 = float64(v9)
                    if true {
                        if false {
                            switch "foo bar" {
                            case "foo bar":
                                if false {
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                } else {
                                    var v74 int16 = v6
                                    v3 = 189
                                    if true {
                                        v6 += 32767
                                        var v75 uint32 = 65535
                                        var v76 bool = true
                                        _ = v75
                                        _ = v76
                                    }
                                    _ = v74
                                }
                            }
                            v8 -= (v8 % ((v8) | 1))
                        } else {
                            v4[0] = (^(v8))
                            v77 := GenS12{F0: 872, F1: -9223372036854775808, F2: true}
                            var v78 float32 = (3.14159 + v9)
                            _ = v77
                            _ = v78
                        }
                        var v79 float32 = (1e10 / ((v9 * 0.1) / (v9 * v9)))
                        if (!(!true)) {
                            var v80 uint32 = v8
                            var v81 uint64 = 2
                            var v82 float32 = ((-2.25 - v9) * float32(3.14159))
                            _ = v80
                            _ = v81
                            _ = v82
                        } else {
                            var v83 int8 = -66
                            v84 := GenS2{F0: 86}
                            _ = v83
                            _ = v84
                        }
                        _ = v79
                    }
                    var v85 uint64 = 18446744073709551615
                    _ = v85
                }
                v86 := GenS1{F0: uint(v3), F1: v6, F2: "a"}
                _ = v86
            case 4294967295:
                v5 = (v5 * 3.14159)
            case -6:
                switch v62 {
                case "\n\t":
                    v6 += v6
                case "foo bar":
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                case "":
                    fmt.Println(v3, v6, v6)
                default:
                    var v87 int64 = -896
                    _ = v87
                }
            }
            _ = fi67
            _ = v73
        }
        _ = fi65
    }
    v6 = v6
    if (true && (true || true)) {
        for fi88 := 0; fi88 < 10; fi88++ {
            if len(v1) > 0 {
                v1[0] = 2147483647
            }
            _ = len(v1)
            _ = fi88
        }
        v89 := GenS13{F0: -725, F1: 9223372036854775807, F2: 297}
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        _ = v89
    }
    switch 568 {
    case -453:
        v90 := GenS9{F0: 702, F1: 42}
        _ = v90
    case 744:
        var v91 int = 179
        _ = v91
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v62:", v62)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v62
    _ = v8
    _ = v9
}
