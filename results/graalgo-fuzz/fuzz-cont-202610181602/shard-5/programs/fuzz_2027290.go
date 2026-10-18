package main

import (
    "fmt"
)

type GenS1 struct { F0 int16 }
func genFn0(p0 bool, p1 int8) uint8 {
    var v7 float64 = (1.5 + 0.0)
    var v8 uint32 = 104
    v8 = v8
    fmt.Println(p0, v7)
    fmt.Println(p0, v8, p1)
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 0
}

func main() {
    var v1 uint8 = 127
    v2 := []int{38, 4294967295, 128, 32768}
    var v3 uint8 = (^(v1))
    var v4 int64 = -9223372036854775808
    var v5 int8 = 100
    v6 := make(map[string]string)
    v4 = v4
    for fi9 := 0; fi9 < 1; fi9++ {
        v3 = (v1 | v1)
        _ = fi9
    }
    v10 := GenS1{F0: -958}
    fmt.Println(v5)
    v1 += (v3 / (((^(v1))) | 1))
    for fi11 := 0; fi11 < 10; fi11++ {
        var v12 float64 = -2.25
        var v13 uint8 = (^(v1))
        _ = fi11
        _ = v12
        _ = v13
    }
    fmt.Println(v5)
    var v14 int16 = -201
    for fi15 := 0; fi15 < 4; fi15++ {
        v16 := GenS1{F0: v14}
        var v17 string = "\n\t"
        var v18 uint64 = 2
        _ = fi15
        _ = v16
        _ = v17
        _ = v18
    }
    switch ("\n\t" + "") {
    case "a":
        var v19 int = 923
        _ = v19
    default:
        v2 = append(v2, 798)
    }
    v20 := GenS1{F0: (v14 << ((1) & 63))}
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("maplen v6:", len(v6))
    _ = v1
    _ = v10
    _ = v14
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
