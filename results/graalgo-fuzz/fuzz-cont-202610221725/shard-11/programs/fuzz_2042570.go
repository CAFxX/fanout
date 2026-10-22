package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 int16 }
type GenS2 struct { F0 string; F1 string; F2 uint32 }
type GenS3 struct { F0 int64 }
type GenS4 struct { F0 string; F1 bool }
type GenS5 struct { F0 uint64; F1 uint; F2 uint }
type GenS6 struct { F0 uint32 }
type GenS7 struct { F0 int }
type GenS8 struct { F0 int16 }
type GenS9 struct { F0 int16; F1 int64; F2 bool }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[int]int)
    var v3 int64 = 42
    switch -1 {
    case 490:
        v4 := GenS1{F0: (0 >= 0), F1: 593, F2: 42}
        _ = v4
    default:
        var v5 float32 = (3.14159 * (0.0 / 1.5))
        _ = v5
    }
    for fi6 := 0; fi6 < 2; fi6++ {
        v7 := GenS1{F0: true, F1: 137, F2: 238}
        for fi8 := 0; fi8 < 12; fi8++ {
            v9 := GenS2{F0: "a", F1: "foo bar", F2: 4294967295}
            fmt.Println(fi8, v3)
            fmt.Println(fi8, v3, fi6)
            v10 := GenS2{F0: ("\n\t" + ""), F1: "hello", F2: 65535}
            switch ("foo bar" + "\n\t") {
            case "":
                var v11 string = ("foo bar" + (("\n\t" + "a") + ""))
                _ = v11
            case "foo bar":
                var v12 int64 = (v3 ^ (-(v3)))
                fmt.Println(v3, fi6, v3)
                _ = v12
            case "hello":
                v13 := GenS3{F0: (v3 % ((v3) | 1))}
                _ = v13
            default:
                v2[fi8] = fi6
            }
            _ = fi8
            _ = v10
            _ = v9
        }
        v2[fi6] = 128
        _ = fi6
        _ = v7
    }
    fmt.Println(v3)
    for fi14 := 0; fi14 < 3; fi14++ {
        v15 := GenS4{F0: (("a" + "") + "\n\t"), F1: (!(-55 > 0))}
        if true {
            for fi16 := 0; fi16 < 17; fi16++ {
                v3 = (v3 | v3)
                _ = fi16
            }
        } else {
            v17 := GenS5{F0: 4294967295, F1: 1, F2: 451}
            _ = v17
        }
        switch (fi14 - 9223372036854775807) {
        case 9223372036854775807:
            var v18 float32 = 0.0
            _ = v18
        default:
            if true {
                var v19 uint = 1
                _ = v19
            } else {
                delete(v1, (fi14 % (((^(fi14))) | 1)))
            }
        }
        _ = fi14
        _ = v15
    }
    var v20 uint16 = 32767
    if false {
        v21 := GenS2{F0: ("hello" + ("\n\t" + ("" + ""))), F1: ("a" + (("a" + "\n\t") + "a")), F2: 414}
        _ = v21
    } else {
        if false {
            v22 := GenS6{F0: 439}
            _ = v22
        }
    }
    v23 := []int{127, 187, 736}
    var v24 int64 = (v3 >> ((1) & 63))
    v3 = v3
    var v25 int32 = 1
    if ((("\n\t" + "foo bar") + ("" + "")) > "foo bar") {
        switch -411 {
        case -1:
            var v26 int32 = -2147483648
            _ = v26
        case -22:
            switch 0 {
            case 9223372036854775807:
                for fi27 := 0; fi27 < 2; fi27++ {
                    for fi28 := 0; fi28 < 2; fi28++ {
                        fmt.Println(v20, v25)
                        _ = fi28
                    }
                    _ = fi27
                }
            case -284:
                v3 -= v3
            case -9223372036854775808:
                switch -184 {
                case 9223372036854775807:
                    if (1 != 0) {
                        var v29 int16 = 42
                        var v30 float64 = -2.25
                        v31 := GenS7{F0: -510}
                        _ = v29
                        _ = v30
                        _ = v31
                    } else {
                        if (!false) {
                            for fi32 := 0; fi32 < 8; fi32++ {
                                var v33 uint = 42
                                v24 -= v24
                                for fi34 := 0; fi34 < 13; fi34++ {
                                    if false {
                                        var v35 int = fi32
                                        if true {
                                            if (((^(v20)) > (^(v20))) || (-9223372036854775808 < -72)) {
                                                var v36 int = (fi34 >> (((v33 << ((v33) & 63))) & 63))
                                                var v37 int8 = -22
                                                _ = v36
                                                _ = v37
                                            } else {
                                                switch 32768 {
                                                case -2147483648:
                                                    v33 -= 575
                                                }
                                                v38 := v1[2147483647]
                                                var v39 int32 = v25
                                                _ = v38
                                                _ = v39
                                            }
                                        }
                                        _ = v35
                                    } else {
                                        v24 -= v24
                                        v40 := GenS4{F0: "\n\t", F1: false}
                                        v41 := v2[(fi34 + fi34)]
                                        _ = v40
                                        _ = v41
                                    }
                                    v42 := GenS7{F0: (^(fi34))}
                                    _ = fi34
                                    _ = v42
                                }
                                _ = fi32
                                _ = v33
                            }
                            for fi43 := 0; fi43 < 7; fi43++ {
                                v24 = 1
                                if (-128 >= 100) {
                                    for fi44 := 0; fi44 < 14; fi44++ {
                                        v45 := GenS3{F0: v24}
                                        var v46 float32 = 3.14159
                                        _ = fi44
                                        _ = v45
                                        _ = v46
                                    }
                                    v3 -= (v3 >> ((1) & 63))
                                }
                                _ = fi43
                            }
                            fmt.Println(v25, v20)
                        } else {
                            v47 := GenS8{F0: -109}
                            if true {
                                fmt.Println(v25)
                                if true {
                                    v25 += v25
                                    for fi48 := 0; fi48 < 18; fi48++ {
                                        var v49 int16 = -32768
                                        _ = fi48
                                        _ = v49
                                    }
                                    if len(v23) > 0 {
                                        v23[0] = 2147483647
                                    }
                                } else {
                                    v50 := GenS1{F0: true, F1: v20, F2: 32767}
                                    v24 += (v3 | -9223372036854775808)
                                    for fi51 := 0; fi51 < 16; fi51++ {
                                        v52 := GenS5{F0: 0, F1: 255, F2: 638}
                                        for fi53 := 0; fi53 < 7; fi53++ {
                                            var v54 uint8 = 42
                                            _ = fi53
                                            _ = v54
                                        }
                                        var v55 float32 = (0.1 * 1e10)
                                        _ = fi51
                                        _ = v52
                                        _ = v55
                                    }
                                    _ = v50
                                }
                            }
                            _ = v47
                        }
                        v56 := GenS3{F0: v3}
                        _ = v56
                    }
                }
            default:
                for fi57 := 0; fi57 < 2; fi57++ {
                    v3 -= 9223372036854775807
                    v58 := GenS9{F0: -32768, F1: (v24 - v3), F2: false}
                    _ = fi57
                    _ = v58
                }
            }
        case 127:
            v25 += (v25 * -1)
        }
        v3 = -610
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v23:", len(v23), v23)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v20
    _ = v23
    _ = v24
    _ = v25
    _ = v3
}
