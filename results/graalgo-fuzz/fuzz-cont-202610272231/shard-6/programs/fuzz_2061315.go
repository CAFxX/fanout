package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 uint32; F2 uint32 }
type GenS2 struct { F0 int64 }
func genFn0() int32 {
    v6 := GenS1{F0: 127, F1: 2147483648, F2: 2}
    switch ("foo bar" + "a") {
    case "foo bar":
        v7 := []int{171, 127, -262, 128}
        _ = v7
    default:
        var v8 uint16 = 1
        _ = v8
    }
    v9 := []int{-673, -1, 409, 127}
    switch 4294967295 {
    case 32767:
        if len(v9) >= 2 {
            v9 = v9[1:len(v9)]
        }
    }
    switch -82 {
    case -939:
        var v10 uint8 = 1
        _ = v10
    default:
        var v11 uint32 = 2
        _ = v11
    }
    _ = v6
    _ = v9
    return 0
}

func main() {
    var v1 uint16 = 65535
    v2 := make(map[string]string)
    var v3 uint8 = 2
    v4 := make(map[uint32]bool)
    var v5 float64 = (-2.25 - 1.5)
    v12 := []int{}
    v13 := GenS2{F0: 1}
    var v14 float32 = (float32(1e10) - (0.1 * 3.14159))
    var v15 uint16 = (v1 * (v1 % ((42) | 1)))
    fmt.Println(v5, v3)
    for fi16 := 0; fi16 < 8; fi16++ {
        if false {
            v12 = append(v12, fi16)
        } else {
            var v17 uint = 65535
            _ = v17
        }
        var v18 uint64 = 2
        for fi19 := 0; fi19 < 18; fi19++ {
            v5 = v5
            v15 -= v15
            var v20 int16 = 42
            _ = fi19
            _ = v20
        }
        _ = fi16
        _ = v18
    }
    v1 += (^(v1))
    fmt.Println(v3)
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v12:", len(v12), v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
}
