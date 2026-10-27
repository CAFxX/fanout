package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 int16 }
type GenS3 struct { F0 string }
func genFn0(p0 int64) int {
    fmt.Println(p0)
    v4 := []int{42, 128, -1, 1}
    _ = p0
    _ = v4
    return -1
}
func genFn1(p0 bool, p1 uint64, p2 uint8) uint64 {
    v5 := []int{513, -822, 128, 0}
    fmt.Println(p1, p1, p0)
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return (p1 % (((p1 % ((p1) | 1))) | 1))
}

func main() {
    var v1 int16 = 603
    v2 := []uint32{2, 408}
    var v3 int32 = 2147483647
    v6 := GenS1{F0: 32768}
    v7 := GenS2{F0: (v1 * (^(v1)))}
    fmt.Println(v1, v1)
    fmt.Println(v1, v3)
    _ = len(v2)
    var v8 uint16 = 1
    var v9 string = ""
    var v10 int16 = 32767
    fmt.Println(v3, v3)
    for fi11 := 0; fi11 < 13; fi11++ {
        var v12 bool = (-182 <= (v10 | 1000))
        v3 += v3
        v13 := GenS3{F0: ((v9 + v9) + v9)}
        _ = fi11
        _ = v12
        _ = v13
    }
    v3 = (v3 * (-(v3)))
    var v14 int32 = (v3 << ((18446744073709551615) & 63))
    v9 = v9
    v14 += (v14 >> ((740) & 63))
    var v15 float64 = 0.1
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
