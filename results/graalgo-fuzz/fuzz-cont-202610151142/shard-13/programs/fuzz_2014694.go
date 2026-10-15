package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 string; F2 string }
func genFn0() bool {
    var v5 uint = 794
    v6 := GenS1{F0: 57, F1: ("" + ("" + ("hello" + "\n\t"))), F2: ("" + ("foo bar" + ""))}
    _ = v5
    _ = v6
    return true
}
func genFn1(p0 uint32) uint64 {
    var v7 int = 0
    p0 -= p0
    _ = p0
    _ = v7
    return 4294967295
}

func main() {
    var v1 uint32 = 0
    var v2 int32 = -1
    var v3 uint = 746
    var v4 uint = (v3 - (v3 >> ((v3) & 63)))
    var v8 int32 = v2
    var v9 uint = (v4 << ((v4) & 63))
    var v10 float64 = 0.1
    v1 -= (v1 + v1)
    var v11 float64 = (((v10 * v10) * (((v10 - v10) - v10) / float64((1e10 + float32(v10))))) / v10)
    fmt.Println(v4, v9, v8)
    v10 -= 1.5
    v9 += (^(v9))
    var v12 uint = (v4 % ((v4) | 1))
    v4 += v12
    var v13 uint16 = 65535
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v8
    _ = v9
}
