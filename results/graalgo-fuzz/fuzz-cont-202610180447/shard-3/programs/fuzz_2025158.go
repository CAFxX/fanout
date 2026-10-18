package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int64; F2 int16 }
type GenS2 struct { F0 int16 }

func main() {
    v1 := []uint32{42, 0, 4294967295, 42}
    var v2 uint32 = 2147483648
    var v3 float32 = float32(0.0)
    v4 := make(map[int]string)
    v2 += (v2 ^ (v2 - 4294967295))
    fmt.Println(v2, v3)
    if (-1 > -9223372036854775808) {
        v2 = v2
        switch "" {
        case "hello":
            switch 2147483647 {
            case 32768:
                for fi5 := 0; fi5 < 19; fi5++ {
                    if true {
                        var v6 uint16 = 2
                        v7 := GenS1{F0: (v2 - (v2 * v2)), F1: 42, F2: 356}
                        _ = v6
                        _ = v7
                    }
                    _ = fi5
                }
            case -125:
                v8 := GenS1{F0: v2, F1: 42, F2: 0}
                _ = v8
            default:
                var v9 uint64 = 137
                _ = v9
            }
        }
        if false {
            fmt.Println(v2)
        } else {
            fmt.Println(v3)
            if false {
                for fi10 := 0; fi10 < 7; fi10++ {
                    v2 = (v2 | v2)
                    var v11 int16 = 1000
                    _ = fi10
                    _ = v11
                }
                v3 = (((v3 / v3) / v3) - float32(0.1))
            } else {
                v3 = (-2.25 * v3)
                var v12 uint16 = 0
                fmt.Println(v2, v3, v3)
                _ = v12
            }
        }
    } else {
        var v13 int8 = -37
        var v14 float32 = v3
        _ = v13
        _ = v14
    }
    var v15 float32 = v3
    var v16 float64 = float64(float32(1e10))
    v17 := GenS2{F0: 42}
    for fi18 := 0; fi18 < 13; fi18++ {
        var v19 uint8 = 230
        _ = fi18
        _ = v19
    }
    v3 += (v15 / 1.5)
    var v20 int64 = -769
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v3
    _ = v4
}
