package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int64; F2 int16 }
type GenS2 struct { F0 string; F1 int16; F2 int }
func genFn0(p0 uint16, p1 int32, p2 uint32) uint8 {
    p0 = p0
    var v7 float64 = 3.14159
    fmt.Println(p1, v7, p0)
    if (p0 <= (p0 | (p0 + (p0 ^ p0)))) {
        if (-128 >= 42) {
            if false {
                for fi8 := 0; fi8 < 20; fi8++ {
                    var v9 uint16 = (p0 ^ 32767)
                    switch (fi8 % ((42) | 1)) {
                    case -144:
                        var v10 uint16 = v9
                        _ = v10
                    case -9223372036854775808:
                        p1 += (p1 * p1)
                    }
                    p2 = p2
                    _ = fi8
                    _ = v9
                }
            }
        }
        var v11 string = ""
        v12 := GenS1{F0: 898, F1: 9223372036854775807, F2: 1000}
        _ = v11
        _ = v12
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v7
    return 255
}

func main() {
    var v1 uint32 = 2
    v2 := []int{42, 4294967295, -9223372036854775808}
    var v3 int32 = 42
    var v4 int64 = 42
    v5 := []int{284, 0}
    var v6 int = 42
    var v13 int = v6
    var v14 string = ""
    v2 = append(v2, v6)
    for fi15 := 0; fi15 < 3; fi15++ {
        v14 = "hello"
        for fi16 := 0; fi16 < 12; fi16++ {
            var v17 bool = false
            if (false || v17) {
                var v18 uint16 = 255
                v19 := GenS2{F0: v14, F1: 42, F2: 128}
                _ = v18
                _ = v19
            }
            _ = fi16
            _ = v17
        }
        v1 = v1
        _ = fi15
    }
    _ = len(v5)
    if len(v5) > 0 {
        v5[0] = (-(v13))
    }
    var v20 uint8 = 205
    if (!false) {
        if len(v5) > 0 {
            v5[0] = (v13 | (v13 | (v6 % (((v13 | (v6 & v6))) | 1))))
        }
        v13 -= (v13 ^ v13)
    }
    v3 = (v3 ^ -2147483648)
    v21 := GenS2{F0: "a", F1: -32768, F2: v6}
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v13
    _ = v14
    _ = v2
    _ = v20
    _ = v21
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
