package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint8; F2 bool }
type GenS2 struct { F0 uint16; F1 uint }
func genFn0() int {
    var v6 uint = 0
    v6 = v6
    fmt.Println(v6)
    v7 := []int{32768, -602}
    _ = v6
    _ = v7
    return 127
}

func main() {
    var v1 uint8 = 255
    var v2 int8 = -34
    v3 := make(map[string]int)
    var v4 int16 = 1000
    var v5 int64 = -889
    fmt.Println(v5, v4, v4)
    v8 := []int{-231}
    var v9 float32 = 0.1
    v1 = (^(v1))
    v5 += 835
    v1 = (v1 >> ((255) & 63))
    v4 = v4
    switch "a" {
    case "a":
        for fi10 := 0; fi10 < 20; fi10++ {
            v3[("a" + "")] = 573
            for fi11 := 0; fi11 < 20; fi11++ {
                v5 = -1
                _ = fi11
            }
            v12 := GenS1{F0: (fi10 << ((1) & 63)), F1: v1, F2: ((float32(float64(v9)) + (v9 / (float32(((-2.25 / ((1.5 / -2.25) - (((0.1 + (1e10 - (-2.25 * ((1e10 * float64(0.0)) / 3.14159)))) + -2.25) - (0.0 / 1.0)))) * ((1e10 / ((1e10 - -2.25) + ((float64((0.1 * v9)) * float64(((1.5 * float32((((float64(v9) * 3.14159) * 0.0) - float64((float32((float64(v9) + 1e10)) + (v9 / float32(0.1))))))) + float32(-2.25)))) * 0.0))) * 0.1))) - 1.5))) == v9)}
            _ = fi10
            _ = v12
        }
    case "":
        for fi13 := 0; fi13 < 1; fi13++ {
            v14 := GenS2{F0: 932, F1: 969}
            var v15 uint32 = 2147483647
            _ = fi13
            _ = v14
            _ = v15
        }
    default:
        delete(v3, "a")
    }
    v9 = (v9 - v9)
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v8:", len(v8), v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
