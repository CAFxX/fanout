package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 uint16 }
type GenS2 struct { F0 int64; F1 uint; F2 int16 }
type GenS3 struct { F0 string; F1 int; F2 uint8 }
type GenS4 struct { F0 int16; F1 int64 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 bool }
type GenS7 struct { F0 uint32 }
type GenS8 struct { F0 uint32 }
type GenS9 struct { F0 uint16 }
type GenS10 struct { F0 int64; F1 uint32; F2 int64 }
func genFn0() int64 {
    var v5 int = -9223372036854775808
    v6 := []int{v5, (-(v5)), 4294967295, (v5 << ((230) & 63))}
    for fi7 := 0; fi7 < 4; fi7++ {
        v8 := GenS1{F0: 42, F1: 255, F2: 192}
        _ = fi7
        _ = v8
    }
    _ = v5
    _ = v6
    return -502
}

func main() {
    var v1 uint32 = 2
    v2 := []uint32{2147483648, (v1 & (v1 + v1)), 2, v1}
    var v3 float32 = (0.1 + (1.5 - (0.1 - 3.14159)))
    var v4 int16 = -333
    if (32767 >= 42) {
        v9 := GenS1{F0: (v4 | (v4 * v4)), F1: 4294967295, F2: 336}
        _ = v9
    } else {
        for fi10 := 0; fi10 < 11; fi10++ {
            switch ("" + ("\n\t" + "foo bar")) {
            case "":
                v11 := GenS1{F0: 1, F1: 65535, F2: 42}
                _ = v11
            case "foo bar":
                if (32767 != 255) {
                    switch ("foo bar" + "a") {
                    case "foo bar":
                        switch "a" {
                        case "\n\t":
                            v1 += (v1 >> ((255) & 63))
                        case "hello":
                            for fi12 := 0; fi12 < 8; fi12++ {
                                v13 := GenS1{F0: 123, F1: 4294967295, F2: 2}
                                v14 := GenS2{F0: genFn0(), F1: 18446744073709551615, F2: int16(v1)}
                                _ = fi12
                                _ = v13
                                _ = v14
                            }
                        default:
                            for fi15 := 0; fi15 < 10; fi15++ {
                                var v16 int = 794
                                _ = fi15
                                _ = v16
                            }
                        }
                    }
                    v17 := GenS3{F0: "hello", F1: 453, F2: 1}
                    _ = v17
                }
            case "hello":
                v4 = (v4 << ((125) & 63))
            }
            _ = fi10
        }
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v18 := GenS4{F0: v4, F1: 42}
    if ((true || false) || true) {
        fmt.Println(v3)
    } else {
        fmt.Println(v4, v3, v3)
    }
    if (-2.25 >= (v3 * 0.0)) {
        var v19 string = "foo bar"
        v4 -= -1
        for fi20 := 0; fi20 < 14; fi20++ {
            v21 := GenS3{F0: (v19 + ""), F1: 42, F2: 2}
            switch 1 {
            case -1:
                for fi22 := 0; fi22 < 5; fi22++ {
                    switch 4294967295 {
                    case -1:
                        var v23 uint16 = 879
                        _ = v23
                    default:
                        v24 := GenS5{F0: 42}
                        _ = v24
                    }
                    var v25 int8 = -128
                    _ = len(v2)
                    _ = fi22
                    _ = v25
                }
            case 2147483647:
                var v26 uint16 = 255
                _ = v26
            default:
                _ = len(v2)
            }
            _ = fi20
            _ = v21
        }
        _ = v19
    }
    if true {
        for fi27 := 0; fi27 < 10; fi27++ {
            var v28 int = fi27
            if (v1 >= 4294967295) {
                switch -9223372036854775808 {
                case -1:
                    for fi29 := 0; fi29 < 9; fi29++ {
                        switch (v28 << ((42) & 63)) {
                        case 85:
                            v30 := GenS1{F0: v4, F1: 255, F2: 1}
                            _ = v30
                        default:
                            v1 = 65535
                        }
                        var v31 bool = true
                        if ((!v31) || false) {
                            switch ("hello" + "\n\t") {
                            case "\n\t":
                                switch "foo bar" {
                                case "foo bar":
                                    switch v28 {
                                    case -646:
                                        v32 := GenS3{F0: "a", F1: fi27, F2: 128}
                                        _ = v32
                                    case 2147483647:
                                        switch fi27 {
                                        case 710:
                                            var v33 uint8 = 209
                                            _ = v33
                                        case -1:
                                            if v31 {
                                                v34 := GenS6{F0: v31}
                                                _ = v34
                                            }
                                        case 0:
                                            v31 = (-752 >= 0)
                                        }
                                    }
                                    switch 42 {
                                    case 128:
                                        if ((42 >= -2147483648) && ((9223372036854775807 == 419) && false)) {
                                            v35 := GenS2{F0: genFn0(), F1: 204, F2: (v4 / (((v4 ^ 1)) | 1))}
                                            var v36 int64 = genFn0()
                                            if true {
                                                if len(v2) >= 2 {
                                                    v2 = v2[1:len(v2)]
                                                }
                                                v37 := GenS1{F0: 42, F1: 18446744073709551615, F2: 85}
                                                _ = v37
                                            } else {
                                                var v38 int16 = v4
                                                if (((-2.25 + float64((1e10 * 3.14159))) * (3.14159 - 1.5)) == -2.25) {
                                                    var v39 int64 = v36
                                                    v31 = v31
                                                    _ = v39
                                                } else {
                                                    var v40 int8 = 127
                                                    if (v3 >= (((-2.25 + v3) + 0.0) - ((1e10 + v3) * 0.1))) {
                                                        var v41 float64 = float64((3.14159 + 1e10))
                                                        _ = v41
                                                    } else {
                                                        var v42 uint = 255
                                                        v43 := GenS7{F0: (v1 >> ((v42) & 63))}
                                                        _ = v42
                                                        _ = v43
                                                    }
                                                    _ = v40
                                                }
                                                _ = v38
                                            }
                                            _ = v35
                                            _ = v36
                                        } else {
                                            var v44 int16 = (v4 + v4)
                                            v45 := GenS3{F0: "a", F1: (fi27 << ((65535) & 63)), F2: 195}
                                            if false {
                                                var v46 int16 = -103
                                                var v47 uint8 = 197
                                                if v31 {
                                                    var v48 string = ""
                                                    _ = v48
                                                } else {
                                                    if true {
                                                        v31 = true
                                                        v3 += (-2.25 * v3)
                                                        if v31 {
                                                            v47 = (v47 & (v47 - v47))
                                                            v47 = (v47 - (v47 * 128))
                                                            v31 = v31
                                                        } else {
                                                            v44 -= -32768
                                                            if len(v2) >= 2 {
                                                                v2 = v2[1:len(v2)]
                                                            }
                                                        }
                                                        v1 += 448
                                                        if ((v47 % (((v47 ^ v47)) | 1)) <= v47) {
                                                            v31 = v31
                                                            switch ("hello" + (("" + ("" + "a")) + ("hello" + ((("\n\t" + ("hello" + ("foo bar" + "a"))) + "") + (("a" + "") + "\n\t"))))) {
                                                            case "":
                                                                v3 = ((3.14159 + v3) * (v3 - (0.1 / float32((-2.25 - 0.0)))))
                                                            case "foo bar":
                                                                if v31 {
                                                                    if len(v2) > 0 {
                                                                        v2[0] = v1
                                                                    }
                                                                    fmt.Println(v47)
                                                                    v2 = append(v2, (v1 << ((255) & 63)))
                                                                    var v49 float32 = 1e10
                                                                    var v50 uint32 = v1
                                                                    _ = v49
                                                                    _ = v50
                                                                } else {
                                                                    v1 += (v1 + v1)
                                                                    v51 := GenS8{F0: v1}
                                                                    _ = v51
                                                                }
                                                            case "a":
                                                                var v52 int32 = -103
                                                                _ = v52
                                                            }
                                                            v47 -= v47
                                                        }
                                                    } else {
                                                        v44 -= (v4 - (v44 >> ((uint(v1)) & 63)))
                                                        if (v31 || (!v31)) {
                                                            if v31 {
                                                                var v53 uint32 = 4294967295
                                                                if v31 {
                                                                    if len(v2) >= 2 {
                                                                        v2 = v2[1:len(v2)]
                                                                    }
                                                                }
                                                                _ = v53
                                                            } else {
                                                                v31 = (-633 == int32(v47))
                                                                if len(v2) > 0 {
                                                                    v2[0] = (v1 % ((2147483648) | 1))
                                                                }
                                                                fmt.Println(v3, v3, v46)
                                                            }
                                                        } else {
                                                            _ = len(v2)
                                                        }
                                                        v54 := GenS4{F0: v4, F1: -857}
                                                        _ = v54
                                                    }
                                                }
                                                _ = v46
                                                _ = v47
                                            } else {
                                                v31 = (-102 > v4)
                                                v55 := GenS9{F0: 2}
                                                var v56 bool = v31
                                                _ = v55
                                                _ = v56
                                            }
                                            _ = v44
                                            _ = v45
                                        }
                                    case 2147483647:
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    default:
                                        v31 = (555 >= 9223372036854775807)
                                    }
                                }
                            default:
                                if v31 {
                                    v1 = (^(v1))
                                } else {
                                    v57 := GenS7{F0: v1}
                                    _ = v57
                                }
                            }
                        }
                        _ = fi29
                        _ = v31
                    }
                case -994:
                    var v58 uint8 = 2
                    _ = v58
                case 9223372036854775807:
                    if true {
                        v3 -= (1.5 - v3)
                    } else {
                        v3 = v3
                        switch "" {
                        case "":
                            switch ("hello" + "a") {
                            case "a":
                                if len(v2) > 0 {
                                    v2[0] = (v1 / ((v1) | 1))
                                }
                            }
                        case "hello":
                            switch "\n\t" {
                            case "foo bar":
                                var v59 int16 = v4
                                _ = v59
                            case "":
                                v1 -= (v1 << ((302) & 63))
                            case "a":
                                v3 = v3
                            default:
                                var v60 int = 4294967295
                                _ = v60
                            }
                        default:
                            var v61 int = (v28 * (fi27 | 128))
                            _ = v61
                        }
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                    }
                }
                switch (v28 * fi27) {
                case 2147483647:
                    if false {
                        var v62 float32 = v3
                        if len(v2) >= 2 {
                            v2 = v2[1:len(v2)]
                        }
                        _ = v62
                    }
                case 32767:
                    switch "a" {
                    case "foo bar":
                        if false {
                            v3 = v3
                        }
                    case "hello":
                        var v63 int16 = (v4 / (((v4 & (v4 - 0))) | 1))
                        _ = v63
                    default:
                        for fi64 := 0; fi64 < 17; fi64++ {
                            v3 = v3
                            v4 = v4
                            fmt.Println(v28, v3, v3)
                            _ = fi64
                        }
                    }
                case -880:
                    var v65 int = -2147483648
                    _ = v65
                default:
                    if false {
                        fmt.Println(v3, v3)
                    }
                }
            } else {
                switch "" {
                case "":
                    v66 := GenS2{F0: 432, F1: 337, F2: v4}
                    _ = v66
                default:
                    if (127 <= -59) {
                        var v67 uint32 = 1
                        switch (-(fi27)) {
                        case 529:
                            for fi68 := 0; fi68 < 18; fi68++ {
                                v1 = 2147483647
                                _ = fi68
                            }
                        default:
                            if (1 == 4294967295) {
                                for fi69 := 0; fi69 < 14; fi69++ {
                                    v70 := GenS3{F0: "foo bar", F1: -121, F2: 0}
                                    var v71 int64 = 262
                                    v3 += 3.14159
                                    _ = fi69
                                    _ = v70
                                    _ = v71
                                }
                            } else {
                                v67 = (v67 + v67)
                                for fi72 := 0; fi72 < 10; fi72++ {
                                    switch ((("a" + "") + "\n\t") + "\n\t") {
                                    case "":
                                        var v73 int64 = -1
                                        var v74 float64 = (-2.25 + ((float64(v3) + -2.25) * (1e10 * (-2.25 * (3.14159 - float64(v3))))))
                                        v75 := GenS8{F0: 65535}
                                        _ = v73
                                        _ = v74
                                        _ = v75
                                    default:
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    }
                                    _ = fi72
                                }
                                if true {
                                    v76 := GenS3{F0: "hello", F1: -253, F2: 255}
                                    _ = v76
                                } else {
                                    fmt.Println(v1)
                                    if (false && false) {
                                        var v77 uint8 = 128
                                        _ = v77
                                    } else {
                                        var v78 uint64 = 4294967295
                                        if (!true) {
                                            var v79 int8 = 127
                                            v28 -= (fi27 >> ((1) & 63))
                                            _ = v79
                                        }
                                        _ = v78
                                    }
                                    for fi80 := 0; fi80 < 2; fi80++ {
                                        v67 = v1
                                        fmt.Println(fi80, v4)
                                        v81 := GenS9{F0: 298}
                                        _ = fi80
                                        _ = v81
                                    }
                                }
                            }
                        }
                        switch "" {
                        case "hello":
                            for fi82 := 0; fi82 < 7; fi82++ {
                                v83 := GenS3{F0: "foo bar", F1: (^(v28)), F2: 34}
                                if (!(false && ((1.5 * 1e10) >= ((0.1 + ((0.0 / float64(-2.25)) + (0.0 + (0.1 / 0.1)))) / float64((-2.25 * ((0.0 / (-2.25 * (0.0 - 0.1))) + v3))))))) {
                                    if (true && (1e10 <= ((0.1 / (0.0 - (float64(float32((float64(v3) / 1e10))) + -2.25))) - float64(v3)))) {
                                        if len(v2) > 0 {
                                            v2[0] = v67
                                        }
                                        var v84 bool = false
                                        var v85 uint8 = 42
                                        if len(v2) > 0 {
                                            v2[0] = (v67 ^ 65535)
                                        }
                                        _ = v84
                                        _ = v85
                                    }
                                }
                                var v86 float64 = (-2.25 * -2.25)
                                _ = fi82
                                _ = v83
                                _ = v86
                            }
                        case "\n\t":
                            var v87 int8 = 42
                            _ = v87
                        case "foo bar":
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                        }
                        _ = v67
                    } else {
                        fmt.Println(fi27, v28)
                        var v88 uint16 = 65535
                        _ = v88
                    }
                }
                switch 0 {
                case 112:
                    v3 = (((((1.5 / (v3 * (v3 * v3))) / v3) + (v3 / float32(1.5))) + (v3 * -2.25)) - ((float32(1.5) - 3.14159) / v3))
                default:
                    _ = len(v2)
                }
            }
            v89 := GenS10{F0: -1, F1: v1, F2: 1}
            _ = fi27
            _ = v28
            _ = v89
        }
        if (v3 == -2.25) {
            v1 -= v1
        }
        var v90 uint32 = (v1 - v1)
        _ = v90
    }
    fmt.Println(v1, v3, v3)
    v4 = v4
    if (-1 < 2147483647) {
        v91 := GenS9{F0: 42}
        fmt.Println(v3, v1)
        _ = v91
    }
    if false {
        fmt.Println(v3, v1)
        var v92 int32 = 0
        _ = v92
    } else {
        for fi93 := 0; fi93 < 7; fi93++ {
            v1 = 1
            if false {
                var v94 int16 = v4
                _ = len(v2)
                _ = len(v2)
                _ = v94
            } else {
                var v95 float32 = v3
                if false {
                    v96 := GenS9{F0: 1}
                    var v97 float64 = (0.1 + float64(v95))
                    _ = v96
                    _ = v97
                }
                _ = v95
            }
            _ = fi93
        }
    }
    v2 = append(v2, 2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
