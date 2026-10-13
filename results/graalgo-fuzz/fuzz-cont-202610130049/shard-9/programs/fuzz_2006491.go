package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32 }
type GenS2 struct { F0 uint; F1 int; F2 uint32 }
type GenS3 struct { F0 int16; F1 uint16; F2 int }
type GenS4 struct { F0 uint64; F1 int16 }
type GenS5 struct { F0 uint; F1 uint; F2 int }
type GenS6 struct { F0 uint16; F1 uint8 }
type GenS7 struct { F0 string }
type GenS8 struct { F0 uint64; F1 uint8; F2 bool }

func main() {
    v1 := make(map[uint32]string)
    var v2 bool = (42 == 9223372036854775807)
    v3 := make(map[string]string)
    v4 := []string{"hello", "\n\t"}
    var v5 int16 = 743
    var v6 int32 = 0
    delete(v3, ("" + ""))
    for fi7 := 0; fi7 < 10; fi7++ {
        if (((((1e10 * 0.1) / 1.5) <= (0.1 * 1e10)) && (!(42 <= 0))) || (v5 <= (v5 ^ -1))) {
            if v2 {
                v2 = v2
                fmt.Println(fi7, v2)
                v2 = (146 != 2147483648)
            } else {
                v2 = (v2 || v2)
                for fi8 := 0; fi8 < 10; fi8++ {
                    if len(v4) > 0 {
                        v4[0] = "a"
                    }
                    _ = fi8
                }
                var v9 int64 = 23
                _ = v9
            }
        } else {
            switch (^(fi7)) {
            case -2147483648:
                if (!(!(v2 && false))) {
                    var v10 uint64 = 0
                    if (1 <= -91) {
                        for fi11 := 0; fi11 < 20; fi11++ {
                            if len(v4) > 0 {
                                v4[0] = "foo bar"
                            }
                            v12 := GenS1{F0: 1, F1: 65535}
                            _ = fi11
                            _ = v12
                        }
                        switch (^(fi7)) {
                        case 1:
                            for fi13 := 0; fi13 < 12; fi13++ {
                                v6 -= (v6 + v6)
                                if (!false) {
                                    v4 = append(v4, ("\n\t" + "hello"))
                                    var v14 bool = true
                                    _ = v14
                                } else {
                                    if false {
                                        v15 := GenS1{F0: 74, F1: 554}
                                        v16 := GenS2{F0: 942, F1: (fi7 & (fi13 | (fi13 >> ((65535) & 63)))), F2: 42}
                                        _ = v15
                                        _ = v16
                                    } else {
                                        if v2 {
                                            v2 = v2
                                            fmt.Println(v5)
                                        } else {
                                            if v2 {
                                                fmt.Println(fi7, v10, v10)
                                            } else {
                                                v17 := GenS2{F0: 0, F1: fi13, F2: 65535}
                                                v2 = true
                                                switch "foo bar" {
                                                case "hello":
                                                    v4 = append(v4, "a")
                                                case "\n\t":
                                                    v18 := GenS3{F0: -32768, F1: 32767, F2: -185}
                                                    _ = v18
                                                case "a":
                                                    var v19 bool = (0 != 1)
                                                    _ = v19
                                                default:
                                                    v10 = (v10 & (v10 / ((2) | 1)))
                                                }
                                                _ = v17
                                            }
                                        }
                                        if v2 {
                                            v10 -= (^(v10))
                                        } else {
                                            v6 += v6
                                            v20 := GenS2{F0: 1, F1: fi13, F2: 2147483647}
                                            _ = v20
                                        }
                                    }
                                    switch "foo bar" {
                                    case "\n\t":
                                        v5 += -162
                                    case "foo bar":
                                        var v21 uint8 = uint8(v10)
                                        _ = v21
                                    default:
                                        var v22 bool = (!((v10 << ((2) & 63)) <= 42))
                                        _ = v22
                                    }
                                    v23 := GenS1{F0: 142, F1: 660}
                                    _ = v23
                                }
                                _ = fi13
                            }
                        case 4294967295:
                            v24 := GenS3{F0: (^(v5)), F1: 32767, F2: (fi7 >> ((0) & 63))}
                            _ = v24
                        case -345:
                            if (v2 || v2) {
                                var v25 uint16 = 65535
                                delete(v3, ((("\n\t" + "\n\t") + "\n\t") + ("" + ("foo bar" + "\n\t"))))
                                var v26 bool = v2
                                _ = v25
                                _ = v26
                            }
                        default:
                            if v2 {
                                switch (("hello" + "") + ("\n\t" + "a")) {
                                case "":
                                    var v27 int16 = (v5 % (((v5 & v5)) | 1))
                                    _ = v27
                                case "hello":
                                    fmt.Println(v5, v5)
                                case "a":
                                    switch ("a" + "") {
                                    case "hello":
                                        var v28 float64 = 1e10
                                        _ = v28
                                    case "foo bar":
                                        v29 := GenS2{F0: 255, F1: fi7, F2: 0}
                                        _ = v29
                                    case "":
                                        v2 = v2
                                    }
                                }
                            }
                        }
                    } else {
                        var v30 int8 = 77
                        if len(v4) > 0 {
                            v4[0] = "\n\t"
                        }
                        switch (fi7 >> ((2) & 63)) {
                        case -2147483648:
                            v31 := GenS2{F0: 828, F1: (fi7 + int(v10)), F2: 4294967295}
                            _ = v31
                        case 32767:
                            var v32 uint16 = 0
                            _ = v32
                        case 127:
                            v5 = -1
                        }
                        _ = v30
                    }
                    fmt.Println(fi7)
                    _ = v10
                }
            case -521:
                fmt.Println(fi7)
            case -1:
                _ = len(v3)
            default:
                v33 := GenS4{F0: 0, F1: (^(v5))}
                _ = v33
            }
            switch "\n\t" {
            case "\n\t":
                if v2 {
                    var v34 int16 = -32768
                    switch "foo bar" {
                    case "\n\t":
                        v6 -= 0
                    }
                    _ = v34
                }
            case "a":
                if (32767 == 65535) {
                    v4 = append(v4, "\n\t")
                }
            case "hello":
                var v35 int16 = (v5 ^ v5)
                _ = v35
            default:
                v5 -= -173
            }
        }
        for fi36 := 0; fi36 < 5; fi36++ {
            v5 += v5
            _ = fi36
        }
        var v37 uint32 = 2
        _ = fi7
        _ = v37
    }
    for fi38 := 0; fi38 < 8; fi38++ {
        v39 := GenS1{F0: 127, F1: 2147483647}
        _ = fi38
        _ = v39
    }
    switch "\n\t" {
    case "hello":
        for fi40 := 0; fi40 < 5; fi40++ {
            if (255 != 102) {
                fmt.Println(fi40, v5)
            }
            _ = fi40
        }
    case "\n\t":
        fmt.Println(v5, v5)
    default:
        v41 := v3[""]
        _ = v41
    }
    if ((2 >= 42) && ((1e10 * -2.25) == (0.1 / float64((1.5 - 0.1))))) {
        var v42 uint32 = 42
        _ = v42
    }
    v2 = v2
    v43 := GenS4{F0: 2, F1: (v5 & (v5 + v5))}
    switch (("foo bar" + "a") + "") {
    case "a":
        v2 = true
    }
    for fi44 := 0; fi44 < 20; fi44++ {
        v6 += (v6 % (((v6 * v6)) | 1))
        _ = fi44
    }
    switch -681 {
    case 586:
        var v45 uint32 = 2147483647
        _ = v45
    default:
        v2 = (!(0 < 42))
    }
    var v46 int32 = (v6 >> ((1) & 63))
    fmt.Println(v6, v6, v5)
    v4 = append(v4, "a")
    switch "foo bar" {
    case "a":
        v47 := GenS3{F0: (v5 >> ((545) & 63)), F1: 966, F2: 2147483647}
        _ = v47
    }
    var v48 uint = 528
    for fi49 := 0; fi49 < 11; fi49++ {
        fmt.Println(v2)
        var v50 uint64 = 18446744073709551615
        if true {
            var v51 uint = (v48 ^ (v48 * v48))
            if ((v50 & (^(v50))) == (v50 - (^(v50)))) {
                _ = len(v3)
            } else {
                for fi52 := 0; fi52 < 13; fi52++ {
                    fmt.Println(v50)
                    v1[1] = ""
                    _ = fi52
                }
                if v2 {
                    var v53 uint8 = 127
                    if v2 {
                        fmt.Println(v5, v5)
                        for fi54 := 0; fi54 < 4; fi54++ {
                            switch 1 {
                            case 9223372036854775807:
                                v46 = v6
                            default:
                                if ("foo bar" != "hello") {
                                    var v55 int64 = 1
                                    _ = v55
                                } else {
                                    v56 := GenS2{F0: v48, F1: 4294967295, F2: 42}
                                    _ = v56
                                }
                            }
                            v57 := GenS4{F0: (v50 ^ v50), F1: v5}
                            if (255 == v48) {
                                v58 := GenS3{F0: (v5 << ((1) & 63)), F1: 32767, F2: fi49}
                                switch fi49 {
                                case 2147483647:
                                    v59 := GenS5{F0: (v51 & 1), F1: (v51 ^ v48), F2: fi54}
                                    _ = v59
                                case -587:
                                    if (!v2) {
                                        var v60 float64 = 0.1
                                        var v61 uint = v48
                                        var v62 uint = (v61 ^ v51)
                                        if false {
                                            var v63 bool = ((((fi54 * (fi49 ^ int(v53))) > 9223372036854775807) && (v60 == v60)) || v2)
                                            v4 = append(v4, "a")
                                            _ = v63
                                        } else {
                                            v6 = v46
                                            var v64 int8 = -26
                                            _ = v64
                                        }
                                        v65 := GenS6{F0: 65535, F1: (v53 | v53)}
                                        _ = v60
                                        _ = v61
                                        _ = v62
                                        _ = v65
                                    } else {
                                        v50 = (v50 >> ((v51) & 63))
                                        var v66 float32 = -2.25
                                        _ = v66
                                    }
                                default:
                                    var v67 float32 = 3.14159
                                    _ = v67
                                }
                                _ = len(v4)
                                _ = v58
                            } else {
                                v53 += 255
                                v5 += (v5 / ((int16(v50)) | 1))
                                _ = len(v3)
                            }
                            _ = fi54
                            _ = v57
                        }
                    }
                    _ = v53
                }
            }
            _ = v51
        } else {
            if (v2 || (false && v2)) {
                v68 := v1[929]
                if v2 {
                    for fi69 := 0; fi69 < 6; fi69++ {
                        v4 = append(v4, v68)
                        _ = fi69
                    }
                    fmt.Println(v2)
                } else {
                    for fi70 := 0; fi70 < 9; fi70++ {
                        if false {
                            if false {
                                switch 4294967295 {
                                case 4294967295:
                                    v6 -= 0
                                case 32768:
                                    var v71 uint = (v48 + v48)
                                    _ = v71
                                case 153:
                                    v72 := GenS6{F0: 65535, F1: 229}
                                    _ = v72
                                }
                                v6 -= (v46 & (^(v6)))
                                fmt.Println(v6)
                            } else {
                                var v73 uint16 = 65535
                                fmt.Println(v48, fi49)
                                v2 = false
                                _ = v73
                            }
                        }
                        var v74 uint8 = 1
                        _ = fi70
                        _ = v74
                    }
                    for fi75 := 0; fi75 < 9; fi75++ {
                        fmt.Println(v6)
                        if (false || v2) {
                            v3[""] = ("\n\t" + v68)
                            if len(v4) > 0 {
                                v4[0] = v68
                            }
                            var v76 string = v68
                            fmt.Println(v68)
                            v77 := GenS5{F0: (v48 << ((1) & 63)), F1: uint(v50), F2: (^(fi75))}
                            var v78 string = v68
                            _ = v76
                            _ = v77
                            _ = v78
                        } else {
                            var v79 int8 = 100
                            v80 := GenS7{F0: "hello"}
                            _ = v79
                            _ = v80
                        }
                        v4 = append(v4, v68)
                        _ = fi75
                    }
                    for fi81 := 0; fi81 < 10; fi81++ {
                        v4 = append(v4, v68)
                        v82 := GenS2{F0: 42, F1: fi81, F2: 42}
                        _ = fi81
                        _ = v82
                    }
                }
                v83 := GenS8{F0: v50, F1: 128, F2: false}
                _ = v68
                _ = v83
            } else {
                v2 = v2
                switch fi49 {
                case 2147483647:
                    for fi84 := 0; fi84 < 14; fi84++ {
                        v85 := GenS5{F0: v48, F1: (^(v48)), F2: (fi49 % ((fi49) | 1))}
                        if v2 {
                            v50 = (^(v50))
                            switch fi84 {
                            case 9223372036854775807:
                                var v86 float64 = (0.0 * (1e10 / 1e10))
                                _ = v86
                            case 0:
                                if len(v4) > 0 {
                                    v4[0] = ""
                                }
                                v87 := GenS2{F0: (v48 | (v48 ^ 906)), F1: (fi49 & fi49), F2: 1}
                                _ = v87
                            case 380:
                                v6 = v6
                            }
                            var v88 bool = true
                            switch (fi49 ^ fi84) {
                            case -834:
                                if false {
                                    fmt.Println(v5)
                                }
                            case -150:
                                v89 := GenS8{F0: v50, F1: 42, F2: (!v88)}
                                var v90 uint64 = uint64(v48)
                                var v91 string = (("hello" + ("hello" + ("foo bar" + "a"))) + "foo bar")
                                _ = v89
                                _ = v90
                                _ = v91
                            }
                            if len(v4) > 0 {
                                v4[0] = ""
                            }
                            _ = v88
                        }
                        _ = fi84
                        _ = v85
                    }
                case -1:
                    v48 -= (v48 + uint(v50))
                default:
                    for fi92 := 0; fi92 < 4; fi92++ {
                        var v93 int8 = 1
                        fmt.Println(v2, v93)
                        _ = fi92
                        _ = v93
                    }
                }
            }
            if v2 {
                var v94 bool = v2
                _ = v94
            } else {
                fmt.Println(v5, v50)
                fmt.Println(v48, v46, v6)
            }
            v2 = false
        }
        _ = fi49
        _ = v50
    }
    var v95 string = (("" + "foo bar") + (("\n\t" + "") + ("" + "")))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v46:", v46)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v95:", v95)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v43
    _ = v46
    _ = v48
    _ = v5
    _ = v6
    _ = v95
}
