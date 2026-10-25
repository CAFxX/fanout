package main

import (
    "fmt"
)

type GenS1 struct { F0 int32; F1 bool; F2 int64 }
type GenS2 struct { F0 uint32 }
func genFn0(p0 uint) int64 {
    var v5 uint64 = 962
    if false {
        v5 += (v5 - v5)
        var v6 uint8 = 127
        _ = v6
    } else {
        var v7 int8 = 127
        var v8 uint16 = 1
        _ = v7
        _ = v8
    }
    v9 := []int{1, 1, -848}
    v9 = append(v9, -69)
    p0 += p0
    _ = p0
    _ = v5
    _ = v9
    return int64(v5)
}

func main() {
    var v1 uint16 = 566
    v2 := []int{127, -2147483648}
    var v3 uint8 = 215
    var v4 bool = ((v1 * v1) < (v1 << ((1) & 63)))
    fmt.Println(v1, v3, v1)
    fmt.Println(v1, v3, v1)
    var v10 int32 = -2147483648
    if (!(!v4)) {
        switch "foo bar" {
        case "hello":
            v10 -= (v10 + v10)
        default:
            var v11 int = 570
            _ = v11
        }
    }
    v3 += (^(v3))
    var v12 int32 = v10
    if v4 {
        var v13 uint16 = 2
        _ = v13
    }
    v14 := GenS1{F0: v10, F1: (2147483647 == 251), F2: -364}
    if false {
        v15 := GenS2{F0: 2147483648}
        _ = v15
    } else {
        var v16 float32 = 1.5
        var v17 string = "foo bar"
        _ = v16
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v10:", v10)
    fmt.Println("v12:", v12)
    fmt.Println("slicelen v2:", len(v2), v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    _ = v1
    _ = v10
    _ = v12
    _ = v14
    _ = v2
    _ = v3
    _ = v4
}
