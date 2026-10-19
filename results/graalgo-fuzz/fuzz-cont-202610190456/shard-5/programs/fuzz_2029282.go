package main

import (
    "fmt"
)

type GenS1 struct { F0 bool; F1 string; F2 int16 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int32; F1 string }
type GenS4 struct { F0 int; F1 int8; F2 int }
type GenS5 struct { F0 int64; F1 int64 }
func genFn0(p0 int8) int64 {
    if false {
        v5 := GenS1{F0: (!((true && (-513 == 42)) || (true || (true || false)))), F1: "foo bar", F2: 32767}
        _ = v5
    }
    if (true && false) {
        var v6 uint32 = 0
        _ = v6
    }
    p0 -= p0
    v7 := GenS2{F0: 461}
    _ = p0
    _ = v7
    return 42
}

func main() {
    var v1 bool = false
    var v2 float64 = float64((((0.1 / 3.14159) / 3.14159) * 3.14159))
    v3 := make(map[int]string)
    v4 := []uint32{524}
    for fi8 := 0; fi8 < 2; fi8++ {
        v2 = v2
        var v9 int8 = -128
        v2 -= v2
        _ = fi8
        _ = v9
    }
    fmt.Println(v2)
    v10 := GenS3{F0: -2147483648, F1: "foo bar"}
    if v1 {
        var v11 uint16 = 2
        v12 := GenS4{F0: -265, F1: -1, F2: 878}
        v13 := GenS1{F0: (!((42 >= 2) || v1)), F1: ("a" + "foo bar"), F2: 32767}
        _ = v11
        _ = v12
        _ = v13
    }
    var v14 uint8 = 127
    v15 := GenS3{F0: 2147483647, F1: ""}
    v2 = -2.25
    if (638 <= 42) {
        v16 := GenS5{F0: -1, F1: 10}
        _ = v16
    }
    fmt.Println(v2)
    var v17 int32 = -485
    v14 = v14
    fmt.Println("v1:", v1)
    fmt.Println("v14:", v14)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("maplen v3:", len(v3))
    fmt.Println("slicelen v4:", len(v4), v4)
    _ = v1
    _ = v10
    _ = v14
    _ = v15
    _ = v17
    _ = v2
    _ = v3
    _ = v4
}
