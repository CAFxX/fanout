package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16; F2 int8 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 string }
type GenS4 struct { F0 uint16; F1 uint16; F2 uint16 }
type GenS5 struct { F0 uint; F1 string }
type GenS6 struct { F0 int64; F1 string }
type GenS7 struct { F0 int64; F1 int8 }
type GenS8 struct { F0 int32; F1 int32; F2 uint64 }
type GenS9 struct { F0 int16; F1 int8 }
type GenS10 struct { F0 string; F1 bool; F2 uint32 }
type GenS11 struct { F0 int16 }

func main() {
    v1 := make(map[string]string)
    var v2 float64 = (0.0 * (float64(3.14159) / (3.14159 / (3.14159 * -2.25))))
    v3 := []int{}
    var v4 bool = (0.1 < (0.0 + (0.1 / 1.0)))
    var v5 int64 = 42
    if (450 != 42) {
        v2 = v2
        if v4 {
            switch 1 {
            case 9223372036854775807:
                var v6 uint64 = 18446744073709551615
                _ = v6
            case 127:
                v5 = v5
            default:
                v7 := GenS1{F0: 1, F1: 42, F2: -1}
                _ = v7
            }
            if ((2 != 0) || (42 < 1)) {
                var v8 int32 = -1
                var v9 int64 = (v5 | (-(v5)))
                for fi10 := 0; fi10 < 7; fi10++ {
                    v5 = v5
                    v11 := GenS2{F0: 127}
                    fmt.Println(v8, v4)
                    _ = fi10
                    _ = v11
                }
                _ = v8
                _ = v9
            }
            v4 = v4
        }
    }
    var v12 uint32 = 535
    if true {
        if false {
            v13 := GenS2{F0: 242}
            for fi14 := 0; fi14 < 8; fi14++ {
                switch 0 {
                case 127:
                    v15 := GenS2{F0: 89}
                    _ = v15
                case 3:
                    fmt.Println(fi14, fi14, v4)
                case -9223372036854775808:
                    var v16 uint16 = 65535
                    _ = v16
                }
                for fi17 := 0; fi17 < 10; fi17++ {
                    switch ("" + "foo bar") {
                    case "hello":
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                    }
                    v5 -= v5
                    _ = len(v3)
                    _ = fi17
                }
                _ = fi14
            }
            v18 := GenS2{F0: 182}
            _ = v13
            _ = v18
        } else {
            switch ("hello" + (("hello" + "hello") + ("foo bar" + ("\n\t" + "hello")))) {
            case "":
                var v19 uint32 = v12
                _ = v19
            default:
                var v20 int16 = 604
                _ = v20
            }
            for fi21 := 0; fi21 < 14; fi21++ {
                if v4 {
                    for fi22 := 0; fi22 < 16; fi22++ {
                        v12 = 65535
                        var v23 uint32 = v12
                        v24 := GenS2{F0: 1}
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                } else {
                    if true {
                        var v25 uint8 = 2
                        _ = v25
                    }
                }
                for fi26 := 0; fi26 < 9; fi26++ {
                    v1[((("hello" + "") + "a") + "a")] = "foo bar"
                    _ = fi26
                }
                if (fi21 > fi21) {
                    fmt.Println(fi21, v2, fi21)
                    switch "\n\t" {
                    case "":
                        var v27 int32 = 42
                        _ = v27
                    case "a":
                        var v28 int16 = -32768
                        _ = v28
                    }
                    for fi29 := 0; fi29 < 10; fi29++ {
                        v4 = v4
                        _ = fi29
                    }
                } else {
                    v4 = v4
                    v5 = 0
                }
                _ = fi21
            }
        }
        v2 += 0.0
    }
    var v30 uint64 = 42
    delete(v1, "")
    if len(v3) > 0 {
        v3[0] = 42
    }
    var v31 uint16 = 524
    v32 := GenS3{F0: "hello"}
    switch "\n\t" {
    case "foo bar":
        v31 += v31
    case "\n\t":
        v31 += v31
    }
    if false {
        switch 4294967295 {
        case 128:
            if true {
                if (!(v4 && v4)) {
                    v33 := v1["foo bar"]
                    for fi34 := 0; fi34 < 17; fi34++ {
                        fmt.Println(v30)
                        fmt.Println(v30, v33, v2)
                        if (float32(float64(float32(3.14159))) != (3.14159 * float32(v2))) {
                            switch v33 {
                            case "foo bar":
                                if ((((!(v4 && (true && false))) || (255 >= 255)) && (128 == 127)) && (v4 || v4)) {
                                    switch ((((v33 + (v33 + "foo bar")) + "a") + v33) + (v33 + "a")) {
                                    case "hello":
                                        if ((v4 || (-128 != -128)) || (v4 && true)) {
                                            for fi35 := 0; fi35 < 15; fi35++ {
                                                if (140 != 0) {
                                                    fmt.Println(v12)
                                                    if v4 {
                                                        switch ((("\n\t" + v33) + ("" + (v33 + v33))) + "foo bar") {
                                                        case "hello":
                                                            if (!v4) {
                                                                fmt.Println(v33, v31, fi35)
                                                                v12 -= (^(v12))
                                                                var v36 float64 = (-2.25 - v2)
                                                                _ = v36
                                                            }
                                                        case "foo bar":
                                                            var v37 int8 = 113
                                                            _ = v37
                                                        }
                                                    } else {
                                                        v38 := GenS4{F0: 781, F1: v31, F2: 302}
                                                        _ = v38
                                                    }
                                                } else {
                                                    fmt.Println(v31, v31)
                                                    v33 = "hello"
                                                    if (42 <= (^(v30))) {
                                                        fmt.Println(v31)
                                                    } else {
                                                        fmt.Println(v5, v4)
                                                        fmt.Println(fi35)
                                                        var v39 bool = v4
                                                        _ = v39
                                                    }
                                                }
                                                v40 := GenS2{F0: 1}
                                                fmt.Println(v2, v4, v31)
                                                _ = len(v3)
                                                _ = fi35
                                                _ = v40
                                            }
                                        } else {
                                            v5 = (v5 / ((v5) | 1))
                                            if v4 {
                                                v4 = (v4 || ((v31 >> ((255) & 63)) != v31))
                                            } else {
                                                var v41 uint = 1
                                                v12 = v12
                                                _ = v41
                                            }
                                        }
                                    case "a":
                                        var v42 uint64 = 18446744073709551615
                                        _ = v42
                                    case "":
                                        if v4 {
                                            _ = len(v1)
                                        }
                                    default:
                                        for fi43 := 0; fi43 < 17; fi43++ {
                                            var v44 float32 = -2.25
                                            if (!v4) {
                                                if (!v4) {
                                                    fmt.Println(v33, v30)
                                                }
                                            }
                                            _ = fi43
                                            _ = v44
                                        }
                                    }
                                    var v45 int32 = -446
                                    for fi46 := 0; fi46 < 6; fi46++ {
                                        if true {
                                            v5 += -9223372036854775808
                                            if (true && false) {
                                                var v47 string = "foo bar"
                                                if ((!(true || v4)) && (-108 >= -39)) {
                                                    v5 = (v5 & 0)
                                                    var v48 uint32 = (v12 << ((2) & 63))
                                                    v49 := GenS3{F0: ""}
                                                    fmt.Println(v2, v48, v5)
                                                    _ = v48
                                                    _ = v49
                                                }
                                                fmt.Println(v33, v31, v31)
                                                _ = v47
                                            } else {
                                                v2 = v2
                                                v50 := GenS5{F0: 1, F1: ((v33 + v33) + "hello")}
                                                _ = v50
                                            }
                                        }
                                        v51 := GenS6{F0: v5, F1: v33}
                                        _ = fi46
                                        _ = v51
                                    }
                                    _ = v45
                                }
                            }
                        }
                        _ = fi34
                    }
                    var v52 string = ("a" + v33)
                    _ = v33
                    _ = v52
                }
                v53 := GenS7{F0: 9223372036854775807, F1: 73}
                if (!(v2 <= -2.25)) {
                    v54 := GenS3{F0: ("" + "a")}
                    for fi55 := 0; fi55 < 2; fi55++ {
                        switch ("" + "a") {
                        case "foo bar":
                            v31 -= v31
                        case "\n\t":
                            for fi56 := 0; fi56 < 7; fi56++ {
                                v31 += v31
                                v57 := GenS8{F0: 2147483647, F1: 2147483647, F2: 0}
                                v58 := GenS5{F0: 0, F1: "a"}
                                _ = fi56
                                _ = v57
                                _ = v58
                            }
                        case "":
                            v30 -= uint64(v31)
                        }
                        _ = fi55
                    }
                    _ = v54
                } else {
                    fmt.Println(v2, v12)
                    var v59 uint64 = (v30 * (v30 / ((v30) | 1)))
                    v60 := GenS1{F0: 42, F1: (v31 >> ((694) & 63)), F2: 1}
                    _ = v59
                    _ = v60
                }
                _ = v53
            } else {
                for fi61 := 0; fi61 < 11; fi61++ {
                    _ = len(v3)
                    _ = fi61
                }
                for fi62 := 0; fi62 < 18; fi62++ {
                    for fi63 := 0; fi63 < 17; fi63++ {
                        v64 := GenS9{F0: 1, F1: 91}
                        _ = fi63
                        _ = v64
                    }
                    _ = fi62
                }
            }
        case 9223372036854775807:
            var v65 int32 = -1
            _ = v65
        default:
            v31 += (v31 >> ((2) & 63))
        }
        if (!v4) {
            var v66 int16 = -575
            var v67 float32 = (1e10 + ((float32(v2) + 3.14159) - 1e10))
            _ = v66
            _ = v67
        } else {
            if v4 {
                var v68 int16 = 622
                v4 = false
                var v69 uint16 = v31
                _ = v68
                _ = v69
            }
            switch "\n\t" {
            case "hello":
                for fi70 := 0; fi70 < 19; fi70++ {
                    if false {
                        if ((true || (true && true)) || (true || ((!v4) && (v4 && v4)))) {
                            if v4 {
                                var v71 uint8 = 127
                                v12 -= uint32(v71)
                                var v72 float64 = v2
                                _ = v71
                                _ = v72
                            }
                        } else {
                            switch ((("foo bar" + "\n\t") + "\n\t") + "") {
                            case "hello":
                                v30 -= (v30 << ((2) & 63))
                            }
                            var v73 uint64 = v30
                            _ = v73
                        }
                        var v74 uint32 = v12
                        var v75 float64 = v2
                        _ = v74
                        _ = v75
                    }
                    if len(v3) > 0 {
                        v3[0] = -901
                    }
                    _ = fi70
                }
            case "foo bar":
                v3 = append(v3, 4294967295)
            case "":
                v2 += 1e10
            }
            for fi76 := 0; fi76 < 13; fi76++ {
                for fi77 := 0; fi77 < 20; fi77++ {
                    v4 = (65535 > 42)
                    _ = fi77
                }
                v4 = v4
                _ = fi76
            }
        }
    } else {
        if v4 {
            v31 += (v31 << ((0) & 63))
        }
        if (v31 <= v31) {
            var v78 int64 = 42
            v79 := GenS4{F0: 221, F1: 65535, F2: (v31 & 65535)}
            _ = v78
            _ = v79
        } else {
            var v80 int32 = -1
            _ = v80
        }
        var v81 uint64 = (v30 - (v30 * (v30 << ((311) & 63))))
        _ = v81
    }
    if (v4 || v4) {
        v82 := GenS6{F0: v5, F1: ""}
        var v83 int8 = 1
        if ("a" < "foo bar") {
            v84 := GenS10{F0: ("a" + ""), F1: (2 != 42), F2: v12}
            v85 := GenS4{F0: v31, F1: (^(v31)), F2: (v31 & 32767)}
            _ = v84
            _ = v85
        }
        _ = v82
        _ = v83
    }
    v86 := GenS11{F0: 32767}
    switch -498 {
    case 356:
        var v87 int32 = -2147483648
        _ = v87
    case 4294967295:
        switch "a" {
        case "hello":
            for fi88 := 0; fi88 < 15; fi88++ {
                if v4 {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                }
                _ = fi88
            }
        default:
            switch -996 {
            case -9223372036854775808:
                var v89 uint8 = 42
                _ = v89
            case 32767:
                var v90 bool = (0 < 2147483647)
                _ = v90
            default:
                switch "hello" {
                case "hello":
                    v91 := GenS9{F0: 32767, F1: -40}
                    _ = v91
                case "\n\t":
                    var v92 uint32 = 554
                    _ = v92
                }
            }
        }
    case 32768:
        v30 -= (v30 - 0)
    default:
        v4 = v4
    }
    _ = len(v3)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v30:", v30)
    fmt.Println("v31:", v31)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v86
}
