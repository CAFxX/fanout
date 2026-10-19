package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int16 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint16; F1 uint64 }

func main() {
    var v1 float32 = 0.0
    var v2 float64 = (float64((((float32(float64(3.14159)) - v1) / v1) + v1)) * 0.1)
    var v3 uint64 = 552
    var v4 float64 = ((v2 * (1.5 * ((0.1 + v2) * -2.25))) * (((v2 - ((0.0 + ((v2 / 0.1) + v2)) / float64(-2.25))) / float64(v1)) * v2))
    for fi5 := 0; fi5 < 8; fi5++ {
        switch "\n\t" {
        case "hello":
            v6 := []int{fi5, (fi5 << ((1) & 63)), 9223372036854775807, (fi5 - 42)}
            _ = v6
        case "\n\t":
            fmt.Println(v4, v4, v4)
        case "a":
            var v7 string = "\n\t"
            _ = v7
        }
        _ = fi5
    }
    for fi8 := 0; fi8 < 1; fi8++ {
        switch fi8 {
        case 726:
            var v9 uint8 = 234
            _ = v9
        case -9223372036854775808:
            v10 := GenS1{F0: 32767, F1: 1}
            _ = v10
        case -960:
            fmt.Println(v2, v4, v1)
        }
        _ = fi8
    }
    v3 = v3
    var v11 int8 = 127
    switch ((("hello" + ("a" + "hello")) + "hello") + "") {
    case "hello":
        var v12 uint8 = 2
        _ = v12
    case "":
        var v13 int = 882
        _ = v13
    default:
        v14 := []int{-860, -922, 1}
        _ = v14
    }
    v15 := []int{1, 132, 4294967295, 143}
    switch "\n\t" {
    case "\n\t":
        for fi16 := 0; fi16 < 5; fi16++ {
            var v17 uint64 = 334
            _ = fi16
            _ = v17
        }
    case "a":
        switch 368 {
        case 0:
            fmt.Println(v3)
        default:
            for fi18 := 0; fi18 < 13; fi18++ {
                fmt.Println(v11, v3, v2)
                if (!false) {
                    for fi19 := 0; fi19 < 3; fi19++ {
                        if (v3 < 4294967295) {
                            switch "" {
                            case "foo bar":
                                if (-32768 != 831) {
                                    v20 := GenS1{F0: 942, F1: -897}
                                    _ = v20
                                }
                            }
                        }
                        v1 = (1.5 + (v1 + v1))
                        _ = fi19
                    }
                    for fi21 := 0; fi21 < 8; fi21++ {
                        if true {
                            v2 = ((v4 - (float64(v1) - v4)) / (((0.1 - v4) - (((1.5 - 1.5) / float64(v1)) / (0.0 * v2))) + v2))
                        } else {
                            v22 := GenS1{F0: 2, F1: -247}
                            var v23 float64 = v4
                            var v24 int64 = 0
                            fmt.Println(v23)
                            _ = v22
                            _ = v23
                            _ = v24
                        }
                        if len(v15) > 0 {
                            v15[0] = (fi18 - 127)
                        }
                        switch fi18 {
                        case -1:
                            v25 := GenS2{F0: 74}
                            _ = v25
                        case 0:
                            v1 -= v1
                        case 2147483647:
                            if false {
                                var v26 string = "a"
                                if len(v15) > 0 {
                                    v15[0] = (fi18 % (((fi21 | (fi21 | 127))) | 1))
                                }
                                var v27 string = ("\n\t" + ("\n\t" + v26))
                                _ = v26
                                _ = v27
                            }
                        default:
                            _ = len(v15)
                            var v28 uint8 = 128
                            switch (((("a" + "a") + "hello") + "a") + "hello") {
                            case "foo bar":
                                if len(v15) > 0 {
                                    v15[0] = fi18
                                }
                            case "a":
                                if false {
                                    v15 = append(v15, fi18)
                                }
                            case "":
                                if true {
                                    fmt.Println(v28)
                                    if len(v15) >= 2 {
                                        v15 = v15[1:len(v15)]
                                    }
                                } else {
                                    switch "foo bar" {
                                    case "":
                                        v2 = 3.14159
                                    }
                                    var v29 int32 = 651
                                    var v30 string = "hello"
                                    _ = v29
                                    _ = v30
                                }
                                if false {
                                    var v31 int16 = 42
                                    _ = v31
                                } else {
                                    v15 = append(v15, (fi21 | (fi18 ^ (-(fi21)))))
                                    var v32 uint = uint(v3)
                                    _ = v32
                                }
                                switch "\n\t" {
                                case "foo bar":
                                    var v33 int64 = 9223372036854775807
                                    var v34 uint32 = 0
                                    _ = len(v15)
                                    var v35 int32 = -1
                                    if len(v15) >= 2 {
                                        v15 = v15[1:len(v15)]
                                    }
                                    v36 := GenS2{F0: (v11 << ((4294967295) & 63))}
                                    _ = v33
                                    _ = v34
                                    _ = v35
                                    _ = v36
                                default:
                                    v15 = append(v15, (^(fi18)))
                                }
                                v28 = (v28 * v28)
                            default:
                                v11 += (v11 + 100)
                            }
                            _ = v28
                        }
                        _ = fi21
                    }
                }
                _ = fi18
            }
        }
    default:
        var v37 uint = 129
        _ = v37
    }
    v38 := GenS3{F0: 42, F1: (v3 % ((v3) | 1))}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v15
    _ = v2
    _ = v3
    _ = v38
    _ = v4
}
