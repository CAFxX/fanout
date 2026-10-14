package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0(p0 string, p1 int64, p2 int8) int16 {
    var v5 int8 = p2
    var v6 uint16 = 2
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    _ = v6
    return 42
}
func genFn1(p0 uint16, p1 uint16, p2 uint64) uint8 {
    var v7 uint8 = 94
    var v8 uint8 = (v7 - (v7 | (v7 << ((42) & 63))))
    for fi9 := 0; fi9 < 4; fi9++ {
        v10 := []int{(fi9 + 9223372036854775807), (^(fi9)), -1}
        _ = fi9
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return (^(v8))
}

func main() {
    var v1 float32 = -2.25
    var v2 float32 = v1
    var v3 string = "a"
    var v4 uint32 = 872
    switch 2147483647 {
    case 128:
        v4 = (v4 / ((1) | 1))
    case -562:
        v3 = ("a" + v3)
    case 0:
        v2 += v1
    }
    var v11 uint16 = 1
    v12 := []int{}
    if len(v12) > 0 {
        v12[0] = 318
    }
    var v13 int16 = 1000
    if (0 <= (v4 % ((v4) | 1))) {
        for fi14 := 0; fi14 < 3; fi14++ {
            v12 = append(v12, fi14)
            v15 := GenS1{F0: false}
            _ = fi14
            _ = v15
        }
        v13 += (v13 ^ (v13 >> ((953) & 63)))
    }
    if len(v12) > 0 {
        v12[0] = 32767
    }
    if len(v12) >= 2 {
        v12 = v12[1:len(v12)]
    }
    v16 := GenS1{F0: ((v3 < v3) && true)}
    v2 = 0.1
    if (-1 > 100) {
        var v17 float32 = v2
        if len(v12) > 0 {
            v12[0] = -2147483648
        }
        fmt.Println(v17, v2)
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v16
    _ = v2
    _ = v3
    _ = v4
}
