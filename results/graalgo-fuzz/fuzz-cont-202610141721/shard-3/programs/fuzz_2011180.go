package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint8; F1 uint8; F2 uint8 }
type GenS3 struct { F0 uint16; F1 int8; F2 uint64 }
type GenS4 struct { F0 int16; F1 bool; F2 int32 }
type GenS5 struct { F0 int; F1 uint32 }
type GenS6 struct { F0 int; F1 int8; F2 string }

func main() {
    v1 := []int{}
    var v2 int32 = 954
    v3 := []string{}
    var v4 int = -2147483648
    var v5 int8 = -128
    v6 := make(map[int]uint32)
    _ = len(v1)
    var v7 float64 = -2.25
    var v8 int = v4
    for fi9 := 0; fi9 < 2; fi9++ {
        if (false || false) {
            switch (-(fi9)) {
            case -622:
                v10 := GenS1{F0: 32767}
                _ = v10
            case -1:
                var v11 int16 = 1
                _ = v11
            default:
                var v12 int64 = 9223372036854775807
                _ = v12
            }
            v13 := GenS2{F0: 2, F1: 153, F2: 227}
            _ = v13
        }
        v7 = 0.0
        var v14 int = (v8 % ((v4) | 1))
        _ = fi9
        _ = v14
    }
    var v15 uint16 = 42
    switch v4 {
    case 32767:
        var v16 uint64 = 2
        _ = v16
    case 127:
        var v17 int32 = v2
        _ = v17
    case 32768:
        v18 := v6[v8]
        _ = v18
    }
    if true {
        fmt.Println(v2)
        if (v2 == -64) {
            for fi19 := 0; fi19 < 18; fi19++ {
                v6[(fi19 + (v8 / ((fi19) | 1)))] = 65535
                _ = fi19
            }
            for fi20 := 0; fi20 < 13; fi20++ {
                fmt.Println(v15, v2)
                _ = fi20
            }
            for fi21 := 0; fi21 < 8; fi21++ {
                v5 -= (^(v5))
                for fi22 := 0; fi22 < 19; fi22++ {
                    fmt.Println(v7, v7, fi21)
                    v23 := GenS3{F0: (v15 + (v15 % ((v15) | 1))), F1: v5, F2: 281}
                    _ = fi22
                    _ = v23
                }
                _ = fi21
            }
        } else {
            if true {
                if false {
                    switch "foo bar" {
                    case "foo bar":
                        if len(v1) > 0 {
                            v1[0] = int(v2)
                        }
                    case "":
                        v8 = v4
                    case "hello":
                        v4 += 88
                    }
                }
                var v24 float32 = 1e10
                _ = v24
            } else {
                switch v4 {
                case 9223372036854775807:
                    _ = len(v1)
                case -562:
                    v5 = v5
                }
                if (true && false) {
                    v25 := GenS4{F0: 928, F1: false, F2: v2}
                    _ = v25
                } else {
                    for fi26 := 0; fi26 < 6; fi26++ {
                        var v27 float64 = (float64(float32((v7 / v7))) - ((((v7 / -2.25) / ((((v7 - float64(0.1)) / v7) / 1.0) + v7)) + v7) / (1e10 / (1.5 / v7))))
                        v5 = v5
                        _ = fi26
                        _ = v27
                    }
                    var v28 uint8 = 127
                    _ = v28
                }
            }
            v29 := GenS4{F0: -216, F1: (false && (-32768 <= 1000)), F2: 258}
            _ = v29
        }
        v30 := GenS5{F0: 32768, F1: 2147483648}
        _ = v30
    }
    for fi31 := 0; fi31 < 16; fi31++ {
        v32 := GenS6{F0: (fi31 ^ v4), F1: -6, F2: "hello"}
        v5 = v5
        _ = fi31
        _ = v32
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
