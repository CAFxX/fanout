package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 uint64 }
type GenS2 struct { F0 int32; F1 uint8; F2 uint64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 int16; F1 int }
type GenS5 struct { F0 uint64; F1 bool }
type GenS6 struct { F0 bool; F1 string; F2 bool }
type GenS7 struct { F0 string; F1 uint64 }
type GenS8 struct { F0 int8; F1 uint; F2 int8 }
type GenS9 struct { F0 uint8; F1 uint32 }
type GenS10 struct { F0 int16; F1 string; F2 int16 }

func main() {
    var v1 int64 = 336
    var v2 bool = false
    v3 := make(map[string]bool)
    v4 := []int{-492, -9223372036854775808}
    v5 := GenS1{F0: (283 == 42), F1: 32767, F2: 100}
    v6 := GenS2{F0: 0, F1: 42, F2: 2}
    fmt.Println(v2, v1)
    var v7 float64 = 0.1
    v8 := GenS1{F0: ((!v2) && v2), F1: 0, F2: 4294967295}
    fmt.Println(v7)
    for fi9 := 0; fi9 < 10; fi9++ {
        v3[((("\n\t" + "a") + ((("" + ("foo bar" + "\n\t")) + ("\n\t" + "\n\t")) + "foo bar")) + "a")] = (18446744073709551615 > 2)
        switch "foo bar" {
        case "foo bar":
            v10 := GenS1{F0: (197 < 896), F1: -32768, F2: 4294967295}
            _ = v10
        case "a":
            var v11 uint64 = 696
            _ = v11
        case "hello":
            var v12 int8 = 42
            _ = v12
        }
        switch fi9 {
        case 9223372036854775807:
            for fi13 := 0; fi13 < 6; fi13++ {
                fmt.Println(v7)
                fmt.Println(v7)
                switch 0 {
                case 32768:
                    v1 = (v1 & v1)
                default:
                    if (!v2) {
                        v7 += v7
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                    } else {
                        v14 := GenS3{F0: v1}
                        switch fi9 {
                        case -2147483648:
                            if (42 >= 2) {
                                switch "foo bar" {
                                case "foo bar":
                                    if v2 {
                                        var v15 uint16 = 255
                                        var v16 float64 = v7
                                        _ = v15
                                        _ = v16
                                    } else {
                                        if ((315 < 255) || false) {
                                            var v17 uint32 = 42
                                            var v18 float32 = 3.14159
                                            _ = v17
                                            _ = v18
                                        } else {
                                            v7 = 1.5
                                            if ("a" <= ("" + (("foo bar" + "hello") + ("foo bar" + "a")))) {
                                                v19 := GenS2{F0: -308, F1: 0, F2: 891}
                                                _ = v19
                                            } else {
                                                v20 := GenS4{F0: -152, F1: (fi9 ^ (fi9 % (((fi13 - (fi13 + -450))) | 1)))}
                                                _ = v20
                                            }
                                        }
                                        if (!(!(float32(v7) < float32(3.14159)))) {
                                            if (0 > 18446744073709551615) {
                                                v4 = append(v4, -417)
                                                if (128 != 1) {
                                                    if len(v4) >= 2 {
                                                        v4 = v4[1:len(v4)]
                                                    }
                                                    v3["hello"] = true
                                                } else {
                                                    if ((v1 % (((-(v1))) | 1)) >= 42) {
                                                        var v21 uint8 = 127
                                                        _ = v21
                                                    } else {
                                                        v2 = true
                                                    }
                                                    switch -49 {
                                                    case 32768:
                                                        var v22 uint = 1
                                                        _ = v22
                                                    }
                                                    var v23 int64 = v1
                                                    _ = v23
                                                }
                                                v7 = ((1.5 / (((v7 * 0.0) / v7) / (0.0 - (3.14159 - v7)))) * (v7 * v7))
                                            }
                                            _ = len(v3)
                                            v2 = (103 <= 127)
                                        }
                                    }
                                case "a":
                                    var v24 string = "\n\t"
                                    _ = v24
                                case "":
                                    v25 := GenS5{F0: 2, F1: (4294967295 != 65535)}
                                    _ = v25
                                }
                                v2 = (v2 || (!v2))
                                v26 := GenS6{F0: v2, F1: "\n\t", F2: (!((true || (61 < 127)) || v2))}
                                _ = v26
                            }
                        default:
                            v2 = (v2 && v2)
                        }
                        v7 = -2.25
                        _ = v14
                    }
                    v1 = int64(fi9)
                }
                switch -451 {
                case -542:
                    v27 := v3["foo bar"]
                    _ = v27
                case -9223372036854775808:
                    v1 += int64(fi9)
                case 298:
                    var v28 int8 = 0
                    if v2 {
                        v2 = v2
                        var v29 int64 = v1
                        if (0 > 286) {
                            var v30 float32 = (1.5 / 1.0)
                            _ = v30
                        } else {
                            var v31 uint32 = 1
                            var v32 int8 = v28
                            fmt.Println(v28)
                            if (42 <= -759) {
                                if false {
                                    v2 = v2
                                }
                                fmt.Println(fi13)
                            } else {
                                if (((!true) && v2) && (2147483647 > 1)) {
                                    var v33 int = int(v31)
                                    var v34 int64 = (v1 - v29)
                                    _ = v33
                                    _ = v34
                                }
                                var v35 int64 = (v1 >> ((407) & 63))
                                switch fi13 {
                                case 0:
                                    fmt.Println(fi9, v31)
                                case -142:
                                    var v36 uint = 2
                                    _ = v36
                                default:
                                    v7 += v7
                                }
                                _ = v35
                            }
                            _ = v31
                            _ = v32
                        }
                        fmt.Println(fi9, v2)
                        _ = v29
                    }
                    _ = v28
                }
                v7 += v7
                fmt.Println(fi9, fi13)
                switch fi9 {
                case 32767:
                    var v37 string = "a"
                    _ = v37
                case 2147483647:
                    v38 := GenS3{F0: v1}
                    _ = v38
                case 127:
                    fmt.Println(v7, v1)
                default:
                    v7 += 0.0
                }
                _ = fi13
            }
        default:
            var v39 float64 = v7
            _ = v39
        }
        _ = fi9
    }
    v40 := GenS6{F0: (!v2), F1: "", F2: v2}
    v7 = ((v7 * v7) / (0.1 * v7))
    for fi41 := 0; fi41 < 15; fi41++ {
        var v42 int8 = -128
        fmt.Println(v7)
        _ = fi41
        _ = v42
    }
    fmt.Println(v1, v1)
    _ = len(v4)
    var v43 float32 = 1.5
    var v44 int8 = 1
    switch 938 {
    case -2147483648:
        if len(v4) > 0 {
            v4[0] = -623
        }
    case -9223372036854775808:
        if v2 {
            if v2 {
                v44 += v44
            } else {
                v45 := GenS5{F0: 427, F1: (42 <= 42)}
                _ = v45
            }
        }
    default:
        switch "hello" {
        case "\n\t":
            if v2 {
                v46 := GenS7{F0: (("\n\t" + "foo bar") + "\n\t"), F1: 4294967295}
                v47 := GenS2{F0: -1, F1: 0, F2: 365}
                if ((0 >= 223) && v2) {
                    if ((true || v2) && (v2 || v2)) {
                        if v2 {
                            for fi48 := 0; fi48 < 9; fi48++ {
                                fmt.Println(v44)
                                v49 := v3[("" + "foo bar")]
                                _ = fi48
                                _ = v49
                            }
                            _ = len(v3)
                            switch "a" {
                            case "":
                                var v50 uint16 = 1
                                _ = v50
                            case "hello":
                                if v2 {
                                    v51 := GenS5{F0: 0, F1: (0 <= 255)}
                                    _ = v51
                                }
                            }
                        } else {
                            if v2 {
                                switch 2147483647 {
                                case -9223372036854775808:
                                    switch ("" + "a") {
                                    case "a":
                                        var v52 string = "\n\t"
                                        _ = v52
                                    }
                                case 127:
                                    v4 = append(v4, -616)
                                case -2147483648:
                                    for fi53 := 0; fi53 < 8; fi53++ {
                                        var v54 bool = v2
                                        _ = fi53
                                        _ = v54
                                    }
                                default:
                                    v55 := GenS6{F0: v2, F1: "foo bar", F2: ((!(952 != 0)) || (0 != 41))}
                                    _ = v55
                                }
                            } else {
                                fmt.Println(v7, v44, v1)
                                fmt.Println(v1)
                            }
                            switch (("a" + "\n\t") + "foo bar") {
                            case "foo bar":
                                for fi56 := 0; fi56 < 15; fi56++ {
                                    v57 := v3["foo bar"]
                                    _ = fi56
                                    _ = v57
                                }
                            case "":
                                fmt.Println(v43, v44, v7)
                            }
                            if (!true) {
                                v2 = false
                                if ((v43 * -2.25) < v43) {
                                    v58 := GenS1{F0: v2, F1: 541, F2: 1}
                                    if (false || (2147483647 <= 1)) {
                                        for fi59 := 0; fi59 < 10; fi59++ {
                                            v4 = append(v4, (fi59 ^ (fi59 / ((fi59) | 1))))
                                            if v2 {
                                                var v60 int16 = 32767
                                                v44 = (v44 % ((v44) | 1))
                                                _ = v60
                                            }
                                            _ = fi59
                                        }
                                        switch 128 {
                                        case 602:
                                            v61 := v3[("foo bar" + "hello")]
                                            _ = v61
                                        case 432:
                                            if ((982 >= 4294967295) && (42 <= 4294967295)) {
                                                if (42 > 2147483647) {
                                                    switch ("" + "") {
                                                    case "foo bar":
                                                        if (1 != 43) {
                                                            if len(v4) >= 2 {
                                                                v4 = v4[1:len(v4)]
                                                            }
                                                            for fi62 := 0; fi62 < 12; fi62++ {
                                                                v1 = (v1 * v1)
                                                                _ = fi62
                                                            }
                                                            if (v44 >= 6) {
                                                                var v63 int16 = 42
                                                                var v64 uint8 = 2
                                                                _ = v63
                                                                _ = v64
                                                            } else {
                                                                switch ("hello" + "a") {
                                                                case "a":
                                                                    var v65 float64 = (v7 + 1e10)
                                                                    _ = v65
                                                                case "hello":
                                                                    var v66 int64 = (v1 * 42)
                                                                    _ = v66
                                                                }
                                                                v67 := GenS3{F0: v1}
                                                                var v68 string = "a"
                                                                _ = v67
                                                                _ = v68
                                                            }
                                                        }
                                                    }
                                                } else {
                                                    fmt.Println(v43, v44)
                                                    v43 = 3.14159
                                                }
                                            } else {
                                                switch 922 {
                                                case 0:
                                                    v3[("foo bar" + "foo bar")] = (!(255 > 2))
                                                case 32767:
                                                    if (!v2) {
                                                        for fi69 := 0; fi69 < 17; fi69++ {
                                                            _ = len(v3)
                                                            _ = fi69
                                                        }
                                                        if v2 {
                                                            var v70 int16 = 98
                                                            for fi71 := 0; fi71 < 12; fi71++ {
                                                                v2 = (!(((v43 * v43) - (3.14159 / 0.1)) == 3.14159))
                                                                var v72 int8 = (v44 / (((-(v44))) | 1))
                                                                _ = fi71
                                                                _ = v72
                                                            }
                                                            v3[("\n\t" + "hello")] = true
                                                            _ = v70
                                                        }
                                                    }
                                                case 21:
                                                    v4 = append(v4, 42)
                                                }
                                                if len(v4) > 0 {
                                                    v4[0] = 0
                                                }
                                                fmt.Println(v44, v43)
                                            }
                                        default:
                                            switch 716 {
                                            case 1:
                                                if v2 {
                                                    delete(v3, ("hello" + "hello"))
                                                    v44 = v44
                                                    switch ("\n\t" + "hello") {
                                                    case "\n\t":
                                                        if ((!(false || (!(75 >= 408)))) && (1.5 == v43)) {
                                                            var v73 int32 = -2147483648
                                                            for fi74 := 0; fi74 < 6; fi74++ {
                                                                fmt.Println(v1, fi74)
                                                                for fi75 := 0; fi75 < 11; fi75++ {
                                                                    var v76 bool = v2
                                                                    fmt.Println(fi75, v43)
                                                                    if len(v4) >= 2 {
                                                                        v4 = v4[1:len(v4)]
                                                                    }
                                                                    if v76 {
                                                                        v73 += (v73 % ((-2147483648) | 1))
                                                                    } else {
                                                                        switch "foo bar" {
                                                                        case "\n\t":
                                                                            var v77 int8 = v44
                                                                            _ = v77
                                                                        case "foo bar":
                                                                            var v78 uint8 = 42
                                                                            switch fi74 {
                                                                            case 0:
                                                                                v79 := GenS8{F0: v44, F1: 0, F2: 10}
                                                                                _ = v79
                                                                            case -1:
                                                                                var v80 bool = false
                                                                                _ = v80
                                                                            case 128:
                                                                                var v81 uint16 = 2
                                                                                _ = v81
                                                                            }
                                                                            switch (fi74 * fi75) {
                                                                            case -807:
                                                                                v82 := GenS5{F0: 18446744073709551615, F1: v76}
                                                                                _ = v82
                                                                            default:
                                                                                v78 += v78
                                                                            }
                                                                            _ = v78
                                                                        default:
                                                                            v1 -= -1
                                                                        }
                                                                        switch "foo bar" {
                                                                        case "hello":
                                                                            v44 = (^(v44))
                                                                        }
                                                                    }
                                                                    var v83 uint32 = 2147483648
                                                                    _ = fi75
                                                                    _ = v76
                                                                    _ = v83
                                                                }
                                                                _ = fi74
                                                            }
                                                            delete(v3, "hello")
                                                            _ = v73
                                                        } else {
                                                            if true {
                                                                v1 += v1
                                                            }
                                                            v1 -= v1
                                                        }
                                                    case "":
                                                        var v84 uint32 = 2
                                                        _ = v84
                                                    case "a":
                                                        if v2 {
                                                            v85 := GenS6{F0: (0 == -1), F1: "a", F2: false}
                                                            v1 += v1
                                                            _ = v85
                                                        }
                                                    }
                                                } else {
                                                    delete(v3, "foo bar")
                                                    v86 := GenS9{F0: 33, F1: 424}
                                                    _ = v86
                                                }
                                            }
                                        }
                                    }
                                    var v87 uint64 = 18446744073709551615
                                    _ = v58
                                    _ = v87
                                }
                            }
                        }
                        if (586 > 2147483647) {
                            fmt.Println(v44, v2, v7)
                            for fi88 := 0; fi88 < 20; fi88++ {
                                if ("\n\t" >= "hello") {
                                    v43 = float32(-2.25)
                                } else {
                                    delete(v3, "a")
                                }
                                _ = fi88
                            }
                            for fi89 := 0; fi89 < 18; fi89++ {
                                fmt.Println(v44)
                                v90 := v3["\n\t"]
                                if ((-(v44)) > v44) {
                                    switch fi89 {
                                    case 1:
                                        fmt.Println(v1)
                                    case 42:
                                        _ = len(v3)
                                    }
                                    v91 := GenS10{F0: 78, F1: "foo bar", F2: 42}
                                    v2 = true
                                    _ = v91
                                }
                                _ = fi89
                                _ = v90
                            }
                        } else {
                            for fi92 := 0; fi92 < 14; fi92++ {
                                v44 = (-(v44))
                                if (-1 != -1) {
                                    fmt.Println(v1, v2)
                                    var v93 bool = ((!(!v2)) || (1 >= 42))
                                    _ = v93
                                }
                                _ = fi92
                            }
                        }
                    }
                } else {
                    v44 = v44
                }
                _ = v46
                _ = v47
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v43:", v43)
    fmt.Println("v44:", v44)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v40
    _ = v43
    _ = v44
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
