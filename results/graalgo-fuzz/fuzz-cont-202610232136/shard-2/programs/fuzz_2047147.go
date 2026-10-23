package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint16 }
type GenS2 struct { F0 uint16; F1 int64 }
func genFn0(p0 uint8) bool {
    p0 += (p0 / ((p0) | 1))
    if true {
        for fi7 := 0; fi7 < 10; fi7++ {
            v8 := GenS1{F0: -32768, F1: 144}
            _ = fi7
            _ = v8
        }
        var v9 bool = true
        switch 0 {
        case 145:
            var v10 float64 = (float64(((1.5 * ((-2.25 - (1e10 - 1e10)) * 1.5)) / (0.0 - -2.25))) * 1e10)
            _ = v10
        default:
            for fi11 := 0; fi11 < 7; fi11++ {
                fmt.Println(fi11, p0)
                _ = fi11
            }
        }
        _ = v9
    } else {
        v12 := []int{127, 567, 128, 9223372036854775807}
        var v13 int16 = -32768
        var v14 int = 42
        _ = v12
        _ = v13
        _ = v14
    }
    _ = p0
    return (434 == uint32(p0))
}

func main() {
    v1 := make(map[int]uint32)
    v2 := make(map[string]int)
    var v3 int = -1
    var v4 string = "foo bar"
    var v5 uint8 = 144
    v6 := make(map[uint32]uint32)
    var v15 bool = (9223372036854775807 < -800)
    v15 = (v15 && true)
    var v16 uint32 = 753
    var v17 string = "hello"
    v18 := []int{}
    var v19 float32 = (0.0 * 3.14159)
    v20 := GenS1{F0: 0, F1: 65535}
    var v21 bool = v15
    switch "\n\t" {
    case "foo bar":
        v22 := GenS2{F0: 90, F1: 0}
        _ = v22
    default:
        v3 -= 1
    }
    var v23 uint16 = 32767
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
