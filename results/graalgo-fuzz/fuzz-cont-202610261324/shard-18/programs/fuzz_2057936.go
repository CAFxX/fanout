package main

import (
    "fmt"
)

type GenS1 struct { F0 uint8; F1 uint }
func genFn0(p0 uint32) uint8 {
    v5 := []int{-1, 42, -1}
    v6 := GenS1{F0: 42, F1: 4294967295}
    if ((-679 != 9223372036854775807) && (-624 == 0)) {
        var v7 int64 = 1
        for fi8 := 0; fi8 < 10; fi8++ {
            p0 = 168
            v5 = append(v5, (fi8 + fi8))
            _ = fi8
        }
        if len(v5) >= 2 {
            v5 = v5[1:len(v5)]
        }
        _ = v7
    } else {
        p0 -= 1
        fmt.Println(p0)
    }
    _ = p0
    _ = v5
    _ = v6
    return 1
}

func main() {
    v1 := []int{}
    var v2 int = -9223372036854775808
    v3 := make(map[int]string)
    var v4 int64 = 42
    v9 := v3[v2]
    v3[v2] = (("\n\t" + (v9 + ("\n\t" + v9))) + v9)
    v1 = append(v1, (v2 / (((v2 * 42)) | 1)))
    v9 = (v9 + v9)
    var v10 int32 = 2147483647
    v11 := v3[(v2 & v2)]
    v12 := GenS1{F0: 2, F1: 298}
    var v13 int64 = v4
    for fi14 := 0; fi14 < 20; fi14++ {
        v3[(v2 / (((v2 & int(v13))) | 1))] = (v9 + v9)
        var v15 int64 = 620
        _ = fi14
        _ = v15
    }
    if true {
        fmt.Println(v4, v11, v13)
    }
    switch v9 {
    case "a":
        for fi16 := 0; fi16 < 20; fi16++ {
            v2 = 32768
            _ = fi16
        }
    case "\n\t":
        v13 += v4
    default:
        v9 = v11
    }
    v17 := GenS1{F0: 192, F1: 691}
    switch "\n\t" {
    case "":
        v11 = v9
    case "a":
        v9 = ("" + (v11 + v11))
    }
    if (2 != 1) {
        v2 = 1
        switch 9223372036854775807 {
        case 32767:
            v2 = -9223372036854775808
        }
        v11 = ((v11 + v11) + v9)
    } else {
        v10 += v10
    }
    v18 := GenS1{F0: genFn0(0), F1: 2}
    fmt.Println("slicelen v1:", len(v1), v1)
    fmt.Println("v10:", v10)
    fmt.Println("v11:", v11)
    fmt.Println("v13:", v13)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("v4:", v4)
    fmt.Println("v9:", v9)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v17
    _ = v18
    _ = v2
    _ = v3
    _ = v4
    _ = v9
}
