package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 bool; F2 uint }
type GenS2 struct { F0 int32; F1 uint64; F2 uint8 }
type GenS3 struct { F0 uint32; F1 uint8; F2 int8 }
type GenS4 struct { F0 int16; F1 uint; F2 int }
type GenS5 struct { F0 bool; F1 int8; F2 int64 }
type GenS6 struct { F0 uint; F1 uint16 }

func main() {
    v1 := make(map[uint32]int)
    v2 := make(map[uint32]int)
    var v3 uint32 = 200
    var v4 int16 = 1000
    var v5 uint64 = 4294967295
    v3 -= (v3 / ((v3) | 1))
    var v6 int = 0
    var v7 float64 = 1e10
    v8 := []int{(^(v6)), v6, -2147483648}
    if (3.14159 > (float32(1e10) / 3.14159)) {
        if false {
            v4 += (v4 >> ((803) & 63))
            var v9 uint32 = (v3 / (((v3 >> ((4294967295) & 63))) | 1))
            var v10 int8 = -1
            _ = v10
            _ = v9
        }
        v5 = v5
        v11 := GenS1{F0: 1000, F1: (((-191 <= -88) || true) && true), F2: 131}
        _ = v11
    } else {
        for fi12 := 0; fi12 < 9; fi12++ {
            for fi13 := 0; fi13 < 2; fi13++ {
                v5 -= v5
                if len(v8) >= 2 {
                    v8 = v8[1:len(v8)]
                }
                if (1 < (v5 | v5)) {
                    fmt.Println(fi13)
                    fmt.Println(v6, v7)
                    switch "\n\t" {
                    case "hello":
                        v14 := GenS2{F0: 42, F1: 2, F2: 127}
                        _ = v14
                    }
                }
                _ = fi13
            }
            v3 -= 0
            _ = fi12
        }
        v15 := v2[(v3 + (v3 * v3))]
        _ = v15
    }
    switch (v6 << ((2) & 63)) {
    case -697:
        v3 = (v3 + 4294967295)
    case 4294967295:
        v16 := GenS2{F0: -770, F1: v5, F2: 42}
        _ = v16
    default:
        v4 -= v4
    }
    if len(v8) >= 2 {
        v8 = v8[1:len(v8)]
    }
    var v17 int8 = -1
    v18 := GenS3{F0: v3, F1: 1, F2: (v17 - v17)}
    var v19 float64 = float64((1.5 * ((float32((v7 - -2.25)) / -2.25) / 3.14159)))
    v20 := GenS4{F0: (v4 | v4), F1: 789, F2: (v6 ^ (v6 * v6))}
    if (!true) {
        if true {
            if true {
                for fi21 := 0; fi21 < 8; fi21++ {
                    _ = len(v8)
                    var v22 uint8 = 1
                    _ = fi21
                    _ = v22
                }
            } else {
                v23 := GenS5{F0: ((0 != 0) && ((-1 <= -9223372036854775808) || ((239 <= 2) && (!false)))), F1: v17, F2: 9223372036854775807}
                _ = v23
            }
            v4 -= (^(v4))
            v7 -= (-2.25 - (v19 + float64((1e10 / (1e10 * (0.1 - (1e10 - float32(v19))))))))
        } else {
            v24 := v2[1]
            _ = v24
        }
    }
    v25 := GenS3{F0: (v3 / ((v3) | 1)), F1: 2, F2: (^(v17))}
    v19 = v7
    var v26 int64 = -691
    v27 := GenS6{F0: uint(v3), F1: 589}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
