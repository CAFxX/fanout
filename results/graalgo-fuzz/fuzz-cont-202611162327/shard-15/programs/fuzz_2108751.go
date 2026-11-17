package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint }
type GenS2 struct { F0 uint16; F1 int16 }
type GenS3 struct { F0 int }
func genFn0(p0 uint32, p1 uint16) int16 {
    for fi5 := 0; fi5 < 5; fi5++ {
        v6 := GenS1{F0: 499, F1: 257}
        _ = fi5
        _ = v6
    }
    if true {
        var v7 int16 = -61
        v8 := GenS2{F0: p1, F1: (v7 ^ (-(v7)))}
        var v9 uint64 = 489
        _ = v7
        _ = v8
        _ = v9
    } else {
        var v10 float32 = float32((3.14159 + float64(float32((1.5 + ((0.1 + -2.25) * float64(((((1e10 + (1e10 / 0.1)) - 0.1) / 1.0) * 1e10))))))))
        _ = v10
    }
    p1 = (^(p1))
    v11 := GenS3{F0: 128}
    var v12 uint8 = 127
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    return 1000
}

func main() {
    var v1 float32 = (1e10 + (3.14159 / (0.0 - 1.5)))
    var v2 uint = 65535
    var v3 int8 = 127
    v4 := make(map[uint32]uint32)
    fmt.Println(v1)
    fmt.Println(v2, v2, v1)
    v1 = (v1 / (v1 - v1))
    if (!true) {
        v3 -= 80
    } else {
        v2 = 2
    }
    v13 := GenS1{F0: -524, F1: (v2 * 65535)}
    delete(v4, 168)
    v2 += (v2 & v2)
    var v14 uint = 4294967295
    switch 42 {
    case 4294967295:
        var v15 string = ""
        _ = v15
    }
    fmt.Println(v3, v3)
    fmt.Println(v14)
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
