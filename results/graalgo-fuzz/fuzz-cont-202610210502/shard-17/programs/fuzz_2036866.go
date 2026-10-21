package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 bool; F2 string }
type GenS2 struct { F0 uint8; F1 int64 }
type GenS3 struct { F0 bool; F1 int32; F2 uint32 }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint; F1 string }
type GenS6 struct { F0 int16; F1 int8; F2 uint32 }
func genFn0() int64 {
    var v6 uint16 = 299
    for fi7 := 0; fi7 < 10; fi7++ {
        for fi8 := 0; fi8 < 11; fi8++ {
            v6 += (v6 - (^(v6)))
            _ = fi8
        }
        var v9 uint8 = 70
        fmt.Println(v6, v9, v9)
        _ = fi7
        _ = v9
    }
    if (!(!(42 <= 1))) {
        var v10 int32 = -1
        v11 := GenS1{F0: false, F1: true, F2: "\n\t"}
        _ = v10
        _ = v11
    } else {
        var v12 uint64 = 292
        var v13 uint64 = v12
        v12 -= (v12 % ((v12) | 1))
        _ = v12
        _ = v13
    }
    _ = v6
    return -892
}
func genFn1() uint32 {
    v14 := []int{-590}
    for fi15 := 0; fi15 < 10; fi15++ {
        for fi16 := 0; fi16 < 10; fi16++ {
            var v17 uint8 = 189
            var v18 string = (("hello" + ("a" + ("hello" + "hello"))) + "\n\t")
            v19 := GenS2{F0: v17, F1: -9223372036854775808}
            _ = fi16
            _ = v17
            _ = v18
            _ = v19
        }
        _ = fi15
    }
    _ = v14
    return 994
}

func main() {
    var v1 uint64 = 1
    var v2 int16 = 529
    var v3 float32 = 0.0
    var v4 uint32 = 42
    v5 := make(map[int]bool)
    switch 1 {
    case -92:
        if ((v2 + 1) == (v2 * v2)) {
            v20 := GenS2{F0: 74, F1: 9223372036854775807}
            switch ("foo bar" + (("" + "\n\t") + "foo bar")) {
            case "":
                if false {
                    var v21 uint64 = v1
                    if false {
                        for fi22 := 0; fi22 < 10; fi22++ {
                            switch (fi22 * (fi22 << ((65535) & 63))) {
                            case -998:
                                var v23 uint64 = 987
                                _ = v23
                            case -9223372036854775808:
                                v24 := v5[fi22]
                                _ = v24
                            }
                            fmt.Println(v2, v4)
                            if false {
                                var v25 string = ("hello" + "\n\t")
                                if ((!true) && ((!(float32(1e10) == 0.1)) || (0 == 18446744073709551615))) {
                                    var v26 uint = 552
                                    fmt.Println(v3)
                                    v25 = v25
                                    _ = v26
                                }
                                _ = v25
                            } else {
                                fmt.Println(v21)
                                var v27 uint16 = 878
                                _ = v27
                            }
                            _ = fi22
                        }
                    } else {
                        fmt.Println(v3, v4, v4)
                    }
                    _ = v21
                }
            case "a":
                if true {
                    switch 4294967295 {
                    case -511:
                        if ((3.14159 - 0.1) > ((((1.5 * 3.14159) / 1.0) + 0.0) * (-2.25 / (((float64((float32(((((-2.25 * 1e10) / (0.1 * float64(-2.25))) * float64(1.5)) + 1e10)) * v3)) / 0.1) - 0.1) / (float64(v3) - -2.25))))) {
                            switch "hello" {
                            case "\n\t":
                                v2 += v2
                            default:
                                fmt.Println(v3, v1)
                            }
                            v5[4294967295] = (true || false)
                        } else {
                            fmt.Println(v3, v1)
                            var v28 float64 = float64(-2.25)
                            _ = v28
                        }
                    case -301:
                        var v29 int32 = 635
                        _ = v29
                    default:
                        v30 := GenS1{F0: (true && false), F1: true, F2: "foo bar"}
                        _ = v30
                    }
                }
            case "\n\t":
                v2 = v2
            default:
                v31 := []int{}
                _ = v31
            }
            _ = v20
        }
    case 32768:
        if (false && (v2 < 32767)) {
            v2 += 1000
        }
    }
    v3 = (((v3 - 0.0) - ((0.0 - v3) / 1e10)) * float32((0.0 * (float64((v3 + 3.14159)) / 1e10))))
    v1 += 1
    v32 := []int{162, 128, 32767, -1}
    if (false || ((float32(1e10) <= v3) || ("hello" >= "a"))) {
        if len(v32) > 0 {
            v32[0] = 0
        }
        for fi33 := 0; fi33 < 10; fi33++ {
            var v34 int8 = 1
            v3 = v3
            _ = fi33
            _ = v34
        }
    } else {
        switch "hello" {
        case "\n\t":
            for fi35 := 0; fi35 < 4; fi35++ {
                var v36 uint64 = v1
                v37 := GenS3{F0: false, F1: 216, F2: (^(v4))}
                _ = fi35
                _ = v36
                _ = v37
            }
        case "foo bar":
            v2 = (v2 << ((42) & 63))
        case "":
            fmt.Println(v1)
        default:
            var v38 int8 = 0
            _ = v38
        }
    }
    var v39 int16 = (^(v2))
    if (genFn0() > genFn0()) {
        if (((1 > 32767) && (v4 >= v4)) || (v1 != (v1 / ((v1) | 1)))) {
            for fi40 := 0; fi40 < 11; fi40++ {
                if (201 != 0) {
                    var v41 float32 = v3
                    _ = v41
                }
                _ = fi40
            }
            if len(v32) >= 2 {
                v32 = v32[1:len(v32)]
            }
            v2 += v39
        } else {
            fmt.Println(v2, v2)
            v42 := GenS1{F0: true, F1: (-2147483648 < -1), F2: ("foo bar" + ("" + ""))}
            _ = v42
        }
        v43 := GenS2{F0: 74, F1: 9223372036854775807}
        var v44 int = 32768
        _ = v43
        _ = v44
    } else {
        if false {
            var v45 int64 = 1
            v3 += 1.5
            _ = v45
        } else {
            var v46 uint64 = 42
            var v47 uint16 = 107
            for fi48 := 0; fi48 < 14; fi48++ {
                var v49 uint64 = (v46 << ((255) & 63))
                var v50 int64 = 9223372036854775807
                switch fi48 {
                case 9223372036854775807:
                    var v51 string = ("hello" + "foo bar")
                    _ = v51
                default:
                    var v52 float32 = 1.5
                    _ = v52
                }
                _ = fi48
                _ = v49
                _ = v50
            }
            _ = v46
            _ = v47
        }
    }
    if true {
        var v53 uint16 = 2
        var v54 int64 = 655
        _ = v53
        _ = v54
    }
    var v55 string = "foo bar"
    switch v55 {
    case "\n\t":
        v1 = v1
    case "":
        v56 := GenS3{F0: (1 < 2), F1: -874, F2: 0}
        _ = v56
    default:
        for fi57 := 0; fi57 < 6; fi57++ {
            switch -9223372036854775808 {
            case 4294967295:
                var v58 uint = 18446744073709551615
                _ = v58
            case 9223372036854775807:
                var v59 uint32 = v4
                _ = v59
            case 2147483647:
                v55 = ("" + v55)
            default:
                var v60 float32 = v3
                _ = v60
            }
            v39 += v39
            v61 := GenS4{F0: (("hello" + v55) + v55)}
            _ = fi57
            _ = v61
        }
    }
    if true {
        v4 += (v4 & v4)
        var v62 uint64 = (^(v1))
        v1 -= v62
        _ = v62
    }
    for fi63 := 0; fi63 < 4; fi63++ {
        fmt.Println(v3)
        var v64 uint16 = 965
        _ = fi63
        _ = v64
    }
    if false {
        delete(v5, 127)
        v39 = (v39 % (((v39 + int16(v1))) | 1))
        fmt.Println(v3)
    }
    var v65 uint16 = 65535
    if (false && true) {
        delete(v5, 108)
        if true {
            for fi66 := 0; fi66 < 13; fi66++ {
                var v67 int64 = 739
                switch -540 {
                case -9223372036854775808:
                    var v68 uint16 = (v65 % (((v65 % ((uint16(v1)) | 1))) | 1))
                    _ = v68
                }
                _ = fi66
                _ = v67
            }
            if false {
                if (!(v1 <= (v1 | (^(v1))))) {
                    var v69 int64 = -499
                    if false {
                        var v70 bool = (true || (v39 == (v2 ^ (v2 % ((1) | 1)))))
                        for fi71 := 0; fi71 < 5; fi71++ {
                            v55 = "foo bar"
                            if (!("\n\t" <= v55)) {
                                v72 := GenS1{F0: ((3.14159 > (-2.25 + 3.14159)) || true), F1: v70, F2: v55}
                                v73 := GenS3{F0: ((v70 || v70) && (!((v55 >= v55) || v70))), F1: 2147483647, F2: (v4 ^ 449)}
                                _ = v72
                                _ = v73
                            } else {
                                switch "a" {
                                case "\n\t":
                                    switch "" {
                                    case "":
                                        if len(v32) >= 2 {
                                            v32 = v32[1:len(v32)]
                                        }
                                    case "a":
                                        var v74 int16 = -321
                                        _ = v74
                                    }
                                default:
                                    var v75 uint32 = (v4 % ((genFn1()) | 1))
                                    _ = v75
                                }
                                v76 := GenS4{F0: v55}
                                _ = v76
                            }
                            if false {
                                v77 := GenS5{F0: 600, F1: (v55 + v55)}
                                _ = v77
                            }
                            _ = fi71
                        }
                        _ = v70
                    }
                    _ = v69
                } else {
                    fmt.Println(v2)
                }
            } else {
                var v78 uint8 = 2
                fmt.Println(v65, v78)
                _ = v78
            }
            v4 = 1
        } else {
            v79 := GenS3{F0: true, F1: int32(v2), F2: genFn1()}
            v4 -= (^(v4))
            for fi80 := 0; fi80 < 13; fi80++ {
                v4 -= 2147483647
                v81 := GenS6{F0: 326, F1: -1, F2: v4}
                v65 += v65
                _ = fi80
                _ = v81
            }
            _ = v79
        }
    } else {
        var v82 int8 = -128
        var v83 float32 = v3
        _ = v82
        _ = v83
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v32:", len(v32), v32)
    fmt.Println("v39:", v39)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v55:", v55)
    fmt.Println("v65:", v65)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v39
    _ = v4
    _ = v5
    _ = v55
    _ = v65
}
