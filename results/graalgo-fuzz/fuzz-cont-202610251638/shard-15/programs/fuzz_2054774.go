package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int8 }
type GenS2 struct { F0 int32; F1 uint64; F2 int }
func genFn0() uint {
    v4 := GenS1{F0: -1, F1: 0}
    var v5 int = -437
    _ = v4
    _ = v5
    return 4294967295
}
func genFn1(p0 int64) string {
    p0 += (p0 & (^(p0)))
    for fi6 := 0; fi6 < 1; fi6++ {
        var v7 float64 = 1e10
        _ = fi6
        _ = v7
    }
    if (42 <= 612) {
        for fi8 := 0; fi8 < 15; fi8++ {
            v9 := []int{(fi8 ^ (fi8 ^ fi8)), (-(fi8))}
            _ = fi8
            _ = v9
        }
        fmt.Println(p0)
        v10 := GenS1{F0: -9223372036854775808, F1: 42}
        _ = v10
    }
    switch 0 {
    case -2147483648:
        var v11 float64 = 0.0
        _ = v11
    case 127:
        if (1 > 42) {
            v12 := []int{-2147483648, 9223372036854775807, 32768}
            if len(v12) > 0 {
                v12[0] = 539
            }
            if true {
                for fi13 := 0; fi13 < 14; fi13++ {
                    var v14 float64 = ((0.0 / 1.0) + 1e10)
                    if len(v12) > 0 {
                        v12[0] = (fi13 >> ((154) & 63))
                    }
                    _ = fi13
                    _ = v14
                }
            } else {
                if true {
                    var v15 string = "\n\t"
                    switch v15 {
                    case "a":
                        p0 -= (p0 >> ((1) & 63))
                    case "\n\t":
                        var v16 uint64 = 18446744073709551615
                        _ = v16
                    }
                    _ = v15
                }
            }
            _ = v12
        }
    case -1:
        var v17 string = "foo bar"
        _ = v17
    }
    _ = p0
    return "\n\t"
}

func main() {
    v1 := make(map[string]uint32)
    var v2 float32 = 0.0
    var v3 int32 = -1
    v18 := GenS2{F0: (^(v3)), F1: 18446744073709551615, F2: -9223372036854775808}
    for fi19 := 0; fi19 < 5; fi19++ {
        v3 -= (v3 - (^(v3)))
        var v20 int = fi19
        var v21 uint64 = 796
        _ = fi19
        _ = v20
        _ = v21
    }
    for fi22 := 0; fi22 < 7; fi22++ {
        var v23 uint32 = 2147483647
        fmt.Println(v2)
        _ = fi22
        _ = v23
    }
    var v24 uint32 = 0
    v25 := v1["a"]
    var v26 uint = 255
    v27 := GenS1{F0: int(v26), F1: 42}
    for fi28 := 0; fi28 < 2; fi28++ {
        v29 := []int{(^(fi28))}
        var v30 float64 = ((((3.14159 - 1.5) - 1.5) / 1.0) * float64(v2))
        v3 += v3
        _ = fi28
        _ = v29
        _ = v30
    }
    _ = len(v1)
    fmt.Println(v24)
    for fi31 := 0; fi31 < 11; fi31++ {
        v24 += v24
        _ = fi31
    }
    for fi32 := 0; fi32 < 1; fi32++ {
        v26 = genFn0()
        v2 = ((0.0 + ((3.14159 / ((v2 + ((float32(1.5) / 1.0) / ((float32((0.0 * 0.0)) / v2) / (v2 / (v2 * v2))))) * v2)) / v2)) + 3.14159)
        _ = fi32
    }
    var v33 int = 855
    v34 := GenS2{F0: v3, F1: 18446744073709551615, F2: -9223372036854775808}
    for fi35 := 0; fi35 < 12; fi35++ {
        fmt.Println(fi35)
        _ = fi35
    }
    _ = len(v1)
    fmt.Println(v2, v33)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v33:", v33)
    _ = v1
    _ = v18
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v3
    _ = v33
    _ = v34
}
