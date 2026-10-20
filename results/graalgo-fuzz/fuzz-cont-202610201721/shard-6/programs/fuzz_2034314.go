package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 string; F1 uint64; F2 uint64 }
type GenS5 struct { F0 uint16; F1 uint8 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 string; F1 uint64 }
type GenS8 struct { F0 bool; F1 uint64 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 int16 }
type GenS11 struct { F0 int16 }
func genFn0(p0 uint8, p1 int8, p2 int32) uint8 {
    switch "" {
    case "foo bar":
        switch ("foo bar" + "\n\t") {
        case "\n\t":
            v5 := GenS1{F0: 18446744073709551615}
            _ = v5
        case "hello":
            fmt.Println(p1, p0, p2)
        default:
            var v6 uint16 = 32767
            _ = v6
        }
    case "":
        switch "foo bar" {
        case "\n\t":
            if ((423 != 1) && (838 >= 2)) {
                v7 := GenS1{F0: 42}
                v8 := []int{}
                _ = v7
                _ = v8
            }
        case "hello":
            v9 := GenS1{F0: 2}
            _ = v9
        }
    default:
        var v10 uint8 = p0
        _ = v10
    }
    for fi11 := 0; fi11 < 3; fi11++ {
        for fi12 := 0; fi12 < 20; fi12++ {
            var v13 uint8 = 149
            _ = fi12
            _ = v13
        }
        v14 := GenS1{F0: 255}
        var v15 float64 = 3.14159
        _ = fi11
        _ = v14
        _ = v15
    }
    v16 := GenS1{F0: 477}
    _ = p0
    _ = p1
    _ = p2
    _ = v16
    return p0
}
func genFn1(p0 uint64, p1 uint32, p2 uint32) int32 {
    var v17 int64 = 0
    if true {
        if false {
            if false {
                for fi18 := 0; fi18 < 6; fi18++ {
                    if (true && true) {
                        switch 128 {
                        case 4294967295:
                            var v19 int64 = v17
                            _ = v19
                        case 128:
                            if true {
                                fmt.Println(p0, p1, p2)
                                v20 := []int{32767, fi18, (fi18 * fi18), (fi18 / ((-361) | 1))}
                                _ = v20
                            } else {
                                v17 += v17
                            }
                        case 127:
                            switch "a" {
                            case "\n\t":
                                var v21 uint8 = 239
                                _ = v21
                            default:
                                v22 := []int{fi18, fi18}
                                _ = v22
                            }
                        default:
                            var v23 string = ("" + ("" + "\n\t"))
                            _ = v23
                        }
                    } else {
                        p2 += p2
                        if true {
                            switch fi18 {
                            case -132:
                                var v24 int = (fi18 + fi18)
                                _ = v24
                            case -2147483648:
                                p2 -= uint32(p0)
                            default:
                                v17 = v17
                            }
                        } else {
                            var v25 int8 = -37
                            _ = v25
                        }
                    }
                    switch fi18 {
                    case 32767:
                        var v26 int64 = v17
                        _ = v26
                    case -1:
                        fmt.Println(v17, fi18)
                    }
                    _ = fi18
                }
            } else {
                if (("\n\t" + ("\n\t" + "")) == "foo bar") {
                    switch "" {
                    case "foo bar":
                        v27 := []int{4294967295}
                        _ = v27
                    case "hello":
                        var v28 int8 = -128
                        _ = v28
                    case "a":
                        for fi29 := 0; fi29 < 1; fi29++ {
                            var v30 int16 = 0
                            _ = fi29
                            _ = v30
                        }
                    default:
                        for fi31 := 0; fi31 < 3; fi31++ {
                            var v32 uint = 255
                            var v33 string = ((("\n\t" + "hello") + (((((("" + ("\n\t" + ("foo bar" + "\n\t"))) + "a") + ("a" + "hello")) + "foo bar") + "\n\t") + (("" + ((((("foo bar" + "hello") + "\n\t") + "") + "foo bar") + ("foo bar" + "\n\t"))) + "\n\t"))) + "hello")
                            _ = fi31
                            _ = v32
                            _ = v33
                        }
                    }
                } else {
                    for fi34 := 0; fi34 < 15; fi34++ {
                        var v35 int64 = -300
                        _ = fi34
                        _ = v35
                    }
                    v36 := GenS2{F0: false}
                    v37 := []int{1, -9223372036854775808, -9223372036854775808}
                    _ = v36
                    _ = v37
                }
                v38 := []int{568}
                _ = v38
            }
        }
        if (623 <= 255) {
            var v39 int64 = 9223372036854775807
            _ = v39
        }
        for fi40 := 0; fi40 < 12; fi40++ {
            p2 += (^(p1))
            p0 -= (p0 * 677)
            v41 := GenS1{F0: 65535}
            _ = fi40
            _ = v41
        }
    }
    var v42 float64 = (1e10 * (3.14159 - (3.14159 - 0.0)))
    for fi43 := 0; fi43 < 4; fi43++ {
        v44 := []int{(-(fi43)), 9223372036854775807}
        var v45 uint = 18446744073709551615
        p1 -= (p2 << ((v45) & 63))
        _ = fi43
        _ = v44
        _ = v45
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v17
    _ = v42
    return 0
}

func main() {
    var v1 uint16 = 933
    var v2 int32 = -2147483648
    var v3 uint64 = 862
    v4 := []string{"hello", "a"}
    if false {
        v1 = v1
    }
    for fi46 := 0; fi46 < 13; fi46++ {
        v47 := GenS3{F0: 0}
        _ = fi46
        _ = v47
    }
    switch ("a" + "a") {
    case "":
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
    case "\n\t":
        v3 += v3
    case "foo bar":
        for fi48 := 0; fi48 < 20; fi48++ {
            for fi49 := 0; fi49 < 11; fi49++ {
                if (!(("hello" + ("hello" + ("\n\t" + ""))) >= "hello")) {
                    switch ("\n\t" + "") {
                    case "\n\t":
                        switch (fi49 / ((741) | 1)) {
                        case 128:
                            if false {
                                if (2 != 181) {
                                    v50 := GenS2{F0: true}
                                    v1 = v1
                                    v1 -= (v1 >> ((18446744073709551615) & 63))
                                    v51 := GenS4{F0: "hello", F1: (^(v3)), F2: v3}
                                    v52 := GenS2{F0: (0.1 != (float32(0.0) - ((1e10 + 1.5) * float32(((float64((0.0 - (float32(0.1) - (3.14159 + 1.5)))) - (-2.25 / 0.1)) / (0.1 * 3.14159))))))}
                                    fmt.Println(v2)
                                    v53 := GenS5{F0: (v1 * (v1 ^ (v1 + (^(v1))))), F1: genFn0(42, -70, (v2 | (-(v2))))}
                                    _ = v50
                                    _ = v51
                                    _ = v52
                                    _ = v53
                                } else {
                                    switch fi49 {
                                    case -797:
                                        v54 := GenS3{F0: -867}
                                        _ = v54
                                    default:
                                        switch fi49 {
                                        case -1:
                                            var v55 uint64 = (v3 ^ 994)
                                            _ = v55
                                        case 1:
                                            v56 := GenS1{F0: 255}
                                            _ = v56
                                        default:
                                            fmt.Println(fi48, v2, fi48)
                                        }
                                        v4 = append(v4, "\n\t")
                                        v57 := GenS3{F0: 42}
                                        if len(v4) > 0 {
                                            v4[0] = ((((("a" + "foo bar") + "foo bar") + ((("hello" + "\n\t") + "foo bar") + "hello")) + "") + (("foo bar" + "hello") + ("" + "foo bar")))
                                        }
                                        _ = v57
                                    }
                                    if (!false) {
                                        if (!false) {
                                            v2 = v2
                                            if len(v4) > 0 {
                                                v4[0] = "hello"
                                            }
                                            v1 = (v1 + 13)
                                        }
                                        if (v3 == (v3 ^ 2)) {
                                            v2 = genFn1(42, 42, 4294967295)
                                            if true {
                                                var v58 uint = 1
                                                fmt.Println(fi49, v58, v3)
                                                _ = v58
                                            } else {
                                                var v59 uint32 = 65535
                                                if (!true) {
                                                    v3 = v3
                                                    var v60 int16 = 1000
                                                    _ = v60
                                                } else {
                                                    v61 := GenS6{F0: true}
                                                    fmt.Println(v3, fi49, fi49)
                                                    _ = v61
                                                }
                                                _ = v59
                                            }
                                            switch "" {
                                            case "foo bar":
                                                var v62 uint = 1
                                                _ = v62
                                            case "a":
                                                var v63 int = (fi48 | -2147483648)
                                                _ = v63
                                            default:
                                                _ = len(v4)
                                            }
                                            if true {
                                                if false {
                                                    switch "a" {
                                                    case "a":
                                                        v2 += (v2 - (v2 % ((v2) | 1)))
                                                    case "hello":
                                                        v2 = 0
                                                    }
                                                } else {
                                                    v64 := GenS7{F0: ("" + "\n\t"), F1: 18446744073709551615}
                                                    _ = v64
                                                }
                                            }
                                            var v65 uint8 = 0
                                            v1 = (^(v1))
                                            _ = v65
                                        }
                                        v1 -= 32767
                                    } else {
                                        v1 += (v1 / (((v1 / ((783) | 1))) | 1))
                                    }
                                    v66 := GenS8{F0: (-128 != 100), F1: (^(v3))}
                                    var v67 uint8 = 155
                                    if len(v4) > 0 {
                                        v4[0] = (("foo bar" + ("hello" + "hello")) + ("a" + ("hello" + "hello")))
                                    }
                                    fmt.Println(v3)
                                    v2 = genFn1((v3 << ((932) & 63)), 1, 4294967295)
                                    if (42 > 42) {
                                        v67 = v67
                                        if ((0.1 - 0.1) == 0.0) {
                                            var v68 int32 = 42
                                            var v69 uint = 0
                                            v4 = append(v4, "hello")
                                            _ = v68
                                            _ = v69
                                        }
                                        var v70 int16 = 32767
                                        v71 := GenS1{F0: 103}
                                        _ = v70
                                        _ = v71
                                    }
                                    switch ("" + ("" + "hello")) {
                                    case "foo bar":
                                        switch fi49 {
                                        case 32768:
                                            if true {
                                                _ = len(v4)
                                            }
                                        default:
                                            if false {
                                                switch 32768 {
                                                case 127:
                                                    var v72 uint32 = 545
                                                    _ = v72
                                                case -2147483648:
                                                    var v73 int64 = -229
                                                    _ = v73
                                                case 9223372036854775807:
                                                    if (2147483648 > 2147483648) {
                                                        v4 = append(v4, (("a" + "") + "foo bar"))
                                                    }
                                                default:
                                                    if len(v4) > 0 {
                                                        v4[0] = "hello"
                                                    }
                                                }
                                            } else {
                                                if len(v4) > 0 {
                                                    v4[0] = "\n\t"
                                                }
                                                v74 := GenS9{F0: (v1 ^ v1)}
                                                v1 -= v1
                                                _ = v74
                                            }
                                        }
                                    case "hello":
                                        v75 := GenS10{F0: 1000}
                                        _ = v75
                                    default:
                                        v76 := GenS4{F0: "a", F1: 377, F2: (^(v3))}
                                        _ = v76
                                    }
                                    _ = v66
                                    _ = v67
                                }
                            }
                        case -241:
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        default:
                            if true {
                                v2 -= genFn1((v3 + (v3 % (((v3 - 4294967295)) | 1))), 754, 2147483647)
                                v1 -= (v1 | v1)
                            } else {
                                v77 := GenS8{F0: (((("hello" + "hello") + "") + ("\n\t" + ("" + ("a" + ("hello" + ("foo bar" + "\n\t")))))) < "\n\t"), F1: (^(v3))}
                                _ = v77
                            }
                        }
                    default:
                        switch "\n\t" {
                        case "":
                            fmt.Println(v3, fi49)
                        }
                    }
                } else {
                    if false {
                        v2 += -182
                        switch "a" {
                        case "foo bar":
                            if false {
                                v3 = (v3 * (v3 >> ((304) & 63)))
                                v2 += 42
                                var v78 uint8 = 0
                                v79 := GenS11{F0: -573}
                                _ = v78
                                _ = v79
                            }
                        case "a":
                            switch (fi49 ^ 713) {
                            case 404:
                                var v80 float64 = float64(3.14159)
                                _ = v80
                            case 2147483647:
                                fmt.Println(fi49)
                            default:
                                fmt.Println(fi49, fi48, fi49)
                            }
                        }
                        fmt.Println(v2, v2, v1)
                    }
                    if ((fi49 - fi49) >= (fi49 ^ (fi48 * (fi49 ^ fi48)))) {
                        switch ("foo bar" + (("\n\t" + "hello") + (("" + ("hello" + "a")) + ("hello" + ("a" + ("\n\t" + ("foo bar" + "hello"))))))) {
                        case "a":
                            var v81 uint64 = 0
                            _ = v81
                        case "hello":
                            v2 += v2
                        default:
                            _ = len(v4)
                        }
                        var v82 bool = false
                        _ = v82
                    }
                }
                _ = fi49
            }
            _ = fi48
        }
    }
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v83 int32 = (v2 >> ((18446744073709551615) & 63))
    for fi84 := 0; fi84 < 5; fi84++ {
        for fi85 := 0; fi85 < 10; fi85++ {
            var v86 int32 = v2
            var v87 bool = true
            switch ("foo bar" + "\n\t") {
            case "\n\t":
                if (((^(v3)) <= v3) || (-128 != -1)) {
                    if len(v4) > 0 {
                        v4[0] = "\n\t"
                    }
                } else {
                    var v88 uint8 = 14
                    var v89 uint = 0
                    var v90 uint32 = 268
                    _ = v88
                    _ = v89
                    _ = v90
                }
            default:
                fmt.Println(v87)
                if v87 {
                    v4 = append(v4, (("" + "") + "foo bar"))
                    var v91 int64 = -854
                    _ = v91
                } else {
                    v4 = append(v4, "hello")
                    switch "foo bar" {
                    case "foo bar":
                        var v92 uint = 2
                        _ = v92
                    case "\n\t":
                        if v87 {
                            v1 = uint16(v3)
                            v2 += genFn1(v3, 2, 2147483647)
                        }
                    case "a":
                        var v93 uint8 = 243
                        _ = v93
                    }
                    var v94 bool = false
                    _ = v94
                }
                fmt.Println(v1, v1)
                fmt.Println(v3)
            }
            var v95 int16 = -1
            _ = fi85
            _ = v86
            _ = v87
            _ = v95
        }
        _ = fi84
    }
    switch "\n\t" {
    case "a":
        for fi96 := 0; fi96 < 16; fi96++ {
            var v97 uint8 = 143
            _ = fi96
            _ = v97
        }
    case "":
        fmt.Println(v1, v3)
    default:
        fmt.Println(v1, v3, v1)
    }
    var v98 uint8 = 136
    v83 = -828
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v83:", v83)
    fmt.Println("v98:", v98)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v83
    _ = v98
}
