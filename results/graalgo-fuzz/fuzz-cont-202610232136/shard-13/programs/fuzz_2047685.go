package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int16; F1 bool; F2 uint16 }
type GenS4 struct { F0 uint32; F1 uint64 }
type GenS5 struct { F0 int32; F1 uint64; F2 int64 }

func main() {
    v1 := make(map[int]int)
    var v2 uint64 = 18446744073709551615
    var v3 int = -2147483648
    for fi4 := 0; fi4 < 1; fi4++ {
        var v5 float64 = ((-2.25 / 0.1) + float64(0.0))
        _ = fi4
        _ = v5
    }
    var v6 float32 = (0.0 / -2.25)
    var v7 uint = 109
    if true {
        v2 = v2
        for fi8 := 0; fi8 < 12; fi8++ {
            fmt.Println(fi8, v2, fi8)
            _ = fi8
        }
        v9 := GenS1{F0: 2}
        _ = v9
    }
    v10 := v1[v3]
    v7 = v7
    v3 -= (^(v10))
    for fi11 := 0; fi11 < 18; fi11++ {
        if true {
            for fi12 := 0; fi12 < 6; fi12++ {
                if false {
                    fmt.Println(fi12)
                    var v13 int = 1
                    _ = v13
                }
                switch "a" {
                case "":
                    v14 := GenS2{F0: -32768}
                    _ = v14
                case "foo bar":
                    switch "foo bar" {
                    case "foo bar":
                        switch "a" {
                        case "foo bar":
                            v15 := GenS3{F0: 0, F1: true, F2: uint16(v7)}
                            _ = v15
                        case "hello":
                            switch "a" {
                            case "a":
                                var v16 uint = (v7 * 42)
                                v17 := GenS4{F0: 65535, F1: (v2 % ((v2) | 1))}
                                if true {
                                    fmt.Println(v16)
                                    switch "" {
                                    case "foo bar":
                                        v18 := v1[0]
                                        _ = v18
                                    case "\n\t":
                                        v2 = 2
                                    default:
                                        var v19 bool = (false || false)
                                        _ = v19
                                    }
                                    var v20 int8 = 0
                                    _ = v20
                                } else {
                                    v21 := GenS5{F0: 585, F1: v2, F2: 78}
                                    _ = v21
                                }
                                _ = v16
                                _ = v17
                            }
                        default:
                            v2 = (v2 << (((v7 + (v7 * (v7 * v7)))) & 63))
                            fmt.Println(v10, v2, fi11)
                        }
                    case "\n\t":
                        v2 += (v2 + v2)
                    default:
                        var v22 uint64 = v2
                        if true {
                            v3 = v3
                        } else {
                            v6 -= (3.14159 / (3.14159 * v6))
                            var v23 float32 = v6
                            _ = v23
                        }
                        _ = v22
                    }
                }
                v1[-585] = (-(fi11))
                _ = fi12
            }
            if false {
                var v24 uint64 = v2
                var v25 uint16 = 2
                _ = v24
                _ = v25
            } else {
                var v26 uint = (v7 & (^(v7)))
                v7 += v7
                v1[(^(v3))] = (v3 % ((fi11) | 1))
                _ = v26
            }
        }
        fmt.Println(v6, v3, v3)
        _ = fi11
    }
    v6 = 3.14159
    v2 = v2
    var v27 int8 = -3
    for fi28 := 0; fi28 < 19; fi28++ {
        fmt.Println(v2, fi28)
        fmt.Println(v2)
        _ = fi28
    }
    switch 32768 {
    case 4294967295:
        fmt.Println(v7, v10, v7)
    case -9223372036854775808:
        if true {
            v27 += (v27 % ((1) | 1))
            v3 = 1
        }
    case -255:
        v1[32767] = v10
    default:
        v29 := []int{v3, (^(v3)), (v10 % ((457) | 1)), v3}
        _ = v29
    }
    fmt.Println(v7)
    if (false && (!false)) {
        _ = len(v1)
        var v30 uint32 = 0
        var v31 int = (v10 - (v10 & (v10 / (((^(v3))) | 1))))
        _ = v30
        _ = v31
    }
    var v32 int32 = -2147483648
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v2:", v2)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v32:", v32)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v10
    _ = v2
    _ = v27
    _ = v3
    _ = v32
    _ = v6
    _ = v7
}
