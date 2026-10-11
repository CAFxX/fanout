package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint32 }
type GenS2 struct { F0 int32; F1 string }
func genFn0(p0 int16, p1 uint16) int64 {
    var v6 int8 = -9
    var v7 float64 = 1e10
    for fi8 := 0; fi8 < 14; fi8++ {
        for fi9 := 0; fi9 < 1; fi9++ {
            v7 += v7
            v10 := []int{}
            _ = fi9
            _ = v10
        }
        _ = fi8
    }
    p0 = (p0 & int16(p1))
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 64
}

func main() {
    var v1 float64 = (0.0 / 1.0)
    var v2 int8 = 0
    var v3 int8 = (v2 * (v2 % (((v2 ^ (^(v2)))) | 1)))
    var v4 int8 = 1
    var v5 int16 = 42
    fmt.Println(v1, v2)
    for fi11 := 0; fi11 < 3; fi11++ {
        var v12 uint = 65535
        var v13 int = (-(fi11))
        v14 := GenS1{F0: 2, F1: 42}
        _ = fi11
        _ = v12
        _ = v13
        _ = v14
    }
    var v15 uint32 = 0
    fmt.Println(v3, v2, v1)
    v16 := []int{682, -2147483648}
    v16 = append(v16, 32767)
    if len(v16) >= 2 {
        v16 = v16[1:len(v16)]
    }
    var v17 bool = ((2 < 523) && false)
    fmt.Println(v3, v5)
    if len(v16) > 0 {
        v16[0] = 32767
    }
    if true {
        v5 -= v5
    } else {
        v18 := GenS1{F0: v15, F1: (v15 >> ((18446744073709551615) & 63))}
        _ = v18
    }
    v19 := GenS2{F0: 2147483647, F1: ("hello" + ("a" + "\n\t"))}
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v16
    _ = v17
    _ = v19
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
