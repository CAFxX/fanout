package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int8; F2 uint16 }
type GenS2 struct { F0 uint16; F1 int }
type GenS3 struct { F0 bool; F1 uint16; F2 int16 }
type GenS4 struct { F0 uint8; F1 int32; F2 uint64 }
type GenS5 struct { F0 int8; F1 bool }

func main() {
    var v1 uint = 18446744073709551615
    v2 := []string{("\n\t" + "foo bar")}
    var v3 float64 = ((0.0 * 0.1) - 1e10)
    v4 := []int{-677, 42}
    switch int(v1) {
    case 128:
        if len(v2) >= 2 {
            v2 = v2[1:len(v2)]
        }
    }
    var v5 float32 = 3.14159
    v5 = 1e10
    v1 -= v1
    if true {
        switch "hello" {
        case "":
            fmt.Println(v3, v3)
        default:
            if ((9223372036854775807 > 0) || true) {
                v1 += v1
            } else {
                v1 -= v1
            }
        }
    } else {
        if len(v4) > 0 {
            v4[0] = 4294967295
        }
    }
    switch "\n\t" {
    case "hello":
        var v6 int8 = 42
        _ = v6
    case "":
        switch 9223372036854775807 {
        case 469:
            v1 -= (v1 / ((v1) | 1))
        default:
            if (0 >= -314) {
                if true {
                    for fi7 := 0; fi7 < 7; fi7++ {
                        if false {
                            v3 += v3
                            var v8 uint64 = 18446744073709551615
                            var v9 int16 = 348
                            _ = v8
                            _ = v9
                        } else {
                            for fi10 := 0; fi10 < 3; fi10++ {
                                v11 := GenS1{F0: 42, F1: 127, F2: 2}
                                switch "foo bar" {
                                case "a":
                                    var v12 int64 = 42
                                    _ = v12
                                case "foo bar":
                                    var v13 int32 = 181
                                    _ = v13
                                }
                                if ((1e10 - v3) == 1e10) {
                                    if false {
                                        v3 = v3
                                        fmt.Println(v5)
                                        v14 := GenS2{F0: 1, F1: (fi10 << ((989) & 63))}
                                        _ = v14
                                    } else {
                                        var v15 uint8 = 0
                                        v3 = float64(v5)
                                        _ = v15
                                    }
                                    if false {
                                        if (!false) {
                                            var v16 int = fi10
                                            var v17 uint64 = 18446744073709551615
                                            var v18 int32 = 309
                                            _ = v16
                                            _ = v17
                                            _ = v18
                                        }
                                    } else {
                                        fmt.Println(v1, v3, v5)
                                        v19 := GenS3{F0: false, F1: uint16(v1), F2: 1}
                                        _ = v19
                                    }
                                    v1 = 42
                                }
                                v5 += ((0.1 - ((3.14159 / 0.1) / ((((v5 + (3.14159 * (0.0 * 0.1))) * ((v5 - 0.1) * v5)) + (v5 - v5)) * v5))) + (v5 - v5))
                                if (false || ((^(v1)) != (v1 << ((0) & 63)))) {
                                    v20 := GenS2{F0: 815, F1: -9223372036854775808}
                                    _ = v20
                                } else {
                                    fmt.Println(v1, fi10)
                                }
                                _ = fi10
                                _ = v11
                            }
                        }
                        _ = fi7
                    }
                    if false {
                        v3 = (v3 / float64(v5))
                    } else {
                        for fi21 := 0; fi21 < 9; fi21++ {
                            for fi22 := 0; fi22 < 16; fi22++ {
                                v3 += (-2.25 - float64((((1.5 + ((v5 + (-2.25 / 3.14159)) / v5)) / v5) * -2.25)))
                                _ = fi22
                            }
                            _ = fi21
                        }
                        var v23 float32 = (v5 * 1e10)
                        _ = v23
                    }
                    var v24 uint = 42
                    _ = v24
                }
            } else {
                switch -2147483648 {
                case 32768:
                    v4 = append(v4, 1)
                default:
                    for fi25 := 0; fi25 < 9; fi25++ {
                        v5 += 3.14159
                        switch 9223372036854775807 {
                        case -451:
                            var v26 uint64 = 18446744073709551615
                            _ = v26
                        case 42:
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                        default:
                            _ = len(v4)
                        }
                        _ = fi25
                    }
                }
            }
        }
    case "foo bar":
        v3 -= v3
    }
    v5 = v5
    var v27 uint64 = 741
    v28 := GenS2{F0: uint16(v1), F1: 32768}
    var v29 uint = v1
    fmt.Println(v5, v1)
    v27 += (v27 + (v27 % ((v27) | 1)))
    v30 := GenS4{F0: 228, F1: 1, F2: (v27 / (((v27 * (v27 << (((^(v29))) & 63)))) | 1))}
    for fi31 := 0; fi31 < 11; fi31++ {
        if false {
            if (65535 < 0) {
                switch "\n\t" {
                case "":
                    var v32 int16 = -32768
                    _ = v32
                case "foo bar":
                    var v33 uint8 = 198
                    _ = v33
                default:
                    switch -2147483648 {
                    case 128:
                        switch fi31 {
                        case -471:
                            v2 = append(v2, "a")
                        case -9223372036854775808:
                            fmt.Println(fi31, v27)
                        default:
                            var v34 uint8 = 52
                            _ = v34
                        }
                    case 32767:
                        var v35 int = 32768
                        _ = v35
                    }
                }
                v29 = (v29 & v29)
                v1 = (^(v1))
            } else {
                var v36 int32 = 398
                _ = v36
            }
            if (896 == 927) {
                v5 = float32(1e10)
                if (true || false) {
                    v29 += 42
                    v5 = 0.1
                }
            }
            var v37 uint = (v1 * (v1 & 4294967295))
            _ = v37
        }
        if false {
            var v38 int16 = 1
            var v39 bool = ((((v3 - (v3 * v3)) <= v3) && (!(!true))) && false)
            _ = v38
            _ = v39
        } else {
            fmt.Println(fi31)
            v40 := GenS5{F0: 42, F1: (3.14159 == (1e10 / ((0.1 / float32(1.5)) / v5)))}
            switch (fi31 * fi31) {
            case -2147483648:
                var v41 uint32 = 42
                _ = v41
            case -1:
                v29 += (v1 & 985)
            case -541:
                v4 = append(v4, fi31)
            }
            _ = v40
        }
        _ = fi31
    }
    fmt.Println(v5)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v27:", v27)
    fmt.Println("v29:", v29)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v2
    _ = v27
    _ = v28
    _ = v29
    _ = v3
    _ = v30
    _ = v4
    _ = v5
}
