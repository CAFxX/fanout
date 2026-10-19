package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint64; F1 int }
type GenS3 struct { F0 uint16; F1 uint32 }
type GenS4 struct { F0 int; F1 uint64 }
func genFn0() uint32 {
    v5 := GenS1{F0: 2}
    for fi6 := 0; fi6 < 20; fi6++ {
        v7 := []int{fi6, fi6}
        switch ("" + "hello") {
        case "foo bar":
            for fi8 := 0; fi8 < 10; fi8++ {
                switch fi8 {
                case 42:
                    var v9 uint32 = 868
                    _ = v9
                }
                switch "" {
                case "foo bar":
                    v10 := GenS2{F0: 18446744073709551615, F1: -1}
                    _ = v10
                default:
                    if true {
                        _ = len(v7)
                        if false {
                            fmt.Println(fi8, fi8)
                            switch 128 {
                            case -2147483648:
                                var v11 uint16 = 255
                                _ = v11
                            case -828:
                                switch ("foo bar" + "a") {
                                case "a":
                                    var v12 string = "foo bar"
                                    _ = v12
                                case "":
                                    switch 2147483647 {
                                    case 4294967295:
                                        var v13 float32 = -2.25
                                        _ = v13
                                    case -2147483648:
                                        if (4294967295 != 759) {
                                            v14 := GenS3{F0: 2, F1: 4294967295}
                                            _ = v14
                                        }
                                    default:
                                        switch fi8 {
                                        case 397:
                                            v15 := GenS4{F0: (fi6 | (^(fi8))), F1: 2}
                                            _ = v15
                                        default:
                                            var v16 int32 = 1
                                            _ = v16
                                        }
                                    }
                                default:
                                    if false {
                                        var v17 string = "a"
                                        _ = v17
                                    }
                                }
                            }
                        } else {
                            var v18 uint32 = 0
                            v18 = 497
                            var v19 int8 = -128
                            v18 += v18
                            v7 = append(v7, 134)
                            _ = v18
                            _ = v19
                        }
                    }
                }
                fmt.Println(fi6, fi8)
                _ = fi8
            }
        case "":
            fmt.Println(fi6)
        case "\n\t":
            v20 := GenS3{F0: 745, F1: 2}
            _ = v20
        default:
            var v21 int = (fi6 + fi6)
            _ = v21
        }
        _ = fi6
        _ = v7
    }
    var v22 int16 = 32767
    var v23 int16 = (v22 - v22)
    _ = v22
    _ = v23
    _ = v5
    return 65535
}

func main() {
    var v1 float32 = float32(1.5)
    var v2 int64 = 0
    v3 := make(map[uint32]int)
    var v4 int = -582
    if false {
        fmt.Println(v2, v1, v2)
    }
    v24 := []int{}
    v1 += v1
    v1 = (((v1 + -2.25) / v1) + float32(-2.25))
    switch (v4 | (v4 * (v4 >> ((65535) & 63)))) {
    case 42:
        v2 = 1
    case 194:
        v4 += 9223372036854775807
    case 127:
        var v25 float64 = 1.5
        _ = v25
    default:
        v26 := GenS4{F0: v4, F1: 18446744073709551615}
        _ = v26
    }
    v24 = append(v24, v4)
    var v27 int32 = -257
    v28 := GenS3{F0: 803, F1: 65535}
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v24:", len(v24), v24)
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v2
    _ = v24
    _ = v27
    _ = v28
    _ = v3
    _ = v4
}
