package main

import (
    "fmt"
)

type GenS1 struct { F0 int; F1 int16; F2 bool }
type GenS2 struct { F0 bool; F1 int; F2 int32 }
type GenS3 struct { F0 int64 }
func genFn0(p0 string) uint8 {
    if (2 >= 480) {
        fmt.Println(p0)
    } else {
        v4 := GenS1{F0: -372, F1: -923, F2: false}
        p0 = (p0 + (p0 + ((p0 + p0) + p0)))
        v5 := GenS2{F0: (42 != -548), F1: -2147483648, F2: 1}
        _ = v4
        _ = v5
    }
    for fi6 := 0; fi6 < 14; fi6++ {
        fmt.Println(p0, fi6)
        if true {
            if (11 >= 1) {
                var v7 float64 = 1.5
                _ = v7
            }
        }
        var v8 string = "a"
        _ = fi6
        _ = v8
    }
    _ = p0
    return 1
}
func genFn1(p0 uint64, p1 uint16, p2 uint64) int16 {
    var v9 int64 = -374
    fmt.Println(p2)
    v10 := GenS2{F0: (!(2147483648 < 575)), F1: 0, F2: 882}
    _ = p0
    _ = p1
    _ = p2
    _ = v10
    _ = v9
    return 1000
}

func main() {
    var v1 bool = ((false && false) && true)
    v2 := make(map[int]int)
    var v3 uint16 = 42
    switch 9223372036854775807 {
    case 32768:
        fmt.Println(v1)
    case 9223372036854775807:
        var v11 int16 = 1000
        _ = v11
    case -685:
        v12 := GenS1{F0: 42, F1: 42, F2: false}
        _ = v12
    }
    var v13 int16 = -270
    fmt.Println(v3)
    var v14 int64 = 0
    var v15 int32 = 869
    fmt.Println(v13, v14)
    v1 = v1
    if (v15 != -2147483648) {
        v16 := GenS2{F0: (18446744073709551615 <= 2), F1: 9223372036854775807, F2: -2147483648}
        _ = len(v2)
        v14 -= 0
        _ = v16
    } else {
        v17 := GenS3{F0: (v14 | v14)}
        switch ("\n\t" + "\n\t") {
        case "foo bar":
            v1 = ((v3 * v3) == v3)
        case "":
            v18 := []int{-585, 1, 217}
            _ = v18
        case "hello":
            v1 = (!(v14 > (v14 >> ((540) & 63))))
        default:
            if (2 >= 4294967295) {
                v1 = false
                var v19 uint8 = 127
                _ = v19
            }
        }
        _ = v17
    }
    fmt.Println("v1:", v1)
    fmt.Println("v13:", v13)
    fmt.Println("v14:", v14)
    fmt.Println("v15:", v15)
    fmt.Println("maplen v2:", len(v2))
    fmt.Println("v3:", v3)
    _ = v1
    _ = v13
    _ = v14
    _ = v15
    _ = v2
    _ = v3
}
