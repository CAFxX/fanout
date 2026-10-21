package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int }
type GenS2 struct { F0 uint32 }
type GenS3 struct { F0 string; F1 uint }
type GenS4 struct { F0 uint }
type GenS5 struct { F0 uint32; F1 string }
type GenS6 struct { F0 int; F1 uint16 }
func genFn0() int16 {
    for fi6 := 0; fi6 < 15; fi6++ {
        v7 := []int{(fi6 & (fi6 - (fi6 >> ((18446744073709551615) & 63)))), fi6}
        var v8 uint32 = 2147483648
        v9 := GenS1{F0: -1, F1: (^(fi6))}
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    var v10 uint32 = 2
    _ = v10
    return -32768
}
func genFn1(p0 int32, p1 uint, p2 uint64) uint64 {
    switch "foo bar" {
    case "foo bar":
        var v11 uint16 = 42
        _ = v11
    default:
        fmt.Println(p1, p1)
    }
    var v12 uint16 = 0
    if true {
        var v13 int32 = -740
        switch "hello" {
        case "hello":
            p2 += (p2 * p2)
        default:
            v14 := []int{383, 128, -2147483648, -656}
            _ = v14
        }
        v15 := []int{32767}
        _ = v13
        _ = v15
    }
    fmt.Println(v12, p0, v12)
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 2
}

func main() {
    var v1 bool = (65535 >= 2)
    var v2 int = 282
    var v3 uint8 = 127
    v4 := make(map[uint32]string)
    var v5 uint32 = 602
    v1 = true
    var v16 string = ""
    v17 := GenS2{F0: v5}
    var v18 int32 = 1
    if v1 {
        for fi19 := 0; fi19 < 20; fi19++ {
            v1 = (-1 != -1)
            v20 := GenS3{F0: (((v16 + v16) + v16) + v16), F1: 65535}
            if false {
                fmt.Println(v5, v2, v2)
                v3 -= v3
            } else {
                fmt.Println(v16, v16, v3)
                v21 := GenS2{F0: (v5 ^ v5)}
                for fi22 := 0; fi22 < 8; fi22++ {
                    v4[v5] = v16
                    var v23 float64 = 3.14159
                    _ = fi22
                    _ = v23
                }
                _ = v21
            }
            _ = fi19
            _ = v20
        }
        var v24 int32 = -185
        _ = v24
    }
    v25 := []int{32768, -478, v2}
    fmt.Println(v1, v2, v5)
    var v26 bool = v1
    _ = len(v4)
    if v26 {
        if false {
            switch ((v16 + (((v16 + "foo bar") + "foo bar") + v16)) + (((((v16 + "") + v16) + v16) + (v16 + (("foo bar" + v16) + "\n\t"))) + v16)) {
            case "foo bar":
                v1 = v26
            case "hello":
                if (-375 > (v2 * 32767)) {
                    for fi27 := 0; fi27 < 14; fi27++ {
                        var v28 uint32 = (v5 + v5)
                        v16 = (v16 + "foo bar")
                        var v29 int16 = 32767
                        _ = fi27
                        _ = v28
                        _ = v29
                    }
                    v30 := GenS1{F0: 789, F1: 32767}
                    _ = v30
                }
            case "\n\t":
                fmt.Println(v1)
            }
            switch v16 {
            case "":
                fmt.Println(v2)
            case "hello":
                var v31 uint16 = 382
                _ = v31
            case "a":
                v32 := GenS1{F0: 42, F1: (v2 | v2)}
                _ = v32
            default:
                switch (v2 >> ((18446744073709551615) & 63)) {
                case 689:
                    v2 -= (v2 - v2)
                }
            }
            var v33 float64 = (-2.25 - 3.14159)
            _ = v33
        } else {
            for fi34 := 0; fi34 < 16; fi34++ {
                if v1 {
                    if v26 {
                        var v35 int8 = 26
                        var v36 uint = 0
                        _ = v35
                        _ = v36
                    }
                    for fi37 := 0; fi37 < 9; fi37++ {
                        v18 = (v18 * (v18 << ((uint(v3)) & 63)))
                        _ = len(v25)
                        _ = fi37
                    }
                    v38 := GenS3{F0: (v16 + ((("\n\t" + v16) + v16) + ("a" + (("\n\t" + v16) + v16)))), F1: 255}
                    _ = v38
                }
                _ = fi34
            }
        }
        if len(v25) > 0 {
            v25[0] = 0
        }
        _ = len(v4)
    }
    v4[v5] = (v16 + "hello")
    if (v2 >= int(v3)) {
        for fi39 := 0; fi39 < 13; fi39++ {
            var v40 bool = v26
            _ = fi39
            _ = v40
        }
        v1 = (0 != 1)
        if (((true && v1) && (uint(v3) < 826)) && (984 > 255)) {
            fmt.Println(v26)
            v41 := GenS2{F0: (v5 / (((v5 % ((v5) | 1))) | 1))}
            _ = v41
        } else {
            var v42 float32 = 0.0
            v4[(v5 ^ uint32(v3))] = (v16 + v16)
            for fi43 := 0; fi43 < 16; fi43++ {
                fmt.Println(v5)
                _ = fi43
            }
            _ = v42
        }
    }
    fmt.Println(v2, v16)
    v18 -= (v18 % ((v18) | 1))
    var v44 uint16 = 2
    if v26 {
        fmt.Println(v44)
        if ((18446744073709551615 < 18446744073709551615) || true) {
            var v45 int32 = (^(v18))
            _ = v45
        }
        if len(v25) > 0 {
            v25[0] = (v2 & (v2 / ((4294967295) | 1)))
        }
    } else {
        v46 := GenS4{F0: 2}
        for fi47 := 0; fi47 < 2; fi47++ {
            var v48 string = v16
            if ((125 == 42) && v1) {
                if v26 {
                    if (float64(1e10) >= 1.5) {
                        var v49 float64 = 1.5
                        v25 = append(v25, (v2 % ((213) | 1)))
                        _ = v49
                    }
                    v44 -= (v44 | 233)
                    var v50 int8 = 42
                    _ = v50
                }
                var v51 int = 42
                delete(v4, v5)
                _ = v51
            }
            v3 = (v3 % ((v3) | 1))
            _ = fi47
            _ = v48
        }
        _ = v46
    }
    for fi52 := 0; fi52 < 18; fi52++ {
        switch (fi52 * 127) {
        case 457:
            var v53 uint64 = uint64(v5)
            _ = v53
        case -747:
            if false {
                switch v16 {
                case "a":
                    for fi54 := 0; fi54 < 5; fi54++ {
                        _ = len(v4)
                        if true {
                            v4[1] = v16
                            fmt.Println(fi52)
                        } else {
                            v25 = append(v25, (-(v2)))
                            fmt.Println(v26, v18, v26)
                            switch v16 {
                            case "foo bar":
                                delete(v4, v5)
                            case "a":
                                if ((((v18 | (v18 / (((v18 * (v18 + (v18 + v18)))) | 1))) >= (v18 >> ((431) & 63))) || v1) && (v18 == -1)) {
                                    var v55 float64 = -2.25
                                    v16 = (v16 + v16)
                                    _ = v55
                                }
                            case "hello":
                                v5 = 285
                            default:
                                v18 -= 0
                            }
                        }
                        fmt.Println(fi54, v1)
                        _ = fi54
                    }
                case "hello":
                    fmt.Println(v26)
                case "":
                    if len(v25) > 0 {
                        v25[0] = -9223372036854775808
                    }
                default:
                    v56 := GenS5{F0: v5, F1: ((v16 + "a") + v16)}
                    _ = v56
                }
                v57 := GenS2{F0: (v5 + uint32(v3))}
                if (4294967295 <= 4294967295) {
                    var v58 uint64 = 324
                    var v59 float64 = float64(3.14159)
                    v5 = (v5 / ((v5) | 1))
                    _ = v58
                    _ = v59
                } else {
                    if true {
                        v60 := GenS4{F0: 0}
                        v61 := GenS3{F0: v16, F1: 0}
                        v3 = (v3 ^ v3)
                        _ = v60
                        _ = v61
                    }
                }
                _ = v57
            }
        default:
            if false {
                v5 = 963
                fmt.Println(v18, v18)
            } else {
                switch v2 {
                case 128:
                    if v1 {
                        var v62 int64 = 517
                        _ = v62
                    } else {
                        for fi63 := 0; fi63 < 15; fi63++ {
                            fmt.Println(v2, v1)
                            if false {
                                v64 := v4[v5]
                                _ = v64
                            }
                            v65 := GenS4{F0: 255}
                            _ = fi63
                            _ = v65
                        }
                    }
                case 9223372036854775807:
                    v66 := GenS6{F0: fi52, F1: (v44 << ((uint(v5)) & 63))}
                    _ = v66
                default:
                    var v67 uint64 = 42
                    _ = v67
                }
                if v1 {
                    _ = len(v4)
                    v5 -= (v5 | v5)
                } else {
                    if (v3 <= 0) {
                        v16 = ""
                    }
                    for fi68 := 0; fi68 < 7; fi68++ {
                        v69 := GenS2{F0: v5}
                        v4[(v5 << ((2) & 63))] = v16
                        _ = fi68
                        _ = v69
                    }
                    if ((339 != 2) || (-54 == -554)) {
                        switch "a" {
                        case "foo bar":
                            fmt.Println(v3, fi52)
                        case "hello":
                            if v26 {
                                v70 := GenS5{F0: (^(v5)), F1: ((v16 + v16) + "")}
                                _ = v70
                            }
                        case "":
                            var v71 uint8 = v3
                            _ = v71
                        }
                    }
                }
                switch v16 {
                case "foo bar":
                    fmt.Println(v5, v44, fi52)
                case "a":
                    var v72 int16 = -32768
                    _ = v72
                case "":
                    for fi73 := 0; fi73 < 11; fi73++ {
                        v4[(v5 | v5)] = v16
                        v3 = v3
                        _ = fi73
                    }
                default:
                    if false {
                        v74 := GenS4{F0: 4294967295}
                        v2 -= (fi52 ^ (v2 * fi52))
                        v2 += -601
                        _ = v74
                    }
                }
            }
        }
        var v75 uint16 = v44
        v3 -= 255
        _ = fi52
        _ = v75
    }
    v76 := GenS4{F0: 65535}
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v25:", len(v25), v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v44:", v44)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v44
    _ = v5
    _ = v76
}
