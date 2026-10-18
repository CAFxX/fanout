package main

import (
    "fmt"
)

type GenS1 struct { F0 int16; F1 uint64 }
func genFn0(p0 string) uint64 {
    var v6 uint64 = 126
    v7 := []int{0, 0, -822}
    v8 := GenS1{F0: 1000, F1: (v6 % ((v6) | 1))}
    var v9 uint16 = 838
    _ = p0
    _ = v6
    _ = v7
    _ = v8
    _ = v9
    return (v6 ^ v6)
}
func genFn1(p0 int, p1 int32, p2 int8) bool {
    var v10 int8 = (p2 - p2)
    fmt.Println(p2, p0)
    var v11 int64 = -9223372036854775808
    fmt.Println(v11)
    v10 = int8(v11)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return (-9223372036854775808 != v11)
}

func main() {
    var v1 bool = (!false)
    v2 := make(map[uint32]string)
    v3 := make(map[int]uint32)
    var v4 uint16 = 65535
    v5 := make(map[uint32]uint32)
    var v12 float64 = ((1e10 + 0.1) - ((0.0 / float64(0.1)) + 3.14159))
    var v13 int64 = 9223372036854775807
    v14 := v5[2]
    v14 = (v14 & (v14 - v14))
    var v15 int32 = 1
    for fi16 := 0; fi16 < 14; fi16++ {
        if true {
            if true {
                v4 -= (v4 >> ((255) & 63))
                delete(v2, 2)
            } else {
                fmt.Println(v14, v13, v12)
                v15 = (v15 + -2147483648)
                var v17 uint8 = 42
                _ = v17
            }
            v13 += v13
        }
        _ = fi16
    }
    v18 := []int{}
    _ = len(v5)
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v18:", len(v18), v18)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
