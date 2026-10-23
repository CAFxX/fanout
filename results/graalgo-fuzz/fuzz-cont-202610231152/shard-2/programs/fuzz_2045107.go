package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 int8; F1 int16 }
func genFn0(p0 bool, p1 int16) uint64 {
    var v5 string = "a"
    var v6 uint32 = 722
    v7 := []int{-1, 398, 2147483647}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 18446744073709551615
}

func main() {
    v1 := make(map[int]string)
    v2 := make(map[uint32]uint32)
    v3 := make(map[string]bool)
    var v4 uint64 = 42
    v4 += v4
    delete(v1, 42)
    for fi8 := 0; fi8 < 5; fi8++ {
        v4 += v4
        var v9 uint32 = 2147483648
        _ = fi8
        _ = v9
    }
    for fi10 := 0; fi10 < 11; fi10++ {
        var v11 float64 = -2.25
        _ = fi10
        _ = v11
    }
    switch 91 {
    case -523:
        for fi12 := 0; fi12 < 4; fi12++ {
            v4 -= 611
            v4 -= (v4 ^ (v4 >> ((42) & 63)))
            var v13 float64 = 3.14159
            _ = fi12
            _ = v13
        }
    case -2147483648:
        v4 -= v4
    case -49:
        var v14 uint8 = 2
        _ = v14
    default:
        switch 782 {
        case 127:
            v15 := GenS1{F0: 255}
            _ = v15
        }
    }
    if false {
        fmt.Println(v4)
        var v16 uint32 = 42
        _ = v16
    }
    v17 := GenS2{F0: 61, F1: -1}
    fmt.Println(v4)
    v4 += genFn0(true, 1)
    if false {
        fmt.Println(v4)
    }
    for fi18 := 0; fi18 < 2; fi18++ {
        var v19 bool = ((-2.25 - (1e10 - 0.1)) != -2.25)
        _ = fi18
        _ = v19
    }
    fmt.Println(v4)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
