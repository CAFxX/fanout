package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 int8; F2 uint32 }
type GenS2 struct { F0 int8 }
func genFn0(p0 uint64, p1 bool, p2 int) uint8 {
    p1 = (42 >= -43)
    var v4 int16 = 1000
    v5 := GenS1{F0: 0, F1: -61, F2: 4294967295}
    fmt.Println(p1)
    p2 = (p2 * 9223372036854775807)
    _ = p0
    _ = p1
    _ = p2
    _ = v4
    _ = v5
    return 1
}
func genFn1(p0 bool, p1 uint32) uint64 {
    v6 := GenS2{F0: -78}
    var v7 uint8 = 128
    p0 = p0
    _ = p0
    _ = p1
    _ = v6
    _ = v7
    return 18446744073709551615
}

func main() {
    var v1 float32 = (1.5 * -2.25)
    var v2 bool = true
    var v3 float64 = float64(0.1)
    for fi8 := 0; fi8 < 9; fi8++ {
        var v9 int32 = 2147483647
        _ = fi8
        _ = v9
    }
    var v10 int8 = -1
    v2 = ((2147483647 == -706) && (!false))
    v3 -= 0.1
    if (("" + "\n\t") <= "foo bar") {
        fmt.Println(v2)
        v3 -= float64(-2.25)
        var v11 int = 854
        _ = v11
    }
    var v12 int16 = 32767
    switch 879 {
    case 872:
        v13 := []int{32767, 383, 127, 915}
        _ = v13
    }
    v14 := []int{32768, 542, 0}
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v14:", len(v14), v14)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v10
    _ = v12
    _ = v14
    _ = v2
    _ = v3
}
