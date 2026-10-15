package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint; F2 uint }
func genFn0() int16 {
    var v5 uint = 1
    fmt.Println(v5)
    v6 := GenS1{F0: v5, F1: 681, F2: v5}
    fmt.Println(v5)
    var v7 uint16 = 816
    _ = v5
    _ = v6
    _ = v7
    return 472
}
func genFn1(p0 bool, p1 int, p2 int16) int32 {
    switch (p1 & (^(p1))) {
    case 42:
        for fi8 := 0; fi8 < 5; fi8++ {
            p1 += (fi8 & p1)
            fmt.Println(p2, p2, p1)
            _ = fi8
        }
    case 275:
        for fi9 := 0; fi9 < 14; fi9++ {
            fmt.Println(fi9, fi9)
            switch (fi9 ^ fi9) {
            case -670:
                var v10 bool = false
                _ = v10
            case 128:
                var v11 uint = 199
                _ = v11
            case 4294967295:
                var v12 int = (fi9 ^ p1)
                _ = v12
            default:
                p2 = (-(p2))
            }
            fmt.Println(p2, p2, p2)
            _ = fi9
        }
    case 248:
        p2 = (p2 / ((p2) | 1))
    default:
        p1 += p1
    }
    var v13 uint8 = 127
    if true {
        switch "hello" {
        case "hello":
            for fi14 := 0; fi14 < 13; fi14++ {
                p0 = (p0 && p0)
                _ = fi14
            }
        case "\n\t":
            p1 = p1
        }
    }
    v13 += v13
    p0 = (!(!((3.14159 - (1e10 * -2.25)) != 3.14159)))
    _ = p0
    _ = p1
    _ = p2
    _ = v13
    return 2147483647
}

func main() {
    v1 := make(map[int]int)
    var v2 int8 = -1
    var v3 int = 128
    v4 := []string{"foo bar", "a"}
    var v15 bool = false
    _ = len(v1)
    if v15 {
        if true {
            for fi16 := 0; fi16 < 19; fi16++ {
                fmt.Println(fi16)
                _ = fi16
            }
            if v15 {
                var v17 float64 = 1.5
                var v18 float32 = (((0.1 + 0.1) * 0.1) - 1.5)
                v17 = v17
                _ = v17
                _ = v18
            }
            v1[(v3 + (v3 & (v3 | v3)))] = 311
        } else {
            v2 = v2
            for fi19 := 0; fi19 < 20; fi19++ {
                var v20 float64 = 3.14159
                var v21 float32 = 3.14159
                if (("a" + "foo bar") <= "") {
                    delete(v1, 32768)
                    if true {
                        v21 = v21
                    } else {
                        switch "" {
                        case "hello":
                            if v15 {
                                fmt.Println(v21)
                            }
                        case "\n\t":
                            v22 := GenS1{F0: 0, F1: 1, F2: 4294967295}
                            _ = v22
                        case "foo bar":
                            fmt.Println(v2, v20)
                        default:
                            for fi23 := 0; fi23 < 9; fi23++ {
                                var v24 uint8 = 92
                                v24 -= v24
                                var v25 int = -673
                                var v26 uint64 = 717
                                v27 := v1[(^(fi23))]
                                _ = fi23
                                _ = v24
                                _ = v25
                                _ = v26
                                _ = v27
                            }
                        }
                    }
                    for fi28 := 0; fi28 < 4; fi28++ {
                        var v29 int32 = 760
                        var v30 float64 = (1.5 + v20)
                        _ = fi28
                        _ = v29
                        _ = v30
                    }
                } else {
                    var v31 int16 = 1
                    v32 := v1[(fi19 % ((965) | 1))]
                    _ = v31
                    _ = v32
                }
                _ = fi19
                _ = v20
                _ = v21
            }
        }
        v33 := GenS1{F0: 65535, F1: 65535, F2: 2}
        _ = v33
    } else {
        if ((!(0 < 104)) || (-2.25 < 3.14159)) {
            fmt.Println(v15)
            v2 += (v2 ^ v2)
        }
    }
    for fi34 := 0; fi34 < 5; fi34++ {
        fmt.Println(v2, fi34)
        fmt.Println(v2, v15, v15)
        _ = fi34
    }
    for fi35 := 0; fi35 < 12; fi35++ {
        v2 -= (v2 << ((42) & 63))
        fmt.Println(v3, v2)
        var v36 int64 = 276
        _ = fi35
        _ = v36
    }
    var v37 uint32 = 1
    var v38 uint32 = v37
    fmt.Println(v15, v38)
    v15 = ((3.14159 - 1e10) < (3.14159 * ((((1e10 + 3.14159) * 3.14159) - float64((((float32(0.0) / ((1.5 + (0.1 / float32(1.5))) * 0.1)) / 1e10) - (0.1 - 1.5)))) * -2.25)))
    _ = len(v1)
    for fi39 := 0; fi39 < 1; fi39++ {
        v2 = (v2 >> ((18446744073709551615) & 63))
        v40 := GenS1{F0: 18446744073709551615, F1: 255, F2: 2}
        _ = fi39
        _ = v40
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("v38:", v38)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v37
    _ = v38
    _ = v4
}
