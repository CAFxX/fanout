package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 uint }
type GenS2 struct { F0 int8; F1 uint32 }

func main() {
    v1 := []uint32{0}
    v2 := []int{-75, 2147483647, 613}
    var v3 uint16 = 0
    v4 := []uint32{}
    var v5 uint64 = 1
    v3 += 32767
    for fi6 := 0; fi6 < 13; fi6++ {
        for fi7 := 0; fi7 < 6; fi7++ {
            if false {
                fmt.Println(fi6, v5, fi6)
                if (0.1 < (0.0 - 0.1)) {
                    v8 := GenS1{F0: 255, F1: 1}
                    v5 += v5
                    _ = v8
                }
            } else {
                v3 = v3
            }
            fmt.Println(v5)
            _ = fi7
        }
        _ = fi6
    }
    var v9 int8 = 1
    v3 = v3
    for fi10 := 0; fi10 < 19; fi10++ {
        switch "foo bar" {
        case "\n\t":
            var v11 uint = 0
            _ = v11
        case "hello":
            fmt.Println(v9, v5, fi10)
        case "foo bar":
            var v12 float64 = (3.14159 * -2.25)
            _ = v12
        default:
            v1 = append(v1, 42)
        }
        _ = fi10
    }
    v5 = v5
    var v13 int32 = 2147483647
    if len(v4) >= 2 {
        v4 = v4[1:len(v4)]
    }
    var v14 uint16 = 42
    var v15 uint16 = uint16(v5)
    for fi16 := 0; fi16 < 13; fi16++ {
        fmt.Println(v14, v13)
        _ = fi16
    }
    for fi17 := 0; fi17 < 11; fi17++ {
        v13 -= -369
        var v18 int16 = -518
        _ = fi17
        _ = v18
    }
    v19 := GenS2{F0: (v9 | v9), F1: 4294967295}
    var v20 float64 = (float64(-2.25) / ((-2.25 / -2.25) + 0.1))
    v14 += (v15 + v14)
    fmt.Println(v15, v20, v20)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v9
}
