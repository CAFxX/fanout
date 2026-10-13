package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 int32 }
func genFn0(p0 uint) int {
    for fi7 := 0; fi7 < 19; fi7++ {
        for fi8 := 0; fi8 < 17; fi8++ {
            var v9 uint32 = 707
            _ = fi8
            _ = v9
        }
        v10 := GenS1{F0: 1, F1: -2147483648}
        _ = fi7
        _ = v10
    }
    fmt.Println(p0)
    _ = p0
    return 42
}
func genFn1(p0 int64, p1 uint16, p2 uint) int64 {
    v11 := GenS1{F0: p0, F1: -2147483648}
    v12 := []int{2147483647}
    _ = p0
    _ = p1
    _ = p2
    _ = v11
    _ = v12
    return (-(p0))
}

func main() {
    v1 := make(map[uint32]uint32)
    var v2 bool = (0.1 >= 1e10)
    var v3 int32 = -2147483648
    v4 := []string{"foo bar", "", ("hello" + "hello"), ("foo bar" + "\n\t")}
    var v5 int16 = 1
    v6 := make(map[uint32]bool)
    v2 = v2
    _ = len(v4)
    var v13 uint32 = 65535
    v3 += v3
    v3 = (v3 % ((-1) | 1))
    var v14 int8 = -128
    v4 = append(v4, "a")
    v15 := v1[2147483647]
    var v16 uint64 = 18446744073709551615
    fmt.Println(v14, v3)
    if len(v4) > 0 {
        v4[0] = ("hello" + "a")
    }
    v14 -= (-(v14))
    var v17 int64 = -915
    v2 = (!(!(892 != 526)))
    fmt.Println(v15, v2, v14)
    v3 = 42
    fmt.Println(v16, v2, v5)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
