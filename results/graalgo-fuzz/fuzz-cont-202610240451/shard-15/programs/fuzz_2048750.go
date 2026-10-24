package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int }
type GenS2 struct { F0 uint64; F1 int32; F2 uint32 }

func main() {
    var v1 uint32 = 420
    var v2 uint64 = 42
    var v3 uint8 = 255
    if ((1 == -577) && ((v1 ^ 2147483647) < v1)) {
        v4 := GenS1{F0: 116, F1: -2147483648}
        v5 := []int{0, 128, 471}
        var v6 uint8 = (v3 + (v3 + v3))
        _ = v4
        _ = v5
        _ = v6
    }
    for fi7 := 0; fi7 < 17; fi7++ {
        var v8 int32 = 42
        _ = fi7
        _ = v8
    }
    var v9 uint = 440
    fmt.Println(v9)
    v1 = v1
    v10 := []int{-475, -662, 4294967295, 2147483647}
    for fi11 := 0; fi11 < 3; fi11++ {
        if (true || (!false)) {
            v9 = (v9 - (^(v9)))
        }
        if true {
            switch 263 {
            case -1:
                var v12 string = "a"
                _ = v12
            case 9223372036854775807:
                fmt.Println(v3, v3, v3)
            case 42:
                for fi13 := 0; fi13 < 12; fi13++ {
                    var v14 uint8 = (v3 % ((v3) | 1))
                    _ = fi13
                    _ = v14
                }
            }
            v1 -= (v1 % ((87) | 1))
            for fi15 := 0; fi15 < 8; fi15++ {
                v16 := GenS2{F0: v2, F1: -1, F2: (v1 | (v1 << (((v9 * v9)) & 63)))}
                var v17 uint32 = (^(v1))
                var v18 uint64 = v2
                v9 += (v9 & v9)
                fmt.Println(v2, fi15, v1)
                _ = fi15
                _ = v16
                _ = v17
                _ = v18
            }
        }
        _ = fi11
    }
    for fi19 := 0; fi19 < 9; fi19++ {
        v20 := GenS2{F0: (v2 * v2), F1: -2147483648, F2: (v1 + (v1 - uint32(v2)))}
        _ = fi19
        _ = v20
    }
    switch "hello" {
    case "\n\t":
        var v21 float32 = (((0.1 - 0.0) / (float32(((1e10 - 1e10) - (3.14159 + 1e10))) - (3.14159 - 1.5))) * (3.14159 - 1e10))
        _ = v21
    case "hello":
        switch 260 {
        case 128:
            for fi22 := 0; fi22 < 13; fi22++ {
                for fi23 := 0; fi23 < 16; fi23++ {
                    var v24 int = fi22
                    _ = fi23
                    _ = v24
                }
                v25 := GenS1{F0: 127, F1: 0}
                _ = fi22
                _ = v25
            }
        case -2147483648:
            var v26 int8 = 1
            _ = v26
        case 32767:
            for fi27 := 0; fi27 < 18; fi27++ {
                for fi28 := 0; fi28 < 4; fi28++ {
                    v3 = v3
                    v1 = (v1 >> ((v9) & 63))
                    v2 = v2
                    _ = fi28
                }
                switch fi27 {
                case 128:
                    var v29 string = "hello"
                    _ = v29
                case 127:
                    v10 = append(v10, int(v9))
                case 4294967295:
                    for fi30 := 0; fi30 < 2; fi30++ {
                        v1 += 1
                        _ = fi30
                    }
                default:
                    var v31 float64 = float64(1.5)
                    _ = v31
                }
                v32 := GenS1{F0: int8(fi27), F1: fi27}
                _ = fi27
                _ = v32
            }
        }
    case "":
        v1 -= (v1 / ((v1) | 1))
    }
    var v33 uint32 = 1
    for fi34 := 0; fi34 < 14; fi34++ {
        var v35 float32 = -2.25
        v9 += (v9 + (v9 * v9))
        var v36 string = "foo bar"
        _ = fi34
        _ = v35
        _ = v36
    }
    v33 = v33
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v33
    _ = v9
}
