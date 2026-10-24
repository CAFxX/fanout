package main

import (
    "fmt"
)

type GenS1 struct { F0 int64; F1 uint16; F2 uint64 }
func genFn0(p0 string, p1 uint64) int16 {
    v7 := []int{-621, 1, 4294967295}
    if len(v7) >= 2 {
        v7 = v7[1:len(v7)]
    }
    _ = p0
    _ = p1
    _ = v7
    return 32767
}
func genFn1(p0 bool, p1 string, p2 int64) int16 {
    p1 = p1
    var v8 int = 2147483647
    var v9 float32 = float32(3.14159)
    if (v8 < v8) {
        fmt.Println(p0, p1, v9)
    } else {
        var v10 int32 = -2147483648
        v11 := []int{}
        _ = v10
        _ = v11
    }
    _ = p0
    _ = p1
    _ = p2
    _ = v8
    _ = v9
    return -32768
}

func main() {
    var v1 int8 = -128
    var v2 uint = 570
    var v3 bool = (true || (true || false))
    var v4 int32 = -666
    v5 := make(map[int]string)
    var v6 uint = v2
    var v12 string = "a"
    fmt.Println(v12, v4, v1)
    v6 -= 2
    v1 -= -1
    for fi13 := 0; fi13 < 12; fi13++ {
        v3 = false
        _ = fi13
    }
    v12 = "\n\t"
    v14 := GenS1{F0: 0, F1: 255, F2: 1}
    fmt.Println(v2, v12, v6)
    var v15 int32 = v4
    for fi16 := 0; fi16 < 6; fi16++ {
        if v3 {
            var v17 uint16 = 174
            _ = v17
        } else {
            v18 := []int{(fi16 / ((fi16) | 1)), (fi16 & (fi16 << ((v2) & 63))), -2147483648, fi16}
            for fi19 := 0; fi19 < 19; fi19++ {
                v20 := v5[(-(fi19))]
                v1 -= v1
                _ = fi19
                _ = v20
            }
            _ = v18
        }
        var v21 int32 = v4
        _ = fi16
        _ = v21
    }
    fmt.Println(v2, v15, v6)
    switch v12 {
    case "":
        v22 := GenS1{F0: 29, F1: 983, F2: 0}
        _ = v22
    case "a":
        v23 := GenS1{F0: 0, F1: 141, F2: 2}
        _ = v23
    case "foo bar":
        var v24 float64 = ((0.0 * float64(3.14159)) - float64(3.14159))
        _ = v24
    default:
        var v25 uint32 = 2
        _ = v25
    }
    v26 := GenS1{F0: 9223372036854775807, F1: 900, F2: 18446744073709551615}
    fmt.Println("v1:", v1)
    fmt.Println("v12:", v12)
    fmt.Println("v15:", v15)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    fmt.Println("v6:", v6)
    _ = v1
    _ = v12
    _ = v14
    _ = v15
    _ = v2
    _ = v26
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
