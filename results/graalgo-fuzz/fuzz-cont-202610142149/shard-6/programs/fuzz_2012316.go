package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 uint8 }
type GenS3 struct { F0 int64; F1 int64; F2 int64 }
type GenS4 struct { F0 int32 }

func main() {
    var v1 uint16 = 32767
    v2 := []int{-571, 32767, 32767, 2147483647}
    var v3 int32 = -2147483648
    v1 -= 405
    v4 := GenS1{F0: (v3 >> ((415) & 63))}
    for fi5 := 0; fi5 < 19; fi5++ {
        switch (fi5 >> ((1) & 63)) {
        case 2147483647:
            switch fi5 {
            case 18:
                var v6 int64 = 1
                _ = v6
            case -1:
                fmt.Println(v3)
            case 42:
                if true {
                    fmt.Println(v3)
                    var v7 uint = 42
                    var v8 int64 = 462
                    _ = v7
                    _ = v8
                } else {
                    for fi9 := 0; fi9 < 3; fi9++ {
                        var v10 uint32 = 2
                        switch "a" {
                        case "":
                            var v11 int64 = 9223372036854775807
                            v3 -= v3
                            v12 := GenS2{F0: 185}
                            var v13 float64 = (float64(1e10) * 0.1)
                            if false {
                                fmt.Println(fi9)
                            } else {
                                v14 := GenS1{F0: -2147483648}
                                var v15 int64 = v11
                                _ = v14
                                _ = v15
                            }
                            v3 = v3
                            var v16 int64 = v11
                            var v17 string = "foo bar"
                            _ = v11
                            _ = v12
                            _ = v13
                            _ = v16
                            _ = v17
                        case "a":
                            v18 := GenS3{F0: 9223372036854775807, F1: -9223372036854775808, F2: 819}
                            _ = v18
                        case "hello":
                            var v19 uint8 = 166
                            _ = v19
                        }
                        v3 -= v3
                        v1 = (v1 - (v1 + (v1 - v1)))
                        v1 = 2
                        _ = fi9
                        _ = v10
                    }
                }
            default:
                v2 = append(v2, fi5)
            }
        case 359:
            switch -2147483648 {
            case -833:
                if (!((42 <= 2) && true)) {
                    v20 := GenS1{F0: v3}
                    v1 += (v1 % (((v1 << ((uint(v1)) & 63))) | 1))
                    var v21 bool = (107 < 2147483647)
                    _ = v20
                    _ = v21
                }
            case 127:
                for fi22 := 0; fi22 < 6; fi22++ {
                    v2 = append(v2, (fi22 / ((9223372036854775807) | 1)))
                    var v23 int16 = 1
                    if len(v2) > 0 {
                        v2[0] = fi22
                    }
                    _ = fi22
                    _ = v23
                }
            case -1:
                v3 -= v3
            }
        default:
            v24 := GenS1{F0: v3}
            _ = v24
        }
        if (("a" >= "\n\t") || (255 != (v1 / ((620) | 1)))) {
            v3 = v3
            v3 += (-(v3))
        } else {
            _ = len(v2)
            var v25 int64 = 226
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            _ = v25
        }
        for fi26 := 0; fi26 < 9; fi26++ {
            v1 -= (v1 + v1)
            var v27 string = "a"
            _ = fi26
            _ = v27
        }
        _ = fi5
    }
    v3 = v3
    if true {
        fmt.Println(v1, v3)
    }
    var v28 bool = (127 <= 42)
    if true {
        for fi29 := 0; fi29 < 4; fi29++ {
            if len(v2) >= 2 {
                v2 = v2[1:len(v2)]
            }
            var v30 float32 = 3.14159
            _ = fi29
            _ = v30
        }
    }
    v31 := GenS1{F0: 1}
    if (128 >= 1) {
        fmt.Println(v1, v28, v28)
        var v32 int8 = -128
        v33 := GenS4{F0: (v3 & (v3 % ((v3) | 1)))}
        _ = v32
        _ = v33
    } else {
        fmt.Println(v1, v3, v3)
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v28:", v28)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v2
    _ = v28
    _ = v3
    _ = v31
    _ = v4
}
