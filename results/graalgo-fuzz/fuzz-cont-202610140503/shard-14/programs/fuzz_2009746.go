package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
func genFn0(p0 uint8, p1 int8, p2 int64) int8 {
    switch 801 {
    case 32768:
        var v4 int32 = -2147483648
        _ = v4
    default:
        p2 = (-(p2))
    }
    var v5 uint8 = p0
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return (p1 | p1)
}

func main() {
    v1 := []uint32{926, 0, 42, 1}
    var v2 float32 = ((0.1 / (3.14159 + 0.1)) + (3.14159 * (0.1 + (0.0 + 3.14159))))
    v3 := []int{}
    var v6 uint16 = 963
    if false {
        switch "foo bar" {
        case "":
            v7 := GenS1{F0: 56}
            _ = v7
        case "\n\t":
            var v8 int16 = 0
            _ = v8
        }
    }
    if ((56 >= 424) && false) {
        v2 += v2
        v6 = 42
        v2 += (v2 + (float32(-2.25) - 0.1))
    }
    switch 4294967295 {
    case 32768:
        switch 0 {
        case 32767:
            v1 = append(v1, 683)
        }
    case 127:
        var v9 int = 32767
        _ = v9
    }
    fmt.Println(v6, v2)
    fmt.Println(v6, v6)
    v10 := GenS1{F0: 32768}
    v3 = append(v3, -1)
    for fi11 := 0; fi11 < 18; fi11++ {
        for fi12 := 0; fi12 < 8; fi12++ {
            var v13 int64 = 42
            _ = fi12
            _ = v13
        }
        var v14 float64 = (float64(v2) + 1.5)
        _ = fi11
        _ = v14
    }
    for fi15 := 0; fi15 < 15; fi15++ {
        if len(v1) >= 2 {
            v1 = v1[1:len(v1)]
        }
        var v16 float32 = v2
        _ = fi15
        _ = v16
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v2
    _ = v3
    _ = v6
}
