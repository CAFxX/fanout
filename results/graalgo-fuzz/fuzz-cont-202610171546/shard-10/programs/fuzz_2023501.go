package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint16 }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 int; F1 int16 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 string; F1 uint32 }
type GenS6 struct { F0 bool; F1 int }
type GenS7 struct { F0 int8; F1 uint64 }
type GenS8 struct { F0 int8; F1 int8 }
type GenS9 struct { F0 uint16; F1 uint8; F2 uint }
type GenS10 struct { F0 uint64; F1 uint16 }
type GenS11 struct { F0 uint16 }
type GenS12 struct { F0 int32 }
type GenS13 struct { F0 int8 }
type GenS14 struct { F0 int16; F1 uint32 }
type GenS15 struct { F0 int; F1 uint64 }
func genFn0() uint32 {
    for fi6 := 0; fi6 < 5; fi6++ {
        v7 := GenS1{F0: 382, F1: 32767}
        var v8 int16 = 32767
        v9 := []int{}
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    switch ("foo bar" + ("a" + "\n\t")) {
    case "":
        var v10 int32 = -2147483648
        _ = v10
    case "foo bar":
        for fi11 := 0; fi11 < 8; fi11++ {
            var v12 int8 = 66
            var v13 int16 = -1
            _ = fi11
            _ = v12
            _ = v13
        }
    }
    var v14 uint16 = 32767
    v15 := GenS2{F0: 2147483648}
    _ = v14
    _ = v15
    return 0
}
func genFn1(p0 uint) string {
    var v16 uint64 = 18446744073709551615
    v16 = (^(v16))
    p0 += p0
    _ = p0
    _ = v16
    return "\n\t"
}

func main() {
    v1 := make(map[string]int)
    v2 := make(map[string]string)
    var v3 int = 127
    var v4 int32 = 1
    var v5 int64 = 303
    for fi17 := 0; fi17 < 13; fi17++ {
        if false {
            fmt.Println(v4)
            fmt.Println(v5)
            v18 := GenS3{F0: (fi17 & (v3 + (v3 / (((v3 & (fi17 * v3))) | 1)))), F1: 1}
            _ = v18
        } else {
            v19 := []int{}
            _ = v19
        }
        v3 += (v3 % ((fi17) | 1))
        _ = fi17
    }
    fmt.Println(v4, v5, v4)
    var v20 bool = true
    v3 -= (v3 + -819)
    v20 = (false || (((v20 || (((1.5 + -2.25) + float32(0.0)) == (0.0 * 1e10))) || (0 >= 575)) && v20))
    fmt.Println(v4, v20, v4)
    v20 = v20
    switch v3 {
    case 32767:
        switch "\n\t" {
        case "hello":
            v21 := GenS4{F0: ((("hello" + "foo bar") + (("a" + "a") + (((("foo bar" + "\n\t") + (("" + (((("" + "") + ("\n\t" + ((((((("a" + ("hello" + "foo bar")) + "foo bar") + (("a" + "a") + "a")) + "a") + "foo bar") + "a") + "\n\t"))) + "hello") + "a")) + "foo bar")) + "\n\t") + ""))) + "\n\t")}
            _ = v21
        case "":
            fmt.Println(v3)
        case "a":
            for fi22 := 0; fi22 < 2; fi22++ {
                if ((!v20) && true) {
                    delete(v1, "foo bar")
                }
                _ = fi22
            }
        default:
            v23 := GenS5{F0: "foo bar", F1: 1}
            _ = v23
        }
    case 9223372036854775807:
        if false {
            delete(v1, "")
            if true {
                v24 := GenS6{F0: v20, F1: v3}
                fmt.Println(v20)
                _ = v24
            } else {
                for fi25 := 0; fi25 < 7; fi25++ {
                    for fi26 := 0; fi26 < 17; fi26++ {
                        if (false || (!((true || (32767 != 65535)) || (2147483647 > -744)))) {
                            var v27 uint64 = 4294967295
                            var v28 int16 = -644
                            _ = v27
                            _ = v28
                        } else {
                            v3 -= -1
                            v29 := []int{}
                            _ = v29
                        }
                        _ = fi26
                    }
                    v30 := GenS6{F0: (-2147483648 >= -587), F1: fi25}
                    _ = fi25
                    _ = v30
                }
                v4 += (-(v4))
                fmt.Println(v20, v20, v4)
            }
        } else {
            switch "a" {
            case "\n\t":
                if (18446744073709551615 == 18446744073709551615) {
                    v5 = v5
                } else {
                    v31 := []int{v3, (v3 / ((v3) | 1)), v3}
                    _ = v31
                }
            }
            if false {
                var v32 bool = v20
                var v33 uint = 2
                _ = v32
                _ = v33
            } else {
                v4 += (v4 ^ 0)
            }
        }
    default:
        if v20 {
            fmt.Println(v3)
            v2[("a" + ("\n\t" + (((((("a" + ("\n\t" + ((("hello" + "hello") + "hello") + "\n\t"))) + "hello") + "foo bar") + "") + "") + "hello")))] = ((("" + ("a" + "foo bar")) + "foo bar") + ("hello" + (("" + "\n\t") + ("hello" + ""))))
            var v34 bool = ((100 >= -57) || v20)
            _ = v34
        }
    }
    if (!(!true)) {
        if (0.1 <= float64(-2.25)) {
            if v20 {
                v35 := []int{}
                v3 -= (v3 / ((v3) | 1))
                _ = v35
            } else {
                var v36 uint = 65535
                _ = v36
            }
            v37 := v2["\n\t"]
            v3 = v3
            _ = v37
        } else {
            v38 := GenS7{F0: 100, F1: 1}
            v39 := []int{v3, v3}
            _ = v38
            _ = v39
        }
        for fi40 := 0; fi40 < 3; fi40++ {
            var v41 int64 = v5
            v4 -= v4
            _ = fi40
            _ = v41
        }
        v42 := GenS2{F0: 4294967295}
        _ = v42
    } else {
        var v43 uint8 = 42
        switch "foo bar" {
        case "a":
            if v20 {
                var v44 float64 = 1e10
                var v45 uint8 = v43
                _ = v44
                _ = v45
            }
        }
        _ = v43
    }
    fmt.Println(v3, v20, v5)
    v20 = v20
    v46 := GenS4{F0: (("hello" + "hello") + "hello")}
    if (128 > 110) {
        switch (v3 * v3) {
        case 0:
            v20 = (0 == 1)
        case 9223372036854775807:
            var v47 bool = (-128 != 42)
            _ = v47
        default:
            if (v5 >= (-(v5))) {
                v1["foo bar"] = v3
                var v48 uint64 = 2
                _ = v48
            } else {
                v5 += v5
                for fi49 := 0; fi49 < 4; fi49++ {
                    switch 4294967295 {
                    case -536:
                        v50 := v1["\n\t"]
                        _ = v50
                    }
                    v51 := GenS8{F0: -128, F1: -1}
                    _ = fi49
                    _ = v51
                }
                v3 = (v3 ^ (^(v3)))
            }
        }
        fmt.Println(v4, v20)
    } else {
        switch (v3 % (((v3 * 127)) | 1)) {
        case 0:
            if v20 {
                v5 = 1
            } else {
                if (!(!true)) {
                    fmt.Println(v4)
                    for fi52 := 0; fi52 < 16; fi52++ {
                        v20 = (-1 <= 0)
                        v53 := GenS3{F0: (v3 + (fi52 >> ((42) & 63))), F1: -922}
                        _ = fi52
                        _ = v53
                    }
                    v3 = (^(v3))
                }
            }
        case 25:
            delete(v1, (((("foo bar" + (("\n\t" + "a") + "hello")) + "a") + ("" + ("foo bar" + "hello"))) + ((("foo bar" + "foo bar") + ("hello" + "hello")) + "a")))
        case -9223372036854775808:
            var v54 int64 = -9223372036854775808
            _ = v54
        }
        fmt.Println(v5, v5, v5)
    }
    if (612 > 451) {
        for fi55 := 0; fi55 < 7; fi55++ {
            v3 -= fi55
            switch -974 {
            case 884:
                var v56 uint = 65535
                _ = v56
            }
            var v57 string = "foo bar"
            _ = fi55
            _ = v57
        }
    }
    if (!(((v20 || (0.0 > -2.25)) || v20) || true)) {
        v58 := GenS4{F0: "a"}
        v59 := []int{v3, v3, -430, (v3 & (v3 | 42))}
        _ = v58
        _ = v59
    } else {
        v20 = (!v20)
        switch v3 {
        case -114:
            v60 := GenS9{F0: 65535, F1: 1, F2: 4294967295}
            _ = v60
        case 4294967295:
            if (true && ((v5 + (v5 / (((v5 >> ((18446744073709551615) & 63))) | 1))) == (^(v5)))) {
                switch "\n\t" {
                case "":
                    for fi61 := 0; fi61 < 14; fi61++ {
                        var v62 uint64 = 4294967295
                        _ = fi61
                        _ = v62
                    }
                case "a":
                    v63 := GenS4{F0: ("\n\t" + "\n\t")}
                    _ = v63
                case "hello":
                    v4 = -1
                default:
                    if (false && ((v3 + v3) == -553)) {
                        v64 := []int{(v3 << ((260) & 63)), (v3 | 128), -112}
                        switch (("a" + "\n\t") + ("hello" + "a")) {
                        case "foo bar":
                            var v65 int16 = 1000
                            _ = v65
                        default:
                            v64 = append(v64, 260)
                        }
                        for fi66 := 0; fi66 < 11; fi66++ {
                            switch 4294967295 {
                            case -2147483648:
                                v4 = (v4 % ((v4) | 1))
                            case -993:
                                for fi67 := 0; fi67 < 14; fi67++ {
                                    switch v3 {
                                    case -9223372036854775808:
                                        var v68 float64 = 1e10
                                        _ = v68
                                    default:
                                        switch v3 {
                                        case 42:
                                            v4 += (v4 + 0)
                                        case -651:
                                            switch "\n\t" {
                                            case "":
                                                var v69 int8 = -74
                                                _ = v69
                                            default:
                                                var v70 int32 = -2147483648
                                                _ = v70
                                            }
                                        case -198:
                                            var v71 int32 = (v4 ^ v4)
                                            _ = v71
                                        }
                                    }
                                    fmt.Println(v20, v20)
                                    if (!(((32767 < 2) || v20) && true)) {
                                        if v20 {
                                            var v72 uint64 = 4294967295
                                            _ = len(v64)
                                            var v73 float64 = ((0.1 - (-2.25 * (-2.25 - 1e10))) * 0.1)
                                            _ = v72
                                            _ = v73
                                        } else {
                                            switch (fi66 >> ((465) & 63)) {
                                            case 32767:
                                                var v74 int8 = 1
                                                _ = v74
                                            case -570:
                                                v75 := GenS10{F0: 308, F1: 42}
                                                _ = v75
                                            case 127:
                                                var v76 uint16 = 65535
                                                _ = v76
                                            }
                                            v77 := GenS4{F0: ""}
                                            _ = v77
                                        }
                                        switch "foo bar" {
                                        case "\n\t":
                                            v78 := GenS3{F0: fi66, F1: -599}
                                            _ = v78
                                        case "hello":
                                            _ = len(v64)
                                        }
                                        fmt.Println(v5)
                                        v5 -= int64(v4)
                                        switch fi66 {
                                        case 9223372036854775807:
                                            v79 := v1["foo bar"]
                                            _ = v79
                                        case -9223372036854775808:
                                            var v80 float32 = 3.14159
                                            _ = v80
                                        case -580:
                                            v81 := v1["\n\t"]
                                            _ = v81
                                        }
                                        v82 := GenS3{F0: (v3 ^ fi67), F1: 32767}
                                        if ((-2.25 + 3.14159) < 0.1) {
                                            v83 := v1[("foo bar" + (("foo bar" + "") + "hello"))]
                                            var v84 uint16 = 827
                                            var v85 string = ("hello" + "a")
                                            if v20 {
                                                v84 += (v84 % ((v84) | 1))
                                            } else {
                                                var v86 bool = v20
                                                v87 := GenS2{F0: 1}
                                                _ = v86
                                                _ = v87
                                            }
                                            v88 := GenS10{F0: 0, F1: (v84 % (((^(v84))) | 1))}
                                            _ = len(v2)
                                            _ = v83
                                            _ = v84
                                            _ = v85
                                            _ = v88
                                        }
                                        var v89 int32 = v4
                                        _ = v82
                                        _ = v89
                                    } else {
                                        switch (v3 ^ (-(fi66))) {
                                        case 128:
                                            v5 = 360
                                        default:
                                            switch "foo bar" {
                                            case "hello":
                                                v20 = (1.5 >= (float64(3.14159) / float64(1e10)))
                                            case "a":
                                                if true {
                                                    v90 := GenS11{F0: 32767}
                                                    _ = v90
                                                } else {
                                                    var v91 uint64 = 0
                                                    v91 = 590
                                                    var v92 float64 = 1e10
                                                    _ = v91
                                                    _ = v92
                                                }
                                            }
                                        }
                                    }
                                    _ = fi67
                                }
                            case 4294967295:
                                switch (v3 & 823) {
                                case -2147483648:
                                    for fi93 := 0; fi93 < 1; fi93++ {
                                        if (!false) {
                                            fmt.Println(v5, v3, fi93)
                                            var v94 string = "a"
                                            var v95 float64 = float64(float32(0.1))
                                            v94 = ("a" + "foo bar")
                                            _ = v94
                                            _ = v95
                                        }
                                        if v20 {
                                            var v96 uint16 = 65535
                                            _ = v96
                                        }
                                        v97 := GenS4{F0: "foo bar"}
                                        _ = fi93
                                        _ = v97
                                    }
                                case -1:
                                    if v20 {
                                        v98 := GenS12{F0: 1}
                                        v4 = v4
                                        _ = v98
                                    }
                                case 42:
                                    if ((!((!true) || (-32768 > -32768))) || ((v3 - fi66) >= fi66)) {
                                        if (!v20) {
                                            switch ("foo bar" + "hello") {
                                            case "a":
                                                if (!v20) {
                                                    v5 = v5
                                                    fmt.Println(v4, v4)
                                                    if (42 >= 32767) {
                                                        v4 -= -918
                                                    }
                                                }
                                            }
                                            fmt.Println(v3, v4)
                                            switch v3 {
                                            case 755:
                                                var v99 int32 = (v4 & -1)
                                                _ = v99
                                            case 128:
                                                if ((false && (42 != 4294967295)) && v20) {
                                                    v5 -= v5
                                                }
                                            case -1:
                                                var v100 float32 = float32(1e10)
                                                _ = v100
                                            }
                                        }
                                        _ = len(v1)
                                        switch "" {
                                        case "a":
                                            fmt.Println(v3, v5)
                                        case "foo bar":
                                            var v101 int32 = 42
                                            _ = v101
                                        }
                                    } else {
                                        var v102 int16 = 32767
                                        v103 := GenS13{F0: -1}
                                        _ = v102
                                        _ = v103
                                    }
                                default:
                                    switch v3 {
                                    case 25:
                                        var v104 float32 = 0.1
                                        _ = v104
                                    case 321:
                                        switch 355 {
                                        case 1:
                                            fmt.Println(v5, v4)
                                        }
                                    default:
                                        var v105 uint8 = 42
                                        _ = v105
                                    }
                                }
                            }
                            v106 := GenS14{F0: 596, F1: 2147483648}
                            v64 = append(v64, (^(fi66)))
                            _ = fi66
                            _ = v106
                        }
                        _ = v64
                    }
                }
                delete(v1, ("\n\t" + ("foo bar" + "hello")))
                v107 := GenS13{F0: 42}
                _ = v107
            }
        }
        v108 := []int{32767, v3, -9223372036854775808}
        _ = v108
    }
    v109 := GenS15{F0: v3, F1: 42}
    v3 -= v3
    fmt.Println(v20)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v109
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v46
    _ = v5
}
