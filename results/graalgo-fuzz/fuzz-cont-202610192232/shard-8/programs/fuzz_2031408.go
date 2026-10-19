package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 int64 }
type GenS2 struct { F0 uint; F1 uint32; F2 bool }
type GenS3 struct { F0 uint16; F1 uint8 }
type GenS4 struct { F0 bool; F1 uint32; F2 uint8 }
type GenS5 struct { F0 uint }
type GenS6 struct { F0 uint32; F1 uint64 }
type GenS7 struct { F0 int32 }

func main() {
    var v1 int8 = 1
    v2 := make(map[uint32]bool)
    v3 := make(map[int]string)
    var v4 int8 = -22
    v5 := v3[9223372036854775807]
    v6 := GenS1{F0: (^(v1)), F1: ((v5 + ("" + ((v5 + v5) + (v5 + ((v5 + "") + "a"))))) + ""), F2: -1}
    v7 := v3[-1]
    if true {
        v4 += v1
        v7 = (v5 + "a")
        v4 = 100
    }
    switch (v5 + "a") {
    case "":
        var v8 int16 = 0
        _ = v8
    case "hello":
        var v9 int = 0
        _ = v9
    default:
        var v10 uint32 = 4294967295
        _ = v10
    }
    var v11 int32 = 1
    var v12 uint8 = 42
    v13 := []int{}
    var v14 int16 = -257
    v5 = v5
    if (1 < 4294967295) {
        switch 127 {
        case 9223372036854775807:
            v15 := GenS2{F0: 529, F1: 916, F2: (1 != 0)}
            _ = v15
        case 128:
            var v16 uint8 = v12
            _ = v16
        case 4294967295:
            if (((1.5 * (0.0 / 1.0)) - 0.1) > 0.0) {
                if false {
                    v17 := GenS3{F0: 32767, F1: v12}
                    var v18 uint8 = (^(v12))
                    _ = v17
                    _ = v18
                }
                v12 = (v12 >> ((955) & 63))
            }
        }
        switch "" {
        case "":
            if (0 > (v11 / (((v11 >> ((57) & 63))) | 1))) {
                v12 += v12
                v19 := GenS1{F0: v1, F1: v5, F2: -1}
                _ = v19
            }
        case "hello":
            var v20 uint8 = (^(v12))
            _ = v20
        }
    } else {
        switch v5 {
        case "":
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
        case "foo bar":
            if (((1.5 - (3.14159 + (0.1 / (((1e10 / 1.5) * (0.1 * 1e10)) + -2.25)))) == -2.25) || false) {
                if (!false) {
                    if len(v13) > 0 {
                        v13[0] = 2147483647
                    }
                }
            } else {
                if false {
                    v14 = (v14 - v14)
                } else {
                    for fi21 := 0; fi21 < 12; fi21++ {
                        for fi22 := 0; fi22 < 7; fi22++ {
                            var v23 int = (fi22 << ((0) & 63))
                            v24 := v2[42]
                            _ = fi22
                            _ = v23
                            _ = v24
                        }
                        for fi25 := 0; fi25 < 7; fi25++ {
                            if false {
                                var v26 uint64 = 513
                                v27 := GenS1{F0: v1, F1: v5, F2: 1}
                                switch "hello" {
                                case "\n\t":
                                    v4 -= v1
                                default:
                                    switch v5 {
                                    case "foo bar":
                                        var v28 int64 = -989
                                        _ = v28
                                    case "a":
                                        var v29 uint16 = 505
                                        _ = v29
                                    }
                                }
                                v30 := GenS4{F0: true, F1: 288, F2: (v12 << ((2) & 63))}
                                v31 := GenS5{F0: 4294967295}
                                _ = v26
                                _ = v27
                                _ = v30
                                _ = v31
                            }
                            v11 -= -123
                            var v32 uint64 = 2
                            v33 := GenS6{F0: 554, F1: (v32 / (((^(v32))) | 1))}
                            var v34 bool = (false && true)
                            v13 = append(v13, fi21)
                            _ = fi25
                            _ = v32
                            _ = v33
                            _ = v34
                        }
                        for fi35 := 0; fi35 < 9; fi35++ {
                            v36 := v2[534]
                            _ = fi35
                            _ = v36
                        }
                        _ = fi21
                    }
                }
            }
        case "\n\t":
            if len(v13) >= 2 {
                v13 = v13[1:len(v13)]
            }
        default:
            var v37 float32 = 3.14159
            _ = v37
        }
        if false {
            if false {
                switch 261 {
                case 0:
                    fmt.Println(v14)
                default:
                    _ = len(v2)
                }
                v4 -= (v4 / (((v1 - v4)) | 1))
            }
            v38 := GenS7{F0: v11}
            _ = v38
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
