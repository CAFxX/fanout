package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 int32 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 string; F1 uint; F2 int16 }

func main() {
    var v1 uint64 = 1
    v2 := make(map[uint32]int)
    var v3 int8 = 42
    var v4 float32 = 0.1
    v5 := []int{-519, -872, 32768}
    var v6 float64 = -2.25
    _ = len(v5)
    switch "foo bar" {
    case "":
        if true {
            delete(v2, 713)
        } else {
            v2[0] = -2147483648
            var v7 uint = 0
            _ = v7
        }
    case "hello":
        switch "hello" {
        case "a":
            v8 := GenS1{F0: 42, F1: 0, F2: -2147483648}
            _ = v8
        case "foo bar":
            switch "\n\t" {
            case "foo bar":
                v9 := GenS2{F0: (v1 / ((v1) | 1))}
                _ = v9
            case "":
                var v10 float32 = v4
                _ = v10
            }
        case "":
            v5 = append(v5, 4294967295)
        default:
            var v11 int = 1
            _ = v11
        }
    case "foo bar":
        if (1 != 9223372036854775807) {
            if (255 < 2) {
                var v12 uint32 = 845
                delete(v2, (^(v12)))
                var v13 uint64 = (v1 + (v1 >> ((155) & 63)))
                _ = v12
                _ = v13
            } else {
                if ((v3 << ((58) & 63)) >= v3) {
                    var v14 string = "hello"
                    _ = v14
                } else {
                    if false {
                        var v15 bool = true
                        v4 = ((1e10 / 1e10) * (-2.25 - v4))
                        switch int(v3) {
                        case 4294967295:
                            v16 := GenS3{F0: ("\n\t" + "a"), F1: 65535, F2: -32768}
                            _ = v16
                        case 9223372036854775807:
                            var v17 int64 = 1
                            _ = v17
                        case -119:
                            switch 645 {
                            case 768:
                                for fi18 := 0; fi18 < 8; fi18++ {
                                    if true {
                                        var v19 int16 = 32767
                                        _ = v19
                                    } else {
                                        var v20 int = fi18
                                        var v21 uint64 = (v1 % ((v1) | 1))
                                        var v22 int16 = 654
                                        _ = v20
                                        _ = v21
                                        _ = v22
                                    }
                                    var v23 uint64 = (v1 ^ v1)
                                    v15 = (1 < 42)
                                    _ = fi18
                                    _ = v23
                                }
                            default:
                                var v24 int = 0
                                _ = v24
                            }
                        default:
                            v2[964] = 516
                        }
                        _ = v15
                    }
                }
                var v25 uint = 562
                _ = v25
            }
            fmt.Println(v3, v3)
        }
    }
    v4 = (((v4 - (0.1 + (float32((v6 / ((0.0 + -2.25) / float64((v4 * (0.0 - v4)))))) - v4))) * (0.1 - ((((v4 - v4) / (0.0 * v4)) + -2.25) / v4))) * 0.1)
    if true {
        v4 += (v4 * v4)
    }
    v4 -= float32(float64(v4))
    fmt.Println(v6, v4, v3)
    var v26 int32 = 1
    v27 := GenS2{F0: 2}
    for fi28 := 0; fi28 < 15; fi28++ {
        fmt.Println(v1)
        _ = fi28
    }
    _ = len(v2)
    if (false || ((v26 > -161) || false)) {
        switch ("a" + "\n\t") {
        case "\n\t":
            v29 := GenS1{F0: 1000, F1: 255, F2: (v26 ^ 42)}
            _ = v29
        case "hello":
            switch ((("hello" + (("\n\t" + "\n\t") + "\n\t")) + ("hello" + "foo bar")) + (("" + "\n\t") + "foo bar")) {
            case "\n\t":
                if true {
                    fmt.Println(v26, v4)
                    for fi30 := 0; fi30 < 9; fi30++ {
                        switch fi30 {
                        case 4294967295:
                            _ = len(v2)
                        case -350:
                            v4 = v4
                        }
                        for fi31 := 0; fi31 < 12; fi31++ {
                            switch "\n\t" {
                            case "\n\t":
                                if len(v5) > 0 {
                                    v5[0] = fi31
                                }
                            case "hello":
                                var v32 int64 = -20
                                _ = v32
                            case "a":
                                switch "a" {
                                case "foo bar":
                                    v26 -= -2147483648
                                case "hello":
                                    if len(v5) >= 2 {
                                        v5 = v5[1:len(v5)]
                                    }
                                case "":
                                    v1 += v1
                                default:
                                    switch (fi30 % ((fi30) | 1)) {
                                    case -9223372036854775808:
                                        v33 := GenS2{F0: 624}
                                        _ = v33
                                    default:
                                        if len(v5) >= 2 {
                                            v5 = v5[1:len(v5)]
                                        }
                                    }
                                }
                            }
                            _ = fi31
                        }
                        v34 := v2[1]
                        _ = fi30
                        _ = v34
                    }
                    v4 += v4
                }
            default:
                delete(v2, 255)
            }
        case "":
            v26 = int32(v3)
        }
    }
    var v35 uint32 = 2
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v35:", v35)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v35
    _ = v4
    _ = v5
    _ = v6
}
