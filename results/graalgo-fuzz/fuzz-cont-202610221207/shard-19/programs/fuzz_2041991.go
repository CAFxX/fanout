package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint32; F1 uint }
type GenS3 struct { F0 uint; F1 uint8; F2 int }
type GenS4 struct { F0 int16; F1 int8 }
type GenS5 struct { F0 int8 }
type GenS6 struct { F0 int16; F1 int16; F2 string }
type GenS7 struct { F0 bool }
type GenS8 struct { F0 uint32; F1 uint }

func main() {
    var v1 uint64 = 4294967295
    v2 := make(map[string]uint32)
    v3 := make(map[string]int)
    v4 := []uint32{216, 4294967295, 978}
    v5 := make(map[int]uint32)
    v6 := make(map[uint32]string)
    fmt.Println(v1)
    var v7 float32 = -2.25
    for fi8 := 0; fi8 < 8; fi8++ {
        var v9 uint64 = (v1 & 0)
        var v10 int16 = 32767
        if false {
            if len(v4) > 0 {
                v4[0] = 0
            }
            v7 += float32(0.0)
            for fi11 := 0; fi11 < 7; fi11++ {
                v7 = (v7 + float32(((-2.25 + (-2.25 + (((0.0 + (-2.25 / (1e10 + 0.0))) - 0.1) / (((0.0 * -2.25) / 1e10) - (float64(1e10) + 1.5))))) - -2.25)))
                _ = fi11
            }
        } else {
            v9 = v9
            if (true || (3.14159 != (((v7 / v7) + float32(3.14159)) + (((float32(1.5) - v7) + float32(1e10)) / ((((0.1 + (1e10 - v7)) * 1e10) * 1.5) / ((float32(3.14159) - v7) + v7)))))) {
                if ((!((^(v10)) < 1000)) || ("\n\t" != "\n\t")) {
                    var v12 uint32 = 1
                    _ = v12
                }
                switch "hello" {
                case "foo bar":
                    fmt.Println(v1, v10)
                case "\n\t":
                    for fi13 := 0; fi13 < 5; fi13++ {
                        fmt.Println(fi8)
                        var v14 uint32 = 2
                        v1 += 802
                        _ = fi13
                        _ = v14
                    }
                default:
                    if (!(1 != 255)) {
                        for fi15 := 0; fi15 < 4; fi15++ {
                            v3["foo bar"] = fi8
                            var v16 float64 = 0.0
                            if len(v4) >= 2 {
                                v4 = v4[1:len(v4)]
                            }
                            _ = fi15
                            _ = v16
                        }
                    } else {
                        v17 := GenS1{F0: "a"}
                        switch 9223372036854775807 {
                        case 128:
                            if (!(-2147483648 > 0)) {
                                var v18 float64 = 3.14159
                                switch 425 {
                                case -655:
                                    v19 := GenS2{F0: 2147483648, F1: 400}
                                    _ = v19
                                default:
                                    v20 := GenS1{F0: "\n\t"}
                                    _ = v20
                                }
                                _ = v18
                            }
                        default:
                            v6[2147483647] = "a"
                        }
                        v10 = 0
                        _ = v17
                    }
                }
            }
            if true {
                var v21 uint8 = 128
                _ = v21
            } else {
                delete(v3, "a")
                var v22 uint16 = 0
                if false {
                    v23 := GenS1{F0: ("" + "a")}
                    _ = v23
                }
                _ = v22
            }
        }
        _ = fi8
        _ = v10
        _ = v9
    }
    v7 = v7
    for fi24 := 0; fi24 < 13; fi24++ {
        switch "" {
        case "":
            var v25 int = (fi24 * fi24)
            _ = v25
        case "\n\t":
            _ = len(v4)
        default:
            var v26 uint32 = 335
            _ = v26
        }
        if len(v4) > 0 {
            v4[0] = 1
        }
        _ = fi24
    }
    v5[1] = 2147483647
    for fi27 := 0; fi27 < 19; fi27++ {
        _ = len(v6)
        for fi28 := 0; fi28 < 9; fi28++ {
            if true {
                if (!false) {
                    v29 := GenS3{F0: 581, F1: 101, F2: fi27}
                    v1 += v1
                    _ = v29
                } else {
                    var v30 float64 = 0.1
                    switch "a" {
                    case "hello":
                        fmt.Println(v1, v7)
                    case "a":
                        v31 := GenS4{F0: 1, F1: -1}
                        if (fi27 != fi27) {
                            _ = len(v2)
                            v30 = v30
                            v32 := GenS5{F0: 42}
                            _ = len(v6)
                            _ = len(v3)
                            _ = v32
                        } else {
                            var v33 uint8 = 128
                            switch ((("a" + "\n\t") + ("foo bar" + "a")) + "\n\t") {
                            case "":
                                v30 += v30
                            default:
                                if true {
                                    var v34 int8 = -126
                                    v35 := GenS6{F0: 32767, F1: 393, F2: "hello"}
                                    v1 += 4294967295
                                    _ = v34
                                    _ = v35
                                } else {
                                    if len(v4) >= 2 {
                                        v4 = v4[1:len(v4)]
                                    }
                                    var v36 int16 = 1000
                                    v33 -= 129
                                    _ = v36
                                }
                            }
                            var v37 bool = ((42 <= int64(fi27)) && true)
                            _ = v33
                            _ = v37
                        }
                        _ = v31
                    }
                    if true {
                        v6[2147483647] = "foo bar"
                    }
                    _ = v30
                }
                fmt.Println(v7)
            }
            switch "" {
            case "\n\t":
                var v38 int64 = 455
                _ = v38
            }
            _ = fi28
        }
        v39 := GenS7{F0: ("\n\t" >= "foo bar")}
        _ = fi27
        _ = v39
    }
    if false {
        delete(v6, 381)
        v40 := GenS8{F0: 42, F1: 65535}
        fmt.Println(v7, v7)
        _ = v40
    } else {
        if true {
            var v41 uint8 = 0
            fmt.Println(v1)
            _ = v41
        } else {
            switch -753 {
            case 103:
                v7 += float32(float64((v7 + 1.5)))
            case -335:
                for fi42 := 0; fi42 < 12; fi42++ {
                    var v43 uint64 = 0
                    _ = fi42
                    _ = v43
                }
            case 660:
                if true {
                    v7 = v7
                    for fi44 := 0; fi44 < 7; fi44++ {
                        if true {
                            for fi45 := 0; fi45 < 1; fi45++ {
                                if (false || (float32(((((3.14159 * ((1.5 * (0.1 / (0.0 + (1.5 + (-2.25 + (0.1 * 1e10)))))) / (-2.25 * ((-2.25 + ((0.0 * (1e10 * 0.0)) / float64(-2.25))) + (-2.25 - float64(((v7 - v7) * 0.1))))))) - -2.25) + -2.25) * 3.14159)) < 0.0)) {
                                    var v46 int = -1
                                    _ = v46
                                } else {
                                    var v47 string = "a"
                                    v7 += v7
                                    if len(v4) > 0 {
                                        v4[0] = 126
                                    }
                                    _ = v47
                                }
                                _ = fi45
                            }
                            for fi48 := 0; fi48 < 14; fi48++ {
                                v4 = append(v4, 601)
                                _ = fi48
                            }
                            if false {
                                fmt.Println(v7, v1)
                            } else {
                                v1 -= v1
                                fmt.Println(v7)
                                for fi49 := 0; fi49 < 9; fi49++ {
                                    _ = len(v4)
                                    _ = fi49
                                }
                            }
                        }
                        v4 = append(v4, 184)
                        _ = fi44
                    }
                    var v50 int8 = 100
                    _ = v50
                } else {
                    switch 32767 {
                    case 32767:
                        v7 += v7
                    }
                    v7 += float32(-2.25)
                }
            }
        }
        var v51 uint64 = 4294967295
        fmt.Println(v51, v51)
        _ = v51
    }
    fmt.Println(v1, v1)
    var v52 float64 = ((float64(1.5) + ((3.14159 * 1.5) / ((0.1 / 0.1) / (1e10 * 1e10)))) - (0.0 * (float64(v7) + 3.14159)))
    var v53 uint16 = 2
    fmt.Println(v53, v7)
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v52:", v52)
    fmt.Println("v53:", v53)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v52
    _ = v53
    _ = v6
    _ = v7
}
