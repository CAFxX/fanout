package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int64 }
func genFn0(p0 uint64, p1 int64, p2 uint64) uint64 {
    for fi5 := 0; fi5 < 9; fi5++ {
        if true {
            for fi6 := 0; fi6 < 6; fi6++ {
                v7 := []int{(fi6 % ((-59) | 1)), (fi5 << ((255) & 63))}
                var v8 uint = 4294967295
                p2 -= 18446744073709551615
                _ = fi6
                _ = v7
                _ = v8
            }
            if (-900 >= -2147483648) {
                p1 -= (p1 % ((p1) | 1))
            }
            var v9 uint8 = 102
            _ = v9
        } else {
            fmt.Println(p0)
        }
        _ = fi5
    }
    fmt.Println(p1, p0, p2)
    v10 := []int{2147483647}
    v11 := GenS1{F0: 4294967295, F1: (p1 >> ((2) & 63))}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    return 549
}

func main() {
    v1 := []uint32{907, 42}
    v2 := make(map[int]uint32)
    v3 := []uint32{}
    var v4 uint32 = 0
    v4 += (v4 / ((4294967295) | 1))
    var v12 uint64 = 0
    var v13 bool = (true && false)
    var v14 uint32 = (v4 - 2)
    fmt.Println(v4)
    for fi15 := 0; fi15 < 6; fi15++ {
        switch "a" {
        case "foo bar":
            fmt.Println(v14, v13, v12)
        case "a":
            if ((!v13) && ("hello" >= "")) {
                fmt.Println(v13)
            } else {
                v1 = append(v1, 4294967295)
                v13 = ((!(v13 || true)) && ((256 < 31) || false))
                v4 += v14
            }
        }
        v14 = v14
        _ = fi15
    }
    var v16 int32 = -737
    fmt.Println(v12)
    if len(v3) > 0 {
        v3[0] = v4
    }
    v17 := GenS1{F0: (v12 | (v12 + v12)), F1: -1}
    v18 := GenS1{F0: v12, F1: 879}
    var v19 int32 = v16
    v13 = v13
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v16:", v16)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v3
    _ = v4
}
