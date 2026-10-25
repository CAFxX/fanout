package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 string }
type GenS3 struct { F0 uint16; F1 string; F2 uint8 }
type GenS4 struct { F0 bool; F1 uint }
type GenS5 struct { F0 uint8; F1 int8; F2 int }
type GenS6 struct { F0 int8; F1 int64 }
type GenS7 struct { F0 int32; F1 uint16; F2 uint8 }
type GenS8 struct { F0 uint16; F1 uint64; F2 uint32 }
func genFn0() bool {
    for fi6 := 0; fi6 < 16; fi6++ {
        v7 := GenS1{F0: 466}
        var v8 int16 = 42
        if (-103 == 1) {
            v9 := []int{}
            if len(v9) >= 2 {
                v9 = v9[1:len(v9)]
            }
            _ = len(v9)
            _ = v9
        } else {
            if (879 <= 65535) {
                v8 -= (v8 % ((v8) | 1))
            }
            fmt.Println(fi6)
            var v10 bool = (false && (v8 > (v8 << ((185) & 63))))
            _ = v10
        }
        _ = fi6
        _ = v7
        _ = v8
    }
    var v11 float32 = -2.25
    for fi12 := 0; fi12 < 16; fi12++ {
        v13 := []int{(-(fi12)), fi12, 32767, fi12}
        _ = fi12
        _ = v13
    }
    fmt.Println(v11)
    v14 := GenS1{F0: 185}
    _ = v11
    _ = v14
    return true
}

func main() {
    var v1 float64 = float64(1e10)
    v2 := []int{-2147483648}
    var v3 int64 = 9223372036854775807
    v4 := []string{"a", "\n\t"}
    var v5 bool = false
    switch "foo bar" {
    case "hello":
        if v5 {
            var v15 uint16 = 569
            switch "" {
            case "hello":
                if (int16(v15) >= -343) {
                    v16 := GenS1{F0: 688}
                    _ = v16
                }
            case "foo bar":
                if true {
                    if v5 {
                        for fi17 := 0; fi17 < 19; fi17++ {
                            v15 = 1
                            v18 := GenS2{F0: "foo bar"}
                            if false {
                                v19 := GenS2{F0: ("hello" + "")}
                                for fi20 := 0; fi20 < 2; fi20++ {
                                    v5 = true
                                    _ = fi20
                                }
                                _ = v19
                            } else {
                                if true {
                                    var v21 uint64 = uint64(v15)
                                    for fi22 := 0; fi22 < 19; fi22++ {
                                        var v23 uint64 = v21
                                        _ = fi22
                                        _ = v23
                                    }
                                    var v24 int32 = 42
                                    _ = v21
                                    _ = v24
                                }
                            }
                            _ = fi17
                            _ = v18
                        }
                    } else {
                        var v25 int16 = 32767
                        _ = v25
                    }
                    if (0.0 >= ((1.5 + -2.25) + 1e10)) {
                        v5 = (!((v3 + v3) == -75))
                        var v26 bool = (true && (-1 <= (^(v3))))
                        v15 = v15
                        _ = v26
                    } else {
                        for fi27 := 0; fi27 < 3; fi27++ {
                            var v28 int8 = 0
                            if v5 {
                                if true {
                                    v29 := GenS3{F0: (v15 - (v15 | v15)), F1: ("\n\t" + "foo bar"), F2: 43}
                                    switch "hello" {
                                    case "foo bar":
                                        v3 = v3
                                    case "":
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    case "hello":
                                        var v30 uint16 = v15
                                        _ = v30
                                    default:
                                        var v31 int64 = v3
                                        _ = v31
                                    }
                                    v32 := GenS4{F0: v5, F1: 784}
                                    _ = v29
                                    _ = v32
                                } else {
                                    if v5 {
                                        fmt.Println(v3, v28, v15)
                                    }
                                    if v5 {
                                        if len(v4) >= 2 {
                                            v4 = v4[1:len(v4)]
                                        }
                                        v2 = append(v2, (fi27 ^ fi27))
                                        var v33 int64 = (v3 - v3)
                                        _ = v33
                                    }
                                    for fi34 := 0; fi34 < 19; fi34++ {
                                        if (v5 && (18446744073709551615 > 0)) {
                                            var v35 float64 = 0.1
                                            _ = v35
                                        }
                                        fmt.Println(v5, v5, fi27)
                                        var v36 int = fi34
                                        switch 32767 {
                                        case 602:
                                            var v37 uint64 = 1
                                            _ = v37
                                        case 0:
                                            v38 := GenS4{F0: false, F1: 65535}
                                            _ = v38
                                        case 4294967295:
                                            var v39 uint8 = 128
                                            _ = v39
                                        default:
                                            var v40 int = (fi34 - v36)
                                            _ = v40
                                        }
                                        _ = fi34
                                        _ = v36
                                    }
                                }
                                fmt.Println(v15, v15)
                                var v41 int16 = 0
                                _ = v41
                            }
                            v3 = (v3 + (v3 | -9223372036854775808))
                            _ = fi27
                            _ = v28
                        }
                        v3 -= v3
                        v42 := GenS4{F0: ((0 == 18446744073709551615) && (1 >= 120)), F1: 255}
                        _ = v42
                    }
                    var v43 uint32 = 65535
                    _ = v43
                }
            default:
                var v44 int64 = (-(v3))
                _ = v44
            }
            _ = v15
        }
    case "":
        for fi45 := 0; fi45 < 20; fi45++ {
            v3 = (v3 + 0)
            var v46 int = (^(fi45))
            if false {
                var v47 int16 = -639
                for fi48 := 0; fi48 < 7; fi48++ {
                    var v49 uint8 = 0
                    _ = fi48
                    _ = v49
                }
                _ = v47
            } else {
                var v50 uint16 = 807
                if v5 {
                    if len(v2) > 0 {
                        v2[0] = int(v50)
                    }
                    if (1 != 0) {
                        for fi51 := 0; fi51 < 15; fi51++ {
                            var v52 int64 = (^(v3))
                            fmt.Println(v50, v50, v3)
                            _ = fi51
                            _ = v52
                        }
                        v50 = 612
                        v2 = append(v2, (v46 - fi45))
                    } else {
                        if v5 {
                            if (-1 > -1) {
                                var v53 uint16 = (^(v50))
                                _ = v53
                            } else {
                                if len(v2) > 0 {
                                    v2[0] = (v46 & v46)
                                }
                                v46 += 9223372036854775807
                            }
                            var v54 uint64 = 848
                            if len(v2) >= 2 {
                                v2 = v2[1:len(v2)]
                            }
                            _ = v54
                        } else {
                            for fi55 := 0; fi55 < 20; fi55++ {
                                fmt.Println(v5)
                                if ((!((fi45 / (((v46 - 32768)) | 1)) == 32767)) || (v50 < v50)) {
                                    v56 := GenS5{F0: 92, F1: 1, F2: v46}
                                    v50 += v50
                                    var v57 float64 = v1
                                    _ = v56
                                    _ = v57
                                }
                                _ = fi55
                            }
                            if (!v5) {
                                if len(v4) > 0 {
                                    v4[0] = ""
                                }
                            }
                            fmt.Println(fi45, v3)
                        }
                        v4 = append(v4, "\n\t")
                    }
                    for fi58 := 0; fi58 < 12; fi58++ {
                        v3 += (^(v3))
                        v46 -= (fi45 << ((815) & 63))
                        _ = fi58
                    }
                }
                _ = v50
            }
            _ = fi45
            _ = v46
        }
    default:
        switch "a" {
        case "":
            if v5 {
                for fi59 := 0; fi59 < 12; fi59++ {
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                    _ = fi59
                }
                var v60 bool = v5
                for fi61 := 0; fi61 < 7; fi61++ {
                    v2 = append(v2, int(v3))
                    _ = fi61
                }
                _ = v60
            } else {
                if v5 {
                    if len(v4) > 0 {
                        v4[0] = ("a" + "")
                    }
                }
            }
        }
    }
    if v5 {
        var v62 int64 = v3
        _ = v62
    } else {
        v2 = append(v2, -2147483648)
        v3 = 0
    }
    if v5 {
        for fi63 := 0; fi63 < 17; fi63++ {
            var v64 uint = 226
            v65 := GenS6{F0: -63, F1: (v3 / (((-(v3))) | 1))}
            _ = fi63
            _ = v64
            _ = v65
        }
    }
    for fi66 := 0; fi66 < 6; fi66++ {
        v1 -= (0.0 * -2.25)
        _ = fi66
    }
    if (127 >= 9223372036854775807) {
        if v5 {
            v2 = append(v2, -2147483648)
            for fi67 := 0; fi67 < 15; fi67++ {
                fmt.Println(fi67, v5)
                _ = fi67
            }
        } else {
            fmt.Println(v5, v3, v3)
            v5 = (-66 > 127)
        }
        v68 := GenS7{F0: 0, F1: 1, F2: 171}
        _ = v68
    } else {
        var v69 int32 = 0
        fmt.Println(v3, v69)
        _ = v69
    }
    fmt.Println(v5, v5)
    switch "a" {
    case "hello":
        for fi70 := 0; fi70 < 19; fi70++ {
            var v71 uint64 = 0
            fmt.Println(v1)
            v5 = false
            _ = fi70
            _ = v71
        }
    }
    v3 = -1
    v72 := GenS7{F0: -711, F1: 42, F2: 127}
    if (!v5) {
        if len(v4) > 0 {
            v4[0] = ("\n\t" + "hello")
        }
        _ = len(v2)
        v73 := GenS8{F0: 255, F1: 363, F2: 1}
        _ = v73
    } else {
        _ = len(v4)
        if (!(v3 < -1)) {
            switch -374 {
            case 479:
                for fi74 := 0; fi74 < 1; fi74++ {
                    v5 = (129 != 984)
                    _ = fi74
                }
            default:
                v5 = v5
            }
        } else {
            for fi75 := 0; fi75 < 13; fi75++ {
                if len(v2) >= 2 {
                    v2 = v2[1:len(v2)]
                }
                var v76 int32 = -1
                var v77 uint8 = 0
                _ = fi75
                _ = v76
                _ = v77
            }
        }
        v1 = float64(float32(-2.25))
    }
    v78 := GenS4{F0: (2147483647 > 1), F1: 255}
    v2 = append(v2, 128)
    for fi79 := 0; fi79 < 8; fi79++ {
        v2 = append(v2, (^(fi79)))
        _ = fi79
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v72
    _ = v78
}
