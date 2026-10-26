package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 uint64 }

func main() {
    v1 := make(map[uint32]int)
    var v2 uint8 = 116
    var v3 int = -315
    fmt.Println(v2, v3)
    var v4 float32 = (0.1 + (0.1 - (3.14159 + 1e10)))
    var v5 uint = 555
    v2 = v2
    for fi6 := 0; fi6 < 1; fi6++ {
        var v7 string = "foo bar"
        var v8 uint16 = 358
        var v9 int64 = 1
        _ = fi6
        _ = v7
        _ = v8
        _ = v9
    }
    var v10 uint16 = 2
    if (!true) {
        v3 -= (v3 >> (((^(v5))) & 63))
    } else {
        v10 = 0
        v11 := []int{v3, v3, (v3 >> ((v5) & 63))}
        _ = v11
    }
    _ = len(v1)
    if (!(false || false)) {
        v10 = 560
        v4 -= float32(float64(3.14159))
    }
    v4 -= float32(3.14159)
    v12 := []int{}
    var v13 float32 = 0.0
    fmt.Println(v2, v4, v2)
    var v14 uint8 = 2
    v15 := GenS1{F0: 331, F1: 0}
    fmt.Println(v14, v14, v10)
    v4 += -2.25
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
