package main

import (
    "fmt"
)

type GenS1 struct { F0 uint64; F1 bool }
type GenS2 struct { F0 bool; F1 uint16 }
func genFn0(p0 bool, p1 int32) uint64 {
    v7 := []int{}
    if true {
        p1 -= p1
    }
    v8 := GenS1{F0: 4294967295, F1: (false || (1 == 2))}
    switch 0 {
    case 128:
        var v9 int16 = 613
        _ = v9
    }
    _ = p0
    _ = p1
    _ = v7
    _ = v8
    return 42
}
func genFn1() int {
    if ((18446744073709551615 <= 985) || false) {
        var v10 bool = ((!true) || (0 != 0))
        if true {
            v10 = (230 != 847)
            v10 = ((!(true || (4294967295 <= 18446744073709551615))) && v10)
            if true {
                var v11 int64 = 1
                _ = v11
            }
        }
        v12 := []int{-9223372036854775808, -898, 2147483647, -28}
        _ = v10
        _ = v12
    } else {
        var v13 int32 = 873
        var v14 string = "hello"
        _ = v13
        _ = v14
    }
    if (4294967295 != 65535) {
        switch 638 {
        case 32768:
            var v15 string = "foo bar"
            _ = v15
        }
    } else {
        var v16 uint32 = 42
        _ = v16
    }
    var v17 uint16 = 2
    _ = v17
    return 42
}

func main() {
    var v1 uint = 255
    var v2 float64 = (0.0 - 0.0)
    var v3 string = "\n\t"
    v4 := make(map[int]uint32)
    var v5 int64 = 1
    var v6 int8 = 0
    switch ((v3 + "") + (v3 + v3)) {
    case "a":
        v2 = v2
    case "\n\t":
        fmt.Println(v2)
    }
    var v18 float32 = 1.5
    var v19 float64 = (v2 + 1e10)
    v20 := []int{9223372036854775807, 127, 831}
    var v21 uint64 = 438
    fmt.Println(v5, v3)
    for fi22 := 0; fi22 < 9; fi22++ {
        v21 -= 1
        _ = fi22
    }
    switch -874 {
    case 32767:
        v21 = uint64(v1)
    }
    switch 882 {
    case 9223372036854775807:
        v23 := GenS2{F0: false, F1: 65535}
        _ = v23
    default:
        v6 = (v6 * -128)
    }
    switch 0 {
    case 127:
        var v24 int32 = 1
        _ = v24
    case -501:
        switch 42 {
        case 173:
            v4[-2147483648] = 2
        case 9223372036854775807:
            v25 := GenS2{F0: ((v19 >= (v19 + 0.0)) && (((v21 | v21) <= v21) || false)), F1: 255}
            _ = v25
        default:
            v19 -= v19
        }
    case -2147483648:
        v1 -= v1
    default:
        var v26 string = v3
        _ = v26
    }
    v5 = 0
    var v27 uint64 = v21
    fmt.Println(v3)
    v20 = append(v20, 0)
    fmt.Println("v1:", v1)
    fmt.Println("v18:", v18)
    fmt.Println("v19:", v19)
    fmt.Println("v2:", v2)
    fmt.Println("slicelen v20:", len(v20), v20)
    fmt.Println("v21:", v21)
    fmt.Println("v27:", v27)
    fmt.Println("v3:", v3)
    fmt.Println("maplen v4:", len(v4))
    fmt.Println("v5:", v5)
    fmt.Println("v6:", v6)
    _ = v1
    _ = v18
    _ = v19
    _ = v2
    _ = v20
    _ = v21
    _ = v27
    _ = v3
    _ = v4
    _ = v5
    _ = v6
}
