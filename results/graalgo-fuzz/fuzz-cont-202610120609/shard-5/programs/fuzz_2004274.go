package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 uint64 }
type GenS2 struct { F0 bool; F1 int16 }
type GenS3 struct { F0 int }
func genFn0(p0 uint32) string {
    var v6 bool = true
    fmt.Println(p0, v6)
    var v7 uint = 42
    if v6 {
        v8 := []int{537, -2147483648}
        var v9 uint8 = 128
        _ = v8
        _ = v9
    } else {
        p0 = (p0 - 42)
        var v10 uint = (v7 << ((681) & 63))
        _ = v10
    }
    p0 = (^(p0))
    _ = p0
    _ = v6
    _ = v7
    return "\n\t"
}

func main() {
    v1 := make(map[int]bool)
    var v2 int8 = 127
    var v3 uint = 463
    var v4 float64 = 0.1
    v5 := make(map[int]uint32)
    switch "" {
    case "foo bar":
        v11 := GenS1{F0: false, F1: 0}
        _ = v11
    }
    fmt.Println(v2, v4, v4)
    v4 += v4
    var v12 uint32 = 845
    delete(v1, -9223372036854775808)
    var v13 string = ""
    if true {
        delete(v1, -143)
        var v14 string = v13
        var v15 uint8 = 127
        _ = v14
        _ = v15
    } else {
        var v16 int32 = 8
        var v17 int8 = -1
        _ = v16
        _ = v17
    }
    v3 -= (v3 - 760)
    var v18 float32 = -2.25
    if (((1e10 - v4) * v4) >= -2.25) {
        for fi19 := 0; fi19 < 16; fi19++ {
            v20 := GenS2{F0: ((!false) || false), F1: -91}
            var v21 float32 = 1e10
            _ = fi19
            _ = v20
            _ = v21
        }
        v3 += (^(v3))
        v22 := v5[-319]
        _ = v22
    } else {
        var v23 float32 = (((3.14159 * v18) + (v18 + 1e10)) / v18)
        _ = v23
    }
    var v24 uint = v3
    v25 := GenS3{F0: 178}
    v4 = -2.25
    var v26 int64 = 0
    fmt.Println(v3, v12)
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v13:", v13)
    fmt.Println("v18:", v18)
    fmt.Println("v2:", v2)
    fmt.Println("v24:", v24)
    fmt.Println("v26:", v26)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("maplen v5:", len(v5))
    _ = v1
    _ = v12
    _ = v13
    _ = v18
    _ = v2
    _ = v24
    _ = v25
    _ = v26
    _ = v3
    _ = v4
    _ = v5
}
