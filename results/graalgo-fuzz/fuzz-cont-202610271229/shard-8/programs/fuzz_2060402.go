package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint64 }
func genFn0(p0 int, p1 uint16, p2 uint64) int {
    var v5 uint64 = (p2 | (p2 ^ (p2 ^ p2)))
    var v6 uint64 = v5
    var v7 float32 = (1e10 * -2.25)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    _ = v7
    return (p0 & p0)
}
func genFn1(p0 int32, p1 uint16) int8 {
    var v8 uint8 = 54
    p1 = p1
    _ = p0
    _ = p1
    _ = v8
    return -128
}

func main() {
    v1 := []uint32{}
    var v2 float32 = 1e10
    var v3 int16 = 244
    var v4 uint16 = 42
    if true {
        switch -64 {
        case 575:
            if false {
                var v9 int8 = -1
                var v10 int64 = 0
                _ = v10
                _ = v9
            }
        }
    }
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    var v11 int16 = v3
    var v12 uint32 = 0
    var v13 int8 = genFn1(42, (v4 / ((v4) | 1)))
    for fi14 := 0; fi14 < 14; fi14++ {
        v15 := GenS1{F0: 42, F1: 2}
        for fi16 := 0; fi16 < 10; fi16++ {
            v4 -= 65535
            if len(v1) > 0 {
                v1[0] = v12
            }
            fmt.Println(v4)
            _ = fi16
        }
        _ = fi14
        _ = v15
    }
    var v17 bool = (v4 <= 346)
    var v18 bool = v17
    v17 = (!true)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
}
