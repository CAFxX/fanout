package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 string; F2 int64 }

func main() {
    v1 := make(map[int]string)
    v2 := make(map[int]int)
    v3 := make(map[int]string)
    v4 := []string{}
    v5 := make(map[string]bool)
    var v6 uint32 = 328
    var v7 uint64 = 1
    if true {
        fmt.Println(v7)
        if (0 >= -1) {
            if ("\n\t" != ("foo bar" + ("" + "\n\t"))) {
                v7 += 4294967295
                v8 := GenS1{F0: 1, F1: "hello", F2: 42}
                _ = v8
            }
        } else {
            for fi9 := 0; fi9 < 13; fi9++ {
                for fi10 := 0; fi10 < 12; fi10++ {
                    fmt.Println(fi9, v6)
                    _ = fi10
                }
                _ = fi9
            }
            v3[32767] = ("hello" + (("\n\t" + "hello") + "\n\t"))
            var v11 int = 128
            _ = v11
        }
    }
    for fi12 := 0; fi12 < 1; fi12++ {
        _ = len(v5)
        var v13 uint64 = v7
        _ = fi12
        _ = v13
    }
    var v14 uint16 = 0
    var v15 uint16 = 65535
    v7 += v7
    v16 := GenS1{F0: 672, F1: "a", F2: 9223372036854775807}
    var v17 float64 = 1e10
    v6 = v6
    v7 = v7
    var v18 float32 = (float32(0.0) * float32((float64((((0.0 - (1.5 / 3.14159)) - 0.0) + 0.1)) * v17)))
    v19 := v1[4294967295]
    if true {
        v14 -= uint16(v7)
        v17 = float64(0.1)
        var v20 int = 188
        _ = v20
    }
    if true {
        v17 -= (v17 + -2.25)
        switch -524 {
        case 2147483647:
            v14 = v14
        default:
            v21 := GenS1{F0: 42, F1: v19, F2: 9223372036854775807}
            _ = v21
        }
    }
    fmt.Println(v18)
    v22 := GenS1{F0: 42, F1: (v19 + v19), F2: 9223372036854775807}
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    fmt.Println("v7:", v7)
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v22
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
}
