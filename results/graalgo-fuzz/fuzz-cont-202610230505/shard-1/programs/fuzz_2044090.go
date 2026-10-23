package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int; F2 uint16 }
type GenS2 struct { F0 uint64 }
func genFn0(p0 int64, p1 int) uint64 {
    fmt.Println(p0, p1)
    switch (p1 << ((18446744073709551615) & 63)) {
    case -1:
        if (!true) {
            v6 := []int{p1, p1, (p1 | (-(p1))), 522}
            _ = len(v6)
            if ((3.14159 != (1.5 / 1.0)) || ((3.14159 >= 3.14159) && (255 > 255))) {
                var v7 int = (p1 / ((42) | 1))
                var v8 uint8 = 0
                var v9 float64 = 0.1
                _ = v7
                _ = v8
                _ = v9
            }
            _ = v6
        }
    case 127:
        var v10 string = (("\n\t" + (("foo bar" + ("foo bar" + "hello")) + ("\n\t" + ""))) + "")
        _ = v10
    default:
        var v11 int32 = 0
        _ = v11
    }
    var v12 int8 = -128
    _ = p0
    _ = p1
    _ = v12
    return 4294967295
}

func main() {
    v1 := make(map[string]string)
    var v2 int = 574
    v3 := []int{v2}
    var v4 uint16 = 32767
    var v5 float32 = 0.1
    var v13 int32 = -1
    v13 -= v13
    v4 += v4
    v13 += (v13 >> ((1) & 63))
    for fi14 := 0; fi14 < 6; fi14++ {
        v13 -= (v13 & (v13 >> ((1) & 63)))
        delete(v1, "a")
        v2 = (-(fi14))
        _ = fi14
    }
    var v15 uint32 = 2147483647
    fmt.Println(v2, v2, v2)
    v13 = v13
    var v16 float32 = (v5 + v5)
    v17 := GenS1{F0: 42, F1: (v2 * (v2 ^ v2)), F2: v4}
    v18 := GenS1{F0: 485, F1: v2, F2: (v4 - v4)}
    v19 := GenS2{F0: 2}
    var v20 uint32 = 951
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v13
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
