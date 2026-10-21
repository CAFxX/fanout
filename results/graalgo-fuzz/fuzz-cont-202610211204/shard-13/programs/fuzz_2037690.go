package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
type GenS2 struct { F0 int32; F1 bool }

func main() {
    var v1 float64 = (-2.25 - (0.1 * 3.14159))
    var v2 int16 = 32767
    v3 := make(map[string]uint32)
    var v4 float64 = 1e10
    var v5 int64 = 1
    for fi6 := 0; fi6 < 3; fi6++ {
        v7 := GenS1{F0: 1, F1: (-(v5))}
        fmt.Println(v4, v2)
        _ = fi6
        _ = v7
    }
    delete(v3, "")
    v8 := []int{42}
    v1 = v4
    v9 := GenS1{F0: 815, F1: (v5 * -1)}
    for fi10 := 0; fi10 < 13; fi10++ {
        fmt.Println(v4, v2, v4)
        var v11 uint64 = 2
        var v12 int16 = 42
        _ = fi10
        _ = v11
        _ = v12
    }
    if true {
        v13 := GenS2{F0: 811, F1: false}
        var v14 uint32 = 2
        _ = v13
        _ = v14
    } else {
        for fi15 := 0; fi15 < 18; fi15++ {
            for fi16 := 0; fi16 < 6; fi16++ {
                var v17 float32 = 0.1
                switch fi15 {
                case 1:
                    switch fi15 {
                    case -833:
                        var v18 float64 = v4
                        _ = v18
                    default:
                        fmt.Println(v5)
                        fmt.Println(v5)
                        _ = len(v3)
                    }
                case -955:
                    v2 = (v2 >> ((42) & 63))
                default:
                    v17 += float32(float64((v17 / 1.5)))
                }
                var v19 uint8 = 64
                if (!true) {
                    var v20 bool = true
                    _ = v20
                } else {
                    v21 := GenS2{F0: 607, F1: (18446744073709551615 == 255)}
                    v5 = 42
                    v17 += (((1.5 / v17) - (((-2.25 - ((v17 * (((0.1 / float32(3.14159)) / v17) * v17)) + float32(-2.25))) * (v17 - v17)) * (-2.25 + (float32(v4) / float32((3.14159 / float64(-2.25))))))) + (float32(v4) * (0.1 / v17)))
                    _ = v21
                }
                _ = fi16
                _ = v17
                _ = v19
            }
            _ = fi15
        }
        for fi22 := 0; fi22 < 14; fi22++ {
            var v23 uint8 = 42
            for fi24 := 0; fi24 < 18; fi24++ {
                var v25 string = "hello"
                var v26 uint = uint(v23)
                _ = fi24
                _ = v25
                _ = v26
            }
            _ = fi22
            _ = v23
        }
    }
    v5 = 105
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
