package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 string; F1 int16 }
type GenS3 struct { F0 uint8; F1 int16; F2 int32 }

func main() {
    var v1 float64 = 3.14159
    var v2 bool = false
    var v3 int64 = -1
    var v4 string = "a"
    v5 := []int{-664, 662}
    v6 := make(map[string]uint32)
    v7 := GenS1{F0: -1, F1: -118}
    v3 += (v3 << ((65535) & 63))
    v2 = (v3 < -350)
    if v2 {
        var v8 int8 = -65
        _ = v8
    }
    for fi9 := 0; fi9 < 3; fi9++ {
        for fi10 := 0; fi10 < 1; fi10++ {
            var v11 float32 = 1e10
            v12 := GenS2{F0: v4, F1: 1}
            var v13 int16 = 32767
            _ = fi10
            _ = v11
            _ = v12
            _ = v13
        }
        v4 = ("\n\t" + "foo bar")
        _ = fi9
    }
    fmt.Println(v4, v1)
    v3 = v3
    var v14 uint8 = 207
    if v2 {
        var v15 int16 = 0
        v3 = (v3 >> ((630) & 63))
        _ = v15
    }
    if (!((true || false) && v2)) {
        if len(v5) > 0 {
            v5[0] = -897
        }
        for fi16 := 0; fi16 < 12; fi16++ {
            var v17 bool = true
            switch fi16 {
            case 845:
                v18 := GenS3{F0: (v14 / ((v14) | 1)), F1: -865, F2: 54}
                _ = v18
            }
            _ = fi16
            _ = v17
        }
    }
    v3 -= -1
    switch 9223372036854775807 {
    case -9223372036854775808:
        var v19 float32 = float32((1e10 * float64(3.14159)))
        _ = v19
    case 469:
        var v20 float32 = (-2.25 + 1.5)
        _ = v20
    case 128:
        if (v3 < (^(v3))) {
            fmt.Println(v4, v4, v14)
            var v21 float32 = (-2.25 * float32(float64(0.1)))
            _ = v21
        }
    default:
        v5 = append(v5, 853)
    }
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
