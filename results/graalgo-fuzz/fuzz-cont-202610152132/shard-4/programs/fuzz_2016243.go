package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 uint64; F2 int32 }
type GenS2 struct { F0 int8 }
type GenS3 struct { F0 uint64; F1 uint16; F2 uint32 }
func genFn0(p0 uint, p1 int64) int16 {
    v5 := []int{}
    var v6 uint = (p0 | 68)
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    return 1000
}

func main() {
    var v1 bool = false
    v2 := make(map[uint32]uint32)
    var v3 float32 = 1e10
    var v4 uint64 = 611
    var v7 uint8 = 128
    v8 := GenS1{F0: 127, F1: v4, F2: 507}
    switch "\n\t" {
    case "hello":
        v3 = (v3 / ((v3 + (1.5 * ((((v3 - v3) * (1.5 - ((float32(float64(v3)) / (float32((1.5 / 3.14159)) * (1e10 + v3))) - (-2.25 * float32((1.5 - 3.14159)))))) * float32(float64(v3))) - -2.25))) * (v3 - 1.5)))
    case "a":
        v9 := GenS2{F0: -22}
        _ = v9
    case "":
        v10 := GenS1{F0: -1, F1: (^(v4)), F2: 42}
        _ = v10
    default:
        if ((255 != 1) && (!(155 >= 1))) {
            fmt.Println(v4, v4)
        }
    }
    for fi11 := 0; fi11 < 8; fi11++ {
        for fi12 := 0; fi12 < 6; fi12++ {
            var v13 int = (fi11 - fi11)
            var v14 uint8 = 0
            _ = fi12
            _ = v13
            _ = v14
        }
        _ = fi11
    }
    for fi15 := 0; fi15 < 3; fi15++ {
        if ((0.1 / float32(float64(v3))) <= v3) {
            v16 := v2[65535]
            _ = v16
        } else {
            var v17 bool = v1
            _ = v17
        }
        _ = fi15
    }
    fmt.Println(v1, v4)
    v18 := GenS3{F0: 3, F1: 0, F2: 2147483647}
    for fi19 := 0; fi19 < 5; fi19++ {
        v1 = (102 >= 2147483647)
        _ = fi19
    }
    fmt.Println("v1:", v1)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v7
    _ = v8
}
