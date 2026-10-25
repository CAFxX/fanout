package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16 }
type GenS2 struct { F0 uint; F1 int8; F2 int16 }
func genFn0(p0 uint64, p1 uint) uint64 {
    var v5 uint64 = 0
    for fi6 := 0; fi6 < 6; fi6++ {
        for fi7 := 0; fi7 < 5; fi7++ {
            switch (fi6 & (fi7 % (((fi7 + (fi7 + fi7))) | 1))) {
            case 952:
                v8 := []int{}
                _ = v8
            case 237:
                p0 = (^(p0))
            default:
                v9 := GenS1{F0: 1, F1: 758}
                _ = v9
            }
            fmt.Println(v5)
            p1 -= (p1 * 255)
            _ = fi7
        }
        _ = fi6
    }
    v10 := GenS1{F0: 9223372036854775807, F1: 42}
    p1 = p1
    _ = p0
    _ = p1
    _ = v10
    _ = v5
    return (p0 * 1)
}
func genFn1(p0 int16, p1 bool, p2 int64) uint16 {
    p2 += p2
    v11 := GenS2{F0: 2, F1: 42, F2: -293}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    return 42
}

func main() {
    var v1 uint8 = 189
    v2 := make(map[int]string)
    var v3 string = ""
    v4 := make(map[string]int)
    switch 127 {
    case 4294967295:
        for fi12 := 0; fi12 < 15; fi12++ {
            var v13 string = v3
            _ = fi12
            _ = v13
        }
    case 128:
        for fi14 := 0; fi14 < 12; fi14++ {
            v3 = ((((v3 + "\n\t") + v3) + v3) + v3)
            _ = fi14
        }
    }
    fmt.Println(v1)
    v3 = "a"
    v1 += v1
    if false {
        v1 += v1
    } else {
        var v15 float32 = float32(((3.14159 + 3.14159) / 1.0))
        v1 = (v1 + v1)
        _ = v15
    }
    delete(v4, v3)
    var v16 float64 = -2.25
    fmt.Println(v1, v3)
    var v17 bool = (false && (false || (-9223372036854775808 == 42)))
    v18 := GenS1{F0: -9223372036854775808, F1: 710}
    fmt.Println(v17, v17)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
