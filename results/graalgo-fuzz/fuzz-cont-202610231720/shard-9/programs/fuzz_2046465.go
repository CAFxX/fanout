package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }

func main() {
    var v1 uint32 = 227
    var v2 int32 = -2147483648
    var v3 float64 = (-2.25 / float64(3.14159))
    v4 := make(map[int]int)
    v5 := make(map[int]int)
    v1 = v1
    _ = len(v4)
    delete(v4, 158)
    for fi6 := 0; fi6 < 19; fi6++ {
        switch (fi6 << ((1) & 63)) {
        case -396:
            v7 := GenS1{F0: ((^(fi6)) < (fi6 / ((fi6) | 1)))}
            _ = v7
        case 9223372036854775807:
            v8 := []int{(fi6 % ((fi6) | 1))}
            _ = v8
        case 597:
            v9 := []int{}
            _ = v9
        default:
            v3 -= v3
        }
        var v10 int64 = -20
        var v11 int16 = 533
        _ = fi6
        _ = v10
        _ = v11
    }
    for fi12 := 0; fi12 < 5; fi12++ {
        v13 := GenS1{F0: (984 == 768)}
        for fi14 := 0; fi14 < 12; fi14++ {
            var v15 uint8 = 2
            var v16 string = "hello"
            _ = fi14
            _ = v15
            _ = v16
        }
        _ = fi12
        _ = v13
    }
    var v17 uint8 = 2
    v3 += (1.5 * (v3 - -2.25))
    v18 := []int{0}
    v19 := GenS1{F0: false}
    var v20 float64 = v3
    v2 += v2
    if len(v18) > 0 {
        v18[0] = 42
    }
    fmt.Println("v1:", v1)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
