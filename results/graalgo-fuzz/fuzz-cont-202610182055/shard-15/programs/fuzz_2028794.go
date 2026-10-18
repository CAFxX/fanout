package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool; F2 int32 }
func genFn0(p0 string) int32 {
    v6 := []int{}
    for fi7 := 0; fi7 < 17; fi7++ {
        fmt.Println(fi7, fi7)
        for fi8 := 0; fi8 < 13; fi8++ {
            var v9 int8 = 127
            if true {
                var v10 uint = 4294967295
                _ = v10
            }
            _ = len(v6)
            _ = fi8
            _ = v9
        }
        if len(v6) >= 2 {
            v6 = v6[1:len(v6)]
        }
        _ = fi7
    }
    fmt.Println(p0)
    _ = len(v6)
    _ = p0
    _ = v6
    return 1
}

func main() {
    var v1 int32 = -664
    var v2 int32 = (^(v1))
    var v3 float32 = (0.0 * (((0.0 * -2.25) / 1.0) + -2.25))
    var v4 int = 652
    v5 := []uint32{4294967295, 2, 2147483648}
    v4 = (v4 % (((v4 - (v4 >> ((65535) & 63)))) | 1))
    v2 += (v2 + v2)
    v11 := GenS1{F0: 1, F1: (2 >= 140), F2: 0}
    v4 = (v4 % (((v4 << ((562) & 63))) | 1))
    fmt.Println(v2, v3, v4)
    var v12 int16 = 547
    var v13 uint8 = 128
    v3 += (1.5 - (v3 - (float32(0.1) - (((-2.25 - 1e10) / v3) * -2.25))))
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
