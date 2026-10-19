package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 uint32; F1 int8 }
type GenS4 struct { F0 uint16; F1 int64; F2 int16 }
type GenS5 struct { F0 int32; F1 uint8 }
type GenS6 struct { F0 int64 }
type GenS7 struct { F0 int16; F1 int64 }
type GenS8 struct { F0 string; F1 uint32 }
type GenS9 struct { F0 uint64; F1 string }

func main() {
    var v1 string = "hello"
    v2 := []uint32{}
    var v3 bool = (333 >= 744)
    for fi4 := 0; fi4 < 9; fi4++ {
        v3 = v3
        fmt.Println(v1, v3)
        if (!v3) {
            fmt.Println(fi4, v3)
        }
        _ = fi4
    }
    v5 := GenS1{F0: -1, F1: 70}
    switch v1 {
    case "foo bar":
        v6 := GenS1{F0: -1, F1: 95}
        _ = v6
    default:
        v1 = ((("\n\t" + (((v1 + v1) + (("" + (v1 + v1)) + v1)) + ("\n\t" + v1))) + "") + v1)
    }
    fmt.Println(v1)
    v7 := GenS1{F0: 272, F1: -128}
    if (816 != 18446744073709551615) {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        v2 = append(v2, 2147483648)
    } else {
        var v8 int64 = 9223372036854775807
        switch -9223372036854775808 {
        case 127:
            var v9 uint8 = 0
            _ = v9
        }
        if (((830 > -475) && (v8 > v8)) || true) {
            for fi10 := 0; fi10 < 13; fi10++ {
                for fi11 := 0; fi11 < 14; fi11++ {
                    if len(v2) > 0 {
                        v2[0] = 42
                    }
                    _ = fi11
                }
                _ = fi10
            }
        }
        _ = v8
    }
    for fi12 := 0; fi12 < 7; fi12++ {
        var v13 bool = v3
        _ = fi12
        _ = v13
    }
    for fi14 := 0; fi14 < 18; fi14++ {
        v15 := GenS2{F0: fi14}
        if v3 {
            fmt.Println(fi14, v3, v3)
            for fi16 := 0; fi16 < 9; fi16++ {
                var v17 string = v1
                if len(v2) > 0 {
                    v2[0] = 2
                }
                _ = fi16
                _ = v17
            }
        }
        _ = fi14
        _ = v15
    }
    v18 := GenS2{F0: -1}
    switch "\n\t" {
    case "":
        v1 = ("" + "hello")
    default:
        v1 = "\n\t"
    }
    var v19 uint8 = 127
    v19 -= v19
    if ((42 == 0) && true) {
        var v20 int64 = 9223372036854775807
        _ = v20
    }
    if (0.0 < (1.5 / ((0.1 + (1e10 + 1.5)) - 0.1))) {
        switch -9223372036854775808 {
        case 4294967295:
            for fi21 := 0; fi21 < 20; fi21++ {
                for fi22 := 0; fi22 < 11; fi22++ {
                    v19 = (v19 + v19)
                    switch 9223372036854775807 {
                    case 125:
                        var v23 float32 = 0.0
                        v23 -= ((float32((1.5 * 1.5)) * v23) * v23)
                        _ = v23
                    default:
                        var v24 uint32 = 67
                        _ = v24
                    }
                    if (!(v3 && ((3.14159 * -2.25) <= (float32((0.1 / 0.1)) * 3.14159)))) {
                        v25 := GenS1{F0: -1, F1: 1}
                        v26 := GenS2{F0: (fi22 + fi21)}
                        switch v1 {
                        case "foo bar":
                            var v27 int = (fi22 % ((fi21) | 1))
                            _ = v27
                        case "\n\t":
                            var v28 int16 = 0
                            _ = v28
                        case "a":
                            var v29 string = "\n\t"
                            _ = v29
                        default:
                            if (0 < 13) {
                                v2 = append(v2, 1)
                                switch (fi22 + -561) {
                                case 325:
                                    var v30 int64 = -1
                                    if (!true) {
                                        v31 := GenS3{F0: 42, F1: 18}
                                        _ = v31
                                    } else {
                                        v30 = (v30 - v30)
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                    }
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    fmt.Println(fi21)
                                    _ = v30
                                case -23:
                                    var v32 uint8 = v19
                                    _ = v32
                                case 32768:
                                    switch fi22 {
                                    case 793:
                                        if v3 {
                                            if (true && (42 != 457)) {
                                                if v3 {
                                                    v33 := GenS4{F0: 255, F1: 831, F2: -485}
                                                    v3 = (float32(3.14159) == (((1e10 * 1e10) + -2.25) + -2.25))
                                                    _ = v33
                                                }
                                            }
                                            switch fi22 {
                                            case 9223372036854775807:
                                                v2 = append(v2, 0)
                                            default:
                                                var v34 uint16 = 42
                                                _ = v34
                                            }
                                            v35 := GenS5{F0: 1, F1: v19}
                                            _ = v35
                                        }
                                        if (0 != 608) {
                                            v1 = v1
                                            v2 = append(v2, 2147483648)
                                            var v36 uint = 933
                                            _ = v36
                                        }
                                    }
                                default:
                                    v1 = (v1 + (v1 + (v1 + v1)))
                                }
                                v37 := GenS6{F0: 917}
                                _ = v37
                            }
                        }
                        switch fi21 {
                        case 42:
                            v38 := GenS7{F0: -1, F1: -902}
                            _ = v38
                        }
                        if v3 {
                            fmt.Println(fi22)
                            _ = len(v2)
                            v39 := GenS8{F0: (v1 + v1), F1: 449}
                            _ = v39
                        }
                        _ = v25
                        _ = v26
                    }
                    _ = fi22
                }
                _ = fi21
            }
        case -575:
            var v40 float64 = 1e10
            _ = v40
        case -504:
            v2 = append(v2, 333)
        }
        if (v1 == ((v1 + "") + "hello")) {
            v41 := GenS9{F0: 2, F1: v1}
            v3 = (!(0 >= 42))
            for fi42 := 0; fi42 < 1; fi42++ {
                v43 := GenS4{F0: 32, F1: -9223372036854775808, F2: 32767}
                _ = fi42
                _ = v43
            }
            _ = v41
        }
        _ = len(v2)
    }
    fmt.Println(v3, v19)
    for fi44 := 0; fi44 < 8; fi44++ {
        switch (v1 + v1) {
        case "":
            var v45 uint16 = 65535
            _ = v45
        }
        _ = fi44
    }
    for fi46 := 0; fi46 < 5; fi46++ {
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
        var v47 uint32 = 455
        _ = fi46
        _ = v47
    }
    for fi48 := 0; fi48 < 7; fi48++ {
        var v49 uint64 = 0
        fmt.Println(v49, fi48)
        v50 := GenS5{F0: 568, F1: 1}
        _ = fi48
        _ = v49
        _ = v50
    }
    fmt.Println("v1:", v1)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v5
    _ = v7
}
