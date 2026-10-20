package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 bool; F2 int64 }
type GenS2 struct { F0 uint; F1 uint8; F2 bool }
func genFn0(p0 uint64, p1 int64) bool {
    p1 = -9223372036854775808
    p1 -= p1
    var v7 int32 = 965
    v7 += 94
    fmt.Println(v7)
    _ = p0
    _ = p1
    _ = v7
    return false
}

func main() {
    var v1 float32 = 0.0
    v2 := make(map[string]bool)
    v3 := []string{"a"}
    v4 := make(map[int]bool)
    v5 := []int{0, -940}
    var v6 float64 = float64(((-2.25 * ((0.0 - v1) - v1)) + v1))
    v8 := GenS1{F0: 65, F1: false, F2: -1}
    v1 -= v1
    v1 -= (float32(v6) * v1)
    if false {
        fmt.Println(v6, v1)
        var v9 bool = true
        v10 := GenS1{F0: -253, F1: v9, F2: -1}
        _ = v10
        _ = v9
    }
    v11 := GenS2{F0: 4294967295, F1: 2, F2: (!(-2147483648 > 552))}
    var v12 uint32 = 2147483647
    var v13 uint16 = 117
    var v14 string = ""
    if false {
        for fi15 := 0; fi15 < 4; fi15++ {
            v13 = (^(v13))
            v14 = v14
            if (2147483648 == (v12 & v12)) {
                var v16 uint32 = 2147483647
                _ = v16
            }
            _ = fi15
        }
        var v17 int16 = 1000
        fmt.Println(v6, v1)
        _ = v17
    }
    if false {
        v1 = 3.14159
        var v18 int16 = 42
        var v19 int64 = 42
        _ = v18
        _ = v19
    } else {
        if false {
            v20 := GenS1{F0: 128, F1: true, F2: 1}
            _ = v20
        }
        v13 -= (v13 - v13)
    }
    if ((42 < 1) || (2 > 899)) {
        if true {
            v21 := GenS1{F0: -9223372036854775808, F1: ((!(-2.25 == 0.0)) || (!(v13 < v13))), F2: 0}
            _ = len(v5)
            fmt.Println(v14, v6)
            _ = v21
        }
        v14 = ("\n\t" + "hello")
        var v22 string = ""
        _ = v22
    } else {
        v4[532] = false
        var v23 float32 = v1
        _ = v23
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("slicelen v3:", len(v3), v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("slicelen v5:", len(v5), v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v11
    _ = v12
    _ = v13
    _ = v14
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v8
}
