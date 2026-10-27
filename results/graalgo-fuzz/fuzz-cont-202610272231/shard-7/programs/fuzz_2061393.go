package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16 }
type GenS2 struct { F0 bool }
func genFn0(p0 int64, p1 bool, p2 int64) uint16 {
    p1 = p1
    fmt.Println(p2)
    p1 = (-9223372036854775808 < -163)
    v6 := GenS1{F0: 1}
    var v7 int8 = 122
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 255
}
func genFn1() uint8 {
    v8 := GenS2{F0: (18446744073709551615 != 4294967295)}
    var v9 uint8 = 1
    _ = v8
    _ = v9
    return v9
}

func main() {
    var v1 int64 = 1
    v2 := make(map[int]int)
    v3 := make(map[int]bool)
    var v4 float64 = 0.1
    v5 := make(map[uint32]bool)
    for fi10 := 0; fi10 < 7; fi10++ {
        fmt.Println(fi10)
        fmt.Println(fi10)
        _ = fi10
    }
    var v11 float64 = float64((1e10 / -2.25))
    v12 := GenS1{F0: 65535}
    delete(v2, 32767)
    v4 += ((v11 / (1.5 - v4)) + ((3.14159 - 0.0) / v11))
    v13 := v3[257]
    v11 = v11
    var v14 int64 = (^(v1))
    v15 := GenS1{F0: 354}
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
