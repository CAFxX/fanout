package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 uint32 }
type GenS2 struct { F0 string; F1 uint8; F2 uint }
func genFn0(p0 uint16, p1 bool) uint64 {
    v5 := GenS1{F0: 91, F1: 808}
    v6 := GenS1{F0: -2147483648, F1: 2147483647}
    v7 := GenS1{F0: 206, F1: 659}
    _ = p0
    _ = p1
    _ = v5
    _ = v6
    _ = v7
    return 4294967295
}
func genFn1(p0 int32, p1 bool, p2 int64) uint16 {
    v8 := []int{}
    var v9 bool = true
    var v10 uint32 = 2147483648
    var v11 int64 = (p2 - 42)
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v11
    _ = v8
    _ = v9
    return 2
}

func main() {
    v1 := make(map[int]int)
    v2 := make(map[int]uint32)
    var v3 bool = ((0.0 * 1e10) >= 1.5)
    var v4 bool = (false && (255 > 748))
    var v12 float64 = 1e10
    v13 := GenS1{F0: -1, F1: 834}
    for fi14 := 0; fi14 < 6; fi14++ {
        for fi15 := 0; fi15 < 2; fi15++ {
            v16 := []int{fi14}
            v12 -= v12
            fmt.Println(v3, v12, fi14)
            v17 := GenS1{F0: 42, F1: 1}
            _ = fi15
            _ = v16
            _ = v17
        }
        _ = fi14
    }
    v3 = v3
    var v18 uint64 = genFn0(658, v4)
    var v19 bool = (-2147483648 != 534)
    v2[127] = 2
    v12 = v12
    var v20 int = 137
    switch "hello" {
    case "foo bar":
        v21 := []int{v20, v20, (^(v20))}
        _ = v21
    case "":
        v22 := []int{4294967295}
        _ = v22
    case "hello":
        for fi23 := 0; fi23 < 2; fi23++ {
            v19 = (42 <= 29)
            _ = fi23
        }
    default:
        if v19 {
            var v24 uint8 = 0
            v25 := []int{32767, -9223372036854775808, (-(v20))}
            _ = v24
            _ = v25
        }
    }
    v26 := GenS2{F0: (("\n\t" + ("foo bar" + "foo bar")) + "foo bar"), F1: 2, F2: 348}
    var v27 int16 = 0
    delete(v1, 2147483647)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v26
    _ = v27
    _ = v3
    _ = v4
}
