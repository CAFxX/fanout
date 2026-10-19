package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 string; F2 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 int32; F1 bool; F2 int32 }
type GenS4 struct { F0 uint8; F1 int8 }
type GenS5 struct { F0 int8; F1 uint; F2 uint }
type GenS6 struct { F0 uint }
type GenS7 struct { F0 uint32 }
func genFn0(p0 uint64) int8 {
    fmt.Println(p0)
    p0 = 0
    _ = p0
    return 1
}
func genFn1(p0 int64) int16 {
    p0 -= p0
    var v6 int16 = -906
    var v7 string = "\n\t"
    switch "" {
    case "a":
        for fi8 := 0; fi8 < 17; fi8++ {
            if (42 < 735) {
                switch fi8 {
                case 662:
                    fmt.Println(v7)
                case -61:
                    if true {
                        var v9 float32 = 0.1
                        _ = v9
                    }
                }
                for fi10 := 0; fi10 < 14; fi10++ {
                    fmt.Println(v6, v7)
                    switch v7 {
                    case "foo bar":
                        v11 := []int{42, fi10, fi10, 1}
                        _ = v11
                    case "\n\t":
                        var v12 uint16 = 517
                        _ = v12
                    case "a":
                        var v13 float64 = (3.14159 * 0.0)
                        _ = v13
                    default:
                        if (!false) {
                            p0 -= int64(fi8)
                            var v14 int8 = -1
                            v15 := []int{32767, (fi10 >> ((4294967295) & 63)), 563, -1}
                            if (0.0 >= 0.1) {
                                switch ("a" + "") {
                                case "hello":
                                    switch fi10 {
                                    case -9223372036854775808:
                                        v15 = append(v15, fi8)
                                    case 128:
                                        var v16 int32 = -502
                                        _ = v16
                                    case -1:
                                        var v17 uint = 188
                                        var v18 float64 = 0.0
                                        _ = v17
                                        _ = v18
                                    }
                                case "foo bar":
                                    var v19 string = v7
                                    _ = v19
                                default:
                                    v20 := GenS1{F0: -1, F1: v7, F2: v6}
                                    _ = v20
                                }
                                v21 := GenS1{F0: 1, F1: (v7 + (("hello" + (v7 + "")) + ("foo bar" + v7))), F2: (v6 + v6)}
                                _ = v21
                            } else {
                                if len(v15) >= 2 {
                                    v15 = v15[1:len(v15)]
                                }
                            }
                            if ((false && false) && true) {
                                v22 := GenS2{F0: v14}
                                v23 := GenS3{F0: -2147483648, F1: (false || true), F2: -2147483648}
                                _ = v22
                                _ = v23
                            } else {
                                v7 = v7
                                var v24 int8 = (v14 * v14)
                                if len(v15) >= 2 {
                                    v15 = v15[1:len(v15)]
                                }
                                v14 += 89
                                var v25 int32 = -2147483648
                                _ = len(v15)
                                _ = v24
                                _ = v25
                            }
                            _ = v14
                            _ = v15
                        }
                    }
                    v7 = ("hello" + v7)
                    _ = fi10
                }
                if (34 != 74) {
                    var v26 int8 = -128
                    var v27 int16 = (v6 << ((255) & 63))
                    _ = v26
                    _ = v27
                } else {
                    fmt.Println(p0, p0)
                }
            }
            _ = fi8
        }
    case "":
        p0 -= 9223372036854775807
    }
    for fi28 := 0; fi28 < 7; fi28++ {
        v29 := GenS2{F0: 100}
        switch "" {
        case "\n\t":
            v30 := GenS2{F0: -1}
            _ = v30
        }
        fmt.Println(v7, fi28)
        _ = fi28
        _ = v29
    }
    _ = p0
    _ = v6
    _ = v7
    return v6
}

func main() {
    var v1 float64 = ((float64(1.5) - 1e10) - 1e10)
    var v2 string = (("" + "\n\t") + ("" + ""))
    var v3 uint16 = 377
    var v4 int32 = 42
    var v5 uint8 = 2
    v3 = 42
    v31 := []int{42, 2147483647}
    v31 = append(v31, -9223372036854775808)
    v32 := GenS3{F0: v4, F1: false, F2: v4}
    v33 := GenS4{F0: (v5 - 96), F1: 100}
    if (0 >= 42) {
        fmt.Println(v4, v4, v1)
    } else {
        v2 = ((v2 + "a") + v2)
        v1 = (((v1 / 1.5) - v1) + (((1.5 + v1) + float64(1.5)) - 0.0))
        v5 = v5
    }
    if true {
        if ((false || true) || false) {
            v1 = 0.1
            for fi34 := 0; fi34 < 14; fi34++ {
                if false {
                    if true {
                        switch v2 {
                        case "":
                            for fi35 := 0; fi35 < 6; fi35++ {
                                v2 = ""
                                if (194 < 449) {
                                    v36 := GenS2{F0: -75}
                                    v4 -= (v4 - -119)
                                    var v37 uint8 = v5
                                    _ = v36
                                    _ = v37
                                } else {
                                    v4 += int32(v3)
                                    v38 := GenS5{F0: 127, F1: 4294967295, F2: 18446744073709551615}
                                    if len(v31) >= 2 {
                                        v31 = v31[1:len(v31)]
                                    }
                                    _ = v38
                                }
                                var v39 uint32 = 567
                                v40 := GenS6{F0: 42}
                                v41 := GenS5{F0: 0, F1: 18446744073709551615, F2: 491}
                                _ = fi35
                                _ = v39
                                _ = v40
                                _ = v41
                            }
                        default:
                            fmt.Println(v3, v3, fi34)
                        }
                        for fi42 := 0; fi42 < 2; fi42++ {
                            v5 -= v5
                            _ = fi42
                        }
                    } else {
                        var v43 int32 = v4
                        var v44 int = fi34
                        switch v2 {
                        case "a":
                            var v45 float32 = 1e10
                            _ = v45
                        case "hello":
                            v43 = (v4 / (((-(v43))) | 1))
                        case "foo bar":
                            var v46 uint = 1
                            _ = v46
                        default:
                            var v47 int64 = -9223372036854775808
                            _ = v47
                        }
                        _ = v43
                        _ = v44
                    }
                }
                _ = fi34
            }
            v48 := GenS4{F0: v5, F1: 68}
            _ = v48
        }
    }
    var v49 string = "a"
    for fi50 := 0; fi50 < 12; fi50++ {
        if true {
            fmt.Println(v2, v1, v1)
        } else {
            switch fi50 {
            case 32768:
                for fi51 := 0; fi51 < 9; fi51++ {
                    var v52 int8 = 52
                    v53 := GenS4{F0: 142, F1: int8(v3)}
                    _ = fi51
                    _ = v52
                    _ = v53
                }
            case 4294967295:
                v2 = v2
            case 2147483647:
                for fi54 := 0; fi54 < 6; fi54++ {
                    v5 -= v5
                    _ = fi54
                }
            default:
                if (-492 != -38) {
                    v55 := GenS3{F0: -122, F1: true, F2: -1}
                    v56 := GenS5{F0: -125, F1: 1, F2: 255}
                    _ = v55
                    _ = v56
                }
            }
            if (!(!(1 >= 18446744073709551615))) {
                v31 = append(v31, 643)
                if len(v31) > 0 {
                    v31[0] = fi50
                }
            }
            switch (("" + v49) + (v49 + v49)) {
            case "hello":
                var v57 int8 = 43
                _ = v57
            case "\n\t":
                for fi58 := 0; fi58 < 5; fi58++ {
                    if false {
                        fmt.Println(v1, fi50)
                        var v59 int32 = -1
                        v59 -= 467
                        _ = v59
                    }
                    v2 = ((((v2 + (v49 + "\n\t")) + ((v2 + "\n\t") + "")) + v2) + (v49 + v49))
                    if false {
                        v60 := GenS7{F0: 745}
                        var v61 int16 = genFn1(1)
                        v49 = v2
                        if len(v31) >= 2 {
                            v31 = v31[1:len(v31)]
                        }
                        _ = v60
                        _ = v61
                    } else {
                        if len(v31) >= 2 {
                            v31 = v31[1:len(v31)]
                        }
                        var v62 int = 845
                        _ = v62
                    }
                    _ = fi58
                }
            case "":
                if len(v31) > 0 {
                    v31[0] = (fi50 * fi50)
                }
            }
        }
        var v63 int16 = 42
        if false {
            for fi64 := 0; fi64 < 17; fi64++ {
                var v65 int = (fi50 >> ((906) & 63))
                var v66 uint16 = v3
                _ = fi64
                _ = v65
                _ = v66
            }
            fmt.Println(v63)
        } else {
            if false {
                if true {
                    var v67 uint8 = (v5 + v5)
                    var v68 float32 = 3.14159
                    v4 -= v4
                    _ = v67
                    _ = v68
                }
                var v69 int32 = v4
                _ = v69
            } else {
                var v70 uint64 = uint64(v5)
                _ = v70
            }
            v1 -= (((1.5 / v1) / 1.5) / v1)
            v63 = v63
        }
        _ = fi50
        _ = v63
    }
    var v71 int16 = 42
    fmt.Println(v4, v2)
    var v72 uint32 = 2147483647
    for fi73 := 0; fi73 < 11; fi73++ {
        _ = len(v31)
        var v74 uint = 818
        _ = fi73
        _ = v74
    }
    fmt.Println(v4, v1, v2)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v4:", v4)
    fmt.Println("v49:", v49)
    fmt.Println("v5:", v5)
    fmt.Println("v71:", v71)
    fmt.Println("v72:", v72)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v49
    _ = v5
    _ = v71
    _ = v72
}
