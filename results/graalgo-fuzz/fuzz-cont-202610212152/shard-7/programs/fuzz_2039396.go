package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 int32; F2 bool }
type GenS2 struct { F0 int; F1 uint64 }
func genFn0(p0 string) int64 {
    if false {
        var v4 uint8 = 127
        _ = v4
    }
    var v5 uint8 = 43
    for fi6 := 0; fi6 < 4; fi6++ {
        p0 = ""
        for fi7 := 0; fi7 < 2; fi7++ {
            v8 := GenS1{F0: 4294967295, F1: 922, F2: false}
            _ = fi7
            _ = v8
        }
        var v9 bool = (255 > 255)
        _ = fi6
        _ = v9
    }
    var v10 float64 = ((1.5 + 1e10) * 1.5)
    _ = p0
    _ = v10
    _ = v5
    return 9223372036854775807
}

func main() {
    v1 := make(map[string]uint32)
    v2 := make(map[uint32]int)
    var v3 uint32 = 2
    var v11 uint8 = 127
    var v12 int8 = 100
    delete(v2, v3)
    v13 := GenS2{F0: -1, F1: 4294967295}
    var v14 bool = (true && (804 <= 459))
    v12 -= (^(v12))
    v3 = (v3 | v3)
    switch "\n\t" {
    case "foo bar":
        v11 = v11
    default:
        v14 = (!(916 == 661))
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
}
