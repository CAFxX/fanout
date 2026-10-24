package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int; F2 int16 }
type GenS2 struct { F0 uint; F1 string }
type GenS3 struct { F0 int8 }

func main() {
    var v1 uint16 = 42
    var v2 uint32 = 0
    v3 := make(map[uint32]int)
    var v4 float32 = (float32(((0.1 / (0.1 - (0.1 * 1.5))) - float64((float32(float64((3.14159 / 1.0))) / (((((0.0 + 1e10) - (0.1 + 1e10)) - ((float32(-2.25) / 1.5) / 3.14159)) - float32(0.0)) - 3.14159))))) + -2.25)
    v4 = (float32(1.5) * (v4 + v4))
    var v5 string = (("a" + ("a" + (("hello" + "a") + ("\n\t" + ("hello" + ("" + "")))))) + "hello")
    var v6 int64 = -1
    var v7 int64 = 0
    var v8 float32 = v4
    v4 = (((v4 - v8) + v4) + -2.25)
    v9 := GenS1{F0: (((false && (!false)) || false) && false), F1: -787, F2: 0}
    v7 = 964
    v10 := GenS2{F0: 379, F1: v5}
    if true {
        v11 := []int{128, -409, 0}
        if len(v11) > 0 {
            v11[0] = -365
        }
        for fi12 := 0; fi12 < 19; fi12++ {
            for fi13 := 0; fi13 < 16; fi13++ {
                var v14 int16 = 1000
                v6 = (v6 >> ((1) & 63))
                _ = fi13
                _ = v14
            }
            var v15 int8 = 127
            _ = fi12
            _ = v15
        }
        _ = v11
    } else {
        var v16 float32 = v4
        _ = v16
    }
    if true {
        v6 += 1
        var v17 int16 = 1
        v17 = (-(v17))
        _ = v17
    }
    fmt.Println(v6)
    if (1000 <= 1) {
        v18 := GenS2{F0: 0, F1: ("a" + v5)}
        var v19 uint32 = 230
        _ = v18
        _ = v19
    }
    if (!false) {
        var v20 string = "hello"
        var v21 uint16 = (v1 << ((733) & 63))
        _ = v20
        _ = v21
    } else {
        v22 := GenS3{F0: 64}
        v2 += 42
        _ = v22
    }
    if true {
        fmt.Println(v7, v1, v6)
        fmt.Println(v2, v2, v5)
        var v23 int = -813
        _ = v23
    }
    var v24 uint32 = (v2 << ((65535) & 63))
    v4 -= -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v2
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
