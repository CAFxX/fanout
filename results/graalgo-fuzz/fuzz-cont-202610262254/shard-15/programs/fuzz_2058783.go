package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint32; F1 uint64 }
type GenS4 struct { F0 uint32; F1 int8 }

func main() {
    var v1 int32 = 2147483647
    var v2 uint64 = 620
    v3 := make(map[int]string)
    var v4 int = 32767
    var v5 float64 = float64(1.5)
    v6 := GenS1{F0: 915, F1: 1}
    var v7 bool = false
    v4 -= (v4 * (v4 % (((v4 / (((v4 + 1)) | 1))) | 1)))
    v8 := GenS2{F0: -70}
    v9 := []int{int(v1), 928, v4}
    switch "a" {
    case "foo bar":
        var v10 bool = v7
        _ = v10
    default:
        var v11 uint8 = 255
        _ = v11
    }
    v1 += -1
    v12 := GenS3{F0: 0, F1: v2}
    if ((42 == 0) && v7) {
        for fi13 := 0; fi13 < 10; fi13++ {
            if (true && (float32(v5) == -2.25)) {
                fmt.Println(v5, v4, v1)
            } else {
                v5 = v5
            }
            for fi14 := 0; fi14 < 10; fi14++ {
                v3[fi14] = "a"
                fmt.Println(v5)
                var v15 int = fi13
                _ = fi14
                _ = v15
            }
            switch 10 {
            case -2147483648:
                delete(v3, (fi13 | fi13))
            case -9223372036854775808:
                fmt.Println(v2, v7)
            default:
                var v16 int = 127
                _ = v16
            }
            _ = fi13
        }
        var v17 int8 = -20
        _ = v17
    } else {
        v2 -= v2
    }
    v5 = (v5 / v5)
    switch ((("" + "a") + "a") + "foo bar") {
    case "hello":
        for fi18 := 0; fi18 < 9; fi18++ {
            var v19 uint = 18446744073709551615
            fmt.Println(v2, v7)
            _ = fi18
            _ = v19
        }
    case "a":
        for fi20 := 0; fi20 < 15; fi20++ {
            switch v4 {
            case 590:
                v2 = 39
            default:
                if v7 {
                    switch "a" {
                    case "hello":
                        var v21 int = v4
                        _ = v21
                    case "\n\t":
                        for fi22 := 0; fi22 < 14; fi22++ {
                            if len(v9) >= 2 {
                                v9 = v9[1:len(v9)]
                            }
                            fmt.Println(fi22, v2)
                            switch -131 {
                            case -9223372036854775808:
                                switch fi20 {
                                case -940:
                                    var v23 uint16 = 591
                                    _ = v23
                                default:
                                    v24 := GenS1{F0: 179, F1: 2}
                                    _ = v24
                                }
                            case 9223372036854775807:
                                v25 := GenS1{F0: -97, F1: 794}
                                _ = v25
                            default:
                                if true {
                                    v4 = (fi22 & fi22)
                                    var v26 uint8 = 165
                                    v27 := GenS3{F0: 2, F1: (v2 * v2)}
                                    v28 := GenS1{F0: -245, F1: 2}
                                    v7 = false
                                    delete(v3, fi20)
                                    _ = v26
                                    _ = v27
                                    _ = v28
                                } else {
                                    if (28 < 822) {
                                        v5 = -2.25
                                    }
                                    if true {
                                        fmt.Println(v2, v5)
                                        v7 = v7
                                        v2 += v2
                                    }
                                    v29 := GenS1{F0: 918, F1: 18446744073709551615}
                                    _ = v29
                                }
                            }
                            if v7 {
                                fmt.Println(v2)
                            }
                            _ = fi22
                        }
                    default:
                        v1 = (v1 | v1)
                    }
                    v7 = (v1 == -1)
                } else {
                    if true {
                        v30 := GenS4{F0: 4294967295, F1: 88}
                        switch (v4 << ((1) & 63)) {
                        case -1:
                            v1 -= (v1 + (v1 + v1))
                        }
                        _ = v30
                    }
                }
            }
            _ = fi20
        }
    case "\n\t":
        if (v4 <= (v4 - -236)) {
            var v31 uint16 = 724
            v4 -= v4
            _ = v31
        }
    default:
        fmt.Println(v2, v5, v1)
    }
    var v32 uint = 65535
    var v33 int64 = -347
    v34 := GenS3{F0: 2, F1: v2}
    v1 = v1
    v1 -= v1
    switch (-(v4)) {
    case 1:
        v35 := GenS2{F0: 1}
        _ = v35
    default:
        v36 := GenS2{F0: -12}
        _ = v36
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v12
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
