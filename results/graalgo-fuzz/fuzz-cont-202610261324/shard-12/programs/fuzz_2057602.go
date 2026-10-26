package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 string }
type GenS2 struct { F0 int64; F1 int16 }
type GenS3 struct { F0 int16 }
type GenS4 struct { F0 uint32; F1 uint64; F2 uint64 }
type GenS5 struct { F0 uint64; F1 uint8; F2 int }

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]int)
    var v3 int16 = 42
    v4 := GenS1{F0: 2, F1: "hello", F2: (("\n\t" + "\n\t") + "a")}
    v3 = v3
    if true {
        v5 := []int{127, 0, 32768}
        switch 1 {
        case 1:
            var v6 uint64 = 616
            _ = v6
        case -1:
            for fi7 := 0; fi7 < 9; fi7++ {
                v8 := GenS2{F0: 900, F1: 1}
                _ = fi7
                _ = v8
            }
        case 127:
            v3 += v3
        default:
            if (false && (v3 == 0)) {
                v9 := GenS3{F0: (^(v3))}
                var v10 uint8 = 214
                var v11 int = 4294967295
                _ = v10
                _ = v11
                _ = v9
            }
        }
        switch ("foo bar" + "a") {
        case "":
            if len(v5) > 0 {
                v5[0] = 2147483647
            }
        case "hello":
            v12 := GenS4{F0: 4294967295, F1: 42, F2: 2}
            _ = v12
        default:
            switch -1 {
            case -9223372036854775808:
                if true {
                    switch -9223372036854775808 {
                    case -9223372036854775808:
                        for fi13 := 0; fi13 < 16; fi13++ {
                            v14 := GenS1{F0: 65535, F1: "foo bar", F2: ("" + ((((("foo bar" + "foo bar") + "a") + "\n\t") + ("\n\t" + "\n\t")) + "\n\t"))}
                            _ = fi13
                            _ = v14
                        }
                    }
                    for fi15 := 0; fi15 < 13; fi15++ {
                        v16 := GenS2{F0: -1, F1: v3}
                        v5 = append(v5, -575)
                        v3 = (v3 % ((v3) | 1))
                        _ = fi15
                        _ = v16
                    }
                    if ("foo bar" > "foo bar") {
                        for fi17 := 0; fi17 < 5; fi17++ {
                            for fi18 := 0; fi18 < 19; fi18++ {
                                v3 = (v3 - v3)
                                _ = fi18
                            }
                            for fi19 := 0; fi19 < 9; fi19++ {
                                _ = len(v5)
                                var v20 uint = 1
                                var v21 float64 = (float64((-2.25 - 1.5)) + 0.0)
                                _ = fi19
                                _ = v20
                                _ = v21
                            }
                            for fi22 := 0; fi22 < 9; fi22++ {
                                var v23 float32 = 1.5
                                fmt.Println(fi17)
                                if len(v5) >= 2 {
                                    v5 = v5[1:len(v5)]
                                }
                                v23 -= (((float32((0.1 * ((((1e10 * 1.5) + ((0.0 + (float64(((-2.25 / (v23 - 0.0)) + (((((v23 + v23) * 3.14159) + float32(0.1)) * v23) / float32(((-2.25 + float64((-2.25 / 1.0))) - float64((float32(1.5) / float32((0.0 - 3.14159))))))))) + ((0.1 / (float64((v23 + float32(0.1))) + 0.0)) * 1.5))) + 0.1)) / float64((v23 / v23))) - ((float64((((1.5 * v23) - float32(float64((v23 + 1.5)))) + v23)) - (0.1 - 0.0)) + (1.5 * (1e10 - 1.5)))))) * 3.14159) - -2.25) / v23)
                                _ = fi22
                                _ = v23
                            }
                            _ = fi17
                        }
                    }
                }
            case 0:
                for fi24 := 0; fi24 < 7; fi24++ {
                    if len(v5) > 0 {
                        v5[0] = fi24
                    }
                    if true {
                        var v25 int32 = -2147483648
                        var v26 float64 = 0.1
                        v3 = (v3 & (v3 | (v3 >> ((42) & 63))))
                        _ = v25
                        _ = v26
                    }
                    switch (fi24 / (((fi24 - 9223372036854775807)) | 1)) {
                    case 127:
                        for fi27 := 0; fi27 < 19; fi27++ {
                            if false {
                                var v28 uint8 = 2
                                v3 -= (v3 - -1)
                                _ = len(v5)
                                _ = v28
                            } else {
                                switch "foo bar" {
                                case "hello":
                                    _ = len(v5)
                                case "foo bar":
                                    fmt.Println(fi27, v3, v3)
                                default:
                                    v29 := GenS1{F0: 632, F1: "", F2: "\n\t"}
                                    _ = v29
                                }
                                v30 := v1[4294967295]
                                switch int(v3) {
                                case 4294967295:
                                    var v31 float64 = (1e10 / (float64((float32((1e10 / (1e10 / 0.1))) - 0.0)) + 1e10))
                                    _ = v31
                                default:
                                    if len(v5) > 0 {
                                        v5[0] = (-(fi27))
                                    }
                                }
                                _ = v30
                            }
                            _ = fi27
                        }
                    case 4294967295:
                        var v32 uint16 = 1
                        _ = v32
                    default:
                        var v33 string = "a"
                        _ = v33
                    }
                    _ = fi24
                }
            default:
                var v34 float32 = 0.0
                _ = v34
            }
        }
        _ = v5
    } else {
        fmt.Println(v3)
        v35 := GenS1{F0: 71, F1: (("\n\t" + "foo bar") + "hello"), F2: "a"}
        _ = v35
    }
    v36 := GenS5{F0: 42, F1: 255, F2: -88}
    fmt.Println(v3)
    fmt.Println(v3)
    var v37 bool = (!(false && false))
    v38 := []int{9223372036854775807, 145, -83, 9223372036854775807}
    if true {
        v37 = (255 == 64)
    }
    v39 := GenS5{F0: 42, F1: 109, F2: 128}
    switch 32768 {
    case 4294967295:
        var v40 float64 = (0.1 - (0.0 * (1e10 - 0.0)))
        _ = v40
    }
    v37 = v37
    v41 := GenS5{F0: 18446744073709551615, F1: 128, F2: -935}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v37:", v37)
    fmt.Println("slicelen v38:", len(v38), v38)
    _ = v1
    _ = v2
    _ = v3
    _ = v36
    _ = v37
    _ = v38
    _ = v39
    _ = v4
    _ = v41
}
