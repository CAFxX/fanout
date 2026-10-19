package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint64; F2 string }
type GenS2 struct { F0 uint8; F1 int64 }
func genFn0(p0 uint64) string {
    switch 0 {
    case -1:
        var v4 int8 = 1
        _ = v4
    default:
        var v5 string = "a"
        _ = v5
    }
    fmt.Println(p0)
    for fi6 := 0; fi6 < 16; fi6++ {
        if (0 == -1) {
            for fi7 := 0; fi7 < 8; fi7++ {
                v8 := GenS1{F0: -9223372036854775808, F1: 4294967295, F2: ((("a" + "\n\t") + ("" + "a")) + "hello")}
                var v9 uint8 = 1
                _ = fi7
                _ = v8
                _ = v9
            }
            v10 := GenS1{F0: -9223372036854775808, F1: p0, F2: ""}
            v11 := []int{}
            _ = v10
            _ = v11
        }
        _ = fi6
    }
    var v12 uint64 = 173
    _ = p0
    _ = v12
    return ("\n\t" + "a")
}

func main() {
    v1 := make(map[string]uint32)
    v2 := []uint32{}
    var v3 int = -9223372036854775808
    var v13 uint8 = 255
    var v14 uint32 = 2147483647
    for fi15 := 0; fi15 < 17; fi15++ {
        fmt.Println(v14, v13, v13)
        _ = fi15
    }
    v16 := GenS2{F0: (v13 / ((v13) | 1)), F1: 9223372036854775807}
    var v17 float32 = (((0.0 * 0.1) * -2.25) / 1.5)
    for fi18 := 0; fi18 < 20; fi18++ {
        var v19 int64 = 1
        v17 = 3.14159
        _ = fi18
        _ = v19
    }
    v20 := GenS2{F0: v13, F1: 822}
    fmt.Println(v17, v13)
    fmt.Println(v17, v14)
    v2 = append(v2, v14)
    var v21 int8 = -128
    var v22 bool = (33 == 1)
    v1["foo bar"] = (^(v14))
    v2 = append(v2, uint32(v13))
    v3 += v3
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v16
    _ = v17
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v3
}
