package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 string; F2 bool }

func main() {
    v1 := []uint32{1, 2147483648, 2}
    var v2 int32 = 630
    v3 := []string{("foo bar" + "foo bar"), "a", ("hello" + ("" + ("\n\t" + "hello")))}
    v4 := make(map[int]bool)
    var v5 int = 0
    var v6 uint16 = 2
    var v7 int64 = 42
    var v8 float32 = 1e10
    var v9 int = 1
    v10 := v4[v9]
    var v11 float32 = (((3.14159 - (v8 / ((v8 / float32(1.5)) * (1.5 * ((float32((3.14159 * float64(-2.25))) / v8) - (1e10 * ((float32((1.5 - -2.25)) / v8) + v8))))))) + float32(1.5)) / v8)
    v12 := GenS1{F0: 623, F1: (("" + "foo bar") + ("\n\t" + "\n\t")), F2: (501 > 783)}
    if len(v1) > 0 {
        v1[0] = uint32(v6)
    }
    for fi13 := 0; fi13 < 2; fi13++ {
        for fi14 := 0; fi14 < 8; fi14++ {
            v1 = append(v1, 42)
            _ = fi14
        }
        for fi15 := 0; fi15 < 11; fi15++ {
            v10 = true
            _ = fi15
        }
        _ = fi13
    }
    var v16 float64 = 3.14159
    v6 += (^(v6))
    if v10 {
        var v17 uint32 = 2147483647
        v18 := GenS1{F0: 2, F1: "\n\t", F2: v10}
        var v19 uint64 = 4294967295
        _ = v17
        _ = v18
        _ = v19
    } else {
        v16 -= (0.0 + 3.14159)
        switch v5 {
        case -2147483648:
            for fi20 := 0; fi20 < 15; fi20++ {
                var v21 int32 = 1
                v22 := GenS1{F0: 127, F1: ("hello" + "\n\t"), F2: (v16 < v16)}
                for fi23 := 0; fi23 < 14; fi23++ {
                    v24 := GenS1{F0: 18446744073709551615, F1: "foo bar", F2: v10}
                    v10 = (2 >= 136)
                    switch "" {
                    case "foo bar":
                        var v25 bool = v10
                        _ = v25
                    case "hello":
                        v21 = 2147483647
                    default:
                        var v26 int8 = 117
                        _ = v26
                    }
                    _ = fi23
                    _ = v24
                }
                _ = fi20
                _ = v21
                _ = v22
            }
        }
        var v27 float64 = (v16 - (v16 - 3.14159))
        _ = v27
    }
    var v28 int8 = -1
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v16:", v16)
    fmt.Println("v2:", v2)
    fmt.Println("v28:", v28)
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v16
    _ = v2
    _ = v28
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
