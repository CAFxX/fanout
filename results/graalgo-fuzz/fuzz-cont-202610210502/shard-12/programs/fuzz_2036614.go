package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint; F1 int64; F2 int }
type GenS3 struct { F0 int32 }
type GenS4 struct { F0 int32 }
type GenS5 struct { F0 uint32 }
type GenS6 struct { F0 int16; F1 uint32; F2 string }

func main() {
    v1 := make(map[uint32]int)
    var v2 int32 = -390
    var v3 bool = true
    var v4 float64 = 1e10
    v5 := make(map[int]bool)
    var v6 int16 = 536
    v2 = v2
    v4 -= 1.5
    v7 := GenS1{F0: -698}
    var v8 float64 = 1e10
    v4 = -2.25
    fmt.Println(v2)
    switch -2147483648 {
    case 532:
        for fi9 := 0; fi9 < 12; fi9++ {
            v3 = (-9223372036854775808 <= 566)
            var v10 int = fi9
            _ = fi9
            _ = v10
        }
    }
    switch 4294967295 {
    case -2147483648:
        if (0.0 >= float64(0.0)) {
            var v11 float32 = 1e10
            switch 0 {
            case 4294967295:
                v2 = (-(v2))
            case -754:
                v6 += v6
            default:
                for fi12 := 0; fi12 < 12; fi12++ {
                    v6 += -288
                    v11 += ((v11 * (v11 * 0.0)) * (v11 + v11))
                    switch "\n\t" {
                    case "hello":
                        v13 := GenS2{F0: 979, F1: -288, F2: fi12}
                        _ = v13
                    }
                    _ = fi12
                }
            }
            v14 := GenS3{F0: 2147483647}
            _ = v11
            _ = v14
        }
    default:
        v6 += v6
    }
    fmt.Println(v6, v8)
    v3 = (1 != 255)
    v15 := GenS4{F0: 2147483647}
    var v16 int16 = (v6 ^ (v6 & (v6 % ((32767) | 1))))
    v17 := GenS2{F0: 815, F1: -9223372036854775808, F2: 127}
    if true {
        if ((1e10 < (0.0 + 0.1)) || true) {
            var v18 bool = v3
            var v19 uint8 = 123
            v20 := GenS5{F0: 109}
            _ = v18
            _ = v19
            _ = v20
        }
        v21 := GenS6{F0: v6, F1: 0, F2: "hello"}
        _ = v21
    }
    var v22 float64 = v4
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
}
