package main

import (
    "fmt"
)

type GenS1 struct { F0 string; F1 uint; F2 uint8 }
func genFn0(p0 string, p1 uint8, p2 uint) uint8 {
    var v7 uint16 = 42
    v7 -= (v7 * v7)
    if false {
        switch 445 {
        case -451:
            p0 = p0
        case 0:
            var v8 bool = false
            _ = v8
        default:
            v9 := GenS1{F0: "a", F1: p2, F2: 1}
            _ = v9
        }
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return (p1 / (((p1 + (p1 % (((p1 << ((p2) & 63))) | 1)))) | 1))
}

func main() {
    v1 := []uint32{2147483647, 2, 1}
    var v2 int8 = 100
    v3 := make(map[int]int)
    var v4 uint16 = 255
    var v5 uint8 = 2
    var v6 uint16 = v4
    for fi10 := 0; fi10 < 9; fi10++ {
        var v11 int = 412
        _ = fi10
        _ = v11
    }
    for fi12 := 0; fi12 < 4; fi12++ {
        if len(v1) > 0 {
            v1[0] = uint32(v5)
        }
        _ = fi12
    }
    var v13 uint64 = 0
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v13 -= 42
    v6 += (v6 * (v4 - 2))
    v1 = append(v1, 4294967295)
    var v14 uint32 = 2
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
