package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint64; F1 uint8 }
type GenS4 struct { F0 uint32; F1 uint; F2 string }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint32; F1 string; F2 uint32 }
type GenS8 struct { F0 int32; F1 uint64 }
type GenS9 struct { F0 int16; F1 bool }
type GenS10 struct { F0 uint }
type GenS11 struct { F0 uint64; F1 int8; F2 int64 }
type GenS12 struct { F0 uint; F1 int64; F2 int }
type GenS13 struct { F0 uint64 }
func genFn0(p0 int16, p1 string, p2 int64) int32 {
    fmt.Println(p1)
    fmt.Println(p1, p1)
    var v5 bool = (!false)
    v6 := []int{-2147483648, 127}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 748
}
func genFn1() uint64 {
    var v7 uint64 = 2
    fmt.Println(v7)
    var v8 uint = 709
    _ = v7
    _ = v8
    return (v7 ^ uint64(v8))
}

func main() {
    var v1 uint32 = 2147483648
    var v2 uint = 42
    v3 := []string{}
    var v4 uint32 = (^(v1))
    v1 += (v1 + v4)
    v4 = (v4 / ((v1) | 1))
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    if true {
        v1 = (v1 + (^(v4)))
        v9 := GenS1{F0: 32}
        _ = v9
    }
    v10 := GenS1{F0: genFn0(42, "", -1)}
    if (!true) {
        for fi11 := 0; fi11 < 17; fi11++ {
            var v12 int8 = 42
            for fi13 := 0; fi13 < 15; fi13++ {
                fmt.Println(v12, v12)
                _ = fi13
            }
            _ = fi11
            _ = v12
        }
    } else {
        if false {
            if false {
                if (true && (!(true && ((^(v1)) < v1)))) {
                    v3 = append(v3, (("foo bar" + "hello") + "foo bar"))
                    if ((!true) || false) {
                        v14 := GenS2{F0: (("a" + "") + ("\n\t" + "a"))}
                        var v15 uint8 = 128
                        _ = v14
                        _ = v15
                    } else {
                        var v16 string = ("a" + "\n\t")
                        _ = v16
                    }
                }
                v4 += (v4 >> (((v2 % (((v2 | v2)) | 1))) & 63))
            } else {
                if (4294967295 > 435) {
                    var v17 int32 = genFn0(1000, "\n\t", -9223372036854775808)
                    var v18 uint = 2
                    if true {
                        fmt.Println(v17, v1)
                        v4 += v1
                    } else {
                        if len(v3) > 0 {
                            v3[0] = "hello"
                        }
                    }
                    _ = v17
                    _ = v18
                } else {
                    var v19 uint32 = (v1 & v1)
                    v20 := GenS2{F0: "a"}
                    _ = v19
                    _ = v20
                }
            }
            v21 := GenS2{F0: "hello"}
            v22 := GenS1{F0: genFn0(42, "", -1)}
            _ = v21
            _ = v22
        }
        v23 := GenS1{F0: 0}
        var v24 int = 128
        _ = v23
        _ = v24
    }
    var v25 uint16 = 68
    if true {
        if false {
            fmt.Println(v4, v4, v25)
            _ = len(v3)
            for fi26 := 0; fi26 < 14; fi26++ {
                switch "\n\t" {
                case "foo bar":
                    var v27 uint64 = 551
                    _ = v27
                case "":
                    switch ("\n\t" + "") {
                    case "":
                        v3 = append(v3, "")
                    case "foo bar":
                        v3 = append(v3, "hello")
                    case "\n\t":
                        var v28 int16 = -1
                        _ = v28
                    default:
                        switch "foo bar" {
                        case "a":
                            var v29 int8 = 100
                            _ = v29
                        case "hello":
                            v30 := GenS1{F0: 42}
                            _ = v30
                        case "foo bar":
                            fmt.Println(v2)
                        default:
                            var v31 bool = (1000 < 1)
                            _ = v31
                        }
                    }
                }
                var v32 uint = 2
                _ = fi26
                _ = v32
            }
        } else {
            if (32767 >= 952) {
                for fi33 := 0; fi33 < 14; fi33++ {
                    v34 := GenS3{F0: uint64(v4), F1: 127}
                    if true {
                        var v35 uint8 = 0
                        _ = v35
                    } else {
                        for fi36 := 0; fi36 < 9; fi36++ {
                            if true {
                                var v37 bool = true
                                if (212 > 724) {
                                    fmt.Println(v37)
                                    _ = len(v3)
                                    switch 128 {
                                    case -421:
                                        fmt.Println(v2)
                                    default:
                                        switch fi33 {
                                        case 525:
                                            if (fi36 != fi33) {
                                                switch fi33 {
                                                case 42:
                                                    if v37 {
                                                        var v38 int32 = 644
                                                        _ = v38
                                                    }
                                                    switch int(v1) {
                                                    case 32768:
                                                        v25 = (^(v25))
                                                    }
                                                    v37 = (false || ("hello" <= "a"))
                                                case 2147483647:
                                                    if len(v3) > 0 {
                                                        v3[0] = "a"
                                                    }
                                                }
                                                var v39 float64 = -2.25
                                                fmt.Println(v4, v25)
                                                _ = v39
                                            } else {
                                                v3 = append(v3, "a")
                                            }
                                        case 32767:
                                            if len(v3) > 0 {
                                                v3[0] = "hello"
                                            }
                                        default:
                                            fmt.Println(v2, fi36)
                                        }
                                    }
                                } else {
                                    v25 += v25
                                    v40 := GenS4{F0: v1, F1: (v2 | (v2 % ((v2) | 1))), F2: ""}
                                    v2 += v2
                                    v3 = append(v3, "a")
                                    var v41 uint16 = v25
                                    switch -2147483648 {
                                    case 42:
                                        if len(v3) > 0 {
                                            v3[0] = "foo bar"
                                        }
                                    }
                                    if false {
                                        var v42 int = fi33
                                        if (v37 && true) {
                                            var v43 int16 = -1
                                            var v44 uint16 = 122
                                            _ = v43
                                            _ = v44
                                        }
                                        v45 := GenS1{F0: genFn0(1, "foo bar", 0)}
                                        _ = v42
                                        _ = v45
                                    } else {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    }
                                    if v37 {
                                        var v46 float32 = 0.1
                                        fmt.Println(v37, v2)
                                        fmt.Println(v1, v1, v37)
                                        _ = v46
                                    }
                                    _ = v40
                                    _ = v41
                                }
                                v47 := GenS5{F0: fi36}
                                _ = v37
                                _ = v47
                            }
                            switch (^(fi33)) {
                            case 4294967295:
                                if true {
                                    var v48 float32 = 0.0
                                    switch (fi36 << (((v2 << (((v2 >> ((v2) & 63))) & 63))) & 63)) {
                                    case 32767:
                                        v49 := GenS4{F0: v1, F1: v2, F2: ("foo bar" + "hello")}
                                        _ = v49
                                    }
                                    _ = v48
                                } else {
                                    if true {
                                        _ = len(v3)
                                    }
                                    v50 := GenS3{F0: genFn1(), F1: 42}
                                    v1 -= v4
                                    var v51 uint32 = v4
                                    var v52 float64 = float64(1.5)
                                    _ = v50
                                    _ = v51
                                    _ = v52
                                }
                            default:
                                switch "\n\t" {
                                case "\n\t":
                                    v4 -= 42
                                case "":
                                    if len(v3) > 0 {
                                        v3[0] = ("a" + "a")
                                    }
                                case "hello":
                                    var v53 float32 = -2.25
                                    _ = v53
                                default:
                                    fmt.Println(fi36)
                                    switch (((("" + "foo bar") + "") + "a") + "") {
                                    case "\n\t":
                                        if true {
                                            var v54 int64 = 1
                                            v55 := GenS4{F0: (v4 % ((0) | 1)), F1: (v2 | uint(v1)), F2: "a"}
                                            switch 32768 {
                                            case -2147483648:
                                                v2 = (v2 >> ((v2) & 63))
                                            case -1:
                                                var v56 int32 = genFn0(32767, (("\n\t" + ("\n\t" + "foo bar")) + "a"), -1)
                                                _ = v56
                                            default:
                                                switch ((("foo bar" + "") + ("" + "a")) + "a") {
                                                case "a":
                                                    if (v2 == 42) {
                                                        v57 := GenS3{F0: genFn1(), F1: 0}
                                                        _ = v57
                                                    }
                                                case "hello":
                                                    if false {
                                                        v58 := GenS1{F0: genFn0(1, "\n\t", int64(v2))}
                                                        if false {
                                                            if ((!((v1 * v4) > v1)) || false) {
                                                                fmt.Println(v54, fi33)
                                                            }
                                                        }
                                                        var v59 int8 = 100
                                                        _ = v58
                                                        _ = v59
                                                    } else {
                                                        var v60 bool = (1.5 <= ((-2.25 + -2.25) * 1.5))
                                                        _ = len(v3)
                                                        if (0 < 42) {
                                                            if v60 {
                                                                v61 := GenS2{F0: ""}
                                                                var v62 int64 = (v54 % ((v54) | 1))
                                                                _ = v61
                                                                _ = v62
                                                            } else {
                                                                v63 := GenS3{F0: 0, F1: 20}
                                                                if (65535 >= v25) {
                                                                    _ = len(v3)
                                                                    if len(v3) > 0 {
                                                                        v3[0] = "a"
                                                                    }
                                                                } else {
                                                                    var v64 int16 = 565
                                                                    _ = v64
                                                                }
                                                                switch "\n\t" {
                                                                case "a":
                                                                    if (-1 < -1) {
                                                                        v65 := GenS6{F0: v25}
                                                                        _ = v65
                                                                    } else {
                                                                        if len(v3) > 0 {
                                                                            v3[0] = "\n\t"
                                                                        }
                                                                        var v66 string = ("a" + ("" + "\n\t"))
                                                                        var v67 uint8 = 127
                                                                        _ = v66
                                                                        _ = v67
                                                                    }
                                                                }
                                                                _ = v63
                                                            }
                                                        }
                                                        _ = v60
                                                    }
                                                case "foo bar":
                                                    if true {
                                                        v68 := GenS7{F0: (v1 * (v1 ^ 42)), F1: ("hello" + "foo bar"), F2: v1}
                                                        switch -9223372036854775808 {
                                                        case 127:
                                                            var v69 int8 = 66
                                                            _ = v69
                                                        case -970:
                                                            var v70 int = int(v25)
                                                            _ = v70
                                                        case 1:
                                                            var v71 string = (("hello" + ("" + "hello")) + "hello")
                                                            _ = v71
                                                        default:
                                                            if false {
                                                                var v72 int16 = -32768
                                                                _ = len(v3)
                                                                if true {
                                                                    v73 := GenS8{F0: 864, F1: 306}
                                                                    v1 -= v4
                                                                    _ = v73
                                                                } else {
                                                                    fmt.Println(fi36, v2, fi33)
                                                                    v1 = 943
                                                                }
                                                                _ = v72
                                                            } else {
                                                                var v74 int8 = -1
                                                                v54 += (v54 - (v54 ^ (v54 >> ((v2) & 63))))
                                                                if len(v3) > 0 {
                                                                    v3[0] = "foo bar"
                                                                }
                                                                _ = v74
                                                            }
                                                        }
                                                        _ = v68
                                                    } else {
                                                        var v75 int32 = 1
                                                        switch "\n\t" {
                                                        case "foo bar":
                                                            v2 += (v2 >> ((65535) & 63))
                                                        case "hello":
                                                            var v76 int64 = (v54 & v54)
                                                            _ = v76
                                                        case "\n\t":
                                                            v77 := GenS2{F0: ""}
                                                            _ = v77
                                                        default:
                                                            if false {
                                                                v4 = (v4 << (((v2 % ((4294967295) | 1))) & 63))
                                                                var v78 uint8 = 128
                                                                _ = v78
                                                            }
                                                        }
                                                        _ = v75
                                                    }
                                                }
                                            }
                                            v4 = v4
                                            if (0.0 < ((1e10 - 3.14159) / float32(3.14159))) {
                                                var v79 float64 = float64(((1.5 * (1.5 + 1e10)) / 1.0))
                                                _ = v79
                                            }
                                            _ = v54
                                            _ = v55
                                        } else {
                                            switch fi36 {
                                            case 32767:
                                                v4 = uint32(v25)
                                            case 42:
                                                v80 := GenS2{F0: ""}
                                                _ = v80
                                            default:
                                                switch "\n\t" {
                                                case "hello":
                                                    var v81 int64 = -1
                                                    _ = v81
                                                case "foo bar":
                                                    v25 += 184
                                                    switch (fi36 >> ((v2) & 63)) {
                                                    case -317:
                                                        v82 := GenS7{F0: (v4 ^ (^(v1))), F1: "\n\t", F2: (v1 * v1)}
                                                        _ = v82
                                                    case 2147483647:
                                                        var v83 uint32 = (v4 & v1)
                                                        _ = v83
                                                    }
                                                    fmt.Println(v1, fi36)
                                                default:
                                                    v1 += 2147483647
                                                }
                                            }
                                        }
                                    case "foo bar":
                                        switch (fi36 * fi33) {
                                        case 1:
                                            v84 := GenS9{F0: -659, F1: false}
                                            _ = v84
                                        case 32767:
                                            v85 := GenS10{F0: uint(v1)}
                                            _ = v85
                                        case -683:
                                            v86 := GenS6{F0: (^(v25))}
                                            _ = v86
                                        }
                                    }
                                }
                            }
                            if (v4 > (^(v1))) {
                                v4 = (v4 / ((v1) | 1))
                                var v87 int64 = 0
                                if (153 <= -897) {
                                    switch (("a" + "") + (("hello" + "hello") + "hello")) {
                                    case "hello":
                                        var v88 string = ("hello" + "hello")
                                        _ = v88
                                    case "\n\t":
                                        v89 := GenS11{F0: 1, F1: -128, F2: v87}
                                        if true {
                                            var v90 uint = v2
                                            var v91 uint64 = 18446744073709551615
                                            v92 := GenS12{F0: v90, F1: (v87 ^ v87), F2: (fi36 | -1)}
                                            _ = v90
                                            _ = v91
                                            _ = v92
                                        } else {
                                            v2 = (v2 * 2)
                                            v3 = append(v3, "\n\t")
                                            switch "hello" {
                                            case "\n\t":
                                                v93 := GenS12{F0: (^(v2)), F1: v87, F2: (fi33 % (((fi36 + (fi33 << (((v2 >> ((v2) & 63))) & 63)))) | 1))}
                                                _ = v93
                                            case "":
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                            case "a":
                                                _ = len(v3)
                                            }
                                            _ = len(v3)
                                            var v94 int = (fi33 % ((fi36) | 1))
                                            v1 += 854
                                            var v95 int16 = 1000
                                            var v96 uint32 = v4
                                            var v97 string = ((("" + (("a" + ("\n\t" + ("a" + (("hello" + "\n\t") + "a")))) + "")) + ("hello" + ("" + ("\n\t" + "foo bar")))) + "hello")
                                            _ = v94
                                            _ = v95
                                            _ = v96
                                            _ = v97
                                        }
                                        _ = v89
                                    default:
                                        v2 = (v2 & uint(v25))
                                    }
                                }
                                _ = v87
                            } else {
                                switch ("" + "hello") {
                                case "":
                                    v4 += (v4 | v1)
                                }
                                if false {
                                    var v98 uint32 = v4
                                    v2 += (^(v2))
                                    _ = v98
                                } else {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                }
                                if (42 < fi33) {
                                    v4 = 2147483647
                                    if true {
                                        var v99 uint8 = 236
                                        _ = len(v3)
                                        _ = v99
                                    }
                                } else {
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                    switch ("" + ("foo bar" + "")) {
                                    case "\n\t":
                                        v1 += v4
                                    case "foo bar":
                                        v2 += v2
                                    case "hello":
                                        if false {
                                            v2 += 2
                                            if true {
                                                fmt.Println(v2)
                                            }
                                            switch (("foo bar" + ("a" + ("\n\t" + "a"))) + "foo bar") {
                                            case "hello":
                                                v1 = (^(v4))
                                            case "foo bar":
                                                v1 += (v1 + (v1 - (^(v4))))
                                            case "a":
                                                var v100 float64 = 3.14159
                                                _ = v100
                                            }
                                        }
                                    default:
                                        var v101 float32 = -2.25
                                        _ = v101
                                    }
                                    if len(v3) >= 2 {
                                        v3 = v3[1:len(v3)]
                                    }
                                }
                                var v102 int64 = 206
                                _ = v102
                            }
                            _ = fi36
                        }
                        v103 := GenS3{F0: 712, F1: 127}
                        v4 = v4
                        _ = v103
                    }
                    _ = fi33
                    _ = v34
                }
                var v104 int16 = 1000
                v3 = append(v3, "hello")
                _ = v104
            }
        }
    } else {
        if true {
            if (!true) {
                var v105 uint8 = 255
                fmt.Println(v1, v25, v105)
                _ = v105
            }
            v25 = v25
        } else {
            if false {
                v4 = (v4 / ((v4) | 1))
                v4 += v4
                v3 = append(v3, "\n\t")
            }
            v25 = v25
            for fi106 := 0; fi106 < 12; fi106++ {
                var v107 uint16 = (v25 << ((1) & 63))
                _ = fi106
                _ = v107
            }
        }
        v108 := GenS10{F0: (^(v2))}
        v109 := GenS13{F0: 18446744073709551615}
        _ = v108
        _ = v109
    }
    _ = len(v3)
    v110 := GenS11{F0: 0, F1: 127, F2: 1}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v25:", v25)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v110
    _ = v2
    _ = v25
    _ = v3
    _ = v4
}
