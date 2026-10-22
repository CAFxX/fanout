package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 uint32; F1 int }
type GenS3 struct { F0 int64; F1 int32; F2 int }
type GenS4 struct { F0 int; F1 bool; F2 int16 }
type GenS5 struct { F0 uint32; F1 bool; F2 uint32 }
type GenS6 struct { F0 int16; F1 int32; F2 bool }
type GenS7 struct { F0 int8; F1 bool }
type GenS8 struct { F0 uint32; F1 bool; F2 uint16 }
func genFn0(p0 uint8, p1 uint) uint16 {
    for fi5 := 0; fi5 < 15; fi5++ {
        for fi6 := 0; fi6 < 1; fi6++ {
            v7 := GenS1{F0: -128, F1: -1}
            switch "a" {
            case "":
                var v8 int = 127
                _ = v8
            case "a":
                v9 := []int{4294967295}
                _ = v9
            case "foo bar":
                v10 := []int{fi5}
                _ = v10
            default:
                p1 = (p1 ^ (p1 + p1))
            }
            var v11 float64 = ((1e10 - 0.0) * 0.1)
            _ = fi6
            _ = v11
            _ = v7
        }
        _ = fi5
    }
    v12 := GenS1{F0: 100, F1: -375}
    p0 = p0
    for fi13 := 0; fi13 < 15; fi13++ {
        p1 += (p1 >> ((p1) & 63))
        _ = fi13
    }
    _ = p0
    _ = p1
    _ = v12
    return 42
}
func genFn1(p0 int8, p1 bool, p2 uint32) int32 {
    p1 = false
    if (p1 || (float32(3.14159) <= 0.1)) {
        for fi14 := 0; fi14 < 16; fi14++ {
            fmt.Println(p2, p0, p1)
            _ = fi14
        }
    }
    var v15 uint16 = 42
    for fi16 := 0; fi16 < 19; fi16++ {
        var v17 int32 = -11
        var v18 int64 = 0
        _ = fi16
        _ = v17
        _ = v18
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v15
    return 348
}

func main() {
    var v1 int8 = 100
    v2 := make(map[string]uint32)
    var v3 int8 = (v1 >> ((18446744073709551615) & 63))
    v4 := []uint32{2147483648, 510}
    if (true || (0 != 18446744073709551615)) {
        v3 -= (v1 % (((v3 | v3)) | 1))
    } else {
        var v19 uint32 = 84
        for fi20 := 0; fi20 < 1; fi20++ {
            v21 := v2["hello"]
            v21 = v21
            _ = fi20
            _ = v21
        }
        if (true || false) {
            switch (("foo bar" + "a") + "\n\t") {
            case "a":
                var v22 uint64 = 42
                _ = v22
            case "hello":
                _ = len(v2)
            case "\n\t":
                v3 = 98
            }
            _ = len(v4)
        }
        _ = v19
    }
    v1 = v3
    if ((0.1 * ((0.0 / 3.14159) - 1.5)) < ((1e10 * (float64(0.1) / float64((3.14159 / 3.14159)))) * 1e10)) {
        var v23 uint = 42
        _ = v23
    } else {
        for fi24 := 0; fi24 < 13; fi24++ {
            switch ("" + "") {
            case "foo bar":
                v25 := GenS1{F0: (v1 & (v1 ^ v3)), F1: 0}
                _ = v25
            case "a":
                if true {
                    for fi26 := 0; fi26 < 15; fi26++ {
                        var v27 int = (^(fi26))
                        _ = fi26
                        _ = v27
                    }
                    switch "hello" {
                    case "foo bar":
                        switch fi24 {
                        case -2147483648:
                            if false {
                                switch (fi24 & 584) {
                                case 595:
                                    switch 127 {
                                    case -2147483648:
                                        for fi28 := 0; fi28 < 14; fi28++ {
                                            if false {
                                                if true {
                                                    var v29 int8 = v3
                                                    fmt.Println(v3, fi28, fi24)
                                                    var v30 int32 = 171
                                                    _ = v29
                                                    _ = v30
                                                }
                                                var v31 string = ("a" + (("hello" + "a") + "a"))
                                                _ = v31
                                            }
                                            var v32 int64 = -1
                                            _ = fi28
                                            _ = v32
                                        }
                                    }
                                case -2147483648:
                                    for fi33 := 0; fi33 < 9; fi33++ {
                                        if true {
                                            fmt.Println(fi33)
                                            var v34 int32 = 0
                                            var v35 uint = 18446744073709551615
                                            _ = v34
                                            _ = v35
                                        }
                                        fmt.Println(fi33, v3)
                                        _ = fi33
                                    }
                                default:
                                    fmt.Println(v1, fi24, v3)
                                }
                            } else {
                                v36 := GenS1{F0: v3, F1: -9223372036854775808}
                                _ = v36
                            }
                        }
                    case "\n\t":
                        var v37 uint64 = 906
                        _ = v37
                    }
                } else {
                    for fi38 := 0; fi38 < 12; fi38++ {
                        var v39 uint16 = 42
                        _ = fi38
                        _ = v39
                    }
                    var v40 uint16 = 2
                    for fi41 := 0; fi41 < 14; fi41++ {
                        v2[("hello" + (("hello" + "") + (("foo bar" + "a") + ("a" + ("hello" + ("a" + "\n\t"))))))] = 65535
                        v42 := GenS2{F0: 2147483647, F1: int(v40)}
                        _ = fi41
                        _ = v42
                    }
                    _ = v40
                }
            case "":
                v3 -= (v1 & (v1 % ((104) | 1)))
            default:
                v43 := GenS2{F0: 2, F1: (fi24 - fi24)}
                _ = v43
            }
            for fi44 := 0; fi44 < 4; fi44++ {
                v3 = v1
                fmt.Println(v3)
                fmt.Println(v3, fi44, fi44)
                _ = fi44
            }
            _ = fi24
        }
    }
    switch 42 {
    case 42:
        fmt.Println(v3, v3)
    case 128:
        for fi45 := 0; fi45 < 18; fi45++ {
            var v46 int64 = 9223372036854775807
            v46 -= (v46 % ((0) | 1))
            _ = fi45
            _ = v46
        }
    case -819:
        v3 = -77
    default:
        if false {
            v3 -= 67
            var v47 bool = false
            switch "" {
            case "\n\t":
                switch -1 {
                case 42:
                    for fi48 := 0; fi48 < 15; fi48++ {
                        for fi49 := 0; fi49 < 20; fi49++ {
                            v50 := GenS2{F0: 1, F1: fi49}
                            v1 -= (v3 - (v1 >> ((2) & 63)))
                            if (-2.25 == 0.0) {
                                v51 := GenS3{F0: -9223372036854775808, F1: 1, F2: (fi48 << ((2) & 63))}
                                v2[(("\n\t" + (("" + "a") + ("" + "foo bar"))) + "a")] = 398
                                var v52 uint16 = 1
                                var v53 string = "a"
                                _ = v51
                                _ = v52
                                _ = v53
                            } else {
                                if (genFn1((v1 / ((0) | 1)), (v47 && true), 4294967295) != 617) {
                                    v54 := GenS4{F0: (fi49 << ((42) & 63)), F1: v47, F2: -321}
                                    v47 = v47
                                    _ = v54
                                }
                                fmt.Println(v3, fi49, fi49)
                                v55 := v2["hello"]
                                _ = v55
                            }
                            var v56 int8 = (^(v3))
                            _ = fi49
                            _ = v50
                            _ = v56
                        }
                        switch "foo bar" {
                        case "hello":
                            v47 = (v47 && v47)
                        case "\n\t":
                            fmt.Println(v47, v3, v3)
                        default:
                            v1 += 127
                        }
                        _ = fi48
                    }
                case -825:
                    fmt.Println(v47, v47, v3)
                case 32767:
                    switch "hello" {
                    case "a":
                        if (349 > 329) {
                            v47 = v47
                        } else {
                            if (42 < 992) {
                                if (-9223372036854775808 == -185) {
                                    fmt.Println(v1, v47)
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    var v57 float32 = 0.0
                                    _ = v57
                                }
                                v47 = (true && false)
                                if false {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    if v47 {
                                        v58 := GenS5{F0: 65535, F1: (167 <= 32767), F2: 739}
                                        v47 = false
                                        _ = len(v4)
                                        _ = v58
                                    }
                                } else {
                                    fmt.Println(v3, v1, v47)
                                }
                            } else {
                                var v59 float64 = 1.5
                                var v60 uint16 = 959
                                _ = v59
                                _ = v60
                            }
                            v61 := GenS2{F0: 614, F1: 884}
                            if ((442 <= -2147483648) && (v47 && v47)) {
                                v1 = (v1 / ((v3) | 1))
                                fmt.Println(v47)
                                switch -1 {
                                case 776:
                                    var v62 int = 406
                                    _ = v62
                                case -2147483648:
                                    for fi63 := 0; fi63 < 8; fi63++ {
                                        for fi64 := 0; fi64 < 17; fi64++ {
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            v1 = 119
                                            _ = fi64
                                        }
                                        _ = fi63
                                    }
                                case -9223372036854775808:
                                    v47 = v47
                                default:
                                    var v65 uint = 18446744073709551615
                                    _ = v65
                                }
                            }
                            _ = v61
                        }
                    case "":
                        _ = len(v2)
                    }
                default:
                    if false {
                        var v66 uint = 952
                        v3 += (v1 & v1)
                        v1 = (v3 - -80)
                        _ = v66
                    }
                }
            case "hello":
                _ = len(v4)
            case "":
                delete(v2, ("foo bar" + ("foo bar" + "hello")))
            default:
                v67 := GenS1{F0: -107, F1: 0}
                _ = v67
            }
            _ = v47
        } else {
            v1 = v3
        }
    }
    v3 = (v3 << ((477) & 63))
    fmt.Println(v1, v3)
    var v68 uint = 2
    var v69 bool = ((1e10 / (1e10 * float32(-2.25))) > ((float32((0.0 + 1.5)) + -2.25) * (((1e10 - 1.5) / 1.0) + float32(((((3.14159 - -2.25) / float64(1e10)) * 1e10) / (3.14159 + -2.25))))))
    v70 := GenS3{F0: 1, F1: 42, F2: 141}
    var v71 int32 = 2147483647
    switch (("" + ("\n\t" + "a")) + (("" + ("a" + "")) + "")) {
    case "":
        switch ((((("a" + "hello") + ("" + ("a" + "\n\t"))) + "foo bar") + "\n\t") + "\n\t") {
        case "a":
            v72 := GenS5{F0: 2, F1: (969 < 315), F2: 44}
            _ = v72
        case "foo bar":
            var v73 int32 = (v71 / ((v71) | 1))
            _ = v73
        case "\n\t":
            for fi74 := 0; fi74 < 17; fi74++ {
                v3 = (v1 + v1)
                _ = fi74
            }
        default:
            var v75 uint64 = 537
            _ = v75
        }
    case "a":
        v76 := GenS3{F0: 470, F1: -1, F2: -1}
        _ = v76
    case "\n\t":
        if (-1 > 475) {
            var v77 float64 = (float64(-2.25) / 1.5)
            if (!(!v69)) {
                if v69 {
                    switch "a" {
                    case "\n\t":
                        v69 = false
                    case "foo bar":
                        var v78 uint = v68
                        _ = v78
                    default:
                        if v69 {
                            v2["\n\t"] = 2147483647
                            switch "a" {
                            case "hello":
                                v79 := GenS2{F0: 208, F1: 127}
                                _ = v79
                            default:
                                var v80 float64 = 1e10
                                _ = v80
                            }
                        }
                    }
                    var v81 int8 = (v1 << ((v68) & 63))
                    _ = v81
                }
            } else {
                var v82 uint16 = 204
                switch -412 {
                case 501:
                    switch "a" {
                    case "\n\t":
                        switch ("\n\t" + "hello") {
                        case "foo bar":
                            if v69 {
                                var v83 int = -395
                                _ = v83
                            } else {
                                var v84 uint8 = uint8(v82)
                                v3 = v1
                                v1 = v1
                                _ = v84
                            }
                        case "hello":
                            fmt.Println(v71, v68, v77)
                        case "a":
                            v85 := GenS2{F0: 65535, F1: -9223372036854775808}
                            _ = v85
                        }
                    case "":
                        if ((v71 - 1) > v71) {
                            for fi86 := 0; fi86 < 13; fi86++ {
                                v87 := GenS4{F0: (fi86 - (fi86 ^ (fi86 / ((486) | 1)))), F1: (42 != (v82 * 2)), F2: 902}
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                                _ = fi86
                                _ = v87
                            }
                            v68 -= (v68 | (v68 << (((v68 & 42)) & 63)))
                            if (v3 != v1) {
                                if v69 {
                                    fmt.Println(v3, v68, v3)
                                    v82 = v82
                                } else {
                                    v77 = ((float64(-2.25) - (3.14159 + (float64(float32((v77 / 1.5))) - 1.5))) + float64(-2.25))
                                }
                                var v88 uint = 255
                                _ = v88
                            } else {
                                _ = len(v4)
                                for fi89 := 0; fi89 < 16; fi89++ {
                                    v77 -= v77
                                    _ = fi89
                                }
                            }
                        } else {
                            if (-1 <= 1) {
                                v69 = v69
                            } else {
                                for fi90 := 0; fi90 < 14; fi90++ {
                                    if (!((v82 >> ((v68) & 63)) == 375)) {
                                        var v91 float64 = v77
                                        _ = v91
                                    }
                                    if ((3.14159 >= -2.25) && v69) {
                                        for fi92 := 0; fi92 < 15; fi92++ {
                                            v71 += (v71 & v71)
                                            var v93 string = "foo bar"
                                            _ = fi92
                                            _ = v93
                                        }
                                        v94 := GenS6{F0: 1, F1: v71, F2: false}
                                        _ = v94
                                    }
                                    _ = fi90
                                }
                            }
                            var v95 int64 = -792
                            _ = v95
                        }
                    case "a":
                        if true {
                            v71 += genFn1(0, v69, 65535)
                        }
                    }
                default:
                    v96 := GenS7{F0: (v1 * 1), F1: v69}
                    _ = v96
                }
                _ = v82
            }
            v97 := GenS5{F0: uint32(v68), F1: ((1.5 != (-2.25 - float32(v77))) && ((90 < 42) && (!(2147483647 >= (v71 << ((v68) & 63)))))), F2: 0}
            _ = v77
            _ = v97
        }
    default:
        v98 := GenS8{F0: 4294967295, F1: (32767 >= -786), F2: genFn0(42, v68)}
        _ = v98
    }
    fmt.Println(v69, v3)
    var v99 uint32 = 2
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v100 := GenS3{F0: -1, F1: v71, F2: -615}
    var v101 int8 = 127
    var v102 int8 = v101
    v102 = 1
    fmt.Println("v1:", v1)
    fmt.Println("v101:", v101)
    fmt.Println("v102:", v102)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v68:", v68)
    fmt.Println("v69:", v69)
    fmt.Println("v71:", v71)
    fmt.Println("v99:", v99)
    _ = v1
    _ = v100
    _ = v101
    _ = v102
    _ = v2
    _ = v3
    _ = v4
    _ = v68
    _ = v69
    _ = v70
    _ = v71
    _ = v99
}
