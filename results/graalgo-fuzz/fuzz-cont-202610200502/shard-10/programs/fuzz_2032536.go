package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64 }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int32; F1 bool; F2 int64 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 int8; F1 uint32 }
type GenS7 struct { F0 int32; F1 int; F2 uint16 }
func genFn0(p0 uint32, p1 string, p2 uint8) uint16 {
    p0 += (p0 ^ p0)
    v5 := []int{}
    p0 = 42
    p0 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 42
}
func genFn1(p0 int32, p1 uint16, p2 uint8) int16 {
    for fi6 := 0; fi6 < 20; fi6++ {
        p2 = (p2 & p2)
        _ = fi6
    }
    p2 = (p2 + 0)
    var v7 int8 = 29
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 980
}

func main() {
    var v1 uint = 29
    v2 := make(map[int]bool)
    v3 := []uint32{42, 0, 2147483647, 2147483647}
    v4 := make(map[int]uint32)
    for fi8 := 0; fi8 < 9; fi8++ {
        var v9 uint = 7
        if false {
            switch (^(fi8)) {
            case 2147483647:
                v9 = (^(v9))
            }
        }
        _ = fi8
        _ = v9
    }
    fmt.Println(v1)
    var v10 int = -937
    for fi11 := 0; fi11 < 16; fi11++ {
        var v12 bool = (!false)
        v13 := GenS1{F0: 1, F1: 334}
        _ = fi11
        _ = v12
        _ = v13
    }
    if true {
        v1 -= (v1 / ((255) | 1))
    } else {
        var v14 float64 = 3.14159
        for fi15 := 0; fi15 < 7; fi15++ {
            v1 -= (v1 >> ((v1) & 63))
            _ = fi15
        }
        switch (-(v10)) {
        case 42:
            var v16 int64 = 135
            _ = v16
        case -573:
            var v17 int8 = 76
            _ = v17
        default:
            switch "" {
            case "a":
                v14 = float64(3.14159)
            case "foo bar":
                var v18 uint8 = 9
                _ = v18
            case "hello":
                v1 = (^(v1))
            }
        }
        _ = v14
    }
    v19 := GenS2{F0: true, F1: -1}
    var v20 float32 = (((float32((-2.25 - 1e10)) * (3.14159 - (3.14159 / 1.5))) + 1e10) + float32(((1.5 + (float64((((((1e10 / (-2.25 + -2.25)) + (0.1 * 1.5)) * (((float32(0.0) * 0.1) + 3.14159) / 1e10)) + (-2.25 + 3.14159)) * -2.25)) / (1e10 * 0.1))) + 1.5)))
    v20 += v20
    var v21 float64 = (1e10 * 0.1)
    _ = len(v4)
    var v22 int8 = 1
    var v23 uint16 = 1
    switch (("a" + ("foo bar" + (("" + "foo bar") + ("a" + (("hello" + "") + (("foo bar" + (((("\n\t" + "foo bar") + "foo bar") + "") + "foo bar")) + "foo bar")))))) + ("" + ("" + "foo bar"))) {
    case "hello":
        for fi24 := 0; fi24 < 15; fi24++ {
            v23 = (v23 * 255)
            var v25 uint8 = 123
            _ = fi24
            _ = v25
        }
    case "":
        for fi26 := 0; fi26 < 19; fi26++ {
            v27 := GenS2{F0: (1 >= -548), F1: 2147483647}
            var v28 int16 = 0
            var v29 float32 = (1.5 / (((v20 / ((1e10 - float32(v21)) / (float32(((1.5 * (((((v21 / v21) * float64(3.14159)) + v21) - (((v21 - float64(0.1)) / 3.14159) / v21)) + -2.25)) + (v21 + (v21 / (v21 / 3.14159))))) * (v20 + -2.25)))) * v20) * v20))
            _ = fi26
            _ = v27
            _ = v28
            _ = v29
        }
    default:
        var v30 int16 = 42
        _ = v30
    }
    for fi31 := 0; fi31 < 11; fi31++ {
        v32 := GenS2{F0: (!((!((!(101 >= 0)) || (("hello" + ("a" + "a")) < "\n\t"))) && false)), F1: 2147483647}
        for fi33 := 0; fi33 < 17; fi33++ {
            if false {
                v34 := GenS1{F0: (v22 << (((v1 | (v1 >> (((v1 - v1)) & 63)))) & 63)), F1: 2}
                v35 := GenS3{F0: 128}
                _ = v34
                _ = v35
            } else {
                var v36 uint = v1
                _ = v36
            }
            fmt.Println(v21, fi31, v10)
            if false {
                var v37 int16 = 704
                if false {
                    v38 := GenS1{F0: v22, F1: 774}
                    if (128 >= 9223372036854775807) {
                        var v39 uint32 = 2147483648
                        _ = v39
                    }
                    var v40 uint = v1
                    _ = v38
                    _ = v40
                } else {
                    var v41 int32 = -1
                    _ = v41
                }
                v21 = (v21 - 0.1)
                if false {
                    v22 += (v22 >> (((v1 % ((v1) | 1))) & 63))
                    v20 += v20
                    v42 := GenS2{F0: (1 < 2), F1: -924}
                    _ = v42
                } else {
                    if false {
                        var v43 int = (v10 - 1)
                        if (2 == 128) {
                            switch (^(v43)) {
                            case 42:
                                v44 := GenS4{F0: 1, F1: (v23 <= v23), F2: -299}
                                _ = v44
                            case 128:
                                v45 := GenS5{F0: 10}
                                _ = v45
                            case 9223372036854775807:
                                v37 = genFn1(78, (v23 / ((genFn0(2147483647, "foo bar", uint8(v1))) | 1)), 122)
                            default:
                                if true {
                                    if len(v3) > 0 {
                                        v3[0] = 1
                                    }
                                    var v46 float32 = v20
                                    _ = v46
                                } else {
                                    v47 := GenS6{F0: v22, F1: 2147483648}
                                    switch "\n\t" {
                                    case "foo bar":
                                        v21 += v21
                                    case "a":
                                        _ = len(v2)
                                    default:
                                        var v48 float64 = (((v21 / v21) + 0.0) * v21)
                                        _ = v48
                                    }
                                    v49 := GenS4{F0: 2147483647, F1: (42 >= 42), F2: -9223372036854775808}
                                    _ = v47
                                    _ = v49
                                }
                                v3 = append(v3, 2147483648)
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                var v50 uint32 = 2147483648
                                _ = v50
                            }
                        }
                        var v51 int16 = v37
                        if len(v3) > 0 {
                            v3[0] = 2
                        }
                        _ = v43
                        _ = v51
                    }
                }
                fmt.Println(v10)
                _ = v37
            } else {
                v1 = v1
            }
            _ = fi33
        }
        _ = fi31
        _ = v32
    }
    for fi52 := 0; fi52 < 11; fi52++ {
        if true {
            var v53 uint = (^(v1))
            var v54 uint64 = 1
            var v55 int64 = 235
            _ = v53
            _ = v54
            _ = v55
        } else {
            if false {
                v23 = (v23 / ((32767) | 1))
                delete(v2, (fi52 % ((v10) | 1)))
            }
            if len(v3) > 0 {
                v3[0] = 2147483647
            }
            var v56 int16 = 1
            _ = v56
        }
        v57 := GenS7{F0: 848, F1: v10, F2: (^(v23))}
        _ = fi52
        _ = v57
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
}
