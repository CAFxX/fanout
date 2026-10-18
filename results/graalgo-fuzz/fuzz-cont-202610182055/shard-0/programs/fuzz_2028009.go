package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint16 }
type GenS2 struct { F0 bool; F1 int32 }
type GenS3 struct { F0 int }
type GenS4 struct { F0 uint; F1 int8; F2 uint8 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 string; F1 uint8 }
func genFn0(p0 int64, p1 bool, p2 int16) bool {
    v4 := GenS1{F0: -1, F1: 0}
    v5 := GenS1{F0: 127, F1: 255}
    var v6 uint = 20
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    _ = v6
    return ("hello" == ("\n\t" + "\n\t"))
}
func genFn1(p0 uint64) uint8 {
    switch ("\n\t" + "") {
    case "\n\t":
        switch 32767 {
        case -9223372036854775808:
            v7 := []int{-9223372036854775808, 32768, -300}
            _ = v7
        case 1:
            v8 := GenS1{F0: -16, F1: 32767}
            _ = v8
        }
    case "":
        var v9 int32 = 0
        _ = v9
    default:
        if (false || (0.0 <= (((((float64(1e10) / -2.25) / 1e10) + -2.25) - -2.25) - 0.1))) {
            var v10 uint64 = p0
            var v11 float64 = ((3.14159 / 3.14159) - (0.0 * 0.0))
            var v12 bool = true
            _ = v10
            _ = v11
            _ = v12
        } else {
            if true {
                p0 = (p0 / ((p0) | 1))
                switch "a" {
                case "hello":
                    fmt.Println(p0)
                case "foo bar":
                    var v13 int16 = 1000
                    _ = v13
                case "a":
                    var v14 int32 = 2147483647
                    _ = v14
                }
            }
            for fi15 := 0; fi15 < 6; fi15++ {
                v16 := []int{fi15, (fi15 / ((fi15) | 1)), fi15}
                if len(v16) > 0 {
                    v16[0] = (fi15 >> ((42) & 63))
                }
                _ = len(v16)
                _ = fi15
                _ = v16
            }
        }
    }
    switch 9223372036854775807 {
    case 32767:
        switch "" {
        case "":
            p0 -= p0
        case "foo bar":
            switch -2147483648 {
            case 4294967295:
                p0 = (p0 | p0)
            case 127:
                p0 = p0
            default:
                if false {
                    var v17 int = 912
                    var v18 int16 = 1000
                    v18 = (v18 + v18)
                    _ = v17
                    _ = v18
                }
            }
        default:
            p0 = (p0 - p0)
        }
    case -2147483648:
        v19 := []int{9223372036854775807, -9223372036854775808}
        _ = v19
    case 4294967295:
        p0 = 414
    default:
        p0 = (p0 - (p0 ^ (p0 / ((476) | 1))))
    }
    var v20 float32 = (0.0 + 0.1)
    v21 := []int{}
    _ = p0
    _ = v20
    _ = v21
    return 221
}

func main() {
    var v1 int8 = 42
    v2 := make(map[int]string)
    v3 := []string{"", "a", "foo bar"}
    if (2147483647 < 0) {
        v3 = append(v3, "foo bar")
    }
    for fi22 := 0; fi22 < 9; fi22++ {
        for fi23 := 0; fi23 < 1; fi23++ {
            if false {
                if (1 < -637) {
                    _ = len(v3)
                    fmt.Println(fi22)
                } else {
                    fmt.Println(fi22, v1, v1)
                    if (1 != 9223372036854775807) {
                        if len(v3) >= 2 {
                            v3 = v3[1:len(v3)]
                        }
                        v24 := v2[fi22]
                        fmt.Println(fi22)
                        _ = v24
                    }
                    v1 = -1
                }
                var v25 int64 = -558
                var v26 float64 = 0.1
                _ = v25
                _ = v26
            }
            _ = fi23
        }
        v27 := GenS1{F0: (v1 ^ -34), F1: 1}
        _ = fi22
        _ = v27
    }
    v1 += v1
    v28 := GenS1{F0: (v1 * v1), F1: 1}
    var v29 string = "\n\t"
    if true {
        if (false || false) {
            if false {
                for fi30 := 0; fi30 < 8; fi30++ {
                    v29 = (v29 + "foo bar")
                    for fi31 := 0; fi31 < 8; fi31++ {
                        v29 = "hello"
                        switch ("hello" + (v29 + "a")) {
                        case "hello":
                            switch ("a" + "a") {
                            case "hello":
                                fmt.Println(fi31, v1)
                            default:
                                var v32 uint = 1
                                _ = v32
                            }
                        case "\n\t":
                            v1 = (v1 & v1)
                        case "":
                            v1 -= -1
                            switch "\n\t" {
                            case "a":
                                if false {
                                    if true {
                                        if len(v3) >= 2 {
                                            v3 = v3[1:len(v3)]
                                        }
                                    }
                                }
                            case "\n\t":
                                v2[(fi31 - 26)] = v29
                            case "":
                                v33 := GenS2{F0: false, F1: 2147483647}
                                _ = v33
                            default:
                                v1 -= (v1 - v1)
                                switch fi30 {
                                case 127:
                                    switch v29 {
                                    case "hello":
                                        v34 := GenS3{F0: (fi31 / ((fi30) | 1))}
                                        _ = v34
                                    case "a":
                                        var v35 uint16 = 65535
                                        _ = v35
                                    default:
                                        v36 := GenS4{F0: 0, F1: (v1 % ((v1) | 1)), F2: 0}
                                        _ = v36
                                    }
                                case 211:
                                    switch fi31 {
                                    case 2147483647:
                                        var v37 uint8 = 42
                                        _ = v37
                                    case 9223372036854775807:
                                        v38 := GenS5{F0: 42}
                                        _ = v38
                                    case 127:
                                        v39 := GenS5{F0: 4294967295}
                                        _ = v39
                                    }
                                default:
                                    v1 += (v1 & (v1 ^ -128))
                                }
                                fmt.Println(v1, fi30, fi30)
                            }
                        }
                        v40 := GenS4{F0: 65535, F1: 0, F2: 209}
                        _ = fi31
                        _ = v40
                    }
                    _ = fi30
                }
                v41 := GenS1{F0: -1, F1: 32767}
                _ = v41
            }
            fmt.Println(v1, v1)
        } else {
            v1 = (v1 * 1)
            fmt.Println(v1)
            if ((1.5 < ((3.14159 - (1e10 + (1.5 * 0.0))) - 0.0)) || (65535 > 18446744073709551615)) {
                if true {
                    var v42 float64 = float64(((float32(-2.25) * -2.25) + 0.0))
                    var v43 float64 = (float64(-2.25) / 0.1)
                    switch 32767 {
                    case 1:
                        if true {
                            for fi44 := 0; fi44 < 7; fi44++ {
                                if true {
                                    fmt.Println(v43)
                                }
                                _ = fi44
                            }
                            v42 = v43
                        }
                    case -2147483648:
                        v45 := GenS6{F0: ((("" + "\n\t") + v29) + v29), F1: 42}
                        _ = v45
                    }
                    _ = v42
                    _ = v43
                }
            } else {
                for fi46 := 0; fi46 < 8; fi46++ {
                    var v47 float32 = 0.0
                    fmt.Println(v1)
                    _ = len(v3)
                    _ = fi46
                    _ = v47
                }
                var v48 int64 = -601
                _ = v48
            }
        }
    }
    v49 := GenS1{F0: 42, F1: 42}
    if (!false) {
        var v50 int16 = 1
        var v51 uint32 = 4294967295
        _ = v50
        _ = v51
    } else {
        v1 = v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v29:", v29)
    fmt.Println("slicelen v3:", len(v3), v3)
    _ = v1
    _ = v2
    _ = v28
    _ = v29
    _ = v3
    _ = v49
}
