package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int8 }
type GenS2 struct { F0 uint16 }
func genFn0(p0 int64, p1 bool) uint {
    var v7 float64 = 1.5
    p0 -= -553
    v8 := []int{32767, 0, 1, -125}
    v8 = append(v8, 370)
    v8 = append(v8, 0)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 255
}
func genFn1(p0 int32, p1 string, p2 int) uint16 {
    p1 = p1
    p0 = (p0 + (p0 * 1))
    v9 := []int{(p2 >> ((356) & 63)), p2, (p2 >> ((28) & 63)), (-(p2))}
    if ((!(255 > 2)) && (32767 >= 0)) {
        p2 -= (p2 & 1)
    } else {
        v10 := GenS1{F0: 127, F1: 18}
        _ = len(v9)
        _ = v10
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v9
    return 65535
}

func main() {
    v1 := make(map[int]bool)
    var v2 int64 = 282
    var v3 uint8 = 2
    var v4 int64 = v2
    var v5 int16 = 1
    var v6 bool = (!false)
    v6 = v6
    v3 = v3
    v1[-219] = ((!(genFn1(2147483647, "hello", -643) != 1)) || v6)
    var v11 bool = ((1e10 / 0.1) != ((3.14159 * (0.1 * (1e10 * (1e10 / 1.0)))) + 1e10))
    var v12 uint16 = 0
    for fi13 := 0; fi13 < 14; fi13++ {
        v5 += 1
        v14 := GenS2{F0: v12}
        var v15 uint32 = 65535
        _ = fi13
        _ = v14
        _ = v15
    }
    v4 -= v4
    fmt.Println(v12, v5)
    var v16 float32 = (0.0 + ((((float32(-2.25) - 3.14159) * 1.5) * 0.1) - 1.5))
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
