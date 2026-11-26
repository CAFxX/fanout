package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint64; F1 int16 }
type GenS3 struct { F0 uint16; F1 uint64 }
type GenS4 struct { F0 uint; F1 uint }
type GenS5 struct { F0 int16; F1 bool; F2 int32 }
type GenS6 struct { F0 uint8 }
type GenS7 struct { F0 uint16; F1 uint64 }
type GenS8 struct { F0 uint8; F1 uint8; F2 string }
func genFn0(p0 uint16) int {
    for fi5 := 0; fi5 < 7; fi5++ {
        var v6 int = fi5
        _ = fi5
        _ = v6
    }
    p0 -= (p0 * (^(p0)))
    _ = p0
    return 0
}
func genFn1(p0 int32, p1 int8) uint64 {
    v7 := GenS1{F0: 42}
    fmt.Println(p0)
    var v8 bool = (!(("" + "\n\t") < ("foo bar" + "foo bar")))
    for fi9 := 0; fi9 < 18; fi9++ {
        p1 = (p1 ^ 1)
        fmt.Println(p0)
        _ = fi9
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 2
}

func main() {
    var v1 int32 = 1
    var v2 uint16 = 32767
    var v3 uint32 = 294
    v4 := make(map[uint32]bool)
    var v10 int64 = -9223372036854775808
    var v11 uint64 = genFn1(v1, 1)
    var v12 int32 = 0
    v2 -= 42
    for fi13 := 0; fi13 < 6; fi13++ {
        for fi14 := 0; fi14 < 16; fi14++ {
            var v15 bool = true
            v16 := []int{fi13, genFn0((^(v2))), fi14}
            _ = fi14
            _ = v15
            _ = v16
        }
        if ("" >= ("a" + "\n\t")) {
            switch (-(fi13)) {
            case 42:
                var v17 int8 = 42
                _ = v17
            case -874:
                v18 := GenS1{F0: 4294967295}
                _ = v18
            case 0:
                if (1 < (^(v11))) {
                    if true {
                        var v19 float64 = float64(float32(0.0))
                        v11 += v11
                        v20 := []int{4294967295, fi13, fi13}
                        _ = v19
                        _ = v20
                    }
                    for fi21 := 0; fi21 < 19; fi21++ {
                        v11 -= 2
                        v22 := []int{}
                        _ = fi21
                        _ = v22
                    }
                }
            default:
                var v23 string = "hello"
                _ = v23
            }
        } else {
            var v24 int8 = 25
            _ = v24
        }
        if true {
            if false {
                v25 := []int{int(v2), (fi13 - fi13), 32768}
                if len(v25) >= 2 {
                    v25 = v25[1:len(v25)]
                }
                _ = v25
            }
        } else {
            var v26 float64 = 0.0
            switch (fi13 | genFn0(uint16(v3))) {
            case 9223372036854775807:
                if (0 >= 127) {
                    switch "" {
                    case "":
                        var v27 int = 32767
                        _ = v27
                    }
                    v28 := []int{}
                    _ = v28
                }
            case 32768:
                var v29 int16 = 32767
                _ = v29
            default:
                if false {
                    var v30 bool = false
                    _ = v30
                } else {
                    v26 = v26
                    if (128 >= 0) {
                        for fi31 := 0; fi31 < 10; fi31++ {
                            if true {
                                var v32 uint = 2
                                var v33 string = (("\n\t" + "\n\t") + "a")
                                if true {
                                    v1 -= 154
                                    v10 -= v10
                                    fmt.Println(v12, v12, v32)
                                }
                                _ = v32
                                _ = v33
                            }
                            _ = fi31
                        }
                        if (!(4294967295 != (^(v3)))) {
                            switch (fi13 * genFn0(65535)) {
                            case 1:
                                if false {
                                    if true {
                                        v26 = v26
                                        if (!(v11 == (v11 | (v11 & (v11 >> ((1) & 63)))))) {
                                            v2 += v2
                                            v34 := []int{(-(fi13))}
                                            if false {
                                                var v35 float32 = 0.1
                                                _ = v35
                                            }
                                            _ = v34
                                        } else {
                                            switch (fi13 << ((uint(v3)) & 63)) {
                                            case 4294967295:
                                                _ = len(v4)
                                            case 166:
                                                fmt.Println(fi13, v3, v10)
                                            case 2147483647:
                                                switch (fi13 << ((102) & 63)) {
                                                case 814:
                                                    for fi36 := 0; fi36 < 16; fi36++ {
                                                        v11 = uint64(v2)
                                                        _ = fi36
                                                    }
                                                case 42:
                                                    if true {
                                                        var v37 float32 = (0.0 - 3.14159)
                                                        _ = v37
                                                    }
                                                case 32767:
                                                    v38 := GenS1{F0: 255}
                                                    _ = v38
                                                default:
                                                    fmt.Println(v10)
                                                }
                                            }
                                        }
                                        v4[(v3 | (^(v3)))] = false
                                    }
                                }
                            default:
                                var v39 uint = 272
                                _ = v39
                            }
                            switch "foo bar" {
                            case "foo bar":
                                switch (fi13 << ((255) & 63)) {
                                case -739:
                                    var v40 float64 = v26
                                    _ = v40
                                case -187:
                                    var v41 string = ((("" + "a") + ("" + (("foo bar" + "\n\t") + ("\n\t" + "a")))) + "a")
                                    _ = v41
                                case 455:
                                    switch genFn0((v2 - v2)) {
                                    case 32768:
                                        v42 := GenS1{F0: 255}
                                        _ = v42
                                    }
                                default:
                                    v4[(^(v3))] = (!(!((("" + "a") + "foo bar") >= "\n\t")))
                                }
                            }
                        }
                    } else {
                        v26 += (3.14159 - 3.14159)
                        var v43 int8 = -46
                        _ = len(v4)
                        _ = v43
                    }
                    if ((v10 >= -41) && true) {
                        for fi44 := 0; fi44 < 16; fi44++ {
                            fmt.Println(v3, v11, v26)
                            var v45 uint8 = 154
                            v12 = (v1 | (v1 ^ 230))
                            _ = fi44
                            _ = v45
                        }
                        fmt.Println(v26, v1, v1)
                    } else {
                        for fi46 := 0; fi46 < 2; fi46++ {
                            var v47 int32 = (v12 | (v12 / ((v12) | 1)))
                            _ = fi46
                            _ = v47
                        }
                    }
                }
            }
            v2 += (v2 + (v2 >> ((0) & 63)))
            _ = v26
        }
        _ = fi13
    }
    if true {
        fmt.Println(v11)
    } else {
        for fi48 := 0; fi48 < 4; fi48++ {
            if true {
                var v49 uint8 = 217
                var v50 uint16 = v2
                _ = v49
                _ = v50
            }
            v51 := GenS2{F0: genFn1(-548, 0), F1: 1}
            _ = fi48
            _ = v51
        }
        switch 128 {
        case 0:
            var v52 int32 = (v12 + 19)
            _ = v52
        }
    }
    var v53 int8 = 0
    for fi54 := 0; fi54 < 1; fi54++ {
        v1 += v1
        for fi55 := 0; fi55 < 12; fi55++ {
            v56 := GenS3{F0: 42, F1: (v11 % (((v11 & uint64(v2))) | 1))}
            v57 := v4[v3]
            var v58 float64 = 3.14159
            _ = fi55
            _ = v56
            _ = v57
            _ = v58
        }
        fmt.Println(v3, v10)
        _ = fi54
    }
    v59 := GenS4{F0: 285, F1: 65535}
    if true {
        v2 -= (v2 * (v2 >> ((184) & 63)))
        if true {
            var v60 bool = true
            switch ("\n\t" + (((((("a" + "foo bar") + "a") + ("foo bar" + "")) + "\n\t") + "") + ("" + ("foo bar" + "\n\t")))) {
            case "hello":
                v61 := GenS3{F0: v2, F1: (v11 / ((0) | 1))}
                _ = v61
            default:
                if (v11 != (v11 << ((326) & 63))) {
                    if v60 {
                        if v60 {
                            for fi62 := 0; fi62 < 15; fi62++ {
                                v63 := GenS5{F0: int16(v10), F1: v60, F2: v1}
                                var v64 bool = (1.5 != ((3.14159 - ((0.0 - (((0.1 - (0.1 - 1e10)) / (((float32(0.1) + 1e10) + 1.5) + float32(1e10))) * ((((0.0 * -2.25) * ((((0.1 * (1e10 * (-2.25 + 1.5))) + ((-2.25 - -2.25) / (-2.25 * (0.0 - float32(float64(-2.25)))))) / 1.5) / (float32((float64(1e10) + (float64(-2.25) / 1.5))) / float32(-2.25)))) - 0.1) / 1.0))) / 3.14159)) - (-2.25 + 3.14159)))
                                v65 := GenS3{F0: v2, F1: 2}
                                _ = fi62
                                _ = v63
                                _ = v64
                                _ = v65
                            }
                        }
                        for fi66 := 0; fi66 < 18; fi66++ {
                            v53 += 1
                            _ = fi66
                        }
                    }
                } else {
                    for fi67 := 0; fi67 < 2; fi67++ {
                        switch (fi67 / ((int(v11)) | 1)) {
                        case 1:
                            for fi68 := 0; fi68 < 10; fi68++ {
                                v69 := []int{(fi67 % ((genFn0((v2 + (v2 + (v2 >> ((762) & 63)))))) | 1))}
                                v60 = (18446744073709551615 > uint(v2))
                                _ = fi68
                                _ = v69
                            }
                        case 111:
                            switch (fi67 >> ((uint(v2)) & 63)) {
                            case 9223372036854775807:
                                v70 := GenS6{F0: 42}
                                _ = v70
                            case -891:
                                for fi71 := 0; fi71 < 18; fi71++ {
                                    if ("\n\t" > "") {
                                        if (255 != 1) {
                                            v72 := []int{}
                                            _ = v72
                                        }
                                        switch fi71 {
                                        case -516:
                                            var v73 int8 = v53
                                            _ = v73
                                        case 162:
                                            v60 = v60
                                        case 128:
                                            switch (fi71 | fi67) {
                                            case 127:
                                                var v74 uint64 = (v11 / ((42) | 1))
                                                _ = v74
                                            case 32767:
                                                if (413 >= 1) {
                                                    if (v60 && ((255 < 42) && (44 < 42))) {
                                                        switch "hello" {
                                                        case "hello":
                                                            var v75 int = (fi67 | fi67)
                                                            _ = v75
                                                        case "":
                                                            var v76 int64 = (v10 / (((v10 >> ((0) & 63))) | 1))
                                                            _ = v76
                                                        default:
                                                            v12 = (v1 & -861)
                                                        }
                                                        switch (-(fi67)) {
                                                        case 750:
                                                            var v77 float32 = -2.25
                                                            _ = v77
                                                        case -183:
                                                            switch "foo bar" {
                                                            case "foo bar":
                                                                var v78 int8 = (v53 & v53)
                                                                _ = v78
                                                            default:
                                                                v79 := GenS2{F0: v11, F1: 26}
                                                                _ = v79
                                                            }
                                                        case -9223372036854775808:
                                                            v11 += 2
                                                        }
                                                        v60 = (v60 && v60)
                                                    }
                                                }
                                            case -1:
                                                var v80 uint = 959
                                                _ = v80
                                            }
                                        }
                                        v81 := GenS7{F0: v2, F1: (v11 / (((^(v11))) | 1))}
                                        _ = v81
                                    }
                                    var v82 uint = 65535
                                    _ = fi71
                                    _ = v82
                                }
                            case 2147483647:
                                v3 = (v3 << ((0) & 63))
                            default:
                                for fi83 := 0; fi83 < 19; fi83++ {
                                    var v84 int32 = (v1 + (v1 / (((^(v1))) | 1)))
                                    v85 := []int{}
                                    _ = len(v4)
                                    switch (-(fi83)) {
                                    case 905:
                                        v86 := GenS8{F0: 1, F1: 157, F2: "hello"}
                                        _ = v86
                                    case 32767:
                                        v84 += v1
                                    }
                                    v87 := GenS1{F0: 4294967295}
                                    _ = fi83
                                    _ = v84
                                    _ = v85
                                    _ = v87
                                }
                            }
                        case 127:
                            v88 := GenS3{F0: v2, F1: (^(v11))}
                            _ = v88
                        default:
                            var v89 float64 = (3.14159 / 3.14159)
                            _ = v89
                        }
                        _ = fi67
                    }
                    v90 := []int{128, -1, -971, 570}
                    var v91 int16 = 0
                    _ = v90
                    _ = v91
                }
            }
            _ = v60
        }
    }
    var v92 float32 = 1.5
    var v93 int32 = int32(v11)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v53:", v53)
    fmt.Println("v92:", v92)
    fmt.Println("v93:", v93)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v53
    _ = v59
    _ = v92
    _ = v93
}
