package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint8; F2 uint16 }
type GenS2 struct { F0 int32; F1 string; F2 string }
func genFn0(p0 string) int16 {
    if false {
        var v6 float64 = (((3.14159 / (0.1 / 1e10)) - 1.5) / 0.1)
        v7 := GenS1{F0: true, F1: 42, F2: 0}
        _ = v6
        _ = v7
    } else {
        for fi8 := 0; fi8 < 8; fi8++ {
            v9 := GenS2{F0: 329, F1: (p0 + (p0 + "foo bar")), F2: p0}
            var v10 int = (fi8 + (-(fi8)))
            _ = fi8
            _ = v10
            _ = v9
        }
        var v11 int32 = -796
        _ = v11
    }
    var v12 uint64 = 1
    _ = p0
    _ = v12
    return -32768
}

func main() {
    var v1 float64 = 1.5
    var v2 uint8 = 2
    var v3 uint8 = (v2 & (v2 ^ (v2 + 2)))
    var v4 int32 = -554
    var v5 uint = 42
    v5 = (v5 ^ (v5 << ((2) & 63)))
    var v13 string = "foo bar"
    v14 := []int{-1, 683, 254}
    var v15 int16 = genFn0(v13)
    fmt.Println(v1, v3, v1)
    switch -429 {
    case 1:
        fmt.Println(v1, v2, v15)
    default:
        var v16 float64 = v1
        _ = v16
    }
    v17 := GenS2{F0: 2147483647, F1: v13, F2: (v13 + v13)}
    v3 = v3
    for fi18 := 0; fi18 < 5; fi18++ {
        v13 = v13
        v14 = append(v14, fi18)
        _ = fi18
    }
    fmt.Println(v3)
    for fi19 := 0; fi19 < 11; fi19++ {
        if len(v14) >= 2 {
            v14 = v14[1:len(v14)]
        }
        v20 := GenS1{F0: true, F1: v3, F2: 12}
        _ = fi19
        _ = v20
    }
    fmt.Println(v13, v3, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
