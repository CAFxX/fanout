package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 string; F1 bool }
type GenS3 struct { F0 uint8; F1 uint8; F2 int }

func main() {
    var v1 int64 = -9223372036854775808
    var v2 uint8 = 128
    var v3 bool = true
    v4 := make(map[int]uint32)
    var v5 uint8 = 187
    v6 := GenS1{F0: -128}
    v3 = ((!true) && v3)
    if (127 > v5) {
        v3 = v3
        fmt.Println(v5, v2)
    } else {
        v2 += 1
    }
    var v7 float64 = 0.1
    if true {
        v1 -= -532
    }
    for fi8 := 0; fi8 < 5; fi8++ {
        var v9 uint64 = 326
        v10 := []int{}
        var v11 uint = 570
        _ = fi8
        _ = v10
        _ = v11
        _ = v9
    }
    v3 = (156 <= 2)
    switch -199 {
    case 32768:
        v12 := GenS2{F0: "\n\t", F1: true}
        _ = v12
    }
    var v13 string = "foo bar"
    var v14 int16 = -1
    for fi15 := 0; fi15 < 11; fi15++ {
        for fi16 := 0; fi16 < 19; fi16++ {
            var v17 float64 = (v7 / (v7 * ((v7 / (float64((3.14159 / float32(v7))) - v7)) - v7)))
            if true {
                if (!false) {
                    var v18 uint = 722
                    v7 += float64(0.0)
                    _ = v18
                } else {
                    var v19 bool = (!(true || false))
                    switch -9223372036854775808 {
                    case 127:
                        if (1 >= 2) {
                            delete(v4, fi16)
                        } else {
                            switch fi15 {
                            case 21:
                                var v20 int64 = v1
                                _ = v20
                            case 127:
                                var v21 uint32 = 325
                                _ = v21
                            default:
                                fmt.Println(fi15, v5)
                            }
                            switch "hello" {
                            case "hello":
                                var v22 float64 = 1.5
                                _ = v22
                            case "\n\t":
                                var v23 int32 = 2147483647
                                _ = v23
                            case "foo bar":
                                var v24 uint8 = (v5 & (v2 / ((v2) | 1)))
                                _ = v24
                            }
                            v1 += (v1 & 422)
                        }
                    case 32768:
                        v25 := GenS1{F0: 11}
                        _ = v25
                    default:
                        v19 = (!(-762 == 0))
                    }
                    v3 = v19
                    _ = v19
                }
            } else {
                v26 := GenS1{F0: 1}
                if (-32768 < (v14 % (((v14 * (v14 ^ 32767))) | 1))) {
                    switch ("\n\t" + v13) {
                    case "foo bar":
                        switch v13 {
                        case "a":
                            v3 = v3
                        case "foo bar":
                            v27 := []int{42, 9223372036854775807}
                            _ = v27
                        default:
                            v28 := GenS2{F0: v13, F1: v3}
                            _ = v28
                        }
                    case "\n\t":
                        fmt.Println(v7, v17, v14)
                    case "hello":
                        fmt.Println(v2, v7)
                    default:
                        v2 = v5
                    }
                    v29 := []int{fi16}
                    if v3 {
                        v5 = (v5 / ((v2) | 1))
                    } else {
                        var v30 int8 = -116
                        v31 := GenS1{F0: 74}
                        var v32 bool = (4294967295 < 4294967295)
                        _ = v30
                        _ = v31
                        _ = v32
                    }
                    _ = v29
                }
                _ = v26
            }
            v33 := v4[fi16]
            _ = fi16
            _ = v17
            _ = v33
        }
        v34 := GenS2{F0: "a", F1: v3}
        _ = fi15
        _ = v34
    }
    v35 := []int{-650, 32768, 69}
    var v36 uint16 = 312
    for fi37 := 0; fi37 < 17; fi37++ {
        if v3 {
            fmt.Println(v3, v36, fi37)
            if ((v36 ^ v36) >= (v36 + v36)) {
                _ = len(v35)
            } else {
                if len(v35) >= 2 {
                    v35 = v35[1:len(v35)]
                }
                v7 += (-2.25 * (float64(3.14159) / float64(-2.25)))
                if (42 <= 65535) {
                    var v38 int32 = -2147483648
                    v39 := GenS3{F0: v5, F1: v5, F2: fi37}
                    _ = v38
                    _ = v39
                } else {
                    _ = len(v35)
                }
            }
        }
        v14 = -766
        _ = fi37
    }
    switch v13 {
    case "foo bar":
        for fi40 := 0; fi40 < 6; fi40++ {
            v36 = v36
            _ = len(v4)
            _ = fi40
        }
    case "a":
        if len(v35) > 0 {
            v35[0] = 866
        }
    default:
        var v41 float32 = 1.5
        _ = v41
    }
    var v42 float64 = 0.1
    var v43 int64 = (v1 - (v1 % (((v1 % ((v1) | 1))) | 1)))
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v36:", v36)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v42:", v42)
    fmt.Println("v43:", v43)
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v35
    _ = v36
    _ = v4
    _ = v42
    _ = v43
    _ = v5
    _ = v6
    _ = v7
}
