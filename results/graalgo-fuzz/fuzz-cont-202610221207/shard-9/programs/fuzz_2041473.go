package main

import (
    "fmt"
)

type GenS1 struct { F0 uint; F1 int16; F2 int16 }
func genFn0(p0 uint32, p1 uint8, p2 int32) uint {
    var v6 int32 = 1
    v7 := []int{42}
    _ = p0
    _ = p1
    _ = p2
    _ = v6
    _ = v7
    return 60
}

func main() {
    var v1 uint = 65535
    var v2 int64 = 9223372036854775807
    var v3 bool = false
    var v4 int32 = -259
    v5 := make(map[int]bool)
    v1 -= 65535
    v2 += v2
    v8 := []int{372, 0, -440}
    v2 = v2
    if v3 {
        v9 := GenS1{F0: v1, F1: 1, F2: -854}
        _ = v9
    } else {
        v3 = (!(!v3))
        var v10 float32 = ((0.1 + float32((1.5 + 1.5))) + ((((-2.25 + ((1.5 * ((((0.0 * 1e10) / -2.25) * 1e10) + 0.0)) + ((1.5 + 1e10) - 0.1))) + 0.0) * 0.0) * 1e10))
        _ = v10
    }
    if ((!false) && ("a" <= "foo bar")) {
        var v11 int32 = (v4 % (((v4 << ((v1) & 63))) | 1))
        _ = v11
    }
    var v12 uint64 = 0
    var v13 uint64 = 194
    var v14 int32 = (v4 & -768)
    fmt.Println(v14, v4)
    var v15 float32 = 1.5
    for fi16 := 0; fi16 < 10; fi16++ {
        switch "foo bar" {
        case "foo bar":
            delete(v5, (fi16 << ((v1) & 63)))
        default:
            var v17 int32 = 19
            _ = v17
        }
        v18 := GenS1{F0: v1, F1: 32767, F2: -142}
        _ = fi16
        _ = v18
    }
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("slicelen v8:", len(v8), v8)
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
    _ = v4
    _ = v5
    _ = v8
}
