package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 uint16 }
type GenS2 struct { F0 int16; F1 uint32; F2 bool }
func genFn0(p0 uint16, p1 int64, p2 uint16) int {
    p2 = 382
    for fi7 := 0; fi7 < 13; fi7++ {
        v8 := GenS1{F0: true, F1: 1000, F2: (p2 / (((^(p0))) | 1))}
        var v9 bool = true
        v10 := []int{32767, (fi7 << ((65535) & 63)), (fi7 << ((1) & 63)), (^(fi7))}
        _ = fi7
        _ = v10
        _ = v8
        _ = v9
    }
    _ = p0
    _ = p1
    _ = p2
    return 32768
}

func main() {
    v1 := make(map[uint32]string)
    var v2 float32 = (3.14159 / 1e10)
    var v3 int16 = 300
    var v4 int32 = -2147483648
    var v5 float64 = (((float64(float32(0.1)) * -2.25) * float64(1.5)) - (3.14159 * 1.5))
    var v6 float32 = 0.1
    fmt.Println(v4, v6, v3)
    var v11 float32 = float32(0.0)
    var v12 int8 = 0
    v1[2] = "\n\t"
    v13 := v1[2]
    v14 := GenS2{F0: 462, F1: 535, F2: (v12 > -88)}
    fmt.Println(v3, v3)
    v15 := GenS2{F0: -800, F1: 130, F2: true}
    delete(v1, 0)
    v16 := v1[1]
    if (-838 == int64(v3)) {
        switch 0 {
        case -2147483648:
            var v17 int8 = (v12 ^ (^(v12)))
            _ = v17
        }
        if ((^(v4)) < (^(v4))) {
            _ = len(v1)
        } else {
            fmt.Println(v6)
            var v18 int16 = 550
            _ = v18
        }
        var v19 int16 = (-(v3))
        _ = v19
    } else {
        delete(v1, 211)
    }
    if true {
        v20 := []int{286, 127, 461}
        v13 = ""
        _ = v20
    }
    var v21 int8 = v12
    v16 = (v16 + v16)
    var v22 uint16 = 291
    var v23 uint64 = 1
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
