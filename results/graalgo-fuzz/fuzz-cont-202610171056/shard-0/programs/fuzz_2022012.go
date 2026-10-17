package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint8; F1 int }
type GenS3 struct { F0 string; F1 uint }
type GenS4 struct { F0 bool; F1 uint; F2 int8 }
type GenS5 struct { F0 uint16 }
type GenS6 struct { F0 uint8; F1 uint64; F2 uint8 }
func genFn0(p0 uint) int64 {
    if false {
        for fi4 := 0; fi4 < 2; fi4++ {
            v5 := GenS1{F0: 2147483647}
            _ = fi4
            _ = v5
        }
    } else {
        v6 := GenS2{F0: 52, F1: -968}
        _ = v6
    }
    v7 := GenS3{F0: ("\n\t" + "foo bar"), F1: p0}
    var v8 uint32 = 65535
    fmt.Println(v8)
    _ = p0
    _ = v7
    _ = v8
    return 997
}
func genFn1() uint {
    var v9 int = 9223372036854775807
    v9 -= (v9 << ((18446744073709551615) & 63))
    _ = v9
    return 42
}

func main() {
    var v1 bool = false
    v2 := []int{546}
    var v3 uint8 = 1
    var v10 int32 = 314
    for fi11 := 0; fi11 < 18; fi11++ {
        switch 267 {
        case 32768:
            v12 := GenS3{F0: ("a" + ("foo bar" + ("foo bar" + "foo bar"))), F1: 1}
            _ = v12
        case 0:
            v13 := GenS4{F0: true, F1: 18446744073709551615, F2: -101}
            _ = v13
        }
        if len(v2) > 0 {
            v2[0] = (fi11 ^ fi11)
        }
        if (2 != 0) {
            fmt.Println(fi11, v3, v3)
            if v1 {
                v10 += v10
                for fi14 := 0; fi14 < 11; fi14++ {
                    v15 := GenS4{F0: v1, F1: 620, F2: 1}
                    _ = fi14
                    _ = v15
                }
            } else {
                for fi16 := 0; fi16 < 10; fi16++ {
                    switch (("" + "\n\t") + "foo bar") {
                    case "":
                        var v17 int8 = 1
                        _ = v17
                    case "\n\t":
                        fmt.Println(fi16, v3, fi16)
                    default:
                        switch fi11 {
                        case 1:
                            switch (fi16 | fi11) {
                            case -9223372036854775808:
                                if v1 {
                                    fmt.Println(fi16)
                                } else {
                                    if v1 {
                                        if v1 {
                                            v10 += (v10 - (^(v10)))
                                            var v18 int = (fi11 & 32767)
                                            var v19 string = ("" + "a")
                                            _ = v18
                                            _ = v19
                                        } else {
                                            var v20 int8 = -63
                                            if v1 {
                                                if (-553 != fi11) {
                                                    if v1 {
                                                        var v21 uint64 = 384
                                                        _ = v21
                                                    }
                                                    var v22 uint32 = 2147483647
                                                    fmt.Println(fi11)
                                                    switch (("" + "\n\t") + "a") {
                                                    case "hello":
                                                        if len(v2) >= 2 {
                                                            v2 = v2[1:len(v2)]
                                                        }
                                                    }
                                                    _ = v22
                                                } else {
                                                    if len(v2) >= 2 {
                                                        v2 = v2[1:len(v2)]
                                                    }
                                                    v1 = (!v1)
                                                    var v23 int8 = 1
                                                    _ = v23
                                                }
                                            } else {
                                                v10 += -2147483648
                                                var v24 int = -1
                                                switch (fi11 & (v24 - 42)) {
                                                case -9223372036854775808:
                                                    switch "\n\t" {
                                                    case "":
                                                        switch "a" {
                                                        case "\n\t":
                                                            v1 = (!v1)
                                                        case "foo bar":
                                                            var v25 string = ""
                                                            _ = v25
                                                        }
                                                    case "\n\t":
                                                        v20 = v20
                                                    case "a":
                                                        var v26 uint = 981
                                                        _ = v26
                                                    default:
                                                        if (-999 <= 0) {
                                                            var v27 int8 = v20
                                                            switch (-(fi16)) {
                                                            case 32767:
                                                                if (1 != 30) {
                                                                    v10 -= v10
                                                                    var v28 float32 = 0.1
                                                                    v2 = append(v2, fi16)
                                                                    _ = v28
                                                                } else {
                                                                    if (128 < v3) {
                                                                        if v1 {
                                                                            v29 := GenS3{F0: "", F1: 42}
                                                                            var v30 uint64 = 563
                                                                            _ = v29
                                                                            _ = v30
                                                                        }
                                                                        var v31 int8 = v27
                                                                        _ = v31
                                                                    }
                                                                    v1 = (0 <= 18446744073709551615)
                                                                    switch v24 {
                                                                    case 32768:
                                                                        v3 = v3
                                                                    default:
                                                                        v1 = (!(3.14159 != 0.0))
                                                                    }
                                                                }
                                                            default:
                                                                if v1 {
                                                                    v10 += -1
                                                                }
                                                            }
                                                            _ = v27
                                                        }
                                                    }
                                                case -2147483648:
                                                    v1 = (genFn0(255) >= genFn0(0))
                                                default:
                                                    v24 += 127
                                                }
                                                if len(v2) > 0 {
                                                    v2[0] = fi16
                                                }
                                                _ = v24
                                            }
                                            _ = v20
                                        }
                                    } else {
                                        v3 = v3
                                        fmt.Println(v3, v10, v10)
                                    }
                                    switch (fi16 << ((genFn1()) & 63)) {
                                    case -9223372036854775808:
                                        var v32 int16 = -32768
                                        _ = v32
                                    case 0:
                                        v3 += 2
                                    case 996:
                                        fmt.Println(v1, v3)
                                    }
                                    if (((3.14159 / ((((float64(1e10) / (float64((-2.25 / 1.5)) * float64(1e10))) + float64((((((((-2.25 + (3.14159 * 1e10)) / (-2.25 / float32(0.1))) - ((((3.14159 + ((1.5 / -2.25) + 0.1)) + (0.1 * 1e10)) - 0.0) + (float32(((float64(1e10) + -2.25) * (1.5 * 3.14159))) - 3.14159))) / 1.0) + 3.14159) * -2.25) / 1e10))) + 0.0) * 3.14159)) / 1e10) == (0.1 - ((0.1 / ((1e10 * 3.14159) * (((-2.25 - (1e10 * -2.25)) + float64(-2.25)) * float64((1e10 + (1.5 / (0.0 - ((((3.14159 - 1.5) * ((float32(float64(0.1)) / 1.5) * (((0.0 + 3.14159) / 1.5) * 1.5))) - 0.0) + (0.0 * 1.5))))))))) / 1e10))) {
                                        v33 := GenS5{F0: 148}
                                        if (18446744073709551615 == 4294967295) {
                                            v10 = v10
                                        }
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        switch fi16 {
                                        case 2147483647:
                                            var v34 float64 = 0.1
                                            _ = v34
                                        case 32768:
                                            _ = len(v2)
                                        }
                                        if len(v2) >= 2 {
                                            v2 = v2[1:len(v2)]
                                        }
                                        _ = v33
                                    }
                                }
                            default:
                                v1 = false
                                var v35 uint = 1
                                _ = v35
                            }
                        case -1:
                            var v36 uint32 = 961
                            _ = v36
                        case 9223372036854775807:
                            v3 = v3
                        }
                    }
                    switch fi16 {
                    case -1:
                        v37 := GenS6{F0: (v3 - v3), F1: 42, F2: (v3 ^ v3)}
                        _ = v37
                    default:
                        var v38 float32 = 0.1
                        _ = v38
                    }
                    _ = fi16
                }
                v39 := GenS1{F0: (fi11 << ((18446744073709551615) & 63))}
                _ = v39
            }
        } else {
            v40 := GenS2{F0: v3, F1: fi11}
            _ = v40
        }
        _ = fi11
    }
    var v41 int = 2147483647
    v41 = 127
    switch 935 {
    case -145:
        var v42 uint = 151
        _ = v42
    }
    v41 += -9223372036854775808
    v41 += (^(v41))
    var v43 uint64 = 42
    v41 = (-(v41))
    for fi44 := 0; fi44 < 13; fi44++ {
        var v45 float32 = 1.5
        v3 += v3
        _ = fi44
        _ = v45
    }
    fmt.Println(v10, v10, v41)
    v46 := GenS6{F0: v3, F1: v43, F2: v3}
    if true {
        fmt.Println(v10, v41)
        fmt.Println(v41)
    }
    switch (v41 & int(v43)) {
    case 42:
        var v47 int32 = (v10 % ((int32(v41)) | 1))
        _ = v47
    case -428:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v48 := GenS5{F0: 255}
    if ((2 < 2147483647) || (v1 || v1)) {
        v49 := GenS4{F0: (!v1), F1: 4294967295, F2: 3}
        v50 := GenS4{F0: true, F1: 255, F2: 0}
        _ = v49
        _ = v50
    } else {
        fmt.Println(v3, v3)
        if v1 {
            v10 = (-(v10))
            fmt.Println(v41)
            _ = len(v2)
        }
        var v51 uint32 = 255
        _ = v51
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v41:", v41)
    fmt.Println("v43:", v43)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v41
    _ = v43
    _ = v46
    _ = v48
}
