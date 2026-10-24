package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint16 }
type GenS2 struct { F0 uint; F1 int; F2 bool }

func main() {
    var v1 int8 = -1
    v2 := make(map[uint32]bool)
    v3 := make(map[uint32]string)
    var v4 int8 = -128
    v5 := make(map[string]string)
    var v6 float32 = (float32(float64(0.0)) * float32(-2.25))
    v1 = 0
    v1 = 42
    if (((v6 - ((0.0 + ((float32((1.5 / (((1.5 * 0.1) + float64(v6)) + (1.5 - float64(((-2.25 * (v6 - v6)) / v6)))))) + -2.25) / float32(1.5))) - -2.25)) + (v6 - (3.14159 / v6))) >= (float32(1.5) - v6)) {
        v7 := GenS1{F0: 1, F1: 42}
        v8 := GenS1{F0: 0, F1: 42}
        _ = v7
        _ = v8
    }
    var v9 int32 = 1
    fmt.Println(v4, v4, v1)
    if (!(("\n\t" >= "a") || (!true))) {
        var v10 int8 = (v4 ^ v4)
        switch 127 {
        case 127:
            v6 = v6
        case 32768:
            if true {
                var v11 int16 = 42
                _ = v11
            } else {
                v12 := []int{9223372036854775807, -2147483648}
                v13 := GenS2{F0: 591, F1: 2147483647, F2: true}
                _ = v12
                _ = v13
            }
        case 42:
            v14 := GenS1{F0: 106, F1: 795}
            _ = v14
        default:
            v15 := GenS2{F0: 255, F1: 579, F2: (!((0.0 >= 0.1) || true))}
            _ = v15
        }
        _ = v10
    }
    var v16 uint32 = 799
    v1 = v1
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
