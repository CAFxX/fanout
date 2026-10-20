package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 uint32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int8; F1 uint; F2 int32 }
type GenS4 struct { F0 int32; F1 uint8; F2 uint }
type GenS5 struct { F0 uint; F1 uint16; F2 uint32 }
type GenS6 struct { F0 uint16; F1 string }
type GenS7 struct { F0 int8; F1 uint; F2 uint16 }
type GenS8 struct { F0 uint8 }
type GenS9 struct { F0 uint }
type GenS10 struct { F0 string }
func genFn0() string {
    var v7 int8 = 0
    var v8 float32 = 0.1
    _ = v7
    _ = v8
    return ("hello" + ("foo bar" + (("foo bar" + (("a" + ("foo bar" + "foo bar")) + "\n\t")) + "")))
}
func genFn1() string {
    var v9 int = -2147483648
    v10 := []int{}
    _ = v10
    _ = v9
    return ("foo bar" + "foo bar")
}

func main() {
    var v1 uint = 644
    v2 := make(map[int]bool)
    var v3 float32 = (1.5 + (1.5 * (0.0 + -2.25)))
    var v4 string = "hello"
    var v5 bool = ((float64((v3 / ((v3 * v3) - (((-2.25 * (float32((float64(float32(0.0)) * (1.5 + (0.1 * (float64(float32(1e10)) + float64(-2.25)))))) * (3.14159 + 0.0))) - ((1.5 * (v3 * (1.5 + v3))) * v3)) * float32(1e10))))) < 0.0) || false)
    v6 := []uint32{42, 4294967295, 42, 245}
    v11 := GenS1{F0: 0, F1: (v1 / ((v1) | 1)), F2: 2}
    var v12 uint64 = 845
    switch 2147483647 {
    case -533:
        if (0 == 192) {
            if (!(255 > 75)) {
                v1 += 175
            }
            var v13 uint64 = v12
            _ = v13
        } else {
            var v14 int16 = 42
            if true {
                if false {
                    v15 := GenS2{F0: (^(v12))}
                    switch 42 {
                    case -9223372036854775808:
                        v16 := GenS1{F0: 9223372036854775807, F1: v1, F2: 0}
                        _ = v16
                    }
                    var v17 int16 = (^(v14))
                    _ = v15
                    _ = v17
                } else {
                    v4 = v4
                    var v18 float32 = (1.5 / 1.0)
                    _ = v18
                }
                v19 := GenS2{F0: v12}
                _ = v19
            } else {
                v20 := GenS3{F0: 127, F1: (v1 * 631), F2: -1}
                if false {
                    for fi21 := 0; fi21 < 8; fi21++ {
                        switch 322 {
                        case 42:
                            switch (-(fi21)) {
                            case 127:
                                v2[(fi21 * (fi21 & -768))] = (!true)
                            case -1:
                                for fi22 := 0; fi22 < 6; fi22++ {
                                    v23 := GenS3{F0: 43, F1: 1, F2: 1}
                                    _ = fi22
                                    _ = v23
                                }
                            case 32768:
                                fmt.Println(v5, v3)
                            }
                        case -215:
                            v12 = (v12 % (((v12 << (((v1 << ((v1) & 63))) & 63))) | 1))
                        case -388:
                            v14 = v14
                        }
                        var v24 bool = (v5 && v5)
                        _ = fi21
                        _ = v24
                    }
                    if v5 {
                        switch ("\n\t" + v4) {
                        case "hello":
                            if v5 {
                                _ = len(v6)
                            } else {
                                fmt.Println(v3)
                                v25 := GenS4{F0: -1, F1: 127, F2: 65535}
                                switch 9223372036854775807 {
                                case 128:
                                    var v26 float64 = -2.25
                                    _ = v26
                                case 2147483647:
                                    var v27 uint = (v1 / ((2) | 1))
                                    _ = v27
                                case 397:
                                    v3 = (v3 * (v3 / v3))
                                default:
                                    if (((v12 - v12) != v12) || (!v5)) {
                                        for fi28 := 0; fi28 < 8; fi28++ {
                                            for fi29 := 0; fi29 < 20; fi29++ {
                                                fmt.Println(v4)
                                                switch fi29 {
                                                case -156:
                                                    if v5 {
                                                        v4 = v4
                                                    } else {
                                                        var v30 bool = true
                                                        v31 := GenS5{F0: v1, F1: 2, F2: 65535}
                                                        v6 = append(v6, 2)
                                                        _ = v30
                                                        _ = v31
                                                    }
                                                case -9223372036854775808:
                                                    var v32 uint = (v1 - v1)
                                                    _ = v32
                                                default:
                                                    v33 := GenS4{F0: -957, F1: 44, F2: v1}
                                                    _ = v33
                                                }
                                                _ = fi29
                                            }
                                            _ = fi28
                                        }
                                    }
                                }
                                _ = v25
                            }
                        case "":
                            fmt.Println(v1, v14)
                        case "foo bar":
                            v34 := GenS6{F0: 2, F1: "a"}
                            _ = v34
                        }
                        v1 = 2
                        switch 237 {
                        case 618:
                            var v35 uint16 = 75
                            _ = v35
                        case 127:
                            if ((!false) && v5) {
                                var v36 uint32 = 329
                                v14 -= v14
                                _ = v36
                            }
                        }
                    }
                    switch "" {
                    case "foo bar":
                        if v5 {
                            v37 := GenS7{F0: -15, F1: v1, F2: 0}
                            _ = v37
                        }
                    default:
                        var v38 uint32 = 0
                        _ = v38
                    }
                } else {
                    var v39 int8 = -1
                    _ = v39
                }
                _ = v20
            }
            _ = v14
        }
    case 348:
        v40 := GenS8{F0: 255}
        _ = v40
    case -767:
        if v5 {
            if (float64(v3) != 0.0) {
                v41 := GenS7{F0: 100, F1: (v1 / ((1) | 1)), F2: 89}
                _ = len(v6)
                _ = v41
            } else {
                if v5 {
                    fmt.Println(v4, v5)
                }
            }
            switch 0 {
            case 591:
                var v42 float32 = 3.14159
                _ = v42
            default:
                if (v5 || (574 > v12)) {
                    if true {
                        v12 = v12
                        if (!false) {
                            if ((v3 == 0.1) || true) {
                                for fi43 := 0; fi43 < 14; fi43++ {
                                    v5 = (127 == 128)
                                    _ = fi43
                                }
                            } else {
                                for fi44 := 0; fi44 < 6; fi44++ {
                                    for fi45 := 0; fi45 < 20; fi45++ {
                                        if (0 < -9223372036854775808) {
                                            var v46 uint8 = 180
                                            _ = v46
                                        } else {
                                            var v47 int = fi44
                                            var v48 int8 = 105
                                            _ = len(v6)
                                            _ = len(v2)
                                            switch 547 {
                                            case 2147483647:
                                                var v49 float32 = (v3 + 0.1)
                                                _ = v49
                                            case 112:
                                                if v5 {
                                                    if len(v6) >= 2 {
                                                        v6 = v6[1:len(v6)]
                                                    }
                                                    fmt.Println(v5, v4)
                                                    v5 = (42 <= 0)
                                                } else {
                                                    var v50 uint8 = 127
                                                    switch v4 {
                                                    case "foo bar":
                                                        var v51 int32 = 42
                                                        _ = v51
                                                    case "a":
                                                        switch (fi44 / (((v47 / (((fi44 % ((v47) | 1))) | 1))) | 1)) {
                                                        case 1:
                                                            var v52 float64 = -2.25
                                                            _ = v52
                                                        case 4294967295:
                                                            var v53 int8 = 100
                                                            _ = v53
                                                        }
                                                    default:
                                                        if (0 > 756) {
                                                            if (v5 || v5) {
                                                                var v54 bool = v5
                                                                _ = v54
                                                            } else {
                                                                var v55 uint = 446
                                                                fmt.Println(v5)
                                                                v3 = v3
                                                                _ = v55
                                                            }
                                                            v50 += (^(v50))
                                                        }
                                                    }
                                                    _ = v50
                                                }
                                            case 128:
                                                v56 := GenS1{F0: 1, F1: (v1 | (v1 * (v1 % ((v1) | 1)))), F2: 42}
                                                v57 := GenS7{F0: (v48 & 1), F1: 0, F2: 2}
                                                if len(v6) >= 2 {
                                                    v6 = v6[1:len(v6)]
                                                }
                                                v3 += ((v3 + v3) + float32(1e10))
                                                _ = v56
                                                _ = v57
                                            default:
                                                v58 := GenS3{F0: v48, F1: (v1 / ((v1) | 1)), F2: -1}
                                                _ = v58
                                            }
                                            var v59 uint32 = 140
                                            _ = v47
                                            _ = v48
                                            _ = v59
                                        }
                                        _ = fi45
                                    }
                                    _ = fi44
                                }
                                v1 = (v1 << ((v1) & 63))
                            }
                            v4 = (v4 + ("hello" + ("\n\t" + (v4 + "\n\t"))))
                            v1 = v1
                        }
                    } else {
                        switch "foo bar" {
                        case "hello":
                            if true {
                                for fi60 := 0; fi60 < 2; fi60++ {
                                    fmt.Println(v3, v4)
                                    v4 = v4
                                    _ = fi60
                                }
                            } else {
                                var v61 uint = 551
                                _ = v61
                            }
                        case "":
                            var v62 uint64 = (v12 ^ v12)
                            _ = v62
                        case "foo bar":
                            var v63 float64 = 0.1
                            _ = v63
                        }
                        var v64 uint = v1
                        _ = v64
                    }
                }
            }
            _ = len(v6)
        } else {
            v12 -= 621
            switch (("" + v4) + v4) {
            case "":
                var v65 int32 = 42
                _ = v65
            case "\n\t":
                for fi66 := 0; fi66 < 5; fi66++ {
                    v12 = (v12 | (^(v12)))
                    _ = fi66
                }
            case "foo bar":
                for fi67 := 0; fi67 < 5; fi67++ {
                    v5 = v5
                    v5 = (v1 > 535)
                    _ = fi67
                }
            default:
                var v68 string = v4
                _ = v68
            }
        }
    }
    if true {
        if ((!false) && ((1e10 + 0.0) >= float64((float32(float64(v3)) + (v3 * v3))))) {
            var v69 int = -194
            _ = v69
        } else {
            if ((42 > -648) && ((42 < -96) || false)) {
                v1 = (v1 & 4294967295)
            }
            var v70 uint = (v1 - (v1 ^ v1))
            switch 0 {
            case 2147483647:
                v1 = 734
            case -2147483648:
                v71 := GenS9{F0: v1}
                _ = v71
            case 127:
                v4 = v4
            }
            _ = v70
        }
        fmt.Println(v3, v3)
    } else {
        v72 := GenS10{F0: (v4 + "\n\t")}
        _ = v72
    }
    v1 = (v1 & v1)
    v2[802] = (v3 == v3)
    var v73 int64 = 42
    switch 9223372036854775807 {
    case 127:
        fmt.Println(v12, v1)
    case 929:
        if (!false) {
            var v74 bool = (v73 < v73)
            _ = v74
        }
    }
    for fi75 := 0; fi75 < 2; fi75++ {
        v6 = append(v6, 756)
        _ = fi75
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    fmt.Println("v73:", v73)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v73
}
