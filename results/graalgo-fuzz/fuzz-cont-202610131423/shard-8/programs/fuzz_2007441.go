package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 int16 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 string; F1 uint32 }
type GenS4 struct { F0 string; F1 uint32 }
type GenS5 struct { F0 uint16; F1 int8; F2 bool }
type GenS6 struct { F0 uint8; F1 int64; F2 uint32 }
type GenS7 struct { F0 uint }
type GenS8 struct { F0 int8 }
type GenS9 struct { F0 uint64; F1 int16 }
type GenS10 struct { F0 uint32; F1 string }
func genFn0() uint64 {
    var v5 uint32 = 2147483648
    fmt.Println(v5)
    var v6 int = 9223372036854775807
    var v7 int16 = 42
    var v8 int8 = 1
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    return 42
}

func main() {
    v1 := make(map[int]bool)
    var v2 uint64 = 4294967295
    var v3 bool = (!false)
    var v4 uint32 = 0
    v9 := v1[-236]
    var v10 int32 = 1
    var v11 int16 = 345
    v12 := GenS1{F0: "hello", F1: int16(v10)}
    v1[-1] = true
    v10 = (v10 | (-(v10)))
    fmt.Println(v10, v11, v2)
    var v13 string = ("a" + "a")
    v14 := GenS1{F0: "hello", F1: (v11 >> ((4294967295) & 63))}
    var v15 uint16 = 42
    var v16 uint8 = 0
    v10 -= (^(v10))
    for fi17 := 0; fi17 < 4; fi17++ {
        for fi18 := 0; fi18 < 5; fi18++ {
            fmt.Println(v4, v10, v16)
            _ = fi18
        }
        _ = fi17
    }
    for fi19 := 0; fi19 < 6; fi19++ {
        switch int(v11) {
        case 42:
            if (1 >= (v15 - v15)) {
                for fi20 := 0; fi20 < 2; fi20++ {
                    var v21 uint32 = v4
                    if v3 {
                        if v3 {
                            if v9 {
                                v22 := GenS1{F0: "", F1: (v11 ^ 1000)}
                                _ = v22
                            } else {
                                switch (fi19 * (^(fi20))) {
                                case 4294967295:
                                    var v23 int32 = -1
                                    _ = v23
                                }
                                var v24 float64 = 1.5
                                v15 -= 1
                                switch v13 {
                                case "foo bar":
                                    var v25 string = (v13 + v13)
                                    _ = v25
                                case "a":
                                    switch v13 {
                                    case "":
                                        var v26 bool = (-123 < fi20)
                                        _ = v26
                                    case "foo bar":
                                        v9 = (!v3)
                                        v13 = v13
                                        v27 := GenS2{F0: v3}
                                        _ = v27
                                    }
                                case "\n\t":
                                    v10 = 0
                                }
                                switch v13 {
                                case "hello":
                                    var v28 float64 = v24
                                    _ = v28
                                case "foo bar":
                                    v13 = (v13 + v13)
                                }
                                _ = v24
                            }
                        } else {
                            v29 := GenS3{F0: ("foo bar" + (("" + "hello") + "a")), F1: (v21 << ((255) & 63))}
                            v11 += (v11 | v11)
                            switch fi19 {
                            case 913:
                                delete(v1, fi20)
                            case -9223372036854775808:
                                var v30 int16 = (v11 + 32767)
                                _ = v30
                            case -857:
                                v21 += (^(v4))
                            default:
                                v31 := []int{(fi19 * (fi19 * fi19)), (fi19 + fi19)}
                                _ = v31
                            }
                            v32 := []int{2147483647}
                            v33 := GenS4{F0: v13, F1: (v21 + (v4 * v4))}
                            if v3 {
                                if ((!v9) || (421 > (v10 * (v10 << ((4294967295) & 63))))) {
                                    v34 := GenS4{F0: (((v13 + ("" + v13)) + "foo bar") + ((v13 + ((v13 + ("foo bar" + v13)) + v13)) + ("" + v13))), F1: v4}
                                    var v35 uint16 = 1
                                    v11 += v11
                                    _ = v34
                                    _ = v35
                                }
                                v36 := GenS5{F0: 32767, F1: 42, F2: true}
                                v16 -= 127
                                var v37 uint = 0
                                switch v13 {
                                case "hello":
                                    if ((!v3) && v3) {
                                        if v3 {
                                            var v38 float32 = 3.14159
                                            fmt.Println(v2, fi19)
                                            _ = v38
                                        }
                                        v9 = v9
                                        _ = len(v1)
                                        var v39 uint16 = (v15 ^ v15)
                                        if (-2.25 <= (1.5 / 1.0)) {
                                            if v3 {
                                                if (v9 && v9) {
                                                    v3 = v3
                                                }
                                            }
                                            _ = len(v32)
                                        }
                                        _ = v39
                                    }
                                default:
                                    v4 = v4
                                }
                                v15 = v15
                                _ = v36
                                _ = v37
                            }
                            fmt.Println(v3, v16)
                            _ = v29
                            _ = v32
                            _ = v33
                        }
                        if v9 {
                            v40 := GenS5{F0: (v15 * v15), F1: -121, F2: v3}
                            fmt.Println(v10)
                            _ = v40
                        }
                        if false {
                            var v41 float64 = (1e10 / (((3.14159 - float64(-2.25)) + -2.25) + 1e10))
                            fmt.Println(v15, fi19, fi19)
                            _ = v41
                        }
                    } else {
                        if v9 {
                            if v3 {
                                switch 128 {
                                case 865:
                                    v4 -= (v21 + (v21 | (v4 - 42)))
                                case -2147483648:
                                    var v42 uint8 = v16
                                    _ = v42
                                case 617:
                                    if v3 {
                                        if (9223372036854775807 != 0) {
                                            var v43 string = v13
                                            var v44 bool = (42 != (-(v11)))
                                            var v45 float64 = 0.1
                                            v46 := []int{(fi20 ^ 406), (fi20 ^ (^(fi20))), fi19}
                                            v9 = v44
                                            v43 = v13
                                            v46 = append(v46, fi20)
                                            _ = v43
                                            _ = v44
                                            _ = v45
                                            _ = v46
                                        }
                                        var v47 float64 = 3.14159
                                        fmt.Println(v47, v15, v9)
                                        _ = v47
                                    }
                                }
                            } else {
                                v48 := []int{int(v11), fi19}
                                v4 += v21
                                v16 = v16
                                _ = v48
                            }
                            v49 := v1[-1]
                            switch (fi19 ^ fi19) {
                            case -9223372036854775808:
                                v50 := GenS6{F0: (^(v16)), F1: 0, F2: uint32(v2)}
                                _ = v50
                            }
                            _ = v49
                        }
                        if (705 <= 1) {
                            fmt.Println(v11)
                        } else {
                            v51 := GenS5{F0: v15, F1: 50, F2: v3}
                            v52 := GenS6{F0: (v16 / (((v16 / ((0) | 1))) | 1)), F1: 1, F2: (v4 << ((uint(v15)) & 63))}
                            _ = v51
                            _ = v52
                        }
                    }
                    v3 = (!(32767 > (^(fi20))))
                    var v53 uint32 = 991
                    _ = fi20
                    _ = v21
                    _ = v53
                }
                _ = len(v1)
            } else {
                switch fi19 {
                case 694:
                    v54 := v1[880]
                    _ = v54
                case -2147483648:
                    if ((v4 * uint32(v2)) >= uint32(v15)) {
                        var v55 int16 = (v11 - 32767)
                        v56 := GenS7{F0: 4294967295}
                        if (fi19 > fi19) {
                            v1[(fi19 * (fi19 / ((fi19) | 1)))] = (false && (4294967295 <= 700))
                            if (42 == 984) {
                                v15 += 801
                                switch (fi19 + (fi19 - -158)) {
                                case -2147483648:
                                    var v57 uint32 = v4
                                    _ = v57
                                }
                                for fi58 := 0; fi58 < 18; fi58++ {
                                    if ((!(v9 && (false && (v3 && ((v15 > v15) || v9))))) || v9) {
                                        fmt.Println(v3, v10, v11)
                                        v59 := GenS7{F0: 2}
                                        _ = v59
                                    }
                                    _ = fi58
                                }
                            } else {
                                switch 4294967295 {
                                case 127:
                                    v3 = (4294967295 <= v2)
                                case 4294967295:
                                    v3 = (-2.25 > ((1e10 + 0.0) - float64(1e10)))
                                case 9223372036854775807:
                                    v1[fi19] = v9
                                }
                                for fi60 := 0; fi60 < 17; fi60++ {
                                    v61 := GenS8{F0: 42}
                                    fmt.Println(v9)
                                    _ = fi60
                                    _ = v61
                                }
                                v62 := []int{fi19, fi19, 4294967295}
                                _ = v62
                            }
                            v55 = -32768
                        } else {
                            var v63 int16 = v55
                            _ = v63
                        }
                        _ = v55
                        _ = v56
                    }
                case -598:
                    v64 := v1[(fi19 ^ (fi19 ^ fi19))]
                    _ = v64
                }
                fmt.Println(fi19)
                v10 += (^(v10))
            }
        case -1:
            for fi65 := 0; fi65 < 15; fi65++ {
                v4 += v4
                if (!v9) {
                    switch (fi19 >> ((1) & 63)) {
                    case 127:
                        delete(v1, (fi19 >> ((255) & 63)))
                    }
                    if v3 {
                        v66 := GenS9{F0: 185, F1: (v11 | (-(v11)))}
                        if (!(fi65 <= fi19)) {
                            if false {
                                v11 = 32767
                                var v67 uint32 = v4
                                _ = v67
                            }
                            if v9 {
                                if v3 {
                                    v10 += int32(v15)
                                    fmt.Println(fi19, v3)
                                    v68 := []int{884, 4294967295}
                                    _ = v68
                                } else {
                                    if (18446744073709551615 != 365) {
                                        v69 := GenS6{F0: v16, F1: -1, F2: (v4 + v4)}
                                        var v70 uint64 = v2
                                        v1[(fi19 - (fi65 - -2147483648))] = (((fi65 * (^(fi65))) <= fi19) || v9)
                                        _ = v69
                                        _ = v70
                                    } else {
                                        fmt.Println(v16, v15)
                                        var v71 int8 = 42
                                        _ = v71
                                    }
                                }
                                if (1e10 > 0.1) {
                                    if false {
                                        v3 = false
                                        v72 := []int{int(v2), (fi19 % (((fi19 & 4294967295)) | 1))}
                                        _ = v72
                                    } else {
                                        v11 += 0
                                        var v73 uint32 = 2147483647
                                        var v74 int64 = 478
                                        fmt.Println(v11, v13, v13)
                                        v75 := GenS10{F0: 2, F1: (v13 + "hello")}
                                        _ = v73
                                        _ = v74
                                        _ = v75
                                    }
                                    v2 -= genFn0()
                                    switch (fi65 % ((fi65) | 1)) {
                                    case -9223372036854775808:
                                        var v76 int = fi65
                                        _ = v76
                                    case 127:
                                        var v77 int = fi19
                                        _ = v77
                                    case 42:
                                        var v78 string = ""
                                        _ = v78
                                    default:
                                        var v79 bool = (128 >= 127)
                                        _ = v79
                                    }
                                } else {
                                    var v80 uint64 = (v2 % (((^(v2))) | 1))
                                    v3 = (!v3)
                                    v1[fi19] = false
                                    _ = v80
                                }
                                v15 = 0
                            } else {
                                _ = len(v1)
                                var v81 uint8 = v16
                                switch 4294967295 {
                                case 127:
                                    if false {
                                        switch fi65 {
                                        case -2147483648:
                                            v82 := []int{-960, fi19, 0, -523}
                                            _ = v82
                                        case -489:
                                            fmt.Println(v2, v4, v3)
                                        case 9223372036854775807:
                                            v1[-2147483648] = (v9 && (!(!((v3 || (((int16(v4) != v11) || (2 == 42)) || v3)) || (true || v3)))))
                                        default:
                                            v9 = true
                                        }
                                        var v83 int16 = (v11 * (v11 % ((0) | 1)))
                                        _ = v83
                                    }
                                case 128:
                                    var v84 int8 = int8(v2)
                                    _ = v84
                                }
                                _ = v81
                            }
                        }
                        _ = v66
                    }
                }
                switch fi65 {
                case 52:
                    v85 := []int{}
                    _ = v85
                case 663:
                    v3 = (!v3)
                }
                _ = fi65
            }
        }
        v86 := GenS8{F0: 100}
        v3 = true
        _ = fi19
        _ = v86
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
