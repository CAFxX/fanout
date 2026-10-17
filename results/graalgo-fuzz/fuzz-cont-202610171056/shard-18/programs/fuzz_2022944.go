package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 string; F2 uint32 }
type GenS2 struct { F0 int64; F1 uint32 }
type GenS3 struct { F0 uint64 }
type GenS4 struct { F0 int; F1 uint16; F2 uint32 }
type GenS5 struct { F0 int; F1 int; F2 uint64 }
type GenS6 struct { F0 int64; F1 int16; F2 bool }
type GenS7 struct { F0 int }
type GenS8 struct { F0 uint8; F1 uint8 }
func genFn0() int64 {
    if true {
        var v5 float64 = 0.0
        var v6 int8 = -1
        _ = v5
        _ = v6
    }
    if false {
        var v7 string = "foo bar"
        v8 := []int{-1, -9223372036854775808, 42, 1}
        _ = len(v8)
        _ = v7
        _ = v8
    } else {
        for fi9 := 0; fi9 < 1; fi9++ {
            if (!true) {
                v10 := []int{(fi9 ^ (fi9 * fi9)), (fi9 * fi9), (fi9 * 2147483647), (-(fi9))}
                var v11 uint32 = 1
                var v12 uint64 = 42
                _ = v10
                _ = v11
                _ = v12
            } else {
                if ((float32(0.1) / 1.0) != ((-2.25 * 3.14159) * (((0.0 / -2.25) - ((0.1 - ((0.0 - (-2.25 / ((1e10 / (0.1 * 0.1)) / (((float32(0.0) / 3.14159) - 3.14159) - 1.5)))) + 3.14159)) / 1e10)) * 1.5))) {
                    v13 := GenS1{F0: 128, F1: "hello", F2: 311}
                    fmt.Println(fi9)
                    var v14 uint = 42
                    _ = v13
                    _ = v14
                }
            }
            if (!(2 < 1)) {
                if (!(128 <= 196)) {
                    var v15 uint32 = 1
                    for fi16 := 0; fi16 < 3; fi16++ {
                        if false {
                            if true {
                                v15 = 945
                                switch fi16 {
                                case 42:
                                    if false {
                                        v15 += (v15 ^ (v15 | 65535))
                                        v15 += (v15 * 2)
                                    } else {
                                        v15 = 4294967295
                                        v17 := []int{(fi9 / (((fi16 / ((fi16) | 1))) | 1)), fi16, fi16, fi16}
                                        _ = v17
                                    }
                                case -128:
                                    var v18 float64 = (3.14159 / 0.1)
                                    _ = v18
                                case 137:
                                    switch (("\n\t" + "\n\t") + "\n\t") {
                                    case "hello":
                                        v15 += (v15 + 0)
                                        v15 -= (v15 >> ((1) & 63))
                                    }
                                }
                            }
                        }
                        _ = fi16
                    }
                    _ = v15
                } else {
                    var v19 int8 = 42
                    if (1.5 <= 0.1) {
                        switch "" {
                        case "":
                            var v20 int = 4294967295
                            _ = v20
                        case "foo bar":
                            v21 := GenS2{F0: 796, F1: 2}
                            _ = v21
                        default:
                            v19 = -21
                        }
                        v19 += 42
                    } else {
                        v19 = v19
                    }
                    _ = v19
                }
                if ((42 < -128) && (4294967295 > 0)) {
                    v22 := []int{697, 128, -9223372036854775808, (fi9 | (fi9 + (fi9 - (-(fi9)))))}
                    v23 := GenS1{F0: 42, F1: "", F2: 2}
                    if ((-2147483648 > -2147483648) && true) {
                        var v24 int8 = 127
                        _ = v24
                    }
                    _ = v22
                    _ = v23
                }
            }
            _ = fi9
        }
        var v25 int8 = 1
        _ = v25
    }
    var v26 int32 = -1
    for fi27 := 0; fi27 < 12; fi27++ {
        v28 := GenS3{F0: 4294967295}
        v29 := []int{}
        v26 -= -234
        _ = fi27
        _ = v28
        _ = v29
    }
    for fi30 := 0; fi30 < 18; fi30++ {
        v31 := []int{fi30}
        switch (("foo bar" + "a") + "a") {
        case "":
            switch "" {
            case "hello":
                _ = len(v31)
            case "foo bar":
                if false {
                    switch ((("hello" + "a") + ("foo bar" + "foo bar")) + ("foo bar" + ("a" + "a"))) {
                    case "\n\t":
                        switch (fi30 * (-(fi30))) {
                        case 593:
                            for fi32 := 0; fi32 < 7; fi32++ {
                                v33 := GenS1{F0: 2, F1: "\n\t", F2: 2147483648}
                                _ = fi32
                                _ = v33
                            }
                        case 2147483647:
                            var v34 float32 = 1.5
                            _ = v34
                        case -2147483648:
                            _ = len(v31)
                        default:
                            v31 = append(v31, (fi30 % ((32767) | 1)))
                        }
                    case "hello":
                        v35 := GenS3{F0: 2}
                        _ = v35
                    case "a":
                        var v36 uint64 = 2
                        _ = v36
                    }
                    v26 += 976
                }
            case "":
                if ((-1 < -592) || false) {
                    v26 = (v26 | v26)
                    var v37 int8 = 56
                    _ = v37
                }
            }
        case "a":
            v38 := GenS2{F0: 9223372036854775807, F1: 1}
            _ = v38
        }
        _ = len(v31)
        _ = fi30
        _ = v31
    }
    _ = v26
    return 42
}

func main() {
    var v1 int8 = 100
    var v2 int = -207
    var v3 uint8 = 0
    v4 := []int{}
    v1 -= v1
    var v39 float32 = float32((3.14159 / -2.25))
    fmt.Println(v39)
    _ = len(v4)
    var v40 int32 = -178
    v41 := GenS3{F0: 2}
    fmt.Println(v3, v40)
    v1 = (v1 - 1)
    if false {
        if true {
            v2 = (v2 - (v2 % (((v2 / ((v2) | 1))) | 1)))
            var v42 int16 = -32768
            _ = v42
        } else {
            for fi43 := 0; fi43 < 6; fi43++ {
                v4 = append(v4, 32768)
                _ = fi43
            }
            v1 -= (v1 | v1)
        }
        if (!false) {
            if len(v4) > 0 {
                v4[0] = v2
            }
            var v44 int = (v2 % ((127) | 1))
            v40 = (v40 % (((v40 << ((65535) & 63))) | 1))
            _ = v44
        } else {
            var v45 uint8 = 128
            if (((3.14159 / 3.14159) + 1.5) != (0.1 - (v39 + v39))) {
                if false {
                    switch ((("a" + "\n\t") + ("a" + ("foo bar" + "\n\t"))) + "a") {
                    case "\n\t":
                        v39 += 1.5
                    case "":
                        for fi46 := 0; fi46 < 14; fi46++ {
                            v1 += v1
                            v47 := GenS2{F0: -9223372036854775808, F1: 42}
                            v45 += v45
                            _ = fi46
                            _ = v47
                        }
                    case "hello":
                        var v48 int8 = -125
                        _ = v48
                    }
                    v49 := GenS3{F0: 18446744073709551615}
                    for fi50 := 0; fi50 < 12; fi50++ {
                        if len(v4) >= 2 {
                            v4 = v4[1:len(v4)]
                        }
                        var v51 uint16 = 42
                        v52 := GenS2{F0: int64(v45), F1: 2147483648}
                        _ = fi50
                        _ = v51
                        _ = v52
                    }
                    _ = v49
                }
                var v53 float64 = (-2.25 * 1.5)
                _ = v53
            }
            _ = v45
        }
        fmt.Println(v40, v2)
    } else {
        var v54 int16 = -32768
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        if ((!(v54 < (v54 << ((255) & 63)))) || false) {
            if (("\n\t" + ((("\n\t" + "") + "hello") + "")) == "a") {
                for fi55 := 0; fi55 < 4; fi55++ {
                    if (false || true) {
                        if len(v4) > 0 {
                            v4[0] = v2
                        }
                    }
                    if ((v54 ^ 42) <= 42) {
                        v56 := GenS3{F0: 18446744073709551615}
                        var v57 float32 = 0.0
                        var v58 float32 = (-2.25 - (v57 + v57))
                        _ = v56
                        _ = v57
                        _ = v58
                    }
                    v3 += (v3 ^ 255)
                    _ = fi55
                }
            } else {
                var v59 uint8 = v3
                if ((77 > 2147483648) && false) {
                    v1 -= (v1 * (-(v1)))
                    for fi60 := 0; fi60 < 4; fi60++ {
                        var v61 bool = (true || (true && (!(42 != 444))))
                        _ = fi60
                        _ = v61
                    }
                }
                for fi62 := 0; fi62 < 6; fi62++ {
                    for fi63 := 0; fi63 < 14; fi63++ {
                        var v64 int64 = 9223372036854775807
                        _ = fi63
                        _ = v64
                    }
                    _ = fi62
                }
                _ = v59
            }
        } else {
            v39 += v39
            v4 = append(v4, int(v3))
            if false {
                v3 = (^(v3))
            } else {
                v1 += (v1 + 26)
                for fi65 := 0; fi65 < 12; fi65++ {
                    var v66 uint16 = 65535
                    fmt.Println(v39, v66, v2)
                    if true {
                        v54 = (^(v54))
                    } else {
                        v4 = append(v4, fi65)
                        for fi67 := 0; fi67 < 18; fi67++ {
                            switch -2147483648 {
                            case -285:
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            case -745:
                                if false {
                                    v3 += (^(v3))
                                    var v68 float32 = v39
                                    fmt.Println(v66, fi65)
                                    switch (v2 / ((fi65) | 1)) {
                                    case 32768:
                                        v4 = append(v4, fi65)
                                    case 128:
                                        if (v68 < v39) {
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                        } else {
                                            switch ("foo bar" + ("\n\t" + "")) {
                                            case "\n\t":
                                                v40 = (-(v40))
                                            case "":
                                                var v69 int8 = (v1 - v1)
                                                _ = v69
                                            default:
                                                var v70 uint = 584
                                                _ = v70
                                            }
                                            v71 := GenS4{F0: (fi67 | (fi67 - fi67)), F1: uint16(v3), F2: 2147483647}
                                            _ = v71
                                        }
                                    case -1:
                                        v72 := GenS3{F0: 465}
                                        fmt.Println(fi67, v1, v40)
                                        var v73 uint8 = v3
                                        _ = v72
                                        _ = v73
                                    }
                                    _ = v68
                                }
                            default:
                                switch "hello" {
                                case "a":
                                    v3 -= (v3 % (((v3 - 127)) | 1))
                                case "hello":
                                    var v74 int = (fi65 & fi65)
                                    _ = v74
                                case "\n\t":
                                    if false {
                                        var v75 float64 = (float64(3.14159) * float64((1e10 - v39)))
                                        _ = v75
                                    }
                                }
                            }
                            _ = len(v4)
                            if false {
                                if (1e10 > float32(-2.25)) {
                                    fmt.Println(v3, v3, v2)
                                    switch "foo bar" {
                                    case "a":
                                        v54 = 352
                                    }
                                    var v76 uint8 = (^(v3))
                                    _ = v76
                                }
                            } else {
                                var v77 uint8 = v3
                                if len(v4) > 0 {
                                    v4[0] = fi67
                                }
                                v78 := GenS5{F0: v2, F1: (-(fi65)), F2: 18446744073709551615}
                                _ = v77
                                _ = v78
                            }
                            _ = fi67
                        }
                    }
                    _ = fi65
                    _ = v66
                }
                var v79 uint16 = 65535
                _ = v79
            }
        }
        _ = v54
    }
    v40 = (v40 / ((v40) | 1))
    for fi80 := 0; fi80 < 4; fi80++ {
        v81 := GenS6{F0: 176, F1: -953, F2: true}
        var v82 float64 = float64(v39)
        _ = fi80
        _ = v81
        _ = v82
    }
    v4 = append(v4, (v2 * (v2 | (v2 | v2))))
    v39 += v39
    fmt.Println(v1)
    v1 -= 50
    v83 := GenS7{F0: v2}
    for fi84 := 0; fi84 < 6; fi84++ {
        for fi85 := 0; fi85 < 18; fi85++ {
            fmt.Println(v40, v40, fi85)
            if true {
                var v86 string = ""
                switch v86 {
                case "hello":
                    v1 = 1
                case "\n\t":
                    v39 -= (v39 + ((v39 * ((((v39 + v39) / v39) / v39) + 3.14159)) - ((v39 + float32(3.14159)) - (0.0 / v39))))
                case "a":
                    v86 = ((v86 + v86) + v86)
                default:
                    v87 := GenS3{F0: 645}
                    _ = v87
                }
                var v88 string = ("" + ((v86 + "foo bar") + "\n\t"))
                _ = v86
                _ = v88
            } else {
                v2 += fi84
                var v89 int8 = v1
                _ = len(v4)
                _ = v89
            }
            var v90 uint = 2
            var v91 int16 = 42
            if (((v40 + (v40 << ((999) & 63))) <= int32(v90)) && false) {
                v92 := GenS8{F0: (^(v3)), F1: v3}
                v1 = (v1 >> (((v90 + (v90 * (v90 + v90)))) & 63))
                _ = v92
            } else {
                v4 = append(v4, 0)
                if false {
                    var v93 uint16 = uint16(v90)
                    _ = v93
                } else {
                    switch 9223372036854775807 {
                    case 32767:
                        var v94 int = (fi85 >> ((4294967295) & 63))
                        _ = v94
                    case 546:
                        v39 = v39
                    }
                    var v95 uint32 = 989
                    v3 = v3
                    _ = v95
                }
            }
            _ = fi85
            _ = v90
            _ = v91
        }
        v1 += v1
        for fi96 := 0; fi96 < 20; fi96++ {
            v39 += (v39 * v39)
            _ = fi96
        }
        _ = fi84
    }
    fmt.Println(v1, v39, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v40:", v40)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v40
    _ = v41
    _ = v83
}
