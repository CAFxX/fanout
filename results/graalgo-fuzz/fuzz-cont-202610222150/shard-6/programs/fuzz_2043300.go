package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 string; F2 uint64 }
type GenS2 struct { F0 uint32; F1 bool; F2 int8 }
type GenS3 struct { F0 uint64; F1 uint8 }
type GenS4 struct { F0 int64; F1 uint32 }
type GenS5 struct { F0 int16; F1 uint16; F2 uint64 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int; F1 uint16 }
type GenS8 struct { F0 uint16; F1 uint; F2 int64 }
type GenS9 struct { F0 int; F1 string }

func main() {
    v1 := make(map[string]int)
    var v2 uint = 1
    var v3 float32 = 3.14159
    var v4 uint16 = 1
    delete(v1, "foo bar")
    var v5 uint32 = 2147483648
    var v6 float32 = float32(1e10)
    if true {
        for fi7 := 0; fi7 < 3; fi7++ {
            fmt.Println(v3, v3, fi7)
            v8 := []int{(fi7 << (((v2 % ((uint(v5)) | 1))) & 63))}
            _ = fi7
            _ = v8
        }
        if true {
            v9 := GenS1{F0: (v5 + v5), F1: (("hello" + "") + "\n\t"), F2: 18446744073709551615}
            _ = v9
        }
        for fi10 := 0; fi10 < 11; fi10++ {
            v2 -= v2
            for fi11 := 0; fi11 < 9; fi11++ {
                fmt.Println(v2)
                var v12 string = ("" + "a")
                fmt.Println(v5, v4, v5)
                v6 += v3
                fmt.Println(v2, v3, v12)
                if true {
                    v13 := GenS1{F0: v5, F1: v12, F2: 42}
                    _ = v13
                }
                v14 := []int{2147483647, (fi10 & fi10)}
                _ = fi11
                _ = v12
                _ = v14
            }
            _ = fi10
        }
    }
    for fi15 := 0; fi15 < 12; fi15++ {
        for fi16 := 0; fi16 < 7; fi16++ {
            if (-2.25 > 1.5) {
                v6 += float32(0.0)
            } else {
                switch ("a" + "foo bar") {
                case "foo bar":
                    _ = len(v1)
                case "":
                    delete(v1, "\n\t")
                default:
                    v4 += (v4 ^ 452)
                }
                var v17 int8 = 127
                _ = v17
            }
            _ = fi16
        }
        fmt.Println(v6, v6)
        var v18 uint = (v2 - v2)
        _ = fi15
        _ = v18
    }
    for fi19 := 0; fi19 < 13; fi19++ {
        v4 += v4
        if (v5 <= v5) {
            for fi20 := 0; fi20 < 20; fi20++ {
                var v21 int8 = -74
                var v22 float64 = 0.0
                var v23 uint8 = 127
                _ = fi20
                _ = v21
                _ = v22
                _ = v23
            }
            if (!(!true)) {
                var v24 uint = v2
                _ = v24
            } else {
                if ((((("" + "\n\t") + ("a" + ((("hello" + "") + ("a" + "hello")) + "a"))) + "foo bar") != "\n\t") || false) {
                    var v25 float64 = (0.0 / 0.1)
                    switch -120 {
                    case 1:
                        var v26 bool = true
                        _ = v26
                    default:
                        for fi27 := 0; fi27 < 9; fi27++ {
                            v28 := GenS2{F0: v5, F1: true, F2: 42}
                            switch ("\n\t" + "\n\t") {
                            case "\n\t":
                                if true {
                                    v29 := GenS2{F0: (^(v5)), F1: (!(0 <= 937)), F2: -128}
                                    _ = v29
                                } else {
                                    fmt.Println(v3, v6, v6)
                                    var v30 uint16 = (v4 * (^(v4)))
                                    _ = v30
                                }
                            case "a":
                                if (!true) {
                                    if true {
                                        v31 := []int{}
                                        if (9223372036854775807 > 1) {
                                            v5 = (v5 - v5)
                                            switch "a" {
                                            case "":
                                                v25 = v25
                                            case "\n\t":
                                                var v32 string = "\n\t"
                                                var v33 uint64 = 695
                                                _ = len(v31)
                                                _ = v32
                                                _ = v33
                                            case "a":
                                                fmt.Println(fi27, fi27, v25)
                                            }
                                            if false {
                                                var v34 int16 = int16(v2)
                                                v34 -= 603
                                                _ = v34
                                            } else {
                                                switch fi27 {
                                                case 127:
                                                    v6 = (v3 / v3)
                                                case 108:
                                                    v35 := GenS3{F0: 4294967295, F1: 255}
                                                    _ = v35
                                                }
                                                v36 := GenS1{F0: 80, F1: "", F2: 1}
                                                _ = len(v31)
                                                _ = v36
                                            }
                                        } else {
                                            var v37 uint8 = 155
                                            var v38 uint64 = uint64(v37)
                                            v38 -= v38
                                            _ = v37
                                            _ = v38
                                        }
                                        _ = v31
                                    }
                                } else {
                                    v39 := GenS2{F0: 88, F1: (!((!(2 <= uint64(v2))) && (-1 >= -1))), F2: 42}
                                    if ((false || false) && (!(42 < 32767))) {
                                        fmt.Println(fi27, v4)
                                    } else {
                                        if (-2.25 > v25) {
                                            var v40 int = fi27
                                            _ = v40
                                        }
                                        v4 += (v4 & v4)
                                    }
                                    v41 := GenS4{F0: -378, F1: v5}
                                    _ = v39
                                    _ = v41
                                }
                            default:
                                if false {
                                    v3 = (float32(v25) + 1.5)
                                    v42 := []int{(fi19 << ((v2) & 63)), 994, (fi27 % ((fi27) | 1))}
                                    v43 := GenS5{F0: 42, F1: v4, F2: 282}
                                    if (!(147 < 127)) {
                                        v44 := GenS1{F0: v5, F1: ("foo bar" + "foo bar"), F2: 519}
                                        _ = v44
                                    } else {
                                        if ((^(v5)) != 1) {
                                            fmt.Println(v5)
                                        }
                                        if (42 >= 9223372036854775807) {
                                            if false {
                                                if true {
                                                    switch "\n\t" {
                                                    case "":
                                                        var v45 string = (("" + (((("foo bar" + "a") + "a") + "") + "a")) + (("" + "\n\t") + ("a" + "hello")))
                                                        _ = v45
                                                    case "hello":
                                                        switch (fi19 % ((fi19) | 1)) {
                                                        case -135:
                                                            _ = len(v1)
                                                        case 2147483647:
                                                            fmt.Println(v4, v4, v3)
                                                        case 4294967295:
                                                            v4 += v4
                                                        default:
                                                            if (!(212 != uint8(v5))) {
                                                                switch (fi19 & fi27) {
                                                                case 42:
                                                                    if (!(!((761 < (fi27 % (((fi27 / ((32768) | 1))) | 1))) && false))) {
                                                                        if ((3.14159 / 1.0) <= v25) {
                                                                            v25 -= v25
                                                                            switch ("a" + ("hello" + "")) {
                                                                            case "foo bar":
                                                                                if (!false) {
                                                                                    var v46 int16 = 1000
                                                                                    v6 = ((0.0 - v3) - 0.1)
                                                                                    if (!((fi27 ^ -251) == 642)) {
                                                                                        _ = len(v1)
                                                                                    }
                                                                                    _ = v46
                                                                                } else {
                                                                                    switch fi19 {
                                                                                    case 0:
                                                                                        v25 += (((v25 - v25) / float64(float32((1e10 * v25)))) + v25)
                                                                                    case 32767:
                                                                                        if false {
                                                                                            var v47 uint16 = v4
                                                                                            if false {
                                                                                                switch ("\n\t" + ((("foo bar" + (("a" + "a") + "foo bar")) + ((("foo bar" + "hello") + "\n\t") + "hello")) + "")) {
                                                                                                case "":
                                                                                                    v3 = v3
                                                                                                default:
                                                                                                    v4 = v47
                                                                                                }
                                                                                                var v48 int16 = -1
                                                                                                _ = v48
                                                                                            } else {
                                                                                                fmt.Println(v4)
                                                                                            }
                                                                                            _ = v47
                                                                                        } else {
                                                                                            var v49 int8 = -128
                                                                                            _ = v49
                                                                                        }
                                                                                    case 625:
                                                                                        v2 += 42
                                                                                    default:
                                                                                        fmt.Println(v5, v2)
                                                                                    }
                                                                                }
                                                                                if true {
                                                                                    fmt.Println(v2, fi27, v6)
                                                                                    if true {
                                                                                        fmt.Println(fi19, fi19)
                                                                                        v3 -= v3
                                                                                    } else {
                                                                                        if (false || true) {
                                                                                            v50 := GenS3{F0: 18446744073709551615, F1: 1}
                                                                                            v51 := GenS3{F0: 154, F1: 128}
                                                                                            var v52 int16 = -68
                                                                                            _ = v50
                                                                                            _ = v51
                                                                                            _ = v52
                                                                                        } else {
                                                                                            fmt.Println(v3, v4, v25)
                                                                                            switch fi19 {
                                                                                            case 32768:
                                                                                                var v53 int = 128
                                                                                                _ = v53
                                                                                            case -1:
                                                                                                v5 -= (v5 % ((0) | 1))
                                                                                            case -797:
                                                                                                switch ("\n\t" + "a") {
                                                                                                case "":
                                                                                                    var v54 int8 = -128
                                                                                                    _ = v54
                                                                                                case "\n\t":
                                                                                                    v4 = (^(v4))
                                                                                                case "a":
                                                                                                    var v55 float64 = v25
                                                                                                    _ = v55
                                                                                                }
                                                                                            }
                                                                                            switch ("a" + "foo bar") {
                                                                                            case "":
                                                                                                v56 := GenS1{F0: (^(v5)), F1: ("a" + "\n\t"), F2: 163}
                                                                                                _ = v56
                                                                                            case "a":
                                                                                                switch fi19 {
                                                                                                case -2147483648:
                                                                                                    var v57 int64 = -94
                                                                                                    _ = v57
                                                                                                default:
                                                                                                    v2 = v2
                                                                                                }
                                                                                            case "hello":
                                                                                                if true {
                                                                                                    var v58 bool = true
                                                                                                    v59 := GenS2{F0: (v5 % ((v5) | 1)), F1: v58, F2: 1}
                                                                                                    _ = v58
                                                                                                    _ = v59
                                                                                                } else {
                                                                                                    v6 = v3
                                                                                                }
                                                                                            default:
                                                                                                v42 = append(v42, fi27)
                                                                                            }
                                                                                            var v60 int16 = 1000
                                                                                            _ = v60
                                                                                        }
                                                                                        var v61 uint = 876
                                                                                        _ = v61
                                                                                    }
                                                                                    switch "\n\t" {
                                                                                    case "":
                                                                                        var v62 float64 = ((0.0 - (((v25 - (v25 + (1e10 - (v25 / v25)))) + v25) * v25)) + v25)
                                                                                        if true {
                                                                                            if (128 == 42) {
                                                                                                if false {
                                                                                                    switch fi27 {
                                                                                                    case 128:
                                                                                                        _ = len(v42)
                                                                                                    case 127:
                                                                                                        fmt.Println(v25, v3, v6)
                                                                                                        var v63 bool = (true || false)
                                                                                                        _ = v63
                                                                                                    }
                                                                                                    if true {
                                                                                                        switch 32768 {
                                                                                                        case -2147483648:
                                                                                                            v4 += (v4 >> ((uint(v4)) & 63))
                                                                                                        default:
                                                                                                            fmt.Println(v62, v4, fi27)
                                                                                                        }
                                                                                                    } else {
                                                                                                        fmt.Println(v3)
                                                                                                        _ = len(v42)
                                                                                                    }
                                                                                                    v62 = 3.14159
                                                                                                    if (true && (2147483647 < -632)) {
                                                                                                        if true {
                                                                                                            switch fi19 {
                                                                                                            case 128:
                                                                                                                switch fi27 {
                                                                                                                case 4294967295:
                                                                                                                    v3 = 1.5
                                                                                                                default:
                                                                                                                    fmt.Println(fi19, v25)
                                                                                                                }
                                                                                                            }
                                                                                                            v62 = (v62 / ((1.5 - v62) - ((v25 - 0.1) + float64(((((v3 / 1.0) + 0.1) + v6) + 3.14159)))))
                                                                                                            v25 -= v62
                                                                                                            v6 = v6
                                                                                                            fmt.Println(v4)
                                                                                                            v42 = append(v42, 32768)
                                                                                                            switch ("a" + ("a" + "hello")) {
                                                                                                            case "hello":
                                                                                                                var v64 string = ("\n\t" + "foo bar")
                                                                                                                _ = v64
                                                                                                            case "a":
                                                                                                                var v65 float32 = float32(v62)
                                                                                                                _ = v65
                                                                                                            default:
                                                                                                                v3 -= (float32(v62) - (float32(((v25 - ((float64(v6) / 0.1) * (v62 - (v25 * (v62 / v25))))) * ((0.0 + (v62 + v62)) + (0.1 * v25)))) - 0.1))
                                                                                                            }
                                                                                                            v66 := v1[(("foo bar" + ("hello" + "hello")) + "hello")]
                                                                                                            var v67 int32 = 1
                                                                                                            _ = v66
                                                                                                            _ = v67
                                                                                                        } else {
                                                                                                            var v68 int64 = -809
                                                                                                            _ = v68
                                                                                                        }
                                                                                                        v6 += 0.0
                                                                                                    } else {
                                                                                                        if len(v42) > 0 {
                                                                                                            v42[0] = (fi27 & fi27)
                                                                                                        }
                                                                                                        v3 = (0.1 * v3)
                                                                                                        v69 := GenS5{F0: 0, F1: 192, F2: 2}
                                                                                                        var v70 int32 = 42
                                                                                                        v62 = (float64((v6 * float32(v25))) * 0.0)
                                                                                                        var v71 uint = v2
                                                                                                        _ = v69
                                                                                                        _ = v70
                                                                                                        _ = v71
                                                                                                    }
                                                                                                } else {
                                                                                                    switch (fi19 & -2147483648) {
                                                                                                    case -2147483648:
                                                                                                        v62 = 0.0
                                                                                                    default:
                                                                                                        v72 := GenS2{F0: (v5 - v5), F1: false, F2: 0}
                                                                                                        v2 = 42
                                                                                                        var v73 uint32 = v5
                                                                                                        if len(v42) > 0 {
                                                                                                            v42[0] = (^(fi19))
                                                                                                        }
                                                                                                        _ = v72
                                                                                                        _ = v73
                                                                                                    }
                                                                                                    fmt.Println(v5, v4)
                                                                                                    v74 := GenS6{F0: (v5 / ((v5) | 1))}
                                                                                                    _ = v74
                                                                                                }
                                                                                                switch 24 {
                                                                                                case 921:
                                                                                                    v75 := GenS7{F0: fi19, F1: v4}
                                                                                                    _ = v75
                                                                                                case -2147483648:
                                                                                                    if (!false) {
                                                                                                        var v76 uint32 = v5
                                                                                                        _ = v76
                                                                                                    }
                                                                                                case 127:
                                                                                                    var v77 uint64 = 2
                                                                                                    _ = v77
                                                                                                }
                                                                                                if len(v42) >= 2 {
                                                                                                    v42 = v42[1:len(v42)]
                                                                                                }
                                                                                                var v78 uint16 = v4
                                                                                                var v79 uint64 = 18446744073709551615
                                                                                                _ = v78
                                                                                                _ = v79
                                                                                            }
                                                                                            v6 = ((v3 + (v6 * float32((-2.25 - v62)))) * (((float32(v25) + 1e10) - (((((float32((v62 * v62)) / 1e10) + -2.25) / float32(v25)) - v3) + (v3 + ((float32(v62) + v6) * (v3 / v3))))) - v3))
                                                                                        } else {
                                                                                            if len(v42) > 0 {
                                                                                                v42[0] = 32767
                                                                                            }
                                                                                        }
                                                                                        _ = v62
                                                                                    case "hello":
                                                                                        var v80 uint = v2
                                                                                        v2 -= 42
                                                                                        fmt.Println(v25)
                                                                                        v2 += (^(v80))
                                                                                        v81 := v1[""]
                                                                                        _ = v80
                                                                                        _ = v81
                                                                                    case "\n\t":
                                                                                        if len(v42) > 0 {
                                                                                            v42[0] = (fi19 & fi27)
                                                                                        }
                                                                                    }
                                                                                }
                                                                                fmt.Println(fi27)
                                                                            }
                                                                        } else {
                                                                            var v82 uint16 = v4
                                                                            switch -173 {
                                                                            case -362:
                                                                                if ((^(v2)) > (^(v2))) {
                                                                                    var v83 int16 = -25
                                                                                    v82 -= v82
                                                                                    _ = v83
                                                                                }
                                                                                v3 = ((((3.14159 - (((3.14159 + (1e10 - v6)) * ((v6 * v6) / 3.14159)) - (v6 + 0.1))) + (-2.25 + v3)) * (v3 * (((-2.25 + 0.0) * 3.14159) - (v3 + v3)))) / -2.25)
                                                                                v82 = (v4 % ((v82) | 1))
                                                                            default:
                                                                                var v84 int8 = -121
                                                                                _ = v84
                                                                            }
                                                                            v85 := GenS1{F0: v5, F1: "", F2: 968}
                                                                            _ = v82
                                                                            _ = v85
                                                                        }
                                                                        if false {
                                                                            v86 := GenS8{F0: 1, F1: 255, F2: 9223372036854775807}
                                                                            _ = v86
                                                                        }
                                                                        var v87 uint32 = (^(v5))
                                                                        _ = v87
                                                                    } else {
                                                                        var v88 int16 = 318
                                                                        v89 := GenS3{F0: 4294967295, F1: 127}
                                                                        var v90 int = fi27
                                                                        _ = v88
                                                                        _ = v89
                                                                        _ = v90
                                                                    }
                                                                case 0:
                                                                    if (float32(float64(v3)) != 0.1) {
                                                                        v25 -= v25
                                                                        switch 172 {
                                                                        case 1:
                                                                            v4 -= 382
                                                                        case 836:
                                                                            v5 -= (^(v5))
                                                                        case 0:
                                                                            var v91 string = ("hello" + "")
                                                                            _ = v91
                                                                        default:
                                                                            if (0 > 42) {
                                                                                var v92 float64 = 1.5
                                                                                if len(v42) >= 2 {
                                                                                    v42 = v42[1:len(v42)]
                                                                                }
                                                                                _ = v92
                                                                            } else {
                                                                                v3 -= 0.0
                                                                                var v93 bool = (-683 > 0)
                                                                                v6 = (v3 / v6)
                                                                                _ = v93
                                                                            }
                                                                        }
                                                                        v94 := GenS5{F0: 1000, F1: (v4 << (((v2 / ((180) | 1))) & 63)), F2: 4294967295}
                                                                        _ = v94
                                                                    }
                                                                default:
                                                                    v6 -= ((float32(((v25 - 0.0) * (((((((v25 / v25) / v25) - 0.1) * (v25 * 3.14159)) + v25) / float64(3.14159)) * 0.1))) - (0.1 + float32(((float64(3.14159) / v25) + float64(1.5))))) - (v3 * 3.14159))
                                                                }
                                                                v2 = uint(v5)
                                                            } else {
                                                                v5 = v5
                                                                var v95 string = "foo bar"
                                                                switch (("a" + "a") + (v95 + ("" + v95))) {
                                                                case "hello":
                                                                    var v96 string = "hello"
                                                                    _ = v96
                                                                case "":
                                                                    switch (v95 + ("a" + ("foo bar" + (v95 + v95)))) {
                                                                    case "hello":
                                                                        if false {
                                                                            if (65535 <= 0) {
                                                                                var v97 int32 = 42
                                                                                v98 := GenS8{F0: (v4 / (((^(v4))) | 1)), F1: 255, F2: 369}
                                                                                _ = v97
                                                                                _ = v98
                                                                            } else {
                                                                                var v99 int16 = 1
                                                                                v42 = append(v42, -2147483648)
                                                                                v100 := v1[v95]
                                                                                _ = v100
                                                                                _ = v99
                                                                            }
                                                                        } else {
                                                                            v101 := GenS2{F0: 4294967295, F1: ("\n\t" >= v95), F2: 0}
                                                                            v3 -= v3
                                                                            _ = v101
                                                                        }
                                                                    default:
                                                                        v3 += v3
                                                                        v25 += v25
                                                                    }
                                                                case "\n\t":
                                                                    v102 := v1[(v95 + ("\n\t" + "a"))]
                                                                    if len(v42) > 0 {
                                                                        v42[0] = fi27
                                                                    }
                                                                    v103 := GenS1{F0: 336, F1: ("foo bar" + v95), F2: 42}
                                                                    var v104 uint = v2
                                                                    v25 -= (v25 * v25)
                                                                    _ = v102
                                                                    _ = v103
                                                                    _ = v104
                                                                }
                                                                _ = v95
                                                            }
                                                        }
                                                    }
                                                    if len(v42) >= 2 {
                                                        v42 = v42[1:len(v42)]
                                                    }
                                                    if len(v42) >= 2 {
                                                        v42 = v42[1:len(v42)]
                                                    }
                                                } else {
                                                    var v105 uint8 = 2
                                                    if true {
                                                        fmt.Println(v6)
                                                        if (v2 != (v2 >> ((uint(v4)) & 63))) {
                                                            var v106 float32 = (1.5 / v6)
                                                            var v107 bool = true
                                                            fmt.Println(v107, v6, v2)
                                                            _ = v106
                                                            _ = v107
                                                        } else {
                                                            var v108 float64 = 1e10
                                                            var v109 uint64 = 4294967295
                                                            switch -1 {
                                                            case -9223372036854775808:
                                                                if true {
                                                                    if len(v42) >= 2 {
                                                                        v42 = v42[1:len(v42)]
                                                                    }
                                                                    v3 = ((-2.25 - (3.14159 - float32(1e10))) - (0.1 * 1.5))
                                                                    v3 = v6
                                                                    v4 = (v4 | v4)
                                                                }
                                                            case -904:
                                                                var v110 int = (-(fi27))
                                                                _ = v110
                                                            case 42:
                                                                v111 := GenS5{F0: -784, F1: v4, F2: 18446744073709551615}
                                                                _ = v111
                                                            default:
                                                                fmt.Println(fi27)
                                                            }
                                                            v112 := GenS3{F0: uint64(v4), F1: (v105 ^ (v105 - uint8(v4)))}
                                                            v25 -= v25
                                                            _ = v108
                                                            _ = v109
                                                            _ = v112
                                                        }
                                                        var v113 uint8 = v105
                                                        _ = v113
                                                    } else {
                                                        v114 := GenS1{F0: (v5 | v5), F1: ("hello" + "foo bar"), F2: uint64(v105)}
                                                        var v115 uint32 = (v5 + v5)
                                                        _ = v114
                                                        _ = v115
                                                    }
                                                    _ = v105
                                                }
                                                var v116 float64 = v25
                                                _ = v116
                                            }
                                            var v117 bool = false
                                            var v118 float64 = v25
                                            _ = v117
                                            _ = v118
                                        } else {
                                            delete(v1, ((("hello" + ("\n\t" + "foo bar")) + ("" + "foo bar")) + "\n\t"))
                                        }
                                        fmt.Println(v4, v3)
                                    }
                                    v42 = append(v42, 2147483647)
                                    v119 := GenS9{F0: (fi19 * (fi27 / (((fi19 & fi19)) | 1))), F1: "hello"}
                                    _ = v119
                                    _ = v42
                                    _ = v43
                                } else {
                                    fmt.Println(fi19)
                                }
                            }
                            _ = fi27
                            _ = v28
                        }
                    }
                    _ = v25
                }
                v120 := GenS7{F0: fi19, F1: (v4 / (((^(v4))) | 1))}
                _ = v120
            }
            fmt.Println(v4, v4, fi19)
        }
        _ = fi19
    }
    v5 -= 859
    v121 := []int{}
    if true {
        v4 = 774
        fmt.Println(v5, v4)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v121:", len(v121), v121)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v121
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
