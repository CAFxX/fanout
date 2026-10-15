package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 uint16 }
func genFn0() string {
    var v6 uint8 = 128
    fmt.Println(v6)
    v7 := []int{1}
    v6 = v6
    var v8 uint = 65535
    _ = v6
    _ = v7
    _ = v8
    return "\n\t"
}

func main() {
    var v1 int16 = 759
    var v2 float32 = -2.25
    v3 := make(map[int]uint32)
    var v4 int32 = 206
    v5 := []uint32{1, 42}
    var v9 int16 = 1000
    if len(v5) >= 2 {
        v5 = v5[1:len(v5)]
    }
    if false {
        var v10 uint64 = 2
        _ = v10
    } else {
        if (32768 < 827) {
            var v11 int8 = 127
            _ = v11
        }
    }
    v1 = 42
    for fi12 := 0; fi12 < 17; fi12++ {
        fmt.Println(v1, fi12, v1)
        _ = fi12
    }
    v4 = v4
    var v13 float64 = ((1e10 / (-2.25 - float64(float32(1e10)))) * ((1e10 + (1e10 + (0.1 + float64(v2)))) + ((1.5 * 1e10) / 0.1)))
    for fi14 := 0; fi14 < 20; fi14++ {
        fmt.Println(v1, v1, v1)
        v4 += (v4 >> ((1) & 63))
        _ = fi14
    }
    if (0.1 != (v13 / v13)) {
        v3[-762] = 65535
        switch "a" {
        case "a":
            for fi15 := 0; fi15 < 10; fi15++ {
                var v16 uint8 = 128
                v9 += (v1 >> ((182) & 63))
                var v17 int32 = (v4 * v4)
                _ = fi15
                _ = v16
                _ = v17
            }
        }
        if (true || (true && (-172 >= -9223372036854775808))) {
            var v18 uint64 = 546
            _ = v18
        } else {
            var v19 int32 = v4
            var v20 int8 = 127
            _ = v19
            _ = v20
        }
    }
    for fi21 := 0; fi21 < 19; fi21++ {
        for fi22 := 0; fi22 < 10; fi22++ {
            v23 := GenS1{F0: 2147483647, F1: 42, F2: 700}
            v13 += v13
            fmt.Println(v4, v9)
            _ = fi22
            _ = v23
        }
        _ = fi21
    }
    if len(v5) > 0 {
        v5[0] = 635
    }
    v3[0] = 469
    v2 = (((v2 + v2) * 3.14159) / v2)
    v13 += 0.0
    if ((v4 - (v4 - 42)) <= v4) {
        var v24 int = 127
        _ = v24
    } else {
        var v25 uint8 = 128
        _ = v25
    }
    v2 -= ((v2 * (v2 / (float32(1.5) / v2))) / v2)
    if (-78 != 101) {
        v4 += (v4 - v4)
        v1 -= v1
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
