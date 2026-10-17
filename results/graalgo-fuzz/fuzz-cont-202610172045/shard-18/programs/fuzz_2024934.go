package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint16 }
type GenS2 struct { F0 bool; F1 uint8; F2 int16 }
type GenS3 struct { F0 string; F1 int; F2 int64 }
type GenS4 struct { F0 int; F1 uint64 }
type GenS5 struct { F0 uint16; F1 uint8 }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 int64; F1 uint32 }
type GenS8 struct { F0 int64 }
type GenS9 struct { F0 uint8; F1 int16; F2 int64 }
type GenS10 struct { F0 int64 }
type GenS11 struct { F0 int64 }
func genFn0(p0 int16, p1 uint32, p2 string) int {
    fmt.Println(p0)
    var v7 string = "hello"
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 486
}

func main() {
    v1 := []string{("a" + ("\n\t" + "a")), ("a" + "foo bar"), "a", ""}
    var v2 uint64 = 910
    var v3 uint8 = 127
    v4 := make(map[int]string)
    v5 := []uint32{265, 1, 103}
    v6 := make(map[string]string)
    v2 -= v2
    v8 := GenS1{F0: 2147483647, F1: 1}
    v3 = 128
    if (!true) {
        var v9 uint8 = 127
        v3 = v9
        v2 += v2
        _ = v9
    } else {
        v2 += (v2 << ((18446744073709551615) & 63))
        var v10 uint64 = v2
        _ = v10
    }
    for fi11 := 0; fi11 < 20; fi11++ {
        v3 = 0
        switch "foo bar" {
        case "":
            v1 = append(v1, ("a" + (("hello" + "hello") + ("foo bar" + "\n\t"))))
        }
        _ = fi11
    }
    var v12 int16 = 93
    for fi13 := 0; fi13 < 10; fi13++ {
        var v14 int8 = 23
        _ = fi13
        _ = v14
    }
    v15 := GenS1{F0: 369, F1: 255}
    if true {
        v12 = (v12 << ((0) & 63))
    }
    for fi16 := 0; fi16 < 2; fi16++ {
        if true {
            for fi17 := 0; fi17 < 5; fi17++ {
                v2 = v2
                _ = fi17
            }
            for fi18 := 0; fi18 < 16; fi18++ {
                v3 = 205
                var v19 uint = 752
                _ = fi18
                _ = v19
            }
        } else {
            var v20 uint16 = 32767
            if len(v1) >= 2 {
                v1 = v1[1:len(v1)]
            }
            if (933 != 2147483648) {
                v21 := v4[genFn0(v12, 2, "hello")]
                var v22 int64 = -1
                var v23 int16 = v12
                _ = v21
                _ = v22
                _ = v23
            } else {
                fmt.Println(v3, v12, fi16)
                var v24 uint = 255
                fmt.Println(v24, fi16)
                _ = v24
            }
            _ = v20
        }
        fmt.Println(v12, v2)
        if (!(false || (17 > 2))) {
            fmt.Println(v2, v2)
            fmt.Println(fi16, v12, v12)
        }
        _ = fi16
    }
    for fi25 := 0; fi25 < 2; fi25++ {
        for fi26 := 0; fi26 < 12; fi26++ {
            var v27 bool = ("" <= "hello")
            v28 := GenS1{F0: (^(fi26)), F1: 42}
            _ = fi26
            _ = v27
            _ = v28
        }
        _ = fi25
    }
    var v29 uint16 = 42
    v2 -= (^(v2))
    for fi30 := 0; fi30 < 14; fi30++ {
        v2 = 42
        var v31 int32 = -1
        switch (fi30 * (fi30 | genFn0(int16(v31), 648, "foo bar"))) {
        case 32767:
            v3 -= (v3 >> ((1) & 63))
        case -2147483648:
            if true {
                v31 -= -2147483648
                v12 += (v12 - v12)
                var v32 int8 = 1
                _ = v32
            } else {
                if (53 == 2147483648) {
                    switch 1 {
                    case -854:
                        var v33 uint32 = 4294967295
                        _ = v33
                    case 1:
                        var v34 uint8 = v3
                        _ = v34
                    case 103:
                        switch fi30 {
                        case 9223372036854775807:
                            for fi35 := 0; fi35 < 4; fi35++ {
                                switch fi35 {
                                case -2147483648:
                                    switch (-(fi30)) {
                                    case 128:
                                        switch ((("a" + "a") + ("a" + "")) + ((("hello" + "hello") + "foo bar") + "hello")) {
                                        case "":
                                            var v36 uint16 = (v29 + 779)
                                            _ = v36
                                        case "\n\t":
                                            v37 := GenS2{F0: (true || (!(116 == 100))), F1: 1, F2: (-(v12))}
                                            _ = v37
                                        default:
                                            v3 = 46
                                        }
                                    case -791:
                                        if (9223372036854775807 == 0) {
                                            v38 := v4[fi30]
                                            v39 := GenS2{F0: (-119 == int8(v3)), F1: v3, F2: (-(v12))}
                                            v40 := GenS2{F0: false, F1: v3, F2: (v12 * -36)}
                                            _ = v38
                                            _ = v39
                                            _ = v40
                                        } else {
                                            switch -314 {
                                            case -1:
                                                v41 := GenS1{F0: (fi35 % ((-965) | 1)), F1: v29}
                                                _ = v41
                                            default:
                                                v31 = (v31 + v31)
                                            }
                                        }
                                        v3 -= 251
                                    }
                                default:
                                    var v42 uint8 = v3
                                    _ = v42
                                }
                                _ = fi35
                            }
                        case 0:
                            for fi43 := 0; fi43 < 14; fi43++ {
                                v44 := GenS2{F0: (!false), F1: (v3 % ((255) | 1)), F2: (v12 * (v12 - 1))}
                                _ = fi43
                                _ = v44
                            }
                        default:
                            var v45 bool = false
                            _ = v45
                        }
                    default:
                        v3 += v3
                    }
                } else {
                    var v46 int16 = (v12 + v12)
                    if (fi30 != (^(fi30))) {
                        if false {
                            var v47 int8 = 87
                            switch (^(fi30)) {
                            case -2147483648:
                                v12 += v12
                            case 126:
                                _ = len(v5)
                            case 128:
                                if (!(!true)) {
                                    v3 = v3
                                }
                            default:
                                v29 += uint16(v2)
                            }
                            _ = v47
                        }
                        var v48 int = (-(fi30))
                        if ((!true) || true) {
                            fmt.Println(v12)
                            if len(v1) > 0 {
                                v1[0] = "hello"
                            }
                        } else {
                            v12 += (v12 >> ((18446744073709551615) & 63))
                            v49 := GenS2{F0: true, F1: v3, F2: v12}
                            var v50 int64 = -1
                            _ = v49
                            _ = v50
                        }
                        _ = v48
                    }
                    v3 = (v3 << ((0) & 63))
                    _ = v46
                }
                for fi51 := 0; fi51 < 1; fi51++ {
                    v52 := GenS3{F0: "a", F1: 268, F2: -1}
                    v53 := GenS4{F0: fi51, F1: v2}
                    fmt.Println(v3, v29, v12)
                    _ = fi51
                    _ = v52
                    _ = v53
                }
                switch (^(fi30)) {
                case -9223372036854775808:
                    switch (^(fi30)) {
                    case 1:
                        v2 = 236
                    case -213:
                        if true {
                            if (1.5 >= 0.0) {
                                var v54 int8 = 1
                                v55 := GenS5{F0: v29, F1: v3}
                                _ = len(v6)
                                _ = v54
                                _ = v55
                            }
                            v56 := GenS2{F0: false, F1: (^(v3)), F2: -32768}
                            var v57 float32 = 1e10
                            _ = v56
                            _ = v57
                        } else {
                            for fi58 := 0; fi58 < 20; fi58++ {
                                v31 -= 2147483647
                                _ = fi58
                            }
                            var v59 uint8 = (v3 | uint8(v29))
                            v59 = 117
                            _ = v59
                        }
                    }
                case 128:
                    if false {
                        v60 := GenS6{F0: -57}
                        _ = v60
                    } else {
                        v31 = (v31 << ((2) & 63))
                        switch ("\n\t" + "\n\t") {
                        case "\n\t":
                            v31 = (v31 / ((v31) | 1))
                        case "a":
                            for fi61 := 0; fi61 < 6; fi61++ {
                                if (!(false || (10 >= 43))) {
                                    if (255 <= uint16(v3)) {
                                        fmt.Println(v29)
                                    }
                                } else {
                                    switch "foo bar" {
                                    case "":
                                        switch (("a" + "hello") + "\n\t") {
                                        case "\n\t":
                                            v62 := v6["hello"]
                                            _ = v62
                                        default:
                                            v3 = v3
                                        }
                                    case "foo bar":
                                        switch "foo bar" {
                                        case "hello":
                                            if false {
                                                v2 = v2
                                                if (v31 != (v31 + v31)) {
                                                    v63 := GenS2{F0: true, F1: v3, F2: (^(v12))}
                                                    if false {
                                                        var v64 string = "\n\t"
                                                        _ = v64
                                                    } else {
                                                        v2 = (v2 | 486)
                                                        fmt.Println(v29)
                                                        var v65 float32 = (3.14159 / 1.0)
                                                        _ = v65
                                                    }
                                                    if true {
                                                        var v66 uint64 = (v2 | (v2 >> ((2) & 63)))
                                                        var v67 int16 = 1
                                                        _ = v66
                                                        _ = v67
                                                    } else {
                                                        if false {
                                                            if (fi30 != 127) {
                                                                v68 := GenS7{F0: 0, F1: 65535}
                                                                fmt.Println(v3, v2, v2)
                                                                _ = v68
                                                            } else {
                                                                var v69 uint16 = (v29 / ((2) | 1))
                                                                v2 = (v2 % ((v2) | 1))
                                                                _ = v69
                                                            }
                                                            fmt.Println(v12, v31)
                                                            var v70 int = genFn0(v12, 42, "")
                                                            var v71 int = (fi61 / (((fi61 | -1)) | 1))
                                                            v3 += (^(v3))
                                                            if (!(true && (127 != -11))) {
                                                                v72 := GenS8{F0: 0}
                                                                _ = v72
                                                            }
                                                            v12 += v12
                                                            switch ("a" + "a") {
                                                            case "foo bar":
                                                                v2 += (v2 ^ (v2 ^ v2))
                                                            default:
                                                                var v73 float64 = (0.1 / 1.5)
                                                                _ = v73
                                                            }
                                                            fmt.Println(v12)
                                                            switch "\n\t" {
                                                            case "hello":
                                                                v74 := GenS9{F0: uint8(v2), F1: v12, F2: 42}
                                                                _ = v74
                                                            case "\n\t":
                                                                if true {
                                                                    var v75 int8 = -66
                                                                    v4[4294967295] = "foo bar"
                                                                    _ = v75
                                                                }
                                                            case "foo bar":
                                                                fmt.Println(v70)
                                                            default:
                                                                v3 += (v3 / ((v3) | 1))
                                                            }
                                                            _ = v70
                                                            _ = v71
                                                        }
                                                    }
                                                    _ = v63
                                                } else {
                                                    _ = len(v6)
                                                }
                                            }
                                        case "foo bar":
                                            if (((^(v2)) == 502) || ((v12 != (v12 % (((v12 / ((448) | 1))) | 1))) || (fi61 <= 32767))) {
                                                if false {
                                                    var v76 uint8 = v3
                                                    _ = v76
                                                } else {
                                                    if false {
                                                        _ = len(v4)
                                                        var v77 bool = (!false)
                                                        switch -2147483648 {
                                                        case 4294967295:
                                                            v1 = append(v1, ("a" + "foo bar"))
                                                        case 944:
                                                            v3 -= 97
                                                        case -976:
                                                            var v78 uint = 42
                                                            _ = v78
                                                        }
                                                        _ = v77
                                                    } else {
                                                        if false {
                                                            v3 = (v3 * v3)
                                                        } else {
                                                            v12 += v12
                                                            var v79 int8 = -75
                                                            v3 += 255
                                                            _ = v79
                                                        }
                                                        v80 := GenS4{F0: (^(fi61)), F1: 42}
                                                        _ = v80
                                                    }
                                                    v31 -= v31
                                                    var v81 int32 = 1
                                                    fmt.Println(fi30, fi30, fi30)
                                                    _ = v81
                                                }
                                            }
                                        default:
                                            v82 := GenS5{F0: (v29 / ((0) | 1)), F1: (^(v3))}
                                            _ = v82
                                        }
                                    default:
                                        v83 := GenS7{F0: -1, F1: 955}
                                        _ = v83
                                    }
                                }
                                if false {
                                    if (!(((0.0 - 3.14159) != 1e10) || false)) {
                                        fmt.Println(v31, v31)
                                    }
                                } else {
                                    v84 := GenS10{F0: 0}
                                    switch ("\n\t" + "hello") {
                                    case "":
                                        var v85 uint16 = 648
                                        _ = v85
                                    default:
                                        var v86 int64 = 9223372036854775807
                                        _ = v86
                                    }
                                    v87 := GenS4{F0: genFn0(1000, 2, "\n\t"), F1: (v2 | (v2 | v2))}
                                    switch 0 {
                                    case 1:
                                        var v88 int16 = (^(v12))
                                        _ = v88
                                    case 4294967295:
                                        if (true || false) {
                                            var v89 bool = (fi30 >= genFn0(v12, 42, ("" + "\n\t")))
                                            if (-1 < 0) {
                                                v89 = true
                                                v6[("foo bar" + "hello")] = ((("hello" + ("hello" + "hello")) + "foo bar") + "\n\t")
                                            } else {
                                                fmt.Println(v12, v89, v89)
                                                switch (("hello" + ("foo bar" + "a")) + "hello") {
                                                case "":
                                                    if v89 {
                                                        v90 := GenS1{F0: (fi61 * 0), F1: v29}
                                                        switch (fi30 + 562) {
                                                        case 42:
                                                            switch (fi61 << ((65535) & 63)) {
                                                            case 32767:
                                                                v12 = v12
                                                            case 509:
                                                                var v91 string = (("" + "\n\t") + "\n\t")
                                                                _ = v91
                                                            default:
                                                                switch 32768 {
                                                                case -2147483648:
                                                                    if (!(uint8(v29) <= 1)) {
                                                                        v29 += 42
                                                                    }
                                                                default:
                                                                    v92 := v6[(("a" + "") + "hello")]
                                                                    _ = v92
                                                                }
                                                            }
                                                        case 1:
                                                            var v93 bool = ((v3 * v3) > (v3 ^ 108))
                                                            _ = v93
                                                        }
                                                        _ = v90
                                                    }
                                                case "a":
                                                    _ = len(v5)
                                                case "hello":
                                                    v2 = v2
                                                default:
                                                    fmt.Println(v31)
                                                }
                                            }
                                            var v94 int = genFn0(67, 2147483647, ("a" + ((((("hello" + "") + "foo bar") + "") + "foo bar") + "a")))
                                            _ = v89
                                            _ = v94
                                        }
                                    default:
                                        v95 := GenS11{F0: -1}
                                        _ = v95
                                    }
                                    var v96 int32 = (v31 - v31)
                                    _ = v84
                                    _ = v87
                                    _ = v96
                                }
                                if false {
                                    delete(v4, int(v31))
                                    v31 -= v31
                                    v12 = v12
                                }
                                _ = fi61
                            }
                        default:
                            _ = len(v1)
                        }
                    }
                case 891:
                    for fi97 := 0; fi97 < 6; fi97++ {
                        var v98 uint64 = v2
                        _ = fi97
                        _ = v98
                    }
                default:
                    fmt.Println(v2)
                }
            }
        default:
            switch "hello" {
            case "hello":
                v31 = v31
            case "a":
                v31 -= 2147483647
            case "":
                var v99 uint8 = 2
                _ = v99
            }
        }
        _ = fi30
        _ = v31
    }
    v100 := GenS2{F0: true, F1: v3, F2: (v12 % ((32767) | 1))}
    v101 := v4[1]
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v101:", v101)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v100
    _ = v101
    _ = v12
    _ = v15
    _ = v2
    _ = v29
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
