package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint; F2 uint8 }
func genFn0() int8 {
    v6 := GenS1{F0: -353, F1: 2, F2: 1}
    var v7 int = 32767
    v7 -= v7
    _ = v6
    _ = v7
    return -57
}

func main() {
    var v1 int16 = 42
    var v2 float64 = 0.1
    var v3 float64 = 3.14159
    v4 := make(map[uint32]uint32)
    v5 := []string{}
    if (950 <= 65535) {
        for fi8 := 0; fi8 < 17; fi8++ {
            var v9 int64 = 272
            _ = fi8
            _ = v9
        }
        fmt.Println(v2, v2, v3)
    }
    for fi10 := 0; fi10 < 13; fi10++ {
        v2 -= v3
        switch fi10 {
        case -453:
            var v11 float32 = (-2.25 - (-2.25 / ((float32(v3) - 0.1) / ((1e10 * -2.25) - 3.14159))))
            _ = v11
        case 855:
            v3 = 1.5
        case -9223372036854775808:
            if len(v5) >= 2 {
                v5 = v5[1:len(v5)]
            }
        }
        var v12 bool = (!(-2.25 != 0.1))
        _ = fi10
        _ = v12
    }
    var v13 uint = 857
    v13 = v13
    var v14 float64 = 0.0
    v2 = (((v14 * ((0.1 * float64(3.14159)) + float64(0.1))) / (float64((1.5 + 0.0)) + ((v14 * v3) * v3))) - v3)
    fmt.Println(v3)
    fmt.Println(v14, v1, v14)
    var v15 bool = true
    v2 += float64(1.5)
    fmt.Println(v15)
    v16 := GenS1{F0: 0, F1: v13, F2: 141}
    var v17 uint16 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
