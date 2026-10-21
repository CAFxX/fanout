package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64 }
type GenS2 struct { F0 int; F1 int64; F2 uint }
type GenS3 struct { F0 bool }
type GenS4 struct { F0 bool }
type GenS5 struct { F0 int; F1 uint8 }
type GenS6 struct { F0 uint16; F1 string }
type GenS7 struct { F0 uint64 }

func main() {
    v1 := make(map[uint32]bool)
    v2 := make(map[int]bool)
    var v3 string = "hello"
    var v4 int64 = 1
    v5 := make(map[int]string)
    for fi6 := 0; fi6 < 14; fi6++ {
        _ = len(v5)
        delete(v5, (fi6 - fi6))
        v7 := GenS1{F0: 349}
        _ = fi6
        _ = v7
    }
    v5[-354] = v3
    v5[832] = ((v3 + ((v3 + v3) + "hello")) + v3)
    fmt.Println(v3)
    v8 := GenS1{F0: 671}
    for fi9 := 0; fi9 < 17; fi9++ {
        for fi10 := 0; fi10 < 17; fi10++ {
            v11 := GenS2{F0: fi9, F1: v4, F2: 1}
            _ = fi10
            _ = v11
        }
        if true {
            fmt.Println(v4)
            v12 := GenS3{F0: ((false || false) || false)}
            _ = v12
        }
        _ = fi9
    }
    var v13 float32 = -2.25
    v14 := GenS4{F0: true}
    var v15 float32 = 1e10
    v16 := GenS5{F0: 632, F1: 255}
    if false {
        switch -943 {
        case 42:
            v17 := []int{}
            _ = v17
        }
        fmt.Println(v15, v15, v4)
        switch "a" {
        case "\n\t":
            v18 := GenS6{F0: 65535, F1: v3}
            _ = v18
        case "hello":
            if false {
                v3 = v3
                fmt.Println(v4)
            } else {
                for fi19 := 0; fi19 < 11; fi19++ {
                    for fi20 := 0; fi20 < 7; fi20++ {
                        if true {
                            v21 := []int{fi20, (fi19 & 1)}
                            switch v3 {
                            case "":
                                var v22 float64 = 0.1
                                _ = v22
                            default:
                                v3 = "hello"
                                v15 += v13
                                v23 := GenS7{F0: 18446744073709551615}
                                switch (v3 + "a") {
                                case "a":
                                    v24 := v5[32767]
                                    _ = v24
                                case "hello":
                                    v13 = float32(3.14159)
                                case "":
                                    var v25 float32 = (v15 / (-2.25 + float32(((0.0 * ((0.1 * (float64((((-2.25 * float32(-2.25)) * v15) / v13)) / (((1.5 * float64(v13)) * (float64(-2.25) + float64((0.0 + v15)))) - (3.14159 + 1e10)))) * (float64(3.14159) * float64(((v15 * 3.14159) + (float32(1.5) + v13)))))) + ((3.14159 - 1.5) + (3.14159 + float64((v13 + v15))))))))
                                    _ = v25
                                default:
                                    var v26 uint16 = 291
                                    _ = v26
                                }
                                _ = v23
                            }
                            _ = v21
                        }
                        var v27 int8 = -1
                        var v28 int32 = -181
                        _ = fi20
                        _ = v27
                        _ = v28
                    }
                    v4 -= (v4 * v4)
                    _ = fi19
                }
                v3 = ("\n\t" + (v3 + (v3 + (("\n\t" + ("hello" + (v3 + v3))) + "\n\t"))))
            }
        case "":
            var v29 float32 = v13
            _ = v29
        }
    } else {
        fmt.Println(v15, v15, v13)
        v3 = v3
        fmt.Println(v3, v4)
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
