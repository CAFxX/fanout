package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint; F1 int64; F2 int64 }
type GenS4 struct { F0 int32; F1 bool }
type GenS5 struct { F0 int8; F1 int32; F2 int }

func main() {
    v1 := make(map[uint32]bool)
    var v2 bool = (!true)
    v3 := []uint32{642}
    v4 := make(map[int]int)
    if len(v3) >= 2 {
        v3 = v3[1:len(v3)]
    }
    fmt.Println(v2)
    v5 := v4[128]
    v2 = v2
    var v6 float32 = (3.14159 / (1e10 + float32(float64(1e10))))
    if len(v3) > 0 {
        v3[0] = 590
    }
    if (!true) {
        if v2 {
            fmt.Println(v6)
            if false {
                v7 := GenS1{F0: 42}
                _ = v7
            }
        }
    } else {
        var v8 int8 = 42
        v8 += (v8 / (((v8 / ((v8) | 1))) | 1))
        if (((!(65535 < 2147483647)) && false) || false) {
            v5 = v5
            if (459 == 624) {
                v9 := GenS2{F0: true}
                v8 = v8
                _ = v9
            } else {
                _ = len(v1)
                if (128 >= 0) {
                    v10 := GenS3{F0: 695, F1: 9223372036854775807, F2: -9223372036854775808}
                    _ = v10
                }
                var v11 uint8 = 42
                _ = v11
            }
            if true {
                if len(v3) > 0 {
                    v3[0] = 1
                }
            } else {
                v12 := GenS4{F0: -699, F1: (65535 < 1)}
                switch "hello" {
                case "a":
                    v5 = (v5 - -993)
                case "\n\t":
                    for fi13 := 0; fi13 < 13; fi13++ {
                        if false {
                            for fi14 := 0; fi14 < 19; fi14++ {
                                v15 := GenS5{F0: 42, F1: 42, F2: fi13}
                                v2 = true
                                if len(v3) >= 2 {
                                    v3 = v3[1:len(v3)]
                                }
                                _ = fi14
                                _ = v15
                            }
                            v16 := GenS5{F0: v8, F1: -1, F2: -899}
                            fmt.Println(v2, v8, fi13)
                            _ = v16
                        } else {
                            v17 := GenS2{F0: (v5 < fi13)}
                            _ = v17
                        }
                        var v18 uint16 = 0
                        v6 = v6
                        _ = fi13
                        _ = v18
                    }
                default:
                    v5 = v5
                }
                fmt.Println(v5, v5)
                _ = v12
            }
        }
        _ = v8
    }
    if (true || true) {
        var v19 float32 = 1.5
        switch v5 {
        case 1:
            for fi20 := 0; fi20 < 14; fi20++ {
                if len(v3) > 0 {
                    v3[0] = 128
                }
                _ = len(v1)
                _ = fi20
            }
        default:
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        }
        v19 = -2.25
        _ = v19
    }
    v1[65535] = v2
    _ = len(v3)
    var v21 int16 = -546
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
