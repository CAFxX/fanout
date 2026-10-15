package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 int = -285
    v2 := make(map[string]uint32)
    v3 := make(map[string]bool)
    v4 := []int{(v1 >> ((2) & 63)), v1, (v1 | 2147483647)}
    v5 := []uint32{130}
    if ((0.1 / (float32((0.0 / (3.14159 + 0.0))) - float32(((float64((3.14159 + 0.1)) / 1.0) * 3.14159)))) != ((float32(((0.0 + ((1e10 * 0.1) + 3.14159)) - 3.14159)) - 0.1) / 1.0)) {
        var v6 bool = (true || ((false || (false || false)) || (!false)))
        _ = v6
    }
    if ((1.5 >= (((1e10 / 1e10) * 0.0) - -2.25)) || false) {
        var v7 float32 = -2.25
        _ = v7
    }
    fmt.Println(v1)
    for fi8 := 0; fi8 < 19; fi8++ {
        var v9 int = (v1 ^ (fi8 >> ((752) & 63)))
        _ = fi8
        _ = v9
    }
    v10 := GenS1{F0: 301}
    var v11 string = "hello"
    fmt.Println(v11, v1)
    _ = len(v3)
    if true {
        var v12 uint8 = 42
        _ = v12
    } else {
        v13 := GenS1{F0: 248}
        _ = v13
    }
    var v14 bool = false
    if v14 {
        var v15 string = v11
        v1 -= v1
        _ = v15
    }
    v1 = (v1 - v1)
    for fi16 := 0; fi16 < 10; fi16++ {
        if v14 {
            for fi17 := 0; fi17 < 20; fi17++ {
                if (0 <= 491) {
                    switch fi16 {
                    case 1:
                        var v18 int = (v1 ^ v1)
                        _ = v18
                    case 127:
                        var v19 uint64 = 0
                        _ = v19
                    default:
                        var v20 uint16 = 1
                        _ = v20
                    }
                    v21 := GenS1{F0: 420}
                    v11 = ("" + v11)
                    _ = v21
                } else {
                    var v22 uint64 = 2
                    _ = v22
                }
                var v23 uint8 = 178
                _ = fi17
                _ = v23
            }
            var v24 uint = 652
            v1 = (v1 << ((v24) & 63))
            _ = v24
        }
        fmt.Println(fi16, v14)
        var v25 int32 = 0
        _ = fi16
        _ = v25
    }
    v11 = ("hello" + v11)
    v26 := GenS1{F0: 935}
    var v27 float32 = 1e10
    fmt.Println(v27, v1, v14)
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v27:", v27)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v14
    _ = v2
    _ = v26
    _ = v27
    _ = v3
    _ = v4
    _ = v5
}
