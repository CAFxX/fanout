package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 string; F1 uint8; F2 uint16 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 int64; F1 int64 }
type GenS7 struct { F0 uint16; F1 bool }
type GenS8 struct { F0 uint16; F1 int16 }
type GenS9 struct { F0 int32; F1 int16; F2 int }
type GenS10 struct { F0 uint8; F1 string }
type GenS11 struct { F0 int }
type GenS12 struct { F0 string }
type GenS13 struct { F0 string }
type GenS14 struct { F0 uint64; F1 uint16; F2 int16 }
type GenS15 struct { F0 uint16 }

func main() {
    v1 := []int{-463, 42, -2147483648, -923}
    v2 := []uint32{}
    v3 := []string{}
    _ = len(v2)
    v4 := GenS1{F0: 722, F1: -722}
    for fi5 := 0; fi5 < 16; fi5++ {
        switch (((("foo bar" + "") + ("" + "")) + "a") + (("foo bar" + (("a" + ("hello" + (("foo bar" + "a") + ""))) + "hello")) + "foo bar")) {
        case "":
            if len(v1) > 0 {
                v1[0] = (fi5 - (fi5 - 1))
            }
        case "foo bar":
            if true {
                var v6 uint64 = 42
                for fi7 := 0; fi7 < 17; fi7++ {
                    if ((!true) || (((!(false || ((!(!(!false))) && (!(128 < 127))))) || true) || true)) {
                        if (-525 > 0) {
                            switch "" {
                            case "a":
                                if false {
                                    v6 = v6
                                    v6 = v6
                                    var v8 int = (fi5 / ((fi5) | 1))
                                    v8 += (fi7 / ((fi7) | 1))
                                    _ = v8
                                } else {
                                    var v9 string = "hello"
                                    var v10 uint32 = 4294967295
                                    _ = v10
                                    _ = v9
                                }
                            case "":
                                v6 = 0
                            case "hello":
                                var v11 int8 = -128
                                _ = v11
                            }
                            _ = len(v3)
                        } else {
                            v6 = v6
                            var v12 int8 = -128
                            if (0.1 >= 1.5) {
                                v6 = v6
                                if (1.5 > -2.25) {
                                    switch ("foo bar" + "a") {
                                    case "":
                                        var v13 int8 = (^(v12))
                                        _ = v13
                                    case "\n\t":
                                        v12 = 0
                                    }
                                    v14 := GenS1{F0: 0, F1: 11}
                                    _ = v14
                                } else {
                                    _ = len(v3)
                                    fmt.Println(fi5, fi5)
                                    v15 := GenS2{F0: true}
                                    var v16 float32 = 1e10
                                    fmt.Println(fi5)
                                    _ = v15
                                    _ = v16
                                }
                                if (false && (true || true)) {
                                    v17 := GenS1{F0: -1, F1: -358}
                                    var v18 int8 = v12
                                    fmt.Println(v12)
                                    _ = v17
                                    _ = v18
                                } else {
                                    switch ("foo bar" + ("hello" + ("hello" + "\n\t"))) {
                                    case "":
                                        v6 += v6
                                    case "\n\t":
                                        var v19 int32 = 42
                                        _ = v19
                                    case "a":
                                        v6 += 2
                                    }
                                    v3 = append(v3, "foo bar")
                                }
                            } else {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                switch (fi5 ^ (fi7 + fi5)) {
                                case -1:
                                    _ = len(v2)
                                }
                                if true {
                                    if false {
                                        var v20 uint8 = 1
                                        if len(v3) > 0 {
                                            v3[0] = "foo bar"
                                        }
                                        if ((1e10 <= float64((1.5 / 1.0))) || true) {
                                            var v21 bool = (((-2.25 - 1.5) + 1e10) < 1e10)
                                            if v21 {
                                                v22 := GenS3{F0: "", F1: (v20 | v20), F2: uint16(v20)}
                                                _ = v22
                                            } else {
                                                var v23 uint32 = 4294967295
                                                var v24 uint64 = 18446744073709551615
                                                var v25 int16 = -1
                                                _ = v23
                                                _ = v24
                                                _ = v25
                                            }
                                            _ = v21
                                        } else {
                                            var v26 float32 = -2.25
                                            var v27 int8 = 1
                                            _ = len(v1)
                                            if (2 > 32767) {
                                                switch "\n\t" {
                                                case "a":
                                                    v28 := GenS1{F0: -9223372036854775808, F1: 2147483647}
                                                    _ = v28
                                                case "\n\t":
                                                    var v29 uint8 = 128
                                                    _ = v29
                                                case "":
                                                    var v30 int16 = 32767
                                                    _ = v30
                                                }
                                                switch (fi7 << ((2) & 63)) {
                                                case 9223372036854775807:
                                                    v26 = v26
                                                case -515:
                                                    v27 = (v12 & (v27 ^ v12))
                                                }
                                            } else {
                                                v1 = append(v1, 0)
                                                if len(v2) > 0 {
                                                    v2[0] = 4294967295
                                                }
                                                if false {
                                                    if false {
                                                        v31 := GenS2{F0: true}
                                                        _ = v31
                                                    }
                                                    v26 -= -2.25
                                                    var v32 uint = 2
                                                    var v33 string = ""
                                                    var v34 int64 = 840
                                                    v1 = append(v1, (fi5 >> ((4294967295) & 63)))
                                                    v35 := GenS4{F0: (^(v20))}
                                                    _ = v32
                                                    _ = v33
                                                    _ = v34
                                                    _ = v35
                                                } else {
                                                    var v36 int64 = 0
                                                    _ = v36
                                                }
                                            }
                                            _ = v26
                                            _ = v27
                                        }
                                        _ = v20
                                    }
                                    fmt.Println(v12)
                                } else {
                                    switch ("hello" + "foo bar") {
                                    case "foo bar":
                                        fmt.Println(v6, v12)
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    case "\n\t":
                                        var v37 bool = true
                                        _ = v37
                                    case "hello":
                                        if false {
                                            var v38 uint64 = (v6 & v6)
                                            if true {
                                                if ("hello" <= "a") {
                                                    var v39 float64 = 0.0
                                                    v6 -= 18446744073709551615
                                                    _ = v39
                                                }
                                                if false {
                                                    v6 += (v38 ^ (v6 >> ((255) & 63)))
                                                    if true {
                                                        if len(v3) > 0 {
                                                            v3[0] = "foo bar"
                                                        }
                                                        fmt.Println(v12)
                                                        if len(v1) > 0 {
                                                            v1[0] = (fi5 - 32768)
                                                        }
                                                    } else {
                                                        if false {
                                                            _ = len(v3)
                                                            switch fi5 {
                                                            case -524:
                                                                v12 -= v12
                                                                v12 -= v12
                                                            case 9223372036854775807:
                                                                if true {
                                                                    if len(v3) > 0 {
                                                                        v3[0] = ""
                                                                    }
                                                                    switch int(v6) {
                                                                    case 302:
                                                                        fmt.Println(v38)
                                                                    }
                                                                    v40 := GenS5{F0: 128}
                                                                    v41 := GenS6{F0: 42, F1: -1000}
                                                                    _ = v40
                                                                    _ = v41
                                                                }
                                                            }
                                                            v12 += v12
                                                        }
                                                        v12 += v12
                                                    }
                                                }
                                                v12 += (v12 ^ (v12 + (^(v12))))
                                            }
                                            _ = len(v1)
                                            _ = v38
                                        } else {
                                            switch "" {
                                            case "foo bar":
                                                v6 = (v6 & (v6 ^ (v6 | (^(v6)))))
                                            case "":
                                                v42 := GenS7{F0: 65535, F1: false}
                                                _ = v42
                                            case "hello":
                                                v6 += (v6 / (((v6 | v6)) | 1))
                                            }
                                        }
                                    default:
                                        var v43 uint64 = v6
                                        _ = v43
                                    }
                                    var v44 uint = 695
                                    v12 -= 1
                                    fmt.Println(v44, fi5)
                                    _ = v44
                                }
                                if len(v2) > 0 {
                                    v2[0] = 225
                                }
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                if false {
                                    v45 := GenS7{F0: 1, F1: (v6 > 954)}
                                    var v46 uint32 = 2
                                    _ = v45
                                    _ = v46
                                }
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                var v47 int64 = 42
                                var v48 float64 = 1.5
                                if (((v12 == v12) && ((true || true) || (0 > 126))) || false) {
                                    switch fi5 {
                                    case 878:
                                        v6 = 739
                                    default:
                                        fmt.Println(v47)
                                    }
                                }
                                _ = v47
                                _ = v48
                            }
                            _ = v12
                        }
                    }
                    if ((419 >= 2) && true) {
                        v49 := GenS8{F0: 1, F1: -497}
                        v50 := GenS8{F0: 200, F1: 1000}
                        var v51 uint8 = 0
                        _ = v49
                        _ = v50
                        _ = v51
                    }
                    switch 1 {
                    case 127:
                        var v52 uint32 = 1
                        _ = v52
                    }
                    _ = fi7
                }
                for fi53 := 0; fi53 < 1; fi53++ {
                    if (!true) {
                        v54 := GenS9{F0: -361, F1: 1, F2: fi53}
                        if (1 < 0) {
                            switch (^(fi53)) {
                            case 127:
                                switch "\n\t" {
                                case "hello":
                                    var v55 float64 = (3.14159 - -2.25)
                                    _ = v55
                                case "a":
                                    _ = len(v1)
                                }
                            }
                            var v56 uint = 18446744073709551615
                            fmt.Println(v56)
                            _ = v56
                        } else {
                            var v57 int32 = -448
                            v58 := GenS8{F0: 255, F1: -32768}
                            v6 = (v6 >> ((4294967295) & 63))
                            _ = v57
                            _ = v58
                        }
                        if (!true) {
                            if true {
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                            } else {
                                fmt.Println(v6, v6, v6)
                                var v59 string = ("a" + (("hello" + "a") + "\n\t"))
                                if len(v2) >= 2 {
                                    v2 = v2[1:len(v2)]
                                }
                                _ = v59
                            }
                            if (!(18446744073709551615 <= 1)) {
                                fmt.Println(fi53)
                            }
                        } else {
                            v6 = v6
                            v6 += v6
                            v60 := GenS2{F0: (!(!false))}
                            _ = v60
                        }
                        _ = v54
                    } else {
                        v6 = (^(v6))
                    }
                    if true {
                        fmt.Println(fi5)
                        v61 := GenS10{F0: 42, F1: "hello"}
                        _ = v61
                    } else {
                        fmt.Println(fi5, v6)
                        var v62 uint8 = 207
                        _ = v62
                    }
                    _ = fi53
                }
                _ = v6
            } else {
                v63 := GenS11{F0: (fi5 - (fi5 << ((65535) & 63)))}
                _ = v63
            }
        case "a":
            var v64 int64 = -670
            _ = v64
        }
        for fi65 := 0; fi65 < 18; fi65++ {
            var v66 bool = false
            _ = fi65
            _ = v66
        }
        _ = fi5
    }
    var v67 int8 = 100
    if (0 <= 65535) {
        var v68 uint16 = 460
        _ = v68
    } else {
        switch "\n\t" {
        case "foo bar":
            var v69 string = "hello"
            _ = v69
        default:
            v70 := GenS12{F0: "foo bar"}
            _ = v70
        }
        for fi71 := 0; fi71 < 2; fi71++ {
            _ = len(v1)
            _ = fi71
        }
        for fi72 := 0; fi72 < 15; fi72++ {
            var v73 int = fi72
            var v74 int64 = 43
            switch (fi72 + 0) {
            case 4294967295:
                fmt.Println(v73, v73, v74)
            case -9223372036854775808:
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            default:
                switch fi72 {
                case 0:
                    v67 = (^(v67))
                case -9223372036854775808:
                    var v75 int8 = 127
                    _ = v75
                default:
                    v74 += 1
                }
            }
            _ = fi72
            _ = v73
            _ = v74
        }
    }
    if (v67 >= (v67 + (v67 ^ 0))) {
        v76 := GenS13{F0: "a"}
        _ = v76
    }
    var v77 bool = (42 <= 1)
    fmt.Println(v77, v77)
    _ = len(v2)
    var v78 uint16 = 1
    v79 := GenS8{F0: v78, F1: 0}
    if len(v1) > 0 {
        v1[0] = 128
    }
    var v80 float64 = 1.5
    if v77 {
        fmt.Println(v77, v78)
        switch "a" {
        case "":
            v1 = append(v1, -2147483648)
        case "a":
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        if (9223372036854775807 <= int(v67)) {
            for fi81 := 0; fi81 < 12; fi81++ {
                var v82 uint32 = 8
                _ = fi81
                _ = v82
            }
            var v83 int64 = -615
            _ = v83
        } else {
            for fi84 := 0; fi84 < 7; fi84++ {
                if v77 {
                    v80 = 1.5
                } else {
                    if len(v1) > 0 {
                        v1[0] = (fi84 - fi84)
                    }
                    if v77 {
                        v78 = v78
                        v77 = v77
                        for fi85 := 0; fi85 < 13; fi85++ {
                            v86 := GenS8{F0: (v78 / ((v78) | 1)), F1: 32767}
                            v77 = false
                            _ = fi85
                            _ = v86
                        }
                    } else {
                        var v87 uint32 = 2
                        var v88 int8 = (v67 | v67)
                        if ((127 == uint8(v87)) || v77) {
                            for fi89 := 0; fi89 < 10; fi89++ {
                                switch "\n\t" {
                                case "":
                                    switch ("\n\t" + "foo bar") {
                                    case "hello":
                                        v78 = 690
                                    default:
                                        if (((("foo bar" >= "foo bar") || (v77 && false)) || (false && false)) || (!false)) {
                                            if len(v1) >= 2 {
                                                v1 = v1[1:len(v1)]
                                            }
                                            if (!false) {
                                                fmt.Println(v88, v88)
                                                v77 = ((!(((4294967295 <= 65535) && v77) || ((fi89 / ((fi89) | 1)) < fi84))) && (v87 < (^(v87))))
                                                v90 := GenS5{F0: 255}
                                                var v91 int16 = int16(fi84)
                                                _ = v90
                                                _ = v91
                                            } else {
                                                v88 = (v88 % ((int8(v87)) | 1))
                                            }
                                            fmt.Println(v80, v78, v87)
                                        } else {
                                            var v92 uint = 18446744073709551615
                                            if (!((-(v88)) < (-(v67)))) {
                                                v78 = uint16(v92)
                                            }
                                            switch "a" {
                                            case "foo bar":
                                                var v93 uint64 = 1
                                                _ = v93
                                            default:
                                                var v94 uint32 = (v87 * (^(v87)))
                                                var v95 uint = 373
                                                _ = v94
                                                _ = v95
                                            }
                                            fmt.Println(v92, fi89, v80)
                                            _ = v92
                                        }
                                    }
                                case "foo bar":
                                    v78 = v78
                                }
                                fmt.Println(v87)
                                v96 := GenS12{F0: ("a" + "\n\t")}
                                _ = fi89
                                _ = v96
                            }
                            var v97 int8 = v88
                            v77 = (((1e10 * 3.14159) + -2.25) == -2.25)
                            _ = v97
                        } else {
                            v3 = append(v3, "")
                            fmt.Println(v77)
                            var v98 int8 = (v67 >> ((42) & 63))
                            _ = v98
                        }
                        _ = v87
                        _ = v88
                    }
                }
                if (454 < 1) {
                    var v99 uint32 = 65535
                    for fi100 := 0; fi100 < 1; fi100++ {
                        v67 -= -80
                        var v101 float64 = float64(((0.0 / 1e10) - 1e10))
                        _ = fi100
                        _ = v101
                    }
                    _ = v99
                } else {
                    v102 := GenS14{F0: 204, F1: (v78 & v78), F2: 456}
                    v103 := GenS15{F0: v78}
                    v1 = append(v1, 32767)
                    _ = v102
                    _ = v103
                }
                _ = fi84
            }
            if true {
                switch 0 {
                case 0:
                    v67 = v67
                case -1:
                    fmt.Println(v67)
                default:
                    if len(v2) > 0 {
                        v2[0] = uint32(v78)
                    }
                }
                v77 = (!false)
                var v104 uint16 = 213
                _ = v104
            }
        }
    }
    if v77 {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        v67 = -128
        v80 = v80
    } else {
        v1 = append(v1, 1)
    }
    v1 = append(v1, -133)
    if ((0 < 18446744073709551615) || (!true)) {
        var v105 uint64 = 2
        if v77 {
            if false {
                for fi106 := 0; fi106 < 3; fi106++ {
                    fmt.Println(v105, v78, v78)
                    _ = fi106
                }
                switch 127 {
                case 2147483647:
                    v78 -= 1
                default:
                    if v77 {
                        _ = len(v2)
                    } else {
                        var v107 int32 = -2147483648
                        var v108 int16 = 42
                        var v109 int16 = v108
                        _ = v107
                        _ = v108
                        _ = v109
                    }
                }
                for fi110 := 0; fi110 < 9; fi110++ {
                    v105 -= (v105 + v105)
                    if v77 {
                        var v111 float64 = v80
                        _ = v111
                    }
                    _ = fi110
                }
            }
            v3 = append(v3, "a")
        } else {
            v78 = 32767
        }
        var v112 int = 32768
        _ = v105
        _ = v112
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v67:", v67)
    fmt.Println("v77:", v77)
    fmt.Println("v78:", v78)
    fmt.Println("v80:", v80)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v67
    _ = v77
    _ = v78
    _ = v79
    _ = v80
}
