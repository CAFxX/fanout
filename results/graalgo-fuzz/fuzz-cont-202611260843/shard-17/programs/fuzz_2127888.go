package main

import (
    "fmt"
)

type GenS1 struct { F0 uint32; F1 uint8 }
type GenS2 struct { F0 uint }
type GenS3 struct { F0 int64 }
func genFn0(p0 int8) int64 {
    var v5 uint8 = 128
    v6 := GenS1{F0: 42, F1: 62}
    switch 127 {
    case 32768:
        v7 := GenS2{F0: 18446744073709551615}
        _ = v7
    case 128:
        var v8 int64 = -1
        _ = v8
    default:
        var v9 int16 = 710
        _ = v9
    }
    _ = p0
    _ = v5
    _ = v6
    return -1
}
func genFn1(p0 int8, p1 uint) int64 {
    v10 := []int{0}
    if (-1 != p0) {
        v11 := GenS2{F0: (^(p1))}
        _ = v11
    }
    _ = p0
    _ = p1
    _ = v10
    return 9223372036854775807
}

func main() {
    v1 := make(map[string]int)
    var v2 uint64 = 2
    var v3 string = "foo bar"
    v4 := make(map[string]uint32)
    var v12 uint64 = (v2 / ((v2) | 1))
    fmt.Println(v2)
    v13 := GenS3{F0: 0}
    v3 = v3
    var v14 float64 = ((1e10 / (1e10 - 3.14159)) * (-2.25 + 3.14159))
    switch -2147483648 {
    case -1:
        v12 = v2
    }
    switch 127 {
    case -174:
        fmt.Println(v3, v3)
    }
    var v15 int64 = 9223372036854775807
    var v16 uint16 = 0
    v12 = v2
    var v17 int8 = 33
    v17 += v17
    for fi18 := 0; fi18 < 6; fi18++ {
        var v19 int16 = 42
        var v20 float64 = v14
        _ = fi18
        _ = v19
        _ = v20
    }
    v21 := GenS3{F0: (v15 * v15)}
    v17 = 1
    for fi22 := 0; fi22 < 10; fi22++ {
        v23 := []int{(-(fi22))}
        var v24 uint64 = v2
        v25 := GenS2{F0: 2}
        _ = fi22
        _ = v23
        _ = v24
        _ = v25
    }
    fmt.Println("maplen v1:", len(v1))
    fmt.Println("v12:", v12)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    _ = v1
    _ = v12
    _ = v13
    _ = v14
    _ = v15
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v3
    _ = v4
}
