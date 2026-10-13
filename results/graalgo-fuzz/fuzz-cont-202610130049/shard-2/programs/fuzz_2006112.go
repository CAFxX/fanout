package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint8; F2 int16 }
type GenS2 struct { F0 uint; F1 bool }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 int; F1 int32; F2 uint16 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int16; F1 uint; F2 int16 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 int8; F1 int8 }
func genFn0(p0 int64, p1 int64) uint16 {
    var v5 uint = 0
    var v6 uint64 = 0
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 1
}

func main() {
    var v1 bool = false
    var v2 float64 = (((1e10 + 0.0) - ((-2.25 * (0.0 / -2.25)) / float64(1.5))) + (1.5 * (0.0 * ((1e10 + 1.5) - ((0.0 * float64(0.1)) / 1e10)))))
    v3 := []uint32{0}
    var v4 int64 = 863
    if v1 {
        if len(v3) > 0 {
            v3[0] = 2
        }
        var v7 int16 = 1
        _ = v7
    } else {
        for fi8 := 0; fi8 < 16; fi8++ {
            v9 := GenS1{F0: v4, F1: 2, F2: -1}
            _ = fi8
            _ = v9
        }
        if true {
            if (true && v1) {
                var v10 uint16 = 1
                fmt.Println(v10)
                _ = v10
            } else {
                _ = len(v3)
                v1 = (128 >= 42)
            }
            if (v1 || (-1 == -2147483648)) {
                var v11 uint16 = 32767
                _ = v11
            } else {
                var v12 int = 127
                fmt.Println(v1, v4, v4)
                _ = v12
            }
        }
    }
    fmt.Println(v1, v1)
    var v13 int = -534
    var v14 int32 = 0
    fmt.Println(v2, v1, v2)
    if (42 <= 42) {
        var v15 int = (v13 / ((v13) | 1))
        if v1 {
            for fi16 := 0; fi16 < 18; fi16++ {
                v17 := GenS1{F0: v4, F1: 34, F2: -32768}
                _ = fi16
                _ = v17
            }
            for fi18 := 0; fi18 < 3; fi18++ {
                if false {
                    for fi19 := 0; fi19 < 8; fi19++ {
                        v3 = append(v3, 1)
                        var v20 string = (("foo bar" + "") + ("foo bar" + "a"))
                        _ = fi19
                        _ = v20
                    }
                }
                for fi21 := 0; fi21 < 10; fi21++ {
                    var v22 int16 = -711
                    var v23 int8 = 80
                    _ = fi21
                    _ = v22
                    _ = v23
                }
                var v24 int16 = 32767
                _ = fi18
                _ = v24
            }
        }
        _ = v15
    }
    v25 := GenS1{F0: (^(v4)), F1: 214, F2: 1}
    switch v13 {
    case -9223372036854775808:
        for fi26 := 0; fi26 < 12; fi26++ {
            var v27 uint64 = 93
            _ = fi26
            _ = v27
        }
    case -789:
        v28 := GenS1{F0: v4, F1: 42, F2: -32768}
        _ = v28
    case 2147483647:
        for fi29 := 0; fi29 < 2; fi29++ {
            v4 -= v4
            _ = fi29
        }
    default:
        for fi30 := 0; fi30 < 4; fi30++ {
            switch -987 {
            case 406:
                v1 = (230 < 42)
            }
            _ = fi30
        }
    }
    for fi31 := 0; fi31 < 14; fi31++ {
        var v32 float32 = -2.25
        var v33 int8 = 42
        _ = fi31
        _ = v32
        _ = v33
    }
    var v34 float64 = (v2 * v2)
    v2 = 1.5
    switch ("" + (("foo bar" + (("" + "foo bar") + "foo bar")) + "\n\t")) {
    case "a":
        v34 += (v2 * (v34 + float64(1.5)))
    case "foo bar":
        v2 += float64((float32(((v34 / v34) - 1e10)) * float32((v2 + float64(float32(1e10))))))
    case "":
        v14 = v14
    default:
        var v35 float32 = (float32(float64(((float32(v2) + (((0.1 - 0.1) - ((1.5 * (float32(float64(((float32(1e10) / (((1.5 + 1e10) - 0.0) - (3.14159 + -2.25))) * (1e10 * float32(0.1))))) / 3.14159)) / (((((1.5 + 1e10) - float32(v34)) * 1e10) + ((0.1 + 3.14159) - 1e10)) / (0.1 - (0.0 - 0.1))))) + (0.1 / 1.5))) - 1e10))) * (0.0 / 1.5))
        _ = v35
    }
    switch -1 {
    case -2147483648:
        if ("foo bar" < "\n\t") {
            v36 := GenS1{F0: 0, F1: 2, F2: 377}
            _ = v36
        } else {
            if len(v3) > 0 {
                v3[0] = 2147483648
            }
            var v37 int32 = (v14 - (v14 << ((233) & 63)))
            _ = v37
        }
    case -9223372036854775808:
        fmt.Println(v14)
    case -714:
        v14 = 1
    default:
        if false {
            v4 += (v4 & v4)
        } else {
            switch "\n\t" {
            case "a":
                v4 = v4
            case "hello":
                if v1 {
                    fmt.Println(v1, v2)
                    v38 := GenS1{F0: (v4 ^ v4), F1: 154, F2: -32768}
                    for fi39 := 0; fi39 < 17; fi39++ {
                        fmt.Println(fi39, v34, fi39)
                        if (114 == 98) {
                            for fi40 := 0; fi40 < 20; fi40++ {
                                var v41 int8 = -33
                                if true {
                                    if v1 {
                                        switch fi39 {
                                        case 32768:
                                            if (v2 < v34) {
                                                if len(v3) >= 2 {
                                                    v3 = v3[1:len(v3)]
                                                }
                                                var v42 string = ("\n\t" + "\n\t")
                                                if (!v1) {
                                                    fmt.Println(v34)
                                                }
                                                _ = v42
                                            } else {
                                                switch "" {
                                                case "foo bar":
                                                    fmt.Println(v2)
                                                case "\n\t":
                                                    v14 = v14
                                                default:
                                                    var v43 bool = (((fi40 * 0) <= v13) && (!(!v1)))
                                                    v13 = (v13 / (((v13 * 127)) | 1))
                                                    _ = len(v3)
                                                    _ = v43
                                                }
                                                if v1 {
                                                    var v44 string = "a"
                                                    _ = v44
                                                }
                                            }
                                            v45 := GenS1{F0: v4, F1: 128, F2: 264}
                                            switch 4294967295 {
                                            case 127:
                                                v34 -= 0.1
                                            default:
                                                v1 = v1
                                            }
                                            if v1 {
                                                if (false && v1) {
                                                    fmt.Println(v34, v41, v14)
                                                    if len(v3) > 0 {
                                                        v3[0] = 65535
                                                    }
                                                    var v46 int16 = 524
                                                    v47 := GenS2{F0: 42, F1: v1}
                                                    var v48 string = "\n\t"
                                                    if false {
                                                        var v49 int16 = (v46 >> ((2) & 63))
                                                        _ = v49
                                                    }
                                                    v3 = append(v3, 218)
                                                    if true {
                                                        v41 = v41
                                                    }
                                                    var v50 uint = 721
                                                    v48 = (v48 + (v48 + "a"))
                                                    if v1 {
                                                        v14 = int32(v4)
                                                        var v51 uint = v50
                                                        _ = v51
                                                    }
                                                    var v52 uint64 = 0
                                                    _ = v46
                                                    _ = v47
                                                    _ = v48
                                                    _ = v50
                                                    _ = v52
                                                }
                                                var v53 bool = ((1.5 - 3.14159) == 3.14159)
                                                _ = len(v3)
                                                _ = v53
                                            }
                                            v34 -= (3.14159 + v2)
                                            v2 += v2
                                            switch ("hello" + ("hello" + "\n\t")) {
                                            case "hello":
                                                var v54 uint16 = 706
                                                _ = v54
                                            case "a":
                                                if len(v3) > 0 {
                                                    v3[0] = 691
                                                }
                                            }
                                            switch "hello" {
                                            case "foo bar":
                                                fmt.Println(v4, v4)
                                            case "hello":
                                                var v55 float32 = 1.5
                                                _ = v55
                                            case "":
                                                switch "foo bar" {
                                                case "hello":
                                                    v56 := GenS3{F0: 1}
                                                    _ = v56
                                                default:
                                                    if v1 {
                                                        v14 += v14
                                                        v3 = append(v3, 117)
                                                    }
                                                }
                                            default:
                                                v57 := GenS4{F0: 4294967295, F1: (v14 * v14), F2: 1}
                                                _ = v57
                                            }
                                            v3 = append(v3, 376)
                                            _ = v45
                                        case -189:
                                            fmt.Println(fi39, v1)
                                        default:
                                            fmt.Println(v34, v13, fi39)
                                        }
                                    }
                                    fmt.Println(v13)
                                    var v58 int8 = int8(v4)
                                    _ = v58
                                } else {
                                    v14 = (v14 / ((-812) | 1))
                                }
                                var v59 int = v13
                                _ = len(v3)
                                v60 := GenS1{F0: v4, F1: 128, F2: 327}
                                var v61 float64 = (v34 + v34)
                                var v62 float64 = float64(1e10)
                                _ = fi40
                                _ = v41
                                _ = v59
                                _ = v60
                                _ = v61
                                _ = v62
                            }
                            var v63 int8 = -128
                            switch 91 {
                            case 128:
                                var v64 uint64 = 701
                                _ = v64
                            case -2147483648:
                                if ((!(!true)) || false) {
                                    if len(v3) > 0 {
                                        v3[0] = 987
                                    }
                                    for fi65 := 0; fi65 < 4; fi65++ {
                                        v66 := GenS5{F0: 455}
                                        v2 -= v2
                                        _ = fi65
                                        _ = v66
                                    }
                                } else {
                                    v13 -= (^(fi39))
                                    if (!v1) {
                                        v67 := GenS6{F0: -474, F1: 554, F2: 1000}
                                        v34 -= 3.14159
                                        _ = v67
                                    } else {
                                        for fi68 := 0; fi68 < 9; fi68++ {
                                            v69 := GenS7{F0: 32767}
                                            _ = fi68
                                            _ = v69
                                        }
                                    }
                                    switch (("foo bar" + "foo bar") + ("hello" + "hello")) {
                                    case "":
                                        v70 := GenS8{F0: (v63 + v63), F1: v63}
                                        _ = v70
                                    case "\n\t":
                                        v34 += float64(1e10)
                                    default:
                                        var v71 uint8 = 1
                                        _ = v71
                                    }
                                }
                            }
                            _ = v63
                        }
                        var v72 uint64 = 42
                        _ = fi39
                        _ = v72
                    }
                    _ = v38
                } else {
                    v73 := GenS5{F0: 417}
                    var v74 int = (v13 - v13)
                    var v75 uint8 = 127
                    _ = v73
                    _ = v74
                    _ = v75
                }
            case "":
                v13 = 42
            default:
                switch ("hello" + ("" + (("" + "") + ("a" + ("hello" + ("" + (("foo bar" + "hello") + ("foo bar" + "")))))))) {
                case "\n\t":
                    var v76 uint16 = 42
                    _ = v76
                case "":
                    for fi77 := 0; fi77 < 15; fi77++ {
                        v3 = append(v3, 0)
                        _ = fi77
                    }
                }
            }
            for fi78 := 0; fi78 < 19; fi78++ {
                var v79 float32 = (1e10 - -2.25)
                _ = fi78
                _ = v79
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v34:", v34)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v25
    _ = v3
    _ = v34
    _ = v4
}
