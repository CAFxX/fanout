package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64; F2 int8 }
type GenS2 struct { F0 bool }
func genFn0() int {
    var v4 uint32 = 0
    var v5 uint8 = 127
    fmt.Println(v4)
    _ = v4
    _ = v5
    return -639
}

func main() {
    v1 := []string{"\n\t", (("foo bar" + "foo bar") + "foo bar"), ""}
    var v2 uint32 = 1
    v3 := make(map[int]int)
    v2 = (v2 << ((255) & 63))
    switch 795 {
    case -9223372036854775808:
        v2 = 4294967295
    case 32767:
        for fi6 := 0; fi6 < 15; fi6++ {
            for fi7 := 0; fi7 < 18; fi7++ {
                switch "foo bar" {
                case "\n\t":
                    var v8 int = 562
                    _ = v8
                }
                _ = fi7
            }
            _ = len(v3)
            var v9 int16 = 1000
            _ = fi6
            _ = v9
        }
    case -432:
        if true {
            v2 = (v2 / (((^(v2))) | 1))
        } else {
            var v10 uint8 = 2
            var v11 float32 = 1.5
            var v12 float64 = float64(((v11 * (v11 + (3.14159 + 0.1))) * (v11 + ((((v11 * (v11 - float32(-2.25))) * (v11 / 1.0)) + (v11 * float32(0.1))) - 0.0))))
            _ = v10
            _ = v11
            _ = v12
        }
    }
    v13 := GenS1{F0: -128, F1: 363, F2: -1}
    var v14 int = 1
    v15 := GenS1{F0: 1, F1: 1, F2: 48}
    v16 := GenS1{F0: -36, F1: -9223372036854775808, F2: int8(v14)}
    for fi17 := 0; fi17 < 20; fi17++ {
        var v18 float64 = -2.25
        _ = fi17
        _ = v18
    }
    v19 := GenS1{F0: 127, F1: 9223372036854775807, F2: 0}
    v20 := GenS2{F0: false}
    fmt.Println(v14)
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v19
    _ = v2
    _ = v20
    _ = v3
}
