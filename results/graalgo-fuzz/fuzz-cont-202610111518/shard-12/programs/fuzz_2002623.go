package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint; F1 int8 }
type GenS3 struct { F0 int32 }

func main() {
    var v1 int16 = -165
    var v2 uint8 = 128
    v3 := make(map[string]int)
    v4 := []int{-3, -2147483648, -1}
    v2 -= v2
    fmt.Println(v1, v1)
    for fi5 := 0; fi5 < 19; fi5++ {
        var v6 int32 = -97
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi5
        _ = v6
    }
    v1 = v1
    var v7 int64 = -9223372036854775808
    if false {
        for fi8 := 0; fi8 < 6; fi8++ {
            for fi9 := 0; fi9 < 11; fi9++ {
                v1 -= 42
                switch "hello" {
                case "\n\t":
                    var v10 uint32 = 2147483647
                    _ = v10
                }
                _ = fi9
            }
            for fi11 := 0; fi11 < 3; fi11++ {
                if false {
                    var v12 int16 = (-(v1))
                    if false {
                        v1 -= (v12 | (v12 >> ((65535) & 63)))
                        var v13 int32 = -2147483648
                        var v14 uint16 = 573
                        _ = v13
                        _ = v14
                    }
                    v15 := GenS1{F0: 374}
                    _ = v12
                    _ = v15
                } else {
                    fmt.Println(v7, fi11)
                    var v16 float64 = 0.1
                    _ = v16
                }
                if false {
                    v2 -= v2
                    var v17 uint = 42
                    _ = v17
                }
                switch "\n\t" {
                case "a":
                    if false {
                        var v18 bool = false
                        _ = v18
                    } else {
                        if (0.1 == (0.0 + 1e10)) {
                            fmt.Println(v2, v1)
                            if ((v2 % ((42) | 1)) <= v2) {
                                v1 -= v1
                                fmt.Println(v1, v1, v1)
                                v19 := GenS1{F0: 481}
                                _ = v19
                            } else {
                                v2 += 255
                                switch "" {
                                case "a":
                                    var v20 uint64 = 0
                                    v20 = v20
                                    _ = v20
                                case "":
                                    var v21 float32 = float32(3.14159)
                                    _ = v21
                                default:
                                    v22 := GenS2{F0: 42, F1: 0}
                                    _ = v22
                                }
                            }
                            var v23 uint32 = 551
                            _ = v23
                        } else {
                            var v24 bool = true
                            _ = v24
                        }
                        if (42 > 2) {
                            v25 := GenS1{F0: 728}
                            v1 += (v1 << ((2) & 63))
                            v2 = (v2 | v2)
                            _ = v25
                        } else {
                            v26 := v3["\n\t"]
                            fmt.Println(v26, v1)
                            _ = v26
                        }
                    }
                    fmt.Println(v7, v2, fi8)
                    if (1 == 4294967295) {
                        var v27 float32 = 0.1
                        _ = v27
                    } else {
                        v1 += v1
                        fmt.Println(v2, v1)
                    }
                case "\n\t":
                    var v28 uint8 = (^(v2))
                    _ = v28
                case "":
                    v2 = 255
                default:
                    var v29 uint32 = 42
                    _ = v29
                }
                _ = fi11
            }
            _ = fi8
        }
        v30 := GenS3{F0: -2147483648}
        v31 := GenS3{F0: 2147483647}
        _ = v30
        _ = v31
    }
    var v32 string = "foo bar"
    var v33 int32 = -234
    var v34 string = (v32 + (((v32 + "\n\t") + v32) + "\n\t"))
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v32:", v32)
    fmt.Println("v33:", v33)
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v2
    _ = v3
    _ = v32
    _ = v33
    _ = v34
    _ = v4
    _ = v7
}
