package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint16; F2 uint16 }
type GenS2 struct { F0 uint16; F1 uint16; F2 uint64 }
type GenS3 struct { F0 string; F1 uint64; F2 int }
type GenS4 struct { F0 int32; F1 bool; F2 int8 }
type GenS5 struct { F0 int16; F1 uint16 }
type GenS6 struct { F0 uint8; F1 int8; F2 uint8 }
type GenS7 struct { F0 int64; F1 int32; F2 uint64 }
type GenS8 struct { F0 bool; F1 string }
type GenS9 struct { F0 string; F1 int8 }
type GenS10 struct { F0 int32; F1 uint; F2 int16 }
func genFn0(p0 uint32, p1 uint16) bool {
    p1 -= p1
    switch "a" {
    case "":
        p1 = 65535
    }
    p0 += 4294967295
    switch "foo bar" {
    case "hello":
        v6 := []int{0, 128, 689, 2147483647}
        _ = v6
    }
    _ = p0
    _ = p1
    return (!(42 == 103))
}
func genFn1(p0 uint16, p1 uint16, p2 uint16) int16 {
    for fi7 := 0; fi7 < 17; fi7++ {
        fmt.Println(fi7)
        p1 -= (p0 % ((p2) | 1))
        _ = fi7
    }
    if (!(!false)) {
        v8 := []int{28, 144, 42, -206}
        switch 42 {
        case 128:
            if false {
                var v9 int16 = 42
                _ = v9
            }
        case 1:
            if (2 == 65535) {
                v10 := GenS1{F0: 2, F1: (^(p2)), F2: p2}
                var v11 uint16 = 1
                var v12 int = int(p1)
                _ = v10
                _ = v11
                _ = v12
            } else {
                p0 = (p1 << ((255) & 63))
                for fi13 := 0; fi13 < 7; fi13++ {
                    v14 := GenS1{F0: 224, F1: (p0 * (p1 | p0)), F2: (p1 * (p2 + p0))}
                    var v15 int = (fi13 | (fi13 * -88))
                    if len(v8) >= 2 {
                        v8 = v8[1:len(v8)]
                    }
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                switch "\n\t" {
                case "\n\t":
                    var v16 uint64 = 1
                    _ = v16
                case "foo bar":
                    p2 = p0
                case "":
                    v17 := GenS2{F0: (p0 - p2), F1: (p1 % ((p2) | 1)), F2: 193}
                    _ = v17
                }
            }
        }
        var v18 int32 = 1
        _ = v18
        _ = v8
    } else {
        var v19 bool = ((-2147483648 < 1) && ((0.0 <= 0.1) || (false && (1 == 1))))
        v20 := []int{}
        v20 = append(v20, -1)
        _ = v19
        _ = v20
    }
    v21 := []int{}
    if true {
        if true {
            switch 9223372036854775807 {
            case 32768:
                var v22 uint16 = 255
                _ = v22
            case -1:
                v23 := GenS3{F0: ("hello" + "a"), F1: uint64(p2), F2: 968}
                _ = v23
            case -582:
                v24 := GenS1{F0: 889, F1: (p1 + (p1 / (((p1 | (p0 / ((p0) | 1)))) | 1))), F2: p1}
                _ = v24
            default:
                if len(v21) > 0 {
                    v21[0] = -2147483648
                }
            }
            for fi25 := 0; fi25 < 1; fi25++ {
                for fi26 := 0; fi26 < 19; fi26++ {
                    var v27 float32 = 3.14159
                    fmt.Println(p0)
                    var v28 uint = 1
                    _ = fi26
                    _ = v27
                    _ = v28
                }
                _ = fi25
            }
        } else {
            switch "foo bar" {
            case "hello":
                p2 += p0
            default:
                var v29 uint = 2
                _ = v29
            }
            if (!false) {
                v30 := GenS4{F0: -1, F1: true, F2: 100}
                var v31 int8 = -45
                _ = v30
                _ = v31
            }
        }
        if len(v21) >= 2 {
            v21 = v21[1:len(v21)]
        }
    } else {
        var v32 uint64 = 18446744073709551615
        _ = v32
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v21
    return 1
}

func main() {
    var v1 string = ""
    v2 := []int{9223372036854775807, -2147483648, 296}
    v3 := make(map[uint32]int)
    var v4 uint16 = 65535
    var v5 bool = (!false)
    var v33 float32 = 1e10
    if true {
        fmt.Println(v4, v4)
    } else {
        fmt.Println(v1)
        var v34 string = (v1 + "a")
        _ = v34
    }
    v1 = (v1 + ("\n\t" + ((v1 + (v1 + (v1 + v1))) + "a")))
    if v5 {
        switch v1 {
        case "a":
            var v35 int32 = -2147483648
            _ = v35
        case "\n\t":
            var v36 int8 = -13
            _ = v36
        case "hello":
            v37 := GenS5{F0: -696, F1: 42}
            _ = v37
        default:
            var v38 float64 = ((1e10 + 3.14159) - 0.0)
            _ = v38
        }
        switch 9223372036854775807 {
        case -890:
            fmt.Println(v1, v1, v4)
        case 4294967295:
            v1 = ("\n\t" + "a")
        case 9223372036854775807:
            for fi39 := 0; fi39 < 19; fi39++ {
                v40 := GenS3{F0: ("hello" + (("a" + (v1 + v1)) + v1)), F1: uint64(v4), F2: fi39}
                for fi41 := 0; fi41 < 4; fi41++ {
                    _ = len(v2)
                    if (42 != 18446744073709551615) {
                        v33 += 0.1
                        var v42 int = 128
                        _ = v42
                    }
                    v3[2147483647] = int(v4)
                    _ = len(v2)
                    if false {
                        v43 := GenS6{F0: 130, F1: -123, F2: 1}
                        _ = v43
                    }
                    _ = fi41
                }
                for fi44 := 0; fi44 < 5; fi44++ {
                    var v45 int = fi44
                    fmt.Println(v5, v45)
                    _ = fi44
                    _ = v45
                }
                _ = fi39
                _ = v40
            }
        }
    } else {
        var v46 uint8 = 151
        if (v1 > v1) {
            v46 = (v46 << ((1) & 63))
            v47 := GenS1{F0: 0, F1: v4, F2: (v4 << ((uint(v46)) & 63))}
            _ = v47
        }
        v3[1] = -314
        _ = v46
    }
    var v48 int = -9223372036854775808
    if true {
        var v49 uint64 = 2
        for fi50 := 0; fi50 < 9; fi50++ {
            var v51 int16 = 42
            for fi52 := 0; fi52 < 5; fi52++ {
                var v53 uint = 0
                v49 += v49
                delete(v3, 2147483648)
                _ = fi52
                _ = v53
            }
            _ = fi50
            _ = v51
        }
        _ = v49
    } else {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v54 int8 = -100
        switch v48 {
        case -1:
            fmt.Println(v54, v54, v1)
        case -822:
            v55 := GenS7{F0: -354, F1: 42, F2: 18446744073709551615}
            _ = v55
        }
        _ = v54
    }
    v2 = append(v2, v48)
    fmt.Println(v48)
    var v56 uint8 = 127
    switch "" {
    case "a":
        if (0 < 619) {
            v57 := GenS8{F0: (!((!v5) || (380 <= 918))), F1: ""}
            _ = v57
        } else {
            if true {
                switch (v48 << ((18446744073709551615) & 63)) {
                case 9223372036854775807:
                    var v58 uint32 = 99
                    _ = v58
                }
                fmt.Println(v4, v33, v48)
                _ = len(v2)
            } else {
                switch v1 {
                case "\n\t":
                    for fi59 := 0; fi59 < 6; fi59++ {
                        switch (v1 + (v1 + "hello")) {
                        case "":
                            switch (v1 + ("hello" + v1)) {
                            case "hello":
                                var v60 uint16 = (v4 ^ v4)
                                _ = v60
                            case "\n\t":
                                for fi61 := 0; fi61 < 12; fi61++ {
                                    v48 = (-(v48))
                                    if (65535 > 1) {
                                        delete(v3, 67)
                                        v4 += uint16(v56)
                                        var v62 int64 = -474
                                        v4 = 466
                                        _ = v62
                                    } else {
                                        v56 -= uint8(v4)
                                        fmt.Println(v1, v4, v48)
                                    }
                                    _ = fi61
                                }
                            }
                        }
                        v63 := GenS9{F0: v1, F1: 42}
                        var v64 uint8 = v56
                        _ = fi59
                        _ = v63
                        _ = v64
                    }
                case "":
                    v65 := GenS7{F0: -1, F1: 1, F2: 42}
                    _ = v65
                case "hello":
                    for fi66 := 0; fi66 < 9; fi66++ {
                        v48 += (fi66 / ((fi66) | 1))
                        v56 -= 42
                        _ = fi66
                    }
                }
                var v67 bool = (false && v5)
                delete(v3, 50)
                _ = v67
            }
        }
    default:
        fmt.Println(v5, v56)
    }
    var v68 uint32 = 932
    switch 127 {
    case 2147483647:
        if (0 != 0) {
            if true {
                for fi69 := 0; fi69 < 14; fi69++ {
                    _ = len(v2)
                    _ = len(v2)
                    _ = fi69
                }
            }
            fmt.Println(v5, v4)
            var v70 bool = v5
            _ = v70
        }
    }
    fmt.Println(v4, v33)
    for fi71 := 0; fi71 < 5; fi71++ {
        _ = len(v2)
        switch "a" {
        case "hello":
            v68 = v68
        case "a":
            if (42 != 2) {
                delete(v3, (v68 >> ((4294967295) & 63)))
                for fi72 := 0; fi72 < 19; fi72++ {
                    switch int(v68) {
                    case 930:
                        v73 := GenS5{F0: -733, F1: v4}
                        v3[402] = fi71
                        _ = len(v2)
                        _ = v73
                    case 9223372036854775807:
                        v33 -= 3.14159
                    case -1:
                        var v74 int8 = 100
                        _ = v74
                    default:
                        switch 1 {
                        case 42:
                            v5 = (1 >= 65535)
                        case 106:
                            v1 = "\n\t"
                        }
                    }
                    if (154 == 0) {
                        fmt.Println(fi72, fi72, v68)
                    } else {
                        var v75 int = fi71
                        if v5 {
                            fmt.Println(v5, v48)
                        }
                        var v76 uint32 = 4294967295
                        _ = v75
                        _ = v76
                    }
                    delete(v3, 436)
                    _ = fi72
                }
                fmt.Println(v4, fi71)
            }
        case "":
            fmt.Println(v68, v4)
        }
        switch "foo bar" {
        case "\n\t":
            switch v48 {
            case 128:
                var v77 uint32 = 252
                _ = v77
            default:
                switch (v1 + ("hello" + (v1 + v1))) {
                case "foo bar":
                    v33 -= 1.5
                case "":
                    if v5 {
                        if (v5 && (0 != int8(v48))) {
                            var v78 int32 = -705
                            _ = len(v3)
                            _ = v78
                        } else {
                            for fi79 := 0; fi79 < 10; fi79++ {
                                fmt.Println(v48, v1)
                                var v80 int32 = 42
                                _ = fi79
                                _ = v80
                            }
                            if v5 {
                                v2 = append(v2, v48)
                                for fi81 := 0; fi81 < 13; fi81++ {
                                    if v5 {
                                        fmt.Println(v56, v68)
                                        v82 := GenS10{F0: -2147483648, F1: 1, F2: 1000}
                                        if (!(("" <= v1) || v5)) {
                                            switch (v1 + v1) {
                                            case "hello":
                                                switch v1 {
                                                case "a":
                                                    if false {
                                                        var v83 uint8 = 41
                                                        v4 = v4
                                                        var v84 int16 = 1
                                                        switch (v1 + ("" + (v1 + (((v1 + ("\n\t" + "foo bar")) + (v1 + v1)) + (v1 + "\n\t"))))) {
                                                        case "a":
                                                            if (!v5) {
                                                                v33 += v33
                                                                var v85 string = "foo bar"
                                                                switch "hello" {
                                                                case "hello":
                                                                    v5 = false
                                                                case "foo bar":
                                                                    v3[(v68 % ((v68) | 1))] = 42
                                                                }
                                                                _ = v85
                                                            } else {
                                                                v86 := v3[v68]
                                                                if (100 <= 1) {
                                                                    fmt.Println(v1)
                                                                    v2 = append(v2, (^(v48)))
                                                                }
                                                                switch 756 {
                                                                case -2147483648:
                                                                    v48 += (fi81 & (v86 & (-(v86))))
                                                                case 505:
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                case -941:
                                                                    var v87 int64 = 0
                                                                    _ = v87
                                                                }
                                                                _ = v86
                                                            }
                                                        case "":
                                                            v84 = int16(v83)
                                                        }
                                                        _ = v83
                                                        _ = v84
                                                    } else {
                                                        var v88 float32 = (float32(1.5) - v33)
                                                        var v89 bool = (121 <= 100)
                                                        _ = len(v3)
                                                        _ = v88
                                                        _ = v89
                                                    }
                                                case "hello":
                                                    if false {
                                                        v2 = append(v2, fi81)
                                                        v3[v68] = v48
                                                    }
                                                }
                                            case "\n\t":
                                                v90 := GenS8{F0: true, F1: ("a" + v1)}
                                                _ = v90
                                            case "":
                                                var v91 string = (v1 + ((v1 + (v1 + "a")) + v1))
                                                _ = v91
                                            default:
                                                if ((v4 >> ((255) & 63)) < v4) {
                                                    fmt.Println(v5, fi71)
                                                }
                                            }
                                        }
                                        _ = v82
                                    } else {
                                        fmt.Println(v33, fi71, v68)
                                    }
                                    _ = fi81
                                }
                                for fi92 := 0; fi92 < 13; fi92++ {
                                    v5 = v5
                                    delete(v3, v68)
                                    switch ("\n\t" + v1) {
                                    case "hello":
                                        var v93 uint8 = v56
                                        _ = v93
                                    case "foo bar":
                                        v68 -= 2
                                    case "":
                                        if (32767 <= int16(v4)) {
                                            delete(v3, (v68 << ((18446744073709551615) & 63)))
                                        } else {
                                            v3[(v68 / ((uint32(v56)) | 1))] = int(v4)
                                        }
                                    default:
                                        delete(v3, (v68 - 55))
                                    }
                                    var v94 uint = 640
                                    v2 = append(v2, (fi71 & fi71))
                                    _ = fi92
                                    _ = v94
                                }
                            }
                        }
                    }
                case "\n\t":
                    var v95 bool = v5
                    _ = v95
                default:
                    fmt.Println(v1)
                }
            }
        case "":
            var v96 uint32 = 65535
            _ = v96
        default:
            var v97 int16 = 42
            _ = v97
        }
        _ = fi71
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v48:", v48)
    fmt.Println("v5:", v5)
    fmt.Println("v56:", v56)
    fmt.Println("v68:", v68)
    _ = v1
    _ = v2
    _ = v3
    _ = v33
    _ = v4
    _ = v48
    _ = v5
    _ = v56
    _ = v68
}
