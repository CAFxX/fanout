package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16; F1 int16; F2 int32 }
type GenS3 struct { F0 int8; F1 int8; F2 uint64 }
type GenS4 struct { F0 uint64; F1 uint32 }
type GenS5 struct { F0 uint8 }
type GenS6 struct { F0 uint8; F1 uint64; F2 bool }
type GenS7 struct { F0 int8; F1 bool; F2 bool }
type GenS8 struct { F0 uint32; F1 int64 }
type GenS9 struct { F0 uint16; F1 int; F2 int }
type GenS10 struct { F0 bool; F1 uint8; F2 bool }
type GenS11 struct { F0 int16; F1 uint16; F2 uint64 }
type GenS12 struct { F0 int; F1 uint; F2 string }
type GenS13 struct { F0 uint16; F1 uint8; F2 bool }
type GenS14 struct { F0 string }
type GenS15 struct { F0 int64 }
type GenS16 struct { F0 uint; F1 uint; F2 uint32 }
type GenS17 struct { F0 uint }
type GenS18 struct { F0 uint32; F1 int }
type GenS19 struct { F0 string; F1 bool }
type GenS20 struct { F0 uint64 }
type GenS21 struct { F0 uint16; F1 uint16 }
func genFn0(p0 int32) int {
    fmt.Println(p0)
    fmt.Println(p0)
    v5 := GenS1{F0: 2}
    fmt.Println(p0)
    _ = p0
    _ = v5
    return 9223372036854775807
}

func main() {
    var v1 bool = false
    v2 := make(map[string]string)
    v3 := []int{127, -392, 42}
    var v4 int16 = 42
    switch 128 {
    case 745:
        v2["a"] = ("" + "\n\t")
    case 9223372036854775807:
        v3 = append(v3, 0)
    }
    v6 := GenS2{F0: -1, F1: -678, F2: -856}
    var v7 uint8 = 1
    fmt.Println(v1, v7)
    v4 = 1
    for fi8 := 0; fi8 < 3; fi8++ {
        var v9 int64 = -486
        _ = fi8
        _ = v9
    }
    if v1 {
        _ = len(v3)
        switch "\n\t" {
        case "a":
            var v10 uint = 409
            _ = v10
        case "foo bar":
            var v11 int8 = 127
            _ = v11
        case "\n\t":
            v12 := v2[("foo bar" + ((((("\n\t" + "a") + "hello") + "foo bar") + "hello") + "hello"))]
            _ = v12
        default:
            v3 = append(v3, -14)
        }
        var v13 bool = (2 <= 1)
        _ = v13
    }
    _ = len(v2)
    if v1 {
        v14 := GenS1{F0: 320}
        if (65535 <= 353) {
            v4 -= v4
            for fi15 := 0; fi15 < 20; fi15++ {
                v7 -= v7
                _ = fi15
            }
            var v16 uint = 65535
            _ = v16
        }
        if (-2.25 != float32(float64(1e10))) {
            switch "hello" {
            case "foo bar":
                v2[("\n\t" + "hello")] = "a"
            case "hello":
                v7 += (v7 & (v7 - 255))
            case "":
                for fi17 := 0; fi17 < 16; fi17++ {
                    if (26 != (fi17 % ((fi17) | 1))) {
                        if (!(!v1)) {
                            for fi18 := 0; fi18 < 18; fi18++ {
                                if v1 {
                                    var v19 uint = 42
                                    if v1 {
                                        v4 += (v4 >> (((v19 >> ((v19) & 63))) & 63))
                                        _ = len(v3)
                                        v4 -= 0
                                    }
                                    _ = v19
                                }
                                v7 = 238
                                _ = fi18
                            }
                            delete(v2, ("" + "\n\t"))
                        }
                        var v20 bool = v1
                        _ = v20
                    } else {
                        if v1 {
                            var v21 uint = 2
                            fmt.Println(v1)
                            _ = v21
                        } else {
                            delete(v2, ("" + "\n\t"))
                        }
                        if v1 {
                            var v22 int32 = -2147483648
                            _ = v22
                        }
                        for fi23 := 0; fi23 < 4; fi23++ {
                            switch ("foo bar" + ("" + "\n\t")) {
                            case "hello":
                                var v24 uint16 = 812
                                _ = v24
                            case "foo bar":
                                fmt.Println(v4)
                            default:
                                if (v1 && v1) {
                                    var v25 int64 = -1
                                    _ = v25
                                } else {
                                    v26 := GenS1{F0: 32767}
                                    _ = v26
                                }
                            }
                            if v1 {
                                var v27 uint64 = 32
                                var v28 uint = 4294967295
                                _ = v27
                                _ = v28
                            }
                            if true {
                                v29 := GenS3{F0: 47, F1: -1, F2: 42}
                                if (11 >= 65535) {
                                    var v30 int8 = 1
                                    v7 = v7
                                    _ = v30
                                }
                                if false {
                                    if v1 {
                                        if false {
                                            v31 := GenS4{F0: 4294967295, F1: 368}
                                            switch ("" + "foo bar") {
                                            case "hello":
                                                v4 = v4
                                            case "a":
                                                v4 = (v4 % ((v4) | 1))
                                            default:
                                                v32 := GenS5{F0: (v7 % (((^(v7))) | 1))}
                                                _ = v32
                                            }
                                            var v33 int64 = 9223372036854775807
                                            _ = v31
                                            _ = v33
                                        } else {
                                            switch 42 {
                                            case 0:
                                                var v34 uint = 65535
                                                _ = v34
                                            case 32767:
                                                var v35 int64 = 244
                                                _ = v35
                                            }
                                            v36 := GenS4{F0: 420, F1: 65535}
                                            fmt.Println(v1, v4, fi23)
                                            _ = v36
                                        }
                                        var v37 uint64 = 19
                                        v38 := GenS6{F0: uint8(v37), F1: (v37 ^ (^(v37))), F2: (v4 < v4)}
                                        _ = v37
                                        _ = v38
                                    } else {
                                        v39 := GenS7{F0: 33, F1: ((fi17 << ((65535) & 63)) == (fi17 << ((381) & 63))), F2: (fi17 != (fi17 >> ((42) & 63)))}
                                        _ = v39
                                    }
                                } else {
                                    var v40 float64 = -2.25
                                    _ = v40
                                }
                                v4 -= v4
                                if ((v7 << ((231) & 63)) >= (v7 * (^(v7)))) {
                                    if (4294967295 > 42) {
                                        _ = len(v3)
                                        switch fi17 {
                                        case -115:
                                            _ = len(v2)
                                        case 42:
                                            v41 := GenS8{F0: 1, F1: 9223372036854775807}
                                            _ = v41
                                        }
                                    } else {
                                        v42 := GenS2{F0: (v4 << ((uint(v7)) & 63)), F1: (v4 * (v4 >> ((0) & 63))), F2: -1}
                                        v7 -= v7
                                        if (true || (!false)) {
                                            var v43 bool = (((((("a" + ("" + "foo bar")) + "a") + ("a" + ("\n\t" + "a"))) + "hello") == (("a" + ("hello" + "\n\t")) + ("hello" + ("\n\t" + "a")))) || v1)
                                            if (false || v43) {
                                                v44 := GenS8{F0: 4294967295, F1: 9223372036854775807}
                                                _ = v44
                                            } else {
                                                fmt.Println(v43, v43, v4)
                                            }
                                            var v45 uint64 = 18446744073709551615
                                            _ = v43
                                            _ = v45
                                        } else {
                                            if v1 {
                                                if false {
                                                    if (65535 < 0) {
                                                        fmt.Println(fi17, v7)
                                                    }
                                                    switch (("\n\t" + "foo bar") + "foo bar") {
                                                    case "\n\t":
                                                        fmt.Println(fi23)
                                                    case "":
                                                        fmt.Println(v1, v1, fi17)
                                                    default:
                                                        v7 = 22
                                                    }
                                                    fmt.Println(v4)
                                                    var v46 int64 = -1
                                                    v4 += v4
                                                    _ = v46
                                                } else {
                                                    if len(v3) >= 2 {
                                                        v3 = v3[1:len(v3)]
                                                    }
                                                    if true {
                                                        v47 := GenS1{F0: 1}
                                                        v48 := GenS9{F0: 32767, F1: 4294967295, F2: 127}
                                                        _ = v47
                                                        _ = v48
                                                    }
                                                }
                                            } else {
                                                var v49 int16 = -886
                                                if (32767 != 2) {
                                                    if len(v3) > 0 {
                                                        v3[0] = fi17
                                                    }
                                                }
                                                if v1 {
                                                    fmt.Println(fi17, fi23, v1)
                                                }
                                                _ = v49
                                            }
                                            v50 := GenS2{F0: (v4 % ((v4) | 1)), F1: v4, F2: 2147483647}
                                            if (false && (((^(v7)) >= (v7 + (v7 ^ v7))) && v1)) {
                                                delete(v2, "")
                                                v1 = (!v1)
                                                switch "" {
                                                case "foo bar":
                                                    v51 := GenS10{F0: (genFn0(-1) < fi23), F1: v7, F2: (!(v1 || true))}
                                                    _ = v51
                                                case "hello":
                                                    switch "" {
                                                    case "foo bar":
                                                        v4 -= v4
                                                    case "":
                                                        switch "a" {
                                                        case "a":
                                                            fmt.Println(v4, fi17, v1)
                                                        case "\n\t":
                                                            v1 = (fi23 > (-(fi17)))
                                                        case "foo bar":
                                                            var v52 int32 = 2147483647
                                                            _ = v52
                                                        default:
                                                            var v53 int64 = 537
                                                            _ = v53
                                                        }
                                                    case "\n\t":
                                                        v7 = (v7 << ((255) & 63))
                                                    default:
                                                        switch (fi17 >> ((255) & 63)) {
                                                        case -2147483648:
                                                            var v54 int = 9223372036854775807
                                                            _ = v54
                                                        }
                                                    }
                                                case "\n\t":
                                                    v4 -= (v4 - (v4 + -453))
                                                }
                                                switch (^(fi23)) {
                                                case 127:
                                                    v1 = false
                                                case 0:
                                                    fmt.Println(fi17, fi17)
                                                case -2147483648:
                                                    if len(v3) > 0 {
                                                        v3[0] = (fi17 & genFn0(-1))
                                                    }
                                                }
                                            }
                                            _ = v50
                                        }
                                        _ = v42
                                    }
                                    v55 := v2["\n\t"]
                                    v3 = append(v3, (^(fi17)))
                                    v56 := GenS3{F0: -1, F1: 126, F2: 703}
                                    fmt.Println(v7, v4)
                                    v57 := GenS11{F0: 1, F1: 947, F2: 0}
                                    var v58 bool = true
                                    v55 = v55
                                    _ = len(v2)
                                    _ = v55
                                    _ = v56
                                    _ = v57
                                    _ = v58
                                }
                                v2[(("hello" + "a") + ("" + "\n\t"))] = ("" + "")
                                _ = v29
                            } else {
                                if (v1 && (false && (!v1))) {
                                    fmt.Println(v7)
                                    v1 = v1
                                } else {
                                    switch "\n\t" {
                                    case "\n\t":
                                        if false {
                                            v59 := GenS11{F0: v4, F1: 1, F2: 0}
                                            _ = v59
                                        } else {
                                            var v60 int64 = -1
                                            _ = v60
                                        }
                                    case "":
                                        if true {
                                            v4 = int16(fi17)
                                            switch "hello" {
                                            case "foo bar":
                                                var v61 int16 = v4
                                                _ = v61
                                            case "a":
                                                v7 = v7
                                            case "\n\t":
                                                v62 := GenS2{F0: 1, F1: v4, F2: 0}
                                                _ = v62
                                            }
                                            fmt.Println(v4, fi23, v4)
                                            fmt.Println(v1, fi23)
                                            v1 = (uint32(v7) >= 407)
                                        }
                                    }
                                }
                            }
                            _ = fi23
                        }
                    }
                    v63 := GenS12{F0: (fi17 / (((fi17 & fi17)) | 1)), F1: 1, F2: ""}
                    _ = fi17
                    _ = v63
                }
            default:
                v4 = v4
            }
            var v64 int64 = 541
            _ = v64
        }
        _ = v14
    } else {
        switch "\n\t" {
        case "foo bar":
            for fi65 := 0; fi65 < 2; fi65++ {
                v3 = append(v3, genFn0(0))
                fmt.Println(v4)
                if len(v3) >= 2 {
                    v3 = v3[1:len(v3)]
                }
                _ = fi65
            }
        case "":
            v7 = v7
        }
        var v66 uint = 0
        _ = v66
    }
    _ = len(v2)
    var v67 string = (((("a" + ("a" + (("" + ("a" + "a")) + "\n\t"))) + ("a" + ("\n\t" + "foo bar"))) + "a") + "hello")
    for fi68 := 0; fi68 < 10; fi68++ {
        for fi69 := 0; fi69 < 13; fi69++ {
            var v70 uint = 1
            v71 := GenS13{F0: 304, F1: 255, F2: v1}
            if ((!((100 >= -1) && true)) && (((-2.25 + (0.0 * (3.14159 / (0.1 / 3.14159)))) - 3.14159) < 1.5)) {
                if (!v1) {
                    var v72 uint32 = 51
                    _ = v72
                }
                delete(v2, v67)
            }
            _ = fi69
            _ = v70
            _ = v71
        }
        switch genFn0(-989) {
        case 580:
            for fi73 := 0; fi73 < 15; fi73++ {
                switch "hello" {
                case "a":
                    if (-223 <= 2147483647) {
                        _ = len(v2)
                        fmt.Println(v67, fi73, v1)
                        v3 = append(v3, 42)
                    }
                default:
                    v74 := GenS7{F0: 111, F1: true, F2: v1}
                    var v75 int32 = 248
                    if v1 {
                        var v76 uint = 18446744073709551615
                        if (v1 && (v1 || (0.0 < ((1.5 - 3.14159) - 1e10)))) {
                            var v77 string = "hello"
                            _ = v77
                        } else {
                            var v78 string = (v67 + (v67 + ("hello" + v67)))
                            if (v67 <= v78) {
                                if false {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    if (-128 >= -1) {
                                        switch fi68 {
                                        case 0:
                                            v79 := GenS8{F0: 0, F1: 0}
                                            _ = v79
                                        case 32767:
                                            v75 = v75
                                        default:
                                            v67 = ("" + v78)
                                        }
                                    } else {
                                        switch (fi73 & (fi73 % (((fi73 * (fi68 / ((fi68) | 1)))) | 1))) {
                                        case 9223372036854775807:
                                            if v1 {
                                                v4 = v4
                                            } else {
                                                switch int(v76) {
                                                case -412:
                                                    var v80 int16 = v4
                                                    _ = v80
                                                case -9223372036854775808:
                                                    v3 = append(v3, genFn0(42))
                                                case 977:
                                                    v81 := GenS14{F0: (("" + v78) + "a")}
                                                    _ = v81
                                                default:
                                                    v82 := GenS15{F0: -9223372036854775808}
                                                    _ = v82
                                                }
                                                var v83 uint16 = 1
                                                var v84 uint64 = 2
                                                _ = v83
                                                _ = v84
                                            }
                                        case 917:
                                            var v85 float64 = 1.5
                                            fmt.Println(v85, v7)
                                            if v1 {
                                                _ = len(v3)
                                                v86 := GenS1{F0: 255}
                                                v76 += (v76 & (v76 - (v76 + (v76 | (v76 - v76)))))
                                                _ = v86
                                            }
                                            fmt.Println(fi68, v75)
                                            v3 = append(v3, (fi73 & fi73))
                                            _ = v85
                                        }
                                        var v87 uint32 = 0
                                        if (!true) {
                                            switch (^(fi73)) {
                                            case -755:
                                                switch v78 {
                                                case "foo bar":
                                                    v78 = "foo bar"
                                                    v3 = append(v3, (^(fi68)))
                                                case "a":
                                                    if (!(true || (!(v87 <= (v87 >> ((993) & 63)))))) {
                                                        var v88 uint64 = 610
                                                        v89 := GenS9{F0: 525, F1: fi73, F2: 0}
                                                        var v90 int64 = -315
                                                        _ = v88
                                                        _ = v89
                                                        _ = v90
                                                    }
                                                }
                                            default:
                                                var v91 int16 = (^(v4))
                                                v92 := GenS14{F0: ""}
                                                switch genFn0((v75 / ((v75) | 1))) {
                                                case 841:
                                                    v93 := GenS4{F0: 1, F1: (v87 * 2147483647)}
                                                    _ = v93
                                                }
                                                fmt.Println(v78)
                                                v2["foo bar"] = v67
                                                _ = v91
                                                _ = v92
                                            }
                                        }
                                        v94 := GenS16{F0: (v76 - v76), F1: v76, F2: (v87 | 2147483647)}
                                        v3 = append(v3, (fi73 >> ((4294967295) & 63)))
                                        v95 := v2[(v67 + ((v67 + ("foo bar" + (v78 + (v67 + v67)))) + ("\n\t" + (v78 + "foo bar"))))]
                                        _ = v87
                                        _ = v94
                                        _ = v95
                                    }
                                }
                                _ = len(v3)
                            } else {
                                switch "a" {
                                case "\n\t":
                                    v4 = (v4 & v4)
                                case "a":
                                    v96 := v2[v67]
                                    _ = v96
                                case "foo bar":
                                    switch ((v78 + "") + v78) {
                                    case "\n\t":
                                        switch ("\n\t" + v78) {
                                        case "":
                                            if true {
                                                var v97 int8 = -28
                                                _ = v97
                                            }
                                        case "a":
                                            switch ("\n\t" + ("hello" + ((v67 + "a") + "a"))) {
                                            case "\n\t":
                                                v98 := GenS17{F0: 255}
                                                _ = v98
                                            case "a":
                                                _ = len(v3)
                                            case "":
                                                switch "hello" {
                                                case "a":
                                                    if (0.1 >= 1.5) {
                                                        v99 := GenS18{F0: 192, F1: fi73}
                                                        _ = v99
                                                    } else {
                                                        v100 := GenS13{F0: 834, F1: v7, F2: ((v4 ^ (v4 / ((v4) | 1))) > (v4 & (v4 & (v4 >> (((v76 * v76)) & 63)))))}
                                                        v78 = "hello"
                                                        var v101 string = "hello"
                                                        _ = v100
                                                        _ = v101
                                                    }
                                                default:
                                                    v102 := GenS13{F0: 0, F1: v7, F2: (1 < 1)}
                                                    _ = v102
                                                }
                                            default:
                                                if true {
                                                    _ = len(v2)
                                                } else {
                                                    delete(v2, "")
                                                }
                                            }
                                        case "foo bar":
                                            if len(v3) > 0 {
                                                v3[0] = 192
                                            }
                                        default:
                                            var v103 uint32 = 42
                                            if (false && (170 > genFn0(-1))) {
                                                v104 := GenS19{F0: "hello", F1: (false || v1)}
                                                var v105 uint = v76
                                                if ((!(false || v1)) || v1) {
                                                    var v106 float64 = float64(0.1)
                                                    if (0 > uint16(v103)) {
                                                        v7 += 255
                                                    }
                                                    _ = v106
                                                } else {
                                                    switch v78 {
                                                    case "\n\t":
                                                        v4 = v4
                                                        switch v78 {
                                                        case "hello":
                                                            var v107 uint16 = 32767
                                                            _ = v107
                                                        default:
                                                            v4 += (v4 << (((v76 - v105)) & 63))
                                                        }
                                                    case "foo bar":
                                                        switch (fi73 * genFn0(2147483647)) {
                                                        case 4294967295:
                                                            v4 = (v4 << ((1) & 63))
                                                        case 678:
                                                            var v108 uint = 0
                                                            _ = v108
                                                        case 150:
                                                            fmt.Println(v103)
                                                        default:
                                                            if v1 {
                                                                var v109 int32 = (v75 / ((v75) | 1))
                                                                v105 = 18446744073709551615
                                                                fmt.Println(v1, v76)
                                                                _ = v109
                                                            }
                                                        }
                                                    }
                                                    _ = len(v3)
                                                    _ = len(v3)
                                                    if v1 {
                                                        var v110 int64 = 0
                                                        _ = v110
                                                    } else {
                                                        v111 := GenS20{F0: 0}
                                                        _ = v111
                                                    }
                                                }
                                                _ = v104
                                                _ = v105
                                            }
                                            _ = v103
                                        }
                                    case "foo bar":
                                        v112 := GenS21{F0: 0, F1: 65535}
                                        _ = v112
                                    default:
                                        v78 = ""
                                    }
                                }
                            }
                            if v1 {
                                if (!(174 <= (v7 << (((v76 | (^(v76)))) & 63)))) {
                                    fmt.Println(v67, fi73)
                                }
                                v113 := GenS20{F0: 18446744073709551615}
                                _ = v113
                            } else {
                                if len(v3) > 0 {
                                    v3[0] = fi68
                                }
                            }
                            _ = v78
                        }
                        _ = v76
                    } else {
                        var v114 int = int(v7)
                        _ = v114
                    }
                    _ = v74
                    _ = v75
                }
                _ = fi73
            }
        case 32767:
            _ = len(v2)
        case -9223372036854775808:
            v67 = "foo bar"
        }
        _ = fi68
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v67:", v67)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v6
    _ = v67
    _ = v7
}
