package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 bool; F1 int16 }

func main() {
    var v1 uint8 = 2
    var v2 uint8 = (v1 / ((v1) | 1))
    v3 := make(map[string]int)
    var v4 float64 = 3.14159
    var v5 bool = (true && (4294967295 != 566))
    var v6 int32 = 0
    fmt.Println(v1, v5, v5)
    switch "\n\t" {
    case "a":
        var v7 int32 = (v6 ^ -428)
        _ = v7
    case "foo bar":
        v6 = 0
    case "":
        v1 -= (v1 ^ 42)
    }
    if (36 <= 0) {
        v8 := GenS1{F0: 0}
        v9 := GenS1{F0: -2147483648}
        if (65535 <= 0) {
            for fi10 := 0; fi10 < 10; fi10++ {
                v1 += v2
                _ = fi10
            }
            v2 = 2
        }
        _ = v8
        _ = v9
    }
    v6 = v6
    v1 += (v1 << ((875) & 63))
    var v11 int64 = -9223372036854775808
    for fi12 := 0; fi12 < 13; fi12++ {
        v13 := GenS2{F0: v5}
        v14 := GenS1{F0: 127}
        _ = fi12
        _ = v13
        _ = v14
    }
    v4 = v4
    v5 = ((float32((1.5 * 3.14159)) / 1e10) >= 0.1)
    var v15 int16 = -32768
    v16 := GenS2{F0: (v11 == v11)}
    v11 = v11
    for fi17 := 0; fi17 < 2; fi17++ {
        for fi18 := 0; fi18 < 7; fi18++ {
            var v19 uint = 515
            var v20 uint64 = 600
            _ = fi18
            _ = v19
            _ = v20
        }
        fmt.Println(fi17)
        _ = fi17
    }
    v21 := GenS3{F0: v5, F1: v15}
    v1 = v2
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
