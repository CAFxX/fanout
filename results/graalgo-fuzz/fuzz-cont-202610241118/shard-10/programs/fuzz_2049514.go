package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
func genFn0(p0 int, p1 int32, p2 int) uint16 {
    var v7 uint64 = 0
    var v8 int8 = 0
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    _ = v8
    return uint16(v7)
}
func genFn1() uint {
    if (true || false) {
        var v9 int16 = -32768
        if (164 > 42) {
            for fi10 := 0; fi10 < 13; fi10++ {
                v9 -= v9
                _ = fi10
            }
        }
        _ = v9
    }
    var v11 int = -9223372036854775808
    var v12 int8 = -128
    v11 += (v11 ^ v11)
    _ = v11
    _ = v12
    return 42
}

func main() {
    var v1 uint16 = 1
    var v2 int = 32768
    var v3 uint16 = v1
    var v4 int = (v2 / (((v2 << ((18446744073709551615) & 63))) | 1))
    var v5 uint = 469
    v6 := make(map[uint32]bool)
    var v13 string = ""
    v14 := []int{(v2 ^ 0)}
    v15 := GenS1{F0: false}
    v2 = v2
    v5 -= (v5 % ((242) | 1))
    var v16 int8 = -128
    v16 = (v16 / (((v16 ^ (v16 / ((100) | 1)))) | 1))
    fmt.Println(v4, v13)
    v4 -= 32767
    v1 += (^(v1))
    if true {
        v5 -= (v5 * genFn1())
        var v17 uint8 = 2
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
