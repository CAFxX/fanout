package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16; F2 int16 }
type GenS2 struct { F0 uint32 }
func genFn0() bool {
    var v6 int32 = -2147483648
    v7 := GenS1{F0: (-2147483648 > 176), F1: 32767, F2: -787}
    v6 += (v6 + (v6 & v6))
    var v8 float32 = float32(1.5)
    for fi9 := 0; fi9 < 6; fi9++ {
        v10 := []int{-9223372036854775808, -970, 2147483647}
        if len(v10) >= 2 {
            v10 = v10[1:len(v10)]
        }
        _ = fi9
        _ = v10
    }
    _ = v6
    _ = v7
    _ = v8
    return ((v6 + (v6 * (v6 >> ((255) & 63)))) != v6)
}

func main() {
    var v1 float32 = 1e10
    var v2 uint64 = 957
    v3 := make(map[string]int)
    v4 := make(map[int]bool)
    var v5 uint64 = v2
    v11 := GenS1{F0: (-1 != -1), F1: 32767, F2: -406}
    var v12 uint8 = 59
    var v13 uint64 = 2
    for fi14 := 0; fi14 < 13; fi14++ {
        var v15 uint64 = (^(v5))
        switch "a" {
        case "":
            if false {
                v12 += v12
                v16 := []int{-9223372036854775808}
                _ = v16
            }
        }
        _ = fi14
        _ = v15
    }
    switch -854 {
    case 579:
        v1 += ((v1 - (v1 * (v1 * -2.25))) / 1e10)
    case 128:
        v17 := GenS2{F0: 2}
        _ = v17
    case 502:
        v12 = (^(v12))
    default:
        fmt.Println(v2, v1, v1)
    }
    v18 := GenS1{F0: true, F1: 0, F2: 42}
    for fi19 := 0; fi19 < 10; fi19++ {
        v13 = (v2 ^ 42)
        v20 := []int{(fi19 / ((-1) | 1))}
        _ = fi19
        _ = v20
    }
    for fi21 := 0; fi21 < 6; fi21++ {
        v4[(fi21 + 2147483647)] = ((^(v2)) > (v2 - (v13 << ((42) & 63))))
        if true {
            v5 -= v13
            delete(v4, fi21)
        } else {
            fmt.Println(fi21)
        }
        delete(v4, -850)
        _ = fi21
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
