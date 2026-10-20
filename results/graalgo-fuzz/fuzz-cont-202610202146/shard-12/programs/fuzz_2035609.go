package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint; F2 uint16 }
type GenS2 struct { F0 int64 }
type GenS3 struct { F0 int32; F1 uint8 }
type GenS4 struct { F0 int; F1 uint8; F2 uint }
type GenS5 struct { F0 uint16; F1 int64 }

func main() {
    v1 := make(map[string]int)
    var v2 uint64 = 2
    var v3 int8 = 127
    var v4 uint = 281
    v5 := GenS1{F0: 128, F1: 255, F2: 65535}
    v6 := GenS2{F0: 120}
    if (-18 != 0) {
        var v7 uint16 = 32767
        v4 = v4
        v2 -= v2
        _ = v7
    } else {
        var v8 int64 = -719
        var v9 uint64 = (v2 + 18446744073709551615)
        _ = v8
        _ = v9
    }
    var v10 uint16 = 2
    var v11 float32 = (0.0 + 1e10)
    var v12 float32 = v11
    v13 := GenS2{F0: 1000}
    v10 = v10
    for fi14 := 0; fi14 < 3; fi14++ {
        for fi15 := 0; fi15 < 19; fi15++ {
            var v16 uint64 = v2
            var v17 int16 = 0
            v18 := GenS2{F0: -9223372036854775808}
            var v19 float64 = 0.1
            var v20 float32 = v12
            _ = fi15
            _ = v16
            _ = v17
            _ = v18
            _ = v19
            _ = v20
        }
        _ = fi14
    }
    var v21 uint16 = (v10 % ((255) | 1))
    if ((1e10 / -2.25) == -2.25) {
        v12 = v12
        v12 = v12
    } else {
        v22 := GenS3{F0: 1, F1: 42}
        _ = v22
    }
    for fi23 := 0; fi23 < 3; fi23++ {
        if (false && (!false)) {
            switch (fi23 + (fi23 | fi23)) {
            case -1:
                v24 := GenS4{F0: 0, F1: 2, F2: (v4 + (v4 % (((v4 ^ (^(v4)))) | 1)))}
                _ = v24
            case 910:
                v25 := GenS5{F0: (v21 / ((v21) | 1)), F1: 9223372036854775807}
                _ = v25
            }
        }
        var v26 uint8 = 128
        _ = fi23
        _ = v26
    }
    fmt.Println(v12, v12)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
