package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint64; F2 uint16 }
type GenS2 struct { F0 int64; F1 int64; F2 int }
type GenS3 struct { F0 uint64; F1 uint8 }
type GenS4 struct { F0 uint16; F1 string }
type GenS5 struct { F0 uint64; F1 int16 }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 uint8; F1 int64 }
type GenS8 struct { F0 int32; F1 uint64 }
type GenS9 struct { F0 uint64; F1 string; F2 int32 }
type GenS10 struct { F0 int64 }
func genFn0(p0 uint8, p1 int8) uint8 {
    if (!true) {
        fmt.Println(p1, p0)
        p1 = 100
        var v7 string = "hello"
        _ = v7
    }
    p0 -= p0
    _ = p0
    _ = p1
    return (p0 ^ (p0 & (^(p0))))
}

func main() {
    v1 := []int{}
    v2 := []string{("" + "foo bar"), "a", ("a" + ("" + "hello"))}
    var v3 string = "\n\t"
    var v4 int16 = 32767
    v5 := make(map[uint32]uint32)
    v6 := make(map[string]int)
    v6["hello"] = 0
    for fi8 := 0; fi8 < 15; fi8++ {
        var v9 int16 = (v4 - (v4 + 298))
        _ = fi8
        _ = v9
    }
    var v10 string = "a"
    switch 295 {
    case 0:
        var v11 bool = false
        _ = v11
    default:
        v10 = v10
    }
    switch 1 {
    case 128:
        v10 = v10
    case 127:
        if (1 > 2) {
            v12 := GenS1{F0: "\n\t", F1: 1, F2: 32767}
            _ = v12
        }
    }
    v13 := GenS2{F0: 9223372036854775807, F1: 1, F2: 0}
    if true {
        for fi14 := 0; fi14 < 8; fi14++ {
            v6[((((v10 + "foo bar") + (v3 + v10)) + (v10 + v3)) + "a")] = fi14
            v15 := v6[v10]
            _ = fi14
            _ = v15
        }
        v3 = v10
        v16 := GenS3{F0: 0, F1: 56}
        _ = v16
    } else {
        switch 32768 {
        case -2147483648:
            for fi17 := 0; fi17 < 9; fi17++ {
                if (!false) {
                    if true {
                        if ((((!false) && (127 == 100)) && (true || true)) && (v3 >= ("\n\t" + v3))) {
                            var v18 float32 = (-2.25 * float32(-2.25))
                            v3 = (("" + v3) + v3)
                            v1 = append(v1, -317)
                            _ = v18
                        }
                        if true {
                            switch (fi17 & fi17) {
                            case 32767:
                                fmt.Println(v3, v3, v10)
                            case -1:
                                v10 = v3
                            default:
                                for fi19 := 0; fi19 < 1; fi19++ {
                                    if (42 > 0) {
                                        v20 := GenS1{F0: "foo bar", F1: 273, F2: 42}
                                        v5[65535] = 0
                                        v5[127] = 136
                                        var v21 int32 = -2147483648
                                        v22 := GenS4{F0: 65535, F1: (((v10 + "a") + v3) + "")}
                                        _ = v20
                                        _ = v21
                                        _ = v22
                                    } else {
                                        v23 := GenS3{F0: 344, F1: 42}
                                        if false {
                                            switch fi19 {
                                            case -344:
                                                if true {
                                                    v1 = append(v1, (-(fi19)))
                                                }
                                            case 283:
                                                v6[(("hello" + "a") + "\n\t")] = (fi19 * (fi19 / ((fi19) | 1)))
                                            case 42:
                                                switch 4294967295 {
                                                case 1:
                                                    if (138 >= 255) {
                                                        var v24 uint8 = 128
                                                        v24 = (^(v24))
                                                        v25 := v5[337]
                                                        _ = v24
                                                        _ = v25
                                                    } else {
                                                        var v26 string = "foo bar"
                                                        var v27 int16 = (^(v4))
                                                        _ = v26
                                                        _ = v27
                                                    }
                                                case 32768:
                                                    v6[v3] = 9223372036854775807
                                                }
                                            default:
                                                fmt.Println(fi17, v10)
                                            }
                                        } else {
                                            var v28 uint32 = 932
                                            v29 := GenS5{F0: 18446744073709551615, F1: -1}
                                            v30 := GenS6{F0: v28}
                                            _ = v28
                                            _ = v29
                                            _ = v30
                                        }
                                        _ = v23
                                    }
                                    switch 127 {
                                    case -603:
                                        switch 127 {
                                        case -9223372036854775808:
                                            v2 = append(v2, "hello")
                                        }
                                    case 32767:
                                        switch fi19 {
                                        case -1:
                                            var v31 uint = 65535
                                            _ = v31
                                        case 9223372036854775807:
                                            var v32 string = ("\n\t" + v10)
                                            _ = v32
                                        case -796:
                                            switch fi17 {
                                            case 4294967295:
                                                var v33 string = v10
                                                _ = v33
                                            case -9223372036854775808:
                                                v34 := GenS2{F0: -31, F1: -415, F2: (fi19 ^ fi19)}
                                                _ = v34
                                            }
                                        default:
                                            if false {
                                                _ = len(v5)
                                                v3 = ((v3 + ((((v10 + "foo bar") + "foo bar") + v10) + (((v3 + v10) + v3) + ("hello" + v10)))) + v10)
                                                v35 := GenS7{F0: genFn0(211, -74), F1: -9223372036854775808}
                                                _ = v35
                                            } else {
                                                v10 = "\n\t"
                                                var v36 int64 = 587
                                                _ = v36
                                            }
                                        }
                                    default:
                                        v10 = "\n\t"
                                    }
                                    if (0 <= 42) {
                                        v37 := GenS4{F0: 32767, F1: (((v3 + v10) + v3) + ((v10 + (v10 + v10)) + (v10 + v10)))}
                                        var v38 float32 = ((0.0 + 3.14159) - (0.1 / 1.5))
                                        _ = v37
                                        _ = v38
                                    }
                                    _ = fi19
                                }
                            }
                            var v39 int64 = 510
                            var v40 float32 = 3.14159
                            _ = v39
                            _ = v40
                        }
                    }
                    v10 = v3
                    switch v3 {
                    case "foo bar":
                        var v41 uint32 = 1
                        _ = v41
                    case "hello":
                        switch fi17 {
                        case -2147483648:
                            for fi42 := 0; fi42 < 4; fi42++ {
                                var v43 uint8 = 231
                                if ((v43 * v43) != (v43 | (v43 + v43))) {
                                    fmt.Println(v4)
                                    switch fi17 {
                                    case 879:
                                        switch (fi17 - (fi42 ^ -9223372036854775808)) {
                                        case -950:
                                            var v44 string = (("" + "foo bar") + (v3 + "a"))
                                            _ = v44
                                        default:
                                            fmt.Println(v3, v4, v4)
                                        }
                                    case -381:
                                        v3 = (v3 + "a")
                                        v45 := GenS3{F0: 925, F1: 128}
                                        var v46 string = v3
                                        var v47 uint32 = 691
                                        switch fi17 {
                                        case 363:
                                            if len(v1) > 0 {
                                                v1[0] = (fi42 ^ fi17)
                                            }
                                        case -844:
                                            var v48 string = ("hello" + "")
                                            _ = v48
                                        case 42:
                                            v10 = v46
                                        default:
                                            var v49 uint32 = (^(v47))
                                            _ = v49
                                        }
                                        _ = v45
                                        _ = v46
                                        _ = v47
                                    }
                                    var v50 uint8 = v43
                                    _ = v50
                                }
                                if true {
                                    switch (fi42 << ((1) & 63)) {
                                    case -9223372036854775808:
                                        v51 := GenS5{F0: 845, F1: v4}
                                        _ = v51
                                    case -1:
                                        v52 := v5[0]
                                        _ = v52
                                    case -686:
                                        if len(v1) > 0 {
                                            v1[0] = (-(fi17))
                                        }
                                    default:
                                        var v53 int8 = -49
                                        _ = v53
                                    }
                                    var v54 int16 = v4
                                    v55 := GenS7{F0: (v43 << ((38) & 63)), F1: 1}
                                    _ = v54
                                    _ = v55
                                }
                                v56 := GenS8{F0: -56, F1: 4294967295}
                                var v57 int = (fi42 ^ (^(fi42)))
                                fmt.Println(fi42)
                                if (!(255 < 763)) {
                                    if len(v1) > 0 {
                                        v1[0] = 486
                                    }
                                    v1 = append(v1, (-(v57)))
                                    if len(v1) >= 2 {
                                        v1 = v1[1:len(v1)]
                                    }
                                } else {
                                    var v58 int32 = int32(v43)
                                    if true {
                                        switch v3 {
                                        case "foo bar":
                                            if false {
                                                if (147 <= 2147483647) {
                                                    v58 = (v58 << ((2) & 63))
                                                    v43 = genFn0(genFn0(127, 42), 1)
                                                    var v59 int16 = (^(v4))
                                                    _ = v59
                                                } else {
                                                    if (903 != 4294967295) {
                                                        switch 4294967295 {
                                                        case 9223372036854775807:
                                                            v60 := GenS7{F0: v43, F1: 9223372036854775807}
                                                            v61 := GenS8{F0: v58, F1: 1}
                                                            v3 = (((v3 + v10) + (("a" + v10) + "\n\t")) + v3)
                                                            switch fi42 {
                                                            case 385:
                                                                delete(v5, 366)
                                                            default:
                                                                v57 += (fi17 * (fi42 | fi17))
                                                                v62 := GenS3{F0: 870, F1: (v43 % ((42) | 1))}
                                                                _ = v62
                                                            }
                                                            _ = v60
                                                            _ = v61
                                                        case -9223372036854775808:
                                                            v57 += (fi42 % (((fi42 << ((0) & 63))) | 1))
                                                        case 42:
                                                            var v63 uint32 = 2147483648
                                                            _ = v63
                                                        }
                                                        v3 = ("" + v10)
                                                        var v64 int16 = v4
                                                        _ = v64
                                                    }
                                                    var v65 uint16 = 0
                                                    _ = v65
                                                }
                                                var v66 int64 = 1
                                                _ = v66
                                            }
                                        case "\n\t":
                                            var v67 uint32 = 2
                                            _ = v67
                                        default:
                                            v68 := GenS1{F0: "foo bar", F1: 0, F2: 255}
                                            _ = v68
                                        }
                                        v3 = v10
                                        fmt.Println(v43, v43)
                                    }
                                    v1 = append(v1, (fi17 >> ((255) & 63)))
                                    _ = v58
                                }
                                if false {
                                    var v69 float64 = (3.14159 + (3.14159 - 1.5))
                                    delete(v6, v10)
                                    switch ((v3 + v10) + ("" + v3)) {
                                    case "hello":
                                        _ = len(v6)
                                    case "a":
                                        fmt.Println(fi42, v4)
                                    case "\n\t":
                                        v70 := GenS6{F0: 2147483648}
                                        _ = v70
                                    default:
                                        v71 := GenS9{F0: 417, F1: "\n\t", F2: -312}
                                        _ = v71
                                    }
                                    _ = v69
                                }
                                switch (fi17 & -9223372036854775808) {
                                case 128:
                                    var v72 bool = false
                                    _ = v72
                                case 380:
                                    var v73 bool = false
                                    _ = v73
                                default:
                                    fmt.Println(fi17)
                                }
                                _ = fi42
                                _ = v43
                                _ = v56
                                _ = v57
                            }
                        case 9223372036854775807:
                            v10 = v10
                        case 42:
                            v4 = v4
                        default:
                            v1 = append(v1, 164)
                        }
                    default:
                        switch (fi17 * (fi17 & (fi17 & fi17))) {
                        case 128:
                            if false {
                                switch fi17 {
                                case 0:
                                    var v74 uint32 = 794
                                    _ = v74
                                case 32767:
                                    v10 = ("a" + ("\n\t" + ""))
                                case 288:
                                    for fi75 := 0; fi75 < 1; fi75++ {
                                        fmt.Println(v4, fi17, fi75)
                                        var v76 int8 = -79
                                        if (!true) {
                                            v4 += v4
                                            _ = len(v6)
                                            v4 -= (v4 ^ v4)
                                        }
                                        _ = fi75
                                        _ = v76
                                    }
                                default:
                                    v77 := GenS2{F0: 0, F1: 0, F2: fi17}
                                    _ = v77
                                }
                                v4 += 464
                            }
                        case -2147483648:
                            var v78 int8 = 127
                            _ = v78
                        }
                    }
                } else {
                    if true {
                        v10 = v10
                    } else {
                        fmt.Println(v4, v3, v10)
                        for fi79 := 0; fi79 < 9; fi79++ {
                            v4 = (v4 - -517)
                            var v80 int16 = v4
                            _ = fi79
                            _ = v80
                        }
                    }
                    if true {
                        if (-2.25 < float64(1.5)) {
                            if (0 >= -128) {
                                switch 2147483647 {
                                case 4294967295:
                                    switch (v10 + (("" + "hello") + "")) {
                                    case "foo bar":
                                        var v81 int32 = -747
                                        _ = v81
                                    }
                                case 2147483647:
                                    v82 := GenS2{F0: 42, F1: 179, F2: 42}
                                    _ = v82
                                }
                                var v83 uint32 = 818
                                _ = v83
                            } else {
                                fmt.Println(fi17)
                                _ = len(v5)
                            }
                            v10 = ""
                        }
                    } else {
                        var v84 string = "a"
                        _ = v84
                    }
                    var v85 uint32 = 0
                    _ = v85
                }
                switch ("\n\t" + v3) {
                case "foo bar":
                    v86 := GenS10{F0: 1}
                    _ = v86
                case "hello":
                    if true {
                        switch v10 {
                        case "foo bar":
                            for fi87 := 0; fi87 < 6; fi87++ {
                                var v88 uint32 = 2147483647
                                _ = fi87
                                _ = v88
                            }
                        }
                    }
                default:
                    v10 = (v3 + "\n\t")
                }
                _ = fi17
            }
        case 9223372036854775807:
            if len(v2) > 0 {
                v2[0] = (v3 + v3)
            }
        case 169:
            fmt.Println(v10)
        }
    }
    v10 = v3
    v4 -= (-(v4))
    v4 = v4
    if len(v2) > 0 {
        v2[0] = "foo bar"
    }
    if ((^(v4)) <= (^(v4))) {
        var v89 int8 = 106
        if len(v1) > 0 {
            v1[0] = 128
        }
        _ = v89
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
