package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int64; F2 int16 }
type GenS2 struct { F0 uint64; F1 int8; F2 int8 }

func main() {
    v1 := []int{1, 9223372036854775807, 4294967295}
    var v2 int8 = 1
    v3 := make(map[int]int)
    var v4 uint8 = 128
    v5 := GenS1{F0: 42, F1: -9223372036854775808, F2: 32767}
    fmt.Println(v4, v4)
    switch "a" {
    case "hello":
        var v6 uint64 = 18446744073709551615
        _ = v6
    case "foo bar":
        switch 2147483647 {
        case 1:
            fmt.Println(v2, v2)
        case 32767:
            if (v4 > (v4 % ((v4) | 1))) {
                for fi7 := 0; fi7 < 2; fi7++ {
                    v8 := GenS1{F0: (fi7 / ((fi7) | 1)), F1: 1, F2: 427}
                    _ = fi7
                    _ = v8
                }
                fmt.Println(v2, v4)
                if true {
                    v4 += (v4 & v4)
                    v2 += -86
                    if (!true) {
                        v2 = 42
                        var v9 float64 = (float64((3.14159 * 1e10)) / 0.1)
                        _ = v9
                    }
                }
            } else {
                var v10 float64 = 0.1
                v11 := v3[-583]
                _ = v10
                _ = v11
            }
        default:
            _ = len(v1)
        }
    }
    v12 := GenS1{F0: 32767, F1: 80, F2: -1}
    var v13 uint16 = 846
    var v14 int64 = 9223372036854775807
    v4 = (v4 & (v4 >> ((18446744073709551615) & 63)))
    for fi15 := 0; fi15 < 3; fi15++ {
        v16 := GenS2{F0: 0, F1: (v2 | (-(v2))), F2: v2}
        var v17 int64 = v14
        _ = fi15
        _ = v16
        _ = v17
    }
    if true {
        v13 = 32767
        v18 := GenS1{F0: -1, F1: v14, F2: -32768}
        _ = v18
    } else {
        _ = len(v1)
    }
    fmt.Println(v13)
    fmt.Println(v2, v2, v13)
    v19 := GenS2{F0: 63, F1: (v2 | 100), F2: 127}
    var v20 string = (("a" + ("" + "")) + "")
    switch 32768 {
    case 4294967295:
        v2 -= v2
    }
    v3[227] = -489
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v20:", v20)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v19
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
}
