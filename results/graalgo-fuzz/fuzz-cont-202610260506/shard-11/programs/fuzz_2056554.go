package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint16 }
type GenS3 struct { F0 int32; F1 int64; F2 string }
type GenS4 struct { F0 uint16; F1 uint; F2 int8 }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 int32; F1 uint; F2 uint }
type GenS7 struct { F0 uint16 }
type GenS8 struct { F0 uint64; F1 int32 }
type GenS9 struct { F0 uint16; F1 int }
func genFn0(p0 uint, p1 string, p2 int32) int {
    v5 := []int{2147483647, -151}
    switch -584 {
    case 942:
        var v6 uint8 = 2
        _ = v6
    case 374:
        if true {
            var v7 int8 = -1
            _ = v7
        }
    case -9223372036854775808:
        p2 += (p2 - 42)
    }
    if (!false) {
        if true {
            for fi8 := 0; fi8 < 4; fi8++ {
                for fi9 := 0; fi9 < 14; fi9++ {
                    v5 = append(v5, -2147483648)
                    v10 := GenS1{F0: (fi8 | fi9)}
                    _ = fi9
                    _ = v10
                }
                _ = fi8
            }
            var v11 uint8 = 17
            switch "\n\t" {
            case "foo bar":
                var v12 uint32 = 2147483647
                _ = v12
            case "":
                fmt.Println(p1, v11)
            case "\n\t":
                fmt.Println(p0, p2)
            }
            _ = v11
        }
        switch -2147483648 {
        case -195:
            for fi13 := 0; fi13 < 7; fi13++ {
                v14 := GenS2{F0: 65535}
                if true {
                    p0 += (p0 >> (((p0 * p0)) & 63))
                } else {
                    if true {
                        fmt.Println(p0)
                        fmt.Println(p1, p2)
                    }
                }
                _ = fi13
                _ = v14
            }
        case 4294967295:
            switch "a" {
            case "hello":
                v15 := GenS3{F0: -1, F1: -318, F2: (p1 + p1)}
                _ = v15
            case "\n\t":
                for fi16 := 0; fi16 < 14; fi16++ {
                    var v17 uint64 = 4294967295
                    _ = fi16
                    _ = v17
                }
            case "":
                var v18 int64 = 42
                _ = v18
            default:
                switch "hello" {
                case "\n\t":
                    var v19 int = 0
                    _ = v19
                case "a":
                    p1 = p1
                }
            }
        }
    } else {
        p1 = ""
    }
    p2 = p2
    p2 = (p2 >> (((p0 * (^(p0)))) & 63))
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return -933
}
func genFn1() int8 {
    v20 := []int{-209, 656, 127, 128}
    var v21 uint = 322
    for fi22 := 0; fi22 < 12; fi22++ {
        v23 := GenS4{F0: 637, F1: 4294967295, F2: 127}
        if true {
            if true {
                v21 -= v21
                v24 := GenS4{F0: 2, F1: (v21 - 65535), F2: 1}
                _ = v24
            }
            var v25 int8 = 0
            _ = v25
        }
        var v26 uint64 = 356
        _ = fi22
        _ = v23
        _ = v26
    }
    var v27 uint8 = 127
    _ = v20
    _ = v21
    _ = v27
    return -103
}

func main() {
    var v1 float32 = float32(((-2.25 + 0.1) - 1.5))
    var v2 uint16 = 32767
    v3 := make(map[int]string)
    var v4 bool = (false || true)
    for fi28 := 0; fi28 < 8; fi28++ {
        var v29 int16 = 1000
        var v30 float64 = 0.1
        for fi31 := 0; fi31 < 12; fi31++ {
            if true {
                switch "foo bar" {
                case "foo bar":
                    var v32 uint = 1
                    _ = v32
                case "hello":
                    if v4 {
                        if v4 {
                            if (false || (!false)) {
                                v33 := GenS5{F0: (v29 + v29)}
                                v30 -= v30
                                _ = v33
                            } else {
                                switch "\n\t" {
                                case "hello":
                                    v34 := GenS5{F0: v29}
                                    _ = v34
                                case "":
                                    v35 := []int{-1}
                                    _ = v35
                                case "\n\t":
                                    v36 := GenS2{F0: (v2 ^ v2)}
                                    _ = v36
                                default:
                                    v37 := GenS6{F0: -771, F1: 0, F2: 65535}
                                    _ = v37
                                }
                            }
                            v4 = true
                            if v4 {
                                v3[128] = "\n\t"
                                fmt.Println(fi31)
                                v29 -= 0
                            }
                        } else {
                            v1 -= -2.25
                            if ((65535 != 2147483647) && v4) {
                                var v38 int8 = 127
                                v39 := GenS5{F0: (v29 << ((255) & 63))}
                                _ = v38
                                _ = v39
                            } else {
                                switch genFn0(255, (("" + (("" + "") + ("" + ""))) + "a"), -415) {
                                case 2147483647:
                                    switch "\n\t" {
                                    case "foo bar":
                                        var v40 string = "hello"
                                        _ = v40
                                    case "a":
                                        v41 := []int{(fi31 | genFn0(42, "foo bar", 69))}
                                        _ = v41
                                    }
                                case 0:
                                    v30 += v30
                                default:
                                    v42 := GenS1{F0: 4294967295}
                                    _ = v42
                                }
                                var v43 int8 = genFn1()
                                v1 -= 1e10
                                _ = v43
                            }
                        }
                        v1 = -2.25
                    }
                case "":
                    if v4 {
                        var v44 uint64 = 39
                        v1 -= v1
                        v2 -= (v2 % (((v2 / ((v2) | 1))) | 1))
                        _ = v44
                    } else {
                        switch fi28 {
                        case -1:
                            var v45 bool = true
                            _ = v45
                        case 9223372036854775807:
                            v46 := []int{}
                            _ = v46
                        }
                        v47 := v3[(^(fi31))]
                        _ = v47
                    }
                default:
                    if (v1 <= v1) {
                        var v48 int32 = -1
                        _ = v48
                    } else {
                        v4 = false
                        var v49 uint64 = 0
                        _ = v49
                    }
                }
                var v50 int = genFn0(806, "foo bar", -2147483648)
                v51 := GenS7{F0: 42}
                _ = v50
                _ = v51
            }
            v1 -= 1.5
            _ = fi31
        }
        _ = fi28
        _ = v29
        _ = v30
    }
    fmt.Println(v1)
    for fi52 := 0; fi52 < 4; fi52++ {
        _ = len(v3)
        _ = fi52
    }
    v4 = v4
    _ = len(v3)
    switch 42 {
    case -9223372036854775808:
        v4 = true
    case 32768:
        var v53 uint16 = 583
        _ = v53
    case -2147483648:
        var v54 int64 = 42
        _ = v54
    }
    _ = len(v3)
    v55 := GenS3{F0: -239, F1: 0, F2: "\n\t"}
    var v56 int = 1
    for fi57 := 0; fi57 < 17; fi57++ {
        var v58 int64 = 42
        switch "" {
        case "foo bar":
            v4 = (!((32767 >= 1) || (100 != -61)))
        case "\n\t":
            var v59 int = (-(v56))
            _ = v59
        default:
            var v60 int16 = 1000
            _ = v60
        }
        _ = fi57
        _ = v58
    }
    for fi61 := 0; fi61 < 7; fi61++ {
        if false {
            switch (("foo bar" + "foo bar") + ((("\n\t" + "\n\t") + "a") + "foo bar")) {
            case "":
                fmt.Println(v2, v56)
            case "\n\t":
                if true {
                    v62 := GenS8{F0: 2, F1: -111}
                    if v4 {
                        v4 = (!((genFn1() == 0) || v4))
                        for fi63 := 0; fi63 < 11; fi63++ {
                            if v4 {
                                v64 := GenS4{F0: (v2 - 2), F1: 84, F2: 125}
                                v2 -= (v2 & (v2 << ((813) & 63)))
                                v65 := []int{}
                                _ = v64
                                _ = v65
                            }
                            _ = fi63
                        }
                    }
                    switch "\n\t" {
                    case "hello":
                        for fi66 := 0; fi66 < 19; fi66++ {
                            v67 := []int{(fi66 - (fi66 % (((v56 << ((65535) & 63))) | 1))), 32768, (fi66 * (fi61 / (((v56 + genFn0(18446744073709551615, "\n\t", 470))) | 1)))}
                            v68 := GenS6{F0: 0, F1: 1, F2: 18446744073709551615}
                            _ = fi66
                            _ = v67
                            _ = v68
                        }
                    default:
                        if false {
                            switch fi61 {
                            case 180:
                                v4 = true
                            case 923:
                                v69 := GenS3{F0: 42, F1: -1, F2: ("\n\t" + "foo bar")}
                                _ = v69
                            case 128:
                                var v70 int16 = 533
                                _ = v70
                            }
                            v71 := []int{}
                            var v72 uint32 = 1
                            _ = v71
                            _ = v72
                        } else {
                            var v73 int16 = 703
                            _ = v73
                        }
                    }
                    _ = v62
                } else {
                    var v74 int8 = 112
                    v2 += v2
                    v1 = v1
                    _ = v74
                }
            }
            var v75 int16 = -32768
            _ = v75
        } else {
            v4 = v4
        }
        _ = len(v3)
        _ = fi61
    }
    _ = len(v3)
    if v4 {
        v56 = (^(v56))
    }
    for fi76 := 0; fi76 < 1; fi76++ {
        var v77 int8 = -20
        if v4 {
            v56 += -2147483648
            v4 = ((v2 / (((v2 / ((196) | 1))) | 1)) == 966)
        }
        v78 := GenS9{F0: 982, F1: v56}
        _ = fi76
        _ = v77
        _ = v78
    }
    v79 := GenS9{F0: 1, F1: v56}
    fmt.Println(v4)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v56:", v56)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v55
    _ = v56
    _ = v79
}
