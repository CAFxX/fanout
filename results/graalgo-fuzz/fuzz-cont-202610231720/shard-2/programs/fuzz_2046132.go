package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16 }
type GenS2 struct { F0 int32; F1 int16 }
func genFn0() uint8 {
    switch "foo bar" {
    case "a":
        var v6 uint16 = 42
        _ = v6
    case "":
        var v7 int32 = -2147483648
        _ = v7
    case "\n\t":
        for fi8 := 0; fi8 < 14; fi8++ {
            switch -198 {
            case 32767:
                if ((fi8 & (fi8 ^ (-(fi8)))) == fi8) {
                    var v9 uint8 = 128
                    _ = v9
                }
            case 94:
                var v10 int32 = 1
                _ = v10
            default:
                var v11 int = (fi8 >> ((200) & 63))
                _ = v11
            }
            _ = fi8
        }
    }
    var v12 uint32 = 119
    _ = v12
    return 208
}
func genFn1(p0 uint64) uint8 {
    v13 := []int{-122, -177, -153, 1}
    fmt.Println(p0)
    v13 = append(v13, 0)
    p0 += 1
    _ = p0
    _ = v13
    return 127
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 float32 = (0.1 / 1e10)
    var v3 uint32 = 0
    v4 := make(map[int]int)
    var v5 int32 = 2147483647
    fmt.Println(v5, v5, v2)
    var v14 uint32 = v3
    switch 0 {
    case 849:
        var v15 bool = false
        _ = v15
    case 209:
        v16 := GenS1{F0: (true && false), F1: 1}
        _ = v16
    case 2147483647:
        switch 127 {
        case -914:
            v3 += 2147483648
        case 32768:
            for fi17 := 0; fi17 < 12; fi17++ {
                var v18 float64 = 1e10
                for fi19 := 0; fi19 < 8; fi19++ {
                    v20 := GenS1{F0: ((2 >= 127) && true), F1: 42}
                    _ = fi19
                    _ = v20
                }
                _ = fi17
                _ = v18
            }
        default:
            fmt.Println(v2)
        }
    }
    v21 := GenS2{F0: v5, F1: -32768}
    v2 -= v2
    v22 := GenS1{F0: true, F1: -533}
    var v23 bool = (581 >= 1000)
    fmt.Println(v2, v2, v23)
    v5 = -743
    v1[v14] = false
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v14
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
