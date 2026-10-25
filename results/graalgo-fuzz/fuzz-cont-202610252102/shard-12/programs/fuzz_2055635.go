package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 bool }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 uint16; F1 int8; F2 uint8 }
type GenS5 struct { F0 uint32; F1 int8 }
type GenS6 struct { F0 int8; F1 bool; F2 uint64 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 bool; F1 int32 }
type GenS9 struct { F0 string; F1 int64 }
type GenS10 struct { F0 int16; F1 uint64; F2 int8 }
type GenS11 struct { F0 string; F1 string; F2 int32 }

func main() {
    var v1 uint32 = 0
    v2 := make(map[int]bool)
    v3 := make(map[int]bool)
    var v4 float64 = 3.14159
    v5 := []string{("a" + "\n\t")}
    var v6 bool = ("a" >= ("hello" + "\n\t"))
    var v7 uint64 = 865
    v4 = 0.0
    for fi8 := 0; fi8 < 1; fi8++ {
        for fi9 := 0; fi9 < 12; fi9++ {
            if ((42 >= 32767) && (v6 || (570 > 2))) {
                var v10 uint64 = v7
                var v11 float32 = 0.0
                if (!(788 >= -9223372036854775808)) {
                    var v12 bool = ((true || v6) || ((v6 && v6) && v6))
                    _ = v12
                }
                _ = v10
                _ = v11
            } else {
                v6 = false
                var v13 uint16 = 917
                var v14 float32 = 1.5
                v15 := GenS1{F0: 127, F1: (385 != (^(v7)))}
                _ = v13
                _ = v14
                _ = v15
            }
            _ = fi9
        }
        _ = fi8
    }
    v6 = (!true)
    if (2147483647 != 2147483647) {
        v1 += v1
        v16 := GenS2{F0: v6}
        switch "a" {
        case "\n\t":
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        case "a":
            for fi17 := 0; fi17 < 19; fi17++ {
                if v6 {
                    v7 = v7
                    fmt.Println(v4)
                    for fi18 := 0; fi18 < 19; fi18++ {
                        if len(v5) > 0 {
                            v5[0] = "a"
                        }
                        _ = fi18
                    }
                } else {
                    for fi19 := 0; fi19 < 11; fi19++ {
                        v4 = ((v4 / v4) / ((v4 / v4) + (v4 / 1.0)))
                        v4 = -2.25
                        switch fi19 {
                        case 9223372036854775807:
                            if (false && true) {
                                if v6 {
                                    v5 = append(v5, ("\n\t" + ("hello" + "")))
                                    switch "" {
                                    case "":
                                        var v20 uint64 = (^(v7))
                                        _ = v20
                                    case "\n\t":
                                        v4 = v4
                                        if (v6 || v6) {
                                            if v6 {
                                                if v6 {
                                                    var v21 uint = 65535
                                                    fmt.Println(v4, v4)
                                                    _ = len(v2)
                                                    v6 = (1 == 81)
                                                    v22 := GenS3{F0: 4294967295}
                                                    _ = v21
                                                    _ = v22
                                                } else {
                                                    v6 = (uint64(v1) != v7)
                                                    v7 -= (v7 % (((v7 & 4294967295)) | 1))
                                                    v5 = append(v5, "hello")
                                                    if (!false) {
                                                        v23 := GenS1{F0: 128, F1: true}
                                                        v5 = append(v5, ("\n\t" + "a"))
                                                        _ = v23
                                                    }
                                                    var v24 string = ""
                                                    _ = v24
                                                }
                                                v25 := v2[fi19]
                                                v26 := GenS4{F0: 255, F1: 94, F2: 255}
                                                _ = v25
                                                _ = v26
                                            }
                                        }
                                    case "a":
                                        var v27 int64 = 42
                                        _ = v27
                                    default:
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                    }
                                } else {
                                    if (65535 <= 53) {
                                        fmt.Println(v7, fi17)
                                        var v28 int = 42
                                        _ = v28
                                    }
                                    var v29 string = "a"
                                    fmt.Println(v7, v6, v4)
                                    _ = v29
                                }
                                if len(v5) > 0 {
                                    v5[0] = "foo bar"
                                }
                            }
                        case -1:
                            var v30 uint32 = (v1 / ((0) | 1))
                            _ = v30
                        default:
                            v5 = append(v5, "foo bar")
                        }
                        _ = fi19
                    }
                }
                _ = len(v2)
                _ = fi17
            }
        case "foo bar":
            v1 = v1
        default:
            fmt.Println(v1, v6, v7)
        }
        _ = v16
    } else {
        switch int(v1) {
        case 32767:
            v1 -= 4294967295
        case -406:
            switch (("\n\t" + "hello") + "hello") {
            case "hello":
                if (1 == -701) {
                    if ((-2147483648 <= 0) && (-116 <= 100)) {
                        switch 9223372036854775807 {
                        case 9223372036854775807:
                            if v6 {
                                switch 42 {
                                case 32768:
                                    fmt.Println(v4, v4, v7)
                                case 914:
                                    switch ((("foo bar" + "\n\t") + ("hello" + "\n\t")) + ("\n\t" + ("a" + ((("foo bar" + "foo bar") + "foo bar") + ("hello" + (("hello" + "\n\t") + (("a" + "hello") + ""))))))) {
                                    case "\n\t":
                                        fmt.Println(v4, v4)
                                    case "a":
                                        var v31 int = 127
                                        _ = v31
                                    }
                                case -836:
                                    if v6 {
                                        var v32 int8 = 86
                                        v4 = v4
                                        _ = v32
                                    } else {
                                        v6 = (568 < 0)
                                        switch -1 {
                                        case -2147483648:
                                            if (-1 > -384) {
                                                if ((v6 && (v6 && false)) && (!(42 >= -1))) {
                                                    delete(v2, 128)
                                                }
                                                v6 = false
                                            }
                                        default:
                                            v33 := GenS4{F0: 1, F1: 100, F2: 128}
                                            _ = v33
                                        }
                                        switch 32768 {
                                        case 9223372036854775807:
                                            if (127 < -1) {
                                                v6 = (3.14159 == ((float32(-2.25) - 3.14159) + 1.5))
                                                if (!(false || (0 == 513))) {
                                                    for fi34 := 0; fi34 < 5; fi34++ {
                                                        var v35 int8 = -128
                                                        _ = fi34
                                                        _ = v35
                                                    }
                                                    var v36 uint = 0
                                                    switch 42 {
                                                    case -225:
                                                        switch ("" + "hello") {
                                                        case "":
                                                            switch "a" {
                                                            case "foo bar":
                                                                fmt.Println(v4, v1, v6)
                                                            case "hello":
                                                                var v37 float64 = (v4 * v4)
                                                                _ = v37
                                                            default:
                                                                if (!false) {
                                                                    var v38 uint32 = uint32(v36)
                                                                    var v39 uint32 = v1
                                                                    _ = v38
                                                                    _ = v39
                                                                } else {
                                                                    if (v6 && false) {
                                                                        v40 := GenS4{F0: 42, F1: 127, F2: 167}
                                                                        for fi41 := 0; fi41 < 16; fi41++ {
                                                                            v36 = 2
                                                                            _ = fi41
                                                                        }
                                                                        switch "hello" {
                                                                        case "foo bar":
                                                                            var v42 int8 = -128
                                                                            _ = v42
                                                                        case "hello":
                                                                            v1 += (v1 % ((v1) | 1))
                                                                        default:
                                                                            if v6 {
                                                                                v43 := GenS4{F0: 42, F1: 127, F2: 42}
                                                                                _ = v43
                                                                            } else {
                                                                                fmt.Println(v36)
                                                                                var v44 float64 = (v4 - (0.1 - v4))
                                                                                _ = v44
                                                                            }
                                                                        }
                                                                        _ = v40
                                                                    } else {
                                                                        v7 -= v7
                                                                        var v45 int32 = 980
                                                                        delete(v2, 128)
                                                                        _ = v45
                                                                    }
                                                                    switch 1 {
                                                                    case 718:
                                                                        var v46 uint64 = (^(v7))
                                                                        _ = v46
                                                                    case 1:
                                                                        for fi47 := 0; fi47 < 12; fi47++ {
                                                                            delete(v3, (fi47 << ((65535) & 63)))
                                                                            _ = fi47
                                                                        }
                                                                    default:
                                                                        for fi48 := 0; fi48 < 18; fi48++ {
                                                                            v2[fi48] = (1 == 174)
                                                                            fmt.Println(v36)
                                                                            _ = fi48
                                                                        }
                                                                    }
                                                                }
                                                            }
                                                        case "foo bar":
                                                            if len(v5) >= 2 {
                                                                v5 = v5[1:len(v5)]
                                                            }
                                                        }
                                                    case 42:
                                                        switch 0 {
                                                        case 32767:
                                                            if (float64(((1e10 - 3.14159) + (-2.25 + (-2.25 / (1e10 * ((((0.0 - 3.14159) + (float32(((-2.25 * v4) + v4)) + 3.14159)) - 0.0) - 1e10)))))) >= 3.14159) {
                                                                switch "a" {
                                                                case "a":
                                                                    switch 698 {
                                                                    case 127:
                                                                        switch 433 {
                                                                        case 2147483647:
                                                                            v49 := GenS2{F0: ("hello" == (((("foo bar" + ("\n\t" + "foo bar")) + "hello") + "hello") + "foo bar"))}
                                                                            _ = v49
                                                                        default:
                                                                            var v50 uint16 = 65535
                                                                            _ = v50
                                                                        }
                                                                    case 9223372036854775807:
                                                                        if v6 {
                                                                            if (2147483647 != -755) {
                                                                                v1 -= 2147483647
                                                                            }
                                                                            fmt.Println(v7)
                                                                        } else {
                                                                            switch (("\n\t" + ("" + "\n\t")) + "hello") {
                                                                            case "hello":
                                                                                var v51 uint = (v36 * (v36 * (v36 % ((v36) | 1))))
                                                                                _ = v51
                                                                            case "":
                                                                                var v52 int16 = 32767
                                                                                _ = v52
                                                                            default:
                                                                                if false {
                                                                                    if (v6 || (v4 != -2.25)) {
                                                                                        v36 += v36
                                                                                    } else {
                                                                                        v7 += (v7 << ((v36) & 63))
                                                                                    }
                                                                                    var v53 int32 = -31
                                                                                    _ = v53
                                                                                } else {
                                                                                    switch 1 {
                                                                                    case -701:
                                                                                        var v54 string = (("hello" + "hello") + "hello")
                                                                                        _ = v54
                                                                                    case 9223372036854775807:
                                                                                        for fi55 := 0; fi55 < 13; fi55++ {
                                                                                            _ = len(v5)
                                                                                            if v6 {
                                                                                                for fi56 := 0; fi56 < 19; fi56++ {
                                                                                                    v57 := GenS2{F0: (v6 || (v7 != 654))}
                                                                                                    _ = fi56
                                                                                                    _ = v57
                                                                                                }
                                                                                                var v58 int8 = -74
                                                                                                _ = v58
                                                                                            } else {
                                                                                                v59 := GenS5{F0: v1, F1: 100}
                                                                                                switch 836 {
                                                                                                case 32768:
                                                                                                    if (true && true) {
                                                                                                        v4 -= v4
                                                                                                        v4 += ((float64((1e10 * 0.1)) / -2.25) / v4)
                                                                                                        v60 := GenS6{F0: 1, F1: v6, F2: v7}
                                                                                                        _ = v60
                                                                                                    } else {
                                                                                                        for fi61 := 0; fi61 < 16; fi61++ {
                                                                                                            v4 -= 0.0
                                                                                                            _ = fi61
                                                                                                        }
                                                                                                        v6 = (!v6)
                                                                                                        var v62 float64 = v4
                                                                                                        _ = v62
                                                                                                    }
                                                                                                case 42:
                                                                                                    v6 = (!true)
                                                                                                default:
                                                                                                    switch (("hello" + "foo bar") + "hello") {
                                                                                                    case "foo bar":
                                                                                                        v1 = 38
                                                                                                    default:
                                                                                                        v36 = 65535
                                                                                                    }
                                                                                                }
                                                                                                var v63 int32 = 912
                                                                                                _ = v59
                                                                                                _ = v63
                                                                                            }
                                                                                            _ = fi55
                                                                                        }
                                                                                    case 128:
                                                                                        v64 := GenS6{F0: 100, F1: v6, F2: v7}
                                                                                        _ = v64
                                                                                    default:
                                                                                        for fi65 := 0; fi65 < 13; fi65++ {
                                                                                            if v6 {
                                                                                                _ = len(v5)
                                                                                                v2[fi65] = false
                                                                                            } else {
                                                                                                v66 := GenS7{F0: (("foo bar" + "foo bar") + "")}
                                                                                                var v67 float32 = 3.14159
                                                                                                var v68 int32 = 0
                                                                                                _ = v66
                                                                                                _ = v67
                                                                                                _ = v68
                                                                                            }
                                                                                            _ = fi65
                                                                                        }
                                                                                    }
                                                                                    var v69 int = -619
                                                                                    if v6 {
                                                                                        v5 = append(v5, "hello")
                                                                                    }
                                                                                    _ = v69
                                                                                }
                                                                            }
                                                                            switch 128 {
                                                                            case 9223372036854775807:
                                                                                for fi70 := 0; fi70 < 9; fi70++ {
                                                                                    var v71 float32 = float32(v4)
                                                                                    _ = fi70
                                                                                    _ = v71
                                                                                }
                                                                            case 32767:
                                                                                fmt.Println(v7, v36, v36)
                                                                            case -2147483648:
                                                                                v36 -= (v36 * (v36 % ((v36) | 1)))
                                                                            }
                                                                            v6 = (v6 || ((v36 * v36) <= 18446744073709551615))
                                                                        }
                                                                    case 2147483647:
                                                                        v72 := GenS8{F0: (565 < 32767), F1: -2147483648}
                                                                        _ = v72
                                                                    }
                                                                case "hello":
                                                                    if v6 {
                                                                        v4 = v4
                                                                        switch 42 {
                                                                        case 2147483647:
                                                                            for fi73 := 0; fi73 < 15; fi73++ {
                                                                                var v74 bool = v6
                                                                                if true {
                                                                                    v7 -= 42
                                                                                }
                                                                                _ = fi73
                                                                                _ = v74
                                                                            }
                                                                        case 128:
                                                                            v1 -= v1
                                                                        }
                                                                    } else {
                                                                        var v75 int8 = -1
                                                                        var v76 uint16 = 1
                                                                        v77 := GenS9{F0: "foo bar", F1: 42}
                                                                        _ = v75
                                                                        _ = v76
                                                                        _ = v77
                                                                    }
                                                                default:
                                                                    var v78 int = 4294967295
                                                                    _ = v78
                                                                }
                                                            }
                                                        case 27:
                                                            var v79 uint8 = 0
                                                            _ = v79
                                                        case 0:
                                                            v80 := GenS7{F0: "a"}
                                                            _ = v80
                                                        }
                                                    case 0:
                                                        if len(v5) >= 2 {
                                                            v5 = v5[1:len(v5)]
                                                        }
                                                    }
                                                    _ = v36
                                                }
                                            }
                                        case 474:
                                            if len(v5) > 0 {
                                                v5[0] = (("a" + "a") + "hello")
                                            }
                                        case 2147483647:
                                            var v81 int8 = -128
                                            _ = v81
                                        }
                                    }
                                default:
                                    if v6 {
                                        fmt.Println(v7)
                                        if (0 > uint8(v1)) {
                                            var v82 uint64 = v7
                                            v2[-533] = (!(2147483647 >= -2147483648))
                                            _ = v82
                                        } else {
                                            fmt.Println(v6)
                                            v7 = 42
                                            v7 = 18446744073709551615
                                        }
                                        fmt.Println(v4, v6)
                                    }
                                }
                                v4 += v4
                            }
                        case -9223372036854775808:
                            v83 := GenS3{F0: 0}
                            _ = v83
                        }
                    }
                } else {
                    if (v6 || v6) {
                        v7 += 222
                        var v84 uint32 = v1
                        _ = v84
                    }
                    var v85 uint64 = (^(v7))
                    var v86 uint8 = 2
                    _ = v85
                    _ = v86
                }
            case "\n\t":
                var v87 uint16 = 1
                _ = v87
            default:
                v4 += 1.5
            }
        case 2147483647:
            if v6 {
                var v88 int = 128
                _ = v88
            } else {
                if true {
                    var v89 int = 128
                    _ = v89
                }
            }
        default:
            v90 := GenS10{F0: 42, F1: v7, F2: 0}
            _ = v90
        }
        var v91 int = 2147483647
        if ("a" != ("" + "a")) {
            v91 = (v91 >> ((2) & 63))
        } else {
            var v92 string = ("hello" + (("foo bar" + "a") + "a"))
            _ = v92
        }
        _ = v91
    }
    v4 = 0.1
    var v93 uint = 2
    if (v6 && (-2.25 <= float32(v4))) {
        v94 := GenS11{F0: ("hello" + "hello"), F1: "a", F2: -2147483648}
        v1 -= 2
        _ = v94
    } else {
        var v95 uint32 = v1
        _ = v95
    }
    v93 -= 4294967295
    v1 = v1
    fmt.Println(v7)
    var v96 string = "hello"
    var v97 int8 = -1
    var v98 uint32 = uint32(v7)
    v7 = v7
    if (v7 >= 18446744073709551615) {
        v99 := GenS8{F0: ((v98 - (v98 << ((v93) & 63))) == (v1 * 4294967295)), F1: 1}
        v1 -= uint32(v7)
        v1 = 0
        _ = v99
    } else {
        v5 = append(v5, (((((v96 + v96) + "\n\t") + v96) + "foo bar") + ("hello" + (v96 + "foo bar"))))
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v93:", v93)
    fmt.Println("v96:", v96)
    fmt.Println("v97:", v97)
    fmt.Println("v98:", v98)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v93
    _ = v96
    _ = v97
    _ = v98
}
