package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }

func main() {
    var v1 uint = 29
    v2 := make(map[int]uint32)
    v3 := []int{}
    v4 := []int{401, -431, -9223372036854775808, 128}
    v5 := GenS1{F0: 32767}
    v6 := GenS1{F0: 42}
    for fi7 := 0; fi7 < 15; fi7++ {
        var v8 uint16 = 255
        var v9 int32 = 510
        _ = fi7
        _ = v8
        _ = v9
    }
    var v10 int64 = 0
    var v11 int8 = 0
    delete(v2, 2147483647)
    var v12 float32 = (float32(1.5) / 1.0)
    var v13 uint8 = 128
    fmt.Println(v13, v12, v11)
    v4 = append(v4, 9223372036854775807)
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    v1 -= (v1 >> (((v1 << ((4294967295) & 63))) & 63))
    for fi14 := 0; fi14 < 1; fi14++ {
        v13 -= (v13 - 236)
        _ = fi14
    }
    if true {
        if true {
            var v15 int64 = v10
            v11 -= v11
            if len(v4) >= 2 {
                v4 = v4[1:len(v4)]
            }
            _ = v15
        }
    }
    v12 -= ((float32(0.1) / (v12 * v12)) / (-2.25 / float32(-2.25)))
    var v16 int64 = int64(v1)
    _ = len(v2)
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
