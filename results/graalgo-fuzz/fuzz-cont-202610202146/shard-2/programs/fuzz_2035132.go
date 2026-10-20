package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32 }
type GenS2 struct { F0 uint; F1 int8 }
func genFn0() uint64 {
    var v7 int32 = 2147483647
    fmt.Println(v7)
    _ = v7
    return 18446744073709551615
}

func main() {
    var v1 int32 = 0
    v2 := make(map[string]string)
    v3 := make(map[string]int)
    v4 := make(map[uint32]int)
    v5 := []uint32{2, 532, 2147483648}
    var v6 int64 = 0
    var v8 float64 = -2.25
    v9 := v2["\n\t"]
    var v10 string = v9
    fmt.Println(v10, v6)
    v11 := GenS1{F0: 2147483647}
    v2[v10] = (v9 + v10)
    v12 := v2[v10]
    fmt.Println(v10)
    for fi13 := 0; fi13 < 17; fi13++ {
        switch fi13 {
        case -9223372036854775808:
            v14 := GenS1{F0: 2}
            _ = v14
        case 4294967295:
            for fi15 := 0; fi15 < 14; fi15++ {
                var v16 uint8 = 0
                fmt.Println(fi15, v10, v10)
                v16 = (v16 + v16)
                v16 = (^(v16))
                switch v9 {
                case "foo bar":
                    var v17 bool = false
                    _ = v17
                }
                _ = fi15
                _ = v16
            }
        default:
            var v18 float32 = 0.0
            _ = v18
        }
        v19 := GenS2{F0: 18446744073709551615, F1: -128}
        var v20 uint32 = 664
        _ = fi13
        _ = v19
        _ = v20
    }
    v3[v10] = -897
    var v21 int8 = 1
    var v22 uint16 = 42
    var v23 float32 = (-2.25 * float32((v8 * (v8 - 0.1))))
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v21:", v21)
    fmt.Println("v22:", v22)
    fmt.Println("v23:", v23)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    fmt.Println("v8:", v8)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v21
    _ = v22
    _ = v23
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
    _ = v9
}
