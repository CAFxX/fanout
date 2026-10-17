package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int }
type GenS2 struct { F0 int64; F1 uint32; F2 uint }

func main() {
    var v1 int = 946
    v2 := make(map[string]int)
    v3 := make(map[uint32]bool)
    for fi4 := 0; fi4 < 13; fi4++ {
        switch "a" {
        case "foo bar":
            v5 := GenS1{F0: 156, F1: 127}
            _ = v5
        case "a":
            var v6 float64 = ((((3.14159 * 1.5) / (3.14159 - 0.0)) / 0.1) / 1.0)
            _ = v6
        default:
            switch "\n\t" {
            case "a":
                for fi7 := 0; fi7 < 10; fi7++ {
                    v1 = (fi4 & fi7)
                    v8 := []int{(fi7 & -379)}
                    var v9 uint64 = 18446744073709551615
                    _ = fi7
                    _ = v8
                    _ = v9
                }
            }
        }
        for fi10 := 0; fi10 < 8; fi10++ {
            fmt.Println(fi10)
            var v11 uint8 = 127
            var v12 uint32 = 896
            _ = fi10
            _ = v11
            _ = v12
        }
        var v13 int64 = 0
        _ = fi4
        _ = v13
    }
    var v14 uint8 = 127
    var v15 int = 201
    v14 -= (v14 >> ((18446744073709551615) & 63))
    v16 := []int{1}
    v17 := v3[2147483647]
    var v18 int16 = -80
    var v19 string = "a"
    v20 := GenS1{F0: 284, F1: v15}
    v1 = v1
    var v21 int16 = (v18 - v18)
    var v22 uint32 = 65535
    for fi23 := 0; fi23 < 17; fi23++ {
        if v17 {
            var v24 float32 = ((1e10 + -2.25) / (((1e10 - 3.14159) / 1.5) - (float32((float64((-2.25 + -2.25)) * float64(0.0))) * (-2.25 + (3.14159 * 3.14159)))))
            _ = v24
        } else {
            v16 = append(v16, -204)
            for fi25 := 0; fi25 < 16; fi25++ {
                fmt.Println(v14, v17)
                var v26 uint8 = v14
                _ = fi25
                _ = v26
            }
        }
        var v27 string = "hello"
        fmt.Println(v14, v22)
        _ = fi23
        _ = v27
    }
    switch int(v18) {
    case -765:
        v15 += (v15 + v1)
    case 9223372036854775807:
        v16 = append(v16, (^(v15)))
    }
    v28 := GenS2{F0: -9223372036854775808, F1: v22, F2: 18446744073709551615}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("slicelen v16:", len(v16), v16)
    fmt.Println("v17:", v17)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v22
    _ = v28
    _ = v3
}
