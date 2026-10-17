package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 int16; F2 bool }
func genFn0(p0 int64, p1 int32, p2 uint8) int {
    var v6 int8 = -12
    if (42 > 218) {
        var v7 int = 672
        v8 := []int{(v7 / ((v7) | 1)), (-(v7)), v7}
        var v9 int8 = v6
        _ = v7
        _ = v8
        _ = v9
    }
    p2 -= p2
    v10 := GenS1{F0: false, F1: 0, F2: ((float64(1.5) + 0.0) >= 0.1)}
    if true {
        p0 += p0
        p2 += (p2 + (p2 % ((p2) | 1)))
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v6
    return 254
}
func genFn1(p0 uint32, p1 uint64) uint16 {
    var v11 bool = (false && (2147483647 >= -172))
    var v12 uint8 = 255
    v13 := []int{1, 42, 4294967295}
    for fi14 := 0; fi14 < 18; fi14++ {
        p1 += (p1 << ((342) & 63))
        v15 := GenS1{F0: (16 == 519), F1: -1, F2: ((!(!v11)) || (65535 == 1))}
        fmt.Println(p0, v11)
        _ = fi14
        _ = v15
    }
    _ = p0
    _ = p1
    _ = v11
    _ = v12
    _ = v13
    return 42
}

func main() {
    v1 := []int{1, 363, -2147483648, 32767}
    var v2 int8 = -58
    var v3 float32 = -2.25
    v4 := make(map[uint32]int)
    var v5 int64 = 9223372036854775807
    delete(v4, 1)
    switch "" {
    case "\n\t":
        var v16 string = ""
        _ = v16
    case "hello":
        var v17 int16 = -1
        _ = v17
    case "a":
        v1 = append(v1, genFn0(v5, 42, 1))
    }
    var v18 int16 = -370
    var v19 int8 = v2
    var v20 string = "foo bar"
    v21 := GenS1{F0: (v3 <= 3.14159), F1: (v18 / ((v18) | 1)), F2: ((2 == 1) || (!(1 != 1)))}
    v3 -= 0.0
    var v22 int8 = (^(v19))
    var v23 int16 = -32768
    for fi24 := 0; fi24 < 17; fi24++ {
        v18 += v18
        v1 = append(v1, (fi24 - fi24))
        _ = fi24
    }
    v3 = (v3 * (-2.25 / 1.5))
    v1 = append(v1, 995)
    switch 32768 {
    case 32768:
        var v25 bool = true
        _ = v25
    }
    v3 = (float32((-2.25 - 1e10)) / v3)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
}
