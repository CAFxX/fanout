package main

import (
    "fmt"
)

type GenS1 struct { F0 int64 }
type GenS2 struct { F0 uint8; F1 uint32; F2 bool }

func main() {
    var v1 float32 = float32((0.0 / (1e10 / (-2.25 * 0.1))))
    var v2 uint16 = 255
    v3 := make(map[int]uint32)
    var v4 string = ("hello" + "hello")
    var v5 int32 = 106
    v6 := make(map[uint32]uint32)
    for fi7 := 0; fi7 < 20; fi7++ {
        fmt.Println(v2)
        _ = fi7
    }
    v8 := GenS1{F0: 1}
    var v9 uint16 = (v2 / ((v2) | 1))
    var v10 string = ((((("" + (v4 + ((v4 + (v4 + "hello")) + "hello"))) + "a") + v4) + "hello") + v4)
    if (-364 < -601) {
        fmt.Println(v5, v2)
    } else {
        for fi11 := 0; fi11 < 12; fi11++ {
            fmt.Println(v1, v1)
            switch -41 {
            case 4294967295:
                var v12 bool = (1 <= 1)
                _ = v12
            case -699:
                var v13 bool = true
                _ = v13
            default:
                v4 = ("\n\t" + (v10 + v10))
            }
            _ = fi11
        }
    }
    var v14 bool = true
    v9 = (v2 >> ((65535) & 63))
    v2 = (^(v2))
    v15 := GenS2{F0: 127, F1: 718, F2: (!(v14 || ((true || v14) || false)))}
    for fi16 := 0; fi16 < 3; fi16++ {
        var v17 float32 = v1
        var v18 uint32 = 861
        _ = fi16
        _ = v17
        _ = v18
    }
    v14 = (-2.25 > (0.1 + 1e10))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
