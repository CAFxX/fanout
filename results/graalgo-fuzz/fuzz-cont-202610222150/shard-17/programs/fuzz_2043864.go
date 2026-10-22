package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int8; F2 uint64 }

func main() {
    v1 := []string{"hello", "a", ("foo bar" + "a")}
    v2 := []int{-960, 2147483647, 591, 127}
    v3 := make(map[int]bool)
    v4 := make(map[string]uint32)
    var v5 uint32 = 65535
    v6 := GenS1{F0: 255, F1: -1, F2: 0}
    var v7 uint32 = (v5 ^ v5)
    if len(v1) >= 2 {
        v1 = v1[1:len(v1)]
    }
    v8 := GenS1{F0: 1, F1: 1, F2: 2}
    var v9 int = 127
    v10 := v4[("" + "foo bar")]
    v9 -= v9
    var v11 int32 = 0
    switch "hello" {
    case "":
        switch 0 {
        case 841:
            var v12 bool = (false || ((v11 ^ 42) >= v11))
            _ = v12
        default:
            v13 := GenS1{F0: 255, F1: 127, F2: 428}
            _ = v13
        }
    case "foo bar":
        var v14 float64 = float64(1e10)
        _ = v14
    case "a":
        for fi15 := 0; fi15 < 4; fi15++ {
            if false {
                var v16 uint = 2
                var v17 int16 = 0
                fmt.Println(v10)
                _ = v16
                _ = v17
            } else {
                v18 := GenS1{F0: 18446744073709551615, F1: -128, F2: 42}
                switch 476 {
                case 636:
                    v3[(fi15 / ((343) | 1))] = (-9223372036854775808 == 42)
                }
                v10 -= v7
                _ = v18
            }
            _ = fi15
        }
    default:
        var v19 uint64 = 4294967295
        _ = v19
    }
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v7:", v7)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v7
    _ = v8
    _ = v9
}
