package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 uint16 }
type GenS2 struct { F0 int64; F1 int64; F2 string }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int; F1 uint }
type GenS5 struct { F0 uint16; F1 bool }
type GenS6 struct { F0 int64; F1 int64; F2 int64 }
type GenS7 struct { F0 uint64; F1 uint16 }
func genFn0(p0 uint, p1 int8) uint32 {
    if true {
        fmt.Println(p0, p1)
        if false {
            v6 := []int{}
            var v7 int8 = (^(p1))
            _ = v6
            _ = v7
        }
    } else {
        var v8 int8 = 127
        for fi9 := 0; fi9 < 19; fi9++ {
            v10 := GenS1{F0: 9223372036854775807, F1: 32767, F2: 2}
            fmt.Println(p1)
            _ = fi9
            _ = v10
        }
        _ = v8
    }
    for fi11 := 0; fi11 < 15; fi11++ {
        p0 = p0
        if (!(4294967295 != 42)) {
            if false {
                p0 -= p0
            } else {
                var v12 int32 = -2147483648
                _ = v12
            }
            if false {
                p0 = p0
            }
            p1 = (p1 | p1)
        } else {
            var v13 uint8 = 2
            for fi14 := 0; fi14 < 17; fi14++ {
                var v15 uint16 = 65535
                _ = fi14
                _ = v15
            }
            v16 := GenS2{F0: 42, F1: 9223372036854775807, F2: "hello"}
            _ = v13
            _ = v16
        }
        var v17 int16 = -1
        _ = fi11
        _ = v17
    }
    if (545 > 2) {
        if (317 <= 32768) {
            if true {
                p0 = 18446744073709551615
                for fi18 := 0; fi18 < 7; fi18++ {
                    p1 = 42
                    _ = fi18
                }
                switch "foo bar" {
                case "hello":
                    var v19 string = ("a" + "")
                    _ = v19
                }
            } else {
                fmt.Println(p1, p0)
                if (float64(0.0) >= (((float64(0.0) * 3.14159) / 1.5) / -2.25)) {
                    var v20 int = -510
                    _ = v20
                }
            }
        }
        fmt.Println(p0)
    }
    switch "a" {
    case "foo bar":
        var v21 uint8 = 42
        _ = v21
    default:
        p0 = (^(p0))
    }
    _ = p0
    _ = p1
    return 4294967295
}

func main() {
    v1 := []int{}
    var v2 int64 = -242
    v3 := make(map[string]int)
    var v4 uint64 = 18446744073709551615
    var v5 uint8 = 127
    switch -2147483648 {
    case 1:
        fmt.Println(v5, v2, v5)
    default:
        switch "foo bar" {
        case "foo bar":
            var v22 int64 = v2
            _ = v22
        case "\n\t":
            v23 := GenS1{F0: v2, F1: 1000, F2: 2}
            _ = v23
        default:
            if (true || (0 == 0)) {
                v5 += (v5 >> ((4294967295) & 63))
            } else {
                switch ("a" + "\n\t") {
                case "hello":
                    v4 += v4
                case "a":
                    v5 -= uint8(v4)
                case "foo bar":
                    v24 := GenS3{F0: v2}
                    _ = v24
                default:
                    v3["\n\t"] = -1
                }
                v25 := GenS3{F0: -1}
                _ = v25
            }
        }
    }
    var v26 int32 = -206
    if false {
        var v27 uint16 = 139
        for fi28 := 0; fi28 < 16; fi28++ {
            fmt.Println(v5, v27)
            if true {
                var v29 float64 = 0.1
                for fi30 := 0; fi30 < 3; fi30++ {
                    v26 -= (v26 ^ v26)
                    switch (fi30 - (^(fi30))) {
                    case 4294967295:
                        v31 := GenS4{F0: 4294967295, F1: 2}
                        _ = v31
                    case 0:
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    default:
                        v27 += (v27 % (((v27 << ((42) & 63))) | 1))
                    }
                    if (0.0 == (v29 + 0.1)) {
                        v2 -= (v2 ^ (v2 + v2))
                    }
                    _ = fi30
                }
                switch (("\n\t" + "hello") + "") {
                case "\n\t":
                    var v32 float32 = 3.14159
                    _ = v32
                case "hello":
                    var v33 int32 = (v26 & v26)
                    _ = v33
                default:
                    switch "\n\t" {
                    case "foo bar":
                        switch "\n\t" {
                        case "a":
                            var v34 int16 = 42
                            _ = v34
                        case "foo bar":
                            if false {
                                if true {
                                    var v35 int16 = int16(v2)
                                    _ = v35
                                }
                                var v36 int = (fi28 - fi28)
                                _ = v36
                            }
                        case "":
                            for fi37 := 0; fi37 < 4; fi37++ {
                                var v38 int64 = (^(v2))
                                v2 = (v2 + 916)
                                if (18446744073709551615 < 4294967295) {
                                    v29 = (1.5 * v29)
                                    var v39 uint16 = (v27 << ((831) & 63))
                                    var v40 float64 = (0.1 - v29)
                                    _ = v39
                                    _ = v40
                                } else {
                                    var v41 int8 = 123
                                    switch (fi28 & 1) {
                                    case 1:
                                        fmt.Println(v26)
                                    }
                                    v42 := GenS4{F0: (fi37 ^ -9223372036854775808), F1: 4294967295}
                                    _ = v41
                                    _ = v42
                                }
                                _ = fi37
                                _ = v38
                            }
                        default:
                            _ = len(v1)
                        }
                    }
                }
                _ = v29
            } else {
                if true {
                    v26 -= v26
                    var v43 int16 = -337
                    _ = v43
                } else {
                    v3["\n\t"] = 0
                    var v44 float64 = (1.5 - 0.1)
                    _ = v44
                }
                for fi45 := 0; fi45 < 14; fi45++ {
                    v46 := GenS2{F0: -312, F1: 1, F2: ("a" + ("foo bar" + "\n\t"))}
                    _ = fi45
                    _ = v46
                }
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            }
            _ = fi28
        }
        var v47 int = 9223372036854775807
        _ = v27
        _ = v47
    } else {
        v4 = v4
        var v48 uint = 750
        _ = v48
    }
    switch "foo bar" {
    case "hello":
        v49 := GenS5{F0: 2, F1: (true && false)}
        _ = v49
    case "a":
        v50 := GenS2{F0: (v2 + -519), F1: 0, F2: ("a" + "")}
        _ = v50
    case "foo bar":
        var v51 string = ("\n\t" + "")
        _ = v51
    default:
        switch -1 {
        case 32768:
            v4 += (v4 * v4)
        }
    }
    v52 := GenS6{F0: 1, F1: (v2 >> ((18446744073709551615) & 63)), F2: (v2 & (^(v2)))}
    v4 = v4
    switch 1 {
    case -9223372036854775808:
        switch ("" + "") {
        case "":
            v4 += (v4 | (v4 * v4))
        case "foo bar":
            var v53 int = 128
            _ = v53
        case "a":
            var v54 int16 = -1
            _ = v54
        default:
            var v55 int64 = -754
            _ = v55
        }
    }
    v4 -= (v4 - v4)
    if false {
        var v56 string = "a"
        _ = v56
    }
    var v57 int8 = 1
    if (!true) {
        if (true && true) {
            v58 := GenS4{F0: int(v5), F1: 0}
            if (true && true) {
                v2 = v2
                var v59 string = "hello"
                for fi60 := 0; fi60 < 1; fi60++ {
                    var v61 int32 = v26
                    fmt.Println(v4, v57)
                    var v62 int = fi60
                    _ = fi60
                    _ = v61
                    _ = v62
                }
                _ = v59
            } else {
                v4 = v4
                if true {
                    var v63 bool = (!true)
                    var v64 uint64 = v4
                    _ = v63
                    _ = v64
                } else {
                    if false {
                        fmt.Println(v2)
                        fmt.Println(v5, v26, v57)
                        if (!true) {
                            for fi65 := 0; fi65 < 20; fi65++ {
                                var v66 int8 = (-(v57))
                                _ = fi65
                                _ = v66
                            }
                            v1 = append(v1, 706)
                        }
                    } else {
                        if (!true) {
                            delete(v3, (("a" + "\n\t") + "foo bar"))
                            var v67 uint32 = genFn0(564, -1)
                            _ = v67
                        } else {
                            v68 := GenS5{F0: 24, F1: (!(((((0.0 - float64(float32(0.0))) / 1e10) + 3.14159) + -2.25) < 0.1))}
                            if (-89 >= (v57 | (v57 & (^(v57))))) {
                                var v69 string = ""
                                _ = v69
                            }
                            var v70 int = -2147483648
                            _ = v68
                            _ = v70
                        }
                        var v71 float32 = float32((float64(-2.25) * (float64(1.5) + (1e10 + -2.25))))
                        if (-1 != -462) {
                            switch 127 {
                            case 9223372036854775807:
                                fmt.Println(v4, v26, v71)
                            case 655:
                                if true {
                                    v72 := GenS2{F0: (v2 & v2), F1: (v2 & (v2 << ((1) & 63))), F2: ""}
                                    _ = v72
                                }
                            }
                            switch "" {
                            case "a":
                                var v73 bool = ((uint32(v5) >= 42) && true)
                                _ = v73
                            case "\n\t":
                                v26 -= v26
                            case "foo bar":
                                if true {
                                    for fi74 := 0; fi74 < 12; fi74++ {
                                        var v75 uint8 = (v5 >> ((940) & 63))
                                        _ = fi74
                                        _ = v75
                                    }
                                    v71 = (v71 * v71)
                                } else {
                                    v76 := v3["foo bar"]
                                    switch int(v26) {
                                    case 0:
                                        v76 = v76
                                    }
                                    if false {
                                        _ = len(v3)
                                        v57 += -1
                                    }
                                    _ = v76
                                }
                            default:
                                var v77 string = "foo bar"
                                _ = v77
                            }
                        } else {
                            var v78 float64 = ((1e10 + float64(1e10)) - (0.1 + (float64(1.5) - 0.1)))
                            _ = v78
                        }
                        _ = v71
                    }
                    v79 := GenS7{F0: v4, F1: 274}
                    if true {
                        fmt.Println(v2, v2, v5)
                        v80 := GenS3{F0: 711}
                        _ = v80
                    } else {
                        var v81 int8 = 127
                        if false {
                            for fi82 := 0; fi82 < 7; fi82++ {
                                v5 = (v5 >> ((546) & 63))
                                v83 := GenS6{F0: (^(v2)), F1: 42, F2: v2}
                                v5 = (v5 - (v5 << ((2) & 63)))
                                _ = fi82
                                _ = v83
                            }
                            var v84 int = 32768
                            _ = v84
                        }
                        v5 += 61
                        _ = v81
                    }
                    _ = v79
                }
            }
            var v85 int = -9223372036854775808
            _ = v58
            _ = v85
        } else {
            var v86 int = 32768
            _ = v86
        }
        if false {
            var v87 int16 = 1
            _ = v87
        } else {
            v5 += (^(v5))
            v5 = (v5 / ((v5) | 1))
            var v88 uint8 = (v5 ^ (v5 >> ((708) & 63)))
            _ = v88
        }
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v26:", v26)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v57:", v57)
    _ = v1
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v57
}
