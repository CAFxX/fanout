package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint64; F2 int8 }
type GenS2 struct { F0 int8 }

func main() {
    var v1 int32 = 42
    var v2 string = ("" + "foo bar")
    var v3 string = "a"
    v4 := []uint32{479, 112, 830}
    var v5 int64 = 9223372036854775807
    for fi6 := 0; fi6 < 8; fi6++ {
        switch fi6 {
        case 128:
            var v7 int16 = -32768
            _ = v7
        case 496:
            for fi8 := 0; fi8 < 14; fi8++ {
                v2 = ("foo bar" + v2)
                _ = fi8
            }
        case 32767:
            v9 := GenS1{F0: 0, F1: 0, F2: -128}
            _ = v9
        }
        var v10 float32 = 0.1
        v1 -= (-(v1))
        _ = fi6
        _ = v10
    }
    var v11 int64 = 42
    for fi12 := 0; fi12 < 5; fi12++ {
        fmt.Println(v1)
        _ = len(v4)
        v2 = (v3 + "foo bar")
        _ = fi12
    }
    v13 := GenS2{F0: -1}
    if (!((false || (1 == (v1 / (((v1 | (v1 * (v1 * v1)))) | 1)))) || false)) {
        if len(v4) > 0 {
            v4[0] = 1
        }
    } else {
        _ = len(v4)
        v11 = (v5 / (((v5 * v5)) | 1))
    }
    var v14 uint8 = 128
    var v15 bool = false
    fmt.Println(v3)
    v14 = (v14 | 50)
    if v15 {
        _ = len(v4)
        for fi16 := 0; fi16 < 15; fi16++ {
            switch v2 {
            case "a":
                if v15 {
                    var v17 uint8 = 1
                    _ = v17
                } else {
                    var v18 string = v2
                    v11 += 0
                    v5 = v11
                    _ = v18
                }
            default:
                for fi19 := 0; fi19 < 12; fi19++ {
                    var v20 int8 = -1
                    var v21 int16 = 1000
                    v15 = (1.5 == (float32(float64(float32((1.5 + 1e10)))) * (1e10 + 0.0)))
                    _ = fi19
                    _ = v20
                    _ = v21
                }
            }
            v22 := GenS2{F0: -128}
            v14 += v14
            _ = fi16
            _ = v22
        }
    }
    var v23 int32 = 2147483647
    v3 = ("foo bar" + (v3 + ""))
    v23 = v1
    fmt.Println(v5, v23)
    v4 = append(v4, 2)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
