package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 float64 = (float64(0.0) - -2.25)
    v2 := make(map[uint32]uint32)
    v3 := make(map[uint32]int)
    if (!false) {
        switch -740 {
        case 127:
            var v4 int = -900
            _ = v4
        case 1:
            var v5 uint32 = 1
            _ = v5
        default:
            var v6 uint = 0
            _ = v6
        }
        var v7 int8 = 127
        _ = v7
    }
    v1 = (float64((-2.25 / (1e10 - (1.5 - float32(-2.25))))) - 3.14159)
    v1 = (v1 - float64((3.14159 - 1.5)))
    var v8 uint = 1
    v9 := GenS1{F0: 42}
    var v10 string = "\n\t"
    v10 = v10
    if true {
        for fi11 := 0; fi11 < 18; fi11++ {
            _ = len(v2)
            _ = fi11
        }
        for fi12 := 0; fi12 < 9; fi12++ {
            v13 := v2[65535]
            var v14 uint16 = 0
            v8 = 567
            _ = fi12
            _ = v13
            _ = v14
        }
        var v15 uint = (^(v8))
        _ = v15
    } else {
        v2[2] = 2
        var v16 uint = v8
        v10 = v10
        _ = v16
    }
    v1 = v1
    var v17 float32 = 0.1
    v1 = (float64(v17) * v1)
    v18 := []int{1, 2147483647, 267, -716}
    switch -420 {
    case 4294967295:
        fmt.Println(v10)
    }
    var v19 int16 = -1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v8
    _ = v9
}
