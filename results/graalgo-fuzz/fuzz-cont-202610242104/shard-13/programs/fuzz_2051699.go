package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int; F2 uint32 }
func genFn0(p0 int, p1 int8, p2 uint8) uint {
    v5 := []int{(^(p0)), p0, (p0 | 42)}
    for fi6 := 0; fi6 < 5; fi6++ {
        p0 += (-(p0))
        if false {
            p1 += -17
            v7 := GenS1{F0: 1, F1: (fi6 << ((4294967295) & 63)), F2: 65535}
            _ = v7
        } else {
            var v8 int8 = (-(p1))
            _ = v8
        }
        _ = fi6
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v5
    return 255
}

func main() {
    v1 := []int{32767, 4294967295, -44, -528}
    v2 := []string{}
    var v3 int32 = -604
    v4 := make(map[int]bool)
    v9 := GenS1{F0: 491, F1: 9223372036854775807, F2: 166}
    var v10 int8 = -128
    var v11 uint64 = 4294967295
    v11 = v11
    var v12 int32 = (v3 + v3)
    for fi13 := 0; fi13 < 7; fi13++ {
        if (!true) {
            if (!(0 != v11)) {
                v11 = v11
                v10 += -116
            } else {
                var v14 int8 = -128
                switch (fi13 | (fi13 / ((128) | 1))) {
                case 963:
                    var v15 uint32 = 4294967295
                    _ = v15
                case 128:
                    var v16 bool = false
                    _ = v16
                case 9223372036854775807:
                    delete(v4, 2147483647)
                default:
                    if len(v1) >= 2 {
                        v1 = v1[1:len(v1)]
                    }
                }
                _ = v14
            }
            var v17 uint32 = 408
            var v18 int64 = int64(v3)
            _ = v17
            _ = v18
        }
        v3 -= (-(v12))
        _ = fi13
    }
    var v19 int64 = int64(v12)
    v19 = (v19 + 9223372036854775807)
    v11 -= 473
    v1 = append(v1, 42)
    var v20 uint16 = 255
    if (((1 <= 393) || false) || ((!false) && (204 > 131))) {
        v21 := GenS1{F0: 695, F1: 128, F2: 2}
        v22 := GenS1{F0: 18446744073709551615, F1: 32768, F2: 2147483648}
        _ = v21
        _ = v22
    } else {
        _ = len(v2)
        for fi23 := 0; fi23 < 20; fi23++ {
            fmt.Println(v20, v3, v12)
            _ = fi23
        }
        v3 = v12
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v19:", v19)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v9
}
