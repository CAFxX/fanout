package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8 }
type GenS2 struct { F0 int8; F1 string; F2 string }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 uint16; F1 int8 }
type GenS5 struct { F0 uint32; F1 int32; F2 string }
type GenS6 struct { F0 uint16 }
type GenS7 struct { F0 uint64; F1 int16; F2 int8 }
type GenS8 struct { F0 string; F1 int16 }
type GenS9 struct { F0 int; F1 uint16 }
type GenS10 struct { F0 uint64; F1 int; F2 int32 }
type GenS11 struct { F0 int64 }
type GenS12 struct { F0 bool; F1 uint64 }
func genFn0(p0 int32) uint32 {
    v7 := []int{128, -943, -1, 9223372036854775807}
    if true {
        for fi8 := 0; fi8 < 16; fi8++ {
            v9 := GenS1{F0: 0}
            if (!(true || false)) {
                for fi10 := 0; fi10 < 2; fi10++ {
                    v11 := GenS2{F0: 100, F1: ("a" + "\n\t"), F2: ""}
                    fmt.Println(p0, fi10, p0)
                    p0 += p0
                    _ = fi10
                    _ = v11
                }
                var v12 uint32 = 1
                _ = v12
            }
            var v13 uint = 1
            _ = fi8
            _ = v13
            _ = v9
        }
        v7 = append(v7, 947)
    }
    _ = p0
    _ = v7
    return 1
}
func genFn1(p0 uint8, p1 int8) int {
    p1 = (p1 & p1)
    var v14 int64 = 0
    if false {
        v15 := []int{32768, -9223372036854775808}
        _ = v15
    } else {
        var v16 int8 = 0
        _ = v16
    }
    var v17 uint8 = p0
    var v18 int64 = v14
    _ = p0
    _ = p1
    _ = v14
    _ = v17
    _ = v18
    return 245
}

func main() {
    v1 := []int{984, 403, -333, 32767}
    v2 := []string{("a" + "hello"), "hello"}
    v3 := []uint32{}
    var v4 int32 = -787
    var v5 int = 2147483647
    var v6 uint8 = 127
    var v19 float32 = ((1.5 / 1.5) + 0.0)
    var v20 int16 = 1000
    v4 = v4
    var v21 uint8 = v6
    if len(v2) > 0 {
        v2[0] = ""
    }
    if (950 < 238) {
        v22 := GenS3{F0: 128}
        _ = len(v3)
        switch genFn1(v6, 112) {
        case -903:
            if (genFn1(160, 100) < v5) {
                v23 := GenS2{F0: 100, F1: "", F2: "foo bar"}
                if false {
                    if false {
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                        if false {
                            for fi24 := 0; fi24 < 12; fi24++ {
                                switch "a" {
                                case "a":
                                    for fi25 := 0; fi25 < 3; fi25++ {
                                        if ((v6 | v6) <= (v6 - (v21 * (v21 >> ((65535) & 63))))) {
                                            v3 = append(v3, 825)
                                            var v26 int32 = (v4 * (v4 << ((337) & 63)))
                                            _ = v26
                                        }
                                        if true {
                                            fmt.Println(fi25)
                                        } else {
                                            if (0.0 < v19) {
                                                if (v4 == (v4 / ((-2147483648) | 1))) {
                                                    v6 = v21
                                                }
                                            } else {
                                                v27 := GenS4{F0: 522, F1: 1}
                                                if len(v1) > 0 {
                                                    v1[0] = (fi25 ^ fi24)
                                                }
                                                if true {
                                                    v20 = (v20 / (((v20 * int16(v4))) | 1))
                                                    v5 = 32767
                                                    if false {
                                                        switch "" {
                                                        case "hello":
                                                            switch (^(v5)) {
                                                            case 2147483647:
                                                                if len(v2) > 0 {
                                                                    v2[0] = ""
                                                                }
                                                                switch (fi25 | 1) {
                                                                case -1:
                                                                    if len(v3) > 0 {
                                                                        v3[0] = 2
                                                                    }
                                                                    v21 = (v21 / ((128) | 1))
                                                                    v4 = -10
                                                                    if true {
                                                                        var v28 int16 = v20
                                                                        v29 := GenS1{F0: (v21 >> ((1) & 63))}
                                                                        var v30 float64 = float64((v19 / v19))
                                                                        v4 = (v4 % ((v4) | 1))
                                                                        _ = v28
                                                                        _ = v29
                                                                        _ = v30
                                                                    } else {
                                                                        if true {
                                                                            fmt.Println(v6, v6)
                                                                            switch "foo bar" {
                                                                            case "foo bar":
                                                                                var v31 int32 = 618
                                                                                _ = v31
                                                                            case "":
                                                                                var v32 int32 = v4
                                                                                _ = v32
                                                                            case "hello":
                                                                                var v33 float32 = 1e10
                                                                                _ = v33
                                                                            }
                                                                            v6 -= v21
                                                                        }
                                                                        switch "\n\t" {
                                                                        case "foo bar":
                                                                            v21 = (v21 & 127)
                                                                        case "\n\t":
                                                                            if len(v2) > 0 {
                                                                                v2[0] = ("\n\t" + "")
                                                                            }
                                                                        case "":
                                                                            v21 = (v21 | 127)
                                                                        }
                                                                    }
                                                                    if false {
                                                                        v6 = 90
                                                                    }
                                                                    if (true && ((("foo bar" + ((("a" + ("" + "hello")) + "") + ((("\n\t" + (((("" + ("foo bar" + "a")) + "a") + "a") + "hello")) + "foo bar") + "\n\t"))) + "a") != ("hello" + ("a" + "hello")))) {
                                                                        v34 := GenS5{F0: 42, F1: v4, F2: "\n\t"}
                                                                        _ = len(v1)
                                                                        switch ("foo bar" + (("" + "hello") + "a")) {
                                                                        case "foo bar":
                                                                            v35 := GenS3{F0: v21}
                                                                            _ = v35
                                                                        default:
                                                                            v2 = append(v2, ("hello" + (((("a" + (("foo bar" + "a") + "foo bar")) + "") + (("\n\t" + ((("foo bar" + "a") + "") + ("hello" + ("\n\t" + ("foo bar" + "hello"))))) + ("\n\t" + "hello"))) + ("\n\t" + "a"))))
                                                                        }
                                                                        _ = v34
                                                                    }
                                                                    if true {
                                                                        fmt.Println(v20, v5)
                                                                        v4 = v4
                                                                        var v36 int16 = v20
                                                                        _ = v36
                                                                    } else {
                                                                        switch (fi25 - -9223372036854775808) {
                                                                        case 2147483647:
                                                                            v4 += 1
                                                                        case 128:
                                                                            var v37 string = "foo bar"
                                                                            v3 = append(v3, 1)
                                                                            if len(v3) > 0 {
                                                                                v3[0] = 400
                                                                            }
                                                                            _ = v37
                                                                        case 0:
                                                                            v1 = append(v1, v5)
                                                                        }
                                                                        v5 = (fi25 << ((18446744073709551615) & 63))
                                                                    }
                                                                    var v38 bool = (v21 == v6)
                                                                    _ = v38
                                                                case -2147483648:
                                                                    if len(v3) > 0 {
                                                                        v3[0] = 1
                                                                    }
                                                                case 32768:
                                                                    if false {
                                                                        var v39 uint64 = 2
                                                                        _ = v39
                                                                    } else {
                                                                        v20 += (v20 % (((v20 + (v20 & v20))) | 1))
                                                                        if (uint(v6) != 241) {
                                                                            if len(v1) > 0 {
                                                                                v1[0] = (v5 - (-(fi24)))
                                                                            }
                                                                        } else {
                                                                            v19 = 3.14159
                                                                            var v40 string = "a"
                                                                            switch int(v20) {
                                                                            case 222:
                                                                                fmt.Println(fi24, v40)
                                                                            case 127:
                                                                                var v41 int = (v5 ^ (fi25 % ((-2147483648) | 1)))
                                                                                _ = v41
                                                                            case 32767:
                                                                                fmt.Println(v5)
                                                                            }
                                                                            _ = v40
                                                                        }
                                                                    }
                                                                default:
                                                                    v6 -= v6
                                                                }
                                                            case -683:
                                                                var v42 int16 = (v20 ^ (v20 + (v20 & (^(v20)))))
                                                                _ = v42
                                                            default:
                                                                fmt.Println(v4)
                                                            }
                                                        case "\n\t":
                                                            var v43 int8 = 0
                                                            if true {
                                                                v2 = append(v2, "")
                                                                var v44 uint32 = genFn0(v4)
                                                                switch ("" + "foo bar") {
                                                                case "foo bar":
                                                                    var v45 uint32 = 885
                                                                    _ = v45
                                                                }
                                                                v20 -= v20
                                                                fmt.Println(v44, v44)
                                                                _ = v44
                                                            } else {
                                                                var v46 uint8 = v6
                                                                _ = v46
                                                            }
                                                            var v47 int64 = int64(v20)
                                                            fmt.Println(v4, v43)
                                                            v19 = (((v19 * v19) * v19) / 1.0)
                                                            _ = v43
                                                            _ = v47
                                                        case "a":
                                                            var v48 int8 = 42
                                                            _ = v48
                                                        default:
                                                            v1 = append(v1, (fi25 >> ((0) & 63)))
                                                        }
                                                        if false {
                                                            v49 := GenS4{F0: 2, F1: 36}
                                                            if true {
                                                                if (v4 < 2147483647) {
                                                                    fmt.Println(fi25, v20, v5)
                                                                    switch v5 {
                                                                    case -9223372036854775808:
                                                                        var v50 int64 = 9223372036854775807
                                                                        _ = v50
                                                                    case 32768:
                                                                        v6 = v21
                                                                    }
                                                                }
                                                                if true {
                                                                    var v51 int32 = (v4 - v4)
                                                                    var v52 uint16 = 32767
                                                                    v2 = append(v2, "a")
                                                                    v53 := GenS5{F0: 76, F1: v4, F2: "\n\t"}
                                                                    switch (fi25 >> ((18446744073709551615) & 63)) {
                                                                    case -1:
                                                                        var v54 string = (("" + "foo bar") + ("" + ""))
                                                                        _ = v54
                                                                    case 42:
                                                                        switch (((("foo bar" + "hello") + "hello") + "a") + (("" + ("" + ("a" + ("a" + "foo bar")))) + "\n\t")) {
                                                                        case "\n\t":
                                                                            var v55 int8 = 1
                                                                            _ = v55
                                                                        case "a":
                                                                            if (v51 <= (v4 - (v51 + (v4 ^ v51)))) {
                                                                                if (((false || (!true)) && true) || false) {
                                                                                    if len(v1) > 0 {
                                                                                        v1[0] = -346
                                                                                    }
                                                                                } else {
                                                                                    fmt.Println(v51, v4)
                                                                                    var v56 int8 = 88
                                                                                    if true {
                                                                                        var v57 bool = true
                                                                                        fmt.Println(v5, v56)
                                                                                        var v58 float32 = ((1.5 * ((v19 - 0.1) * (v19 / (v19 / v19)))) / float32(((1e10 - 0.0) - 0.0)))
                                                                                        _ = v57
                                                                                        _ = v58
                                                                                    } else {
                                                                                        var v59 int64 = 0
                                                                                        v6 -= v21
                                                                                        v20 = int16(v56)
                                                                                        _ = v59
                                                                                    }
                                                                                    _ = v56
                                                                                }
                                                                            }
                                                                        default:
                                                                            if len(v3) >= 2 {
                                                                                v3 = v3[1:len(v3)]
                                                                            }
                                                                        }
                                                                    default:
                                                                        v60 := GenS6{F0: v52}
                                                                        _ = v60
                                                                    }
                                                                    v3 = append(v3, 2)
                                                                    _ = v51
                                                                    _ = v52
                                                                    _ = v53
                                                                } else {
                                                                    switch "a" {
                                                                    case "a":
                                                                        var v61 int = (^(fi25))
                                                                        _ = v61
                                                                    case "\n\t":
                                                                        v19 += 3.14159
                                                                    case "":
                                                                        var v62 int64 = -504
                                                                        _ = v62
                                                                    }
                                                                    if false {
                                                                        var v63 float64 = 3.14159
                                                                        switch (("" + ("a" + "hello")) + (("a" + ((("a" + ("hello" + "foo bar")) + "a") + "a")) + ("\n\t" + "a"))) {
                                                                        case "hello":
                                                                            if (!false) {
                                                                                v21 = v6
                                                                                if len(v3) >= 2 {
                                                                                    v3 = v3[1:len(v3)]
                                                                                }
                                                                                v64 := GenS1{F0: (v21 - (v6 << ((412) & 63)))}
                                                                                _ = v64
                                                                            }
                                                                        case "foo bar":
                                                                            switch 32768 {
                                                                            case -439:
                                                                                v6 += v6
                                                                            case -2147483648:
                                                                                if len(v1) > 0 {
                                                                                    v1[0] = fi25
                                                                                }
                                                                            }
                                                                            if false {
                                                                                if true {
                                                                                    if false {
                                                                                        v2 = append(v2, ("foo bar" + "foo bar"))
                                                                                        if true {
                                                                                            var v65 int = (fi24 / ((fi24) | 1))
                                                                                            var v66 uint16 = uint16(v21)
                                                                                            v4 -= v4
                                                                                            _ = v65
                                                                                            _ = v66
                                                                                        } else {
                                                                                            var v67 int16 = v20
                                                                                            _ = v67
                                                                                        }
                                                                                        if (-862 > 0) {
                                                                                            v68 := GenS4{F0: 2, F1: 100}
                                                                                            if len(v3) >= 2 {
                                                                                                v3 = v3[1:len(v3)]
                                                                                            }
                                                                                            _ = v68
                                                                                        } else {
                                                                                            v19 -= 3.14159
                                                                                        }
                                                                                    }
                                                                                }
                                                                                if false {
                                                                                    switch genFn1(0, 42) {
                                                                                    case -697:
                                                                                        if len(v2) > 0 {
                                                                                            v2[0] = ("a" + "")
                                                                                        }
                                                                                    }
                                                                                    _ = len(v2)
                                                                                    if true {
                                                                                        var v69 float64 = v63
                                                                                        _ = v69
                                                                                    } else {
                                                                                        v70 := GenS7{F0: 684, F1: (v20 / ((v20) | 1)), F2: 99}
                                                                                        var v71 bool = (!((((int32(v21) != (v4 >> ((44) & 63))) || true) && (42 != -1)) || true))
                                                                                        _ = v70
                                                                                        _ = v71
                                                                                    }
                                                                                }
                                                                                if (!true) {
                                                                                    v4 = 991
                                                                                }
                                                                            } else {
                                                                                fmt.Println(fi24)
                                                                                if false {
                                                                                    v4 += (v4 / ((v4) | 1))
                                                                                    fmt.Println(v5, fi25)
                                                                                }
                                                                            }
                                                                        case "":
                                                                            v3 = append(v3, genFn0((v4 << ((1) & 63))))
                                                                        }
                                                                        v72 := GenS5{F0: genFn0((v4 >> ((42) & 63))), F1: (v4 * (v4 * v4)), F2: ("hello" + (("hello" + (("foo bar" + "hello") + "hello")) + "hello"))}
                                                                        _ = v63
                                                                        _ = v72
                                                                    }
                                                                    var v73 float64 = 3.14159
                                                                    _ = v73
                                                                }
                                                                v6 = 20
                                                                v74 := GenS8{F0: "foo bar", F1: v20}
                                                                _ = v74
                                                            } else {
                                                                var v75 int64 = 0
                                                                v76 := GenS9{F0: fi24, F1: 799}
                                                                _ = v75
                                                                _ = v76
                                                            }
                                                            v20 += (v20 >> ((169) & 63))
                                                            var v77 int8 = 42
                                                            var v78 uint8 = v6
                                                            _ = v49
                                                            _ = v77
                                                            _ = v78
                                                        } else {
                                                            switch fi24 {
                                                            case 1:
                                                                v3 = append(v3, 0)
                                                            default:
                                                                var v79 uint32 = 4294967295
                                                                _ = v79
                                                            }
                                                        }
                                                        _ = len(v2)
                                                    }
                                                }
                                                if false {
                                                    if (4294967295 < 18446744073709551615) {
                                                        v80 := GenS8{F0: "hello", F1: 0}
                                                        var v81 int32 = v4
                                                        v6 = (v21 << ((2) & 63))
                                                        _ = v80
                                                        _ = v81
                                                    }
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    v21 = 0
                                                } else {
                                                    v20 = v20
                                                }
                                                _ = v27
                                            }
                                            v6 = v21
                                            var v82 int = v5
                                            _ = v82
                                        }
                                        var v83 int = (fi25 / ((fi24) | 1))
                                        _ = fi25
                                        _ = v83
                                    }
                                }
                                if true {
                                    v84 := GenS1{F0: (v6 | v6)}
                                    _ = v84
                                } else {
                                    v21 += v6
                                    var v85 uint32 = 2147483647
                                    _ = v85
                                }
                                v86 := GenS5{F0: 65535, F1: v4, F2: "\n\t"}
                                _ = fi24
                                _ = v86
                            }
                            fmt.Println(v6, v4)
                        }
                        switch (v5 - v5) {
                        case -26:
                            v87 := GenS10{F0: 212, F1: 42, F2: 2147483647}
                            _ = v87
                        }
                    } else {
                        for fi88 := 0; fi88 < 15; fi88++ {
                            v89 := GenS5{F0: 65535, F1: (v4 & -2147483648), F2: "foo bar"}
                            switch "\n\t" {
                            case "\n\t":
                                for fi90 := 0; fi90 < 5; fi90++ {
                                    v19 += 0.0
                                    v91 := GenS6{F0: 65535}
                                    _ = fi90
                                    _ = v91
                                }
                            }
                            _ = fi88
                            _ = v89
                        }
                        if false {
                            v19 = v19
                            for fi92 := 0; fi92 < 10; fi92++ {
                                var v93 uint = 42
                                v94 := GenS2{F0: -1, F1: "foo bar", F2: "\n\t"}
                                _ = fi92
                                _ = v93
                                _ = v94
                            }
                            if (-47 == 100) {
                                if false {
                                    for fi95 := 0; fi95 < 14; fi95++ {
                                        switch (fi95 - fi95) {
                                        case 2147483647:
                                            switch 9223372036854775807 {
                                            case 921:
                                                v4 += (v4 % ((v4) | 1))
                                            default:
                                                fmt.Println(v6)
                                            }
                                        case -353:
                                            for fi96 := 0; fi96 < 11; fi96++ {
                                                v5 -= (fi95 * (-(fi95)))
                                                _ = fi96
                                            }
                                        }
                                        if false {
                                            var v97 int8 = -60
                                            _ = v97
                                        }
                                        for fi98 := 0; fi98 < 7; fi98++ {
                                            var v99 bool = true
                                            var v100 uint16 = 255
                                            _ = fi98
                                            _ = v100
                                            _ = v99
                                        }
                                        _ = fi95
                                    }
                                } else {
                                    v20 += (v20 - v20)
                                    v20 -= v20
                                }
                            } else {
                                fmt.Println(v19, v5, v19)
                                var v101 int16 = (v20 - (v20 * 42))
                                _ = v101
                            }
                        }
                        if len(v1) >= 2 {
                            v1 = v1[1:len(v1)]
                        }
                    }
                    var v102 string = ""
                    _ = v102
                } else {
                    var v103 uint8 = (v6 - 10)
                    var v104 float64 = (1.5 * (0.1 / 1.0))
                    v5 -= 32767
                    _ = v103
                    _ = v104
                }
                v19 = (float32(-2.25) + ((((float32(1.5) - v19) + v19) - (1.5 * (-2.25 / v19))) - ((3.14159 + (3.14159 / 1.5)) - float32(3.14159))))
                _ = v23
            }
        }
        _ = v22
    }
    fmt.Println(v20, v19, v19)
    switch (v5 ^ (v5 + (v5 ^ (v5 / (((^(v5))) | 1))))) {
    case 127:
        for fi105 := 0; fi105 < 7; fi105++ {
            v106 := GenS11{F0: 9223372036854775807}
            if true {
                switch "\n\t" {
                case "hello":
                    v107 := GenS9{F0: (v5 >> ((1) & 63)), F1: 603}
                    _ = v107
                case "a":
                    for fi108 := 0; fi108 < 8; fi108++ {
                        v20 -= v20
                        _ = fi108
                    }
                }
            } else {
                var v109 bool = (42 != 2147483648)
                fmt.Println(v109, v19, v109)
                _ = v109
            }
            _ = fi105
            _ = v106
        }
    case 32768:
        v110 := GenS11{F0: 9223372036854775807}
        _ = v110
    case 128:
        var v111 uint64 = 42
        _ = v111
    default:
        var v112 int = (^(v5))
        _ = v112
    }
    if (18446744073709551615 != 4294967295) {
        fmt.Println(v21, v21, v20)
        var v113 int = genFn1(v21, 0)
        v19 -= ((v19 - v19) + ((float32((-2.25 + 0.1)) - v19) * v19))
        _ = v113
    }
    v114 := GenS11{F0: -523}
    var v115 uint = 1
    v4 = v4
    v6 = (^(v21))
    v116 := GenS11{F0: 9223372036854775807}
    v117 := GenS12{F0: false, F1: 42}
    switch (v5 - v5) {
    case 42:
        var v118 uint64 = 251
        _ = v118
    case -935:
        v119 := GenS5{F0: 2147483647, F1: v4, F2: ("a" + "a")}
        _ = v119
    case 128:
        var v120 float64 = ((1e10 / 1.0) * float64(v19))
        _ = v120
    default:
        v19 = (((v19 / v19) + ((v19 * (v19 - v19)) - (float32(0.1) / v19))) / 1.0)
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v115:", v115)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v114
    _ = v115
    _ = v116
    _ = v117
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
