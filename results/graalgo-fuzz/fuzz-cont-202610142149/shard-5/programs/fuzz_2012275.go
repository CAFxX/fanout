package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 uint16 }
type GenS2 struct { F0 uint16; F1 uint32; F2 uint16 }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 int64; F1 bool }
type GenS5 struct { F0 int64; F1 string; F2 bool }
func genFn0() int64 {
    var v4 uint32 = 731
    v5 := []int{2147483647, 0, 1}
    v5 = append(v5, 0)
    if false {
        var v6 int32 = 0
        _ = v6
    } else {
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        for fi7 := 0; fi7 < 6; fi7++ {
            if false {
                switch "foo bar" {
                case "foo bar":
                    fmt.Println(v4)
                case "hello":
                    v4 -= 590
                case "a":
                    switch ("\n\t" + "") {
                    case "a":
                        for fi8 := 0; fi8 < 16; fi8++ {
                            fmt.Println(v4, fi8, v4)
                            _ = fi8
                        }
                    case "\n\t":
                        switch "\n\t" {
                        case "hello":
                            for fi9 := 0; fi9 < 2; fi9++ {
                                if true {
                                    fmt.Println(fi7, fi7)
                                    var v10 int8 = -128
                                    _ = v10
                                } else {
                                    var v11 uint8 = 42
                                    _ = len(v5)
                                    if len(v5) > 0 {
                                        v5[0] = fi9
                                    }
                                    _ = v11
                                }
                                if true {
                                    v4 -= 454
                                    if (-2.25 == (0.0 + 0.1)) {
                                        v4 += v4
                                    } else {
                                        fmt.Println(v4, fi7, fi7)
                                        v4 = v4
                                        v5 = append(v5, 1)
                                        v4 = (v4 % (((v4 & v4)) | 1))
                                        var v12 int = 2147483647
                                        var v13 string = "\n\t"
                                        _ = v12
                                        _ = v13
                                    }
                                } else {
                                    var v14 int8 = 54
                                    v4 = (v4 << ((0) & 63))
                                    _ = len(v5)
                                    var v15 float32 = 0.1
                                    _ = v14
                                    _ = v15
                                }
                                v16 := GenS1{F0: ("\n\t" + "a"), F1: false, F2: 1}
                                if (!false) {
                                    if false {
                                        switch (fi9 / ((fi7) | 1)) {
                                        case -2147483648:
                                            v5 = append(v5, fi7)
                                        default:
                                            v4 -= v4
                                        }
                                        if true {
                                            v4 += (^(v4))
                                            fmt.Println(fi7, fi7)
                                            v4 -= v4
                                        } else {
                                            var v17 uint = 2
                                            if len(v5) >= 2 {
                                                v5 = v5[1:len(v5)]
                                            }
                                            _ = v17
                                        }
                                        v4 -= v4
                                    }
                                    var v18 string = "a"
                                    fmt.Println(fi9, fi7, fi7)
                                    _ = v18
                                } else {
                                    v4 = 500
                                    v19 := GenS1{F0: "hello", F1: true, F2: 650}
                                    _ = v19
                                }
                                _ = fi9
                                _ = v16
                            }
                        case "\n\t":
                            var v20 int64 = -178
                            _ = v20
                        }
                    case "foo bar":
                        v4 = v4
                    default:
                        for fi21 := 0; fi21 < 10; fi21++ {
                            v22 := GenS1{F0: (("a" + "") + ("foo bar" + "")), F1: (((float32(float64(((((float32(3.14159) + float32(0.1)) / 1.5) * 1e10) + (1.5 * 0.1)))) * (0.1 - 0.1)) / (float32(((0.1 - float64(((((1.5 / 1e10) - 1e10) - -2.25) - 3.14159))) * float64(0.1))) * (float32(((float64(0.1) / -2.25) * 0.0)) - (float32(3.14159) - 1.5)))) < (((0.0 + ((((1.5 + 3.14159) + -2.25) / 1.0) * float32((1.5 + (0.1 + 0.0))))) + 1e10) / 1.5)), F2: 766}
                            v4 += v4
                            _ = fi21
                            _ = v22
                        }
                    }
                }
                var v23 string = "\n\t"
                switch "" {
                case "hello":
                    var v24 int8 = 82
                    _ = v24
                case "":
                    v23 = v23
                }
                _ = v23
            }
            fmt.Println(v4)
            for fi25 := 0; fi25 < 9; fi25++ {
                v26 := GenS2{F0: 65535, F1: 0, F2: 255}
                var v27 uint64 = 2
                _ = fi25
                _ = v26
                _ = v27
            }
            _ = fi7
        }
        var v28 int8 = -1
        _ = v28
    }
    fmt.Println(v4)
    _ = v4
    _ = v5
    return 9223372036854775807
}
func genFn1(p0 string) uint {
    var v29 int32 = 0
    v30 := GenS1{F0: ("hello" + p0), F1: true, F2: 0}
    _ = p0
    _ = v29
    _ = v30
    return 0
}

func main() {
    v1 := make(map[int]string)
    v2 := []string{}
    v3 := []int{32767, 9223372036854775807}
    v31 := GenS3{F0: (105 == 255)}
    for fi32 := 0; fi32 < 16; fi32++ {
        switch fi32 {
        case 32768:
            delete(v1, fi32)
        }
        _ = fi32
    }
    v33 := GenS4{F0: -9223372036854775808, F1: false}
    if false {
        _ = len(v2)
        v34 := GenS2{F0: 0, F1: 1, F2: 32767}
        for fi35 := 0; fi35 < 18; fi35++ {
            switch "foo bar" {
            case "hello":
                v36 := GenS3{F0: false}
                _ = v36
            case "\n\t":
                v1[-2147483648] = "a"
            default:
                var v37 int8 = 120
                _ = v37
            }
            if true {
                fmt.Println(fi35)
                var v38 int16 = 0
                fmt.Println(fi35)
                _ = v38
            }
            if false {
                _ = len(v1)
            }
            _ = fi35
        }
        _ = v34
    }
    var v39 float64 = 3.14159
    v39 -= (1e10 + v39)
    v39 += ((0.0 - 1.5) / v39)
    if true {
        for fi40 := 0; fi40 < 19; fi40++ {
            switch "" {
            case "\n\t":
                v41 := GenS5{F0: 42, F1: "", F2: false}
                _ = v41
            default:
                v39 = 1.5
            }
            _ = fi40
        }
        var v42 uint = 210
        _ = v42
    } else {
        fmt.Println(v39)
        v43 := GenS3{F0: true}
        if len(v3) > 0 {
            v3[0] = 128
        }
        _ = v43
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v33
    _ = v39
}
