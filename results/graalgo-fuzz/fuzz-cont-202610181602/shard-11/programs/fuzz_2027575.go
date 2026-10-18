package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
type GenS2 struct { F0 uint; F1 int }
type GenS3 struct { F0 uint; F1 uint16 }
type GenS4 struct { F0 uint8; F1 int64 }
type GenS5 struct { F0 int64; F1 int; F2 uint32 }
type GenS6 struct { F0 int32; F1 uint16; F2 uint32 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int64; F1 int; F2 uint }
type GenS9 struct { F0 int64; F1 uint64; F2 uint8 }
type GenS10 struct { F0 uint16; F1 bool; F2 int64 }
type GenS11 struct { F0 uint }
type GenS12 struct { F0 int; F1 uint64; F2 uint16 }

func main() {
    v1 := make(map[string]uint32)
    var v2 int8 = 0
    var v3 bool = true
    var v4 int64 = 9223372036854775807
    for fi5 := 0; fi5 < 4; fi5++ {
        if v3 {
            v2 += v2
            var v6 float32 = 0.1
            switch fi5 {
            case 629:
                v7 := GenS1{F0: 32767, F1: 42}
                _ = v7
            case 32768:
                v8 := GenS1{F0: 510, F1: 18446744073709551615}
                _ = v8
            case 128:
                if true {
                    switch fi5 {
                    case -9223372036854775808:
                        switch "" {
                        case "\n\t":
                            fmt.Println(fi5)
                        default:
                            if v3 {
                                if v3 {
                                    for fi9 := 0; fi9 < 9; fi9++ {
                                        if (!(((v4 != (v4 + v4)) || (255 >= 4294967295)) || (int32(v4) >= 2147483647))) {
                                            switch ("foo bar" + "hello") {
                                            case "a":
                                                v10 := GenS2{F0: 18446744073709551615, F1: fi9}
                                                _ = v10
                                            case "":
                                                v6 += v6
                                            default:
                                                var v11 uint8 = 42
                                                _ = v11
                                            }
                                        }
                                        if (!v3) {
                                            if ((-(v4)) == (^(v4))) {
                                                delete(v1, "hello")
                                            } else {
                                                if (v2 == -1) {
                                                    if (1 >= -973) {
                                                        v3 = v3
                                                        v12 := GenS2{F0: 102, F1: 42}
                                                        _ = v12
                                                    } else {
                                                        v13 := GenS3{F0: 255, F1: 42}
                                                        var v14 uint64 = 1
                                                        v4 = v4
                                                        _ = v13
                                                        _ = v14
                                                    }
                                                    switch ((("a" + "") + "") + "\n\t") {
                                                    case "hello":
                                                        var v15 uint8 = 255
                                                        _ = v15
                                                    case "\n\t":
                                                        _ = len(v1)
                                                    }
                                                    v16 := GenS4{F0: 128, F1: 9223372036854775807}
                                                    _ = v16
                                                }
                                            }
                                            switch "hello" {
                                            case "a":
                                                var v17 string = ("a" + ("foo bar" + "foo bar"))
                                                switch "foo bar" {
                                                case "a":
                                                    fmt.Println(v17, fi5)
                                                case "\n\t":
                                                    v18 := GenS1{F0: 157, F1: 2}
                                                    _ = v18
                                                }
                                                var v19 uint = 255
                                                var v20 float32 = v6
                                                switch v17 {
                                                case "":
                                                    delete(v1, "a")
                                                case "a":
                                                    v1[""] = 829
                                                }
                                                _ = v17
                                                _ = v19
                                                _ = v20
                                            case "":
                                                var v21 float32 = (float32(-2.25) * (((((v6 / float32(1.5)) - float32((1.5 / 3.14159))) * v6) - (v6 / v6)) / v6))
                                                _ = v21
                                            }
                                            v6 = v6
                                        } else {
                                            var v22 int = (fi9 & (fi5 - (fi5 * fi5)))
                                            var v23 uint16 = 65535
                                            v3 = v3
                                            _ = v22
                                            _ = v23
                                        }
                                        _ = fi9
                                    }
                                    v3 = v3
                                }
                            }
                        }
                    case 4294967295:
                        var v24 int32 = 1
                        _ = v24
                    }
                    v25 := v1[(("a" + "foo bar") + ("hello" + "hello"))]
                    _ = v25
                } else {
                    v4 += v4
                    v26 := GenS4{F0: 127, F1: (v4 << ((695) & 63))}
                    for fi27 := 0; fi27 < 8; fi27++ {
                        var v28 string = ("hello" + "")
                        v2 += 1
                        if v3 {
                            switch fi5 {
                            case 4294967295:
                                v2 = (v2 & v2)
                            default:
                                switch "" {
                                case "\n\t":
                                    delete(v1, v28)
                                case "hello":
                                    var v29 uint32 = 4294967295
                                    _ = v29
                                }
                            }
                            v3 = v3
                            var v30 string = ("a" + ((v28 + "") + v28))
                            _ = v30
                        } else {
                            var v31 bool = (0 >= 65535)
                            fmt.Println(v28, fi27, fi5)
                            if (0 > 98) {
                                if false {
                                    v4 = (v4 ^ (^(v4)))
                                    v32 := GenS4{F0: 255, F1: (v4 / ((v4) | 1))}
                                    v6 = (v6 * v6)
                                    v33 := v1[(v28 + ("\n\t" + v28))]
                                    if v31 {
                                        var v34 uint8 = 127
                                        v2 += 22
                                        if true {
                                            fmt.Println(v31, v34)
                                        }
                                        _ = v34
                                    } else {
                                        switch "hello" {
                                        case "a":
                                            v33 += 439
                                            switch (fi5 >> ((18446744073709551615) & 63)) {
                                            case 407:
                                                if (!true) {
                                                    if (336 >= 328) {
                                                        if (453 >= 0) {
                                                            fmt.Println(v2, fi27, v2)
                                                            if v31 {
                                                                v35 := v1[(v28 + v28)]
                                                                v35 += 2
                                                                _ = v35
                                                            }
                                                        }
                                                    }
                                                }
                                            default:
                                                var v36 uint32 = 4294967295
                                                _ = v36
                                            }
                                            fmt.Println(v2)
                                        case "\n\t":
                                            v4 = v4
                                        case "foo bar":
                                            v37 := GenS5{F0: v4, F1: (fi27 | fi27), F2: v33}
                                            _ = v37
                                        default:
                                            v38 := GenS6{F0: -898, F1: 0, F2: (v33 + (v33 | v33))}
                                            _ = v38
                                        }
                                        var v39 int64 = (v4 ^ v4)
                                        var v40 int16 = -1
                                        switch (v28 + v28) {
                                        case "\n\t":
                                            fmt.Println(v28, v33)
                                        case "":
                                            var v41 bool = v3
                                            if (!v3) {
                                                v31 = (((v33 << ((18446744073709551615) & 63)) != v33) || true)
                                            } else {
                                                switch (fi27 % (((fi5 * (fi27 + (fi5 & fi27)))) | 1)) {
                                                case 434:
                                                    var v42 string = v28
                                                    _ = v42
                                                }
                                            }
                                            _ = len(v1)
                                            _ = v41
                                        case "hello":
                                            v4 -= v4
                                        default:
                                            v39 -= (v39 & v39)
                                        }
                                        _ = v39
                                        _ = v40
                                    }
                                    _ = v32
                                    _ = v33
                                }
                                var v43 int64 = 0
                                _ = v43
                            } else {
                                var v44 int32 = 0
                                v44 = v44
                                v28 = v28
                                var v45 uint8 = 248
                                v46 := []int{int(v45)}
                                v31 = (3.14159 == v6)
                                _ = v44
                                _ = v45
                                _ = v46
                            }
                            switch v28 {
                            case "foo bar":
                                var v47 uint32 = 0
                                _ = v47
                            case "\n\t":
                                v48 := GenS6{F0: 105, F1: 320, F2: 65535}
                                _ = v48
                            case "":
                                var v49 bool = (v31 && (1e10 >= 0.1))
                                switch (fi27 % ((9223372036854775807) | 1)) {
                                case 38:
                                    var v50 uint = 65535
                                    _ = v50
                                case 128:
                                    v51 := GenS4{F0: 127, F1: (v4 << ((679) & 63))}
                                    var v52 uint8 = 42
                                    switch fi27 {
                                    case 619:
                                        v53 := GenS2{F0: 1, F1: fi27}
                                        _ = v53
                                    case -9223372036854775808:
                                        v31 = (538 != 255)
                                    case -1:
                                        _ = len(v1)
                                    default:
                                        var v54 float32 = (1e10 / 1e10)
                                        if (!v31) {
                                            switch (fi27 % ((fi5) | 1)) {
                                            case 32767:
                                                switch (fi5 + (fi5 ^ 127)) {
                                                case -9223372036854775808:
                                                    v1["foo bar"] = 1
                                                case -2147483648:
                                                    if (false || v31) {
                                                        var v55 int32 = 853
                                                        switch fi5 {
                                                        case 32767:
                                                            v55 -= int32(v4)
                                                        case -996:
                                                            if (v4 <= -9223372036854775808) {
                                                                v2 -= (v2 / ((72) | 1))
                                                            } else {
                                                                v56 := GenS7{F0: 768}
                                                                _ = v56
                                                            }
                                                            v6 = v54
                                                        default:
                                                            v28 = v28
                                                        }
                                                        v57 := GenS8{F0: (v4 * 0), F1: (fi5 + -2147483648), F2: 613}
                                                        v3 = (-2147483648 != v55)
                                                        _ = v55
                                                        _ = v57
                                                    }
                                                case 725:
                                                    if v49 {
                                                        v4 = v4
                                                    } else {
                                                        if v49 {
                                                            v58 := v1[v28]
                                                            _ = v58
                                                        } else {
                                                            v54 -= float32((0.0 - ((1.5 * ((3.14159 + 1e10) - float64((0.0 * (1e10 * ((0.0 - v54) + float32((0.0 / 1e10)))))))) - float64((v6 - ((-2.25 / (v6 - 3.14159)) / float32(0.1)))))))
                                                            v6 = v54
                                                        }
                                                    }
                                                }
                                            }
                                        }
                                        if true {
                                            v54 -= (v6 / ((v6 - (float32(0.0) / v6)) * 3.14159))
                                            v59 := GenS9{F0: v4, F1: 2, F2: v52}
                                            _ = len(v1)
                                            var v60 uint8 = (v52 - (^(v52)))
                                            _ = v59
                                            _ = v60
                                        }
                                        _ = v54
                                    }
                                    fmt.Println(v6, v49, v52)
                                    var v61 int8 = v2
                                    var v62 string = v28
                                    _ = v51
                                    _ = v52
                                    _ = v61
                                    _ = v62
                                case 0:
                                    var v63 uint64 = 42
                                    _ = v63
                                default:
                                    v31 = ((0.1 != 1e10) && ((("a" + "\n\t") + (v28 + "foo bar")) != v28))
                                }
                                if v49 {
                                    var v64 float32 = v6
                                    var v65 uint16 = 42
                                    _ = v64
                                    _ = v65
                                }
                                _ = v49
                            default:
                                _ = len(v1)
                            }
                            if v31 {
                                v66 := []int{fi27, fi5, (fi27 & fi27), (fi27 >> ((1) & 63))}
                                v3 = true
                                v67 := GenS10{F0: 255, F1: v3, F2: (^(v4))}
                                _ = v66
                                _ = v67
                            } else {
                                var v68 int16 = 1
                                var v69 float32 = v6
                                v28 = v28
                                _ = v68
                                _ = v69
                            }
                            _ = v31
                        }
                        var v70 int = 1
                        var v71 int64 = v4
                        if false {
                            v71 += (v71 / (((v4 | v4)) | 1))
                        } else {
                            v70 = (^(fi5))
                        }
                        _ = fi27
                        _ = v28
                        _ = v70
                        _ = v71
                    }
                    _ = v26
                }
            }
            _ = v6
        }
        _ = fi5
    }
    switch "a" {
    case "a":
        switch "a" {
        case "hello":
            if (!false) {
                v3 = (!true)
            }
        case "":
            v72 := []int{4294967295}
            _ = v72
        default:
            switch 32767 {
            case 32767:
                v73 := v1[""]
                _ = v73
            default:
                v74 := GenS11{F0: 18446744073709551615}
                _ = v74
            }
        }
    case "":
        _ = len(v1)
    case "\n\t":
        v75 := []int{32768}
        _ = v75
    }
    for fi76 := 0; fi76 < 3; fi76++ {
        v77 := []int{0}
        _ = fi76
        _ = v77
    }
    for fi78 := 0; fi78 < 2; fi78++ {
        v3 = (1 <= 2)
        v2 = (^(v2))
        var v79 int8 = 1
        _ = fi78
        _ = v79
    }
    if (!v3) {
        v2 = v2
        for fi80 := 0; fi80 < 7; fi80++ {
            var v81 int16 = -32768
            _ = fi80
            _ = v81
        }
    }
    v82 := GenS12{F0: 1, F1: 4294967295, F2: 32767}
    var v83 uint8 = 42
    var v84 float32 = 1.5
    v85 := GenS10{F0: 690, F1: (!(v2 == v2)), F2: v4}
    v4 = (-(v4))
    delete(v1, ("" + "a"))
    if v3 {
        var v86 string = "\n\t"
        v2 += 42
        _ = v86
    } else {
        fmt.Println(v4, v3, v4)
        v87 := []int{127, -692, 127, 9223372036854775807}
        var v88 string = ("hello" + ("a" + "a"))
        _ = v87
        _ = v88
    }
    var v89 uint8 = (v83 - 255)
    if false {
        for fi90 := 0; fi90 < 5; fi90++ {
            var v91 uint32 = 0
            var v92 uint16 = 1
            _ = fi90
            _ = v91
            _ = v92
        }
    } else {
        var v93 int16 = -32768
        for fi94 := 0; fi94 < 16; fi94++ {
            var v95 bool = v3
            _ = fi94
            _ = v95
        }
        _ = len(v1)
        _ = v93
    }
    v89 = (v83 << ((931) & 63))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v83:", v83)
    fmt.Println("v84:", v84)
    fmt.Println("v89:", v89)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v82
    _ = v83
    _ = v84
    _ = v85
    _ = v89
}
