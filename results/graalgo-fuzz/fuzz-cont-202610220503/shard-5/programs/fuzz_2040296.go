package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 uint32 }
func genFn0(p0 uint64, p1 int16, p2 bool) uint64 {
    if p2 {
        p2 = (!((65535 == 42) || (("" + "\n\t") <= "hello")))
    }
    for fi4 := 0; fi4 < 1; fi4++ {
        v5 := GenS1{F0: (-(fi4))}
        var v6 uint64 = (p0 / (((p0 + p0)) | 1))
        fmt.Println(p0)
        _ = fi4
        _ = v5
        _ = v6
    }
    v7 := GenS2{F0: 1}
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return p0
}

func main() {
    v1 := make(map[uint32]bool)
    var v2 uint32 = 753
    var v3 bool = (false || ((0 != 330) && true))
    var v8 int8 = -1
    switch 2147483647 {
    case 32768:
        v3 = v3
    case 470:
        var v9 int32 = -407
        _ = v9
    case -859:
        v8 += v8
    }
    var v10 uint16 = 331
    fmt.Println(v8, v3)
    if ((genFn0(18446744073709551615, 42, (-324 >= 127)) > 350) || (((v2 != 42) && v3) && v3)) {
        fmt.Println(v3, v10)
        _ = len(v1)
    } else {
        var v11 uint = 4294967295
        _ = v11
    }
    fmt.Println(v10, v8)
    if (18446744073709551615 <= 42) {
        var v12 float32 = (0.1 - -2.25)
        _ = v12
    }
    v13 := []int{910}
    for fi14 := 0; fi14 < 14; fi14++ {
        v15 := GenS2{F0: (v2 >> ((65535) & 63))}
        var v16 float64 = float64(3.14159)
        _ = fi14
        _ = v15
        _ = v16
    }
    var v17 bool = v3
    fmt.Println(v2, v17, v10)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v10:", v10)
    fmt.Println("slicelen v13:", len(v13), v13)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v8:", v8)
    _ = v1
    _ = v10
    _ = v13
    _ = v17
    _ = v2
    _ = v3
    _ = v8
}
