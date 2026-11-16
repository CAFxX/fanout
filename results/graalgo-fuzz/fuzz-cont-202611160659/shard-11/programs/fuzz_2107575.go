package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 bool }

func main() {
    var v1 uint32 = 332
    v2 := make(map[int]int)
    var v3 float32 = 1e10
    var v4 float32 = -2.25
    var v5 int = -2147483648
    var v6 float32 = (v3 + -2.25)
    for fi7 := 0; fi7 < 11; fi7++ {
        for fi8 := 0; fi8 < 8; fi8++ {
            v9 := GenS1{F0: 100, F1: true}
            v10 := v2[v5]
            if true {
                v6 -= v4
                v5 += v10
                fmt.Println(v6)
            } else {
                fmt.Println(v6)
                v2[42] = (fi8 / ((fi8) | 1))
                fmt.Println(fi8, fi8)
            }
            _ = fi8
            _ = v10
            _ = v9
        }
        fmt.Println(fi7)
        switch fi7 {
        case -2147483648:
            if (-731 >= 32767) {
                fmt.Println(v1)
            } else {
                var v11 uint32 = (v1 / ((2147483648) | 1))
                _ = v11
            }
        }
        _ = fi7
    }
    var v12 string = "foo bar"
    for fi13 := 0; fi13 < 17; fi13++ {
        var v14 int32 = -596
        _ = fi13
        _ = v14
    }
    v15 := v2[324]
    fmt.Println(v5, v3)
    var v16 string = (((v12 + v12) + v12) + "a")
    var v17 float32 = (((1.5 * 1.5) + v3) - float32(1.5))
    v18 := []int{v5, v5}
    if true {
        v16 = "\n\t"
    } else {
        v17 += 0.1
    }
    for fi19 := 0; fi19 < 11; fi19++ {
        v20 := GenS1{F0: -127, F1: false}
        _ = fi19
        _ = v20
    }
    v4 += 3.14159
    fmt.Println(v3, v1, v6)
    v16 = "foo bar"
    fmt.Println(v1)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
