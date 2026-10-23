package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint32; F2 uint32 }
type GenS2 struct { F0 int16 }
func genFn0(p0 int, p1 bool, p2 bool) int8 {
    if false {
        v5 := GenS1{F0: 127, F1: 2147483647, F2: 42}
        _ = v5
    } else {
        p1 = p1
        v6 := []int{p0, p0, 219, (p0 / ((407) | 1))}
        _ = v6
    }
    for fi7 := 0; fi7 < 12; fi7++ {
        v8 := GenS2{F0: 1}
        for fi9 := 0; fi9 < 5; fi9++ {
            var v10 uint = 42
            _ = fi9
            _ = v10
        }
        _ = fi7
        _ = v8
    }
    switch "" {
    case "a":
        var v11 int64 = 459
        _ = v11
    }
    v12 := []int{p0, (p0 << ((255) & 63))}
    _ = p0
    _ = p1
    _ = p2
    _ = v12
    return 127
}

func main() {
    var v1 uint8 = 2
    var v2 int = 4294967295
    var v3 int64 = 42
    var v4 uint8 = v1
    fmt.Println(v3)
    v13 := []int{(-(v2)), v2, v2, v2}
    v14 := GenS2{F0: -362}
    v4 += v4
    v2 += v2
    _ = len(v13)
    for fi15 := 0; fi15 < 17; fi15++ {
        fmt.Println(v2)
        _ = len(v13)
        fmt.Println(v4)
        _ = fi15
    }
    var v16 int32 = -1
    _ = len(v13)
    var v17 uint16 = 42
    if len(v13) >= 2 {
        v13 = v13[1:len(v13)]
    }
    v16 = v16
    var v18 int32 = (-(v16))
    fmt.Println(v17, v2)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v13
    _ = v14
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
