package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint16; F2 int64 }
type GenS2 struct { F0 int64; F1 int32; F2 int16 }
type GenS3 struct { F0 uint }
type GenS4 struct { F0 bool; F1 int; F2 int16 }
func genFn0(p0 uint8, p1 uint16, p2 int32) int16 {
    fmt.Println(p1, p2, p0)
    switch "\n\t" {
    case "\n\t":
        p2 = -1
    case "a":
        var v6 uint16 = p1
        _ = v6
    case "hello":
        var v7 int = 485
        _ = v7
    default:
        for fi8 := 0; fi8 < 6; fi8++ {
            v9 := GenS1{F0: ("a" + ("hello" + ("\n\t" + "hello"))), F1: (p1 - 65535), F2: 1}
            if (!(262 < 1)) {
                v10 := []int{267, fi8, (fi8 + fi8)}
                for fi11 := 0; fi11 < 10; fi11++ {
                    var v12 uint = 498
                    p2 += p2
                    if (!true) {
                        v13 := GenS1{F0: "\n\t", F1: 0, F2: 42}
                        _ = v13
                    } else {
                        if (!(p1 <= 725)) {
                            v12 = v12
                            p2 = p2
                            var v14 uint = (^(v12))
                            _ = v14
                        } else {
                            p0 = (^(p0))
                        }
                    }
                    _ = fi11
                    _ = v12
                }
                if len(v10) > 0 {
                    v10[0] = fi8
                }
                _ = v10
            }
            v15 := []int{fi8, (fi8 - (fi8 >> ((uint(p0)) & 63))), fi8, (fi8 - -307)}
            _ = fi8
            _ = v15
            _ = v9
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 42
}
func genFn1(p0 int8, p1 int) uint32 {
    p0 = p0
    for fi16 := 0; fi16 < 3; fi16++ {
        var v17 int8 = (-(p0))
        _ = fi16
        _ = v17
    }
    p0 = p0
    _ = p0
    _ = p1
    return 65535
}

func main() {
    var v1 int8 = -22
    var v2 uint = 0
    var v3 int64 = 9223372036854775807
    var v4 uint64 = 647
    var v5 string = ("a" + ("hello" + ("a" + ((("hello" + "a") + "foo bar") + ("a" + "hello")))))
    if false {
        for fi18 := 0; fi18 < 12; fi18++ {
            v19 := []int{fi18, fi18, fi18}
            var v20 uint16 = 32767
            v3 += (v3 | (v3 ^ v3))
            _ = fi18
            _ = v19
            _ = v20
        }
        var v21 int64 = v3
        _ = v21
    } else {
        v3 -= 9223372036854775807
    }
    fmt.Println(v1)
    v22 := GenS1{F0: ("\n\t" + ("foo bar" + "foo bar")), F1: 317, F2: v3}
    fmt.Println(v5, v1, v2)
    v1 = v1
    fmt.Println(v1, v3, v4)
    v5 = "foo bar"
    fmt.Println(v1)
    v4 = (v4 & (v4 & (^(v4))))
    switch 1 {
    case 32768:
        for fi23 := 0; fi23 < 18; fi23++ {
            v4 = v4
            v24 := GenS1{F0: v5, F1: 2, F2: -1}
            for fi25 := 0; fi25 < 1; fi25++ {
                v3 = (v3 + (v3 - -1))
                _ = fi25
            }
            _ = fi23
            _ = v24
        }
    default:
        switch v5 {
        case "\n\t":
            v26 := []int{-9223372036854775808, -639, int(v1), 32767}
            _ = v26
        }
    }
    v2 = v2
    v1 += (v1 | (-(v1)))
    v1 = 100
    if ("hello" > (v5 + v5)) {
        switch 449 {
        case -2147483648:
            v2 -= v2
        case 1:
            v27 := GenS1{F0: "\n\t", F1: 926, F2: v3}
            _ = v27
        case 4294967295:
            v28 := GenS1{F0: ("\n\t" + "a"), F1: 849, F2: (v3 ^ (v3 >> ((v2) & 63)))}
            _ = v28
        default:
            v29 := []int{-9223372036854775808, 32768, 7, 606}
            _ = v29
        }
    } else {
        v3 += 919
        fmt.Println(v4, v1)
    }
    v2 = 1
    for fi30 := 0; fi30 < 4; fi30++ {
        var v31 float64 = (float64(1.5) / (0.1 / 1.5))
        switch fi30 {
        case -2147483648:
            if (genFn0(188, 3, 108) != 42) {
                var v32 int = 327
                v31 = v31
                fmt.Println(fi30)
                _ = v32
            }
        case 2147483647:
            var v33 uint = (v2 << ((uint(v4)) & 63))
            _ = v33
        case 0:
            v3 = v3
        }
        _ = fi30
        _ = v31
    }
    if true {
        var v34 float64 = float64(1.5)
        for fi35 := 0; fi35 < 17; fi35++ {
            v36 := GenS2{F0: v3, F1: 145, F2: -416}
            if false {
                v37 := GenS1{F0: v5, F1: 32767, F2: (v3 - v3)}
                switch fi35 {
                case 32768:
                    if false {
                        var v38 bool = (!(!false))
                        _ = v38
                    }
                case 2147483647:
                    switch v5 {
                    case "foo bar":
                        for fi39 := 0; fi39 < 1; fi39++ {
                            v5 = "hello"
                            v5 = v5
                            v34 = float64((((((0.0 - 0.1) * 1.5) + ((1.5 + -2.25) - 3.14159)) * 1e10) * 3.14159))
                            _ = fi39
                        }
                    default:
                        for fi40 := 0; fi40 < 14; fi40++ {
                            v41 := []int{(fi40 * (fi35 * fi35)), 4294967295, fi40}
                            switch "" {
                            case "hello":
                                if true {
                                    var v42 int8 = (-(v1))
                                    _ = v42
                                } else {
                                    if false {
                                        if len(v41) > 0 {
                                            v41[0] = fi35
                                        }
                                        v34 += 0.1
                                        v43 := GenS1{F0: v5, F1: 1, F2: (v3 % ((v3) | 1))}
                                        _ = v43
                                    } else {
                                        if (3.14159 != float32((3.14159 - float64((3.14159 / 1.0))))) {
                                            v1 -= v1
                                            v5 = ""
                                        } else {
                                            v1 += v1
                                        }
                                    }
                                    var v44 int32 = 0
                                    _ = v44
                                }
                            case "":
                                _ = len(v41)
                                _ = len(v41)
                                v5 = "\n\t"
                            case "\n\t":
                                v2 += (v2 % (((^(v2))) | 1))
                                var v45 uint16 = 255
                                v46 := GenS3{F0: (v2 - (v2 << (((v2 << (((v2 >> ((v2) & 63))) & 63))) & 63)))}
                                _ = v45
                                _ = v46
                            default:
                                v47 := GenS1{F0: "hello", F1: 32767, F2: -9223372036854775808}
                                if true {
                                    v41 = append(v41, fi35)
                                    var v48 int32 = 0
                                    v1 -= int8(v4)
                                    _ = v48
                                }
                                _ = v47
                            }
                            v1 = (v1 >> (((v2 / ((18446744073709551615) | 1))) & 63))
                            var v49 int16 = int16(v1)
                            v50 := GenS1{F0: v5, F1: 65535, F2: v3}
                            _ = fi40
                            _ = v41
                            _ = v49
                            _ = v50
                        }
                    }
                default:
                    for fi51 := 0; fi51 < 16; fi51++ {
                        fmt.Println(v4)
                        _ = fi51
                    }
                }
                _ = v37
            } else {
                var v52 int32 = 42
                _ = v52
            }
            if (v34 <= v34) {
                v5 = v5
                v53 := GenS4{F0: (v2 <= 65535), F1: 42, F2: -345}
                _ = v53
            }
            _ = fi35
            _ = v36
        }
        v1 += v1
        _ = v34
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
}
