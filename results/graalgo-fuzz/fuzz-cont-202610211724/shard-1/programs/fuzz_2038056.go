package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint16 }
func genFn0(p0 int, p1 int) uint16 {
    v7 := []int{(p0 ^ (p0 / (((p1 / ((p1) | 1))) | 1)))}
    _ = len(v7)
    p1 -= (p1 + 894)
    for fi8 := 0; fi8 < 19; fi8++ {
        if false {
            p0 = (fi8 % ((p0) | 1))
            for fi9 := 0; fi9 < 7; fi9++ {
                if true {
                    p0 = (fi9 * -509)
                    fmt.Println(fi9)
                    var v10 uint = 42
                    _ = v10
                }
                _ = len(v7)
                _ = fi9
            }
            p0 = 9223372036854775807
        }
        _ = fi8
    }
    _ = p0
    _ = p1
    _ = v7
    return 974
}

func main() {
    var v1 uint = 18446744073709551615
    var v2 uint32 = 65535
    var v3 int = 128
    var v4 uint64 = 0
    var v5 bool = false
    var v6 float32 = ((3.14159 + 1.5) / (1.5 / 1.0))
    v1 = (^(v1))
    v1 -= v1
    v2 -= v2
    v11 := GenS1{F0: false, F1: 2}
    var v12 float32 = (3.14159 - (v6 * (v6 - ((v6 * (1.5 / 1e10)) / ((0.0 * v6) + v6)))))
    for fi13 := 0; fi13 < 4; fi13++ {
        fmt.Println(v4, v5, v1)
        fmt.Println(v2)
        v3 = (v3 | (fi13 - v3))
        _ = fi13
    }
    var v14 int8 = 22
    v15 := []int{}
    var v16 int64 = 907
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v15:", len(v15), v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
