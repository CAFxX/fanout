package main

import (
    "fmt"
)

type GenS1 struct { F0 int8; F1 int64 }
type GenS2 struct { F0 int32; F1 int8; F2 uint }
func genFn0(p0 int64, p1 int32) int {
    if (p1 <= p1) {
        v7 := GenS1{F0: 0, F1: (p0 + p0)}
        v8 := GenS1{F0: -26, F1: -390}
        _ = v7
        _ = v8
    }
    fmt.Println(p0, p1)
    _ = p0
    _ = p1
    return 42
}

func main() {
    var v1 uint = 94
    v2 := make(map[uint32]uint32)
    var v3 float64 = ((1.5 + 1.5) / ((1.5 / 0.1) / (float64(0.1) / float64(3.14159))))
    var v4 uint = (v1 * v1)
    var v5 bool = (!(0 != 42))
    v6 := []string{"foo bar"}
    fmt.Println(v5, v3, v1)
    v2[2147483648] = 529
    v5 = v5
    v9 := GenS1{F0: -1, F1: 1}
    v10 := GenS2{F0: -2147483648, F1: 1, F2: 552}
    var v11 int8 = -128
    delete(v2, 2)
    var v12 int16 = -411
    fmt.Println(v11, v12)
    v12 -= (v12 & v12)
    for fi13 := 0; fi13 < 3; fi13++ {
        var v14 float64 = (v3 * v3)
        _ = fi13
        _ = v14
    }
    if (true || (32768 != 978)) {
        fmt.Println(v12)
        var v15 string = "hello"
        switch ("a" + ((v15 + "hello") + v15)) {
        case "foo bar":
            v11 = (v11 | 42)
        case "a":
            if (42 < 911) {
                v16 := GenS2{F0: 0, F1: v11, F2: 18446744073709551615}
                fmt.Println(v15, v1)
                _ = v16
            } else {
                fmt.Println(v1, v3)
            }
        }
        _ = v15
    }
    if true {
        v17 := GenS1{F0: v11, F1: 1}
        v5 = v5
        _ = v17
    }
    for fi18 := 0; fi18 < 13; fi18++ {
        var v19 int = 0
        fmt.Println(v19, v3)
        _ = fi18
        _ = v19
    }
    var v20 int64 = 0
    fmt.Println("v1:", v1)
    fmt.Println("v11:", v11)
    fmt.Println("v12:", v12)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v20:", v20)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("slicelen v6:", len(v6), v6)
    _ = v1
    _ = v10
    _ = v11
    _ = v12
    _ = v2
    _ = v20
    _ = v3
    _ = v4
    _ = v5
    _ = v6
    _ = v9
}
