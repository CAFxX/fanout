package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int8; F2 int64 }

func main() {
    var v1 int64 = -764
    var v2 int64 = v1
    var v3 float64 = 0.1
    v4 := []int{1, -376}
    var v5 uint64 = 249
    for fi6 := 0; fi6 < 12; fi6++ {
        if len(v4) > 0 {
            v4[0] = fi6
        }
        _ = fi6
    }
    for fi7 := 0; fi7 < 4; fi7++ {
        fmt.Println(v1, v2)
        fmt.Println(v2)
        _ = fi7
    }
    switch "hello" {
    case "a":
        for fi8 := 0; fi8 < 9; fi8++ {
            fmt.Println(v1, fi8, v5)
            fmt.Println(v3)
            var v9 uint = 930
            _ = fi8
            _ = v9
        }
    }
    var v10 uint = 883
    _ = len(v4)
    if len(v4) > 0 {
        v4[0] = 32768
    }
    var v11 float64 = v3
    for fi12 := 0; fi12 < 6; fi12++ {
        v11 += (float64(0.1) * ((((0.1 + 1e10) + v11) - (float64((float32(((v3 + v11) - v11)) + float32(v11))) - ((float64(float32(v3)) + v3) * v11))) + ((v11 - v11) / (v11 - v3))))
        if len(v4) >= 2 {
            v4 = v4[1:len(v4)]
        }
        _ = fi12
    }
    v2 = v2
    v11 = v11
    if (false || true) {
        if len(v4) > 0 {
            v4[0] = -933
        }
        if true {
            if false {
                v13 := GenS1{F0: 2147483647, F1: 127, F2: 116}
                var v14 uint8 = 185
                _ = v13
                _ = v14
            } else {
                var v15 int16 = 365
                _ = v15
            }
        }
    }
    var v16 int64 = v2
    var v17 uint8 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
