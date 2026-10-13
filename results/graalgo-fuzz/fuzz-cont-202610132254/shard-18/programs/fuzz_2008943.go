package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16; F1 string; F2 string }
type GenS4 struct { F0 uint8 }
type GenS5 struct { F0 string; F1 uint16; F2 uint64 }
type GenS6 struct { F0 int }
func genFn0(p0 uint8, p1 uint8, p2 string) int16 {
    v5 := []int{-9223372036854775808}
    var v6 float64 = -2.25
    if (519 > 42) {
        switch -63 {
        case 571:
            if (1e10 > (v6 / (1e10 + 3.14159))) {
                if true {
                    var v7 float64 = float64(1.5)
                    _ = v7
                }
                for fi8 := 0; fi8 < 17; fi8++ {
                    switch (fi8 * fi8) {
                    case 128:
                        v9 := GenS1{F0: false, F1: 554}
                        _ = v9
                    }
                    fmt.Println(fi8)
                    for fi10 := 0; fi10 < 13; fi10++ {
                        if false {
                            p0 += 0
                            var v11 string = p2
                            var v12 int8 = 0
                            _ = v11
                            _ = v12
                        } else {
                            var v13 float32 = 3.14159
                            p1 += p0
                            _ = v13
                        }
                        var v14 int16 = 1
                        _ = fi10
                        _ = v14
                    }
                    _ = fi8
                }
                v6 = (3.14159 * v6)
            }
        case 9:
            _ = len(v5)
        default:
            p0 = 101
        }
        p1 += (p1 & p1)
    }
    switch 128 {
    case 4294967295:
        var v15 uint8 = 34
        _ = v15
    case -1:
        switch 127 {
        case 32768:
            p0 = (p1 % ((127) | 1))
        case 434:
            _ = len(v5)
        case -1:
            v16 := GenS2{F0: 32767}
            _ = v16
        default:
            v17 := GenS2{F0: -1}
            _ = v17
        }
    default:
        fmt.Println(p2)
    }
    p1 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return -32768
}
func genFn1(p0 int64) uint32 {
    p0 = (p0 << ((65535) & 63))
    var v18 int8 = 127
    switch "hello" {
    case "a":
        p0 += (p0 >> ((255) & 63))
    case "":
        var v19 int16 = -392
        _ = v19
    }
    v20 := GenS3{F0: 1, F1: ("foo bar" + "hello"), F2: ""}
    v21 := []int{128, int(v18)}
    _ = p0
    _ = v18
    _ = v20
    _ = v21
    return 2
}

func main() {
    var v1 int8 = 127
    var v2 uint8 = 75
    v3 := make(map[string]string)
    v4 := []string{(("" + "foo bar") + "\n\t"), "\n\t", "\n\t", "foo bar"}
    if true {
        var v22 uint64 = 2
        _ = v22
    } else {
        if (-2147483648 == 1) {
            v2 -= (v2 / (((^(v2))) | 1))
            if (1e10 <= (float32(((float64(1.5) * 3.14159) / 3.14159)) / -2.25)) {
                delete(v3, "foo bar")
            } else {
                if false {
                    if true {
                        var v23 uint8 = (v2 | (^(v2)))
                        for fi24 := 0; fi24 < 8; fi24++ {
                            for fi25 := 0; fi25 < 16; fi25++ {
                                if true {
                                    var v26 int16 = 14
                                    var v27 float32 = 3.14159
                                    _ = v26
                                    _ = v27
                                } else {
                                    v1 += v1
                                }
                                v28 := GenS2{F0: -193}
                                if (!false) {
                                    switch "" {
                                    case "a":
                                        if (true || (-785 >= 831)) {
                                            var v29 uint16 = 1
                                            var v30 string = ""
                                            if len(v4) > 0 {
                                                v4[0] = v30
                                            }
                                            _ = v29
                                            _ = v30
                                        }
                                    case "foo bar":
                                        v31 := GenS4{F0: (v2 >> ((0) & 63))}
                                        _ = v31
                                    case "\n\t":
                                        v2 = v23
                                    }
                                    var v32 bool = ((0.1 * float64((0.1 * (-2.25 * float32((1.5 * 1.5)))))) < 3.14159)
                                    switch "foo bar" {
                                    case "a":
                                        _ = len(v3)
                                    case "foo bar":
                                        v1 = v1
                                    case "":
                                        v2 = v2
                                    }
                                    _ = v32
                                } else {
                                    var v33 int = fi25
                                    var v34 int64 = 42
                                    _ = len(v3)
                                    _ = v33
                                    _ = v34
                                }
                                v2 += v2
                                _ = fi25
                                _ = v28
                            }
                            _ = fi24
                        }
                        _ = v23
                    } else {
                        for fi35 := 0; fi35 < 17; fi35++ {
                            v1 += v1
                            _ = fi35
                        }
                    }
                } else {
                    v1 -= (^(v1))
                    v2 = 129
                    v2 -= (v2 >> ((4294967295) & 63))
                }
                fmt.Println(v2, v1)
            }
        } else {
            v2 = (^(v2))
            for fi36 := 0; fi36 < 2; fi36++ {
                fmt.Println(fi36, fi36)
                v2 -= v2
                _ = len(v3)
                _ = fi36
            }
            v2 += 0
        }
        for fi37 := 0; fi37 < 6; fi37++ {
            _ = len(v4)
            _ = fi37
        }
    }
    fmt.Println(v2, v1)
    for fi38 := 0; fi38 < 13; fi38++ {
        for fi39 := 0; fi39 < 14; fi39++ {
            var v40 uint8 = v2
            var v41 uint64 = 0
            _ = fi39
            _ = v40
            _ = v41
        }
        if true {
            v42 := GenS5{F0: ("a" + "a"), F1: 788, F2: 4294967295}
            _ = v42
        }
        _ = fi38
    }
    for fi43 := 0; fi43 < 13; fi43++ {
        v4 = append(v4, "\n\t")
        _ = fi43
    }
    v2 = (v2 << ((uint(v2)) & 63))
    if ((-2.25 * (0.1 + 0.1)) > float64((-2.25 * 0.1))) {
        v44 := GenS3{F0: -1, F1: "a", F2: ("\n\t" + ("" + ("a" + ("" + "a"))))}
        fmt.Println(v2)
        _ = v44
    } else {
        var v45 uint32 = uint32(v2)
        var v46 bool = true
        _ = v45
        _ = v46
    }
    for fi47 := 0; fi47 < 12; fi47++ {
        for fi48 := 0; fi48 < 16; fi48++ {
            var v49 int16 = -311
            var v50 float32 = (0.0 / 3.14159)
            _ = fi48
            _ = v49
            _ = v50
        }
        var v51 uint16 = 568
        _ = fi47
        _ = v51
    }
    fmt.Println(v1, v1)
    var v52 float32 = 0.0
    v53 := GenS6{F0: 9223372036854775807}
    if ((-1 < -942) && false) {
        var v54 int64 = -9223372036854775808
        _ = v54
    }
    var v55 int = 816
    v1 = (v1 ^ -1)
    for fi56 := 0; fi56 < 5; fi56++ {
        var v57 int = 4294967295
        v4 = append(v4, "\n\t")
        for fi58 := 0; fi58 < 17; fi58++ {
            v52 = v52
            _ = fi58
        }
        _ = fi56
        _ = v57
    }
    fmt.Println(v1)
    fmt.Println(v52, v2, v1)
    v59 := v3[(("foo bar" + "a") + "a")]
    if len(v4) > 0 {
        v4[0] = v59
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v52:", v52)
    fmt.Println("v55:", v55)
    fmt.Println("v59:", v59)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v52
    _ = v53
    _ = v55
    _ = v59
}
