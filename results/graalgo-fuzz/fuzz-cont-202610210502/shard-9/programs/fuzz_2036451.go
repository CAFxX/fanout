package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 uint8 }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 int32; F1 int8 }
type GenS4 struct { F0 uint8; F1 int16 }

func main() {
    var v1 uint64 = 2
    v2 := make(map[uint32]string)
    var v3 int64 = -9223372036854775808
    var v4 float64 = (float64((3.14159 - (3.14159 + float32(1e10)))) / -2.25)
    var v5 int32 = -307
    var v6 float64 = v4
    for fi7 := 0; fi7 < 2; fi7++ {
        var v8 int64 = v3
        v9 := GenS1{F0: false, F1: ("a" + ("hello" + "")), F2: 42}
        _ = fi7
        _ = v8
        _ = v9
    }
    fmt.Println(v1, v4)
    var v10 uint = 1
    var v11 int16 = -32768
    v6 += 3.14159
    var v12 uint8 = 51
    v3 = (v3 * (v3 >> (((^(v10))) & 63)))
    for fi13 := 0; fi13 < 1; fi13++ {
        v4 = (0.1 + v4)
        fmt.Println(v1)
        for fi14 := 0; fi14 < 12; fi14++ {
            v12 -= (v12 >> (((v10 * v10)) & 63))
            v11 = v11
            _ = fi14
        }
        _ = fi13
    }
    var v15 uint64 = (v1 % (((v1 >> ((v10) & 63))) | 1))
    switch -531 {
    case -2147483648:
        v16 := GenS1{F0: (int64(v15) > 0), F1: "\n\t", F2: (^(v12))}
        _ = v16
    case 4294967295:
        v17 := GenS2{F0: v11}
        _ = v17
    }
    v4 -= ((-2.25 - (v6 * (((v4 / (v6 / v6)) / (v4 / 3.14159)) + v6))) - v4)
    v18 := GenS3{F0: v5, F1: -1}
    v6 += (v4 - (float64((1e10 + 0.1)) - 1e10))
    v19 := GenS4{F0: (v12 - (v12 % ((v12) | 1))), F1: v11}
    var v20 bool = true
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
