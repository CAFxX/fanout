package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int64 }
type GenS2 struct { F0 uint; F1 uint }
type GenS3 struct { F0 int16; F1 uint32 }

func main() {
    var v1 float32 = 1.5
    v2 := []uint32{209, 930}
    v3 := []uint32{}
    v4 := []uint32{}
    var v5 uint = 2
    v5 -= (v5 ^ (v5 & v5))
    v5 -= 642
    v1 = (1.5 * float32((-2.25 + (0.1 - ((1e10 * ((1.5 * 0.0) + 0.1)) - 1.5)))))
    v5 = v5
    v1 -= v1
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    for fi6 := 0; fi6 < 19; fi6++ {
        v5 -= v5
        v1 = (v1 + v1)
        _ = fi6
    }
    if true {
        var v7 uint64 = 18446744073709551615
        for fi8 := 0; fi8 < 14; fi8++ {
            switch "a" {
            case "":
                for fi9 := 0; fi9 < 3; fi9++ {
                    if len(v3) >= 2 {
                        v3 = v3[1:len(v3)]
                    }
                    _ = fi9
                }
            case "\n\t":
                fmt.Println(v7, v5, fi8)
            }
            _ = fi8
        }
        var v10 float32 = v1
        _ = v10
        _ = v7
    }
    v11 := GenS1{F0: 0, F1: -1}
    switch ("" + "hello") {
    case "hello":
        v1 += 1.5
    case "":
        if true {
            switch "hello" {
            case "hello":
                switch ("foo bar" + ("hello" + "a")) {
                case "foo bar":
                    fmt.Println(v5)
                case "hello":
                    if false {
                        for fi12 := 0; fi12 < 8; fi12++ {
                            if (true && (!false)) {
                                v13 := GenS1{F0: 175, F1: 513}
                                _ = v13
                            }
                            var v14 uint8 = 0
                            v14 += 92
                            _ = fi12
                            _ = v14
                        }
                    } else {
                        if false {
                            if true {
                                if (1 == 42) {
                                    v15 := GenS2{F0: v5, F1: 255}
                                    _ = v15
                                } else {
                                    var v16 string = ""
                                    if (!true) {
                                        var v17 int = 1
                                        for fi18 := 0; fi18 < 6; fi18++ {
                                            for fi19 := 0; fi19 < 17; fi19++ {
                                                fmt.Println(v1, fi18)
                                                var v20 int = 4294967295
                                                _ = fi19
                                                _ = v20
                                            }
                                            _ = fi18
                                        }
                                        v21 := GenS1{F0: -1, F1: -611}
                                        _ = v17
                                        _ = v21
                                    }
                                    _ = len(v2)
                                    _ = v16
                                }
                                v1 += ((1e10 + v1) / ((float32(0.0) * v1) * v1))
                                var v22 float64 = 0.1
                                _ = v22
                            }
                            v1 -= v1
                            var v23 int32 = -395
                            _ = v23
                        } else {
                            v5 += v5
                            fmt.Println(v5, v5)
                        }
                        v24 := GenS2{F0: v5, F1: v5}
                        _ = v24
                    }
                case "a":
                    v25 := GenS3{F0: -970, F1: 65535}
                    _ = v25
                default:
                    v26 := GenS2{F0: v5, F1: v5}
                    _ = v26
                }
            case "":
                fmt.Println(v1, v5)
            }
            v27 := GenS3{F0: 129, F1: 65535}
            _ = len(v3)
            _ = v27
        } else {
            var v28 int = -88
            for fi29 := 0; fi29 < 19; fi29++ {
                for fi30 := 0; fi30 < 15; fi30++ {
                    if (v5 > v5) {
                        v1 = v1
                    }
                    var v31 int64 = -523
                    if false {
                        var v32 float32 = v1
                        v5 = 265
                        if (-304 > -1) {
                            switch -226 {
                            case 4294967295:
                                v1 = ((3.14159 - v32) + ((((-2.25 + (1.5 * 0.1)) + v32) / (-2.25 - 3.14159)) - v1))
                            }
                        } else {
                            if true {
                                v1 = (v32 / 1e10)
                            } else {
                                var v33 float64 = ((0.1 / 1.5) - -2.25)
                                var v34 uint16 = 65535
                                v35 := GenS3{F0: 232, F1: 2}
                                _ = v33
                                _ = v34
                                _ = v35
                            }
                        }
                        _ = v32
                    }
                    _ = fi30
                    _ = v31
                }
                v3 = append(v3, 221)
                var v36 uint64 = 2
                _ = fi29
                _ = v36
            }
            v37 := GenS1{F0: -136, F1: -552}
            _ = v28
            _ = v37
        }
    case "\n\t":
        var v38 float32 = (3.14159 / float32((1e10 - (float64(v1) - ((0.0 + 0.0) / 1.0)))))
        _ = v38
    default:
        v1 -= (float32(0.0) - (v1 / v1))
    }
    if len(v2) > 0 {
        v2[0] = 2147483647
    }
    var v39 int8 = -128
    v1 += (0.1 * (0.0 - (1e10 * (v1 - v1))))
    v2 = append(v2, 65535)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v39:", v39)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v39
    _ = v4
    _ = v5
}
