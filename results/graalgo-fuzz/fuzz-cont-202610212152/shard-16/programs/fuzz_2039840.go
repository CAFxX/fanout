package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int32; F2 string }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint32; F1 uint64 }
type GenS4 struct { F0 bool; F1 string }
type GenS5 struct { F0 uint8; F1 string; F2 int64 }
type GenS6 struct { F0 uint8; F1 uint8 }
type GenS7 struct { F0 uint; F1 int64; F2 int64 }
type GenS8 struct { F0 uint64; F1 string; F2 int8 }
type GenS9 struct { F0 uint8; F1 uint }
type GenS10 struct { F0 string }
type GenS11 struct { F0 bool; F1 int; F2 uint16 }
type GenS12 struct { F0 uint; F1 uint32; F2 int }
type GenS13 struct { F0 uint32 }
type GenS14 struct { F0 uint64; F1 int64; F2 bool }

func main() {
    v1 := []int{4294967295, 209}
    var v2 int8 = -1
    v3 := []uint32{2147483647, 2, 2147483647, 65535}
    var v4 int16 = -667
    var v5 uint = 4294967295
    v6 := []string{"\n\t"}
    fmt.Println(v2, v2, v4)
    if (32767 < 9223372036854775807) {
        v5 = 2
    }
    if (true && (0 < -9223372036854775808)) {
        v4 += (v4 / ((v4) | 1))
        switch "a" {
        case "hello":
            switch -20 {
            case -9223372036854775808:
                fmt.Println(v5, v5)
            case 9223372036854775807:
                v2 = v2
            default:
                v2 += (v2 * 0)
            }
        }
    }
    v7 := GenS1{F0: 2, F1: -28, F2: ""}
    switch 42 {
    case 4294967295:
        v4 -= v4
    case -621:
        if (0 <= uint64(v5)) {
            switch int(v2) {
            case -939:
                if (true && false) {
                    v5 += 1
                    var v8 int = 9223372036854775807
                    _ = v8
                } else {
                    for fi9 := 0; fi9 < 20; fi9++ {
                        var v10 float32 = (1.5 - 1.5)
                        v11 := GenS2{F0: (("a" + "hello") + "a")}
                        v5 = 18446744073709551615
                        _ = fi9
                        _ = v10
                        _ = v11
                    }
                    for fi12 := 0; fi12 < 12; fi12++ {
                        if (v5 >= v5) {
                            if true {
                                for fi13 := 0; fi13 < 1; fi13++ {
                                    if len(v1) > 0 {
                                        v1[0] = 32768
                                    }
                                    var v14 float32 = 1e10
                                    _ = fi13
                                    _ = v14
                                }
                            }
                            if (2147483647 < 4294967295) {
                                var v15 int32 = -1
                                if true {
                                    v4 = (-(v4))
                                    if true {
                                        if len(v6) > 0 {
                                            v6[0] = "\n\t"
                                        }
                                        fmt.Println(fi12)
                                    } else {
                                        v1 = append(v1, (^(fi12)))
                                        for fi16 := 0; fi16 < 13; fi16++ {
                                            v5 += v5
                                            _ = fi16
                                        }
                                    }
                                } else {
                                    switch "a" {
                                    case "hello":
                                        for fi17 := 0; fi17 < 4; fi17++ {
                                            if ("hello" <= "\n\t") {
                                                var v18 float64 = (0.0 + ((0.1 * (1.5 / float64((-2.25 + 1.5)))) - 0.0))
                                                v5 += v5
                                                _ = v18
                                            } else {
                                                if true {
                                                    if false {
                                                        var v19 uint32 = 756
                                                        _ = v19
                                                    }
                                                    if false {
                                                        if (0.0 < ((-2.25 + (1e10 / 3.14159)) - 3.14159)) {
                                                            v4 = v4
                                                            v4 += v4
                                                        } else {
                                                            v2 -= (v2 - 100)
                                                            v20 := GenS2{F0: "hello"}
                                                            _ = v20
                                                        }
                                                        if true {
                                                            v3 = append(v3, 42)
                                                            fmt.Println(v4, v2)
                                                        }
                                                        if (v15 == (v15 - (v15 & 42))) {
                                                            v4 -= -32768
                                                        } else {
                                                            if (true && (!(9223372036854775807 == fi17))) {
                                                                switch (fi12 % ((fi17) | 1)) {
                                                                case 9223372036854775807:
                                                                    var v21 uint64 = 4294967295
                                                                    _ = v21
                                                                case 1:
                                                                    v22 := GenS3{F0: 4294967295, F1: 42}
                                                                    _ = v22
                                                                case -2147483648:
                                                                    v23 := GenS4{F0: (42 > 1), F1: ""}
                                                                    _ = v23
                                                                }
                                                                v2 = (v2 % ((v2) | 1))
                                                                v24 := GenS1{F0: 255, F1: (v15 & 1), F2: "foo bar"}
                                                                if (825 > 729) {
                                                                    fmt.Println(v4, v15, v2)
                                                                } else {
                                                                    v5 = 65535
                                                                    if (false || true) {
                                                                        v4 = v4
                                                                        if ((v4 % (((^(v4))) | 1)) > v4) {
                                                                            switch "" {
                                                                            case "a":
                                                                                if (65535 == 65535) {
                                                                                    var v25 int8 = -1
                                                                                    switch fi12 {
                                                                                    case -9223372036854775808:
                                                                                        v25 += (^(v2))
                                                                                    case 236:
                                                                                        v6 = append(v6, "hello")
                                                                                    }
                                                                                    _ = v25
                                                                                }
                                                                            default:
                                                                                fmt.Println(v5)
                                                                            }
                                                                            if len(v1) > 0 {
                                                                                v1[0] = 9223372036854775807
                                                                            }
                                                                        }
                                                                    } else {
                                                                        if len(v1) > 0 {
                                                                            v1[0] = (^(fi17))
                                                                        }
                                                                        var v26 float64 = 3.14159
                                                                        _ = len(v3)
                                                                        var v27 bool = (fi17 < (fi17 - fi17))
                                                                        _ = v26
                                                                        _ = v27
                                                                    }
                                                                }
                                                                v5 = (v5 * v5)
                                                                _ = v24
                                                            } else {
                                                                switch fi12 {
                                                                case -754:
                                                                    fmt.Println(v5, v2)
                                                                case 720:
                                                                    switch ("\n\t" + "\n\t") {
                                                                    case "\n\t":
                                                                        var v28 uint16 = 65535
                                                                        _ = v28
                                                                    case "":
                                                                        if len(v6) > 0 {
                                                                            v6[0] = "hello"
                                                                        }
                                                                    default:
                                                                        if (0 <= (^(v15))) {
                                                                            if false {
                                                                                v4 -= v4
                                                                            } else {
                                                                                v4 -= (v4 | v4)
                                                                            }
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                            if true {
                                                                var v29 uint8 = 42
                                                                _ = v29
                                                            }
                                                            v1 = append(v1, -9223372036854775808)
                                                        }
                                                        v30 := GenS5{F0: 64, F1: "foo bar", F2: -9223372036854775808}
                                                        v2 += 42
                                                        _ = v30
                                                    }
                                                }
                                                var v31 uint16 = 1
                                                v32 := GenS3{F0: 65535, F1: 0}
                                                _ = v31
                                                _ = v32
                                            }
                                            var v33 uint = (^(v5))
                                            if len(v3) >= 2 {
                                                v3 = v3[1:len(v3)]
                                            }
                                            _ = fi17
                                            _ = v33
                                        }
                                    }
                                }
                                var v34 uint = v5
                                _ = v15
                                _ = v34
                            }
                            fmt.Println(v5, v2, v2)
                        } else {
                            fmt.Println(v5, v4)
                            if (538 > 65535) {
                                var v35 string = ("a" + "\n\t")
                                for fi36 := 0; fi36 < 1; fi36++ {
                                    switch (fi36 & (fi12 / ((-2147483648) | 1))) {
                                    case 1:
                                        var v37 string = "\n\t"
                                        _ = v37
                                    }
                                    fmt.Println(v5, v35)
                                    _ = fi36
                                }
                                if false {
                                    if true {
                                        v38 := GenS4{F0: true, F1: (v35 + "")}
                                        _ = v38
                                    } else {
                                        var v39 int64 = -9223372036854775808
                                        if false {
                                            var v40 int64 = 1
                                            var v41 int = (fi12 >> ((v5) & 63))
                                            _ = v40
                                            _ = v41
                                        } else {
                                            v42 := GenS6{F0: 0, F1: 73}
                                            if (!(false && false)) {
                                                if len(v6) >= 2 {
                                                    v6 = v6[1:len(v6)]
                                                }
                                                var v43 uint16 = 0
                                                _ = v43
                                            }
                                            _ = v42
                                        }
                                        fmt.Println(v5, fi12, v35)
                                        _ = v39
                                    }
                                    v6 = append(v6, ((v35 + v35) + "hello"))
                                }
                                _ = v35
                            } else {
                                switch (fi12 - fi12) {
                                case -2147483648:
                                    v2 -= v2
                                case 2147483647:
                                    var v44 bool = false
                                    _ = v44
                                }
                            }
                        }
                        _ = fi12
                    }
                }
            case 385:
                v45 := GenS7{F0: 1, F1: 9223372036854775807, F2: 0}
                _ = v45
            }
            v46 := GenS3{F0: 2, F1: 4294967295}
            _ = v46
        }
    case 405:
        for fi47 := 0; fi47 < 14; fi47++ {
            for fi48 := 0; fi48 < 2; fi48++ {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                var v49 bool = (127 < 0)
                _ = fi48
                _ = v49
            }
            v5 -= 65535
            _ = fi47
        }
    }
    v3 = append(v3, 507)
    var v50 uint8 = 255
    var v51 float32 = 1.5
    fmt.Println(v51, v2, v5)
    for fi52 := 0; fi52 < 19; fi52++ {
        v3 = append(v3, 2)
        _ = fi52
    }
    fmt.Println(v50, v4)
    switch -722 {
    case 763:
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
    case -1:
        if false {
            switch "a" {
            case "foo bar":
                v53 := GenS1{F0: 1, F1: -449, F2: ""}
                _ = v53
            case "hello":
                switch 42 {
                case -396:
                    v2 += v2
                case 128:
                    var v54 uint = 255
                    _ = v54
                case 4294967295:
                    if false {
                        for fi55 := 0; fi55 < 13; fi55++ {
                            for fi56 := 0; fi56 < 16; fi56++ {
                                if (v5 >= 2) {
                                    var v57 int = fi55
                                    if ((0.0 * 0.1) <= 1.5) {
                                        v57 = fi55
                                        var v58 float32 = 0.0
                                        var v59 string = "\n\t"
                                        _ = v58
                                        _ = v59
                                    } else {
                                        fmt.Println(v2)
                                        v60 := GenS7{F0: v5, F1: -1, F2: -1}
                                        v51 += v51
                                        _ = v60
                                    }
                                    _ = len(v3)
                                    v57 += fi55
                                    if (true && true) {
                                        if len(v1) > 0 {
                                            v1[0] = fi56
                                        }
                                        v61 := GenS8{F0: 355, F1: "", F2: v2}
                                        _ = v61
                                    } else {
                                        _ = len(v3)
                                        switch 127 {
                                        case 128:
                                            if false {
                                                if false {
                                                    if len(v3) > 0 {
                                                        v3[0] = 37
                                                    }
                                                    v62 := GenS5{F0: (v50 + 230), F1: ("a" + ""), F2: -9223372036854775808}
                                                    if false {
                                                        fmt.Println(v2, v2)
                                                        v5 = v5
                                                    }
                                                    _ = v62
                                                }
                                            } else {
                                                var v63 uint32 = 2147483647
                                                var v64 int8 = -1
                                                _ = v63
                                                _ = v64
                                            }
                                        case 2147483647:
                                            if len(v6) >= 2 {
                                                v6 = v6[1:len(v6)]
                                            }
                                        default:
                                            var v65 string = "hello"
                                            _ = v65
                                        }
                                        if false {
                                            fmt.Println(v51, v50, fi56)
                                            if (0.0 == 1e10) {
                                                var v66 uint16 = 32767
                                                _ = v66
                                            } else {
                                                var v67 float32 = v51
                                                v68 := GenS9{F0: 1, F1: v5}
                                                v50 = v50
                                                _ = v67
                                                _ = v68
                                            }
                                            if (-1 <= v2) {
                                                v5 -= (v5 << ((v5) & 63))
                                                v4 += int16(v50)
                                                if (((v2 >= v2) || (42 >= 100)) || (!false)) {
                                                    var v69 uint8 = 1
                                                    var v70 int32 = 42
                                                    _ = v69
                                                    _ = v70
                                                } else {
                                                    if true {
                                                        v50 -= (v50 / (((v50 * v50)) | 1))
                                                        v71 := GenS10{F0: "foo bar"}
                                                        v72 := GenS11{F0: true, F1: (^(v57)), F2: 208}
                                                        _ = v71
                                                        _ = v72
                                                    }
                                                }
                                                if false {
                                                    var v73 int = 430
                                                    v57 += (fi56 | fi55)
                                                    switch ("foo bar" + "") {
                                                    case "a":
                                                        switch 42 {
                                                        case -2147483648:
                                                            var v74 uint16 = 32767
                                                            _ = v74
                                                        case 32767:
                                                            switch ("a" + "\n\t") {
                                                            case "hello":
                                                                fmt.Println(v2)
                                                            default:
                                                                if ((!false) || true) {
                                                                    if len(v6) >= 2 {
                                                                        v6 = v6[1:len(v6)]
                                                                    }
                                                                }
                                                            }
                                                        default:
                                                            v2 = v2
                                                        }
                                                    case "foo bar":
                                                        var v75 int64 = -9223372036854775808
                                                        _ = v75
                                                    case "":
                                                        var v76 int32 = 1
                                                        _ = v76
                                                    default:
                                                        if false {
                                                            v77 := GenS8{F0: 18446744073709551615, F1: ("foo bar" + (("hello" + ("foo bar" + ("a" + "foo bar"))) + "a")), F2: v2}
                                                            v57 -= v73
                                                            _ = v77
                                                        } else {
                                                            if false {
                                                                var v78 int16 = (v4 - (v4 << ((v5) & 63)))
                                                                _ = v78
                                                            }
                                                        }
                                                    }
                                                    _ = v73
                                                } else {
                                                    fmt.Println(v2, fi55)
                                                    fmt.Println(v51)
                                                    var v79 bool = (0.0 <= v51)
                                                    _ = v79
                                                }
                                                v50 += 0
                                                var v80 uint = v5
                                                switch 9223372036854775807 {
                                                case 9223372036854775807:
                                                    _ = len(v6)
                                                }
                                                var v81 int64 = 9223372036854775807
                                                var v82 int64 = -1
                                                if len(v1) >= 2 {
                                                    v1 = v1[1:len(v1)]
                                                }
                                                switch ("hello" + ("foo bar" + "\n\t")) {
                                                case "foo bar":
                                                    var v83 bool = (true && (int16(v80) != 42))
                                                    _ = v83
                                                }
                                                _ = v80
                                                _ = v81
                                                _ = v82
                                            }
                                            v84 := GenS6{F0: 2, F1: v50}
                                            _ = v84
                                        } else {
                                            v85 := GenS12{F0: v5, F1: 2147483648, F2: (fi56 >> ((99) & 63))}
                                            v86 := GenS13{F0: 729}
                                            v50 = 49
                                            v87 := GenS5{F0: v50, F1: "foo bar", F2: 0}
                                            _ = v85
                                            _ = v86
                                            _ = v87
                                        }
                                    }
                                    _ = v57
                                }
                                v2 += int8(v50)
                                fmt.Println(v5, v50)
                                _ = fi56
                            }
                            v88 := GenS9{F0: v50, F1: (v5 << ((4294967295) & 63))}
                            _ = fi55
                            _ = v88
                        }
                    } else {
                        for fi89 := 0; fi89 < 17; fi89++ {
                            if len(v3) >= 2 {
                                v3 = v3[1:len(v3)]
                            }
                            _ = fi89
                        }
                        if false {
                            if len(v6) >= 2 {
                                v6 = v6[1:len(v6)]
                            }
                            switch "foo bar" {
                            case "foo bar":
                                _ = len(v6)
                            case "\n\t":
                                if len(v3) > 0 {
                                    v3[0] = 0
                                }
                            default:
                                v90 := GenS4{F0: (!false), F1: "\n\t"}
                                _ = v90
                            }
                        } else {
                            var v91 int64 = 42
                            v92 := GenS13{F0: 854}
                            fmt.Println(v91)
                            _ = v91
                            _ = v92
                        }
                    }
                default:
                    switch ((("" + "hello") + "\n\t") + "\n\t") {
                    case "":
                        var v93 uint = v5
                        _ = v93
                    case "foo bar":
                        v2 = (^(v2))
                    default:
                        var v94 float64 = ((3.14159 / 1e10) - (((1.5 + 0.1) * (-2.25 * (1.5 * 0.0))) / 1.5))
                        _ = v94
                    }
                }
            default:
                fmt.Println(v51, v2, v2)
            }
        }
    }
    var v95 int32 = -462
    v96 := GenS14{F0: uint64(v5), F1: 0, F2: false}
    v4 -= (v4 & v4)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v50:", v50)
    fmt.Println("v51:", v51)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v95:", v95)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v50
    _ = v51
    _ = v6
    _ = v7
    _ = v95
    _ = v96
}
