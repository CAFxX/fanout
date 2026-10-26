package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int; F2 uint16 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 bool; F1 int64 }
type GenS4 struct { F0 uint; F1 string; F2 bool }
type GenS5 struct { F0 int64; F1 int32 }
type GenS6 struct { F0 int16; F1 int16; F2 bool }
func genFn0() int16 {
    for fi4 := 0; fi4 < 19; fi4++ {
        var v5 uint64 = 458
        _ = fi4
        _ = v5
    }
    if ((1.5 >= float32(float64(1e10))) && (("" == "foo bar") || true)) {
        for fi6 := 0; fi6 < 7; fi6++ {
            var v7 uint8 = 126
            var v8 uint64 = 2
            _ = fi6
            _ = v7
            _ = v8
        }
    } else {
        for fi9 := 0; fi9 < 15; fi9++ {
            fmt.Println(fi9)
            if false {
                fmt.Println(fi9)
                var v10 int32 = -518
                _ = v10
            } else {
                var v11 int = (fi9 | 9223372036854775807)
                v11 -= 421
                v11 = 9223372036854775807
                _ = v11
            }
            _ = fi9
        }
    }
    var v12 bool = true
    _ = v12
    return 32767
}
func genFn1(p0 string) string {
    p0 = p0
    v13 := []int{2147483647, 9223372036854775807, -426, -2147483648}
    _ = p0
    _ = v13
    return ((p0 + "") + p0)
}

func main() {
    var v1 uint16 = 0
    var v2 string = "hello"
    var v3 string = "foo bar"
    v14 := GenS1{F0: 876, F1: -361, F2: (v1 + (^(v1)))}
    for fi15 := 0; fi15 < 18; fi15++ {
        if (579 <= 0) {
            v16 := GenS2{F0: 88}
            var v17 uint = 675
            var v18 int = (-(fi15))
            _ = v16
            _ = v17
            _ = v18
        } else {
            v3 = v2
            v19 := []int{(fi15 % ((fi15) | 1)), -9223372036854775808}
            _ = v19
        }
        v20 := GenS2{F0: 159}
        if (1 >= 2147483647) {
            if (-9223372036854775808 > -659) {
                v21 := GenS1{F0: 4294967295, F1: fi15, F2: v1}
                switch ((v3 + "") + (("foo bar" + ((v3 + v2) + ((v2 + (v2 + v3)) + (("\n\t" + v3) + v2)))) + v2)) {
                case "\n\t":
                    if false {
                        v3 = ((("hello" + v2) + v2) + "")
                        v3 = ((v3 + (v2 + "")) + v2)
                        v22 := []int{(fi15 ^ fi15), 32768, fi15, fi15}
                        _ = v22
                    } else {
                        var v23 float32 = 0.1
                        v1 += 2
                        v24 := GenS1{F0: 0, F1: 577, F2: (v1 >> ((367) & 63))}
                        _ = v23
                        _ = v24
                    }
                case "hello":
                    v25 := []int{-343, (-(fi15))}
                    _ = v25
                case "a":
                    var v26 bool = true
                    _ = v26
                default:
                    var v27 int = (^(fi15))
                    _ = v27
                }
                _ = v21
            }
            if (686 <= 4294967295) {
                v28 := []int{9223372036854775807, (fi15 << ((0) & 63))}
                switch "foo bar" {
                case "foo bar":
                    if len(v28) > 0 {
                        v28[0] = fi15
                    }
                default:
                    fmt.Println(v2)
                }
                var v29 string = (v3 + (("\n\t" + v3) + v2))
                _ = v28
                _ = v29
            }
            v2 = "hello"
        }
        _ = fi15
        _ = v20
    }
    if ((true || true) && true) {
        v30 := []int{0}
        fmt.Println(v3)
        v3 = v3
        _ = v30
    }
    v31 := []int{387, -9223372036854775808, 509, 812}
    for fi32 := 0; fi32 < 9; fi32++ {
        var v33 float64 = 3.14159
        _ = fi32
        _ = v33
    }
    switch v2 {
    case "foo bar":
        v34 := GenS2{F0: 255}
        _ = v34
    case "hello":
        if len(v31) >= 2 {
            v31 = v31[1:len(v31)]
        }
    }
    v1 = (^(v1))
    var v35 int8 = 100
    v35 += (v35 ^ (v35 + v35))
    v35 = v35
    var v36 int32 = 1
    if len(v31) > 0 {
        v31[0] = 824
    }
    for fi37 := 0; fi37 < 4; fi37++ {
        if (265 > 2147483647) {
            switch fi37 {
            case 950:
                switch fi37 {
                case 2147483647:
                    v36 = (v36 + (v36 << ((0) & 63)))
                default:
                    v35 -= v35
                }
            case 128:
                v38 := GenS3{F0: (-32768 != -32768), F1: int64(v36)}
                _ = v38
            default:
                v2 = v3
            }
            v36 -= v36
        }
        for fi39 := 0; fi39 < 11; fi39++ {
            v2 = (("\n\t" + "hello") + v2)
            var v40 int = fi39
            _ = fi39
            _ = v40
        }
        var v41 float32 = float32((3.14159 - (1.5 / float64((1e10 * ((((0.1 / 1.5) + (1.5 / 0.1)) + 1.5) + (((-2.25 + 3.14159) / (1.5 * -2.25)) + -2.25)))))))
        _ = fi37
        _ = v41
    }
    v36 = (v36 % ((0) | 1))
    v31 = append(v31, 4294967295)
    var v42 int32 = v36
    v43 := GenS2{F0: 2}
    for fi44 := 0; fi44 < 15; fi44++ {
        v1 = 0
        var v45 float32 = float32(1.5)
        if (false && (!true)) {
            if ((true || false) || true) {
                if len(v31) >= 2 {
                    v31 = v31[1:len(v31)]
                }
                switch v2 {
                case "a":
                    if len(v31) >= 2 {
                        v31 = v31[1:len(v31)]
                    }
                case "hello":
                    var v46 int16 = 870
                    _ = v46
                }
                var v47 float32 = v45
                _ = v47
            } else {
                if false {
                    v48 := GenS4{F0: 4294967295, F1: v2, F2: false}
                    if (!false) {
                        var v49 float64 = 1.5
                        v35 += (v35 >> ((4294967295) & 63))
                        var v50 bool = (2 == 594)
                        _ = v49
                        _ = v50
                    } else {
                        v51 := GenS5{F0: -1, F1: 1}
                        v52 := GenS6{F0: 1, F1: -32768, F2: (!(true || true))}
                        _ = v51
                        _ = v52
                    }
                    _ = v48
                } else {
                    var v53 int = (fi44 & fi44)
                    v54 := GenS6{F0: -32768, F1: 163, F2: true}
                    _ = v53
                    _ = v54
                }
                v31 = append(v31, 32767)
                if true {
                    var v55 uint64 = 293
                    var v56 uint = 1
                    _ = v55
                    _ = v56
                } else {
                    for fi57 := 0; fi57 < 6; fi57++ {
                        switch (fi44 | (fi44 << ((1) & 63))) {
                        case 32767:
                            var v58 uint = 226
                            _ = v58
                        case 983:
                            var v59 int = -346
                            _ = v59
                        default:
                            switch ((v2 + v3) + ("\n\t" + "hello")) {
                            case "foo bar":
                                if len(v31) >= 2 {
                                    v31 = v31[1:len(v31)]
                                }
                            case "a":
                                var v60 uint64 = 735
                                if (v35 >= -128) {
                                    var v61 bool = true
                                    _ = v61
                                }
                                _ = v60
                            default:
                                if len(v31) >= 2 {
                                    v31 = v31[1:len(v31)]
                                }
                            }
                        }
                        _ = fi57
                    }
                    v36 += 711
                }
            }
        }
        _ = fi44
        _ = v45
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v31:", len(v31), v31)
    fmt.Println("v35:", v35)
    fmt.Println("v36:", v36)
    fmt.Println("v42:", v42)
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v31
    _ = v35
    _ = v36
    _ = v42
    _ = v43
}
