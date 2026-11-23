package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int32 }
type GenS3 struct { F0 int32; F1 int16; F2 int64 }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint8; F1 int32; F2 int64 }
type GenS7 struct { F0 string; F1 bool }
type GenS8 struct { F0 uint16; F1 bool }
func genFn0(p0 bool, p1 int32, p2 uint16) string {
    switch "a" {
    case "":
        p0 = (!(!p0))
    case "\n\t":
        v6 := []int{1, 127, 9223372036854775807, 32767}
        _ = v6
    case "foo bar":
        fmt.Println(p0, p1)
    }
    switch 32767 {
    case -917:
        v7 := []int{4294967295}
        _ = v7
    default:
        v8 := GenS1{F0: 65535}
        _ = v8
    }
    _ = p0
    _ = p1
    _ = p2
    return "foo bar"
}
func genFn1(p0 int, p1 string, p2 int64) int {
    v9 := GenS1{F0: 2147483647}
    for fi10 := 0; fi10 < 5; fi10++ {
        for fi11 := 0; fi11 < 2; fi11++ {
            p0 = 1
            _ = fi11
        }
        p0 -= -1
        _ = fi10
    }
    if (((((p1 + "\n\t") + p1) + (p1 + "a")) + p1) > p1) {
        p1 = p1
        fmt.Println(p2, p2)
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return p0
}

func main() {
    v1 := []string{("foo bar" + ("\n\t" + ("" + "hello"))), "a"}
    v2 := make(map[string]uint32)
    var v3 int64 = 562
    v4 := []uint32{2147483648, 4294967295, 862}
    v5 := make(map[uint32]bool)
    fmt.Println(v3)
    var v12 float32 = 3.14159
    for fi13 := 0; fi13 < 9; fi13++ {
        v5[50] = false
        fmt.Println(fi13, fi13)
        _ = fi13
    }
    for fi14 := 0; fi14 < 18; fi14++ {
        for fi15 := 0; fi15 < 14; fi15++ {
            var v16 uint = 0
            v5[690] = (9223372036854775807 == -420)
            _ = fi15
            _ = v16
        }
        var v17 uint = 42
        _ = fi14
        _ = v17
    }
    switch "foo bar" {
    case "a":
        if true {
            fmt.Println(v3, v3)
            for fi18 := 0; fi18 < 12; fi18++ {
                if true {
                    switch (fi18 / (((-(fi18))) | 1)) {
                    case 128:
                        for fi19 := 0; fi19 < 1; fi19++ {
                            if false {
                                v12 = v12
                                if false {
                                    v5[65535] = (!true)
                                }
                                switch "a" {
                                case "foo bar":
                                    fmt.Println(v12, fi19, v3)
                                }
                            } else {
                                if (18446744073709551615 <= 4294967295) {
                                    v3 += (^(v3))
                                    var v20 uint64 = 0
                                    if len(v4) > 0 {
                                        v4[0] = 4294967295
                                    }
                                    var v21 bool = ((1e10 != 3.14159) && (("foo bar" + ("a" + "\n\t")) != "hello"))
                                    _ = v20
                                    _ = v21
                                } else {
                                    if (fi19 >= fi18) {
                                        v22 := GenS1{F0: 1}
                                        _ = v22
                                    }
                                }
                            }
                            v23 := v2["\n\t"]
                            delete(v2, "a")
                            if (false && (!(380 <= 1))) {
                                var v24 uint = 18446744073709551615
                                v24 += v24
                                v24 = (v24 / (((v24 << (((v24 % (((v24 % ((v24) | 1))) | 1))) & 63))) | 1))
                                v23 += 579
                                _ = v24
                            } else {
                                v12 = (float32(3.14159) / v12)
                                var v25 float32 = float32((0.0 + (0.1 - float64((1.5 * 1e10)))))
                                _ = len(v5)
                                _ = v25
                            }
                            var v26 float32 = v12
                            _ = fi19
                            _ = v23
                            _ = v26
                        }
                    case -637:
                        var v27 int = fi18
                        _ = v27
                    default:
                        v4 = append(v4, 65535)
                    }
                    v12 = (((v12 / ((v12 * 0.1) - 1.5)) / v12) * v12)
                }
                _ = fi18
            }
        }
    case "foo bar":
        var v28 string = "hello"
        _ = v28
    default:
        if (!true) {
            v29 := GenS1{F0: 600}
            if false {
                for fi30 := 0; fi30 < 11; fi30++ {
                    for fi31 := 0; fi31 < 12; fi31++ {
                        switch fi31 {
                        case 1:
                            v12 = 0.0
                        }
                        if ((("a" + "\n\t") + "") <= ("a" + ((("\n\t" + "hello") + "foo bar") + ("foo bar" + "foo bar")))) {
                            if false {
                                var v32 int16 = -1
                                v1 = append(v1, "a")
                                if (!(true || true)) {
                                    var v33 uint64 = 42
                                    _ = v33
                                }
                                _ = v32
                            }
                            var v34 bool = false
                            _ = v34
                        } else {
                            var v35 uint32 = 2147483647
                            _ = v35
                        }
                        v3 += v3
                        _ = fi31
                    }
                    if (0 != 1) {
                        var v36 int32 = 602
                        _ = v36
                    } else {
                        for fi37 := 0; fi37 < 2; fi37++ {
                            v5[720] = true
                            _ = fi37
                        }
                        if false {
                            v12 -= v12
                            v3 -= 42
                            var v38 int = 708
                            _ = v38
                        } else {
                            fmt.Println(v12)
                        }
                    }
                    _ = fi30
                }
            }
            _ = v29
        } else {
            _ = len(v2)
        }
    }
    for fi39 := 0; fi39 < 4; fi39++ {
        fmt.Println(fi39, fi39, v3)
        var v40 int = fi39
        _ = fi39
        _ = v40
    }
    for fi41 := 0; fi41 < 5; fi41++ {
        var v42 uint8 = 20
        v43 := GenS2{F0: 1}
        _ = fi41
        _ = v42
        _ = v43
    }
    var v44 int = 203
    if (!false) {
        for fi45 := 0; fi45 < 12; fi45++ {
            for fi46 := 0; fi46 < 11; fi46++ {
                v47 := GenS2{F0: 2147483647}
                switch "\n\t" {
                case "a":
                    v44 -= (-(fi45))
                case "hello":
                    v48 := GenS3{F0: 1, F1: 1, F2: (^(v3))}
                    _ = v48
                case "foo bar":
                    var v49 int8 = 127
                    _ = v49
                }
                if (true && (-1 >= 42)) {
                    var v50 uint = 723
                    v51 := GenS1{F0: 0}
                    _ = v50
                    _ = v51
                }
                v3 += (v3 >> ((668) & 63))
                _ = fi46
                _ = v47
            }
            v44 = genFn1(900, "hello", 42)
            v44 = (fi45 % ((fi45) | 1))
            _ = fi45
        }
        v3 += v3
        v52 := GenS2{F0: 2147483647}
        _ = v52
    } else {
        delete(v2, "hello")
        for fi53 := 0; fi53 < 3; fi53++ {
            for fi54 := 0; fi54 < 15; fi54++ {
                switch ("hello" + "foo bar") {
                case "foo bar":
                    var v55 float64 = 3.14159
                    _ = v55
                case "a":
                    v3 += v3
                }
                if true {
                    var v56 bool = false
                    var v57 bool = (v56 && ((100 < -63) || ((!true) || (!v56))))
                    var v58 int8 = 100
                    _ = v56
                    _ = v57
                    _ = v58
                }
                var v59 int8 = -84
                v59 = v59
                if (true && false) {
                    fmt.Println(v12)
                    if (!(777 <= 42)) {
                        v4 = append(v4, 712)
                    } else {
                        if (!false) {
                            v3 += (v3 - v3)
                            _ = len(v1)
                            v12 = float32(float64((((float32(((0.0 * -2.25) + 0.1)) * (v12 * (1e10 - v12))) / v12) * (v12 / v12))))
                            _ = len(v4)
                            if (65535 >= 4294967295) {
                                if (!false) {
                                    if false {
                                        v44 = genFn1((v44 / (((fi53 ^ genFn1(fi53, ("" + (((("" + "a") + ("\n\t" + ("foo bar" + "\n\t"))) + "\n\t") + "a")), 96))) | 1)), ("" + "a"), v3)
                                        v60 := GenS3{F0: 1, F1: -936, F2: (v3 >> ((18446744073709551615) & 63))}
                                        var v61 float32 = v12
                                        _ = v60
                                        _ = v61
                                    }
                                    if ((0 <= 42) && (false && (int64(v59) >= v3))) {
                                        var v62 uint16 = 0
                                        fmt.Println(v62, v44, v59)
                                        var v63 int64 = v3
                                        _ = v62
                                        _ = v63
                                    } else {
                                        var v64 int64 = v3
                                        _ = v64
                                    }
                                    if false {
                                        if (("\n\t" + "hello") > "hello") {
                                            v5[2147483647] = ((!true) || false)
                                        } else {
                                            if true {
                                                switch ((((("a" + (("foo bar" + "\n\t") + "a")) + "a") + (("" + ((((("" + ("hello" + "hello")) + "\n\t") + "hello") + "foo bar") + "foo bar")) + "")) + "foo bar") + "hello") {
                                                case "foo bar":
                                                    fmt.Println(v59, v3, fi53)
                                                case "\n\t":
                                                    v65 := GenS3{F0: 42, F1: 1000, F2: v3}
                                                    _ = v65
                                                case "hello":
                                                    var v66 uint32 = 995
                                                    _ = v66
                                                default:
                                                    if false {
                                                        _ = len(v1)
                                                        if (!(0.0 >= 1e10)) {
                                                            v67 := v2["\n\t"]
                                                            fmt.Println(fi53)
                                                            _ = v67
                                                        } else {
                                                            switch ("a" + "hello") {
                                                            case "\n\t":
                                                                _ = len(v4)
                                                            case "":
                                                                switch (fi54 & fi54) {
                                                                case -2147483648:
                                                                    switch "" {
                                                                    case "foo bar":
                                                                        v12 = v12
                                                                    case "hello":
                                                                        if true {
                                                                            v68 := GenS2{F0: 2147483647}
                                                                            v3 = v3
                                                                            _ = v68
                                                                        }
                                                                    default:
                                                                        if true {
                                                                            v59 = v59
                                                                            v2["foo bar"] = 228
                                                                        }
                                                                    }
                                                                }
                                                            default:
                                                                var v69 float64 = 3.14159
                                                                _ = v69
                                                            }
                                                            var v70 float32 = v12
                                                            _ = v70
                                                        }
                                                        fmt.Println(v44)
                                                    }
                                                }
                                            } else {
                                                var v71 bool = ((65535 != 2147483648) && (-886 != 83))
                                                var v72 uint32 = 42
                                                _ = v71
                                                _ = v72
                                            }
                                            if true {
                                                switch "foo bar" {
                                                case "foo bar":
                                                    v59 += (^(v59))
                                                    var v73 int64 = (v3 << ((255) & 63))
                                                    _ = v73
                                                case "hello":
                                                    var v74 bool = true
                                                    _ = v74
                                                }
                                                v75 := GenS3{F0: -2147483648, F1: -1, F2: (v3 + v3)}
                                                v76 := GenS4{F0: 42}
                                                _ = v75
                                                _ = v76
                                            } else {
                                                fmt.Println(v3)
                                                fmt.Println(fi53)
                                                v3 = -1
                                            }
                                            switch 42 {
                                            case 32768:
                                                v3 = (v3 | v3)
                                            }
                                        }
                                        var v77 int16 = 1
                                        fmt.Println(v77, v59)
                                        _ = v77
                                    } else {
                                        var v78 int = genFn1(fi53, "\n\t", (v3 ^ v3))
                                        v78 = genFn1((v44 + v44), ("\n\t" + "hello"), (v3 << ((4294967295) & 63)))
                                        _ = v78
                                    }
                                    fmt.Println(v12, v59, v12)
                                } else {
                                    var v79 uint64 = 2
                                    delete(v2, ("\n\t" + ("foo bar" + "\n\t")))
                                    v12 -= v12
                                    _ = v79
                                }
                            } else {
                                if false {
                                    _ = len(v5)
                                    var v80 float64 = 1.5
                                    var v81 string = "a"
                                    _ = v80
                                    _ = v81
                                } else {
                                    var v82 int16 = 32767
                                    _ = v82
                                }
                                var v83 int64 = v3
                                _ = v83
                            }
                            if false {
                                v3 += (v3 & v3)
                                switch "a" {
                                case "\n\t":
                                    if (-1 > 2147483647) {
                                        switch "foo bar" {
                                        case "foo bar":
                                            switch ((("foo bar" + "") + "a") + ((((((((("hello" + "\n\t") + "foo bar") + "foo bar") + ("\n\t" + "hello")) + "") + "foo bar") + "hello") + "") + ("\n\t" + "foo bar"))) {
                                            case "hello":
                                                switch "hello" {
                                                case "":
                                                    if true {
                                                        delete(v5, 488)
                                                    } else {
                                                        var v84 float32 = 0.1
                                                        fmt.Println(v12)
                                                        _ = v84
                                                    }
                                                case "\n\t":
                                                    var v85 uint64 = 42
                                                    _ = v85
                                                case "foo bar":
                                                    if ((390 >= 2147483648) && false) {
                                                        switch "\n\t" {
                                                        case "hello":
                                                            v86 := GenS5{F0: 0}
                                                            _ = v86
                                                        case "foo bar":
                                                            v12 = 3.14159
                                                        }
                                                    }
                                                default:
                                                    v3 = 915
                                                }
                                            case "\n\t":
                                                _ = len(v5)
                                            default:
                                                var v87 bool = true
                                                _ = v87
                                            }
                                        }
                                    } else {
                                        v59 = v59
                                    }
                                case "hello":
                                    var v88 int64 = (v3 - (^(v3)))
                                    _ = v88
                                }
                            } else {
                                if (!true) {
                                    var v89 int8 = (v59 * -1)
                                    v90 := GenS5{F0: 290}
                                    v91 := GenS6{F0: 0, F1: 42, F2: (v3 - 1)}
                                    _ = v89
                                    _ = v90
                                    _ = v91
                                }
                                fmt.Println(fi54, fi54)
                            }
                        } else {
                            v92 := GenS6{F0: 255, F1: 2147483647, F2: -1}
                            switch 32767 {
                            case 128:
                                if false {
                                    var v93 string = "\n\t"
                                    switch (^(v44)) {
                                    case -1:
                                        var v94 uint16 = 58
                                        _ = v94
                                    case 2147483647:
                                        _ = len(v1)
                                    case -9223372036854775808:
                                        switch (v44 % ((genFn1((fi54 << ((137) & 63)), v93, v3)) | 1)) {
                                        case -1:
                                            fmt.Println(v44, fi53, v59)
                                        default:
                                            switch (v93 + v93) {
                                            case "foo bar":
                                                if true {
                                                    if true {
                                                        if true {
                                                            if true {
                                                                var v95 int8 = 0
                                                                if len(v1) >= 2 {
                                                                    v1 = v1[1:len(v1)]
                                                                }
                                                                _ = v95
                                                            } else {
                                                                v59 += v59
                                                                var v96 uint16 = 255
                                                                _ = v96
                                                            }
                                                            var v97 float32 = float32(((0.0 * (-2.25 - (0.1 - 1.5))) * (((0.1 - (0.0 + 3.14159)) * (((0.1 * (float64(v12) / 3.14159)) + (1.5 * 0.0)) + (-2.25 * (1.5 * 1.5)))) - 0.1)))
                                                            if (!true) {
                                                                switch 128 {
                                                                case -893:
                                                                    var v98 bool = true
                                                                    _ = v98
                                                                case 128:
                                                                    fmt.Println(v12, v93, v12)
                                                                case 4294967295:
                                                                    switch v93 {
                                                                    case "\n\t":
                                                                        if false {
                                                                            fmt.Println(v3, fi53)
                                                                            v99 := GenS7{F0: "a", F1: false}
                                                                            _ = v99
                                                                        }
                                                                    case "hello":
                                                                        if (false && (683 == 18446744073709551615)) {
                                                                            var v100 uint = 771
                                                                            _ = v100
                                                                        }
                                                                    }
                                                                }
                                                            } else {
                                                                v97 = v97
                                                                v97 = 3.14159
                                                                v93 = (("hello" + v93) + v93)
                                                            }
                                                            v59 = v59
                                                            _ = v97
                                                        }
                                                    } else {
                                                        fmt.Println(v3)
                                                    }
                                                    var v101 int = fi53
                                                    v3 -= v3
                                                    _ = v101
                                                }
                                            }
                                        }
                                        v59 -= v59
                                    default:
                                        var v102 int32 = 2147483647
                                        _ = v102
                                    }
                                    _ = v93
                                }
                            default:
                                fmt.Println(v44, fi53, fi54)
                            }
                            _ = v92
                        }
                        v103 := GenS8{F0: 65535, F1: false}
                        if true {
                            var v104 float64 = ((0.0 + 1e10) + (1.5 / (0.1 - (float64(v12) * -2.25))))
                            delete(v2, ("\n\t" + ("\n\t" + "")))
                            var v105 uint8 = 2
                            _ = v104
                            _ = v105
                        }
                        _ = v103
                    }
                    v12 = -2.25
                } else {
                    v106 := GenS1{F0: 42}
                    v59 = (v59 << ((65535) & 63))
                    _ = v106
                }
                _ = fi54
                _ = v59
            }
            var v107 uint64 = 0
            _ = fi53
            _ = v107
        }
    }
    v44 = genFn1(0, (((("a" + ("foo bar" + "")) + "") + ("hello" + "")) + "foo bar"), (v3 << ((4294967295) & 63)))
    fmt.Println(v12)
    _ = len(v2)
    for fi108 := 0; fi108 < 5; fi108++ {
        _ = len(v4)
        var v109 uint = 204
        v109 += v109
        _ = fi108
        _ = v109
    }
    var v110 int32 = 359
    v111 := GenS6{F0: 128, F1: (v110 << ((2) & 63)), F2: (v3 ^ (v3 - (v3 / ((1) | 1))))}
    v44 = v44
    fmt.Println(v44)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v110:", v110)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v44:", v44)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v110
    _ = v111
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v44
    _ = v5
}
