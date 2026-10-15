package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16; F2 int16 }
type GenS2 struct { F0 uint32 }

func main() {
    v1 := make(map[int]bool)
    var v2 bool = (!true)
    var v3 int32 = 879
    v4 := make(map[string]int)
    v5 := make(map[string]uint32)
    v6 := []int{}
    switch 320 {
    case -436:
        v3 = (v3 << ((184) & 63))
    default:
        var v7 uint32 = 42
        _ = v7
    }
    v3 = (v3 & v3)
    for fi8 := 0; fi8 < 2; fi8++ {
        v3 = (v3 | 2147483647)
        _ = fi8
    }
    switch "a" {
    case "foo bar":
        _ = len(v4)
    case "\n\t":
        for fi9 := 0; fi9 < 15; fi9++ {
            var v10 string = "hello"
            _ = fi9
            _ = v10
        }
    }
    var v11 uint8 = 133
    for fi12 := 0; fi12 < 6; fi12++ {
        if (((42 > 42) && v2) || false) {
            if v2 {
                for fi13 := 0; fi13 < 4; fi13++ {
                    var v14 int32 = v3
                    v15 := GenS1{F0: -1, F1: 264, F2: 1}
                    _ = fi13
                    _ = v14
                    _ = v15
                }
                v16 := GenS1{F0: 32767, F1: 255, F2: 723}
                _ = v16
            }
            var v17 bool = true
            v18 := GenS1{F0: 42, F1: 255, F2: -32768}
            _ = v17
            _ = v18
        } else {
            switch "a" {
            case "a":
                v19 := GenS1{F0: -32768, F1: 65535, F2: 49}
                _ = v19
            default:
                v11 -= (v11 / ((v11) | 1))
            }
            v2 = (0.0 > ((float64(((((3.14159 - 3.14159) * ((1.5 - 0.1) + 1.5)) * float32((float64(0.1) + -2.25))) + ((((((float32(((3.14159 - ((-2.25 / ((3.14159 * -2.25) + (-2.25 / ((1.5 / 1.0) * 0.1)))) / 1.0)) * float64((((0.0 + 1.5) * (0.1 + 0.0)) + (float32((1e10 * float64(0.1))) / (1e10 / 0.1)))))) / ((-2.25 - float32(0.1)) * (((((float32(((float64(-2.25) - -2.25) / (3.14159 + float64(((1e10 + 1e10) - ((0.0 + (-2.25 * (0.0 + -2.25))) * 1e10)))))) * 3.14159) / (3.14159 * -2.25)) + -2.25) / 1e10) + (1.5 / -2.25)))) / 3.14159) + 3.14159) + (((0.0 - (((1e10 + -2.25) * ((0.1 - 1.5) / 3.14159)) - -2.25)) * float32(float64(float32(0.1)))) + 0.1)) / 1.5) - 0.1))) * 1.5) - 0.0))
        }
        fmt.Println(fi12, v11)
        _ = fi12
    }
    v3 = v3
    for fi20 := 0; fi20 < 14; fi20++ {
        for fi21 := 0; fi21 < 14; fi21++ {
            if v2 {
                if (1.5 != (1e10 / 3.14159)) {
                    v22 := GenS2{F0: 1}
                    _ = v22
                }
                var v23 bool = (true && ((fi21 / ((fi21) | 1)) > (fi20 ^ (fi21 | (fi21 % ((-2147483648) | 1))))))
                _ = v23
            }
            _ = fi21
        }
        var v24 uint32 = 2
        _ = fi20
        _ = v24
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
