package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 int16; F2 uint32 }

func main() {
    var v1 uint16 = 155
    v2 := make(map[uint32]int)
    var v3 int8 = -58
    var v4 float32 = float32((0.1 / (float64(0.0) - 0.1)))
    v2[915] = 128
    for fi5 := 0; fi5 < 9; fi5++ {
        var v6 int8 = (v3 ^ (v3 * 127))
        fmt.Println(v3, fi5, v1)
        _ = fi5
        _ = v6
    }
    switch ("foo bar" + "hello") {
    case "a":
        if true {
            for fi7 := 0; fi7 < 9; fi7++ {
                v3 = 42
                var v8 int = fi7
                _ = fi7
                _ = v8
            }
            v9 := []int{}
            _ = v9
        } else {
            var v10 int32 = 0
            _ = v10
        }
    default:
        var v11 string = "a"
        _ = v11
    }
    for fi12 := 0; fi12 < 14; fi12++ {
        var v13 float32 = ((v4 + (1e10 + 1.5)) / ((0.1 - float32(1e10)) * -2.25))
        v1 -= (v1 / ((32767) | 1))
        _ = fi12
        _ = v13
    }
    v4 += 0.0
    var v14 int8 = (v3 % ((0) | 1))
    v4 += v4
    var v15 int32 = -548
    if (!true) {
        fmt.Println(v1, v4)
        var v16 int16 = 346
        _ = v16
    }
    var v17 int = 4294967295
    v18 := GenS1{F0: v15, F1: 0, F2: 656}
    if true {
        v14 += v14
        v3 = 100
        v4 += (v4 - -2.25)
    } else {
        var v19 string = ("" + "\n\t")
        var v20 uint16 = (^(v1))
        _ = v19
        _ = v20
    }
    v4 -= float32(0.1)
    v4 = -2.25
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v14
    _ = v15
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
