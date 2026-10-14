package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool; F1 uint64 }
type GenS3 struct { F0 uint16; F1 int32; F2 uint8 }
type GenS4 struct { F0 uint32 }
type GenS5 struct { F0 bool }
type GenS6 struct { F0 int }
type GenS7 struct { F0 string }
type GenS8 struct { F0 int8 }
func genFn0(p0 string) uint32 {
    switch ("" + p0) {
    case "":
        var v5 uint64 = 994
        _ = v5
    case "a":
        p0 = p0
    case "foo bar":
        v6 := GenS1{F0: 42}
        _ = v6
    default:
        p0 = ("a" + p0)
    }
    for fi7 := 0; fi7 < 10; fi7++ {
        v8 := GenS1{F0: (fi7 & (-(fi7)))}
        _ = fi7
        _ = v8
    }
    _ = p0
    return 732
}
func genFn1(p0 uint16) uint32 {
    v9 := GenS1{F0: 128}
    if true {
        v10 := []int{-2147483648, -663}
        _ = v10
    }
    _ = p0
    _ = v9
    return genFn0("")
}

func main() {
    var v1 float32 = float32(1e10)
    v2 := []string{"a", "hello", "hello", "hello"}
    var v3 string = "hello"
    var v4 int = 128
    var v11 uint64 = 42
    var v12 int8 = -128
    v13 := GenS1{F0: v4}
    if ((v12 + (v12 >> ((2) & 63))) < (v12 >> ((255) & 63))) {
        fmt.Println(v1)
        v11 = 18446744073709551615
    } else {
        fmt.Println(v11, v1, v3)
        for fi14 := 0; fi14 < 2; fi14++ {
            var v15 uint8 = 85
            _ = fi14
            _ = v15
        }
        v16 := GenS1{F0: 127}
        _ = v16
    }
    v17 := GenS1{F0: 32767}
    if (!false) {
        for fi18 := 0; fi18 < 19; fi18++ {
            v2 = append(v2, "")
            _ = fi18
        }
        var v19 float64 = float64(-2.25)
        _ = v19
    } else {
        v20 := GenS1{F0: -1}
        if (42 < 32767) {
            if (v4 > v4) {
                v3 = v3
                v12 = -128
            }
        } else {
            v11 = 0
        }
        v12 += (-(v12))
        _ = v20
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    switch v4 {
    case 9223372036854775807:
        for fi21 := 0; fi21 < 1; fi21++ {
            var v22 int64 = 472
            switch "" {
            case "hello":
                if (!(65535 >= 767)) {
                    v23 := GenS1{F0: v4}
                    fmt.Println(v4)
                    _ = v23
                } else {
                    v22 = (v22 >> ((4294967295) & 63))
                    fmt.Println(v22, v3, v1)
                }
            }
            _ = fi21
            _ = v22
        }
    case 128:
        fmt.Println(v4, v11, v1)
    case -499:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    default:
        v24 := GenS1{F0: v4}
        _ = v24
    }
    v3 = v3
    if (!("" != v3)) {
        for fi25 := 0; fi25 < 7; fi25++ {
            if true {
                var v26 uint = 65535
                var v27 int = fi25
                var v28 int32 = 0
                _ = v26
                _ = v27
                _ = v28
            }
            v3 = (v3 + v3)
            _ = fi25
        }
    }
    switch -849 {
    case 183:
        for fi29 := 0; fi29 < 18; fi29++ {
            switch int(v11) {
            case -107:
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
            case 229:
                v2 = append(v2, "")
            case 32767:
                if false {
                    for fi30 := 0; fi30 < 5; fi30++ {
                        if true {
                            if (0 > 42) {
                                if false {
                                    var v31 int = (fi30 << ((65535) & 63))
                                    _ = v31
                                }
                                var v32 bool = true
                                _ = v32
                            }
                            v1 = v1
                        }
                        fmt.Println(v1, v3, v11)
                        v1 += (v1 + 0.1)
                        var v33 float64 = 3.14159
                        _ = fi30
                        _ = v33
                    }
                } else {
                    if false {
                        var v34 uint16 = 255
                        _ = v34
                    }
                    v2 = append(v2, "foo bar")
                    v4 -= (fi29 / ((v4) | 1))
                }
            default:
                _ = len(v2)
            }
            for fi35 := 0; fi35 < 19; fi35++ {
                v1 -= float32(-2.25)
                v11 = 42
                _ = fi35
            }
            _ = fi29
        }
    case 2147483647:
        if (1 <= 4294967295) {
            switch (v3 + v3) {
            case "":
                for fi36 := 0; fi36 < 14; fi36++ {
                    var v37 uint8 = 25
                    if (true && false) {
                        _ = len(v2)
                    } else {
                        v11 = v11
                        var v38 uint32 = 101
                        _ = v38
                    }
                    v11 = v11
                    _ = fi36
                    _ = v37
                }
            default:
                if true {
                    v4 -= (^(v4))
                    if true {
                        switch (v3 + (v3 + "a")) {
                        case "":
                            for fi39 := 0; fi39 < 7; fi39++ {
                                v4 = (fi39 & (-(fi39)))
                                var v40 uint8 = 42
                                if (!(0 > 0)) {
                                    var v41 float64 = ((((-2.25 + float64(0.1)) / 1.5) / 1.0) - 1e10)
                                    v42 := GenS1{F0: 128}
                                    _ = v41
                                    _ = v42
                                }
                                _ = fi39
                                _ = v40
                            }
                        case "foo bar":
                            v43 := GenS1{F0: (v4 * (v4 / (((v4 % ((v4) | 1))) | 1)))}
                            _ = v43
                        default:
                            v3 = ("" + "hello")
                        }
                    } else {
                        fmt.Println(v1, v3)
                        if (false && ((65535 > 2) && (2 != 163))) {
                            var v44 int32 = -1
                            switch (v4 * (v4 - -1)) {
                            case -667:
                                switch (("\n\t" + v3) + ((v3 + (v3 + "")) + (v3 + v3))) {
                                case "":
                                    if len(v2) > 0 {
                                        v2[0] = v3
                                    }
                                case "foo bar":
                                    v1 -= (v1 + v1)
                                default:
                                    v12 -= 1
                                }
                            case 1:
                                v3 = "foo bar"
                            }
                            _ = v44
                        } else {
                            v11 -= 0
                            if true {
                                for fi45 := 0; fi45 < 18; fi45++ {
                                    var v46 bool = true
                                    v47 := GenS1{F0: 127}
                                    fmt.Println(v3)
                                    _ = fi45
                                    _ = v46
                                    _ = v47
                                }
                                v4 -= (v4 + -72)
                            }
                            v11 += 18446744073709551615
                        }
                    }
                    _ = len(v2)
                } else {
                    for fi48 := 0; fi48 < 17; fi48++ {
                        switch (v3 + v3) {
                        case "":
                            fmt.Println(v4, v12)
                        case "hello":
                            for fi49 := 0; fi49 < 14; fi49++ {
                                v3 = (v3 + v3)
                                switch v3 {
                                case "":
                                    var v50 uint32 = 2147483647
                                    _ = v50
                                case "\n\t":
                                    var v51 int32 = 64
                                    _ = v51
                                }
                                _ = fi49
                            }
                        default:
                            for fi52 := 0; fi52 < 18; fi52++ {
                                if (42 >= 2) {
                                    if (false || (4294967295 > 4294967295)) {
                                        fmt.Println(fi52, fi48)
                                    } else {
                                        switch (v3 + (((v3 + v3) + (v3 + (((v3 + "\n\t") + v3) + ""))) + "\n\t")) {
                                        case "foo bar":
                                            v12 = (v12 - 127)
                                            switch ((v3 + "") + v3) {
                                            case "a":
                                                if false {
                                                    v3 = "foo bar"
                                                    v2 = append(v2, ((v3 + "foo bar") + v3))
                                                    if false {
                                                        v2 = append(v2, ((v3 + (v3 + v3)) + "foo bar"))
                                                        var v53 string = v3
                                                        _ = v53
                                                    }
                                                    var v54 float32 = v1
                                                    switch "foo bar" {
                                                    case "":
                                                        var v55 uint32 = 0
                                                        _ = v55
                                                    case "a":
                                                        v12 -= 1
                                                    }
                                                    var v56 float64 = 0.0
                                                    v56 = v56
                                                    _ = v54
                                                    _ = v56
                                                } else {
                                                    fmt.Println(v3, v11, fi52)
                                                    var v57 int32 = 1
                                                    _ = v57
                                                }
                                            case "foo bar":
                                                var v58 int8 = (v12 | (v12 + (v12 * (v12 >> ((795) & 63)))))
                                                _ = v58
                                            }
                                        case "":
                                            fmt.Println(fi48, v4, fi52)
                                        }
                                    }
                                } else {
                                    fmt.Println(v1, v12, v12)
                                    v59 := GenS2{F0: (true && false), F1: 4294967295}
                                    var v60 int32 = 42
                                    _ = v59
                                    _ = v60
                                }
                                switch fi52 {
                                case -9223372036854775808:
                                    var v61 int16 = 930
                                    _ = v61
                                case 0:
                                    var v62 uint16 = 2
                                    _ = v62
                                case -3:
                                    v12 = 100
                                }
                                v63 := GenS3{F0: 2, F1: 2147483647, F2: 128}
                                _ = fi52
                                _ = v63
                            }
                        }
                        _ = fi48
                    }
                    if true {
                        v4 -= -223
                    } else {
                        switch v3 {
                        case "\n\t":
                            switch (v3 + v3) {
                            case "\n\t":
                                v11 += v11
                            case "a":
                                if (0 > uint16(v11)) {
                                    var v64 bool = (127 < 189)
                                    _ = v64
                                } else {
                                    fmt.Println(v11, v4, v12)
                                    for fi65 := 0; fi65 < 3; fi65++ {
                                        if false {
                                            v66 := GenS3{F0: 109, F1: 2147483647, F2: 42}
                                            _ = v66
                                        }
                                        v11 = (v11 >> ((577) & 63))
                                        _ = fi65
                                    }
                                    var v67 int16 = 412
                                    _ = v67
                                }
                            }
                        default:
                            v68 := GenS4{F0: 224}
                            _ = v68
                        }
                        fmt.Println(v3, v11, v3)
                        if len(v2) > 0 {
                            v2[0] = ("hello" + (v3 + (v3 + (v3 + "foo bar"))))
                        }
                    }
                }
            }
        }
    }
    switch "foo bar" {
    case "\n\t":
        for fi69 := 0; fi69 < 11; fi69++ {
            switch (fi69 % (((v4 * v4)) | 1)) {
            case 32768:
                switch fi69 {
                case -403:
                    v11 = v11
                case 9223372036854775807:
                    var v70 float32 = v1
                    _ = v70
                case 32768:
                    if true {
                        v12 = (v12 % ((v12) | 1))
                        if false {
                            if false {
                                var v71 int = v4
                                if false {
                                    v72 := GenS5{F0: false}
                                    var v73 float32 = float32(1.5)
                                    _ = v72
                                    _ = v73
                                }
                                v74 := GenS6{F0: (-(v71))}
                                _ = v71
                                _ = v74
                            }
                        } else {
                            if true {
                                for fi75 := 0; fi75 < 19; fi75++ {
                                    v76 := GenS2{F0: false, F1: v11}
                                    v77 := GenS5{F0: (false || true)}
                                    v4 = (fi69 | (v4 % ((fi69) | 1)))
                                    _ = fi75
                                    _ = v76
                                    _ = v77
                                }
                                if false {
                                    v3 = ""
                                    var v78 float32 = ((v1 * v1) + (v1 + 0.1))
                                    var v79 uint32 = 2
                                    _ = v78
                                    _ = v79
                                } else {
                                    for fi80 := 0; fi80 < 2; fi80++ {
                                        if true {
                                            var v81 int32 = -2147483648
                                            _ = v81
                                        } else {
                                            v12 -= 42
                                        }
                                        _ = fi80
                                    }
                                    var v82 bool = false
                                    var v83 bool = v82
                                    _ = v82
                                    _ = v83
                                }
                            } else {
                                fmt.Println(v4)
                            }
                            if (42 != 32767) {
                                for fi84 := 0; fi84 < 5; fi84++ {
                                    switch (("foo bar" + v3) + "foo bar") {
                                    case "\n\t":
                                        switch (v3 + v3) {
                                        case "\n\t":
                                            switch ("\n\t" + "foo bar") {
                                            case "\n\t":
                                                fmt.Println(fi69, v3, fi84)
                                            default:
                                                v1 = v1
                                            }
                                        case "foo bar":
                                            fmt.Println(v3)
                                        }
                                    case "":
                                        v11 += 2
                                    case "a":
                                        var v85 uint64 = v11
                                        var v86 int32 = -770
                                        _ = v85
                                        _ = v86
                                    default:
                                        var v87 float64 = 1e10
                                        _ = v87
                                    }
                                    if true {
                                        v2 = append(v2, (v3 + v3))
                                    } else {
                                        v3 = v3
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        v11 = (^(v11))
                                        switch int(v11) {
                                        case -9223372036854775808:
                                            v88 := GenS7{F0: ("foo bar" + (v3 + v3))}
                                            _ = v88
                                        default:
                                            if false {
                                                v4 += (^(fi84))
                                            }
                                        }
                                    }
                                    _ = fi84
                                }
                                v89 := GenS8{F0: v12}
                                _ = v89
                            }
                        }
                        var v90 uint = 4294967295
                        _ = v90
                    }
                }
            case 491:
                var v91 uint16 = 65535
                _ = v91
            default:
                v12 = v12
            }
            v92 := GenS3{F0: 1, F1: -2147483648, F2: 255}
            _ = fi69
            _ = v92
        }
    default:
        var v93 int8 = v12
        _ = v93
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
