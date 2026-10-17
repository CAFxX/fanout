package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int8; F2 uint32 }
type GenS2 struct { F0 uint64; F1 uint; F2 int }
func genFn0(p0 int) int8 {
    p0 -= p0
    v6 := GenS1{F0: (true || (344 < 950)), F1: -128, F2: 65535}
    var v7 uint16 = 65535
    fmt.Println(v7, p0)
    if (!(-2.25 < (0.1 - 0.0))) {
        v7 = 32767
    } else {
        var v8 bool = true
        _ = v8
    }
    _ = p0
    _ = v6
    _ = v7
    return 1
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]string)
    var v3 int = -1
    var v4 int16 = 1000
    var v5 float64 = 0.0
    switch 1 {
    case 1:
        v4 -= v4
    case 32767:
        if (-1 >= 0) {
            v9 := []int{1}
            _ = v9
        }
    }
    var v10 int32 = 785
    var v11 bool = false
    v12 := GenS2{F0: 18446744073709551615, F1: 117, F2: 32767}
    var v13 int64 = -114
    var v14 float64 = float64(1e10)
    v5 -= v5
    v4 = (v4 * (v4 << ((1) & 63)))
    var v15 int = 597
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v10
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
