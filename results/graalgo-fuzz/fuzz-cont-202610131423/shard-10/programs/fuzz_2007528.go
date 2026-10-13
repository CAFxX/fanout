package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint16; F2 uint }
type GenS2 struct { F0 int8; F1 uint16; F2 uint64 }
type GenS3 struct { F0 string; F1 int64 }
func genFn0(p0 uint8, p1 uint, p2 uint) int32 {
    var v7 int = 2147483647
    var v8 int32 = -2147483648
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 443
}
func genFn1(p0 int64, p1 uint32, p2 string) string {
    v9 := []int{963, -9223372036854775808, -9223372036854775808}
    for fi10 := 0; fi10 < 3; fi10++ {
        p0 = p0
        for fi11 := 0; fi11 < 9; fi11++ {
            v9 = append(v9, (fi11 << ((93) & 63)))
            if false {
                v12 := GenS1{F0: 396, F1: 0, F2: 18446744073709551615}
                _ = v12
            } else {
                switch fi11 {
                case 820:
                    if (2 != 0) {
                        p2 = "\n\t"
                        fmt.Println(p1)
                    } else {
                        var v13 float64 = float64((-2.25 + 0.0))
                        var v14 uint16 = 2
                        var v15 int = fi10
                        _ = v13
                        _ = v14
                        _ = v15
                    }
                case 128:
                    if (!((float32(1.5) * (0.0 + 0.0)) <= ((float32((3.14159 - float64(-2.25))) * 0.1) * float32(1.5)))) {
                        fmt.Println(p2)
                    } else {
                        var v16 int32 = 42
                        var v17 int = fi11
                        if (false || (true || (false || (127 < -75)))) {
                            var v18 uint64 = 936
                            fmt.Println(v17, v16, fi10)
                            v19 := GenS1{F0: 4294967295, F1: 255, F2: 255}
                            p1 = (p1 * (p1 + (p1 - p1)))
                            _ = v18
                            _ = v19
                        }
                        _ = v16
                        _ = v17
                    }
                case 706:
                    var v20 float32 = -2.25
                    _ = v20
                default:
                    if (2 > 962) {
                        p1 += p1
                    } else {
                        var v21 string = "a"
                        p2 = "a"
                        _ = v21
                    }
                }
                p0 += p0
            }
            _ = fi11
        }
        _ = fi10
    }
    var v22 float32 = 0.0
    var v23 int8 = 114
    _ = p0
    _ = p1
    _ = p2
    _ = v22
    _ = v23
    _ = v9
    return ((p2 + p2) + p2)
}

func main() {
    v1 := make(map[int]bool)
    v2 := []int{32768}
    v3 := make(map[int]string)
    var v4 float32 = 0.0
    var v5 int = 766
    var v6 int8 = 127
    if false {
        _ = len(v3)
    } else {
        switch int(v6) {
        case 9223372036854775807:
            var v24 uint = 0
            _ = v24
        }
    }
    if ((!false) && (true || (4294967295 < 728))) {
        switch v5 {
        case 99:
            var v25 float32 = 1e10
            _ = v25
        case -1:
            var v26 int = v5
            _ = v26
        }
        var v27 int16 = 0
        switch v5 {
        case -103:
            for fi28 := 0; fi28 < 2; fi28++ {
                v29 := v3[fi28]
                v29 = "\n\t"
                for fi30 := 0; fi30 < 18; fi30++ {
                    v31 := GenS1{F0: 295, F1: 8, F2: 1}
                    v6 = (v6 * v6)
                    v32 := GenS2{F0: -128, F1: 65535, F2: 472}
                    v27 = (^(v27))
                    switch (^(fi28)) {
                    case 32767:
                        v6 = v6
                    case 32768:
                        if (!(!(!(!(((3.14159 / (((((((v4 / ((0.1 + ((v4 + -2.25) / v4)) + (v4 + v4))) * v4) * float32(-2.25)) + v4) * -2.25) * (v4 - v4)) * v4)) + (v4 / v4)) < ((-2.25 + ((((float32((3.14159 - (1.5 - (0.1 + (0.0 + 3.14159))))) - float32(float64(float32((0.0 / (3.14159 - (1.5 * 1.5))))))) - v4) - 1e10) * -2.25)) + 3.14159)))))) {
                            if false {
                                var v33 uint32 = 2147483647
                                _ = v33
                            }
                            v4 = v4
                            var v34 uint64 = 0
                            _ = v34
                        }
                    case -2147483648:
                        var v35 float64 = 1.5
                        _ = v35
                    default:
                        v36 := GenS2{F0: (v6 & (^(v6))), F1: 42, F2: 2}
                        v27 = (-(v27))
                        _ = v36
                    }
                    _ = fi30
                    _ = v31
                    _ = v32
                }
                _ = fi28
                _ = v29
            }
        }
        _ = v27
    } else {
        var v37 int = (^(v5))
        var v38 int8 = (-(v6))
        _ = v37
        _ = v38
    }
    delete(v3, (-(v5)))
    v39 := GenS1{F0: 255, F1: 459, F2: 2}
    for fi40 := 0; fi40 < 4; fi40++ {
        fmt.Println(fi40, v5, v5)
        v41 := GenS3{F0: "\n\t", F1: -1}
        _ = fi40
        _ = v41
    }
    _ = len(v2)
    v4 -= 3.14159
    if false {
        fmt.Println(v5, v5)
    } else {
        for fi42 := 0; fi42 < 20; fi42++ {
            fmt.Println(v5, v4)
            for fi43 := 0; fi43 < 14; fi43++ {
                if (!(652 >= 33)) {
                    v44 := GenS3{F0: "\n\t", F1: -32}
                    var v45 int64 = 42
                    _ = v44
                    _ = v45
                }
                var v46 int = fi42
                v6 = -63
                _ = fi43
                _ = v46
            }
            _ = fi42
        }
    }
    switch "hello" {
    case "a":
        _ = len(v2)
    case "":
        for fi47 := 0; fi47 < 19; fi47++ {
            v4 = v4
            _ = fi47
        }
    default:
        if true {
            if (!false) {
                if ((v5 >> ((710) & 63)) <= (v5 >> ((18446744073709551615) & 63))) {
                    fmt.Println(v4)
                } else {
                    v6 = v6
                    for fi48 := 0; fi48 < 15; fi48++ {
                        for fi49 := 0; fi49 < 18; fi49++ {
                            if ((-2.25 - 1e10) > 1.5) {
                                v50 := GenS1{F0: 255, F1: 997, F2: 201}
                                switch 128 {
                                case 530:
                                    fmt.Println(v4, fi48)
                                case 0:
                                    if false {
                                        fmt.Println(fi48, v5, v5)
                                    }
                                }
                                _ = len(v1)
                                _ = v50
                            }
                            _ = fi49
                        }
                        var v51 float32 = float32(3.14159)
                        _ = fi48
                        _ = v51
                    }
                }
                if ((v6 <= (v6 - -128)) && false) {
                    if (!((v5 << ((1) & 63)) < v5)) {
                        for fi52 := 0; fi52 < 7; fi52++ {
                            switch (fi52 % ((2147483647) | 1)) {
                            case 980:
                                v4 = v4
                            case 0:
                                v6 += 1
                            }
                            v5 = -870
                            _ = fi52
                        }
                        _ = len(v2)
                        for fi53 := 0; fi53 < 14; fi53++ {
                            switch (^(fi53)) {
                            case 127:
                                fmt.Println(v4, v5, v6)
                            case -9223372036854775808:
                                var v54 int16 = 381
                                _ = v54
                            case 4294967295:
                                var v55 uint32 = 0
                                _ = v55
                            default:
                                if true {
                                    if len(v2) > 0 {
                                        v2[0] = int(v6)
                                    }
                                } else {
                                    v56 := GenS2{F0: v6, F1: 42, F2: 0}
                                    for fi57 := 0; fi57 < 10; fi57++ {
                                        v6 = (v6 << ((18446744073709551615) & 63))
                                        _ = fi57
                                    }
                                    _ = v56
                                }
                            }
                            _ = fi53
                        }
                    } else {
                        var v58 uint16 = 116
                        _ = v58
                    }
                    if (-1 > -404) {
                        var v59 uint32 = 2147483648
                        var v60 int32 = 0
                        _ = v59
                        _ = v60
                    } else {
                        if (-2147483648 >= 681) {
                            v61 := GenS3{F0: "", F1: 1}
                            var v62 int16 = -796
                            v63 := GenS1{F0: 1, F1: 876, F2: 0}
                            _ = v61
                            _ = v62
                            _ = v63
                        }
                        for fi64 := 0; fi64 < 11; fi64++ {
                            if len(v2) > 0 {
                                v2[0] = v5
                            }
                            var v65 int32 = 1
                            _ = fi64
                            _ = v65
                        }
                    }
                } else {
                    fmt.Println(v5, v4, v6)
                    var v66 int16 = -1
                    _ = v66
                }
                v4 -= float32(-2.25)
            } else {
                switch (("\n\t" + ("a" + "\n\t")) + "a") {
                case "a":
                    _ = len(v2)
                case "foo bar":
                    fmt.Println(v5, v4)
                default:
                    var v67 int32 = genFn0(1, 1, 18446744073709551615)
                    _ = v67
                }
            }
        } else {
            var v68 uint32 = 65535
            _ = v68
        }
    }
    v5 -= v5
    v69 := GenS2{F0: (-(v6)), F1: 2, F2: 1}
    switch "foo bar" {
    case "a":
        for fi70 := 0; fi70 < 11; fi70++ {
            fmt.Println(fi70, v6)
            v4 -= float32(float64(float32(((1.5 / 1.0) + float64(v4)))))
            _ = fi70
        }
    default:
        v4 -= 3.14159
    }
    var v71 int8 = v6
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v71:", v71)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v5
    _ = v6
    _ = v69
    _ = v71
}
