package main

import (
    "fmt"
)

type GenS1 struct { F0 string }
type GenS2 struct { F0 uint; F1 string }
func genFn0(p0 string, p1 string, p2 uint8) uint64 {
    var v7 float64 = -2.25
    var v8 int = 2147483647
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return 2
}

func main() {
    var v1 uint32 = 620
    var v2 int16 = 286
    var v3 uint = 124
    var v4 float64 = 3.14159
    var v5 uint32 = (v1 | uint32(v3))
    v6 := make(map[int]bool)
    v9 := []int{32768, -945, 0}
    for fi10 := 0; fi10 < 13; fi10++ {
        var v11 int = (fi10 ^ 32767)
        v5 = uint32(v3)
        _ = fi10
        _ = v11
    }
    v1 -= (v1 - (v5 | v5))
    v1 = (^(v1))
    if true {
        v12 := GenS1{F0: ("hello" + "\n\t")}
        for fi13 := 0; fi13 < 11; fi13++ {
            v5 += 331
            for fi14 := 0; fi14 < 7; fi14++ {
                v6[fi14] = false
                v1 += (v1 % ((v5) | 1))
                _ = fi14
            }
            fmt.Println(v1, fi13)
            _ = fi13
        }
        _ = v12
    }
    fmt.Println(v1, v4)
    if len(v9) >= 2 {
        v9 = v9[1:len(v9)]
    }
    for fi15 := 0; fi15 < 17; fi15++ {
        var v16 uint32 = (v1 | (v5 ^ v5))
        fmt.Println(v4)
        _ = fi15
        _ = v16
    }
    v3 += v3
    if (2147483647 > v5) {
        v17 := GenS2{F0: (v3 | (v3 & (v3 + v3))), F1: "foo bar"}
        v3 = (^(v3))
        _ = v17
    }
    if len(v9) > 0 {
        v9[0] = 127
    }
    v6[127] = false
    var v18 uint = (^(v3))
    for fi19 := 0; fi19 < 11; fi19++ {
        v20 := GenS1{F0: ""}
        for fi21 := 0; fi21 < 19; fi21++ {
            v4 = v4
            _ = fi21
        }
        _ = fi19
        _ = v20
    }
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    fmt.Println("slicelen v9:", len(v9), v9)
    _ = v1
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
