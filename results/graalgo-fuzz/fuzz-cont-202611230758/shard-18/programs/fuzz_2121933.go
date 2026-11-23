package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 uint64 }
type GenS2 struct { F0 int }
type GenS3 struct { F0 int16; F1 int8 }
type GenS4 struct { F0 string; F1 int64 }
type GenS5 struct { F0 uint64; F1 uint }

func main() {
    v1 := make(map[int]string)
    var v2 uint16 = 255
    var v3 uint8 = 128
    v3 -= v3
    v4 := []int{42, 9, -9223372036854775808}
    var v5 uint = 2
    v5 += uint(v2)
    if false {
        v2 += (^(v2))
        v5 = (^(v5))
        v6 := GenS1{F0: 42, F1: -1, F2: uint64(v5)}
        _ = v6
    } else {
        if len(v4) > 0 {
            v4[0] = 1
        }
        v7 := GenS1{F0: 32767, F1: 9223372036854775807, F2: 743}
        _ = v7
    }
    var v8 uint8 = 128
    v2 += (v2 / ((v2) | 1))
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    if false {
        _ = len(v1)
        for fi9 := 0; fi9 < 5; fi9++ {
            fmt.Println(v5, v2, fi9)
            _ = fi9
        }
        v1[-349] = "hello"
    } else {
        var v10 string = "foo bar"
        v10 = v10
        v8 += 114
        _ = v10
    }
    v11 := v1[128]
    v12 := GenS2{F0: 793}
    switch -1 {
    case 918:
        v5 += v5
    case 59:
        v5 = (^(v5))
    case 13:
        if (true || false) {
            _ = len(v1)
            for fi13 := 0; fi13 < 4; fi13++ {
                switch "hello" {
                case "\n\t":
                    for fi14 := 0; fi14 < 12; fi14++ {
                        _ = len(v4)
                        var v15 int8 = 100
                        _ = fi14
                        _ = v15
                    }
                default:
                    _ = len(v4)
                }
                if true {
                    v2 -= v2
                }
                var v16 bool = true
                _ = fi13
                _ = v16
            }
        }
    default:
        _ = len(v1)
    }
    for fi17 := 0; fi17 < 2; fi17++ {
        v18 := GenS3{F0: 1, F1: 100}
        v19 := GenS4{F0: v11, F1: -1}
        _ = fi17
        _ = v18
        _ = v19
    }
    v20 := GenS5{F0: 4294967295, F1: (v5 | 1)}
    if (!(65535 == 65535)) {
        v8 = (v3 >> (((v5 >> ((v5) & 63))) & 63))
        v11 = ""
        for fi21 := 0; fi21 < 7; fi21++ {
            delete(v1, 498)
            if true {
                switch v11 {
                case "\n\t":
                    if (false || true) {
                        v5 += (v5 ^ v5)
                        if (1e10 == 0.1) {
                            v8 = (v8 >> ((v5) & 63))
                        } else {
                            for fi22 := 0; fi22 < 5; fi22++ {
                                if false {
                                    var v23 string = ((v11 + "a") + "hello")
                                    var v24 float32 = 3.14159
                                    v8 -= (v3 + v8)
                                    _ = v23
                                    _ = v24
                                }
                                _ = fi22
                            }
                        }
                    } else {
                        var v25 uint64 = 1
                        var v26 uint8 = (v8 - (v3 << (((v5 ^ v5)) & 63)))
                        v8 = uint8(v2)
                        _ = v25
                        _ = v26
                    }
                }
                v27 := GenS4{F0: "hello", F1: 88}
                _ = v27
            } else {
                var v28 int64 = 721
                switch (v11 + "") {
                case "\n\t":
                    switch -707 {
                    case 892:
                        var v29 int8 = -23
                        _ = v29
                    default:
                        if ((float64(-2.25) / 3.14159) == 1e10) {
                            delete(v1, fi21)
                        } else {
                            var v30 uint32 = 37
                            v31 := GenS2{F0: 128}
                            v32 := GenS1{F0: -289, F1: fi21, F2: uint64(v2)}
                            _ = v30
                            _ = v31
                            _ = v32
                        }
                    }
                case "foo bar":
                    if true {
                        fmt.Println(v28, v2)
                    }
                default:
                    if len(v4) >= 2 {
                        v4 = v4[1:len(v4)]
                    }
                }
                _ = v28
            }
            v11 = v11
            _ = fi21
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v12
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
