package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16 }
type GenS2 struct { F0 uint; F1 int32 }
func genFn0(p0 int, p1 int) int8 {
    for fi5 := 0; fi5 < 19; fi5++ {
        for fi6 := 0; fi6 < 5; fi6++ {
            var v7 uint32 = 2
            v8 := GenS1{F0: 42, F1: -32768}
            var v9 uint64 = 2
            var v10 uint8 = 133
            _ = fi6
            _ = v10
            _ = v7
            _ = v8
            _ = v9
        }
        _ = fi5
    }
    p1 -= p1
    _ = p0
    _ = p1
    return 53
}

func main() {
    var v1 string = "a"
    var v2 float32 = 1e10
    v3 := make(map[string]string)
    v4 := make(map[int]int)
    var v11 float64 = 0.1
    v12 := GenS2{F0: 311, F1: 2147483647}
    v11 = float64((((((1e10 * (v2 / v2)) / float32((0.0 - v11))) + v2) / (float32(v11) * v2)) + float32(0.0)))
    if (genFn0(32767, 9223372036854775807) > 0) {
        v11 = v11
        fmt.Println(v1, v1)
    } else {
        for fi13 := 0; fi13 < 19; fi13++ {
            for fi14 := 0; fi14 < 5; fi14++ {
                v1 = ""
                v1 = v1
                var v15 uint8 = 128
                _ = fi14
                _ = v15
            }
            _ = fi13
        }
    }
    var v16 float32 = ((v2 + 3.14159) + v2)
    var v17 int8 = 127
    v11 -= 1.5
    for fi18 := 0; fi18 < 12; fi18++ {
        if (684 >= 1) {
            for fi19 := 0; fi19 < 12; fi19++ {
                fmt.Println(v1)
                v11 = (((v11 + float64(v2)) + 1.5) * float64(float32(v11)))
                var v20 int32 = 86
                _ = fi19
                _ = v20
            }
        }
        _ = fi18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v11
    _ = v12
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
