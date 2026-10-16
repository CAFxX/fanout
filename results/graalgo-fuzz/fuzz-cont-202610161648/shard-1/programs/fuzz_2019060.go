package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint32; F1 uint32; F2 bool }
type GenS3 struct { F0 int64; F1 uint64; F2 uint64 }
type GenS4 struct { F0 bool; F1 uint8; F2 bool }
type GenS5 struct { F0 bool; F1 int; F2 uint }
type GenS6 struct { F0 int8 }
type GenS7 struct { F0 uint32; F1 uint; F2 int64 }
type GenS8 struct { F0 uint; F1 int32; F2 int }
type GenS9 struct { F0 uint64; F1 string; F2 uint32 }
type GenS10 struct { F0 uint64; F1 bool }
type GenS11 struct { F0 int32; F1 string; F2 uint64 }
type GenS12 struct { F0 string; F1 bool }
func genFn0() uint32 {
    v6 := GenS1{F0: 444}
    v7 := GenS2{F0: 528, F1: 815, F2: ((-30 > 1) || ("hello" >= "a"))}
    for fi8 := 0; fi8 < 5; fi8++ {
        v9 := []int{}
        fmt.Println(fi8)
        _ = fi8
        _ = v9
    }
    v10 := []int{281, 0}
    _ = v10
    _ = v6
    _ = v7
    return 65535
}

func main() {
    var v1 float64 = 1.5
    var v2 uint8 = 1
    var v3 float32 = -2.25
    var v4 float32 = -2.25
    var v5 uint32 = 309
    v11 := GenS1{F0: 42}
    switch 2147483647 {
    case 32767:
        var v12 int64 = 160
        _ = v12
    case 1:
        var v13 int64 = 9223372036854775807
        _ = v13
    default:
        switch 2147483647 {
        case -2147483648:
            for fi14 := 0; fi14 < 19; fi14++ {
                var v15 bool = false
                _ = fi14
                _ = v15
            }
        }
    }
    fmt.Println(v4, v4, v2)
    var v16 uint8 = v2
    var v17 int32 = 42
    var v18 float32 = (0.1 + (0.0 * (v4 / (v4 - v4))))
    v19 := []int{-712, -2147483648, 42, 4294967295}
    v20 := GenS1{F0: 1000}
    var v21 int16 = 888
    if true {
        var v22 uint64 = 62
        var v23 int16 = (v21 - (v21 & (v21 & 0)))
        switch -878 {
        case -525:
            if len(v19) > 0 {
                v19[0] = 2147483647
            }
        case 32768:
            fmt.Println(v18)
        case 9223372036854775807:
            for fi24 := 0; fi24 < 17; fi24++ {
                if (!true) {
                    if len(v19) >= 2 {
                        v19 = v19[1:len(v19)]
                    }
                    v23 += v23
                    v3 = ((v3 / v18) + v4)
                } else {
                    var v25 int = fi24
                    var v26 int16 = v21
                    _ = v25
                    _ = v26
                }
                _ = fi24
            }
        default:
            v27 := GenS2{F0: 42, F1: genFn0(), F2: false}
            _ = v27
        }
        _ = v22
        _ = v23
    }
    var v28 bool = ((v16 & 128) >= v16)
    for fi29 := 0; fi29 < 13; fi29++ {
        v30 := GenS1{F0: (v21 * 770)}
        _ = len(v19)
        _ = fi29
        _ = v30
    }
    if v28 {
        switch ("a" + "hello") {
        case "\n\t":
            fmt.Println(v2, v18)
        }
        for fi31 := 0; fi31 < 8; fi31++ {
            _ = len(v19)
            if len(v19) > 0 {
                v19[0] = (fi31 - fi31)
            }
            for fi32 := 0; fi32 < 14; fi32++ {
                fmt.Println(fi32, fi31, v18)
                var v33 float64 = ((((v1 / v1) / ((v1 * ((1.5 / v1) - ((v1 - v1) / (((-2.25 - v1) - v1) * (v1 * float64(float32(0.1))))))) * 3.14159)) / 0.1) / -2.25)
                if true {
                    switch (fi31 << ((197) & 63)) {
                    case 128:
                        switch fi31 {
                        case -820:
                            switch "" {
                            case "":
                                switch fi32 {
                                case 4294967295:
                                    if (255 >= 846) {
                                        switch fi31 {
                                        case -848:
                                            _ = len(v19)
                                        case 703:
                                            v34 := GenS3{F0: -9223372036854775808, F1: 2, F2: 4294967295}
                                            _ = v34
                                        case 32767:
                                            fmt.Println(v2, v16)
                                        }
                                        var v35 uint16 = 611
                                        _ = v35
                                    }
                                }
                                v21 = v21
                            }
                        }
                    case 127:
                        _ = len(v19)
                    case 253:
                        switch "a" {
                        case "\n\t":
                            if v28 {
                                if len(v19) >= 2 {
                                    v19 = v19[1:len(v19)]
                                }
                                if ((!v28) || true) {
                                    _ = len(v19)
                                    v36 := GenS2{F0: (v5 ^ v5), F1: (v5 << ((255) & 63)), F2: ("" == "hello")}
                                    _ = v36
                                } else {
                                    if len(v19) > 0 {
                                        v19[0] = (fi31 >> ((4294967295) & 63))
                                    }
                                    var v37 int8 = int8(v21)
                                    _ = v37
                                }
                                fmt.Println(fi31, v4)
                                v38 := GenS3{F0: 9223372036854775807, F1: 2, F2: 2}
                                fmt.Println(v21)
                                _ = v38
                            } else {
                                _ = len(v19)
                                switch (fi31 / ((fi31) | 1)) {
                                case 4294967295:
                                    fmt.Println(v1, v21)
                                    fmt.Println(v33)
                                    v39 := GenS3{F0: 9223372036854775807, F1: 2, F2: 153}
                                    v21 = v21
                                    _ = v39
                                case -186:
                                    if len(v19) > 0 {
                                        v19[0] = fi32
                                    }
                                case 9223372036854775807:
                                    switch "hello" {
                                    case "":
                                        if true {
                                            v40 := GenS4{F0: (!v28), F1: (^(v2)), F2: (false && (113 != 42))}
                                            v41 := GenS5{F0: (!(v1 == v33)), F1: (fi32 ^ fi32), F2: 18446744073709551615}
                                            _ = v40
                                            _ = v41
                                        } else {
                                            var v42 int16 = v21
                                            v1 = (float64(v18) + v33)
                                            fmt.Println(v16)
                                            _ = v42
                                        }
                                    case "\n\t":
                                        var v43 bool = v28
                                        _ = v43
                                    default:
                                        if true {
                                            fmt.Println(fi31, v16, v1)
                                        } else {
                                            v44 := GenS6{F0: -128}
                                            _ = v44
                                        }
                                    }
                                }
                                v45 := GenS7{F0: v5, F1: 65535, F2: -398}
                                _ = v45
                            }
                        default:
                            v46 := GenS6{F0: 61}
                            _ = v46
                        }
                        if true {
                            fmt.Println(v2, v21, v3)
                            var v47 uint16 = 255
                            _ = v47
                        } else {
                            var v48 uint32 = 0
                            _ = v48
                        }
                    default:
                        var v49 int = fi32
                        if len(v19) > 0 {
                            v19[0] = 32768
                        }
                        _ = v49
                    }
                    fmt.Println(v5, v2, v33)
                    v4 = ((-2.25 / 3.14159) * float32(((v1 * v1) * v1)))
                    v33 -= (float64(-2.25) / ((v1 + 0.1) * ((float64((((((v3 / 1e10) - 1e10) + 1.5) * v3) - (v4 * ((float32(v33) / (v4 - (v4 * v4))) * (v4 - (float32(v1) * (3.14159 - (3.14159 / v18)))))))) / (v1 / 1.5)) / v1)))
                    var v50 bool = (!(!(v28 || v28)))
                    var v51 float64 = (0.0 + v33)
                    _ = v50
                    _ = v51
                }
                _ = fi32
                _ = v33
            }
            _ = fi31
        }
    } else {
        v19 = append(v19, -9223372036854775808)
    }
    switch 1 {
    case -751:
        for fi52 := 0; fi52 < 16; fi52++ {
            if len(v19) >= 2 {
                v19 = v19[1:len(v19)]
            }
            for fi53 := 0; fi53 < 4; fi53++ {
                fmt.Println(v1, v1)
                switch (^(fi53)) {
                case 128:
                    if v28 {
                        v18 -= v18
                        switch "hello" {
                        case "hello":
                            v4 += v3
                        case "\n\t":
                            var v54 uint64 = 695
                            _ = v54
                        }
                        v55 := GenS8{F0: 0, F1: (v17 ^ int32(v21)), F2: fi53}
                        _ = v55
                    }
                case -569:
                    switch "foo bar" {
                    case "foo bar":
                        if (1 <= -221) {
                            switch "hello" {
                            case "":
                                var v56 int16 = (v21 & v21)
                                _ = v56
                            case "a":
                                v4 -= v4
                            case "hello":
                                var v57 uint16 = 291
                                _ = v57
                            }
                            var v58 uint32 = 981
                            v17 = (v17 / (((v17 - v17)) | 1))
                            _ = v58
                        }
                    case "\n\t":
                        v16 = (v2 | (v2 + 224))
                        v16 = (v16 / ((v16) | 1))
                        v59 := GenS5{F0: true, F1: 0, F2: 1}
                        _ = v59
                    default:
                        fmt.Println(v5, v16)
                    }
                }
                _ = fi53
            }
            if true {
                var v60 float64 = v1
                for fi61 := 0; fi61 < 18; fi61++ {
                    switch (fi52 - (fi52 ^ fi52)) {
                    case -706:
                        if false {
                            _ = len(v19)
                            fmt.Println(v2)
                            switch (("a" + "foo bar") + "\n\t") {
                            case "\n\t":
                                if (65535 != 0) {
                                    var v62 int8 = -2
                                    _ = len(v19)
                                    var v63 int = -732
                                    switch fi52 {
                                    case -76:
                                        v5 = (v5 - (v5 + (v5 % ((v5) | 1))))
                                    case -1:
                                        if (!v28) {
                                            var v64 int32 = v17
                                            v65 := GenS9{F0: 18446744073709551615, F1: "\n\t", F2: v5}
                                            _ = v64
                                            _ = v65
                                        } else {
                                            if (4294967295 < 0) {
                                                var v66 string = "\n\t"
                                                v67 := GenS10{F0: 0, F1: v28}
                                                if len(v19) > 0 {
                                                    v19[0] = (fi61 + v63)
                                                }
                                                _ = v66
                                                _ = v67
                                            }
                                            fmt.Println(v62)
                                            v68 := GenS2{F0: v5, F1: 181, F2: (v28 || true)}
                                            _ = v68
                                        }
                                    case 32768:
                                        v17 = v17
                                    }
                                    _ = v62
                                    _ = v63
                                }
                                var v69 int16 = (v21 + (v21 + v21))
                                _ = v69
                            default:
                                switch ((("hello" + "a") + "\n\t") + ("" + ("hello" + "foo bar"))) {
                                case "a":
                                    var v70 uint32 = 1
                                    _ = v70
                                case "\n\t":
                                    v4 += v4
                                case "foo bar":
                                    if len(v19) >= 2 {
                                        v19 = v19[1:len(v19)]
                                    }
                                default:
                                    fmt.Println(v4)
                                }
                            }
                            switch "foo bar" {
                            case "hello":
                                v71 := GenS11{F0: v17, F1: "a", F2: 42}
                                var v72 uint16 = 42
                                _ = v71
                                _ = v72
                            case "a":
                                if len(v19) > 0 {
                                    v19[0] = fi52
                                }
                            case "foo bar":
                                if v28 {
                                    if (v16 == v16) {
                                        var v73 bool = v28
                                        v2 = v16
                                        v21 = v21
                                        _ = v73
                                    }
                                }
                            default:
                                v74 := GenS12{F0: "\n\t", F1: ((v17 ^ v17) == 256)}
                                _ = v74
                            }
                            if v28 {
                                v75 := GenS4{F0: ((v21 < (v21 ^ (v21 | 911))) && (!((v21 % ((-1) | 1)) == v21))), F1: (v16 + (v2 % (((v2 + (v2 | v16))) | 1))), F2: ((v21 * (v21 / ((int16(v16)) | 1))) < (v21 & (v21 ^ 32767)))}
                                _ = v75
                            }
                        }
                    default:
                        if len(v19) > 0 {
                            v19[0] = (fi61 % ((fi61) | 1))
                        }
                    }
                    _ = len(v19)
                    var v76 int32 = (v17 >> ((42) & 63))
                    _ = fi61
                    _ = v76
                }
                var v77 int16 = (v21 << ((18446744073709551615) & 63))
                _ = v60
                _ = v77
            }
            _ = fi52
        }
    }
    for fi78 := 0; fi78 < 15; fi78++ {
        for fi79 := 0; fi79 < 12; fi79++ {
            v16 += (v2 >> ((59) & 63))
            if v28 {
                var v80 uint = 2
                _ = v80
            }
            _ = fi79
        }
        for fi81 := 0; fi81 < 9; fi81++ {
            v28 = v28
            var v82 bool = ((-(v21)) > (v21 + v21))
            v16 = v2
            _ = fi81
            _ = v82
        }
        _ = len(v19)
        _ = fi78
    }
    if ((true || v28) || v28) {
        v16 = (v2 / ((v2) | 1))
    }
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v19:", len(v19), v19)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
