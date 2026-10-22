package main

import (
    "fmt"
)

type GenS1 struct { F0 uint }

func main() {
    var v1 uint32 = 65535
    v2 := make(map[uint32]int)
    v3 := []int{}
    var v4 int8 = -52
    _ = len(v2)
    for fi5 := 0; fi5 < 14; fi5++ {
        for fi6 := 0; fi6 < 13; fi6++ {
            delete(v2, v1)
            _ = fi6
        }
        _ = fi5
    }
    for fi7 := 0; fi7 < 17; fi7++ {
        fmt.Println(v1, v1, fi7)
        var v8 uint32 = 42
        _ = fi7
        _ = v8
    }
    switch 720 {
    case 4294967295:
        v1 -= 2147483647
    case 1:
        v9 := GenS1{F0: 413}
        _ = v9
    default:
        v10 := GenS1{F0: 42}
        _ = v10
    }
    v3 = append(v3, -419)
    v11 := GenS1{F0: 255}
    v4 = (^(v4))
    v4 = v4
    v4 += (-(v4))
    v1 = v1
    var v12 float64 = (((((1e10 * 1e10) / 0.1) + float64(3.14159)) - -2.25) - float64((float32((-2.25 + (float64(1e10) - 0.1))) - 3.14159)))
    v1 = (v1 | (v1 * 0))
    delete(v2, v1)
    var v13 bool = (!true)
    for fi14 := 0; fi14 < 4; fi14++ {
        if v13 {
            var v15 float32 = 0.1
            v1 = (^(v1))
            var v16 float64 = v12
            _ = v15
            _ = v16
        } else {
            fmt.Println(v12, v1)
        }
        _ = fi14
    }
    if (v1 >= v1) {
        var v17 string = ""
        _ = v17
    } else {
        v18 := GenS1{F0: 773}
        for fi19 := 0; fi19 < 12; fi19++ {
            v12 = v12
            _ = fi19
        }
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
}
