package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 bool; F2 uint32 }
type GenS2 struct { F0 uint64 }

func main() {
    var v1 uint64 = 361
    var v2 bool = false
    v3 := make(map[int]int)
    v4 := []uint32{481, 641}
    var v5 uint16 = 255
    if len(v4) > 0 {
        v4[0] = 518
    }
    if v2 {
        if (1.5 >= (1.5 + float32((0.0 + 3.14159)))) {
            v1 += v1
            var v6 int16 = -1
            _ = v6
        }
    } else {
        v5 = (v5 * 32767)
        for fi7 := 0; fi7 < 4; fi7++ {
            var v8 uint8 = 2
            for fi9 := 0; fi9 < 13; fi9++ {
                fmt.Println(v8)
                _ = fi9
            }
            _ = fi7
            _ = v8
        }
    }
    var v10 float32 = -2.25
    var v11 bool = (!(42 == (v1 / ((v1) | 1))))
    v12 := GenS1{F0: "foo bar", F1: (!(!false)), F2: 1}
    var v13 uint16 = v5
    v10 -= (float32((((-2.25 + float64(((v10 - float32(1.5)) - 1e10))) * -2.25) + 1e10)) - float32(3.14159))
    switch 0 {
    case -2147483648:
        v14 := GenS2{F0: v1}
        _ = v14
    case 32768:
        switch "\n\t" {
        case "foo bar":
            for fi15 := 0; fi15 < 7; fi15++ {
                var v16 float32 = 0.1
                _ = fi15
                _ = v16
            }
        }
    case -694:
        var v17 float64 = (0.0 * 1.5)
        _ = v17
    default:
        v1 += 2
    }
    var v18 bool = v11
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
