package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }
type GenS2 struct { F0 uint32; F1 uint32; F2 int32 }
type GenS3 struct { F0 uint16; F1 int64; F2 int16 }
type GenS4 struct { F0 uint8; F1 int16; F2 int64 }
type GenS5 struct { F0 uint32; F1 uint16 }
type GenS6 struct { F0 int16; F1 string; F2 uint64 }
type GenS7 struct { F0 uint64; F1 bool; F2 uint16 }
type GenS8 struct { F0 uint; F1 uint8; F2 uint }
func genFn0(p0 uint64, p1 uint32, p2 uint64) int64 {
    v7 := []int{1, -2147483648, int(p0), 2147483647}
    v8 := GenS1{F0: 18446744073709551615}
    if false {
        for fi9 := 0; fi9 < 17; fi9++ {
            p0 += 956
            _ = fi9
        }
    }
    for fi10 := 0; fi10 < 17; fi10++ {
        v11 := GenS1{F0: 706}
        p0 += 66
        p1 = p1
        _ = fi10
        _ = v11
    }
    var v12 int16 = 436
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    _ = v7
    _ = v8
    return 34
}
func genFn1(p0 uint8, p1 bool, p2 int64) int32 {
    v13 := GenS1{F0: 18446744073709551615}
    if (p1 || (!false)) {
        for fi14 := 0; fi14 < 8; fi14++ {
            v15 := GenS1{F0: 108}
            _ = fi14
            _ = v15
        }
        fmt.Println(p0, p0)
        p0 += 255
    }
    v16 := []int{}
    if (p1 || (!(1 < 2147483647))) {
        for fi17 := 0; fi17 < 8; fi17++ {
            v18 := GenS2{F0: 1, F1: 4294967295, F2: 42}
            _ = fi17
            _ = v18
        }
        fmt.Println(p2)
    } else {
        if (2 >= 832) {
            if (4294967295 == 344) {
                v19 := GenS2{F0: 955, F1: 869, F2: 2147483647}
                _ = v19
            } else {
                switch "a" {
                case "hello":
                    if (!(true && ((2 > 1) && (1 > 191)))) {
                        fmt.Println(p1)
                        p1 = p1
                        if (3.14159 == 0.1) {
                            var v20 uint64 = 2
                            if len(v16) > 0 {
                                v16[0] = 127
                            }
                            _ = v20
                        }
                    }
                case "":
                    v21 := GenS3{F0: 754, F1: -511, F2: 0}
                    _ = v21
                case "\n\t":
                    if (2 <= 2147483647) {
                        var v22 int16 = 42
                        _ = v22
                    } else {
                        if (519 == 1000) {
                            if false {
                                var v23 int64 = p2
                                if (p1 || false) {
                                    v23 += (p2 ^ (v23 ^ v23))
                                    var v24 uint64 = 42
                                    switch -626 {
                                    case -194:
                                        var v25 uint = 4294967295
                                        _ = v25
                                    case 4294967295:
                                        switch ("a" + (("foo bar" + "") + (("a" + "hello") + "\n\t"))) {
                                        case "a":
                                            switch -178 {
                                            case 4294967295:
                                                fmt.Println(p0, v23, v23)
                                            default:
                                                var v26 int32 = 25
                                                _ = v26
                                            }
                                        case "foo bar":
                                            for fi27 := 0; fi27 < 9; fi27++ {
                                                var v28 int8 = -118
                                                if p1 {
                                                    v28 = (v28 << ((168) & 63))
                                                } else {
                                                    v16 = append(v16, (fi27 / (((fi27 + (^(fi27)))) | 1)))
                                                }
                                                p0 -= (p0 | p0)
                                                _ = fi27
                                                _ = v28
                                            }
                                        case "\n\t":
                                            p1 = false
                                        }
                                    default:
                                        var v29 uint8 = 1
                                        _ = v29
                                    }
                                    _ = v24
                                } else {
                                    if len(v16) > 0 {
                                        v16[0] = -377
                                    }
                                    p0 -= (p0 | (p0 >> ((4294967295) & 63)))
                                    if len(v16) > 0 {
                                        v16[0] = 2147483647
                                    }
                                }
                                _ = v23
                            } else {
                                var v30 int16 = -264
                                _ = v30
                            }
                            for fi31 := 0; fi31 < 18; fi31++ {
                                var v32 uint32 = 2147483648
                                switch (fi31 >> ((693) & 63)) {
                                case -265:
                                    v33 := GenS2{F0: (v32 ^ 946), F1: 662, F2: -2147483648}
                                    _ = v33
                                case -1:
                                    switch (fi31 * (fi31 << ((0) & 63))) {
                                    case 9223372036854775807:
                                        v34 := GenS4{F0: (^(p0)), F1: -32768, F2: genFn0(94, (v32 - 42), 717)}
                                        _ = v34
                                    case 1:
                                        v16 = append(v16, fi31)
                                    default:
                                        p1 = (-1 <= 127)
                                    }
                                case -585:
                                    if false {
                                        fmt.Println(v32, p2, p0)
                                        if p1 {
                                            p1 = false
                                            for fi35 := 0; fi35 < 4; fi35++ {
                                                var v36 uint32 = (v32 ^ (v32 | v32))
                                                v37 := GenS4{F0: 233, F1: 1, F2: (p2 >> ((554) & 63))}
                                                fmt.Println(p0, fi31, fi35)
                                                _ = fi35
                                                _ = v36
                                                _ = v37
                                            }
                                            for fi38 := 0; fi38 < 13; fi38++ {
                                                v39 := GenS5{F0: v32, F1: 255}
                                                p2 = (^(p2))
                                                var v40 uint = 599
                                                _ = fi38
                                                _ = v39
                                                _ = v40
                                            }
                                        }
                                        var v41 float32 = -2.25
                                        _ = v41
                                    }
                                default:
                                    p1 = ((100 >= int8(p2)) && ((!p1) && (4294967295 <= 1)))
                                }
                                _ = len(v16)
                                _ = fi31
                                _ = v32
                            }
                            var v42 uint16 = 701
                            _ = v42
                        } else {
                            p1 = (0 < 4294967295)
                            v43 := GenS6{F0: 967, F1: ("\n\t" + ""), F2: 4294967295}
                            _ = v43
                        }
                    }
                }
            }
            p1 = p1
            for fi44 := 0; fi44 < 8; fi44++ {
                p2 += (p2 & (p2 + -379))
                _ = fi44
            }
        } else {
            var v45 int64 = 747
            _ = v45
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    _ = v16
    return 1
}

func main() {
    var v1 uint64 = 42
    var v2 float64 = 1e10
    v3 := make(map[string]string)
    v4 := make(map[int]uint32)
    var v5 int64 = 0
    v6 := make(map[int]uint32)
    v1 += v1
    if (1 != 42) {
        var v46 uint32 = 420
        var v47 uint64 = 0
        _ = v46
        _ = v47
    } else {
        var v48 bool = false
        v49 := GenS4{F0: 128, F1: -822, F2: v5}
        _ = v48
        _ = v49
    }
    fmt.Println(v1)
    for fi50 := 0; fi50 < 14; fi50++ {
        if false {
            if true {
                if true {
                    fmt.Println(fi50)
                    switch (fi50 ^ fi50) {
                    case 589:
                        fmt.Println(v1, fi50)
                    case -578:
                        if (!false) {
                            v51 := GenS1{F0: 2}
                            var v52 uint16 = 388
                            _ = v51
                            _ = v52
                        }
                    default:
                        fmt.Println(v5, fi50, v5)
                    }
                    var v53 int32 = -782
                    _ = v53
                }
            }
            var v54 float32 = 3.14159
            v55 := GenS7{F0: (v1 * v1), F1: ((v1 >= v1) && false), F2: 65535}
            _ = v54
            _ = v55
        } else {
            var v56 uint16 = uint16(v1)
            v1 -= 18446744073709551615
            _ = v56
        }
        var v57 int8 = 42
        _ = fi50
        _ = v57
    }
    fmt.Println(v2, v1, v5)
    fmt.Println(v2)
    v58 := GenS2{F0: 1, F1: uint32(v1), F2: -144}
    for fi59 := 0; fi59 < 19; fi59++ {
        v60 := GenS8{F0: 42, F1: 224, F2: 1}
        _ = fi59
        _ = v60
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v58
    _ = v6
}
