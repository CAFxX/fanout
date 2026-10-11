package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 uint32 }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int16; F1 bool; F2 uint64 }
func genFn0(p0 string) int8 {
    p0 = ("" + p0)
    p0 = (p0 + "foo bar")
    _ = p0
    return 0
}

func main() {
    var v1 float32 = 1e10
    var v2 int64 = -9223372036854775808
    var v3 string = ""
    v4 := []int{771}
    var v5 bool = (((true || (!("\n\t" != "hello"))) || (-2.25 > float64(float32((-2.25 - (float64((0.0 + 3.14159)) - (((1.5 / (0.0 / (1e10 * (float64(v1) - (-2.25 * 0.1))))) - 3.14159) + ((0.0 / 1.5) - 3.14159)))))))) && false)
    v6 := GenS1{F0: false}
    if false {
        if (!(0 != 104)) {
            for fi7 := 0; fi7 < 11; fi7++ {
                if true {
                    if (v5 && v5) {
                        _ = len(v4)
                        if (!v5) {
                            if len(v4) > 0 {
                                v4[0] = fi7
                            }
                            v1 -= float32(1.5)
                            for fi8 := 0; fi8 < 8; fi8++ {
                                if true {
                                    v9 := GenS2{F0: 195}
                                    _ = v9
                                } else {
                                    if (237 >= 830) {
                                        v10 := GenS2{F0: 0}
                                        fmt.Println(v2, v5)
                                        switch v3 {
                                        case "foo bar":
                                            _ = len(v4)
                                        case "":
                                            var v11 uint = 255
                                            _ = v11
                                        }
                                        if (fi7 == fi7) {
                                            var v12 string = v3
                                            _ = v12
                                        } else {
                                            if len(v4) >= 2 {
                                                v4 = v4[1:len(v4)]
                                            }
                                            v5 = (v1 < float32(0.1))
                                        }
                                        _ = v10
                                    } else {
                                        var v13 int16 = 312
                                        v14 := GenS3{F0: 42}
                                        if v5 {
                                            var v15 int64 = v2
                                            _ = v15
                                        } else {
                                            v16 := GenS4{F0: (4294967295 < 141)}
                                            var v17 int8 = 42
                                            _ = v16
                                            _ = v17
                                        }
                                        _ = v13
                                        _ = v14
                                    }
                                }
                                switch (-(fi7)) {
                                case 1:
                                    v1 += float32(1.5)
                                case -2147483648:
                                    var v18 int32 = 42
                                    _ = v18
                                }
                                _ = fi8
                            }
                        }
                    } else {
                        if (!false) {
                            switch v3 {
                            case "a":
                                var v19 float64 = -2.25
                                _ = v19
                            case "":
                                v4 = append(v4, -2147483648)
                            case "hello":
                                if len(v4) >= 2 {
                                    v4 = v4[1:len(v4)]
                                }
                            }
                        } else {
                            var v20 int32 = -659
                            _ = v20
                        }
                    }
                    if v5 {
                        v4 = append(v4, -624)
                    } else {
                        var v21 int32 = -2147483648
                        _ = v21
                    }
                } else {
                    var v22 uint64 = 42
                    _ = v22
                }
                for fi23 := 0; fi23 < 10; fi23++ {
                    var v24 float64 = (((1e10 / (float64(3.14159) * float64((((((1e10 / (3.14159 + v1)) * (v1 / (((v1 - v1) + v1) - (((0.1 / (float32(1.5) - ((v1 / v1) / -2.25))) - ((v1 - (0.0 * 1e10)) * 0.0)) * (v1 + v1))))) + (v1 + 1.5)) + 0.0) * ((((v1 - (((3.14159 + float32((0.1 + (float64(v1) * ((3.14159 / ((float64(v1) * (0.1 / float64(((((v1 / (((v1 - (-2.25 + v1)) * 1e10) + (float32(-2.25) + float32(1e10)))) - v1) - -2.25) / 1.5)))) / (((-2.25 / float64(v1)) / ((-2.25 - (0.1 / 1.5)) - 1e10)) * -2.25))) * (0.1 - (((1.5 + ((-2.25 - (0.1 - (((0.0 - 1.5) * 0.0) * 1e10))) * ((0.0 / (0.0 + ((3.14159 + ((1e10 * 0.0) - (0.0 / ((-2.25 + (0.0 * 3.14159)) - 0.1)))) - (((-2.25 + -2.25) - (0.0 + (0.0 / 1.5))) - 1.5)))) + 3.14159))) * 1.5) - 1.5))))))) / (((((v1 - 1.5) - v1) + 0.1) + 1.5) / (((1e10 * v1) - v1) * (1.5 - (float32(float64((1e10 / 0.1))) - (float32(0.0) - 0.0)))))) - v1)) - 3.14159) * (3.14159 * (v1 * (v1 / (float32((float64((v1 / v1)) + 0.0)) / (v1 * ((((v1 * ((v1 - ((3.14159 / v1) - (v1 - -2.25))) + v1)) - v1) - v1) * (v1 + 1e10)))))))) * float32((1e10 + (float64((0.0 - 0.0)) + 0.0)))))))) - 1e10) + 3.14159)
                    switch v3 {
                    case "foo bar":
                        switch 0 {
                        case 127:
                            var v25 int16 = 32767
                            _ = v25
                        case 128:
                            switch -1 {
                            case 187:
                                fmt.Println(v5)
                            case 2147483647:
                                var v26 int8 = 2
                                _ = v26
                            default:
                                fmt.Println(fi7, v3)
                            }
                        default:
                            if true {
                                switch fi7 {
                                case 1:
                                    v27 := GenS3{F0: 42}
                                    _ = v27
                                default:
                                    fmt.Println(fi23, v5, v24)
                                }
                            }
                        }
                    }
                    _ = fi23
                    _ = v24
                }
                _ = fi7
            }
        }
        if len(v4) > 0 {
            v4[0] = -51
        }
    }
    fmt.Println(v5, v5, v3)
    var v28 int16 = -1
    for fi29 := 0; fi29 < 3; fi29++ {
        if len(v4) > 0 {
            v4[0] = fi29
        }
        for fi30 := 0; fi30 < 7; fi30++ {
            v3 = ""
            _ = fi30
        }
        switch "" {
        case "a":
            v31 := GenS1{F0: (1 == -2147483648)}
            _ = v31
        case "foo bar":
            v2 += (v2 % (((v2 ^ int64(v28))) | 1))
        case "hello":
            v1 -= 0.1
        default:
            if (((!v5) || v5) && (!(((((("hello" + v3) + v3) < (v3 + v3)) && true) || ((-101 == -58) || (!false))) && (v1 == v1)))) {
                for fi32 := 0; fi32 < 12; fi32++ {
                    if v5 {
                        switch -543 {
                        case -866:
                            var v33 uint16 = 0
                            v34 := GenS2{F0: 143}
                            _ = v33
                            _ = v34
                        case -322:
                            var v35 int8 = 60
                            _ = v35
                        default:
                            v36 := GenS4{F0: v5}
                            _ = v36
                        }
                        if len(v4) > 0 {
                            v4[0] = (fi32 - 32768)
                        }
                        v37 := GenS5{F0: (v28 + (-(v28))), F1: (42 > -55), F2: 4294967295}
                        _ = v37
                    } else {
                        v28 -= v28
                        v38 := GenS5{F0: v28, F1: (9223372036854775807 > v2), F2: 120}
                        _ = v38
                    }
                    v4 = append(v4, -9223372036854775808)
                    v28 = 42
                    _ = fi32
                }
            }
        }
        _ = fi29
    }
    if len(v4) > 0 {
        v4[0] = 128
    }
    switch ("a" + v3) {
    case "":
        var v39 int = 32767
        _ = v39
    case "foo bar":
        switch "\n\t" {
        case "hello":
            v40 := GenS4{F0: false}
            _ = v40
        case "foo bar":
            switch v3 {
            case "foo bar":
                for fi41 := 0; fi41 < 13; fi41++ {
                    var v42 int16 = -32768
                    _ = fi41
                    _ = v42
                }
            case "":
                switch v3 {
                case "":
                    v5 = (!v5)
                case "hello":
                    var v43 uint = 542
                    _ = v43
                }
            default:
                v1 = ((1.5 / (v1 * ((v1 + float32(1.5)) * v1))) + ((3.14159 + ((1.5 * v1) / 3.14159)) / (-2.25 + v1)))
            }
        }
    case "hello":
        fmt.Println(v3)
    default:
        if false {
            var v44 float32 = (0.1 - v1)
            _ = v44
        } else {
            fmt.Println(v3, v5)
            v3 = ("foo bar" + "hello")
            var v45 int64 = (-(v2))
            _ = v45
        }
    }
    var v46 bool = v5
    v2 -= 9223372036854775807
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v46:", v46)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v46
    _ = v5
    _ = v6
}
