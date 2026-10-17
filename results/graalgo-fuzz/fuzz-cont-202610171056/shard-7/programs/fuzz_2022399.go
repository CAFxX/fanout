package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32; F1 uint }
type GenS4 struct { F0 string }
type GenS5 struct { F0 uint8; F1 int8 }
type GenS6 struct { F0 int16; F1 int32; F2 uint16 }
type GenS7 struct { F0 uint; F1 int8 }
type GenS8 struct { F0 int64; F1 int64 }
func genFn0(p0 string, p1 uint32, p2 int16) int64 {
    p1 += 65535
    p2 += p2
    var v5 int64 = 667
    v6 := []int{687, 2147483647, 2147483647}
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return (v5 % ((1) | 1))
}

func main() {
    v1 := []uint32{68, 4294967295}
    var v2 uint = 0
    var v3 float32 = ((-2.25 / 1.5) - 3.14159)
    v4 := make(map[uint32]bool)
    for fi7 := 0; fi7 < 8; fi7++ {
        var v8 int64 = 1
        if (32767 <= 36) {
            var v9 uint64 = 592
            v1 = append(v1, 42)
            _ = v9
        }
        _ = fi7
        _ = v8
    }
    for fi10 := 0; fi10 < 17; fi10++ {
        v11 := GenS1{F0: (fi10 << (((v2 * 1)) & 63)), F1: 248, F2: 18446744073709551615}
        if (true || ((float64((v3 / v3)) + 1.5) >= ((1e10 * 0.1) + -2.25))) {
            _ = len(v1)
            var v12 uint16 = 59
            _ = v12
        }
        _ = fi10
        _ = v11
    }
    v13 := GenS2{F0: 32767}
    fmt.Println(v3, v3)
    for fi14 := 0; fi14 < 8; fi14++ {
        v4[467] = false
        _ = fi14
    }
    fmt.Println(v2, v3)
    if false {
        _ = len(v1)
        switch "\n\t" {
        case "foo bar":
            if ((^(v2)) <= v2) {
                var v15 uint16 = 42
                if false {
                    var v16 int16 = 1000
                    var v17 string = "foo bar"
                    _ = v16
                    _ = v17
                }
                _ = v15
            }
        case "a":
            for fi18 := 0; fi18 < 1; fi18++ {
                v3 += v3
                _ = fi18
            }
        case "hello":
            switch "hello" {
            case "a":
                for fi19 := 0; fi19 < 18; fi19++ {
                    if false {
                        v20 := GenS3{F0: 1, F1: v2}
                        _ = v20
                    }
                    v3 = v3
                    _ = fi19
                }
            case "\n\t":
                v3 = 0.1
            case "foo bar":
                if len(v1) >= 2 {
                    v1 = v1[1:len(v1)]
                }
            default:
                if (!true) {
                    v2 -= 4294967295
                } else {
                    if (-2147483648 != 4294967295) {
                        fmt.Println(v2, v3)
                    }
                }
            }
        default:
            for fi21 := 0; fi21 < 7; fi21++ {
                for fi22 := 0; fi22 < 10; fi22++ {
                    var v23 int32 = 2147483647
                    _ = fi22
                    _ = v23
                }
                var v24 uint16 = 2
                for fi25 := 0; fi25 < 13; fi25++ {
                    if (((^(v24)) <= 29) || true) {
                        v26 := GenS2{F0: 42}
                        v27 := GenS4{F0: (("" + "\n\t") + (("" + "\n\t") + ("foo bar" + "")))}
                        _ = v26
                        _ = v27
                    } else {
                        v24 += (^(v24))
                        v3 += float32(3.14159)
                        v28 := v4[321]
                        _ = v28
                    }
                    v4[1] = (!true)
                    _ = fi25
                }
                _ = fi21
                _ = v24
            }
        }
        var v29 uint32 = 0
        _ = v29
    } else {
        v30 := GenS2{F0: 0}
        if true {
            for fi31 := 0; fi31 < 16; fi31++ {
                for fi32 := 0; fi32 < 2; fi32++ {
                    fmt.Println(fi31)
                    switch "hello" {
                    case "a":
                        v33 := v4[1]
                        _ = v33
                    }
                    _ = fi32
                }
                _ = fi31
            }
        }
        _ = v30
    }
    for fi34 := 0; fi34 < 7; fi34++ {
        switch ("foo bar" + ("\n\t" + "hello")) {
        case "a":
            v35 := GenS5{F0: 191, F1: 42}
            _ = v35
        case "foo bar":
            v36 := GenS6{F0: 1000, F1: -2147483648, F2: 386}
            _ = v36
        case "":
            fmt.Println(v3)
        }
        v37 := GenS6{F0: 1000, F1: 2147483647, F2: 65535}
        _ = fi34
        _ = v37
    }
    _ = len(v1)
    if (1.5 <= float64(1.5)) {
        var v38 bool = ((42 < -1) || true)
        _ = v38
    } else {
        if (18446744073709551615 >= 2) {
            var v39 bool = ("" <= "")
            _ = v39
        }
        v2 = 153
        if (false || ((^(v2)) == 1)) {
            switch 696 {
            case 127:
                if (-1 < -128) {
                    var v40 uint16 = 32767
                    if true {
                        var v41 int16 = -32768
                        var v42 int16 = (v41 - (v41 << ((0) & 63)))
                        v43 := GenS7{F0: 0, F1: 127}
                        _ = v41
                        _ = v42
                        _ = v43
                    } else {
                        v2 = (v2 ^ 65535)
                    }
                    _ = v40
                } else {
                    for fi44 := 0; fi44 < 11; fi44++ {
                        v2 = (v2 ^ 255)
                        _ = fi44
                    }
                    delete(v4, 1)
                }
            case 128:
                var v45 float32 = v3
                _ = v45
            case -846:
                var v46 uint64 = 4294967295
                _ = v46
            default:
                switch "\n\t" {
                case "":
                    if len(v1) > 0 {
                        v1[0] = 42
                    }
                case "\n\t":
                    switch -446 {
                    case 32768:
                        var v47 string = "foo bar"
                        _ = v47
                    case 0:
                        for fi48 := 0; fi48 < 13; fi48++ {
                            for fi49 := 0; fi49 < 12; fi49++ {
                                if (v3 == 1.5) {
                                    var v50 uint32 = 2147483648
                                    if true {
                                        v2 = v2
                                        if ((fi49 % ((2147483647) | 1)) == 9223372036854775807) {
                                            if false {
                                                switch "\n\t" {
                                                case "foo bar":
                                                    if false {
                                                        var v51 int16 = -65
                                                        _ = v51
                                                    }
                                                case "a":
                                                    var v52 int32 = -611
                                                    _ = v52
                                                }
                                            } else {
                                                v50 -= v50
                                                v2 += v2
                                                fmt.Println(fi49, fi48, fi49)
                                            }
                                            var v53 bool = (v3 < 0.1)
                                            _ = v53
                                        }
                                        switch (fi48 & fi48) {
                                        case -925:
                                            v54 := v4[(v50 + 4294967295)]
                                            _ = v54
                                        default:
                                            fmt.Println(v2)
                                        }
                                        var v55 float32 = ((float32(float64(-2.25)) * 3.14159) / 1.0)
                                        _ = v55
                                    }
                                    fmt.Println(v3)
                                    _ = v50
                                }
                                delete(v4, 4294967295)
                                _ = fi49
                            }
                            v2 = v2
                            v3 += (((-2.25 - v3) - float32(0.0)) / (0.1 / 1e10))
                            _ = fi48
                        }
                    case -1:
                        v3 = (v3 / v3)
                    }
                case "foo bar":
                    v2 -= (v2 ^ 4294967295)
                default:
                    v2 -= (v2 % ((42) | 1))
                }
            }
            v3 = float32((-2.25 + (float64(v3) * -2.25)))
            fmt.Println(v3, v3)
        } else {
            v2 -= 850
            v3 = v3
            var v56 int8 = 31
            _ = v56
        }
    }
    v2 += (^(v2))
    for fi57 := 0; fi57 < 18; fi57++ {
        if (42 <= 42) {
            _ = len(v4)
        } else {
            v58 := GenS4{F0: "hello"}
            _ = v58
        }
        _ = fi57
    }
    v59 := GenS4{F0: "foo bar"}
    v1 = append(v1, 939)
    switch "hello" {
    case "foo bar":
        v60 := v4[75]
        _ = v60
    case "\n\t":
        for fi61 := 0; fi61 < 7; fi61++ {
            if true {
                var v62 uint8 = 1
                v63 := GenS8{F0: 9223372036854775807, F1: 0}
                fmt.Println(v3, v3, v62)
                _ = v62
                _ = v63
            }
            _ = fi61
        }
    default:
        switch "a" {
        case "\n\t":
            v2 = v2
        case "hello":
            for fi64 := 0; fi64 < 19; fi64++ {
                var v65 uint64 = 332
                fmt.Println(v65, v65, v2)
                v66 := GenS8{F0: 449, F1: 161}
                _ = fi64
                _ = v65
                _ = v66
            }
        }
    }
    var v67 uint = (v2 + (v2 & (v2 / (((^(v2))) | 1))))
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v67:", v67)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v59
    _ = v67
}
