package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int64; F1 bool }
type GenS3 struct { F0 uint; F1 int8 }
type GenS4 struct { F0 int8; F1 int16; F2 uint64 }
type GenS5 struct { F0 int64 }
type GenS6 struct { F0 uint16; F1 uint8 }
type GenS7 struct { F0 uint64; F1 bool; F2 int8 }

func main() {
    var v1 uint = 792
    v2 := make(map[string]string)
    v3 := make(map[uint32]string)
    var v4 uint = 18446744073709551615
    v5 := make(map[int]bool)
    switch "foo bar" {
    case "\n\t":
        v6 := GenS1{F0: (1 >= 32767)}
        _ = v6
    case "hello":
        for fi7 := 0; fi7 < 9; fi7++ {
            var v8 int16 = 21
            v4 -= v4
            _ = fi7
            _ = v8
        }
    case "foo bar":
        fmt.Println(v4, v1)
    default:
        if (!false) {
            v1 += (^(v1))
            v9 := GenS1{F0: (!(-245 < -301))}
            if (2 < 4294967295) {
                v3[2147483648] = "\n\t"
                if false {
                    for fi10 := 0; fi10 < 10; fi10++ {
                        var v11 uint8 = 0
                        if (((1 <= 2147483647) || true) && true) {
                            delete(v2, "")
                        } else {
                            v4 = (v1 >> ((32) & 63))
                        }
                        _ = fi10
                        _ = v11
                    }
                    var v12 uint = (v1 + (v1 << ((v1) & 63)))
                    v13 := GenS1{F0: false}
                    _ = v12
                    _ = v13
                } else {
                    v14 := GenS2{F0: 577, F1: (65535 < 109)}
                    var v15 uint8 = 0
                    var v16 float64 = (((0.1 - 3.14159) * (float64((0.0 + (3.14159 / 1.0))) * (3.14159 * (float64(float32(-2.25)) - float64(float32(0.1)))))) + 0.0)
                    _ = v14
                    _ = v15
                    _ = v16
                }
            } else {
                fmt.Println(v4)
                var v17 uint64 = 512
                v18 := GenS1{F0: (!(0 == 4294967295))}
                _ = v17
                _ = v18
            }
            _ = v9
        }
    }
    if true {
        var v19 int8 = -128
        var v20 float32 = 0.1
        for fi21 := 0; fi21 < 16; fi21++ {
            v22 := GenS3{F0: v4, F1: v19}
            if false {
                v23 := GenS2{F0: 9223372036854775807, F1: false}
                _ = v23
            }
            if (-217 < 920) {
                v5[fi21] = false
                v24 := []int{fi21, 127, (fi21 >> (((v1 / (((^(v4))) | 1))) & 63)), fi21}
                _ = len(v2)
                _ = v24
            } else {
                if (2147483647 <= 1) {
                    var v25 uint32 = 1
                    if (-696 > 2147483647) {
                        v26 := GenS4{F0: 116, F1: 42, F2: 18446744073709551615}
                        var v27 uint64 = 2
                        _ = v26
                        _ = v27
                    } else {
                        if (fi21 == fi21) {
                            v28 := GenS5{F0: 35}
                            _ = v28
                        }
                    }
                    _ = v25
                }
            }
            _ = fi21
            _ = v22
        }
        _ = v19
        _ = v20
    }
    for fi29 := 0; fi29 < 2; fi29++ {
        v30 := v2[""]
        _ = fi29
        _ = v30
    }
    v31 := GenS6{F0: 2, F1: uint8(v1)}
    if false {
        fmt.Println(v4, v4)
    }
    v1 = (v1 + (v4 >> (((^(v1))) & 63)))
    v32 := GenS7{F0: 42, F1: false, F2: 25}
    v5[127] = false
    v33 := GenS3{F0: 0, F1: 0}
    v4 -= 134
    var v34 bool = (-2.25 <= (1e10 * 1e10))
    v35 := []int{-2147483648, 1}
    var v36 float64 = float64((0.1 + 3.14159))
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v34:", v34)
    fmt.Println("slicelen v35:", len(v35), v35)
    fmt.Println("v36:", v36)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v2
    _ = v3
    _ = v31
    _ = v32
    _ = v33
    _ = v34
    _ = v35
    _ = v36
    _ = v4
    _ = v5
}
