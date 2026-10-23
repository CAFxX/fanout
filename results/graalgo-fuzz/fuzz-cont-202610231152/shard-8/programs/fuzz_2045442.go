package main

import (
    "fmt"
)

type GenS1 struct { F0 int32 }
type GenS2 struct { F0 int; F1 int; F2 uint32 }
func genFn0(p0 string, p1 bool) uint16 {
    var v4 uint = 65535
    for fi5 := 0; fi5 < 5; fi5++ {
        v6 := []int{}
        var v7 uint16 = 255
        var v8 int = fi5
        _ = fi5
        _ = v6
        _ = v7
        _ = v8
    }
    v4 = 358
    _ = p0
    _ = p1
    _ = v4
    return 905
}
func genFn1() uint32 {
    var v9 int64 = -9223372036854775808
    v10 := GenS1{F0: 42}
    v11 := GenS2{F0: -460, F1: -1, F2: 633}
    _ = v10
    _ = v11
    _ = v9
    return 939
}

func main() {
    var v1 bool = (!false)
    var v2 float32 = (1e10 / -2.25)
    var v3 uint16 = 339
    var v12 string = "hello"
    var v13 uint16 = (v3 | v3)
    fmt.Println(v1, v2)
    v14 := GenS2{F0: 2147483647, F1: 1, F2: uint32(v3)}
    switch -1 {
    case 128:
        var v15 uint64 = 42
        _ = v15
    case 42:
        v2 = (((0.1 + 1.5) - float32(1.5)) - float32(1.5))
    }
    fmt.Println(v13)
    switch "foo bar" {
    case "hello":
        v12 = "\n\t"
    default:
        switch "hello" {
        case "a":
            if false {
                if (v2 > (0.0 + ((0.1 - float32((0.1 - 0.0))) + v2))) {
                    var v16 int8 = 100
                    var v17 uint8 = 42
                    _ = v16
                    _ = v17
                }
            }
        default:
            fmt.Println(v2, v1)
        }
    }
    v2 += float32((float64(v2) + ((((1e10 * 3.14159) / -2.25) * float64(v2)) - (3.14159 + 0.1))))
    v18 := []int{70, 286}
    switch 127 {
    case 318:
        v19 := GenS2{F0: -9223372036854775808, F1: 898, F2: 184}
        _ = v19
    }
    v1 = (2 >= 224)
    v13 += (v13 / ((genFn0(v12, (v12 == v12))) | 1))
    v2 -= (1e10 * v2)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v18
    _ = v2
    _ = v3
}
