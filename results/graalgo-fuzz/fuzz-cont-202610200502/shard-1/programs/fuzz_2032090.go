package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 uint64; F1 uint; F2 string }
type GenS4 struct { F0 uint32; F1 uint32 }
type GenS5 struct { F0 string }
type GenS6 struct { F0 uint; F1 int16; F2 int }
type GenS7 struct { F0 uint8; F1 uint8 }
type GenS8 struct { F0 string; F1 uint16 }
type GenS9 struct { F0 int64; F1 int64 }
type GenS10 struct { F0 string; F1 int }
type GenS11 struct { F0 uint; F1 uint; F2 int }
func genFn0() int64 {
    for fi6 := 0; fi6 < 8; fi6++ {
        switch "a" {
        case "\n\t":
            fmt.Println(fi6)
        case "hello":
            var v7 uint = 42
            _ = v7
        case "a":
            var v8 int32 = -1
            _ = v8
        default:
            v9 := GenS1{F0: 266}
            _ = v9
        }
        var v10 int64 = 1
        _ = fi6
        _ = v10
    }
    for fi11 := 0; fi11 < 16; fi11++ {
        var v12 int32 = 661
        fmt.Println(fi11)
        switch "hello" {
        case "":
            if (!true) {
                if false {
                    v13 := GenS2{F0: 0}
                    _ = v13
                }
                switch "\n\t" {
                case "a":
                    for fi14 := 0; fi14 < 2; fi14++ {
                        var v15 uint32 = 65535
                        fmt.Println(v12, v12)
                        fmt.Println(v15, v15, v15)
                        _ = fi14
                        _ = v15
                    }
                case "":
                    var v16 int8 = 100
                    _ = v16
                case "\n\t":
                    var v17 uint = 255
                    _ = v17
                }
            }
        case "\n\t":
            fmt.Println(fi11)
        case "hello":
            if (true && (3.14159 == 1.5)) {
                var v18 int64 = 0
                for fi19 := 0; fi19 < 20; fi19++ {
                    switch "hello" {
                    case "foo bar":
                        var v20 string = ("" + "hello")
                        _ = v20
                    case "a":
                        v18 += v18
                    default:
                        var v21 float64 = 1.5
                        _ = v21
                    }
                    switch "a" {
                    case "\n\t":
                        v12 -= v12
                    default:
                        v22 := []int{1}
                        _ = v22
                    }
                    _ = fi19
                }
                _ = v18
            } else {
                v23 := []int{}
                v24 := GenS3{F0: 42, F1: 4294967295, F2: "hello"}
                _ = v23
                _ = v24
            }
        }
        _ = fi11
        _ = v12
    }
    var v25 uint8 = 0
    v25 = (^(v25))
    v25 += (v25 / ((v25) | 1))
    _ = v25
    return -930
}
func genFn1(p0 string, p1 string, p2 string) int {
    if true {
        v26 := []int{692, -761, 1}
        for fi27 := 0; fi27 < 6; fi27++ {
            switch fi27 {
            case 888:
                for fi28 := 0; fi28 < 5; fi28++ {
                    if true {
                        var v29 uint8 = 42
                        p1 = p0
                        if ((42 == 171) || (!true)) {
                            switch (p1 + p0) {
                            case "\n\t":
                                p2 = "foo bar"
                            default:
                                var v30 uint16 = uint16(v29)
                                _ = v30
                            }
                            var v31 string = ((p1 + "foo bar") + p2)
                            _ = len(v26)
                            _ = v31
                        }
                        _ = v29
                    }
                    var v32 uint16 = 156
                    v32 = (^(v32))
                    _ = fi28
                    _ = v32
                }
            }
            _ = fi27
        }
        for fi33 := 0; fi33 < 7; fi33++ {
            for fi34 := 0; fi34 < 14; fi34++ {
                if true {
                    p2 = ("hello" + p2)
                } else {
                    var v35 uint64 = 42
                    p0 = ""
                    if false {
                        var v36 string = p2
                        if (42 <= 496) {
                            if len(v26) >= 2 {
                                v26 = v26[1:len(v26)]
                            }
                            v37 := GenS4{F0: 45, F1: 665}
                            _ = v37
                        } else {
                            v26 = append(v26, (fi33 >> ((2) & 63)))
                        }
                        p1 = v36
                        _ = v36
                    }
                    _ = v35
                }
                if true {
                    p2 = ((("" + "hello") + p0) + ("foo bar" + ""))
                    var v38 uint = 42
                    if (-128 <= 1) {
                        var v39 int16 = 1000
                        fmt.Println(fi34, v39)
                        _ = v39
                    }
                    if false {
                        fmt.Println(p2, p0, fi34)
                        if false {
                            if true {
                                var v40 int8 = 78
                                var v41 int32 = -945
                                _ = v40
                                _ = v41
                            } else {
                                fmt.Println(v38, p0)
                                var v42 float32 = 3.14159
                                _ = v42
                            }
                        }
                        if false {
                            v43 := GenS3{F0: 2, F1: (v38 * v38), F2: "foo bar"}
                            v44 := GenS4{F0: 677, F1: 42}
                            _ = v43
                            _ = v44
                        }
                    } else {
                        switch 32767 {
                        case 2147483647:
                            v45 := GenS5{F0: p2}
                            _ = v45
                        case 127:
                            v38 = (^(v38))
                        }
                        v38 += (v38 | v38)
                    }
                    if false {
                        p1 = (p0 + "a")
                    } else {
                        v46 := GenS6{F0: (v38 + v38), F1: 1000, F2: fi34}
                        _ = v46
                    }
                    p2 = (p1 + p0)
                    if len(v26) > 0 {
                        v26[0] = -975
                    }
                    _ = v38
                }
                p2 = p2
                _ = fi34
            }
            var v47 uint64 = 4294967295
            _ = fi33
            _ = v47
        }
        _ = v26
    } else {
        p1 = p2
        for fi48 := 0; fi48 < 11; fi48++ {
            for fi49 := 0; fi49 < 16; fi49++ {
                v50 := GenS7{F0: 1, F1: 1}
                _ = fi49
                _ = v50
            }
            if false {
                var v51 float32 = ((1e10 - -2.25) + 0.0)
                _ = v51
            }
            _ = fi48
        }
        v52 := []int{}
        _ = v52
    }
    var v53 uint64 = 18446744073709551615
    _ = p0
    _ = p1
    _ = p2
    _ = v53
    return -175
}

func main() {
    var v1 int16 = -3
    var v2 int8 = 100
    v3 := make(map[int]string)
    v4 := []string{}
    var v5 int16 = v1
    var v54 string = "foo bar"
    v5 = (v5 - v5)
    for fi55 := 0; fi55 < 12; fi55++ {
        switch (v54 + v54) {
        case "hello":
            v1 -= (v1 ^ 32767)
        case "":
            switch (fi55 >> ((18446744073709551615) & 63)) {
            case -9223372036854775808:
                for fi56 := 0; fi56 < 3; fi56++ {
                    fmt.Println(v2, v54, fi55)
                    fmt.Println(fi56)
                    _ = fi56
                }
            case -2147483648:
                var v57 uint8 = 75
                _ = v57
            case 1:
                v58 := GenS3{F0: 18446744073709551615, F1: 2, F2: "hello"}
                _ = v58
            }
        }
        fmt.Println(fi55, fi55, v2)
        _ = fi55
    }
    if true {
        if (9223372036854775807 == 9223372036854775807) {
            var v59 bool = (1.5 <= 0.1)
            switch "hello" {
            case "hello":
                v2 += 42
            case "":
                for fi60 := 0; fi60 < 14; fi60++ {
                    fmt.Println(fi60, v59, v1)
                    _ = fi60
                }
            case "foo bar":
                v5 += (v5 ^ v5)
            }
            _ = v59
        }
        if false {
            var v61 int8 = (v2 >> ((65535) & 63))
            var v62 uint8 = 1
            var v63 float32 = ((1e10 * (float32(0.1) + float32(0.0))) + (0.1 * 0.1))
            _ = v61
            _ = v62
            _ = v63
        }
        var v64 uint8 = 79
        _ = v64
    }
    fmt.Println(v54)
    var v65 float64 = 1.5
    v2 += v2
    fmt.Println(v54, v65, v65)
    if (false && (-9223372036854775808 < -210)) {
        delete(v3, -441)
    } else {
        var v66 int8 = -128
        var v67 int32 = 2147483647
        _ = v66
        _ = v67
    }
    for fi68 := 0; fi68 < 2; fi68++ {
        if false {
            fmt.Println(v65, v65, v5)
            if (true && true) {
                if (true || false) {
                    switch (("\n\t" + ("\n\t" + v54)) + (v54 + (v54 + "foo bar"))) {
                    case "foo bar":
                        if (!(1 == 42)) {
                            var v69 uint32 = 717
                            _ = v69
                        } else {
                            v54 = v54
                            if false {
                                v70 := GenS2{F0: v5}
                                var v71 string = (v54 + v54)
                                for fi72 := 0; fi72 < 4; fi72++ {
                                    if true {
                                        v73 := GenS8{F0: v71, F1: 200}
                                        fmt.Println(v2)
                                        _ = v73
                                    } else {
                                        var v74 uint32 = 1
                                        v75 := GenS5{F0: (v54 + "foo bar")}
                                        switch v54 {
                                        case "foo bar":
                                            v65 += v65
                                        }
                                        _ = v74
                                        _ = v75
                                    }
                                    _ = fi72
                                }
                                _ = v70
                                _ = v71
                            } else {
                                var v76 float64 = v65
                                switch (fi68 >> ((255) & 63)) {
                                case 4294967295:
                                    var v77 uint8 = 1
                                    _ = v77
                                case 9223372036854775807:
                                    _ = len(v3)
                                case -1:
                                    switch ((("\n\t" + (("\n\t" + ((v54 + (v54 + "")) + "foo bar")) + v54)) + (("hello" + (v54 + v54)) + "\n\t")) + v54) {
                                    case "foo bar":
                                        v1 = -211
                                    default:
                                        delete(v3, fi68)
                                    }
                                default:
                                    if (!(!false)) {
                                        v54 = ""
                                        if false {
                                            v1 += (v5 * (^(v5)))
                                        }
                                    } else {
                                        v2 += (v2 - 0)
                                        switch -2147483648 {
                                        case 564:
                                            v1 += v5
                                        case 2147483647:
                                            for fi78 := 0; fi78 < 16; fi78++ {
                                                v79 := GenS1{F0: 644}
                                                _ = fi78
                                                _ = v79
                                            }
                                        case 32768:
                                            v65 -= v65
                                        default:
                                            fmt.Println(v1)
                                        }
                                        for fi80 := 0; fi80 < 8; fi80++ {
                                            fmt.Println(fi68, v5, v65)
                                            _ = fi80
                                        }
                                    }
                                }
                                _ = v76
                            }
                        }
                    case "hello":
                        if ((-2147483648 > -2147483648) && ((((v65 / float64(1e10)) * v65) - (v65 + 3.14159)) >= ((1e10 / v65) - v65))) {
                            if false {
                                if true {
                                    switch "" {
                                    case "a":
                                        var v81 int16 = (v5 % (((v5 + v5)) | 1))
                                        _ = v81
                                    }
                                }
                            } else {
                                if false {
                                    var v82 float64 = v65
                                    _ = v82
                                } else {
                                    var v83 int8 = v2
                                    _ = v83
                                }
                            }
                        } else {
                            var v84 string = v54
                            v5 = (v1 ^ v5)
                            _ = v84
                        }
                    case "\n\t":
                        v65 += (v65 / (v65 / (3.14159 / (float64(0.0) * (0.1 / v65)))))
                    default:
                        var v85 uint8 = 26
                        _ = v85
                    }
                } else {
                    v4 = append(v4, (v54 + v54))
                    var v86 uint32 = 42
                    v87 := GenS9{F0: 590, F1: -141}
                    _ = v86
                    _ = v87
                }
                for fi88 := 0; fi88 < 1; fi88++ {
                    var v89 int32 = 266
                    var v90 int = fi88
                    v91 := GenS10{F0: v54, F1: fi88}
                    _ = fi88
                    _ = v89
                    _ = v90
                    _ = v91
                }
            } else {
                if (false || true) {
                    switch "\n\t" {
                    case "hello":
                        v54 = v54
                    }
                    v92 := GenS11{F0: 65535, F1: 0, F2: genFn1("hello", ((((v54 + v54) + v54) + v54) + (v54 + ((v54 + "\n\t") + "a"))), v54)}
                    v5 = 1000
                    _ = v92
                } else {
                    var v93 uint16 = 255
                    var v94 int32 = -2147483648
                    fmt.Println(v2, v2)
                    _ = v93
                    _ = v94
                }
                v5 -= v5
                v3[-514] = v54
            }
        } else {
            fmt.Println(fi68, fi68)
        }
        for fi95 := 0; fi95 < 14; fi95++ {
            v2 -= (^(v2))
            v5 = (v5 + v1)
            _ = fi95
        }
        v96 := GenS8{F0: "a", F1: 255}
        _ = fi68
        _ = v96
    }
    var v97 bool = true
    switch v54 {
    case "\n\t":
        v65 -= 0.0
    default:
        if ((v54 != v54) && (18446744073709551615 <= 0)) {
            v97 = (!false)
            var v98 float32 = (3.14159 - (0.1 + -2.25))
            v99 := GenS3{F0: 2, F1: 255, F2: "hello"}
            _ = v98
            _ = v99
        } else {
            if (4294967295 == 2) {
                fmt.Println(v2)
            } else {
                switch v54 {
                case "a":
                    fmt.Println(v54)
                default:
                    v100 := GenS11{F0: 255, F1: 855, F2: 0}
                    _ = v100
                }
                v4 = append(v4, "hello")
                var v101 int = 0
                _ = v101
            }
            v2 = v2
        }
    }
    delete(v3, 0)
    for fi102 := 0; fi102 < 18; fi102++ {
        var v103 int32 = 42
        _ = fi102
        _ = v103
    }
    v3[-856] = (("foo bar" + "") + v54)
    var v104 int64 = 193
    var v105 uint = 4294967295
    fmt.Println("v1:", v1)
    fmt.Println("v104:", v104)
    fmt.Println("v105:", v105)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v54:", v54)
    fmt.Println("v65:", v65)
    fmt.Println("v97:", v97)
    _ = v1
    _ = v104
    _ = v105
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v54
    _ = v65
    _ = v97
}
