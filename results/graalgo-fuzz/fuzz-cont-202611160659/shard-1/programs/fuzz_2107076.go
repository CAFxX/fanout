package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 uint16 }
type GenS2 struct { F0 int8; F1 int8; F2 int }
func genFn0(p0 uint32, p1 uint64) uint {
    p1 = (p1 | 4294967295)
    v6 := GenS1{F0: 1, F1: 865}
    var v7 int64 = -9223372036854775808
    p1 -= p1
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 65535
}

func main() {
    var v1 string = ""
    var v2 uint16 = 504
    var v3 float32 = float32(-2.25)
    var v4 int64 = 335
    var v5 uint64 = 42
    var v8 uint64 = (^(v5))
    v9 := GenS1{F0: (v2 / ((v2) | 1)), F1: (v2 - 575)}
    v5 -= v5
    if (false && (1000 == 32767)) {
        var v10 int8 = 127
        _ = v10
    } else {
        v11 := []int{42, 2147483647}
        var v12 int16 = 32767
        _ = v11
        _ = v12
    }
    v13 := GenS1{F0: 32767, F1: (v2 - 0)}
    v14 := GenS2{F0: 0, F1: 42, F2: 2147483647}
    v5 = v5
    var v15 uint = 310
    fmt.Println(v2, v4)
    v3 -= (v3 + ((v3 - v3) - ((v3 / ((v3 + (v3 / 1.0)) - v3)) + 1.5)))
    v15 -= 812
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
    _ = v9
}
