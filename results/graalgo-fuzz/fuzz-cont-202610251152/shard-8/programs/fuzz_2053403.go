package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 uint8; F1 uint64 }
func genFn0(p0 uint32, p1 int16, p2 uint16) bool {
    if false {
        fmt.Println(p2, p2)
        p2 += 465
    }
    switch 2147483647 {
    case 1:
        fmt.Println(p1)
    case -1:
        for fi4 := 0; fi4 < 19; fi4++ {
            fmt.Println(p0, fi4)
            if false {
                p0 = 2
                v5 := []int{-631, 845, 32}
                v6 := GenS1{F0: uint16(p0)}
                _ = v5
                _ = v6
            }
            var v7 int64 = -1
            _ = fi4
            _ = v7
        }
    default:
        var v8 uint32 = (p0 << ((710) & 63))
        _ = v8
    }
    fmt.Println(p2, p2)
    v9 := GenS1{F0: p2}
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return true
}
func genFn1(p0 bool) uint {
    for fi10 := 0; fi10 < 19; fi10++ {
        p0 = p0
        fmt.Println(fi10, p0)
        _ = fi10
    }
    v11 := GenS1{F0: 1}
    _ = p0
    _ = v11
    return 255
}

func main() {
    var v1 int8 = -82
    var v2 float64 = 1.5
    var v3 uint8 = 0
    v12 := []int{}
    v13 := GenS2{F0: (v3 + v3), F1: 358}
    v2 -= (v2 - float64((3.14159 / 3.14159)))
    v1 = v1
    var v14 uint = 0
    var v15 uint32 = 2147483648
    v16 := GenS2{F0: 2, F1: uint64(v3)}
    var v17 uint = v14
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
}
