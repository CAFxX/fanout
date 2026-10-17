package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 int8 }
type GenS3 struct { F0 int32; F1 int32; F2 string }
type GenS4 struct { F0 string; F1 uint }
type GenS5 struct { F0 uint16; F1 uint16; F2 bool }
type GenS6 struct { F0 uint32; F1 uint8; F2 int }
type GenS7 struct { F0 int8; F1 int64; F2 string }
type GenS8 struct { F0 bool; F1 int16; F2 int8 }
type GenS9 struct { F0 bool; F1 uint8; F2 uint32 }
type GenS10 struct { F0 uint16; F1 bool; F2 uint }
func genFn0(p0 uint8, p1 bool) uint32 {
    if p1 {
        v7 := []int{145, -121, 32768}
        p0 += (p0 >> ((953) & 63))
        _ = v7
    }
    v8 := []int{}
    _ = p0
    _ = p1
    _ = v8
    return 59
}
func genFn1(p0 int64) uint64 {
    v9 := []int{561}
    v10 := GenS1{F0: 1000}
    v11 := GenS2{F0: 50, F1: -128}
    fmt.Println(p0)
    _ = p0
    _ = v10
    _ = v11
    _ = v9
    return 658
}

func main() {
    var v1 int = 127
    v2 := []string{(("hello" + ("\n\t" + "\n\t")) + "a"), ("foo bar" + "foo bar"), ("foo bar" + "\n\t")}
    var v3 float32 = (-2.25 * 0.0)
    var v4 bool = (1 != 0)
    var v5 uint8 = 127
    v6 := make(map[uint32]bool)
    for fi12 := 0; fi12 < 2; fi12++ {
        switch "" {
        case "\n\t":
            if v4 {
                for fi13 := 0; fi13 < 11; fi13++ {
                    v14 := v6[42]
                    v15 := GenS2{F0: v5, F1: 100}
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                var v16 string = "foo bar"
                v1 += (fi12 + v1)
                _ = v16
            }
        default:
            switch "" {
            case "hello":
                var v17 bool = v4
                _ = v17
            case "\n\t":
                switch 9223372036854775807 {
                case 127:
                    switch ("" + ((("a" + "\n\t") + "a") + "\n\t")) {
                    case "a":
                        for fi18 := 0; fi18 < 16; fi18++ {
                            v19 := GenS2{F0: 124, F1: 11}
                            v4 = (4294967295 == 0)
                            v20 := GenS1{F0: 1000}
                            _ = fi18
                            _ = v19
                            _ = v20
                        }
                    case "hello":
                        v21 := GenS3{F0: -2147483648, F1: -886, F2: (("a" + "\n\t") + "hello")}
                        _ = v21
                    }
                default:
                    for fi22 := 0; fi22 < 19; fi22++ {
                        if false {
                            v4 = (v3 <= v3)
                        } else {
                            v23 := GenS4{F0: "hello", F1: 283}
                            _ = v23
                        }
                        var v24 int8 = int8(v5)
                        _ = fi22
                        _ = v24
                    }
                }
            case "":
                if ((2147483647 == -750) || (!((((("foo bar" + "foo bar") + (("a" + "foo bar") + "a")) + "") + "") >= ("" + "")))) {
                    fmt.Println(fi12, fi12)
                    switch ("" + "hello") {
                    case "hello":
                        for fi25 := 0; fi25 < 8; fi25++ {
                            v5 += (v5 + (v5 & (^(v5))))
                            v1 = (fi25 * (fi25 >> ((216) & 63)))
                            _ = fi25
                        }
                    case "\n\t":
                        var v26 uint16 = 0
                        _ = v26
                    default:
                        var v27 uint8 = v5
                        _ = v27
                    }
                } else {
                    v28 := GenS2{F0: v5, F1: 87}
                    _ = v28
                }
            }
        }
        var v29 uint8 = v5
        fmt.Println(v3, fi12)
        _ = fi12
        _ = v29
    }
    v30 := GenS2{F0: v5, F1: -3}
    _ = len(v2)
    v31 := GenS4{F0: ("foo bar" + "a"), F1: 812}
    v32 := GenS5{F0: uint16(v5), F1: 255, F2: (!v4)}
    if (!(true && (((!(v3 > 0.0)) || true) && ((v5 >> ((1) & 63)) > (v5 << ((65535) & 63)))))) {
        var v33 uint64 = 2
        for fi34 := 0; fi34 < 16; fi34++ {
            if (!(v5 > v5)) {
                v3 = v3
                for fi35 := 0; fi35 < 13; fi35++ {
                    if (4294967295 > genFn0(v5, v4)) {
                        if ("\n\t" < ("foo bar" + "foo bar")) {
                            if ((v5 ^ (^(v5))) != (v5 + (^(v5)))) {
                                if v4 {
                                    switch fi35 {
                                    case 42:
                                        var v36 uint16 = 2
                                        _ = v36
                                    case -453:
                                        var v37 int64 = -9223372036854775808
                                        _ = v37
                                    case 32767:
                                        switch (v1 ^ -537) {
                                        case 42:
                                            if (-1 >= 100) {
                                                v38 := v6[242]
                                                _ = v38
                                            }
                                        case 4294967295:
                                            v39 := v6[131]
                                            _ = v39
                                        default:
                                            var v40 int = (-(fi35))
                                            _ = v40
                                        }
                                    default:
                                        v4 = true
                                    }
                                    v1 = (fi34 + 2147483647)
                                }
                                switch ("hello" + ("hello" + "a")) {
                                case "":
                                    v3 = v3
                                case "hello":
                                    fmt.Println(v4, fi35, v1)
                                }
                                v3 = v3
                            } else {
                                var v41 uint64 = 18446744073709551615
                                if len(v2) > 0 {
                                    v2[0] = "hello"
                                }
                                fmt.Println(v5, fi35, v1)
                                v42 := GenS3{F0: -613, F1: 2147483647, F2: "foo bar"}
                                v43 := GenS2{F0: v5, F1: -1}
                                v6[42] = (!(1.5 <= 3.14159))
                                _ = v41
                                _ = v42
                                _ = v43
                            }
                            switch ("hello" + "hello") {
                            case "a":
                                v44 := GenS5{F0: 2, F1: 201, F2: v4}
                                _ = v44
                            case "foo bar":
                                var v45 float32 = ((v3 * -2.25) / float32(1e10))
                                _ = v45
                            case "":
                                delete(v6, 107)
                                v2 = append(v2, "foo bar")
                                var v46 uint8 = v5
                                _ = v46
                            }
                        } else {
                            v5 = (^(v5))
                            var v47 float64 = 1e10
                            v33 = 369
                            v48 := GenS2{F0: 165, F1: 1}
                            v49 := GenS2{F0: (v5 & v5), F1: -87}
                            if (!false) {
                                _ = len(v6)
                                if v4 {
                                    v50 := GenS4{F0: "\n\t", F1: 2}
                                    var v51 bool = (!(!v4))
                                    switch 922 {
                                    case 42:
                                        if (v33 >= (v33 + v33)) {
                                            var v52 float32 = ((-2.25 + 0.1) / v3)
                                            v3 = v3
                                            var v53 float32 = (((v3 - (float32(v47) * ((3.14159 / (v3 * (-2.25 * float32((v47 + v47))))) / v3))) + v52) / (1e10 * (((v3 + v52) * ((v3 * v3) + v3)) * 0.1)))
                                            _ = v52
                                            _ = v53
                                        }
                                    }
                                    _ = v50
                                    _ = v51
                                }
                                var v54 uint32 = 0
                                _ = v54
                            }
                            v5 = (v5 / (((v5 / ((0) | 1))) | 1))
                            if (((v33 & (v33 - v33)) <= (v33 - (v33 | v33))) || false) {
                                v55 := v6[genFn0((v5 << ((622) & 63)), (100 != 42))]
                                v56 := GenS3{F0: 262, F1: -1, F2: "foo bar"}
                                _ = v55
                                _ = v56
                            }
                            _ = v47
                            _ = v48
                            _ = v49
                        }
                        v57 := GenS3{F0: 2147483647, F1: 2147483647, F2: "a"}
                        _ = v57
                    }
                    var v58 uint16 = 256
                    v3 = v3
                    _ = len(v2)
                    var v59 float64 = 0.1
                    _ = fi35
                    _ = v58
                    _ = v59
                }
            } else {
                var v60 int32 = 42
                v5 = (^(v5))
                if v4 {
                    v33 = v33
                    v6[1] = (!(2147483648 > 2147483647))
                } else {
                    v61 := GenS6{F0: 18, F1: (^(v5)), F2: v1}
                    fmt.Println(v5, v60, v1)
                    if (1000 >= 376) {
                        v33 = (v33 >> ((1) & 63))
                        v5 = (v5 & (v5 % ((uint8(v33)) | 1)))
                    }
                    _ = v61
                }
                _ = v60
            }
            if (900 == 4294967295) {
                v33 -= (v33 / ((v33) | 1))
                var v62 string = ("foo bar" + "")
                _ = v62
            } else {
                v63 := GenS7{F0: 127, F1: -1, F2: "foo bar"}
                v3 = (v3 / v3)
                _ = v63
            }
            if true {
                for fi64 := 0; fi64 < 9; fi64++ {
                    _ = len(v2)
                    var v65 uint8 = (v5 >> ((1) & 63))
                    if false {
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    } else {
                        var v66 uint64 = (v33 ^ (v33 / (((v33 | (v33 - 1))) | 1)))
                        if (9223372036854775807 <= 42) {
                            if (fi64 == v1) {
                                delete(v6, 4294967295)
                                switch 32768 {
                                case -1:
                                    v4 = ((v3 - (v3 / (0.0 * (float32(1e10) / (v3 - v3))))) < 1e10)
                                case 1:
                                    if v4 {
                                        v67 := GenS7{F0: -80, F1: 1, F2: ""}
                                        var v68 int8 = 100
                                        if false {
                                            fmt.Println(v68, v5)
                                            if len(v2) >= 2 {
                                                v2 = v2[1:len(v2)]
                                            }
                                            if len(v2) > 0 {
                                                v2[0] = ("\n\t" + "hello")
                                            }
                                        } else {
                                            v69 := GenS8{F0: (2 > 996), F1: 32767, F2: 42}
                                            v1 -= (^(v1))
                                            _ = v69
                                        }
                                        _ = v67
                                        _ = v68
                                    } else {
                                        var v70 uint64 = v33
                                        _ = v70
                                    }
                                case 9223372036854775807:
                                    v1 = 881
                                }
                            } else {
                                fmt.Println(v3)
                                v5 -= 1
                            }
                            delete(v6, 477)
                            switch "a" {
                            case "\n\t":
                                fmt.Println(fi64)
                                v33 -= (v66 & v33)
                                switch ((("foo bar" + "\n\t") + "foo bar") + (("a" + "hello") + "\n\t")) {
                                case "hello":
                                    fmt.Println(fi64)
                                default:
                                    _ = len(v6)
                                }
                            case "a":
                                var v71 int = fi34
                                _ = v71
                            }
                        } else {
                            v4 = (4294967295 != 2147483647)
                            v65 = (v65 + v5)
                            var v72 int32 = 1
                            _ = v72
                        }
                        _ = v66
                    }
                    _ = fi64
                    _ = v65
                }
                v73 := GenS2{F0: (v5 | (v5 << ((704) & 63))), F1: 100}
                _ = v73
            }
            _ = fi34
        }
        _ = v33
    }
    var v74 int64 = -1
    for fi75 := 0; fi75 < 12; fi75++ {
        v76 := GenS1{F0: -901}
        _ = fi75
        _ = v76
    }
    for fi77 := 0; fi77 < 14; fi77++ {
        fmt.Println(v5)
        _ = fi77
    }
    switch "" {
    case "\n\t":
        for fi78 := 0; fi78 < 1; fi78++ {
            var v79 int16 = 32767
            var v80 int32 = 42
            _ = fi78
            _ = v79
            _ = v80
        }
    }
    v6[736] = (("foo bar" + (("" + "\n\t") + ((("foo bar" + (("" + "\n\t") + "")) + ("\n\t" + "foo bar")) + ("hello" + "")))) > "")
    v3 += float32(float64(3.14159))
    for fi81 := 0; fi81 < 6; fi81++ {
        if ((v4 && (!(0 < 0))) || v4) {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
        }
        var v82 float64 = 1.5
        _ = fi81
        _ = v82
    }
    for fi83 := 0; fi83 < 4; fi83++ {
        v1 = (fi83 << ((0) & 63))
        switch "\n\t" {
        case "\n\t":
            v5 = (v5 / ((v5) | 1))
        default:
            switch (v1 ^ fi83) {
            case 761:
                if v4 {
                    for fi84 := 0; fi84 < 9; fi84++ {
                        var v85 string = "foo bar"
                        _ = fi84
                        _ = v85
                    }
                    _ = len(v6)
                }
            case -2147483648:
                _ = len(v6)
            case 32768:
                switch (v1 / ((v1) | 1)) {
                case 0:
                    for fi86 := 0; fi86 < 14; fi86++ {
                        v3 -= (((-2.25 + v3) / v3) + v3)
                        _ = fi86
                    }
                case -9223372036854775808:
                    var v87 int = fi83
                    _ = v87
                default:
                    switch (("" + "foo bar") + "") {
                    case "hello":
                        for fi88 := 0; fi88 < 16; fi88++ {
                            v89 := GenS8{F0: false, F1: 1000, F2: -1}
                            _ = fi88
                            _ = v89
                        }
                    case "a":
                        for fi90 := 0; fi90 < 1; fi90++ {
                            v3 = ((0.1 - float32((-2.25 / float64(v3)))) / v3)
                            _ = fi90
                        }
                    case "":
                        var v91 uint32 = 673
                        _ = v91
                    default:
                        for fi92 := 0; fi92 < 5; fi92++ {
                            v93 := GenS9{F0: v4, F1: (^(v5)), F2: 2}
                            _ = fi92
                            _ = v93
                        }
                    }
                }
            default:
                fmt.Println(v3)
            }
        }
        for fi94 := 0; fi94 < 8; fi94++ {
            v95 := GenS10{F0: 255, F1: v4, F2: 18446744073709551615}
            v96 := GenS10{F0: 42, F1: (v4 && (!false)), F2: 721}
            _ = fi94
            _ = v95
            _ = v96
        }
        _ = fi83
    }
    if v4 {
        switch "hello" {
        case "":
            v4 = false
        case "\n\t":
            v5 += v5
        case "a":
            fmt.Println(v4)
        }
    } else {
        _ = len(v2)
    }
    if ((v74 - v74) != (^(v74))) {
        fmt.Println(v3, v5)
    }
    if v4 {
        v97 := GenS9{F0: (0 != 42), F1: v5, F2: 2}
        _ = v97
    } else {
        var v98 bool = (365 == 65535)
        v99 := GenS4{F0: "foo bar", F1: 215}
        _ = v98
        _ = v99
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
    _ = v74
}
