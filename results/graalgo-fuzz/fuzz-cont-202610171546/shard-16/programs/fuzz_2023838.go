package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 int16; F1 uint16 }
func genFn0(p0 int) int16 {
    var v7 int16 = 0
    v8 := []int{(p0 % (((p0 - p0)) | 1))}
    v7 -= (v7 % (((-(v7))) | 1))
    v7 = (v7 - 0)
    _ = p0
    _ = v7
    _ = v8
    return -32768
}
func genFn1(p0 int8, p1 bool, p2 int32) int8 {
    fmt.Println(p0, p0, p1)
    v9 := []int{4294967295, 0, 127}
    var v10 float32 = (((((-2.25 / 3.14159) / -2.25) + (0.0 - ((1.5 - 0.1) - 1.5))) - 0.1) * ((0.0 * ((-2.25 / (float32(1.5) / 0.1)) - 3.14159)) / 1.0))
    var v11 uint64 = 625
    for fi12 := 0; fi12 < 19; fi12++ {
        v10 = 0.0
        fmt.Println(p1)
        _ = fi12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v9
    return 100
}

func main() {
    var v1 uint8 = 128
    var v2 uint64 = 0
    var v3 int32 = 385
    v4 := make(map[uint32]string)
    v5 := make(map[int]uint32)
    var v6 uint16 = 623
    v6 = (v6 % (((v6 - (v6 >> ((65535) & 63)))) | 1))
    var v13 int16 = 42
    v14 := GenS1{F0: 993}
    v15 := []int{2147483647, 42, 0}
    if len(v15) > 0 {
        v15[0] = 0
    }
    v16 := GenS2{F0: v13, F1: (^(v6))}
    var v17 uint8 = 2
    switch ("" + "hello") {
    case "":
        var v18 int = 127
        _ = v18
    }
    if (-2147483648 > -784) {
        v3 += -451
        fmt.Println(v6)
        for fi19 := 0; fi19 < 2; fi19++ {
            var v20 int = (-(fi19))
            var v21 uint64 = 728
            v22 := GenS2{F0: (v13 - v13), F1: (v6 >> ((255) & 63))}
            _ = fi19
            _ = v20
            _ = v21
            _ = v22
        }
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
