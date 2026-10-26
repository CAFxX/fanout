package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint32 }
type GenS2 struct { F0 string }
func genFn0(p0 bool) uint64 {
    v7 := GenS1{F0: (!p0), F1: 668}
    var v8 int8 = 100
    v9 := GenS1{F0: p0, F1: 2147483648}
    v10 := GenS2{F0: ("hello" + "\n\t")}
    _ = p0
    _ = v10
    _ = v7
    _ = v8
    _ = v9
    return 4294967295
}

func main() {
    var v1 uint = 18446744073709551615
    v2 := []uint32{}
    var v3 uint16 = 1
    v4 := []int{4294967295, 9223372036854775807, 517, -9223372036854775808}
    v5 := []string{}
    var v6 uint64 = 2
    var v11 bool = (4294967295 > 631)
    v11 = (1 >= 42)
    for fi12 := 0; fi12 < 15; fi12++ {
        var v13 string = "\n\t"
        if (376 >= 1) {
            for fi14 := 0; fi14 < 13; fi14++ {
                var v15 string = (v13 + v13)
                _ = fi14
                _ = v15
            }
        }
        _ = fi12
        _ = v13
    }
    fmt.Println(v3)
    _ = len(v4)
    if false {
        var v16 float64 = ((1e10 + (0.0 * (1.5 / 3.14159))) * (0.0 / 1e10))
        _ = v16
    } else {
        var v17 int8 = -1
        var v18 uint64 = (v6 % (((v6 << ((uint(v3)) & 63))) | 1))
        for fi19 := 0; fi19 < 12; fi19++ {
            v6 = (v6 ^ (v18 ^ 4294967295))
            var v20 int32 = -2147483648
            for fi21 := 0; fi21 < 6; fi21++ {
                v11 = (v3 > v3)
                var v22 float32 = -2.25
                _ = fi21
                _ = v22
            }
            _ = fi19
            _ = v20
        }
        _ = v17
        _ = v18
    }
    v11 = ((!v11) || v11)
    var v23 int64 = 419
    var v24 int32 = 157
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v23:", v23)
    fmt.Println("v24:", v24)
    fmt.Println("v3:", v3)
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v2
    _ = v23
    _ = v24
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
