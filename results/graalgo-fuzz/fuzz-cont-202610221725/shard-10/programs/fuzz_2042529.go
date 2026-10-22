package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 string; F1 int8 }
type GenS4 struct { F0 int16; F1 int32; F2 uint16 }
type GenS5 struct { F0 int16; F1 string }
type GenS6 struct { F0 uint16; F1 int16 }

func main() {
    var v1 float64 = 0.0
    v2 := []int{636, 1, 619}
    var v3 int = 1
    v4 := make(map[int]bool)
    var v5 uint = 4294967295
    if len(v2) > 0 {
        v2[0] = -641
    }
    if true {
        v1 = v1
        switch 128 {
        case 128:
            if false {
                fmt.Println(v5, v3, v3)
                v3 -= v3
            }
        case 967:
            v3 += -1
        }
        for fi6 := 0; fi6 < 12; fi6++ {
            fmt.Println(fi6, v1, v5)
            _ = fi6
        }
    }
    var v7 int8 = 0
    v8 := GenS1{F0: 538, F1: 1}
    if len(v2) > 0 {
        v2[0] = -672
    }
    var v9 uint = v5
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    fmt.Println(v1, v9, v9)
    for fi10 := 0; fi10 < 3; fi10++ {
        v5 -= (v9 | (v9 + (^(v5))))
        if false {
            if true {
                v5 = 0
                for fi11 := 0; fi11 < 8; fi11++ {
                    if false {
                        switch "a" {
                        case "a":
                            v12 := v4[fi10]
                            _ = v12
                        case "foo bar":
                            if (127 < (fi11 >> (((v5 ^ (v9 << (((v5 / (((v5 / (((v9 ^ 0)) | 1))) | 1))) & 63)))) & 63))) {
                                var v13 int32 = 1
                                switch "a" {
                                case "foo bar":
                                    var v14 float64 = (((((float64(0.1) / (v1 / v1)) * 1e10) - v1) - ((-2.25 + 3.14159) / -2.25)) * v1)
                                    _ = v14
                                case "\n\t":
                                    switch (("a" + "hello") + "") {
                                    case "hello":
                                        v9 = v5
                                    case "\n\t":
                                        var v15 bool = ((!((false || false) || (-1 >= 0))) && false)
                                        _ = v15
                                    case "a":
                                        if true {
                                            var v16 int = -544
                                            if true {
                                                var v17 uint8 = 42
                                                switch (("foo bar" + "hello") + ("a" + "\n\t")) {
                                                case "hello":
                                                    v16 -= 0
                                                default:
                                                    var v18 int16 = -1
                                                    _ = v18
                                                }
                                                _ = v17
                                            } else {
                                                delete(v4, (v3 + (fi11 ^ (v3 % ((9223372036854775807) | 1)))))
                                                v1 = (float64(1e10) + float64(float32(v1)))
                                                v19 := GenS2{F0: -1}
                                                _ = v19
                                            }
                                            var v20 uint8 = 128
                                            _ = v16
                                            _ = v20
                                        } else {
                                            v21 := GenS3{F0: "foo bar", F1: v7}
                                            _ = v21
                                        }
                                    }
                                case "":
                                    if (4294967295 > 140) {
                                        delete(v4, (fi10 << (((v9 % ((v5) | 1))) & 63)))
                                        fmt.Println(v1, fi10)
                                    }
                                }
                                _ = v13
                            } else {
                                v4[(fi11 - 716)] = false
                                v7 = v7
                                if ((0 <= 42) || (!false)) {
                                    if len(v2) >= 2 {
                                        v2 = v2[1:len(v2)]
                                    }
                                    if len(v2) > 0 {
                                        v2[0] = (fi10 % ((fi10) | 1))
                                    }
                                } else {
                                    var v22 uint64 = 611
                                    var v23 int16 = -32768
                                    v24 := GenS4{F0: -439, F1: 0, F2: 255}
                                    _ = v22
                                    _ = v23
                                    _ = v24
                                }
                                var v25 uint = v5
                                v26 := GenS5{F0: 1000, F1: "\n\t"}
                                _ = v25
                                _ = v26
                            }
                        case "hello":
                            var v27 int = fi11
                            _ = v27
                        }
                        v9 -= v9
                    } else {
                        fmt.Println(v9, v9)
                    }
                    if len(v2) > 0 {
                        v2[0] = (fi11 % ((v3) | 1))
                    }
                    _ = fi11
                }
                if false {
                    v28 := GenS4{F0: 1, F1: -1, F2: 0}
                    for fi29 := 0; fi29 < 12; fi29++ {
                        fmt.Println(v9)
                        _ = fi29
                    }
                    _ = v28
                }
            }
            v1 = ((v1 * v1) / v1)
        }
        if (!true) {
            var v30 int64 = 331
            _ = v30
        }
        _ = fi10
    }
    fmt.Println(v7, v7, v9)
    v31 := GenS6{F0: 2, F1: 1}
    var v32 float32 = -2.25
    v5 = 1
    v33 := GenS3{F0: ("foo bar" + "foo bar"), F1: 42}
    v9 += v9
    fmt.Println(v9, v32)
    for fi34 := 0; fi34 < 4; fi34++ {
        v32 -= v32
        var v35 float32 = v32
        _ = fi34
        _ = v35
    }
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v4
    _ = v5
    _ = v7
    _ = v8
    _ = v9
}
