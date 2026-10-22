package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 int; F1 int64 }
type GenS5 struct { F0 int8; F1 string }
func genFn0(p0 int8, p1 uint) int16 {
    var v7 int8 = (p0 / ((p0) | 1))
    v8 := []int{-808, 4294967295, 9223372036854775807}
    if len(v8) > 0 {
        v8[0] = 42
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 0
}
func genFn1(p0 bool, p1 uint64) uint {
    var v9 int32 = 2147483647
    v9 = v9
    _ = p0
    _ = p1
    _ = v9
    return 65535
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[string]string)
    var v3 float64 = 1e10
    var v4 int16 = 32767
    var v5 uint16 = 255
    var v6 float32 = -2.25
    fmt.Println(v6, v3, v5)
    if (8 > 2) {
        var v10 uint16 = v5
        _ = v10
    }
    for fi11 := 0; fi11 < 12; fi11++ {
        var v12 float32 = (v6 * v6)
        _ = fi11
        _ = v12
    }
    var v13 float32 = v6
    v3 = (1.5 * v3)
    for fi14 := 0; fi14 < 5; fi14++ {
        var v15 float64 = (v3 / (float64(v6) + (v3 * float64(v13))))
        _ = fi14
        _ = v15
    }
    if true {
        for fi16 := 0; fi16 < 1; fi16++ {
            var v17 int16 = v4
            switch "" {
            case "\n\t":
                var v18 int32 = 1
                _ = v18
            case "foo bar":
                switch "a" {
                case "foo bar":
                    delete(v2, ("hello" + ""))
                case "a":
                    var v19 int32 = 438
                    _ = v19
                case "":
                    v13 = v6
                default:
                    var v20 bool = (255 < 179)
                    _ = v20
                }
            }
            if (46 >= 1) {
                v13 = (float32((0.1 - (v3 - (float64(v6) - ((float64(v13) * v3) * (v3 / 1.5)))))) / 1.0)
                delete(v2, "")
            }
            _ = fi16
            _ = v17
        }
    } else {
        if (867 <= 4294967295) {
            v21 := GenS1{F0: 482, F1: 9223372036854775807}
            _ = v21
        } else {
            for fi22 := 0; fi22 < 8; fi22++ {
                fmt.Println(v13, v13)
                for fi23 := 0; fi23 < 14; fi23++ {
                    switch "foo bar" {
                    case "\n\t":
                        if false {
                            v5 = v5
                        } else {
                            v4 = 237
                        }
                    case "":
                        var v24 uint = 0
                        _ = v24
                    case "foo bar":
                        v25 := GenS1{F0: uint32(v5), F1: 42}
                        _ = v25
                    }
                    fmt.Println(v5, fi22, v6)
                    _ = fi23
                }
                _ = fi22
            }
            fmt.Println(v6, v6, v5)
            if (false && (false && (1 > 2))) {
                switch int(v4) {
                case 870:
                    v6 = v6
                case 0:
                    v5 = 1
                case 705:
                    if ((-(v4)) < v4) {
                        v5 = v5
                    } else {
                        v3 += -2.25
                    }
                default:
                    switch 2147483647 {
                    case -497:
                        v26 := []int{127, 1, 42}
                        _ = v26
                    }
                }
                v13 += 1.5
                var v27 int64 = -1
                _ = v27
            } else {
                v28 := v1[199]
                _ = v28
            }
        }
        if true {
            v29 := []int{4294967295}
            if true {
                if false {
                    v5 = (v5 / ((v5) | 1))
                    v30 := GenS1{F0: 0, F1: 1}
                    for fi31 := 0; fi31 < 18; fi31++ {
                        var v32 int = (fi31 % (((fi31 | -201)) | 1))
                        for fi33 := 0; fi33 < 15; fi33++ {
                            if (int16(fi33) != 0) {
                                v3 += v3
                                if false {
                                    if (!(fi31 <= fi33)) {
                                        fmt.Println(v32, fi33, v13)
                                        var v34 string = "a"
                                        _ = v34
                                    } else {
                                        v35 := GenS1{F0: 68, F1: -9223372036854775808}
                                        var v36 uint32 = 0
                                        _ = v35
                                        _ = v36
                                    }
                                    if (!((!true) && false)) {
                                        switch "hello" {
                                        case "foo bar":
                                            if len(v29) >= 2 {
                                                v29 = v29[1:len(v29)]
                                            }
                                        default:
                                            switch 32767 {
                                            case 127:
                                                var v37 int8 = 127
                                                _ = v37
                                            }
                                            v29 = append(v29, fi33)
                                            var v38 uint = 748
                                            _ = v38
                                        }
                                        var v39 uint16 = (v5 | 42)
                                        v3 = (-2.25 / v3)
                                        if false {
                                            if true {
                                                v32 = fi33
                                            } else {
                                                var v40 int64 = -9223372036854775808
                                                _ = v40
                                            }
                                            if true {
                                                delete(v1, v32)
                                                if (true || (int16(fi31) < (-(v4)))) {
                                                    v41 := GenS1{F0: 42, F1: 566}
                                                    v3 += ((-2.25 * -2.25) * float64((v13 + 1.5)))
                                                    _ = v41
                                                } else {
                                                    if len(v29) >= 2 {
                                                        v29 = v29[1:len(v29)]
                                                    }
                                                    var v42 float64 = v3
                                                    _ = v42
                                                }
                                                v4 = -648
                                                var v43 float32 = 0.1
                                                _ = v43
                                            }
                                        }
                                        _ = v39
                                    }
                                    var v44 int8 = -1
                                    fmt.Println(v3, v6, v44)
                                    var v45 bool = (42 == v32)
                                    _ = v44
                                    _ = v45
                                } else {
                                    var v46 bool = false
                                    _ = v46
                                }
                            } else {
                                var v47 int8 = 42
                                _ = v47
                            }
                            v13 = 1e10
                            _ = fi33
                        }
                        _ = fi31
                        _ = v32
                    }
                    _ = v30
                } else {
                    delete(v2, "\n\t")
                    v4 = genFn0(127, genFn1(true, 2))
                }
                v5 = (v5 | 32767)
            } else {
                v48 := GenS1{F0: 505, F1: -9223372036854775808}
                v49 := GenS1{F0: 761, F1: -916}
                if true {
                    if (!(genFn1(true, 18446744073709551615) >= uint(v5))) {
                        if ((42 == 128) || true) {
                            v3 += v3
                        } else {
                            if ((v5 >> ((255) & 63)) >= (v5 - v5)) {
                                var v50 int64 = -1
                                for fi51 := 0; fi51 < 3; fi51++ {
                                    var v52 float32 = (v6 + v13)
                                    var v53 uint32 = 47
                                    _ = fi51
                                    _ = v52
                                    _ = v53
                                }
                                v13 = v13
                                _ = v50
                            }
                            var v54 float32 = -2.25
                            var v55 uint8 = 1
                            _ = v54
                            _ = v55
                        }
                        if true {
                            fmt.Println(v4)
                            for fi56 := 0; fi56 < 4; fi56++ {
                                var v57 int64 = -9223372036854775808
                                var v58 string = "a"
                                _ = fi56
                                _ = v57
                                _ = v58
                            }
                        } else {
                            v59 := GenS1{F0: 1, F1: 948}
                            var v60 string = ""
                            var v61 float32 = (0.1 * float32(v3))
                            _ = v59
                            _ = v60
                            _ = v61
                        }
                    }
                    if (false && false) {
                        if (255 != 2) {
                            var v62 int32 = 2147483647
                            var v63 uint8 = 42
                            if (false && ((4294967295 != 911) && (!false))) {
                                fmt.Println(v3)
                                v64 := GenS2{F0: 22}
                                _ = v64
                            }
                            _ = v62
                            _ = v63
                        } else {
                            switch -9223372036854775808 {
                            case 128:
                                _ = len(v29)
                            case 4294967295:
                                v3 = v3
                            }
                            v13 -= ((v13 + (v6 - (v6 + v6))) + v6)
                            if (-1 >= int32(v5)) {
                                if true {
                                    if (42 > 70) {
                                        delete(v1, 54)
                                        var v65 uint16 = v5
                                        for fi66 := 0; fi66 < 4; fi66++ {
                                            if len(v29) >= 2 {
                                                v29 = v29[1:len(v29)]
                                            }
                                            v67 := GenS2{F0: -71}
                                            _ = fi66
                                            _ = v67
                                        }
                                        _ = v65
                                    }
                                    switch "" {
                                    case "foo bar":
                                        v68 := v2[("a" + "\n\t")]
                                        _ = v68
                                    }
                                }
                                v69 := GenS3{F0: (v4 << ((57) & 63))}
                                fmt.Println(v5, v6, v5)
                                _ = v69
                            } else {
                                for fi70 := 0; fi70 < 1; fi70++ {
                                    var v71 float64 = (float64(float32((float64(float32(v3)) * v3))) / v3)
                                    _ = fi70
                                    _ = v71
                                }
                            }
                        }
                        v4 = v4
                    }
                }
                _ = v48
                _ = v49
            }
            var v72 string = ""
            _ = v29
            _ = v72
        }
    }
    for fi73 := 0; fi73 < 6; fi73++ {
        if (32767 > (v5 % ((v5) | 1))) {
            if true {
                var v74 uint32 = 65535
                var v75 uint16 = (^(v5))
                var v76 uint16 = (v5 ^ v5)
                _ = v74
                _ = v75
                _ = v76
            }
        } else {
            v77 := v2[""]
            _ = v77
        }
        v3 = (0.1 - v3)
        var v78 bool = (1 <= 1)
        _ = fi73
        _ = v78
    }
    v79 := GenS4{F0: 127, F1: 1}
    v80 := []int{-881, 162}
    v81 := GenS5{F0: -30, F1: ("" + "hello")}
    var v82 uint8 = 1
    _ = len(v80)
    fmt.Println(v5, v6, v6)
    v83 := GenS1{F0: 2147483647, F1: 0}
    var v84 uint32 = 2
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("slicelen v80:", len(v80), v80)
    fmt.Println("v82:", v82)
    fmt.Println("v84:", v84)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v79
    _ = v80
    _ = v81
    _ = v82
    _ = v83
    _ = v84
}
