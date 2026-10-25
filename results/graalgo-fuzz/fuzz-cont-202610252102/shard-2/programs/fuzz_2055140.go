package main

import (
    "fmt"
)

type GenS1 struct { F0 bool }
type GenS2 struct { F0 int16 }
func genFn0() uint8 {
    switch "" {
    case "":
        for fi6 := 0; fi6 < 9; fi6++ {
            var v7 int64 = -1
            v8 := GenS1{F0: (654 < (v7 / (((^(v7))) | 1)))}
            v9 := []int{(fi6 << ((255) & 63)), fi6, -9223372036854775808}
            _ = fi6
            _ = v7
            _ = v8
            _ = v9
        }
    }
    v10 := GenS1{F0: true}
    if true {
        v11 := []int{-1, -9223372036854775808}
        if true {
            var v12 float64 = float64((0.1 / 0.1))
            var v13 float32 = (0.0 * (0.0 * 3.14159))
            var v14 int64 = 1
            _ = v12
            _ = v13
            _ = v14
        }
        if len(v11) > 0 {
            v11[0] = 0
        }
        _ = v11
    }
    _ = v10
    return 110
}

func main() {
    var v1 int8 = -106
    v2 := []uint32{4294967295, 2147483648, 465, 1}
    v3 := make(map[string]bool)
    var v4 uint64 = 1
    var v5 int8 = -55
    var v15 float32 = float32((-2.25 / 0.1))
    v1 -= (v5 - v5)
    v15 = 1e10
    if (false && true) {
        v4 += (v4 + (v4 << ((18446744073709551615) & 63)))
        for fi16 := 0; fi16 < 18; fi16++ {
            switch fi16 {
            case 9223372036854775807:
                v4 -= v4
            }
            _ = len(v2)
            _ = fi16
        }
    } else {
        v4 = v4
        var v17 float32 = v15
        _ = v17
    }
    v15 = v15
    var v18 uint8 = 1
    v19 := GenS1{F0: true}
    switch 32768 {
    case 595:
        fmt.Println(v1, v4, v15)
    case 805:
        switch 2147483647 {
        case 9223372036854775807:
            var v20 uint16 = 668
            _ = v20
        }
    }
    v21 := GenS2{F0: -713}
    for fi22 := 0; fi22 < 9; fi22++ {
        var v23 uint = 2
        _ = fi22
        _ = v23
    }
    fmt.Println(v1, v5)
    v24 := v3["a"]
    var v25 string = ("hello" + ("a" + "a"))
    delete(v3, "a")
    var v26 uint16 = 42
    v27 := v3[v25]
    var v28 int8 = -128
    fmt.Println("v1:", v1)
    fmt.Println("v15:", v15)
    fmt.Println("v18:", v18)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v24:", v24)
    fmt.Println("v25:", v25)
    fmt.Println("v26:", v26)
    fmt.Println("v27:", v27)
    fmt.Println("v28:", v28)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    _ = v1
    _ = v15
    _ = v18
    _ = v19
    _ = v2
    _ = v21
    _ = v24
    _ = v25
    _ = v26
    _ = v27
    _ = v28
    _ = v3
    _ = v4
    _ = v5
}
