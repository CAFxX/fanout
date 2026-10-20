package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 int16; F2 uint }
func genFn0(p0 uint64) uint {
    for fi4 := 0; fi4 < 4; fi4++ {
        for fi5 := 0; fi5 < 11; fi5++ {
            switch "foo bar" {
            case "foo bar":
                var v6 int32 = 627
                _ = v6
            }
            v7 := []int{fi5, (fi4 / ((-424) | 1)), fi4, fi5}
            _ = fi5
            _ = v7
        }
        _ = fi4
    }
    var v8 int32 = -1
    v9 := GenS1{F0: 128, F1: 1000, F2: 18446744073709551615}
    p0 += 561
    var v10 float32 = 1e10
    _ = p0
    _ = v10
    _ = v8
    _ = v9
    return 749
}
func genFn1(p0 int32) int16 {
    fmt.Println(p0)
    v11 := []int{4294967295, -9223372036854775808, 496, -1}
    p0 = p0
    _ = len(v11)
    _ = p0
    _ = v11
    return 924
}

func main() {
    v1 := make(map[string]int)
    var v2 uint = 12
    var v3 uint = (v2 & (v2 ^ (v2 >> ((18446744073709551615) & 63))))
    v12 := []int{32767}
    if (-9223372036854775808 >= 1) {
        v13 := GenS1{F0: 1, F1: genFn1(-521), F2: genFn0(4294967295)}
        _ = v13
    }
    v14 := GenS1{F0: 128, F1: 1000, F2: v2}
    var v15 int8 = 42
    v3 -= 233
    if (((("a" + "") + (("foo bar" + "foo bar") + "")) != "foo bar") && ((324 == 4294967295) && (true || (18 > 42)))) {
        v16 := GenS1{F0: 1, F1: 32767, F2: 18446744073709551615}
        _ = len(v1)
        _ = v16
    } else {
        var v17 float32 = ((3.14159 * 0.0) + float32(-2.25))
        _ = len(v1)
        v17 = 0.1
        _ = v17
    }
    v15 -= v15
    for fi18 := 0; fi18 < 15; fi18++ {
        var v19 int8 = v15
        v19 = (^(v19))
        _ = fi18
        _ = v19
    }
    if len(v12) > 0 {
        v12[0] = 907
    }
    v2 -= 515
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v12
    _ = v14
    _ = v15
    _ = v2
    _ = v3
}
