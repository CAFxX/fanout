package main

import (
    "fmt"
)

type GenS1 struct { F0 int }
type GenS2 struct { F0 bool; F1 int8; F2 uint32 }
type GenS3 struct { F0 bool }
func genFn0(p0 int64, p1 string, p2 string) uint8 {
    switch "a" {
    case "hello":
        var v7 float64 = 1.5
        _ = v7
    }
    p1 = ((p1 + p1) + p1)
    switch -541 {
    case -2147483648:
        v8 := []int{410}
        _ = v8
    case 32768:
        var v9 int64 = (p0 & 42)
        _ = v9
    case 446:
        if (1 != 1) {
            v10 := GenS1{F0: 127}
            v11 := []int{}
            _ = v10
            _ = v11
        } else {
            v12 := GenS2{F0: true, F1: 107, F2: 2147483648}
            var v13 uint64 = 787
            if true {
                p1 = (p1 + ("foo bar" + p2))
            } else {
                fmt.Println(p2, p0)
                p1 = "\n\t"
            }
            _ = v12
            _ = v13
        }
    }
    for fi14 := 0; fi14 < 9; fi14++ {
        v15 := []int{fi14, fi14, (fi14 * fi14)}
        _ = fi14
        _ = v15
    }
    p1 = ("foo bar" + p2)
    _ = p0
    _ = p1
    _ = p2
    return 178
}

func main() {
    var v1 int8 = 0
    var v2 uint = 1
    var v3 string = "foo bar"
    var v4 int64 = 9223372036854775807
    var v5 uint = (v2 % ((v2) | 1))
    var v6 int16 = -1
    var v16 float64 = ((1.5 + (float64((1.5 + (0.1 + (-2.25 + 0.1)))) - 3.14159)) - float64(3.14159))
    var v17 uint = v5
    fmt.Println(v3, v4)
    fmt.Println(v5, v1)
    switch 1 {
    case -2147483648:
        switch -9223372036854775808 {
        case 0:
            var v18 uint64 = 4294967295
            _ = v18
        case 9223372036854775807:
            v19 := []int{}
            _ = v19
        default:
            var v20 uint16 = 1
            _ = v20
        }
    }
    v21 := []int{1, -1}
    fmt.Println(v2)
    for fi22 := 0; fi22 < 14; fi22++ {
        var v23 string = "\n\t"
        v24 := GenS3{F0: false}
        _ = fi22
        _ = v23
        _ = v24
    }
    v25 := GenS1{F0: -706}
    if (128 <= -2147483648) {
        var v26 int64 = v4
        _ = v26
    }
    v1 -= v1
    _ = len(v21)
    fmt.Println(v4, v2, v1)
    fmt.Println("v1:", v1)
    fmt.Println("v16:", v16)
    fmt.Println("v17:", v17)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v21:", len(v21), v21)
    fmt.Println("v3:", v3)
    fmt.Println("v4:", v4)
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v16
    _ = v17
    _ = v2
    _ = v21
    _ = v25
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
