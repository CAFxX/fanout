package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int8; F1 uint32; F2 bool }
func genFn0(p0 int, p1 uint32, p2 int8) uint64 {
    p1 = p1
    v7 := GenS1{F0: p0}
    var v8 uint32 = 4294967295
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 2
}

func main() {
    v1 := make(map[int]bool)
    v2 := make(map[int]string)
    var v3 float32 = ((-2.25 + 0.1) + 0.0)
    var v4 int16 = 1000
    v5 := []uint32{}
    var v6 uint8 = 104
    if true {
        switch "a" {
        case "\n\t":
            var v9 int64 = 467
            _ = v9
        case "hello":
            for fi10 := 0; fi10 < 7; fi10++ {
                if false {
                    var v11 string = "foo bar"
                    v12 := GenS1{F0: (fi10 % (((^(fi10))) | 1))}
                    for fi13 := 0; fi13 < 1; fi13++ {
                        v3 += -2.25
                        v14 := GenS1{F0: 2147483647}
                        _ = fi13
                        _ = v14
                    }
                    _ = v11
                    _ = v12
                }
                _ = fi10
            }
        case "a":
            for fi15 := 0; fi15 < 16; fi15++ {
                v6 = (^(v6))
                _ = fi15
            }
        default:
            v6 += (^(v6))
        }
        if true {
            var v16 int8 = 100
            v6 = 147
            if true {
                switch 42 {
                case 827:
                    var v17 bool = (756 <= 2)
                    _ = v17
                default:
                    v5 = append(v5, 65535)
                }
                var v18 int32 = 1
                if false {
                    v18 += v18
                    v2[42] = "foo bar"
                } else {
                    _ = len(v5)
                }
                _ = v18
            }
            _ = v16
        } else {
            v3 = ((((((((-2.25 / 0.1) / (float32(float64(v3)) * ((1.5 / 1.0) - ((((((-2.25 + float32(1e10)) / (v3 / 1e10)) + float32((3.14159 - float64(v3)))) / (v3 * (1.5 + 1e10))) / v3) / 1.5)))) - 0.1) - v3) + v3) * v3) - 0.1) / 3.14159)
            var v19 string = ("\n\t" + "")
            var v20 string = v19
            _ = v19
            _ = v20
        }
    } else {
        var v21 uint = 573
        _ = v21
    }
    var v22 int32 = 42
    delete(v1, 68)
    v23 := GenS1{F0: 2147483647}
    if (v22 <= (-(v22))) {
        switch 0 {
        case 1:
            var v24 uint16 = 983
            _ = v24
        case -1:
            var v25 int32 = (-(v22))
            _ = v25
        case 278:
            var v26 uint32 = 1
            _ = v26
        default:
            v22 = 2147483647
        }
        v27 := v1[128]
        var v28 uint64 = 4294967295
        _ = v27
        _ = v28
    } else {
        var v29 uint64 = 0
        v22 -= (-(v22))
        _ = v29
    }
    var v30 uint8 = (v6 | (v6 * (^(v6))))
    fmt.Println(v4, v22, v22)
    fmt.Println(v3, v4)
    v31 := GenS1{F0: 2147483647}
    fmt.Println(v6)
    fmt.Println(v22, v3)
    v32 := GenS2{F0: 3, F1: 931, F2: true}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v30:", v30)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v2
    _ = v22
    _ = v23
    _ = v3
    _ = v30
    _ = v31
    _ = v32
    _ = v4
    _ = v5
    _ = v6
}
