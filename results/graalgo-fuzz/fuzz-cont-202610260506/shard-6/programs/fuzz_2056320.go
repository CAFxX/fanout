package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
type GenS2 struct { F0 uint64 }
type GenS3 struct { F0 int32; F1 uint }

func main() {
    v1 := make(map[string]int)
    var v2 uint32 = 4294967295
    var v3 uint8 = 58
    var v4 uint = 74
    v5 := make(map[int]string)
    var v6 float32 = -2.25
    switch 128 {
    case 42:
        if (32767 <= 78) {
            var v7 int32 = 1
            fmt.Println(v7)
            _ = v7
        }
    default:
        var v8 uint64 = 581
        _ = v8
    }
    var v9 int = 42
    v10 := GenS1{F0: 32767}
    v11 := GenS2{F0: 15}
    delete(v5, 1)
    v12 := GenS3{F0: -1, F1: 970}
    for fi13 := 0; fi13 < 15; fi13++ {
        v6 = (v6 / (v6 * v6))
        _ = fi13
    }
    v2 += v2
    if ((0.0 + -2.25) >= (float32(((float64(-2.25) / 0.1) - 1.5)) * (3.14159 / v6))) {
        var v14 uint = 65535
        switch ("\n\t" + ("hello" + "")) {
        case "":
            v4 = (^(v14))
        default:
            var v15 uint32 = v2
            _ = v15
        }
        var v16 uint64 = uint64(v3)
        _ = v14
        _ = v16
    }
    v17 := GenS1{F0: 1}
    fmt.Println(v9, v2)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
