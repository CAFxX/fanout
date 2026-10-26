package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint32 }
type GenS2 struct { F0 int8; F1 uint16 }
type GenS3 struct { F0 uint8 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint; F1 uint16; F2 int8 }
type GenS6 struct { F0 uint32; F1 string }
type GenS7 struct { F0 int16 }

func main() {
    v1 := make(map[uint32]string)
    var v2 int64 = 42
    var v3 float64 = (-2.25 + -2.25)
    for fi4 := 0; fi4 < 8; fi4++ {
        switch fi4 {
        case 127:
            v5 := GenS1{F0: -492, F1: 65535}
            _ = v5
        case 32768:
            _ = len(v1)
        default:
            v6 := GenS1{F0: -1, F1: 65535}
            _ = v6
        }
        var v7 uint8 = 1
        var v8 int = (fi4 * fi4)
        _ = fi4
        _ = v7
        _ = v8
    }
    v3 -= v3
    v3 += v3
    v2 = (^(v2))
    v3 = (-2.25 - v3)
    if false {
        if true {
            v2 = v2
            if (((false || false) || ((v2 | v2) <= -1)) && (-13 < 100)) {
                if (v3 >= (v3 * 0.0)) {
                    for fi9 := 0; fi9 < 13; fi9++ {
                        v3 = v3
                        var v10 uint32 = 967
                        _ = fi9
                        _ = v10
                    }
                    for fi11 := 0; fi11 < 1; fi11++ {
                        var v12 int64 = (v2 - (v2 / ((v2) | 1)))
                        if (true && true) {
                            v3 += float64(1e10)
                            v1[543] = ""
                            for fi13 := 0; fi13 < 4; fi13++ {
                                var v14 int = -9223372036854775808
                                v12 -= (v12 + v2)
                                fmt.Println(v3)
                                _ = fi13
                                _ = v14
                            }
                        }
                        if (v3 < ((v3 + ((v3 - (0.0 / 1.0)) - (3.14159 - v3))) + -2.25)) {
                            if false {
                                fmt.Println(v12, v2, v2)
                            }
                            for fi15 := 0; fi15 < 10; fi15++ {
                                var v16 uint32 = 65535
                                v16 -= (v16 * (v16 | 2))
                                _ = fi15
                                _ = v16
                            }
                            fmt.Println(v3)
                        } else {
                            _ = len(v1)
                            v2 += v12
                            switch ("foo bar" + "a") {
                            case "foo bar":
                                var v17 int = fi11
                                _ = v17
                            case "\n\t":
                                if false {
                                    v18 := GenS2{F0: 127, F1: 255}
                                    var v19 uint8 = 255
                                    _ = v18
                                    _ = v19
                                }
                            case "":
                                fmt.Println(v2)
                            default:
                                for fi20 := 0; fi20 < 18; fi20++ {
                                    var v21 uint64 = 2
                                    fmt.Println(v2, v2)
                                    _ = fi20
                                    _ = v21
                                }
                            }
                        }
                        _ = fi11
                        _ = v12
                    }
                    for fi22 := 0; fi22 < 4; fi22++ {
                        v23 := GenS3{F0: 2}
                        if false {
                            var v24 uint8 = 255
                            v2 += (v2 + (v2 | v2))
                            switch "foo bar" {
                            case "hello":
                                v2 += (v2 ^ (v2 % ((9223372036854775807) | 1)))
                            case "a":
                                if false {
                                    v1[2147483648] = "a"
                                    fmt.Println(v24, v3)
                                    delete(v1, 993)
                                }
                            case "":
                                if false {
                                    v3 -= (v3 - -2.25)
                                }
                            }
                            _ = v24
                        } else {
                            var v25 int16 = -1
                            for fi26 := 0; fi26 < 14; fi26++ {
                                fmt.Println(v2, v2, v3)
                                switch "" {
                                case "\n\t":
                                    var v27 int32 = -332
                                    _ = v27
                                case "a":
                                    v2 = (v2 >> ((2) & 63))
                                default:
                                    var v28 float64 = (v3 * ((float64(1.5) - v3) + ((v3 - (v3 / (v3 - (v3 - float64(float32(-2.25)))))) / 1.0)))
                                    _ = v28
                                }
                                var v29 float64 = (((v3 + v3) * (float64(((0.0 + 1.5) / 1e10)) * ((float64(1e10) / v3) + 0.1))) - -2.25)
                                _ = fi26
                                _ = v29
                            }
                            _ = v25
                        }
                        var v30 uint = 42
                        _ = fi22
                        _ = v23
                        _ = v30
                    }
                } else {
                    fmt.Println(v2, v3)
                }
                fmt.Println(v2, v3)
            }
            v3 = v3
        } else {
            v31 := GenS4{F0: -2147483648}
            var v32 string = ""
            v33 := GenS5{F0: 18446744073709551615, F1: 1, F2: 1}
            _ = v31
            _ = v32
            _ = v33
        }
        fmt.Println(v3, v3)
    } else {
        if (65535 > 1) {
            v34 := []int{128}
            _ = v34
        } else {
            var v35 float32 = 1.5
            v36 := []int{710, -516, -9223372036854775808}
            v37 := GenS6{F0: 249, F1: "a"}
            _ = v35
            _ = v36
            _ = v37
        }
    }
    switch 127 {
    case 890:
        v38 := GenS7{F0: 32767}
        _ = v38
    }
    v39 := GenS7{F0: -32768}
    v2 = -9223372036854775808
    v40 := []int{442}
    v41 := GenS5{F0: 18446744073709551615, F1: 794, F2: 1}
    v1[2147483648] = ((("\n\t" + ("a" + "hello")) + "hello") + "hello")
    switch "" {
    case "\n\t":
        var v42 int64 = (v2 ^ 9223372036854775807)
        _ = v42
    case "a":
        var v43 uint32 = 42
        _ = v43
    case "hello":
        v2 = v2
    }
    if len(v40) > 0 {
        v40[0] = -154
    }
    var v44 int32 = 1
    for fi45 := 0; fi45 < 18; fi45++ {
        var v46 int8 = 127
        var v47 uint = 42
        var v48 uint64 = 4294967295
        _ = fi45
        _ = v46
        _ = v47
        _ = v48
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v40:", len(v40), v40)
    fmt.Println("v44:", v44)
    _ = v1
    _ = v2
    _ = v3
    _ = v39
    _ = v40
    _ = v41
    _ = v44
}
