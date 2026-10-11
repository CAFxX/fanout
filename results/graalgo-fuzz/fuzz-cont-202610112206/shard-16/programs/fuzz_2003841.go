package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 bool }
type GenS2 struct { F0 int64; F1 bool; F2 string }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 uint16; F1 bool }
type GenS5 struct { F0 uint32; F1 bool }
type GenS6 struct { F0 int64; F1 bool; F2 int16 }
type GenS7 struct { F0 int; F1 uint8 }
type GenS8 struct { F0 uint; F1 int8 }
type GenS9 struct { F0 uint32; F1 uint16 }
type GenS10 struct { F0 uint16; F1 int16 }
type GenS11 struct { F0 bool; F1 int; F2 int32 }
type GenS12 struct { F0 int64; F1 uint16; F2 int8 }
func genFn0() string {
    v7 := GenS1{F0: 32767, F1: true}
    var v8 int32 = -464
    for fi9 := 0; fi9 < 7; fi9++ {
        if false {
            switch fi9 {
            case 1:
                var v10 uint32 = 655
                _ = v10
            case 42:
                v11 := GenS2{F0: 106, F1: (1e10 < 0.0), F2: ""}
                _ = v11
            case -148:
                var v12 float64 = (1.5 + 0.1)
                _ = v12
            default:
                var v13 int32 = 1
                _ = v13
            }
        }
        v8 += -2147483648
        var v14 int32 = v8
        _ = fi9
        _ = v14
    }
    v15 := GenS2{F0: -1, F1: (true || false), F2: "\n\t"}
    _ = v15
    _ = v7
    _ = v8
    return (("hello" + ("" + ("" + "\n\t"))) + "a")
}
func genFn1() int64 {
    v16 := []int{32767, -2147483648}
    var v17 uint8 = 2
    _ = v16
    _ = v17
    return -1
}

func main() {
    var v1 int64 = -1
    var v2 int32 = -115
    var v3 uint32 = 4294967295
    v4 := make(map[uint32]int)
    var v5 int32 = -1
    var v6 float32 = 1.5
    for fi18 := 0; fi18 < 7; fi18++ {
        var v19 int32 = (v2 + (v2 * v5))
        _ = fi18
        _ = v19
    }
    v20 := []int{}
    var v21 uint = 4294967295
    v22 := GenS3{F0: 1}
    fmt.Println(v3, v2, v3)
    for fi23 := 0; fi23 < 14; fi23++ {
        switch ("foo bar" + "\n\t") {
        case "\n\t":
            v24 := GenS4{F0: 2, F1: (v3 != (v3 >> ((v21) & 63)))}
            _ = v24
        case "foo bar":
            if false {
                switch "foo bar" {
                case "\n\t":
                    for fi25 := 0; fi25 < 13; fi25++ {
                        fmt.Println(fi25)
                        _ = fi25
                    }
                case "":
                    for fi26 := 0; fi26 < 4; fi26++ {
                        if ((v5 + v5) < (v2 - (v2 + (v5 * -982)))) {
                            var v27 uint8 = 42
                            _ = v27
                        } else {
                            if true {
                                v5 += v5
                                v28 := GenS5{F0: (v3 << ((0) & 63)), F1: false}
                                _ = v28
                            }
                            var v29 int8 = 42
                            var v30 uint32 = v3
                            if true {
                                v20 = append(v20, fi26)
                            } else {
                                var v31 uint16 = 151
                                v2 = (^(v2))
                                _ = v31
                            }
                            var v32 int64 = v1
                            v20 = append(v20, 32768)
                            var v33 float32 = (float32(-2.25) / v6)
                            _ = v29
                            _ = v30
                            _ = v32
                            _ = v33
                        }
                        if (143 >= 1) {
                            if len(v20) >= 2 {
                                v20 = v20[1:len(v20)]
                            }
                        }
                        if len(v20) > 0 {
                            v20[0] = 1
                        }
                        _ = fi26
                    }
                }
                v3 = v3
                switch (fi23 % (((fi23 + 42)) | 1)) {
                case 739:
                    switch ("foo bar" + "foo bar") {
                    case "\n\t":
                        v6 -= 1e10
                    case "":
                        var v34 uint32 = (v3 ^ 332)
                        _ = v34
                    default:
                        v20 = append(v20, (fi23 >> (((^(v21))) & 63)))
                    }
                default:
                    if (1.5 > 0.1) {
                        if true {
                            for fi35 := 0; fi35 < 5; fi35++ {
                                v3 = (v3 - 4294967295)
                                _ = fi35
                            }
                            var v36 uint64 = 0
                            _ = v36
                        }
                        delete(v4, v3)
                        switch ("hello" + "\n\t") {
                        case "foo bar":
                            if true {
                                fmt.Println(v1)
                                var v37 int16 = 0
                                _ = v37
                            } else {
                                v2 += 0
                                v38 := GenS6{F0: genFn1(), F1: true, F2: 42}
                                v20 = append(v20, 127)
                                _ = v38
                            }
                        case "hello":
                            v6 = -2.25
                        case "a":
                            for fi39 := 0; fi39 < 12; fi39++ {
                                if len(v20) >= 2 {
                                    v20 = v20[1:len(v20)]
                                }
                                switch (-(fi39)) {
                                case -198:
                                    _ = len(v20)
                                }
                                if (!true) {
                                    delete(v4, v3)
                                    var v40 uint32 = (^(v3))
                                    _ = v40
                                } else {
                                    v41 := GenS5{F0: v3, F1: false}
                                    if ((18 > 0) || (true || true)) {
                                        switch "hello" {
                                        case "\n\t":
                                            var v42 uint64 = 2
                                            _ = v42
                                        case "a":
                                            v6 += (v6 / float32(1e10))
                                        }
                                        var v43 uint32 = (v3 ^ v3)
                                        v3 -= (v43 * 4294967295)
                                        _ = v43
                                    } else {
                                        v5 = -1
                                    }
                                    v44 := GenS7{F0: fi23, F1: 128}
                                    _ = v41
                                    _ = v44
                                }
                                v21 += (^(v21))
                                if true {
                                    if len(v20) >= 2 {
                                        v20 = v20[1:len(v20)]
                                    }
                                }
                                _ = fi39
                            }
                        }
                    }
                }
            } else {
                for fi45 := 0; fi45 < 5; fi45++ {
                    var v46 bool = (!(true && false))
                    if v46 {
                        v20 = append(v20, (fi45 >> (((^(v21))) & 63)))
                    }
                    fmt.Println(v3, v2, v1)
                    _ = fi45
                    _ = v46
                }
                fmt.Println(v21, v5, v5)
                v6 = ((1.5 / v6) / v6)
            }
        }
        switch int(v3) {
        case 128:
            v47 := GenS6{F0: (v1 & v1), F1: (!true), F2: 1}
            _ = v47
        case 32767:
            if true {
                v6 = (v6 + (v6 - (0.1 + v6)))
                switch (fi23 ^ (fi23 - int(v1))) {
                case 42:
                    for fi48 := 0; fi48 < 13; fi48++ {
                        var v49 string = ""
                        v1 -= v1
                        _ = fi48
                        _ = v49
                    }
                case -26:
                    var v50 uint = (^(v21))
                    _ = v50
                default:
                    var v51 uint8 = 154
                    _ = v51
                }
                if true {
                    v52 := GenS8{F0: v21, F1: 0}
                    v21 += (v21 ^ v21)
                    var v53 uint8 = 229
                    _ = v52
                    _ = v53
                } else {
                    v21 = 65535
                    v5 = 0
                }
            } else {
                var v54 int32 = v5
                _ = v54
            }
        default:
            v2 = v2
        }
        _ = fi23
    }
    if true {
        if true {
            v1 = 0
        }
        for fi55 := 0; fi55 < 20; fi55++ {
            var v56 uint16 = 2
            if false {
                var v57 uint64 = 4294967295
                for fi58 := 0; fi58 < 11; fi58++ {
                    var v59 uint8 = 42
                    _ = fi58
                    _ = v59
                }
                var v60 float64 = ((3.14159 / float64(1e10)) / 3.14159)
                _ = v57
                _ = v60
            } else {
                v6 -= 0.1
                if false {
                    switch (-(fi55)) {
                    case 32768:
                        delete(v4, (v3 >> ((0) & 63)))
                    case 32767:
                        var v61 uint8 = 255
                        _ = v61
                    case -651:
                        var v62 float32 = 3.14159
                        _ = v62
                    }
                    v6 += 1e10
                    if false {
                        if true {
                            v63 := GenS1{F0: v56, F1: true}
                            for fi64 := 0; fi64 < 8; fi64++ {
                                if (genFn1() <= (v1 - v1)) {
                                    var v65 uint = v21
                                    _ = v65
                                }
                                fmt.Println(v21)
                                if true {
                                    var v66 uint64 = 1
                                    fmt.Println(v2, v5)
                                    _ = v66
                                } else {
                                    switch fi55 {
                                    case 2147483647:
                                        switch 128 {
                                        case 4294967295:
                                            v4[(v3 / ((v3) | 1))] = fi64
                                        case 32767:
                                            switch "" {
                                            case "foo bar":
                                                var v67 float32 = (float32(((float64(v6) + -2.25) / 3.14159)) - v6)
                                                _ = v67
                                            case "\n\t":
                                                var v68 int16 = int16(fi64)
                                                _ = v68
                                            case "":
                                                v69 := GenS7{F0: (fi55 * fi64), F1: 42}
                                                _ = v69
                                            default:
                                                v70 := GenS4{F0: (v56 << ((uint(v56)) & 63)), F1: false}
                                                _ = v70
                                            }
                                        }
                                    case 127:
                                        switch (fi64 << (((v21 + 2)) & 63)) {
                                        case 42:
                                            var v71 float64 = -2.25
                                            _ = v71
                                        case -1:
                                            switch fi55 {
                                            case -334:
                                                if false {
                                                    var v72 string = ("" + "a")
                                                    var v73 int = 2147483647
                                                    _ = v72
                                                    _ = v73
                                                } else {
                                                    v6 = (v6 - ((v6 + float32(float64(float32(0.1)))) / v6))
                                                    var v74 int = 1
                                                    _ = v74
                                                }
                                            default:
                                                v75 := GenS9{F0: 2147483647, F1: v56}
                                                _ = v75
                                            }
                                            var v76 string = "foo bar"
                                            v77 := GenS10{F0: (^(v56)), F1: -32768}
                                            v1 = (v1 ^ v1)
                                            v1 = int64(v56)
                                            v20 = append(v20, 127)
                                            _ = v76
                                            _ = v77
                                        case 127:
                                            if len(v20) > 0 {
                                                v20[0] = 42
                                            }
                                        }
                                    case 293:
                                        v21 = v21
                                    }
                                    if true {
                                        v21 += v21
                                        _ = len(v4)
                                        v78 := GenS11{F0: (false && false), F1: (fi55 / (((fi55 * fi64)) | 1)), F2: 42}
                                        _ = v78
                                    }
                                    var v79 float32 = 1e10
                                    var v80 string = "foo bar"
                                    _ = v79
                                    _ = v80
                                }
                                _ = fi64
                            }
                            _ = v63
                        } else {
                            v81 := GenS4{F0: v56, F1: ((0 >= 42) || false)}
                            _ = v81
                        }
                        for fi82 := 0; fi82 < 17; fi82++ {
                            if len(v20) >= 2 {
                                v20 = v20[1:len(v20)]
                            }
                            if len(v20) >= 2 {
                                v20 = v20[1:len(v20)]
                            }
                            _ = fi82
                        }
                    }
                } else {
                    switch (fi55 | fi55) {
                    case 906:
                        v5 -= v5
                    case 1:
                        v5 = v5
                    case -660:
                        var v83 uint16 = (^(v56))
                        _ = v83
                    }
                    v21 = 6
                    if len(v20) > 0 {
                        v20[0] = (fi55 / ((fi55) | 1))
                    }
                }
            }
            _ = fi55
            _ = v56
        }
        v21 -= v21
    } else {
        v84 := GenS7{F0: 1, F1: 127}
        _ = v84
    }
    if (!false) {
        if (false || (true && true)) {
            fmt.Println(v2)
        }
        fmt.Println(v21)
    } else {
        v5 -= (-(v2))
        switch -2147483648 {
        case 4294967295:
            v4[v3] = 9223372036854775807
        case -1:
            v21 = 187
        case -248:
            switch 2147483647 {
            case 57:
                fmt.Println(v5, v21, v5)
            case 2147483647:
                switch "" {
                case "\n\t":
                    for fi85 := 0; fi85 < 15; fi85++ {
                        fmt.Println(v1)
                        _ = fi85
                    }
                case "a":
                    var v86 string = "foo bar"
                    _ = v86
                case "hello":
                    switch "hello" {
                    case "foo bar":
                        v21 -= v21
                    case "a":
                        var v87 bool = (true || false)
                        _ = v87
                    case "\n\t":
                        if true {
                            var v88 bool = (false || ((v21 + 988) == (v21 % ((4294967295) | 1))))
                            fmt.Println(v3, v5)
                            _ = v88
                        }
                    }
                default:
                    switch ("a" + ("a" + "foo bar")) {
                    case "hello":
                        v5 = v5
                    default:
                        v1 = v1
                    }
                }
            case 961:
                v21 = 4294967295
            default:
                v1 = 9223372036854775807
            }
        }
    }
    v6 += v6
    var v89 int64 = genFn1()
    for fi90 := 0; fi90 < 18; fi90++ {
        fmt.Println(fi90)
        _ = fi90
    }
    fmt.Println(v1)
    var v91 int32 = v2
    switch ((("\n\t" + "a") + "hello") + "\n\t") {
    case "hello":
        v92 := GenS2{F0: (v89 ^ (v1 * 0)), F1: false, F2: ("\n\t" + (((("a" + "hello") + "foo bar") + ("hello" + ("\n\t" + "foo bar"))) + ("a" + "a")))}
        _ = v92
    case "a":
        for fi93 := 0; fi93 < 16; fi93++ {
            var v94 int16 = 1000
            v3 = 42
            for fi95 := 0; fi95 < 2; fi95++ {
                fmt.Println(fi93)
                v91 = v5
                var v96 int8 = 127
                _ = fi95
                _ = v96
            }
            _ = fi93
            _ = v94
        }
    case "":
        v97 := GenS12{F0: (v89 & (v89 + v1)), F1: 0, F2: -97}
        _ = v97
    }
    switch "hello" {
    case "\n\t":
        v21 = v21
    default:
        for fi98 := 0; fi98 < 16; fi98++ {
            v21 = v21
            _ = fi98
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v89:", v89)
    fmt.Println("v91:", v91)
    _ = v1
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v89
    _ = v91
}
