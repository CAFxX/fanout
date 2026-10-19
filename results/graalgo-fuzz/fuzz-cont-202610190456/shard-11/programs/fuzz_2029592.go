package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int32; F2 int }
type GenS2 struct { F0 uint16; F1 uint16 }
type GenS3 struct { F0 uint16; F1 int; F2 int8 }
type GenS4 struct { F0 uint8; F1 int64; F2 int32 }
type GenS5 struct { F0 int; F1 uint32 }
type GenS6 struct { F0 int; F1 uint8; F2 uint32 }
type GenS7 struct { F0 int32; F1 int32 }
type GenS8 struct { F0 string; F1 int16 }
type GenS9 struct { F0 uint; F1 uint; F2 uint }
type GenS10 struct { F0 int8; F1 uint32; F2 string }
type GenS11 struct { F0 uint32 }
func genFn0(p0 int8, p1 int32, p2 int16) uint64 {
    p2 = p2
    fmt.Println(p1)
    if (false || true) {
        v6 := GenS1{F0: (^(p1)), F1: (p1 % ((p1) | 1)), F2: 1}
        v7 := GenS2{F0: 255, F1: 512}
        _ = v6
        _ = v7
    } else {
        var v8 float32 = 1.5
        var v9 int16 = (p2 / ((p2) | 1))
        _ = v8
        _ = v9
    }
    var v10 uint16 = 42
    v11 := []int{42, 2147483647, -9223372036854775808}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return 18446744073709551615
}
func genFn1(p0 uint32) int32 {
    p0 = (p0 % ((p0) | 1))
    var v12 uint8 = 8
    v13 := []int{4294967295, -2147483648, 0}
    switch ("\n\t" + ("hello" + ("\n\t" + "hello"))) {
    case "hello":
        if (1e10 == 1.5) {
            switch (("" + "") + "hello") {
            case "\n\t":
                switch "" {
                case "\n\t":
                    var v14 int32 = -2147483648
                    _ = v14
                case "":
                    v12 -= v12
                default:
                    switch 2147483647 {
                    case 822:
                        var v15 string = ""
                        _ = v15
                    case 127:
                        if (!(-2147483648 < 2147483647)) {
                            fmt.Println(v12, v12)
                        }
                    default:
                        fmt.Println(v12)
                    }
                }
            default:
                p0 = uint32(v12)
            }
        } else {
            var v16 int8 = -128
            if (!true) {
                v17 := GenS1{F0: 2147483647, F1: -917, F2: -1}
                fmt.Println(v12)
                _ = v17
            }
            _ = v16
        }
    case "\n\t":
        fmt.Println(v12, p0)
    }
    p0 = (p0 >> ((65535) & 63))
    _ = p0
    _ = v12
    _ = v13
    return 526
}

func main() {
    var v1 uint64 = 4294967295
    var v2 string = ("" + (("" + "") + "hello"))
    v3 := make(map[string]string)
    v4 := []string{"foo bar", "hello", "\n\t", (((v2 + ("" + v2)) + v2) + ("\n\t" + v2))}
    var v5 float64 = -2.25
    _ = len(v4)
    switch -2147483648 {
    case 128:
        var v18 int32 = 1
        _ = v18
    case -949:
        v19 := v3[v2]
        _ = v19
    case 2147483647:
        for fi20 := 0; fi20 < 13; fi20++ {
            var v21 uint = 136
            v5 = (((v5 * v5) + v5) + v5)
            if (((0 >= 0) || true) || (255 <= 463)) {
                v21 -= (v21 >> ((931) & 63))
            } else {
                fmt.Println(v2, fi20, v21)
            }
            _ = fi20
            _ = v21
        }
    }
    if (1 >= -2147483648) {
        if len(v4) > 0 {
            v4[0] = (((("hello" + v2) + v2) + (v2 + v2)) + (((("foo bar" + "foo bar") + "") + v2) + "\n\t"))
        }
        v22 := GenS2{F0: 770, F1: 2}
        switch -1 {
        case -637:
            v23 := GenS3{F0: 1, F1: -735, F2: 97}
            _ = v23
        case 127:
            var v24 uint = 42
            _ = v24
        }
        _ = v22
    } else {
        for fi25 := 0; fi25 < 9; fi25++ {
            var v26 int64 = -1
            _ = fi25
            _ = v26
        }
    }
    if true {
        v27 := GenS1{F0: 1, F1: 0, F2: 0}
        _ = v27
    }
    delete(v3, v2)
    if (true && false) {
        if (!(42 < -794)) {
            for fi28 := 0; fi28 < 2; fi28++ {
                if false {
                    v5 = (3.14159 + (v5 - (((1e10 * ((v5 + 3.14159) / ((0.0 + v5) - (v5 - (v5 + 0.1))))) - v5) + 0.1)))
                    if (!true) {
                        v5 -= 1.5
                        v5 += v5
                    } else {
                        v29 := GenS4{F0: 0, F1: -9223372036854775808, F2: 2147483647}
                        if true {
                            v30 := v3[v2]
                            _ = v30
                        }
                        v1 -= 583
                        _ = v29
                    }
                }
                _ = fi28
            }
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            v31 := GenS3{F0: 892, F1: 32768, F2: 8}
            _ = v31
        } else {
            if false {
                var v32 string = v2
                _ = v32
            } else {
                v5 = 0.0
            }
        }
        for fi33 := 0; fi33 < 1; fi33++ {
            if false {
                if ((-2.25 / (((((1.5 - -2.25) / 1.5) - 0.0) + 1e10) - float32((((v5 - 1e10) / -2.25) * -2.25)))) != 0.0) {
                    _ = len(v4)
                    var v34 int8 = 127
                    _ = v34
                } else {
                    if ((776 <= 42) || (0 != 2)) {
                        for fi35 := 0; fi35 < 13; fi35++ {
                            if len(v4) > 0 {
                                v4[0] = "foo bar"
                            }
                            v3["\n\t"] = v2
                            if false {
                                switch ("foo bar" + "a") {
                                case "foo bar":
                                    v36 := GenS2{F0: 48, F1: 32767}
                                    _ = v36
                                case "\n\t":
                                    _ = len(v4)
                                }
                                v1 += (v1 - v1)
                                var v37 float64 = (v5 / (1e10 + v5))
                                _ = v37
                            } else {
                                fmt.Println(v5, fi35)
                                switch "\n\t" {
                                case "foo bar":
                                    v38 := GenS5{F0: fi33, F1: 485}
                                    v39 := GenS1{F0: genFn1(2), F1: -1, F2: (-(fi35))}
                                    _ = v38
                                    _ = v39
                                case "hello":
                                    v5 -= v5
                                case "\n\t":
                                    v1 += (^(v1))
                                default:
                                    var v40 uint8 = 0
                                    _ = v40
                                }
                            }
                            _ = fi35
                        }
                    } else {
                        v41 := GenS3{F0: 65535, F1: (fi33 >> ((38) & 63)), F2: 1}
                        _ = v41
                    }
                    if (344 != 32767) {
                        v42 := GenS6{F0: fi33, F1: 1, F2: 1}
                        if (!true) {
                            var v43 int16 = -1
                            _ = v43
                        } else {
                            v44 := GenS7{F0: 394, F1: 90}
                            var v45 uint = 1
                            _ = v44
                            _ = v45
                        }
                        _ = v42
                    } else {
                        var v46 uint32 = 232
                        _ = v46
                    }
                    var v47 uint8 = 128
                    _ = v47
                }
                v1 = v1
                if (!(true && (-2147483648 != -555))) {
                    var v48 uint8 = 39
                    _ = v48
                }
            } else {
                v4 = append(v4, (((("hello" + "a") + v2) + ((v2 + ("a" + v2)) + v2)) + v2))
                if (false && ((1.5 < 0.1) || false)) {
                    var v49 int16 = 1
                    for fi50 := 0; fi50 < 20; fi50++ {
                        v51 := GenS7{F0: 1, F1: 554}
                        if (-9223372036854775808 > 1) {
                            var v52 int8 = 0
                            var v53 int64 = -1
                            var v54 int = fi33
                            _ = v52
                            _ = v53
                            _ = v54
                        } else {
                            v55 := v3[v2]
                            v56 := GenS8{F0: ("foo bar" + "foo bar"), F1: (-(v49))}
                            _ = v55
                            _ = v56
                        }
                        switch v2 {
                        case "":
                            if ((((1.5 + 1e10) * ((0.0 - 3.14159) + 0.1)) / 3.14159) > float32((1.5 / v5))) {
                                switch "hello" {
                                case "hello":
                                    if (82 < 127) {
                                        var v57 uint = 262
                                        _ = v57
                                    }
                                case "":
                                    if true {
                                        if ((!(true || true)) || (v49 > -32768)) {
                                            var v58 float64 = v5
                                            _ = v58
                                        } else {
                                            var v59 string = (v2 + "foo bar")
                                            _ = v59
                                        }
                                        var v60 float32 = 0.1
                                        _ = v60
                                    }
                                    fmt.Println(v49, v5, fi33)
                                    var v61 uint32 = 694
                                    var v62 int16 = 1
                                    var v63 uint64 = (v1 % ((genFn0(-128, -647, (v62 | v49))) | 1))
                                    _ = v61
                                    _ = v62
                                    _ = v63
                                case "\n\t":
                                    switch "\n\t" {
                                    case "foo bar":
                                        v4 = append(v4, ((v2 + v2) + v2))
                                    case "hello":
                                        if (42 >= 4294967295) {
                                            fmt.Println(v1)
                                            switch (v2 + v2) {
                                            case "hello":
                                                var v64 string = v2
                                                if true {
                                                    v64 = v2
                                                }
                                                _ = v64
                                            default:
                                                v49 = 1000
                                            }
                                            if false {
                                                v65 := v3[""]
                                                _ = v65
                                            }
                                        } else {
                                            fmt.Println(v2, v49)
                                            v1 = (v1 & v1)
                                            v2 = (v2 + v2)
                                        }
                                    case "a":
                                        v2 = v2
                                    default:
                                        var v66 int8 = 100
                                        _ = v66
                                    }
                                default:
                                    var v67 int8 = 0
                                    delete(v3, ("\n\t" + (v2 + ("a" + v2))))
                                    if ("foo bar" == (("" + "\n\t") + v2)) {
                                        fmt.Println(v5, v67, v49)
                                        var v68 float32 = float32(v5)
                                        var v69 float32 = 3.14159
                                        _ = v68
                                        _ = v69
                                    } else {
                                        fmt.Println(fi50, v5)
                                        v70 := GenS1{F0: 2147483647, F1: int32(v67), F2: int(v1)}
                                        switch ("foo bar" + ((v2 + v2) + "\n\t")) {
                                        case "a":
                                            var v71 float32 = float32(((v5 + (v5 - v5)) + 0.0))
                                            _ = v71
                                        case "":
                                            var v72 int16 = (v49 ^ v49)
                                            _ = v72
                                        default:
                                            var v73 uint = 2
                                            _ = v73
                                        }
                                        _ = v70
                                    }
                                    _ = v67
                                }
                                v3[("" + v2)] = "\n\t"
                                var v74 uint64 = (^(v1))
                                _ = v74
                            } else {
                                var v75 uint16 = 2
                                if true {
                                    var v76 uint = uint(v75)
                                    _ = v76
                                }
                                _ = v75
                            }
                        case "a":
                            _ = len(v3)
                        default:
                            var v77 int32 = -1
                            _ = v77
                        }
                        _ = fi50
                        _ = v51
                    }
                    _ = v49
                }
            }
            v78 := GenS9{F0: 2, F1: 42, F2: 4294967295}
            _ = fi33
            _ = v78
        }
    }
    switch -201 {
    case 0:
        v2 = "a"
    case 9223372036854775807:
        if (0 < 127) {
            if (v5 <= 1e10) {
                fmt.Println(v2)
                v79 := GenS10{F0: -87, F1: 933, F2: v2}
                _ = v79
            }
            if false {
                v80 := GenS4{F0: 42, F1: 0, F2: 2147483647}
                for fi81 := 0; fi81 < 1; fi81++ {
                    var v82 int = fi81
                    if true {
                        if ((0 > 1) && true) {
                            switch v82 {
                            case 32767:
                                var v83 int32 = -1
                                _ = v83
                            default:
                                var v84 int8 = -128
                                _ = v84
                            }
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        } else {
                            v5 -= float64(1.5)
                            if true {
                                var v85 int32 = 0
                                _ = v85
                            } else {
                                v86 := GenS4{F0: 127, F1: -924, F2: genFn1(379)}
                                _ = v86
                            }
                            var v87 float64 = ((v5 - v5) - 0.0)
                            _ = v87
                        }
                        v88 := GenS11{F0: 2}
                        _ = v88
                    } else {
                        v1 = genFn0(1, 0, 927)
                        var v89 uint8 = 42
                        _ = v89
                    }
                    if ((fi81 > (v82 / ((42) | 1))) && ((0.1 / 1.0) < (0.0 / -2.25))) {
                        if true {
                            var v90 uint64 = 598
                            _ = v90
                        } else {
                            v91 := GenS5{F0: (fi81 * (^(fi81))), F1: 2}
                            _ = v91
                        }
                        if (!false) {
                            var v92 uint32 = 42
                            if ("hello" > ("foo bar" + v2)) {
                                fmt.Println(v2, v1, fi81)
                                v5 = 0.0
                            } else {
                                var v93 uint16 = 626
                                _ = v93
                            }
                            _ = v92
                        }
                    }
                    _ = fi81
                    _ = v82
                }
                _ = v80
            } else {
                switch "a" {
                case "":
                    v1 += (v1 | 439)
                default:
                    delete(v3, v2)
                }
            }
        } else {
            if true {
                v94 := GenS11{F0: 2147483647}
                v95 := v3[(v2 + (v2 + v2))]
                _ = v94
                _ = v95
            } else {
                switch 9223372036854775807 {
                case -672:
                    for fi96 := 0; fi96 < 3; fi96++ {
                        fmt.Println(v2)
                        v3[(v2 + v2)] = v2
                        _ = len(v3)
                        _ = fi96
                    }
                case -9223372036854775808:
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                default:
                    fmt.Println(v1, v2, v1)
                }
                v2 = ((v2 + "\n\t") + v2)
            }
            var v97 int8 = -1
            v2 = v2
            _ = v97
        }
    }
    v3[("hello" + v2)] = (v2 + ("" + ("foo bar" + "hello")))
    fmt.Println(v5, v1, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
