package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint8; F2 uint16 }
func genFn0(p0 uint, p1 uint16) int8 {
    p0 -= (p0 | p0)
    for fi5 := 0; fi5 < 9; fi5++ {
        v6 := []int{fi5, (fi5 % ((fi5) | 1))}
        v7 := GenS1{F0: 0, F1: 84, F2: (p1 << ((p0) & 63))}
        _ = fi5
        _ = v6
        _ = v7
    }
    switch 524 {
    case 4294967295:
        p0 += p0
    case -1:
        for fi8 := 0; fi8 < 15; fi8++ {
            for fi9 := 0; fi9 < 8; fi9++ {
                v10 := []int{fi8, (fi9 ^ fi8)}
                _ = fi9
                _ = v10
            }
            var v11 uint16 = (p1 ^ p1)
            for fi12 := 0; fi12 < 18; fi12++ {
                if false {
                    var v13 uint64 = 450
                    _ = v13
                } else {
                    p0 += (^(p0))
                }
                v14 := GenS1{F0: 65535, F1: 42, F2: v11}
                _ = fi12
                _ = v14
            }
            _ = fi8
            _ = v11
        }
    }
    var v15 uint8 = 127
    _ = p0
    _ = p1
    _ = v15
    return 127
}

func main() {
    var v1 uint = 0
    var v2 float32 = 3.14159
    var v3 int16 = 0
    var v4 float64 = 3.14159
    var v16 int16 = (v3 & v3)
    v4 = 3.14159
    var v17 uint32 = 65535
    v18 := []int{-245}
    var v19 uint64 = 0
    var v20 int8 = 1
    v19 -= (v19 - 1)
    v3 = (-(v16))
    fmt.Println(v1, v20, v4)
    var v21 uint32 = (^(v17))
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v21:", v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
}
