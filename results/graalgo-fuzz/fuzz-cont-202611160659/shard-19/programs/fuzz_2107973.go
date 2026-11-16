package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
type GenS2 struct { F0 int16; F1 int16; F2 uint16 }
type GenS3 struct { F0 int64; F1 string }

func main() {
    v1 := make(map[uint32]bool)
    var v2 int16 = 42
    v3 := []uint32{302}
    var v4 bool = false
    var v5 int32 = -578
    fmt.Println(v4)
    for fi6 := 0; fi6 < 18; fi6++ {
        var v7 uint = 4294967295
        v2 = (v2 % ((373) | 1))
        _ = fi6
        _ = v7
    }
    var v8 int16 = -1
    if (false || (v5 <= 958)) {
        v8 = int16(v5)
    }
    var v9 float32 = float32(float64((1.5 + (float32(float64(1e10)) - (-2.25 + (1.5 + (3.14159 + (float32(((1e10 * (float64(1.5) * 0.0)) * (-2.25 + float64(0.0)))) * 1e10))))))))
    fmt.Println(v4)
    v5 = v5
    if len(v3) > 0 {
        v3[0] = 1
    }
    v10 := GenS1{F0: 128, F1: 65535}
    v8 -= 766
    for fi11 := 0; fi11 < 8; fi11++ {
        fmt.Println(v2, v2)
        v4 = ((!false) && v4)
        switch ("foo bar" + "\n\t") {
        case "hello":
            v12 := GenS2{F0: v8, F1: (v2 / ((v8) | 1)), F2: 65535}
            _ = v12
        case "a":
            var v13 float32 = float32(3.14159)
            _ = v13
        case "foo bar":
            var v14 int8 = 49
            _ = v14
        default:
            var v15 float64 = float64(v9)
            _ = v15
        }
        _ = fi11
    }
    v16 := GenS3{F0: -1, F1: ("\n\t" + ("hello" + ("hello" + "\n\t")))}
    fmt.Println(v4, v2)
    var v17 uint32 = 2147483647
    v9 = 1e10
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
