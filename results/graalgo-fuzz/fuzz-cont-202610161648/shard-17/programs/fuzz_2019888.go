package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int; F1 string; F2 uint8 }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 int16 }
type GenS8 struct { F0 int32; F1 uint32 }
type GenS9 struct { F0 uint8; F1 int64 }
func genFn0(p0 string, p1 int16) uint64 {
    p1 += (-(p1))
    var v7 bool = (p0 <= (((p0 + p0) + p0) + "a"))
    fmt.Println(p1)
    var v8 int = 127
    if (true || (100 < 1)) {
        v8 = v8
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 0
}
func genFn1() string {
    v9 := []int{9223372036854775807, -1}
    if (!false) {
        var v10 int64 = 42
        var v11 float32 = float32(1e10)
        _ = v10
        _ = v11
    } else {
        v12 := GenS1{F0: 100}
        _ = v12
    }
    for fi13 := 0; fi13 < 13; fi13++ {
        if (127 > -1) {
            var v14 uint8 = 0
            v15 := GenS2{F0: 0}
            if (false && true) {
                if (32767 > -32768) {
                    var v16 uint = 0
                    var v17 uint8 = (v14 % ((v14) | 1))
                    var v18 uint16 = 448
                    _ = v16
                    _ = v17
                    _ = v18
                } else {
                    if len(v9) > 0 {
                        v9[0] = fi13
                    }
                    switch (fi13 - (fi13 >> ((65535) & 63))) {
                    case 4294967295:
                        if false {
                            v9 = append(v9, fi13)
                        }
                    case -9223372036854775808:
                        v14 -= 127
                    default:
                        if len(v9) >= 2 {
                            v9 = v9[1:len(v9)]
                        }
                    }
                    for fi19 := 0; fi19 < 11; fi19++ {
                        v20 := GenS3{F0: 1}
                        var v21 string = "a"
                        if false {
                            v22 := GenS1{F0: 1}
                            _ = v22
                        }
                        v23 := GenS1{F0: 0}
                        _ = fi19
                        _ = v20
                        _ = v21
                        _ = v23
                    }
                }
                for fi24 := 0; fi24 < 14; fi24++ {
                    v14 = (v14 | v14)
                    _ = fi24
                }
                v14 -= (v14 >> ((4294967295) & 63))
            }
            _ = v14
            _ = v15
        }
        fmt.Println(fi13)
        _ = fi13
    }
    var v25 int32 = 642
    switch 127 {
    case 42:
        if (-1 >= -32768) {
            for fi26 := 0; fi26 < 11; fi26++ {
                v25 = (-(v25))
                var v27 int8 = -1
                _ = fi26
                _ = v27
            }
            for fi28 := 0; fi28 < 9; fi28++ {
                var v29 int16 = 1000
                var v30 uint8 = 0
                _ = fi28
                _ = v29
                _ = v30
            }
        } else {
            var v31 int64 = 0
            v9 = append(v9, -909)
            _ = v31
        }
    case 1:
        for fi32 := 0; fi32 < 6; fi32++ {
            _ = len(v9)
            v33 := GenS4{F0: 737}
            if (!((("foo bar" + "") + "\n\t") > ("\n\t" + (("" + "hello") + "hello")))) {
                var v34 string = "\n\t"
                switch fi32 {
                case 9223372036854775807:
                    var v35 float32 = -2.25
                    _ = v35
                default:
                    fmt.Println(v34)
                }
                _ = v34
            }
            _ = fi32
            _ = v33
        }
    default:
        v25 = (v25 ^ v25)
    }
    _ = v25
    _ = v9
    return ("hello" + "foo bar")
}

func main() {
    var v1 float32 = ((3.14159 * 0.1) + 3.14159)
    var v2 uint = 1
    var v3 int64 = 1
    var v4 int8 = -15
    v5 := []uint32{}
    var v6 int16 = 910
    if (50 > 926) {
        for fi36 := 0; fi36 < 16; fi36++ {
            v6 = (v6 / ((v6) | 1))
            switch (("foo bar" + "foo bar") + "foo bar") {
            case "\n\t":
                v6 += (v6 / ((v6) | 1))
            default:
                fmt.Println(fi36)
            }
            var v37 int32 = -485
            _ = fi36
            _ = v37
        }
    } else {
        if (1 != v4) {
            v38 := GenS3{F0: v3}
            _ = v38
        } else {
            if false {
                _ = len(v5)
                if (("foo bar" + ((((("\n\t" + ("foo bar" + ("\n\t" + ("a" + "a")))) + "foo bar") + ("foo bar" + "\n\t")) + "foo bar") + "foo bar")) > ("\n\t" + (("a" + "foo bar") + "\n\t"))) {
                    v39 := GenS1{F0: (v4 | v4)}
                    _ = v39
                } else {
                    for fi40 := 0; fi40 < 16; fi40++ {
                        if (!true) {
                            if (!(524 == 1)) {
                                var v41 uint8 = 1
                                _ = v41
                            } else {
                                var v42 float32 = 0.1
                                var v43 int32 = int32(v4)
                                _ = v42
                                _ = v43
                            }
                            v1 += ((((3.14159 + 0.0) - v1) + 0.0) / v1)
                        } else {
                            v44 := GenS4{F0: v2}
                            if len(v5) >= 2 {
                                v5 = v5[1:len(v5)]
                            }
                            switch "a" {
                            case "":
                                v6 = v6
                            case "\n\t":
                                fmt.Println(fi40, v1)
                            default:
                                v1 = v1
                            }
                            _ = v44
                        }
                        fmt.Println(v3, v6, v1)
                        _ = fi40
                    }
                    var v45 bool = (true || (true || (!false)))
                    for fi46 := 0; fi46 < 1; fi46++ {
                        var v47 uint16 = 42
                        var v48 bool = (!(!false))
                        _ = fi46
                        _ = v47
                        _ = v48
                    }
                    _ = v45
                }
                var v49 int = -9223372036854775808
                _ = v49
            } else {
                switch "" {
                case "":
                    for fi50 := 0; fi50 < 8; fi50++ {
                        if len(v5) > 0 {
                            v5[0] = 0
                        }
                        _ = fi50
                    }
                case "hello":
                    var v51 float32 = (v1 / v1)
                    _ = v51
                case "foo bar":
                    fmt.Println(v3, v3, v4)
                }
            }
            for fi52 := 0; fi52 < 9; fi52++ {
                if false {
                    v4 = (v4 | (v4 / ((96) | 1)))
                    _ = len(v5)
                    v6 = (v6 + v6)
                }
                _ = fi52
            }
            v1 -= 1e10
        }
        switch (("\n\t" + "hello") + ("\n\t" + "a")) {
        case "foo bar":
            var v53 float32 = 1e10
            _ = v53
        case "":
            for fi54 := 0; fi54 < 18; fi54++ {
                if ((int64(v6) != (v3 + 9223372036854775807)) || (999 != 2)) {
                    var v55 string = "hello"
                    _ = v55
                } else {
                    v2 = (v2 / (((v2 ^ v2)) | 1))
                    if false {
                        v56 := GenS5{F0: (-(fi54)), F1: ("a" + ""), F2: 1}
                        _ = v56
                    }
                    v57 := GenS6{F0: 1}
                    _ = v57
                }
                _ = fi54
            }
        default:
            for fi58 := 0; fi58 < 19; fi58++ {
                v4 -= -1
                switch "foo bar" {
                case "foo bar":
                    switch fi58 {
                    case 2147483647:
                        v4 = v4
                    case 9223372036854775807:
                        var v59 int64 = v3
                        _ = v59
                    default:
                        if len(v5) > 0 {
                            v5[0] = 0
                        }
                    }
                default:
                    v60 := GenS1{F0: v4}
                    _ = v60
                }
                var v61 int = fi58
                _ = fi58
                _ = v61
            }
        }
    }
    if (false || (v1 > 3.14159)) {
        v1 -= v1
        v62 := GenS1{F0: 71}
        _ = v62
    } else {
        var v63 uint32 = 2147483647
        fmt.Println(v6, v3, v1)
        _ = v63
    }
    _ = len(v5)
    v4 = (v4 >> (((v2 ^ v2)) & 63))
    v1 = v1
    switch "hello" {
    case "":
        if len(v5) > 0 {
            v5[0] = 392
        }
    case "foo bar":
        v4 = 0
    }
    var v64 uint16 = uint16(v2)
    var v65 bool = (!(false && (false && (genFn0("hello", (v6 >> ((v2) & 63))) != 858))))
    if v65 {
        for fi66 := 0; fi66 < 5; fi66++ {
            v67 := GenS4{F0: (v2 * 255)}
            _ = fi66
            _ = v67
        }
    }
    v3 = v3
    if v65 {
        for fi68 := 0; fi68 < 10; fi68++ {
            if (!(v65 && (v2 != v2))) {
                var v69 int16 = (^(v6))
                _ = v69
            }
            var v70 int = fi68
            for fi71 := 0; fi71 < 10; fi71++ {
                if (v1 < 0.1) {
                    fmt.Println(v6, v2, v70)
                } else {
                    if (-9223372036854775808 > (v3 / ((0) | 1))) {
                        switch fi68 {
                        case 891:
                            v72 := GenS1{F0: v4}
                            _ = v72
                        case 859:
                            var v73 float32 = 1e10
                            _ = v73
                        case -9223372036854775808:
                            var v74 bool = v65
                            var v75 uint8 = 127
                            var v76 string = "hello"
                            if (uint64(v64) == 18446744073709551615) {
                                v77 := GenS6{F0: v64}
                                if true {
                                    v4 -= v4
                                }
                                switch (v70 + int(v4)) {
                                case 4294967295:
                                    switch (("\n\t" + "\n\t") + ("foo bar" + "a")) {
                                    case "":
                                        v64 += 918
                                    default:
                                        v3 += (v3 / ((-1) | 1))
                                    }
                                case 0:
                                    switch fi68 {
                                    case 32767:
                                        v78 := GenS7{F0: v6}
                                        _ = v78
                                    case -266:
                                        _ = len(v5)
                                    case 702:
                                        v75 = v75
                                    }
                                    fmt.Println(v4, v70, fi71)
                                default:
                                    if false {
                                        v65 = (!false)
                                        v79 := GenS8{F0: int32(v75), F1: 548}
                                        _ = v79
                                    }
                                }
                                _ = v77
                            } else {
                                v80 := GenS9{F0: (v75 * v75), F1: (v3 / ((v3) | 1))}
                                _ = v80
                            }
                            _ = v74
                            _ = v75
                            _ = v76
                        }
                    }
                    v3 = v3
                }
                var v81 uint64 = 18446744073709551615
                var v82 uint = v2
                _ = fi71
                _ = v81
                _ = v82
            }
            _ = fi68
            _ = v70
        }
        var v83 uint32 = 2
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = v83
    }
    switch 128 {
    case 4294967295:
        if (9223372036854775807 >= v3) {
            var v84 uint32 = 2
            var v85 int = 9223372036854775807
            var v86 int8 = (v4 % ((v4) | 1))
            _ = v84
            _ = v85
            _ = v86
        } else {
            var v87 uint = (v2 | (v2 & (v2 / ((uint(v64)) | 1))))
            if len(v5) > 0 {
                v5[0] = 2
            }
            _ = v87
        }
    default:
        var v88 string = ""
        _ = v88
    }
    for fi89 := 0; fi89 < 15; fi89++ {
        for fi90 := 0; fi90 < 8; fi90++ {
            v4 = v4
            var v91 int32 = 459
            v64 -= 65535
            _ = fi90
            _ = v91
        }
        fmt.Println(v6, v2)
        _ = fi89
    }
    if ((-2.25 / 1.0) < 0.0) {
        for fi92 := 0; fi92 < 15; fi92++ {
            fmt.Println(v65)
            _ = fi92
        }
        for fi93 := 0; fi93 < 14; fi93++ {
            switch ("a" + "\n\t") {
            case "":
                switch ("" + "\n\t") {
                case "foo bar":
                    v94 := GenS6{F0: (v64 & 0)}
                    _ = v94
                case "\n\t":
                    v1 += v1
                }
            case "hello":
                var v95 int = fi93
                _ = v95
            default:
                var v96 int64 = (v3 / (((-(v3))) | 1))
                _ = v96
            }
            _ = fi93
        }
        v97 := GenS6{F0: v64}
        _ = v97
    } else {
        var v98 int16 = v6
        _ = v98
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v64:", v64)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v64
    _ = v65
}
