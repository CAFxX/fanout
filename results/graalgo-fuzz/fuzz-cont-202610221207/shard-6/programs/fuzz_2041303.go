package main

import (
    "fmt"
)

type GenS1 struct { F0 int8 }
type GenS2 struct { F0 uint16; F1 int8 }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 bool; F1 string; F2 int16 }

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]string)
    v3 := make(map[string]uint32)
    var v4 uint32 = 877
    var v5 float32 = ((-2.25 - 1e10) / 1.5)
    v6 := make(map[uint32]string)
    v7 := GenS1{F0: -128}
    var v8 uint8 = 42
    for fi9 := 0; fi9 < 4; fi9++ {
        var v10 int32 = -2147483648
        v8 = 40
        _ = fi9
        _ = v10
    }
    var v11 float32 = (float32((float64(float32(0.0)) + 1e10)) * (v5 - v5))
    if false {
        for fi12 := 0; fi12 < 2; fi12++ {
            var v13 int16 = 1
            _ = fi12
            _ = v13
        }
        if (v4 > v4) {
            v14 := GenS2{F0: 2, F1: 1}
            switch "\n\t" {
            case "foo bar":
                var v15 float64 = 1e10
                _ = v15
            case "a":
                v11 += (3.14159 * v5)
            case "":
                v5 += (v11 + float32(float64((v11 / v11))))
            }
            fmt.Println(v11, v11)
            _ = v14
        } else {
            v16 := GenS2{F0: 2, F1: 127}
            _ = v16
        }
    }
    v17 := GenS1{F0: 1}
    if false {
        for fi18 := 0; fi18 < 2; fi18++ {
            if (!(42 > -117)) {
                v19 := GenS1{F0: -1}
                _ = v19
            }
            _ = fi18
        }
        var v20 int64 = -9223372036854775808
        _ = v20
    }
    switch "" {
    case "hello":
        var v21 int8 = -1
        _ = v21
    case "\n\t":
        var v22 uint16 = 32767
        _ = v22
    default:
        if false {
            for fi23 := 0; fi23 < 12; fi23++ {
                fmt.Println(v8, v5)
                v24 := GenS3{F0: -686}
                var v25 uint = 1
                _ = fi23
                _ = v24
                _ = v25
            }
            fmt.Println(v5)
        } else {
            v26 := GenS4{F0: true, F1: "", F2: -922}
            _ = v26
        }
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v11
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
