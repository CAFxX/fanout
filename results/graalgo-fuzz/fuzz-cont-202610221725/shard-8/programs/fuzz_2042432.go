package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint; F2 uint32 }
type GenS2 struct { F0 string; F1 int32 }
type GenS3 struct { F0 int16; F1 uint32; F2 uint32 }
type GenS4 struct { F0 int16; F1 uint8 }
type GenS5 struct { F0 uint; F1 uint8; F2 uint32 }
type GenS6 struct { F0 string; F1 int8 }
type GenS7 struct { F0 bool; F1 uint32 }
type GenS8 struct { F0 uint32 }

func main() {
    var v1 float64 = 3.14159
    var v2 uint8 = 127
    v3 := []int{4294967295, 32768, 1}
    var v4 int16 = 1
    var v5 float32 = ((float32(v1) - 0.1) * 1e10)
    fmt.Println(v5, v4)
    v3 = append(v3, 42)
    for fi6 := 0; fi6 < 2; fi6++ {
        v4 = (v4 ^ (v4 / ((-1) | 1)))
        v2 -= (v2 / (((^(v2))) | 1))
        if false {
            v1 -= (v1 + (v1 * float64(3.14159)))
            _ = len(v3)
            for fi7 := 0; fi7 < 18; fi7++ {
                v8 := GenS1{F0: (v4 - v4), F1: 2, F2: 976}
                _ = fi7
                _ = v8
            }
        }
        _ = fi6
    }
    var v9 int16 = v4
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    v5 += ((0.0 + 0.0) - float32((v1 - (v1 + v1))))
    var v10 bool = true
    v11 := GenS2{F0: "a", F1: 851}
    var v12 float64 = v1
    var v13 float32 = v5
    for fi14 := 0; fi14 < 12; fi14++ {
        var v15 uint64 = uint64(v2)
        _ = fi14
        _ = v15
    }
    if (false || (v10 || false)) {
        v3 = append(v3, 32768)
        for fi16 := 0; fi16 < 19; fi16++ {
            switch 192 {
            case 269:
                v17 := GenS3{F0: v9, F1: 0, F2: 2147483647}
                _ = v17
            case -9223372036854775808:
                var v18 string = "a"
                _ = v18
            default:
                v19 := GenS4{F0: (v9 * -1), F1: 127}
                _ = v19
            }
            v4 += (v4 ^ (-(v9)))
            var v20 int32 = -2147483648
            _ = fi16
            _ = v20
        }
        v3 = append(v3, 127)
    } else {
        var v21 int64 = 9223372036854775807
        v22 := GenS1{F0: (v4 - v4), F1: 852, F2: 1}
        _ = v21
        _ = v22
    }
    var v23 int8 = -128
    for fi24 := 0; fi24 < 16; fi24++ {
        var v25 float64 = (0.1 + ((v12 * 1e10) * v12))
        for fi26 := 0; fi26 < 8; fi26++ {
            v1 += v12
            v9 = (v4 + v4)
            _ = fi26
        }
        if len(v3) >= 2 {
            v3 = v3[1:len(v3)]
        }
        _ = fi24
        _ = v25
    }
    if len(v3) > 0 {
        v3[0] = -935
    }
    switch (("foo bar" + ((((("hello" + "\n\t") + (("a" + "\n\t") + ("\n\t" + ((("a" + "") + "foo bar") + ("foo bar" + "a"))))) + (("\n\t" + "a") + "\n\t")) + "a") + ((("a" + "\n\t") + (("hello" + "a") + "a")) + "\n\t"))) + (("\n\t" + "") + ((("hello" + (("\n\t" + ("foo bar" + ((((("\n\t" + ("a" + "a")) + "") + ("" + "")) + "hello") + "a"))) + "a")) + "") + ("\n\t" + "foo bar")))) {
    case "":
        switch 2147483647 {
        case 1:
            v27 := GenS1{F0: (v9 * (v4 * v9)), F1: 119, F2: 1}
            _ = v27
        case -383:
            for fi28 := 0; fi28 < 9; fi28++ {
                v4 -= 813
                if len(v3) > 0 {
                    v3[0] = fi28
                }
                _ = fi28
            }
        default:
            v5 = (v5 - (float32((float64(v5) * ((1e10 + float64(-2.25)) - (((v1 / (v1 - (v12 / (float64(0.1) * -2.25)))) / 1e10) + (0.1 * 1.5))))) * v13))
        }
    case "hello":
        if (v2 != (v2 >> ((1) & 63))) {
            for fi29 := 0; fi29 < 9; fi29++ {
                if ((255 == 255) && (true && true)) {
                    v1 = ((0.0 + v12) / float64((0.1 - 3.14159)))
                    v30 := GenS1{F0: (v9 ^ v9), F1: 836, F2: 717}
                    v31 := GenS3{F0: 1000, F1: 1, F2: 1}
                    _ = v30
                    _ = v31
                } else {
                    for fi32 := 0; fi32 < 8; fi32++ {
                        var v33 int32 = -929
                        var v34 int8 = 42
                        v34 = 0
                        _ = fi32
                        _ = v33
                        _ = v34
                    }
                    v35 := GenS5{F0: 67, F1: (^(v2)), F2: 0}
                    _ = v35
                }
                if (18446744073709551615 != 0) {
                    v5 -= 3.14159
                    v36 := GenS4{F0: v4, F1: 255}
                    for fi37 := 0; fi37 < 12; fi37++ {
                        var v38 int64 = -9223372036854775808
                        if ((35 != 4294967295) && true) {
                            v1 -= v1
                            v39 := GenS5{F0: 935, F1: 253, F2: 397}
                            fmt.Println(v1, fi29)
                            _ = v39
                        } else {
                            switch fi37 {
                            case -1:
                                var v40 bool = false
                                _ = v40
                            case -837:
                                v41 := GenS6{F0: "foo bar", F1: v23}
                                fmt.Println(v38)
                                if (!v10) {
                                    var v42 int8 = 0
                                    if len(v3) > 0 {
                                        v3[0] = (-(fi37))
                                    }
                                    v1 += v12
                                    _ = v42
                                } else {
                                    var v43 float64 = float64(1e10)
                                    fmt.Println(v43)
                                    v44 := GenS7{F0: (2147483647 != 1), F1: 223}
                                    _ = v43
                                    _ = v44
                                }
                                _ = v41
                            case 2147483647:
                                v2 = v2
                            }
                            var v45 uint32 = 0
                            if v10 {
                                if (-2147483648 != 672) {
                                    var v46 uint64 = 42
                                    _ = v46
                                }
                                if true {
                                    v9 = (v9 ^ 1)
                                    var v47 uint = 4294967295
                                    _ = v47
                                }
                            }
                            v4 = (v4 - v9)
                            _ = len(v3)
                            _ = v45
                        }
                        _ = fi37
                        _ = v38
                    }
                    _ = v36
                } else {
                    for fi48 := 0; fi48 < 17; fi48++ {
                        if v10 {
                            v10 = v10
                            if false {
                                var v49 int32 = 240
                                if v10 {
                                    fmt.Println(fi29)
                                }
                                var v50 int32 = v49
                                _ = v49
                                _ = v50
                            }
                            var v51 uint64 = 384
                            _ = v51
                        }
                        _ = fi48
                    }
                    v23 -= (v23 ^ v23)
                    v52 := GenS2{F0: ((("foo bar" + ("a" + "\n\t")) + "hello") + "foo bar"), F1: 1}
                    _ = v52
                }
                v53 := GenS8{F0: 4294967295}
                _ = fi29
                _ = v53
            }
            if (v2 != (v2 - v2)) {
                var v54 uint64 = 4294967295
                var v55 float32 = (1.5 * ((v5 + ((float32(v1) * -2.25) - ((float32(float64(v13)) + (v13 / v13)) / float32(v12)))) + 0.1))
                _ = v54
                _ = v55
            }
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
