package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 int; F2 int64 }
func genFn0(p0 int, p1 string, p2 uint) int64 {
    p1 = ("a" + p1)
    for fi4 := 0; fi4 < 12; fi4++ {
        if false {
            v5 := []int{p0, p0}
            _ = v5
        } else {
            v6 := []int{-823, int(p2), 561}
            for fi7 := 0; fi7 < 13; fi7++ {
                if len(v6) >= 2 {
                    v6 = v6[1:len(v6)]
                }
                _ = fi7
            }
            var v8 uint8 = 42
            _ = v6
            _ = v8
        }
        _ = fi4
    }
    _ = p0
    _ = p1
    _ = p2
    return -1
}
func genFn1(p0 int8, p1 int8) uint8 {
    fmt.Println(p0)
    v9 := []int{-577, 2147483647, 417, 128}
    var v10 int64 = -1
    _ = p0
    _ = p1
    _ = v10
    _ = v9
    return 127
}

func main() {
    var v1 int = 127
    v2 := []int{}
    var v3 float64 = 1.5
    v3 -= v3
    for fi11 := 0; fi11 < 2; fi11++ {
        var v12 int = (v1 + v1)
        _ = fi11
        _ = v12
    }
    if len(v2) >= 2 {
        v2 = v2[1:len(v2)]
    }
    v3 += v3
    v13 := GenS1{F0: -1, F1: (v1 % (((v1 / (((-(v1))) | 1))) | 1)), F2: 778}
    v1 -= 0
    var v14 string = "a"
    v14 = (v14 + (v14 + v14))
    v2 = append(v2, (^(v1)))
    v1 = v1
    v3 = 0.1
    v1 -= (v1 - 350)
    var v15 int16 = 42
    v1 = (-(v1))
    v16 := GenS1{F0: (v15 << ((2) & 63)), F1: (v1 | (v1 ^ 32767)), F2: -244}
    var v17 int = v1
    var v18 int8 = -128
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v2
    _ = v3
}
