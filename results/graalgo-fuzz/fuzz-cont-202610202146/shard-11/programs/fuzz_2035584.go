package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int }
type GenS2 struct { F0 uint8; F1 uint64; F2 string }
type GenS3 struct { F0 int8; F1 int8; F2 int16 }
type GenS4 struct { F0 int64; F1 int8; F2 int64 }
type GenS5 struct { F0 int }
type GenS6 struct { F0 uint32; F1 uint8 }
func genFn0(p0 int16) string {
    v6 := GenS1{F0: true, F1: -2147483648}
    p0 += 32767
    _ = p0
    _ = v6
    return "foo bar"
}

func main() {
    v1 := make(map[string]bool)
    v2 := []int{}
    v3 := make(map[uint32]uint32)
    var v4 int16 = 786
    var v5 uint16 = 939
    for fi7 := 0; fi7 < 14; fi7++ {
        v4 -= v4
        for fi8 := 0; fi8 < 14; fi8++ {
            v4 = v4
            _ = fi8
        }
        _ = fi7
    }
    var v9 uint8 = 128
    if true {
        v4 = (v4 << ((2) & 63))
        v5 = (v5 - v5)
        var v10 uint64 = 221
        _ = v10
    } else {
        for fi11 := 0; fi11 < 17; fi11++ {
            var v12 int16 = v4
            for fi13 := 0; fi13 < 19; fi13++ {
                v4 = v12
                v3[363] = 1
                _ = fi13
            }
            _ = fi11
            _ = v12
        }
    }
    v3[0] = 42
    v4 -= (v4 & 0)
    for fi14 := 0; fi14 < 2; fi14++ {
        v9 = 42
        switch (-(fi14)) {
        case 792:
            switch fi14 {
            case 2147483647:
                var v15 uint8 = (v9 << ((65535) & 63))
                _ = v15
            case 4294967295:
                switch "hello" {
                case "":
                    for fi16 := 0; fi16 < 19; fi16++ {
                        v17 := GenS1{F0: true, F1: (-(fi14))}
                        var v18 float32 = float32(((1.5 - 3.14159) / float64(1.5)))
                        _ = fi16
                        _ = v17
                        _ = v18
                    }
                }
            case 838:
                if false {
                    v19 := GenS2{F0: v9, F1: 18446744073709551615, F2: "\n\t"}
                    var v20 int64 = 606
                    if (276 == 298) {
                        v21 := GenS1{F0: (!true), F1: fi14}
                        for fi22 := 0; fi22 < 8; fi22++ {
                            switch (fi22 & -629) {
                            case 9223372036854775807:
                                if false {
                                    v9 = 127
                                    var v23 uint8 = v9
                                    _ = v23
                                } else {
                                    var v24 uint64 = 375
                                    _ = v24
                                }
                            case -261:
                                switch "foo bar" {
                                case "\n\t":
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                case "foo bar":
                                    var v25 string = ""
                                    _ = v25
                                case "":
                                    switch (-(fi14)) {
                                    case 9223372036854775807:
                                        if len(v2) > 0 {
                                            v2[0] = (^(fi14))
                                        }
                                    default:
                                        v26 := GenS1{F0: (1 != (^(v4))), F1: fi22}
                                        _ = v26
                                    }
                                    var v27 string = "a"
                                    fmt.Println(v20, v5)
                                    _ = v27
                                }
                            default:
                                v1["\n\t"] = true
                            }
                            _ = fi22
                        }
                        _ = v21
                    } else {
                        for fi28 := 0; fi28 < 11; fi28++ {
                            v9 = (v9 * 255)
                            if false {
                                var v29 float32 = -2.25
                                _ = v29
                            } else {
                                v30 := GenS2{F0: 0, F1: 0, F2: "\n\t"}
                                v5 = v5
                                if (270 >= 1) {
                                    if len(v2) > 0 {
                                        v2[0] = fi14
                                    }
                                    var v31 int64 = (-(v20))
                                    _ = v31
                                } else {
                                    fmt.Println(v5, fi28)
                                    if len(v2) > 0 {
                                        v2[0] = (fi28 * (^(fi28)))
                                    }
                                    var v32 bool = (42 == 286)
                                    var v33 int64 = (v20 % (((v20 + -970)) | 1))
                                    _ = v32
                                    _ = v33
                                }
                                _ = v30
                            }
                            v34 := GenS3{F0: int8(v4), F1: 90, F2: 1000}
                            if len(v2) > 0 {
                                v2[0] = fi14
                            }
                            fmt.Println(v20)
                            if (!(18446744073709551615 <= 18446744073709551615)) {
                                v4 = (v4 >> ((uint(v9)) & 63))
                            } else {
                                var v35 uint = 4294967295
                                if (0.1 <= (-2.25 * (0.0 / 1e10))) {
                                    var v36 int = fi28
                                    _ = v36
                                } else {
                                    fmt.Println(v4, v4, v20)
                                    if true {
                                        var v37 int16 = -32768
                                        if len(v2) > 0 {
                                            v2[0] = fi14
                                        }
                                        _ = v37
                                    } else {
                                        switch fi28 {
                                        case 957:
                                            var v38 string = ("hello" + "a")
                                            _ = v38
                                        }
                                        v39 := GenS3{F0: 100, F1: 0, F2: (v4 & v4)}
                                        v40 := GenS4{F0: (v20 * 0), F1: 100, F2: -266}
                                        _ = v39
                                        _ = v40
                                    }
                                }
                                var v41 uint = v35
                                if false {
                                    v5 += (v5 * 32767)
                                    v42 := v1["\n\t"]
                                    v42 = (!v42)
                                    _ = v42
                                }
                                _ = v35
                                _ = v41
                            }
                            _ = fi28
                            _ = v34
                        }
                    }
                    _ = v19
                    _ = v20
                }
            default:
                for fi43 := 0; fi43 < 16; fi43++ {
                    _ = len(v3)
                    v2 = append(v2, (fi43 - (fi14 << ((18446744073709551615) & 63))))
                    _ = fi43
                }
            }
        default:
            for fi44 := 0; fi44 < 3; fi44++ {
                var v45 float32 = -2.25
                v5 += v5
                _ = fi44
                _ = v45
            }
        }
        fmt.Println(fi14, v4)
        _ = fi14
    }
    var v46 uint32 = 710
    v5 += 0
    if false {
        for fi47 := 0; fi47 < 3; fi47++ {
            var v48 int32 = -1
            var v49 float32 = 1.5
            _ = fi47
            _ = v48
            _ = v49
        }
        switch 42 {
        case 492:
            if true {
                var v50 int64 = -850
                _ = v50
            } else {
                _ = len(v2)
            }
        default:
            for fi51 := 0; fi51 < 5; fi51++ {
                for fi52 := 0; fi52 < 19; fi52++ {
                    _ = len(v1)
                    fmt.Println(fi51, v5)
                    switch (fi52 % ((fi52) | 1)) {
                    case -824:
                        var v53 int16 = v4
                        switch 128 {
                        case -471:
                            var v54 float64 = float64(1e10)
                            _ = v54
                        case 42:
                            var v55 int64 = 9223372036854775807
                            _ = v55
                        default:
                            v56 := GenS3{F0: 100, F1: 1, F2: (v4 >> ((46) & 63))}
                            _ = v56
                        }
                        _ = v53
                    }
                    _ = fi52
                }
                v46 = (v46 | (v46 - (v46 >> ((2) & 63))))
                _ = fi51
            }
        }
    }
    _ = len(v2)
    for fi57 := 0; fi57 < 3; fi57++ {
        if len(v2) > 0 {
            v2[0] = (fi57 % (((fi57 - fi57)) | 1))
        }
        v58 := GenS4{F0: 42, F1: 127, F2: 470}
        v46 = (v46 - 65535)
        _ = fi57
        _ = v58
    }
    fmt.Println(v46)
    for fi59 := 0; fi59 < 2; fi59++ {
        if false {
            switch "\n\t" {
            case "foo bar":
                if ((^(v46)) >= (v46 % (((v46 >> ((18446744073709551615) & 63))) | 1))) {
                    v4 = v4
                    if false {
                        v60 := GenS4{F0: 42, F1: -62, F2: -952}
                        if false {
                            for fi61 := 0; fi61 < 4; fi61++ {
                                v62 := GenS2{F0: v9, F1: 4294967295, F2: "foo bar"}
                                _ = fi61
                                _ = v62
                            }
                            v5 = (v5 / (((v5 % ((2) | 1))) | 1))
                        }
                        _ = v60
                    } else {
                        var v63 string = ("a" + ("foo bar" + "foo bar"))
                        for fi64 := 0; fi64 < 18; fi64++ {
                            v65 := GenS2{F0: 1, F1: 842, F2: v63}
                            _ = fi64
                            _ = v65
                        }
                        var v66 int16 = 1000
                        _ = v63
                        _ = v66
                    }
                }
            case "hello":
                v67 := v3[(v46 * v46)]
                _ = v67
            default:
                v4 = (v4 - -32768)
            }
            v68 := GenS5{F0: (-(fi59))}
            if ((true && (42 > -25)) || ((!true) || (!((((3.14159 + (0.0 / (3.14159 / -2.25))) - (-2.25 / 0.1)) - 3.14159) <= (1.5 + -2.25))))) {
                v69 := GenS6{F0: v46, F1: (^(v9))}
                _ = v69
            }
            _ = v68
        }
        _ = fi59
    }
    fmt.Println(v4, v4)
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    if false {
        var v70 float64 = -2.25
        switch "a" {
        case "hello":
            v71 := GenS3{F0: 0, F1: 1, F2: (v4 | (v4 >> ((185) & 63)))}
            _ = v71
        }
        if (!(true || false)) {
            _ = len(v3)
            v4 -= -5
            var v72 bool = (((v70 + float64(-2.25)) - (v70 * ((v70 + v70) - float64(3.14159)))) != v70)
            _ = v72
        }
        _ = v70
    }
    var v73 uint64 = 2
    v5 = (v5 - (v5 & (v5 % ((v5) | 1))))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    fmt.Println("v73:", v73)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v73
    _ = v9
}
