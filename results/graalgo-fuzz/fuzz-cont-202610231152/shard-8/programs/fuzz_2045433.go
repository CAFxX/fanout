package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 uint32; F1 int16; F2 uint }

func main() {
    var v1 uint16 = 32767
    var v2 float32 = float32(1e10)
    v3 := make(map[int]string)
    v1 += v1
    if (!((false || true) || (!true))) {
        var v4 uint = 868
        for fi5 := 0; fi5 < 4; fi5++ {
            fmt.Println(v2, v2)
            fmt.Println(v1)
            _ = fi5
        }
        v1 = uint16(v4)
        _ = v4
    } else {
        fmt.Println(v1)
    }
    for fi6 := 0; fi6 < 3; fi6++ {
        v7 := []int{}
        var v8 float64 = (0.0 + 0.1)
        fmt.Println(v8)
        _ = fi6
        _ = v7
        _ = v8
    }
    _ = len(v3)
    var v9 int = 471
    switch ("foo bar" + "hello") {
    case "a":
        if (247 >= 1) {
            var v10 bool = false
            v11 := GenS1{F0: v1, F1: v1}
            var v12 int32 = 64
            _ = v10
            _ = v11
            _ = v12
        } else {
            delete(v3, v9)
        }
    }
    v3[v9] = (("" + "a") + "hello")
    for fi13 := 0; fi13 < 18; fi13++ {
        v14 := GenS2{F0: 0, F1: 1, F2: 0}
        v2 = (0.1 - ((1e10 - float32((1.5 + 0.1))) + v2))
        _ = fi13
        _ = v14
    }
    fmt.Println(v2, v2)
    v15 := GenS1{F0: v1, F1: (^(v1))}
    v9 += v9
    for fi16 := 0; fi16 < 16; fi16++ {
        v9 = 42
        _ = fi16
    }
    fmt.Println("v1:", v1)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v9:", v9)
    _ = v1
    _ = v15
    _ = v2
    _ = v3
    _ = v9
}
