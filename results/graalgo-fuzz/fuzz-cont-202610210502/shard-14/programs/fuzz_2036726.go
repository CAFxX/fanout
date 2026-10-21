package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int32 }
type GenS2 struct { F0 bool }
type GenS3 struct { F0 uint16; F1 int32; F2 int32 }
func genFn0(p0 int) int32 {
    for fi6 := 0; fi6 < 19; fi6++ {
        v7 := GenS1{F0: p0, F1: 1}
        _ = fi6
        _ = v7
    }
    p0 += 822
    _ = p0
    return -132
}

func main() {
    var v1 int64 = 0
    v2 := make(map[string]bool)
    v3 := []uint32{397, 4294967295, 65535, 0}
    var v4 uint16 = 63
    v5 := make(map[int]string)
    v8 := GenS2{F0: false}
    var v9 float32 = float32(0.1)
    _ = len(v2)
    var v10 int16 = 42
    v11 := GenS1{F0: 128, F1: -368}
    v12 := GenS3{F0: v4, F1: 2147483647, F2: -930}
    v13 := GenS2{F0: true}
    if (0.1 != 1e10) {
        v14 := GenS2{F0: true}
        if (v9 <= ((v9 + v9) + (1e10 / v9))) {
            var v15 float32 = v9
            switch -521 {
            case 127:
                var v16 uint8 = 42
                _ = v16
            case -1:
                var v17 int64 = (v1 | (^(v1)))
                _ = v17
            case -606:
                var v18 bool = (((4294967295 > 512) && true) || true)
                _ = v18
            }
            v1 += (v1 & (v1 + v1))
            _ = v15
        }
        fmt.Println(v1)
        _ = v14
    }
    v9 += (float32(float64(3.14159)) * 1.5)
    if true {
        v9 = v9
        v4 = (v4 * v4)
    } else {
        var v19 float32 = 3.14159
        v20 := GenS2{F0: (557 < 255)}
        _ = v19
        _ = v20
    }
    switch (("hello" + "") + "a") {
    case "hello":
        if (!false) {
            if len(v3) >= 2 {
                v3 = v3[1:len(v3)]
            }
        } else {
            fmt.Println(v9, v1, v4)
        }
    default:
        v21 := GenS2{F0: false}
        _ = v21
    }
    fmt.Println(v1, v10, v9)
    v1 -= 9223372036854775807
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
