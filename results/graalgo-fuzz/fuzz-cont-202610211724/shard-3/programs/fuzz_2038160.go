package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 bool }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 string; F1 string; F2 uint32 }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 int64; F1 uint64; F2 uint16 }
type GenS6 struct { F0 bool; F1 int64 }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 int16; F1 int32; F2 int64 }
type GenS9 struct { F0 uint32; F1 int16; F2 int32 }
type GenS10 struct { F0 uint8; F1 uint }
type GenS11 struct { F0 bool; F1 uint8 }
type GenS12 struct { F0 int; F1 uint64 }
type GenS13 struct { F0 bool; F1 int; F2 uint16 }

func main() {
    v1 := []int{-9223372036854775808, 32768, 2147483647, 32768}
    var v2 int64 = 42
    v3 := make(map[string]uint32)
    v4 := make(map[string]bool)
    v5 := []uint32{733, 2147483648, 1, 0}
    v6 := make(map[uint32]int)
    v7 := GenS1{F0: 255, F1: (!false)}
    v8 := GenS1{F0: 750, F1: (true || false)}
    switch 42 {
    case 4294967295:
        v2 += 9223372036854775807
    default:
        if true {
            switch -2147483648 {
            case 9223372036854775807:
                switch 844 {
                case 460:
                    v9 := GenS1{F0: 259, F1: false}
                    _ = v9
                }
            }
            fmt.Println(v2)
            var v10 bool = (3.14159 == 3.14159)
            _ = v10
        }
    }
    for fi11 := 0; fi11 < 8; fi11++ {
        for fi12 := 0; fi12 < 3; fi12++ {
            v13 := GenS1{F0: 2, F1: false}
            fmt.Println(v2, fi11, fi12)
            _ = fi12
            _ = v13
        }
        for fi14 := 0; fi14 < 9; fi14++ {
            v2 += (v2 % ((-1) | 1))
            var v15 bool = true
            if (!v15) {
                fmt.Println(v15, v15, fi11)
                if (32767 == 32767) {
                    switch (fi11 | fi14) {
                    case 299:
                        switch "a" {
                        case "a":
                            if v15 {
                                var v16 float64 = 1.5
                                v16 = 0.1
                                switch "hello" {
                                case "a":
                                    switch (fi14 | -9223372036854775808) {
                                    case 732:
                                        if false {
                                            var v17 bool = false
                                            _ = v17
                                        }
                                    case 2147483647:
                                        v2 += (^(v2))
                                    case -9223372036854775808:
                                        var v18 uint64 = 198
                                        _ = v18
                                    }
                                default:
                                    v19 := GenS1{F0: 883, F1: v15}
                                    _ = v19
                                }
                                _ = v16
                            }
                        default:
                            var v20 uint64 = 0
                            _ = v20
                        }
                    case 32767:
                        v2 += (-(v2))
                    case 2147483647:
                        v1 = append(v1, 975)
                    }
                    v2 = v2
                } else {
                    v21 := GenS1{F0: 77, F1: (94 < 292)}
                    var v22 uint64 = 18446744073709551615
                    v1 = append(v1, -266)
                    fmt.Println(v15, fi11, fi11)
                    var v23 float32 = (((((0.0 / 1e10) * 0.1) * 1e10) + float32(0.1)) + ((0.1 * (0.0 * (((1e10 * 3.14159) - (0.0 + float32((float64((1.5 * 0.1)) + -2.25)))) - (0.1 / 1.0)))) / -2.25))
                    _ = v21
                    _ = v22
                    _ = v23
                }
                var v24 uint8 = 128
                _ = v24
            } else {
                v2 += v2
                if v15 {
                    if v15 {
                        var v25 bool = ("hello" != "a")
                        v15 = v15
                        v26 := GenS2{F0: (-1 != 97)}
                        v25 = true
                        _ = v25
                        _ = v26
                    } else {
                        switch "foo bar" {
                        case "a":
                            if (1 > 0) {
                                var v27 bool = false
                                _ = v27
                            }
                        case "foo bar":
                            v28 := GenS3{F0: (("" + ("foo bar" + "a")) + ""), F1: ("" + "hello"), F2: 1}
                            _ = v28
                        case "hello":
                            v2 -= -9223372036854775808
                        default:
                            switch "\n\t" {
                            case "\n\t":
                                if (-1 < -139) {
                                    v29 := GenS3{F0: ("a" + "\n\t"), F1: (("hello" + "a") + (("\n\t" + ("foo bar" + ("foo bar" + "a"))) + "\n\t")), F2: 42}
                                    switch (-(fi14)) {
                                    case 95:
                                        switch fi14 {
                                        case 4294967295:
                                            if (2 <= 153) {
                                                v2 += (v2 - (v2 & v2))
                                                var v30 uint32 = 340
                                                _ = v30
                                            }
                                        default:
                                            var v31 int = fi14
                                            _ = v31
                                        }
                                    default:
                                        switch "a" {
                                        case "hello":
                                            v2 += 1
                                        case "foo bar":
                                            fmt.Println(fi14)
                                            fmt.Println(v2, fi11)
                                            v2 += v2
                                            fmt.Println(fi11, fi14, v15)
                                            v2 = (v2 % (((v2 - v2)) | 1))
                                            switch fi14 {
                                            case -9223372036854775808:
                                                var v32 uint64 = 557
                                                switch (("foo bar" + "\n\t") + "hello") {
                                                case "a":
                                                    switch 2147483647 {
                                                    case 9223372036854775807:
                                                        var v33 uint64 = v32
                                                        _ = v33
                                                    case -796:
                                                        var v34 uint32 = 65535
                                                        _ = v34
                                                    default:
                                                        if v15 {
                                                            v35 := GenS3{F0: "a", F1: (("" + "hello") + "\n\t"), F2: 803}
                                                            if (!(v15 && v15)) {
                                                                if (v15 || (!(v15 && (v32 != v32)))) {
                                                                    var v36 bool = (v15 || true)
                                                                    _ = v36
                                                                }
                                                                v2 += (v2 - (v2 / ((1) | 1)))
                                                                var v37 uint16 = 2
                                                                _ = v37
                                                            }
                                                            if (false || (128 < 208)) {
                                                                v1 = append(v1, -9223372036854775808)
                                                                v38 := GenS4{F0: 65535}
                                                                v39 := GenS1{F0: 2, F1: true}
                                                                _ = v38
                                                                _ = v39
                                                            } else {
                                                                v40 := GenS5{F0: v2, F1: 1, F2: 1}
                                                                _ = v40
                                                            }
                                                            var v41 string = (("\n\t" + "hello") + ("\n\t" + "\n\t"))
                                                            _ = v35
                                                            _ = v41
                                                        } else {
                                                            var v42 int64 = v2
                                                            _ = len(v5)
                                                            _ = v42
                                                        }
                                                        switch fi11 {
                                                        case 32768:
                                                            fmt.Println(v32, v15)
                                                        default:
                                                            v43 := GenS6{F0: v15, F1: 925}
                                                            _ = v43
                                                        }
                                                        switch fi14 {
                                                        case -785:
                                                            switch (^(fi14)) {
                                                            case -133:
                                                                if v15 {
                                                                    delete(v6, 181)
                                                                } else {
                                                                    switch "hello" {
                                                                    case "hello":
                                                                        fmt.Println(v32, fi11, v15)
                                                                    case "":
                                                                        if ((18446744073709551615 > (^(v32))) || v15) {
                                                                            switch "" {
                                                                            case "":
                                                                                v44 := GenS7{F0: ((v2 << ((18446744073709551615) & 63)) > (v2 ^ v2))}
                                                                                _ = v44
                                                                            case "a":
                                                                                v45 := v6[42]
                                                                                _ = v45
                                                                            }
                                                                            fmt.Println(fi11, fi14)
                                                                            v46 := GenS8{F0: 1000, F1: -2147483648, F2: v2}
                                                                            _ = v46
                                                                        } else {
                                                                            if (!true) {
                                                                                var v47 int16 = -32768
                                                                                _ = v47
                                                                            } else {
                                                                                if (false || (!true)) {
                                                                                    if ((v2 >> ((2) & 63)) >= v2) {
                                                                                        v48 := GenS6{F0: (!false), F1: (v2 >> ((65535) & 63))}
                                                                                        var v49 bool = (3.14159 != (3.14159 - ((1.5 + (1e10 * (3.14159 / 3.14159))) / -2.25)))
                                                                                        _ = v48
                                                                                        _ = v49
                                                                                    }
                                                                                    if (false && false) {
                                                                                        v15 = false
                                                                                        _ = len(v3)
                                                                                        v15 = ((0.1 - 1e10) == 0.0)
                                                                                    }
                                                                                }
                                                                                var v50 int32 = 42
                                                                                var v51 int = fi11
                                                                                _ = v50
                                                                                _ = v51
                                                                            }
                                                                            fmt.Println(v15, v15, v2)
                                                                        }
                                                                    case "\n\t":
                                                                        if (v15 || v15) {
                                                                            delete(v4, ("foo bar" + ("" + ((("\n\t" + "\n\t") + "hello") + (("hello" + "foo bar") + "hello")))))
                                                                            if (255 > 128) {
                                                                                var v52 float64 = (0.0 / (1.5 / 1.0))
                                                                                v32 = 880
                                                                                _ = v52
                                                                            }
                                                                        } else {
                                                                            var v53 int16 = -32768
                                                                            fmt.Println(v32, fi11, fi14)
                                                                            v53 = v53
                                                                            if (803 > 2) {
                                                                                var v54 int32 = 42
                                                                                _ = v54
                                                                            } else {
                                                                                var v55 bool = true
                                                                                _ = v55
                                                                            }
                                                                            _ = v53
                                                                        }
                                                                    default:
                                                                        if (32767 >= -32768) {
                                                                            var v56 int = -9223372036854775808
                                                                            _ = v56
                                                                        } else {
                                                                            v1 = append(v1, (^(fi14)))
                                                                        }
                                                                    }
                                                                    _ = len(v4)
                                                                    v57 := v3["foo bar"]
                                                                    _ = v57
                                                                }
                                                            case 9223372036854775807:
                                                                v58 := GenS9{F0: 42, F1: -619, F2: 2147483647}
                                                                var v59 string = "a"
                                                                var v60 string = v59
                                                                _ = v58
                                                                _ = v59
                                                                _ = v60
                                                            case 32767:
                                                                v61 := GenS7{F0: (613 == 1)}
                                                                _ = v61
                                                            default:
                                                                var v62 float32 = (1.5 - 3.14159)
                                                                _ = v62
                                                            }
                                                        case 128:
                                                            v2 += (v2 & v2)
                                                        }
                                                    }
                                                }
                                                v15 = (-9223372036854775808 < fi11)
                                                if (!((100 >= 127) || v15)) {
                                                    fmt.Println(fi14)
                                                    v1 = append(v1, (fi11 >> ((18446744073709551615) & 63)))
                                                    if ((63 < 1) || true) {
                                                        var v63 uint64 = 238
                                                        if len(v1) > 0 {
                                                            v1[0] = (fi11 % (((fi14 & fi11)) | 1))
                                                        }
                                                        v2 -= -948
                                                        _ = v63
                                                    }
                                                }
                                                v64 := GenS2{F0: false}
                                                _ = v32
                                                _ = v64
                                            default:
                                                if v15 {
                                                    v15 = (!v15)
                                                } else {
                                                    var v65 float64 = (0.0 - (0.1 * ((((float64((float32(float64(float32(1.5))) + 0.0)) - 1e10) + 0.1) / ((float64((1.5 - (1e10 / (((3.14159 - -2.25) / 3.14159) + (0.0 - (((float32((0.1 + (float64(0.0) - 0.1))) / 1e10) - (1e10 - 0.0)) - (-2.25 + ((0.0 / 1.5) - (0.1 / 1.0))))))))) - 0.0) / 1e10)) * 1e10)))
                                                    v66 := GenS1{F0: 0, F1: false}
                                                    switch "" {
                                                    case "":
                                                        if (1 >= 128) {
                                                            delete(v6, 0)
                                                        }
                                                        var v67 int8 = -1
                                                        v2 = v2
                                                        _ = v67
                                                    default:
                                                        v15 = (!(-79 >= 0))
                                                    }
                                                    _ = v65
                                                    _ = v66
                                                }
                                            }
                                        }
                                    }
                                    v2 = (v2 + v2)
                                    _ = v29
                                } else {
                                    var v68 int64 = v2
                                    v2 = (v68 << ((65535) & 63))
                                    fmt.Println(v68)
                                    _ = v68
                                }
                            default:
                                v69 := v3[""]
                                _ = v69
                            }
                        }
                        switch "a" {
                        case "":
                            fmt.Println(fi11)
                        case "hello":
                            switch ("\n\t" + ("\n\t" + "hello")) {
                            case "":
                                if (0 < 255) {
                                    v5 = append(v5, 2147483648)
                                } else {
                                    delete(v4, ("foo bar" + ("a" + ("a" + "\n\t"))))
                                }
                            case "hello":
                                v70 := GenS10{F0: 42, F1: 2}
                                _ = v70
                            default:
                                v15 = false
                            }
                            v15 = false
                            switch "\n\t" {
                            case "\n\t":
                                v6[65535] = (fi11 * (fi11 ^ fi11))
                            default:
                                if len(v5) > 0 {
                                    v5[0] = 4294967295
                                }
                            }
                        case "a":
                            var v71 uint16 = 2
                            _ = v71
                        default:
                            v15 = v15
                        }
                    }
                    var v72 float32 = (0.0 * 0.0)
                    _ = v72
                } else {
                    v15 = (!(!((3.14159 / ((-2.25 / 1.0) - 1.5)) < 1.5)))
                }
            }
            _ = fi14
            _ = v15
        }
        var v73 bool = (!((fi11 - fi11) < fi11))
        _ = fi11
        _ = v73
    }
    v74 := GenS1{F0: 18446744073709551615, F1: false}
    _ = len(v5)
    if len(v1) > 0 {
        v1[0] = 228
    }
    v75 := GenS9{F0: 329, F1: 1, F2: -140}
    for fi76 := 0; fi76 < 4; fi76++ {
        switch fi76 {
        case 4294967295:
            if true {
                if false {
                    switch (fi76 + fi76) {
                    case 4294967295:
                        fmt.Println(v2, fi76)
                    case 127:
                        var v77 uint = 1
                        _ = v77
                    }
                } else {
                    if (false || true) {
                        if (!(int64(fi76) == v2)) {
                            v78 := GenS1{F0: 1, F1: ((240 >= 4294967295) || true)}
                            v79 := GenS11{F0: ((42 > 483) || (!((42 == 1) && (((true || (0.0 >= float64(1e10))) && (2147483647 != 876)) || true)))), F1: 255}
                            _ = v78
                            _ = v79
                        }
                    }
                    for fi80 := 0; fi80 < 2; fi80++ {
                        v81 := GenS12{F0: fi80, F1: 926}
                        v82 := GenS4{F0: 473}
                        if true {
                            v2 -= (v2 >> ((255) & 63))
                            v83 := GenS7{F0: (!((4294967295 != 0) && (65535 > 0)))}
                            fmt.Println(fi76)
                            _ = v83
                        }
                        _ = fi80
                        _ = v81
                        _ = v82
                    }
                    for fi84 := 0; fi84 < 7; fi84++ {
                        fmt.Println(v2, v2, v2)
                        v2 = (v2 - (v2 | (^(v2))))
                        _ = fi84
                    }
                }
                v2 = v2
                v85 := GenS13{F0: false, F1: (fi76 / ((fi76) | 1)), F2: 283}
                _ = v85
            } else {
                fmt.Println(fi76)
                v2 = (v2 * -1)
            }
        case 128:
            fmt.Println(v2, fi76)
        case 2147483647:
            fmt.Println(v2, fi76)
        }
        _ = fi76
    }
    fmt.Println(v2)
    for fi86 := 0; fi86 < 16; fi86++ {
        v2 = 782
        _ = fi86
    }
    var v87 int = 42
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v87:", v87)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v74
    _ = v75
    _ = v8
    _ = v87
}
