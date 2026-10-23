package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int }
type GenS2 struct { F0 int16; F1 int64 }
type GenS3 struct { F0 int64; F1 int32 }
type GenS4 struct { F0 int32; F1 int }
type GenS5 struct { F0 int16 }
type GenS6 struct { F0 uint32; F1 uint64 }
type GenS7 struct { F0 uint8; F1 int32; F2 int32 }
func genFn0(p0 uint32, p1 int16, p2 int64) uint8 {
    for fi6 := 0; fi6 < 5; fi6++ {
        p1 = -1
        p2 = 0
        _ = fi6
    }
    switch 775 {
    case -150:
        var v7 float64 = 0.0
        _ = v7
    case 2147483647:
        var v8 uint32 = p0
        _ = v8
    case 593:
        if true {
            v9 := GenS1{F0: 222, F1: 4294967295}
            var v10 int64 = 42
            _ = v10
            _ = v9
        } else {
            var v11 int32 = -379
            p0 = (p0 + (p0 & (^(p0))))
            _ = v11
        }
    }
    _ = p0
    _ = p1
    _ = p2
    return 1
}

func main() {
    var v1 uint = 0
    v2 := make(map[string]int)
    v3 := []string{"foo bar", ("" + "\n\t")}
    var v4 uint32 = 304
    var v5 string = ("\n\t" + "hello")
    v2[v5] = 32767
    if ((true && false) || false) {
        v12 := GenS2{F0: 42, F1: 1}
        _ = v12
    }
    if (1 == 255) {
        switch "" {
        case "":
            v13 := GenS3{F0: int64(v1), F1: 1}
            _ = v13
        case "hello":
            switch -279 {
            case 32768:
                for fi14 := 0; fi14 < 12; fi14++ {
                    if (2 == 42) {
                        var v15 uint64 = 190
                        _ = v15
                    } else {
                        switch int(v1) {
                        case 9223372036854775807:
                            v1 = 42
                        case 127:
                            var v16 int64 = 9223372036854775807
                            _ = v16
                        case 4294967295:
                            v1 += v1
                        }
                        for fi17 := 0; fi17 < 15; fi17++ {
                            var v18 int32 = 0
                            v5 = ""
                            if (-1 >= 0) {
                                v19 := GenS3{F0: 42, F1: (v18 >> (((v1 / (((v1 >> ((255) & 63))) | 1))) & 63))}
                                _ = len(v3)
                                var v20 float32 = float32((((3.14159 + ((1.5 + 0.0) * 3.14159)) * 0.0) - 0.1))
                                v1 = (v1 % (((v1 - (^(v1)))) | 1))
                                var v21 int16 = 0
                                _ = v19
                                _ = v20
                                _ = v21
                            }
                            _ = fi17
                            _ = v18
                        }
                        delete(v2, v5)
                    }
                    for fi22 := 0; fi22 < 16; fi22++ {
                        v1 += 0
                        v23 := GenS4{F0: -396, F1: (fi14 + -1)}
                        var v24 int32 = -2147483648
                        _ = fi22
                        _ = v23
                        _ = v24
                    }
                    for fi25 := 0; fi25 < 6; fi25++ {
                        v26 := GenS5{F0: -32768}
                        _ = fi25
                        _ = v26
                    }
                    _ = fi14
                }
            case 4294967295:
                if (42 == int8(v1)) {
                    var v27 int8 = 80
                    if ((true || ((^(v27)) > v27)) && false) {
                        v5 = v5
                        if true {
                            var v28 string = v5
                            _ = v28
                        }
                    }
                    _ = v27
                }
            case 1:
                switch -9223372036854775808 {
                case 588:
                    v29 := GenS1{F0: 337, F1: 4294967295}
                    _ = v29
                case 0:
                    v30 := GenS3{F0: 0, F1: 606}
                    _ = v30
                case 4294967295:
                    fmt.Println(v4, v5, v4)
                default:
                    var v31 uint32 = 0
                    _ = v31
                }
            default:
                v2[v5] = 42
            }
        case "a":
            if (!true) {
                if (127 > 2147483647) {
                    var v32 uint64 = 64
                    v33 := GenS3{F0: -118, F1: 1}
                    var v34 uint16 = 0
                    _ = v32
                    _ = v33
                    _ = v34
                }
                fmt.Println(v4, v5, v4)
            } else {
                for fi35 := 0; fi35 < 2; fi35++ {
                    var v36 float64 = (0.0 + ((1.5 * (1.5 / 1e10)) * ((-2.25 - 0.0) + (3.14159 * (0.1 * 0.0)))))
                    _ = fi35
                    _ = v36
                }
                if false {
                    fmt.Println(v4, v5)
                    fmt.Println(v1, v1, v1)
                }
            }
        }
        if (false && (!true)) {
            v5 = "a"
            delete(v2, "\n\t")
            for fi37 := 0; fi37 < 2; fi37++ {
                v38 := GenS5{F0: 1000}
                _ = fi37
                _ = v38
            }
        } else {
            var v39 int32 = -690
            _ = v39
        }
        switch -9223372036854775808 {
        case 1:
            var v40 uint16 = 2
            _ = v40
        default:
            v41 := GenS6{F0: (v4 / (((^(v4))) | 1)), F1: 4294967295}
            _ = v41
        }
    } else {
        v4 = v4
    }
    v1 = v1
    var v42 uint16 = 255
    for fi43 := 0; fi43 < 19; fi43++ {
        var v44 float64 = (-2.25 - 0.1)
        switch fi43 {
        case 1:
            for fi45 := 0; fi45 < 15; fi45++ {
                v4 -= v4
                var v46 float32 = -2.25
                _ = fi45
                _ = v46
            }
        case 566:
            v47 := GenS7{F0: 255, F1: -2147483648, F2: 708}
            _ = v47
        default:
            for fi48 := 0; fi48 < 13; fi48++ {
                v49 := GenS6{F0: v4, F1: 182}
                fmt.Println(v1, v44)
                var v50 float64 = v44
                _ = fi48
                _ = v49
                _ = v50
            }
        }
        if false {
            v44 += 0.1
        }
        _ = fi43
        _ = v44
    }
    v3 = append(v3, v5)
    var v51 uint16 = 65535
    v52 := GenS2{F0: 1000, F1: 9223372036854775807}
    var v53 int32 = 0
    v51 = v51
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v42:", v42)
    fmt.Println("v5:", v5)
    fmt.Println("v51:", v51)
    fmt.Println("v53:", v53)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v42
    _ = v5
    _ = v51
    _ = v52
    _ = v53
}
