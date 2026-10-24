package main

import (
    "fmt"
)

type GenS1 struct { F0 uint16; F1 int8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 uint32; F1 uint32 }
type GenS4 struct { F0 int8 }
func genFn0() int16 {
    var v4 uint64 = 42
    v5 := []int{503, 42, -793}
    _ = v4
    _ = v5
    return 0
}

func main() {
    var v1 int64 = -1
    var v2 uint8 = 42
    var v3 uint16 = uint16(v2)
    var v6 int8 = 127
    switch "foo bar" {
    case "a":
        v7 := GenS1{F0: v3, F1: (v6 << ((2) & 63))}
        _ = v7
    case "foo bar":
        var v8 int64 = (v1 * v1)
        _ = v8
    default:
        v9 := []int{2147483647, 0, 4294967295, -2147483648}
        _ = v9
    }
    v10 := []int{740}
    var v11 float64 = 0.1
    v10 = append(v10, 614)
    var v12 uint16 = v3
    v13 := GenS2{F0: 42}
    var v14 int = 128
    fmt.Println(v2, v14)
    if (0 <= (v2 << ((18446744073709551615) & 63))) {
        var v15 float64 = 1e10
        _ = v15
    }
    v16 := GenS2{F0: 255}
    fmt.Println(v2)
    if true {
        var v17 string = "a"
        v18 := GenS3{F0: 668, F1: 793}
        fmt.Println(v17, v11, v17)
        _ = v17
        _ = v18
    }
    if true {
        _ = len(v10)
        fmt.Println(v6, v11)
    } else {
        var v19 int16 = -32768
        for fi20 := 0; fi20 < 12; fi20++ {
            fmt.Println(v19)
            if ((^(v2)) >= v2) {
                if false {
                    var v21 uint8 = (v2 + v2)
                    _ = v21
                }
            }
            _ = fi20
        }
        _ = v19
    }
    if true {
        switch ("" + "") {
        case "a":
            v1 += 0
        }
    } else {
        v22 := GenS4{F0: (v6 ^ -128)}
        _ = v22
    }
    var v23 uint64 = 496
    fmt.Println("v1:", v1)
    fmt.Println("slicelen v10:", len(v10), v10)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v2:", v2)
    fmt.Println("v23:", v23)
    fmt.Println("v3:", v3)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v16
    _ = v2
    _ = v23
    _ = v3
    _ = v6
}
