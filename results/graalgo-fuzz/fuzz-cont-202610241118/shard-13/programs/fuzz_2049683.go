package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint32; F2 uint }
type GenS2 struct { F0 int16; F1 uint8; F2 int }
type GenS3 struct { F0 int32; F1 uint16; F2 bool }
type GenS4 struct { F0 uint32; F1 int16 }
type GenS5 struct { F0 int32; F1 string; F2 int8 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 int16; F1 int32 }
type GenS8 struct { F0 uint8; F1 uint16; F2 uint64 }
type GenS9 struct { F0 int8; F1 uint32 }
type GenS10 struct { F0 int8; F1 uint }
type GenS11 struct { F0 uint8; F1 bool; F2 int8 }
type GenS12 struct { F0 bool; F1 uint16 }
type GenS13 struct { F0 uint16 }
type GenS14 struct { F0 int }
type GenS15 struct { F0 string }
type GenS16 struct { F0 uint32; F1 int16 }
type GenS17 struct { F0 bool }
type GenS18 struct { F0 bool; F1 bool }
type GenS19 struct { F0 uint64 }
type GenS20 struct { F0 int; F1 int8; F2 uint32 }
type GenS21 struct { F0 uint16; F1 uint32 }
type GenS22 struct { F0 string; F1 uint32; F2 int16 }
type GenS23 struct { F0 string; F1 uint8; F2 uint }
type GenS24 struct { F0 bool }
type GenS25 struct { F0 uint32; F1 int8; F2 int32 }
func genFn0(p0 uint, p1 string, p2 uint32) int8 {
    if ((0.1 - ((3.14159 + -2.25) / (0.0 - float32(3.14159)))) == -2.25) {
        switch 127 {
        case 2147483647:
            v4 := []int{-9223372036854775808, 9223372036854775807}
            _ = v4
        case 127:
            var v5 bool = false
            _ = v5
        case -603:
            for fi6 := 0; fi6 < 20; fi6++ {
                var v7 float32 = (1.5 + 1.5)
                p1 = p1
                for fi8 := 0; fi8 < 3; fi8++ {
                    var v9 uint16 = 255
                    _ = fi8
                    _ = v9
                }
                _ = fi6
                _ = v7
            }
        default:
            var v10 int16 = 830
            _ = v10
        }
    }
    if (!true) {
        p0 -= 42
    }
    v11 := GenS1{F0: "foo bar", F1: p2, F2: (^(p0))}
    switch "foo bar" {
    case "a":
        v12 := []int{4294967295, 552, -2147483648}
        _ = v12
    case "":
        var v13 string = "hello"
        _ = v13
    case "\n\t":
        for fi14 := 0; fi14 < 18; fi14++ {
            for fi15 := 0; fi15 < 10; fi15++ {
                fmt.Println(fi14)
                var v16 int64 = -1
                _ = fi15
                _ = v16
            }
            _ = fi14
        }
    }
    for fi17 := 0; fi17 < 15; fi17++ {
        p1 = p1
        _ = fi17
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 1
}

func main() {
    var v1 int8 = -1
    v2 := []string{"\n\t"}
    var v3 uint = 456
    for fi18 := 0; fi18 < 5; fi18++ {
        v19 := GenS2{F0: 32767, F1: 0, F2: (fi18 % ((fi18) | 1))}
        _ = fi18
        _ = v19
    }
    v3 = 4294967295
    if len(v2) > 0 {
        v2[0] = (("" + "a") + "foo bar")
    }
    if true {
        switch 47 {
        case 4294967295:
            switch 325 {
            case 42:
                switch 617 {
                case 2147483647:
                    switch "\n\t" {
                    case "a":
                        if false {
                            var v20 int32 = -2147483648
                            var v21 int16 = 32767
                            if len(v2) > 0 {
                                v2[0] = "a"
                            }
                            _ = v20
                            _ = v21
                        }
                    case "\n\t":
                        switch ("foo bar" + "\n\t") {
                        case "a":
                            v22 := GenS1{F0: "a", F1: 2147483647, F2: 255}
                            _ = v22
                        default:
                            for fi23 := 0; fi23 < 13; fi23++ {
                                fmt.Println(v3, fi23)
                                _ = fi23
                            }
                        }
                    default:
                        var v24 uint64 = 18446744073709551615
                        _ = v24
                    }
                case 946:
                    for fi25 := 0; fi25 < 20; fi25++ {
                        v1 += v1
                        switch fi25 {
                        case -2147483648:
                            for fi26 := 0; fi26 < 13; fi26++ {
                                fmt.Println(fi25, fi26)
                                _ = fi26
                            }
                        case 9223372036854775807:
                            v27 := GenS2{F0: -490, F1: 42, F2: (fi25 - fi25)}
                            _ = v27
                        case 128:
                            for fi28 := 0; fi28 < 13; fi28++ {
                                switch "hello" {
                                case "":
                                    var v29 uint64 = 1
                                    _ = v29
                                case "a":
                                    v30 := GenS3{F0: 0, F1: 1, F2: false}
                                    _ = v30
                                case "hello":
                                    var v31 uint16 = 0
                                    _ = v31
                                }
                                v3 = (v3 * v3)
                                if false {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    if true {
                                        if (!false) {
                                            if false {
                                                if len(v2) > 0 {
                                                    v2[0] = (("a" + ("foo bar" + ("\n\t" + "hello"))) + "hello")
                                                }
                                                switch (fi28 & (fi28 ^ (fi25 ^ (fi25 / (((^(fi25))) | 1))))) {
                                                case 2:
                                                    switch "" {
                                                    case "a":
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                    case "foo bar":
                                                        var v32 int16 = 215
                                                        _ = v32
                                                    }
                                                case -28:
                                                    var v33 int64 = -814
                                                    _ = v33
                                                case -117:
                                                    if false {
                                                        var v34 uint = 1
                                                        if false {
                                                            v35 := GenS3{F0: 822, F1: 42, F2: (true || ("hello" > "\n\t"))}
                                                            if len(v2) > 0 {
                                                                v2[0] = (("\n\t" + "\n\t") + "")
                                                            }
                                                            _ = v35
                                                        } else {
                                                            switch (-(fi28)) {
                                                            case 128:
                                                                var v36 string = "\n\t"
                                                                _ = v36
                                                            case 1:
                                                                var v37 uint8 = 42
                                                                _ = v37
                                                            }
                                                            var v38 int8 = v1
                                                            _ = v38
                                                        }
                                                        _ = v34
                                                    }
                                                default:
                                                    v2 = append(v2, ("hello" + "\n\t"))
                                                }
                                            } else {
                                                v39 := GenS3{F0: 1, F1: 65535, F2: false}
                                                _ = v39
                                            }
                                            var v40 int = -9223372036854775808
                                            _ = v40
                                        } else {
                                            v3 = v3
                                        }
                                        var v41 int16 = 32767
                                        _ = v41
                                    } else {
                                        v42 := GenS2{F0: 446, F1: 128, F2: fi28}
                                        v3 -= (v3 ^ (v3 & (^(v3))))
                                        v3 -= 0
                                        _ = v42
                                    }
                                    v3 = v3
                                    switch "\n\t" {
                                    case "":
                                        v1 = v1
                                    case "\n\t":
                                        if false {
                                            var v43 uint = 42
                                            if (511 > 2147483647) {
                                                if true {
                                                    v3 = 65535
                                                }
                                                switch (fi25 & (fi28 * 32767)) {
                                                case 42:
                                                    var v44 uint64 = 1
                                                    _ = v44
                                                case 1:
                                                    var v45 int8 = genFn0((v3 % (((v3 >> ((965) & 63))) | 1)), ("hello" + ""), 1)
                                                    _ = v45
                                                case 32767:
                                                    if ((((1e10 / float32((-2.25 * 1e10))) / float32(float64(1.5))) + 0.1) != 1e10) {
                                                        if (false || (((!((true && (!(0 < 42))) || true)) && (287 > 663)) && (2147483647 == 1))) {
                                                            switch (^(fi25)) {
                                                            case 32767:
                                                                v1 += (v1 / ((v1) | 1))
                                                            case 127:
                                                                fmt.Println(v43, v43, fi28)
                                                            case 42:
                                                                var v46 int64 = -9223372036854775808
                                                                _ = v46
                                                            default:
                                                                switch fi28 {
                                                                case 127:
                                                                    var v47 string = "a"
                                                                    _ = v47
                                                                case 9223372036854775807:
                                                                    v1 += v1
                                                                case -1:
                                                                    switch (fi28 + fi25) {
                                                                    case 674:
                                                                        fmt.Println(v1, v1)
                                                                    case -2147483648:
                                                                        v3 = v43
                                                                        if true {
                                                                            v48 := GenS4{F0: 2147483647, F1: 32767}
                                                                            v1 += v1
                                                                            _ = v48
                                                                        } else {
                                                                            v2 = append(v2, "\n\t")
                                                                        }
                                                                        fmt.Println(v1, fi28)
                                                                    case 1:
                                                                        if len(v2) >= 2 {
                                                                            v2 = v2[1:len(v2)]
                                                                        }
                                                                    }
                                                                default:
                                                                    v49 := GenS3{F0: 702, F1: 65535, F2: (fi25 == (fi25 - fi28))}
                                                                    _ = v49
                                                                }
                                                            }
                                                        } else {
                                                            var v50 uint8 = 42
                                                            _ = v50
                                                        }
                                                    }
                                                default:
                                                    var v51 int64 = 839
                                                    _ = v51
                                                }
                                                v3 += (v43 % (((v3 << (((v43 * (v3 | v43))) & 63))) | 1))
                                            } else {
                                                if true {
                                                    if ("a" >= "foo bar") {
                                                        v3 += (^(v43))
                                                        v2 = append(v2, "")
                                                    } else {
                                                        v3 += 42
                                                    }
                                                    var v52 string = ("a" + "a")
                                                    if len(v2) > 0 {
                                                        v2[0] = v52
                                                    }
                                                    _ = v52
                                                } else {
                                                    _ = len(v2)
                                                    v1 = (v1 % ((127) | 1))
                                                    var v53 float32 = ((3.14159 + -2.25) / (1.5 * ((-2.25 / (3.14159 - 0.1)) - 0.1)))
                                                    _ = v53
                                                }
                                            }
                                            var v54 int32 = 2147483647
                                            _ = v43
                                            _ = v54
                                        } else {
                                            if len(v2) > 0 {
                                                v2[0] = ("a" + ("a" + ""))
                                            }
                                        }
                                    default:
                                        v55 := GenS5{F0: 42, F1: (("" + ("" + (("" + (("\n\t" + "foo bar") + "\n\t")) + (("hello" + "hello") + ("foo bar" + ("foo bar" + "")))))) + "a"), F2: -128}
                                        _ = v55
                                    }
                                    if (((((0.0 - 1e10) * -2.25) * 1e10) / 1.5) >= (-2.25 - 0.1)) {
                                        var v56 int = -193
                                        var v57 float32 = 1.5
                                        var v58 bool = false
                                        if (439 >= 127) {
                                            fmt.Println(v56)
                                            var v59 uint32 = 4294967295
                                            _ = len(v2)
                                            _ = v59
                                        } else {
                                            v60 := GenS6{F0: 127}
                                            var v61 int32 = 343
                                            _ = v60
                                            _ = v61
                                        }
                                        _ = v56
                                        _ = v57
                                        _ = v58
                                    } else {
                                        var v62 uint64 = 4294967295
                                        var v63 string = ("hello" + "hello")
                                        switch "a" {
                                        case "hello":
                                            v62 += (v62 & v62)
                                            v62 = (^(v62))
                                            var v64 float64 = float64(-2.25)
                                            _ = v64
                                        default:
                                            v1 += -1
                                        }
                                        _ = v62
                                        _ = v63
                                    }
                                    if (!(((false && (-513 >= 42)) && true) && true)) {
                                        fmt.Println(fi28, fi25)
                                        var v65 int = fi25
                                        _ = v65
                                    } else {
                                        var v66 uint32 = 829
                                        switch -667 {
                                        case 42:
                                            v1 -= -125
                                        case 128:
                                            if (255 == 255) {
                                                fmt.Println(v66)
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                            } else {
                                                if ((v3 != 18446744073709551615) || false) {
                                                    v67 := GenS6{F0: 1}
                                                    v3 = v3
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    _ = v67
                                                }
                                                switch "foo bar" {
                                                case "":
                                                    var v68 uint8 = 1
                                                    _ = v68
                                                case "a":
                                                    if (351 == 433) {
                                                        fmt.Println(v3)
                                                    }
                                                case "foo bar":
                                                    if (42 != 188) {
                                                        var v69 int = (fi28 << ((763) & 63))
                                                        switch "" {
                                                        case "\n\t":
                                                            var v70 bool = true
                                                            _ = v70
                                                        case "a":
                                                            switch 32768 {
                                                            case 42:
                                                                var v71 float64 = float64(0.0)
                                                                _ = v71
                                                            default:
                                                                if true {
                                                                    if true {
                                                                        v72 := GenS3{F0: -1, F1: 738, F2: (false || (!true))}
                                                                        if true {
                                                                            fmt.Println(v69)
                                                                            v2 = append(v2, ("hello" + "hello"))
                                                                            if false {
                                                                                v69 -= (fi25 ^ fi25)
                                                                                if false {
                                                                                    v1 += (v1 ^ v1)
                                                                                    var v73 float64 = float64((3.14159 / (1.5 * (1.5 - 1e10))))
                                                                                    if (!true) {
                                                                                        v73 = (0.0 - (1.5 + v73))
                                                                                    } else {
                                                                                        if true {
                                                                                            if len(v2) > 0 {
                                                                                                v2[0] = ("\n\t" + "\n\t")
                                                                                            }
                                                                                        }
                                                                                        if (2147483647 > 554) {
                                                                                            _ = len(v2)
                                                                                            v74 := GenS7{F0: 32767, F1: 2147483647}
                                                                                            if len(v2) > 0 {
                                                                                                v2[0] = "\n\t"
                                                                                            }
                                                                                            _ = v74
                                                                                        }
                                                                                        v75 := GenS8{F0: 2, F1: 204, F2: 48}
                                                                                        _ = v75
                                                                                    }
                                                                                    var v76 uint8 = 210
                                                                                    _ = v73
                                                                                    _ = v76
                                                                                } else {
                                                                                    if ("a" == ("foo bar" + "foo bar")) {
                                                                                        v1 = genFn0((v3 % (((^(v3))) | 1)), "foo bar", v66)
                                                                                    }
                                                                                    v2 = append(v2, "")
                                                                                }
                                                                                if false {
                                                                                    v69 -= -456
                                                                                    fmt.Println(v66, fi28, v69)
                                                                                    fmt.Println(v69)
                                                                                    var v77 float32 = 1e10
                                                                                    _ = v77
                                                                                } else {
                                                                                    var v78 int16 = 1000
                                                                                    v69 -= 2147483647
                                                                                    v69 -= (-(fi25))
                                                                                    _ = v78
                                                                                }
                                                                                if true {
                                                                                    v79 := GenS9{F0: 79, F1: (v66 | (v66 << ((v3) & 63)))}
                                                                                    v80 := GenS10{F0: (v1 - 4), F1: (v3 % ((v3) | 1))}
                                                                                    _ = v79
                                                                                    _ = v80
                                                                                } else {
                                                                                    switch fi25 {
                                                                                    case 4294967295:
                                                                                        v81 := GenS11{F0: 26, F1: (v66 > (v66 * 65535)), F2: (v1 % (((v1 / (((v1 & (-(v1)))) | 1))) | 1))}
                                                                                        _ = v81
                                                                                    case 608:
                                                                                        switch "a" {
                                                                                        case "hello":
                                                                                            v82 := GenS7{F0: 63, F1: -1}
                                                                                            _ = v82
                                                                                        default:
                                                                                            fmt.Println(v69, fi28, v3)
                                                                                        }
                                                                                    }
                                                                                }
                                                                                v83 := GenS11{F0: 183, F1: ((true && true) || ((0.0 - 1.5) <= (-2.25 * (float32(((1e10 + 0.1) / (1e10 + (0.1 - 0.0)))) + -2.25)))), F2: genFn0(255, ("foo bar" + "a"), 2147483648)}
                                                                                _ = v83
                                                                            } else {
                                                                                v84 := GenS12{F0: ((-(v1)) >= v1), F1: 65535}
                                                                                _ = v84
                                                                            }
                                                                        }
                                                                        var v85 int16 = int16(v1)
                                                                        _ = v72
                                                                        _ = v85
                                                                    } else {
                                                                        v1 -= v1
                                                                        if len(v2) > 0 {
                                                                            v2[0] = "foo bar"
                                                                        }
                                                                        switch (fi28 + (v69 % ((-1) | 1))) {
                                                                        case -68:
                                                                            v1 -= (v1 >> ((uint(v66)) & 63))
                                                                            if len(v2) > 0 {
                                                                                v2[0] = ""
                                                                            }
                                                                            var v86 int32 = 0
                                                                            _ = v86
                                                                        case 32768:
                                                                            v2 = append(v2, ("\n\t" + "a"))
                                                                        case 0:
                                                                            if false {
                                                                                fmt.Println(fi28, v69, fi25)
                                                                            } else {
                                                                                var v87 float32 = (0.1 - (3.14159 - float32(0.1)))
                                                                                if (255 < 200) {
                                                                                    v88 := GenS13{F0: uint16(v3)}
                                                                                    fmt.Println(fi25)
                                                                                    _ = v88
                                                                                }
                                                                                _ = v87
                                                                            }
                                                                            v69 = (fi28 + fi25)
                                                                        }
                                                                        if false {
                                                                            v89 := GenS14{F0: (-(fi28))}
                                                                            v2 = append(v2, ("a" + "foo bar"))
                                                                            if (3.14159 > -2.25) {
                                                                                fmt.Println(fi28, fi28)
                                                                                fmt.Println(v69)
                                                                            }
                                                                            if len(v2) > 0 {
                                                                                v2[0] = ""
                                                                            }
                                                                            _ = v89
                                                                        } else {
                                                                            switch "foo bar" {
                                                                            case "":
                                                                                switch "foo bar" {
                                                                                case "\n\t":
                                                                                    v2 = append(v2, ("\n\t" + "a"))
                                                                                case "foo bar":
                                                                                    v90 := GenS15{F0: ("a" + "foo bar")}
                                                                                    _ = v90
                                                                                }
                                                                            case "a":
                                                                                var v91 uint8 = 2
                                                                                _ = v91
                                                                            case "hello":
                                                                                var v92 uint64 = 752
                                                                                _ = v92
                                                                            }
                                                                            v93 := GenS6{F0: 41}
                                                                            _ = v93
                                                                        }
                                                                    }
                                                                    if false {
                                                                        var v94 int64 = 9223372036854775807
                                                                        _ = v94
                                                                    } else {
                                                                        var v95 int32 = 323
                                                                        if (!(fi28 >= -696)) {
                                                                            _ = len(v2)
                                                                            if true {
                                                                                v96 := GenS16{F0: (v66 >> (((v3 << (((^(v3))) & 63))) & 63)), F1: 32767}
                                                                                v3 = (v3 * 595)
                                                                                _ = v96
                                                                            } else {
                                                                                _ = len(v2)
                                                                                if (9223372036854775807 >= 0) {
                                                                                    if len(v2) >= 2 {
                                                                                        v2 = v2[1:len(v2)]
                                                                                    }
                                                                                    var v97 int = int(v95)
                                                                                    var v98 float32 = ((((3.14159 - -2.25) / (-2.25 + 0.1)) / 0.1) / (-2.25 * 3.14159))
                                                                                    _ = v97
                                                                                    _ = v98
                                                                                } else {
                                                                                    var v99 int32 = (v95 % ((v95) | 1))
                                                                                    v100 := GenS17{F0: false}
                                                                                    v66 = v66
                                                                                    var v101 uint32 = v66
                                                                                    v69 += (v69 / (((fi25 - 1)) | 1))
                                                                                    if (317 < 42) {
                                                                                        var v102 int8 = 127
                                                                                        if true {
                                                                                            v101 = 1
                                                                                        }
                                                                                        if (42 == 123) {
                                                                                            var v103 uint16 = 32767
                                                                                            v103 += (v103 / (((^(v103))) | 1))
                                                                                            v99 = (v99 * (v95 + 921))
                                                                                            _ = v103
                                                                                        } else {
                                                                                            var v104 int8 = v1
                                                                                            var v105 int32 = v99
                                                                                            v106 := GenS7{F0: 42, F1: (v105 + v105)}
                                                                                            _ = v104
                                                                                            _ = v105
                                                                                            _ = v106
                                                                                        }
                                                                                        _ = v102
                                                                                    } else {
                                                                                        switch "foo bar" {
                                                                                        case "a":
                                                                                            if len(v2) >= 2 {
                                                                                                v2 = v2[1:len(v2)]
                                                                                            }
                                                                                        case "":
                                                                                            var v107 bool = false
                                                                                            _ = v107
                                                                                        }
                                                                                    }
                                                                                    _ = v100
                                                                                    _ = v101
                                                                                    _ = v99
                                                                                }
                                                                                switch "hello" {
                                                                                case "":
                                                                                    if (!true) {
                                                                                        v108 := GenS14{F0: (v69 / (((fi28 * 623)) | 1))}
                                                                                        v69 -= (v69 + fi28)
                                                                                        fmt.Println(v66, v1, fi28)
                                                                                        _ = v108
                                                                                    } else {
                                                                                        if ((231 == 2) && false) {
                                                                                            if ((-2.25 + (-2.25 + 1.5)) != -2.25) {
                                                                                                v69 = -1
                                                                                                v95 = v95
                                                                                                fmt.Println(v1, v69, fi25)
                                                                                            } else {
                                                                                                var v109 int16 = -32768
                                                                                                var v110 int8 = v1
                                                                                                _ = v109
                                                                                                _ = v110
                                                                                            }
                                                                                            if len(v2) > 0 {
                                                                                                v2[0] = (("foo bar" + "a") + "")
                                                                                            }
                                                                                            var v111 uint16 = uint16(v3)
                                                                                            switch ("" + "a") {
                                                                                            case "":
                                                                                                v3 = 4294967295
                                                                                            }
                                                                                            _ = v111
                                                                                        }
                                                                                        v66 = v66
                                                                                        v112 := GenS1{F0: "\n\t", F1: (v66 ^ 2), F2: v3}
                                                                                        if false {
                                                                                            if true {
                                                                                                _ = len(v2)
                                                                                                v95 += -2147483648
                                                                                                var v113 uint8 = 0
                                                                                                _ = v113
                                                                                            }
                                                                                        } else {
                                                                                            v3 = (v3 / ((uint(v66)) | 1))
                                                                                            if len(v2) > 0 {
                                                                                                v2[0] = "foo bar"
                                                                                            }
                                                                                            fmt.Println(v69, v1)
                                                                                        }
                                                                                        _ = v112
                                                                                    }
                                                                                case "a":
                                                                                    v1 += (-(v1))
                                                                                case "hello":
                                                                                    var v114 int16 = int16(v3)
                                                                                    _ = v114
                                                                                default:
                                                                                    var v115 uint32 = (^(v66))
                                                                                    _ = v115
                                                                                }
                                                                                fmt.Println(v66, v3, fi25)
                                                                                switch (^(v69)) {
                                                                                case 454:
                                                                                    fmt.Println(v1, fi28, v3)
                                                                                case 564:
                                                                                    if len(v2) >= 2 {
                                                                                        v2 = v2[1:len(v2)]
                                                                                    }
                                                                                }
                                                                                switch fi25 {
                                                                                case 4294967295:
                                                                                    var v116 int32 = v95
                                                                                    _ = v116
                                                                                case -9223372036854775808:
                                                                                    v1 -= (v1 * (v1 & genFn0(v3, "a", v66)))
                                                                                default:
                                                                                    if true {
                                                                                        var v117 int64 = -533
                                                                                        _ = v117
                                                                                    }
                                                                                }
                                                                                var v118 int32 = 955
                                                                                v118 = v118
                                                                                var v119 float32 = (3.14159 * -2.25)
                                                                                _ = v118
                                                                                _ = v119
                                                                            }
                                                                            if len(v2) >= 2 {
                                                                                v2 = v2[1:len(v2)]
                                                                            }
                                                                        } else {
                                                                            if len(v2) > 0 {
                                                                                v2[0] = ("foo bar" + "\n\t")
                                                                            }
                                                                            if len(v2) >= 2 {
                                                                                v2 = v2[1:len(v2)]
                                                                            }
                                                                            fmt.Println(v69, v69, fi25)
                                                                        }
                                                                        v69 += fi28
                                                                        _ = v95
                                                                    }
                                                                    fmt.Println(v3, fi25)
                                                                } else {
                                                                    switch fi28 {
                                                                    case -280:
                                                                        switch fi28 {
                                                                        case 4294967295:
                                                                            v120 := GenS2{F0: -74, F1: 255, F2: 0}
                                                                            _ = v120
                                                                        case 487:
                                                                            v69 -= v69
                                                                        }
                                                                    case -2147483648:
                                                                        if len(v2) >= 2 {
                                                                            v2 = v2[1:len(v2)]
                                                                        }
                                                                    case -766:
                                                                        if (false || (1.5 == 0.1)) {
                                                                            var v121 uint = 177
                                                                            _ = v121
                                                                        } else {
                                                                            if (-796 > 955) {
                                                                                var v122 int = fi28
                                                                                var v123 uint32 = v66
                                                                                fmt.Println(v123, v122)
                                                                                _ = v122
                                                                                _ = v123
                                                                            }
                                                                        }
                                                                        v124 := GenS18{F0: false, F1: (32767 >= 576)}
                                                                        _ = v124
                                                                    }
                                                                    var v125 uint64 = 352
                                                                    _ = v125
                                                                }
                                                            }
                                                        case "":
                                                            fmt.Println(fi28)
                                                        }
                                                        if ((!false) || (-1 < -945)) {
                                                            if (1 > (-(v1))) {
                                                                if (fi25 != v69) {
                                                                    if true {
                                                                        var v126 uint8 = 158
                                                                        fmt.Println(fi25)
                                                                        _ = v126
                                                                    }
                                                                    fmt.Println(v69, fi28)
                                                                    if (uint8(v66) > 1) {
                                                                        var v127 uint8 = 1
                                                                        switch (v69 >> ((v3) & 63)) {
                                                                        case 32768:
                                                                            v128 := GenS19{F0: 0}
                                                                            _ = v128
                                                                        }
                                                                        v127 = (v127 >> ((v3) & 63))
                                                                        _ = v127
                                                                    }
                                                                } else {
                                                                    if len(v2) > 0 {
                                                                        v2[0] = "\n\t"
                                                                    }
                                                                    v129 := GenS20{F0: -2147483648, F1: (v1 << ((uint(v66)) & 63)), F2: v66}
                                                                    _ = v129
                                                                }
                                                            } else {
                                                                if len(v2) > 0 {
                                                                    v2[0] = ("hello" + "\n\t")
                                                                }
                                                            }
                                                            var v130 uint32 = (v66 / ((v66) | 1))
                                                            if (!true) {
                                                                v1 = v1
                                                                switch "" {
                                                                case "":
                                                                    if (!(0 <= 32767)) {
                                                                        v1 -= (v1 & (v1 ^ v1))
                                                                    } else {
                                                                        v3 += v3
                                                                        v66 -= 2
                                                                    }
                                                                }
                                                            }
                                                            _ = v130
                                                        }
                                                        _ = v69
                                                    } else {
                                                        if true {
                                                            v1 = 42
                                                            _ = len(v2)
                                                        } else {
                                                            switch fi28 {
                                                            case -2147483648:
                                                                if false {
                                                                    v131 := GenS21{F0: 0, F1: 0}
                                                                    _ = v131
                                                                } else {
                                                                    v132 := GenS22{F0: "hello", F1: v66, F2: 344}
                                                                    if false {
                                                                        fmt.Println(v3, fi25, v3)
                                                                    }
                                                                    var v133 uint = uint(v66)
                                                                    _ = v132
                                                                    _ = v133
                                                                }
                                                            case 84:
                                                                v66 = (^(v66))
                                                            case -9223372036854775808:
                                                                if false {
                                                                    var v134 uint = (^(v3))
                                                                    _ = v134
                                                                } else {
                                                                    v1 += genFn0((^(v3)), "a", 2)
                                                                    var v135 float32 = 0.0
                                                                    var v136 uint = 0
                                                                    v66 = (v66 & v66)
                                                                    _ = v135
                                                                    _ = v136
                                                                }
                                                            default:
                                                                v137 := GenS23{F0: ("foo bar" + "\n\t"), F1: 1, F2: v3}
                                                                _ = v137
                                                            }
                                                        }
                                                    }
                                                default:
                                                    var v138 uint8 = 128
                                                    _ = v138
                                                }
                                                var v139 bool = ((true && (v3 > (v3 + 925))) || (0 > -1))
                                                _ = v139
                                            }
                                        default:
                                            v2 = append(v2, "foo bar")
                                        }
                                        _ = v66
                                    }
                                    _ = len(v2)
                                }
                                _ = fi28
                            }
                        default:
                            if false {
                                switch fi25 {
                                case 32768:
                                    for fi140 := 0; fi140 < 9; fi140++ {
                                        fmt.Println(fi25)
                                        _ = fi140
                                    }
                                }
                            } else {
                                v2 = append(v2, ("foo bar" + "foo bar"))
                                switch (fi25 | (fi25 << ((823) & 63))) {
                                case -2147483648:
                                    v1 = (v1 >> ((v3) & 63))
                                default:
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                }
                                v1 += (-(v1))
                            }
                        }
                        _ = fi25
                    }
                case 280:
                    switch 128 {
                    case -938:
                        var v141 uint = (v3 / ((605) | 1))
                        _ = v141
                    case 0:
                        var v142 int32 = 2147483647
                        _ = v142
                    default:
                        v3 = v3
                    }
                }
            case 9223372036854775807:
                if (0.1 == 1.5) {
                    for fi143 := 0; fi143 < 13; fi143++ {
                        switch fi143 {
                        case 127:
                            for fi144 := 0; fi144 < 15; fi144++ {
                                fmt.Println(fi143, fi143, v1)
                                _ = fi144
                            }
                        }
                        _ = fi143
                    }
                    var v145 uint = (v3 & v3)
                    _ = v145
                }
            case 2147483647:
                if ("foo bar" != "\n\t") {
                    var v146 bool = (true || (3 == 2))
                    _ = v146
                } else {
                    fmt.Println(v1, v3)
                    fmt.Println(v3)
                    v3 = v3
                }
            default:
                for fi147 := 0; fi147 < 16; fi147++ {
                    switch "foo bar" {
                    case "a":
                        switch "a" {
                        case "foo bar":
                            var v148 uint = (^(v3))
                            _ = v148
                        default:
                            _ = len(v2)
                        }
                    case "\n\t":
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    case "foo bar":
                        var v149 int64 = 0
                        _ = v149
                    }
                    var v150 int32 = 0
                    fmt.Println(fi147, v1, fi147)
                    _ = fi147
                    _ = v150
                }
            }
        default:
            fmt.Println(v3)
        }
        v3 = 4294967295
    } else {
        var v151 int32 = -1
        _ = v151
    }
    if false {
        switch 119 {
        case -2147483648:
            v152 := GenS24{F0: (32767 <= 806)}
            _ = v152
        case -1:
            v1 -= (v1 & v1)
        default:
            if false {
                for fi153 := 0; fi153 < 10; fi153++ {
                    v3 = (v3 - v3)
                    var v154 float32 = (float32(1e10) / 1.0)
                    _ = fi153
                    _ = v154
                }
                for fi155 := 0; fi155 < 12; fi155++ {
                    for fi156 := 0; fi156 < 7; fi156++ {
                        var v157 bool = false
                        switch (fi156 >> (((^(v3))) & 63)) {
                        case -2147483648:
                            fmt.Println(v1)
                        }
                        _ = fi156
                        _ = v157
                    }
                    _ = fi155
                }
                var v158 float64 = -2.25
                _ = v158
            } else {
                var v159 int = -1
                var v160 uint64 = 2
                _ = v159
                _ = v160
            }
        }
        v3 += v3
    }
    v161 := GenS25{F0: 0, F1: 42, F2: -1}
    v1 += (v1 ^ v1)
    if false {
        v162 := GenS22{F0: "hello", F1: 2147483647, F2: -160}
        v1 = v1
        _ = v162
    } else {
        fmt.Println(v3)
    }
    fmt.Println(v1, v3)
    fmt.Println(v1)
    v1 -= genFn0(2, (("\n\t" + ("foo bar" + "hello")) + "hello"), 2147483647)
    v163 := GenS12{F0: (!(0 == 9223372036854775807)), F1: 0}
    for fi164 := 0; fi164 < 16; fi164++ {
        switch (^(fi164)) {
        case -666:
            switch "\n\t" {
            case "foo bar":
                var v165 bool = ("hello" > "")
                _ = v165
            case "":
                for fi166 := 0; fi166 < 3; fi166++ {
                    var v167 uint = v3
                    _ = fi166
                    _ = v167
                }
            default:
                for fi168 := 0; fi168 < 1; fi168++ {
                    if (true && true) {
                        var v169 float64 = -2.25
                        _ = v169
                    }
                    v3 -= v3
                    _ = fi168
                }
            }
        default:
            v1 = 100
        }
        if true {
            v3 = 18446744073709551615
            var v170 bool = (!(((0.1 - (float64(0.0) - (0.0 - 0.0))) / 1.0) == (0.1 - -2.25)))
            _ = v170
        }
        _ = fi164
    }
    v171 := GenS20{F0: -554, F1: (-(v1)), F2: 1}
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v3 = v3
    for fi172 := 0; fi172 < 18; fi172++ {
        for fi173 := 0; fi173 < 3; fi173++ {
            if (962 < 0) {
                v174 := GenS22{F0: "a", F1: 2, F2: -278}
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                var v175 uint16 = 32767
                fmt.Println(fi173, v1, v3)
                _ = v174
                _ = v175
            }
            var v176 uint = (v3 >> ((4294967295) & 63))
            _ = fi173
            _ = v176
        }
        _ = fi172
    }
    var v177 float64 = 0.0
    fmt.Println("v1:", v1)
    fmt.Println("v177:", v177)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v161
    _ = v163
    _ = v171
    _ = v177
    _ = v2
    _ = v3
}
