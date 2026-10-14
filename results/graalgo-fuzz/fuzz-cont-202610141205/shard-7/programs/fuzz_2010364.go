package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int16; F2 int32 }
type GenS2 struct { F0 bool }

func main() {
    var v1 int16 = 0
    v2 := make(map[int]bool)
    var v3 uint8 = 99
    var v4 uint16 = 42
    var v5 uint64 = 42
    v6 := make(map[string]uint32)
    v3 = 128
    v7 := []int{127, 127, 612, 9223372036854775807}
    v8 := GenS1{F0: 42, F1: v1, F2: -2147483648}
    var v9 uint16 = (v4 >> ((4294967295) & 63))
    for fi10 := 0; fi10 < 9; fi10++ {
        if true {
            if len(v7) > 0 {
                v7[0] = fi10
            }
            _ = len(v2)
            switch fi10 {
            case 117:
                v3 = (v3 >> ((258) & 63))
            }
        } else {
            if ((v3 <= 2) && ((-1 >= 127) || false)) {
                var v11 float64 = (float64((3.14159 * 0.0)) * (-2.25 / 1e10))
                _ = v11
            } else {
                v12 := GenS1{F0: 260, F1: v1, F2: 2147483647}
                fmt.Println(v3, fi10)
                _ = v12
            }
        }
        v1 += -213
        var v13 int = (fi10 % ((-9223372036854775808) | 1))
        _ = fi10
        _ = v13
    }
    v14 := GenS1{F0: 0, F1: v1, F2: 42}
    for fi15 := 0; fi15 < 10; fi15++ {
        if false {
            if true {
                if false {
                    if (true && (!false)) {
                        var v16 int8 = 42
                        switch 32768 {
                        case 958:
                            v16 -= (v16 + v16)
                        case 42:
                            v17 := GenS2{F0: ((!true) || false)}
                            _ = v17
                        default:
                            var v18 float32 = -2.25
                            _ = v18
                        }
                        for fi19 := 0; fi19 < 19; fi19++ {
                            v3 = v3
                            fmt.Println(v9, v4)
                            v20 := GenS1{F0: -620, F1: (v1 & 32767), F2: 42}
                            _ = fi19
                            _ = v20
                        }
                        _ = v16
                    } else {
                        var v21 int8 = -104
                        _ = v21
                    }
                }
                v9 = v4
                switch "\n\t" {
                case "a":
                    if len(v7) > 0 {
                        v7[0] = fi15
                    }
                case "hello":
                    v5 += uint64(v3)
                case "":
                    var v22 int16 = v1
                    _ = v22
                }
            }
            var v23 int64 = -548
            _ = v23
        } else {
            for fi24 := 0; fi24 < 6; fi24++ {
                v25 := v2[-743]
                _ = fi24
                _ = v25
            }
            v2[(-(fi15))] = true
        }
        _ = fi15
    }
    var v26 int8 = int8(v9)
    var v27 uint64 = (v5 % (((v5 << ((65535) & 63))) | 1))
    fmt.Println(v3, v5, v4)
    v28 := GenS1{F0: int64(v5), F1: (v1 | -32768), F2: int32(v9)}
    v5 = uint64(v4)
    v4 -= (^(v4))
    switch "" {
    case "hello":
        var v29 string = "hello"
        _ = v29
    }
    for fi30 := 0; fi30 < 20; fi30++ {
        fmt.Println(v27, v26, v27)
        _ = fi30
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v7:", len(v7), v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v14
    _ = v2
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
