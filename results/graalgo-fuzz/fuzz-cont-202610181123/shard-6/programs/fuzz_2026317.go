package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int; F1 int16 }
type GenS4 struct { F0 int32; F1 int16; F2 int }
type GenS5 struct { F0 uint8; F1 string; F2 int64 }
type GenS6 struct { F0 int }
type GenS7 struct { F0 int; F1 string; F2 uint }
type GenS8 struct { F0 int16; F1 string }
func genFn0(p0 int8, p1 int8, p2 uint) string {
    p2 -= (p2 ^ (p2 * 458))
    var v7 string = "foo bar"
    v8 := []int{-976, -391, 1}
    v9 := GenS1{F0: 100, F1: 885}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    _ = v9
    return (("" + v7) + "hello")
}
func genFn1(p0 int32) uint {
    for fi10 := 0; fi10 < 10; fi10++ {
        for fi11 := 0; fi11 < 6; fi11++ {
            if (100 <= 64) {
                fmt.Println(fi10, p0, fi11)
            } else {
                p0 = p0
                if true {
                    if (false || true) {
                        p0 -= p0
                    } else {
                        var v12 uint64 = 18446744073709551615
                        var v13 int = fi11
                        _ = v12
                        _ = v13
                    }
                    var v14 float32 = float32(1.5)
                    _ = v14
                } else {
                    v15 := []int{(fi11 & -563)}
                    _ = v15
                }
                if true {
                    if false {
                        p0 += (^(p0))
                    } else {
                        v16 := GenS2{F0: 97}
                        p0 = 1
                        _ = v16
                    }
                    v17 := GenS3{F0: (fi11 << ((2) & 63)), F1: -32768}
                    fmt.Println(fi10)
                    _ = v17
                } else {
                    var v18 int64 = 1
                    v18 = v18
                    var v19 uint = 65535
                    _ = v18
                    _ = v19
                }
                var v20 string = "hello"
                _ = v20
            }
            fmt.Println(p0, fi10)
            v21 := GenS2{F0: 4294967295}
            _ = fi11
            _ = v21
        }
        v22 := []int{}
        p0 += -861
        _ = fi10
        _ = v22
    }
    if false {
        for fi23 := 0; fi23 < 2; fi23++ {
            var v24 float64 = (0.1 - -2.25)
            _ = fi23
            _ = v24
        }
    } else {
        if (!(32768 <= 1)) {
            v25 := []int{-727, 916}
            v26 := GenS3{F0: int(p0), F1: 1}
            for fi27 := 0; fi27 < 7; fi27++ {
                for fi28 := 0; fi28 < 11; fi28++ {
                    if (128 >= 0) {
                        var v29 int8 = 46
                        switch "foo bar" {
                        case "":
                            if len(v25) >= 2 {
                                v25 = v25[1:len(v25)]
                            }
                        case "foo bar":
                            p0 -= p0
                        case "a":
                            switch fi28 {
                            case -83:
                                p0 = p0
                            case 730:
                                v30 := GenS2{F0: 929}
                                _ = v30
                            case -9223372036854775808:
                                switch (("hello" + ("foo bar" + "hello")) + "foo bar") {
                                case "":
                                    var v31 int32 = (p0 - p0)
                                    _ = v31
                                case "hello":
                                    _ = len(v25)
                                case "a":
                                    var v32 uint8 = 0
                                    if len(v25) > 0 {
                                        v25[0] = -9223372036854775808
                                    }
                                    _ = v32
                                }
                            }
                        default:
                            if ("" >= "hello") {
                                if false {
                                    if false {
                                        var v33 bool = false
                                        _ = v33
                                    } else {
                                        switch (fi27 * 4294967295) {
                                        case 0:
                                            fmt.Println(p0, fi28, fi28)
                                        case 42:
                                            _ = len(v25)
                                        }
                                        v29 = 127
                                    }
                                    var v34 int8 = v29
                                    if true {
                                        v29 -= (v34 * 127)
                                        v25 = append(v25, (fi27 % ((1) | 1)))
                                        if true {
                                            _ = len(v25)
                                            var v35 uint8 = 2
                                            v36 := GenS4{F0: 724, F1: 1000, F2: fi28}
                                            _ = v35
                                            _ = v36
                                        } else {
                                            v37 := GenS3{F0: (fi28 | (fi28 * fi27)), F1: 32767}
                                            v29 -= (v29 / (((v29 - (v34 | -1))) | 1))
                                            _ = v37
                                        }
                                    } else {
                                        if ("\n\t" == ("foo bar" + ((("" + "foo bar") + ((("foo bar" + (("" + ("a" + "a")) + "a")) + ("" + "")) + "foo bar")) + "a"))) {
                                            if false {
                                                if true {
                                                    v38 := GenS1{F0: -1, F1: 1000}
                                                    v39 := GenS2{F0: 2}
                                                    _ = v38
                                                    _ = v39
                                                }
                                            }
                                        }
                                    }
                                    var v40 float32 = (0.1 - 0.1)
                                    _ = v34
                                    _ = v40
                                } else {
                                    var v41 float32 = -2.25
                                    v42 := GenS5{F0: 88, F1: "", F2: 9223372036854775807}
                                    _ = v41
                                    _ = v42
                                }
                                var v43 uint16 = 65535
                                v44 := GenS6{F0: 9223372036854775807}
                                _ = v43
                                _ = v44
                            } else {
                                v29 = -4
                            }
                        }
                        _ = v29
                    }
                    if (42 <= 32767) {
                        var v45 float32 = 1e10
                        p0 -= (p0 / (((p0 * p0)) | 1))
                        switch ("foo bar" + "foo bar") {
                        case "hello":
                            var v46 float32 = v45
                            _ = v46
                        default:
                            if (128 < 2) {
                                var v47 bool = (!true)
                                _ = v47
                            }
                        }
                        _ = v45
                    }
                    _ = fi28
                }
                v48 := GenS6{F0: (fi27 % ((-348) | 1))}
                _ = fi27
                _ = v48
            }
            _ = v25
            _ = v26
        } else {
            p0 = p0
            var v49 uint8 = 201
            _ = v49
        }
        p0 = (p0 / ((p0) | 1))
        v50 := []int{4294967295, 42, 497}
        _ = v50
    }
    if false {
        if (float64(1.5) <= 1.5) {
            switch ("" + "a") {
            case "a":
                if ((2 > 0) || true) {
                    var v51 int32 = (-(p0))
                    _ = v51
                } else {
                    switch ("" + "hello") {
                    case "\n\t":
                        p0 = (^(p0))
                    default:
                        p0 = -272
                    }
                }
            case "":
                var v52 uint32 = 42
                _ = v52
            case "foo bar":
                if true {
                    if (false && ((p0 ^ (^(p0))) >= (-(p0)))) {
                        var v53 bool = false
                        p0 = (p0 * -2147483648)
                        v54 := []int{127, 830}
                        _ = v53
                        _ = v54
                    }
                    p0 = p0
                    v55 := []int{-9223372036854775808, -2147483648, 127, 32768}
                    _ = v55
                }
            default:
                var v56 string = "foo bar"
                _ = v56
            }
            var v57 uint16 = 67
            _ = v57
        }
        if (42 > 255) {
            fmt.Println(p0)
        } else {
            p0 += p0
            for fi58 := 0; fi58 < 3; fi58++ {
                switch ("a" + "foo bar") {
                case "":
                    v59 := []int{(^(fi58)), 42, 294, (fi58 ^ (fi58 * (fi58 & fi58)))}
                    _ = v59
                }
                v60 := []int{(fi58 % (((fi58 & 42)) | 1))}
                var v61 string = ("a" + "hello")
                _ = fi58
                _ = v60
                _ = v61
            }
            p0 -= (p0 - 842)
        }
    } else {
        if ((!(0 > 42)) && true) {
            for fi62 := 0; fi62 < 2; fi62++ {
                var v63 bool = (18446744073709551615 == 42)
                if (127 < 138) {
                    for fi64 := 0; fi64 < 20; fi64++ {
                        v65 := GenS5{F0: 2, F1: "a", F2: 42}
                        var v66 bool = (24 < 454)
                        fmt.Println(v66, v66, fi62)
                        _ = fi64
                        _ = v65
                        _ = v66
                    }
                    fmt.Println(v63, fi62)
                }
                _ = fi62
                _ = v63
            }
        } else {
            p0 -= 110
        }
        switch -1 {
        case -1:
            for fi67 := 0; fi67 < 2; fi67++ {
                var v68 int64 = 0
                _ = fi67
                _ = v68
            }
        case -2147483648:
            var v69 int16 = 1
            _ = v69
        case 4294967295:
            for fi70 := 0; fi70 < 20; fi70++ {
                fmt.Println(p0)
                v71 := []int{(fi70 / ((32768) | 1))}
                _ = fi70
                _ = v71
            }
        default:
            for fi72 := 0; fi72 < 16; fi72++ {
                v73 := []int{fi72, fi72}
                fmt.Println(p0, p0)
                _ = fi72
                _ = v73
            }
        }
    }
    p0 = (p0 + 42)
    _ = p0
    return 2
}

func main() {
    v1 := []int{2147483647, 864, -19}
    var v2 int = 319
    var v3 int8 = 127
    v4 := []string{"a", "a", "foo bar"}
    v5 := make(map[string]int)
    var v6 int8 = v3
    fmt.Println(v3)
    var v74 uint8 = 255
    fmt.Println(v3, v74, v2)
    fmt.Println(v2, v3)
    if (-1 < v2) {
        for fi75 := 0; fi75 < 6; fi75++ {
            if true {
                v1 = append(v1, 9223372036854775807)
                v6 = 127
                var v76 uint = 311
                _ = v76
            } else {
                switch "foo bar" {
                case "hello":
                    v2 = (fi75 ^ (v2 & (fi75 / (((-(v2))) | 1))))
                default:
                    for fi77 := 0; fi77 < 4; fi77++ {
                        var v78 int = (-(v2))
                        fmt.Println(fi77, fi75)
                        v79 := v5["a"]
                        _ = fi77
                        _ = v78
                        _ = v79
                    }
                }
                if (!(-907 <= -1)) {
                    switch (-(fi75)) {
                    case 2147483647:
                        if (32767 < -32768) {
                            var v80 float64 = ((-2.25 * float64(0.0)) / 1.0)
                            for fi81 := 0; fi81 < 16; fi81++ {
                                v3 -= v3
                                _ = fi81
                            }
                            _ = v80
                        }
                    case -197:
                        v2 = fi75
                    }
                    v82 := GenS4{F0: 1, F1: -510, F2: (v2 * (^(v2)))}
                    _ = v82
                } else {
                    _ = len(v4)
                    fmt.Println(fi75)
                }
            }
            fmt.Println(v6, v3)
            _ = fi75
        }
        fmt.Println(v74)
        if true {
            var v83 int8 = (^(v6))
            _ = v83
        } else {
            v3 -= v3
            if (1 >= -9223372036854775808) {
                switch v2 {
                case 768:
                    _ = len(v4)
                }
                switch v2 {
                case 410:
                    for fi84 := 0; fi84 < 10; fi84++ {
                        var v85 uint16 = 255
                        var v86 uint8 = v74
                        var v87 float32 = (1.5 - (-2.25 - 1e10))
                        _ = fi84
                        _ = v85
                        _ = v86
                        _ = v87
                    }
                case 32768:
                    if true {
                        v74 -= (^(v74))
                        v6 += (v3 * v6)
                    }
                case 128:
                    switch (v2 + 150) {
                    case 385:
                        v4 = append(v4, ("hello" + ("foo bar" + "hello")))
                    case -9223372036854775808:
                        switch "hello" {
                        case "foo bar":
                            if false {
                                if len(v4) > 0 {
                                    v4[0] = "a"
                                }
                                switch v2 {
                                case 1:
                                    switch 127 {
                                    case -1:
                                        v6 -= v6
                                    default:
                                        fmt.Println(v3, v6, v3)
                                    }
                                case 128:
                                    var v88 int64 = -9223372036854775808
                                    _ = v88
                                default:
                                    v6 = v6
                                }
                                fmt.Println(v3)
                            } else {
                                v89 := GenS1{F0: v3, F1: -1}
                                var v90 int16 = 42
                                _ = v89
                                _ = v90
                            }
                        default:
                            v91 := GenS7{F0: 2147483647, F1: ("hello" + "hello"), F2: 671}
                            _ = v91
                        }
                    case 9223372036854775807:
                        switch ("" + "") {
                        case "a":
                            v1 = append(v1, 4294967295)
                        }
                    }
                }
            } else {
                v92 := GenS3{F0: v2, F1: -411}
                _ = v92
            }
            if false {
                switch (v2 & 1) {
                case 128:
                    switch 1 {
                    case -1:
                        var v93 int64 = 42
                        _ = v93
                    }
                }
            } else {
                var v94 int64 = 42
                _ = v94
            }
        }
    }
    if false {
        _ = len(v5)
    }
    v95 := GenS8{F0: 42, F1: ("a" + ((("" + ("\n\t" + "foo bar")) + "\n\t") + ""))}
    v3 = v6
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v74:", v74)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v74
    _ = v95
}
